/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment, KnownMediaType, pascalCase, camelCase } from '@azure-tools/codegen'
import { ArraySchema, CodeModel, ConstantSchema, DateTimeSchema, DictionarySchema, ImplementationLocation, NumberSchema, Operation, OperationGroup, Parameter, Property, Protocols, Response, Schema, SchemaResponse, SchemaType, SerializationStyle } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { aggregateParameters, isArraySchema, isPageableOperation, isSchemaResponse, PagerInfo, isLROOperation } from '../common/helpers';
import { OperationNaming } from '../transform/namer';
import { contentPreamble, formatParameterTypeName, hasDescription, skipURLEncoding, sortAscending, sortParametersByRequired } from './helpers';
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
    }
    const interfaceText = createInterfaceDefinition(group, imports);
    // stitch it all together
    let text = await contentPreamble(session);
    text += imports.text();
    text += interfaceText;
    text += `// ${clientName} implements the ${group.language.go!.clientName} interface.\n`;
    text += `type ${clientName} struct{\n`;
    text += '\t*Client\n';
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
  }
  if (param.required !== true) {
    if (param.implementation === ImplementationLocation.Method) {
      // optional params at the method level will be in an options struct
      paramName = `*options.${pascalCase(paramName)}`;
    } else {
      paramName = `*${paramName}`;
    }
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
      return `base64.StdEncoding.EncodeToString(${paramName})`;
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
    case SchemaType.UnixTime:
      if (param.required !== true && paramName[0] === '*') {
        // remove the dereference
        paramName = paramName.substr(1);
      }
      return `${paramName}.String()`;
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
      text += `\t\t${name}, err := base64.StdEncoding.DecodeString(val)\n`;
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

function generateOperation(clientName: string, op: Operation, imports: ImportManager): string {
  if (isPageableOperation(op) && op.language.go!.paging.member === op.language.go!.name) {
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
  // split param list into individual params
  const reqParams = getCreateRequestParametersSig(op).split(',');
  // slice off the parameter names from the type/type tuples
  for (let i = 0; i < reqParams.length; ++i) {
    reqParams[i] = reqParams[i].trim().split(' ')[0];
  }
  if (isLROOperation(op)) {
    text += `func (client *${clientName}) Begin${op.language.go!.name}(${params}) (${returns.join(', ')}) {\n`;
    // TODO uncomment the following code to actually implement polling 
    // text += `\treq, err := client.${info.protocolNaming.requestMethod}(${reqParams.join(', ')})\n`;
    // text += `\tif err != nil {\n`;
    // text += `\t\treturn nil, err\n`;
    // text += `\t}\n`;
    // text += `\t// send the first request to initialize the poller\n`;
    // text += `\tresp, err := client.p.Do(ctx, req)\n`;
    // text += `\tif err != nil {\n`;
    // text += `\t\treturn nil, err\n`;
    // text += `\t}\n`;
    // text += `\tpt := pollingTracker${pascalCase(op.requests![0].protocol.http!.method)}{\n`;
    // text += `\t\tpollingTrackerBase: pollingTrackerBase{\n`;
    // text += `\t\t\tresp: resp,\n`;
    // text += `\t\t}}\n`;
    // text += `\terr = pt.initializeState()\n`;
    // text += `\tif err != nil {\n`;
    // text += `\t\treturn nil, err\n`;
    // text += `\t}\n`;
    // closing braces
    text += `\treturn &${op.language.go!.pollerType.name}{\n`;
    // text += `\t\tpt: &pt,\n`;
    text += `\t\tclient: client,\n`;
    text += `\t}, nil\n`;
    text += '}\n\n';
    return text;
  }
  text += `func (client *${clientName}) ${op.language.go!.name}(${params}) (${returns.join(', ')}) {\n`;
  text += `\treq, err := client.${info.protocolNaming.requestMethod}(${reqParams.join(', ')})\n`;
  text += `\tif err != nil {\n`;
  text += `\t\treturn nil, err\n`;
  text += `\t}\n`;
  if (isPageableOperation(op)) {
    text += `\treturn &${camelCase(op.language.go!.pageableType.name)}{\n`;
    text += `\t\tclient: client,\n`;
    text += `\t\trequest: req,\n`;
    text += `\t\tresponder: client.${info.protocolNaming.responseMethod},\n`;
    const pager = <PagerInfo>op.language.go!.pageableType;
    if (op.language.go!.paging.member) {
      reqParams.push(`*resp.${pager.schema.language.go!.name}.${pager.nextLink}`);
      text += `\t\tadvancer: func(resp *${pager.schema.language.go!.responseType.name}) (*azcore.Request, error) {\n`;
      text += `\t\t\treturn client.${camelCase(op.language.go!.paging.member)}CreateRequest(${reqParams.join(', ')})\n`;
      text += '\t\t},\n';
    } else {
      imports.add('fmt');
      imports.add('net/url');
      let resultTypeName = pager.schema.language.go!.name;
      if (pager.schema.serialization?.xml?.name) {
        // xml can specifiy its own name, prefer that if available
        resultTypeName = pager.schema.serialization.xml.name;
      }
      text += `\t\tadvancer: func(resp *${pager.schema.language.go!.responseType.name}) (*azcore.Request, error) {\n`;
      text += `\t\t\tu, err := url.Parse(*resp.${resultTypeName}.${pager.nextLink})\n`;
      text += `\t\t\tif err != nil {\n`;
      text += `\t\t\t\treturn nil, fmt.Errorf("invalid ${pager.nextLink}: %w", err)\n`;
      text += `\t\t\t}\n`;
      text += `\t\t\tif u.Scheme == "" {\n`;
      text += `\t\t\t\treturn nil, fmt.Errorf("no scheme detected in ${pager.nextLink} %s", *resp.${resultTypeName}.${pager.nextLink})\n`;
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
    imports.add('net/url');
    text += `\turlPath := "${op.requests![0].protocol.http!.path}"\n`;
    // replace path parameters
    for (const pp of values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'path'; })) {
      let paramValue = `url.PathEscape(${formatParamValue(pp, imports)})`;
      if (skipURLEncoding(pp)) {
        paramValue = formatParamValue(pp, imports);
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
  const inQueryParams = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'query'; });
  if (inQueryParams.any()) {
    // add query parameters
    text += '\tquery := u.Query()\n';
    for (const qp of values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'query'; })) {
      if (qp.required === true) {
        text += `\tquery.Set("${qp.language.go!.serializedName}", ${formatParamValue(qp, imports)})\n`;
      } else if (qp.schema.type === SchemaType.Constant) {
        // omit this query param. TODO once non-required constants are fixed
      } else if (qp.implementation === ImplementationLocation.Client) {
        // global optional param
        text += `\tif client.${qp.language.go!.name} != nil {\n`;
        text += `\t\tquery.Set("${qp.language.go!.serializedName}", ${formatParamValue(qp, imports)})\n`;
        text += `\t}\n`;
      } else {
        text += `\tif options != nil && options.${pascalCase(qp.language.go!.name)} != nil {\n`;
        text += `\t\tquery.Set("${qp.language.go!.serializedName}", ${formatParamValue(qp, imports)})\n`;
        text += `\t}\n`;
      }
    }
    text += '\tu.RawQuery = query.Encode()\n';
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
      text += `\tif options != nil && options.${pascalCase(header.language.go!.name)} != nil {\n`;
      text += `\t\treq.Header.Set("${header.language.go!.serializedName}", ${formatParamValue(header, imports)})\n`;
      text += `\t}\n`;
    }
  });
  const mediaType = getMediaType(op.requests![0].protocol);
  if (mediaType === 'JSON' || mediaType === 'XML') {
    const bodyParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http!.in === 'body'; }).first();
    // adding this variable to control whether a 'options.' needs to be added before optional body parameters
    let setOptionsPrefix = !bodyParam!.required;
    // default to the body param name
    let body = bodyParam!.language.go!.name;
    let skipOptions = false;
    if (bodyParam!.schema.type === SchemaType.Constant) {
      // if the value is constant, embed it directly
      body = formatConstantValue(<ConstantSchema>bodyParam!.schema);
      // directly assigned boolean values cannot be marshalled and are not set as enumerated types on
      // options structs, therefore would cause an issue when trying to access options.true or options.false
      // skipOptions skips appending an options prefix to these type of variables
      // NOTE: constants are commonly defined as enumerated types which is why an exception if being made for directly returned booleans
      if ((<ConstantSchema>bodyParam!.schema).valueType.type === SchemaType.Boolean) {
        skipOptions = true;
      }
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
      if (!bodyParam!.required) {
        body = `wrapper{${fieldName}: options.${pascalCase(bodyParam!.language.go!.name)}}`;
      } else {
        body = `wrapper{${fieldName}: &${bodyParam!.language.go!.name}}`;
      }
      // wrapper precludes the need for 'options.' prefix
      setOptionsPrefix = false;
    } else if (bodyParam!.schema.type === SchemaType.DateTime && (<DateTimeSchema>bodyParam!.schema).format === 'date-time-rfc1123') {
      // wrap the body in the custom RFC1123 type
      text += `\taux := ${bodyParam!.schema.language.go!.internalTimeType}`;
      if (!bodyParam!.required) {
        text += `(options.${pascalCase(body)})\n`;
      } else {
        text += `(${body})\n`;
      }
      body = 'aux';
      // aux precludes the need for 'options.' prefix
      setOptionsPrefix = false;
    } else if (isArrayOfRFC1123(bodyParam!.schema)) {
      const timeType = (<ArraySchema>bodyParam!.schema).elementType.language.go!.internalTimeType;
      text += `\taux := make([]${timeType}, len(${bodyParam!.language.go!.name}), len(${bodyParam!.language.go!.name}))\n`;
      text += `\tfor i := 0; i < len(${bodyParam!.language.go!.name}); i++ {\n`;
      text += `\t\taux[i] = ${timeType}(${bodyParam!.language.go!.name}[i])\n`;
      text += '\t}\n';
      body = 'aux';
      // aux precludes the need for 'options.' prefix
      setOptionsPrefix = false;
    } else if (isMapOfDateTime(bodyParam!.schema)) {
      const timeType = (<ArraySchema>bodyParam!.schema).elementType.language.go!.internalTimeType;
      text += `\taux := map[string]${timeType}{}\n`;
      text += `\tfor k, v := range ${bodyParam!.language.go!.name} {\n`;
      text += `\t\taux[k] = ${timeType}(v)\n`;
      text += '\t}\n';
      body = 'aux';
      // aux precludes the need for 'options.' prefix
      setOptionsPrefix = false;
    }
    if (setOptionsPrefix === true && !skipOptions) {
      body = `options.${pascalCase(body)}`;
      text += `\tif options != nil {\n`;
      text += `\t\treturn req, req.MarshalAs${mediaType}(${body})\n`;
      text += '\t}\n';
      text += '\treturn req, nil\n';
    } else {
      text += `\treturn req, req.MarshalAs${mediaType}(${body})\n`;
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

function isArrayOfRFC1123(schema: Schema): boolean {
  if (schema.type !== SchemaType.Array) {
    return false;
  }
  const arraySchema = <ArraySchema>schema;
  const arrayElem = <Schema>arraySchema.elementType;
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
    text += '\treturn nil, newError(resp)';
    text += '}\n\n';
    return text;
  }
  const firstResp = op.responses![0];
  text += `\tif !resp.HasStatusCode(${formatStatusCodes(firstResp.protocol.http?.statusCodes)}) {\n`;
  // if the response doesn't define a 'default' section return a generic error
  // TODO: can be multiple exceptions when x-ms-error-response is in use (rare)
  if (!op.exceptions || op.exceptions[0].language.go!.genericError) {
    imports.add('errors');
    text += `\t\treturn nil, errors.New(resp.Status)\n`;
  } else {
    const schemaError = (<SchemaResponse>op.exceptions![0]).schema;
    text += `\t\treturn nil, ${schemaError.language.go!.constructorName}(resp)\n`;
  }
  text += '\t}\n';

  if (!isSchemaResponse(firstResp)) {
    // no response body, return the *http.Response
    text += `\treturn resp.Response, nil\n`;
    text += '}\n\n';
    return text;
  } else if (firstResp.schema.type === SchemaType.DateTime) {
    // use the designated time type for unmarshalling
    text += `\tvar aux *${firstResp.schema.language.go!.internalTimeType}\n`;
    text += `\terr := resp.UnmarshalAs${getMediaType(firstResp.protocol)}(&aux)\n`;
    const resp = `${firstResp.schema.language.go!.responseType.name}{RawResponse: resp.Response, ${firstResp.schema.language.go!.responseType.value}: (*time.Time)(aux)}`;
    text += `\treturn &${resp}, err\n`;
    text += '}\n\n';
    return text;
  } else if (isArrayOfDateTime(firstResp.schema)) {
    // unmarshalling arrays of date/time is a little more involved
    text += `\tvar aux *[]${(<ArraySchema>firstResp.schema).elementType.language.go!.internalTimeType}\n`;
    text += `\tif err := resp.UnmarshalAs${getMediaType(firstResp.protocol)}(&aux); err != nil {\n`;
    text += '\t\treturn nil, err\n';
    text += '\t}\n';
    text += '\tcp := make([]time.Time, len(*aux), len(*aux))\n';
    text += '\tfor i := 0; i < len(*aux); i++ {\n';
    text += '\t\tcp[i] = time.Time((*aux)[i])\n';
    text += '\t}\n';
    const resp = `${firstResp.schema.language.go!.responseType.name}{RawResponse: resp.Response, ${firstResp.schema.language.go!.responseType.value}: &cp}`;
    text += `\treturn &${resp}, nil\n`;
    text += '}\n\n';
    return text;
  } else if (isMapOfDateTime(firstResp.schema)) {
    text += `\taux := map[string]${(<DictionarySchema>firstResp.schema).elementType.language.go!.internalTimeType}{}\n`;
    text += `\tif err := resp.UnmarshalAs${getMediaType(firstResp.protocol)}(&aux); err != nil {\n`;
    text += '\t\treturn nil, err\n';
    text += '\t}\n';
    text += `\tcp := map[string]time.Time{}\n`;
    text += `\tfor k, v := range aux {\n`;
    text += `\t\tcp[k] = time.Time(v)\n`;
    text += `\t}\n`;
    const resp = `${firstResp.schema.language.go!.responseType.name}{RawResponse: resp.Response, ${firstResp.schema.language.go!.responseType.value}: &cp}`;
    text += `\treturn &${resp}, nil\n`;
    text += '}\n\n';
    return text;
  }

  const schemaResponse = <SchemaResponse>firstResp;
  let respObj = `${schemaResponse.schema.language.go!.responseType.name}{RawResponse: resp.Response}`;
  text += `\tresult := ${respObj}\n`;
  // assign any header values
  for (const prop of values(<Array<Property>>schemaResponse.schema.language.go!.properties)) {
    if (prop.language.go!.fromHeader) {
      text += formatHeaderResponseValue(prop.language.go!.name, prop.language.go!.fromHeader, prop.schema, imports, 'result');
    }
  }
  const mediaType = getMediaType(firstResp.protocol);
  if (mediaType === 'none' || mediaType === 'binary') {
    // nothing to unmarshal
    text += '\treturn &result, nil\n';
    text += '}\n\n';
    return text;
  }
  let target = `result.${schemaResponse.schema.language.go!.responseType.value}`;
  // when unmarshalling a wrapped XML array, unmarshal into the response type, not the field
  if (mediaType === 'XML' && schemaResponse.schema.type === SchemaType.Array) {
    target = 'result';
  }
  text += `\treturn &result, resp.UnmarshalAs${mediaType}(&${target})\n`;
  text += '}\n\n';
  return text;
}

