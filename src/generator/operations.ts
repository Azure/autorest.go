/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment, KnownMediaType, pascalCase, camelCase } from '@azure-tools/codegen';
import { ArraySchema, ByteArraySchema, CodeModel, ConstantSchema, DateTimeSchema, DictionarySchema, GroupProperty, ImplementationLocation, NumberSchema, Operation, OperationGroup, Parameter, Property, Protocols, Response, Schema, SchemaResponse, SchemaType } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { aggregateParameters, isArraySchema, isPageableOperation, isSchemaResponse, PagerInfo, isLROOperation, exportClients } from '../common/helpers';
import { OperationNaming } from '../transform/namer';
import { contentPreamble, formatParameterTypeName, formatStatusCodes, getStatusCodes, hasDescription, hasSchemaResponse, skipURLEncoding, sortAscending, getCreateRequestParameters, getCreateRequestParametersSig, getMethodParameters, getParamName, formatParamValue, dateFormat, datetimeRFC1123Format, datetimeRFC3339Format, sortParametersByRequired } from './helpers';
import { ImportManager } from './imports';

// represents the generated content for an operation group
export class OperationGroupContent {
  readonly name: string;
  readonly content: string;

  constructor(name: string, content: string) {
    this.name = name;
    this.content = content;
  }
}

// Creates the content for all <operation>.go files
export async function generateOperations(session: Session<CodeModel>): Promise<OperationGroupContent[]> {
  const isARM = session.model.language.go!.openApiType === 'arm';
  // generate protocol operations
  const operations = new Array<OperationGroupContent>();
  for (const group of values(session.model.operationGroups)) {
    // the list of packages to import
    const imports = new ImportManager();
    // add standard imorts
    imports.add('net/http');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');

    let opText = '';
    group.operations.sort((a: Operation, b: Operation) => { return sortAscending(a.language.go!.name, b.language.go!.name) });
    for (const op of values(group.operations)) {
      // protocol creation can add imports to the list so
      // it must be done before the imports are written out
      if (isARM && isLROOperation(op)) {
        // generate Begin and Resume methods
        opText += generateARMLROBeginMethod(op, imports);
        opText += generateARMLROResumeMethod(op);
      }
      opText += generateOperation(op, imports);
      opText += createProtocolRequest(session.model, op, imports);
      opText += createProtocolResponse(op, imports);
      opText += createProtocolErrHandler(op, imports);
    }
    let interfaceText = '';
    if (isARM) {
      interfaceText = createInterfaceDefinition(group, imports);
    }
    // stitch it all together
    let text = await contentPreamble(session);
    const exportClient = await exportClients(session);
    let client = 'Client';
    let clientName = group.language.go!.clientName;
    if (!exportClient) {
      client = 'client';
    }
    const clientCtor = group.language.go!.clientCtorName;
    text += imports.text();
    text += interfaceText;
    // generate the operation client
    const interfaceName = group.language.go!.interfaceName;
    if (isARM) {
      text += `// ${clientName} implements the ${interfaceName} interface.\n`;
      text += `// Don't use this type directly, use ${clientCtor}() instead.\n`;
    }
    text += `type ${clientName} struct {\n`;
    text += `\t*${client}\n`;
    if (group.language.go!.clientParams) {
      const clientParams = <Array<Parameter>>group.language.go!.clientParams;
      for (const clientParam of values(clientParams)) {
        text += `\t${clientParam.language.go!.name} ${formatParameterTypeName(clientParam)}\n`;
      }
    }
    text += '}\n\n';
    if (isARM) {
      // operation client constructor
      const clientLiterals = [`${client}: c`];
      const methodParams = [`c *${client}`];
      // add client params to the operation client constructor
      if (group.language.go!.clientParams) {
        const clientParams = <Array<Parameter>>group.language.go!.clientParams;
        clientParams.sort(sortParametersByRequired);
        for (const clientParam of values(clientParams)) {
          clientLiterals.push(`${clientParam.language.go!.name}: ${clientParam.language.go!.name}`);
          methodParams.push(`${clientParam.language.go!.name} ${formatParameterTypeName(clientParam)}`);
        }
      }
      text += `// ${clientCtor} creates a new instance of ${clientName} with the specified values.\n`;
      text += `func ${clientCtor}(${methodParams.join(', ')}) ${interfaceName} {\n`;
      text += `\treturn &${clientName}{${clientLiterals.join(', ')}}\n`;
      text += '}\n\n';
    }
    // operation client Do method
    text += '// Do invokes the Do() method on the pipeline associated with this client.\n';
    text += `func (client *${clientName}) Do(req *azcore.Request) (*azcore.Response, error) {\n`;
    text += '\treturn client.p.Do(req)\n';
    text += '}\n\n';
    // add operations content last
    text += opText;
    operations.push(new OperationGroupContent(group.language.go!.name, text));
  }
  return operations;
}

