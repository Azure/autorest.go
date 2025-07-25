/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../codemodel.go/src/index.js';
import { values } from '@azure-tools/linq';
import { capitalize, comment, uncapitalize } from '@azure-tools/codegen';
import { ImportManager } from './imports.js';
import { CodegenError } from './errors.js';

// variable to be used to determine comment length when calling comment from @azure-tools
export const commentLength = 120;

export const dateFormat = '2006-01-02';
export const datetimeRFC3339Format = 'time.RFC3339Nano';
export const datetimeRFC1123Format = 'time.RFC1123';
export const timeRFC3339Format = '15:04:05.999999999Z07:00';

/**
 * returns the common source-file preamble (license comment, package name etc)
 * 
 * @param codeModel the code model being emitted
 * @param doNotEdit when false the 'DO NOT EDIT' clause is omitted
 * @param packageName overrides the package name in the code model
 * @returns the source file preamble
 */
export function contentPreamble(codeModel: go.CodeModel, doNotEdit = true, packageName?: string): string {
  if (!packageName) {
    packageName = codeModel.packageName;
  }
  const headerText = comment(codeModel.options.headerText, '// ');
  let text = headerText;
  if (doNotEdit) {
    // ensure tools recognize the file as generated according to
    // https://pkg.go.dev/cmd/go#hdr-Generate_Go_files_by_processing_source
    text = text.replace(/^\/\/ Code generated .*\.$/m, '$& DO NOT EDIT.');
    if (!text.match(/^\/\/ Code generated .* DO NOT EDIT\.$/m)) {
      text += '\n// Code generated by @autorest/go. DO NOT EDIT.';
    }
  } else {
    // remove the blurb about the changes being lost
    text = text.replace(/^\/\/ Changes may cause incorrect behavior and will be lost if the code is regenerated\.$/m, '');
  }
  text += `\n\npackage ${packageName}\n\n`;
  return text;
}

// used to sort strings in ascending order
export function sortAscending(a: string, b: string): number {
  return a < b ? -1 : a > b ? 1 : 0;
}

// returns the type name with possible * prefix
export function formatParameterTypeName(param: go.ClientParameter | go.ParameterGroup, pkgName?: string): string {
  let typeName: string;
  switch (param.kind) {
    case 'paramGroup':
      typeName = param.groupName;
      if (pkgName) {
        typeName = `${pkgName}.${typeName}`;
      }
      if (param.required) {
        return typeName;
      }
      break;
    default:
      typeName = go.getTypeDeclaration(param.type, pkgName);
      if (parameterByValue(param)) {
        // client parameters with default values aren't emitted as pointer-to-type
        return typeName;
      }
  }
  return `*${typeName}`;
}

export function parameterByValue(param: go.ClientParameter): boolean {
  return go.isRequiredParameter(param) || (param.location === 'client' && go.isClientSideDefault(param.style))
}

// sorts parameters by their required state, ordering required before optional
export function sortParametersByRequired(a: go.ClientParameter | go.ParameterGroup, b: go.ClientParameter | go.ParameterGroup): number {
  let aRequired = false;
  let bRequired = false;

  switch (a.kind) {
    case 'paramGroup':
      aRequired = a.required;
      break;
    default:
      aRequired = go.isRequiredParameter(a);
      break;
  }

  switch (b.kind) {
    case 'paramGroup':
      bRequired = b.required;
      break;
    default:
      bRequired = go.isRequiredParameter(b);
      break;
  }

  if (aRequired === bRequired) {
    return 0;
  } else if (aRequired && !bRequired) {
    return -1;
  }
  return 1;
}

