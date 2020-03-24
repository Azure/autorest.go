/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment, KnownMediaType, pascalCase, camelCase } from '@azure-tools/codegen'
import { ArraySchema, CodeModel, ConstantSchema, DateTimeSchema, ImplementationLocation, Language, NumberSchema, Operation, OperationGroup, Parameter, Property, Protocols, Response, Schema, SchemaResponse, SchemaType, SerializationStyle } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { aggregateParameters, ContentPreamble, formatParamInfoTypeName, generateParamsSig, generateParameterInfo, genereateReturnsInfo, HasDescription, ImportManager, isArraySchema, isPageableOperation, MethodSig, ParamInfo, paramInfo, skipURLEncoding, SortAscending, isSchemaResponse, PagerInfo } from './helpers';
import { OperationNaming } from '../namer/namer';

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
    group.operations.sort((a: Operation, b: Operation) => { return SortAscending(a.language.go!.name, b.language.go!.name) });
    for (const op of values(group.operations)) {
      // protocol creation can add imports to the list so
      // it must be done before the imports are written out
      op.language.go!.protocolSigs = new protocolSigs();
      // TODO: generateOperation depends on some work that happens in
      // protocol request/response creation, fix this
      const reqText = createProtocolRequest(clientName, op, imports);
      const respText = createProtocolResponse(clientName, op, imports);
      opText += generateOperation(clientName, op, imports);
      opText += reqText;
      opText += respText;
    }
    const interfaceText = createInterfaceDefinition(group, imports);
    // stitch it all together
    let text = await ContentPreamble(session);
    text += imports.text();
    text += interfaceText;
    text += `// ${clientName} implements the ${group.language.go!.clientName} interface.\n`;
    text += `type ${clientName} struct{\n`;
    text += '\t*Client\n';
    if (group.language.go!.globals) {
      const globals = <Array<ParamInfo>>group.language.go!.globals;
      for (const global of values(globals)) {
        text += `\t${global.name} ${formatParamInfoTypeName(global)}\n`;
      }
    }
    text += '}\n\n';
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
  let text = '';
  let needsErrorCheck = true;
  const name = camelCase(propName);
  switch (schema.type) {
    case SchemaType.Boolean:
      imports.add('strconv');
      text = `\t${name}, err := strconv.ParseBool(resp.Header.Get("${header}"))\n`;
      break;
    case SchemaType.ByteArray:
      // ByteArray is a base-64 encoded value in string format
      imports.add('encoding/base64');
      text = `\t${name}, err := base64.StdEncoding.DecodeString(resp.Header.Get("${header}"))\n`;
      break;
    case SchemaType.Choice:
    case SchemaType.SealedChoice:
      text = `\t${name} := ${schema.language.go!.name}(resp.Header.Get("${header}"))\n`;
      needsErrorCheck = false;
      break;
    case SchemaType.Constant:
    case SchemaType.String:
      text = `\t${name} := resp.Header.Get("${header}")\n`;
      needsErrorCheck = false;
      break;
    case SchemaType.Date:
      imports.add('time');
      text = `\t${name}, err := time.Parse("${dateFormat}", resp.Header.Get("${header}"))\n`;
      break;
    case SchemaType.DateTime:
      imports.add('time');
      let format = datetimeRFC3339Format;
      const dateTime = <DateTimeSchema>schema;
      if (dateTime.format === 'date-time-rfc1123') {
        format = datetimeRFC1123Format;
      }
      text = `\t${name}, err := time.Parse(${format}, resp.Header.Get("${header}"))\n`;
      break;
    case SchemaType.Duration:
      imports.add('time');
      text = `\t${name}, err := time.ParseDuration(resp.Header.Get("${header}"))\n`;
      break;
    case SchemaType.Integer:
      imports.add('strconv');
      const intNum = <NumberSchema>schema;
      if (intNum.precision === 32) {
        text = `\t${name}32, err := strconv.ParseInt(resp.Header.Get("${header}"), 10, 32)\n`;
        text += `\t${name} := int32(${name}32)\n`;
      } else {
        text = `\t${name}, err := strconv.ParseInt(resp.Header.Get("${header}"), 10, 64)\n`;
      }
      break;
    case SchemaType.Number:
      imports.add('strconv');
      const floatNum = <NumberSchema>schema;
      if (floatNum.precision === 32) {
        text = `\t${name}32, err := strconv.ParseFloat(resp.Header.Get("${header}"), 32)\n`;
        text += `\t${name} := float32(${name}32)\n`;
      } else {
        text = `\t${name}, err := strconv.ParseFloat(resp.Header.Get("${header}"), 64)\n`;
      }
      break;
    default:
      throw console.error(`unsupported header type ${schema.type}`);
  }
  if (needsErrorCheck) {
    text += `\tif err != nil {\n`;
    text += `\t\treturn nil, err\n`;
    text += `\t}\n`;
  }
  text += `\t${respObj}.${propName} = &${name}\n`;
  return text;
}

