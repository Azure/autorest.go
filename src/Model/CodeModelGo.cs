// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core;
using AutoRest.Core.Model;
using AutoRest.Core.Utilities;
using AutoRest.Extensions;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;

namespace AutoRest.Go.Model
{
    public class CodeModelGo : CodeModel
    {
        private Dictionary<FutureTypeGo, FutureTypeGo> _futureTypes;

        private static readonly Regex semVerPattern = new Regex(@"^v?(?<major>\d+)\.(?<minor>\d+)\.(?<patch>\d+)(?:-(?<tag>\S+))?$", RegexOptions.Compiled);

        public CodeModelGo()
        {
            NextMethodUndefined = new List<IModelType>();
            PagedTypes = new Dictionary<IModelType, string>();
            Version = FormatVersion(Settings.Instance.PackageVersion);
            _futureTypes = new Dictionary<FutureTypeGo, FutureTypeGo>();
        }

        public string Version { get; }

        public string UserAgent => $"Azure-SDK-For-Go/{Version} arm-{Namespace}/{ApiVersion}";

        public string ServiceName => CodeNamerGo.Instance.PascalCase(Namespace ?? string.Empty);

        public string BaseClient => "BaseClient";
        public bool IsCustomBaseUri => Extensions.ContainsKey(SwaggerExtensions.ParameterizedHostExtension);

        public string APIType => (string)Settings.Instance.Host?.GetValue<string>("openapi-type").Result;

        public IEnumerable<string> ClientImports
        {
            get
            {
                var imports = new HashSet<string>();
                imports.UnionWith(CodeNamerGo.Instance.AutorestImports);
                var clientMg = MethodGroups.FirstOrDefault(mg => string.IsNullOrEmpty(mg.Name));
                if (clientMg != null)
                {
                    imports.UnionWith(clientMg.Imports);
                }
                foreach (var p in Properties)
                {
                    p.ModelType.AddImports(imports);
                }
                return imports.OrderBy(i => i);
            }
        }

        public string ClientDocumentation => string.Format("{0} is the base client for {1}.", BaseClient, ServiceName);

        public Dictionary<IModelType, string> PagedTypes { get; }

        /// <summary>
        /// Returns an enumerator to the collection of future types; may be empty.
        /// </summary>
        internal IEnumerable<FutureTypeGo> FutureTypes => _futureTypes.Keys;

        // NextMethodUndefined is used to keep track of those models which are returned by paged methods,
        // but the next method is not defined in the service client, so these models need a preparer.
        public List<IModelType> NextMethodUndefined { get; }

        public IEnumerable<string> ModelImports
        {
            get
            {
                // Create an ordered union of the imports each model requires
                var imports = new HashSet<string>();
                if (ModelTypes != null && ModelTypes.Cast<CompositeTypeGo>().Any(mtm => mtm.IsResponseType || mtm.IsWrapperType))
                {
                    imports.Add(PrimaryTypeGo.GetImportLine("github.com/Azure/go-autorest/autorest"));
                }
                if (_futureTypes.Any())
                {
                    imports.Add(PrimaryTypeGo.GetImportLine("github.com/Azure/go-autorest/autorest/azure"));
                    imports.Add(PrimaryTypeGo.GetImportLine("net/http"));
                }
                ModelTypes.Cast<CompositeTypeGo>()
                    .ForEach(mt =>
                    {
                        mt.AddImports(imports);
                        if (NextMethodUndefined.Any())
                        {
                            imports.UnionWith(CodeNamerGo.Instance.PageableImports);
                        }
                    });
                return imports.OrderBy(i => i);
            }
        }

        public virtual IEnumerable<MethodGroupGo> MethodGroups => Operations.Cast<MethodGroupGo>();

        public bool ShouldValidate => (bool)AutoRest.Core.Settings.Instance.Host?.GetValue<bool?>("client-side-validation").Result;

        public string GlobalParameters
        {
            get
            {
                var declarations = new List<string>();
                foreach (var p in Properties)
                {
                    if (!p.SerializedName.IsApiVersion() && p.DefaultValue.FixedValue.IsNullOrEmpty())
                    {
                        declarations.Add(
                                string.Format(
                                        (p.IsRequired || p.ModelType.CanBeEmpty() ? "{0} {1}" : "{0} *{1}"),
                                         p.Name.Value.ToSentence(), p.ModelType.Name));
                    }
                }
                return string.Join(", ", declarations);
            }
        }

        public string HelperGlobalParameters
        {
            get
            {
                var invocationParams = new List<string>();
                foreach (var p in Properties)
                {
                    if (!p.SerializedName.IsApiVersion() && p.DefaultValue.FixedValue.IsNullOrEmpty())
                    {
                        invocationParams.Add(p.Name.Value.ToSentence());
                    }
                }
                return string.Join(", ", invocationParams);
            }
        }

