/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { values } from '@azure-tools/linq';
import * as tcgc from '@azure-tools/typespec-client-generator-core';
import { NoTarget } from '@typespec/compiler';
import * as go from '../../../codemodel.go/src/index.js';
import { capitalize, createOptionsTypeDescription, createResponseEnvelopeDescription, ensureNameCase, getEscapedReservedName, uncapitalize } from '../../../naming.go/src/naming.js';
import { AdapterError } from './errors.js';
import { GoEmitterOptions } from '../lib.js';
import { isTypePassedByValue, typeAdapter } from './types.js';

// used to convert SDK clients and their methods to Go code model types
export class clientAdapter {
  private ta: typeAdapter;
  private opts: GoEmitterOptions;

  // track all of the client and parameter group params across all operations
  // as not every option might contain them, and parameter groups can be shared
  // across multiple operations
  private clientParams: Map<string, go.MethodParameter>;

  constructor(ta: typeAdapter, opts: GoEmitterOptions) {
    this.ta = ta;
    this.opts = opts;
    this.clientParams = new Map<string, go.MethodParameter>();
  }

  // converts all clients and their methods to Go code model types.
  // this includes parameter groups/options types and response envelopes.
  adaptClients(sdkPackage: tcgc.SdkPackage<tcgc.SdkHttpOperation>) {
    if (this.opts['single-client'] && sdkPackage.clients.length > 1) {
      throw new AdapterError('InvalidArgument', 'single-client cannot be enabled when there are multiple clients', NoTarget);
    }
    for (const sdkClient of sdkPackage.clients) {
      // start with instantiable clients and recursively work down
      if (sdkClient.clientInitialization.initializedBy & tcgc.InitializedByFlags.Individually) {
        this.recursiveAdaptClient(sdkClient);
      }
    }
  }

  private recursiveAdaptClient(sdkClient: tcgc.SdkClientType<tcgc.SdkHttpOperation>, parent?: go.Client): go.Client | undefined {
    if (sdkClient.methods.length === 0 && (sdkClient.children === undefined || sdkClient.children.length === 0)) {
      // skip generating empty clients
      return undefined;
    }

    let clientName = ensureNameCase(sdkClient.name);

    // to keep compat with existing ARM packages, don't use hierarchically named clients
    if (parent && this.ta.codeModel.type !== 'azure-arm') {
      // for hierarchical clients, the child client names are built
      // from the parent client name. this is because tsp allows subclients
      // with the same name. consider the following example.
      //
      // namespace Chat {
      //   interface Completions {
      //     ...
      //   }
      // }
      // interface Completions { ... }
      //
      // we want to generate two clients from this,
      // one name ChatCompletions and the other Completions

      // strip off the Client suffix from the parent client name
      clientName = parent.name.substring(0, parent.name.length - 6) + clientName;
    }

    const docs: go.Docs = {
      summary: sdkClient.summary,
      description: sdkClient.doc,
    };

    if (docs.summary) {
      docs.summary = `${clientName} - ${docs.summary}`;
    } else if (docs.description) {
      docs.description = `${clientName} - ${docs.description}`;
    } else if (clientName.length > 6) {
      // strip clientName's "Client" suffix
      const groupName = clientName.substring(0, clientName.length - 6);
      docs.summary = `${clientName} contains the methods for the ${groupName} group.`;
    } else {
      // the client name is simply "Client"
      docs.summary = `${clientName} contains the methods for the service.`;
    }

    const goClient = new go.Client(clientName, docs, go.newClientOptions(this.ta.codeModel.type, clientName));
    goClient.parent = parent;

    // anything other than public means non-instantiable client
    if (sdkClient.clientInitialization.initializedBy & tcgc.InitializedByFlags.Individually) {
      for (const param of sdkClient.clientInitialization.parameters) {
        switch (param.kind) {
          case 'credential':
            // skip this for now as we don't generate client constructors
            continue;
          case 'endpoint': {
            if (this.ta.codeModel.type === 'azure-arm') {
              // for ARM, the endpoint is handled via the azcore/arm.Client
              // so we don't need to adapt it.
              continue;
            }

            let endpointType: tcgc.SdkEndpointType;
            switch (param.type.kind) {
              case 'endpoint':
                // single endpoint without any supplemental path
                endpointType = param.type;
                break;
              case 'union':
                // this is a union of endpoints. the first is the endpoint plus
                // the supplemental path. the second is a "raw" endpoint which
                // requires the caller to provide the complete endpoint. we only
                // expose the former at present. languages that support overloads
                // MAY support both but it's not a requirement.
                endpointType = param.type.variantTypes[0];
            }

            for (let i = 0; i < endpointType.templateArguments.length; ++i) {
              const templateArg = endpointType.templateArguments[i];
              if (i === 0) {
                // the first template arg is always the endpoint parameter.
                // note that the types of the param and the field are different.
                // NOTE: we use param.name here instead of templateArg.name as
                // the former has the fixed name "endpoint" which is what we want.
                const adaptedParam = this.adaptURIParam(templateArg);
                adaptedParam.docs.summary = templateArg.summary;
                adaptedParam.docs.description = templateArg.doc;
                goClient.parameters.push(adaptedParam);

                // if the server's URL is *only* the endpoint parameter then we're done.
                // this is the param.type.kind === 'endpoint' case.
                if (endpointType.serverUrl === `{${templateArg.serializedName}}`) {
                  break;
                }

                goClient.templatedHost = endpointType.serverUrl;
                continue;
              }

              const adaptedParam = this.adaptURIParam(templateArg);
              adaptedParam.docs.summary = templateArg.summary;
              adaptedParam.docs.description = templateArg.doc;
              goClient.parameters.push(adaptedParam);
            }
            break;
          }
          case 'method':
            // some client params, notably api-version, can be explicitly
            // defined in the operation signature:
            // e.g. op withQueryApiVersion(@query("api-version") apiVersion: string)
            // these get propagated to sdkMethod.operation.parameters thus they
            // will be adapted in adaptMethodParameters()
            continue;
        }
      }
    } else if (parent) {
      // this is a sub-client. it will share the client/host params of the parent.
      // NOTE: we must propagate parant params before a potential recursive call
      // to create a child client that will need to inherit our client params.
      goClient.templatedHost = parent.templatedHost;

      // make a copy of the client params. this is to prevent
      // client method params from being shared across clients
      // as not all client method params might be uniform.
      goClient.parameters = new Array<go.ClientParameter>(...parent.parameters);
    } else {
      throw new AdapterError('InternalError', `uninstantiable client ${sdkClient.name} has no parent`, NoTarget);
    }
    if (sdkClient.children && sdkClient.children.length > 0) {
      for (const child of sdkClient.children) {
        const subClient = this.recursiveAdaptClient(child, goClient);
        if (subClient) {
          goClient.clientAccessors.push(new go.ClientAccessor(`New${subClient.name}`, subClient));
        }
      }
    }
    for (const sdkMethod of sdkClient.methods) {
      this.adaptMethod(sdkMethod, goClient);
    }

    if (this.ta.codeModel.type === 'azure-arm' && goClient.clientAccessors.length > 0 && goClient.methods.length === 0) {
      // this is the service client. to keep compat with existing
      // ARM SDKs we skip adding it to the code model in favor of
      // the synthesized client factory.
    } else {
      this.ta.codeModel.clients.push(goClient);
    }
    return goClient;
  }

