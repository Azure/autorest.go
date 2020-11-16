// Copyright (c) Microsoft Open Technologies, Inc. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core.Model;
using AutoRest.Core.Utilities;
using AutoRest.Go.Model;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;

namespace AutoRest.Go
{
    public static class Extensions
    {
        public const string NullConstraint = "Null";

        public const string EmptyConstraint = "Empty";

        public const string ReadOnlyConstraint = "ReadOnly";

        private static readonly Regex IsApiVersionPattern = new Regex(@"^api[^a-zA-Z0-9_]?version", RegexOptions.IgnoreCase);

        private static readonly Regex UnwrapAnchorTagsPattern = new Regex("([^<>]*)<a\\s*.*\\shref\\s*=\\s*[\'\"]([^\'\"]*)[\'\"][^>]*>(.*)</a>");

        private static readonly Regex WordSplitPattern = new Regex(@"(\p{Lu}\p{Ll}+)");

        private static readonly Dictionary<string, string> plural = new Dictionary<string, string>
        {
            { "eventhub", "eventhubs" },
            { "containerservice", "containerservices" }
        };

        /// <summary>
        /// Resets variant static data.  Call between batches.
        /// </summary>
        public static void ResetState()
        {
            s_interfaceNames.Clear();
            s_wordMap.Clear();
        }

        // contains a map from a model type to its corresponding interface name
        private static Dictionary<IModelType, string> s_interfaceNames = new Dictionary<IModelType, string>();

        // contains a map from a one-word string to multiple words e.g. "FooBarBaz" => ["Foo", "Bar", "Baz"]
        private static Dictionary<string, string[]> s_wordMap = new Dictionary<string, string[]>();

        /////////////////////////////////////////////////////////////////////////////////////////
        //
        // General Extensions
        //
        /////////////////////////////////////////////////////////////////////////////////////////

        /// <summary>
        /// This method changes string to sentence where is make the first word
        /// of sentence to lowercase (unless it is an acronym). The sentence is coming directly from swagger.
        /// </summary>
        /// <param name="value"></param>
        /// <returns></returns>
        public static string ToSentence(this string value)
        {
            if (string.IsNullOrWhiteSpace(value))
            {
                return string.Empty;
            }
            else
            {
                value = value.Trim();
                if (value.StartsWithAcronym())
                {
                    return value;
                }
                return value.First().ToString().ToLowerInvariant() + (value.Length > 1 ? value.Substring(1) : "");
            }
        }

        /// <summary>
        /// Determines if the first word in a string is an acronym
        /// (acronym defined as all caps and more than 1 char)
        /// </summary>
        /// <param name="value"></param>
        /// <returns></returns>
        public static bool StartsWithAcronym(this string value)
        {
            string firstWord = value.Trim().Split(' ', '-', '_').First();
            return firstWord.Length > 1 && firstWord.All(char.IsUpper);
        }

        /// <summary>
        /// Makes first word of the sentence as upper case.
        /// </summary>
        /// <param name="value"></param>
        /// <returns></returns>
        public static string Capitalize(this string value)
        {
            return string.IsNullOrWhiteSpace(value)
                    ? string.Empty
                    : value.First()
                           .ToString()
                           .ToUpperInvariant() + (value.Length > 1 ? value.Substring(1) : "");
        }

        /// <summary>
        /// String manipulation function converts all words in a sentence to lowercase.
        /// Refactor -> Namer
        /// </summary>
        /// <param name="value"></param>
        /// <returns></returns>
        public static string ToPhrase(this string value)
        {
            List<string> words = new List<string>(value.ToWords());
            for (int i = 0; i < words.Count; i++)
            {
                words[i] = words[i].ToLowerInvariant();
            }
            return string.Join(" ", words.ToArray());
        }

        /// <summary>
        /// Split sentence into words.
        /// </summary>
        /// <param name="value"></param>
        /// <returns></returns>
        public static string[] ToWords(this string value)
        {
            if (!s_wordMap.ContainsKey(value))
            {
                s_wordMap.Add(value, WordSplitPattern.Split(value).Where(s => !string.IsNullOrEmpty(s)).ToArray());
            }
            return s_wordMap[value];
        }

