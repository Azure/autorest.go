/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { capitalize, ensureNameCase, getEscapedReservedName, uncapitalize } from '../../../naming.go/src/naming.js';
import { isTypePassedByValue, typeAdapter } from './types.js';
import * as go from '../../../codemodel.go/src/gocodemodel.js';
import * as tcgc from '@azure-tools/typespec-client-generator-core';
import { values } from '@azure-tools/linq';

// used to convert SDK clients and their methods to Go code model types
export class clientAdapter {
  private ta: typeAdapter;

  // track all of the client and parameter group params across all operations
  // as not every option might contain them, and parameter groups can be shared
  // across multiple operations
  private clientParams: Map<string, go.Parameter>;
  private paramGroups: Map<string, go.ParameterGroup>;

  constructor(ta: typeAdapter) {
    this.ta = ta;
    this.clientParams = new Map<string, go.Parameter>();
    this.paramGroups = new Map<string, go.ParameterGroup>();
  }

  // converts all clients and their methods to Go code model types.
  // this includes parameter groups/options types and response envelopes.
  adaptClients(sdkPackage: tcgc.SdkPackage<tcgc.SdkHttpOperation>) {
    for (const sdkClient of sdkPackage.clients) {
      if (sdkClient.initialization && values(sdkClient.methods).all((each: tcgc.SdkMethod<tcgc.SdkHttpOperation>) => { return each.kind === 'clientaccessor'; })) {
        // this is a hierarchical client with only client accessors so skip
        // it as we don't currently support hierarchical clients.
        continue;
      }
      let clientName = sdkClient.name;
      if (!clientName.match(/Client$/)) {
        clientName += 'Client';
      }
      const goClient = new go.Client(clientName, sdkPackage.name, `New${clientName}`);
      goClient.host = sdkClient.endpoint;
      if (!this.ta.codeModel.host) {
        this.ta.codeModel.host = goClient.host;
      } else if (this.ta.codeModel.host !== goClient.host) {
        throw new Error(`client ${goClient.clientName} has a conflicting host`);
      }
      for (const sdkMethod of sdkClient.methods) {
        if (sdkMethod.kind === 'clientaccessor') {
          // used for hierarchical clients which isn't currently supported
          continue;
        }
        this.adaptMethod(sdkMethod, goClient);
      }
  
      this.ta.codeModel.clients.push(goClient);
    }
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
  
    const methodName = capitalize(ensureNameCase(sdkMethod.name));

    if (sdkMethod.kind === 'basic') {
      method = new go.Method(methodName, goClient, sdkMethod.operation.path, sdkMethod.operation.verb, getStatusCodes(sdkMethod.operation), naming);
    } else if (sdkMethod.kind === 'paging') {
      method = new go.PageableMethod(methodName, goClient, sdkMethod.operation.path, sdkMethod.operation.verb, getStatusCodes(sdkMethod.operation), naming);
      if (sdkMethod.nextLinkLogicalPath) {
        // TODO: this assumes that nextLink is the first element. and what does it mean to have more than one entry?
        (<go.PageableMethod>method).nextLinkName = capitalize(ensureNameCase(sdkMethod.nextLinkLogicalPath[0]));
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
      const optionalParamsGroupName = `${method.client.clientName}${method.methodName}Options`;
      // TODO: ensure param name is unique
      method.optionalParamsGroup = new go.ParameterGroup('options', optionalParamsGroupName, false, 'method');
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

    if (sdkMethod.operation.bodyParams.length > 1) {
      throw new Error('multipart body NYI');
    } else if (sdkMethod.operation.bodyParams.length === 1) {
      const bodyParam = sdkMethod.operation.bodyParams[0];
      method.parameters.push(this.adaptMethodParameter(bodyParam, optionalGroup));
    }
  
    for (const param of sdkMethod.operation.parameters) {
      const adaptedParam = this.adaptMethodParameter(param, optionalGroup);
      method.parameters.push(adaptedParam);
      if (adaptedParam.location === 'client' && !method.client.parameters.includes(adaptedParam)) {
        method.client.parameters.push(adaptedParam);
      }
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

  private adaptResponseEnvelope(sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>, method: go.Method): go.ResponseEnvelope {
    // TODO: add Envelope suffix if name collides with existing type
    const respEnvName = `${method.client.clientName}${method.methodName}Response`;
    // TODO: proper name for paged methods in doc comment
    const respEnvDesc = `${respEnvName} contains the response from method ${method.client.clientName}.${method.methodName}.`;
    const respEnv = new go.ResponseEnvelope(respEnvName, respEnvDesc, method);
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
      const modelName = ensureNameCase(bodyResponses[0].name);
      for (const model of this.ta.codeModel.models) {
        if (model.name === modelName) {
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
  
  private adaptParameterType(param: tcgc.SdkBodyParameter | tcgc.SdkHeaderParameter | tcgc.SdkPathParameter | tcgc.SdkQueryParameter): go.ParameterType {
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
