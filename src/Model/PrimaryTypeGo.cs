﻿// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Model;
using AutoRest.Core.Utilities;
using AutoRest.Go;
using System;
using System.Collections.Generic;
using System.Text;

namespace AutoRest.Go.Model
{
    public class PrimaryTypeGo : PrimaryType
    {
        /// <summary>
        /// Format name for ETags.
        /// </summary>
        public const string FormatETag = "etag";

        public PrimaryTypeGo() : base()
        {
            Name.OnGet += v =>
            {
                return ImplementationName;
            };
        }

        public PrimaryTypeGo(KnownPrimaryType primaryType) : base(primaryType)
        {
            Name.OnGet += v =>
            {
                return ImplementationName;
            };
        }

        /// <summary>
        /// Add imports for primary type.
        /// </summary>
        /// <param name="imports"></param>
        public void AddImports(HashSet<string> imports)
        {
            if (!string.IsNullOrWhiteSpace(Import))
                imports.Add(Import);
        }

        public virtual string Import
        {
            get
            {
                switch (KnownPrimaryType)
                {
                    case KnownPrimaryType.Date:
                    case KnownPrimaryType.DateTimeRfc1123:
                    case KnownPrimaryType.DateTime:
                        return GetImportLine(package: "time");
                    case KnownPrimaryType.Decimal:
                        return GetImportLine(package: "github.com/shopspring/decimal");
                    case KnownPrimaryType.Stream:
                        return GetImportLine(package: "io");
                    case KnownPrimaryType.UnixTime:
                        return GetImportLine(package: "github.com/Azure/go-autorest/autorest/date");
                    case KnownPrimaryType.Uuid:
                        return GetImportLine(package: "github.com/satori/go.uuid", alias: "uuid");
                    default:
                        return string.Empty;
                }
            }
        }

        public virtual string ImplementationName
        {
            get
            {
                switch (KnownPrimaryType)
                {
                    case KnownPrimaryType.Base64Url:
                        // TODO: add support
                        return "string";

                    case KnownPrimaryType.ByteArray:
                        return "[]byte";

                    case KnownPrimaryType.Boolean:
                        return "bool";

                    case KnownPrimaryType.Date:
                    case KnownPrimaryType.DateTime:
                    case KnownPrimaryType.DateTimeRfc1123:
                        return "time.Time";

                    case KnownPrimaryType.Double:
                        return "float64";

                    case KnownPrimaryType.Decimal:
                        return "decimal.Decimal";

                    case KnownPrimaryType.Int:
                        return "int32";

                    case KnownPrimaryType.Long:
                        return "int64";

                    case KnownPrimaryType.Stream:
                        return "io.ReadCloser";

                    case KnownPrimaryType.String:
                        if (string.Compare(Format, FormatETag, StringComparison.OrdinalIgnoreCase) == 0)
                        {
                            return CodeNamerGo.Instance.ETagTypeName;
                        }
                        return "string";

                    case KnownPrimaryType.TimeSpan:
                        return "string";

                    case KnownPrimaryType.Object:
                        // TODO: is this the correct way to support object types?
                        return "map[string]interface{}";

                    case KnownPrimaryType.UnixTime:
                        return "date.UnixTime";

                    case KnownPrimaryType.Uuid:
                        return "uuid.UUID";

                }
                throw new NotImplementedException($"Primary type {KnownPrimaryType} is not implemented in {GetType().Name}");
            }
        }

        public static string GetImportLine(string package, string alias = default(string))
        {
            var builder = new StringBuilder();
            if(!string.IsNullOrEmpty(alias))
            {
                builder.Append(alias);
                builder.Append(' ');
            }

            builder.Append('"');
            builder.Append(package);
            builder.Append('"');
            return builder.ToString();
        }
    }
}
