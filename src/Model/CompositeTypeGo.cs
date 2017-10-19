// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Model;
using AutoRest.Core.Utilities;
using AutoRest.Extensions;
using System;
using System.Collections.Generic;
using System.Linq;
using static AutoRest.Core.Utilities.DependencyInjection;

namespace AutoRest.Go.Model
{
    /// <summary>
    /// Defines a synthesized composite type that wraps a primary type, array, or dictionary method response.
    /// </summary>
    public class CompositeTypeGo : CompositeType
    {
        private bool _wrapper;

        // True if the type is returned by a method
        public bool IsResponseType;

        // Name of the field containing the URL used to retrieve the next result set
        // (null or empty if the model is not paged).
        public string NextLink;

        public bool PreparerNeeded = false;

        public IEnumerable<CompositeType> DerivedTypes => CodeModel.ModelTypes.Where(t => t.DerivesFrom(this));

        public IEnumerable<CompositeType> SiblingTypes
        {
            get
            {
                var st = (BaseModelType as CompositeTypeGo).DerivedTypes;
                if (BaseModelType.BaseModelType != null && BaseModelType.BaseModelType.IsPolymorphic)
                {
                    st = st.Union((BaseModelType as CompositeTypeGo).SiblingTypes);
                }
                return st;
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

        public EnumType DiscriminatorEnum;

        public string DiscriminatorEnumValue => (DiscriminatorEnum as EnumTypeGo).Constants.FirstOrDefault(c => c.Value.Equals(SerializedName)).Key;

        public CompositeTypeGo()
        {

        }

        public CompositeTypeGo(string name) : base(name)
        {

        }

        public CompositeTypeGo(MethodGo responseToWrap)
        {
            var wrappedType = responseToWrap.ReturnType.Body;
            if (!wrappedType.ShouldBeSyntheticType())
            {
                throw new ArgumentException("{0} is not a valid type for SyntheticType", wrappedType.ToString());
            }

            // gosdk: Ensure the generated name does not collide with existing type names
            BaseType = wrappedType;

            if (wrappedType.XmlIsWrapped)
            {
                Name = wrappedType.XmlName;
                XmlProperties = wrappedType.XmlProperties;
            }
            else
            {
                Name = $"{responseToWrap.Name}Response";
            }

            // don't add the Value field for streams as it just duplicates
            // the response.Body field and doesn't provide any value.
            if (!wrappedType.IsPrimaryType(KnownPrimaryType.Stream))
            {
                // add the wrapped type as a property named Value
                var p = new PropertyGo();
                p.Name = "Value";
                p.SerializedName = "value";
                p.ModelType = wrappedType;
                Add(p);
            }

            _wrapper = true;
        }

        /// <summary>
        /// Returns true if XML serialization is enabled and the type name doesn't match the specified XML name.
        /// </summary>
        private bool NeedsXmlNameField => CodeModel.ShouldGenerateXmlSerialization && string.CompareOrdinal(Name, XmlName) != 0;

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
                    ModelType = DiscriminatorEnum,
                }));
            }            
        }

        public string PolymorphicProperty
        {
            get
            {
                if (!string.IsNullOrEmpty(PolymorphicDiscriminator))
                {
                    return CodeNamerGo.Instance.GetPropertyName(PolymorphicDiscriminator);
                }
                if (BaseModelType != null)
                {
                    return (BaseModelType as CompositeTypeGo).PolymorphicProperty;
                }
                return null;
            }
        }

        public IEnumerable<PropertyGo> AllProperties
        {
            get
            {
                if (BaseModelType != null)
                {
                    return Properties.Cast<PropertyGo>().Concat((BaseModelType as CompositeTypeGo).AllProperties);
                }
                return Properties.Cast<PropertyGo>();
            }
        }

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

            if (NeedsXmlNameField)
            {
                imports.Add(PrimaryTypeGo.GetImportLine(package: "encoding/xml"));
            }

            if (BaseIsPolymorphic && !IsPolymorphic)
            {
                imports.Add("\"encoding/json\"");
                imports.Add("\"errors\"");
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
            if (BaseIsPolymorphic && BaseModelType != null)
            {
                return (BaseModelType as CompositeTypeGo).IsPolymorphicResponse();
            }
            return IsPolymorphic && IsResponseType;
        }

        /// <summary>
        /// Returns the fields defined in this type.
        /// </summary>
        /// <param name="forMarshaller">Indicates if this type is an internal marshaller (usually false).</param>
        public string Fields(bool forMarshaller)
        {
            AddPolymorphicPropertyIfNecessary();
            var indented = new IndentedStringBuilder("    ");

            // check if the XML name matches tyhe type name.
            // if it doesn't then add an XMLName field.
            if (NeedsXmlNameField)
            {
                indented.AppendLine("// XMLName is used for marshalling and is subject to removal in a future release.");
                indented.AppendLine($"XMLName xml.Name `xml:\"{XmlName}\"`");
            }

            var properties = Properties.Cast<PropertyGo>().ToList();

            if (BaseModelType != null)
            {
                indented.Append(((CompositeTypeGo)BaseModelType).Fields(forMarshaller));
            }

            // Emit each property, except for named Enumerated types, as a pointer to the type
            foreach (var property in properties)
            {
                if (!forMarshaller && !string.IsNullOrEmpty(property.Documentation))
                {
                    indented.AppendFormat("// {0} - {1}\n", property.Name, property.Documentation);
                }
                var enumType = property.ModelType as EnumTypeGo;
                if (enumType != null && enumType.IsNamed)
                {
                    indented.AppendFormat("{0} {1} {2}\n",
                                    property.Name,
                                    enumType.Name,
                                    property.Tag());

                }
                else if (property.ModelType is DictionaryType)
                {
                    var typeName = (property.ModelType as DictionaryTypeGo).Name;
                    if (property.IsMetadata)
                    {
                        // use custom type instead of a map[string]string
                        typeName = "Metadata";
                    }
                    indented.AppendFormat("{0} {1} {2}\n", property.Name, typeName, property.Tag());
                }
                else if (property.ModelType.IsPrimaryType(KnownPrimaryType.Object))
                {
                    // TODO: I don't think this is the best way to handle object types
                    indented.AppendFormat("{0} *{1} {2}\n", property.Name, property.ModelType.Name, property.Tag());
                }
                else if (property.ShouldBeFlattened())
                {
                    // embed as an anonymous struct.  note that the ordering of this clause is
                    // important, i.e. we don't want to flatten primary types like dictionaries.
                    indented.AppendFormat("*{0} {1}\n", property.ModelType.Name, property.Tag());
                    property.Extensions[SwaggerExtensions.FlattenOriginalTypeName] = Name;
                }
                else if (!string.IsNullOrEmpty(NextLink) && property.Name.EqualsIgnoreCase(NextLink))
                {
                    // use custom type instead of *string
                    indented.Append($"{NextLink} Marker `xml:\"{NextLink}\"`");
                }
                else if (property.ModelType is CompositeType && (property.ModelType as CompositeTypeGo).IsPolymorphic)
                {
                    indented.AppendFormat("{0} {1} {2}\n", property.Name, property.ModelType.Name, property.Tag());
                }
                else
                {
                    // NextLinks might have differences in casing, but they need to be consistent
                    if (property.Name.EqualsIgnoreCase(NextLink))
                    {
                        property.Name = NextLink;
                    }

                    // to avoid any breaking changes in the v1 template always emit as a pointer type
                    var deref = property.ModelType.CanBeNull() || property.IsRequired ? string.Empty : "*";
                    var typeName = property.ModelType.Name.ToString();
                    if (forMarshaller && property.ModelType.IsDateTimeType())
                    {
                        typeName = "timeRFC3339";
                        if (property.ModelType.IsPrimaryType(KnownPrimaryType.DateTimeRfc1123))
                        {
                            typeName = "timeRFC1123";
                        }
                    }
                    indented.AppendFormat("{0} {3}{1} {2}\n", property.Name, typeName, property.Tag(), deref);
                }
            }

            return indented.ToString();
        }

        public bool IsWrapperType => _wrapper;

        public IModelType BaseType { get; private set; }

        public IModelType GetElementType(IModelType type)
        {
            if (type is SequenceTypeGo)
            {
                Name += "List";
                return GetElementType((type as SequenceType).ElementType);
            }
            else if (type is DictionaryTypeGo)
            {
                Name += "Set";
                return GetElementType(((type as DictionaryTypeGo).ValueType));
            }
            else
            {
                return type;
            }
        }

        public string PreparerMethodName => $"{Name}Preparer";

        public void SetName(string name)
        {
            Name = name;
        }

        /// <summary>
        /// Represents a response value that comes from an HTTP header.
        /// </summary>
        public class HeaderResponse
        {
            /// <summary>
            /// Gets the name of the response method.
            /// </summary>
            public string Name { get; }

            /// <summary>
            /// Gets type information of the response.
            /// </summary>
            public PropertyGo Type { get; }

            public HeaderResponse(PropertyGo pg)
            {
                Name = CodeNamerGo.Instance.GetMethodName(pg.GetClientName());
                Type = pg;
            }
        }

        /// <summary>
        /// Returns the list of values returned via response headers ordered by Name.  Can be empty.
        /// </summary>
        public IEnumerable<HeaderResponse> ResponseHeaders()
        {
            var respHeaders = new List<HeaderResponse>();
            if (IsResponseType)
            {
                // look up the response types
                var methods = CodeModel.Methods.Cast<MethodGo>().Where(m => m.HasReturnValue() && m.ReturnType.Body.Equals(this));
                foreach (var method in methods)
                {
                    if (method.ReturnType.Headers != null)
                    {
                        var headersType = method.ReturnType.Headers as CompositeTypeGo;
                        foreach (var property in headersType.Properties.Cast<PropertyGo>())
                        {
                            if (property.SerializedName.ToString().StartsWith("x-ms-meta", StringComparison.OrdinalIgnoreCase))
                            {
                                // skip the metadata header as it's handled separately
                                continue;
                            }
                            respHeaders.Add(new HeaderResponse(property));
                        }
                    }
                }
            }
            respHeaders.Sort((lhs, rhs) => { return string.Compare(lhs.Name, rhs.Name, StringComparison.OrdinalIgnoreCase); });
            return respHeaders;
        }

        /// <summary>
        /// Returns true if this response type contains a metadata response header.
        /// </summary>
        public bool ResponseIncludesMetadata
        {
            get
            {
                if (!IsResponseType)
                    return false;

                var methods = CodeModel.Methods.Cast<MethodGo>().Where(m => m.HasReturnValue() && m.ReturnType.Body.Equals(this));
                foreach (var method in methods)
                {
                    if (method.ReturnType.Headers != null)
                    {
                        var headersType = method.ReturnType.Headers as CompositeTypeGo;
                        foreach (var property in headersType.Properties.Cast<PropertyGo>())
                        {
                            if (property.SerializedName.ToString().StartsWith("x-ms-meta", StringComparison.OrdinalIgnoreCase))
                            {
                                return true;
                            }
                        }
                    }
                }
                return false;
            }
        }
    }
}