        /// <summary>
        /// Creates a string from the first letter in each word.
        /// E.g. "SomeTypeName" would generate the string "stn".
        /// </summary>
        /// <param name="scope">Provide an instance to ensure variable names are unique within a given scope.</param>
        public static string ToVariableName(this string longName, VariableScopeProvider scope = null)
        {
            var initials = from word in longName.ToWords()
                           select word[0];
            var acronym = string.Concat(initials).ToLowerInvariant();
            var name = CodeNamerGo.Instance.GetVariableName(acronym);
            if (scope != null)
            {
                name = scope.GetVariableName(name);
            }
            return name;
        }

        /// <summary>
        /// Creates a string from the first letter in each word.
        /// E.g. "SomeTypeName" would generate the string "stn".
        /// </summary>
        /// <param name="scope">Provide an instance to ensure variable names are unique within a given scope.</param>
        public static string ToVariableName(this Fixable<string> longName, VariableScopeProvider scope = null)
        {
            return longName.ToString().ToVariableName(scope);
        }

        /// <summary>
        /// This method checks if MethodGroupName is plural of package name.
        /// It returns false for packages not listed in dictionary 'plural'.
        /// Example, group EventHubs in package EventHub.
        /// Refactor -> Namer, but also could be used by the CodeModelTransformer
        /// </summary>
        /// <param name="value"></param>
        /// <param name="packageName"></param>
        /// <returns></returns>
        public static bool IsNamePlural(this string value, string packageName)
        {
            return plural.ContainsKey(packageName) && plural[packageName] == value.ToLower();
        }

        /// <summary>
        /// Gets substring from value string.
        /// </summary>
        /// <param name="value"></param>
        /// <param name="s"></param>
        /// <returns></returns>
        public static string TrimStartsWith(this string value, string s)
        {
            if (!string.IsNullOrEmpty(s) && s.Length < value.Length && value.StartsWith(s, StringComparison.OrdinalIgnoreCase))
            {
                value = value.Substring(s.Length);
            }
            return value;
        }

        /// <summary>
        /// Removes the package name from the beginning of a type or method name.
        /// Refactor -> Namer, but could be used by the CodeModelTransformer too
        /// </summary>
        /// <param name="value">The name of the type or method from which to remove the package name.</param>
        /// <param name="packageName">The name of the package to be removed.</param>
        /// <returns>A string containing the modified name.</returns>
        public static string TrimPackageName(this string value, string packageName)
        {
            // check if the package name straddles a casing boundary, if it
            // does then don't trim the name.  e.g. if value == "SubscriptionState"
            // and packageName == "subscriptions" it would be incorrect to remove
            // the package name from the value.

            bool straddle = value.Length > packageName.Length && !char.IsUpper(value[packageName.Length]);

            var originalLen = value.Length;

            if (!straddle)
                value = value.TrimStartsWith(packageName);

            // if nothing was trimmed and the package name is plural then make the
            // package name singular and try removing that if it's not too short
            if (value.Length == originalLen && packageName.EndsWith("s") && (value.Length - packageName.Length + 1) > 1)
            {
                value = value.TrimPackageName(packageName.Substring(0, packageName.Length - 1));
            }

            return value;
        }

        /// <summary>
        /// Converts List to formatted string of arguments.
        /// Refactor -> Generator
        /// </summary>
        /// <param name="arguments"></param>
        /// <returns></returns>
        public static string EmitAsArguments(this IEnumerable<string> arguments)
        {
            return String.Join(",\n", arguments);
        }


        // This function removes html anchor tags and reformats the comment text.
        // For example, Swagger documentation text --> "This is a documentation text. For information see <a href=LINK">CONTENT.</a>"
        // reformats to  --> "This is a documentation text. For information see CONTENT (LINK)."
        // Refactor -> Namer
        // Still, nobody uses this...
        public static string UnwrapAnchorTags(this string comments)
        {

            Match match = UnwrapAnchorTagsPattern.Match(comments);

            if (match.Success)
            {
                string content = match.Groups[3].Value;
                string link = match.Groups[2].Value;

                return (".?!;:".Contains(content[content.Length - 1])
                        ? match.Groups[1].Value + content.Substring(0, content.Length - 1) + " (" + link + ")" + content[content.Length - 1]
                        : match.Groups[1].Value + content + " (" + link + ")");
            }

            return comments;
        }

