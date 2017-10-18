// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Utilities;
using AutoRest.Go.Model;

namespace AutoRest.Go.Decorators.v1
{
    public class EnumTypeGoDecorator : EnumTypeGo
    {
        /// <summary>
        /// Returns true if all the values for this enum are unique.
        /// </summary>
        public bool HasUniqueNames { get; }

        public EnumTypeGoDecorator(EnumTypeGo etg, bool hasUniqueNames)
        {
            this.LoadFrom(etg);
            HasUniqueNames = hasUniqueNames;
        }
    }
}