// use this to generate the code that will help process values returned in response headers
function formatHeaderResponseValue(propName: string, header: string, schema: Schema, imports: ImportManager, respObj: string): string {
  // dictionaries are handled slightly different so we do that first
  if (schema.type === SchemaType.Dictionary) {
    imports.add('strings');
    let text = '\tfor hh := range resp.Header {\n';
    text += `\t\tif strings.HasPrefix(hh, "${schema.language.go!.headerCollectionPrefix}") {\n`;
    text += `\t\t\tif ${respObj}.Metadata == nil {\n`;
    text += `\t\t\t\t${respObj}.Metadata = &map[string]string{}\n`;
    text += '\t\t\t}\n';
    text += `\t\t\t(*${respObj}.Metadata)[hh[len("${schema.language.go!.headerCollectionPrefix}"):]] = resp.Header.Get(hh)\n`;
    text += '\t\t}\n';
    text += '\t}\n';
    return text;
  }
  let text = `\tif val := resp.Header.Get("${header}"); val != "" {\n`;
  const name = camelCase(propName);
  switch (schema.type) {
    case SchemaType.Boolean:
      imports.add('strconv');
      text += `\t\t${name}, err := strconv.ParseBool(val)\n`;
      break;
    case SchemaType.ByteArray:
      // ByteArray is a base-64 encoded value in string format
      imports.add('encoding/base64');
      let byteFormat = 'Std';
      if ((<ByteArraySchema>schema).format === 'base64url') {
        byteFormat = 'RawURL';
      }
      text += `\t\t${name}, err := base64.${byteFormat}Encoding.DecodeString(val)\n`;
      break;
    case SchemaType.Choice:
    case SchemaType.SealedChoice:
      text += `\t\t${respObj}.${propName} = (*${schema.language.go!.name})(&val)\n`;
      text += '\t}\n';
      return text;
    case SchemaType.Constant:
    case SchemaType.Duration:
    case SchemaType.String:
      text += `\t\t${respObj}.${propName} = &val\n`;
      text += '\t}\n';
      return text;
    case SchemaType.Date:
      imports.add('time');
      text += `\t\t${name}, err := time.Parse("${dateFormat}", val)\n`;
      break;
    case SchemaType.DateTime:
      imports.add('time');
      let format = datetimeRFC3339Format;
      const dateTime = <DateTimeSchema>schema;
      if (dateTime.format === 'date-time-rfc1123') {
        format = datetimeRFC1123Format;
      }
      text += `\t\t${name}, err := time.Parse(${format}, val)\n`;
      break;
    case SchemaType.Integer:
      imports.add('strconv');
      const intNum = <NumberSchema>schema;
      if (intNum.precision === 32) {
        text += `\t\t${name}32, err := strconv.ParseInt(val, 10, 32)\n`;
        text += `\t\t${name} := int32(${name}32)\n`;
      } else {
        text += `\t\t${name}, err := strconv.ParseInt(val, 10, 64)\n`;
      }
      break;
    case SchemaType.Number:
      imports.add('strconv');
      const floatNum = <NumberSchema>schema;
      if (floatNum.precision === 32) {
        text += `\t\t${name}32, err := strconv.ParseFloat(val, 32)\n`;
        text += `\t\t${name} := float32(${name}32)\n`;
      } else {
        text += `\t\t${name}, err := strconv.ParseFloat(val, 64)\n`;
      }
      break;
    default:
      throw console.error(`unsupported header type ${schema.type}`);
  }
  text += `\t\tif err != nil {\n`;
  text += `\t\t\treturn nil, err\n`;
  text += `\t\t}\n`;
  text += `\t\t${respObj}.${propName} = &${name}\n`;
  text += '\t}\n';
  return text;
}

function generateMultiRespComment(op: Operation): string {
  const returnTypes = new Array<string>();
  for (const response of values(op.responses)) {
    returnTypes.push(`*${(<SchemaResponse>response).schema.language.go!.responseType.name}`);
  }
  return `// Possible return types are ${returnTypes.join(', ')}\n`;
}

function generateOperation(op: Operation, imports: ImportManager): string {
  if (op.language.go!.paging && op.language.go!.paging.isNextOp) {
    // don't generate a public API for the methods used to advance pages
    return '';
  }
  const info = <OperationNaming>op.language.go!;
  const params = getAPIParametersSig(op, imports);
  const returns = generateReturnsInfo(op, 'op');
  const clientName = op.language.go!.clientName;
  let text = '';
  if (hasDescription(op.language.go!)) {
    text += `// ${op.language.go!.name} - ${op.language.go!.description} \n`;
  }
  if (isMultiRespOperation(op)) {
    text += generateMultiRespComment(op);
  }
  text += `func (client *${clientName}) ${op.language.go!.name}(${params}) (${returns.join(', ')}) {\n`;
  const reqParams = getCreateRequestParameters(op);
  if (isPageableOperation(op) && !isLROOperation(op)) {
    imports.add('context');
    text += `\treturn &${camelCase(op.language.go!.pageableType.name)}{\n`;
    text += `\t\tpipeline: client.p,\n`;
    text += `\t\trequester: func(ctx context.Context) (*azcore.Request, error) {\n`;
    text += `\t\t\treturn client.${info.protocolNaming.requestMethod}(${reqParams})\n`;
    text += '\t\t},\n';
    text += `\t\tresponder: client.${info.protocolNaming.responseMethod},\n`;
    text += `\t\terrorer:   client.${info.protocolNaming.errorMethod},\n`;
    const pager = <PagerInfo>op.language.go!.pageableType;
    const schemaResponse = <SchemaResponse>pager.op.responses![0];
    const nextLink = pager.op.language.go!.paging.nextLinkName;
    if (op.language.go!.paging.member) {
      const nextOpParams = getCreateRequestParametersSig(op.language.go!.paging.nextLinkOperation).split(',');
      // keep the parameter names from the name/type tuples and find nextLink param
      for (let i = 0; i < nextOpParams.length; ++i) {
        const paramName = nextOpParams[i].trim().split(' ')[0];
        const paramType = nextOpParams[i].trim().split(' ')[1];
        if (paramName.startsWith('next') && paramType === 'string') {
          nextOpParams[i] = `*resp.${schemaResponse.schema.language.go!.name}.${nextLink}`;
        } else {
          nextOpParams[i] = paramName;
        }
      }
      text += `\t\tadvancer: func(ctx context.Context, resp *${schemaResponse.schema.language.go!.responseType.name}) (*azcore.Request, error) {\n`;
      text += `\t\t\treturn client.${op.language.go!.paging.member}CreateRequest(${nextOpParams.join(', ')})\n`;
      text += '\t\t},\n';
    } else {
      let resultTypeName = schemaResponse.schema.language.go!.name;
      if (schemaResponse.schema.serialization?.xml?.name) {
        // xml can specifiy its own name, prefer that if available
        resultTypeName = schemaResponse.schema.serialization.xml.name;
      }
      text += `\t\tadvancer: func(ctx context.Context, resp *${schemaResponse.schema.language.go!.responseType.name}) (*azcore.Request, error) {\n`;
      text += `\t\t\treturn azcore.NewRequest(ctx, http.MethodGet, *resp.${resultTypeName}.${nextLink})\n`;
      text += `\t\t},\n`;
    }
    text += `\t}\n`;
    text += '}\n\n';
    return text;
  }
  text += `\treq, err := client.${info.protocolNaming.requestMethod}(${reqParams})\n`;
  text += `\tif err != nil {\n`;
  text += `\t\treturn nil, err\n`;
  text += `\t}\n`;
  text += `\tresp, err := client.Do(req)\n`;
  text += `\tif err != nil {\n`;
  text += `\t\treturn nil, err\n`;
  text += `\t}\n`;
  const statusCodes = getStatusCodes(op);
  text += `\tif !resp.HasStatusCode(${formatStatusCodes(statusCodes)}) {\n`;
  text += `\t\treturn nil, client.${info.protocolNaming.errorMethod}(resp)\n`;
  text += '\t}\n';
  if (isLROOperation(op)) {
    text += '\t return resp, nil\n';
  } else if (needsResponseHandler(op)) {
    // also cheating here as at present the only param to the responder is an azcore.Response
    text += `\tresult, err := client.${info.protocolNaming.responseMethod}(resp)\n`;
    text += `\tif err != nil {\n`;
    text += `\t\treturn nil, err\n`;
    text += `\t}\n`;
    text += `\treturn result, nil\n`;
  } else {
    text += '\treturn resp.Response, nil\n';
  }
  text += '}\n\n';
  return text;
}