        public string GlobalDefaultParameters
        {
            get
            {
                var declarations = new List<string>();
                foreach (var p in Properties)
                {
                    if (!p.SerializedName.IsApiVersion() && !p.DefaultValue.FixedValue.IsNullOrEmpty())
                    {
                        declarations.Add(
                                string.Format(
                                        (p.IsRequired || p.ModelType.CanBeEmpty() ? "{0} {1}" : "{0} *{1}"),
                                         p.Name.Value.ToSentence(), p.ModelType.Name.Value.ToSentence()));
                    }
                }
                return string.Join(", ", declarations);
            }
        }

        public string HelperGlobalDefaultParameters
        {
            get
            {
                var invocationParams = new List<string>();
                foreach (var p in Properties)
                {
                    if (!p.SerializedName.IsApiVersion() && !p.DefaultValue.FixedValue.IsNullOrEmpty())
                    {
                        invocationParams.Add("Default" + p.Name.Value);
                    }
                }
                return string.Join(", ", invocationParams);
            }
        }

        public string ConstGlobalDefaultParameters
        {
            get
            {
                var constDeclaration = new List<string>();
                foreach (var p in Properties)
                {
                    if (!p.SerializedName.IsApiVersion() && !p.DefaultValue.FixedValue.IsNullOrEmpty())
                    {
                        constDeclaration.Add(string.Format("// Default{0} is the default value for {1}\nDefault{0} = {2}",
                            p.Name.Value,
                            p.Name.Value.ToPhrase(),
                            p.DefaultValue.Value));
                    }
                }
                return string.Join("\n", constDeclaration);
            }
        }

        public string AllGlobalParameters
        {
            get
            {
                if (GlobalParameters.IsNullOrEmpty())
                {
                    return GlobalDefaultParameters;
                }
                if (GlobalDefaultParameters.IsNullOrEmpty())
                {
                    return GlobalParameters;
                }
                return string.Join(", ", GlobalParameters, GlobalDefaultParameters);
            }
        }

        public string HelperAllGlobalParameters
        {
            get
            {
                if (HelperGlobalParameters.IsNullOrEmpty())
                {
                    return HelperGlobalDefaultParameters;
                }
                if (HelperGlobalDefaultParameters.IsNullOrEmpty())
                {
                    return HelperGlobalParameters;
                }
                return string.Join(", ", HelperGlobalParameters, HelperGlobalDefaultParameters);
            }
        }

        // client methods are the ones with no method group
        public IEnumerable<MethodGo> ClientMethods => Methods.Cast<MethodGo>().Where(m => string.IsNullOrEmpty(m.MethodGroup.Name));

        public override string Namespace
        {
            get => string.IsNullOrEmpty(base.Namespace) ? base.Namespace : base.Namespace.ToLowerInvariant();
            set => base.Namespace = value;
        }

        public string GetDocumentation()
        {
            return $"Package {Namespace} implements the Azure ARM {ServiceName} service API version {ApiVersion}.\n\n{(Documentation ?? string.Empty).UnwrapAnchorTags()}";
        }

        /// FormatVersion normalizes a version string into a SemVer if it resembles one. Otherwise,
        /// it returns the original string unmodified. If version is empty or only comprised of
        /// whitespace, 
        public static string FormatVersion(string version)
        {

            if (string.IsNullOrWhiteSpace(version))
            {
                return "0.0.0";
            }

            var semVerMatch = semVerPattern.Match(version);

            if (!semVerMatch.Success)
            {
                return version;
            }

            var builder = new StringBuilder("v");
            builder.Append(semVerMatch.Groups["major"].Value);
            builder.Append('.');
            builder.Append(semVerMatch.Groups["minor"].Value);
            builder.Append('.');
            builder.Append(semVerMatch.Groups["patch"].Value);
            if (semVerMatch.Groups["tag"].Success)
            {
                builder.Append('-');
                builder.Append(semVerMatch.Groups["tag"].Value);
            }
            return builder.ToString();
        }

        /// <summary>
        /// Creates a future for the specified method and updates its return type.
        /// </summary>
        /// <param name="method">The method to be modified.</param>
        internal void CreateFutureTypeForMethod(MethodGo method)
        {
            if (!method.IsLongRunningOperation())
            {
                throw new InvalidOperationException("CreateFutureTypeForMethod requires method to be a long-running operation");
            }

            // don't create duplicate future types
            var future = new FutureTypeGo(method);
            if (_futureTypes.ContainsKey(future))
            {
                future = _futureTypes[future];
            }
            else
            {
                _futureTypes.Add(future, future);
            }

            method.ReturnType = new Response(future, method.ReturnType.Headers);
        }
    }
}