  private adaptURIParam(sdkParam: tcgc.SdkPathParameter): go.URIParameter {
    const paramType = this.ta.getWireType(sdkParam.type, true, false);
    if (go.isURIParameterType(paramType)) {
      // TODO: follow up with tcgc if serializedName should actually be optional
      return new go.URIParameter(sdkParam.name, sdkParam.serializedName ? sdkParam.serializedName : sdkParam.name, paramType,
        this.adaptParameterStyle(sdkParam), isTypePassedByValue(sdkParam.type) || !sdkParam.optional, 'client');
    }
    throw new AdapterError('UnsupportedTsp', `unsupported URI parameter type ${paramType.kind}`, sdkParam.__raw?.node ?? NoTarget);
  }

  private adaptMethod(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, goClient: go.Client) {
    let method: go.MethodType;
    const naming = new go.MethodNaming(getEscapedReservedName(uncapitalize(ensureNameCase(sdkMethod.name)), 'Operation'), ensureNameCase(`${sdkMethod.name}CreateRequest`, true),
      ensureNameCase(`${sdkMethod.name}HandleResponse`, true));

    const getStatusCodes = function (httpOp: tcgc.SdkHttpOperation): Array<number> {
      const statusCodes = new Array<number>();
      for (const response of httpOp.responses) {
        const statusCode = response.statusCodes;
        if (isHttpStatusCodeRange(statusCode)) {
          for (let code = statusCode.start; code <= statusCode.end; ++code) {
            statusCodes.push(code);
          }
        } else {
          statusCodes.push(statusCode);
        }
      }
      return statusCodes;
    };

    let methodName = capitalize(ensureNameCase(sdkMethod.name));
    if (sdkMethod.access === 'internal') {
      // we add internal to the extra list so we don't end up with a method named "internal"
      // which will collide with an unexported field with the same name.
      methodName = getEscapedReservedName(uncapitalize(methodName), 'Method', ['internal']);
    }
    const statusCodes = getStatusCodes(sdkMethod.operation);

    if (sdkMethod.kind === 'basic') {
      method = new go.Method(methodName, goClient, sdkMethod.operation.path, sdkMethod.operation.verb, statusCodes, naming);
    } else if (sdkMethod.kind === 'paging') {
      if (sdkMethod.pagingMetadata.nextLinkReInjectedParametersSegments !== undefined && sdkMethod.pagingMetadata.nextLinkReInjectedParametersSegments.length > 0) {
        throw new AdapterError('UnsupportedTsp', `paging with re-injected parameters is not supported`, sdkMethod.__raw?.node ?? NoTarget);
      }
      method = new go.PageableMethod(methodName, goClient, sdkMethod.operation.path, sdkMethod.operation.verb, statusCodes, naming);
      if (sdkMethod.pagingMetadata.nextLinkSegments) {
        method.nextLinkName = capitalize(sdkMethod.pagingMetadata.nextLinkSegments.map((segment) => {
          if (segment.kind === 'property') {
            return ensureNameCase(segment.name);
          } else {
            throw new AdapterError('UnsupportedTsp', `unsupported next link segment kind ${segment.kind}`, sdkMethod.__raw?.node ?? NoTarget);
          }
        }).join('.'));
      }
    } else if (sdkMethod.kind === 'lro') {
      method = new go.LROMethod(methodName, goClient, sdkMethod.operation.path, sdkMethod.operation.verb, statusCodes, naming);
      const lroOptions = this.hasDecorator('Azure.Core.@useFinalStateVia', sdkMethod.decorators);
      if (lroOptions) {
        method.finalStateVia = <go.FinalStateVia>lroOptions['finalState'];
      }
      if (sdkMethod.lroMetadata.finalResponse?.resultSegments) {
        // 'resultSegments' is designed for furture extensibility, currently only has one segment
        method.operationLocationResultPath = sdkMethod.lroMetadata.finalResponse.resultSegments.map((segment) => {
          return (<tcgc.SdkBodyModelPropertyType>segment).serializationOptions.json?.name;
        }).join('.');
      }
    } else {
      throw new AdapterError('UnsupportedTsp', `unsupported method kind ${sdkMethod.kind}`, sdkMethod.__raw?.node ?? NoTarget);
    }

    method.docs.summary = sdkMethod.summary;
    method.docs.description = sdkMethod.doc;
    goClient.methods.push(method);
    this.populateMethod(sdkMethod, method);
  }