function createProtocolRequest(codeModel: CodeModel, op: Operation, imports: ImportManager): string {
  const info = <OperationNaming>op.language.go!;
  const name = info.protocolNaming.requestMethod;
  for (const param of values(aggregateParameters(op))) {
    if (param.implementation !== ImplementationLocation.Method || param.required !== true) {
      continue;
    }
    imports.addImportForSchemaType(param.schema);
  }
  const returns = ['*azcore.Request', 'error'];
  let text = `${comment(name, '// ')} creates the ${info.name} request.\n`;
  text += `func (client *${op.language.go!.clientName}) ${name}(${getCreateRequestParametersSig(op)}) (${returns.join(', ')}) {\n`;
  // default to host on the client
  let hostParam = 'client.u';
  if (codeModel.language.go!.complexHostParams) {
    imports.add('strings');
    // we have a complex parameterized host
    text += `\thost := "${op.requests![0].protocol.http!.uri}"\n`;
    // get all the host params on the client
    const hostParams = <Array<Parameter>>codeModel.language.go!.hostParams;
    for (const hostParam of values(hostParams)) {
      text += `\thost = strings.ReplaceAll(host, "{${hostParam.language.go!.serializedName}}", client.${hostParam.language.go!.name})\n`;
    }
    // check for any method local host params
    for (const param of values(op.parameters)) {
      if (param.implementation === ImplementationLocation.Method && param.protocol.http!.in === 'uri') {
        text += `\thost = strings.ReplaceAll(host, "{${param.language.go!.serializedName}}", ${param.language.go!.name})\n`;
      }
    }
    hostParam = 'host';
  }
  const hasPathParams = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'path'; }).any();
  // storage needs the client.u to be the source-of-truth for the full path.
  // however, swagger requires that all operations specify a path, which is at odds with storage.
  // to work around this, storage specifies x-ms-path paths with path params but doesn't
  // actually reference the path params (i.e. no params with which to replace the tokens).
  // so, if a path contains tokens but there are no path params, skip emitting the path.
  const pathStr = <string>op.requests![0].protocol.http!.path;
  const pathContainsParms = pathStr.includes('{');
  if (hasPathParams || (!pathContainsParms && pathStr.length > 1)) {
    // there are path params, or the path doesn't contain tokens and is not "/" so emit it
    text += `\turlPath := "${op.requests![0].protocol.http!.path}"\n`;
    hostParam = `azcore.JoinPaths(${hostParam}, urlPath)`;
  }
  if (hasPathParams) {
    // swagger defines path params, emit path and replace tokens
    imports.add('strings');
    // replace path parameters
    for (const pp of values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'path'; })) {
      let paramValue = formatParamValue(pp, imports);
      if (!skipURLEncoding(pp)) {
        imports.add('net/url');
        paramValue = `url.PathEscape(${formatParamValue(pp, imports)})`;
      }
      text += `\turlPath = strings.ReplaceAll(urlPath, "{${pp.language.go!.serializedName}}", ${paramValue})\n`;
    }
  }
  text += `\treq, err := azcore.NewRequest(ctx, http.Method${pascalCase(op.requests![0].protocol.http!.method)}, ${hostParam})\n`;
  text += '\tif err != nil {\n';
  text += '\t\treturn nil, err\n';
  text += '\t}\n';
  const hasQueryParams = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'query'; }).any();
  // helper to build nil checks for param groups
  const emitParamGroupCheck = function (gp: GroupProperty, param: Parameter): string {
    const paramGroupName = camelCase(gp.language.go!.name);
    let optionalParamGroupCheck = `${paramGroupName} != nil && `;
    if (gp.required) {
      optionalParamGroupCheck = '';
    }
    return `\tif ${optionalParamGroupCheck}${paramGroupName}.${pascalCase(param.language.go!.name)} != nil {\n`;
  }
  if (hasQueryParams) {
    // add query parameters
    const encodedParams = new Array<Parameter>();
    const unencodedParams = new Array<Parameter>();
    for (const qp of values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'query'; })) {
      if (skipURLEncoding(qp)) {
        unencodedParams.push(qp);
      } else {
        encodedParams.push(qp);
      }
    }
    const emitQueryParam = function (qp: Parameter, setter: string): string {
      let qpText = '';
      if (qp.required === true) {
        qpText = `\t${setter}\n`;
      } else if (qp.implementation === ImplementationLocation.Client) {
        // global optional param
        qpText = `\tif client.${qp.language.go!.name} != nil {\n`;
        qpText += `\t\t${setter}\n`;
        qpText += `\t}\n`;
      } else {
        qpText = emitParamGroupCheck(<GroupProperty>qp.language.go!.paramGroup, qp);
        qpText += `\t\t${setter}\n`;
        qpText += `\t}\n`;
      }
      return qpText;
    }
    // emit encoded params first
    if (encodedParams.length > 0) {
      text += '\tquery := req.URL.Query()\n';
      for (const qp of values(encodedParams)) {
        let setter: string;
        if (qp.protocol.http?.explode === true) {
          setter = `\tfor _, qv := range ${getParamName(qp)} {\n`;
          setter += `\t\tquery.Add("${qp.language.go!.serializedName}", qv)\n`;
          setter += '\t}';
        } else {
          // cannot initialize setter to this value as formatParamValue() can change imports
          setter = `query.Set("${qp.language.go!.serializedName}", ${formatParamValue(qp, imports)})`;
        }
        text += emitQueryParam(qp, setter);
      }
      text += '\treq.URL.RawQuery = query.Encode()\n';
    }
    // tack on any unencoded params to the end
    if (unencodedParams.length > 0) {
      if (encodedParams.length > 0) {
        text += '\tunencodedParams := []string{req.URL.RawQuery}\n';
      } else {
        text += '\tunencodedParams := []string{}\n';
      }
      for (const qp of values(unencodedParams)) {
        let setter: string;
        if (qp.protocol.http?.explode === true) {
          setter = `\tfor _, qv := range ${getParamName(qp)} {\n`;
          setter += `\t\tunencodedParams = append(unencodedParams, "${qp.language.go!.serializedName}="+qv)\n`;
          setter += '\t}';
        } else {
          setter = `unencodedParams = append(unencodedParams, "${qp.language.go!.serializedName}="+${formatParamValue(qp, imports)})`;
        }
        text += emitQueryParam(qp, `unencodedParams = append(unencodedParams, "${qp.language.go!.serializedName}="+${formatParamValue(qp, imports)})`);
      }
      text += '\treq.URL.RawQuery = strings.Join(unencodedParams, "&")\n';
    }
  }
  if (hasBinaryResponse(op.responses!)) {
    // skip auto-body downloading for binary stream responses
    text += '\treq.SkipBodyDownload()\n';
  }
  // add specific request headers
  const headerParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined; }).where((each: Parameter) => { return each.protocol.http!.in === 'header'; });
  headerParam.forEach(header => {
    const emitHeaderSet = function (headerParam: Parameter, prefix: string): string {
      if (header.schema.language.go!.headerCollectionPrefix) {
        let headerText = `${prefix}for k, v := range ${getParamName(headerParam)} {\n`;
        headerText += `${prefix}\treq.Header.Set("${header.schema.language.go!.headerCollectionPrefix}"+k, v)\n`;
        headerText += `${prefix}}\n`;
        return headerText;
      } else {
        return `${prefix}req.Header.Set("${headerParam.language.go!.serializedName}", ${formatParamValue(headerParam, imports)})\n`;
      }
    }
    if (header.required) {
      text += emitHeaderSet(header, '\t');
    } else {
      text += emitParamGroupCheck(<GroupProperty>header.language.go!.paramGroup, header);
      text += emitHeaderSet(header, '\t\t');
      text += `\t}\n`;
    }
  });
  const mediaType = getMediaType(op.requests![0].protocol);
  if (mediaType === 'JSON' || mediaType === 'XML') {
    const bodyParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http!.in === 'body'; }).first();
    // default to the body param name
    let body = bodyParam!.language.go!.name;
    if (bodyParam!.language.go!.paramGroup) {
      const paramGroup = <GroupProperty>bodyParam!.language.go!.paramGroup;
      body = `${camelCase(paramGroup.language.go!.name)}.${pascalCase(bodyParam!.language.go!.name)}`;
    }
    if (bodyParam!.schema.type === SchemaType.Constant) {
      // if the value is constant, embed it directly
      body = formatConstantValue(<ConstantSchema>bodyParam!.schema);
    } else if (mediaType === 'XML' && bodyParam!.schema.type === SchemaType.Array) {
      // for XML payloads, create a wrapper type if the payload is an array
      imports.add('encoding/xml');
      text += '\ttype wrapper struct {\n';
      let tagName = bodyParam!.schema.language.go!.name;
      if (bodyParam!.schema.serialization?.xml?.name) {
        tagName = bodyParam!.schema.serialization.xml.name;
      }
      text += `\t\tXMLName xml.Name \`xml:"${tagName}"\`\n`;
      let fieldName = bodyParam!.schema.language.go!.name;
      if (isArraySchema(bodyParam!.schema)) {
        fieldName = pascalCase(bodyParam!.language.go!.name);
        let tag = bodyParam!.schema.elementType.language.go!.name;
        if (bodyParam!.schema.elementType.serialization?.xml?.name) {
          tag = bodyParam!.schema.elementType.serialization.xml.name;
        }
        text += `\t\t${fieldName} *${bodyParam!.schema.language.go!.name} \`xml:"${tag}"\`\n`;
      }
      text += '\t}\n';
      let addr = '&';
      if (!bodyParam?.required) {
        addr = '';
      }
      body = `wrapper{${fieldName}: ${addr}${body}}`;
    } else if ((bodyParam!.schema.type === SchemaType.DateTime && (<DateTimeSchema>bodyParam!.schema).format === 'date-time-rfc1123') || bodyParam!.schema.type === SchemaType.UnixTime) {
      // wrap the body in the custom RFC1123 type
      text += `\taux := ${bodyParam!.schema.language.go!.internalTimeType}(${body})\n`;
      body = 'aux';
    } else if (isArrayOfTimesForMarshalling(bodyParam!.schema)) {
      const timeType = (<ArraySchema>bodyParam!.schema).elementType.language.go!.internalTimeType;
      text += `\taux := make([]${timeType}, len(${body}), len(${body}))\n`;
      text += `\tfor i := 0; i < len(${body}); i++ {\n`;
      text += `\t\taux[i] = ${timeType}(${body}[i])\n`;
      text += '\t}\n';
      body = 'aux';
    } else if (isMapOfDateTime(bodyParam!.schema)) {
      const timeType = (<ArraySchema>bodyParam!.schema).elementType.language.go!.internalTimeType;
      text += `\taux := map[string]${timeType}{}\n`;
      text += `\tfor k, v := range ${body} {\n`;
      text += `\t\taux[k] = ${timeType}(v)\n`;
      text += '\t}\n';
      body = 'aux';
    }
    if (bodyParam!.required || bodyParam!.schema.type === SchemaType.Constant) {
      text += `\treturn req, req.MarshalAs${getMediaFormat(bodyParam!.schema, mediaType, body)}\n`;
    } else {
      const paramGroup = <GroupProperty>bodyParam!.language.go!.paramGroup;
      text += `\tif ${camelCase(paramGroup.language.go!.name)} != nil {\n`;
      text += `\t\treturn req, req.MarshalAs${getMediaFormat(bodyParam!.schema, mediaType, body)}\n`;
      text += '\t}\n';
      text += '\treturn req, nil\n';
    }
  } else if (mediaType === 'binary') {
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
    const bodyParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http!.in === 'body'; }).first();
    text += `\treturn req, req.SetBody(${bodyParam?.language.go!.name}, ${contentType})\n`;
  } else if (mediaType === 'text') {
    imports.add('strings');
    let bodyParam = '';
    for (const param of values(op.requests![0].parameters)) {
      if (param.protocol.http!.in === 'body') {
        bodyParam = param.language.go!.name;
      }
    }
    text += `\tbody := azcore.NopCloser(strings.NewReader(${bodyParam}))\n`;
    text += `\treturn req, req.SetBody(body, "text/plain; encoding=UTF-8")\n`;
  } else {
    text += `\treturn req, nil\n`;
  }
  text += '}\n\n';
  return text;
}

