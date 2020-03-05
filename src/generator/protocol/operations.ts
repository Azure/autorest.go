/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment, KnownMediaType, pascalCase } from '@azure-tools/codegen'
import { ArraySchema, CodeModel, ConstantSchema, DateTimeSchema, ImplementationLocation, Language, NumberSchema, Operation, Parameter, Protocols, Response, Schema, SchemaResponse, SchemaType, SerializationStyle } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { aggregateParameters, ContentPreamble, generateParamsSig, generateParameterInfo, genereateReturnsInfo, ImportManager, isArraySchema, LanguageHeader, MethodSig, ParamInfo, paramInfo, SortAscending } from '../common/helpers';
import { OperationNaming } from '../../namer/namer';

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
    imports.add('net/url');
    imports.add('path');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');

    const clientName = group.language.go!.clientName;
    let opText = '';
    group.operations.sort((a: Operation, b: Operation) => { return SortAscending(a.language.go!.name, b.language.go!.name) });
    for (const op of values(group.operations)) {
      // protocol creation can add imports to the list so
      // it must be done before the imports are written out
      op.language.go!.protocolSigs = new protocolSigs();
      opText += createProtocolRequest(clientName, op, imports);
      opText += createProtocolResponse(clientName, op, imports);
    }
    // stitch it all together
    let text = await ContentPreamble(session);
    text += imports.text();
    text += `// ${clientName} contains the methods for the ${group.language.go!.name} group.\n`;
    text += `type ${clientName} struct{}\n\n`;
    text += opText;

    operations.push(new OperationGroupContent(group.language.go!.name, text));
  }
  return operations;
}

// contains method signature information for request and response methods
export interface ProtocolSig extends Language {
  protocolSigs: ProtocolSigs;
}

interface ProtocolSigs {
  requestMethod: MethodSig;
  responseMethod: MethodSig;
}

class protocolSigs implements ProtocolSigs {
  requestMethod: MethodSig;
  responseMethod: MethodSig;
  constructor() {
    this.requestMethod = new methodSig();
    this.responseMethod = new methodSig();
  }
}

class methodSig implements MethodSig {
  params: ParamInfo[];
  returns: string[];
  constructor() {
    this.params = new Array<ParamInfo>();
    this.returns = new Array<string>();
  }
}

