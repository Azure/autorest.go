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
using Newtonsoft.Json.Linq;
using AutoRest.Core.Logging;

namespace AutoRest.Go.Model
{
    public class MethodGo : Method
    {
        internal const string DefaultReturnType = "autorest.Response";

        public string Owner { get; private set; }

        public string PackageName { get; private set; }

        public string APIVersion { get; private set; }

        public bool NextAlreadyDefined { get; private set; }

        /// <summary>
        /// The method name qualified with the client name
        /// </summary>
        public string QualifiedName => $"{Owner}.{Name}";

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

            // Fixing the returnType in modeler
            // find all the non-error responses with non-empty body
            // we have to ignore those non-error responses with empty bodies, because there are quite plenty of swaggers with one response with body and the other without body
            Logger.Instance.Log(Category.Debug, $"All responses of {SerializedName}: {string.Join(", ", Responses.Select(kv => $"{(int)kv.Key}: {kv.Value?.Body?.Name}"))}");
            var nonErrorNonEmptyStatusCodes = NonErrorNonEmptyStatusCodeDict;
            Logger.Instance.Log(Category.Debug, $"Method {SerializedName} has the following non-error & non-empty responses (total {nonErrorNonEmptyStatusCodes.Count()}): {string.Join(", ", nonErrorNonEmptyStatusCodes.Select(resp => resp.Value?.Body?.Name))}");
            // categorize the models by its name
            var nonErrorNonEmptyResponses = NonErrorNonEmptyResponses.ToList();
            Logger.Instance.Log(Category.Debug, $"Non-error & non-empty models: {string.Join(", ", nonErrorNonEmptyResponses.Select(model => model.Body.Name))}");
            // fix the original return type from the modeler
            ReturnType = FixedReturnType();
            Logger.Instance.Log(Category.Debug, $"Return Type of {SerializedName}: {ReturnType?.Body?.Name}");
        }

        /// <summary>
        /// This function should only be used in the transformation to get the non-error return type ahead of time.
        /// </summary>
        /// <returns></returns>
        public Response FixedReturnType()
        {
            var nonErrorNonEmptyResponses = NonErrorNonEmptyResponses;
            switch (nonErrorNonEmptyResponses.Count())
            {
                case 0:
                    return new Response();
                case 1:
                    return nonErrorNonEmptyResponses.First();
                default:
                    // In this case, we do nothing but let the original value in ReturnType to take effect
                    Logger.Instance.Log(Category.Debug, $"we have more than one non-error responses with non-empty but different schemas in operationId {SerializedName}, but in this case we just honor the return type in the modeler");
                    return ReturnType;
            }
        }

        public IEnumerable<KeyValuePair<HttpStatusCode, Response>> NonErrorNonEmptyStatusCodeDict => Responses.Where(kv => !kv.Value.IsErrorResponse() && kv.Value.Body != null);

        public IEnumerable<Response> NonErrorNonEmptyResponses => NonErrorNonEmptyStatusCodeDict.Select(kv => kv.Value).Distinct(new ResponseEqualityComparer());

        public IEnumerable<IModelType> NonErrorNonEmptyResponseModels => NonErrorNonEmptyResponses.Select(r => r.Body);

        private struct ResponseEqualityComparer : IEqualityComparer<Response>
        {

            public bool Equals(Response x, Response y)
            {
                return x?.Body?.Name == y?.Body?.Name;
            }

            public int GetHashCode(Response obj)
            {
                return obj?.Body?.Name.GetHashCode() ?? 0;
            }
        }

        /// <summary>
        /// Returns true if the local parameters contain documentation.
        /// </summary>
        public bool AddParamsDoc
        {
            get
            {
                foreach (var parameter in LocalParameters)
                {
                    if (!string.IsNullOrWhiteSpace(parameter.Documentation.FixedValue))
                    {
                        return true;
                    }
                }
                return false;
            }
        }

