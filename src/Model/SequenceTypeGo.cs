// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Model;
using System.Collections.Generic;

namespace AutoRest.Go.Model
{
    public class SequenceTypeGo : SequenceType
    {
        public SequenceTypeGo() : base()
        {
            Name.OnGet += v => $"[]{ElementType.Name}";
        }

        /// <summary>
        /// Add imports for sequence type.
        /// </summary>
        /// <param name="imports"></param>
        public void AddImports(HashSet<string> imports)
        {
            ElementType.AddImports(imports);
        }

        /// <summary>
        /// Returns the type name of the element (shorthand for ElementType.Name).
        /// </summary>
        public string GetElement => $"{ElementType.Name}";
    }
}
