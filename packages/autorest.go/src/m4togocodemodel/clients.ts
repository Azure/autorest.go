/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

/* eslint-disable @typescript-eslint/no-unsafe-argument */
/* eslint-disable @typescript-eslint/no-unsafe-assignment */
/* eslint-disable @typescript-eslint/no-unsafe-member-access */

import * as m4 from '@autorest/codemodel';
import { KnownMediaType } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import { adaptXMLInfo } from './types.js';
import { adaptWireType, hasDescription } from './types.js';
import * as go from '../../../codemodel.go/src/index.js';
import * as helpers from '../transform/helpers.js';
import { OperationNaming } from '../transform/namer.js';

// track all of the client and parameter group params across all operations
// as not every option might contain them, and parameter groups can be shared
// across multiple operations
const clientParams = new Map<string, go.MethodParameter>();
const paramGroups = new Map<string, go.ParameterGroup>();

export function adaptClients(m4CodeModel: m4.CodeModel, codeModel: go.CodeModel) {
  for (const group of values(m4CodeModel.operationGroups)) {
    const client = adaptClient(codeModel.type, group);

    for (const op of values(group.operations)) {
      const httpPath = <string>op.requests![0].protocol.http!.path;
      const httpMethod = op.requests![0].protocol.http!.method;
      let method: go.MethodType;
      const naming = adaptMethodNaming(op);

      if (helpers.isLROOperation(op) && helpers.isPageableOperation(op)) {
        method = new go.LROPageableMethod(op.language.go!.name, client, httpPath, httpMethod, getStatusCodes(op), naming);
        method.finalStateVia = (op.extensions?.['x-ms-long-running-operation-options']?.['final-state-via']);
        method.nextLinkName = op.language.go!.paging.nextLinkName;
        if (op.language.go!.paging.nextLinkOperation) {
          // adapt the next link operation
          const nextPageMethod = new go.NextPageMethod(op.language.go!.paging.nextLinkOperation.language.go.name, client, httpPath, httpMethod, getStatusCodes(op.language.go!.paging.nextLinkOperation));
          populateMethod(op.language.go!.paging.nextLinkOperation, nextPageMethod, m4CodeModel, codeModel);
          method.nextPageMethod = nextPageMethod;
        }
      } else if (helpers.isLROOperation(op)) {
        method = new go.LROMethod(op.language.go!.name, client, httpPath, httpMethod, getStatusCodes(op), naming);
        method.finalStateVia = (op.extensions?.['x-ms-long-running-operation-options']?.['final-state-via']);
      } else if (helpers.isPageableOperation(op)) {
        if (op.language.go!.paging.isNextOp) {
          continue;
        }
        method = new go.PageableMethod(op.language.go!.name, client, httpPath, httpMethod, getStatusCodes(op), naming);
        method.nextLinkName = op.language.go!.paging.nextLinkName;
        if (op.language.go!.paging.nextLinkOperation) {
          // adapt the next link operation
          const nextPageMethod = adaptNextPageMethod(op, m4CodeModel, client, codeModel);
          method.nextPageMethod = nextPageMethod;
        }
      } else {
        method = new go.SyncMethod(op.language.go!.name, client, httpPath, httpMethod, getStatusCodes(op), naming);
      }

      populateMethod(op, method, m4CodeModel, codeModel);

      client.methods.push(method);
    }

    // if any client parameters were adapted, add them to the client
    if (group.language.go!.clientParams) {
      for (const param of <Array<m4.Parameter>>group.language.go!.clientParams) {
        const adaptedParam = clientParams.get(param.language.go!.name);
        if (!adaptedParam) {
          throw new Error(`missing adapted client parameter ${param.language.go!.name}`);
        }
        client.parameters.push(adaptedParam);
      }
    }

    if (codeModel.type === 'azure-arm') {
      // we don't need the scopes for ARM, it's handled by pipeline policy
      const ctor = new go.Constructor(`New${client.name}`, new go.TokenAuthentication([]));
      // add any modeled parameter first, which should only be the subscriptionID, then add TokenCredential
      for (const param of client.parameters) {
        ctor.parameters.push(param);
      }
      client.constructors.push(ctor);
    }

    codeModel.clients.push(client);
  }
}