export interface HeaderResponse {
  body: string;
  respObj: string;
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
  if (param.required !== true) {
    if (param.implementation === ImplementationLocation.Method) {
      // optional params at the method level will be in an options struct
      paramName = `*options.${pascalCase(paramName)}`;
    } else {
      // optional globals are passed as just another parameter
      paramName = `*${paramName}`;
    }
  }
  switch (param.schema.type) {
    case SchemaType.Array:
      const arraySchema = <ArraySchema>param.schema;
      switch (arraySchema.elementType.type) {
        case SchemaType.Choice:
        case SchemaType.SealedChoice:
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
function formatHeaderResponseValue(header: LanguageHeader, imports: ImportManager, respObj: string): HeaderResponse {
  if (respObj[respObj.length - 1] == '}') {
    respObj = respObj.substring(0, respObj.length - 1);
  }
  let headerText = <HeaderResponse>{};
  let text = ``;
  switch (header.schema.type) {
    case SchemaType.Boolean:
      imports.add('strconv');
      text = `\tval, err := strconv.ParseBool(resp.Header.Get("${header.header}"))\n`;
      text += `\tif err != nil {\n`;
      text += `\t\treturn nil, err\n`;
      text += `\t}\n`;
      headerText.body = text;
      headerText.respObj = respObj + `, ${header.name}: &val}`;
      return headerText;
    case SchemaType.ByteArray:
      // ByteArray is a base-64 encoded value in string format
      imports.add('encoding/base64');
      headerText.body = `\tval := []byte(resp.Header.Get("${header.header}"))\n`;
      headerText.respObj = respObj + `, ${header.name}: &val}`;
      return headerText;
    case SchemaType.Choice:
    case SchemaType.SealedChoice:
      headerText.body = `\tval := ${header.schema.language.go!.name}(resp.Header.Get("${header.header}"))\n`;
      headerText.respObj = respObj + `, ${header.name}: &val}`;
      return headerText;
    case SchemaType.Constant:
    case SchemaType.String:
      headerText.body = `\tval := resp.Header.Get("${header.header}")\n`;
      headerText.respObj = respObj + `, ${header.name}: &val}`;
      return headerText;
    case SchemaType.Date:
      imports.add('time');
      text = `\tval, err := time.Parse("${dateFormat}", resp.Header.Get("${header.header}"))\n`;
      text += `\tif err != nil {\n`;
      text += `\t\treturn nil, err\n`;
      text += `\t}\n`;
      headerText.body = text;
      headerText.respObj = respObj + `, ${header.name}: &val}`;
      return headerText;
    case SchemaType.DateTime:
      imports.add('time');
      let format = datetimeRFC3339Format;
      const dateTime = <DateTimeSchema>header.schema;
      if (dateTime.format === 'date-time-rfc1123') {
        format = datetimeRFC1123Format;
      }
      text = `\tval, err := time.Parse(${format}, resp.Header.Get("${header.header}"))\n`;
      text += `\tif err != nil {\n`;
      text += `\t\treturn nil, err\n`;
      text += `\t}\n`;
      headerText.body = text;
      headerText.respObj = respObj + `, ${header.name}: &val}`;
      return headerText;
    case SchemaType.Duration:
      imports.add('time');
      text = `\tval, err := time.ParseDuration(resp.Header.Get("${header.header}"))\n`;
      text += `\tif err != nil {\n`;
      text += `\t\treturn nil, err\n`;
      text += `\t}\n`;
      headerText.body = text;
      headerText.respObj = respObj + `, ${header.name}: &val}`;
      return headerText;
    case SchemaType.Integer:
      imports.add('strconv');
      const intNum = <NumberSchema>header.schema;
      if (intNum.precision === 32) {
        headerText.body = `\tval32, err := strconv.ParseInt(resp.Header.Get("${header.header}"), 10, 32)\n`;
        headerText.body += `\tval := int32(val32)\n`;
      } else {
        headerText.body = `\tval, err := strconv.ParseInt(resp.Header.Get("${header.header}"), 10, 64)\n`;
      }
      headerText.body += `\tif err != nil {\n`;
      headerText.body += `\t\treturn nil, err\n`;
      headerText.body += `\t}\n`;
      headerText.respObj = respObj + `, ${header.name}: &val}`;
      return headerText;
    case SchemaType.Number:
      imports.add('strconv');
      const floatNum = <NumberSchema>header.schema;
      if (floatNum.precision === 32) {
        headerText.body = `\tval32, err := strconv.ParseFloat(resp.Header.Get("${header.header}"), 32)\n`;
        headerText.body += `\tval := float32(val32)\n`;
      } else {
        headerText.body = `\tval, err := strconv.ParseFloat(resp.Header.Get("${header.header}"), 64)\n`;
      }
      headerText.body += `\tif err != nil {\n`;
      headerText.body += `\t\treturn nil, err\n`;
      headerText.body += `\t}\n`;
      headerText.respObj = respObj + `, ${header.name}: &val}`;
      return headerText;
    default:
      if (respObj[respObj.length - 1] == '}') {
        headerText.respObj = respObj + "}";
      }
      return headerText;
  }
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
  // stick the method signature info into the code model so other generators can access it later
  const sig = <ProtocolSig>op.language.go!;
  sig.protocolSigs.requestMethod.params = [new paramInfo('u', 'url.URL', false, true)].concat(generateParameterInfo(op));
  sig.protocolSigs.requestMethod.returns = ['*azcore.Request', 'error'];
  let text = `${comment(name, '// ')} creates the ${info.name} request.\n`;
  text += `func (${client}) ${name}(${generateParamsSig(sig.protocolSigs.requestMethod.params, true)}) (${sig.protocolSigs.requestMethod.returns.join(', ')}) {\n`;
  text += `\turlPath := "${op.requests![0].protocol.http!.path}"\n`;
  if (values(aggregateParameters(op)).any((each: Parameter) => { return each.protocol.http!.in === 'path' })) {
    // replace path parameters
    imports.add('strings');
    imports.add('net/url');
    for (const pp of values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http!.in === 'path' })) {
      text += `\turlPath = strings.ReplaceAll(urlPath, "{${pp.language.go!.serializedName}}", url.PathEscape(${formatParamValue(pp, imports)}))\n`;
    }
  }
  text += `\tu.Path = path.Join(u.Path, urlPath)\n`;
  if (values(aggregateParameters(op)).any((each: Parameter) => { return each.protocol.http!.in === 'query' })) {
    // add query parameters
    text += '\tquery := u.Query()\n';
    for (const qp of values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http!.in === 'query'; })) {
      if (qp.required === true) {
        text += `\tquery.Set("${qp.language.go!.name}", ${formatParamValue(qp, imports)})\n`;
      } else if (qp.implementation === ImplementationLocation.Client) {
        // global optional param
        text += `\tif ${qp.language.go!.name} != nil {\n`;
        text += `\t\tquery.Set("${qp.language.go!.name}", ${formatParamValue(qp, imports)})\n`;
        text += `\t}\n`;
      } else {
        text += `\tif options != nil && options.${pascalCase(qp.language.go!.name)} != nil {\n`;
        text += `\t\tquery.Set("${qp.language.go!.name}", ${formatParamValue(qp, imports)})\n`;
        text += `\t}\n`;
      }
    }
    text += '\tu.RawQuery = query.Encode()\n';
  }
  text += `\treq := azcore.NewRequest(http.Method${pascalCase(op.requests![0].protocol.http!.method)}, u)\n`;
  if (hasBinaryResponse(op.responses!)) {
    // skip auto-body downloading for binary stream responses
    text += '\treq.SkipBodyDownload()\n';
  }
  // add specific request headers
  const headerParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http!.in === 'header'; });
  headerParam.forEach(header => {
    text += `\treq.Header.Set("${header.language.go!.serializedName}", ${formatParamValue(header, imports)})\n`;
  });
  const mediaType = getMediaType(op.requests![0].protocol);
  if (mediaType === 'JSON' || mediaType === 'XML') {
    const bodyParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http!.in === 'body'; }).first();
    // default to the body param name
    let body = bodyParam!.language.go!.name;
    if (bodyParam!.schema.type === SchemaType.Constant) {
      // if the value is constant, embed it directly
      body = formatConstantValue(<ConstantSchema>bodyParam!.schema);
    } else if (mediaType === 'XML' && bodyParam!.schema.serialization?.xml) {
      // for XML payloads, create a wrapper type in the following cases
      // 1. the type's name doesn't match the XML name (e.g. "Slideshow" vs. "slideshow")
      // 2. the payload is a wrapped array
      if (bodyParam!.schema.language.go!.name !== bodyParam!.schema.serialization.xml.name || bodyParam!.schema.type === SchemaType.Array) {
        imports.add('encoding/xml');
        text += '\ttype wrapper struct {\n';
        text += `\t\tXMLName xml.Name \`xml:"${bodyParam!.schema.serialization.xml.name}"\`\n`;
        let fieldName = bodyParam!.schema.language.go!.name;
        if (isArraySchema(bodyParam!.schema)) {
          fieldName = pascalCase(bodyParam!.language.go!.name);
          let tag = bodyParam!.schema.elementType.language.go!.name;
          if (bodyParam!.schema.elementType.serialization?.xml?.name) {
            tag = bodyParam!.schema.elementType.serialization.xml.name;
          }
          text += `\t\t${fieldName} *${bodyParam!.schema.language.go!.name} \`xml:"${tag}"\`\n`;
        } else {
          // embed as anonymous field
          text += `\t\t*${bodyParam!.schema.language.go!.name}\n`;
        }
        text += '\t}\n';
        body = `wrapper{${fieldName}: &${bodyParam!.language.go!.name}}`;
      }
    } else if (bodyParam!.schema.type === SchemaType.DateTime && (<DateTimeSchema>bodyParam!.schema).format === 'date-time-rfc1123') {
      // wrap the body in the custom RFC1123 type
      text += `\taux := ${bodyParam!.schema.language.go!.internalTimeType}(${body})\n`;
      body = 'aux';
    }
    text += `\terr := req.MarshalAs${mediaType}(${body})\n`;
    text += `\tif err != nil {\n`;
    text += `\t\treturn nil, err\n`;
    text += `\t}\n`;
  }
  text += `\treturn req, nil\n`;
  text += '}\n\n';
  return text;
}

