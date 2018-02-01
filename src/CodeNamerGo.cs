﻿// Copyright (c) Microsoft Corporation. All rights reserved.
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
using System.Text;

namespace AutoRest.Go
{
    public class CodeNamerGo : CodeNamer
    {
        public new static CodeNamerGo Instance
        {
            get
            {
                return (CodeNamerGo)CodeNamer.Instance;
            }
        }

        public virtual IEnumerable<string> PipelineImports
        {
            get
            {
                return new string[] { PrimaryTypeGo.GetImportLine(package: PipelineImportPath) };
            }
        }

        public virtual IEnumerable<string> StandardImports
        {
            get
            {
                var imports = new List<string>();
                imports.Add(PrimaryTypeGo.GetImportLine(package: "net/http"));
                imports.Add(PrimaryTypeGo.GetImportLine(package: "context"));
                return imports;
            }
        }

        public virtual IEnumerable<string> PageableImports => new string[] 
        {
            PrimaryTypeGo.GetImportLine(package: "net/http")
        };

        // CommonInitialisms are those "words" within a name that Golint expects to be uppercase.
        // See https://github.com/golang/lint/blob/master/lint.go for detail.
        private HashSet<string> CommonInitialisms => new HashSet<string>(StringComparer.OrdinalIgnoreCase) {
                                                            "Acl",
                                                            "Api",
                                                            "Ascii",
                                                            "Cpu",
                                                            "Css",
                                                            "Dns",
                                                            "Eof",
                                                            "Guid",
                                                            "Html",
                                                            "Http",
                                                            "Https",
                                                            "Id",
                                                            "Ip",
                                                            "Json",
                                                            "Lhs",
                                                            "Qps",
                                                            "Ram",
                                                            "Rhs",
                                                            "Rpc",
                                                            "Sla",
                                                            "Smtp",
                                                            "Sql",
                                                            "Ssh",
                                                            "Tcp",
                                                            "Tls",
                                                            "Ttl",
                                                            "Udp",
                                                            "Ui",
                                                            "Uid",
                                                            "Uuid",
                                                            "Uri",
                                                            "Url",
                                                            "Utf8",
                                                            "Vm",
                                                            "Xml",
                                                            "Xsrf",
                                                            "Xss",
                                                        };

        public string[] UserDefinedNames => new string[] {
                                                            "UserAgent",
                                                            "Version",
                                                            "APIVersion",
                                                            "DefaultBaseURI",
                                                            "ManagementClient",
                                                            "NewWithBaseURI",
                                                            "New",
                                                            "NewWithoutDefaults",
                                                        };

        public IReadOnlyDictionary<HttpStatusCode, string> StatusCodeToGoString;