// used to ensure unique instances of go.NextPageMethod since they can be shared across operations
// only adaptNextPageMethod should touch this!
const adaptedNextPageMethods = new Map<string, go.NextPageMethod>();

function adaptNextPageMethod(op: m4.Operation, m4CodeModel: m4.CodeModel, client: go.Client, codeModel: go.CodeModel): go.NextPageMethod {
  const nextPageMethodName = op.language.go!.paging.nextLinkOperation.language.go.name;
  let nextPageMethod = adaptedNextPageMethods.get(nextPageMethodName);
  if (!nextPageMethod) {
    const httpPath = <string>op.language.go!.paging.nextLinkOperation.requests![0].protocol.http!.path;
    const httpMethod = <go.HTTPMethod>op.language.go!.paging.nextLinkOperation.requests![0].protocol.http!.method;
    nextPageMethod = new go.NextPageMethod(nextPageMethodName, client, httpPath, httpMethod, getStatusCodes(op.language.go!.paging.nextLinkOperation));
    populateMethod(op.language.go!.paging.nextLinkOperation, nextPageMethod, m4CodeModel, codeModel);
    adaptedNextPageMethods.set(nextPageMethodName, nextPageMethod);
  }
  return nextPageMethod;
}

function populateMethod(op: m4.Operation, method: go.MethodType | go.NextPageMethod, m4CodeModel: m4.CodeModel, codeModel: go.CodeModel) {
  if (method.kind !== 'nextPageMethod') {
    if (hasDescription(op.language.go!)) {
      method.docs.description = op.language.go!.description;
    }

    let optionalParamsGroup = paramGroups.get(op.language.go!.optionalParamGroup.schema.language.go!.name);
    if (!optionalParamsGroup) {
      optionalParamsGroup = adaptParameterGroup('method', op.language.go!.optionalParamGroup);
      paramGroups.set(op.language.go!.optionalParamGroup.schema.language.go!.name, optionalParamsGroup);
    }

    method.optionalParamsGroup = optionalParamsGroup;
    method.returns = adaptResponseEnvelope(m4CodeModel, codeModel, op, method);
  }

  adaptMethodParameters(op, method);

  for (const apiver of values(op.apiVersions)) {
    method.apiVersions.push(apiver.version);
  }
}

function adaptHeaderScalarType(schema: m4.Schema, forParam: boolean): go.HeaderScalarType {
  // for header params, we never pass the element type by pointer
  const type = adaptWireType(schema, forParam);
  if (go.isHeaderScalarType(type)) {
    return type;
  }
  throw new Error(`unexpected header scalar parameter type ${schema.type}`);
}

function adaptPathScalarParameterType(schema: m4.Schema): go.PathScalarParameterType {
  const type = adaptWireType(schema);
  if (go.isPathScalarParameterType(type)) {
    return type;
  }
  throw new Error(`unexpected path scalar parameter type ${schema.type}`);
}

function adaptQueryScalarParameterType(schema: m4.Schema): go.QueryScalarParameterType {
  const type = adaptWireType(schema);
  if (go.isQueryScalarParameterType(type)) {
    return type;
  }
  throw new Error(`unexpected query scalar parameter type ${schema.type}`);
}

function adaptURIPrameterType(schema: m4.Schema): go.URIParameterType {
  const type = adaptWireType(schema);
  if (go.isURIParameterType(type)) {
    return type;
  }
  throw new Error(`unexpected URI parameter type ${schema.type}`);
}

