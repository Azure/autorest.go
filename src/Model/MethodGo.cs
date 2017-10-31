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

namespace AutoRest.Go.Model
{
    public class MethodGo : Method
    {
        private const string DefaultResponseType = "http.Response";

        public new MethodGroupGo MethodGroup { get { return (MethodGroupGo)base.MethodGroup; } }

        public string PackageName { get; private set; }

        public string APIVersion { get; private set; }

        private readonly string lroDescription = " This method may poll for completion. Polling can be canceled by passing the cancel channel argument. " +
                                                 "The channel will be used to cancel polling and any outstanding HTTP requests.";

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

            var parameter = Parameters.ToList().Find(p => p.ModelType.IsPrimaryType(KnownPrimaryType.Stream)
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

            if (IsLongRunningOperation())
            {
                Description += lroDescription;
            }

            // Registering Azure resource providers should only happen with Azure resource manager REST APIs
            // This depends on go-autorest here:
            // https://github.com/Azure/go-autorest/blob/c0eb859387e57a164bf64171da307e2ef8168b58/autorest/azure/rp.go#L30
            // As registering needs the Azure subscription ID, we take it from the operation path, on the
            // assumption that ARM APIs should include the subscription ID right after `subscriptions`
            RegisterRP = cmg.APIType.EqualsIgnoreCase("arm") && Url.Split("/").Any(p => p.EqualsIgnoreCase("subscriptions"));
        }

        public string MethodSignature => $"{Name}({MethodParametersSignature(false)})";
        
        public string MethodParametersSignatureComplete
        {
            get
            {     
                var signature = new StringBuilder("(");
                signature.Append(MethodParametersSignature(false));
                if (!IsLongRunningOperation())
                {
                    if (MethodParametersSignature(false).Length > 0)
                    {
                        signature.Append( ", ");
                    }
                    signature.Append("cancel <-chan struct{}");
                }
                signature.Append(")");
                return signature.ToString();
            }
        }

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
                    if (parameter.ModelType.IsPrimaryType(KnownPrimaryType.Stream))
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
        public string MethodParametersSignature(bool includeCtx)
        {
            List<string> declarations = new List<string>();

            if (includeCtx)
            {
                // add context as first param
                declarations.Add("ctx context.Context");
            }

            foreach (var localParam in LocalParameters)
            {
                if (localParam.ReplaceStreamWithReadSeeker)
                {
                    // body parameters are passed as read seekers
                    declarations.Add("body io.ReadSeeker");
                    continue;
                }
                declarations.Add(string.Format(localParam.IsPassedByValue()
                                                        ? "{0} {1}"
                                                        : "{0} *{1}", localParam.Name, localParam.ModelType.Name));
            }

            return string.Join(", ", declarations);
        }

        /// <summary>
        /// Gets the return type name for this method.
        /// </summary>
        public string MethodReturnType
        {
            get
            {
                var rv = ReturnValue();
                if (rv.Body != null)
                {
                    return rv.Body.Name.ToString();
                }
                else if (rv.Headers != null)
                {
                    return rv.Headers.Name.ToString();
                }
                return DefaultResponseType;
            }
        }

        /// <summary>
        /// Returns true if the method return type is the default response.
        /// </summary>
        public bool IsDefaultResponseType
        {
            get
            {
                return MethodReturnType == DefaultResponseType;
            }
        }

        /// <summary>
        /// Returns the method return signature for this method (e.g. "foo, bar").
        /// </summary>
        /// <param name="helper">Indicates if this method is a helper method (i.e. preparer/sender/responder).</param>
        /// <returns>The method signature for this method.</returns>
        public string MethodReturnSignature(bool helper)
        {
            return $"*{MethodReturnType}, error";
        }

        public string NextMethodName => $"{Name}NextResults";