  /* eslint-disable-next-line @typescript-eslint/no-explicit-any */
  private hasDecorator(name: string, decorators: Array<tcgc.DecoratorInfo>): Record<string, any> | undefined {
    for (const decorator of decorators) {
      if (decorator.name === name) {
        return decorator.arguments;
      }
    }
    return undefined;
  }

  private populateMethod(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, method: go.MethodType | go.NextPageMethod) {
    if (method.kind === 'nextPageMethod') {
      throw new AdapterError('UnsupportedTsp', `unsupported method kind ${sdkMethod.kind}`, sdkMethod.__raw?.node ?? NoTarget);
    }

    let prefix = method.client.name;
    if (this.opts['single-client']) {
      prefix = '';
    }
    if (go.isLROMethod(method)) {
      prefix += 'Begin';
    }
    let optionalParamsGroupName = `${prefix}${method.name}Options`;
    if (sdkMethod.access === 'internal') {
      optionalParamsGroupName = uncapitalize(optionalParamsGroupName);
    }
    let optsGroupName = 'options';
    // if there's an existing required parameter with the name options then pick something else.
    // optional params will be inside the options type, so they can never collide.
    for (const param of sdkMethod.parameters) {
      if (!param.optional && param.name === optsGroupName) {
        optsGroupName = 'opts';
        break;
      }
    }
    method.optionalParamsGroup = new go.ParameterGroup(optsGroupName, optionalParamsGroupName, false, 'method');
    method.optionalParamsGroup.docs.summary = createOptionsTypeDescription(optionalParamsGroupName, this.getMethodNameForDocComment(method));
    method.responseEnvelope = this.adaptResponseEnvelope(sdkMethod, method);

    // find the api version param to use for the doc comment.
    // we can't use sdkMethod.apiVersions as that includes all
    // of the api versions supported by the service.
    for (const opParam of sdkMethod.operation.parameters) {
      if (opParam.isApiVersionParam && opParam.clientDefaultValue) {
        method.apiVersions.push(<string>opParam.clientDefaultValue);
        break;
      }
    }

    const paramMapping = this.adaptMethodParameters(sdkMethod, method);

    // we must do this after adapting method params as it can add optional params
    this.ta.codeModel.paramGroups.push(this.adaptParameterGroup(method.optionalParamsGroup));

    if (this.ta.codeModel.options.generateExamples) {
      this.adaptHttpOperationExamples(sdkMethod, method, paramMapping);
    }
  }

