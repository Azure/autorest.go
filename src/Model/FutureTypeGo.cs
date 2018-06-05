// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Model;
using System;

namespace AutoRest.Go.Model
{
    /// <summary>
    /// Represents a future, which is the return type for long-running operations.
    /// </summary>
    internal class FutureTypeGo : CompositeTypeGo
    {
        /// <summary>
        /// Creates a new future type for the specified method.
        /// </summary>
        /// <param name="method">The method that will return a future.</param>
        public FutureTypeGo(MethodGo method, IModelType resultType) : this(CodeNamerGo.Instance.GetFutureTypeName(method), method, resultType) {}

        /// <summary>
        /// Creates a new future type for the specified method using the specified name.
        /// </summary>
        /// <param name="method">The method that will return a future.</param>
        public FutureTypeGo(string methodName, MethodGo method, IModelType resultType) : base(methodName)
        {
            if (!method.IsLongRunningOperation())
            {
                throw new InvalidOperationException($"method {method.Owner}.{method.Name} is not a long-running operation");
            }

            CodeModel = method.CodeModel;
            Documentation = "An abstraction for monitoring and retrieving the results of a long-running operation.";
            ClientTypeName = method.Owner;
            ResultType = resultType;
            ResponderMethodName = method.ResponderMethodName;
            if (method.Deprecated)
            {
                DeprecationMessage = "The method for this type has been deprecated.";
            }
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
        /// Gets the type of the object that's returned when the operation is complete.
        /// If the future doesn't have a response body this will be null.
        /// </summary>
        public IModelType ResultType { get; }

        /// <summary>
        /// Gets the type name of the object that's returned when the operation is complete.
        /// </summary>
        public string ResultTypeName
        {
            get
            {
                var resTypeName = MethodGo.DefaultReturnType;
                if (ResultType != null)
                {
                    resTypeName = ResultType.Name.ToString();
                }
                return resTypeName;
            }
        }

        /// <summary>
        /// Returns true if the result type is the default response type.
        /// </summary>
        public bool IsDefaultReturnType => ResultType == null;

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