        /// <summary>
        /// Return the separator associated with a given collectionFormat
        /// It looks like other generators use this for split / join operations ?
        /// Refactor -> I think CodeMoedelTransformer
        /// </summary>
        /// <param name="format">The collection format</param>
        /// <returns>The separator</returns>
        public static string GetSeparator(this CollectionFormat format)
        {
            switch (format)
            {
                case CollectionFormat.Csv:
                    return ",";
                case CollectionFormat.Pipes:
                    return "|";
                case CollectionFormat.Ssv:
                    return " ";
                case CollectionFormat.Tsv:
                    return "\t";
                default:
                    throw new NotSupportedException(string.Format("Collection format {0} is not supported.", format));
            }
        }

        /// <summary>
        /// Returns true if the specified CollectionFormat requires a separator.
        /// </summary>
        public static bool CollectionFormatRequiresSeparator(this CollectionFormat format)
        {
            return format != CollectionFormat.None && format != CollectionFormat.Multi;
        }

        public static bool IsApiVersion(this string name)
        {
            return IsApiVersionPattern.IsMatch(name);
        }

        /// <summary>
        /// Transforms the specified text body into a comment block.
        /// If the body contains embedded new-line characters it will be
        /// broken up into multiple comment-lines.
        /// </summary>
        /// <param name="body">The body of text to transform into a comment block.</param>
        /// <returns>The comment block.</returns>
        public static string ToCommentBlock(this string body)
        {
            if (body.IndexOf('\n') > 0)
            {
                body = string.Join("\n// ", body.Split('\n', StringSplitOptions.RemoveEmptyEntries));
            }
            return $"// {body}\n";
        }

        /////////////////////////////////////////////////////////////////////////////////////////
        //
        // Type Extensions
        //
        /////////////////////////////////////////////////////////////////////////////////////////

        public static bool IsStreamType(this IModelType body)
        {
            return body is CompositeTypeGo r && (r.BaseType.PrimaryType(KnownPrimaryType.Stream));
        }

        public static bool PrimaryType(this IModelType type, KnownPrimaryType typeToMatch)
        {
            if (type == null)
            {
                return false;
            }

            return type is PrimaryType primaryType && primaryType.KnownPrimaryType == typeToMatch;
        }

        public static bool CanBeEmpty(this IModelType type)
        {
            return type is DictionaryType
                || (type is PrimaryType primaryType
                 && (primaryType.KnownPrimaryType == KnownPrimaryType.ByteArray
                        || primaryType.KnownPrimaryType == KnownPrimaryType.Stream
                        || primaryType.KnownPrimaryType == KnownPrimaryType.String))
                || type is SequenceType
                || type is EnumType;
        }

        /// <summary>
        /// Returns true if the specified type can be implicitly null.
        /// E.g. things like maps, arrays, interfaces etc can all be null.
        /// </summary>
        /// <param name="type">The type to inspect.</param>
        /// <returns>True if the specified type can be null.</returns>
        public static bool CanBeNull(this IModelType type)
        {
            return 
                type is DictionaryType
                || type is SequenceType
                || (type is PrimaryType primaryType
                    && (primaryType.KnownPrimaryType == KnownPrimaryType.ByteArray
                        || primaryType.KnownPrimaryType == KnownPrimaryType.Stream
                        || primaryType.KnownPrimaryType == KnownPrimaryType.Object));
        }

        /// <summary>
        /// Add imports for a type.
        /// </summary>
        /// <param name="type"></param>
        /// <param name="imports"></param>
        public static void AddImports(this IModelType type, HashSet<string> imports)
        {
            if (type is DictionaryTypeGo dictionaryType)
            {
                dictionaryType.AddImports(imports);
            }
            else if (type is PrimaryTypeGo primaryType)
            {
                primaryType.AddImports(imports);
            }
            else
            {
                (type as SequenceTypeGo)?.AddImports(imports);
            }
        }

        public static bool ShouldBeSyntheticType(this IModelType type)
        {
            return (type is PrimaryType || type is SequenceType || type is DictionaryType || type is EnumType ||
                (type is CompositeType && (type as CompositeTypeGo).IsPolymorphicResponse()));
        }

