/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { CodeModel as M4CodeModel, GroupProperty, ImplementationLocation, Operation, OperationGroup, Parameter as M4Parameter, Property, Protocols, Schema, SchemaType, SerializationStyle, ObjectSchema } from '@autorest/codemodel';
import { Session } from '@autorest/extension-base';
import { KnownMediaType } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import { adaptXMLInfo } from './types';
import { adaptPossibleType, hasDescription } from './types';
import { AnyResult, BinaryResult, BodyFormat, FormBodyParameter, HeadAsBooleanResult, InterfaceType, LiteralValue, ModelResult, MonomorphicResult, MultipartFormBodyParameter, NextPageMethod, PageableMethod, ResultFormat, getTypeDeclaration, isConstantType, isInterfaceType, isLiteralValue, isLiteralValueType, isModelType, isPolymorphicType, isSliceType, isStandardType } from '../gocodemodel/gocodemodel';
import { BodyParameter, Client, ClientSideDefault, GoCodeModel, HeaderParameter, HeaderType, isMethod, Method, ModelType, LROMethod, LROPageableMethod, Parameter, ParameterType, PathParameter, PathParameterType, PossibleType, URIParameter, isPrimitiveType, ParameterGroup, PolymorphicResult, QueryParameter, QueryParameterType, ResponseEnvelope, ResumeTokenParameter, URIParameterType, MethodNaming, HeaderResponse, isMapType, ParameterLocation } from '../gocodemodel/gocodemodel';
import { aggregateParameters, getSchemaResponse, isLROOperation, isMultiRespOperation, isPageableOperation, isSchemaResponse } from '../transform/helpers';
import { OperationNaming } from '../transform/namer';

// track all of the client and parameter group params across all operations
// as not every option might contain them, and parameter groups can be shared
// across multiple operations
const clientParams = new Map<string, Parameter>();
const paramGroups = new Map<string, ParameterGroup>();

export function adaptClients(session: Session<M4CodeModel>, codeModel: GoCodeModel): Array<Client> | undefined {
  const clients = new Array<Client>();
  for (const group of values(session.model.operationGroups)) {
    const client = adaptClient(group);

    for (const op of values(group.operations)) {
      const httpPath = <string>op.requests![0].protocol.http!.path;
      const httpMethod = op.requests![0].protocol.http!.method;
      let method: Method | LROMethod | LROPageableMethod | PageableMethod;
      const naming = adaptMethodNaming(op);

      if (isLROOperation(op) && isPageableOperation(op)) {
        method = new LROPageableMethod(op.language.go!.name, client, httpPath, httpMethod, getStatusCodes(op), naming);
        (<LROPageableMethod>method).finalStateVia = (op.extensions?.['x-ms-long-running-operation-options']?.['final-state-via']);
        (<LROPageableMethod>method).nextLinkName = op.language.go!.paging.nextLinkName;
        if (op.language.go!.paging.nextLinkOperation) {
          // adapt the next link operation
          const nextPageMethod = new NextPageMethod(op.language.go!.paging.nextLinkOperation.language.go.name, client, httpPath, httpMethod, getStatusCodes(op.language.go!.paging.nextLinkOperation));
          populateMethod(op.language.go!.paging.nextLinkOperation, nextPageMethod, session.model, codeModel);
          (<LROPageableMethod>method).nextPageMethod = nextPageMethod;
        }
      } else if (isLROOperation(op)) {
        method = new LROMethod(op.language.go!.name, client, httpPath, httpMethod, getStatusCodes(op), naming);
        (<LROMethod>method).finalStateVia = (op.extensions?.['x-ms-long-running-operation-options']?.['final-state-via']);
      } else if (isPageableOperation(op)) {
        if (op.language.go!.paging.isNextOp) {
          continue;
        }
        method = new PageableMethod(op.language.go!.name, client, httpPath, httpMethod, getStatusCodes(op), naming);
        (<PageableMethod>method).nextLinkName = op.language.go!.paging.nextLinkName;
        if (op.language.go!.paging.nextLinkOperation) {
          // adapt the next link operation
          const httpPath = <string>op.language.go!.paging.nextLinkOperation.requests![0].protocol.http!.path;
          const httpMethod = op.language.go!.paging.nextLinkOperation.requests![0].protocol.http!.method;
          const nextPageMethod = new NextPageMethod(op.language.go!.paging.nextLinkOperation.language.go.name, client, httpPath, httpMethod, getStatusCodes(op.language.go!.paging.nextLinkOperation));
          populateMethod(op.language.go!.paging.nextLinkOperation, nextPageMethod, session.model, codeModel);
          (<PageableMethod>method).nextPageMethod = nextPageMethod;
        }
      } else {
        method = new Method(op.language.go!.name, client, httpPath, httpMethod, getStatusCodes(op), naming);
      }

      populateMethod(op, method, session.model, codeModel);

      client.methods.push(method);
    }

    // if any client parameters were adapted, add them to the client
    if (group.language.go!.clientParams) {
      client.parameters = new Array<Parameter>();
      for (const param of <Array<M4Parameter>>group.language.go!.clientParams) {
        const adaptedParam = clientParams.get(param.language.go!.name);
        if (!adaptedParam) {
          throw new Error(`missing adapted client parameter ${param.language.go!.name}`);
        }
        client.parameters.push(adaptedParam);
      }
    }

    clients.push(client);
  }

  if (clients.length === 0) {
    return undefined;
  }
  return clients;
}