function adaptClient(type: go.CodeModelType, group: m4.OperationGroup): go.Client {
  const description = `${group.language.go!.clientName} contains the methods for the ${group.language.go!.name} group.`;
  const client = new go.Client(group.language.go!.clientName, {description: description}, go.newClientOptions(type, group.language.go!.clientName));
  if (group.language.go!.complexHostParams) {
    client.templatedHost = group.language.go!.host;
  }
  if (group.language.go!.hostParams) {
    for (const hostParam of values(<Array<m4.Parameter>>group.language.go!.hostParams)) {
      const uriParam = new go.URIParameter(hostParam.language.go!.name, hostParam.language.go!.serializedName, adaptURIPrameterType(hostParam.schema),
        adaptParameterStyle(hostParam), hostParam.language.go!.byValue, adaptMethodLocation(hostParam.implementation));
      client.parameters.push(uriParam);
    }
  }

  return client;
}

function adaptMethodLocation(location?: m4.ImplementationLocation): go.ParameterLocation {
  if (!location) {
    return 'method';
  }
  switch (location) {
    case m4.ImplementationLocation.Client:
      return 'client';
    case m4.ImplementationLocation.Method:
      return 'method';
    default:
      throw new Error(`unhandled parameter location type ${location}`);
  }
}

function adaptMethodParameters(op: m4.Operation, method: go.MethodType | go.NextPageMethod) {
  if (!op.parameters) {
    return;
  }

  for (const param of values(helpers.aggregateParameters(op))) {
    const methodParam = adaptMethodParameter(op, param);
    method.parameters.push(methodParam);
  }
}

function adaptResponseEnvelope(m4CodeModel: m4.CodeModel, codeModel: go.CodeModel, op: m4.Operation, forMethod: go.MethodType): go.ResponseEnvelope {
  const respEnvSchema = <m4.ObjectSchema>op.language.go!.responseEnv;
  const respEnv = new go.ResponseEnvelope(respEnvSchema.language.go!.name, {description: respEnvSchema.language.go!.description}, forMethod);

  // add any headers
  for (const prop of values(respEnvSchema.properties)) {
    if (prop.language.go!.fromHeader) {
      let headerResp: go.HeaderScalarResponse | go.HeaderMapResponse;
      if (prop.schema.language.go!.headerCollectionPrefix) {
        const headerType = adaptWireType(prop.schema, false);
        if (headerType.kind !== 'map') {
          throw new Error(`unexpected type ${go.getTypeDeclaration(headerType)} for HeaderMapResponse ${prop.language.go!.name}`);
        }
        headerResp = new go.HeaderMapResponse(prop.language.go!.name, headerType, prop.schema.language.go!.headerCollectionPrefix);
      } else {
        headerResp = new go.HeaderScalarResponse(prop.language.go!.name, adaptHeaderScalarType(prop.schema, false), prop.language.go!.fromHeader, prop.language.go!.byValue);
      }
      if (hasDescription(prop.language.go!)) {
        headerResp.docs.description = prop.language.go!.description;
      }
      respEnv.headers.push(headerResp);
    }
  }

  if (!respEnvSchema.language.go!.resultProp) {
    return respEnv;
  }

  // now add the result field
  const resultProp = <m4.Property>respEnvSchema.language.go!.resultProp;
  if (helpers.isMultiRespOperation(op)) {
    respEnv.result = adaptAnyResult(op);
  } else if (resultProp.schema.type === m4.SchemaType.Binary) {
    respEnv.result = new go.BinaryResult(resultProp.language.go!.name);
  } else if (m4CodeModel.language.go!.headAsBoolean && op.requests![0].protocol.http!.method === 'head') {
    respEnv.result = new go.HeadAsBooleanResult(resultProp.language.go!.name);
  } else if (!resultProp.language.go!.embeddedType) {
    const resultType = adaptWireType(resultProp.schema);
    if (go.isMonomorphicResultType(resultType)) {
      respEnv.result = new go.MonomorphicResult(resultProp.language.go!.name, adaptResultFormat(helpers.getSchemaResponse(op)!.protocol), resultType, resultProp.language.go!.byValue);
      respEnv.result.xml = adaptXMLInfo(resultProp.schema);
    } else {
      throw new Error(`invalid monomorphic result type ${resultType.kind}`);
    }
  } else if (resultProp.isDiscriminator) {
    let ifaceResult: go.Interface | undefined;
    for (const iface of values(codeModel.interfaces)) {
      if (iface.name === resultProp.schema.language.go!.name) {
        ifaceResult = iface;
        break;
      }
    }
    if (!ifaceResult) {
      throw new Error(`didn't find InterfaceType for result property ${resultProp.schema.language.go!.name}`);
    }
    respEnv.result = new go.PolymorphicResult(ifaceResult);
  } else if (helpers.getSchemaResponse(op)) {
  /** 
   * The modelType will be a PolymorphicModel when the response envelope
   * is a concrete type from a polymorphic hierarchy
   */
    let modelType: go.Model | go.PolymorphicModel | undefined;
    for (const model of codeModel.models) {
      if ((model.kind === 'model' || model.kind === 'polymorphicModel') && model.name === resultProp.schema.language.go!.name) {
        modelType = model;
        break;
      }
    }
    if (!modelType) {
      throw new Error(`didn't find type name ${resultProp.schema.language.go!.name} for response envelope ${respEnv.name}`);
    }
    const resultFormat = adaptResultFormat(helpers.getSchemaResponse(op)!.protocol);
    if (resultFormat !== 'JSON' && resultFormat !== 'XML') {
      throw new Error(`unexpected result format ${resultFormat} for model ${modelType.name}`);
    }
    respEnv.result = new go.ModelResult(modelType, resultFormat);
  } else {
    throw new Error(`unhandled result type for operation ${op.language.go!.name}`);
  }

  if (hasDescription(resultProp.language.go!)) {
    respEnv.result.docs.description = resultProp.language.go!.description;
  }

  return respEnv;
}

