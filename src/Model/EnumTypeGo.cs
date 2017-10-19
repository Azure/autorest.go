// Copyright (c) Microsoft Open Technologies, Inc. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Model;
using AutoRest.Core.Utilities;
using System.Linq;

namespace AutoRest.Go.Model
{
    public class EnumTypeGo : EnumType
    {
        public EnumTypeGo()
        {
            // the default value for unnamed enums is "enum"
            Name.OnGet += v => v == "enum" ? "string" : v;
        }

        public EnumTypeGo(EnumType source) : this()
        {
            this.LoadFrom(source);
        }

        /// <summary>
        /// Returns true if this enum has a name (swagger supports "anonymous" enums).
        /// </summary>
        public bool IsNamed => Name != "string" && Values.Any();

        /// <summary>
        /// Gets the doc string for this enum type.
        /// Since swagger doesn't let you define a description for enums we make one up.
        /// </summary>
        public string Documentation => $"{Name} enumerates the values for {Name.FixedValue.ToPhrase()}.";
    }
}
