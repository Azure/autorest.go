// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Model;

namespace AutoRest.Go.Model
{
    /// <summary>
    /// Represents a response from a long-running operation that's also a paged operation.
    /// This is a bit of a hack as the "list all" method isn't part of the code model and doing
    /// it this way was the path of least resistance (we should fix this eventually).
    /// </summary>
    internal class LroPagedResponseGo : Response
    {
        public LroPagedResponseGo(FutureTypeGo returnType, FutureTypeGo listAllReturnType, IModelType headers) : base(returnType, headers)
        {
            ListAllReturnType = listAllReturnType;
        }

        /// <summary>
        /// Gets the future type to be returned from the "list all" method.
        /// </summary>
        public FutureTypeGo ListAllReturnType { get; }
    }
}
