// Copyright (c) Microsoft Open Technologies, Inc. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using System.Collections.Generic;
using System.Linq;
using AutoRest.Core.Model;
using AutoRest.Core.Utilities;
using AutoRest.Extensions;

namespace AutoRest.Go.Model
{
    public class MethodGroupGo : MethodGroup
    {
        private static HashSet<string> s_AllNames = new HashSet<string>();

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

            // use the API version as the prefix to the group name; this is because
            // in batch mode the generator isn't recycled per batch so you can end up
            // with erroneous collisions.  note that for composite packages (web)
            // there will be no API version so fall back to using the batch tag.
            var prefix = cmg.ApiVersion;
            if (string.IsNullOrWhiteSpace(prefix))
            {
                prefix = cmg.Tag;
            }

            // keep a list of all method group names as trimming the package name
            // can introduce collisions.  if there's a collision append "Group" to
            // the name.  unfortunately we can't do this in the namer as we don't
            // have access to the package name.
            if (s_AllNames.Contains($"{prefix}_{Name.Value}"))
            {
                Name += "Group";
            }
            s_AllNames.Add($"{prefix}_{Name.Value}");

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
            imports.UnionWith(CodeNamerGo.Instance.AutorestImports);
            imports.UnionWith(CodeNamerGo.Instance.StandardImports);
            imports.UnionWith(CodeNamerGo.Instance.TracingImports);


            cmg.Methods.Where(m => m.Group.Value == Name)
                .ForEach(m =>
                {
                    var mg = m as MethodGo;
                    if ((CodeModel as CodeModelGo).ShouldValidate && !mg.ParameterValidations.IsNullOrEmpty())
                    {
                        imports.UnionWith(CodeNamerGo.Instance.ValidationImports);
                    }
                    mg.ParametersGo.ForEach(p => p.AddImports(imports));
                    if (mg.HasReturnValue() && !mg.ReturnValue().Body.PrimaryType(KnownPrimaryType.Stream))
                    {
                        mg.ReturnType.Body.AddImports(imports);
                    }
                    if (mg.IsNextMethod)
                    {
                        imports.UnionWith(CodeNamerGo.Instance.PageableImports);
                    }
                });

            foreach (var p in cmg.Properties)
            {
                p.ModelType.AddImports(imports);
            }

            imports.OrderBy(i => i);
            Imports = imports;
        }
    }
}