function adaptMethodNaming(op: m4.Operation): go.MethodNaming {
  const info = <OperationNaming>op.language.go!;
  return new go.MethodNaming(info.protocolNaming.internalMethod, info.protocolNaming.requestMethod, info.protocolNaming.responseMethod);
}

function getStatusCodes(op: m4.Operation): Array<number> {
  // concat all status codes that return the same schema into one array.
  // this is to support operations that specify multiple response codes
  // that return the same schema (or no schema).
  let statusCodes = new Array<number>();
  for (const resp of values(op.responses)) {
    statusCodes = statusCodes.concat(parseInt(resp.protocol.http?.statusCodes));
  }
  if (statusCodes.length === 0) {
    // if the operation defines no status codes (which is non-conformant)
    // then add 200, 201, 202, and 204 to the list.  this is to accomodate
    // some quirky tests in the test server.
    // TODO: https://github.com/Azure/autorest.go/issues/659
    statusCodes = [200, 201, 202, 204];
  }
  return statusCodes;
}

function adaptMethodParameter(op: m4.Operation, param: m4.Parameter): go.MethodParameter {
  let adaptedParam: go.MethodParameter;
  let location: go.ParameterLocation = 'method';
  if (param.implementation === m4.ImplementationLocation.Client) {
    // check if we've already adapted this client parameter
    // TODO: grouped client params
    const clientParam = clientParams.get(param.language.go!.name);
    if (clientParam) {
      return clientParam;
    }
    location = 'client';
  }

  const style = adaptParameterStyle(param);

  // unfortunately param.language.go!.byValue isn't always populated.
  // since we can't trust it, we calculate the value instead.
  const byValue = helpers.isTypePassedByValue(param.schema) ? true : (go.isRequiredParameter(style) || (location === 'client' && go.isClientSideDefault(style)));

  switch (param.protocol.http?.in) {
    case 'body': {
      if (!op.requests![0].protocol.http!.mediaTypes) {
        throw new Error(`no media types defined for operation ${op.operationId}`);
      }
      let contentType = `"${op.requests![0].protocol.http!.mediaTypes[0]}"`;
      if (op.requests![0].protocol.http!.mediaTypes.length > 1) {
        for (const param of values(op.requests![0].parameters)) {
          // If a request defined more than one possible media type, then the param is expected to be synthesized from modelerfour
          // and should be a SealedChoice schema type that account for the acceptable media types defined in the swagger.
          if (param.origin === 'modelerfour:synthesized/content-type' && param.schema.type === m4.SchemaType.SealedChoice) {
            contentType = `string(${param.language.go!.name})`;
          }
        }
      }
      const bodyType = adaptWireType(param.schema);
      if (op.requests![0].protocol.http!.knownMediaType === KnownMediaType.Form) {
        const collectionFormat = adaptCollectionFormat(param);
        if (collectionFormat) {
          if (bodyType.kind !== 'slice') {
            throw new Error(`unexpected type ${go.getTypeDeclaration(bodyType)} for FormBodyCollectionParameter ${param.language.go!.name}`);
          }
          adaptedParam = new go.FormBodyCollectionParameter(param.language.go!.name, param.language.go!.serializedName, bodyType, collectionFormat, style,byValue);
        } else {
          adaptedParam = new go.FormBodyScalarParameter(param.language.go!.name, param.language.go!.serializedName, bodyType, style, byValue);
        }
      } else if (op.requests![0].protocol.http!.knownMediaType === KnownMediaType.Multipart) {
        adaptedParam = new go.MultipartFormBodyParameter(param.language.go!.name, bodyType, style, byValue);
      } else {
        const format = adaptBodyFormat(op.requests![0].protocol);
        adaptedParam = new go.BodyParameter(param.language.go!.name, format, contentType, bodyType, style, byValue);
        adaptedParam.xml = adaptXMLInfo(param.schema);
      }

      break;
    }
    case 'header': {
      const collectionFormat = adaptCollectionFormat(param);
      if (param.schema.language.go!.headerCollectionPrefix) {
        const headerType = adaptWireType(param.schema, true);
        if (headerType.kind !== 'map') {
          throw new Error(`unexpected type ${go.getTypeDeclaration(headerType)} for HeaderMapParameter ${param.language.go!.name}`);
        }
        adaptedParam = new go.HeaderMapParameter(param.language.go!.name, param.schema.language.go!.headerCollectionPrefix, headerType, style, byValue, location);
      } else if (collectionFormat) {
        const headerType = adaptWireType(param.schema, true);
        if (headerType.kind !== 'slice') {
          throw new Error(`unexpected type ${go.getTypeDeclaration(headerType)} for HeaderCollectionParameter ${param.language.go!.name}`);
        }
        adaptedParam = new go.HeaderCollectionParameter(param.language.go!.name, param.language.go!.serializedName, headerType, collectionFormat, style, byValue, location);
      } else {
        adaptedParam = new go.HeaderScalarParameter(param.language.go!.name, param.language.go!.serializedName, adaptHeaderScalarType(param.schema, true),
          style, byValue, location);
      }
      break;
    }
    case 'path': {
      const collectionFormat = adaptCollectionFormat(param);
      if (collectionFormat) {
        const pathType = adaptWireType(param.schema);
        if (pathType.kind !== 'slice') {
          throw new Error(`unexpected type ${go.getTypeDeclaration(pathType)} for PathCollectionParameter ${param.language.go!.name}`);
        }
        adaptedParam = new go.PathCollectionParameter(param.language.go!.name, param.language.go!.serializedName, !skipURLEncoding(param),
          pathType, collectionFormat, style, byValue, location);
      } else {
        const skipUrlEncoding = skipURLEncoding(param);
        adaptedParam = new go.PathScalarParameter(param.language.go!.name, param.language.go!.serializedName, !skipUrlEncoding,
          adaptPathScalarParameterType(param.schema), style, byValue, location);
        // this is a legacy hack to work around the fact that
        // swagger doesn't allow path params to be empty.
        adaptedParam.omitEmptyStringCheck = skipUrlEncoding && (adaptedParam.type.kind === 'string' || (adaptedParam.type.kind === 'constant' && adaptedParam.type.type === 'string'));
      }
      break;
    }
    case 'query': {
      const collectionFormat = adaptExtendedCollectionFormat(param);
      if (collectionFormat) {
        const queryType = adaptWireType(param.schema);
        if (queryType.kind !== 'slice') {
          throw new Error(`unexpected type ${go.getTypeDeclaration(queryType)} for QueryCollectionParameter ${param.language.go!.name}`);
        }
        adaptedParam = new go.QueryCollectionParameter(param.language.go!.name, param.language.go!.serializedName, !skipURLEncoding(param),
          queryType, collectionFormat, style, byValue, location);
      } else {
        adaptedParam = new go.QueryScalarParameter(param.language.go!.name, param.language.go!.serializedName, !skipURLEncoding(param),
          adaptQueryScalarParameterType(param.schema), style, byValue, location);
      }
      break;
    }
    case 'uri':
      adaptedParam = new go.URIParameter(param.language.go!.name, param.language.go!.serializedName, adaptURIPrameterType(param.schema),
        style, byValue, adaptParameterlocation(param));
      break;
    default: {
      if (param.protocol.http?.in) {
        throw new Error(`unhandled parameter location type ${param.protocol.http.in}`);
      }
      // this is a synthesized parameter (e.g. ResumeToken)
      if (param.language.go!.isResumeToken) {
        adaptedParam = new go.ResumeTokenParameter();
      } else {
        throw new Error(`unknown parameter in operation ${op.language.go!.name}`);
      }
    }
  }

  if (hasDescription(param.language.go!)) {
    adaptedParam.docs.description = param.language.go!.description;
  }

  // track client parameter for later use
  if (adaptedParam.location === 'client') {
    clientParams.set(param.language.go!.name, adaptedParam);
  }
  
  if (param.language.go!.paramGroup) {
    const paramGroup = findOrAdaptParamsGroup(param);
    // parameter groups can be shared across methods so don't add any duplicate parameters
    if (values(paramGroup.params).where((each: go.MethodParameter) => { return each.name === adaptedParam.name; }).count() === 0) {
      paramGroup.params.push(adaptedParam);
    }
    if (adaptedParam.style === 'required') {
      // if at least one param within a group is required then the group must be required.
      // however, it's possible that the param group was initially created from a non-required
      // param. so we need to be sure to update it as required.
      paramGroup.required = true;
    }
    adaptedParam.group = paramGroup;
  }

  return adaptedParam;
}