function isArrayOfDateTime(schema: Schema): boolean {
  if (schema.type !== SchemaType.Array) {
    return false;
  }
  const arraySchema = <ArraySchema>schema;
  const arrayElem = <Schema>arraySchema.elementType;
  return arrayElem.type === SchemaType.DateTime;
}

function isMapOfDateTime(schema: Schema): boolean {
  if (schema.type !== SchemaType.Dictionary) {
    return false;
  }
  const dictSchema = <DictionarySchema>schema;
  const dictElem = <Schema>dictSchema.elementType;
  return dictElem.type === SchemaType.DateTime;
}

function createInterfaceDefinition(group: OperationGroup, imports: ImportManager): string {
  let interfaceText = `// ${group.language.go!.clientName} contains the methods for the ${group.language.go!.name} group.\n`;
  interfaceText += `type ${group.language.go!.clientName} interface {\n`;
  for (const op of values(group.operations)) {
    if (isPageableOperation(op) && op.language.go!.paging.member === op.language.go!.name) {
      // don't generate a public API for the methods used to advance pages
      continue;
    }
    if (isLROOperation(op)) {
      op.language.go!.name = `Begin${op.language.go!.name}`;
    }
    for (const param of values(aggregateParameters(op))) {
      if (param.implementation !== ImplementationLocation.Method || param.required !== true) {
        continue;
      }
      imports.addImportForSchemaType(param.schema);
    }
    if (hasDescription(op.language.go!)) {
      interfaceText += `\t// ${op.language.go!.name} - ${op.language.go!.description} \n`;
    }
    const returns = generateReturnsInfo(op, false);
    interfaceText += `\t${op.language.go!.name}(${getAPIParametersSig(op, imports)}) (${returns.join(', ')})\n`;
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
    switch (rawCode) {
      case '200':
        asHTTPStatus.push('http.StatusOK');
        break;
      case '201':
        asHTTPStatus.push('http.StatusCreated');
        break;
      case '202':
        asHTTPStatus.push('http.StatusAccepted');
        break;
      case '204':
        asHTTPStatus.push('http.StatusNoContent');
        break;
      case '300':
        asHTTPStatus.push('http.StatusMultipleChoices');
        break;
      case '301':
        asHTTPStatus.push('http.StatusMovedPermanently');
        break;
      case '302':
        asHTTPStatus.push('http.StatusFound');
        break;
      case '400':
        asHTTPStatus.push('http.StatusBadRequest');
        break;
      default:
        throw console.error(`unhandled status code ${rawCode}`);
    }
  }
  return asHTTPStatus.join(', ');
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
  if (!isPageableOperation(op)) {
    imports.add('context');
    params.push('ctx context.Context');
  }
  for (const methodParam of values(methodParams)) {
    params.push(`${methodParam.language.go!.name} ${formatParameterTypeName(methodParam)}`);
  }
  return params.join(', ');
}