function getMediaFormat(schema: Schema, mediaType: 'JSON' | 'XML', param: string): string {
  let marshaller: 'JSON' | 'XML' | 'ByteArray' = mediaType;
  let format = '';
  if (schema.type === SchemaType.ByteArray) {
    marshaller = 'ByteArray';
    format = ', azcore.Base64StdFormat';
    if ((<ByteArraySchema>schema).format === 'base64url') {
      format = ', azcore.Base64URLFormat';
    }
  }
  return `${marshaller}(${param}${format})`;
}

function isArrayOfTimesForMarshalling(schema: Schema): boolean {
  if (schema.type !== SchemaType.Array) {
    return false;
  }
  const arraySchema = <ArraySchema>schema;
  const arrayElem = <Schema>arraySchema.elementType;
  if (arrayElem.type === SchemaType.UnixTime) {
    return true;
  }
  if (arrayElem.type !== SchemaType.DateTime) {
    return false;
  }
  return (<DateTimeSchema>arrayElem).format === 'date-time-rfc1123';
}

function needsResponseHandler(op: Operation): boolean {
  return hasSchemaResponse(op) || (isLROOperation(op) && hasSchemaResponse(op)) || isPageableOperation(op);
}

function generateResponseUnmarshaller(response: Response, imports: ImportManager): string {
  let unmarshallerText = '';
  if (!isSchemaResponse(response)) {
    throw console.error('TODO');
  } else if (response.schema.type === SchemaType.DateTime || response.schema.type === SchemaType.UnixTime) {
    // use the designated time type for unmarshalling
    unmarshallerText += `\tvar aux *${response.schema.language.go!.internalTimeType}\n`;
    unmarshallerText += `\terr := resp.UnmarshalAs${getMediaType(response.protocol)}(&aux)\n`;
    const resp = `${response.schema.language.go!.responseType.name}{RawResponse: resp.Response, ${response.schema.language.go!.responseType.value}: (*time.Time)(aux)}`;
    unmarshallerText += `\treturn &${resp}, err\n`;
    return unmarshallerText;
  } else if (isArrayOfDateTime(response.schema)) {
    // unmarshalling arrays of date/time is a little more involved
    unmarshallerText += `\tvar aux *[]${(<ArraySchema>response.schema).elementType.language.go!.internalTimeType}\n`;
    unmarshallerText += `\tif err := resp.UnmarshalAs${getMediaType(response.protocol)}(&aux); err != nil {\n`;
    unmarshallerText += '\t\treturn nil, err\n';
    unmarshallerText += '\t}\n';
    unmarshallerText += '\tcp := make([]time.Time, len(*aux), len(*aux))\n';
    unmarshallerText += '\tfor i := 0; i < len(*aux); i++ {\n';
    unmarshallerText += '\t\tcp[i] = time.Time((*aux)[i])\n';
    unmarshallerText += '\t}\n';
    const resp = `${response.schema.language.go!.responseType.name}{RawResponse: resp.Response, ${response.schema.language.go!.responseType.value}: &cp}`;
    unmarshallerText += `\treturn &${resp}, nil\n`;
    return unmarshallerText;
  } else if (isMapOfDateTime(response.schema)) {
    unmarshallerText += `\taux := map[string]${(<DictionarySchema>response.schema).elementType.language.go!.internalTimeType}{}\n`;
    unmarshallerText += `\tif err := resp.UnmarshalAs${getMediaType(response.protocol)}(&aux); err != nil {\n`;
    unmarshallerText += '\t\treturn nil, err\n';
    unmarshallerText += '\t}\n';
    unmarshallerText += `\tcp := map[string]time.Time{}\n`;
    unmarshallerText += `\tfor k, v := range aux {\n`;
    unmarshallerText += `\t\tcp[k] = time.Time(v)\n`;
    unmarshallerText += `\t}\n`;
    const resp = `${response.schema.language.go!.responseType.name}{RawResponse: resp.Response, ${response.schema.language.go!.responseType.value}: &cp}`;
    unmarshallerText += `\treturn &${resp}, nil\n`;
    return unmarshallerText;
  }
  const schemaResponse = <SchemaResponse>response;
  let respObj = `${schemaResponse.schema.language.go!.responseType.name}{RawResponse: resp.Response}`;
  unmarshallerText += `\tresult := ${respObj}\n`;
  // assign any header values
  for (const prop of values(<Array<Property>>schemaResponse.schema.language.go!.properties)) {
    if (prop.language.go!.fromHeader) {
      unmarshallerText += formatHeaderResponseValue(prop.language.go!.name, prop.language.go!.fromHeader, prop.schema, imports, 'result');
    }
  }
  const mediaType = getMediaType(response.protocol);
  if (mediaType === 'JSON' || mediaType === 'XML') {
    let target = `result.${schemaResponse.schema.language.go!.responseType.value}`;
    // when unmarshalling a wrapped XML array or discriminated type, unmarshal into the response type, not the field
    if ((mediaType === 'XML' && schemaResponse.schema.type === SchemaType.Array) || schemaResponse.schema.language.go!.discriminatorInterface) {
      target = 'result';
    }
    unmarshallerText += `\treturn &result, resp.UnmarshalAs${getMediaFormat(response.schema, mediaType, `&${target}`)}\n`;
    return unmarshallerText;
  }
  // nothing to unmarshal
  unmarshallerText += '\treturn &result, nil\n';
  return unmarshallerText;
}

