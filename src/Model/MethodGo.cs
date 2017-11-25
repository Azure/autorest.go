// Copyright (c) Microsoft Open Technologies, Inc. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Go.Properties;
using AutoRest.Core.Utilities;
using AutoRest.Core.Model;
using AutoRest.Extensions;
using AutoRest.Extensions.Azure;
using AutoRest.Extensions.Azure.Model;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using System.Net;
using System.Text;
using Newtonsoft.Json.Linq;

namespace AutoRest.Go.Model
{
    public class MethodGo : Method
    {
        private const string DefaultReturnType = "autorest.Response";

        public string Owner { get; private set; }

        public string PackageName { get; private set; }

        public string APIVersion { get; private set; }

        public bool NextAlreadyDefined { get; private set; }

        public bool IsCustomBaseUri
            => CodeModel.Extensions.ContainsKey(SwaggerExtensions.ParameterizedHostExtension);

        // RegisterRP determines if a DoRetryWithRegistration send decorator will be added to the operation
        // DoRetryWithRegistration retries. Default is generating DoRetryForStatusCodes decorator instead.
        public bool RegisterRP;

        public MethodGo()
        {
            NextAlreadyDefined = true;
        }

        internal void Transform(CodeModelGo cmg)
        {
            Owner = (MethodGroup as MethodGroupGo).ClientName;
            PackageName = cmg.Namespace;
            NextAlreadyDefined = NextMethodExists(cmg.Methods.Cast<MethodGo>());

            var apiVersionParam =
              from p in Parameters
              let name = p.SerializedName
              where name != null && name.IsApiVersion()
              select p.DefaultValue.Value?.Trim(new[] { '"' });

            // When APIVersion is blank, it means that it was unavailable at the method level
            // and we should default back to whatever is present at the client level. However,
            // we will continue embedding that in each method to have broader support.
            APIVersion = apiVersionParam.SingleOrDefault();
            if (APIVersion == default(string))
            {
                APIVersion = cmg.ApiVersion;
            }

            var parameter = Parameters.ToList().Find(p => p.ModelType.PrimaryType(KnownPrimaryType.Stream)
                                                && !(p.Location == ParameterLocation.Body || p.Location == ParameterLocation.FormData));

            if (parameter != null)
            {
                throw new ArgumentException(string.Format(CultureInfo.InvariantCulture,
                    Resources.IllegalStreamingParameter, parameter.Name));
            }
            if (string.IsNullOrEmpty(Description))
            {
                Description = string.Format("sends the {0} request.", Name.ToString().ToPhrase());
            }

            // Registering Azure resource providers should only happen with Azure resource manager REST APIs
            // This depends on go-autorest here:
            // https://github.com/Azure/go-autorest/blob/c0eb859387e57a164bf64171da307e2ef8168b58/autorest/azure/rp.go#L30
            // As registering needs the Azure subscription ID, we take it from the operation path, on the
            // assumption that ARM APIs should include the subscription ID right after `subscriptions`
            RegisterRP = cmg.APIType.EqualsIgnoreCase("arm") && Url.Split("/").Any(p => p.EqualsIgnoreCase("subscriptions"));
        }

        public string MethodSignature => $"{Name}({MethodParametersSignature})";

        public string MethodReturnSignatureComplete
        {
            get
            {
                var signature = new StringBuilder("(<-chan ");
                signature.Append((ListElement.ModelType as SequenceTypeGo).GetElement);
                signature.Append(", <-chan error)");
                return signature.ToString();
            }
        }

        public string ParametersDocumentation
        {
            get
            {
                StringBuilder sb = new StringBuilder();
                foreach (var parameter in LocalParameters)
                {
                    if (!string.IsNullOrEmpty(parameter.Documentation))
                    {
                        sb.Append(parameter.Name);
                        sb.Append(" is ");
                        sb.Append(parameter.Documentation.FixedValue.ToSentence());
                        sb.Append(" ");
                    }
                    if (parameter.ModelType.PrimaryType(KnownPrimaryType.Stream))
                    {
                        sb.Append(parameter.Name);
                        sb.Append(" will be closed upon successful return. Callers should ensure closure when receiving an error.");
                    }
                }
                return sb.ToString();
            }
        }

        public PropertyGo ListElement
        {
            get
            {
                var body = ReturnType.Body as CompositeTypeGo;
                return body.Properties.Where(p => p.ModelType is SequenceTypeGo).FirstOrDefault() as PropertyGo;
            }
        }

        public string ListCompleteMethodName => $"{Name}Complete";

