/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { capitalize, comment, KnownMediaType, uncapitalize } from '@azure-tools/codegen';
import { ApiVersions, ArraySchema, ByteArraySchema, ChoiceSchema, CodeModel, ConstantSchema, DateTimeSchema, DictionarySchema, GroupProperty, ImplementationLocation, NumberSchema, Operation, OperationGroup, Parameter, Property, Protocols, Response, Schema, SchemaResponse, SchemaType } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { aggregateParameters, formatConstantValue, getSchemaResponse, isArraySchema, isBinaryResponseOperation, isMultiRespOperation, isPageableOperation, isSchemaResponse, isTypePassedByValue, isLROOperation, commentLength } from '../common/helpers';
import { OperationNaming } from '../transform/namer';
import { contentPreamble, formatCommentAsBulletItem, formatParameterTypeName, formatStatusCodes, formatValue, getClientDefaultValue, getResponseEnvelope, getResponseEnvelopeName, getResultFieldName, getStatusCodes, hasDescription, hasResultProperty, hasSchemaResponse, skipURLEncoding, sortAscending, getCreateRequestParameters, getCreateRequestParametersSig, getMethodParameters, getParamName, formatParamValue, dateFormat, datetimeRFC1123Format, datetimeRFC3339Format, sortParametersByRequired } from './helpers';
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
  const azureARM = <boolean>session.model.language.go!.azureARM;
  // generate protocol operations
  const operations = new Array<OperationGroupContent>();
  for (const group of values(session.model.operationGroups)) {
    // the list of packages to import
    const imports = new ImportManager();
    // add standard imports
    imports.add('net/http');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/policy');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');

    let clientPkg = 'azcore';
    if (azureARM) {
      clientPkg = 'arm';
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/arm');
    } else {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    }

    // generate client type

    let clientText = '';
    let clientName = group.language.go!.clientName;
    clientText += `// ${clientName} contains the methods for the ${group.language.go!.name} group.\n`;
    clientText += `// Don't use this type directly, use `;
    if (azureARM) {
      clientText += `${group.language.go!.clientCtorName}() instead.\n`;
    } else {
      clientText += 'a constructor function instead.\n';
    }
    clientText += `type ${clientName} struct {\n`;
    clientText += `\tinternal *${clientPkg}.Client\n`;
    if (azureARM) {
      group.language.go!.hostParamName = 'internal.Endpoint()';
    } else if (group.language.go!.complexHostParams) {
      // for the complex case, all the host params must be stashed on
      // the client as the full URL is constructed in the operations.
      // MUST check before non-complex host params case.
      const hostParams = <Array<Parameter>>group.language.go!.hostParams;
      for (const param of values(hostParams)) {
        clientText += `\t${param.language.go!.name} ${param.schema.language.go!.name}\n`;
      }
    } else if (group.language.go!.hostParams) {
      // non-complex case.  the final endpoint URL will be constructed
      // from the host param(s) in the client constructor and placed here.
      group.language.go!.hostParamName = 'endpoint';
      clientText += `\t${group.language.go!.hostParamName} string\n`;
    }

    // check for any optional host params
    const optionalParams = new Array<Parameter>();
    if (group.language.go!.hostParams) {
      // client parameterized host
      const hostParams = <Array<Parameter>>group.language.go!.hostParams;
      for (const param of values(hostParams)) {
        if (param.clientDefaultValue || param.required === false) {
          optionalParams.push(param);
        }
      }
    }

    // now emit any client params (non parameterized host params case)
    if (group.language.go!.clientParams) {
      const clientParams = <Array<Parameter>>group.language.go!.clientParams;
      for (const clientParam of values(clientParams)) {
        clientText += `\t${clientParam.language.go!.name} `;
        if (clientParam.required) {
          clientText += `${clientParam.schema.language.go!.name}\n`;
        } else {
          clientText += `${formatParameterTypeName(clientParam)}\n`;
        }
        if (clientParam.clientDefaultValue || clientParam.required === false) {
          optionalParams.push(clientParam);
        }
      }
    }
    clientText += '}\n\n';

    if (azureARM && optionalParams.length > 0) {
      throw new Error('optional client parameters for ARM is not supported');
    }

    // generate client constructor (we do this only for ARM)

    if (azureARM) {
      // build constructor params
      const emitClientParams = function() {
        if (group.language.go!.clientParams) {
          const clientParams = <Array<Parameter>>group.language.go!.clientParams;
          clientParams.sort(sortParametersByRequired);
          for (const clientParam of values(clientParams)) {
            methodParams.push(`${clientParam.language.go!.name} ${formatParameterTypeName(clientParam)}`);
            if (clientParam.language.go!.description) {
              paramDocs.push(formatCommentAsBulletItem(`${clientParam.language.go!.name} - ${clientParam.language.go!.description}`));
            }
          }
        }
      }

      const methodParams = new Array<string>();
      const paramDocs = new Array<string>();
      // AzureARM is the simplest case, no parametertized host etc
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
      emitClientParams();
      methodParams.push('credential azcore.TokenCredential');
      paramDocs.push(formatCommentAsBulletItem('credential - used to authorize requests. Usually a credential from azidentity.'));
      methodParams.push(`options *arm.ClientOptions`);
      paramDocs.push(formatCommentAsBulletItem('options - pass nil to accept the default values.'));

      // now build constructor
      clientText += `// ${group.language.go!.clientCtorName} creates a new instance of ${clientName} with the specified values.\n`;
      for (const doc of values(paramDocs)) {
        clientText += `${doc}\n`;
      }

      clientText += `func ${group.language.go!.clientCtorName}(${methodParams.join(', ')}) (*${clientName}, error) {\n`;
      clientText += `\tcl, err := ${clientPkg}.NewClient(moduleName+".${clientName}", moduleVersion, credential, options)\n`;
      clientText += "\tif err != nil {\n"
      clientText += '\t\treturn nil, err\n';
      clientText += '\t}\n';

      // construct client literal
      clientText += `\tclient := &${clientName}{\n`;
      for (const clientParam of values(<Array<Parameter>>group.language.go!.clientParams)) {
        clientText += `\t\t${clientParam.language.go!.name}: ${clientParam.language.go!.name},\n`;
      }
      clientText += '\tinternal: cl,\n';
      clientText += '\t}\n';
      clientText += '\treturn client, nil\n';
      clientText += '}\n\n';
    }

    // generate operations
    const injectSpans = await session.getValue('inject-spans', false);
    let opText = '';
    group.operations.sort((a: Operation, b: Operation) => { return sortAscending(a.language.go!.name, b.language.go!.name) });
    for (const op of values(group.operations)) {
      // protocol creation can add imports to the list so
      // it must be done before the imports are written out
      if (isLROOperation(op)) {
        // generate Begin method
        opText += generateLROBeginMethod(op, injectSpans, imports);
      }
      opText += generateOperation(op, injectSpans, imports);
      opText += createProtocolRequest(group, op, imports);
      if (!isLROOperation(op) || isPageableOperation(op)) {
        // LRO responses are handled elsewhere, with the exception of pageable LROs
        opText += createProtocolResponse(op, imports);
      }
    }

    // stitch it all together
    let text = await contentPreamble(session);
    text += imports.text();
    text += clientText;
    text += opText;
    operations.push(new OperationGroupContent(group.language.go!.clientName, text));
  }
  return operations;
}

