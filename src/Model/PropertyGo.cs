// Copyright (c) Microsoft Open Technologies, Inc. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Utilities;
using AutoRest.Core.Model;
using AutoRest.Extensions;
using System;
using System.Text;

namespace AutoRest.Go.Model
{
    public class PropertyGo : Property
    {
        public PropertyGo()
        {
        }

        public string Tag(bool omitEmpty = true)
        {
            // don't emit a tag if this property is part of a parameter group (it's not necessary)
            if (Extensions.ContainsKey(SwaggerExtensions.ParameterGroupExtension))
            {
                return string.Empty;
            }

            if (!Parent.CodeModel.ShouldGenerateXmlSerialization)
            {
                return string.Format("`json:\"{0}{1}\"`", SerializedName, omitEmpty ? ",omitempty" : "");
            }

            var sb = new StringBuilder("`xml:\"");

            bool hasParent = false;
            if (Parent is CompositeTypeGo go && !go.IsWrapperType)
            {
                sb.Append(XmlName);
                hasParent = true;
            }

            if (XmlIsWrapped)
            {
                if (hasParent)
                {
                    sb.Append('>');
                }

                var asSequence = ModelType as SequenceTypeGo;
                sb.Append(asSequence.ElementXmlName);
            }
            else if (XmlIsAttribute)
            {
                sb.Append(",attr");
            }

            sb.Append("\"`");
            return sb.ToString();

        }

        /// <summary>
        /// Returns true if this property represents custom metadata.
        /// </summary>
        public bool IsMetadata => string.Compare(Name, "Metadata", StringComparison.OrdinalIgnoreCase) == 0 && ModelType is DictionaryTypeGo;

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
