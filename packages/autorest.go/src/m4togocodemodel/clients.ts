/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as m4 from '@autorest/codemodel';
import { KnownMediaType } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import { adaptXMLInfo } from './types';
import { adaptPossibleType, hasDescription } from './types';
import * as go from '../gocodemodel/gocodemodel';
import { aggregateParameters, getSchemaResponse, isLROOperation, isMultiRespOperation, isPageableOperation, isSchemaResponse } from '../transform/helpers';
import { OperationNaming } from '../transform/namer';

// track all of the client and parameter group params across all operations
// as not every option might contain them, and parameter groups can be shared
// across multiple operations
const clientParams = new Map<string, go.Parameter>();
const paramGroups = new Map<string, go.ParameterGroup>();

export function adaptClients(m4CodeModel: m4.CodeModel, codeModel: go.GoCodeModel) {
  for (const group of values(m4CodeModel.operationGroups)) {
    const client = adaptClient(group);

    for (const op of values(group.operations)) {
      const httpPath = <string>op.requests![0].protocol.http!.path;
      const httpMethod = op.requests![0].protocol.http!.method;
      let method: go.Method | go.LROMethod | go.LROPageableMethod | go.PageableMethod;
      const naming = adaptMethodNaming(op);

      if (isLROOperation(op) && isPageableOperation(op)) {
        method = new go.LROPageableMethod(op.language.go!.name, client, httpPath, httpMethod, getStatusCodes(op), naming);
        (<go.LROPageableMethod>method).finalStateVia = (op.extensions?.['x-ms-long-running-operation-options']?.['final-state-via']);
        (<go.LROPageableMethod>method).nextLinkName = op.language.go!.paging.nextLinkName;
        if (op.language.go!.paging.nextLinkOperation) {
          // adapt the next link operation
          const nextPageMethod = new go.NextPageMethod(op.language.go!.paging.nextLinkOperation.language.go.name, client, httpPath, httpMethod, getStatusCodes(op.language.go!.paging.nextLinkOperation));
          populateMethod(op.language.go!.paging.nextLinkOperation, nextPageMethod, m4CodeModel, codeModel);
          (<go.LROPageableMethod>method).nextPageMethod = nextPageMethod;
        }
      } else if (isLROOperation(op)) {
        method = new go.LROMethod(op.language.go!.name, client, httpPath, httpMethod, getStatusCodes(op), naming);
        (<go.LROMethod>method).finalStateVia = (op.extensions?.['x-ms-long-running-operation-options']?.['final-state-via']);
      } else if (isPageableOperation(op)) {
        if (op.language.go!.paging.isNextOp) {
          continue;
        }
        method = new go.PageableMethod(op.language.go!.name, client, httpPath, httpMethod, getStatusCodes(op), naming);
        (<go.PageableMethod>method).nextLinkName = op.language.go!.paging.nextLinkName;
        if (op.language.go!.paging.nextLinkOperation) {
          // adapt the next link operation
          const httpPath = <string>op.language.go!.paging.nextLinkOperation.requests![0].protocol.http!.path;
          const httpMethod = op.language.go!.paging.nextLinkOperation.requests![0].protocol.http!.method;
          const nextPageMethod = new go.NextPageMethod(op.language.go!.paging.nextLinkOperation.language.go.name, client, httpPath, httpMethod, getStatusCodes(op.language.go!.paging.nextLinkOperation));
          populateMethod(op.language.go!.paging.nextLinkOperation, nextPageMethod, m4CodeModel, codeModel);
          (<go.PageableMethod>method).nextPageMethod = nextPageMethod;
        }
      } else {
        method = new go.Method(op.language.go!.name, client, httpPath, httpMethod, getStatusCodes(op), naming);
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

    codeModel.clients.push(client);
  }
}

function populateMethod(op: m4.Operation, method: go.Method | go.NextPageMethod, m4CodeModel: m4.CodeModel, codeModel: go.GoCodeModel) {
  if (go.isMethod(method)) {
    if (hasDescription(op.language.go!)) {
      method.description = op.language.go!.description;
    }

    let optionalParamsGroup = paramGroups.get(op.language.go!.optionalParamGroup.schema.language.go!.name);
    if (!optionalParamsGroup) {
      optionalParamsGroup = adaptParameterGroup('method', op.language.go!.optionalParamGroup);
      paramGroups.set(op.language.go!.optionalParamGroup.schema.language.go!.name, optionalParamsGroup);
    }

    method.optionalParamsGroup = optionalParamsGroup;
    method.responseEnvelope = adaptResponseEnvelope(m4CodeModel, codeModel, op, method);
  }

  adaptMethodParameters(op, method);

  for (const apiver of values(op.apiVersions)) {
    method.apiVersions.push(apiver.version);
  }
}

function adaptHeaderType(schema: m4.Schema, forParam: boolean): go.HeaderType {
  // for header params, we never pass the element type by pointer
  const type = adaptPossibleType(schema, forParam);
  if (go.isInterfaceType(type) || go.isModelType(type) || go.isPolymorphicType(type) || go.isStandardType(type)) {
    throw new Error(`unexpected header parameter type ${schema.type}`);
  }
  return type;
}

function adaptPathPrameterType(schema: m4.Schema): go.PathParameterType {
  const type = adaptPossibleType(schema);
  if (go.isMapType(type) || go.isInterfaceType(type) || go.isModelType(type) || go.isPolymorphicType(type) || go.isStandardType(type)) {
    throw new Error(`unexpected path parameter type ${schema.type}`);
  }
  return type;
}

function adaptQueryParameterType(schema: m4.Schema): go.QueryParameterType {
  const type = adaptPossibleType(schema);
  if (go.isMapType(type) || go.isInterfaceType(type) || go.isModelType(type) || go.isPolymorphicType(type) || go.isStandardType(type)) {
    throw new Error(`unexpected query parameter type ${schema.type}`);
  } else if (go.isSliceType(type)) {
    type.elementTypeByValue = true;
  }
  return type;
}

function adaptURIPrameterType(schema: m4.Schema): go.URIParameterType {
  const type = adaptPossibleType(schema);
  if (!go.isConstantType(type) && !go.isPrimitiveType(type)) {
    throw new Error(`unexpected URI parameter type ${schema.type}`);
  }
  return type;
}

function adaptClient(group: m4.OperationGroup): go.Client {
  const client = new go.Client(group.language.go!.clientName, group.language.go!.name, group.language.go!.clientCtorName);

  client.host = group.language.go!.host;
  if (group.language.go!.complexHostParams) {
    client.complexHostParams = true;
  }
  if (group.language.go!.hostParams) {
    for (const hostParam of values(<Array<m4.Parameter>>group.language.go!.hostParams)) {
      const uriParam = new go.URIParameter(hostParam.language.go!.name, hostParam.language.go!.serializedName, adaptURIPrameterType(hostParam.schema),
        adaptParameterType(hostParam), hostParam.language.go!.byValue, adaptMethodLocation(hostParam.implementation));
      client.hostParams.push(uriParam);
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

function adaptMethodParameters(op: m4.Operation, method: go.Method | go.NextPageMethod) {
  if (!op.parameters) {
    return;
  }

  for (const param of values(aggregateParameters(op))) {
    const methodParam = adaptMethodParameter(op, param);
    method.parameters.push(methodParam);
  }
}

function adaptResponseEnvelope(m4CodeModel: m4.CodeModel, codeModel: go.GoCodeModel, op: m4.Operation, forMethod: go.Method): go.ResponseEnvelope {
  const respEnvSchema = <m4.ObjectSchema>op.language.go!.responseEnv;
  const respEnv = new go.ResponseEnvelope(respEnvSchema.language.go!.name, respEnvSchema.language.go!.description, forMethod);

  // add any headers
  for (const prop of values(respEnvSchema.properties)) {
    if (prop.language.go!.fromHeader) {
      const headerResp = new go.HeaderResponse(prop.language.go!.name, adaptHeaderType(prop.schema, false), prop.language.go!.fromHeader, prop.language.go!.byValue);
      if (hasDescription(prop.language.go!)) {
        headerResp.description = prop.language.go!.description;
      }
      if (prop.schema.language.go!.headerCollectionPrefix) {
        headerResp.collectionPrefix = prop.schema.language.go!.headerCollectionPrefix;
      }
      respEnv.headers.push(headerResp);
    }
  }

  if (!respEnvSchema.language.go!.resultProp) {
    return respEnv;
  }

  // now add the result field
  const resultProp = <m4.Property>respEnvSchema.language.go!.resultProp;
  if (isMultiRespOperation(op)) {
    respEnv.result = adaptAnyResult(op);
  } else if (resultProp.schema.type === m4.SchemaType.Binary) {
    respEnv.result = new go.BinaryResult(resultProp.language.go!.name, 'binary');
  } else if (m4CodeModel.language.go!.headAsBoolean && op.requests![0].protocol.http!.method === 'head') {
    respEnv.result = new go.HeadAsBooleanResult(resultProp.language.go!.name);
  } else if (!resultProp.language.go!.embeddedType) {
    const resultType = adaptPossibleType(resultProp.schema);
    if (go.isInterfaceType(resultType) || go.isLiteralValue(resultType) || go.isModelType(resultType) || go.isPolymorphicType(resultType) || go.isStandardType(resultType)) {
      throw new Error(`invalid monomorphic result type ${resultType}`);
    }
    respEnv.result = new go.MonomorphicResult(resultProp.language.go!.name, adaptResultFormat(getSchemaResponse(op)!.protocol), resultType, resultProp.language.go!.byValue);
    respEnv.result.xml = adaptXMLInfo(resultProp.schema);
  } else if (resultProp.isDiscriminator) {
    let ifaceResult: go.InterfaceType | undefined;
    for (const iface of values(codeModel.interfaceTypes)) {
      if (iface.name === resultProp.schema.language.go!.name) {
        ifaceResult = iface;
        break;
      }
    }
    if (!ifaceResult) {
      throw new Error(`didn't find InterfaceType for result property ${resultProp.schema.language.go!.name}`);
    }
    respEnv.result = new go.PolymorphicResult(ifaceResult);
  } else if (getSchemaResponse(op)) {
    let modelType: go.ModelType | undefined;
    for (const model of codeModel.models) {
      if (model.name === resultProp.schema.language.go!.name && go.isModelType(model)) {
        modelType = model;
        break;
      }
    }
    if (!modelType) {
      throw new Error(`didn't find ModelType for response envelope ${respEnv.name}`);
    }
    respEnv.result = new go.ModelResult(modelType, adaptResultFormat(getSchemaResponse(op)!.protocol));
  } else {
    throw new Error(`unhandled result type for operation ${op.language.go!.name}`);
  }

  if (hasDescription(resultProp.language.go!)) {
      respEnv.result!.description = resultProp.language.go!.description;
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

function adaptMethodParameter(op: m4.Operation, param: m4.Parameter): go.Parameter {
  let adaptedParam: go.Parameter;
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

  switch (param.protocol.http?.in) {
    case 'body': {
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
      const bodyType = adaptPossibleType(param.schema);
      const paramType = adaptParameterType(param);
      if (op.requests![0].protocol.http!.knownMediaType === KnownMediaType.Form) {
        adaptedParam = new go.FormBodyParameter(param.language.go!.name, param.language.go!.serializedName, bodyType, paramType, param.language.go!.byValue);
        (<go.FormBodyParameter>adaptedParam).delimiter = adaptParamDelimiter(param);
      } else if (op.requests![0].protocol.http!.knownMediaType === KnownMediaType.Multipart) {
        adaptedParam = new go.MultipartFormBodyParameter(param.language.go!.name, bodyType, paramType, param.language.go!.byValue);
      } else {
        const format = adaptBodyFormat(op.requests![0].protocol);
        adaptedParam = new go.BodyParameter(param.language.go!.name, format, contentType, bodyType, paramType, param.language.go!.byValue);
      }

      adaptedParam.xml = adaptXMLInfo(param.schema);
      break;
    }
    case 'header': {
      const headerType = adaptHeaderType(param.schema, true);
      adaptedParam = new go.HeaderParameter(param.language.go!.name, param.language.go!.serializedName, headerType, adaptParameterType(param),
                param.language.go!.byValue, location);

      if (param.schema.language.go!.headerCollectionPrefix) {
        (<go.HeaderParameter>adaptedParam).collectionPrefix = param.schema.language.go!.headerCollectionPrefix;
      }
      (<go.HeaderParameter>adaptedParam).delimiter = adaptParamDelimiter(param);
      break;
    }
    case 'path':
      adaptedParam = new go.PathParameter(param.language.go!.name, param.language.go!.serializedName, !skipURLEncoding(param),
        adaptPathPrameterType(param.schema), adaptParameterType(param), param.language.go!.byValue, location);
      (<go.PathParameter>adaptedParam).delimiter = adaptParamDelimiter(param);
      break;
    case 'query': {
      const queryType = adaptQueryParameterType(param.schema);
      adaptedParam = new go.QueryParameter(param.language.go!.name, param.language.go!.serializedName, !skipURLEncoding(param),
                param.protocol.http?.explode === true, queryType, adaptParameterType(param), param.language.go!.byValue, location);
      (<go.QueryParameter>adaptedParam).delimiter = adaptParamDelimiter(param);
      break;
    }
    case 'uri':
      adaptedParam = new go.URIParameter(param.language.go!.name, param.language.go!.serializedName, adaptURIPrameterType(param.schema),
        adaptParameterType(param), param.language.go!.byValue, adaptParameterlocation(param));
      break;
    default: {
      if (param.protocol.http?.in) {
        throw new Error(`unhandled parameter location type ${param.protocol.http.in}`);
      }
      // this is a synthesized parameter (e.g. ResumeToken)
      if (param.language.go!.isResumeToken) {
        adaptedParam = new go.ResumeTokenParameter(param.language.go!.name);
      } else {
        const type = adaptPossibleType(param.schema);
        const paramType = adaptParameterType(param);
        const paramLoc = adaptParameterlocation(param);
        adaptedParam = new go.Parameter(param.language.go!.name, type, paramType, param.language.go!.byValue, paramLoc);
      }
    }
  }

  if (hasDescription(param.language.go!)) {
    adaptedParam.description = param.language.go!.description;
  }

  // track client parameter for later use
  if (adaptedParam.location === 'client') {
    clientParams.set(param.language.go!.name, adaptedParam);
  }
  
  if (param.language.go!.paramGroup) {
    const paramGroup = findOrAdaptParamsGroup(param);
    // parameter groups can be shared across methods so don't add any duplicate parameters
    if (values(paramGroup.params).where((each: go.Parameter) => { return each.paramName === adaptedParam.paramName; }).count() === 0) {
      paramGroup.params.push(adaptedParam);
    }
    if (adaptedParam.paramType === 'required') {
      // if at least one param within a group is required then the group must be required.
      // however, it's possible that the param group was initially created from a non-required
      // param. so we need to be sure to update it as required.
      paramGroup.required = true;
    }
    adaptedParam.group = paramGroup;
  }

  return adaptedParam;
}

function adaptParamDelimiter(param: m4.Parameter): '|' | ' ' | '\\t' | undefined {
  switch (param.protocol.http?.style) {
    case m4.SerializationStyle.PipeDelimited:
      return '|';
    case m4.SerializationStyle.SpaceDelimited:
      return ' ';
    case m4.SerializationStyle.TabDelimited:
      return '\\t';
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

function adaptParameterType(param: m4.Parameter): go.ParameterType {
  if (param.clientDefaultValue) {
    const adaptedType = adaptPossibleType(param.schema);
    if (!go.isLiteralValueType(adaptedType)) {
      throw new Error(`unsupported client side default type ${go.getTypeDeclaration(adaptedType)} for parameter ${param.language.go!.name}`);
    }
    return new go.ClientSideDefault(new go.LiteralValue(adaptedType, param.clientDefaultValue));
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
  paramGroup.description = groupProp.language.go!.description;
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
  const resultTypes: Record<number, go.PossibleType> = {};
  for (const resp of values(op.responses)) {
    let possibleType: go.PossibleType;
    if (isSchemaResponse(resp)) {
      possibleType = adaptPossibleType(resp.schema);
    } else {
      // the operation contains a mix of schemas and non-schema responses
      continue;
    }

    resultTypes[parseInt(resp.protocol.http!.statusCodes)] = possibleType;
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
