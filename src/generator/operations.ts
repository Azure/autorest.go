/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment, KnownMediaType, pascalCase, camelCase } from '@azure-tools/codegen'
import { ArraySchema, ByteArraySchema, CodeModel, ConstantSchema, DateTimeSchema, DictionarySchema, GroupProperty, ImplementationLocation, NumberSchema, Operation, OperationGroup, Parameter, Property, Protocols, Response, Schema, SchemaResponse, SchemaType, SerializationStyle } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { aggregateParameters, isArraySchema, isPageableOperation, isSchemaResponse, PagerInfo, isLROOperation } from '../common/helpers';
import { OperationNaming } from '../transform/namer';
import { contentPreamble, formatParameterTypeName, hasDescription, skipURLEncoding, sortAscending, sortParametersByRequired, getCreateRequestParametersSig, getMethodParameters } from './helpers';
import { ImportManager } from './imports';

const dateFormat = '2006-01-02';
const datetimeRFC3339Format = 'time.RFC3339';
const datetimeRFC1123Format = 'time.RFC1123';

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
  // generate protocol operations
  const operations = new Array<OperationGroupContent>();
  for (const group of values(session.model.operationGroups)) {
    // the list of packages to import
    const imports = new ImportManager();
    // add standard imorts
    imports.add('net/http');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');

    const clientName = camelCase(group.language.go!.clientName);
    let opText = '';
    group.operations.sort((a: Operation, b: Operation) => { return sortAscending(a.language.go!.name, b.language.go!.name) });
    for (const op of values(group.operations)) {
      // protocol creation can add imports to the list so
      // it must be done before the imports are written out
      opText += generateOperation(clientName, op, imports);
      opText += createProtocolRequest(clientName, op, imports);
      opText += createProtocolResponse(clientName, op, imports);
      opText += createProtocolErrHandler(clientName, op, imports);
    }
    const interfaceText = createInterfaceDefinition(group, imports);
    // stitch it all together
    let text = await contentPreamble(session);
    const exportClient = await session.getValue('export-client', true);
    let client = 'Client';
    if (!exportClient) {
      client = 'client';
    }
    text += imports.text();
    text += interfaceText;
    text += `// ${clientName} implements the ${group.language.go!.clientName} interface.\n`;
    text += `type ${clientName} struct {\n`;
    text += `\t*${client}\n`;
    if (group.language.go!.clientParams) {
      const clientParams = <Array<Parameter>>group.language.go!.clientParams;
      for (const clientParam of values(clientParams)) {
        text += `\t${clientParam.language.go!.name} ${formatParameterTypeName(clientParam)}\n`;
      }
    }
    text += '}\n\n';
    text += opText;

    operations.push(new OperationGroupContent(group.language.go!.name, text));
  }
  return operations;
}

