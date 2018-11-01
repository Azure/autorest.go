// Copyright (c) Microsoft Open Technologies, Inc. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Utilities;
using AutoRest.Core.Model;
using AutoRest.Extensions;

namespace AutoRest.Go.Model
{
    public class PropertyGo : Property
    {
        private const string Base64UrlDoc = "a URL-encoded base64 string";

        public PropertyGo()
        {
        }

        /// <summary>
        /// Gets or sets the documentation.
        /// </summary>
        public new Fixable<string> Documentation
        {
            get
            {
                if (ModelType is PrimaryTypeGo ptg && ptg.KnownFormat == KnownFormat.base64url)
                {
                    // we don't properly support the base64url type so add some extra
                    // docs stating that the type should be a URL-encoded base64 string.
                    // NOTE: once proper support is added remove this code.
                    if (base.Documentation.IsNullOrEmpty())
                    {
                        return Base64UrlDoc;
                    }
                    else
                    {
                        return $"{base.Documentation} ({Base64UrlDoc})";
                    }
                }
                return base.Documentation;
            }
            set { base.Documentation = value; }
        }

        public string JsonTag(bool omitEmpty = true)
        {
            return string.Format("`json:\"{0}{1}\"`", SerializedName, omitEmpty ? ",omitempty" : "");
        }

        /// <summary>
        /// Gets if the property should be treated as a pointer
        /// </summary>
        public bool IsPointer => !(this.ModelType.HasInterface()
            || (this.ModelType is EnumTypeGo enumType && enumType.IsNamed)
            || this.ModelType is DictionaryTypeGo
            || this.ModelType.PrimaryType(KnownPrimaryType.Object));

        /// <summary>
        /// Gets the field name for this property.  Usually this is Name however for
        /// flattened composite types it will be the name of the composite type.
        /// </summary>
        public string FieldName
        {
            get
            {
                if (ModelType is CompositeTypeGo && this.ShouldBeFlattened())
                {
                    return TypeName;
                }
                return Name;
            }
        }

        /// <summary>
        /// Gets the type name for this property.
        /// </summary>
        private string TypeName
        {
            get
            {
                return ModelType.HasInterface() ? ModelType.GetInterfaceName() : ModelType.Name.Value;
            }
        }

        ///<summary>
        /// Gets the property representation.
        /// </summary>
        public string Field
        {
            get
            {
                // Polymorphic fields are implemented as go interfaces and a pointer to an
                // interface is not implementing the interface.
                var fieldType = IsPointer ? $"*{TypeName}" : $"{TypeName}";
                var jsonTag = ModelType is DictionaryTypeGo ? JsonTag(omitEmpty: false) : JsonTag();

                return ModelType is CompositeTypeGo && this.ShouldBeFlattened()
                    ? $"{fieldType} {jsonTag}"
                    : $"{Name} {fieldType} {jsonTag}";
            }
        }

        /// <summary>
        /// Determiens if this PropertyGo instance is equal to another.
        /// </summary>
        /// <param name="value"></param>
        /// <returns></returns>
        public override bool Equals(object value)
        {
            if (value is PropertyGo goProperty)
            {
                return goProperty.Name == Name;
            }

            return false;
        }

        /// <summary>
        /// Returns the hash code for this instance.
        /// </summary>
        /// <returns>A 32-bit signed integer hash code.</returns>
        public override int GetHashCode()
        {
            return Name.GetHashCode();
        }
    }
}
