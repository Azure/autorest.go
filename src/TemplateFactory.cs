// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core;
using AutoRest.Go.Model;
using System;
using System.Collections.Generic;
using System.Linq;

namespace AutoRest.Go
{
    /// <summary>
    /// Provides factory methods for obtaining templates based on user-supplied version.
    /// </summary>
    class TemplateFactory
    {
        public enum Version
        {
            v1,
            v2
        }

        private static TemplateFactory s_instance;

        /// <summary>
        /// Returns an immutable template factory.
        /// </summary>
        public static TemplateFactory Instance
        {
            get
            {
                if (s_instance == null)
                {
                    s_instance = new TemplateFactory();
                }
                return s_instance;
            }
        }

        public Version TemplateVersion { get; private set; }

        private TemplateFactory()
        {
            // the go-template-ver param is optional so if it wasn't provided default to v1
            var version = Version.v1;
            var ver = Settings.Instance.Host?.GetValue<string>("go-template-ver").Result;
            if (!string.IsNullOrWhiteSpace(ver) && !Enum.TryParse(ver, out version))
            {
                throw new ArgumentException($"bad go template version '{ver}'");
            }
            TemplateVersion = version;
        }

        /// <summary>
        /// Returns a method group template that uses the provided method group model.
        /// </summary>
        /// <param name="methodGroup">The method group model that the template will consume.</param>
        /// <returns>A new instance of a method group template based on the current template version.</returns>
        public ITemplate GetMethodGroupTemplate(MethodGroupGo methodGroup)
        {
            if (methodGroup == null)
            {
                throw new ArgumentNullException(nameof(methodGroup));
            }
            switch (TemplateVersion)
            {
                case Version.v1:
                    return new Templates.v1.MethodGroupTemplate() { Model = methodGroup };
                case Version.v2:
                    return new Templates.v2.MethodGroupTemplate() { Model = methodGroup };
                default:
                    throw new InvalidOperationException($"");
            }
        }

        /// <summary>
        /// Returns a models group template that uses the provided code model.
        /// </summary>
        /// <param name="codeModel">The code model that the template will consume.</param>
        /// <returns>A new instance of a models template based on the current template version.</returns>
        public ITemplate GetModelsTemplate(CodeModelGo codeModel)
        {
            if (codeModel == null)
            {
                throw new ArgumentNullException(nameof(codeModel));
            }
            switch (TemplateVersion)
            {
                case Version.v1:
                    return new Templates.v1.ModelsTemplate() { Model = codeModel };
                case Version.v2:
                    return new Templates.v2.ModelsTemplate() { Model = codeModel };
                default:
                    throw new InvalidOperationException($"");
            }
        }

        /// <summary>
        /// Returns a service client template that uses the provided code model.
        /// </summary>
        /// <param name="codeModel">The code model that the template will consume.</param>
        /// <returns>A new instance of a service client template based on the current template verion.</returns>
        public ITemplate GetServiceClientTemplate(CodeModelGo codeModel)
        {
            if (codeModel == null)
            {
                throw new ArgumentNullException(nameof(codeModel));
            }
            switch (TemplateVersion)
            {
                case Version.v1:
                    return new Templates.v1.ServiceClientTemplate() { Model = codeModel };
                case Version.v2:
                    return new Templates.v2.ServiceClientTemplate() { Model = codeModel };
                default:
                    throw new InvalidOperationException($"");
            }
        }

        /// <summary>
        /// Returns a version template that uses the provided code model.
        /// </summary>
        /// <param name="codeModel">The code moel that the template will consume.</param>
        /// <returns>A new instance of a version template based on the current template version.</returns>
        public ITemplate GetVersionTemplate(CodeModelGo codeModel)
        {
            if (codeModel == null)
            {
                throw new ArgumentNullException(nameof(codeModel));
            }
            switch (TemplateVersion)
            {
                case Version.v1:
                    return new Templates.v1.VersionTemplate() { Model = codeModel };
                case Version.v2:
                    return new Templates.v2.VersionTemplate() { Model = codeModel };
                default:
                    throw new InvalidOperationException($"");
            }
        }

        /// <summary>
        /// Returns a collection of templates to emit.
        /// </summary>
        /// <param name="codeModel">The code moel that the template will consume.</param>
        /// <returns>A collection of templates, may be empty.</returns>
        public IEnumerable<Tuple<ITemplate, string>> GetFixedTemplates(CodeModelGo codeModel)
        {
            var list = new List<Tuple<ITemplate, string>>();
            if (TemplateVersion == Version.v2)
            {
                list.Add(new Tuple<ITemplate, string>(new Templates.v2.ResponderPolicy() { Model = codeModel }, "responder_policy"));
                list.Add(new Tuple<ITemplate, string>(new Templates.v2.ResponseError() { Model = codeModel }, "response_error"));
                list.Add(new Tuple<ITemplate, string>(new Templates.v2.Validation() { Model = codeModel }, "validation"));
            }
            return list;
        }

        /// <summary>
        /// Returns a marshalling template that uses the provided code model.
        /// </summary>
        /// <param name="codeModel">The code moel that the template will consume.</param>
        /// <returns>
        /// A new instance of a marshalling template based on the current
        /// template version or null if no marshalling template is required.
        /// </returns>
        public ITemplate GetMarshallingTemplate(CodeModelGo codeModel)
        {
            if (TemplateVersion == Version.v1 || !codeModel.RequiresMarshallers.Any())
            {
                return null;
            }
            return new Templates.v2.Marshalling() { Model = codeModel };
        }
    }
}
