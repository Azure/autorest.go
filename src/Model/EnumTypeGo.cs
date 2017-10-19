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
        public EnumTypeGo()
        {
            // the default value for unnamed enums is "enum"
            Name.OnGet += v => v == "enum" ? "string" : FormatName(v);
        }

        public EnumTypeGo(EnumType source) : this()
        {
            this.LoadFrom(source);
        }

        /// <summary>
        /// Returns an empty paramater check operation.
        /// </summary>
        /// <param name="valueReference">The parameter to be checked if it's empty.</param>
        /// <param name="required">Indicates if this parameter is optional; optional enum parameters are passed by reference.</param>
        /// <param name="asEmpty">Indicates if we want to check that valueReference is empty or not empty.</param>
        /// <returns></returns>
        public string GetEmptyCheck(string valueReference, bool required, bool asEmpty)
        {
            var comp = asEmpty ? "==" : "!=";

            if (required)
            {
                return $"len({valueReference}) {comp} 0";
            }
            
            var logiclOp = asEmpty ? "||" : "&&";

            return string.Format("{0} {1} nil {2} len(*{0}) {1} 0", valueReference, comp, logiclOp);
        }

        public bool IsNamed => Name != "string" && Values.Any();

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
            // append "Type" to the type name
            if (!rawName.EndsWith("Type"))
            {
                return $"{rawName}Type";
            }
            return rawName;
        }

        private string FormatValue(string rawValue)
        {
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