        /// <summary>
        /// Gets if the type has an interface.
        /// </summary>
        public static bool HasInterface(this IModelType type)
        {
            return (type is CompositeTypeGo compositeType) &&
                (compositeType.IsRootType || compositeType.BaseIsPolymorphic && !compositeType.IsLeafType);
        }

        /// <summary>
        /// Gets the interface name for the type.
        /// </summary>
        /// <param name="includePkgName">Pass true if the type name should include the package prefix.  Defaults to false.</param>
        /// <returns></returns>
        public static string GetInterfaceName(this IModelType type, bool includePkgName = false)
        {
            // this function is called in a *LOT* of places so is perf sensitive
            if (!s_interfaceNames.ContainsKey(type))
            {
                var interfaceName = $"Basic{type.Name}";
                if (type.CodeModel.AllModelTypes.Any(mt => mt.Name == interfaceName))
                {
                    interfaceName = $"Basic{interfaceName}";
                }
                s_interfaceNames.Add(type, interfaceName);
            }
            return includePkgName
                ? $"{type.CodeModel.Namespace}.{s_interfaceNames[type]}"
                : s_interfaceNames[type];
        }

        /// <summary>
        /// Determines whether one composite type derives directly or indirectly from another.
        /// </summary>
        /// <param name="type">Type to test.</param>
        /// <param name="possibleAncestorType">Type that may be an ancestor of this type.</param>
        /// <returns>true if the type is an ancestor, false otherwise.</returns>
        public static bool DerivesFrom(this CompositeType type, CompositeType possibleAncestorType)
        {
            return
                type.BaseModelType != null &&
                (type.BaseModelType.Equals(possibleAncestorType) ||
                 type.BaseModelType.DerivesFrom(possibleAncestorType));
        }

        /// <summary>
        /// Casts the specified IModelType to the specified type paramater or throws an InvalidCastException.
        /// </summary>
        /// <typeparam name="T">The type to cast to.</typeparam>
        /// <param name="type">The type to cast from.</param>
        /// <returns>The type converted to T.</returns>
        public static T Cast<T>(this IModelType type)
        {
            return (T)type;
        }

        /// <summary>
        /// Converts the specified composite type into a page type.
        /// If the type specified is a future it will "unwrap" the page type from it.
        /// Throws if the specified type is not a future or page type.
        /// </summary>
        /// <param name="ctg">The type to convert from.</param>
        /// <returns>The type converted to a page type.</returns>
        internal static PageTypeGo UnwrapPageType(this CompositeTypeGo ctg)
        {
            PageTypeGo result;
            if (ctg is FutureTypeGo ftg)
            {
                result = ftg.ResultType.Cast<PageTypeGo>();
            }
            else if (ctg is PageTypeGo ptg)
            {
                result = ptg;
            }
            else
            {
                throw new InvalidCastException("supplied object is not a FutureTypeGo or PageTypeGo");
            }
            return result;
        }

        /// <summary>
        /// Returns an expression for zero-initializing the specified type.
        /// </summary>
        /// <param name="type">The type for which to create a zero-init expression.</param>
        /// <returns>The zero-init expression.</returns>
        public static string GetZeroInitExpression(this IModelType type)
        {
            if (type is CompositeTypeGo ctg)
            {
                return ctg.ZeroInitExpression;
            }
            else if (type is EnumTypeGo etg)
            {
                return etg.ZeroInitExpression;
            }
            else if (type is PrimaryTypeGo ptg)
            {
                return ptg.ZeroInitExpression;
            }
            else if (type is DictionaryTypeGo dtg)
            {
                return dtg.ZeroInitExpression;
            }
            throw new NotImplementedException($"GetZeroInitExpression for type {type} NYI");
        }

        /// <summary>
        /// Returns true if the format is a date/time.
        /// </summary>
        /// <param name="format">The format type to check.</param>
        /// <returns>True if the format is a date/time.</returns>
        public static bool IsDateTime(this KnownFormat format)
        {
            return format == KnownFormat.date ||
                format == KnownFormat.date_time ||
                format == KnownFormat.date_time_rfc1123;
        }

        /// <summary>
        /// Returns true if the specified type is user-defined.
        /// </summary>
        public static bool IsUserDefinedType(this IModelType type)
        {
            return (type is CompositeTypeGo) || (type is EnumTypeGo etg && etg.IsNamed);
        }

