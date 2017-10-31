// Copyright (c) Microsoft Open Technologies, Inc. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using System;
using System.Collections.Generic;
using System.Linq;
using AutoRest.Core.Model;
using AutoRest.Core.Utilities;
using AutoRest.Go;
using AutoRest.Extensions;

namespace AutoRest.Go.Model
{
    public class MethodGroupGo : MethodGroup
    {
        public string ClientName { get; private set; }
        public string Documentation { get; private set; }
        public string PackageName { get; private set; }
        public string BaseClient { get; private set; }

        public bool IsCustomBaseUri
            => CodeModel.Extensions.ContainsKey(SwaggerExtensions.ParameterizedHostExtension);

        public string GlobalParameters;
        public string HelperGlobalParameters;

        public IEnumerable<string> Imports { get; private set; }

        public MethodGroupGo(string name) : base(name)
        {
        }

        public MethodGroupGo()
        {
        }

        internal void Transform(CodeModelGo cmg)
        {
            var originalName = Name.Value;
            Name = Name.Value.TrimPackageName(cmg.Namespace);
            if (Name != originalName)
            {
                // fix up the method group names
                cmg.Methods.Where(m => m.Group.Value == originalName)
                    .ForEach(m =>
                    {
                        m.Group = Name;
                    });
            }

            ClientName = string.IsNullOrEmpty(Name)
                            ? cmg.BaseClient
                            : TypeName.Value.IsNamePlural(cmg.Namespace)
                                             ? Name + "Client"
                                             : (Name + "Client").TrimPackageName(cmg.Namespace);

            if (!CodeNamerGo.Instance.ExportClientTypes)
            {
                ClientName = ClientName.ToCamelCase();
            }

            Documentation = string.Format("{0} is the {1} ", ClientName,
                                    string.IsNullOrEmpty(cmg.Documentation)
                                        ? string.Format("client for the {0} methods of the {1} service.", TypeName, cmg.ServiceName)
                                        : cmg.Documentation.ToSentence());

            PackageName = cmg.Namespace;
            BaseClient = cmg.BaseClient;
            GlobalParameters = cmg.GlobalParameters;
            HelperGlobalParameters = cmg.HelperGlobalParameters;

            //Imports
            var imports = new HashSet<string>();
            imports.UnionWith(CodeNamerGo.Instance.PipelineImports);
            imports.UnionWith(CodeNamerGo.Instance.StandardImports);

            bool marshalImports = false;
            bool unmarshalImports = false;
            cmg.Methods.Where(m => m.Group.Value == Name)
                .ForEach(m =>
                {
                    var mg = m as MethodGo;
                    foreach (var param in mg.ParametersGo)
                    {
                        param.AddImports(imports);
                    }
                    if (mg.HasReturnValue() && !mg.ReturnValue().Body.IsPrimaryType(KnownPrimaryType.Stream))
                    {
                        mg.ReturnType.Body.AddImports(imports);
                    }
                    if (mg.ReturnValueRequiresUnmarshalling())
                    {
                        unmarshalImports = true;
                    }
                    if (mg.BodyParamNeedsMarshalling())
                    {
                        marshalImports = true;
                    }
                });

            if (marshalImports || unmarshalImports)
            {
                // used by preparers and responders
                var encoding = CodeModel.ShouldGenerateXmlSerialization ? "xml" : "json";
                imports.Add(PrimaryTypeGo.GetImportLine(package: $"encoding/{encoding}"));
            }

            if (unmarshalImports)
            {
                // needed by the responder to read the response body
                imports.Add(PrimaryTypeGo.GetImportLine(package: "io/ioutil"));
            }

            if (marshalImports)
            {
                // needed by the preparer to wrap the request body
                imports.Add(PrimaryTypeGo.GetImportLine(package: "bytes"));
            }

            foreach (var p in cmg.Properties)
            {
                p.ModelType.AddImports(imports);
            }

            imports.OrderBy(i => i);
            Imports = imports;
        }

        /// <summary>
        /// Gets the collection of methods in this group sorted by name.
        /// </summary>
        public IEnumerable<MethodGo> MethodsGo
        {
            get
            {
                return Methods.Cast<MethodGo>().OrderBy(m => m.Name.Value);
            }
        }
    }
}
