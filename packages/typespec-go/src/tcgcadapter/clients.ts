/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { capitalize, createOptionsTypeDescription, createResponseEnvelopeDescription, ensureNameCase, getEscapedReservedName, uncapitalize } from '../../../naming.go/src/naming.js';
import { GoEmitterOptions } from '../lib.js';
import { isTypePassedByValue, typeAdapter } from './types.js';
import * as go from '../../../codemodel.go/src/gocodemodel.js';
import * as tcgc from '@azure-tools/typespec-client-generator-core';
import { values } from '@azure-tools/linq';

// used to convert SDK clients and their methods to Go code model types
export class clientAdapter {
  private ta: typeAdapter;
  private opts: GoEmitterOptions;

  // track all of the client and parameter group params across all operations
  // as not every option might contain them, and parameter groups can be shared
  // across multiple operations
  private clientParams: Map<string, go.Parameter>;
  private paramGroups: Map<string, go.ParameterGroup>;

  constructor(ta: typeAdapter, opts: GoEmitterOptions) {
    this.ta = ta;
    this.opts = opts;
    this.clientParams = new Map<string, go.Parameter>();
    this.paramGroups = new Map<string, go.ParameterGroup>();
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
      return o[i].clientName === clientName;
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
        if (param.kind === 'credential' || param.isApiVersionParam) {
          // skip these for now as we don't generate client constructors
          continue;
        } else if (param.kind === 'endpoint' && param.type.kind === 'constant') {
          // this is the param for the fixed host, don't create a param for it
          goClient.host = <string>param.type.value;
          if (!this.ta.codeModel.host) {
            this.ta.codeModel.host = goClient.host;
          } else if (this.ta.codeModel.host !== goClient.host) {
            throw new Error(`client ${goClient.clientName} has a conflicting host ${goClient.host}`);
          }
          continue;
        } else if (param.kind === 'method') {
          throw new Error('client method params NYI');
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
        goClient.clientAccessors.push(new go.ClientAccessor(subClient));
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

    if (sdkMethod.kind === 'basic') {
      method = new go.Method(methodName, goClient, sdkMethod.operation.path, sdkMethod.operation.verb, getStatusCodes(sdkMethod.operation), naming);
    } else if (sdkMethod.kind === 'paging') {
      method = new go.PageableMethod(methodName, goClient, sdkMethod.operation.path, sdkMethod.operation.verb, getStatusCodes(sdkMethod.operation), naming);
      if (sdkMethod.nextLinkPath) {
        // TODO: handle nested next link
        (<go.PageableMethod>method).nextLinkName = capitalize(ensureNameCase(sdkMethod.nextLinkPath));
      }
    } else {
      throw new Error(`method kind ${sdkMethod.kind} NYI`);
    }
  
    method.description = sdkMethod.description;
    goClient.methods.push(method);
    this.populateMethod(sdkMethod, method);
  }

  private populateMethod(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, method: go.Method | go.NextPageMethod) {
    if (go.isMethod(method)) {
      let prefix = method.client.clientName;
      if (this.opts['single-client']) {
        prefix = '';
      }
      let optionalParamsGroupName = `${prefix}${method.methodName}Options`;
      if (sdkMethod.access === 'internal') {
        optionalParamsGroupName = uncapitalize(optionalParamsGroupName);
      }
      // TODO: ensure param name is unique
      method.optionalParamsGroup = new go.ParameterGroup('options', optionalParamsGroupName, false, 'method');
      method.optionalParamsGroup.description = createOptionsTypeDescription(optionalParamsGroupName, this.getMethodNameForDocComment(method));
      method.responseEnvelope = this.adaptResponseEnvelope(sdkMethod, method);
    } else {
      throw new Error('NYI');
    }
  
    this.adaptMethodParameters(sdkMethod, method);
  
    /*for (const apiver of values(op.apiVersions)) {
      method.apiVersions.push(apiver.version);
    }*/

    // we must do this after adapting method params as it can add optional params
    this.ta.codeModel.paramGroups.push(this.adaptParameterGroup(method.optionalParamsGroup));
  }

  private adaptMethodParameters(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, method: go.Method | go.NextPageMethod) {
    let optionalGroup: go.ParameterGroup | undefined;
    if (go.isMethod(method)) {
      // NextPageMethods don't have optional params
      optionalGroup = method.optionalParamsGroup;
    }
  
    for (const param of sdkMethod.operation.parameters) {
      const adaptedParam = this.adaptMethodParameter(param, optionalGroup);
      method.parameters.push(adaptedParam);
      if (adaptedParam.location === 'client' && !method.client.parameters.includes(adaptedParam)) {
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

  private adaptMethodParameter(param: tcgc.SdkBodyParameter | tcgc.SdkHeaderParameter | tcgc.SdkPathParameter | tcgc.SdkQueryParameter, optionalGroup?: go.ParameterGroup): go.Parameter {
    let location: go.ParameterLocation = 'method';
    if (param.onClient) {
      // check if we've already adapted this client parameter
      // TODO: grouped client params
      const clientParam = this.clientParams.get(param.nameInClient);
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
      // TODO: hard-coded format type
      adaptedParam = new go.BodyParameter(paramName, 'JSON', param.defaultContentType, this.ta.getPossibleType(param.type, false, true), paramType, byVal);
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
      this.clientParams.set(param.nameInClient, adaptedParam);
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
    return `${method.client.clientName}.${go.isPageableMethod(method) && !go.isLROMethod(method) ? `New${method.methodName}Pager` : method.methodName}`;
  }

  private adaptResponseEnvelope(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, method: go.Method): go.ResponseEnvelope {
    // TODO: add Envelope suffix if name collides with existing type
    let prefix = method.client.clientName;
    if (this.opts['single-client']) {
      prefix = '';
    }
    let respEnvName = `${prefix}${method.methodName}Response`;
    if (sdkMethod.access === 'internal') {
      respEnvName = uncapitalize(respEnvName);
    }
    const respEnv = new go.ResponseEnvelope(respEnvName, createResponseEnvelopeDescription(respEnvName, this. getMethodNameForDocComment(method)), method);
    this.ta.codeModel.responseEnvelopes.push(respEnv);
  
    const bodyResponses = new Array<tcgc.SdkType>();

    // add any headers
    for (const httpResp of Object.values(sdkMethod.operation.responses)) {
      const addedHeaders = new Set<string>();
      if (httpResp.type && !bodyResponses.includes(httpResp.type)) {
        bodyResponses.push(httpResp.type);
      }
  
      for (const httpHeader of httpResp.headers) {
        if (addedHeaders.has(httpHeader.serializedName)) {
          continue;
        }
        // TODO: x-ms-header-collection-prefix
        const headerResp = new go.HeaderResponse(ensureNameCase(httpHeader.serializedName), this.adaptHeaderType(httpHeader.type, false), httpHeader.serializedName, isTypePassedByValue(httpHeader.type));
        headerResp.description = httpHeader.description;
        respEnv.headers.push(headerResp);
        addedHeaders.add(httpHeader.serializedName);
      }
    }
  
    if (bodyResponses.length === 0) {
      return respEnv;
    }
  
    if (bodyResponses.length > 1) {
      throw new Error('any response NYI');
    } else if (bodyResponses[0].kind === 'model') {
      let modelType: go.ModelType | undefined;
      const modelName = ensureNameCase(bodyResponses[0].name).toUpperCase();
      for (const model of this.ta.codeModel.models) {
        if (model.name.toUpperCase() === modelName) {
          modelType = model;
          break;
        }
      }
      if (!modelType) {
        throw new Error(`didn't find type name ${bodyResponses[0].name} for response envelope ${respEnv.name}`);
      }
      if (go.isPolymorphicType(modelType)) {
        respEnv.result = new go.PolymorphicResult(modelType.interface);
      } else {
        // TODO: hard-coded JSON
        respEnv.result = new go.ModelResult(modelType, 'JSON');
      }
      respEnv.result.description = bodyResponses[0].description;
    } else {
      const resultType = this.ta.getPossibleType(bodyResponses[0], false, false);
      if (go.isInterfaceType(resultType) || go.isLiteralValue(resultType) || go.isModelType(resultType) || go.isPolymorphicType(resultType) || go.isQualifiedType(resultType)) {
        throw new Error(`invalid monomorphic result type ${go.getTypeDeclaration(resultType)}`);
      }
      // TODO: hard-coded JSON
      respEnv.result = new go.MonomorphicResult('Value', 'JSON', resultType, isTypePassedByValue(bodyResponses[0]));
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
      const field = new go.StructField(param.paramName, param.type, byValue);
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