// returns the parameters for the internal request creator method.
// e.g. "i int, s string"
export function getCreateRequestParametersSig(method: go.MethodType | go.NextPageMethod): string {
  const methodParams = getMethodParameters(method);
  const params = new Array<string>();
  params.push('ctx context.Context');
  for (const methodParam of values(methodParams)) {
    let paramName = uncapitalize(methodParam.name);
    // when creating the method sig for fooCreateRequest, if the options type is empty
    // or only contains the ResumeToken param use _ for the param name to quiet the linter
    if (methodParam.kind === 'paramGroup' && (methodParam.params.length === 0 || (methodParam.params.length === 1 && methodParam.params[0].kind === 'resumeTokenParam'))) {
      paramName = '_';
    }
    params.push(`${paramName} ${formatParameterTypeName(methodParam)}`);
  }
  return params.join(', ');
}

// returns the parameter names for an operation (excludes the param types).
// e.g. "i, s"
export function getCreateRequestParameters(method: go.MethodType): string {
  // NOTE: keep in sync with getCreateRequestParametersSig
  const methodParams = getMethodParameters(method);
  const params = new Array<string>();
  params.push('ctx');
  for (const methodParam of values(methodParams)) {
    params.push(uncapitalize(methodParam.name));
  }
  return params.join(', ');
}

// returns the complete collection of method parameters
export function getMethodParameters(method: go.MethodType | go.NextPageMethod, paramsFilter?: (p: Array<go.MethodParameter>) => Array<go.MethodParameter>): Array<go.MethodParameter | go.ParameterGroup> {
  const params = new Array<go.MethodParameter>();
  const paramGroups = new Array<go.ParameterGroup>();
  let methodParams = method.parameters;
  if (paramsFilter) {
    methodParams = paramsFilter(methodParams);
  }
  for (const param of values(methodParams)) {
    if (param.location === 'client') {
      // client params are passed via the receiver
      // must check before param group as client params can be grouped
      continue;
    } else if (param.group) {
      // param groups will be added after individual params
      if (!paramGroups.includes(param.group)) {
        paramGroups.push(param.group);
      }
    } else if (param.type.kind === 'literal') {
      // don't generate a parameter for a constant
      // NOTE: this check must come last as non-required optional constants
      // in header/query params get dumped into the optional params group
      continue;
    } else {
      params.push(param);
    }
  }
  // move global optional params to the end of the slice
  params.sort(sortParametersByRequired);
  // add any parameter groups.  optional groups go last
  paramGroups.sort((a: go.ParameterGroup, b: go.ParameterGroup) => {
    if (a.required === b.required) {
      return 0;
    }
    if (a.required && !b.required) {
      return -1;
    }
    return 1;
  });
  // add the optional param group last if it's not already in the list.
  if (method.kind !== 'nextPageMethod') {
    if (!values(paramGroups).any(gp => { return gp.groupName === method.optionalParamsGroup.groupName; })) {
      paramGroups.push(method.optionalParamsGroup);
    }
  }
  const combined = new Array<go.MethodParameter | go.ParameterGroup>();
  for (const param of params) {
    combined.push(param);
  }
  for (const paramGroup of paramGroups) {
    combined.push(paramGroup);
  }
  return combined;
}

// returns the fully-qualified parameter name.  this is usually just the name
// but will include the client or optional param group name prefix as required.
export function getParamName(param: go.MethodParameter): string {
  let paramName = param.name;
  // must check paramGroup first as client params can also be grouped
  if (param.group) {
    paramName = `${uncapitalize(param.group.name)}.${capitalize(paramName)}`;
  }
  if (param.location === 'client') {
    paramName = `client.${paramName}`;
  }
  // client parameters with default values aren't emitted as pointer-to-type
  if (!go.isRequiredParameter(param) && !(param.location === 'client' && go.isClientSideDefault(param.style)) && !param.byValue) {
    paramName = `*${paramName}`;
  }
  return paramName;
}

// converts the Go code model encoding type to the type name in the standard library
export function formatBytesEncoding(enc: go.BytesEncoding): string {
  if (enc === 'URL') {
    return 'RawURL';
  }
  return 'Std';
}