  private adaptMethodParameters(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, method: go.MethodType | go.NextPageMethod): Map<tcgc.SdkHttpParameter, Array<go.MethodParameter>> {
    const paramMapping = new Map<tcgc.SdkHttpParameter, Array<go.MethodParameter>>();

    let optionalGroup: go.ParameterGroup | undefined;
    if (method.kind !== 'nextPageMethod') {
      // NextPageMethods don't have optional params
      optionalGroup = method.optionalParamsGroup;
      if (go.isLROMethod(method)) {
        optionalGroup.params.push(new go.ResumeTokenParameter());
      }
    }

    // stuff all of the operation parameters into one array for easy traversal
    type OperationParamType = tcgc.SdkBodyParameter | tcgc.SdkHeaderParameter | tcgc.SdkPathParameter | tcgc.SdkQueryParameter | tcgc.SdkCookieParameter;
    const allOpParams = new Array<OperationParamType>();
    allOpParams.push(...sdkMethod.operation.parameters);
    if (sdkMethod.operation.bodyParam) {
      allOpParams.push(sdkMethod.operation.bodyParam);
    }

    // we must enumerate parameters, not operation.parameters, as it
    // contains the params in tsp order as well as any spread params.
    for (const param of sdkMethod.parameters) {
      // we need to translate from the method param to its underlying operation param.
      // most params have a one-to-one mapping. however, for spread params, there will
      // be a many-to-one mapping. i.e. multiple params will map to the same underlying
      // operation param. each param corresponds to a field within the operation param.
      let opParam = values(allOpParams).where((opParam: OperationParamType) => {
        return values(opParam.correspondingMethodParams).where((methodParam: tcgc.SdkModelPropertyType) => {
          if (param.type.kind === 'model') {
            for (const property of param.type.properties) {
              if (property === methodParam) {
                return true;
              }
            }
          }
          return methodParam === param;
        }).any();
      }).first();

      // special handling for constants that used in path, this will not be in operation parameters since it has been resolved in the url
      if (!opParam && param.type.kind === 'constant') {
        continue;
      }

      // special handling for `@bodyRoot`/`@body` on model param's property
      if (!opParam && param.type.kind === 'model') {
        for (const property of param.type.properties) {
          opParam = values(allOpParams).where((opParam: OperationParamType) => {
            return values(opParam.correspondingMethodParams).where((methodParam: tcgc.SdkModelPropertyType) => {
              return methodParam === property;
            }).any();
          }).first();
          if (opParam) {
            break;
          }
        }
      }

      if (!opParam) {
        throw new AdapterError('InternalError', `didn't find operation parameter for method ${sdkMethod.name} parameter ${param.name}`, sdkMethod.__raw?.node ?? NoTarget);
      }

      if (opParam.kind === 'header' && opParam.serializedName.match(/^content-type$/i) && param.type.kind === 'constant') {
        // if the body param is optional, then the content-type param is also optional.
        // for optional constants, this has the side effect of the param being treated like
        // a flag which isn't what we want. so, we mark it as required. we ONLY do this
        // if the content-type is a constant (i.e. literal value).
        // the content-type will be conditionally set based on the requiredness of the body.
        opParam.optional = false;
      }

      let adaptedParam: go.MethodParameter;
      if (opParam.kind === 'body' && opParam.type.kind === 'model' && opParam.type.kind !== param.type.kind) {
        const paramStyle = this.adaptParameterStyle(param);
        const paramName = getEscapedReservedName(ensureNameCase(param.name, paramStyle === 'required'), 'Param');
        const byVal = isTypePassedByValue(param.type);
        const contentType = this.adaptContentType(opParam.defaultContentType);
        const getSerializedNameFromProperty = function(property: tcgc.SdkBodyModelPropertyType): string | undefined {
          if (contentType === 'JSON') {
            return property.serializationOptions.json?.name;
          }
          if (contentType === 'XML') {
            return property.serializationOptions.xml?.name;
          }
          if (contentType === 'binary') {
            return property.serializationOptions.multipart?.name;
          }
          return undefined;
        };
        switch (contentType) {
          case 'JSON':
          case 'XML': {
            // find the corresponding field within the model param so we can get the serialized name
            let serializedName: string | undefined;
            for (const property of opParam.type.properties) {
              if (property.name === param.name) {
                serializedName = getSerializedNameFromProperty(<tcgc.SdkBodyModelPropertyType>property);
                break;
              }
            }
            if (!serializedName) {
              throw new AdapterError('InternalError', `didn't find body model property for spread parameter ${param.name}`, param.__raw?.node ?? NoTarget);
            }
            adaptedParam = new go.PartialBodyParameter(paramName, serializedName, contentType, this.ta.getWireType(param.type, true, true), paramStyle, byVal);
            break;
          }
          case 'binary':
            if (opParam.defaultContentType.match(/multipart/i)) {
              adaptedParam = new go.MultipartFormBodyParameter(paramName, this.ta.getReadSeekCloser(false), paramStyle, byVal);
            } else {
              adaptedParam = new go.BodyParameter(paramName, contentType, `"${opParam.defaultContentType}"`, this.ta.getReadSeekCloser(false), paramStyle, byVal);
            }
            break;
          default:
            throw new AdapterError('UnsupportedTsp', `unsupported spread param content type ${contentType}`, opParam.__raw?.node ?? NoTarget);
        }
      } else {
        adaptedParam = this.adaptMethodParameter(opParam, method.httpMethod);
      }

      adaptedParam.docs.summary = param.summary;
      adaptedParam.docs.description = param.doc;
      method.parameters.push(adaptedParam);
      if (!paramMapping.has(opParam)) {
        paramMapping.set(opParam, new Array<go.MethodParameter>());
      }
      paramMapping.get(opParam)?.push(adaptedParam);

      if (adaptedParam.style !== 'required' && adaptedParam.style !== 'literal') {
        // add optional method param to the options param group
        if (!optionalGroup) {
          throw new AdapterError('InternalError', `optional parameter ${param.name} has no optional parameter group`, param.__raw?.node ?? NoTarget);
        }
        adaptedParam.group = optionalGroup;
        optionalGroup.params.push(adaptedParam);
      }
    }

    // client params aren't included in method.parameters so
    // look for them in the operation parameters.
    for (const opParam of allOpParams) {
      if (opParam.onClient) {
        const adaptedParam = this.adaptMethodParameter(opParam, method.httpMethod);
        adaptedParam.docs.summary = opParam.summary;
        adaptedParam.docs.description = opParam.doc;
        method.parameters.unshift(adaptedParam);
        if (!paramMapping.has(opParam)) {
          paramMapping.set(opParam, new Array<go.MethodParameter>());
        }
        paramMapping.get(opParam)?.push(adaptedParam);

        // if the adapted client param is a literal then don't add it to
        // the array of client params as it's not a formal parameter.
        if (go.isLiteralParameter(adaptedParam)) {
          continue;
        }

        // we must check via param name and not reference equality. this is because a client param
        // can be used in multiple ways. e.g. a client param "apiVersion" that's used as a path param
        // in one method and a query param in another.
        if (!method.client.parameters.find((v: go.ClientParameter) => {
          return v.name === adaptedParam.name;
        })) {
          method.client.parameters.push(adaptedParam);
        }
      }
    }

    return paramMapping;
  }

  private adaptContentType(contentTypeStr: string): 'binary' | 'JSON' | 'Text' | 'XML' {
    // we only recognize/support JSON, text, and XML content types, so assume anything else is binary
    // NOTE: we check XML before text in order to support text/xml
    let contentType: go.BodyFormat = 'binary';
    if (contentTypeStr.match(/json/i)) {
      contentType = 'JSON';
    } else if (contentTypeStr.match(/xml/i)) {
      contentType = 'XML';
    } else if (contentTypeStr.match(/text/i)) {
      contentType = 'Text';
    }
    return contentType;
  }

