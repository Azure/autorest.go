/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { values } from '@azure-tools/linq';
import * as tcgc from '@azure-tools/typespec-client-generator-core';
import * as go from '../../../codemodel.go/src/index.js';
import { capitalize, createOptionsTypeDescription, createResponseEnvelopeDescription, ensureNameCase, getEscapedReservedName, uncapitalize } from '../../../naming.go/src/naming.js';
import { GoEmitterOptions } from '../lib.js';
import { isTypePassedByValue, typeAdapter } from './types.js';

// used to convert SDK clients and their methods to Go code model types
export class clientAdapter {
  private ta: typeAdapter;
  private opts: GoEmitterOptions;

  // track all of the client and parameter group params across all operations
  // as not every option might contain them, and parameter groups can be shared
  // across multiple operations
  private clientParams: Map<string, go.Parameter>;

  constructor(ta: typeAdapter, opts: GoEmitterOptions) {
    this.ta = ta;
    this.opts = opts;
    this.clientParams = new Map<string, go.Parameter>();
  }

  // converts all clients and their methods to Go code model types.
  // this includes parameter groups/options types and response envelopes.
  adaptClients(sdkPackage: tcgc.SdkPackage<tcgc.SdkHttpOperation>) {
    if (this.opts['single-client'] && sdkPackage.clients.length > 1) {
      throw new Error('single-client cannot be enabled when there are multiple clients');
    }
    for (const sdkClient of sdkPackage.clients) {
      // start with instantiable clients and recursively work down
      if (sdkClient.initialization.access === 'public') {
        this.recursiveAdaptClient(sdkClient);
      }
    }
  }