export function formatParamValue(param: go.MethodParameter, imports: ImportManager): string {
  let paramName = getParamName(param);
  switch (param.kind) {
    case 'formBodyCollectionParam':
    case 'headerCollectionParam':
    case 'pathCollectionParam':
    case 'queryCollectionParam': {
      if (param.collectionFormat === 'multi') {
        throw new CodegenError('InternalError', 'multi collection format should have been previously handled');
      }

      const separator = getDelimiterForCollectionFormat(param.collectionFormat);

      const emitConvertOver = function(paramName: string, format: string): string {
        const encodedVar = `encoded${capitalize(paramName)}`;
        let content = 'strings.Join(func() []string {\n';
        content += `\t\t${encodedVar} := make([]string, len(${paramName}))\n`;
        content += `\t\tfor i := 0; i < len(${paramName}); i++ {\n`;
        content += `\t\t\t${encodedVar}[i] = ${format}\n\t\t}\n`;
        content += `\t\treturn ${encodedVar}\n`;
        content += `\t}(), "${separator}")`;
        return content;
      }

      switch (param.type.elementType.kind) {
        case 'encodedBytes':
          imports.add('encoding/base64');
          imports.add('strings');
          return emitConvertOver(param.name, `base64.${formatBytesEncoding(param.type.elementType.encoding)}Encoding.EncodeToString(${param.name}[i])`);
        case 'string':
          imports.add('strings');
          return `strings.Join(${paramName}, "${separator}")`;
        case 'time':
          imports.add('strings');
          return emitConvertOver(param.name, `${param.type.elementType.format}(${param.name}[i]).String()`);
        default:
          imports.add('fmt');
          imports.add('strings');
          return `strings.Join(strings.Fields(strings.Trim(fmt.Sprint(${paramName}), "[]")), "${separator}")`;
      }
    }
  }

  if (param.type.kind === 'time' && param.type.format !== 'timeUnix') {
    // for most time types we call methods on time.Time which is why we remove the dereference.
    // however, for unix time, we cast to our unixTime helper first so we must keep the dereference.
    if (!go.isRequiredParameter(param) && paramName[0] === '*') {
      // remove the dereference
      paramName = paramName.substring(1);
    }
  }
  return formatValue(paramName, param.type, imports);
}

export function getDelimiterForCollectionFormat(cf: go.CollectionFormat): string {
  switch (cf) {
    case 'csv':
      return ',';
    case 'pipes':
      return '|';
    case 'ssv':
      return ' ';
    case 'tsv':
      return '\\t';
    default:
      throw new CodegenError('InternalError', `unhandled CollectionFormat ${cf}`);
  }
}

export function formatValue(paramName: string, type: go.WireType, imports: ImportManager, defef?: boolean): string {
  // callers don't have enough context to know if paramName needs to be
  // deferenced so we track that here when specified. note that not all
  // cases will require paramName to be dereferenced.
  let star = '';
  if (defef === true) {
    star = '*';
  }

  switch (type.kind) {
    case 'constant':
      if (type.type === 'string') {
        return `string(${star}${paramName})`;
      }
      imports.add('fmt');
      return `fmt.Sprintf("%v", ${star}${paramName})`;
    case 'encodedBytes':
      // a base-64 encoded value in string format
      imports.add('encoding/base64');
      return `base64.${formatBytesEncoding(type.encoding)}Encoding.EncodeToString(${paramName})`;
    case 'literal':
      // cannot use formatLiteralValue() since all values are treated as strings
      return `"${type.literal}"`;
    case 'scalar':
      switch (type.type) {
        case 'bool':
          imports.add('strconv');
          return `strconv.FormatBool(${star}${paramName})`;
        case 'float32':
          imports.add('strconv');
          return `strconv.FormatFloat(float64(${star}${paramName}), 'f', -1, 32)`;
        case 'float64':
          imports.add('strconv');
          return `strconv.FormatFloat(${star}${paramName}, 'f', -1, 64)`;
        case 'int32':
          imports.add('strconv');
          return `strconv.FormatInt(int64(${star}${paramName}), 10)`;
        case 'int64':
          imports.add('strconv');
          return `strconv.FormatInt(${star}${paramName}, 10)`;
        default:
          throw new CodegenError('InternalError', `unhandled scalar type ${type.type}`);
      }
    case 'time':
      switch (type.format) {
        case 'dateTimeRFC1123':
        case 'dateTimeRFC3339':
          imports.add('time');
          return `${paramName}.Format(${type.format === 'dateTimeRFC1123' ? datetimeRFC1123Format : datetimeRFC3339Format})`;
        case 'dateType':
          return `${paramName}.Format("${dateFormat}")`;
        case 'timeRFC3339':
          return `timeRFC3339(${star}${paramName}).String()`;
        case 'timeUnix':
          return `timeUnix(${star}${paramName}).String()`;
      }
    default:
      return `${star}${paramName}`;
  }
}

