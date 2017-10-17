// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core;
using AutoRest.Core.Model;
using AutoRest.Go.Model;
using System;
using System.Linq;
using System.Threading.Tasks;
using System.Collections.Generic;

namespace AutoRest.Go
{
    public class CodeGeneratorGo : CodeGenerator
    {
        public string Name
        {
            get { return "Go"; }
        }

        public override string UsageInstructions
        {
            get { return string.Empty; }
        }

        public override string ImplementationFileExtension
        {
            get { return ".go"; }
        }

        /// <summary>
        /// Generates Go code for service client.
        /// </summary>
        /// <param name="serviceClient"></param>
        /// <returns></returns>
        public override async Task Generate(CodeModel cm)
        {
            var codeModel = cm as CodeModelGo;
            if (codeModel == null)
            {
                throw new Exception("Code model is not a Go Code Model");
            }

            // unfortunately there is an ordering issue here.  during model generation we might
            // flatten some types (but not all depending on type).  then during client generation
            // the validation codegen needs to know if a type was flattened so it can generate
            // the correct code, so we need to generate models before clients.

            // Models
            var modelsTemplate = TemplateFactory.Instance.GetModelsTemplate(codeModel);
            await Write(modelsTemplate, FormatFileName("models"));

            // Service client
            var serviceClientTemplate = TemplateFactory.Instance.GetServiceClientTemplate(codeModel);
            await Write(serviceClientTemplate, FormatFileName("client"));

            // by convention the methods in the method group with an empty
            // name go into the client template so skip them here.
            HashSet<string> ReservedFiles = new HashSet<string>(StringComparer.OrdinalIgnoreCase)
            {
                "models",
                "client",
                "version",
            };

            foreach (var methodGroup in codeModel.MethodGroups.Where(mg => !string.IsNullOrEmpty(mg.Name)))
            {
                if (ReservedFiles.Contains(methodGroup.Name.Value))
                {
                    methodGroup.Name += "group";
                }
                ReservedFiles.Add(methodGroup.Name);
                var methodGroupTemplate = TemplateFactory.Instance.GetMethodGroupTemplate(methodGroup);
                await Write(methodGroupTemplate, FormatFileName(methodGroup.Name));
            }

            // Version
            var versionTemplate = TemplateFactory.Instance.GetVersionTemplate(codeModel);
            await Write(versionTemplate, FormatFileName("version"));

            var fixedTemplates = TemplateFactory.Instance.GetFixedTemplates(codeModel);
            foreach (var template in fixedTemplates)
            {
                await Write(template.Item1, FormatFileName(template.Item2));
            }

            var marshallingTemplate = TemplateFactory.Instance.GetMarshallingTemplate(codeModel);
            if (marshallingTemplate != null)
            {
                await Write(marshallingTemplate, FormatFileName("marshalling"));
            }
        }

        private string FormatFileName(string fileName)
        {
            var prefix = Settings.Instance.Host?.GetValue<string>("file-prefix").Result;
            // if the prefix is already snaked don't double-snake it
            var prefixSnake = string.Empty;
            if (!string.IsNullOrWhiteSpace(prefix) && prefix[prefix.Length - 1] != '_')
            {
                prefixSnake = "_";
            }
            // convert fileName to snake-case, i.e. "FooBar" becomes "foo_bar"
            fileName = string.Join('_', fileName.ToWords()).ToLowerInvariant();
            return $"{prefix}{prefixSnake}{fileName}{ImplementationFileExtension}";
        }
    }
}