function getParamInfo(op: Operation, imports: ImportManager): paramInfo[] {
  let params = generateParameterInfo(op);
  if (!isPageableOperation(op)) {
    imports.add('context');
    params = [new paramInfo('ctx', 'context.Context', false, true)].concat(params);
  }
  return params;
}

function generateOperation(clientName: string, op: Operation, imports: ImportManager): string {
  if (isPageableOperation(op) && op.language.go!.paging.member === op.language.go!.name) {
    // don't generate a public API for the methods used to advance pages
    return '';
  }
  const info = <OperationNaming>op.language.go!;
  const params = getParamInfo(op, imports);
  const returns = genereateReturnsInfo(op, false);
  const protocol = <ProtocolSig>op.language.go!;
  let text = '';
  if (HasDescription(op.language.go!)) {
    text += `// ${op.language.go!.name} - ${op.language.go!.description} \n`;
  }
  text += `func (client *${clientName}) ${op.language.go!.name}(${generateParamsSig(params, false)}) (${returns.join(', ')}) {\n`;
  // slice off the first param returned from extractParamNames as we know it's the URL (cheating a bit...)
  const protocolReqParams = extractParamNames(protocol.protocolSigs.requestMethod.params);
  text += `\treq, err := client.${info.protocolNaming.requestMethod}(${protocolReqParams.join(', ')})\n`;
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
      protocolReqParams.push(`*resp.${pager.schema.language.go!.name}.${pager.nextLink}`);
      text += `\t\tadvancer: func(resp *${pager.schema.language.go!.responseType.name}) (*azcore.Request, error) {\n`;
      text += `\t\t\treturn client.${camelCase(op.language.go!.paging.member)}CreateRequest(${protocolReqParams.join(', ')})\n`;
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
  // stick the method signature info into the code model so other generators can access it later
  const sig = <ProtocolSig>op.language.go!;
  sig.protocolSigs.requestMethod.params = generateParameterInfo(op);
  sig.protocolSigs.requestMethod.returns = ['*azcore.Request', 'error'];
  let text = `${comment(name, '// ')} creates the ${info.name} request.\n`;
  text += `func (client *${client}) ${name}(${generateParamsSig(sig.protocolSigs.requestMethod.params, true)}) (${sig.protocolSigs.requestMethod.returns.join(', ')}) {\n`;
  text += `\turlPath := "${op.requests![0].protocol.http!.path}"\n`;
  const inPathParams = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'path'; });
  if (inPathParams.any()) {
    imports.add('strings');
    imports.add('net/url');
    // replace path parameters
    for (const pp of values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'path'; })) {
      let paramValue = `url.PathEscape(${formatParamValue(pp, imports)})`;
      if (skipURLEncoding(pp)) {
        paramValue = formatParamValue(pp, imports);
      }
      text += `\turlPath = strings.ReplaceAll(urlPath, "{${pp.language.go!.serializedName}}", ${paramValue})\n`;
    }
  }

  text += `\tu, err := client.u.Parse(urlPath)\n`;
  text += '\tif err != nil {\n';
  text += '\t\treturn nil, err\n';
  text += '\t}\n';
  const inQueryParams = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'query'; });
  if (inQueryParams.any()) {
    // add query parameters
    text += '\tquery := u.Query()\n';
    for (const qp of values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'query'; })) {
      if (qp.required === true) {
        text += `\tquery.Set("${qp.language.go!.serializedName}", ${formatParamValue(qp, imports)})\n`;
      } else if (qp.implementation === ImplementationLocation.Client) {
        // global optional param
        text += `\tif ${qp.language.go!.name} != nil {\n`;
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
    }
    if (setOptionsPrefix === true) {
      body = `options.${pascalCase(body)}`;
      text += `\tif options != nil {\n`;
      text += `\t\terr = req.MarshalAs${mediaType}(${body})\n`;
      text += `\t\tif err != nil {\n`;
      text += `\t\t\treturn nil, err\n`;
      text += `\t\t}\n`;
      text += '\t}\n';
    } else {
      text += `\terr = req.MarshalAs${mediaType}(${body})\n`;
      text += `\tif err != nil {\n`;
      text += `\t\treturn nil, err\n`;
      text += `\t}\n`;
    }
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
  sig.protocolSigs.responseMethod.returns = genereateReturnsInfo(op, true);

  const firstResp = op.responses![0];
  let text = `${comment(name, '// ')} handles the ${info.name} response.\n`;
  text += `func (client *${client}) ${name}(${generateParamsSig(sig.protocolSigs.responseMethod.params, true)}) (${sig.protocolSigs.responseMethod.returns.join(', ')}) {\n`;
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
  if (mediaType === 'none') {
    // nothing to unmarshal
    text += '\treturn &result, nil\n';
    text += '}\n\n';
    return text;
  }
  if (schemaResponse.schema.type === SchemaType.DateTime) {
    // use the designated time type for unmarshalling
    text += `\tvar aux *${schemaResponse.schema.language.go!.internalTimeType}\n`;
    text += `\terr := resp.UnmarshalAs${mediaType}(&aux)\n`;
    text += `\tresult.${schemaResponse.schema.language.go!.responseType.value} = (*time.Time)(aux)\n`;
    text += `\treturn &result, err\n`;
  } else {
    let target = `result.${schemaResponse.schema.language.go!.responseType.value}`;
    // when unmarshalling a wrapped XML array, unmarshal into the response type, not the field
    if (mediaType === 'XML' && schemaResponse.schema.type === SchemaType.Array) {
      target = 'result';
    }
    text += `\treturn &result, resp.UnmarshalAs${mediaType}(&${target})\n`;
  }
  text += '}\n\n';
  return text;
}