        /////////////////////////////////////////////////////////////////////////////////////////
        // Validate code
        //
        // This code generates a validation object which is defined in
        // go-autorest/autorest/validation package and is used to validate
        // constraints.
        // See PR: https://github.com/Azure/go-autorest/tree/master/autorest/validation
        //
        /////////////////////////////////////////////////////////////////////////////////////////



        /// <summary>
        /// Return list of validations for primary, map, sequence and rest of the types.
        /// </summary>
        /// <param name="p"></param>
        /// <param name="name"></param>
        /// <param name="method"></param>
        /// <param name="isCompositeProperties"></param>
        /// <returns></returns>
        public static List<string> ValidateType(this IVariable p, string name, HttpMethod method,
            bool isCompositeProperties = false)
        {
            List<string> x = new List<string>();
            if (method != HttpMethod.Patch || !p.IsBodyParameter() || isCompositeProperties)
            {
                x.AddRange(p.Constraints
                    .Where(c => c.IsValidConstraint())
                    .Select(c => GetConstraint(name, p.ModelTypeName, c.Key.ToString(), c.Value)).ToList());
            }

            List<string> y = new List<string>();
            if (x.Count > 0)
            {
                if (p.CheckNull() || isCompositeProperties)
                    y.AddRange(x.AddChain(name, p.ModelTypeName, NullConstraint, p.IsRequired));
                else if (!p.IsRequired && p.ModelType.PrimaryType(KnownPrimaryType.String))
                    y.AddRange(x.AddChain(name, p.ModelTypeName, EmptyConstraint, p.IsRequired));
                else
                    y.AddRange(x);
            }
            else
            {
                if (p.IsRequired && (p.CheckNull() || isCompositeProperties))
                    y.AddNullValidation(name, p.ModelTypeName, p.IsRequired);
            }
            return y;
        }

        /// <summary>
        /// Return list of validations for composite type.
        /// </summary>
        /// <param name="p"></param>
        /// <param name="name"></param>
        /// <param name="method"></param>
        /// <param name="ancestors"></param>
        /// <param name="isCompositeProperties"></param>
        /// <returns></returns>
        public static List<string> ValidateCompositeType(this IVariable p, string name, HttpMethod method, HashSet<string> ancestors,
            bool isCompositeProperties = false)
        {
            List<string> x = new List<string>();
            if (method != HttpMethod.Patch || !p.IsBodyParameter() || isCompositeProperties)
            {
                foreach (var prop in ((CompositeType)p.ModelType).Properties.Cast<PropertyGo>())
                {
                    var primary = prop.ModelType as PrimaryType;
                    var sequence = prop.ModelType as SequenceType;
                    var map = prop.ModelType as DictionaryTypeGo;
                    var composite = prop.ModelType as CompositeType;

                    var propName = prop.FieldName;

                    if (primary != null || sequence != null || map != null)
                    {
                        x.AddRange(prop.ValidateType($"{name}.{propName}", method, true));
                    }
                    else if (composite != null)
                    {
                        if (ancestors.Contains(composite.Name))
                        {
                            x.AddNullValidation($"{name}.{propName}", p.ModelTypeName, p.IsRequired);
                        }
                        else
                        {
                            ancestors.Add(composite.Name);
                            x.AddRange(prop.ValidateCompositeType($"{name}.{propName}", method, ancestors, true));
                            ancestors.Remove(composite.Name);
                        }
                    }
                }
            }

            List<string> y = new List<string>();
            if (x.Count > 0)
            {
                if (p.CheckNull() || isCompositeProperties)
                    y.AddRange(x.AddChain(name, p.ModelTypeName, NullConstraint, p.IsRequired));
                else
                    y.AddRange(x);
            }
            else
            {
                if (p.IsRequired && (p.CheckNull() || isCompositeProperties))
                    y.AddNullValidation(name, p.ModelTypeName, p.IsRequired);
            }
            return y;
        }

        /// <summary>
        /// Add null validation in validation object.
        /// </summary>
        /// <param name="v"></param>
        /// <param name="name"></param>
        /// <param name="isRequired"></param>
        public static void AddNullValidation(this List<string> v, string name, string type, bool isRequired = false)
        {
            v.Add(GetConstraint(name, type, NullConstraint, $"{isRequired}".ToLower()));
        }