// use this to generate the code that will help process values returned in response headers
function formatHeaderResponseValue(propName: string, header: string, schema: Schema, imports: ImportManager, respObj: string, zeroResp: string): string {
  // dictionaries are handled slightly different so we do that first
  if (schema.type === SchemaType.Dictionary) {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
    imports.add('strings');
    const headerPrefix = schema.language.go!.headerCollectionPrefix;
    let text = '\tfor hh := range resp.Header {\n';
    text += `\t\tif len(hh) > len("${headerPrefix}") && strings.EqualFold(hh[:len("${headerPrefix}")], "${headerPrefix}") {\n`;
    text += `\t\t\tif ${respObj}.Metadata == nil {\n`;
    text += `\t\t\t\t${respObj}.Metadata = map[string]*string{}\n`;
    text += '\t\t\t}\n';
    text += `\t\t\t${respObj}.Metadata[hh[len("${headerPrefix}"):]] = to.Ptr(resp.Header.Get(hh))\n`;
    text += '\t\t}\n';
    text += '\t}\n';
    return text;
  }
  let text = `\tif val := resp.Header.Get("${header}"); val != "" {\n`;
  const name = uncapitalize(propName);
  let byRef = '&';
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
      byRef = '';
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
      throw new Error(`unsupported header type ${schema.type}`);
  }
  text += `\t\tif err != nil {\n`;
  text += `\t\t\treturn ${zeroResp}, err\n`;
  text += `\t\t}\n`;
  text += `\t\t${respObj}.${propName} = ${byRef}${name}\n`;
  text += '\t}\n';
  return text;
}

function getZeroReturnValue(op: Operation, apiType: 'api' | 'op' | 'handler'): string {
  let returnType = `${getResponseEnvelopeName(op)}{}`;
  if (isLROOperation(op)) {
    if (apiType === 'api' || apiType === 'op') {
      // the api returns a *Poller[T]
      // the operation returns an *http.Response
      returnType = 'nil';
    } else if (apiType === 'handler' && isPageableOperation(op)) {
      returnType = `${getResponseEnvelopeName(op)}{}`;
    }
  }
  return returnType
}

// returns true if the response contains any headers
function responseHasHeaders(op: Operation): boolean {
  const respEnv = getResponseEnvelope(op);
  for (const prop of values(respEnv.properties)) {
    if (prop.language.go!.fromHeader) {
      return true;
    }
  }
  return false;
}

function emitPagerDefinition(op: Operation, injectSpans: boolean, imports: ImportManager): string {
  const info = <OperationNaming>op.language.go!;
  const nextLink = op.language.go!.paging.nextLinkName;
  imports.add('context');
  let text = `runtime.NewPager(runtime.PagingHandler[${getResponseEnvelopeName(op)}]{\n`;
  text += `\t\tMore: func(page ${getResponseEnvelopeName(op)}) bool {\n`;
  // there is no advancer for single-page pagers
  if (op.language.go!.paging.nextLinkName) {
    text += `\t\t\treturn page.${nextLink} != nil && len(*page.${nextLink}) > 0\n`;
    text += '\t\t},\n';
  } else {
    text += `\t\t\treturn false\n`;
    text += '\t\t},\n';
  }
  text += `\t\tFetcher: func(ctx context.Context, page *${getResponseEnvelopeName(op)}) (${getResponseEnvelopeName(op)}, error) {\n`;
  const reqParams = getCreateRequestParameters(op);
  if (op.language.go!.paging.nextLinkName) {
    const isLRO = isLROOperation(op);
    const defineOrAssign = isLRO ? ':=' : '=';
    if (!isLRO) {
      text += '\t\t\tvar req *policy.Request\n';
      text += '\t\t\tvar err error\n';
      text += '\t\t\tif page == nil {\n';
      text += `\t\t\t\treq, err = client.${info.protocolNaming.requestMethod}(${reqParams})\n`;
      text += '\t\t\t} else {\n';
    }
    // nextLinkOperation might be absent in some cases, see https://github.com/Azure/autorest/issues/4393
    if (op.language.go!.paging.nextLinkOperation) {
      const nextOpParams = getCreateRequestParametersSig(op.language.go!.paging.nextLinkOperation).split(',');
      // keep the parameter names from the name/type tuples and find nextLink param
      for (let i = 0; i < nextOpParams.length; ++i) {
        const paramName = nextOpParams[i].trim().split(' ')[0];
        const paramType = nextOpParams[i].trim().split(' ')[1];
        if (paramName.startsWith('next') && paramType === 'string') {
          nextOpParams[i] = `*page.${nextLink}`;
        } else {
          nextOpParams[i] = paramName;
        }
      }
      text += `\t\t\t\treq, err ${defineOrAssign} client.${op.language.go!.paging.member}CreateRequest(${nextOpParams.join(', ')})\n`;
    } else {
      text += `\t\t\t\treq, err ${defineOrAssign} runtime.NewRequest(ctx, http.MethodGet, *page.${nextLink})\n`;
    }
    if (!isLRO) {
      text += '\t\t\t}\n';
    }
  } else {
    // this is the singular page case
    text += `\t\t\treq, err := client.${info.protocolNaming.requestMethod}(${reqParams})\n`;
  }
  text += '\t\t\tif err != nil {\n';
  text += `\t\t\t\treturn ${getResponseEnvelopeName(op)}{}, err\n`;
  text += '\t\t\t}\n';
  text += '\t\t\tresp, err := client.internal.Pipeline().Do(req)\n';
  text += '\t\t\tif err != nil {\n';
  text += `\t\t\t\treturn ${getResponseEnvelopeName(op)}{}, err\n`;
  text += '\t\t\t}\n';
  text += '\t\t\tif !runtime.HasStatusCode(resp, http.StatusOK) {\n';
  text += `\t\t\t\treturn ${getResponseEnvelopeName(op)}{}, runtime.NewResponseError(resp)\n`;
  text += '\t\t\t}\n';
  text += `\t\t\treturn client.${info.protocolNaming.responseMethod}(resp)\n`;
  text += '\t\t},\n';
  if (injectSpans) {
    text += '\t\tTracer: client.internal.Tracer(),\n';
  }
  text += `\t})\n`;
  return text;
}

