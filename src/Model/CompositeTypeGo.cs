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
        /// True if the type is returned by a method
        /// </summary>
        public bool IsResponseType;

        public EnumTypeGo DiscriminatorEnum;

        private CompositeTypeGo _rootType;

        public CompositeTypeGo()
        {
        }

        public CompositeTypeGo(string name) : base(name)
        {
            if (string.IsNullOrWhiteSpace(Documentation))
            {
                Documentation = "...";
            }
        }

        public CompositeTypeGo(IModelType wrappedType)
        {
            if (!wrappedType.ShouldBeSyntheticType())
            {
                throw new ArgumentException("{0} is not a valid type for SyntheticType", wrappedType.ToString());
            }

            if (string.IsNullOrWhiteSpace(Documentation))
            {
                Documentation = "...";
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

        public string DiscriminatorEnumValue => DiscriminatorEnum?.Values.FirstOrDefault(v => v.SerializedName.Equals(SerializedName))?.Name;

        public PropertyGo AdditionalPropertiesField => AllProperties.FirstOrDefault(p => p.ModelType is DictionaryTypeGo dictionaryType && dictionaryType.SupportsAdditionalProperties);

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

        /// <summary>
        /// Gets if there are any polymorphic fields.
        /// </summary>
        public bool HasPolymorphicFields
        {
            get
            {
                return AllProperties.Any(p =>
                        // polymorphic composite
                        p.ModelType.HasInterface() ||
                        // polymorphic array
                        (p.ModelType is SequenceType sequenceType &&
                         sequenceType.ElementType.HasInterface()));
            }
        }

        /// <summary>
        /// Gets if there are any flattened fields.
        /// </summary>
        public bool HasFlattenedFields => AllProperties.Any(p => p.ModelType is CompositeTypeGo && p.ShouldBeFlattened());

        public string PolymorphicProperty => !string.IsNullOrEmpty(PolymorphicDiscriminator)
            ? CodeNamerGo.Instance.GetPropertyName(PolymorphicDiscriminator)
            : ((CompositeTypeGo)BaseModelType).PolymorphicProperty;

        public IEnumerable<PropertyGo> AllProperties => BaseModelType != null ?
            Properties.Cast<PropertyGo>().Concat(((CompositeTypeGo)BaseModelType).AllProperties) :
            Properties.Cast<PropertyGo>();

        /// <summary>
        /// Returns true if this type requires custom unmarshalling methods to be generated.
        /// </summary>
        public bool NeedsCustomUnmarshalling =>
            HasPolymorphicFields || HasFlattenedFields || AdditionalPropertiesField != null;

        /// <summary>
        /// Returns true if this type requires custom marshalling methods to be generated.
        /// </summary>
        public bool NeedsCustomMarshalling => 
            BaseIsPolymorphic || IsPolymorphic || HasFlattenedFields || AllProperties.Any(p => p.ModelType is DictionaryTypeGo) ||
            AllProperties.Any(p => p.IsReadOnly);

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
                    while (rootModelType.BaseModelType?.BaseIsPolymorphic == true)
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

        public override Property Add(Property item)
        {
            var property = base.Add(item);
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
            if (NeedsCustomMarshalling)
            {
                imports.Add(PrimaryTypeGo.GetImportLine("encoding/json"));
            }
        }

        public string AddHTTPResponse()
        {
            return (IsResponseType || IsWrapperType) ?
                "autorest.Response `json:\"-\"`\n" :
                null;
        }

        public bool IsPolymorphicResponse()
        {
            if (IsPolymorphic && IsResponseType)
            {
                return true;
            }
            if (BaseModelType != null && BaseIsPolymorphic)
            {
                return ((CompositeTypeGo)BaseModelType).IsPolymorphicResponse();
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
            var properties = AllProperties.ToHashSet();

            if (!IsPolymorphic && RootType.IsPolymorphic)
            {
                RootType.AddPolymorphicPropertyIfNecessary();
                properties.Add((PropertyGo)RootType.PolymorphicDiscriminatorProperty);
            }

            // Emit each property, except for named Enumerated types, as a pointer to the type
            foreach (var property in properties)
            {
                if (property.Deprecated)
                {
                    var message = "This property has been deprecated.";
                    if (!string.IsNullOrWhiteSpace(property.DeprecationMessage))
                    {
                        message = property.DeprecationMessage;
                    }
                    indented.Append($"// Deprecated: {message}\n");
                }
                if (!string.IsNullOrWhiteSpace(property.Documentation))
                {
                    var ro = "";
                    if (property.IsReadOnly)
                    {
                        ro = "READ-ONLY; ";
                    }
                    indented.Append($"{property.FieldName} - {ro}{property.Documentation}".ToCommentBlock());
                }
                else if (property.IsReadOnly)
                {
                    indented.Append($"// {property.FieldName} - READ-ONLY\n");
                }

                indented.AppendLine(property.Field);
            }

            return indented.ToString();
        }

        private IModelType GetElementType(IModelType type)
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
        /// Gets the expression for a zero-initialized composite type.
        /// </summary>
        public string ZeroInitExpression => $"{Name}{{}}";

        /// <summary>
        /// If PolymorphicDiscriminator is set, makes sure we have a PolymorphicDiscriminator property.
        /// </summary>
        internal void AddPolymorphicPropertyIfNecessary()
        {
            if (!string.IsNullOrEmpty(PolymorphicDiscriminator) && Properties.All(p => p.SerializedName != PolymorphicDiscriminator))
            {
                base.Add(New<Property>(new
                {
                    Name = PolymorphicDiscriminator,
                    SerializedName = PolymorphicDiscriminator,
                    ModelType = DiscriminatorEnum
                }));
            }
        }
    }
}