// returns the clientDefaultValue of the specified param.
// this is usually the value in quotes (i.e. a string) however
// it could also be a constant.
export function formatLiteralValue(value: go.Literal, withCast: boolean): string {
  switch (value.type.kind) {
    case 'constant':
      return (<go.ConstantValue>value.literal).name;
    case 'encodedBytes':
      return value.literal;
    case 'scalar':
      if (!withCast) {
        return `${value.literal}`;
      }
      switch (value.type.type) {
        case 'float32':
          return `float32(${value.literal})`;
        case 'float64':
          return `float64(${value.literal})`;
        case 'int32':
          return `int32(${value.literal})`;
        case 'int64':
          return `int64(${value.literal})`;
        default:
          return value.literal;
      }
    case 'string':
      if (value.literal[0] === '"') {
        // string is already quoted
        return value.literal;
      }
      return `"${value.literal}"`;
    case 'time':
      return `"${value.literal}"`;
  }
}

// returns true if at least one of the responses has a schema
export function hasSchemaResponse(method: go.MethodType): boolean {
  switch (method.responseEnvelope.result?.kind) {
    case 'anyResult':
    case 'modelResult':
    case 'monomorphicResult':
    case 'polymorphicResult':
      return true;
    default:
      return false;
  }
}

// returns the name of the response field within the response envelope
export function getResultFieldName(method: go.MethodType): string {
  const result = method.responseEnvelope.result;
  if (!result) {
    throw new CodegenError('InternalError', `missing result for method ${method.name}`);
  }
  switch (result.kind) {
    case 'anyResult':
    case 'binaryResult':
    case 'headAsBooleanResult':
    case 'monomorphicResult':
      return result.fieldName;
    case 'modelResult':
      return result.modelType.name;
    case 'polymorphicResult':
      return result.interface.name;
  }
}

export function formatStatusCodes(statusCodes: Array<number>): string {
  const asHTTPStatus = new Array<string>();
  for (const rawCode of statusCodes) {
    asHTTPStatus.push(formatStatusCode(rawCode));
  }
  return asHTTPStatus.join(', ');
}