  private recursiveAdaptClient(sdkClient: tcgc.SdkClientType<tcgc.SdkHttpOperation>, parent?: go.Client): go.Client | undefined {
    if (sdkClient.methods.length === 0) {
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

    if (!clientName.match(/Client$/)) {
      clientName += 'Client';
    }

    let description: string;
    if (sdkClient.description) {
      description = `${clientName} - ${sdkClient.description}`;
    } else {
      // strip clientName's "Client" suffix
      const groupName = clientName.substring(0, clientName.length - 6);
      description = `${clientName} contains the methods for the ${groupName} group.`;
    }

    const goClient = new go.Client(clientName, description, go.newClientOptions(this.ta.codeModel.type, clientName));
    goClient.parent = parent;

    // anything other than public means non-instantiable client
    if (sdkClient.initialization.access === 'public') {
      for (const param of sdkClient.initialization.properties) {
        if (param.kind === 'credential') {
          // skip this for now as we don't generate client constructors
          continue;
        } else if (param.kind === 'endpoint') {
          // for multiple endpoint, we only generate the first one
          let paramType;
          if (param.type.kind === 'endpoint') {
            paramType = param.type;
          } else {
            paramType = param.type.values[0];
          }
          // this will either be a fixed or templated host
          // don't set the fixed host for ARM as it isn't used
          if (this.ta.codeModel.type !== 'azure-arm') {
            goClient.host = paramType.serverUrl;
          }
          if (paramType.templateArguments.length === 0) {
            // this is the param for the fixed host, don't create a param for it
            if (!this.ta.codeModel.host) {
              this.ta.codeModel.host = goClient.host;
            } else if (this.ta.codeModel.host !== goClient.host) {
              throw new Error(`client ${goClient.name} has a conflicting host ${goClient.host}`);
            }
          } else {
            if (this.ta.codeModel.type !== 'azure-arm') {
              goClient.templatedHost = true;
              for (const templateArg of paramType.templateArguments) {
                goClient.parameters.push(this.adaptURIParam(templateArg));
              }
            }
          }
          continue;
        } else if (param.kind === 'method') {
          // some client params, notably api-version, can be explicitly
          // defined in the operation signature:
          // e.g. op withQueryApiVersion(@query("api-version") apiVersion: string)
          // these get propagated to sdkMethod.operation.parameters thus they
          // will be adapted in adaptMethodParameters()
          continue;
        }

        goClient.parameters.push(this.adaptURIParam(param));
      }
    } else if (parent) {
      // this is a sub-client. it will share the client/host params of the parent.
      // NOTE: we must propagate parant params before a potential recursive call
      // to create a child client that will need to inherit our client params.
      goClient.templatedHost = parent.templatedHost;
      goClient.host = parent.host;
      // make a copy of the client params. this is to prevent
      // client method params from being shared across clients
      // as not all client method params might be uniform.
      goClient.parameters = new Array<go.Parameter>(...parent.parameters);
    } else {
      throw new Error(`uninstantiable client ${sdkClient.name} has no parent`);
    }

    for (const sdkMethod of sdkClient.methods) {
      if (sdkMethod.kind === 'clientaccessor') {
        const subClient = this.recursiveAdaptClient(sdkMethod.response, goClient);
        if (subClient) {
          goClient.clientAccessors.push(new go.ClientAccessor(`New${subClient.name}`, subClient));
        }
      } else {
        this.adaptMethod(sdkMethod, goClient);
      }
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

  private adaptURIParam(sdkParam: tcgc.SdkEndpointParameter | tcgc.SdkPathParameter): go.URIParameter {
    const paramType = this.ta.getPossibleType(sdkParam.type, true, false);
    if (!go.isConstantType(paramType) && !go.isPrimitiveType(paramType)) {
      throw new Error(`unexpected URI parameter type ${go.getTypeDeclaration(paramType)}`);
    }
    // TODO: follow up with tcgc if serializedName should actually be optional
    return new go.URIParameter(sdkParam.name, sdkParam.serializedName ? sdkParam.serializedName : sdkParam.name, paramType,
      this.adaptParameterKind(sdkParam), isTypePassedByValue(sdkParam.type) || !sdkParam.optional, 'client');
  }

  private adaptMethod(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, goClient: go.Client) {
    let method: go.Method | go.LROMethod | go.LROPageableMethod | go.PageableMethod;
    const naming = new go.MethodNaming(getEscapedReservedName(uncapitalize(ensureNameCase(sdkMethod.name)), 'Operation'), ensureNameCase(`${sdkMethod.name}CreateRequest`, true),
      ensureNameCase(`${sdkMethod.name}HandleResponse`, true));

    const getStatusCodes = function (httpOp: tcgc.SdkHttpOperation): Array<number> {
      const statusCodes = new Array<number>();
      for (const statusCode of httpOp.responses.keys()) {
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
      method = new go.PageableMethod(methodName, goClient, sdkMethod.operation.path, sdkMethod.operation.verb, statusCodes, naming);
      if (sdkMethod.nextLinkPath) {
        // TODO: handle nested next link
        (<go.PageableMethod>method).nextLinkName = capitalize(ensureNameCase(sdkMethod.nextLinkPath));
      }
    } else if (sdkMethod.kind === 'lro') {
      method = new go.LROMethod(methodName, goClient, sdkMethod.operation.path, sdkMethod.operation.verb, statusCodes, naming);
      const lroOptions = this.hasDecorator('Azure.Core.@useFinalStateVia', sdkMethod.decorators);
      if (lroOptions) {
        (<go.LROMethod>method).finalStateVia = lroOptions['finalState'];
      }
    } else {
      throw new Error(`method kind ${sdkMethod.kind} NYI`);
    }

    method.description = sdkMethod.description;
    goClient.methods.push(method);
    this.populateMethod(sdkMethod, method);
  }

  private hasDecorator(name: string, decorators: Array<tcgc.DecoratorInfo>): Record<string, any> | undefined {
    for (const decorator of decorators) {
      if (decorator.name === name) {
        return decorator.arguments;
      }
    }
    return undefined;
  }

  private populateMethod(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, method: go.Method | go.NextPageMethod) {
    if (go.isMethod(method)) {
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
      // if there's an existing parameter with the name options then pick something else
      for (const param of sdkMethod.parameters) {
        if (param.name === optsGroupName) {
          optsGroupName = 'opts';
          break;
        }
      }
      method.optionalParamsGroup = new go.ParameterGroup(optsGroupName, optionalParamsGroupName, false, 'method');
      method.optionalParamsGroup.description = createOptionsTypeDescription(optionalParamsGroupName, this.getMethodNameForDocComment(method));
      method.responseEnvelope = this.adaptResponseEnvelope(sdkMethod, method);
    } else {
      throw new Error('NYI');
    }

    // find the api version param to use for the doc comment.
    // we can't use sdkMethod.apiVersions as that includes all
    // of the api versions supported by the service.
    for (const opParam of sdkMethod.operation.parameters) {
      if (opParam.isApiVersionParam && opParam.clientDefaultValue) {
        method.apiVersions.push(opParam.clientDefaultValue);
        break;
      }
    }

    const paramMapping = this.adaptMethodParameters(sdkMethod, method);

    // we must do this after adapting method params as it can add optional params
    this.ta.codeModel.paramGroups.push(this.adaptParameterGroup(method.optionalParamsGroup));

    this.adaptHttpOperationExamples(sdkMethod, method, paramMapping);
  }

  private adaptMethodParameters(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, method: go.Method | go.NextPageMethod): Map<tcgc.SdkHttpParameter, Array<go.Parameter>> {
    const paramMapping = new Map<tcgc.SdkHttpParameter, Array<go.Parameter>>();

    let optionalGroup: go.ParameterGroup | undefined;
    if (go.isMethod(method)) {
      // NextPageMethods don't have optional params
      optionalGroup = method.optionalParamsGroup;
      if (go.isLROMethod(method)) {
        optionalGroup.params.push(new go.ResumeTokenParameter());
      }
    }

    // stuff all of the operation parameters into one array for easy traversal
    type OperationParamType = tcgc.SdkBodyParameter | tcgc.SdkHeaderParameter | tcgc.SdkPathParameter | tcgc.SdkQueryParameter;
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
      const opParam = values(allOpParams).where((opParam: OperationParamType) => {
        return values(opParam.correspondingMethodParams).where((methodParam: tcgc.SdkModelPropertyType) => {
          return methodParam.name === param.name;
        }).any();
      }).first();

      // special handling for constants that used in path, this will not be in operation parameters since it has been resolved in the url
      if (!opParam && param.type.kind === 'constant') {
        continue;
      }

      if (!opParam) {
        throw new Error(`didn't find operation parameter for method ${sdkMethod.name} parameter ${param.name}`);
      }

      let adaptedParam: go.Parameter;
      if (opParam.kind === 'body' && opParam.type.kind === 'model' && opParam.type.kind !== param.type.kind) {
        const paramKind = this.adaptParameterKind(param);
        const byVal = isTypePassedByValue(param.type);
        const contentType = this.adaptContentType(opParam.defaultContentType);
        switch (contentType) {
          case 'JSON':
          case 'XML': {
            // find the corresponding field within the model param so we can get the serialized name
            let serializedName: string | undefined;
            for (const property of opParam.type.properties) {
              if (property.name === param.name) {
                serializedName = (<tcgc.SdkBodyModelPropertyType>property).serializedName;
                break;
              }
            }
            if (!serializedName) {
              throw new Error(`didn't find body model property for spread parameter ${param.name}`);
            }
            adaptedParam = new go.PartialBodyParameter(param.name, serializedName, contentType, this.ta.getPossibleType(param.type, true, true), paramKind, byVal);
            break;
          }
          case 'binary':
            if (opParam.defaultContentType.match(/multipart/i)) {
              adaptedParam = new go.MultipartFormBodyParameter(param.name, this.ta.getReadSeekCloser(false), paramKind, byVal);
            } else {
              adaptedParam = new go.BodyParameter(param.name, contentType, `"${opParam.defaultContentType}"`, this.ta.getReadSeekCloser(false), paramKind, byVal);
            }
            break;
          default:
            throw new Error(`unhandled spread param content type ${contentType}`);
        }
      } else {
        adaptedParam = this.adaptMethodParameter(opParam, optionalGroup);
      }

      adaptedParam.description = param.description;
      method.parameters.push(adaptedParam);
      if (!paramMapping.has(opParam)) {
        paramMapping.set(opParam, new Array<go.Parameter>());
      }
      paramMapping.get(opParam)?.push(adaptedParam);

      if (adaptedParam.location === 'client') {
        // we must check via param name and not reference equality. this is because a client param
        // can be used in multiple ways. e.g. a client param "apiVersion" that's used as a path param
        // in one method and a query param in another.
        if (!method.client.parameters.find((v: go.Parameter, i: number, o: Array<go.Parameter>) => {
          return v.name === adaptedParam.name;
        })) {
          method.client.parameters.push(adaptedParam);
        }
      } else if (adaptedParam.kind !== 'required' && adaptedParam.kind !== 'literal') {
        // add optional method param to the options param group
        if (!optionalGroup) {
          throw new Error(`optional parameter ${param.name} has no optional parameter group`);
        }
        adaptedParam.group = optionalGroup;
        optionalGroup.params.push(adaptedParam);
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

  private adaptMethodParameter(param: tcgc.SdkBodyParameter | tcgc.SdkHeaderParameter | tcgc.SdkPathParameter | tcgc.SdkQueryParameter, optionalGroup?: go.ParameterGroup): go.Parameter {
    if (param.isApiVersionParam && param.clientDefaultValue) {
      // we emit the api version param inline as a literal, never as a param.
      // the ClientOptions.APIVersion setting is used to change the version.
      const paramType = new go.LiteralValue(new go.PrimitiveType('string'), param.clientDefaultValue);
      switch (param.kind) {
        case 'header':
          return new go.HeaderParameter(param.name, param.serializedName, paramType, 'literal', true, 'method');
        case 'path':
          return new go.PathParameter(param.name, param.serializedName, true, paramType, 'literal', true, 'method');
        case 'query':
          return new go.QueryParameter(param.name, param.serializedName, true, paramType, 'literal', true, 'method');
        default:
          throw new Error(`unhandled param kind ${param.kind} for API version param`);
      }
    }

    let location: go.ParameterLocation = 'method';
    const getClientParamsKey = function (param: tcgc.SdkBodyParameter | tcgc.SdkHeaderParameter | tcgc.SdkPathParameter | tcgc.SdkQueryParameter): string {
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

    let adaptedParam: go.Parameter;
    const paramName = getEscapedReservedName(ensureNameCase(param.name, true), 'Param');
    const paramKind = this.adaptParameterKind(param);
    const byVal = isTypePassedByValue(param.type);

    if (param.kind === 'body') {
      // TODO: form data? (non-multipart)
      if (param.defaultContentType.match(/multipart/i)) {
        adaptedParam = new go.MultipartFormBodyParameter(paramName, this.ta.getPossibleType(param.type, false, true), paramKind, byVal);
      } else {
        const contentType = this.adaptContentType(param.defaultContentType);
        let bodyType = this.ta.getPossibleType(param.type, false, true);
        if (contentType === 'binary') {
          // tcgc models binary params as 'bytes' but we want an io.ReadSeekCloser
          bodyType = this.ta.getReadSeekCloser(param.type.kind === 'array');
        }
        adaptedParam = new go.BodyParameter(paramName, contentType, `"${param.defaultContentType}"`, bodyType, paramKind, byVal);
      }
    } else if (param.kind === 'header') {
      if (param.collectionFormat) {
        if (param.collectionFormat === 'multi' || param.collectionFormat === 'form') {
          throw new Error('unexpected collection format multi for HeaderCollectionParameter');
        }
        // TODO: is hard-coded false for element type by value correct?
        const type = this.ta.getPossibleType(param.type, true, false);
        if (!go.isSliceType(type)) {
          throw new Error(`unexpected type ${go.getTypeDeclaration(type)} for HeaderCollectionParameter ${param.name}`);
        }
        adaptedParam = new go.HeaderCollectionParameter(paramName, param.serializedName, type, param.collectionFormat === 'simple' ? 'csv' : param.collectionFormat, paramKind, byVal, location);
      } else {
        adaptedParam = new go.HeaderParameter(paramName, param.serializedName, this.adaptHeaderType(param.type, true), paramKind, byVal, location);
      }
    } else if (param.kind === 'path') {
      adaptedParam = new go.PathParameter(paramName, param.serializedName, param.urlEncode || param.allowReserved, this.adaptPathParameterType(param.type), paramKind, byVal, location);
    } else {
      if (param.collectionFormat) {
        const type = this.ta.getPossibleType(param.type, true, false);
        if (!go.isSliceType(type)) {
          throw new Error(`unexpected type ${go.getTypeDeclaration(type)} for QueryCollectionParameter ${param.name}`);
        }
        // TODO: unencoded query param
        adaptedParam = new go.QueryCollectionParameter(paramName, param.serializedName, true, type, param.collectionFormat === 'simple' ? 'csv' : (param.collectionFormat === 'form' ? 'multi' : param.collectionFormat), paramKind, byVal, location);
      } else {
        // TODO: unencoded query param
        adaptedParam = new go.QueryParameter(paramName, param.serializedName, true, this.adaptQueryParameterType(param.type), paramKind, byVal, location);
      }
    }

    if (adaptedParam.location === 'client') {
      // track client parameter for later use
      this.clientParams.set(getClientParamsKey(param), adaptedParam);
    }

    return adaptedParam;
  }

  private getMethodNameForDocComment(method: go.Method): string {
    let methodName = method.name;
    if (go.isLROMethod(method)) {
      methodName = `Begin${methodName}`;
    } else if (go.isPageableMethod(method)) {
      methodName = `New${methodName}Pager`;
    }
    return `${method.client.name}.${methodName}`;
  }

  private adaptResponseEnvelope(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, method: go.Method): go.ResponseEnvelope {
    // TODO: add Envelope suffix if name collides with existing type
    let prefix = method.client.name;
    if (this.opts['single-client']) {
      prefix = '';
    }
    let respEnvName = `${prefix}${method.name}Response`;
    if (sdkMethod.access === 'internal') {
      respEnvName = uncapitalize(respEnvName);
    }
    const respEnv = new go.ResponseEnvelope(respEnvName, createResponseEnvelopeDescription(respEnvName, this.getMethodNameForDocComment(method)), method);
    this.ta.codeModel.responseEnvelopes.push(respEnv);

    // add any headers
    const addedHeaders = new Set<string>();
    for (const httpResp of sdkMethod.operation.responses.values()) {
      for (const httpHeader of httpResp.headers) {
        if (addedHeaders.has(httpHeader.serializedName)) {
          continue;
        } else if (go.isLROMethod(method) && httpHeader.serializedName.match(/Azure-AsyncOperation|Location|Operation-Location|Retry-After/i)) {
          // we omit the LRO polling headers as they aren't useful on the response envelope
          continue;
        }

        // TODO: x-ms-header-collection-prefix
        const headerResp = new go.HeaderResponse(ensureNameCase(httpHeader.serializedName), this.adaptHeaderType(httpHeader.type, false), httpHeader.serializedName, isTypePassedByValue(httpHeader.type));
        headerResp.description = httpHeader.description;
        respEnv.headers.push(headerResp);
        addedHeaders.add(httpHeader.serializedName);
      }
    }

    let sdkResponseType = sdkMethod.response.type;

    // since HEAD requests don't return a type, we must check this before checking sdkResponseType
    if (method.httpMethod === 'head' && this.opts['head-as-boolean'] === true) {
      respEnv.result = new go.HeadAsBooleanResult('Success');
      respEnv.result.description = 'Success indicates if the operation succeeded or failed.';
    }

    if (!sdkResponseType) {
      // method doesn't return a type, so we're done
      return respEnv;
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
      for (const httpResp of sdkMethod.operation.responses.values()) {
        if (!httpResp.type || !httpResp.defaultContentType || httpResp.type.kind !== sdkResponseType.kind) {
          continue;
        }
        contentType = this.adaptContentType(httpResp.defaultContentType);
        foundResp = true;
        break;
      }
      if (!foundResp) {
        throw new Error(`didn't find HTTP response for kind ${sdkResponseType.kind} in method ${method.name}`);
      }
    }

    if (contentType === 'binary') {
      respEnv.result = new go.BinaryResult('Body', 'binary');
      respEnv.result.description = 'Body contains the streaming response.';
      return respEnv;
    } else if (sdkResponseType.kind === 'model') {
      let modelType: go.ModelType | undefined;
      const modelName = ensureNameCase(sdkResponseType.name).toUpperCase();
      for (const model of this.ta.codeModel.models) {
        if (model.name.toUpperCase() === modelName) {
          modelType = model;
          break;
        }
      }
      if (!modelType) {
        throw new Error(`didn't find model type name ${sdkResponseType.name} for response envelope ${respEnv.name}`);
      }
      if (go.isPolymorphicType(modelType)) {
        respEnv.result = new go.PolymorphicResult(modelType.interface);
      } else {
        if (contentType !== 'JSON' && contentType !== 'XML') {
          throw new Error(`unexpected content type ${contentType} for model ${modelType.name}`);
        }
        respEnv.result = new go.ModelResult(modelType, contentType);
      }
      respEnv.result.description = sdkResponseType.description;
    } else {
      const resultType = this.ta.getPossibleType(sdkResponseType, false, false);
      if (go.isInterfaceType(resultType) || go.isLiteralValue(resultType) || go.isModelType(resultType) || go.isPolymorphicType(resultType) || go.isQualifiedType(resultType)) {
        throw new Error(`invalid monomorphic result type ${go.getTypeDeclaration(resultType)}`);
      }
      respEnv.result = new go.MonomorphicResult('Value', contentType, resultType, isTypePassedByValue(sdkResponseType));
    }

    return respEnv;
  }

  private adaptParameterGroup(paramGroup: go.ParameterGroup): go.StructType {
    const structType = new go.StructType(paramGroup.groupName);
    structType.description = paramGroup.description;
    for (const param of paramGroup.params) {
      if (param.kind === 'literal') {
        continue;
      }
      let byValue = param.kind === 'required' || (param.location === 'client' && go.isClientSideDefault(param.kind));
      // if the param isn't required, check if it should be passed by value or not.
      // optional params that are implicitly nil-able shouldn't be pointer-to-type.
      if (!byValue) {
        byValue = param.byValue;
      }
      const field = new go.StructField(param.name, param.type, byValue);
      field.description = param.description;
      structType.fields.push(field);
    }
    return structType;
  }

  private adaptHeaderType(sdkType: tcgc.SdkType, forParam: boolean): go.HeaderType {
    // for header params, we never pass the element type by pointer
    const type = this.ta.getPossibleType(sdkType, forParam, false);
    if (go.isInterfaceType(type) || go.isMapType(type) || go.isModelType(type) || go.isPolymorphicType(type) || go.isSliceType(type) || go.isQualifiedType(type)) {
      throw new Error(`unexpected header parameter type ${sdkType.kind}`);
    }
    return type;
  }

  private adaptPathParameterType(sdkType: tcgc.SdkType): go.PathParameterType {
    const type = this.ta.getPossibleType(sdkType, false, false);
    if (go.isMapType(type) || go.isInterfaceType(type) || go.isModelType(type) || go.isPolymorphicType(type) || go.isSliceType(type) || go.isQualifiedType(type)) {
      throw new Error(`unexpected path parameter type ${sdkType.kind}`);
    }
    return type;
  }

  private adaptQueryParameterType(sdkType: tcgc.SdkType): go.QueryParameterType {
    const type = this.ta.getPossibleType(sdkType, false, false);
    if (go.isMapType(type) || go.isInterfaceType(type) || go.isModelType(type) || go.isPolymorphicType(type) || go.isSliceType(type) || go.isQualifiedType(type)) {
      throw new Error(`unexpected query parameter type ${sdkType.kind}`);
    } else if (go.isSliceType(type)) {
      type.elementTypeByValue = true;
    }
    return type;
  }

  private adaptParameterKind(param: tcgc.SdkBodyParameter | tcgc.SdkEndpointParameter | tcgc.SdkHeaderParameter | tcgc.SdkMethodParameter | tcgc.SdkPathParameter | tcgc.SdkQueryParameter): go.ParameterKind {
    // NOTE: must check for constant type first as it will also set clientDefaultValue
    if (param.type.kind === 'constant') {
      if (param.optional) {
        return 'flag';
      }
      return 'literal';
    } else if (param.clientDefaultValue) {
      const adaptedType = this.ta.getPossibleType(param.type, false, false);
      if (!go.isLiteralValueType(adaptedType)) {
        throw new Error(`unsupported client side default type ${go.getTypeDeclaration(adaptedType)} for parameter ${param.name}`);
      }
      return new go.ClientSideDefault(new go.LiteralValue(adaptedType, param.clientDefaultValue));
    } else if (param.optional) {
      return 'optional';
    } else {
      return 'required';
    }
  }

  private adaptHttpOperationExamples(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, method: go.Method, paramMapping: Map<tcgc.SdkHttpParameter, Array<go.Parameter>>) {
    if (sdkMethod.operation.examples) {
      for (const example of sdkMethod.operation.examples) {
        const goExample = new go.MethodExample(example.name, example.description, example.filePath);
        for (const param of example.parameters) {
          const goParams = paramMapping.get(param.parameter);
          if (!goParams) {
            throw new Error(`can not find go param for example param ${param.parameter.name}`);
          }
          if (goParams.length > 1) {
            // spread case
            for (const goParam of goParams) {
              const propertyValue = (<tcgc.SdkModelExample>param.value).value[(<go.PartialBodyParameter>goParam).serializedName];
              const paramExample = new go.ParameterExample(goParam, this.adaptExampleType(propertyValue, goParam?.type));
              if (goParam?.group) {
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
        const response = example.responses.get(200);
        if (response) {
          goExample.responseEnvelope = new go.ResponseEnvelopeExample(method.responseEnvelope);
          for (const header of response.headers) {
            const goHeader = method.responseEnvelope.headers.find(h => h.headerName === header.header.serializedName);
            if (!goHeader) {
              throw new Error(`can not find go header for example header ${header.header.serializedName}`);
            }
            goExample.responseEnvelope.headers.push(new go.ResponseHeaderExample(goHeader, this.adaptExampleType(header.value, goHeader.type)));
          }
          if (response.bodyValue) {
            if (go.isAnyResult(method.responseEnvelope.result!)) {
              goExample.responseEnvelope.result = this.adaptExampleType(response.bodyValue, new go.PrimitiveType('any'));
            } else if (go.isModelResult(method.responseEnvelope.result!)) {
              goExample.responseEnvelope.result = this.adaptExampleType(response.bodyValue, method.responseEnvelope.result.modelType);
            } else if (go.isBinaryResult(method.responseEnvelope.result!)) {
              goExample.responseEnvelope.result = this.adaptExampleType(response.bodyValue, new go.PrimitiveType('byte'));
            } else if (go.isMonomorphicResult(method.responseEnvelope.result!)) {
              goExample.responseEnvelope.result = this.adaptExampleType(response.bodyValue, method.responseEnvelope.result.monomorphicType);
            } else if (go.isPolymorphicResult(method.responseEnvelope.result!)) {
              goExample.responseEnvelope.result = this.adaptExampleType(response.bodyValue, method.responseEnvelope.result.interfaceType);
            }
          }
        }
        method.examples.push(goExample);
      }
    }
  }

  private adaptExampleType(exampleType: tcgc.SdkTypeExample, goType: go.PossibleType): go.ExampleType {
    switch (exampleType.kind) {
      case 'string':
        if (go.isConstantType(goType) || go.isBytesType(goType) || go.isLiteralValue(goType) || go.isTimeType(goType) || go.isPrimitiveType(goType)) {
          return new go.StringExample(exampleType.value, goType);
        }
        break;
      case 'number':
        if (go.isConstantType(goType) || go.isLiteralValue(goType) || go.isTimeType(goType) || go.isPrimitiveType(goType)) {
          return new go.NumberExample(exampleType.value, goType);
        }
        break;
      case 'boolean':
        if (go.isConstantType(goType) || go.isLiteralValue(goType) || go.isPrimitiveType(goType)) {
          return new go.BooleanExample(exampleType.value, goType);
        }
        break;
      case 'null':
        return new go.NullExample(goType);
      case 'any':
        if (go.isPrimitiveType(goType) && goType.typeName === 'any') {
          return new go.AnyExample(exampleType.value);
        }
        break;
      case 'array':
        if (go.isSliceType(goType)) {
          const ret = new go.ArrayExample(goType);
          for (const v of exampleType.value) {
            ret.value.push(this.adaptExampleType(v, goType.elementType));
          }
          return ret;
        }
        break;
      case 'dict':
        if (go.isMapType(goType)) {
          const ret = new go.DictionaryExample(goType);
          for (const [k, v] of Object.entries(exampleType.value)) {
            ret.value[k] = this.adaptExampleType(v, goType.valueType);
          }
          return ret;
        }
        break;
      case 'union':
        throw new Error('go could not support union for now');
      case 'model':
        if (go.isModelType(goType) || go.isInterfaceType(goType)) {
          let concreteType: go.ModelType | go.PolymorphicType;
          if (go.isInterfaceType(goType)) {
            concreteType = goType.possibleTypes.find(t => t.discriminatorValue?.literal === exampleType.type.discriminatorValue)!;
          } else {
            concreteType = goType;
          }
          const ret = new go.StructExample(concreteType);
          for (const [k, v] of Object.entries(exampleType.value)) {
            const field = concreteType.fields.find(f => f.serializedName === k)!;
            ret.value[field.name] = this.adaptExampleType(v, field.type);
          }
          if (exampleType.additionalPropertiesValue) {
            ret.additionalProperties = {};
            for (const [k, v] of Object.entries(exampleType.additionalPropertiesValue)) {
              ret.additionalProperties[k] = this.adaptExampleType(v, concreteType.fields.find(f => f.annotations.isAdditionalProperties)!.type!);
            }
          }
          return ret;
        }
        break;
    }
    throw new Error(`can not map go type into example type ${exampleType.kind}`);
  }
}

interface HttpStatusCodeRange {
  start: number;
  end: number;
}

function isHttpStatusCodeRange(statusCode: HttpStatusCodeRange | number): statusCode is HttpStatusCodeRange {
  return (<HttpStatusCodeRange>statusCode).start !== undefined;
}
