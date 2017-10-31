// Copyright (c) Microsoft Open Technologies, Inc. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Model;

namespace AutoRest.Go.Model
{
    class EnumValueGo : EnumValue
    {
        /// <summary>
        /// Returns a properly formatted enum value name based on its containing parent.
        /// </summary>
        /// <param name="parent">The EnumTypeGo to which the value belongs.</param>
        /// <param name="value">The EnumValueGo for which a formatted value name will be returned.</param>
        /// <returns>The formatted enum value name.</returns>
        public static string FormatName(EnumTypeGo parent, EnumValueGo value)
        {
            // TODO: ideally the core would set the Parent field to that of the EnumTypeGo
            //       to which this belongs, for now this is how we work around that.
            var parentName = parent.Name.ToString();
            return $"{parentName.Substring(0, parentName.Length - 4)}{value.Name}";
        }
    }
}