export function formatStatusCode(statusCode: number): string {
  switch (statusCode) {
    // 1xx
    case 100:
      return 'http.StatusContinue';
    case 101:
      return 'http.StatusSwitchingProtocols';
    case 102:
      return 'http.StatusProcessing';
    case 103:
      return 'http.StatusEarlyHints';
    // 2xx
    case 200:
      return 'http.StatusOK';
    case 201:
      return 'http.StatusCreated';
    case 202:
      return 'http.StatusAccepted';
    case 203:
      return 'http.StatusNonAuthoritativeInfo';
    case 204:
      return 'http.StatusNoContent';
    case 205:
      return 'http.StatusResetContent';
    case 206:
      return 'http.StatusPartialContent';
    case 207:
      return 'http.StatusMultiStatus';
    case 208:
      return 'http.StatusAlreadyReported';
    case 226:
      return 'http.StatusIMUsed';
    // 3xx
    case 300:
      return 'http.StatusMultipleChoices';
    case 301:
      return 'http.StatusMovedPermanently';
    case 302:
      return 'http.StatusFound';
    case 303:
      return 'http.StatusSeeOther';
    case 304:
      return 'http.StatusNotModified';
    case 305:
      return 'http.StatusUseProxy';
    case 307:
      return 'http.StatusTemporaryRedirect';
    // 4xx
    case 400:
      return 'http.StatusBadRequest';
    case 401:
      return 'http.StatusUnauthorized';
    case 402:
      return 'http.StatusPaymentRequired';
    case 403:
      return 'http.StatusForbidden';
    case 404:
      return 'http.StatusNotFound';
    case 405:
      return 'http.StatusMethodNotAllowed';
    case 406:
      return 'http.StatusNotAcceptable';
    case 407:
      return 'http.StatusProxyAuthRequired';
    case 408:
      return 'http.StatusRequestTimeout';
    case 409:
      return 'http.StatusConflict';
    case 410:
      return 'http.StatusGone';
    case 411:
      return 'http.StatusLengthRequired';
    case 412:
      return 'http.StatusPreconditionFailed';
    case 413:
      return 'http.StatusRequestEntityTooLarge';
    case 414:
      return 'http.StatusRequestURITooLong';
    case 415:
      return 'http.StatusUnsupportedMediaType';
    case 416:
      return 'http.StatusRequestedRangeNotSatisfiable';
    case 417:
      return 'http.StatusExpectationFailed';
    case 418:
      return 'http.StatusTeapot';
    case 421:
      return 'http.StatusMisdirectedRequest';
    case 422:
      return 'http.StatusUnprocessableEntity';
    case 423:
      return 'http.StatusLocked';
    case 424:
      return 'http.StatusFailedDependency';
    case 425:
      return 'http.StatusTooEarly';
    case 426:
      return 'http.StatusUpgradeRequired';
    case 428:
      return 'http.StatusPreconditionRequired';
    case 429:
      return 'http.StatusTooManyRequests';
    case 431:
      return 'http.StatusRequestHeaderFieldsTooLarge';
    case 451:
      return 'http.StatusUnavailableForLegalReasons';
    // 5xx
    case 500:
      return 'http.StatusInternalServerError';
    case 501:
      return 'http.StatusNotImplemented';
    case 502:
      return 'http.StatusBadGateway';
    case 503:
      return 'http.StatusServiceUnavailable';
    case 504:
      return 'http.StatusGatewayTimeout ';
    case 505:
      return 'http.StatusHTTPVersionNotSupported';
    case 506:
      return 'http.StatusVariantAlsoNegotiates';
    case 507:
      return 'http.StatusInsufficientStorage';
    case 508:
      return 'http.StatusLoopDetected';
    case 510:
      return 'http.StatusNotExtended';
    case 511:
      return 'http.StatusNetworkAuthenticationRequired';
    default:
      throw new CodegenError('InternalError', `unhandled status code ${statusCode}`);
  }
}

export function formatCommentAsBulletItem(prefix: string, docs: go.Docs): string {
  // first create the comment block. note that it can be multi-line depending on length:
  //
  // some comment first line
  // and it finishes here.
  let description = formatDocCommentWithPrefix(prefix, docs);
  if (description.length === 0) {
    return '';
  }

  // transform the above to look like this:
  //
  //   - some comment first line
  //     and it finishes here.
  const chunks = description.split('\n');
  for (let i = 0; i < chunks.length; ++i) {
    if (i === 0) {
      chunks[i] = chunks[i].replace('// ', '//   - ');
    } else {
      chunks[i] = chunks[i].replace('// ', '//     ');
    }
  }
  return chunks.join('\n');
}

// conditionally returns a doc comment on an entity that requires a prefix.
// e.g.:
// {Prefix} - {docs.summary}
//
// {docs.description}
export function formatDocCommentWithPrefix(prefix: string, docs: go.Docs): string {
  if (!docs.summary && !docs.description) {
    return '';
  }

  let docComment = '';
  if (docs.summary) {
    docComment = `${comment(`${prefix} - ${docs.summary}`, '//', undefined, commentLength)}\n`;
  }

  if (docs.description) {
    let description = docs.description;
    if (docs.summary) {
      docComment += '//\n';
    } else {
      // only apply the prefix to the description if there was no summary
      description = `${prefix} - ${description}`;
    }
    docComment += `${comment(`${description}`, '//', undefined, commentLength)}\n`;
  }

  return docComment;
}