function formatParamValue(param: Parameter, imports: ImportManager): string {
  let separator = ',';
  switch (param.protocol.http?.style) {
    case SerializationStyle.PipeDelimited:
      separator = '|';
      break;
    case SerializationStyle.SpaceDelimited:
      separator = ' ';
      break;
    case SerializationStyle.TabDelimited:
      separator = '\\t';
      break;
  }
  let paramName = param.language.go!.name;
  if (param.implementation === ImplementationLocation.Client) {
    paramName = `client.${paramName}`;
  } else if (param.language.go!.paramGroup) {
    paramName = `${camelCase(param.language.go!.paramGroup.language.go!.name)}.${pascalCase(paramName)}`;
  }
  if (param.required !== true) {
    paramName = `*${paramName}`;
  }
  switch (param.schema.type) {
    case SchemaType.Array:
      const arraySchema = <ArraySchema>param.schema;
      switch (arraySchema.elementType.type) {
        case SchemaType.String:
          imports.add('strings');
          return `strings.Join(${paramName}, "${separator}")`;
        default:
          imports.add('fmt');
          imports.add('strings');
          return `strings.Join(strings.Fields(strings.Trim(fmt.Sprint(${paramName}), "[]")), "${separator}")`;
      }
    case SchemaType.Boolean:
      imports.add('strconv');
      return `strconv.FormatBool(${paramName})`;
    case SchemaType.ByteArray:
      // ByteArray is a base-64 encoded value in string format
      imports.add('encoding/base64');
      let byteFormat = 'Std';
      if ((<ByteArraySchema>param.schema).format === 'base64url') {
        byteFormat = 'RawURL';
      }
      return `base64.${byteFormat}Encoding.EncodeToString(${paramName})`;
    case SchemaType.Choice:
    case SchemaType.SealedChoice:
      return `string(${paramName})`;
    case SchemaType.Constant:
      const constSchema = <ConstantSchema>param.schema;
      // cannot use formatConstantValue() since all values are treated as strings
      return `"${constSchema.value.value}"`;
    case SchemaType.Date:
      if (param.required !== true && paramName[0] === '*') {
        // remove the dereference
        paramName = paramName.substr(1);
      }
      return `${paramName}.Format("${dateFormat}")`;
    case SchemaType.DateTime:
      imports.add('time');
      if (param.required !== true && paramName[0] === '*') {
        // remove the dereference
        paramName = paramName.substr(1);
      }
      let format = datetimeRFC3339Format;
      const dateTime = <DateTimeSchema>param.schema;
      if (dateTime.format === 'date-time-rfc1123') {
        format = datetimeRFC1123Format;
      }
      return `${paramName}.Format(${format})`;
    case SchemaType.Duration:
      if (param.required !== true && paramName[0] === '*') {
        // remove the dereference
        paramName = paramName.substr(1);
      }
      return `${paramName}.String()`;
    case SchemaType.UnixTime:
      return `timeUnix(${paramName}).String()`;
    case SchemaType.Uri:
      imports.add('net/url');
      if (param.required !== true && paramName[0] === '*') {
        // remove the dereference
        paramName = paramName.substr(1);
      }
      return `${paramName}.String()`;
    case SchemaType.Integer:
      imports.add('strconv');
      const intSchema = <NumberSchema>param.schema;
      let intParam = paramName;
      if (intSchema.precision === 32) {
        intParam = `int64(${intParam})`;
      }
      return `strconv.FormatInt(${intParam}, 10)`;
    case SchemaType.Number:
      imports.add('strconv');
      const numberSchema = <NumberSchema>param.schema;
      let floatParam = paramName;
      if (numberSchema.precision === 32) {
        floatParam = `float64(${floatParam})`;
      }
      return `strconv.FormatFloat(${floatParam}, 'f', -1, ${numberSchema.precision})`;
    default:
      return paramName;
  }
}

