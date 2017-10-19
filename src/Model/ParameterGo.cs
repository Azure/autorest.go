// Copyright (c) Microsoft Open Technologies, Inc. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Utilities;
using AutoRest.Core.Model;
using AutoRest.Extensions;
using AutoRest.Extensions.Azure;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace AutoRest.Go.Model
{
    public class ParameterGo : Parameter
    {
        public const string APIVersionName = "APIVersion";
        public ParameterGo()
        {

        }

        /// <summary>
        /// Add imports for the parameter in parameter type.
        /// </summary>
        /// <param name="parameter"></param>
        /// <param name="imports"></param>
        public void AddImports(HashSet<string> imports)
        {
            ModelType.AddImports(imports);
        }

        /// <summary>
        /// Return string with formatted Go map.
        /// xyz["abc"] = 123
        /// </summary>
        /// <param name="mapVariable"></param>
        /// <returns></returns>
        public string AddToMap(string mapVariable)
        {
            return string.Format("{0}[\"{1}\"] = {2}", mapVariable, NameForMap(), ValueForMap(true));
        }

        public string GetParameterName()
        {
            string retval;
            if (IsAPIVersion)
            {
                retval = APIVersionName;
            }
            else if (IsClientProperty)
            {
                retval = "client." + Name.Value.Capitalize();
            }
            else
            {
                retval = Name.Value;
            }
            return retval;
        }

        public override bool IsClientProperty => base.IsClientProperty == true && !IsAPIVersion;

        public virtual bool IsAPIVersion => SerializedName.IsApiVersion();

        public virtual bool IsAPIHeader => SerializedName.IsApiHeader();

        public virtual bool IsMethodArgument => !IsClientProperty && !IsAPIVersion;

        public string HeaderCollectionPrefix => Extensions.GetValue<string>(SwaggerExtensions.HeaderCollectionPrefix);

        public bool IsHeaderCollection => !string.IsNullOrEmpty(HeaderCollectionPrefix);

        public bool IsCustomMetadata => SerializedName.StartsWith("x-ms-meta", StringComparison.OrdinalIgnoreCase);

        /// <summary>
        /// Get Name for parameter for Go map. 
        /// If parameter is client parameter, then return client.<parametername>
        /// </summary>
        /// <returns></returns>
        public string NameForMap()
        {
            return IsAPIVersion
                     ? AzureExtensions.ApiVersion
                     : SerializedName;
        }

        public bool RequiresUrlEncoding()
        {
            return (Location == Core.Model.ParameterLocation.Query || Location == Core.Model.ParameterLocation.Path) && !Extensions.ContainsKey(SwaggerExtensions.SkipUrlEncodingExtension);
        }

        /// <summary>
        /// Return formatted value string for the parameter.
        /// </summary>
        /// <returns></returns>
        public string ValueForMap(bool emitEncoding)
        {
            if (IsAPIVersion)
            {
                return APIVersionName;
            }

            var value = IsClientProperty
                ? "client." + CodeNamerGo.Instance.GetPropertyName(Name.Value)
                : Name.Value;

            var format = this.Format();

            var s = CollectionFormat != CollectionFormat.None
                                  ? $"{format},\"{CollectionFormat.GetSeparator()}\""
                                  : $"{format}";

            if (emitEncoding)
            {
                return this.EncodedString(s, value);
            }
            return string.Format(s, value);
        }

        public override IModelType ModelType
        {
            get
            {
                if (base.ModelType == null || !IsHeaderCollection)
                {
                    return base.ModelType;
                }
                // this is a header collection so emit it as a map[string]string
                return new DictionaryTypeGo() { ValueType = base.ModelType, CodeModel = base.ModelType.CodeModel, SupportsAdditionalProperties = false };
            }
            set
            {
                base.ModelType = value;
            }
        }

        /// <summary>
        /// Returns the RHS in an optional parameter value check.
        /// This is usually nil but is not in some cases.
        /// </summary>
        public string GetOptionalComparand()
        {
            if (IsRequired)
            {
                throw new Exception($"GetOptionalComparand called on required paramater {Name}");
            }

            if (ModelType is EnumTypeGo)
            {
                var et = ModelType as EnumTypeGo;
                var typeName = et.Name.ToString();
                if (typeName.EndsWith("Type"))
                {
                    typeName = typeName.Substring(0, typeName.Length - 4);
                }
                return $"{typeName}None";
            }
            return "nil";
        }

        /// <summary>
        /// Returns trus if the parameter should be passed by value.
        /// </summary>
        public bool IsPassedByValue()
        {
            return IsRequired || ModelType.CanBeNull() || ModelType is EnumTypeGo;
        }

        /// <summary>
        /// Returns true if the paramater is a stream that should be replaced by an io.ReadSeeker.
        /// </summary>
        public bool ReplaceStreamWithReadSeeker => ModelType.IsPrimaryType(KnownPrimaryType.Stream);
    }

    public static class ParameterGoExtensions
    {
        /// <summary>
        /// Returns the appropriate format string depending on if the paramater is passed by value.
        /// If the parameter is passed by reference then the parameter will need to be dereferenced.
        /// </summary>
        public static string Format(this ParameterGo parameter)
        {
            return  parameter.IsPassedByValue() ? "{0}" : "*{0}";
        }

        /// <summary>
        /// Wraps the parameter in a call to autorest.Encode() if the parameter requires URL encoding.
        /// </summary>
        public static string EncodedString(this ParameterGo parameter, string format, string value)
        {
            return string.Format(
                    parameter.RequiresUrlEncoding()
                        ? $"autorest.Encode(\"{parameter.Location.ToString().ToLower()}\",{format})"
                        : $"{format}",
                    value);
        }

        /// <summary>
        /// Wraps the parameter in a call to one of the string formatting functions (e.g. fmt.Sprintf) depending on the parameter's type.
        /// </summary>
        public static string GetStringFormat(this ParameterGo parameter, string defaultFormat)
        {
            if (parameter.ModelType.IsPrimaryType(KnownPrimaryType.String))
            {
                if (parameter.ModelType.IsETagType())
                {
                    // the etag is a format of the string primary type
                    return $"string({defaultFormat})";
                }
                return defaultFormat;
            }
            else if (parameter.ModelType.IsDateTimeType())
            {
                if (!parameter.IsRequired)
                {
                    // optional parameters are passed by reference, so defaultFormat
                    // will be dereferenced requiring us to surround it in parens.
                    // e.g. (*fooparam).Format(rfc339Format)
                    defaultFormat = $"({defaultFormat})";
                }
                
                if (parameter.ModelType.IsPrimaryType(KnownPrimaryType.DateTimeRfc1123))
                {
                    return $"{defaultFormat}.In(gmt).Format(time.RFC1123)";
                }
                return $"{defaultFormat}.Format(rfc3339Format)";
            }
            else
            {
                return $"fmt.Sprintf(\"%v\", {defaultFormat})";
            }
        }

        /// <summary>
        /// Return a Go map of required parameters.
        // Refactor -> Generator
        /// </summary>
        /// <param name="parameters"></param>
        /// <param name="mapVariable"></param>
        /// <returns></returns>
        public static string BuildParameterMap(this IEnumerable<ParameterGo> parameters, string mapVariable)
        {
            var builder = new StringBuilder();

            builder.Append(mapVariable);
            builder.Append(" := map[string]interface{} {");

            if (parameters.Count() > 0)
            {
                builder.AppendLine();
                var indented = new IndentedStringBuilder("  ");
                var paramsList = parameters
                    .Where(p => p.IsRequired)
                    .OrderBy(p => p.SerializedName.ToString());

                foreach (var p in paramsList)
                {
                    if (mapVariable == "queryParameters" && p.IsConstant)
                    {
                        var val = p.DefaultValue.ToString();
                        if ((p.ModelType.IsPrimaryType(KnownPrimaryType.String) || p.ModelType.IsDateTimeType()) && val[0] != '"')
                        {
                            val = $"\"{val}\"";
                        }
                        indented.AppendLine("\"{0}\": {1},", p.NameForMap(), p.EncodedString(p.Format(), val));
                    }
                    else
                    {
                        indented.AppendLine("\"{0}\": {1},", p.NameForMap(), p.ValueForMap(true));
                    }
                }

                builder.Append(indented);
            }

            builder.AppendLine("}");
            return builder.ToString();
        }

        /// <summary>
        /// Return list of parameters for specified location passed in an argument.
        /// Refactor -> Probably CodeModeltransformer, but even with 5 references, the other mkethods are not used anywhere
        /// </summary>
        /// <param name="parameters"></param>
        /// <param name="location"></param>
        /// <returns></returns>
        public static IEnumerable<ParameterGo> ByLocation(this IEnumerable<ParameterGo> parameters, ParameterLocation location)
        {
            return parameters
                .Where(p => p.Location == location);
        }

        /// <summary>
        /// Return list of retuired parameters for specified location passed in an argument.
        /// Refactor -> CodeModelTransformer, still, 3 erefences, but no one uses the other methods.
        /// </summary>
        /// <param name="parameters"></param>
        /// <param name="location"></param>
        /// <returns></returns>
        public static IEnumerable<ParameterGo> ByLocationAsRequired(this IEnumerable<ParameterGo> parameters, ParameterLocation location, bool isRequired)
        {
            return parameters
                .Where(p => p.Location == location && p.IsRequired == isRequired);
        }

        /// <summary>
        /// Return list of parameters as per their location in request.
        /// </summary>
        /// <param name="parameters"></param>
        /// <returns></returns>
        public static ParameterGo BodyParameter(this IEnumerable<ParameterGo> parameters)
        {
            var bodyParameters = parameters.ByLocation(ParameterLocation.Body);
            return bodyParameters.Any()
                    ? bodyParameters.First()
                    : null;
        }

        public static IEnumerable<ParameterGo> FormDataParameters(this IEnumerable<ParameterGo> parameters)
        {
            return parameters.ByLocation(ParameterLocation.FormData);
        }

        public static IEnumerable<ParameterGo> HeaderParameters(this IEnumerable<ParameterGo> parameters)
        {
            return parameters.ByLocation(ParameterLocation.Header);
        }

        public static IEnumerable<ParameterGo> HeaderParameters(this IEnumerable<ParameterGo> parameters, bool isRequired)
        {
            return parameters.ByLocationAsRequired(ParameterLocation.Header, isRequired);
        }

        public static IEnumerable<ParameterGo> URLParameters(this IEnumerable<ParameterGo> parameters)
        {
            var urlParams = new List<ParameterGo>();
            foreach (ParameterGo p in parameters.ByLocation(ParameterLocation.Path))
            {
                if (p.Method.CodeModel.BaseUrl.Contains(p.SerializedName))
                {
                    urlParams.Add(p);
                }
            }
            return urlParams;
        }

        public static IEnumerable<ParameterGo> PathParameters(this IEnumerable<ParameterGo> parameters)
        {
            var pathParams = new List<ParameterGo>();
            foreach (ParameterGo p in parameters.ByLocation(ParameterLocation.Path))
            {
                if (!p.Method.CodeModel.BaseUrl.Contains(p.SerializedName))
                {
                    pathParams.Add(p);
                }
            }
            return pathParams;
        }

        public static IEnumerable<ParameterGo> QueryParameters(this IEnumerable<ParameterGo> parameters)
        {
            return parameters.ByLocation(ParameterLocation.Query);
        }

        public static IEnumerable<ParameterGo> QueryParameters(this IEnumerable<ParameterGo> parameters, bool isRequired)
        {
            return parameters.ByLocationAsRequired(ParameterLocation.Query, isRequired);
        }

        public static string Validate(this IEnumerable<ParameterGo> parameters, HttpMethod method)
        {
            List<string> v = new List<string>();
            HashSet<string> ancestors = new HashSet<string>();

            foreach (var p in parameters)
            {
                if (p.IsAPIVersion)
                {
                    continue;
                }

                var name = !p.IsClientProperty
                        ? p.Name.Value
                        : "client." + p.Name.Value.Capitalize();

                List<string> x = new List<string>();
                if (p.ModelType is CompositeType)
                {
                    ancestors.Add(p.ModelType.Name);
                    x.AddRange(p.ValidateCompositeType(name, method, ancestors, false));
                    ancestors.Remove(p.ModelType.Name);
                }
                else
                {
                    x.AddRange(p.ValidateType(name, method, false));
                }

                if (x.Count != 0)
                {
                    v.Add($"{{ targetValue: {name},\n constraints: []constraint{{{string.Join(",\n", x)}}}}}");
                }
            }
            return string.Join(",\n", v);
        }
    }
}