  private adaptMethodParameter(param: tcgc.SdkBodyParameter | tcgc.SdkHeaderParameter | tcgc.SdkPathParameter | tcgc.SdkQueryParameter | tcgc.SdkCookieParameter, verb: go.HTTPMethod): go.MethodParameter {
    if (param.isApiVersionParam && param.clientDefaultValue) {
      // we emit the api version param inline as a literal, never as a param.
      // the ClientOptions.APIVersion setting is used to change the version.
      const paramType = new go.Literal(new go.String(), param.clientDefaultValue);
      switch (param.kind) {
        case 'header':
          return new go.HeaderScalarParameter(param.name, param.serializedName, paramType, 'literal', true, 'method');
        case 'path':
          return new go.PathScalarParameter(param.name, param.serializedName, true, paramType, 'literal', true, 'method');
        case 'query':
          return new go.QueryScalarParameter(param.name, param.serializedName, true, paramType, 'literal', true, 'method');
        default:
          throw new AdapterError('UnsupportedTsp', `unsupported API version param kind ${param.kind}`, param.__raw?.node ?? NoTarget);
      }
    }

    let location: go.ParameterLocation = 'method';
    const getClientParamsKey = function (param: tcgc.SdkBodyParameter | tcgc.SdkHeaderParameter | tcgc.SdkPathParameter | tcgc.SdkQueryParameter | tcgc.SdkCookieParameter): string {
      // include the param kind in the key name as a client param can be used
      // in different places across methods (path/query)
      return `${param.name}-${param.kind}`;
    };
    if (param.onClient) {
      // check if we've already adapted this client parameter
      const clientParam = this.clientParams.get(getClientParamsKey(param));
      if (clientParam) {
        return clientParam;
      }
      location = 'client';
    }

    let adaptedParam: go.MethodParameter;
    let paramStyle = this.adaptParameterStyle(param);
    if (param.kind === 'body' && (verb === 'patch' || verb === 'put')) {
      paramStyle = 'required';
    }
    const paramName = getEscapedReservedName(ensureNameCase(param.name, paramStyle === 'required'), 'Param');
    const byVal = isTypePassedByValue(param.type);

    if (param.kind === 'body') {
      // TODO: form data? (non-multipart)
      if (param.defaultContentType.match(/multipart/i)) {
        adaptedParam = new go.MultipartFormBodyParameter(paramName, this.ta.getWireType(param.type, false, true), paramStyle, byVal);
      } else {
        const contentType = this.adaptContentType(param.defaultContentType);
        let bodyType = this.ta.getWireType(param.type, false, true);
        if (contentType === 'binary') {
          // tcgc models binary params as 'bytes' but we want an io.ReadSeekCloser
          bodyType = this.ta.getReadSeekCloser(param.type.kind === 'array');
        }
        adaptedParam = new go.BodyParameter(paramName, contentType, `"${param.defaultContentType}"`, bodyType, paramStyle, byVal);
      }
    } else if (param.kind === 'header') {
      if (param.collectionFormat) {
        if (param.collectionFormat === 'multi' || param.collectionFormat === 'form') {
          throw new AdapterError('InternalError', `unexpected collection format ${param.collectionFormat} for HeaderCollectionParameter`, param.__raw?.node ?? NoTarget);
        }
        // TODO: is hard-coded false for element type by value correct?
        const type = this.ta.getWireType(param.type, true, false);
        if (type.kind !== 'slice') {
          throw new AdapterError('InternalError', `unexpected type ${go.getTypeDeclaration(type)} for HeaderCollectionParameter ${param.name}`, param.__raw?.node ?? NoTarget);
        }
        adaptedParam = new go.HeaderCollectionParameter(paramName, param.serializedName, type, param.collectionFormat === 'simple' ? 'csv' : param.collectionFormat, paramStyle, byVal, location);
      } else {
        adaptedParam = new go.HeaderScalarParameter(paramName, param.serializedName, this.adaptHeaderScalarType(param.type, true), paramStyle, byVal, location);
      }
    } else if (param.kind === 'path') {
      adaptedParam = new go.PathScalarParameter(paramName, param.serializedName, !param.allowReserved, this.adaptPathScalarParameterType(param.type), paramStyle, byVal, location);
    } else if (param.kind === 'cookie') {
      // TODO: currently we don't have Azure service using cookie parameter. need to add support if needed in the future.
      throw new AdapterError('UnsupportedTsp', 'unsupported parameter type cookie', param.__raw?.node ?? NoTarget);
    } else {
      if (param.collectionFormat) {
        const type = this.ta.getWireType(param.type, true, false);
        if (type.kind !== 'slice') {
          throw new AdapterError('InternalError', `unexpected type ${go.getTypeDeclaration(type)} for QueryCollectionParameter ${param.name}`, param.__raw?.node ?? NoTarget);
        }
        // TODO: unencoded query param
        adaptedParam = new go.QueryCollectionParameter(paramName, param.serializedName, true, type, param.collectionFormat === 'simple' ? 'csv' : (param.collectionFormat === 'form' ? 'multi' : param.collectionFormat), paramStyle, byVal, location);
      } else {
        // TODO: unencoded query param
        adaptedParam = new go.QueryScalarParameter(paramName, param.serializedName, true, this.adaptQueryScalarParameterType(param.type), paramStyle, byVal, location);
      }
    }

    if (adaptedParam.location === 'client') {
      // track client parameter for later use
      this.clientParams.set(getClientParamsKey(param), adaptedParam);
    }

    return adaptedParam;
  }

  private getMethodNameForDocComment(method: go.MethodType): string {
    let methodName: string;
    switch (method.kind) {
      case 'lroMethod':
      case 'lroPageableMethod':
        methodName = `Begin${method.name}`;
        break;
      case 'method':
        methodName = method.name;
        break;
      case 'pageableMethod':
        methodName = `New${method.name}Pager`;
        break;
    }
    return `${method.client.name}.${methodName}`;
  }

