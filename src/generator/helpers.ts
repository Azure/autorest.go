/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { values } from '@azure-tools/linq';
import { capitalize, comment, uncapitalize } from '@azure-tools/codegen';
import { aggregateParameters, isSchemaResponse, isMultiRespOperation } from '../common/helpers';
import { ArraySchema, CodeModel, DictionarySchema, Language, Parameter, Schema, SchemaType, ObjectSchema, Operation, Property, GroupProperty, ImplementationLocation, SerializationStyle, ByteArraySchema, ConstantSchema, NumberSchema, DateTimeSchema } from '@autorest/codemodel';
import { ImportManager } from './imports';

export const dateFormat = '2006-01-02';
export const datetimeRFC3339Format = 'time.RFC3339Nano';
export const datetimeRFC1123Format = 'time.RFC1123';

// returns the common source-file preamble (license comment, package name etc)
export async function contentPreamble(session: Session<CodeModel>): Promise<string> {
  const headerText = comment(await session.getValue('header-text', 'MISSING LICENSE HEADER'), '// ');
  let text = `//go:build go1.18\n`;
  text += `// +build go1.18\n\n${headerText}\n// DO NOT EDIT.\n\n`;
  text += `package ${session.model.language.go!.packageName}\n\n`;
  return text;
}

// returns true if the language contains a description
export function hasDescription(lang: Language): boolean {
  return (lang.description !== undefined && lang.description.length > 0 && !lang.description.startsWith('MISSING'));
}

// used to sort strings in ascending order
export function sortAscending(a: string, b: string): number {
  return a < b ? -1 : a > b ? 1 : 0;
}

// returns true if the element type for a parameter should be passed by value
export function elementByValueForParam(param: Parameter): boolean {
  // passing nil for array elements in headers, paths, and query params
  // isn't very useful as we'd just skip nil entries.  so disable it.
  if (param.schema.type === SchemaType.Array) {
    return param.protocol.http!.in === 'header' || param.protocol.http!.in === 'path' || param.protocol.http!.in === 'query';
  } else if (param.schema.type === SchemaType.Dictionary) {
    return param.extensions?.['x-ms-header-collection-prefix'];
  }
  return false;
}

// returns the type name with possible * prefix
export function formatParameterTypeName(param: Parameter): string {
  const typeName = substituteDiscriminator(param.schema, elementByValueForParam(param));
  // client params with default values are treated as optional
  if (param.required && !(param.implementation === ImplementationLocation.Client && param.clientDefaultValue)) {
    return typeName;
  }
  return `*${typeName}`;
}

// returns true if the parameter should not be URL encoded
export function skipURLEncoding(param: Parameter): boolean {
  if (param.extensions) {
    return param.extensions['x-ms-skip-url-encoding'] === true;
  }
  return false;
}

// sorts parameters by their required state, ordering required before optional
export function sortParametersByRequired(a: Parameter, b: Parameter): number {
  if (a.required === b.required) {
    return 0;
  }
  if (a.required && !b.required) {
    return -1;
  }
  return 1;
}

// if a field is a discriminator use the interface type instead.
// elemByValue is to support corner-cases where we explicitly want
// the element type to be passed by value.
export function substituteDiscriminator(schema: Schema, elemByVal: boolean): string {
  switch (schema.type) {
    case SchemaType.Array:
      const arraySchema = <ArraySchema>schema;
      const arrayElem = <Schema>arraySchema.elementType;
      if (<boolean>arraySchema.language.go!.elementIsPtr && !elemByVal) {
        return `[]*${substituteDiscriminator(arrayElem, elemByVal)}`;
      }
      return `[]${substituteDiscriminator(arrayElem, elemByVal)}`;
    case SchemaType.Dictionary:
      const dictSchema = <DictionarySchema>schema;
      const dictElem = <Schema>dictSchema.elementType;
      if (<boolean>dictSchema.language.go!.elementIsPtr && !elemByVal) {
        return `map[string]*${substituteDiscriminator(dictElem, elemByVal)}`;
      }
      return `map[string]${substituteDiscriminator(dictElem, elemByVal)}`;
    case SchemaType.Object:
      if (schema.language.go!.discriminatorInterface) {
        return schema.language.go!.discriminatorInterface;
      }
      return schema.language.go!.name;
    default:
      return schema.language.go!.name;
  }
}