        /// <summary>
        /// Generate the method parameter declaration.
        /// </summary>
        public string MethodParametersSignature
        {
            get
            {
                var declarations = new List<string> {"ctx context.Context"};
                LocalParameters
                    .ForEach(p => declarations.Add(string.Format(
                                                        p.IsRequired || p.ModelType.CanBeEmpty()
                                                            ? "{0} {1}"
                                                            : "{0} *{1}", p.Name, p.ModelType is CompositeTypeGo type && type.IsRootType
                                                                ? p.ModelType.GetInterfaceName()
                                                                : p.ModelType.Name.ToString())));
                return string.Join(", ", declarations);
            }
        }

        /// <summary>
        /// Gets the return type name for this method.
        /// </summary>
        public string MethodReturnType => HasReturnValue() ? ReturnValue().Body.Name.ToString() : DefaultReturnType;

        private string MethodReturnSig(string resultTypeName)
        {
            return $"result {resultTypeName}, err error";
        }

        /// <summary>
        /// Returns the method return signature for this method (e.g. "foo, bar").
        /// For responder methods use ResponderReturnSignature() instead.
        /// </summary>
        /// <returns>The method signature for this method.</returns>
        public string MethodReturnSignature()
        {
            return MethodReturnSig(MethodReturnType);
        }

        /// <summary>
        /// Returns the method return signature for the responder method (e.g. "foo, bar").
        /// </summary>
        /// <returns>The method signature for the responder method.</returns>
        public string ResponderReturnSignature()
        {
            if (IsLongRunningOperation())
            {
                return MethodReturnSig(ReturnType.Body.Cast<FutureTypeGo>().ResultTypeName);
            }
            return MethodReturnSignature();
        }

        public string NextMethodName => $"{Name}NextResults";

        public string PreparerMethodName => $"{Name}Preparer";

        public string SenderMethodName => $"{Name}Sender";

        public string ResponderMethodName => $"{Name}Responder";

        public string HelperInvocationParameters(bool complete)
        {
            var invocationParams = new List<string> {"ctx"};

            foreach (ParameterGo p in LocalParameters)
            {
                if (p.Name.EqualsIgnoreCase("nextlink") && complete)
                {
                    invocationParams.Add(string.Format("*list.{0}", NextLink));
                }
                else
                {
                    invocationParams.Add(p.Name);
                }
            }
            return string.Join(", ", invocationParams);
        }

        /// <summary>
        /// Return the parameters as they appear in the method signature excluding global parameters.
        /// </summary>
        public IEnumerable<ParameterGo> LocalParameters
        {
            get
            {
                return
                    Parameters.Cast<ParameterGo>().Where(
                        p => p != null && p.IsMethodArgument && !string.IsNullOrWhiteSpace(p.Name))
                                .OrderBy(item => !item.IsRequired);
            }
        }

        public IEnumerable<ParameterGo> ParametersGo => Parameters.Cast<ParameterGo>();

        public string ParameterValidations => ParametersGo.Validate(HttpMethod);

        public ParameterGo BodyParameter => ParametersGo.BodyParameter();

        public IEnumerable<ParameterGo> FormDataParameters => ParametersGo.FormDataParameters();

        public IEnumerable<ParameterGo> HeaderParameters => ParametersGo.HeaderParameters();

        public IEnumerable<ParameterGo> OptionalHeaderParameters => ParametersGo.HeaderParameters(false);

        public IEnumerable<ParameterGo> URLParameters => ParametersGo.URLParameters();

        public string URLMap => URLParameters.BuildParameterMap("urlParameters");

        public IEnumerable<ParameterGo> PathParameters => ParametersGo.PathParameters();

        public string PathMap => PathParameters.BuildParameterMap("pathParameters");

        public IEnumerable<ParameterGo> QueryParameters => ParametersGo.QueryParameters();

        public IEnumerable<ParameterGo> OptionalQueryParameters => ParametersGo.QueryParameters(false);

        public string QueryMap => QueryParameters.BuildParameterMap("queryParameters");

        public string FormDataMap => FormDataParameters.BuildParameterMap("formDataParameters");

        public List<string> ResponseCodes
        {
            get
            {
                var codes = new List<string>();
                // Refactor -> CodeModelTransformer
                // Actually, this is the kind of stuff that would be better in the core...
                if (!Responses.ContainsKey(HttpStatusCode.OK))
                {
                    codes.Add(CodeNamerGo.Instance.StatusCodeToGoString[HttpStatusCode.OK]);
                }
                // Refactor -> generator
                foreach (var sc in Responses.Keys)
                {
                    codes.Add(CodeNamerGo.Instance.StatusCodeToGoString[sc]);
                }
                return codes;
            }
        }