function genApiVersionDoc(apiVersions?: ApiVersions): string {
  if (!apiVersions) {
    return '';
  }
  const versions = new Array<string>();
  apiVersions.forEach((val) => {
    versions.push(val.version);
  })
  return `//\n// Generated from API version ${versions.join(',')}\n`;
}

function generateOperation(op: Operation, injectSpans: boolean, imports: ImportManager): string {
  if (op.language.go!.paging && op.language.go!.paging.isNextOp) {
    // don't generate a public API for the methods used to advance pages
    return '';
  }
  const info = <OperationNaming>op.language.go!;
  const params = getAPIParametersSig(op, imports);
  const returns = generateReturnsInfo(op, 'op');
  const clientName = op.language.go!.clientName;
  let opName = op.language.go!.name;
  if(isPageableOperation(op) && !isLROOperation(op)) {
    opName = `New${opName}Pager`;
  }
  let text = '';
  if (hasDescription(op.language.go!)) {
    text += `${comment(`${opName} - ${op.language.go!.description}`, "//", undefined, commentLength)}\n`;
    text += genApiVersionDoc(op.apiVersions);
  }
  if (isLROOperation(op)) {
    opName = info.protocolNaming.internalMethod;
  } else {
    const methodParams = getMethodParameters(op);
    for (const param of values(methodParams)) {
      if (param.language.go!.description) {
        text += `${formatCommentAsBulletItem(`${param.language.go!.name} - ${param.language.go!.description}`)}\n`;
      }
    }
  }
  text += `func (client *${clientName}) ${opName}(${params}) (${returns.join(', ')}) {\n`;
  const reqParams = getCreateRequestParameters(op);
  const statusCodes = getStatusCodes(op);
  if (isPageableOperation(op) && !isLROOperation(op)) {
    text += '\treturn ';
    text += emitPagerDefinition(op, injectSpans, imports);
    text += '}\n\n';
    return text;
  }
  text += '\tvar err error\n';
  if (!isLROOperation(op) && injectSpans) {
    text += `\tctx, endSpan := runtime.StartSpan(ctx, "${clientName}.${opName}", client.internal.Tracer(), nil)\n`;
    text += '\tdefer func() { endSpan(err) }()\n';
  }
  const zeroResp = getZeroReturnValue(op, 'op');
  text += `\treq, err := client.${info.protocolNaming.requestMethod}(${reqParams})\n`;
  text += `\tif err != nil {\n`;
  text += `\t\treturn ${zeroResp}, err\n`;
  text += `\t}\n`;
  text += `\thttpResp, err := client.internal.Pipeline().Do(req)\n`;
  text += `\tif err != nil {\n`;
  text += `\t\treturn ${zeroResp}, err\n`;
  text += `\t}\n`;
  text += `\tif !runtime.HasStatusCode(httpResp, ${formatStatusCodes(statusCodes)}) {\n`;
  text += '\t\terr = runtime.NewResponseError(httpResp)\n';
  text += `\t\treturn ${zeroResp}, err\n`;
  text += '\t}\n';
  // HAB with headers response is handled in protocol responder
  if (op.language.go!.headAsBoolean && !responseHasHeaders(op)) {
    text += `\treturn ${getResponseEnvelopeName(op)}{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil\n`;
  } else {
    if (isLROOperation(op)) {
      text += '\treturn httpResp, nil\n';
    } else if (needsResponseHandler(op)) {
      // also cheating here as at present the only param to the responder is an http.Response
      text += `\tresp, err := client.${info.protocolNaming.responseMethod}(httpResp)\n`;
      text += '\treturn resp, err\n';
    } else if (isBinaryResponseOperation(op)) {
      text += `\treturn ${getResponseEnvelopeName(op)}{Body: httpResp.Body}, nil\n`;
    } else {
      text += `\treturn ${getResponseEnvelopeName(op)}{}, nil\n`;
    }
  }
  text += '}\n\n';
  return text;
}