// if an LRO returns a discriminated type, unmarshall the response into the response envelope, else the property field
export function discriminatorFinalResponse(respEnv: ObjectSchema): string {
  const resultProp = <Property>respEnv.language.go!.resultProp;
  if (resultProp.schema.language.go!.discriminatorInterface) {
    return '';
  }
  return '.' + resultProp.language.go!.name;
}

// returns the parameters for the internal request creator method.
// e.g. "i int, s string"
export function getCreateRequestParametersSig(op: Operation): string {
  const methodParams = getMethodParameters(op);
  const params = new Array<string>();
  params.push('ctx context.Context');
  for (const methodParam of values(methodParams)) {
    params.push(`${uncapitalize(methodParam.language.go!.name)} ${formatParameterTypeName(methodParam)}`);
  }
  return params.join(', ');
}

// returns the parameter names for an operation (excludes the param types).
// e.g. "i, s"
export function getCreateRequestParameters(op: Operation): string {
  // split param list into individual params
  const reqParams = getCreateRequestParametersSig(op).split(',');
  // keep the parameter names from the name/type tuples
  for (let i = 0; i < reqParams.length; ++i) {
    reqParams[i] = reqParams[i].trim().split(' ')[0];
  }
  return reqParams.join(', ');
}

// returns the complete collection of method parameters
export function getMethodParameters(op: Operation): Parameter[] {
  const params = new Array<Parameter>();
  const paramGroups = new Array<GroupProperty>();
  for (const param of values(aggregateParameters(op))) {
    if (param.implementation === ImplementationLocation.Client) {
      // client params are passed via the receiver
      continue;
    } else if (param.language.go!.paramGroup) {
      // param groups will be added after individual params
      if (!paramGroups.includes(param.language.go!.paramGroup)) {
        paramGroups.push(param.language.go!.paramGroup);
      }
      continue;
    } else if (param.schema.type === SchemaType.Constant) {
      // don't generate a parameter for a constant
      // NOTE: this check must come last as non-required optional constants
      // in header/query params get dumped into the optional params group
      continue;
    }
    params.push(param);
  }
  // move global optional params to the end of the slice
  params.sort(sortParametersByRequired);
  // add any parameter groups.  optional groups go last
  paramGroups.sort((a: GroupProperty, b: GroupProperty) => {
    if (a.required === b.required) {
      return 0;
    }
    if (a.required && !b.required) {
      return -1;
    }
    return 1;
  })
  // add the optional param group last if it's not already in the list.
  // all operations should have an optional params type.  the only exception
  // is the next link operation for pageable operations.
  if (op.language.go!.optionalParamGroup && !values(paramGroups).any(gp => { return gp.language.go!.name === op.language.go!.optionalParamGroup.language.go!.name})) {
    paramGroups.push(op.language.go!.optionalParamGroup);
  }
  for (const paramGroup of values(paramGroups)) {
    params.push(paramGroup);
  }
  return params;
}

// returns the fully-qualified parameter name.  this is usually just the name
// but will include the client or optional param group name prefix as required.
// dereference: pass true to dereference an optional param
export function getParamName(param: Parameter): string {
  let paramName = param.language.go!.name;
  if (param.implementation === ImplementationLocation.Client) {
    paramName = `client.${paramName}`;
  } else if (param.language.go!.paramGroup) {
    paramName = `${uncapitalize(param.language.go!.paramGroup.language.go!.name)}.${capitalize(paramName)}`;
  }
  if (param.required !== true && !param.language.go!.byValue) {
    paramName = `*${paramName}`;
  }
  return paramName;
}

export function formatParamValue(param: Parameter, imports: ImportManager): string {
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
  let paramName = getParamName(param);
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
    case SchemaType.Date:
      if (param.required !== true && paramName[0] === '*') {
        // remove the dereference
        paramName = paramName.substring(1);
      }
    case SchemaType.DateTime:
      imports.add('time');
      if (param.required !== true && paramName[0] === '*') {
        // remove the dereference
        paramName = paramName.substring(1);
      }
  }
  return formatValue(paramName, param.schema, imports);
}

