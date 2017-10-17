// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using System;
using System.Collections.Generic;
using System.Globalization;
using AutoRest.Core.Model;


namespace AutoRest.Go.Model
{
    /// <summary>
    /// Defines a synthetic type used to hold an array or dictionary method response.
    /// </summary>
    public class DictionaryTypeGo : DictionaryType
    {
        bool _reqValues;

        // if value type can be implicitly null
        // then don't emit it as a pointer type.
        // only do this for newer templates.
        private string FieldNameFormat => (TemplateFactory.Instance.TemplateVersion != TemplateFactory.Version.v1) &&
            (ValueType.CanBeNull() || _reqValues)
                                ? "map[string]{0}"
                                : "map[string]*{0}";

        public DictionaryTypeGo()
        {
            Name.OnGet += value => string.Format(CultureInfo.InvariantCulture, FieldNameFormat, ValueType.Name);
        }

        public DictionaryTypeGo(bool reqValues) : this()
        {
            // for the v1 template always leave _reqValues as false
            if (TemplateFactory.Instance.TemplateVersion != TemplateFactory.Version.v1)
            {
                _reqValues = reqValues;
            }
        }

        /// <summary>
        /// Add imports for dictionary type.
        /// </summary>
        /// <param name="imports"></param>
        public void AddImports(HashSet<string> imports)
        {
            ValueType.AddImports(imports);
        }

        public string GetEmptyCheck(string valueReference, bool asEmpty)
        {
            return string.Format(asEmpty
                                    ? "{0} == nil || len({0}) == 0"
                                    : "{0} != nil && len({0}) > 0", valueReference);
        }

        /// <summary>
        /// Determines whether the specified object is equal to this object based on the ValueType.
        /// </summary>
        /// <param name="obj">The object to compare with this object.</param>
        /// <returns>true if the specified object is equal to this object; otherwise, false.</returns>
        public override bool Equals(object obj)
        {
            var mapType = obj as DictionaryTypeGo;

            if (mapType != null)
            {
                return mapType.ValueType == ValueType;
            }

            return false;
        }

        /// <summary>
        /// Returns the hash code for this instance.
        /// </summary>
        /// <returns>A 32-bit signed integer hash code.</returns>
        public override int GetHashCode()
        {
            return ValueType.GetHashCode();
        }
    }
}