function createInterfaceDefinition(group: OperationGroup, imports: ImportManager): string {
  let interfaceText = `// ${group.language.go!.clientName} contains the methods for the ${group.language.go!.name} group.\n`;
  interfaceText += `type ${group.language.go!.clientName} interface {\n`;
  for (const op of values(group.operations)) {
    if (isPageableOperation(op) && op.language.go!.paging.member === op.language.go!.name) {
      // don't generate a public API for the methods used to advance pages
      continue;
    }
    for (const param of values(aggregateParameters(op))) {
      if (param.implementation !== ImplementationLocation.Method || param.required !== true) {
        continue;
      }
      imports.addImportForSchemaType(param.schema);
    }
    if (HasDescription(op.language.go!)) {
      interfaceText += `\t// ${op.language.go!.name} - ${op.language.go!.description} \n`;
    }
    const params = getParamInfo(op, imports);
    const returns = genereateReturnsInfo(op, false);
    interfaceText += `\t${op.language.go!.name}(${generateParamsSig(params, false)}) (${returns.join(', ')})\n`;
  }
  interfaceText += '}\n\n';
  return interfaceText;
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

// returns an array of just the parameter names
// e.g. [ 'i', 's', 'b' ]
function extractParamNames(paramInfo: ParamInfo[]): string[] {
  let paramNames = new Array<string>();
  for (const param of values(paramInfo)) {
    let name = param.name;
    if (param.global) {
      name = `client.${name}`;
    }
    paramNames.push(name);
  }
  return paramNames;
}