        public IEnumerable<string> PrepareDecorators
        {
            get
            {
                var decorators = new List<string>();

                if (BodyParameter != null && !BodyParameter.ModelType.PrimaryType(KnownPrimaryType.Stream))
                {
                    decorators.Add("autorest.AsJSON()");
                }

                decorators.Add(HTTPMethodDecorator);
                if (!this.IsCustomBaseUri)
                {
                    decorators.Add(string.Format("autorest.WithBaseURL(client.BaseURI)"));
                }
                else
                {
                    decorators.Add(string.Format("autorest.WithCustomBaseURL(\"{0}\", urlParameters)", CodeModel.BaseUrl));
                }

                decorators.Add(string.Format(PathParameters.Any()
                            ? "autorest.WithPathParameters(\"{0}\",pathParameters)"
                            : "autorest.WithPath(\"{0}\")",
                        Url));

                if (BodyParameter != null && BodyParameter.IsRequired)
                {
                    decorators.Add(string.Format(BodyParameter.ModelType.PrimaryType(KnownPrimaryType.Stream) && BodyParameter.Location == ParameterLocation.Body
                                        ? "autorest.WithFile({0})"
                                        : "autorest.WithJSON({0})",
                                BodyParameter.Name));
                }

                if (QueryParameters.Any())
                {
                    decorators.Add("autorest.WithQueryParameters(queryParameters)");
                }

                if (FormDataParameters.Any())
                {
                    decorators.Add(
                        FormDataParameters.Any(p => p.ModelType.PrimaryType(KnownPrimaryType.Stream))
                            ? "autorest.WithMultiPartFormData(formDataParameters)"
                            : "autorest.WithFormData(autorest.MapToValues(formDataParameters))"
                        );
                }

                if (HeaderParameters.Any())
                {
                    foreach (var param in Parameters.Where(p => p.IsRequired && p.Location == ParameterLocation.Header))
                    {
                        decorators.Add(param.IsClientProperty
                            ? $"autorest.WithHeader(\"{param.SerializedName}\",client.{param.Name.ToPascalCase()})"
                            : $"autorest.WithHeader(\"{param.SerializedName}\",autorest.String({param.Name}))"
                                );
                    }
                }

                return decorators;
            }
        }

        public string HTTPMethodDecorator
        {
            get
            {
                switch (HttpMethod)
                {
                    case HttpMethod.Delete: return "autorest.AsDelete()";
                    case HttpMethod.Get: return "autorest.AsGet()";
                    case HttpMethod.Head: return "autorest.AsHead()";
                    case HttpMethod.Options: return "autorest.AsOptions()";
                    case HttpMethod.Patch: return "autorest.AsPatch()";
                    case HttpMethod.Post: return "autorest.AsPost()";
                    case HttpMethod.Put: return "autorest.AsPut()";
                    default:
                        throw new ArgumentException(string.Format("The HTTP verb {0} is not supported by the Go SDK", HttpMethod));
                }
            }
        }

        public IEnumerable<string> SendDecorators
        {
            get
            {
                var decorators = new List<string>
                {
                    RegisterRP
                        ? "azure.DoRetryWithRegistration(client.Client)"
                        : "autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...)"
                };
                return decorators;
            }
        }

        public IEnumerable<string> RespondDecorators
        {
            get
            {
                var decorators = new List<string>
                {
                    "resp",
                    "client.ByInspecting()",
                    string.Format("azure.WithErrorUnlessStatusCode({0})", string.Join(",", ResponseCodes.ToArray()))
                };

                if (HasReturnValue() && !ReturnValue().Body.IsStreamType() && !LroWrapsDefaultResp())
                {
                    if (((CompositeTypeGo)ReturnValue().Body).IsWrapperType && !((CompositeTypeGo)ReturnValue().Body).HasPolymorphicFields)
                    {
                        decorators.Add("autorest.ByUnmarshallingJSON(&result.Value)");
                    }
                    else
                    {
                        decorators.Add("autorest.ByUnmarshallingJSON(&result)");
                    }
                }

                if (!HasReturnValue() || !ReturnValue().Body.IsStreamType())
                {
                    decorators.Add("autorest.ByClosing()");
                }
                return decorators;
            }
        }

        /// <summary>
        /// Returns true if the future response wraps a default response type.
        /// We need to make this distinction because the default response doesn't
        /// need to be unmarshalled.
        /// </summary>
        /// <returns>True if the return type is a future that wraps the default response type.</returns>
        private bool LroWrapsDefaultResp()
        {
            if (!IsLongRunningOperation())
            {
                return false;
            }

            if (!HasReturnValue())
            {
                return false;
            }

            var retType = ReturnValue().Body as FutureTypeGo;
            return string.CompareOrdinal(retType.ResultTypeName, DefaultReturnType) == 0;
        }