        /// <summary>
        /// Generate the method parameter declaration.
        /// </summary>
        /// <param name="includePkgName">Pass true if the type name should include the package prefix.  Defaults to false.</param>
        public string MethodParametersSignature(bool includePkgName = false)
        {
            var declarations = new List<string> { "ctx context.Context" };
            LocalParameters
                .ForEach(p => declarations.Add(string.Format(
                                                    p.IsRequired || p.ModelType.CanBeEmpty()
                                                        ? "{0} {1}"
                                                        : "{0} *{1}", p.Name, p.ModelType.HasInterface()
                                                            ? p.ModelType.GetInterfaceName(includePkgName)
                                                            : ParameterTypeSig(p.ModelType, includePkgName))));
            return string.Join(", ", declarations);
        }

        private string ParameterTypeSig(IModelType type, bool includePkgName)
        {
            if (includePkgName)
            {
                if (type.IsUserDefinedType())
                    return $"{CodeModel.Namespace}.{type.Name}";
                else if (type is SequenceTypeGo stg)
                    return stg.NameWithPackagePrefix;
                else if (type is DictionaryTypeGo dtg)
                    return dtg.NameWithPackagePrefix;
            }
            return type.Name.ToString();
        }

        /// <summary>
        /// Gets the return type name for this method.
        /// </summary>
        /// <param name="includePkgName">Pass true if the type name should include the package prefix.  Defaults to false.</param>
        public string MethodReturnType(bool includePkgName = false)
        {
            return HasReturnValue()
                ? includePkgName
                    ? $"{CodeModel.Namespace}.{ReturnType.Body.Name}"
                    : ReturnType.Body.Name.ToString()
                : DefaultReturnType;
        }

        private string MethodReturnSig(string resultTypeName)
        {
            return $"result {resultTypeName}, err error";
        }

        /// <summary>
        /// Returns the method return signature for this method (e.g. "foo, bar").
        /// For responder methods use ResponderReturnSignature() instead.
        /// </summary>
        /// <param name="includePkgName">Pass true if the type name should include the package prefix.  Defaults to false.</param>
        /// <returns>The method signature for this method.</returns>
        public string MethodReturnSignature(bool includePkgName = false)
        {
            return MethodReturnSig(MethodReturnType(includePkgName));
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
            else if (IsPageable && !IsNextMethod)
            {
                return MethodReturnSig(ReturnType.Body.Cast<PageTypeGo>().ContentType.Name);
            }
            return MethodReturnSignature();
        }

        /// <summary>
        /// Returns the method return signature for the next results page method (e.g. "foo, bar").
        /// </summary>
        /// <returns>The method signature for the next results page method.</returns>
        public string NextMethodReturnSignature()
        {
            return MethodReturnSig(ReturnType.Body.Cast<CompositeTypeGo>().UnwrapPageType().ContentType.Name);
        }

        /// <summary>
        /// Returns the type name used as the parameter for the "next results" method.
        /// </summary>
        /// <returns>The "next results" method parameter type name.</returns>
        public string LastResultsTypeName()
        {
            var type = ReturnType.Body;
            if (IsLongRunningOperation())
            {
                type = type.Cast<FutureTypeGo>().ResultType;
            }
            if (IsPageable && !IsNextMethod)
            {
                type = type.Cast<PageTypeGo>().ContentType;
            }
            return type.Name;
        }

        public string NextMethodName => $"{Name.ToCamelCase()}NextResults";

        public string PreparerMethodName => $"{Name}Preparer";

        public string SenderMethodName => $"{Name}Sender";

        public string ResponderMethodName => $"{Name}Responder";

        public string ListCompleteMethodName => $"{Name}Complete";

        public string HelperInvocationParameters()
        {
            var invocationParams = new List<string> { "ctx" };

            foreach (ParameterGo p in LocalParameters)
            {
                invocationParams.Add(p.Name);
            }
            return string.Join(", ", invocationParams);
        }