        /// <summary>
        /// Initializes a new instance of CodeNamerGo.
        /// </summary>
        public CodeNamerGo()
        {
            // Create a map from HttpStatusCode to the appropriate Go http.StatusXxxxx string.
            // -- Go does not have constants for the full HttpStatusCode enumeration; this set taken from http://golang.org/pkg/net/http/
            const HttpStatusCode tooManyRequests = (HttpStatusCode)429;
            var statusCodeMap = new Dictionary<HttpStatusCode, string>();
            foreach (var sc in new HttpStatusCode[]{
                HttpStatusCode.Continue,
                HttpStatusCode.SwitchingProtocols,

                HttpStatusCode.OK,
                HttpStatusCode.Created,
                HttpStatusCode.Accepted,
                HttpStatusCode.NonAuthoritativeInformation,
                HttpStatusCode.NoContent,
                HttpStatusCode.ResetContent,
                HttpStatusCode.PartialContent,

                HttpStatusCode.MultipleChoices,
                HttpStatusCode.MovedPermanently,
                HttpStatusCode.Found,
                HttpStatusCode.SeeOther,
                HttpStatusCode.NotModified,
                HttpStatusCode.UseProxy,
                HttpStatusCode.TemporaryRedirect,

                HttpStatusCode.BadRequest,
                HttpStatusCode.Unauthorized,
                HttpStatusCode.PaymentRequired,
                HttpStatusCode.Forbidden,
                HttpStatusCode.NotFound,
                HttpStatusCode.MethodNotAllowed,
                HttpStatusCode.NotAcceptable,
                HttpStatusCode.ProxyAuthenticationRequired,
                HttpStatusCode.RequestTimeout,
                HttpStatusCode.Conflict,
                HttpStatusCode.Gone,
                HttpStatusCode.LengthRequired,
                HttpStatusCode.PreconditionFailed,
                HttpStatusCode.RequestEntityTooLarge,
                HttpStatusCode.RequestUriTooLong,
                HttpStatusCode.UnsupportedMediaType,
                HttpStatusCode.RequestedRangeNotSatisfiable,
                HttpStatusCode.ExpectationFailed,
                tooManyRequests,

                HttpStatusCode.InternalServerError,
                HttpStatusCode.NotImplemented,
                HttpStatusCode.BadGateway,
                HttpStatusCode.ServiceUnavailable,
                HttpStatusCode.GatewayTimeout,
                HttpStatusCode.HttpVersionNotSupported
            })
            {
                statusCodeMap.Add(sc, string.Format("http.Status{0}", sc));
            }

            // Go names some constants slightly differently than the HttpStatusCode enumeration -- correct those
            statusCodeMap[HttpStatusCode.Redirect] = "http.StatusFound";
            statusCodeMap[HttpStatusCode.NonAuthoritativeInformation] = "http.StatusNonAuthoritativeInfo";
            statusCodeMap[HttpStatusCode.ProxyAuthenticationRequired] = "http.StatusProxyAuthRequired";
            statusCodeMap[HttpStatusCode.RequestUriTooLong] = "http.StatusRequestURITooLong";
            statusCodeMap[tooManyRequests] = "http.StatusTooManyRequests";
            statusCodeMap[HttpStatusCode.HttpVersionNotSupported] = "http.StatusHTTPVersionNotSupported";

            StatusCodeToGoString = statusCodeMap;

            ReservedWords.AddRange(
                new[]
                {
                    // Reserved keywords -- list retrieved from http://golang.org/ref/spec#Keywords
                    "break",        "default",      "func",         "interface",    "select",
                    "case",         "defer",        "go",           "map",          "struct",
                    "chan",         "else",         "goto",         "package",      "switch",
                    "const",        "fallthrough",  "if",           "range",        "type",
                    "continue",     "for",          "import",       "return",       "var",        
                
                    // Reserved predeclared identifiers -- list retrieved from http://golang.org/ref/spec#Predeclared_identifiers
                    "bool", "byte",
                    "complex64", "complex128",
                    "error",
                    "float32", "float64",
                    "int", "int8", "int16", "int32", "int64",
                    "rune", "string",
                    "uint", "uint8", "uint16", "uint32", "uint64",
                    "uintptr",

                    "true", "false", "iota",

                    "nil",

                    "append", "cap", "close", "complex", "copy", "delete", "imag", "len", "make", "new", "panic", "print", "println", "real", "recover",


                    // Reserved packages -- list retrieved from http://golang.org/pkg/
                    // -- Since package names serve as partial identifiers, exclude the standard library
                    "archive", "tar", "zip",
                    "bufio",
                    "builtin",
                    "bytes",
                    "compress", "bzip2", "flate", "gzip", "lzw", "zlib",
                    "container", "heap", "list", "ring",
                    "crypto", "aes", "cipher", "des", "dsa", "ecdsa", "elliptic", "hmac", "md5", "rand", "rc4", "rsa", "sha1", "sha256", "sha512", "subtle", "tls", "x509", "pkix",
                    "database", "sql", "driver",
                    "debug", "dwarf", "elf", "gosym", "macho", "pe", "plan9obj",
                    "encoding", "ascii85", "asn1", "base32", "base64", "binary", "csv", "gob", "hex", "json", "pem", "xml",
                    "errors",
                    "expvar",
                    "flag",
                    "fmt",
                    "go", "ast", "build", "constant", "doc", "format", "importer", "parser", "printer", "scanner", "token", "types",
                    "hash", "adler32", "crc32", "crc64", "fnv",
                    "html", "template",
                    "image", "color", "palette", "draw", "gif", "jpeg", "png",
                    "index", "suffixarray",
                    "io", "ioutil",
                    "log", "syslog",
                    "math", "big", "cmplx", "rand",
                    "mime", "multipart", "quotedprintable",
                    "net", "http", "cgi", "cookiejar", "fcgi", "httptest", "httputil", "pprof", "mail", "rpc", "jsonrpc", "smtp", "textproto", "url",
                    "os", "exec", "signal", "user",
                    "path", "filepath",
                    "reflect",
                    "regexp", "syntax",
                    "runtime", "cgo", "debug", "pprof", "race", "trace",
                    "sort",
                    "strconv",
                    "strings",
                    "sync", "atomic",
                    "syscall",
                    "testing", "iotest", "quick",
                    "text", "scanner", "tabwriter", "template", "parse",
                    "time",
                    "unicode", "utf16", "utf8",
                    "unsafe",

                    // Other reserved names and packages (defined by the base libraries this code uses)
                    "autorest", "client", "date", "err", "req", "resp", "result", "sender", "to", "validation"

                });
        }