// use this to generate the code that will help process values returned in response headers
function formatHeaderResponseValue(propName: string, header: string, schema: Schema, imports: ImportManager, respObj: string): string {
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
    case SchemaType.Duration:
      imports.add('time');
      text += `\t\t${name}, err := time.ParseDuration(val)\n`;
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

function generateOperation(clientName: string, op: Operation, imports: ImportManager): string {
  if (op.language.go!.paging && op.language.go!.paging.isNextOp) {
    // don't generate a public API for the methods used to advance pages
    return '';
  }
  const info = <OperationNaming>op.language.go!;
  const params = getAPIParametersSig(op, imports);
  const returns = generateReturnsInfo(op, false);
  let text = '';
  if (hasDescription(op.language.go!)) {
    text += `// ${op.language.go!.name} - ${op.language.go!.description} \n`;
  }
  if (isMultiRespOperation(op)) {
    text += generateMultiRespComment(op);
  }
  if (op.language.go!.methodPrefix) {
    text += `func (client *${clientName}) ${op.language.go!.methodPrefix}${op.language.go!.name}(${params}) (${returns.join(', ')}) {\n`;
  } else {
    text += `func (client *${clientName}) ${op.language.go!.name}(${params}) (${returns.join(', ')}) {\n`;
  }
  // split param list into individual params
  const reqParams = getCreateRequestParametersSig(op).split(',');
  // keep the parameter names from the name/type tuples
  for (let i = 0; i < reqParams.length; ++i) {
    reqParams[i] = reqParams[i].trim().split(' ')[0];
  }
  if (isLROOperation(op)) {
    imports.add('time');
    text += `\treq, err := client.${info.protocolNaming.requestMethod}(${reqParams.join(', ')})\n`;
    text += `\tif err != nil {\n`;
    text += `\t\treturn nil, err\n`;
    text += `\t}\n`;
    text += `\t// send the first request to initialize the poller\n`;
    text += `\tresp, err := client.p.Do(ctx, req)\n`;
    text += `\tif err != nil {\n`;
    text += `\t\treturn nil, err\n`;
    text += `\t}\n`;
    if (!op.responses) {
      text += '\tresult := &HTTPResponse{\n';
      text += '\t\tresult.RawResponse: resp\n';
      text += '\t}\n';
    } else {
      text += `\tresult, err := client.${info.protocolNaming.responseMethod}(resp)\n`;
      text += `\tif err != nil {\n`;
      text += `\t\treturn nil, err\n`;
      text += `\t}\n`;
    }
    // LRO operation might have a special configuration set in x-ms-long-running-operation-options
    // which indicates a specific url to perform the final Get operation on
    let finalState = '';
    if (op.extensions?.['x-ms-long-running-operation-options']?.['final-state-via']) {
      finalState = op.extensions?.['x-ms-long-running-operation-options']?.['final-state-via'];
    }
    text += `\tpt, err := createPollingTracker("${clientName}.${op.language.go!.name}", "${finalState}", resp, client.${info.protocolNaming.errorMethod})\n`;
    text += '\tif err != nil {\n';
    text += '\t\treturn nil, err\n';
    text += '\t}\n';
    text += `\tpoller := &${camelCase(op.language.go!.pollerType.name)}{\n`;
    text += '\t\t\tpt: pt,\n';
    if (isPageableOperation(op)) {
      text += `\t\t\trespHandler: client.${camelCase(op.language.go!.pageableType.name)}HandleResponse,\n`;
    }
    text += '\t\t\tpipeline: client.p,\n';
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
    text += addResumePollerMethod(op, clientName);
    return text;
  }
  text += `\treq, err := client.${info.protocolNaming.requestMethod}(${reqParams.join(', ')})\n`;
  text += `\tif err != nil {\n`;
  text += `\t\treturn nil, err\n`;
  text += `\t}\n`;
  if (isPageableOperation(op)) {
    text += `\treturn &${camelCase(op.language.go!.pageableType.name)}{\n`;
    text += `\t\tpipeline: client.p,\n`;
    text += `\t\trequest: req,\n`;
    text += `\t\tresponder: client.${info.protocolNaming.responseMethod},\n`;
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
      text += `\t\tadvancer: func(resp *${schemaResponse.schema.language.go!.responseType.name}) (*azcore.Request, error) {\n`;
      text += `\t\t\treturn client.${camelCase(op.language.go!.paging.member)}CreateRequest(${nextOpParams.join(', ')})\n`;
      text += '\t\t},\n';
    } else {
      imports.add('fmt');
      imports.add('net/url');
      let resultTypeName = schemaResponse.schema.language.go!.name;
      if (schemaResponse.schema.serialization?.xml?.name) {
        // xml can specifiy its own name, prefer that if available
        resultTypeName = schemaResponse.schema.serialization.xml.name;
      }
      text += `\t\tadvancer: func(resp *${schemaResponse.schema.language.go!.responseType.name}) (*azcore.Request, error) {\n`;
      text += `\t\t\tu, err := url.Parse(*resp.${resultTypeName}.${nextLink})\n`;
      text += `\t\t\tif err != nil {\n`;
      text += `\t\t\t\treturn nil, fmt.Errorf("invalid ${nextLink}: %w", err)\n`;
      text += `\t\t\t}\n`;
      text += `\t\t\tif u.Scheme == "" {\n`;
      text += `\t\t\t\treturn nil, fmt.Errorf("no scheme detected in ${nextLink} %s", *resp.${resultTypeName}.${nextLink})\n`;
      text += `\t\t\t}\n`;
      text += `\t\t\treturn azcore.NewRequest(http.MethodGet, *u), nil\n`;
      text += `\t\t},\n`;
    }
    text += `\t}, nil\n`;
    text += '}\n\n';
    return text;
  }
  text += `\tresp, err := client.p.Do(ctx, req)\n`;
  text += `\tif err != nil {\n`;
  text += `\t\treturn nil, err\n`;
  text += `\t}\n`;
  // also cheating here as at present the only param to the responder is an azcore.Response
  text += `\tresult, err := client.${info.protocolNaming.responseMethod}(resp)\n`;
  text += `\tif err != nil {\n`;
  text += `\t\treturn nil, err\n`;
  text += `\t}\n`;
  text += `\treturn result, nil\n`;
  text += '}\n\n';
  return text;
}

