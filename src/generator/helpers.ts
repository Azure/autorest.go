/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { values } from '@azure-tools/linq';
import { comment, camelCase, pascalCase } from '@azure-tools/codegen';
import { aggregateParameters, isSchemaResponse } from '../common/helpers';
import { ArraySchema, CodeModel, DictionarySchema, Language, Parameter, Schema, SchemaType, Operation, GroupProperty, ImplementationLocation, SerializationStyle, ByteArraySchema, ConstantSchema, NumberSchema, DateTimeSchema } from '@azure-tools/codemodel';
import { ImportManager } from './imports';

export const dateFormat = '2006-01-02';
export const datetimeRFC3339Format = 'time.RFC3339Nano';
export const datetimeRFC1123Format = 'time.RFC1123';

// returns the common source-file preamble (license comment, package name etc)
export async function contentPreamble(session: Session<CodeModel>): Promise<string> {
  const headerText = comment(await session.getValue('header-text', 'MISSING LICENSE HEADER'), '// ');
  let text = `${headerText}\n\n`;
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

// returns the type name with possible * prefix
export function formatParameterTypeName(param: Parameter): string {
  const typeName = substituteDiscriminator(param.schema);
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

// if a field is a discriminator use the interface type instead
export function substituteDiscriminator(schema: Schema): string {
  switch (schema.type) {
    case SchemaType.Array:
      const arraySchema = <ArraySchema>schema;
      const arrayElem = <Schema>arraySchema.elementType;
      return `[]${substituteDiscriminator(arrayElem)}`;
    case SchemaType.Dictionary:
      const dictSchema = <DictionarySchema>schema;
      const dictElem = <Schema>dictSchema.elementType;
      return `map[string]${substituteDiscriminator(dictElem)}`;
    case SchemaType.Object:
      if (schema.language.go!.discriminatorInterface) {
        return schema.language.go!.discriminatorInterface;
      }
      return schema.language.go!.name;
    default:
      return schema.language.go!.name;
  }
}

// returns the parameters for the internal request creator method.
// e.g. "i int, s string"
export function getCreateRequestParametersSig(op: Operation): string {
  const methodParams = getMethodParameters(op);
  const params = new Array<string>();
  params.push('ctx context.Context');
  for (const methodParam of values(methodParams)) {
    params.push(`${camelCase(methodParam.language.go!.name)} ${formatParameterTypeName(methodParam)}`);
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
  // this handles the case where the operation has no optional params
  // but has the optional params placeholder type.
  if (paramGroups.length === 0 && op.language.go!.optionalParamGroup) {
    paramGroups.push(op.language.go!.optionalParamGroup);
  }
  // move global optional params to the end of the slice
  params.sort(sortParametersByRequired);
  // add any parameter groups.  optional group goes last
  paramGroups.sort((a: GroupProperty, b: GroupProperty) => {
    if (a.required === b.required) {
      return 0;
    }
    if (a.required && !b.required) {
      return -1;
    }
    return 1;
  })
  for (const paramGroup of values(paramGroups)) {
    // if there's only one optional param group, name the param "options" instead of its (long) type name
    if (!paramGroup.required && paramGroups.length === 1) {
      paramGroup.language.go!.name = 'options';
    }
    params.push(paramGroup);
  }
  return params;
}

export function getParamName(param: Parameter): string {
  let paramName = param.language.go!.name;
  if (param.implementation === ImplementationLocation.Client) {
    paramName = `client.${paramName}`;
  } else if (param.language.go!.paramGroup) {
    paramName = `${camelCase(param.language.go!.paramGroup.language.go!.name)}.${pascalCase(paramName)}`;
  }
  if (param.required !== true) {
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

// returns true if at least one of the responses has a schema
export function hasSchemaResponse(op: Operation): boolean {
  for (const resp of values(op.responses)) {
    if (isSchemaResponse(resp)) {
      return true;
    }
  }
  return false;
}

export function getStatusCodes(op: Operation): string[] {
  // concat all status codes that return the same schema into one array.
  // this is to support operations that specify multiple response codes
  // that return the same schema (or no schema).
  let statusCodes = new Array<string>();
  for (const resp of values(op.responses)) {
    statusCodes = statusCodes.concat(resp.protocol.http?.statusCodes);
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