function createProtocolRequest(group: OperationGroup, op: Operation, imports: ImportManager): string {
  const info = <OperationNaming>op.language.go!;
  const name = info.protocolNaming.requestMethod;
  for (const param of values(aggregateParameters(op))) {
    if (param.implementation !== ImplementationLocation.Method || param.required !== true) {
      continue;
    }
    imports.addImportForSchemaType(param.schema);
  }
  const returns = ['*policy.Request', 'error'];
  let text = `${comment(name, '// ')} creates the ${info.name} request.\n`;
  text += `func (client *${op.language.go!.clientName}) ${name}(${getCreateRequestParametersSig(op)}) (${returns.join(', ')}) {\n`;
  let hostParam: string;
  if (group.language.go!.complexHostParams) {
    imports.add('strings');
    // we have a complex parameterized host
    text += `\thost := "${op.requests![0].protocol.http!.uri}"\n`;
    // get all the host params on the client
    const hostParams = <Array<Parameter>>group.language.go!.hostParams;
    for (const hostParam of values(hostParams)) {
      text += `\thost = strings.ReplaceAll(host, "{${hostParam.language.go!.serializedName}}", client.${(<string>hostParam.language.go!.name)})\n`;
    }
    // check for any method local host params
    for (const param of values(op.parameters)) {
      if (param.implementation === ImplementationLocation.Method && param.protocol.http!.in === 'uri') {
        text += `\thost = strings.ReplaceAll(host, "{${param.language.go!.serializedName}}", ${param.language.go!.name})\n`;
      }
    }
    hostParam = 'host';
  } else if (group.language.go!.hostParamName) {
    // simple parameterized host case or Azure ARM
    hostParam = 'client.' + group.language.go!.hostParamName;
  } else if (group.language.go!.host) {
    // swagger defines a host, use its const
    hostParam = '\thost';
  } else {
    throw new Error(`no host or endpoint defined for method ${group.language.go!.clientName}.${op.language.go!.name}`);
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
    hostParam = `runtime.JoinPaths(${hostParam}, urlPath)`;
  }
  if (hasPathParams) {
    // swagger defines path params, emit path and replace tokens
    imports.add('strings');
    // replace path parameters
    for (const pp of values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'path'; })) {
      // emit check to ensure path param isn't an empty string.  we only need
      // to do this for params that have an underlying type of string.
      const choiceIsString = function (schema: Schema): boolean {
        if (schema.type === SchemaType.Choice) {
          return (<ChoiceSchema>schema).choiceType.type === SchemaType.String;
        }
        if (schema.type === SchemaType.SealedChoice) {
          return (<ChoiceSchema>schema).choiceType.type === SchemaType.String;
        }
        return false;
      }
      const skipEncoding = skipURLEncoding(pp);
      if ((pp.schema.type === SchemaType.String || choiceIsString(pp.schema)) && !skipEncoding) {
        const paramName = getParamName(pp);
        imports.add('errors');
        text += `\tif ${paramName} == "" {\n`;
        text += `\t\treturn nil, errors.New("parameter ${paramName} cannot be empty")\n`;
        text += '\t}\n';
      }
      let paramValue = formatParamValue(pp, imports);
      if (!skipEncoding) {
        imports.add('net/url');
        paramValue = `url.PathEscape(${formatParamValue(pp, imports)})`;
      }
      text += `\turlPath = strings.ReplaceAll(urlPath, "{${pp.language.go!.serializedName}}", ${paramValue})\n`;
    }
  }
  text += `\treq, err := runtime.NewRequest(ctx, http.Method${capitalize(op.requests![0].protocol.http!.method)}, ${hostParam})\n`;
  text += '\tif err != nil {\n';
  text += '\t\treturn nil, err\n';
  text += '\t}\n';
  const hasQueryParams = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'query'; }).any();
  // helper to build nil checks for param groups
  const emitParamGroupCheck = function (gp: GroupProperty, param: Parameter): string {
    if (param.implementation === ImplementationLocation.Client) {
      return `\tif client.${param.language.go!.name} != nil {\n`;
    }
    const paramGroupName = uncapitalize(gp.language.go!.name);
    let optionalParamGroupCheck = `${paramGroupName} != nil && `;
    if (gp.required) {
      optionalParamGroupCheck = '';
    }
    return `\tif ${optionalParamGroupCheck}${paramGroupName}.${capitalize(param.language.go!.name)} != nil {\n`;
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
      if (qp.clientDefaultValue && qp.implementation === ImplementationLocation.Method) {
        qpText = emitClientSideDefault(qp, (name, val) => { return `\treqQP.Set(${name}, ${val})` }, imports);
      } else if (qp.required === true) {
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
      text += '\treqQP := req.Raw().URL.Query()\n';
      for (const qp of values(encodedParams)) {
        let setter: string;
        if (qp.protocol.http?.explode === true) {
          setter = `\tfor _, qv := range ${getParamName(qp)} {\n`;
          if (qp.schema.type !== SchemaType.Array) {
            throw new Error(`expected SchemaType.Array for query param ${qp.language.go!.name}`);
          }
          // emit a type conversion for the qv based on the array's element type
          let queryVal: string;
          const arrayQP = <ArraySchema>qp.schema;
          switch (arrayQP.elementType.type) {
            case SchemaType.Choice:
            case SchemaType.SealedChoice:
              const ch = <ChoiceSchema>arrayQP.elementType;
              // only string and number types are supported for enums
              if (ch.choiceType.type === SchemaType.String) {
                queryVal = 'string(qv)';
              } else {
                imports.add('fmt');
                queryVal = 'fmt.Sprintf("%d", qv)';
              }
              break;
            case SchemaType.String:
              queryVal = 'qv';
              break;
            default:
              imports.add('fmt');
              queryVal = 'fmt.Sprintf("%v", qv)';
          }
          setter += `\t\treqQP.Add("${qp.language.go!.serializedName}", ${queryVal})\n`;
          setter += '\t}';
        } else {
          // cannot initialize setter to this value as formatParamValue() can change imports
          setter = `reqQP.Set("${qp.language.go!.serializedName}", ${formatParamValue(qp, imports)})`;
        }
        text += emitQueryParam(qp, setter);
      }
      text += '\treq.Raw().URL.RawQuery = reqQP.Encode()\n';
    }
    // tack on any unencoded params to the end
    if (unencodedParams.length > 0) {
      if (encodedParams.length > 0) {
        text += '\tunencodedParams := []string{req.Raw().URL.RawQuery}\n';
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
        text += emitQueryParam(qp, setter);
      }
      imports.add('strings');
      text += '\treq.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")\n';
    }
  }
  if (hasBinaryResponse(op.responses!)) {
    // skip auto-body downloading for binary stream responses
    text += '\truntime.SkipBodyDownload(req)\n';
  }
  // add specific request headers
  const headerParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined; }).where((each: Parameter) => { return each.protocol.http!.in === 'header'; });
  headerParam.forEach(header => {
    const emitHeaderSet = function (headerParam: Parameter, prefix: string): string {
      if (headerParam.clientDefaultValue && headerParam.implementation === ImplementationLocation.Method) {
        return emitClientSideDefault(headerParam, (name, val) => {
          return `${prefix}req.Raw().Header[${name}] = []string{${val}}`;
        }, imports);
      } else if (header.schema.language.go!.headerCollectionPrefix) {
        let headerText = `${prefix}for k, v := range ${getParamName(headerParam)} {\n`;
        headerText += `${prefix}\tif v != nil {\n`;
        headerText += `${prefix}\t\treq.Raw().Header["${header.schema.language.go!.headerCollectionPrefix}"+k] = []string{*v}\n`;
        headerText += `${prefix}}\n`;
        headerText += `${prefix}}\n`;
        return headerText;
      } else {
        return `${prefix}req.Raw().Header["${headerParam.language.go!.serializedName}"] = []string{${formatParamValue(headerParam, imports)}}\n`;
      }
    }
    if (header.required || header.clientDefaultValue) {
      text += emitHeaderSet(header, '\t');
    } else {
      text += emitParamGroupCheck(<GroupProperty>header.language.go!.paramGroup, header);
      text += emitHeaderSet(header, '\t\t');
      text += `\t}\n`;
    }
  });
  const emitSetBodyWithErrCheck = function(setBodyParam: string): string {
    return `if err := ${setBodyParam}; err != nil {\n\treturn nil, err\n}\n`;
  }
  const mediaType = getMediaType(op.requests![0].protocol);
  if (mediaType === 'JSON' || mediaType === 'XML') {
    const bodyParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http!.in === 'body'; }).first();
    // default to the body param name
    let body = getParamName(bodyParam!);
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
        fieldName = capitalize(bodyParam!.language.go!.name);
        let tag = bodyParam!.schema.elementType.language.go!.name;
        if (bodyParam!.schema.elementType.serialization?.xml?.name) {
          tag = bodyParam!.schema.elementType.serialization.xml.name;
        }
        text += `\t\t${fieldName} *${bodyParam!.schema.language.go!.name} \`xml:"${tag}"\`\n`;
      }
      text += '\t}\n';
      let addr = '&';
      if (bodyParam && (!bodyParam.required && !isTypePassedByValue(bodyParam.schema))) {
        addr = '';
      }
      body = `wrapper{${fieldName}: ${addr}${body}}`;
    } else if (bodyParam!.schema.type === SchemaType.Date) {
      // wrap the body in the internal dateType
      body = `dateType(${body})`;
    } else if ((bodyParam!.schema.type === SchemaType.DateTime && (<DateTimeSchema>bodyParam!.schema).format === 'date-time-rfc1123') || bodyParam!.schema.type === SchemaType.UnixTime) {
      // wrap the body in the custom RFC1123 type
      text += `\taux := ${bodyParam!.schema.language.go!.internalTimeType}(${body})\n`;
      body = 'aux';
    } else if (isArrayOfTimesForMarshalling(bodyParam!.schema) || isArrayOfDatesForMarshalling(bodyParam!.schema)) {
      const timeType = (<ArraySchema>bodyParam!.schema).elementType.language.go!.internalTimeType;
      text += `\taux := make([]*${timeType}, len(${body}))\n`;
      text += `\tfor i := 0; i < len(${body}); i++ {\n`;
      text += `\t\taux[i] = (*${timeType})(${body}[i])\n`;
      text += '\t}\n';
      body = 'aux';
    } else if (isMapOfDateTime(bodyParam!.schema) || isMapOfDate(bodyParam!.schema)) {
      const timeType = (<ArraySchema>bodyParam!.schema).elementType.language.go!.internalTimeType;
      text += `\taux := map[string]*${timeType}{}\n`;
      text += `\tfor k, v := range ${body} {\n`;
      text += `\t\taux[k] = (*${timeType})(v)\n`;
      text += '\t}\n';
      body = 'aux';
    }
    let setBody = `runtime.MarshalAs${getMediaFormat(bodyParam!.schema, mediaType, `req, ${body}`)}`;
    if (bodyParam!.schema.language.go!.rawJSONAsBytes) {
      imports.add('bytes');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
      setBody = `req.SetBody(streaming.NopCloser(bytes.NewReader(${body})), "application/${mediaType.toLowerCase()}")`;
    }
    if (bodyParam!.required || bodyParam!.schema.type === SchemaType.Constant) {
      text += `\t${emitSetBodyWithErrCheck(setBody)}`;
      text += '\treturn req, nil\n';
    } else {
      text += emitParamGroupCheck(<GroupProperty>bodyParam!.language.go!.paramGroup, bodyParam!);
      text += `\t\t${emitSetBodyWithErrCheck(setBody)}`;
      text += '\t\treturn req, nil\n';
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
    if (bodyParam!.required) {
      text += `\t${emitSetBodyWithErrCheck(`req.SetBody(${bodyParam?.language.go!.name}, ${contentType})`)}`;
      text += '\treturn req, nil\n';
    } else {
      text += emitParamGroupCheck(<GroupProperty>bodyParam!.language.go!.paramGroup, bodyParam!);
      text += `\t\t${emitSetBodyWithErrCheck(`req.SetBody(${getParamName(bodyParam!)}, ${contentType})`)}`;
      text += '\t\treturn req, nil\n';
      text += '\t}\n';
      text += '\treturn req, nil\n';
    }
  } else if (mediaType === 'text') {
    imports.add('strings');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
    const bodyParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http!.in === 'body'; }).first();
    const contentType = `"${op.requests![0].protocol.http!.mediaTypes[0]}"`;
    if (bodyParam!.required) {
      text += `\tbody := streaming.NopCloser(strings.NewReader(${bodyParam!.language.go!.name}))\n`;
      text += `\t${emitSetBodyWithErrCheck(`req.SetBody(body, ${contentType})`)}\n`;
      text += '\treturn req, nil\n';
    } else {
      text += emitParamGroupCheck(<GroupProperty>bodyParam!.language.go!.paramGroup, bodyParam!);
      text += `\tbody := streaming.NopCloser(strings.NewReader(${getParamName(bodyParam!)}))\n`;
      text += `\t${emitSetBodyWithErrCheck(`req.SetBody(body, ${contentType})`)}`;
      text += '\treturn req, nil\n';
      text += '\t}\n';
      text += '\treturn req, nil\n';
    }
  } else if (mediaType === 'multipart') {
    text += '\tif err := runtime.SetMultipartFormData(req, map[string]any{\n';
    for (const param of values(aggregateParameters(op))) {
      if (param.isPartialBody) {
        text += `\t\t\t"${param.language.go!.name}": ${param.language.go!.name},\n`;
      }
    }
    text += '}); err != nil {'
    text += '\t\treturn nil, err\n';
    text += '\t}\n';
    text += '\treturn req, nil\n';
  } else if (mediaType === 'form') {
    const emitFormData = function (param: Parameter, setter: string): string {
      let formDataText = '';
      if (param.required === true) {
        formDataText = `\t${setter}\n`;
      } else {
        formDataText = emitParamGroupCheck(<GroupProperty>param.language.go!.paramGroup, param);
        formDataText += `\t\t${setter}\n`;
        formDataText += `\t}\n`;
      }
      return formDataText;
    }
    imports.add('net/url');
    imports.add('strings');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
    text += '\tformData := url.Values{}\n';
    for (const param of values(aggregateParameters(op))) {
      if (param.protocol.http?.in === 'body') {
        const setter = `formData.Set("${param.language.go!.serializedName}", ${formatParamValue(param, imports)})`;
        text += emitFormData(param, setter);
      }
    }
    text += `\tbody := streaming.NopCloser(strings.NewReader(formData.Encode()))\n`;
    text += `\t${emitSetBodyWithErrCheck('req.SetBody(body, "application/x-www-form-urlencoded")')}`;
    text += '\treturn req, nil\n';
  } else {
    text += `\treturn req, nil\n`;
  }
  text += '}\n\n';
  return text;
}

