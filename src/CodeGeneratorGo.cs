// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core;
using AutoRest.Core.Model;
using AutoRest.Go.Model;
using AutoRest.Go.Templates;
using System;
using System.Linq;
using System.Threading.Tasks;
using System.Collections.Generic;

namespace AutoRest.Go
{
    public class CodeGeneratorGo : CodeGenerator
    {
        private const string ModelsFileName = "models";
        private const string ClientFileName = "client";
        private const string VersionFileName = "version";
        private const string ResponderFileName = "responder_policy";
        private const string ResponseErrorFileName = "response_error";
        private const string ValidationFileName = "validation";

        private HashSet<string> _reservedFiles;

        public CodeGeneratorGo()
        {
            // by convention the methods in the method group with an empty
            // name go into the client template so skip them here.
            _reservedFiles = new HashSet<string>(StringComparer.OrdinalIgnoreCase)
            {
                ModelsFileName,
                ClientFileName,
                VersionFileName,
                ResponderFileName,
                ResponseErrorFileName,
                ValidationFileName
            };
        }

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
            var modelsTemplate = new ModelsTemplate { Model = codeModel };
            await Write(modelsTemplate, FormatFileName(ModelsFileName));

            // Service client
            var serviceClientTemplate = new ServiceClientTemplate { Model = codeModel };
            await Write(serviceClientTemplate, FormatFileName(ClientFileName));

            foreach (var methodGroup in codeModel.MethodGroups.Where(mg => !string.IsNullOrEmpty(mg.Name)))
            {
                if (_reservedFiles.Contains(methodGroup.Name.Value))
                {
                    methodGroup.Name += "group";
                }
                _reservedFiles.Add(methodGroup.Name);
                var methodGroupTemplate = new MethodGroupTemplate { Model = methodGroup };
                await Write(methodGroupTemplate, FormatFileName(methodGroup.Name));
            }

            // Version
            var versionTemplate = new VersionTemplate { Model = codeModel };
            await Write(versionTemplate, FormatFileName(VersionFileName));

            var fixedTemplates = GetFixedTemplates(codeModel);
            foreach (var template in fixedTemplates)
            {
                await Write(template.Item1, FormatFileName(template.Item2));
            }
        }

        /// <summary>
        /// Returns a collection of templates to emit.
        /// </summary>
        /// <param name="codeModel">The code model that the template will consume.</param>
        /// <returns>A collection of templates.</returns>
        private IEnumerable<Tuple<ITemplate, string>> GetFixedTemplates(CodeModelGo codeModel)
        {
            var list = new List<Tuple<ITemplate, string>>
            {
                new Tuple<ITemplate, string>(new ResponderPolicy() { Model = codeModel }, ResponderFileName),
                new Tuple<ITemplate, string>(new ResponseError() { Model = codeModel }, ResponseErrorFileName),
                new Tuple<ITemplate, string>(new Validation() { Model = codeModel }, ValidationFileName)
            };
            return list;
        }

        private string FormatFileName(string fileName)
        {
            var prefix = Settings.Instance.Host?.GetValue<string>("file-prefix").Result;
            // if the prefix is already snaked don't double-snake it
            var prefixSnake = string.Empty;
            if (!string.IsNullOrWhiteSpace(prefix))
            {
                prefix = prefix.ToLowerInvariant();
                if (prefix[prefix.Length - 1] != '_')
                {
                    prefixSnake = "_";
                }
            }
            // convert fileName to snake-case, i.e. "FooBar" becomes "foo_bar"
            fileName = string.Join('_', fileName.ToWords()).ToLowerInvariant();
            return $"{prefix}{prefixSnake}{fileName}{ImplementationFileExtension}";
        }
    }
}