export function formatValue(paramName: string, schema: Schema, imports: ImportManager): string {
  switch (schema.type) {
    case SchemaType.Array:
      throw new Error(`can't format array without parameter info`);
    case SchemaType.Boolean:
      imports.add('strconv');
      return `strconv.FormatBool(${paramName})`;
    case SchemaType.ByteArray:
      // ByteArray is a base-64 encoded value in string format
      imports.add('encoding/base64');
      let byteFormat = 'Std';
      if ((<ByteArraySchema>schema).format === 'base64url') {
        byteFormat = 'RawURL';
      }
      return `base64.${byteFormat}Encoding.EncodeToString(${paramName})`;
    case SchemaType.Choice:
    case SchemaType.SealedChoice:
      return `string(${paramName})`;
    case SchemaType.Constant:
      const constSchema = <ConstantSchema>schema;
      // cannot use formatConstantValue() since all values are treated as strings
      return `"${constSchema.value.value}"`;
    case SchemaType.Date:
      return `${paramName}.Format("${dateFormat}")`;
    case SchemaType.DateTime:
      imports.add('time');
      let format = datetimeRFC3339Format;
      const dateTime = <DateTimeSchema>schema;
      if (dateTime.format === 'date-time-rfc1123') {
        format = datetimeRFC1123Format;
      }
      return `${paramName}.Format(${format})`;
    case SchemaType.UnixTime:
      return `timeUnix(${paramName}).String()`;
    case SchemaType.Integer:
      imports.add('strconv');
      const intSchema = <NumberSchema>schema;
      let intParam = paramName;
      if (intSchema.precision === 32) {
        intParam = `int64(${intParam})`;
      }
      return `strconv.FormatInt(${intParam}, 10)`;
    case SchemaType.Number:
      imports.add('strconv');
      const numberSchema = <NumberSchema>schema;
      let floatParam = paramName;
      if (numberSchema.precision === 32) {
        floatParam = `float64(${floatParam})`;
      }
      return `strconv.FormatFloat(${floatParam}, 'f', -1, ${numberSchema.precision})`;
    default:
      return paramName;
  }
}

// returns true if at least one of the responses has a schema
export function hasSchemaResponse(op: Operation): boolean {
  for (const resp of values(op.responses)) {
    if (isSchemaResponse(resp)) {
      return true;
    }
  }
  return false;
}

// returns the response envelope type name
export function getResponseEnvelopeName(op: Operation): string {
  return op.language.go!.responseEnv.language.go!.name;
}

// returns the result property for the operation or undefined if it doesn't return a model
export function hasResultProperty(op: Operation): Property | undefined {
  const responseEnv = getResponseEnvelope(op);
  if (responseEnv.language.go!.resultProp) {
    return responseEnv.language.go!.resultProp;
  }
  return undefined;
}

export function getResponseEnvelope(op: Operation): ObjectSchema {
  return op.language.go!.responseEnv;
}

// returns the name of the response field within the response envelope
export function getResultFieldName(op: Operation): string {
  if (isMultiRespOperation(op)) {
    return 'Value';
  }
  let responseEnv = op.language.go!.responseEnv;
  if (responseEnv.language.go!.resultProp.schema.serialization?.xml?.name) {
    // here we use the schema name instead of the result field name as it's anonymously embedded in the response envelope.
    // this is to handle XML cases that specify a custom XML name for the propery within the result field.
    return responseEnv.language.go!.resultProp.schema.language.go!.name;
  }
  return responseEnv.language.go!.resultProp.language.go!.name;
}

export function getStatusCodes(op: Operation): string[] {
  // concat all status codes that return the same schema into one array.
  // this is to support operations that specify multiple response codes
  // that return the same schema (or no schema).
  let statusCodes = new Array<string>();
  for (const resp of values(op.responses)) {
    statusCodes = statusCodes.concat(resp.protocol.http?.statusCodes);
  }
  if (statusCodes.length === 0) {
    // if the operation defines no status codes (which is non-conformant)
    // then add 200, 201, 202, and 204 to the list.  this is to accomodate
    // some quirky tests in the test server.
    // TODO: https://github.com/Azure/autorest.go/issues/659
    statusCodes = ['200', '201', '202', '204'];
  }
  return statusCodes;
}