        /// <summary>
        /// Formats a string to work around golint name stuttering
        /// Refactor -> CodeModelTransformer
        /// </summary>
        /// <param name="name"></param>
        /// <param name="packageName"></param>
        /// <param name="nameInUse"></param>
        /// <param name="attachment"></param>
        /// <returns>The formatted string</returns>
        public static string AttachTypeName(string name, string packageName, bool nameInUse, string attachment)
        {
            return nameInUse
                ? name.Equals(packageName, StringComparison.OrdinalIgnoreCase)
                    ? name
                    : name + attachment
                : name;
        }

        /// <summary>
        /// Formats a string to pascal case using a specific character as splitter
        /// Refactor -> Namer ... Even better if this already exists in the core :D
        /// </summary>
        /// <param name="name"></param>
        /// <param name="splitter"></param>
        /// <returns>The formatted string</returns>
        public override string PascalCase(string name)
        {
            if (string.IsNullOrWhiteSpace(name))
            {
                return name;
            }

            return
                name.Split(new Char[]{'.', '_', '@', '-', ' ', '$'})
                    .Where(s => !string.IsNullOrEmpty(s))
                    .Select(s => char.ToUpperInvariant(s[0]) + s.Substring(1, s.Length - 1))
                    .DefaultIfEmpty("")
                    .Aggregate(string.Concat);
        }

        public override string GetEnumMemberName(string name)
        {
            return EnsureNameCase(base.GetEnumMemberName(name));
        }

        public override string GetFieldName(string name)
        {
            if (string.IsNullOrWhiteSpace(name))
            {
                return name;
            }
            return EnsureNameCase(RemoveInvalidCharacters(PascalCase(GetEscapedReservedName(name, "Field"))));
        }

        public override string GetInterfaceName(string name)
        {
            if (string.IsNullOrWhiteSpace(name))
            {
                return name;
            }
            return EnsureNameCase(RemoveInvalidCharacters(PascalCase(name)));
        }

        /// <summary>
        /// Formats a string for naming a method using Pascal case by default.
        /// </summary>
        /// <param name="name"></param>
        /// <returns>The formatted string.</returns>
        public override string GetMethodName(string name)
        {
            if (string.IsNullOrWhiteSpace(name))
            {
                return name;
            }
            return EnsureNameCase(GetEscapedReservedName(RemoveInvalidCharacters(PascalCase(name)), "Method"));
        }

        public override string GetMethodGroupName(string name)
        {
            if (string.IsNullOrWhiteSpace(name))
            {
                return name;
            }

            return EnsureNameCase(RemoveInvalidCharacters(PascalCase(name)));
        }

        /// <summary>
        /// Formats a string for naming method parameters using Camel case by default.
        /// </summary>
        /// <param name="name"></param>
        /// <returns>The formatted string.</returns>
        public override string GetParameterName(string name)
        {
            if (string.IsNullOrWhiteSpace(name))
            {
                return name;
            }
            if (Extensions.StartsWithAcronym(name))
            {
                return EnsureNameCase(GetEscapedReservedName((RemoveInvalidCharacters(name).ToLower()), "Parameter"));
            }
            return EnsureNameCase(GetEscapedReservedName(CamelCase(RemoveInvalidCharacters(name)), "Parameter"));
        }

        /// <summary>
        /// Formats a string for naming properties using Pascal case by default.
        /// </summary>
        /// <param name="name"></param>
        /// <returns>The formatted string.</returns>
        public override string GetPropertyName(string name)
        {
            if (string.IsNullOrWhiteSpace(name))
            {
                return name;
            }
            return EnsureNameCase(GetEscapedReservedName(RemoveInvalidCharacters(PascalCase(name)), "Property"));
        }

        /// <summary>
        /// Formats a string for naming a Type or Object using Pascal case by default.
        /// </summary>
        /// <param name="name"></param>
        /// <returns>The formatted string.</returns>
        public override string GetTypeName(string name)
        {
            if (string.IsNullOrWhiteSpace(name))
            {
                return name;
            }
            return EnsureNameCase(GetEscapedReservedName(RemoveInvalidCharacters(PascalCase(name)), "Type"));
        }

