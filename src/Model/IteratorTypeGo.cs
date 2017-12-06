// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Model;
using AutoRest.Core.Utilities;
using AutoRest.Extensions.Azure;
using AutoRest.Go;
using System;
using System.Linq;

namespace AutoRest.Go.Model
{
    /// <summary>
    /// An abstraction for seamlessly iterating over a pageable collection.
    /// </summary>
    internal class IteratorTypeGo : CompositeTypeGo
    {
        /// <summary>
        /// Creates a new iterator type for the specified pageable type.
        /// </summary>
        /// <param name="method">The pageable type that requires an iterator.</param>
        public IteratorTypeGo(PageTypeGo pageType) : base(CodeNamerGo.Instance.GetIteratorTypeName(pageType))
        {
            PageType = pageType;
            Documentation = $"Provides access to a complete listing of {PageType.ElementType.Name} values.";
        }

        /// <summary>
        /// Gets the PageTypeGo type associated with this iterator.
        /// </summary>
        public PageTypeGo PageType { get; }

        /// <summary>
        /// Gets the name of the indexer field used to track the current value.
        /// </summary>
        public string IndexField => "i";

        /// <summary>
        /// Gets the name of the page field that contains the current page of values.
        /// </summary>
        public string PageField => "page";

        public override string Fields()
        {
            return $"    {IndexField} int\n    {PageField} {PageType.Name}";
        }
    }
}
