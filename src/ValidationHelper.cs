// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core;
using AutoRest.Core.Utilities;
using AutoRest.Core.Utilities.Collections;
using AutoRest.Core.Model;
using AutoRest.Go.Model;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;

namespace AutoRest.Go
{
    /// <summary>
    /// Provides helper methods for generating validation code based on the template version being used.
    /// For the v1 template all of the validation code comes from the validation package in go-autorest,
    /// so type names will start with an upper-case letter and be fully qualified.
    /// </summary>
    static class ValidationHelper
    {
        private static TemplateFactory.Version s_version = TemplateFactory.Instance.TemplateVersion;

        /// <summary>
        /// Returns true if the validation code is imported from go-autorest.
        /// </summary>
        private static bool UseValidationPackage()
        {
            return s_version == TemplateFactory.Version.v1;
        }

        /// <summary>
        /// Gets the proper type name for validation constraints.
        /// </summary>
        public static string ConstraintTypeName => UseValidationPackage() ? "validation.Constraint" : "constraint";

        /// <summary>
        /// Gets the proper type name fo the null constraint.
        /// </summary>
        public static string NullConstraint => UseValidationPackage() ? "validation.Null" : "null";

        /// <summary>
        /// Returns the proper type name for the specified constraint field.
        /// </summary>
        public static string GetConstraintFieldName(ConstraintFields field)
        {
            if (UseValidationPackage())
            {
                return field.ToString();
            }
            return field.ToString().ToCamelCase();
        }

        /// <summary>
        /// Returns the proper type name for the specified validation field.
        /// </summary>
        public static string GetValidationFieldName(ValidationFields field)
        {
            if (UseValidationPackage())
            {
                return field.ToString();
            }
            return field.ToString().ToCamelCase();
        }

        /// <summary>
        /// Returns a properly cased string depending on if the type is exported or not.
        /// </summary>
        public static string ConstraintCasing(this string s)
        {
            if (UseValidationPackage())
            {
                return s;
            }
            return s.ToCamelCase();
        }
    }

    /// <summary>
    /// Enumerates the fields in the Constraint type.
    /// </summary>
    enum ConstraintFields
    {
        Chain,
        Name,
        Rule,
        Target
    }

    /// <summary>
    /// Enumerates the fields in the Validation type.
    /// </summary>
    enum ValidationFields
    {
        TargetValue,
        Constraints
    }
}