function createProtocolResponse(op: Operation, imports: ImportManager): string {
  if (!needsResponseHandler(op)) {
    return '';
  }
  const info = <OperationNaming>op.language.go!;
  const name = info.protocolNaming.responseMethod;
  const clientName = op.language.go!.clientName;
  let text = `${comment(name, '// ')} handles the ${info.name} response.\n`;
  text += `func (client *${clientName}) ${name}(resp *azcore.Response) (${generateReturnsInfo(op, 'handler').join(', ')}) {\n`;
  if (!isMultiRespOperation(op)) {
    text += generateResponseUnmarshaller(op.responses![0], imports);
  } else {
    imports.add('fmt');
    text += '\tswitch resp.StatusCode {\n';
    for (const response of values(op.responses)) {
      text += `\tcase ${formatStatusCodes(response.protocol.http!.statusCodes)}:\n`
      text += generateResponseUnmarshaller(response, imports);
    }
    text += '\tdefault:\n';
    text += `\t\treturn nil, fmt.Errorf("unhandled HTTP status code %d", resp.StatusCode)\n`;
    text += '\t}\n';
  }
  text += '}\n\n';
  return text;
}

function createProtocolErrHandler(op: Operation, imports: ImportManager): string {
  const info = <OperationNaming>op.language.go!;
  const name = info.protocolNaming.errorMethod;
  let text = `${comment(name, '// ')} handles the ${info.name} error response.\n`;
  text += `func (client *${op.language.go!.clientName}) ${name}(resp *azcore.Response) error {\n`;
  // define a generic error for when there are no exceptions or no error schema
  const generateGenericError = function () {
    imports.add('errors');
    imports.add('io/ioutil');
    imports.add('fmt');
    return `body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
    }
    if len(body) == 0 {
      return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
    }
    return azcore.NewResponseError(errors.New(string(body)), resp.Response)
    `;
  }

  // if the response doesn't define any error types return a generic error
  if (!op.exceptions) {
    text += generateGenericError();
    text += '}\n\n';
    return text;
  }

  const generateUnmarshaller = function (exception: Response, prefix: string) {
    let unmarshaller = '';
    if (exception.language.go!.genericError) {
      unmarshaller += `${prefix}${generateGenericError()}`;
      return unmarshaller;
    }
    const schemaError = (<SchemaResponse>exception).schema;
    const errFormat = <string>schemaError.language.go!.marshallingFormat;
    let typeName = schemaError.language.go!.name;
    if (schemaError.language.go!.internalErrorType) {
      typeName = schemaError.language.go!.internalErrorType;
    }
    unmarshaller += `var err ${typeName}\n`;
    unmarshaller += `${prefix}if err := resp.UnmarshalAs${errFormat.toUpperCase()}(&err); err != nil {\n`;
    unmarshaller += `${prefix}\treturn err\n`;
    unmarshaller += `${prefix}}\n`;
    if (schemaError.language.go!.internalErrorType) {
      // err.wrapped is for discriminated error types, it will already be pointer-to-type
      unmarshaller += `${prefix}return azcore.NewResponseError(err.wrapped, resp.Response)\n`;
    } else if (schemaError.type === SchemaType.Object) {
      // for consistency with success responses, return pointer-to-error type
      unmarshaller += `${prefix}return azcore.NewResponseError(&err, resp.Response)\n`;
    } else {
      imports.add('fmt');
      unmarshaller += `${prefix}return azcore.NewResponseError(fmt.Errorf("%v", err), resp.Response)\n`;
    }
    return unmarshaller;
  };
  if (op.exceptions.length === 1) {
    text += generateUnmarshaller(op.exceptions![0], '\t');
    text += '}\n\n';
    return text;
  }
  text += '\tswitch resp.StatusCode {\n';
  for (const exception of values(op.exceptions)) {
    for (const statusCode of values(<Array<string>>exception.protocol.http!.statusCodes)) {
      if (statusCode === 'default') {
        text += '\tdefault:\n';
        text += generateUnmarshaller(exception, '\t\t');
      } else {
        text += `\tcase ${formatStatusCodes([statusCode])}:\n`;
        text += generateUnmarshaller(exception, '\t\t');
      }
    }
  }
  text += '\t}\n';
  text += '}\n\n';
  return text;
}