// conditionally returns a doc comment
// {docs.summary}
//
// {docs.description}
export function formatDocComment(docs: go.Docs): string {
  if (!docs.summary && !docs.description) {
    return '';
  }

  let docComment = '';
  if (docs.summary) {
    docComment = `${comment(docs.summary, '//', undefined, commentLength)}\n`;
  }

  if (docs.description) {
    if (docs.summary) {
      docComment += '//\n';
    }
    docComment += `${comment(docs.description, '//', undefined, commentLength)}\n`;
  }

  return docComment;
}

export function getParentImport(codeModel: go.CodeModel): string {
  const clientPkg = codeModel.packageName;
  if (codeModel.options.module) {
    return codeModel.options.module;
  } else if (codeModel.options.containingModule) {
    return codeModel.options.containingModule + '/' + clientPkg;
  } else {
    throw new CodegenError('InvalidArgument', 'unable to determine containing module for fakes. specify either the module or containing-module switch');
  }
}

export function getBitSizeForNumber(intSize: 'float32' | 'float64' | 'int8' | 'int16' | 'int32' | 'int64'): string {
  switch (intSize) {
    case 'int8':
      return '8';
    case 'int16':
      return '16';
    case 'int32':
    case 'float32':
      return '32';
    case 'int64':
    case 'float64':
      return '64';
  }
}

// returns the underlying map/slice element/value type
// if item isn't a map or slice, item is returned
export function recursiveUnwrapMapSlice(item: go.WireType): go.WireType {
  switch (item.kind) {
    case 'map':
      return recursiveUnwrapMapSlice(item.valueType);
    case 'slice':
      return recursiveUnwrapMapSlice(item.elementType);
    default:
      return item;
  }
}

// returns a * for optional params
export function star(param: go.MethodParameter): string {
  return go.isRequiredParameter(param) || param.byValue ? '' : '*';
}

export type SerDeFormat = 'JSON' | 'XML';

// used by getSerDeFormat to cache results
const serDeFormatCache = new Map<string, SerDeFormat>();

// returns the wire format for the named model.
// at present this assumes the formats to be mutually exclusive.
export function getSerDeFormat(model: go.Model | go.PolymorphicModel, codeModel: go.CodeModel): SerDeFormat {
  let serDeFormat = serDeFormatCache.get(model.name);
  if (serDeFormat) {
    return serDeFormat;
  }

  // for model-only builds we assume the format to be JSON
  if (codeModel.clients.length === 0) {
    return 'JSON';
  }

  // recursively walks the fields in model, updating serDeFormatCache with the model name and specified format
  const recursiveWalkModelFields = function(type: go.WireType, serDeFormat: SerDeFormat): void {
    type = recursiveUnwrapMapSlice(type);
    switch (type.kind) {
      case 'interface':
        recursiveWalkModelFields(type.rootType, serDeFormat);
        for (const possibleType of type.possibleTypes) {
          recursiveWalkModelFields(possibleType, serDeFormat);
        }
        break;
      case 'model':
      case 'polymorphicModel':
        if (serDeFormatCache.has(type.name)) {
          // we've already processed this type, don't do it again
          return;
        }

        serDeFormatCache.set(type.name, serDeFormat);
        for (const field of type.fields) {
          const fieldType = recursiveUnwrapMapSlice(field.type);
          recursiveWalkModelFields(fieldType, serDeFormat);
        }
        break;
    }
  };

  // walk the methods, indexing the model formats
  for (const client of codeModel.clients) {
    for (const method of client.methods) {
      for (const param of method.parameters) {
        if (param.kind !== 'bodyParam' || (param.bodyFormat !== 'JSON' && param.bodyFormat !== 'XML')) {
          continue;
        }

        recursiveWalkModelFields(param.type, param.bodyFormat);
      }

      const resultType = method.responseEnvelope.result;
      switch (resultType?.kind) {
        case 'anyResult':
          if (resultType.format === 'JSON' || resultType.format === 'XML') {
            for (const type of Object.values(resultType.httpStatusCodeType)) {
              recursiveWalkModelFields(type, resultType.format);
            }
          }
          break;
        case 'modelResult':
          recursiveWalkModelFields(resultType.modelType, resultType.format);
          break;
        case 'monomorphicResult':
          if (resultType.format === 'JSON' || resultType.format === 'XML') {
            recursiveWalkModelFields(resultType.monomorphicType, resultType.format);
          }
          break;
        case 'polymorphicResult':
          recursiveWalkModelFields(resultType.interface, resultType.format);
          break;
      }
    }
  }

  serDeFormat = serDeFormatCache.get(model.name);
  if (!serDeFormat) {
    // if we get here there are two possibilities
    //  - we have a bug in the above indexing
    //  - the model type is unreferenced by any operation
    //
    // while the former is possible, the latter has the potential to be real.
    // regardless of the cause, we will just assume the format to be JSON.
    serDeFormat = 'JSON';
  }
  return serDeFormat;
}