function emitClientSideDefault(param: Parameter, setterFormat: (name: string, val: string) => string, imports: ImportManager): string {
  const defaultVar = uncapitalize(param.language.go!.name) + 'Default';
  let text = `\t${defaultVar} := ${getClientDefaultValue(param)}\n`;
  text += `\tif options != nil && options.${capitalize(param.language.go!.name)} != nil {\n`;
  text += `\t\t${defaultVar} = *options.${capitalize(param.language.go!.name)}\n`;
  text += '}\n';
  text += setterFormat(`"${param.language.go!.serializedName}"`, formatValue(defaultVar, param.schema, imports)) + '\n';
  return text;
}

function getMediaFormat(schema: Schema, mediaType: 'JSON' | 'XML', param: string): string {
  let marshaller: 'JSON' | 'XML' | 'ByteArray' = mediaType;
  let format = '';
  if (schema.type === SchemaType.ByteArray) {
    marshaller = 'ByteArray';
    format = ', runtime.Base64StdFormat';
    if ((<ByteArraySchema>schema).format === 'base64url') {
      format = ', runtime.Base64URLFormat';
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

function isArrayOfDatesForMarshalling(schema: Schema): boolean {
  if (schema.type !== SchemaType.Array) {
    return false;
  }
  const arraySchema = <ArraySchema>schema;
  const arrayElem = <Schema>arraySchema.elementType;
  return arrayElem.type === SchemaType.Date;
}

function needsResponseHandler(op: Operation): boolean {
  return hasSchemaResponse(op) || responseHasHeaders(op) || (isLROOperation(op) && hasResultProperty(op) !== undefined) || isPageableOperation(op);
}

function generateResponseUnmarshaller(op: Operation, response: SchemaResponse, unmarshalTarget: string): string {
  let unmarshallerText = '';
  const zeroValue = getZeroReturnValue(op, 'handler');
  if (response.schema.type === SchemaType.DateTime || response.schema.type === SchemaType.UnixTime || response.schema.type === SchemaType.Date) {
    // use the designated time type for unmarshalling
    unmarshallerText += `\tvar aux *${response.schema.language.go!.internalTimeType}\n`;
    unmarshallerText += `\tif err := runtime.UnmarshalAs${getMediaType(response.protocol)}(resp, &aux); err != nil {\n`;
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += `\tresult.${getResultFieldName(op)} = (*time.Time)(aux)\n`;
    return unmarshallerText;
  } else if (isArrayOfDateTime(response.schema) || isArrayOfDate(response.schema)) {
    // unmarshalling arrays of date/time is a little more involved
    unmarshallerText += `\tvar aux []*${(<ArraySchema>response.schema).elementType.language.go!.internalTimeType}\n`;
    unmarshallerText += `\tif err := runtime.UnmarshalAs${getMediaType(response.protocol)}(resp, &aux); err != nil {\n`;
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += '\tcp := make([]*time.Time, len(aux))\n';
    unmarshallerText += '\tfor i := 0; i < len(aux); i++ {\n';
    unmarshallerText += '\t\tcp[i] = (*time.Time)(aux[i])\n';
    unmarshallerText += '\t}\n';
    unmarshallerText += `\tresult.${getResultFieldName(op)} = cp\n`;
    return unmarshallerText;
  } else if (isMapOfDateTime(response.schema) || isMapOfDate(response.schema)) {
    unmarshallerText += `\taux := map[string]*${(<DictionarySchema>response.schema).elementType.language.go!.internalTimeType}{}\n`;
    unmarshallerText += `\tif err := runtime.UnmarshalAs${getMediaType(response.protocol)}(resp, &aux); err != nil {\n`;
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += `\tcp := map[string]*time.Time{}\n`;
    unmarshallerText += `\tfor k, v := range aux {\n`;
    unmarshallerText += `\t\tcp[k] = (*time.Time)(v)\n`;
    unmarshallerText += `\t}\n`;
    unmarshallerText += `\tresult.${getResultFieldName(op)} = cp\n`;
    return unmarshallerText;
  }
  const mediaType = getMediaType(response.protocol);
  if (mediaType === 'JSON' || mediaType === 'XML') {
    if (response.schema.language.go!.rawJSONAsBytes) {
      unmarshallerText += `\tbody, err := runtime.Payload(resp)\n`;
      unmarshallerText += '\tif err != nil {\n';
      unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
      unmarshallerText += '\t}\n';
      unmarshallerText += `\t${unmarshalTarget} = body\n`;
    } else {
      unmarshallerText += `\tif err := runtime.UnmarshalAs${getMediaFormat(response.schema, mediaType, `resp, &${unmarshalTarget}`)}; err != nil {\n`;
      unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
      unmarshallerText += '\t}\n';
    }
  } else if (mediaType === 'text') {
    unmarshallerText += `\tbody, err := runtime.Payload(resp)\n`;
    unmarshallerText += '\tif err != nil {\n';
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += '\ttxt := string(body)\n';
    unmarshallerText += `\t${unmarshalTarget} = &txt\n`;
  } else {
    // the remaining media types are handled elsewhere
    throw new Error(`unhandled media type ${mediaType} for operation ${op.language.go!.clientName}.${op.language.go!.name}`);
  }
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
  text += `func (client *${clientName}) ${name}(resp *http.Response) (${generateReturnsInfo(op, 'handler').join(', ')}) {\n`;
  const addHeaders = function (props?: Property[]) {
    const headerVals = new Array<Property>();
    for (const prop of values(props)) {
      if (prop.language.go!.fromHeader) {
        headerVals.push(prop);
      }
    }
    for (const headerVal of values(headerVals)) {
      text += formatHeaderResponseValue(headerVal.language.go!.name, headerVal.language.go!.fromHeader, headerVal.schema, imports, 'result', `${getResponseEnvelopeName(op)}{}`);
    }
  }
  if (!isMultiRespOperation(op)) {
    const respEnvName = getResponseEnvelopeName(op);
    text += `\tresult := ${respEnvName}{`;
    if (isBinaryResponseOperation(op)) {
      text += 'Body: resp.Body';
    }
    text += '}\n';
    // we know there's a result envelope at this point
    const respEnv = getResponseEnvelope(op);
    addHeaders(respEnv.properties);
    const schemaResponse = getSchemaResponse(op);
    if (op.language.go!.headAsBoolean === true) {
      text += '\tresult.Success = resp.StatusCode >= 200 && resp.StatusCode < 300\n';
    } else if (schemaResponse) {
      // when unmarshalling a wrapped XML array or discriminated type, unmarshal into the response envelope
      let target = `result.${getResultFieldName(op)}`
      if ((getMediaType(schemaResponse.protocol) === 'XML' && schemaResponse.schema.type === SchemaType.Array) || schemaResponse.schema.language.go!.discriminatorInterface) {
        target = 'result';
      }
      text += generateResponseUnmarshaller(op, schemaResponse, target);
    }
    text += '\treturn result, nil\n';
  } else {
    imports.add('fmt');
    text += `\tresult := ${getResponseEnvelopeName(op)}{}\n`;
    // unmarshal any header values
    const respEnv = getResponseEnvelope(op);
    addHeaders(respEnv.properties);
    text += '\tswitch resp.StatusCode {\n';
    for (const response of values(op.responses)) {
      text += `\tcase ${formatStatusCodes(response.protocol.http!.statusCodes)}:\n`
      if (!isSchemaResponse(response)) {
        // the operation contains a mix of schemas and non-schema responses
        continue;
      }
      text += `\tvar val ${response.schema.language.go!.name}\n`;
      text += generateResponseUnmarshaller(op, response, 'val');
      text += `\tresult.Value = val\n`;
    }
    text += '\tdefault:\n';
    text += `\t\treturn ${getZeroReturnValue(op, 'handler')}, fmt.Errorf("unhandled HTTP status code %d", resp.StatusCode)\n`;
    text += '\t}\n';
    text += '\treturn result, nil\n';
  }
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

function isArrayOfDate(schema: Schema): boolean {
  if (schema.type !== SchemaType.Array) {
    return false;
  }
  const arraySchema = <ArraySchema>schema;
  const arrayElem = <Schema>arraySchema.elementType;
  return arrayElem.type === SchemaType.Date;
}

function isMapOfDateTime(schema: Schema): boolean {
  if (schema.type !== SchemaType.Dictionary) {
    return false;
  }
  const dictSchema = <DictionarySchema>schema;
  const dictElem = <Schema>dictSchema.elementType;
  return dictElem.type === SchemaType.DateTime || dictElem.type === SchemaType.UnixTime;
}

function isMapOfDate(schema: Schema): boolean {
  if (schema.type !== SchemaType.Dictionary) {
    return false;
  }
  const dictSchema = <DictionarySchema>schema;
  const dictElem = <Schema>dictSchema.elementType;
  return dictElem.type === SchemaType.Date;
}

// returns the media type used by the protocol
function getMediaType(protocol: Protocols): 'JSON' | 'XML' | 'binary' | 'text' | 'form' | 'multipart' | 'none' {
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
    case KnownMediaType.Form:
      return 'form';
    case KnownMediaType.Multipart:
      return 'multipart';
    default:
      return 'none';
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
    params.push(`${uncapitalize(methodParam.language.go!.name)} ${formatParameterTypeName(methodParam)}`);
  }
  return params.join(', ');
}

// returns the return signature where each entry is the type name
// e.g. [ '*string', 'error' ]
// apiType describes where the return sig is used.
//   api - for the API definition
//    op - for the operation
// handler - for the response handler
function generateReturnsInfo(op: Operation, apiType: 'api' | 'op' | 'handler'): string[] {
  let returnType = getResponseEnvelopeName(op);
  if (isLROOperation(op)) {
    switch (apiType) {
      case 'api':
        if (isPageableOperation(op)) {
          returnType = `*runtime.Poller[*runtime.Pager[${getResponseEnvelopeName(op)}]]`;
        } else {
          returnType = `*runtime.Poller[${getResponseEnvelopeName(op)}]`;
        }
        break;
      case 'handler':
        // we only have a handler for operations that return a schema
        if (isPageableOperation(op)) {
          // we need to consult the final response type name
          returnType = getResponseEnvelopeName(op);
        } else {
          throw new Error(`handler being generated for non-pageable LRO ${op.language.go!.name} which is unexpected`);
        }
        break;
      case 'op':
        returnType = '*http.Response';
        break;
    }
  } else if (isPageableOperation(op)) {
    switch (apiType) {
      case 'api':
      case 'op':
        // pager operations don't return an error
        return [`*runtime.Pager[${returnType}]`];
    }
  }
  return [returnType, 'error'];
}

function generateLROBeginMethod(op: Operation, injectSpans: boolean, imports: ImportManager): string {
  const info = <OperationNaming>op.language.go!;
  const params = getAPIParametersSig(op, imports);
  const returns = generateReturnsInfo(op, 'api');
  const clientName = op.language.go!.clientName;
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
  let text = '';
  if (hasDescription(op.language.go!)) {
    text += `${comment(`Begin${op.language.go!.name} - ${op.language.go!.description}`, "//", undefined, commentLength)}\n`;
    text += genApiVersionDoc(op.apiVersions);
  }
  const zeroResp = getZeroReturnValue(op, 'api');
  const methodParams = getMethodParameters(op);
  for (const param of values(methodParams)) {
    if (param.language.go!.description) {
      text += `${formatCommentAsBulletItem(`${param.language.go!.name} - ${param.language.go!.description}`)}\n`;
    }
  }
  text += `func (client *${clientName}) Begin${op.language.go!.name}(${params}) (${returns.join(', ')}) {\n`;
  let pollerType = 'nil';
  let pollerTypeParam = `[${getResponseEnvelopeName(op)}]`;
  if (isPageableOperation(op)) {
    // for paged LROs, we construct a pager and pass it to the LRO ctor.
    pollerTypeParam = `[*runtime.Pager${pollerTypeParam}]`;
    pollerType = '&pager';
    text += '\tpager := ';
    text += emitPagerDefinition(op, injectSpans, imports);
  }

  text += '\tif options == nil || options.ResumeToken == "" {\n';
  text += '\t\tvar err error\n';
  if (injectSpans) {
    text += `\t\tctx, endSpan := runtime.StartSpan(ctx, "${clientName}.Begin${op.language.go!.name}", client.internal.Tracer(), nil)\n`;
    text += '\t\tdefer func() { endSpan(err) }()\n';
  }
  // creating the poller from response branch

  let opName = op.language.go!.name;
  opName = info.protocolNaming.internalMethod;
  text += `\t\tresp, err := client.${opName}(${getCreateRequestParameters(op)})\n`;
  text += `\t\tif err != nil {\n`;
  text += `\t\t\treturn ${zeroResp}, err\n`;
  text += `\t\t}\n`;

  let finalStateVia = '';
  // LRO operation might have a special configuration set in x-ms-long-running-operation-options
  // which indicates a specific url to perform the final Get operation on
  if (op.extensions?.['x-ms-long-running-operation-options']?.['final-state-via']) {
    finalStateVia = op.extensions?.['x-ms-long-running-operation-options']?.['final-state-via'];
    switch (finalStateVia) {
      case "azure-async-operation":
        finalStateVia = `runtime.FinalStateViaAzureAsyncOp`;
        break;
      case "location":
        finalStateVia = `runtime.FinalStateViaLocation`;
        break;
      case "original-uri":
        finalStateVia = `runtime.FinalStateViaOriginalURI`;
        break;
      case "operation-location":
        finalStateVia = `runtime.FinalStateViaOpLocation`;
        break;
      default:
        throw new Error(`unhandled final-state-via value ${finalStateVia}`);
    }
  }

  text += `\t\tpoller, err := runtime.NewPoller`;
  if (finalStateVia === '' && pollerType === 'nil') {
    // the generic type param is redundant when it's also specified in the
    // options struct so we only include it when there's no options.
    text += pollerTypeParam;
  }
  text += '(resp, client.internal.Pipeline(), ';
  if (finalStateVia === '' && pollerType === 'nil') {
    // no options
    text += 'nil)\n';
  } else {
    // at least one option
    text += `&runtime.NewPollerOptions${pollerTypeParam}{\n`;
    if (finalStateVia !== '') {
      text += `\t\t\tFinalStateVia: ${finalStateVia},\n`;  
    }
    if (pollerType !== 'nil') {
      text += `\t\t\tResponse: ${pollerType},\n`;
    }
    text += '\t\t})\n';
  }
  text += '\t\treturn poller, err\n';
  text += '\t} else {\n';

  // creating the poller from resume token branch

  text += `\t\treturn runtime.NewPollerFromResumeToken`;
  if (pollerType === 'nil') {
    text += pollerTypeParam;
  }
  text += '(options.ResumeToken, client.internal.Pipeline(), ';
  if (pollerType === 'nil') {
    text += 'nil)\n';
  } else {
    text += `&runtime.NewPollerFromResumeTokenOptions${pollerTypeParam}{\n`;
    text += `\t\t\tResponse: ${pollerType},\n`;
    text  += '\t\t})\n';
  }
  text += '\t}\n';

  text += '}\n\n';
  return text;
}