  private adaptResponseEnvelope(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, method: go.MethodType): go.ResponseEnvelope {
    // TODO: add Envelope suffix if name collides with existing type
    let prefix = method.client.name;
    if (this.opts['single-client']) {
      prefix = '';
    }
    let respEnvName = `${prefix}${method.name}Response`;
    if (sdkMethod.access === 'internal') {
      respEnvName = uncapitalize(respEnvName);
    }
    const respEnv = new go.ResponseEnvelope(respEnvName, {summary: createResponseEnvelopeDescription(respEnvName, this.getMethodNameForDocComment(method))}, method);
    this.ta.codeModel.responseEnvelopes.push(respEnv);

    // add any headers
    const addedHeaders = new Set<string>();
    for (const httpResp of sdkMethod.operation.responses) {
      for (const httpHeader of httpResp.headers) {
        if (addedHeaders.has(httpHeader.serializedName)) {
          continue;
        } else if (go.isLROMethod(method) && httpHeader.serializedName.match(/Azure-AsyncOperation|Location|Operation-Location|Retry-After/i)) {
          // we omit the LRO polling headers as they aren't useful on the response envelope
          continue;
        }

        // TODO: x-ms-header-collection-prefix
        const headerResp = new go.HeaderScalarResponse(ensureNameCase(httpHeader.serializedName), this.adaptHeaderScalarType(httpHeader.type, false), httpHeader.serializedName, isTypePassedByValue(httpHeader.type));
        headerResp.docs.summary = httpHeader.summary;
        headerResp.docs.description = httpHeader.doc;
        respEnv.headers.push(headerResp);
        addedHeaders.add(httpHeader.serializedName);
      }
    }

    let sdkResponseType = sdkMethod.response.type;

    // since HEAD requests don't return a type, we must check this before checking sdkResponseType
    if (method.httpMethod === 'head' && this.opts['head-as-boolean'] === true) {
      respEnv.result = new go.HeadAsBooleanResult('Success');
      respEnv.result.docs.summary = 'Success indicates if the operation succeeded or failed.';
    }

    if (!sdkResponseType) {
      // method doesn't return a type, so we're done
      return respEnv;
    }

    if (sdkResponseType.kind === 'nullable') {
      // unwrap the nullable type, this will only happen for operations with two responses and one of them does not have a body
      sdkResponseType = sdkResponseType.type;
    }

    // for paged methods, tcgc models the method response type as an Array<T>.
    // however, we want the synthesized paged response envelope as that's what Go returns.
    if (sdkMethod.kind === 'paging') {
      // grab the paged response envelope type from the first response
      sdkResponseType = values(sdkMethod.operation.responses).first()!.type!;
    }

    // we have a response type, determine the content type
    let contentType: go.BodyFormat = 'binary';
    if (sdkMethod.kind === 'lro') {
      // we can't grovel through the operation responses for LROs as some of them
      // return only headers, thus have no content type. while it's highly likely
      // to only ever be JSON, this will be broken for LROs that return text/plain
      // or a binary response. the former seems unlikely, the latter though...??
      // TODO: https://github.com/Azure/typespec-azure/issues/535
      contentType = 'JSON';
    } else {
      let foundResp = false;
      for (const httpResp of sdkMethod.operation.responses) {
        if (!httpResp.type || !httpResp.defaultContentType || httpResp.type.kind !== sdkResponseType.kind) {
          continue;
        }
        contentType = this.adaptContentType(httpResp.defaultContentType);
        foundResp = true;
        break;
      }
      if (!foundResp) {
        throw new AdapterError('InternalError', `didn't find HTTP response for kind ${sdkResponseType.kind} in method ${method.name}`, sdkResponseType.__raw?.node ?? NoTarget);
      }
    }

    if (contentType === 'binary') {
      respEnv.result = new go.BinaryResult('Body');
      respEnv.result.docs.summary = 'Body contains the streaming response.';
      return respEnv;
    } else if (sdkResponseType.kind === 'model') {
      let modelType: go.Model | go.PolymorphicModel | undefined;
      const modelName = ensureNameCase(sdkResponseType.name).toUpperCase();
      for (const model of this.ta.codeModel.models) {
        if (model.name.toUpperCase() === modelName) {
          modelType = model;
          break;
        }
      }
      if (!modelType) {
        throw new AdapterError('InternalError', `didn't find model type name ${sdkResponseType.name} for response envelope ${respEnv.name}`, sdkResponseType.__raw?.node ?? NoTarget);
      }
      if (modelType.kind === 'polymorphicModel') {
        respEnv.result = new go.PolymorphicResult(modelType.interface);
      } else {
        if (contentType !== 'JSON' && contentType !== 'XML') {
          throw new AdapterError('InternalError', `unexpected content type ${contentType} for model ${modelType.name}`, NoTarget);
        }
        respEnv.result = new go.ModelResult(modelType, contentType);
      }
      respEnv.result.docs.summary = sdkResponseType.summary;
      respEnv.result.docs.description = sdkResponseType.doc;
    } else {
      const resultType = this.ta.getWireType(sdkResponseType, false, false);
      if (go.isMonomorphicResultType(resultType)) {
        respEnv.result = new go.MonomorphicResult(this.recursiveTypeName(sdkResponseType, false), contentType, resultType, isTypePassedByValue(sdkResponseType));
      } else {
        throw new AdapterError('InternalError', `invalid monomorphic result type ${resultType.kind}`, sdkResponseType.__raw?.node ?? NoTarget);
      }
    }

    return respEnv;
  }