        public string HelperInvocationParameters(bool complete, bool includeCtx)
        {
            List<string> invocationParams = new List<string>();
            if (includeCtx)
            {
                invocationParams.Add("ctx");
            }
            foreach (ParameterGo p in LocalParameters)
            {
                if (p.ReplaceStreamWithReadSeeker)
                {
                    invocationParams.Add("body");
                }
                else if (p.Name.EqualsIgnoreCase("nextlink") && complete)
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
                        p => p != null && p.IsMethodArgument && !string.IsNullOrWhiteSpace(p.Name) &&
                        !(p.Location == ParameterLocation.Query && p.IsConstant))
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

                if (BodyParameter != null && !BodyParameter.ModelType.IsPrimaryType(KnownPrimaryType.Stream))
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
                    decorators.Add(string.Format(BodyParameter.ModelType.IsPrimaryType(KnownPrimaryType.Stream) && BodyParameter.Location == ParameterLocation.Body
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
                        FormDataParameters.Any(p => p.ModelType.IsPrimaryType(KnownPrimaryType.Stream))
                            ? "autorest.WithMultiPartFormData(formDataParameters)"
                            : "autorest.WithFormData(autorest.MapToValues(formDataParameters))"
                        );
                }

                if (HeaderParameters.Any())
                {
                    foreach (var param in Parameters.Where(p => p.IsRequired && p.Location == ParameterLocation.Header))
                    {
                        if (param.IsClientProperty)
                        {
                            decorators.Add(string.Format("autorest.WithHeader(\"{0}\",client.{1})", param.SerializedName, param.Name.ToPascalCase().ToString()));
                        }
                        else
                        {
                            decorators.Add(string.Format("autorest.WithHeader(\"{0}\",autorest.String({1}))", param.SerializedName, param.Name.ToString()));
                        }
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
                var decorators = new List<string>();
                decorators.Add("req");
                if (RegisterRP)
                {
                    decorators.Add("azure.DoRetryWithRegistration(client.Client)");
                } else {
                    decorators.Add("autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...)");
                }
                if (IsLongRunningOperation())
                {
                    decorators.Add("azure.DoPollForAsynchronous(client.PollingDelay)");
                }
                return decorators;
            }
        }

        public IEnumerable<string> RespondDecorators
        {
            get
            {
                var decorators = new List<string>();
                decorators.Add("resp");
                decorators.Add("client.ByInspecting()");
                decorators.Add(string.Format("azure.WithErrorUnlessStatusCode({0})", string.Join(",", ResponseCodes.ToArray())));

                var rvNeedsUnmarshalling = ReturnValueRequiresUnmarshalling();
                if (rvNeedsUnmarshalling && !ReturnValue().Body.IsStreamType())
                {
                    var rv = ReturnValue().Body as CompositeTypeGo;
                    if (rv.IsWrapperType && !rv.HasPolymorphicFields && !rv.XmlIsWrapped)
                    {
                        decorators.Add($"autorest.ByUnmarshallingJSON(&result.Value)");
                    }
                    else
                    {
                        decorators.Add($"autorest.ByUnmarshallingJSON(&result)");
                    }
                }

                if (!rvNeedsUnmarshalling || !ReturnValue().Body.IsStreamType())
                {
                    decorators.Add("autorest.ByClosing()");
                }
                return decorators;
            }
        }

        public string Response
        {
            get
            {
                return HasReturnValue()
                    ? "result.Response = autorest.Response{Response: resp}"
                    : "result.Response = resp";
            }
        }

        /// <summary>
        /// Check if method has a return response.
        /// </summary>
        public bool HasReturnValue()
        {
            return ReturnValue()?.Body != null;
        }

        /// <summary>
        /// Returns true if the return type requires unmarshalling.
        /// </summary>
        public bool ReturnValueRequiresUnmarshalling()
        {
            return HasReturnValue() && ReturnValue().Body is CompositeTypeGo && ((CompositeTypeGo)ReturnValue().Body).Properties.Any();
        }

        /// <summary>
        /// Returns true if the return type is a wrapper type (i.e. "synthetic" type).
        /// </summary>
        public bool ReturnValueIsWrapperType()
        {
            return HasReturnValue() && ReturnValue().Body is CompositeTypeGo && ReturnValue().Body.Cast<CompositeTypeGo>().IsWrapperType;
        }

        /// <summary>
        /// Returns true if the return type is XML-wrapped.
        /// </summary>
        public bool ReturnValueIsXmlWrapped()
        {
            if (!HasReturnValue())
            {
                return false;
            }

            var ctg = ReturnValue().Body as CompositeTypeGo;
            if (ctg == null)
            {
                return false;
            }

            if (ctg.XmlProperties == null)
            {
                return false;
            }
            return ctg.XmlProperties.Wrapped;
        }

        /// <summary>
        /// Return response object for the method.
        /// </summary>
        public Response ReturnValue()
        {
            return ReturnType ?? DefaultResponse;
        }

        /// <summary>
        /// Returns true if this method's body parameter requires marshalling.
        /// </summary>
        public bool BodyParamNeedsMarshalling()
        {
            return BodyParameter != null && !BodyParameter.ModelType.IsPrimaryType(KnownPrimaryType.Stream);
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

        public string NextOperationName
        {
            get
            {
                return NextMethod?.Name.Value;
            }
        }

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
                if (Extensions.ContainsKey(AzureExtensions.PageableExtension))
                {
                    var pageableExtension = Extensions[AzureExtensions.PageableExtension] as Newtonsoft.Json.Linq.JContainer;
                    if (pageableExtension != null)
                    {
                        var nextLink = (string)pageableExtension["nextLinkName"];
                        if (!string.IsNullOrEmpty(nextLink))
                        {
                            return CodeNamerGo.Instance.GetPropertyName(nextLink);
                        }
                    }
                }
                return null;
            }
        }

        /// <summary>
        /// Gets the name of the body parameter to be passed to the marshaller (might be XML wrapped).
        /// </summary>
        public string BodyParamNameForMarshalling
        {
            get
            {
                if (BodyParameter.ModelType.XmlIsWrapped)
                {
                    return $"{BodyParameter.ModelType.XmlName}{{Value: {BodyParameter.Name}}}";
                }
                else
                {
                    return BodyParameter.Name;
                }
            }
        }
    }
}