function adaptCollectionFormat(param: m4.Parameter): go.CollectionFormat | undefined {
  switch (param.protocol.http?.style) {
    case m4.SerializationStyle.PipeDelimited:
      return 'pipes';
    case m4.SerializationStyle.Simple:
      return 'csv';
    case m4.SerializationStyle.SpaceDelimited:
      return 'ssv';
    case m4.SerializationStyle.TabDelimited:
      return 'tsv';
    default:
      return undefined;
  }
}

function adaptExtendedCollectionFormat(param: m4.Parameter): go.ExtendedCollectionFormat | undefined {
  switch (param.protocol.http?.style) {
    case m4.SerializationStyle.Form:
      if (param.protocol.http?.explode === true){
        return 'multi';
      }
      return 'csv';
    case m4.SerializationStyle.PipeDelimited:
      return 'pipes';
    case m4.SerializationStyle.Simple:
      return 'csv';
    case m4.SerializationStyle.SpaceDelimited:
      return 'ssv';
    case m4.SerializationStyle.TabDelimited:
      return 'tsv';
    default:
      return undefined;
  }
}

// returns true if the parameter should not be URL encoded
function skipURLEncoding(param: m4.Parameter): boolean {
  if (param.extensions) {
    return param.extensions['x-ms-skip-url-encoding'] === true;
  }
  return false;
}