        /// <summary>
        /// Gets the appropriate response assignment expression; it can be different depending on the method.
        /// </summary>
        /// <param name="forResponder">Specify true if this expression is inside the responder method.</param>
        /// <returns>The response assignment expression.</returns>
        public string Response(bool forResponder)
        {
            if (!HasReturnValue() || (forResponder && LroWrapsDefaultResp()))
            {
                return "result.Response = resp";
            }
            return "result.Response = autorest.Response{Response: resp}";
        }

        public string AutorestError(string phase, string response = null, string parameter = null)
        {
            return !string.IsNullOrEmpty(parameter)
                        ? string.Format("autorest.NewErrorWithError(err, \"{0}.{1}\", \"{2}\", nil , \"{3}\'{4}\'\")", PackageName, Owner, Name, phase, parameter)
                        : string.IsNullOrEmpty(response)
                                 ? string.Format("autorest.NewErrorWithError(err, \"{0}.{1}\", \"{2}\", nil , \"{3}\")", PackageName, Owner, Name, phase)
                                 : string.Format("autorest.NewErrorWithError(err, \"{0}.{1}\", \"{2}\", {3}, \"{4}\")", PackageName, Owner, Name, response, phase);
        }

        public string ValidationError => $"validation.NewErrorWithValidationError(err, \"{PackageName}.{Owner}\",\"{Name}\")";

        /// <summary>
        /// Check if method has a return response.
        /// </summary>
        /// <returns></returns>
        public bool HasReturnValue()
        {
            return ReturnValue()?.Body != null;
        }

        /// <summary>
        /// Return response object for the method.
        /// </summary>
        /// <returns></returns>
        public Response ReturnValue()
        {
            return ReturnType ?? DefaultResponse;
        }

        /// <summary>
        /// Checks if method has pageable extension (x-ms-pageable) enabled.
        /// </summary>
        /// <returns></returns>

        public bool IsPageable => !string.IsNullOrEmpty(NextLink);

        public bool IsNextMethod => Name.Value.EqualsIgnoreCase(NextOperationName);

        /// <summary>
        /// Checks if method for next page of results on paged methods is already present in the method list.
        /// </summary>
        /// <param name="methods"></param>
        /// <returns></returns>
        public bool NextMethodExists(IEnumerable<MethodGo> methods)
        {
            string next = NextOperationName;
            if (string.IsNullOrEmpty(next))
            {
                return false;
            }
            return methods.Any(m => m.Name.Value.EqualsIgnoreCase(next));
        }

        public MethodGo NextMethod
        {
            get
            {
                if (Extensions.ContainsKey(AzureExtensions.PageableExtension))
                {
                    var pageableExtension = JsonConvert.DeserializeObject<PageableExtension>(Extensions[AzureExtensions.PageableExtension].ToString());
                    if (pageableExtension != null && !string.IsNullOrWhiteSpace(pageableExtension.OperationName))
                    {
                        return (CodeModel.Methods.First(m => m.SerializedName.EqualsIgnoreCase(pageableExtension.OperationName)) as MethodGo);
                    }
                }
                return null;
            }
        }

        public string NextOperationName => NextMethod?.Name.Value;

        /// <summary>
        /// Check if method has long running extension (x-ms-long-running-operation) enabled.
        /// </summary>
        /// <returns></returns>
        public bool IsLongRunningOperation()
        {
            try
            {
                return Extensions.ContainsKey(AzureExtensions.LongRunningExtension) && (bool)Extensions[AzureExtensions.LongRunningExtension];
            }
            catch (InvalidCastException e)
            {
                var message = $@"{
                    e.Message
                    } The value \'{
                    Extensions[AzureExtensions.LongRunningExtension]
                    }\' for extension {
                    AzureExtensions.LongRunningExtension
                    } for method {
                    Group
                    }. {
                    Name
                    } is invalid in Swagger. It should be boolean.";

                throw new InvalidOperationException(message);
            }
        }

        /// <summary>
        /// Add NextLink attribute for pageable extension for the method.
        /// </summary>
        /// <returns></returns>
        public string NextLink
        {
            get
            {
                // Note:
                // Methods can be paged, even if "nextLinkName" is null
                // Paged method just means a method returns an array
                if (!Extensions.ContainsKey(AzureExtensions.PageableExtension))
                {
                    return null;
                }
                if (!(Extensions[AzureExtensions.PageableExtension] is JContainer pageableExtension))
                {
                    return null;
                }

                var nextLink = (string)pageableExtension["nextLinkName"];
                if (string.IsNullOrEmpty(nextLink))
                {
                    return null;

                }
                return CodeNamerGo.Instance.GetPropertyName(nextLink);
            }
        }
    }
}