        /// <summary>
        /// Formats a string for naming a local variable using Camel case by default.
        /// </summary>
        /// <param name="name"></param>
        /// <returns>The formatted string.</returns>
        public override string GetVariableName(string name)
        {
            if (string.IsNullOrWhiteSpace(name))
            {
                return name;
            }
            return EnsureNameCase(GetEscapedReservedName(CamelCase(RemoveInvalidCharacters(name)), "Var"));
        }

        /// <summary>
        /// Converts names the conflict with Go reserved terms by appending the passed appendValue.
        /// </summary>
        /// <param name="name">Name.</param>
        /// <param name="appendValue">String to append.</param>
        /// <returns>The transformed reserved name</returns>
        protected override string GetEscapedReservedName(string name, string appendValue)
        {
            if (name == null)
            {
                throw new ArgumentNullException("name");
            }

            if (appendValue == null)
            {
                throw new ArgumentNullException("appendValue");
            }

            // Use case-sensitive comparisons to reduce generated names
            if (ReservedWords.Contains(name, StringComparer.Ordinal))
            {
                name += appendValue;
            }

            return name;
        }

        // Refactor -> Namer
        public void ReserveNamespace(string ns)
        {
            ReservedWords.Add(ns);
        }

        // EnsureNameCase ensures that all "words" in the passed name adhere to Golint casing expectations.
        // A "word" is a sequence of characters separated by a change in case or underscores. Since this
        // method alters name casing, it should be used after any other method that expects normal
        // camelCase or PascalCase.
        private string EnsureNameCase(string name)
        {
            var builder = new StringBuilder();
            var words = name.ToWords();
            for (int i = 0; i < words.Length; i++)
            {
                string word = words[i];
                if (CommonInitialisms.Contains(word))
                {
                    word = word.ToUpper();
                }
                else if (i < words.Length-1)
                {
                    // This ensures that names like `ClusterUsersGroupDNs`
                    // get propery cased to `ClusterUsersGroupDNS`
                    var concat = words[i] + words[i+1];
                    if (CommonInitialisms.Contains(concat.ToLower()))
                    {
                        word = concat.ToUpper();
                        i++;
                    }
                }
                builder.Append(word);
            }
            return builder.ToString();
        }

        public override string EscapeDefaultValue(string defaultValue, IModelType type)
        {
            if (type == null)
            {
                throw new ArgumentNullException("type");
            }
            PrimaryType primaryType = type as PrimaryType;
            if (defaultValue != null)
            {
                if (type is CompositeType)
                {
                    return type.Name + "{}";
                }
                else if (primaryType != null)
                {
                    if (primaryType.KnownPrimaryType == KnownPrimaryType.String
                        || primaryType.KnownPrimaryType == KnownPrimaryType.Uuid
                        || primaryType.KnownPrimaryType == KnownPrimaryType.TimeSpan
                        || primaryType.IsDateTimeType())
                    {
                        return CodeNamer.Instance.QuoteValue(defaultValue);
                    }
                    else if (primaryType.KnownPrimaryType == KnownPrimaryType.Boolean)
                    {
                        return defaultValue.ToLowerInvariant();
                    }
                    else if (primaryType.KnownPrimaryType == KnownPrimaryType.ByteArray)
                    {
                        return "[]byte(\"" + defaultValue + "\")";
                    }
                    else
                    {
                        //TODO: handle imports for package types.
                    }
                }
            }
            return defaultValue;
        }

        /// <summary>
        /// Returns true if client types should be exported.
        /// </summary>
        public bool ExportClientTypes
        {
            get
            {
                var setting = Settings.Instance.Host?.GetValue<string>("go-export-clients").Result;
                if (string.IsNullOrWhiteSpace(setting))
                {
                    // by default export client types
                    return true;
                }
                return bool.Parse(setting);
            }
        }

        /// <summary>
        /// Returns the value of the go-pipeline-import parameter or the empty string if it wasn't provided.
        /// </summary>
        public string PipelineImportPath
        {
            get
            {
                var path = Settings.Instance.Host?.GetValue<string>("go-pipeline-import").Result;
                if (string.IsNullOrWhiteSpace(path))
                {
                    // default location
                    return "github.com/Azure/azure-pipeline-go/pipeline";
                }
                return path;
            }
        }

        /// <summary>
        /// Gets the type name for the etag type.
        /// </summary>
        public string ETagTypeName => "ETag";
    }
}