function createProtocolResponse(client: string, op: Operation, imports: ImportManager): string {
  const info = <OperationNaming>op.language.go!;
  const name = info.protocolNaming.responseMethod;
  // stick the method signature info into the code model so other generators can access it later
  const sig = <ProtocolSig>op.language.go!;
  sig.protocolSigs.responseMethod.params = [new paramInfo('resp', '*azcore.Response', false, true)];
  sig.protocolSigs.responseMethod.returns = genereateReturnsInfo(op);

  const resp = op.responses![0];
  let text = `${comment(name, '// ')} handles the ${info.name} response.\n`;
  text += `func (${client}) ${name}(${generateParamsSig(sig.protocolSigs.responseMethod.params, true)}) (${sig.protocolSigs.responseMethod.returns.join(', ')}) {\n`;
  text += `\tif !resp.HasStatusCode(${formatStatusCodes(resp.protocol.http?.statusCodes)}) {\n`;
  text += `\t\treturn nil, newError(resp)\n`;
  text += '\t}\n';

  let respObj = `${resp.language.go!.name}{RawResponse: resp.Response}`;
  let headResp = <HeaderResponse>{};
  // check if the response is expecting information from headers
  if (resp.protocol.http!.headers) {
    for (const header of values(resp.protocol.http!.headers)) {
      const head = <LanguageHeader>header;
      headResp = formatHeaderResponseValue(head, imports, respObj);
      // reassign respObj to include the value returned from the headers
      respObj = headResp.respObj;
      // add the code necessary to process data returned in a header
      if (headResp.body) {
        text += headResp.body;
      }
    }
  }
  const mediaType = getMediaType(resp.protocol);
  if (mediaType === 'none') {
    // no response body so nothing to unmarshal
    text += `\treturn &${respObj}, nil\n`;
  } else if ((<SchemaResponse>resp).schema.type === SchemaType.DateTime) {
    // use the designated time type for unmarshalling
    text += `\tvar aux *${(<SchemaResponse>resp).schema.language.go!.internalTimeType}\n`;
    text += `\terr := resp.UnmarshalAs${mediaType}(&aux)\n`;
    text += `\tresult := ${respObj}\n`;
    text += `\tresult.${(<SchemaResponse>resp).schema.language.go!.responseValue} = aux.ToTime()\n`;
    text += `\treturn &result, err\n`;
  } else {
    text += `\tresult := ${respObj}\n`;
    let target = `result.${(<SchemaResponse>resp).schema.language.go!.responseValue}`;
    // when unmarshalling a wrapped XML array, unmarshal into the response type, not the field
    if (mediaType === 'XML' && (<SchemaResponse>resp).schema.type === SchemaType.Array) {
      target = 'result';
    }
    text += `\treturn &result, resp.UnmarshalAs${mediaType}(&${target})\n`;
  }
  text += '}\n\n';
  return text;
}

// returns the media type used by the protocol
function getMediaType(protocol: Protocols): 'JSON' | 'XML' | 'none' {
  // TODO: binary, forms etc
  switch (protocol.http!.knownMediaType) {
    case KnownMediaType.Json:
      return 'JSON';
    case KnownMediaType.Xml:
      return 'XML';
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
