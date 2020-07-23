/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { values } from '@azure-tools/linq';
import { comment, camelCase, pascalCase } from '@azure-tools/codegen';
import { aggregateParameters } from '../common/helpers';
import { ArraySchema, CodeModel, DictionarySchema, Language, Parameter, Schema, SchemaType, Operation, GroupProperty, ImplementationLocation, OperationGroup, SerializationStyle, ByteArraySchema, ConstantSchema, NumberSchema, DateTimeSchema } from '@azure-tools/codemodel';
import { ImportManager } from './imports';

export const dateFormat = '2006-01-02';
export const datetimeRFC3339Format = 'time.RFC3339';
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
  if (param.required) {
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
  for (const methodParam of values(methodParams)) {
    params.push(`${camelCase(methodParam.language.go!.name)} ${formatParameterTypeName(methodParam)}`);
  }
  return params.join(', ');
}

// returns the complete collection of method parameters
export function getMethodParameters(op: Operation): Parameter[] {
  const params = new Array<Parameter>();
  const paramGroups = new Array<GroupProperty>();
  for (const param of values(aggregateParameters(op))) {
    if (param.implementation === ImplementationLocation.Client) {
      // client params are passed via the receiver
      continue;
    } else if (param.schema.type === SchemaType.Constant) {
      // don't generate a parameter for a constant
      continue;
    } else if (param.language.go!.paramGroup) {
      // param groups will be added after individual params
      if (!paramGroups.includes(param.language.go!.paramGroup)) {
        paramGroups.push(param.language.go!.paramGroup);
      }
      continue;
    }
    params.push(param);
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
    let name = camelCase(paramGroup.language.go!.name);
    if (!paramGroup.required) {
      name = 'options';
    }
    params.push(paramGroup);
  }
  return params;
}

export function getParamName(param: Parameter, onClient: boolean): string {
  let paramName = param.language.go!.name;
  if (onClient) {
    paramName = `client.Client.${paramName}`;
  } else if (param.implementation === ImplementationLocation.Client) {
    paramName = `client.${paramName}`;
  } else if (param.language.go!.paramGroup) {
    paramName = `${camelCase(param.language.go!.paramGroup.language.go!.name)}.${pascalCase(paramName)}`;
  }
  if (param.required !== true) {
    paramName = `*${paramName}`;
  }
  return paramName;
}

export function formatParamValue(param: Parameter, imports: ImportManager, onClient: boolean): string {
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
  let paramName = getParamName(param, onClient);
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

export interface ParameterizedHost {
  addParamHost: boolean;
  urlOnClient: boolean;
  clientParams: Array<Parameter>; // contains the list of client params shared across all operation groups
}

// this function checks if parameterized host functionality needs to be added for the service
// and returns two booleans. The first boolean signals if parameterized host should be added or not
// the second boolean signals if all of the parameterized host parameters are on the client or not.
export function addParameterizedHostFunctionality(operationGroups: Array<OperationGroup>): ParameterizedHost {
  // before checking for special parameterized host conditions, we need to search through all of the
  // operaiton groups in order to know if there are different parameterized host implementations in the 
  // package. 
  let separateHosts = false; // this indicates if there are multiple parameterized host implementations
  const paramHost = operationGroups[0].operations[0].requests![0].protocol.http!.uri;
  let allClientParams = new Array<Array<Parameter>>();
  let checkSharedParams = true; // variable to control if we should keep checking for shared client params
  for (const group of values(operationGroups)) {
    const hostURI = group.operations[0].requests![0].protocol.http!.uri;
    // if we find a different parameterized host definition url parsing is done off the client
    if (hostURI !== paramHost) {
      separateHosts = true;
    }
    // store client params in one array to later filter down to all shared client params
    if (group.language.go!.clientParams !== undefined && checkSharedParams) {
      allClientParams.push(group.language.go!.clientParams);
    } else {
      if (allClientParams.length > 0) {
        // wipe all client params since one group does not have any, therefore indicating
        // that there is no client param shared among all operation groups
        allClientParams = new Array<Array<Parameter>>();
        checkSharedParams = false;
      }
    }
  }
  let sharedParams = new Array<Parameter>();
  // if the length of the array is not equal to the number of operation groups, the all operation groups
  // do not share client params. 
  if (allClientParams.length === operationGroups.length && allClientParams.length > 1) {
    // filter down to shared params, reduce repeated instances (unlikely)
    sharedParams = allClientParams.reduce((p, c) => p.filter(e => c.includes(e)));
  } else if (operationGroups.length === 1 && allClientParams.length === 1) {
    sharedParams = allClientParams[0];
  }
  // determine where client params need to be placed, client level or operation group level
  if (separateHosts || !(<string>paramHost).match(/^\{\$?\w+\}$/)) {
    let methodParamsCount = 0;
    for (const p of values(aggregateParameters(operationGroups[0].operations[0])).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'uri'; })) {
      if (!(p.implementation === ImplementationLocation.Client)) {
        methodParamsCount++;
      }
    }
    if (methodParamsCount > 0 || separateHosts) {
      return {
        addParamHost: true,
        urlOnClient: false,
        clientParams: sharedParams,
      };
    } else {
      // if all params are on the client then it could all be handled in the new client with pipeline
      return {
        addParamHost: true,
        urlOnClient: true,
        clientParams: sharedParams,
      };
    }
  }
  return {
    addParamHost: false,
    urlOnClient: false, // leave this as false so it doesn't interact with parameterized host setting that check for this condition
    clientParams: sharedParams,
  };
}