  /**
   * creates the monomorphic response field name based on its type.
   * 
   * for unknown, use Interface or RawJSON if setting is enabled
   * for basic type, map of basic type, map of UDTs, enum, use Value
   * for array of basic type, array of UDTs, use xxxArray
   * 
   * @param type the type for which to create a name
   * @param fromArray indicates if there was recursion from a parent array
   * @returns the name
   */
  private recursiveTypeName(type: tcgc.SdkType, fromArray: boolean): string {
    if (!fromArray) {
      switch (type.kind) {
        case 'array':
          return `${this.recursiveTypeName(type.valueType, true)}Array`;
        case 'nullable':
          return this.recursiveTypeName(type.type, false);
        case 'unknown':
          return this.ta.codeModel.options.rawJSONAsBytes ? 'RawJSON' : 'Interface';
        default:
          return 'Value';
      }
    }

    switch (type.kind) {
      case 'array':
        return `${this.recursiveTypeName(type.valueType, true)}Array`;
      case 'boolean':
        return 'Bool';
      case 'bytes':
        return 'ByteArray';
      case 'enum':
      case 'model':
        return ensureNameCase(type.name);
      case 'utcDateTime':
      case 'offsetDateTime':
        return 'Time';
      case 'decimal':
      case 'decimal128':
        return 'Float64';
      case 'dict':
        return `MapOf${this.recursiveTypeName(type.valueType, fromArray)}`;
      case 'float32':
      case 'float64':
      case 'int16':
      case 'int32':
      case 'int64':
      case 'int8':
        return capitalize(type.kind);
      case 'nullable':
        return this.recursiveTypeName(type.type, fromArray);
      case 'duration':
      case 'string':
      case 'url':
        return 'String';
      case 'unknown':
        return this.ta.codeModel.options.rawJSONAsBytes ? 'RawJSON' : 'Interface';
      default:
        throw new Error(`unhandled monomorphic response type kind ${type.kind}`);
    }
  }

  private adaptParameterGroup(paramGroup: go.ParameterGroup): go.Struct {
    const structType = new go.Struct(paramGroup.groupName);
    structType.docs = paramGroup.docs;
    for (const param of paramGroup.params) {
      if (param.style === 'literal') {
        continue;
      }
      let byValue = param.style === 'required' || (param.location === 'client' && go.isClientSideDefault(param.style));
      // if the param isn't required, check if it should be passed by value or not.
      // optional params that are implicitly nil-able shouldn't be pointer-to-type.
      if (!byValue) {
        byValue = param.byValue;
      }
      const field = new go.StructField(param.name, param.type, byValue);
      field.docs = param.docs;
      structType.fields.push(field);
    }
    return structType;
  }

  private adaptHeaderScalarType(sdkType: tcgc.SdkType, forParam: boolean): go.HeaderScalarType {
    // for header params, we never pass the element type by pointer
    const type = this.ta.getWireType(sdkType, forParam, false);
    if (go.isHeaderScalarType(type)) {
      return type;
    }
    throw new AdapterError('InternalError', `unexpected header scalar parameter type ${sdkType.kind}`, sdkType.__raw?.node ?? NoTarget);
  }

  private adaptPathScalarParameterType(sdkType: tcgc.SdkType): go.PathScalarParameterType {
    const type = this.ta.getWireType(sdkType, false, false);
    if (go.isPathScalarParameterType(type)) {
      return type;
    }
    throw new AdapterError('InternalError', `unexpected path scalar parameter type ${sdkType.kind}`, sdkType.__raw?.node ?? NoTarget);
  }

  private adaptQueryScalarParameterType(sdkType: tcgc.SdkType): go.QueryScalarParameterType {
    const type = this.ta.getWireType(sdkType, false, false);
    if (go.isQueryScalarParameterType(type)) {
      return type;
    }
    throw new AdapterError('InternalError', `unexpected query scalar parameter type ${sdkType.kind}`, sdkType.__raw?.node ?? NoTarget);
  }

  private adaptParameterStyle(param: tcgc.SdkBodyParameter | tcgc.SdkEndpointParameter | tcgc.SdkHeaderParameter | tcgc.SdkMethodParameter | tcgc.SdkPathParameter | tcgc.SdkQueryParameter | tcgc.SdkCookieParameter): go.ParameterStyle {
    // NOTE: must check for constant type first as it will also set clientDefaultValue
    if (param.type.kind === 'constant') {
      if (param.optional) {
        return 'flag';
      }
      return 'literal';
    } else if (param.clientDefaultValue) {
      const adaptedType = this.ta.getWireType(param.type, false, false);
      if (!go.isLiteralValueType(adaptedType)) {
        throw new AdapterError('InternalError', `unexpected client side default type ${go.getTypeDeclaration(adaptedType)} for parameter ${param.name}`, param.__raw?.node ?? NoTarget);
      }
      return new go.ClientSideDefault(new go.Literal(adaptedType, param.clientDefaultValue));
    } else if (param.optional) {
      return 'optional';
    } else {
      return 'required';
    }
  }

