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
        public static readonly string OneVerString = "version.Number";
        private static readonly Regex semVerPattern = new Regex(@"^v?(?<major>\d+)\.(?<minor>\d+)\.(?<patch>\d+)(?:-(?<tag>\S+))?$", RegexOptions.Compiled);

        public CodeModelGo()
        {
            Version = FormatVersion(Settings.Instance.PackageVersion);
            SpecifiedUserAgent = Settings.Instance.Host?.GetValue<string>("user-agent").Result;
        }

        public string Version { get; }

        public string UserAgent
        {
            get => SpecifiedUserAgent ?? DefaultUserAgent;
            set => SpecifiedUserAgent = value;
        }

        /// <summary>
        /// Returns true if the --use-onever flag was specified (off by default).
        /// </summary>
        public bool UseOneVer => Settings.Instance.Host?.GetValue<bool>("use-onever").Result ?? false;

        /// <summary>
        /// Returns the value passed with the --tag option or null if not specified.
        /// </summary>
        public string Tag => Settings.Instance.Host?.GetValue<string>("tag").Result ?? null;

        /// <summary>
        /// Returns the name of the packages version directory, e.g. "2018-02-01", calculated
        /// from the value of the output-folder.  Returns null if output-folder wasn't specified.
        /// </summary>
        private string PackageVerDir()
        {
            var outDir = Settings.Instance.Host?.GetValue<string>("output-folder").Result;
            if (string.IsNullOrWhiteSpace(outDir))
            {
                return null;
            }

            // the output-folder is defined in the config file like this:
            //
            //   output-folder: $(go-sdk-folder)/services/monitor/mgmt/2017-05-01-preview/insights
            //
            // we want the "2017-05-01-preview" portion

            var i = outDir.LastIndexOf('/');
            var j = outDir.LastIndexOf('/', i - 1);
            return outDir.Substring(j + 1, i - j - 1);
        }

        private string DefaultUserAgent
        {
            get
            {
                var verStr = UseOneVer ? $"\" + {OneVerString} + \"" : Version;
                var suffix = "";

                // the API version will not be populated for composite packages that span
                // multiple swaggers.  in that case first try to get the version info from
                // the output directory of the package.  if that fails then try the tag,
                // and if that fails just include the package name.
                if (!string.IsNullOrWhiteSpace(ApiVersion))
                {
                    suffix = $"/{ApiVersion}";
                }
                else if (!string.IsNullOrWhiteSpace(PackageVerDir()))
                {
                    suffix = $"/{PackageVerDir()}";
                }
                else if (!string.IsNullOrWhiteSpace(Tag))
                {
                    suffix = $"/{Tag}";
                }

                return $"Azure-SDK-For-Go/{verStr} {Namespace}{suffix}";
            }
        }

        private string SpecifiedUserAgent
        {
            get;
            set;
        }

        public string ServiceName => CodeNamerGo.Instance.PascalCase(Namespace ?? string.Empty);

        public string BaseClient => "BaseClient";
        public bool IsCustomBaseUri => Extensions.ContainsKey(SwaggerExtensions.ParameterizedHostExtension);

        public string APIType => Settings.Instance.Host?.GetValue<string>("openapi-type").Result;

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
                var futureModels = ModelTypes.Where(mt => mt is FutureTypeGo).Cast<FutureTypeGo>();
                if (futureModels.Any())
                {
                    imports.Add(PrimaryTypeGo.GetImportLine("github.com/Azure/go-autorest/autorest/azure"));
                    // if any of the futures return the non-default response then we need the net/http package
                    if (futureModels.Any(fm => !fm.IsDefaultReturnType))
                    {
                        imports.Add(PrimaryTypeGo.GetImportLine("net/http"));
                    }
                }
                ModelTypes.Cast<CompositeTypeGo>()
                    .ForEach(mt =>
                    {
                        mt.AddImports(imports);
                    });
                // if any paged types need a preparer created add the pageable imports
                if (ModelTypes.Any(mt => mt is PageTypeGo && mt.Cast<PageTypeGo>().PreparerNeeded))
                {
                    imports.UnionWith(CodeNamerGo.Instance.PageableImports);
                }
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
                if (!IsCustomBaseUri)
                {
                    constDeclaration.Add($"// DefaultBaseURI is the default URI used for the service {ServiceName}\nDefaultBaseURI = \"{BaseUrl}\"");
                }
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

        public string GetDocumentation => $"Package {Namespace} implements the Azure ARM {ServiceName} service API version {ApiVersion}.\n\n{(base.Documentation ?? string.Empty).UnwrapAnchorTags()}";

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
        /// Creates a pageable type for the specified method and updates its return type.
        /// </summary>
        /// <param name="method">The method to be modified.</param>
        internal void CreatePageableTypeForMethod(MethodGo method)
        {
            if (!method.IsPageable)
            {
                throw new InvalidOperationException("CreatePageableTypeForMethod requires method to be a pageable operation");
            }

            var page = new PageTypeGo(method);
            if (ModelTypes.Contains(page))
            {
                page = ModelTypes.First(mt => mt.Equals(page)).Cast<PageTypeGo>();
            }
            else
            {
                Add(page);
                Add(page.IteratorType);
            }

            method.ReturnType = new Response(page, method.ReturnType.Headers);
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

            // this is the future to return from the method
            var future = GetOrAddFuture(new FutureTypeGo(method, method.ReturnValue().Body));

            // if this is a pageable method create a future type for the
            // "list all" method wrapped in our custom response type
            if (method.IsPageable)
            {
                var listAllFuture = GetOrAddFuture(new FutureTypeGo(CodeNamerGo.Instance.GetFutureTypeName($"{method.Group}{method.Name}All"), method,
                    ((PageTypeGo)method.ReturnValue().Body).IteratorType));
                method.ReturnType = new LroPagedResponseGo(future, listAllFuture, method.ReturnType.Headers);
            }
            else
            {
                method.ReturnType = new Response(future, method.ReturnType.Headers);
            }
        }

        /// <summary>
        /// Checks if the specified future type already exists, if it does return that one instead.
        /// If it does not exist it is added to the collection of model types and returned.
        /// </summary>
        /// <param name="futureType">The future type to check for and possibly add.</param>
        /// <returns>The existing or added object.</returns>
        private FutureTypeGo GetOrAddFuture(FutureTypeGo futureType)
        {
            // don't create duplicate future types
            if (ModelTypes.Contains(futureType))
            {
                futureType = ModelTypes.First(mt => mt.Equals(futureType)).Cast<FutureTypeGo>();
            }
            else
            {
                Add(futureType);
            }
            return futureType;
        }
    }
}
