// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Model;
using System;

namespace AutoRest.Go.Model
{
    /// <summary>
    /// Represents a future, which is the return type for long-running operations.
    /// </summary>
    class FutureTypeGo : CompositeTypeGo
    {
        /// <summary>
        /// Creates a new future type for the specified method.
        /// </summary>
        /// <param name="method">The method that will return a future.</param>
        public FutureTypeGo(MethodGo method) : base(CodeNamerGo.Instance.GetFutureTypeName(method))
        {
            if (!method.IsLongRunningOperation())
            {
                throw new InvalidOperationException($"method {method.Owner}.{method.Name} is not a long-running operation");
            }

            CodeModel = method.CodeModel;
            Documentation = "An abstraction for monitoring and retrieving the results of a long-running operation.";
            ClientTypeName = method.Owner;
            ResultTypeName = method.MethodReturnType;
            ResponderMethodName = method.ResponderMethodName;
        }

        public override string Fields()
        {
            return "    azure.Future";
        }

        /// <summary>
        /// Gets the client type name associated with this future.
        /// </summary>
        public string ClientTypeName { get; }

        /// <summary>
        /// Gets the type name for the object that's returned when the operation is complete.
        /// </summary>
        public string ResultTypeName { get; }

        /// <summary>
        /// Gets the name of the responder method associated with this future.
        /// </summary>
        public string ResponderMethodName { get; }

        public override bool Equals(object other)
        {
            if (other == null)
            {
                return false;
            }

            var asMyType = other as FutureTypeGo;
            if (asMyType == null)
            {
                return false;
            }

            if (ReferenceEquals(this, other))
            {
                return true;
            }
            return string.Compare(Name, asMyType.Name, StringComparison.Ordinal) == 0;
        }

        public override int GetHashCode()
        {
            return Name.GetHashCode();
        }
    }
}