function adaptParameterlocation(param: m4.Parameter): go.ParameterLocation {
  if (!param.implementation) {
    return 'method';
  }
  switch (param.implementation) {
    case m4.ImplementationLocation.Client:
      return 'client';
    case m4.ImplementationLocation.Method:
      return 'method';
    default:
      throw new Error(`unhandled parameter location ${param.implementation}`);
  }
}

function adaptParameterStyle(param: m4.Parameter): go.ParameterStyle {
  if (param.clientDefaultValue) {
    const adaptedType = adaptWireType(param.schema);
    if (!go.isLiteralValueType(adaptedType)) {
      throw new Error(`unsupported client side default type ${go.getTypeDeclaration(adaptedType)} for parameter ${param.language.go!.name}`);
    }
    return new go.ClientSideDefault(new go.Literal(adaptedType, param.clientDefaultValue));
  } else if (param.schema.type === m4.SchemaType.Constant) {
    if (param.required) {
      return 'literal';
    }
    return 'flag';
  } else if (param.required === true) {
    return 'required';
  } else {
    return 'optional';
  }
}

function findOrAdaptParamsGroup(param: m4.Parameter): go.ParameterGroup {
  const groupProp = <m4.GroupProperty>param.language.go!.paramGroup;
  let paramGroup = paramGroups.get(groupProp.schema.language.go!.name);
  if (!paramGroup) {
    paramGroup = adaptParameterGroup(adaptParameterlocation(param), groupProp);
    paramGroups.set(groupProp.schema.language.go!.name, paramGroup);
  }

  return paramGroup;
}