        /// <summary>
        /// Calculates the args to be passed to the "next method".
        /// </summary>
        /// <param name="nextLink">The arg to be passed in the "next link" param.</param>
        /// <returns>The params string, e.g. "ctx, foo, nextLink".</returns>
        public string NextMethodInvocationParameters(string nextLink)
        {
            // some next methods take the same params as the "list initial" method plus
            // the next link param.  so if the param counts match assume this is the case.
            // to date, the only place where this appears is in the autorest tests.
            if (NextMethod.LocalParameters.Count() == LocalParameters.Count() + 1)
            {
                return $"{HelperInvocationParameters()}, {nextLink}";
            }

            // attempt to match our local params to that of the next method.
            // by convention ctx is always the first parameter.
            // NOTE: the context param is implicit, i.e. it isn't part of the code model
            var invocationParams = new List<string> { "ctx" };

            // short-circuit simple case, if the next method takes
            // one parameter then it can only be nextLink
            if (NextMethod.LocalParameters.Count() == 1)
            {
                invocationParams.Add(nextLink);
            }
            else
            {
                // create param lists so we can walk them by ordinal
                var myMethodParams = LocalParameters.ToList();
                var nextMethodParams = NextMethod.LocalParameters.ToList();

                for (int i = 0; i < nextMethodParams.Count; ++i)
                {
                    if (nextMethodParams[i].Name.EqualsIgnoreCase("nextlink"))
                    {
                        invocationParams.Add(nextLink);
                    }
                    else if (i < myMethodParams.Count && ParameterGo.Match(myMethodParams[i], nextMethodParams[i]))
                    {
                        invocationParams.Add(myMethodParams[i].Name);
                    }
                    else
                    {
                        // try to find a match in our local params
                        var param = myMethodParams
                            .Where(p => ParameterGo.Match(p, nextMethodParams[i]))
                            .FirstOrDefault();

                        if (param == null)
                        {
                            throw new Exception("failed to find a matching local parameter");
                        }

                        invocationParams.Add(param.Name);
                    }
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

        public IEnumerable<ParameterGo> OptionalFormDataParameters => ParametersGo.FormDataParameters(false);

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
                // Only add the response to responseCodes when it is not marked by x-ms-error-response: true
                foreach (var response in Responses)
                {
                    if (!response.Value.IsErrorResponse())
                    {
                        codes.Add(CodeNamerGo.Instance.StatusCodeToGoString[response.Key]);
                    }
                }
                return codes;
            }
        }

        public IEnumerable<string> PrepareDecorators
        {
            get
            {
                var decorators = new List<string>();

                if (BodyParameter != null)
                {
                    decorators.Add($"autorest.AsContentType(\"{RequestContentType}\")");
                }

                decorators.Add(HTTPMethodDecorator);
                if (!IsCustomBaseUri)
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
                    var bodyParam = BodyParameter.Name;
                    if (BodyParameter.IsConstant)
                    {
                        bodyParam = BodyParameter.DefaultValue;
                    }
                    decorators.Add(string.Format(BodyParameter.ModelType.PrimaryType(KnownPrimaryType.Stream) && BodyParameter.Location == ParameterLocation.Body
                                        ? "autorest.WithFile({0})"
                                        : "autorest.WithJSON({0})",
                                bodyParam));
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
                    foreach (var param in ParametersGo.Where(p => p.IsRequired && p.Location == ParameterLocation.Header))
                    {
                        string value;
                        if (param.IsConstant)
                        {
                            value = param.DefaultValueString;
                        }
                        else if (param.IsClientProperty)
                        {
                            value = param.GetClientPropertryName();
                        }
                        else
                        {
                            value = $"autorest.String({param.Name})";
                        }
                        decorators.Add($"autorest.WithHeader(\"{param.SerializedName}\", {value})");
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
                    string.Format("azure.WithErrorUnlessStatusCode({0})", string.Join(",", ResponseCodes.ToArray()))
                };

                if (HasReturnValue() && !ReturnType.Body.IsStreamType() && !LroWrapsDefaultResp())
                {
                    if ((((CompositeTypeGo)ReturnType.Body).IsWrapperType || (ReturnType.Body is FutureTypeGo ftg && ftg.IsResultWrapperType)) && !((CompositeTypeGo)ReturnType.Body).HasPolymorphicFields)
                    {
                        decorators.Add("autorest.ByUnmarshallingJSON(&result.Value)");
                    }
                    else
                    {
                        decorators.Add("autorest.ByUnmarshallingJSON(&result)");
                    }
                }

                if (!HasReturnValue() || !ReturnType.Body.IsStreamType())
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

            var retType = ReturnType.Body as FutureTypeGo;
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
            var lhs = "result.Response";
            if (!forResponder)
            {
                lhs = $"{ResponseAssignTarget}.Response";
            }
            return $"{lhs} = autorest.Response{{Response: resp}}";
        }

        /// <summary>
        /// Gets the left-hand side of the response assignment.
        /// It can be different if the method is pageable.
        /// </summary>
        public string ResponseAssignTarget
        {
            get
            {
                var target = "result";
                if (IsPageable && !IsNextMethod)
                {
                    target = $"result.{ReturnType.Body.Cast<CompositeTypeGo>().UnwrapPageType().ResultFieldName}";
                }
                return target;
            }
        }

        public string AutorestError(string phase, string response = null, string parameter = null, string methodName = null)
        {
            if (methodName == null)
            {
                methodName = Name;
            }
            return !string.IsNullOrEmpty(parameter)
                        ? string.Format("autorest.NewErrorWithError(err, \"{0}.{1}\", \"{2}\", nil , \"{3}\'{4}\'\")", PackageName, Owner, methodName, phase, parameter)
                        : string.IsNullOrEmpty(response)
                                 ? string.Format("autorest.NewErrorWithError(err, \"{0}.{1}\", \"{2}\", nil , \"{3}\")", PackageName, Owner, methodName, phase)
                                 : string.Format("autorest.NewErrorWithError(err, \"{0}.{1}\", \"{2}\", {3}, \"{4}\")", PackageName, Owner, methodName, response, phase);
        }

        public string ValidationError => $"validation.NewError(\"{PackageName}.{Owner}\", \"{Name}\", err.Error())";

        /// <summary>
        /// Check if method has a return response.
        /// </summary>
        /// <returns></returns>
        public bool HasReturnValue()
        {
            return ReturnType?.Body != null;
        }

        /// <summary>
        /// Returns true if method has pageable extension (x-ms-pageable) with a non-null nextLinkName.
        /// </summary>
        public bool IsPageable
        {
            get
            {
                if (!Extensions.ContainsKey(AzureExtensions.PageableExtension))
                {
                    return false;
                }
                // if the nextLinkName field in the swagger has a null value ("nextLinkName": null)
                // then don't treat this operation as pageable.
                var pageableExtension = Extensions[AzureExtensions.PageableExtension] as JContainer;
                return !string.IsNullOrWhiteSpace((string)pageableExtension["nextLinkName"]);
            }
        }

        /// <summary>
        /// Returns true if method is a "next method" as defined in swagger (x-ms-pageable:operationName).
        /// </summary>
        public bool IsNextMethod => Name.Value.EqualsIgnoreCase(NextOperationName);

        /// <summary>
        /// Returns true if a ListComplete method should be generated.
        /// </summary>
        public bool NeedsListComplete => IsPageable && !IsNextMethod;

        /// <summary>
        /// Returns the name of the type returned from a ListComplete method.
        /// This should only be called if NeedsListComplete returns true.
        /// </summary>
        public string ListCompleteReturnSig(bool includePkgName = false)
        {
            var resultTypeName = ReturnType.Body.Cast<CompositeTypeGo>().UnwrapPageType().IteratorType.Name;
            if (IsLongRunningOperation())
            {
                resultTypeName = ((LroPagedResponseGo)ReturnType).ListAllReturnType.Name;
            }
            if (includePkgName)
            {
                resultTypeName = $"{CodeModel.Namespace}.{resultTypeName}";
            }
            return MethodReturnSig(resultTypeName);
        }

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
    }
}