// return combined client parameters for all the clients
export function getAllClientParameters(codeModel: go.CodeModel): Array<go.ClientParameter> {
  const allClientParams = new Array<go.ClientParameter>();
  for (const clients of codeModel.clients) {
    for (const clientParam of values(clients.parameters)) {
      if (values(allClientParams).where(param => param.name === clientParam.name).any()) {
        continue;
      }
      allClientParams.push(clientParam);
    }
  }
  allClientParams.sort(sortParametersByRequired);
  return allClientParams;
}

// returns common client parameters for all the clients
export function getCommonClientParameters(codeModel: go.CodeModel): Array<go.ClientParameter> {
  const paramCount = new Map<string, { uses: number, param: go.ClientParameter }>();
  let numClients = 0; // track client count since we might skip some
  for (const clients of codeModel.clients) {
    // special cases: some ARM clients always don't contain any parameters (OperationsClient will be depracated in the future)
    if (codeModel.type === 'azure-arm' && clients.name.match(/^OperationsClient$/)) {
      continue; 
    }

    ++numClients;
    for (const clientParam of values(clients.parameters)) {
      let entry = paramCount.get(clientParam.name);
      if (!entry) {
        entry = { uses: 0, param: clientParam };
        paramCount.set(clientParam.name, entry);
      }

      ++entry.uses;
    }
  }

  // for each param, if its usage count is equal to the
  // number of clients, then it's common to all clients
  const commonClientParams = new Array<go.ClientParameter>();
  for (const entry of paramCount.values()) {
    if (entry.uses === numClients) {
      commonClientParams.push(entry.param);
    }
  }

  return commonClientParams.sort(sortParametersByRequired);
}

/**
 * groups method parameters based on their kind.
 * note that the body param kinds are mutually exclusive.
 */
export interface MethodParamGroups {
  /** the body parameter if applicable */
  bodyParam?: go.BodyParameter;

  /** encoded query params. can be empty */
  encodedQueryParams: Array<go.QueryParameter>;

  /** form body params. can be empty */
  formBodyParams: Array<go.FormBodyParameter>;

  /** head params. can be empty */
  headerParams: Array<go.HeaderParameter>;

  /** multipart-form body params. can be empty */
  multipartBodyParams: Array<go.MultipartFormBodyParameter>;

  /** path params. can be empty */
  pathParams: Array<go.PathParameter>;

  /** partial body params. can be empty */
  partialBodyParams: Array<go.PartialBodyParameter>;

  /** unencoded query params. can be empty */
  unencodedQueryParams: Array<go.QueryParameter>;
}

/**
 * enumerates method parameters and returns them based on kinds
 * 
 * @param method the method containing the parameters to group
 * @returns the groups of parameters
 */