export function formatStatusCodes(statusCodes: Array<string>): string {
  const asHTTPStatus = new Array<string>();
  for (const rawCode of values(statusCodes)) {
    asHTTPStatus.push(formatStatusCode(rawCode));
  }
  return asHTTPStatus.join(', ');
}

export function formatStatusCode(statusCode: string): string {
  switch (statusCode) {
    // 1xx
    case '100':
      return 'http.StatusContinue';
    case '101':
      return 'http.StatusSwitchingProtocols';
    case '102':
      return 'http.StatusProcessing';
    case '103':
      return 'http.StatusEarlyHints';
    // 2xx
    case '200':
      return 'http.StatusOK';
    case '201':
      return 'http.StatusCreated';
    case '202':
      return 'http.StatusAccepted';
    case '203':
      return 'http.StatusNonAuthoritativeInfo';
    case '204':
      return 'http.StatusNoContent';
    case '205':
      return 'http.StatusResetContent';
    case '206':
      return 'http.StatusPartialContent';
    case '207':
      return 'http.StatusMultiStatus';
    case '208':
      return 'http.StatusAlreadyReported';
    case '226':
      return 'http.StatusIMUsed';
    // 3xx
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
    case '305':
      return 'http.StatusUseProxy';
    case '307':
      return 'http.StatusTemporaryRedirect';
    // 4xx
    case '400':
      return 'http.StatusBadRequest';
    case '401':
      return 'http.StatusUnauthorized';
    case '402':
      return 'http.StatusPaymentRequired';
    case '403':
      return 'http.StatusForbidden';
    case '404':
      return 'http.StatusNotFound';
    case '405':
      return 'http.StatusMethodNotAllowed';
    case '406':
      return 'http.StatusNotAcceptable';
    case '407':
      return 'http.StatusProxyAuthRequired';
    case '408':
      return 'http.StatusRequestTimeout';
    case '409':
      return 'http.StatusConflict';
    case '410':
      return 'http.StatusGone';
    case '411':
      return 'http.StatusLengthRequired';
    case '412':
      return 'http.StatusPreconditionFailed';
    case '413':
      return 'http.StatusRequestEntityTooLarge';
    case '414':
      return 'http.StatusRequestURITooLong';
    case '415':
      return 'http.StatusUnsupportedMediaType';
    case '416':
      return 'http.StatusRequestedRangeNotSatisfiable';
    case '417':
      return 'http.StatusExpectationFailed';
    case '418':
      return 'http.StatusTeapot';
    case '421':
      return 'http.StatusMisdirectedRequest';
    case '422':
      return 'http.StatusUnprocessableEntity';
    case '423':
      return 'http.StatusLocked';
    case '424':
      return 'http.StatusFailedDependency';
    case '425':
      return 'http.StatusTooEarly';
    case '426':
      return 'http.StatusUpgradeRequired';
    case '428':
      return 'http.StatusPreconditionRequired';
    case '429':
      return 'http.StatusTooManyRequests';
    case '431':
      return 'http.StatusRequestHeaderFieldsTooLarge';
    case '451':
      return 'http.StatusUnavailableForLegalReasons';
    // 5xx
    case '500':
      return 'http.StatusInternalServerError';
    case '501':
      return 'http.StatusNotImplemented';
    case '502':
      return 'http.StatusBadGateway';
    case '503':
      return 'http.StatusServiceUnavailable';
    case '504':
      return 'http.StatusGatewayTimeout ';
    case '505':
      return 'http.StatusHTTPVersionNotSupported';
    case '506':
      return 'http.StatusVariantAlsoNegotiates';
    case '507':
      return 'http.StatusInsufficientStorage';
    case '508':
      return 'http.StatusLoopDetected';
    case '510':
      return 'http.StatusNotExtended';
    case '511':
      return 'http.StatusNetworkAuthenticationRequired';
    default:
      throw new Error(`unhandled status code ${statusCode}`);
  }
}