function createProtocolRequest(client: string, op: Operation, imports: ImportManager): string {
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
  text += `func (client *${client}) ${name}(${getCreateRequestParametersSig(op)}) (${returns.join(', ')}) {\n`;
  const inPathParams = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'path'; });
  // storage needs the client.u to be the source-of-truth for the full path.
  // however, swagger requires that all operations specify a path, which is at odds with storage.
  // to work around this, storage specifies x-ms-path paths with path params but doesn't
  // actually reference the path params (i.e. no params with which to replace the tokens).
  // so, if a path contains tokens but there are no path params, skip emitting the path.
  let includeParse = false;
  const pathStr = <string>op.requests![0].protocol.http!.path;
  const pathContainsParms = pathStr.includes('{');
  if (!pathContainsParms && pathStr.length > 1) {
    // path does NOT include path params and is not "/", emit it
    text += `\turlPath := "${op.requests![0].protocol.http!.path}"\n`;
    includeParse = true;
  } else if (inPathParams.any()) {
    // swagger defines path params, emit path and replace tokens
    imports.add('strings');
    text += `\turlPath := "${op.requests![0].protocol.http!.path}"\n`;
    // replace path parameters
    for (const pp of values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'path'; })) {
      let paramValue = `url.PathEscape(${formatParamValue(pp, imports)})`;
      if (skipURLEncoding(pp)) {
        paramValue = formatParamValue(pp, imports);
      } else {
        imports.add('net/url');
      }
      text += `\turlPath = strings.ReplaceAll(urlPath, "{${pp.language.go!.serializedName}}", ${paramValue})\n`;
    }
    includeParse = true;
  }
  if (includeParse) {
    text += `\tu, err := client.u.Parse(urlPath)\n`;
    text += '\tif err != nil {\n';
    text += '\t\treturn nil, err\n';
    text += '\t}\n';
  } else {
    text += `\tu := client.u\n`;
  }
  // helper to build nil checks for param groups
  const emitParamGroupCheck = function (gp: GroupProperty, param: Parameter): string {
    const paramGroupName = camelCase(gp.language.go!.name);
    let optionalParamGroupCheck = `${paramGroupName} != nil && `;
    if (gp.required) {
      optionalParamGroupCheck = '';
    }
    return `\tif ${optionalParamGroupCheck}${paramGroupName}.${pascalCase(param.language.go!.name)} != nil {\n`;
  }
  const inQueryParams = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'query'; });
  if (inQueryParams.any()) {
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
      } else if (qp.schema.type === SchemaType.Constant) {
        // omit this query param. TODO once non-required constants are fixed
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
      text += '\tquery := u.Query()\n';
      for (const qp of values(encodedParams)) {
        text += emitQueryParam(qp, `query.Set("${qp.language.go!.serializedName}", ${formatParamValue(qp, imports)})`);
      }
      text += '\tu.RawQuery = query.Encode()\n';
    }
    // tack on any unencoded params to the end
    if (unencodedParams.length > 0) {
      if (encodedParams.length > 0) {
        text += '\tunencodedParams := []string{u.RawQuery}\n';
      } else {
        text += '\tunencodedParams := []string{}\n';
      }
      for (const qp of values(unencodedParams)) {
        text += emitQueryParam(qp, `unencodedParams = append(unencodedParams, "${qp.language.go!.serializedName}="+${formatParamValue(qp, imports)})`);
      }
      text += '\tu.RawQuery = strings.Join(unencodedParams, "&")\n';
    }
  }
  text += `\treq := azcore.NewRequest(http.Method${pascalCase(op.requests![0].protocol.http!.method)}, *u)\n`;
  if (hasBinaryResponse(op.responses!)) {
    // skip auto-body downloading for binary stream responses
    text += '\treq.SkipBodyDownload()\n';
  }
  // add specific request headers
  const headerParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined; }).where((each: Parameter) => { return each.protocol.http!.in === 'header'; });
  headerParam.forEach(header => {
    if (header.required) {
      text += `\treq.Header.Set("${header.language.go!.serializedName}", ${formatParamValue(header, imports)})\n`;
    } else if (header.schema.type === SchemaType.Constant) {
      // omit this header. TODO once non-required constants are fixed
    } else {
      text += emitParamGroupCheck(<GroupProperty>header.language.go!.paramGroup, header);
      text += `\t\treq.Header.Set("${header.language.go!.serializedName}", ${formatParamValue(header, imports)})\n`;
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
    // TODO once non-required constants are fixed
    if (bodyParam!.required || bodyParam?.schema.type === SchemaType.Constant) {
      text += `\treturn req, req.MarshalAs${getMediaFormat(bodyParam!.schema, mediaType, body)}\n`;
    } else {
      const paramGroup = <GroupProperty>bodyParam!.language.go!.paramGroup;
      text += `\tif ${camelCase(paramGroup.language.go!.name)} != nil {\n`;
      text += `\t\treturn req, req.MarshalAs${getMediaFormat(bodyParam!.schema, mediaType, body)}\n`;
      text += '\t}\n';
      text += '\treturn req, nil\n';
    }
  } else if (mediaType === 'binary') {
    const bodyParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http!.in === 'body'; }).first();
    text += `\treturn req, req.SetBody(${bodyParam?.language.go!.name})\n`;
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

function createProtocolResponse(client: string, op: Operation, imports: ImportManager): string {
  const info = <OperationNaming>op.language.go!;
  const name = info.protocolNaming.responseMethod;
  let text = `${comment(name, '// ')} handles the ${info.name} response.\n`;
  text += `func (client *${client}) ${name}(resp *azcore.Response) (${generateReturnsInfo(op, true).join(', ')}) {\n`;
  if (!op.responses) {
    text += `\treturn nil, client.${info.protocolNaming.errorMethod}(resp)`;
    text += '}\n\n';
    return text;
  }
  const generateResponseUnmarshaller = function (response: Response, isLRO: boolean): string {
    let unmarshallerText = '';
    if (!isSchemaResponse(response)) {
      if (isLRO) {
        unmarshallerText += '\treturn &HTTPPollerResponse{RawResponse: resp.Response}, nil\n';
        return unmarshallerText;
      }
      // no response body, return the *http.Response
      unmarshallerText += `\treturn resp.Response, nil\n`;
      return unmarshallerText;
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
    if (isLRO) {
      unmarshallerText += `\treturn &${schemaResponse.schema.language.go!.lroResponseType.language.go!.name}{RawResponse: resp.Response}, nil\n`;
      return unmarshallerText;
    }
    let respObj = `${schemaResponse.schema.language.go!.responseType.name}{RawResponse: resp.Response}`;
    unmarshallerText += `\tresult := ${respObj}\n`;
    // assign any header values
    for (const prop of values(<Array<Property>>schemaResponse.schema.language.go!.properties)) {
      if (prop.language.go!.fromHeader) {
        unmarshallerText += formatHeaderResponseValue(prop.language.go!.name, prop.language.go!.fromHeader, prop.schema, imports, 'result');
      }
    }
    const mediaType = getMediaType(response.protocol);
    if (mediaType === 'none' || mediaType === 'binary') {
      // nothing to unmarshal
      unmarshallerText += '\treturn &result, nil\n';
      return unmarshallerText;
    }
    let target = `result.${schemaResponse.schema.language.go!.responseType.value}`;
    // when unmarshalling a wrapped XML array or discriminated type, unmarshal into the response type, not the field
    if ((mediaType === 'XML' && schemaResponse.schema.type === SchemaType.Array) || schemaResponse.schema.language.go!.discriminatorInterface) {
      target = 'result';
    }
    unmarshallerText += `\treturn &result, resp.UnmarshalAs${getMediaFormat(response.schema, mediaType, `&${target}`)}\n`;
    return unmarshallerText;
  };
  if (!isMultiRespOperation(op)) {
    // concat all status codes that return the same schema into one array.
    // this is to support operations that specify multiple response codes
    // that return the same schema (or no schema).
    let statusCodes = new Array<string>();
    for (let i = 0; i < op.responses.length; ++i) {
      statusCodes = statusCodes.concat(op.responses[i].protocol.http?.statusCodes);
    }
    if (isLROOperation(op) && statusCodes.find(element => element === '204') === undefined) {
      statusCodes = statusCodes.concat('204');
    }
    text += `\tif !resp.HasStatusCode(${formatStatusCodes(statusCodes)}) {\n`;
    text += `\t\treturn nil, client.${info.protocolNaming.errorMethod}(resp)\n`;
    text += '\t}\n';
    if (isLROOperation(op) && isPageableOperation(op)) {
      text += generateResponseUnmarshaller(op.responses![0], true);
      text += '}\n\n';
      text += `${comment(name, '// ')} handles the ${info.name} response.\n`;
      text += `func (client *${client}) ${camelCase(op.language.go!.pageableType.name)}HandleResponse(resp *azcore.Response) (*${(<SchemaResponse>op.responses![0]).schema.language.go!.responseType.value}Response, error) {\n`;
      const index = statusCodes.indexOf('204');
      if (index > -1) {
        statusCodes.splice(index, 1);
      }
      statusCodes.push('200');
      text += `\tif !resp.HasStatusCode(${formatStatusCodes(statusCodes)}) {\n`;
      text += `\t\treturn nil, client.${info.protocolNaming.errorMethod}(resp)\n`;
      text += '\t}\n';
      text += generateResponseUnmarshaller(op.responses![0], false);
    } else {
      text += generateResponseUnmarshaller(op.responses![0], isLROOperation(op));
    }
  } else {
    text += '\tswitch resp.StatusCode {\n';
    for (const response of values(op.responses)) {
      text += `\tcase ${formatStatusCodes(response.protocol.http!.statusCodes)}:\n`
      text += generateResponseUnmarshaller(response, isLROOperation(op));
    }
    text += '\tdefault:\n';
    text += `\t\treturn nil, client.${info.protocolNaming.errorMethod}(resp)\n`;
    text += '\t}\n';
  }
  text += '}\n\n';
  return text;
}

function createProtocolErrHandler(client: string, op: Operation, imports: ImportManager): string {
  const info = <OperationNaming>op.language.go!;
  const name = info.protocolNaming.errorMethod;
  let text = `${comment(name, '// ')} handles the ${info.name} error response.\n`;
  text += `func (client *${client}) ${name}(resp *azcore.Response) error {\n`;

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
      return errors.New(resp.Status)
    }
    return errors.New(string(body))
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
      unmarshaller += `${prefix}return err.wrapped\n`;
    } else if (schemaError.type === SchemaType.Object) {
      unmarshaller += `${prefix}return err\n`;
    } else {
      imports.add('fmt');
      unmarshaller += `${prefix}return fmt.Errorf("%v", err)\n`;
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
  let interfaceText = `// ${group.language.go!.clientName} contains the methods for the ${group.language.go!.name} group.\n`;
  interfaceText += `type ${group.language.go!.clientName} interface {\n`;
  for (const op of values(group.operations)) {
    let opName = op.language.go!.name;
    if (op.language.go!.paging && op.language.go!.paging.isNextOp) {
      // don't generate a public API for the methods used to advance pages
      continue;
    }
    if (op.language.go!.methodPrefix) {
      opName = `${op.language.go!.methodPrefix}${opName}`;
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
    const returns = generateReturnsInfo(op, false);
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
function getMediaType(protocol: Protocols): 'JSON' | 'XML' | 'binary' | 'none' {
  // TODO: binary, forms etc
  switch (protocol.http!.knownMediaType) {
    case KnownMediaType.Json:
      return 'JSON';
    case KnownMediaType.Xml:
      return 'XML';
    case KnownMediaType.Binary:
      return 'binary';
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

function formatStatusCodes(statusCodes: Array<string>): string {
  const asHTTPStatus = new Array<string>();
  for (const rawCode of values(statusCodes)) {
    asHTTPStatus.push(formatStatusCode(rawCode));
  }
  return asHTTPStatus.join(', ');
}

function formatStatusCode(statusCode: string): string {
  switch (statusCode) {
    case '200':
      return 'http.StatusOK';
    case '201':
      return 'http.StatusCreated';
    case '202':
      return 'http.StatusAccepted';
    case '204':
      return 'http.StatusNoContent';
    case '206':
      return 'http.StatusPartialContent';
    case '300':
      return 'http.StatusMultipleChoices';
    case '301':
      return 'http.StatusMovedPermanently';
    case '302':
      return 'http.StatusFound';
    case '303':
      return 'http.StatusSeeOther';
    case '304':
      return 'http.StatusNotModified';
    case '307':
      return 'http.StatusTemporaryRedirect';
    case '400':
      return 'http.StatusBadRequest';
    case '404':
      return 'http.StatusNotFound';
    case '409':
      return 'http.StatusConflict';
    case '500':
      return 'http.StatusInternalServerError';
    case '501':
      return 'http.StatusNotImplemented';
    default:
      throw console.error(`unhandled status code ${statusCode}`);
  }
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
function generateReturnsInfo(op: Operation, forHandler: boolean): string[] {
  if (!op.responses) {
    return ['*http.Response', 'error'];
  }
  let returnType = '*http.Response';
  if (isMultiRespOperation(op)) {
    returnType = 'interface{}';
  } else {
    const firstResp = <SchemaResponse>op.responses![0];
    // must check pageable first as all pageable operations are also schema responses, 
    // but LRO operations that return a pager are an exception and need to return LRO specific
    // responses
    if (!forHandler && isPageableOperation(op) && !isLROOperation(op)) {
      returnType = op.language.go!.pageableType.name;
    } else if (isSchemaResponse(firstResp)) {
      returnType = '*' + firstResp.schema.language.go!.responseType.name;
      if (isLROOperation(op)) {
        returnType = '*' + firstResp.schema.language.go!.lroResponseType.language.go!.name;
      }
    } else if (isLROOperation(op)) {
      returnType = '*HTTPPollerResponse';
    }
  }
  return [returnType, 'error'];
}

function addResumePollerMethod(op: Operation, clientName: string): string {
  const info = <OperationNaming>op.language.go!;
  let text = `func (client *${clientName}) Resume${op.language.go!.name}(token string) (${op.language.go!.pollerType.name}, error) {\n`;
  text += `\tpt, err := resumePollingTracker("${clientName}.${op.language.go!.name}", token, client.${info.protocolNaming.errorMethod})\n`;
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
