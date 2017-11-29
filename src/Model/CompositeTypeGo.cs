// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Model;
using AutoRest.Core.Utilities;
using AutoRest.Extensions;
using System;
using System.Collections.Generic;
using System.Linq;
using AutoRest.Core.Utilities.Collections;
using static AutoRest.Core.Utilities.DependencyInjection;

namespace AutoRest.Go.Model
{
    /// <summary>
    /// Defines a synthesized composite type that wraps a primary type, array, or dictionary method response.
    /// </summary>
    public class CompositeTypeGo : CompositeType
    {
        /// <summary>
        ///True if the type is returned by a method
        /// </summary>
        public bool IsResponseType;

        /// <summary>
        /// Name of the field containing the URL used to retrieve the next result set (null or empty if the model is not paged).
        /// </summary>
        public string NextLink;

        public bool PreparerNeeded = false;

        public EnumTypeGo DiscriminatorEnum;

        private CompositeTypeGo _rootType;

        public CompositeTypeGo()
        {
        }

        public CompositeTypeGo(string name) : base(name)
        {
        }

        public CompositeTypeGo(IModelType wrappedType)
        {
            if (!wrappedType.ShouldBeSyntheticType())
            {
                throw new ArgumentException("{0} is not a valid type for SyntheticType", wrappedType.ToString());
            }

            // gosdk: Ensure the generated name does not collide with existing type names
            BaseType = wrappedType;

            IModelType elementType = GetElementType(wrappedType);

            if (elementType is PrimaryType)
            {
                var type = (elementType as PrimaryType).KnownPrimaryType;
                switch (type)
                {
                    case KnownPrimaryType.Object:
                        Name += "SetObject";
                        break;

                    case KnownPrimaryType.Boolean:
                        Name += "Bool";
                        break;

                    case KnownPrimaryType.Double:
                        Name += "Float64";
                        break;

                    case KnownPrimaryType.Int:
                        Name += "Int32";
                        break;

                    case KnownPrimaryType.Long:
                        Name += "Int64";
                        break;

                    case KnownPrimaryType.Stream:
                        Name += "ReadCloser";
                        break;

                    default:
                        Name += type.ToString();
                        break;
                }
            }
            else if (elementType is EnumType)
            {
                Name += "String";
            }
            else
            {
                Name += elementType.Name;
            }

            // add the wrapped type as a property named Value
            var p = new PropertyGo
            {
                Name = "Value",
                SerializedName = "value",
                ModelType = wrappedType
            };
            base.Add(p);
            AddPolymorphicPropertyIfNecessary();

            IsWrapperType = true;
        }

        public IEnumerable<CompositeType> DerivedTypes => CodeModel.ModelTypes.Where(t => t.DerivesFrom(this));

        public string DiscriminatorEnumValue => DiscriminatorEnum.Values.FirstOrDefault(v => v.SerializedName.Equals(SerializedName)).Name;

        public string PreparerMethodName => $"{Name}Preparer";

        public bool IsWrapperType { get; }

        public IModelType BaseType { get; }

        public IEnumerable<CompositeType> SiblingTypes
        {
            get
            {

                var siblingTypes = RootType.DerivedTypes;

                if (RootType.IsPolymorphic)
                {
                    siblingTypes = siblingTypes.ConcatSingleItem(RootType);
                }

                return siblingTypes;
            }
        }

        public bool HasPolymorphicFields
        {
            get
            {
                return AllProperties.Any(p =>
                        // polymorphic composite
                        (p.ModelType is CompositeType && (p.ModelType as CompositeTypeGo).IsPolymorphic) ||
                        // polymorphic array
                        (p.ModelType is SequenceType && (p.ModelType as SequenceTypeGo).ElementType is CompositeType &&
                            ((p.ModelType as SequenceTypeGo).ElementType as CompositeType).IsPolymorphic));
            }
        }

        public string PolymorphicProperty => !string.IsNullOrEmpty(PolymorphicDiscriminator) ?
            CodeNamerGo.Instance.GetPropertyName(PolymorphicDiscriminator) :
            (BaseModelType as CompositeTypeGo)?.PolymorphicProperty;

        public IEnumerable<PropertyGo> AllProperties => BaseModelType != null ?
            Properties.Cast<PropertyGo>().Concat((BaseModelType as CompositeTypeGo).AllProperties) :
            Properties.Cast<PropertyGo>();

        /// <summary>
        /// Gets the root type of the inheritance chain.
        /// </summary>
        public CompositeTypeGo RootType
        {
            get
            {
                if (_rootType == null)
                {

                    CompositeType rootModelType = this;
                    while (rootModelType.BaseModelType != null && rootModelType.BaseIsPolymorphic)
                    {
                        rootModelType = rootModelType.BaseModelType;
                    }

                    _rootType = rootModelType as CompositeTypeGo;
                }

                return _rootType;
            }
        }

        /// <summary>
        /// Gets if the type is a root type in an inheritance chain.
        /// </summary>
        public bool IsRootType => IsPolymorphic && RootType == this;

