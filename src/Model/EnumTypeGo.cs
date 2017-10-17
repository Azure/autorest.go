// Copyright (c) Microsoft Open Technologies, Inc. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using System.Collections.Generic;
using System.Linq;
using AutoRest.Core.Model;
using AutoRest.Core.Utilities;

namespace AutoRest.Go.Model
{
    public class EnumTypeGo : EnumType
    {
        public bool HasUniqueNames { get; set; }

        public EnumTypeGo()
        {
            // the default value for unnamed enums is "enum"
            Name.OnGet += v => v == "enum" ? "string" : FormatName(v);

            // Assume members have unique names
            HasUniqueNames = true;
        }

        public EnumTypeGo(EnumType source) : this()
        {
            this.LoadFrom(source);
        }

        public string GetEmptyCheck(string valueReference, bool required, bool asEmpty)
        {
            var comp = asEmpty ? "==" : "!=";
            var logiclOp = asEmpty ? "||" : "&&";
            var deref = required ? string.Empty : "*";

            return string.Format("{0} {1} nil {2} len({3}{0}) {1} 0", valueReference, comp, logiclOp, deref);
        }

        public bool IsNamed => Name != "string" && Values.Any();

        /// <summary>
        /// Returns true if enums use a "none" value instead of nil.
        /// </summary>
        public bool UseNone
        {
            get
            {
                return TemplateFactory.Instance.TemplateVersion != TemplateFactory.Version.v1;
            }
        }

        public IDictionary<string, string> Constants
        {
            get
            {
                var constants = new Dictionary<string, string>();
                Values
                    .ForEach(v =>
                    {
                        constants.Add(v.Name, v.SerializedName);
                    });

                return constants;
            }
        }

        public string Documentation { get; set; }

        private string FormatName(string rawName)
        {
            if (TemplateFactory.Instance.TemplateVersion == TemplateFactory.Version.v1)
            {
                return rawName;
            }
            // append "Type" to the type name
            if (!rawName.EndsWith("Type"))
            {
                return $"{rawName}Type";
            }
            return rawName;
        }

        private string FormatValue(string rawValue)
        {
            if (TemplateFactory.Instance.TemplateVersion == TemplateFactory.Version.v1)
            {
                return rawValue;
            }
            // remove "Type" from the end of the name
            // then append the value name to this string
            var nameAsString = Name.ToString();
            return $"{nameAsString.Substring(0, nameAsString.Length - 4)}{rawValue}";
        }

        /// <summary>
        /// Gets the type name phrase for this enum type.
        /// </summary>
        public string Phrase
        {
            get
            {
                return Name.FixedValue.ToPhrase();
            }
        }
    }
}