  private adaptHttpOperationExamples(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, method: go.MethodType, paramMapping: Map<tcgc.SdkHttpParameter, Array<go.MethodParameter>>) {
    if (sdkMethod.operation.examples) {
      for (const example of sdkMethod.operation.examples) {
        const goExample = new go.MethodExample(example.name, {summary: example.doc}, example.filePath);
        for (const param of example.parameters) {
          if (param.parameter.isApiVersionParam && param.parameter.clientDefaultValue) {
            // skip the api-version param as it's not a formal parameter
            continue;
          }
          const goParams = paramMapping.get(param.parameter);
          if (!goParams) {
            throw new AdapterError('InternalError', `can not find go param for example param ${param.parameter.name}`, NoTarget);
          }
          if (goParams.length > 1) {
            // spread case
            for (const goParam of goParams) {
              const propertyValue = (<tcgc.SdkModelExampleValue>param.value).value[(<go.PartialBodyParameter>goParam).serializedName];
              const paramExample = new go.ParameterExample(goParam, this.adaptExampleType(propertyValue, goParam?.type));
              if (goParam.group) {
                goExample.optionalParamsGroup.push(paramExample);
              } else {
                goExample.parameters.push(paramExample);
              }
            }
          } else {
            const paramExample = new go.ParameterExample(goParams[0], this.adaptExampleType(param.value, goParams[0]?.type));
            if (goParams[0]?.group) {
              goExample.optionalParamsGroup.push(paramExample);
            } else {
              goExample.parameters.push(paramExample);
            }
          }
        }
        // only handle 200 response
        const response = example.responses.find((v) => { return v.statusCode === 200; });
        if (response) {
          goExample.responseEnvelope = new go.ResponseEnvelopeExample(method.responseEnvelope);
          for (const header of response.headers) {
            const goHeader = method.responseEnvelope.headers.find(h => h.headerName === header.header.serializedName);
            if (!goHeader) {
              throw new AdapterError('InternalError', `can not find go header for example header ${header.header.serializedName}`, NoTarget);
            }
            goExample.responseEnvelope.headers.push(new go.ResponseHeaderExample(goHeader, this.adaptExampleType(header.value, goHeader.type)));
          }
          // there are some problems with LROs at present which can cause the result
          // to be undefined even though the operation returns a response.
          // TODO: https://github.com/Azure/typespec-azure/issues/1688
          if (response.bodyValue && method.responseEnvelope.result) {
            switch (method.responseEnvelope.result.kind) {
              case 'anyResult':
                goExample.responseEnvelope.result = this.adaptExampleType(response.bodyValue, new go.Any());
                break;
              case 'binaryResult':
                goExample.responseEnvelope.result = this.adaptExampleType(response.bodyValue, new go.Scalar('byte', false));
                break;
              case 'modelResult':
                goExample.responseEnvelope.result = this.adaptExampleType(response.bodyValue, method.responseEnvelope.result.modelType);
                break;
              case 'monomorphicResult':
                goExample.responseEnvelope.result = this.adaptExampleType(response.bodyValue, method.responseEnvelope.result.monomorphicType);
                break;
              case 'polymorphicResult':
                goExample.responseEnvelope.result = this.adaptExampleType(response.bodyValue, method.responseEnvelope.result.interface);
                break;
            }
          }
        }
        method.examples.push(goExample);
      }
    }
  }

  private adaptExampleType(exampleType: tcgc.SdkExampleValue, goType: go.WireType): go.ExampleType {
    switch (exampleType.kind) {
      case 'string':
        switch (goType.kind) {
          case 'constant':
          case 'encodedBytes':
          case 'literal':
          case 'string':
          case 'time':
            return new go.StringExample(exampleType.value, goType);
          case 'qualifiedType':
            return new go.QualifiedExample(goType, exampleType.value);
        }
        break;
      case 'number':
        switch (goType.kind) {
          case 'constant':
          case 'literal':
          case 'scalar':
          case 'time':
            return new go.NumberExample(exampleType.value, goType);
        }
        break;
      case 'boolean':
        switch (goType.kind) {
          case 'constant':
          case 'literal':
          case 'scalar':
            return new go.BooleanExample(exampleType.value, goType);
        }
        break;
      case 'null':
        return new go.NullExample(goType);
      case 'unknown':
        if (goType.kind === 'any') {
          return new go.AnyExample(exampleType.value);
        }
        break;
      case 'array':
        if (goType.kind === 'slice') {
          const ret = new go.ArrayExample(goType);
          for (const v of exampleType.value) {
            ret.value.push(this.adaptExampleType(v, goType.elementType));
          }
          return ret;
        }
        break;
      case 'dict':
        if (goType.kind === 'map') {
          const ret = new go.DictionaryExample(goType);
          for (const [k, v] of Object.entries(exampleType.value)) {
            ret.value[k] = this.adaptExampleType(v, goType.valueType);
          }
          return ret;
        }
        break;
      case 'union':
        throw new AdapterError('UnsupportedTsp', 'unsupported example type kind union', NoTarget);
      case 'model':
        if (goType.kind === 'interface' || goType.kind === 'model' || goType.kind === 'polymorphicModel') {
          let concreteType: go.Model | go.PolymorphicModel | undefined;
          if (goType.kind === 'interface') {
            /* eslint-disable-next-line @typescript-eslint/no-unsafe-member-access */
            concreteType = goType.possibleTypes.find(t => t.discriminatorValue?.literal === exampleType.type.discriminatorValue || t.discriminatorValue?.literal.value === exampleType.type.discriminatorValue)!;
            if (concreteType === undefined) {
              // can't find the sub type of a discriminated type, fallback to the base type
              concreteType = goType.rootType;
            }
          } else {
            concreteType = goType;
          }
          if (concreteType === undefined) {
            throw new AdapterError('InternalError', `can not find concrete type for example type ${exampleType.type.name}`, NoTarget);
          }
          const ret = new go.StructExample(concreteType);
          for (const [k, v] of Object.entries(exampleType.value)) {
            const field = concreteType.fields.find(f => f.serializedName === k)!;
            ret.value[field.name] = this.adaptExampleType(v, field.type);
          }
          if (exampleType.additionalPropertiesValue) {
            ret.additionalProperties = {};
            for (const [k, v] of Object.entries(exampleType.additionalPropertiesValue)) {
              const filed = concreteType.fields.find(f => f.annotations.isAdditionalProperties)!;
              if (filed.type.kind === 'map') {
                ret.additionalProperties[k] = this.adaptExampleType(v, filed.type.valueType);
              } else {
                throw new AdapterError('InternalError', `additional properties field type should be map type`, NoTarget);
              }
            }
          }
          return ret;
        }
        break;
    }
    throw new AdapterError('InternalError', `can not map go type into example type ${exampleType.kind}`, NoTarget);
  }
}

interface HttpStatusCodeRange {
  start: number;
  end: number;
}

function isHttpStatusCodeRange(statusCode: HttpStatusCodeRange | number): statusCode is HttpStatusCodeRange {
  return (<HttpStatusCodeRange>statusCode).start !== undefined;
}