        /// <summary>
        /// Gets if the type is a leaf type in an inheritance chain.
        /// </summary>
        public bool IsLeafType => BaseIsPolymorphic && DerivedTypes.IsNullOrEmpty();

        /// <summary>
        /// Gets if the type has an interface.
        /// </summary>
        public bool HasInterface => IsRootType || (BaseIsPolymorphic && !IsLeafType);

        public override Property Add(Property item)
        {
            var property = base.Add(item) as PropertyGo;
            AddPolymorphicPropertyIfNecessary();
            return property;
        }

        /// <summary>
        /// Add imports for composite types.
        /// </summary>
        /// <param name="imports"></param>
        public void AddImports(HashSet<string> imports)
        {
            Properties.ForEach(p => p.ModelType.AddImports(imports));
            if (IsPolymorphic)
            {
                imports.Add("\"encoding/json\"");
            }
        }

        public string AddHTTPResponse()
        {
            return (IsResponseType || IsWrapperType) ?
                "autorest.Response `json:\"-\"`\n" :
                null;
        }

        public bool IsPolymorphicResponse() {
            if (IsPolymorphic && IsResponseType)
            {
                return true;
            }
            if (BaseModelType != null && BaseIsPolymorphic)
            {
                return (BaseModelType as CompositeTypeGo).IsPolymorphicResponse();
            }
            return false;
        }

        /// <summary>
        /// Returns all the fields contained in this type in a formatted string.
        /// </summary>
        public virtual string Fields()
        {
            AddPolymorphicPropertyIfNecessary();
            var indented = new IndentedStringBuilder("    ");
            var properties = Properties.Cast<PropertyGo>().ToList();

            if (BaseModelType != null)
            {
                indented.Append(((CompositeTypeGo)BaseModelType).Fields());
            }

            // Emit each property, except for named Enumerated types, as a pointer to the type
            foreach (var property in properties)
            {
                if (property.ModelType is EnumTypeGo enumType && enumType.IsNamed)
                {
                    indented.AppendFormat("{0} {1} {2}\n",
                                    property.Name,
                                    enumType.Name,
                                    property.JsonTag());

                }
                else if (property.ModelType is DictionaryType)
                {
                    indented.AppendFormat("{0} *{1} {2}\n", property.Name, (property.ModelType as DictionaryTypeGo).Name, property.JsonTag());
                }
                else if (property.ModelType.PrimaryType(KnownPrimaryType.Object))
                {
                    // TODO: I don't think this is the best way to handle object types
                    indented.AppendFormat("{0} *{1} {2}\n", property.Name, property.ModelType.Name, property.JsonTag());
                }
                else if (property.ModelType is CompositeTypeGo && property.ShouldBeFlattened())
                {
                    // embed as an anonymous struct.  note that the ordering of this clause is
                    // important, i.e. we don't want to flatten primary types like dictionaries.
                    // Polymorphic fields are implemented as go interfaces and a pointer to an
                    // interface is not implementing the interface.
                    indented.AppendFormat((property.ModelType as CompositeTypeGo).IsPolymorphic ?
                        "{0} {1}\n" :
                        "*{0} {1}\n",
                            property.ModelType.Name, property.JsonTag());
                }
                else if (property.ModelType is CompositeTypeGo && ((CompositeTypeGo) property.ModelType).IsPolymorphic)
                {
                    indented.AppendFormat("{0} {1} {2}\n", property.Name, property.ModelType.GetInterfaceName(), property.JsonTag());
                }
                else
                {
                    // NextLinks might have differences in casing, but they need to be consistent
                    if (property.Name.EqualsIgnoreCase(NextLink))
                    {
                        property.Name = NextLink;
                    }
                    indented.AppendFormat("{0} *{1} {2}\n", property.Name, property.ModelType.Name, property.JsonTag());
                }
            }

            return indented.ToString();
        }

        public IModelType GetElementType(IModelType type)
        {
            if (type is SequenceTypeGo sequenceType)
            {
                Name += "List";
                return GetElementType(sequenceType.ElementType);
            }
            else if (type is DictionaryTypeGo dictionaryType)
            {
                Name += "Set";
                return GetElementType(dictionaryType.ValueType);
            }
            else
            {
                return type;
            }
        }

        public void SetName(string name)
        {
            Name = name;
        }

        /// <summary>
        /// If PolymorphicDiscriminator is set, makes sure we have a PolymorphicDiscriminator property.
        /// </summary>
        private void AddPolymorphicPropertyIfNecessary()
        {
            if (!string.IsNullOrEmpty(PolymorphicDiscriminator) && Properties.All(p => p.SerializedName != PolymorphicDiscriminator))
            {
                base.Add(New<Property>(new
                {
                    Name = CodeNamerGo.Instance.GetPropertyName(PolymorphicDiscriminator),
                    SerializedName = PolymorphicDiscriminator,
                    ModelType = DiscriminatorEnum
                }));
            }
        }
    }
}