function isArrayOfDateTime(schema: Schema): boolean {
  if (schema.type !== SchemaType.Array) {
    return false;
  }
  const arraySchema = <ArraySchema>schema;
  const arrayElem = <Schema>arraySchema.elementType;
  return arrayElem.type === SchemaType.DateTime || arrayElem.type === SchemaType.UnixTime;
}

function isMapOfDateTime(schema: Schema): boolean {
  if (schema.type !== SchemaType.Dictionary) {
    return false;
  }
  const dictSchema = <DictionarySchema>schema;
  const dictElem = <Schema>dictSchema.elementType;
  return dictElem.type === SchemaType.DateTime || dictElem.type === SchemaType.UnixTime;
}

function createInterfaceDefinition(group: OperationGroup, imports: ImportManager): string {
  let interfaceText = `// ${group.language.go!.interfaceName} contains the methods for the ${group.language.go!.name} group.\n`;
  interfaceText += `type ${group.language.go!.interfaceName} interface {\n`;
  for (const op of values(group.operations)) {
    let opName = op.language.go!.name;
    if (op.language.go!.paging && op.language.go!.paging.isNextOp) {
      // don't generate a public API for the methods used to advance pages
      continue;
    }
    if (isLROOperation(op)) {
      opName = `Begin${opName}`;
    }
    for (const param of values(aggregateParameters(op))) {
      if (param.implementation !== ImplementationLocation.Method || param.required !== true) {
        continue;
      }
      imports.addImportForSchemaType(param.schema);
    }
    if (hasDescription(op.language.go!)) {
      interfaceText += `\t// ${opName} - ${op.language.go!.description} \n`;
    }
    if (isMultiRespOperation(op)) {
      interfaceText += generateMultiRespComment(op);
    }
    const returns = generateReturnsInfo(op, 'int');
    interfaceText += `\t${opName}(${getAPIParametersSig(op, imports)}) (${returns.join(', ')})\n`;
    // Add resume LRO poller method for each Begin poller method
    if (isLROOperation(op)) {
      interfaceText += `\t// Resume${op.language.go!.name} - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.\n`;
      interfaceText += `\tResume${op.language.go!.name}(token string) (${op.language.go!.pollerType.name}, error)\n`;
    }
  }
  interfaceText += '}\n\n';
  return interfaceText;
}