function adaptParameterGroup(location: go.ParameterLocation, groupProp: m4.GroupProperty): go.ParameterGroup {
  const paramGroup = new go.ParameterGroup(groupProp.language.go!.name, groupProp.schema.language.go!.name, groupProp.required === true, location);
  paramGroup.docs.description = groupProp.language.go!.description;
  return paramGroup;
}

// returns the media type used by the protocol
function adaptBodyFormat(protocol: m4.Protocols): go.BodyFormat {
  switch (protocol.http!.knownMediaType) {
    case KnownMediaType.Json:
      return 'JSON';
    case KnownMediaType.Xml:
      return 'XML';
    case KnownMediaType.Binary:
      return 'binary';
    case KnownMediaType.Text:
      return 'Text';
    default:
      throw new Error(`unhandled body protocol format ${protocol.http!.knownMediaType}`);
  }
}

function adaptAnyResult(op: m4.Operation): go.AnyResult {
  const resultTypes: Record<number, go.WireType> = {};
  for (const resp of values(op.responses)) {
    let wireType: go.WireType;
    if (helpers.isSchemaResponse(resp)) {
      wireType = adaptWireType(resp.schema);
    } else {
      // the operation contains a mix of schemas and non-schema responses
      continue;
    }

    resultTypes[parseInt(resp.protocol.http!.statusCodes)] = wireType;
  }

  return new go.AnyResult('Value', 'JSON', resultTypes);
}

function adaptResultFormat(protocol: m4.Protocols): go.ResultFormat {
  switch (protocol.http!.knownMediaType) {
    case KnownMediaType.Json:
      return 'JSON';
    case KnownMediaType.Xml:
      return 'XML';
    case KnownMediaType.Text:
      return 'Text';
    default:
      throw new Error(`unhandled result protocol format ${protocol.http!.knownMediaType}`);
  }
}
