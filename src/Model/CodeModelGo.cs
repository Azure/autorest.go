// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using AutoRest.Core;
using AutoRest.Core.Model;
using AutoRest.Core.Utilities;
using AutoRest.Extensions;
using static AutoRest.Core.Utilities.DependencyInjection;

namespace AutoRest.Go.Model
{
    public class CodeModelGo : CodeModel
    {

        private static readonly Regex semVerPattern = new Regex(@"^v?(?<major>\d+)\.(?<minor>\d+)\.(?<patch>\d+)(?:-(?<tag>\S+))?$", RegexOptions.Compiled);
        public string Version { get; }
        public string UserAgent
        {
            get
            {
                return $"Azure-SDK-For-Go/{Version} arm-{Namespace}/{ApiVersion}";
            }
        }

        public CodeModelGo()
        {
            NextMethodUndefined = new List<IModelType>();
            PagedTypes = new Dictionary<IModelType, string>();
            Version = FormatVersion(Settings.Instance.PackageVersion);
        }

        public override string Namespace
        {
            get
            {
                if (string.IsNullOrEmpty(base.Namespace))
                {
                    return base.Namespace;
                }

                return base.Namespace.ToLowerInvariant();
            }
            set
            {
                base.Namespace = value;
            }
        }

        public string ServiceName => CodeNamer.Instance.PascalCase(Namespace ?? string.Empty);

        public string GetDocumentation()
        {
            return $"Package {Namespace} implements the Azure ARM {ServiceName} service API version {ApiVersion}.\n\n{(Documentation ?? string.Empty).UnwrapAnchorTags()}";
        }

        public string BaseClient => CodeNamerGo.Instance.ExportClientTypes ? "ManagementClient" : "managementClient";

        public bool IsCustomBaseUri => Extensions.ContainsKey(SwaggerExtensions.ParameterizedHostExtension);

        public IEnumerable<string> ClientImports
        {
            get
            {
                var imports = new HashSet<string>();
                imports.UnionWith(CodeNamerGo.Instance.AutorestImports);
                var clientMg = MethodGroups.Where(mg => string.IsNullOrEmpty(mg.Name)).FirstOrDefault();
                if (clientMg != null)
                {
                    imports.UnionWith(clientMg.Imports);
                }
                return imports.OrderBy(i => i);
            }
        }

        public string ClientDocumentation => string.Format("{0} is the base client for {1}.", BaseClient, ServiceName);

        public Dictionary<IModelType, string> PagedTypes { get; }

        // NextMethodUndefined is used to keep track of those models which are returned by paged methods,
        // but the next method is not defined in the service client, so these models need a preparer.
        public List<IModelType> NextMethodUndefined { get; }

        public IEnumerable<string> ModelImports
        {
            get
            {
                var addIoImport = false;
                var addStrConvImport = false;
                var addStringsImport = false;
                // Create an ordered union of the imports each model requires
                var imports = new HashSet<string>();
                var isV1Template = TemplateFactory.Instance.TemplateVersion == TemplateFactory.Version.v1;
                if (!isV1Template)
                {
                    imports.Add(PrimaryTypeGo.GetImportLine(package: "net/http"));
                }
                else
                {
                    if (ModelTypes != null && ModelTypes.Cast<CompositeTypeGo>().Any(mtm => mtm.IsResponseType || mtm.IsWrapperType))
                    {
                        imports.Add(PrimaryTypeGo.GetImportLine("github.com/Azure/go-autorest/autorest"));
                    }
                }

                ModelTypes.Cast<CompositeTypeGo>()
                    .ForEach(mt =>
                    {
                        mt.AddImports(imports);
                        if (NextMethodUndefined.Any() && isV1Template)
                        {
                            imports.UnionWith(CodeNamerGo.Instance.PageableImports);
                        }
                        if (!isV1Template && mt.IsResponseType)
                        {
                            foreach (var p in mt.Properties)
                            {
                                if (p.ModelType.IsPrimaryType(KnownPrimaryType.Int) || p.ModelType.IsPrimaryType(KnownPrimaryType.Long))
                                {
                                    addStrConvImport = true;
                                    break;
                                }
                            }
                        }
                        if (mt.ResponseIncludesMetadata)
                        {
                            addStringsImport = true;
                        }
                        if (!isV1Template && mt.IsResponseType && mt.IsStreamType())
                        {
                            addIoImport = true;
                        }
                    });
                if (addIoImport)
                {
                    imports.Add(PrimaryTypeGo.GetImportLine(package: "io"));
                }
                if (addStrConvImport)
                {
                    imports.Add(PrimaryTypeGo.GetImportLine(package: "strconv"));
                }
                if (addStringsImport)
                {
                    imports.Add(PrimaryTypeGo.GetImportLine(package: "strings"));
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
                                        (p.IsRequired || p.ModelType.CanBeNull() ? "{0} {1}" : "{0} *{1}"),
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
                                        (p.IsRequired || p.ModelType.CanBeNull() ? "{0} {1}" : "{0} *{1}"),
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
                return string.Join(", ", new string[] { GlobalParameters, GlobalDefaultParameters });
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
                return string.Join(", ", new string[] { HelperGlobalParameters, HelperGlobalDefaultParameters });
            }
        }

        public IEnumerable<MethodGo> ClientMethods
        {
            get
            {
                // client methods are the ones with no method group
                return Methods.Cast<MethodGo>().Where(m => string.IsNullOrEmpty(m.MethodGroup.Name));
            }
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

            if (semVerMatch.Success)
            {
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

            return version;
        }

        /// <summary>
        /// Returns true if any model types contain a metadata property.
        /// </summary>
        public bool UsesMetadataType
        {
            get
            {
                return ModelTypes.Where(m => m.Properties.Cast<PropertyGo>().Where(p => p.IsMetadata).Any()).Any();
            }
        }

        /// <summary>
        /// Returns true if any model types contain an Etag property.
        /// </summary>
        public bool UsesETags
        {
            get
            {
                return ModelTypes.Where(m => m.Properties.Where(p => p.ModelType.IsETagType()).Any()).Any();
            }
        }

        /// <summary>
        /// Returns a collection of composite types that require custom marshalling and/or
        /// unmarshalling. Can be empty if there are no types requriring marshallers.
        /// </summary>
        public IEnumerable<CompositeTypeGo> RequiresMarshallers
        {
            get
            {
                return ModelTypes.Cast<CompositeTypeGo>().Where(m => m.Properties.Where(p => p.ModelType.IsDateTimeType()).Any());
            }
        }

        /// <summary>
        /// Returns the encoding type used for serialization (e.g. xml or json).
        /// </summary>
        public string Encoding => ShouldGenerateXmlSerialization ? "xml" : "json";

        /// <summary>
        /// Gets the collection of enum types sorted by name.
        /// </summary>
        public IEnumerable<EnumTypeGo> Enums
        {
            get
            {
                return EnumTypes.Cast<EnumTypeGo>().OrderBy(e => e.Name.FixedValue);
            }
        }

        /// <summary>
        /// Gets the collection of model types sorted by name.
        /// </summary>
        public IEnumerable<CompositeTypeGo> Models
        {
            get
            {
                return ModelTypes.Cast<CompositeTypeGo>().OrderBy(m => m.Name.Value);
            }
        }
    }
}