function populateMethod(op: Operation, method: Method | NextPageMethod, m4CodeModel: M4CodeModel, codeModel: GoCodeModel) {
  if (isMethod(method)) {
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

  method.parameters = adaptMethodParameters(op);

  if (op.apiVersions) {
    method.apiVersions = new Array<string>();
    for (const apiver of op.apiVersions) {
      method.apiVersions.push(apiver.version);
    }
  }
}

function adaptHeaderType(schema: Schema, forParam: boolean): HeaderType {
  // for header params, we never pass the element type by pointer
  const type = adaptPossibleType(schema, forParam);
  if (isInterfaceType(type) || isModelType(type) || isPolymorphicType(type) || isStandardType(type)) {
    throw new Error(`unexpected header parameter type ${schema.type}`);
  }
  return type;
}

function adaptPathPrameterType(schema: Schema): PathParameterType {
  const type = adaptPossibleType(schema);
  if (isMapType(type) || isInterfaceType(type) || isModelType(type) || isPolymorphicType(type) || isStandardType(type)) {
    throw new Error(`unexpected path parameter type ${schema.type}`);
  }
  return type;
}

function adaptQueryParameterType(schema: Schema): QueryParameterType {
  const type = adaptPossibleType(schema);
  if (isMapType(type) || isInterfaceType(type) || isModelType(type) || isPolymorphicType(type) || isStandardType(type)) {
    throw new Error(`unexpected query parameter type ${schema.type}`);
  } else if (isSliceType(type)) {
    type.elementTypeByValue = true;
  }
  return type;
}

function adaptURIPrameterType(schema: Schema): URIParameterType {
  const type = adaptPossibleType(schema);
  if (!isConstantType(type) && !isPrimitiveType(type)) {
    throw new Error(`unexpected URI parameter type ${schema.type}`);
  }
  return type;
}

function adaptClient(group: OperationGroup): Client {
  const client = new Client(group.language.go!.clientName, group.language.go!.name, group.language.go!.clientCtorName);

  client.host = group.language.go!.host;
  if (group.language.go!.complexHostParams) {
    client.complexHostParams = true;
  }
  if (group.language.go!.hostParams) {
    const uriParams = new Array<URIParameter>();
    for (const hostParam of values(<Array<M4Parameter>>group.language.go!.hostParams)) {
      const uriParam = new URIParameter(hostParam.language.go!.name, hostParam.language.go!.serializedName, adaptURIPrameterType(hostParam.schema),
        adaptParameterType(hostParam), hostParam.language.go!.byValue, adaptMethodLocation(hostParam.implementation));
      uriParams.push(uriParam);
    }
    client.hostParams = uriParams;
  }

  return client;
}

function adaptMethodLocation(location?: ImplementationLocation): ParameterLocation {
  if (!location) {
    return 'method';
  }
  switch (location) {
    case ImplementationLocation.Client:
      return 'client';
    case ImplementationLocation.Method:
      return 'method';
    default:
      throw new Error(`unhandled parameter location type ${location}`);
  }
}

function adaptMethodParameters(op: Operation): Array<Parameter> | undefined {
  if (!op.parameters) {
    return undefined;
  }

  const methodParams = new Array<Parameter>();
  for (const param of values(aggregateParameters(op))) {
    const methodParam = adaptMethodParameter(op, param);
    methodParams.push(methodParam);
  }

  return methodParams;
}

function adaptResponseEnvelope(m4CodeModel: M4CodeModel, codeModel: GoCodeModel, op: Operation, forMethod: Method): ResponseEnvelope {
  const respEnvSchema = <ObjectSchema>op.language.go!.responseEnv;
  const respEnv = new ResponseEnvelope(respEnvSchema.language.go!.name, respEnvSchema.language.go!.description, forMethod);

  // add any headers
  for (const prop of values(respEnvSchema.properties)) {
    if (prop.language.go!.fromHeader) {
      if (!respEnv.headers) {
        respEnv.headers = new Array<HeaderResponse>();
      }
      const headerResp = new HeaderResponse(prop.language.go!.name, adaptHeaderType(prop.schema, false), prop.language.go!.fromHeader, prop.language.go!.byValue);
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
  const resultProp = <Property>respEnvSchema.language.go!.resultProp;
  if (isMultiRespOperation(op)) {
    respEnv.result = adaptAnyResult(op);
  } else if (resultProp.schema.type === SchemaType.Binary) {
    respEnv.result = new BinaryResult(resultProp.language.go!.name, 'binary');
  } else if (m4CodeModel.language.go!.headAsBoolean && op.requests![0].protocol.http!.method === 'head') {
    respEnv.result = new HeadAsBooleanResult(resultProp.language.go!.name);
  } else if (!resultProp.language.go!.embeddedType) {
    const resultType = adaptPossibleType(resultProp.schema);
    if (isInterfaceType(resultType) || isLiteralValue(resultType) || isModelType(resultType) || isPolymorphicType(resultType) || isStandardType(resultType)) {
      throw new Error(`invalid monomorphic result type ${resultType}`);
    }
    respEnv.result = new MonomorphicResult(resultProp.language.go!.name, adaptResultFormat(getSchemaResponse(op)!.protocol), resultType, resultProp.language.go!.byValue);
    respEnv.result.xml = adaptXMLInfo(resultProp.schema);
  } else if (resultProp.isDiscriminator) {
    let ifaceResult: InterfaceType | undefined;
    for (const iface of values(codeModel.interfaceTypes)) {
      if (iface.name === resultProp.schema.language.go!.name) {
        ifaceResult = iface;
        break;
      }
    }
    if (!ifaceResult) {
      throw new Error(`didn't find InterfaceType for result property ${resultProp.schema.language.go!.name}`);
    }
    respEnv.result = new PolymorphicResult(ifaceResult);
  } else if (getSchemaResponse(op)) {
    let modelType: ModelType | undefined;
    for (const model of codeModel.models) {
      if (model.name === resultProp.schema.language.go!.name && isModelType(model)) {
        modelType = model;
        break;
      }
    }
    if (!modelType) {
      throw new Error(`didn't find ModelType for response envelope ${respEnv.name}`);
    }
    respEnv.result = new ModelResult(modelType, adaptResultFormat(getSchemaResponse(op)!.protocol));
  } else {
    throw new Error(`unhandled result type for operation ${op.language.go!.name}`);
  }

  if (hasDescription(resultProp.language.go!)) {
      respEnv.result!.description = resultProp.language.go!.description;
  }

  return respEnv;
}

function adaptMethodNaming(op: Operation): MethodNaming {
  const info = <OperationNaming>op.language.go!;
  return new MethodNaming(info.protocolNaming.internalMethod, info.protocolNaming.requestMethod, info.protocolNaming.responseMethod);
}

function getStatusCodes(op: Operation): Array<number> {
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

function adaptMethodParameter(op: Operation, param: M4Parameter): Parameter {
  let adaptedParam: Parameter;
  let location: ParameterLocation = 'method';
  if (param.implementation === ImplementationLocation.Client) {
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
          if (param.origin === 'modelerfour:synthesized/content-type' && param.schema.type === SchemaType.SealedChoice) {
            contentType = `string(${param.language.go!.name})`;
          }
        }
      }
      const bodyType = adaptPossibleType(param.schema);
      const paramType = adaptParameterType(param);
      if (op.requests![0].protocol.http!.knownMediaType === KnownMediaType.Form) {
        adaptedParam = new FormBodyParameter(param.language.go!.name, param.language.go!.serializedName, bodyType, paramType, param.language.go!.byValue);
        (<FormBodyParameter>adaptedParam).delimiter = adaptParamDelimiter(param);
      } else if (op.requests![0].protocol.http!.knownMediaType === KnownMediaType.Multipart) {
        adaptedParam = new MultipartFormBodyParameter(param.language.go!.name, bodyType, paramType, param.language.go!.byValue);
      } else {
        const format = adaptBodyFormat(op.requests![0].protocol);
        adaptedParam = new BodyParameter(param.language.go!.name, format, contentType, bodyType, paramType, param.language.go!.byValue);
      }

      adaptedParam.xml = adaptXMLInfo(param.schema);
      break;
    }
    case 'header': {
      const headerType = adaptHeaderType(param.schema, true);
      adaptedParam = new HeaderParameter(param.language.go!.name, param.language.go!.serializedName, headerType, adaptParameterType(param),
                param.language.go!.byValue, location);

      if (param.schema.language.go!.headerCollectionPrefix) {
        (<HeaderParameter>adaptedParam).collectionPrefix = param.schema.language.go!.headerCollectionPrefix;
      }
      (<HeaderParameter>adaptedParam).delimiter = adaptParamDelimiter(param);
      break;
    }
    case 'path':
      adaptedParam = new PathParameter(param.language.go!.name, param.language.go!.serializedName, !skipURLEncoding(param),
        adaptPathPrameterType(param.schema), adaptParameterType(param), param.language.go!.byValue, location);
      (<PathParameter>adaptedParam).delimiter = adaptParamDelimiter(param);
      break;
    case 'query': {
      const queryType = adaptQueryParameterType(param.schema);
      adaptedParam = new QueryParameter(param.language.go!.name, param.language.go!.serializedName, !skipURLEncoding(param),
                param.protocol.http?.explode === true, queryType, adaptParameterType(param), param.language.go!.byValue, location);
      (<QueryParameter>adaptedParam).delimiter = adaptParamDelimiter(param);
      break;
    }
    case 'uri':
      adaptedParam = new URIParameter(param.language.go!.name, param.language.go!.serializedName, adaptURIPrameterType(param.schema),
        adaptParameterType(param), param.language.go!.byValue, adaptParameterlocation(param));
      break;
    default: {
      if (param.protocol.http?.in) {
        throw new Error(`unhandled parameter location type ${param.protocol.http.in}`);
      }
      // this is a synthesized parameter (e.g. ResumeToken)
      if (param.language.go!.isResumeToken) {
        adaptedParam = new ResumeTokenParameter(param.language.go!.name);
      } else {
        const type = adaptPossibleType(param.schema);
        const paramType = adaptParameterType(param);
        const paramLoc = adaptParameterlocation(param);
        adaptedParam = new Parameter(param.language.go!.name, type, paramType, param.language.go!.byValue, paramLoc);
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
    if (!paramGroup.params) {
      paramGroup.params = new Array<Parameter>();
    }

    // parameter groups can be shared across methods so don't add any duplicate parameters
    if (values(paramGroup.params).where((each: Parameter) => { return each.paramName === adaptedParam.paramName; }).count() === 0) {
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

function adaptParamDelimiter(param: M4Parameter): '|' | ' ' | '\\t' | undefined {
  switch (param.protocol.http?.style) {
    case SerializationStyle.PipeDelimited:
      return '|';
    case SerializationStyle.SpaceDelimited:
      return ' ';
    case SerializationStyle.TabDelimited:
      return '\\t';
    default:
      return undefined;
  }
}

// returns true if the parameter should not be URL encoded
function skipURLEncoding(param: M4Parameter): boolean {
  if (param.extensions) {
    return param.extensions['x-ms-skip-url-encoding'] === true;
  }
  return false;
}

function adaptParameterlocation(param: M4Parameter): ParameterLocation {
  if (!param.implementation) {
    return 'method';
  }
  switch (param.implementation) {
    case ImplementationLocation.Client:
      return 'client';
    case ImplementationLocation.Method:
      return 'method';
    default:
      throw new Error(`unhandled parameter location ${param.implementation}`);
  }
}

function adaptParameterType(param: M4Parameter): ParameterType {
  if (param.clientDefaultValue) {
    const adaptedType = adaptPossibleType(param.schema);
    if (!isLiteralValueType(adaptedType)) {
      throw new Error(`unsupported client side default type ${getTypeDeclaration(adaptedType)} for parameter ${param.language.go!.name}`);
    }
    return new ClientSideDefault(new LiteralValue(adaptedType, param.clientDefaultValue));
  } else if (param.schema.type === SchemaType.Constant) {
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

function findOrAdaptParamsGroup(param: M4Parameter): ParameterGroup {
  const groupProp = <GroupProperty>param.language.go!.paramGroup;
  let paramGroup = paramGroups.get(groupProp.schema.language.go!.name);
  if (!paramGroup) {
    paramGroup = adaptParameterGroup(adaptParameterlocation(param), groupProp);
    paramGroups.set(groupProp.schema.language.go!.name, paramGroup);
  }

  return paramGroup;
}

function adaptParameterGroup(location: ParameterLocation, groupProp: GroupProperty): ParameterGroup {
  const paramGroup = new ParameterGroup(groupProp.language.go!.name, groupProp.schema.language.go!.name, groupProp.required === true, location);
  paramGroup.description = groupProp.language.go!.description;
  return paramGroup;
}

// returns the media type used by the protocol
function adaptBodyFormat(protocol: Protocols): BodyFormat {
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

function adaptAnyResult(op: Operation): AnyResult {
  const resultTypes: Record<number, PossibleType> = {};
  for (const resp of values(op.responses)) {
    let possibleType: PossibleType;
    if (isSchemaResponse(resp)) {
      possibleType = adaptPossibleType(resp.schema);
    } else {
      // the operation contains a mix of schemas and non-schema responses
      continue;
    }

    resultTypes[parseInt(resp.protocol.http!.statusCodes)] = possibleType;
  }

  return new AnyResult('Value', 'JSON', resultTypes);
}

function adaptResultFormat(protocol: Protocols): ResultFormat {
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