// returns the media type used by the protocol
function getMediaType(protocol: Protocols): 'JSON' | 'XML' | 'binary' | 'text' | 'none' {
  // TODO: binary, forms etc
  switch (protocol.http!.knownMediaType) {
    case KnownMediaType.Json:
      return 'JSON';
    case KnownMediaType.Xml:
      return 'XML';
    case KnownMediaType.Binary:
      return 'binary';
    case KnownMediaType.Text:
      return 'text';
    default:
      return 'none';
  }
}

function formatConstantValue(schema: ConstantSchema) {
  // null check must come before any type checks
  if (schema.value.value === null) {
    return 'nil';
  }
  if (schema.valueType.type === SchemaType.String) {
    return `"${schema.value.value}"`;
  }
  return schema.value.value;
}

// returns true if any responses are a binary stream
function hasBinaryResponse(responses: Response[]): boolean {
  for (const resp of values(responses)) {
    if (resp.protocol.http!.knownMediaType === KnownMediaType.Binary) {
      return true;
    }
  }
  return false;
}

// returns the parameters for the public API
// e.g. "ctx context.Context, i int, s string"
function getAPIParametersSig(op: Operation, imports: ImportManager): string {
  const methodParams = getMethodParameters(op);
  const params = new Array<string>();
  if (!isPageableOperation(op) || isLROOperation(op)) {
    imports.add('context');
    params.push('ctx context.Context');
  }
  for (const methodParam of values(methodParams)) {
    params.push(`${camelCase(methodParam.language.go!.name)} ${formatParameterTypeName(methodParam)}`);
  }
  return params.join(', ');
}

// returns the return signature where each entry is the type name
// e.g. [ '*string', 'error' ]
// apiType describes where the return sig is used.
//   int - for the interface definition
//    op - for the operation
// handler - for the response handler
function generateReturnsInfo(op: Operation, apiType: 'int' | 'op' | 'handler'): string[] {
  if (!op.responses) {
    return ['*http.Response', 'error'];
  }
  let returnType = '*http.Response';
  if (isLROOperation(op)) {
    switch (apiType) {
      case 'handler':
        returnType = '*' + (<SchemaResponse>op.responses![0]).schema.language.go!.responseType.name;
        break;
      case 'int':
        returnType = '*HTTPPollerResponse';
        if (hasSchemaResponse(op)) {
          returnType = '*' + (<SchemaResponse>op.responses![0]).schema.language.go!.lroResponseType.language.go!.name;
        }
        break;
      case 'op':
        returnType = '*azcore.Response';
        break;
    }
  } else if (isPageableOperation(op)) {
    switch (apiType) {
      case 'handler':
        returnType = '*' + (<SchemaResponse>op.responses![0]).schema.language.go!.responseType.name;
        break;
      case 'int':
      case 'op':
        // pager operations don't return an error
        return [op.language.go!.pageableType.name];
    }
  } else if (isMultiRespOperation(op)) {
    returnType = 'interface{}';
  } else if (hasSchemaResponse(op)) {
    // simple schema response
    returnType = '*' + (<SchemaResponse>op.responses![0]).schema.language.go!.responseType.name;
  }
  return [returnType, 'error'];
}

