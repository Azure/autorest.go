// Copyright (c) Microsoft Open Technologies, Inc. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Utilities;
using AutoRest.Core.Model;
using AutoRest.Extensions;

namespace AutoRest.Go.Model
{
    public class PropertyGo : Property
    {
        public PropertyGo()
        {

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

        ///<summary>
        /// Gets the property representation.
        /// </summary>
        public string Field
        {
            get
            {
                // Polymorphic fields are implemented as go interfaces and a pointer to an
                // interface is not implementing the interface.
                var typeName = this.ModelType.HasInterface() ? this.ModelType.GetInterfaceName() : this.ModelType.Name.Value;
                var fieldType = this.IsPointer ? $"*{typeName}" : $"{typeName}";
                var jsonTag = this.ModelType is DictionaryTypeGo ? this.JsonTag(omitEmpty: false) : this.JsonTag();

                return this.ModelType is CompositeTypeGo && this.ShouldBeFlattened()
                    ? $"{fieldType} {jsonTag}"
                    : $"{this.Name} {fieldType} {jsonTag}";
            }
        }
    }
}
