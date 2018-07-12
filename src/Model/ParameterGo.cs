// Copyright (c) Microsoft Open Technologies, Inc. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Utilities;
using AutoRest.Core.Model;
using AutoRest.Extensions;
using AutoRest.Extensions.Azure;
using System.Collections.Generic;
using System.Globalization;
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
        /// Add imports for the parameter model type.
        /// </summary>
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
        /// <param name="useDefaultValue"></param>
        /// <returns></returns>
        public string AddToMap(string mapVariable, bool useDefaultValue = false)
        {
            return string.Format("{0}[\"{1}\"] = {2}", mapVariable, NameForMap(), ValueForMap(useDefault: useDefaultValue));
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

        public virtual bool IsMethodArgument => !IsClientProperty && !IsAPIVersion && !IsConstant;

        /// <summary>
        /// Returns a properly formatted DefaultValue string.
        /// </summary>
        public string DefaultValueString
        {
            // unfortunately the modeler doesn't uniformly wrap default values in double quotes, so
            // depending on the type's format it might or might not be quoted.  e.g. plain ol' strings
            // will be double-quoted but a string in date/time format will not.  note that there can
            // be other "interesting" default values, e.g. []byte(""), so you can't simply check for
            // the absense of double-quotes and then add them.  right now the only affected type is
            // date/times, if we find more cases this will need to be updated.
            get
            {
                // another irritant is that the javascript front-end to autorest will
                // munge certain decimals/doubles, e.g. it will change 1.034E+20 to
                // 103400000000000000000 which we don't want, so we have to round-trip
                // these types to get the desired output.
                if (string.IsNullOrWhiteSpace(DefaultValue))
                {
                    return DefaultValue;
                }

                if (ModelType is PrimaryTypeGo primaryType)
                {
                    if (primaryType.KnownFormat.IsDateTime())
                    {
                        return $"\"{DefaultValue}\"";
                    }
                    else if (primaryType.KnownPrimaryType == KnownPrimaryType.Decimal)
                    {
                        var asDecimal = decimal.Parse(DefaultValue);
                        return asDecimal.ToString(CultureInfo.InvariantCulture);
                    }
                    else if (primaryType.KnownPrimaryType == KnownPrimaryType.Double)
                    {
                        var asDouble = double.Parse(DefaultValue);
                        return asDouble.ToString(CultureInfo.InvariantCulture);
                    }
                }
                else if (ModelType is EnumTypeGo)
                {
                    return $"\"{DefaultValue}\"";
                }

                return DefaultValue;
            }
        }

        /// <summary>
        /// Get Name for parameter for Go map.
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
        public string ValueForMap(bool useDefault = false)
        {
            if (IsAPIVersion)
            {
                return APIVersionName;
            }

            if (IsConstant)
            {
                return RequiresUrlEncoding() ? 
                    $"autorest.Encode(\"{Location.ToString().ToLower()}\", {DefaultValueString})" :
                    DefaultValueString;
            }

            string value = "";

            if (useDefault)
            {
                value = DefaultValueString;
            }
            else if (IsClientProperty)
            {
                var propName = CodeNamerGo.Instance.GetPropertyName(Name.Value);
                // verify that the calculated name matches the property name
                bool found = false;
                foreach (var clientProp in Method.CodeModel.Properties)
                {
                    // prefer an exact match
                    if (clientProp.Name == propName)
                    {
                        found = true;
                        break;
                    }
                }
                if (!found)
                {
                    // didn't find an exact match, find the closest match.  we hit this case if
                    // the front-end decided to add a suffix (e.g. '1') to the client property name.
                    foreach (var clientProp in Method.CodeModel.Properties)
                    {
                        if (clientProp.Name.ToString().IndexOf(propName) > -1)
                        {
                            propName = clientProp.Name;
                            break;
                        }
                    }
                }
                value = "client." + propName;
            }
            else
            {
                value = Name.Value;
            }

            var format = IsRequired || ModelType.CanBeEmpty() || useDefault
                                          ? "{0}"
                                          : "*{0}";

            var s = CollectionFormat != CollectionFormat.None
                                  ? $"{format},\"{CollectionFormat.GetSeparator()}\""
                                  : $"{format}";

            return string.Format(
                RequiresUrlEncoding()
                    ? $"autorest.Encode(\"{Location.ToString().ToLower()}\",{s})"
                    : $"{s}",
                value);
        }

        public string GetEmptyCheck(string valueReference, bool asEmpty = true)
        {
            if (ModelType is PrimaryTypeGo goPrimaryType)
            {
                return GetPrimaryTypeEmptyCheck(goPrimaryType, valueReference, asEmpty);
            }
            else if (ModelType is SequenceTypeGo)
            {
                return GetSequenceTypeEmptyCheck(valueReference, asEmpty);
            }
            else if (ModelType is DictionaryTypeGo)
            {
                return GetDictionaryEmptyCheck(valueReference, asEmpty);
            }
            else if (ModelType is EnumTypeGo)
            {
                return GetEnumEmptyCheck(valueReference, asEmpty);
            }
            else
            {
                return string.Format(asEmpty
                                        ? "{0} == nil"
                                        : "{0} != nil", valueReference);
            }
        }

        private string GetDictionaryEmptyCheck(string valueReference, bool asEmpty)
        {
            return string.Format(asEmpty
                                    ? "{0} == nil || len({0}) == 0"
                                    : "{0} != nil && len({0}) > 0", valueReference);
        }

        private string GetEnumEmptyCheck(string valueReference, bool asEmpty)
        {
            return string.Format(asEmpty
                                    ? "len(string({0})) == 0"
                                    : "len(string({0})) > 0", valueReference);
        }

        private string GetPrimaryTypeEmptyCheck(PrimaryTypeGo pt, string valueReference, bool asEmpty)
        {
            if (pt.PrimaryType(KnownPrimaryType.ByteArray))
            {
                return string.Format(asEmpty
                                        ? "{0} == nil || len({0}) == 0"
                                        : "{0} != nil && len({0}) > 0", valueReference);
            }
            else if (pt.PrimaryType(KnownPrimaryType.String))
            {
                return string.Format(asEmpty
                                        ? "len({0}) == 0"
                                        : "len({0}) > 0", valueReference);
            }
            else
            {
                return string.Format(asEmpty
                                        ? "{0} == nil"
                                        : "{0} != nil", valueReference);
            }
        }

        private string GetSequenceTypeEmptyCheck(string valueReference, bool asEmpty)
        {
            return string.Format(asEmpty
                                   ? "{0} == nil || len({0}) == 0"
                                   : "{0} != nil && len({0}) > 0", valueReference);
        }

        /// <summary>
        /// Returns true if two parameters are semantically equivalent.
        /// The names match (excluding casing) and the types are identical.
        /// </summary>
        /// <param name="lhs">Left-hand side to compare against.</param>
        /// <param name="rhs">Right-hand side to compare with.</param>
        /// <returns>True if the two are semantically equal.</returns>
        public static bool Match(ParameterGo lhs, ParameterGo rhs)
        {
            return lhs.Name.EqualsIgnoreCase(rhs.Name) && lhs.ModelTypeName.Equals(rhs.ModelTypeName);
        }
    }

    public static class ParameterGoExtensions
    {
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

            if (parameters.Any())
            {
                builder.AppendLine();
                var indented = new IndentedStringBuilder("  ");
                parameters
                    .Where(p => p.IsRequired)
                    .OrderBy(p => p.SerializedName.ToString())
                    .ForEach(p => indented.AppendLine("\"{0}\": {1},", p.NameForMap(), p.ValueForMap()));
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

        public static IEnumerable<ParameterGo> FormDataParameters(this IEnumerable<ParameterGo> parameters, bool isRequired)
        {
            return parameters.ByLocationAsRequired(ParameterLocation.FormData, isRequired);
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
                if (p.IsAPIVersion || p.IsConstant)
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
                    x.AddRange(p.ValidateCompositeType(name, method, ancestors));
                    ancestors.Remove(p.ModelType.Name);
                }
                else
                    x.AddRange(p.ValidateType(name, method));

                if (x.Count != 0)
                    v.Add($"{{ TargetValue: {name},\n Constraints: []validation.Constraint{{{string.Join(",\n", x)}}}}}");
            }
            return string.Join(",\n", v);
        }
    }
}