// returns the parameters for the internal request creator method.
// e.g. "i int, s string"
function getCreateRequestParametersSig(op: Operation): string {
  const methodParams = getMethodParameters(op);
  const params = new Array<string>();
  for (const methodParam of values(methodParams)) {
    params.push(`${methodParam.language.go!.name} ${formatParameterTypeName(methodParam)}`);
  }
  return params.join(', ');
}

// returns the complete collection of method parameters
function getMethodParameters(op: Operation): Parameter[] {
  const params = new Array<Parameter>();
  for (const param of values(aggregateParameters(op))) {
    if (param.implementation === ImplementationLocation.Client) {
      // client params are passed via the receiver
      continue;
    } else if (param.schema.type === SchemaType.Constant) {
      // don't generate a parameter for a constant
      continue;
    } else if (param.implementation === ImplementationLocation.Method && param.required !== true) {
      // omit method-optional params as they're grouped in the optional params type
      continue;
    }
    params.push(param);
  }
  // move global optional params to the end of the slice
  params.sort(sortParametersByRequired);
  // if there's a method-optional params struct add it last
  if (op.requests![0].language.go!.optionalParam) {
    params.push(op.requests![0].language.go!.optionalParam);
  }

  return params;
}

// returns the return signature where each entry is the type name
// e.g. [ '*string', 'error' ]
function generateReturnsInfo(op: Operation, forHandler: boolean): string[] {
  if (!op.responses) {
    return ['*http.Response', 'error'];
  }
  // TODO check this implementation, if any additional return information needs to be included for multiple responses
  const firstResp = op.responses![0];
  let returnType = '*http.Response';
  // must check pageable first as all pageable operations are also schema responses
  if (!forHandler && isPageableOperation(op)) {
    returnType = op.language.go!.pageableType.name;
  } else if (!forHandler && isLROOperation(op)) {
    returnType = pascalCase(op.language.go!.pollerType.name);
  } else if (isSchemaResponse(firstResp)) {
    returnType = '*' + firstResp.schema.language.go!.responseType.name;
  }
  return [returnType, 'error'];
}
