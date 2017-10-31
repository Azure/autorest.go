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

        public string Tag()
        {
            // don't emit a tag if this property is part of a parameter group (it's not necessary)
            if (Extensions.ContainsKey(SwaggerExtensions.ParameterGroupExtension))
            {
                return string.Empty;
            }

            if (Parent.CodeModel.ShouldGenerateXmlSerialization)
            {
                var sb = new StringBuilder("`xml:\"");

                bool hasParent = false;
                if (Parent is CompositeTypeGo && !((CompositeTypeGo)Parent).IsWrapperType)
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

            return string.Format("`json:\"{0},omitempty\"`", SerializedName);
        }

        /// <summary>
        /// Returns true if this property represents custom metadata.
        /// </summary>
        public bool IsMetadata
        {
            get
            {
                // unfortunately we have to use a heuristic
                return string.Compare(Name, "Metadata", StringComparison.OrdinalIgnoreCase) == 0 && ModelType is DictionaryTypeGo;
            }
        }
    }
}