function generateARMLROBeginMethod(op: Operation, imports: ImportManager): string {
  const info = <OperationNaming>op.language.go!;
  const params = getAPIParametersSig(op, imports);
  const returns = generateReturnsInfo(op, 'int');
  const clientName = op.language.go!.clientName;
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/armcore');
  imports.add('time');
  let text = `func (client *${clientName}) Begin${op.language.go!.name}(${params}) (${returns.join(', ')}) {\n`;
  text += `\tresp, err := client.${op.language.go!.name}(${getCreateRequestParameters(op)})\n`;
  text += `\tif err != nil {\n`;
  text += `\t\treturn nil, err\n`;
  text += `\t}\n`;
  if (!op.responses || !isSchemaResponse(op.responses![0])) {
    text += '\tresult := &HTTPPollerResponse{\n';
  } else {
    text += `\tresult := &${(<SchemaResponse>op.responses![0]).schema.language.go!.lroResponseType.language.go!.name}{\n`;
  }
  text += '\t\tRawResponse: resp.Response,\n';
  text += '\t}\n';
  // LRO operation might have a special configuration set in x-ms-long-running-operation-options
  // which indicates a specific url to perform the final Get operation on
  let finalState = '';
  if (op.extensions?.['x-ms-long-running-operation-options']?.['final-state-via']) {
    finalState = op.extensions?.['x-ms-long-running-operation-options']?.['final-state-via'];
  }
  text += `\tpt, err := armcore.NewPoller("${clientName}.${op.language.go!.name}", "${finalState}", resp, client.${info.protocolNaming.errorMethod})\n`;
  text += '\tif err != nil {\n';
  text += '\t\treturn nil, err\n';
  text += '\t}\n';
  text += `\tpoller := &${camelCase(op.language.go!.pollerType.name)}{\n`;
  text += '\t\tpt: pt,\n';
  if (isPageableOperation(op)) {
    const statusCodes = getStatusCodes(op);
    if (statusCodes.indexOf('200') < 0) {
      statusCodes.push('200');
    }
    if (statusCodes.indexOf('204') < 0) {
      statusCodes.push('204');
    }
    statusCodes.sort();
    text += `\t\terrHandler: func(resp *azcore.Response) error {\n`;
    text += `\t\t\tif resp.HasStatusCode(${formatStatusCodes(statusCodes)}) {\n`;
    text += `\t\t\t\treturn nil\n`;
    text += '\t\t\t}\n';
    text += `\t\t\treturn client.${info.protocolNaming.errorMethod}(resp)\n`;
    text += '\t\t},\n';
    text += `\t\trespHandler: func(resp *azcore.Response) (*${(<SchemaResponse>op.responses![0]).schema.language.go!.responseType.name}, error) {\n`;
    text += generateResponseUnmarshaller(op.responses![0], imports);
    text += '\t\t},\n';
  }
  text += '\t\tpipeline: client.p,\n';
  text += '\t}\n';
  text += '\tresult.Poller = poller\n';
  // determine the poller response based on the name and whether is is a pageable operation
  let pollerResponse = '*http.Response';
  if (isPageableOperation(op)) {
    pollerResponse = op.language.go!.pageableType.name;
  } else if (isSchemaResponse(op.responses![0])) {
    pollerResponse = '*' + (<SchemaResponse>op.responses![0]).schema.language.go!.responseType.name;
  }
  text += `\tresult.PollUntilDone = func(ctx context.Context, frequency time.Duration) (${pollerResponse}, error) {\n`;
  text += `\t\treturn poller.pollUntilDone(ctx, frequency)\n`;
  text += `\t}\n`;
  text += `\treturn result, nil\n`;
  // closing braces
  text += '}\n\n';
  return text;
}

function generateARMLROResumeMethod(op: Operation): string {
  const info = <OperationNaming>op.language.go!;
  const clientName = op.language.go!.clientName;
  let text = `func (client *${clientName}) Resume${op.language.go!.name}(token string) (${op.language.go!.pollerType.name}, error) {\n`;
  text += `\tpt, err := armcore.NewPollerFromResumeToken("${clientName}.${op.language.go!.name}", token, client.${info.protocolNaming.errorMethod})\n`;
  text += '\tif err != nil {\n';
  text += '\t\treturn nil, err\n';
  text += '\t}\n';
  text += `\treturn &${camelCase(op.language.go!.pollerType.name)}{\n`;
  text += '\t\tpipeline: client.p,\n';
  text += '\t\tpt: pt,\n';
  text += '\t}, nil\n';
  text += '}\n\n';
  return text;
}

// returns true if the operation returns multiple response types
function isMultiRespOperation(op: Operation): boolean {
  // treat LROs as single-response ops
  if (!op.responses || op.responses?.length === 1 || isLROOperation(op)) {
    return false;
  }
  // count the number of schemas returned by this operation
  let schemaCount = 0;
  let currentResp = op.responses![0];
  if (isSchemaResponse(currentResp)) {
    ++schemaCount;
  }
  // check that all response types are identical
  for (let i = 1; i < op.responses!.length; ++i) {
    const response = op.responses![i];
    if (isSchemaResponse(response) && isSchemaResponse(currentResp)) {
      // both are schema responses, ensure they match
      if ((<SchemaResponse>response).schema !== (<SchemaResponse>currentResp).schema) {
        ++schemaCount;
      }
    } else if (isSchemaResponse(response) && !isSchemaResponse(currentResp)) {
      ++schemaCount;
      // update currentResp to this response so we can compare it against the remaining responses
      currentResp = response;
    }
  }
  return schemaCount > 1;
}