export function getMethodParamGroups(method: go.MethodType | go.NextPageMethod): MethodParamGroups {
  let bodyParam: go.BodyParameter | undefined;
  const encodedQueryParams = new Array<go.QueryParameter>();
  const formBodyParams = new Array<go.FormBodyParameter>();
  const headerParams = new Array<go.HeaderParameter>();
  const multipartBodyParams = new Array<go.MultipartFormBodyParameter>();
  const pathParams = new Array<go.PathParameter>();
  const partialBodyParams = new Array<go.PartialBodyParameter>();
  const unencodedQueryParams = new Array<go.QueryParameter>();

  for (const param of method.parameters) {
    switch (param.kind) {
      case 'bodyParam':
        bodyParam = param;
        break;
      case 'formBodyCollectionParam':
      case 'formBodyScalarParam':
        formBodyParams.push(param);
        break;
      case 'headerCollectionParam':
      case 'headerMapParam':
      case 'headerScalarParam':
        headerParams.push(param);
        break;
      case 'multipartFormBodyParam':
        multipartBodyParams.push(param);
        break;
      case 'partialBodyParam':
        partialBodyParams.push(param);
        break;
      case 'pathCollectionParam':
      case 'pathScalarParam':
        pathParams.push(param);
        break;
      case 'queryCollectionParam':
      case 'queryScalarParam':
        if (param.isEncoded) {
          encodedQueryParams.push(param);
        } else {
          unencodedQueryParams.push(param);
        }
        break;
    }
  }

  return {
    bodyParam,
    encodedQueryParams,
    formBodyParams,
    headerParams,
    multipartBodyParams,
    pathParams,
    partialBodyParams,
    unencodedQueryParams,
  }
}

/** helper for managing indentation levels */
export class indentation {
  private level: number;
  constructor(level?: number) {
    if (level !== undefined) {
      this.level = level;
    } else {
      // default to one level of indentation
      this.level = 1;
    }
  }

  /**
   * returns spaces for the current indentation level
   * 
   * @returns a string with the current indentation level
   */
  get(): string {
    let indent = '';
    for (let i = 0; i < this.level; ++i) {
      indent += '\t';
    }
    return indent;
  }

  /**
   * increments the indentation level
   * 
   * @returns this indentation instance
   */
  push(): indentation {
    ++this.level;
    return this;
  }

  /**
   * decrements the indentation level
   * 
   * @returns this indentation instance
   */
  pop(): indentation {
    --this.level;
    if (this.level < 0) {
      throw new CodegenError('InternalError', 'indentation stack underflow');
    }
    return this;
  }
}

/** the if condition in an if block */
export interface ifBlock {
  /** the condition in the if block */
  condition: string;

  /** the body of the if block */
  body: (indent: indentation) => string;
}

/** the else condition in an if/else block */
export interface elseBlock {
  /** the body of the else block */
  body: (indent: indentation) => string;
}

/**
 * constructs an if block (can expand to include else if as necessary)
 * 
 * @param indent the current indentation helper in scope
 * @param ifBlock the if block definition
 * @param elseBlock optional else block definition
 * @returns the text for the if block
 */
export function buildIfBlock(indent: indentation, ifBlock: ifBlock, elseBlock?: elseBlock): string {
  let body = `if ${ifBlock.condition} {\n`;
  body += ifBlock.body(indent.push());
  body += `${indent.pop().get()}}`;

  if (elseBlock) {
    body += ' else {\n';
    body += elseBlock.body(indent.push());
    body += `${indent.pop().get()}}`;
  }

  return body;
}

/**
 * constructs an "if err != nil { return something }" block
 * 
 * @param indent the current indentation helper in scope
 * @param errVar the name of the error variable used in the condition
 * @param returns the value(s) to return from the control block
 * @returns the text for the error check block
 */
export function buildErrCheck(indent: indentation, errVar: string, returns: string): string {
  let body = `if ${errVar} != nil {\n`;
  body += `${indent.push().get()}return ${returns}\n`;
  body += `${indent.pop().get()}}`;
  return body;
}
