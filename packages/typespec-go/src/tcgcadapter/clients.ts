/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { capitalize, createOptionsTypeDescription, createResponseEnvelopeDescription, ensureNameCase, getEscapedReservedName, uncapitalize } from '../../../naming.go/src/naming.js';
import { GoEmitterOptions } from '../lib.js';
import { isTypePassedByValue, typeAdapter } from './types.js';
import * as go from '../../../codemodel.go/src/gocodemodel.js';
import * as tcgc from '@azure-tools/typespec-client-generator-core';

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
      if (sdkClient.initialization) {
        this.recursiveAdaptClient(sdkClient);
      }
    }
  }

  private recursiveAdaptClient(sdkClient: tcgc.SdkClientType<tcgc.SdkHttpOperation>, parent?: go.Client): go.Client {
    let clientName = sdkClient.name;
    if (!clientName.match(/Client$/)) {
      clientName += 'Client';
    }

    // depending on the hierarchy, we could end up attempting to adapt the same client.
    // if it's already been adapted, return it.
    let goClient = this.ta.codeModel.clients.find((v: go.Client, i: number, o: Array<go.Client>) => {
      return o[i].name === clientName;
    });
    if (goClient) {
      return goClient;
    }

    let description: string;
    if (sdkClient.description) {
      description = `${clientName} - ${sdkClient.description}`;
    } else {
      description = `${clientName} contains the methods for the ${sdkClient.nameSpace} namespace.`;
    }

    goClient = new go.Client(clientName, description, `New${clientName}`);
    goClient.parent = parent;
    goClient.host = sdkClient.endpoint;
    goClient.complexHostParams = sdkClient.hasParameterizedEndpoint;

    if (sdkClient.initialization) {
      for (const param of sdkClient.initialization.properties) {
        if (param.kind === 'credential') {
          // skip this for now as we don't generate client constructors
          continue;
        } else if (param.kind === 'endpoint' && param.type.kind === 'constant') {
          // this is the param for the fixed host, don't create a param for it
          goClient.host = <string>param.type.value;
          if (!this.ta.codeModel.host) {
            this.ta.codeModel.host = goClient.host;
          } else if (this.ta.codeModel.host !== goClient.host) {
            throw new Error(`client ${goClient.name} has a conflicting host ${goClient.host}`);
          }
          continue;
        } else if (param.kind === 'method') {
          // some client params, notably (only?) api-version, can be explicitly
          // defined in the operation signature:
          // e.g. op withQueryApiVersion(@query("api-version") apiVersion: string)
          // these get propagated to sdkMethod.operation.parameters thus they
          // will be adapted in adaptMethodParameters()
          continue;
        }

        const paramType = this.ta.getPossibleType(param.type, true, false);
        if (!go.isConstantType(paramType) && !go.isPrimitiveType(paramType)) {
          throw new Error(`unexpected URI parameter type ${go.getTypeDeclaration(paramType)}`);
        }
        // TODO: follow up with tcgc if serializedName should actually be optional
        const uriParam = new go.URIParameter(param.nameInClient, param.serializedName ? param.serializedName : param.nameInClient, paramType,
          this.adaptParameterType(param), isTypePassedByValue(param.type) || !param.optional, 'client');
        goClient.hostParams.push(uriParam);
      }
    } else if (parent) {
      // this is a sub-client. it will share the client/host params of the parent.
      // NOTE: we must propagate parant params before a potential recursive call
      // to create a child client that will need to inherit our client params.
      goClient.hostParams = parent.hostParams;
      goClient.parameters = parent.parameters;
    }

    for (const sdkMethod of sdkClient.methods) {
      if (sdkMethod.kind === 'clientaccessor') {
        const subClient = this.recursiveAdaptClient(sdkMethod.response, goClient);
        goClient.clientAccessors.push(new go.ClientAccessor(`New${subClient.name}`, subClient));
      } else {
        this.adaptMethod(sdkMethod, goClient);
      }
    }

    this.ta.codeModel.clients.push(goClient);
    return goClient;
  }

  private adaptMethod(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, goClient: go.Client) {
    let method: go.Method | go.LROMethod | go.LROPageableMethod | go.PageableMethod;
    const naming = new go.MethodNaming(getEscapedReservedName(uncapitalize(ensureNameCase(sdkMethod.name)), 'Operation'), ensureNameCase(`${sdkMethod.name}CreateRequest`, true),
      ensureNameCase(`${sdkMethod.name}HandleResponse`, true));
  
    const getStatusCodes = function(httpOp: tcgc.SdkHttpOperation): Array<number> {
      const statusCodes = new Array<number>();
      for (const statusCode of Object.keys(httpOp.responses)) {
        statusCodes.push(parseInt(statusCode));
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
    } else {
      throw new Error(`method kind ${sdkMethod.kind} NYI`);
    }
  
    method.description = sdkMethod.description;
    goClient.methods.push(method);
    this.populateMethod(sdkMethod, method);
  }

  private populateMethod(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, method: go.Method | go.NextPageMethod) {
    if (go.isMethod(method)) {
      let prefix = method.client.name;
      if (this.opts['single-client']) {
        prefix = '';
      }
      let optionalParamsGroupName = `${prefix}${method.name}Options`;
      if (sdkMethod.access === 'internal') {
        optionalParamsGroupName = uncapitalize(optionalParamsGroupName);
      }
      let optsGroupName = 'options';
      // if there's an existing parameter with the name options then pick something else
      for (const param of sdkMethod.parameters) {
        if (param.nameInClient === optsGroupName) {
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
  
    this.adaptMethodParameters(sdkMethod, method);

    // we must do this after adapting method params as it can add optional params
    this.ta.codeModel.paramGroups.push(this.adaptParameterGroup(method.optionalParamsGroup));
  }

  private adaptMethodParameters(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, method: go.Method | go.NextPageMethod) {
    let optionalGroup: go.ParameterGroup | undefined;
    if (go.isMethod(method)) {
      // NextPageMethods don't have optional params
      optionalGroup = method.optionalParamsGroup;
      if (go.isLROMethod(method)) {
        optionalGroup.params.push(new go.ResumeTokenParameter());
      }
    }

    for (const param of sdkMethod.operation.parameters) {
      const adaptedParam = this.adaptMethodParameter(param, optionalGroup);
      method.parameters.push(adaptedParam);
      // we must check via param name and not reference equality. this is because a client param
      // can be used in multiple ways. e.g. a client param "apiVersion" that's used as a path param
      // in one method and a query param in another.
      if (adaptedParam.location === 'client' && !method.client.parameters.find((v: go.Parameter, i: number, o: Array<go.Parameter>) => {
        return v.name === adaptedParam.name;
      })) {
        method.client.parameters.push(adaptedParam);
      }
    }

    // we add the body param after any required params. this way,
    // if the body param is required it shows up last in the list.
    if (sdkMethod.operation.bodyParams.length > 1) {
      throw new Error('multipart body NYI');
    } else if (sdkMethod.operation.bodyParams.length === 1) {
      const bodyParam = sdkMethod.operation.bodyParams[0];
      method.parameters.push(this.adaptMethodParameter(bodyParam, optionalGroup));
    }
  }

  private adaptContentType(contentTypeStr: string): 'binary' | 'JSON' | 'Text' | 'XML' {
    // we only recognize/support JSON, text, and XML content types, so assume anything else is binary
    // TODO: this seems goofy but maybe it's ok?
    let contentType: go.BodyFormat = 'binary';
    if (contentTypeStr.match(/json/i)) {
      contentType = 'JSON';
    } else if (contentTypeStr.match(/text/i)) {
      contentType = 'Text';
    } else if (contentTypeStr.match(/xml/i)) {
      contentType = 'XML';
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
          return new go.HeaderParameter(param.nameInClient, param.serializedName, paramType, 'literal', true, 'method');
        case 'path':
          return new go.PathParameter(param.nameInClient, param.serializedName, true, paramType, 'literal', true, 'method');
        case 'query':
          return new go.QueryParameter(param.nameInClient, param.serializedName, true, paramType, 'literal', true, 'method');
        default:
          throw new Error(`unhandled param kind ${param.kind} for API version param`);
      }
    }

    let location: go.ParameterLocation = 'method';
    const getClientParamsKey = function(param: tcgc.SdkBodyParameter | tcgc.SdkHeaderParameter | tcgc.SdkPathParameter | tcgc.SdkQueryParameter): string {
      // include the param kind in the key name as a client param can be used
      // in different places across methods (path/query)
      return `${param.nameInClient}-${param.kind}`;
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
    const paramName = getEscapedReservedName(ensureNameCase(param.nameInClient, true), 'Param');
    const paramType = this.adaptParameterType(param);
    const byVal = isTypePassedByValue(param.type);

    if (param.kind === 'body') {
      // TODO: form data? (non-multipart)
      const contentType = this.adaptContentType(param.defaultContentType);
      let bodyType = this.ta.getPossibleType(param.type, false, true);
      if (contentType === 'binary') {
        // tcgc models binary params as 'bytes' but we want an io.ReadSeekCloser
        bodyType = this.ta.getReadSeekCloser(param.type.kind === 'array');
      }
      adaptedParam = new go.BodyParameter(paramName, contentType, `"${param.defaultContentType}"`, bodyType, paramType, byVal);
    } else if (param.kind === 'header') {
      if (param.collectionFormat) {
        if (param.collectionFormat === 'multi') {
          throw new Error('unexpected collection format multi for HeaderCollectionParameter');
        }
        // TODO: is hard-coded false for element type by value correct?
        const type = this.ta.getPossibleType(param.type, true, false);
        if (!go.isSliceType(type)) {
          throw new Error(`unexpected type ${go.getTypeDeclaration(type)} for HeaderCollectionParameter ${param.nameInClient}`);
        }
        adaptedParam = new go.HeaderCollectionParameter(paramName, param.serializedName, type, param.collectionFormat, paramType, byVal, location);
      } else {
        adaptedParam = new go.HeaderParameter(paramName, param.serializedName, this.adaptHeaderType(param.type, true), paramType, byVal, location);
      }
    } else if (param.kind === 'path') {
      adaptedParam = new go.PathParameter(paramName, param.serializedName, param.urlEncode, this.adaptPathParameterType(param.type), paramType, byVal, location);
    } else {
      if (param.collectionFormat) {
        const type = this.ta.getPossibleType(param.type, true, false);
        if (!go.isSliceType(type)) {
          throw new Error(`unexpected type ${go.getTypeDeclaration(type)} for QueryCollectionParameter ${param.nameInClient}`);
        }
        // TODO: unencoded query param
        adaptedParam = new go.QueryCollectionParameter(paramName, param.serializedName, true, type, param.collectionFormat, paramType, byVal, location);
      } else {
        // TODO: unencoded query param
        adaptedParam = new go.QueryParameter(paramName, param.serializedName, true, this.adaptQueryParameterType(param.type), paramType, byVal, location);
      }
    }

    adaptedParam.description = param.description;

    if (adaptedParam.location === 'client') {
      // track client parameter for later use
      this.clientParams.set(getClientParamsKey(param), adaptedParam);
    } else if (paramType !== 'required' && paramType !== 'literal') {
      // add optional method param to the options param group
      if (!optionalGroup) {
        throw new Error(`optional parameter ${param.nameInClient} has no optional parameter group`);
      }
      adaptedParam.group = optionalGroup;
      optionalGroup.params.push(adaptedParam);
    }

    return adaptedParam;
  }

  private getMethodNameForDocComment(method: go.Method): string {
    return `${method.client.name}.${go.isPageableMethod(method) && !go.isLROMethod(method) ? `New${method.name}Pager` : method.name}`;
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
    const respEnv = new go.ResponseEnvelope(respEnvName, createResponseEnvelopeDescription(respEnvName, this. getMethodNameForDocComment(method)), method);
    this.ta.codeModel.responseEnvelopes.push(respEnv);

    // add any headers
    const addedHeaders = new Set<string>();
    for (const httpResp of Object.values(sdkMethod.operation.responses)) { 
      for (const httpHeader of httpResp.headers) {
        if (addedHeaders.has(httpHeader.serializedName)) {
          continue;
        } else if (go.isLROMethod(method) && httpHeader.serializedName.match(/Azure-AsyncOperation|Location|Operation-Location/i)) {
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

    if (!sdkResponseType) {
      // method doesn't return a type, so we're done
      return respEnv;
    }

    // workaround until https://github.com/Azure/typespec-azure/issues/124 is fixed
    if (sdkResponseType.kind === 'model' && (sdkResponseType.name === 'OperationStatus' || sdkResponseType.name === 'ResourceOperationStatus')) {
      return respEnv;
    }
    // end workaround

    // for paged methods, tcgc models the method response type as an Array<T>.
    // however, we want the synthesized paged response envelope as that's what Go returns.
    if (sdkMethod.kind === 'paging') {
      // grab the paged response envelope type from the first response
      sdkResponseType = Object.values(sdkMethod.operation.responses)[0].type!;
    }

    // we have a response type, determine the content type
    let contentType: go.BodyFormat = 'binary';
    if (sdkMethod.kind === 'lro') {
      // we can't grovel through the operation responses for LROs as some of them
      // return only headers, thus have no content type. while it's highly likely
      // to only ever be JSON, this will be broken for LROs that return text/plain
      // or a binary response. the former seems unlikely, the latter though...??
      // TODO: follow up with tcgc team about this
      contentType = 'JSON';
    } else {
      let foundResp = false;
      for (const httpResp of Object.values(sdkMethod.operation.responses)) {
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
      if (param.paramType === 'literal') {
        continue;
      }
      let byValue = param.paramType === 'required' || (param.location === 'client' && go.isClientSideDefault(param.paramType));
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
    if (go.isMapType(type) || go.isInterfaceType(type) || go.isModelType(type) || go.isPolymorphicType(type) || go.isSliceType(type)  || go.isQualifiedType(type)) {
      throw new Error(`unexpected path parameter type ${sdkType.kind}`);
    }
    return type;
  }
  
  private adaptQueryParameterType(sdkType: tcgc.SdkType): go.QueryParameterType {
    const type = this.ta.getPossibleType(sdkType, false, false);
    if (go.isMapType(type) || go.isInterfaceType(type) || go.isModelType(type) || go.isPolymorphicType(type) || go.isSliceType(type)  || go.isQualifiedType(type)) {
      throw new Error(`unexpected query parameter type ${sdkType.kind}`);
    } else if (go.isSliceType(type)) {
      type.elementTypeByValue = true;
    }
    return type;
  }
  
  private adaptParameterType(param: tcgc.SdkBodyParameter | tcgc.SdkEndpointParameter | tcgc.SdkHeaderParameter | tcgc.SdkPathParameter | tcgc.SdkQueryParameter): go.ParameterType {
    // NOTE: must check for constant type first as it will also set clientDefaultValue
    if (param.type.kind === 'constant') {
      if (param.optional) {
        return 'flag';
      }
      return 'literal';
    } else if (param.clientDefaultValue) {
      const adaptedType = this.ta.getPossibleType(param.type, false, false);
      if (!go.isLiteralValueType(adaptedType)) {
        throw new Error(`unsupported client side default type ${go.getTypeDeclaration(adaptedType)} for parameter ${param.nameInClient}`);
      }
      return new go.ClientSideDefault(new go.LiteralValue(adaptedType, param.clientDefaultValue));
    } else if (param.optional) {
      return 'optional';
    } else {
      return 'required';
    }
  }
}