        /// <summary>
        /// Add chain of validation for composite type.
        /// </summary>
        /// <param name="x"></param>
        /// <param name="name"></param>
        /// <param name="constraint"></param>
        /// <param name="isRequired"></param>
        /// <returns></returns>
        public static List<string> AddChain(this List<string> x, string name, string type, string constraint, bool isRequired)
        {
            List<string> a = new List<string>
            {
                GetConstraint(name, type, constraint, $"{isRequired}".ToLower(), true),
                $"Chain: []validation.Constraint{{{x[0]}"
            };
            a.AddRange(x.GetRange(1, x.Count - 1));
            a.Add("}}");
            return a;
        }
        /// <summary>
        /// CheckNull
        /// </summary>
        /// <param name="p"></param>
        /// <returns></returns>
        // Check if type is not a null or pointer type.
        public static bool CheckNull(this IVariable p)
        {
            return p is Parameter && (p.ModelType.IsNullValueType() || !(p.IsRequired || p.ModelType.CanBeEmpty()));
        }

        /// <summary>
        /// Check whether a type is nullable type.
        /// </summary>
        /// <param name="t"></param>
        /// <returns></returns>
        public static bool IsNullValueType(this IModelType t)
        {
            return t is DictionaryType
                || (t is PrimaryType primaryType
                   && primaryType.KnownPrimaryType == KnownPrimaryType.ByteArray)
                || t is SequenceType;
        }

        /// <summary>
        /// Returns true if the specified type is a numeric type.
        /// </summary>
        /// <param name="t"></param>
        /// <returns></returns>
        public static bool IsNumericType(this IModelType t)
        {
            return t.PrimaryType(KnownPrimaryType.Decimal) ||
                t.PrimaryType(KnownPrimaryType.Double) ||
                t.PrimaryType(KnownPrimaryType.Int) ||
                t.PrimaryType(KnownPrimaryType.Long);
        }

        /// <summary>
        /// Check if parameter is a body parameter.
        /// </summary>
        /// <param name="p"></param>
        /// <returns></returns>
        public static bool IsBodyParameter(this IVariable p)
        {
            return p is Parameter parameter && parameter.Location == ParameterLocation.Body;
        }

        /// <summary>
        /// Construct validation string for validation object for the passed constraint.
        /// </summary>
        /// <param name="name"></param>
        /// <param name="constraintName"></param>
        /// <param name="constraintValue"></param>
        /// <param name="chain"></param>
        /// <returns></returns>
        public static string GetConstraint(string name, string type, string constraintName, string constraintValue, bool chain = false)
        {
            var value = constraintValue;
            if (constraintName == Constraint.Pattern.ToString())
            {
                value = $"`{constraintValue}`";
            }
            else if (constraintName == Constraint.InclusiveMaximum.ToString() ||
                     constraintName == Constraint.InclusiveMinimum.ToString())
            {
                // swagger spec states that InclusiveMaximum should be a number
                // however the validation code supports int64 and float64.  to be
                // on the safe side handle both cases here.
                switch (type)
                {
                    case "float64":
                        value = $"float64({constraintValue})";
                        break;
                    case "int32":
                    case "int64":
                        value = $"int64({constraintValue})";
                        break;
                    default:
                        throw new InvalidOperationException($"Constraint {constraintName} only supports numbers, but got ${type}");
                }
            }

            return string.Format(chain
                                    ? "\t{{Target: \"{0}\", Name: validation.{1}, Rule: {2} "
                                    : "\t{{Target: \"{0}\", Name: validation.{1}, Rule: {2}, Chain: nil }}",
                                    name, constraintName, value);
        }

        /// <summary>
        /// Returns true if the specified constraint can be expressed in Go.
        /// </summary>
        private static bool IsValidConstraint(this KeyValuePair<Constraint, string> constraint)
        {
            // Go's regex engine doesn't support positive or negative lookaheads or
            // lookbehinds, so if the constraint contain any of them we will omit it.
            if (constraint.Key == Constraint.Pattern &&
                (constraint.Value.Contains("?=") || constraint.Value.Contains("?!") ||
                 constraint.Value.Contains("?<=") || constraint.Value.Contains("?<!")))
            {
                return false;
            }
            return true;
        }
    }
}
