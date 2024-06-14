/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../codemodel.go/src/index.js';
import { capitalize, comment, uncapitalize } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import * as helpers from './helpers.js';
import { ImportManager } from './imports.js';

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
export async function generateOperations(codeModel: go.CodeModel): Promise<Array<OperationGroupContent>> {
  // generate protocol operations
  const operations = new Array<OperationGroupContent>();
  if (codeModel.clients.length === 0) {
    return operations;
  }
  const azureARM = codeModel.type === 'azure-arm';
  for (const client of codeModel.clients) {
    // the list of packages to import
    const imports = new ImportManager();
    if (client.methods.length > 0) {
      // add standard imports for clients with methods.
      // clients that are purely hierarchical (i.e. having no APIs) won't need them.
      imports.add('net/http');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/policy');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
    }

    let clientPkg = 'azcore';
    if (azureARM) {
      clientPkg = 'arm';
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/arm');
      client.constructors.push(createARMClientConstructor(client, imports))
    } else {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    }

    // generate client type

    let clientText = '';
    clientText += `${comment(`${client.description}`, '//', undefined, helpers.commentLength)}\n`;
    clientText += '// Don\'t use this type directly, use ';
    if (client.constructors.length === 1) {
      clientText += `${client.constructors[0].name}() instead.\n`;
    } else if (client.parent) {
      // find the accessor method
      let accessorMethod: string | undefined;
      for (const clientAccessor of client.parent.clientAccessors) {
        if (clientAccessor.subClient === client) {
          accessorMethod = clientAccessor.name;
          break;
        }
      }
      if (!accessorMethod) {
        throw new Error(`didn't find accessor method for client ${client.name} on parent client ${client.parent.name}`);
      }
      clientText += `[${client.parent.name}.${accessorMethod}] instead.\n`;
    } else {
      clientText += 'a constructor function instead.\n';
    }
    clientText += `type ${client.name} struct {\n`;
    clientText += `\tinternal *${clientPkg}.Client\n`;

    // check for any optional host params
    const optionalParams = new Array<go.Parameter>();

    const isParamPointer = function(param: go.Parameter): boolean {
      // for client params, only optional and flag types are passed by pointer
      return param.kind === 'flag' || param.kind === 'optional';
    };

    // now emit any client params (non parameterized host params case)
    if (client.parameters.length > 0) {
      const addedGroups = new Set<string>();
      for (const clientParam of values(client.parameters)) {
        if (go.isLiteralParameter(clientParam)) {
          continue;
        }
        if (clientParam.group) {
          if (!addedGroups.has(clientParam.group.groupName)) {
            clientText += `\t${uncapitalize(clientParam.group.groupName)} ${!isParamPointer(clientParam) ? '' : '*'}${clientParam.group.groupName}\n`;
            addedGroups.add(clientParam.group.groupName);
          }
          continue;
        }
        clientText += `\t${clientParam.name} `;
        if (!isParamPointer(clientParam)) {
          clientText += `${go.getTypeDeclaration(clientParam.type)}\n`;
        } else {
          clientText += `${helpers.formatParameterTypeName(clientParam)}\n`;
        }
        if (!go.isRequiredParameter(clientParam)) {
          optionalParams.push(clientParam);
        }
      }
    }

    // end of client definition
    clientText += '}\n\n';

    if (azureARM && optionalParams.length > 0) {
      throw new Error('optional client parameters for ARM is not supported');
    }

    // generate client constructors
    clientText += generateConstructors(azureARM, client, imports);

    // generate client accessors and operations
    let opText = '';
    for (const clientAccessor of client.clientAccessors) {
      opText += `// ${clientAccessor.name} creates a new instance of [${clientAccessor.subClient.name}].\n`;
      opText += `func (client *${client.name}) ${clientAccessor.name}() *${clientAccessor.subClient.name} {\n`;
      opText += `\treturn &${clientAccessor.subClient.name}{\n`;
      opText += '\t\tinternal: client.internal,\n';
      // propagate all client params
      for (const param of client.parameters) {
        opText += `\t\t${param.name}: client.${param.name},\n`;
      }
      opText += '\t}\n}\n\n';
    }

    const nextPageMethods = new Array<go.NextPageMethod>();
    for (const method of client.methods) {
      // protocol creation can add imports to the list so
      // it must be done before the imports are written out
      if (go.isLROMethod(method)) {
        // generate Begin method
        opText += generateLROBeginMethod(client, method, imports, codeModel.options.injectSpans, codeModel.options.generateFakes);
      }
      opText += generateOperation(client, method, imports, codeModel.options.injectSpans, codeModel.options.generateFakes);
      opText += createProtocolRequest(azureARM, client, method, imports);
      if (!go.isLROMethod(method) || go.isPageableMethod(method)) {
        // LRO responses are handled elsewhere, with the exception of pageable LROs
        opText += createProtocolResponse(client, method, imports);
      }
      if (go.isPageableMethod(method) && method.nextPageMethod && !nextPageMethods.includes(method.nextPageMethod)) {
        // track the next page methods to generate as multiple operations can use the same next page operation
        nextPageMethods.push(method.nextPageMethod);
      }
    }

    for (const method of nextPageMethods) {
      opText += createProtocolRequest(azureARM, client, method, imports);
    }

    // stitch it all together
    let text = helpers.contentPreamble(codeModel);
    text += imports.text();
    text += clientText;
    text += opText;
    operations.push(new OperationGroupContent(client.name, text));
  }
  return operations;
}

// generates all modeled client constructors
function generateConstructors(azureARM: boolean, client: go.Client, imports: ImportManager): string {
  if (client.constructors.length === 0) {
    return '';
  }

  let ctorText = '';
  for (const constructor of client.constructors) {
    const ctorParams = new Array<string>();
    const paramDocs = new Array<string>();

    constructor.parameters.sort(helpers.sortParametersByRequired);
    for (const ctorParam of constructor.parameters) {
      imports.addImportForType(ctorParam.type);
      ctorParams.push(`${ctorParam.name} ${helpers.formatParameterTypeName(ctorParam)}`);
      if (ctorParam.description) {
        paramDocs.push(helpers.formatCommentAsBulletItem(`${ctorParam.name} - ${ctorParam.description}`));
      }
    }

    // add client options last
    ctorParams.push(`${client.options.name} ${helpers.formatParameterTypeName(client.options)}`);
    paramDocs.push(helpers.formatCommentAsBulletItem(`${client.options.name} - ${client.options.description}`));

    ctorText += `// ${constructor.name} creates a new instance of ${client.name} with the specified values.\n`;
    for (const doc of paramDocs) {
      ctorText += `${doc}\n`;
    }

    ctorText += `func ${constructor.name}(${ctorParams.join(', ')}) (*${client.name}, error) {\n`;
    let clientType = 'azcore';
    if (azureARM) {
      clientType = 'arm';
    }

    ctorText += `\tcl, err := ${clientType}.NewClient(moduleName, moduleVersion, credential, options)\n`;
    ctorText += '\tif err != nil {\n';
    ctorText += '\t\treturn nil, err\n';
    ctorText += '\t}\n';

    // construct client literal
    ctorText += `\tclient := &${client.name}{\n`;
    for (const parameter of values(client.parameters)) {
      // each client field will have a matching parameter with the same name
      ctorText += `\t\t${parameter.name}: ${parameter.name},\n`;
    }
    ctorText += '\tinternal: cl,\n';
    ctorText += '\t}\n';
    ctorText += '\treturn client, nil\n';
    ctorText += '}\n\n';
  }

  return ctorText;
}

// creates a modeled constructor for an ARM client
function createARMClientConstructor(client: go.Client, imports: ImportManager): go.Constructor {
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/arm');
  const ctor = new go.Constructor(`New${client.name}`);
  // add any modeled parameter first, which should only be the subscriptionID, then add TokenCredential
  for (const param of client.parameters) {
    ctor.parameters.push(param);
  }
  const tokenCredParam = new go.Parameter('credential', new go.QualifiedType('TokenCredential', 'github.com/Azure/azure-sdk-for-go/sdk/azcore'), 'required', true, 'client');
  tokenCredParam.description = 'used to authorize requests. Usually a credential from azidentity.';
  ctor.parameters.push(tokenCredParam);
  return ctor;
}

// use this to generate the code that will help process values returned in response headers
function formatHeaderResponseValue(headerResp: go.HeaderResponse | go.HeaderMapResponse, imports: ImportManager, respObj: string, zeroResp: string): string {
  // dictionaries are handled slightly different so we do that first
  if (go.isHeaderMapResponse(headerResp)) {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
    imports.add('strings');
    const headerPrefix = headerResp.collectionPrefix;
    let text = '\tfor hh := range resp.Header {\n';
    text += `\t\tif len(hh) > len("${headerPrefix}") && strings.EqualFold(hh[:len("${headerPrefix}")], "${headerPrefix}") {\n`;
    text += `\t\t\tif ${respObj}.${headerResp.fieldName} == nil {\n`;
    text += `\t\t\t\t${respObj}.${headerResp.fieldName} = map[string]*string{}\n`;
    text += '\t\t\t}\n';
    text += `\t\t\t${respObj}.${headerResp.fieldName}[hh[len("${headerPrefix}"):]] = to.Ptr(resp.Header.Get(hh))\n`;
    text += '\t\t}\n';
    text += '\t}\n';
    return text;
  }
  let text = `\tif val := resp.Header.Get("${headerResp.headerName}"); val != "" {\n`;
  let name = uncapitalize(headerResp.fieldName);
  let byRef = '&';
  if (go.isConstantType(headerResp.type)) {
    text += `\t\t${respObj}.${headerResp.fieldName} = (*${headerResp.type.name})(&val)\n`;
    text += '\t}\n';
    return text;
  } else if (go.isPrimitiveType(headerResp.type)) {
    if (headerResp.type.typeName === 'bool') {
      imports.add('strconv');
      text += `\t\t${name}, err := strconv.ParseBool(val)\n`;
    } else if (headerResp.type.typeName === 'int32' || headerResp.type.typeName === 'int64') {
      imports.add('strconv');
      if (headerResp.type.typeName === 'int32') {
        text += `\t\t${name}32, err := strconv.ParseInt(val, 10, 32)\n`;
        text += `\t\t${name} := int32(${name}32)\n`;
      } else {
        text += `\t\t${name}, err := strconv.ParseInt(val, 10, 64)\n`;
      }
    } else if (headerResp.type.typeName === 'float32' || headerResp.type.typeName === 'float64') {
      imports.add('strconv');
      if (headerResp.type.typeName === 'float32') {
        text += `\t\t${name}32, err := strconv.ParseFloat(val, 32)\n`;
        text += `\t\t${name} := float32(${name}32)\n`;
      } else {
        text += `\t\t${name}, err := strconv.ParseFloat(val, 64)\n`;
      }
    } else if (headerResp.type.typeName === 'string') {
      text += `\t\t${respObj}.${headerResp.fieldName} = &val\n`;
      text += '\t}\n';
      return text;
    } else {
      throw new Error(`unhandled primitive type ${headerResp.type.typeName}`);
    }
  } else if (go.isTimeType(headerResp.type)) {
    imports.add('time');
    if (headerResp.type.dateTimeFormat === 'dateType') {
      text += `\t\t${name}, err := time.Parse("${helpers.dateFormat}", val)\n`;
    } else if (headerResp.type.dateTimeFormat === 'timeRFC3339') {
      text += `\t\t${name}, err := time.Parse("${helpers.timeRFC3339Format}", val)\n`;
    } else if (headerResp.type.dateTimeFormat === 'timeUnix') {
      imports.add('strconv');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
      text += '\t\tsec, err := strconv.ParseInt(val, 10, 64)\n';
      name = 'to.Ptr(time.Unix(sec, 0))';
      byRef = '';
    } else {
      let format = helpers.datetimeRFC3339Format;
      if (headerResp.type.dateTimeFormat === 'dateTimeRFC1123') {
        format = helpers.datetimeRFC1123Format;
      }
      text += `\t\t${name}, err := time.Parse(${format}, val)\n`;
    }
  } else if (go.isBytesType(headerResp.type)) {
    // ByteArray is a base-64 encoded value in string format
    imports.add('encoding/base64');
    text += `\t\t${name}, err := base64.${helpers.formatBytesEncoding(headerResp.type.encoding)}Encoding.DecodeString(val)\n`;
    byRef = '';
  } else if (go.isLiteralValue(headerResp.type)) {
    text += `\t\t${respObj}.${headerResp.fieldName} = &val\n`;
    text += '\t}\n';
    return text;
  } else {
    throw new Error(`unsupported header type ${go.getTypeDeclaration(headerResp.type)}`);
  }
  text += '\t\tif err != nil {\n';
  text += `\t\t\treturn ${zeroResp}, err\n`;
  text += '\t\t}\n';
  text += `\t\t${respObj}.${headerResp.fieldName} = ${byRef}${name}\n`;
  text += '\t}\n';
  return text;
}

function getZeroReturnValue(method: go.Method, apiType: 'api' | 'op' | 'handler'): string {
  let returnType = `${method.responseEnvelope.name}{}`;
  if (go.isLROMethod(method)) {
    if (apiType === 'api' || apiType === 'op') {
      // the api returns a *Poller[T]
      // the operation returns an *http.Response
      returnType = 'nil';
    }
  }
  return returnType;
}

function emitPagerDefinition(client: go.Client, method: go.PageableMethod, imports: ImportManager, injectSpans: boolean, generateFakes: boolean): string {
  imports.add('context');
  let text = `runtime.NewPager(runtime.PagingHandler[${method.responseEnvelope.name}]{\n`;
  text += `\t\tMore: func(page ${method.responseEnvelope.name}) bool {\n`;
  // there is no advancer for single-page pagers
  if (method.nextLinkName) {
    text += `\t\t\treturn page.${method.nextLinkName} != nil && len(*page.${method.nextLinkName}) > 0\n`;
    text += '\t\t},\n';
  } else {
    text += '\t\t\treturn false\n';
    text += '\t\t},\n';
  }
  text += `\t\tFetcher: func(ctx context.Context, page *${method.responseEnvelope.name}) (${method.responseEnvelope.name}, error) {\n`;
  const reqParams = helpers.getCreateRequestParameters(method);
  if (generateFakes) {
    text += `\t\tctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "${client.name}.${fixUpMethodName(method)}")\n`;
  }
  if (method.nextLinkName) {
    let nextLinkVar: string;
    if (!go.isLROMethod(method)) {
      text += '\t\t\tnextLink := ""\n';
      nextLinkVar = 'nextLink';
      text += '\t\t\tif page != nil {\n';
      text += `\t\t\t\tnextLink = *page.${method.nextLinkName}\n\t\t\t}\n`;
    } else {
      nextLinkVar = `*page.${method.nextLinkName}`;
    }
    text += `\t\t\tresp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), ${nextLinkVar}, func(ctx context.Context) (*policy.Request, error) {\n`;
    text += `\t\t\t\treturn client.${method.naming.requestMethod}(${reqParams})\n\t\t\t}, `;
    // nextPageMethod might be absent in some cases, see https://github.com/Azure/autorest/issues/4393
    if (method.nextPageMethod) {
      const nextOpParams = helpers.getCreateRequestParametersSig(method.nextPageMethod).split(',');
      // keep the parameter names from the name/type tuples and find nextLink param
      for (let i = 0; i < nextOpParams.length; ++i) {
        const paramName = nextOpParams[i].trim().split(' ')[0];
        const paramType = nextOpParams[i].trim().split(' ')[1];
        if (paramName.startsWith('next') && paramType === 'string') {
          nextOpParams[i] = 'encodedNextLink';
        } else {
          nextOpParams[i] = paramName;
        }
      }
      // add a definition for the nextReq func that uses the nextLinkOperation
      text += '&runtime.FetcherForNextLinkOptions{\n\t\t\t\tNextReq: func(ctx context.Context, encodedNextLink string) (*policy.Request, error) {\n';
      text += `\t\t\t\t\treturn client.${method.nextPageMethod.name}(${nextOpParams.join(', ')})\n\t\t\t\t},\n\t\t\t})\n`;
    } else {
      text += 'nil)\n';
    }
    text += `\t\t\tif err != nil {\n\t\t\t\treturn ${method.responseEnvelope.name}{}, err\n\t\t\t}\n`;
    text += `\t\t\treturn client.${method.naming.responseMethod}(resp)\n`;
    text += '\t\t\t},\n';
  } else {
    // this is the singular page case, no fetcher helper required
    text += `\t\t\treq, err := client.${method.naming.requestMethod}(${reqParams})\n`;
    text += '\t\t\tif err != nil {\n';
    text += `\t\t\t\treturn ${method.responseEnvelope.name}{}, err\n`;
    text += '\t\t\t}\n';
    text += '\t\t\tresp, err := client.internal.Pipeline().Do(req)\n';
    text += '\t\t\tif err != nil {\n';
    text += `\t\t\t\treturn ${method.responseEnvelope.name}{}, err\n`;
    text += '\t\t\t}\n';
    text += '\t\t\tif !runtime.HasStatusCode(resp, http.StatusOK) {\n';
    text += `\t\t\t\treturn ${method.responseEnvelope.name}{}, runtime.NewResponseError(resp)\n`;
    text += '\t\t\t}\n';
    text += `\t\t\treturn client.${method.naming.responseMethod}(resp)\n`;
    text += '\t\t},\n';
  }
  if (injectSpans) {
    text += '\t\tTracer: client.internal.Tracer(),\n';
  }
  text += '\t})\n';
  return text;
}

function genApiVersionDoc(apiVersions: Array<string>): string {
  if (apiVersions.length === 0) {
    return '';
  }
  return `//\n// Generated from API version ${apiVersions.join(', ')}\n`;
}

function genRespErrorDoc(method: go.Method): string {
  if (!(method.responseEnvelope.result && go.isHeadAsBooleanResult(method.responseEnvelope.result)) && !go.isPageableMethod(method)) {
    // when head-as-boolean is enabled, no error is returned for 4xx status codes.
    // pager constructors don't return an error
    return '// If the operation fails it returns an *azcore.ResponseError type.\n';
  }
  return '';
}

function generateOperation(client: go.Client, method: go.Method, imports: ImportManager, injectSpans: boolean, generateFakes: boolean): string {
  const params = getAPIParametersSig(method, imports);
  const returns = generateReturnsInfo(method, 'op');
  let methodName = method.name;
  if(go.isPageableMethod(method) && !go.isLROMethod(method)) {
    methodName = fixUpMethodName(method);
  }
  let text = '';
  const respErrDoc = genRespErrorDoc(method);
  const apiVerDoc = genApiVersionDoc(method.apiVersions);
  if (method.description) {
    text += `${comment(`${methodName} - ${method.description}`, '//', undefined, helpers.commentLength)}\n`;
  } else if (respErrDoc.length > 0 || apiVerDoc.length > 0) {
    // if the method has no doc comment but we're adding other
    // doc comments, add an empty method name comment. this preserves
    // existing behavior and makes the docs look better overall.
    text += `// ${methodName} -\n`;
  }
  text += respErrDoc;
  text += apiVerDoc;
  if (go.isLROMethod(method)) {
    methodName = method.naming.internalMethod;
  } else {
    for (const param of values(helpers.getMethodParameters(method))) {
      if (param.description) {
        text += `${helpers.formatCommentAsBulletItem(`${param.name} - ${param.description}`)}\n`;
      }
    }
  }
  text += `func (client *${client.name}) ${methodName}(${params}) (${returns.join(', ')}) {\n`;
  const reqParams = helpers.getCreateRequestParameters(method);
  if (go.isPageableMethod(method) && !go.isLROMethod(method)) {
    text += '\treturn ';
    text += emitPagerDefinition(client, method, imports, injectSpans, generateFakes);
    text += '}\n\n';
    return text;
  }
  text += '\tvar err error\n';
  let operationName = `"${client.name}.${fixUpMethodName(method)}"`;
  if (generateFakes && injectSpans) {
    text += `\tconst operationName = ${operationName}\n`;
    operationName = 'operationName';
  }
  if (generateFakes) {
    text += `\tctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, ${operationName})\n`;
  }
  if (injectSpans) {
    text += `\tctx, endSpan := runtime.StartSpan(ctx, ${operationName}, client.internal.Tracer(), nil)\n`;
    text += '\tdefer func() { endSpan(err) }()\n';
  }
  const zeroResp = getZeroReturnValue(method, 'op');
  text += `\treq, err := client.${method.naming.requestMethod}(${reqParams})\n`;
  text += '\tif err != nil {\n';
  text += `\t\treturn ${zeroResp}, err\n`;
  text += '\t}\n';
  text += '\thttpResp, err := client.internal.Pipeline().Do(req)\n';
  text += '\tif err != nil {\n';
  text += `\t\treturn ${zeroResp}, err\n`;
  text += '\t}\n';
  text += `\tif !runtime.HasStatusCode(httpResp, ${helpers.formatStatusCodes(method.httpStatusCodes)}) {\n`;
  text += '\t\terr = runtime.NewResponseError(httpResp)\n';
  text += `\t\treturn ${zeroResp}, err\n`;
  text += '\t}\n';
  // HAB with headers response is handled in protocol responder
  if (method.responseEnvelope.result && go.isHeadAsBooleanResult(method.responseEnvelope.result) && method.responseEnvelope.headers.length === 0) {
    text += `\treturn ${method.responseEnvelope.name}{${method.responseEnvelope.result.fieldName}: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil\n`;
  } else {
    if (go.isLROMethod(method)) {
      text += '\treturn httpResp, nil\n';
    } else if (needsResponseHandler(method)) {
      // also cheating here as at present the only param to the responder is an http.Response
      text += `\tresp, err := client.${method.naming.responseMethod}(httpResp)\n`;
      text += '\treturn resp, err\n';
    } else if (method.responseEnvelope.result && go.isBinaryResult(method.responseEnvelope.result)) {
      text += `\treturn ${method.responseEnvelope.name}{${method.responseEnvelope.result.fieldName}: httpResp.Body}, nil\n`;
    } else {
      text += `\treturn ${method.responseEnvelope.name}{}, nil\n`;
    }
  }
  text += '}\n\n';
  return text;
}

function createProtocolRequest(azureARM: boolean, client: go.Client, method: go.Method | go.NextPageMethod, imports: ImportManager): string {
  let name = method.name;
  if (go.isMethod(method)) {
    name = method.naming.requestMethod;
  }

  for (const param of values(method.parameters)) {
    if (param.location !== 'method' || !go.isRequiredParameter(param)) {
      continue;
    }
    imports.addImportForType(param.type);
  }

  const returns = ['*policy.Request', 'error'];
  let text = `${comment(name, '// ')} creates the ${method.name} request.\n`;
  text += `func (client *${client.name}) ${name}(${helpers.getCreateRequestParametersSig(method)}) (${returns.join(', ')}) {\n`;

  const hostParams = new Array<go.URIParameter>();
  for (const parameter of client.parameters) {
    if (go.isURIParameter(parameter)) {
      hostParams.push(parameter);
    }
  }

  let hostParam: string;
  if (azureARM) {
    hostParam = 'client.internal.Endpoint()';
  } else if (client.templatedHost) {
    imports.add('strings');
    // we have a templated host
    text += `\thost := "${client.host!}"\n`;
    // get all the host params on the client
    for (const hostParam of hostParams) {
      text += `\thost = strings.ReplaceAll(host, "{${hostParam.uriPathSegment}}", ${helpers.formatValue(`client.${(<string>hostParam.name)}`, hostParam.type, imports)})\n`;
    }
    // check for any method local host params
    for (const param of values(method.parameters)) {
      if (param.location === 'method' && go.isURIParameter(param)) {
        text += `\thost = strings.ReplaceAll(host, "{${param.uriPathSegment}}", ${helpers.formatValue(helpers.getParamName(param), param.type, imports)})\n`;
      }
    }
    hostParam = 'host';
  } else if (hostParams.length === 1) {
    // simple parameterized host case
    hostParam = 'client.' + hostParams[0].name;
  } else if (client.host) {
    // swagger defines a host, use its const
    hostParam = '\thost';
  } else {
    throw new Error(`no host or endpoint defined for method ${client.name}.${method.name}`);
  }
  const hasPathParams = values(method.parameters).where((each: go.Parameter) => { return go.isPathParameter(each); }).any();
  // storage needs the client.u to be the source-of-truth for the full path.
  // however, swagger requires that all operations specify a path, which is at odds with storage.
  // to work around this, storage specifies x-ms-path paths with path params but doesn't
  // actually reference the path params (i.e. no params with which to replace the tokens).
  // so, if a path contains tokens but there are no path params, skip emitting the path.
  const pathStr = method.httpPath;
  const pathContainsParms = pathStr.includes('{');
  if (hasPathParams || (!pathContainsParms && pathStr.length > 1)) {
    // there are path params, or the path doesn't contain tokens and is not "/" so emit it
    text += `\turlPath := "${method.httpPath}"\n`;
    hostParam = `runtime.JoinPaths(${hostParam}, urlPath)`;
  }
  if (hasPathParams) {
    // swagger defines path params, emit path and replace tokens
    imports.add('strings');
    // replace path parameters
    for (const pp of values(method.parameters)) {
      if (!go.isPathParameter(pp)) {
        continue;
      }
      // emit check to ensure path param isn't an empty string.  we only need
      // to do this for params that have an underlying type of string.
      const choiceIsString = function (type: go.PathParameterType): boolean {
        if (!go.isConstantType(type)) {
          return false;
        }
        return type.type === 'string';
      };
      if (((go.isPrimitiveType(pp.type) && pp.type.typeName === 'string') || choiceIsString(pp.type)) && pp.isEncoded) {
        const paramName = helpers.getParamName(pp);
        imports.add('errors');
        text += `\tif ${paramName} == "" {\n`;
        text += `\t\treturn nil, errors.New("parameter ${paramName} cannot be empty")\n`;
        text += '\t}\n';
      }
      let paramValue = helpers.formatParamValue(pp, imports);
      if (pp.isEncoded) {
        imports.add('net/url');
        paramValue = `url.PathEscape(${helpers.formatParamValue(pp, imports)})`;
      }
      text += `\turlPath = strings.ReplaceAll(urlPath, "{${pp.pathSegment}}", ${paramValue})\n`;
    }
  }
  text += `\treq, err := runtime.NewRequest(ctx, http.Method${capitalize(method.httpMethod)}, ${hostParam})\n`;
  text += '\tif err != nil {\n';
  text += '\t\treturn nil, err\n';
  text += '\t}\n';
  // helper to build nil checks for param groups
  const emitParamGroupCheck = function (param: go.Parameter): string {
    if (!param.group) {
      throw new Error(`emitParamGroupCheck called for ungrouped parameter ${param.name}`);
    }
    let client = '';
    if (param.location === 'client') {
      client = 'client.';
    }
    const paramGroupName = uncapitalize(param.group.name);
    let optionalParamGroupCheck = `${client}${paramGroupName} != nil && `;
    if (param.group.required) {
      optionalParamGroupCheck = '';
    }
    return `\tif ${optionalParamGroupCheck}${client}${paramGroupName}.${capitalize(param.name)} != nil {\n`;
  };
  // add query parameters
  const encodedParams = new Array<go.QueryParameter>();
  const unencodedParams = new Array<go.QueryParameter>();
  for (const qp of values(method.parameters)) {
    if (!go.isQueryParameter(qp)) {
      continue;
    }
    if (qp.isEncoded) {
      encodedParams.push(qp);
    } else {
      unencodedParams.push(qp);
    }
  }
  const emitQueryParam = function (qp: go.QueryParameter, setter: string): string {
    let qpText = '';
    if (qp.location === 'method' && go.isClientSideDefault(qp.kind)) {
      qpText = emitClientSideDefault(qp, qp.kind, (name, val) => { return `\treqQP.Set(${name}, ${val})`; }, imports);
    } else if (go.isRequiredParameter(qp) || go.isLiteralParameter(qp) || (qp.location === 'client' && go.isClientSideDefault(qp.kind))) {
      qpText = `\t${setter}\n`;
    } else if (qp.location === 'client' && !qp.group) {
      // global optional param
      qpText = `\tif client.${qp.name} != nil {\n`;
      qpText += `\t\t${setter}\n`;
      qpText += '\t}\n';
    } else {
      qpText = emitParamGroupCheck(qp);
      qpText += `\t\t${setter}\n`;
      qpText += '\t}\n';
    }
    return qpText;
  };
  // emit encoded params first
  if (encodedParams.length > 0) {
    text += '\treqQP := req.Raw().URL.Query()\n';
    for (const qp of values(encodedParams.sort((a: go.QueryParameter, b: go.QueryParameter) => { return helpers.sortAscending(a.queryParameter, b.queryParameter); }))) {
      let setter: string;
      if (go.isQueryCollectionParameter(qp) && qp.collectionFormat === 'multi') {
        setter = `\tfor _, qv := range ${helpers.getParamName(qp)} {\n`;
        // emit a type conversion for the qv based on the array's element type
        let queryVal: string;
        const arrayQP = qp.type;
        if (go.isConstantType(arrayQP.elementType)) {
          const ch = arrayQP.elementType;
          // only string and number types are supported for enums
          if (ch.type === 'string') {
            queryVal = 'string(qv)';
          } else {
            imports.add('fmt');
            queryVal = 'fmt.Sprintf("%d", qv)';
          }
        } else if (go.isPrimitiveType(arrayQP.elementType) && arrayQP.elementType.typeName === 'string') {
          queryVal = 'qv';
        } else {
          imports.add('fmt');
          queryVal = 'fmt.Sprintf("%v", qv)';
        }
        setter += `\t\treqQP.Add("${qp.queryParameter}", ${queryVal})\n`;
        setter += '\t}';
      } else {
        // cannot initialize setter to this value as helpers.formatParamValue() can change imports
        setter = `reqQP.Set("${qp.queryParameter}", ${helpers.formatParamValue(qp, imports)})`;
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
    for (const qp of values(unencodedParams.sort((a: go.QueryParameter, b: go.QueryParameter) => { return helpers.sortAscending(a.queryParameter, b.queryParameter); }))) {
      let setter: string;
      if (go.isQueryCollectionParameter(qp) && qp.collectionFormat === 'multi') {
        setter = `\tfor _, qv := range ${helpers.getParamName(qp)} {\n`;
        setter += `\t\tunencodedParams = append(unencodedParams, "${qp.queryParameter}="+qv)\n`;
        setter += '\t}';
      } else {
        setter = `unencodedParams = append(unencodedParams, "${qp.queryParameter}="+${helpers.formatParamValue(qp, imports)})`;
      }
      text += emitQueryParam(qp, setter);
    }
    imports.add('strings');
    text += '\treq.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")\n';
  }
  if (go.isMethod(method) && method.responseEnvelope.result && go.isBinaryResult(method.responseEnvelope.result)) {
    // skip auto-body downloading for binary stream responses
    text += '\truntime.SkipBodyDownload(req)\n';
  }
  // add specific request headers
  const emitHeaderSet = function (headerParam: go.HeaderParameter, prefix: string): string {
    if (headerParam.location === 'method' && go.isClientSideDefault(headerParam.kind)) {
      return emitClientSideDefault(headerParam, headerParam.kind, (name, val) => {
        return `${prefix}req.Raw().Header[${name}] = []string{${val}}`;
      }, imports);
    } else if (go.isHeaderMapParameter(headerParam)) {
      let headerText = `${prefix}for k, v := range ${helpers.getParamName(headerParam)} {\n`;
      headerText += `${prefix}\tif v != nil {\n`;
      headerText += `${prefix}\t\treq.Raw().Header["${headerParam.collectionPrefix}"+k] = []string{*v}\n`;
      headerText += `${prefix}}\n`;
      headerText += `${prefix}}\n`;
      return headerText;
    } else {
      return `${prefix}req.Raw().Header["${headerParam.headerName}"] = []string{${helpers.formatParamValue(headerParam, imports)}}\n`;
    }
  };
  const headerParams = new Array<go.HeaderParameter>();
  for (const param of values(method.parameters)) {
    if (go.isHeaderParameter(param)) {
      headerParams.push(param);
    }
  }
  for (const param of headerParams.sort((a: go.HeaderParameter, b: go.HeaderParameter) => { return helpers.sortAscending(a.headerName, b.headerName);})) {
    if (param.headerName.match(/^content-type$/)) {
      // canonicalize content-type as req.SetBody checks for it via its canonicalized name :(
      param.headerName = 'Content-Type';
    }
    if (go.isRequiredParameter(param) || go.isLiteralParameter(param) || go.isClientSideDefault(param.kind)) {
      text += emitHeaderSet(param, '\t');
    } else if (param.location === 'client' && !param.group) {
      // global optional param
      text += `\tif client.${param.name} != nil {\n`;
      text += emitHeaderSet(param, '\t');
      text += '\t}\n';
    } else {
      text += emitParamGroupCheck(param);
      text += emitHeaderSet(param, '\t\t');
      text += '\t}\n';
    }
  }

  const partialBodyParams = values(method.parameters).where((param: go.Parameter) => { return go.isPartialBodyParameter(param); }).toArray();
  const bodyParam = <go.BodyParameter | undefined>values(method.parameters).where((each: go.Parameter) => { return go.isBodyParameter(each) || go.isFormBodyParameter(each) || go.isMultipartFormBodyParameter(each); }).first();
  const emitSetBodyWithErrCheck = function(setBodyParam: string): string {
    return `if err := ${setBodyParam}; err != nil {\n\treturn nil, err\n}\n`;
  };

  if (partialBodyParams.length > 0) {
    // partial body params are discrete params that are all fields within an internal struct.
    // define and instantiate an instance of the wire type, using the values from each param.
    text += '\tbody := struct {\n';
    for (const partialBodyParam of <Array<go.PartialBodyParameter>>partialBodyParams) {
      text += `\t\t${capitalize(partialBodyParam.serializedName)} ${helpers.star(partialBodyParam)}${go.getTypeDeclaration(partialBodyParam.type)} \`${partialBodyParam.format.toLowerCase()}:"${partialBodyParam.serializedName}"\`\n`;
    }
    text += '\t}{\n';
    for (const partialBodyParam of <Array<go.PartialBodyParameter>>partialBodyParams) {
      let addr = '&';
      if (go.isRequiredParameter(partialBodyParam)) {
        addr = '';
      }
      text += `\t\t${capitalize(partialBodyParam.serializedName)}: ${addr}${uncapitalize(partialBodyParam.name)},\n`;
    }
    text += '\t}\n\tif err := runtime.MarshalAsJSON(req, body); err != nil {\n\t\treturn nil, err\n\t}\n';
    text += '\treturn req, nil\n';
  } else if (!bodyParam) {
    text += '\treturn req, nil\n';
  } else if (bodyParam.bodyFormat === 'JSON' || bodyParam.bodyFormat === 'XML') {
    // default to the body param name
    let body = helpers.getParamName(bodyParam);
    if (go.isLiteralValue(bodyParam.type)) {
      // if the value is constant, embed it directly
      body = helpers.formatLiteralValue(bodyParam.type, true);
    } else if (bodyParam.bodyFormat === 'XML' && go.isSliceType(bodyParam.type)) {
      // for XML payloads, create a wrapper type if the payload is an array
      imports.add('encoding/xml');
      text += '\ttype wrapper struct {\n';
      let tagName = go.getTypeDeclaration(bodyParam.type);
      if (bodyParam.xml?.name) {
        tagName = bodyParam.xml.name;
      }
      text += `\t\tXMLName xml.Name \`xml:"${tagName}"\`\n`;
      const fieldName = capitalize(bodyParam.name);
      let tag = go.getTypeDeclaration(bodyParam.type.elementType);
      if (go.isModelType(bodyParam.type.elementType) && bodyParam.type.elementType.xml?.name) {
        tag = bodyParam.type.elementType.xml.name;
      }
      text += `\t\t${fieldName} *${go.getTypeDeclaration(bodyParam.type)} \`xml:"${tag}"\`\n`;
      text += '\t}\n';
      let addr = '&';
      if (!go.isRequiredParameter(bodyParam) && !bodyParam.byValue) {
        addr = '';
      }
      body = `wrapper{${fieldName}: ${addr}${body}}`;
    } else if (go.isTimeType(bodyParam.type) && bodyParam.type.dateTimeFormat !== 'dateTimeRFC3339') {
      // wrap the body in the internal time type
      // no need for dateTimeRFC3339 as the JSON marshaler defaults to that.
      body = `${bodyParam.type.dateTimeFormat}(${body})`;
    } else if (isArrayOfDateTimeForMarshalling(bodyParam.type)) {
      const timeInfo = isArrayOfDateTimeForMarshalling(bodyParam.type);
      let elementPtr = '*';
      if (timeInfo?.elemByVal) {
        elementPtr = '';
      }
      text += `\taux := make([]${elementPtr}${timeInfo?.format}, len(${body}))\n`;
      text += `\tfor i := 0; i < len(${body}); i++ {\n`;
      text += `\t\taux[i] = (${elementPtr}${timeInfo?.format})(${body}[i])\n`;
      text += '\t}\n';
      body = 'aux';
    } else if (isMapOfDateTime(bodyParam.type)) {
      const timeType = isMapOfDateTime(bodyParam.type);
      text += `\taux := map[string]*${timeType}{}\n`;
      text += `\tfor k, v := range ${body} {\n`;
      text += `\t\taux[k] = (*${timeType})(v)\n`;
      text += '\t}\n';
      body = 'aux';
    }
    let setBody = `runtime.MarshalAs${getMediaFormat(bodyParam.type, bodyParam.bodyFormat, `req, ${body}`)}`;
    if (go.isSliceType(bodyParam.type) && bodyParam.type.rawJSONAsBytes) {
      imports.add('bytes');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
      setBody = `req.SetBody(streaming.NopCloser(bytes.NewReader(${body})), "application/${bodyParam.bodyFormat.toLowerCase()}")`;
    }
    if (go.isRequiredParameter(bodyParam) || go.isLiteralParameter(bodyParam)) {
      text += `\t${emitSetBodyWithErrCheck(setBody)}`;
      text += '\treturn req, nil\n';
    } else {
      text += emitParamGroupCheck(bodyParam);
      text += `\t${emitSetBodyWithErrCheck(setBody)}`;
      text += '\t\treturn req, nil\n';
      text += '\t}\n';
      text += '\treturn req, nil\n';
    }
  } else if (bodyParam.bodyFormat === 'binary') {
    if (go.isRequiredParameter(bodyParam)) {
      text += `\t${emitSetBodyWithErrCheck(`req.SetBody(${bodyParam.name}, ${bodyParam.contentType})`)}`;
      text += '\treturn req, nil\n';
    } else {
      text += emitParamGroupCheck(bodyParam);
      text += `\t${emitSetBodyWithErrCheck(`req.SetBody(${helpers.getParamName(bodyParam)}, ${bodyParam.contentType})`)}`;
      text += '\treturn req, nil\n';
      text += '\t}\n';
      text += '\treturn req, nil\n';
    }
  } else if (bodyParam.bodyFormat === 'Text') {
    imports.add('strings');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
    const bodyParam = <go.BodyParameter>values(method.parameters).where((each: go.Parameter) => { return go.isBodyParameter(each); }).first();
    if (go.isRequiredParameter(bodyParam)) {
      text += `\tbody := streaming.NopCloser(strings.NewReader(${bodyParam.name}))\n`;
      text += `\t${emitSetBodyWithErrCheck(`req.SetBody(body, ${bodyParam.contentType})`)}`;
      text += '\treturn req, nil\n';
    } else {
      text += emitParamGroupCheck(bodyParam);
      text += `\tbody := streaming.NopCloser(strings.NewReader(${helpers.getParamName(bodyParam)}))\n`;
      text += `\t${emitSetBodyWithErrCheck(`req.SetBody(body, ${bodyParam.contentType})`)}`;
      text += '\treturn req, nil\n';
      text += '\t}\n';
      text += '\treturn req, nil\n';
    }
  } else if (go.isMultipartFormBodyParameter(bodyParam)) {
    if (go.isModelType(bodyParam.type) && bodyParam.type.annotations.multipartFormData) {
      text += `\tformData, err := ${bodyParam.name}.toMultipartFormData()\n`;
      text += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    } else {
      text += '\tformData := map[string]any{}\n';
      for (const param of values(method.parameters)) {
        if (!go.isMultipartFormBodyParameter(param)) {
          continue;
        }
        const setter = `formData["${param.name}"] = ${helpers.getParamName(param)}`;
        if (go.isRequiredParameter(param)) {
          text += `\t${setter}\n`;
        } else {
          text += emitParamGroupCheck(param);
          text += `\t${setter}\n\t}\n`;
        }
      }
    }
    text += '\tif err := runtime.SetMultipartFormData(req, formData); err != nil {\n\t\treturn nil, err\n\t}\n';
    text += '\treturn req, nil\n';
  } else if (go.isFormBodyParameter(bodyParam)) {
    const emitFormData = function (param: go.Parameter, setter: string): string {
      let formDataText = '';
      if (go.isRequiredParameter(param)) {
        formDataText = `\t${setter}\n`;
      } else {
        formDataText = emitParamGroupCheck(param);
        formDataText += `\t\t${setter}\n`;
        formDataText += '\t}\n';
      }
      return formDataText;
    };
    imports.add('net/url');
    imports.add('strings');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
    text += '\tformData := url.Values{}\n';
    // find all the form body params
    for (const param of values(method.parameters)) {
      if (go.isFormBodyParameter(param)) {
        const setter = `formData.Set("${param.formDataName}", ${helpers.formatParamValue(param, imports)})`;
        text += emitFormData(param, setter);
      }
    }
    text += '\tbody := streaming.NopCloser(strings.NewReader(formData.Encode()))\n';
    text += `\t${emitSetBodyWithErrCheck('req.SetBody(body, "application/x-www-form-urlencoded")')}`;
    text += '\treturn req, nil\n';
  } else {
    text += '\treturn req, nil\n';
  }
  text += '}\n\n';
  return text;
}

function emitClientSideDefault(param: go.HeaderParameter | go.QueryParameter, csd: go.ClientSideDefault, setterFormat: (name: string, val: string) => string, imports: ImportManager): string {
  const defaultVar = uncapitalize(param.name) + 'Default';
  let text = `\t${defaultVar} := ${helpers.formatLiteralValue(csd.defaultValue, true)}\n`;
  text += `\tif options != nil && options.${capitalize(param.name)} != nil {\n`;
  text += `\t\t${defaultVar} = *options.${capitalize(param.name)}\n`;
  text += '}\n';
  let serializedName: string;
  if (go.isHeaderParameter(param)) {
    serializedName = param.headerName;
  } else {
    serializedName = param.queryParameter;
  }
  text += setterFormat(`"${serializedName}"`, helpers.formatValue(defaultVar, param.type, imports)) + '\n';
  return text;
}

function getMediaFormat(type: go.PossibleType, mediaType: 'JSON' | 'XML', param: string): string {
  let marshaller: 'JSON' | 'XML' | 'ByteArray' = mediaType;
  let format = '';
  if (go.isBytesType(type)) {
    marshaller = 'ByteArray';
    format = `, runtime.Base64${type.encoding}Format`;
  }
  return `${marshaller}(${param}${format})`;
}

function isArrayOfDateTimeForMarshalling(paramType: go.PossibleType): { format: go.DateTimeFormat, elemByVal: boolean } | undefined {
  if (!go.isSliceType(paramType)) {
    return undefined;
  }
  if (!go.isTimeType(paramType.elementType)) {
    return undefined;
  }
  switch (paramType.elementType.dateTimeFormat) {
    case 'dateType':
    case 'dateTimeRFC1123':
    case 'timeRFC3339':
    case 'timeUnix':
      return {
        format: paramType.elementType.dateTimeFormat,
        elemByVal: paramType.elementTypeByValue
      };
    default:
      // dateTimeRFC3339 uses the default marshaller
      return undefined;
  }
}

// returns true if the method requires a response handler.
// this is used to unmarshal the response body, parse response headers, or both.
function needsResponseHandler(method: go.Method): boolean {
  return helpers.hasSchemaResponse(method) || method.responseEnvelope.headers.length > 0;
}

function generateResponseUnmarshaller(method: go.Method, type: go.PossibleType, format: go.ResultFormat, unmarshalTarget: string): string {
  let unmarshallerText = '';
  const zeroValue = getZeroReturnValue(method, 'handler');
  if (go.isTimeType(type)) {
    // use the designated time type for unmarshalling
    unmarshallerText += `\tvar aux *${type.dateTimeFormat}\n`;
    unmarshallerText += `\tif err := runtime.UnmarshalAs${format}(resp, &aux); err != nil {\n`;
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += `\tresult.${helpers.getResultFieldName(method)} = (*time.Time)(aux)\n`;
    return unmarshallerText;
  } else if (isArrayOfDateTime(type)) {
    // unmarshalling arrays of date/time is a little more involved
    const timeInfo = isArrayOfDateTime(type);
    let elementPtr = '*';
    if (timeInfo?.elemByVal) {
      elementPtr = '';
    }
    unmarshallerText += `\tvar aux []${elementPtr}${timeInfo?.format}\n`;
    unmarshallerText += `\tif err := runtime.UnmarshalAs${format}(resp, &aux); err != nil {\n`;
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += `\tcp := make([]${elementPtr}time.Time, len(aux))\n`;
    unmarshallerText += '\tfor i := 0; i < len(aux); i++ {\n';
    unmarshallerText += `\t\tcp[i] = (${elementPtr}time.Time)(aux[i])\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += `\tresult.${helpers.getResultFieldName(method)} = cp\n`;
    return unmarshallerText;
  } else if (isMapOfDateTime(type)) {
    unmarshallerText += `\taux := map[string]*${isMapOfDateTime(type)}{}\n`;
    unmarshallerText += `\tif err := runtime.UnmarshalAs${format}(resp, &aux); err != nil {\n`;
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += '\tcp := map[string]*time.Time{}\n';
    unmarshallerText += '\tfor k, v := range aux {\n';
    unmarshallerText += '\t\tcp[k] = (*time.Time)(v)\n';
    unmarshallerText += '\t}\n';
    unmarshallerText += `\tresult.${helpers.getResultFieldName(method)} = cp\n`;
    return unmarshallerText;
  }
  if (format === 'JSON' || format === 'XML') {
    if (go.isSliceType(type) && type.rawJSONAsBytes) {
      unmarshallerText += '\tbody, err := runtime.Payload(resp)\n';
      unmarshallerText += '\tif err != nil {\n';
      unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
      unmarshallerText += '\t}\n';
      unmarshallerText += `\t${unmarshalTarget} = body\n`;
    } else {
      unmarshallerText += `\tif err := runtime.UnmarshalAs${getMediaFormat(type, format, `resp, &${unmarshalTarget}`)}; err != nil {\n`;
      unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
      unmarshallerText += '\t}\n';
    }
  } else if (format === 'Text') {
    unmarshallerText += '\tbody, err := runtime.Payload(resp)\n';
    unmarshallerText += '\tif err != nil {\n';
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += '\ttxt := string(body)\n';
    unmarshallerText += `\t${unmarshalTarget} = &txt\n`;
  } else {
    // the remaining formats should have been handled elsewhere
    throw new Error(`unhandled format ${format} for operation ${method.client.name}.${method.name}`);
  }
  return unmarshallerText;
}

function createProtocolResponse(client: go.Client, method: go.Method, imports: ImportManager): string {
  if (!needsResponseHandler(method)) {
    return '';
  }
  const name = method.naming.responseMethod;
  let text = `${comment(name, '// ')} handles the ${method.name} response.\n`;
  text += `func (client *${client.name}) ${name}(resp *http.Response) (${generateReturnsInfo(method, 'handler').join(', ')}) {\n`;

  const addHeaders = function (headers: Array<go.HeaderResponse | go.HeaderMapResponse>) {
    for (const header of values(headers)) {
      text += formatHeaderResponseValue(header, imports, 'result', `${method.responseEnvelope.name}{}`);
    }
  };

  const result = method.responseEnvelope.result;
  if (!result) {
    // only headers
    text += `\tresult := ${method.responseEnvelope.name}{}\n`;
    addHeaders(method.responseEnvelope.headers);
  } else if (go.isAnyResult(result)) {
    imports.add('fmt');
    text += `\tresult := ${method.responseEnvelope.name}{}\n`;
    addHeaders(method.responseEnvelope.headers);
    text += '\tswitch resp.StatusCode {\n';
    for (const statusCode of method.httpStatusCodes) {
      text += `\tcase ${helpers.formatStatusCodes([statusCode])}:\n`;
      const resultType = result.httpStatusCodeType[statusCode];
      if (!resultType) {
        // the operation contains a mix of schemas and non-schema responses
        continue;
      }
      text += `\tvar val ${go.getTypeDeclaration(resultType)}\n`;
      text += generateResponseUnmarshaller(method, resultType, result.format, 'val');
      text += '\tresult.Value = val\n';
    }
    text += '\tdefault:\n';
    text += `\t\treturn ${getZeroReturnValue(method, 'handler')}, fmt.Errorf("unhandled HTTP status code %d", resp.StatusCode)\n`;
    text += '\t}\n';
  } else if (go.isBinaryResult(result)) {
    text += `\tresult := ${method.responseEnvelope.name}{${result.fieldName}: resp.Body}\n`;
    addHeaders(method.responseEnvelope.headers);
  } else if (go.isHeadAsBooleanResult(result)) { 
    text += `\tresult := ${method.responseEnvelope.name}{${result.fieldName}: resp.StatusCode >= 200 && resp.StatusCode < 300}\n`;
    addHeaders(method.responseEnvelope.headers);
  } else if (go.isMonomorphicResult(result)) {
    text += `\tresult := ${method.responseEnvelope.name}{}\n`;
    addHeaders(method.responseEnvelope.headers);
    let target = `result.${helpers.getResultFieldName(method)}`;
    // when unmarshalling a wrapped XML array, unmarshal into the response envelope
    if (result.format === 'XML' && go.isSliceType(result.monomorphicType)) {
      target = 'result';
    }
    text += generateResponseUnmarshaller(method, result.monomorphicType, result.format, target);
  } else if (go.isPolymorphicResult(result)) {
    text += `\tresult := ${method.responseEnvelope.name}{}\n`;
    addHeaders(method.responseEnvelope.headers);
    text += generateResponseUnmarshaller(method, result.interfaceType, result.format, 'result');
  } else if (go.isModelResult(result)) {
    text += `\tresult := ${method.responseEnvelope.name}{}\n`;
    addHeaders(method.responseEnvelope.headers);
    text += generateResponseUnmarshaller(method, result.modelType, result.format, `result.${helpers.getResultFieldName(method)}`);
  } else {
    throw new Error(`unhandled result type for ${client.name}.${method.name}`);
  }

  text += '\treturn result, nil\n';
  text += '}\n\n';
  return text;
}

function isArrayOfDateTime(paramType: go.PossibleType): { format: go.DateTimeFormat, elemByVal: boolean } | undefined {
  if (!go.isSliceType(paramType)) {
    return undefined;
  }
  if (!go.isTimeType(paramType.elementType)) {
    return undefined;
  }
  return {
    format: paramType.elementType.dateTimeFormat,
    elemByVal: paramType.elementTypeByValue
  };
}

function isMapOfDateTime(paramType: go.PossibleType): string | undefined {
  if (!go.isMapType(paramType)) {
    return undefined;
  }
  if (!go.isTimeType(paramType.valueType)) {
    return undefined;
  }
  return paramType.valueType.dateTimeFormat;
}

// returns the parameters for the public API
// e.g. "ctx context.Context, i int, s string"
function getAPIParametersSig(method: go.Method, imports: ImportManager, pkgName?: string): string {
  const methodParams = helpers.getMethodParameters(method);
  const params = new Array<string>();
  if (!go.isPageableMethod(method) || go.isLROMethod(method)) {
    imports.add('context');
    params.push('ctx context.Context');
  }
  for (const methodParam of values(methodParams)) {
    params.push(`${uncapitalize(methodParam.name)} ${helpers.formatParameterTypeName(methodParam, pkgName)}`);
  }
  return params.join(', ');
}

// returns the return signature where each entry is the type name
// e.g. [ '*string', 'error' ]
// apiType describes where the return sig is used.
//   api - for the API definition
//    op - for the operation
// handler - for the response handler
function generateReturnsInfo(method: go.Method, apiType: 'api' | 'op' | 'handler'): Array<string> {
  let returnType = method.responseEnvelope.name;
  if (go.isLROMethod(method)) {
    switch (apiType) {
      case 'api':
        if (go.isPageableMethod(method)) {
          returnType = `*runtime.Poller[*runtime.Pager[${returnType}]]`;
        } else {
          returnType = `*runtime.Poller[${returnType}]`;
        }
        break;
      case 'handler':
        // we only have a handler for operations that return a schema
        if (!go.isPageableMethod(method)) {
          throw new Error(`handler being generated for non-pageable LRO ${method.name} which is unexpected`);
        }
        break;
      case 'op':
        returnType = '*http.Response';
        break;
    }
  } else if (go.isPageableMethod(method)) {
    switch (apiType) {
      case 'api':
      case 'op':
        // pager operations don't return an error
        return [`*runtime.Pager[${returnType}]`];
    }
  }
  return [returnType, 'error'];
}

function generateLROBeginMethod(client: go.Client, method: go.LROMethod, imports: ImportManager, injectSpans: boolean, generateFakes: boolean): string {
  const params = getAPIParametersSig(method, imports);
  const returns = generateReturnsInfo(method, 'api');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
  let text = '';
  if (method.description) {
    text += `${comment(`${fixUpMethodName(method)} - ${method.description}`, '//', undefined, helpers.commentLength)}\n`;
    text += genRespErrorDoc(method);
    text += genApiVersionDoc(method.apiVersions);
  }
  const zeroResp = getZeroReturnValue(method, 'api');
  const methodParams = helpers.getMethodParameters(method);
  for (const param of values(methodParams)) {
    if (param.description) {
      text += `${helpers.formatCommentAsBulletItem(`${param.name} - ${param.description}`)}\n`;
    }
  }
  text += `func (client *${client.name}) ${fixUpMethodName(method)}(${params}) (${returns.join(', ')}) {\n`;
  let pollerType = 'nil';
  let pollerTypeParam = `[${method.responseEnvelope.name}]`;
  if (go.isPageableMethod(method)) {
    // for paged LROs, we construct a pager and pass it to the LRO ctor.
    pollerTypeParam = `[*runtime.Pager${pollerTypeParam}]`;
    pollerType = '&pager';
    text += '\tpager := ';
    text += emitPagerDefinition(client, method, imports, injectSpans, generateFakes);
  }

  text += '\tif options == nil || options.ResumeToken == "" {\n';

  // creating the poller from response branch

  const opName = method.naming.internalMethod;
  text += `\t\tresp, err := client.${opName}(${helpers.getCreateRequestParameters(method)})\n`;
  text += '\t\tif err != nil {\n';
  text += `\t\t\treturn ${zeroResp}, err\n`;
  text += '\t\t}\n';

  let finalStateVia = '';
  // LRO operation might have a special configuration set in x-ms-long-running-operation-options
  // which indicates a specific url to perform the final Get operation on
  if (method.finalStateVia) {
    switch (method.finalStateVia) {
      case 'azure-async-operation':
        finalStateVia = 'runtime.FinalStateViaAzureAsyncOp';
        break;
      case 'location':
        finalStateVia = 'runtime.FinalStateViaLocation';
        break;
      case 'original-uri':
        finalStateVia = 'runtime.FinalStateViaOriginalURI';
        break;
      case 'operation-location':
        finalStateVia = 'runtime.FinalStateViaOpLocation';
        break;
      default:
        throw new Error(`unhandled final-state-via value ${finalStateVia}`);
    }
  }

  text += '\t\tpoller, err := runtime.NewPoller';
  if (finalStateVia === '' && pollerType === 'nil' && !injectSpans) {
    // the generic type param is redundant when it's also specified in the
    // options struct so we only include it when there's no options.
    text += pollerTypeParam;
  }
  text += '(resp, client.internal.Pipeline(), ';
  if (finalStateVia === '' && pollerType === 'nil' && !injectSpans) {
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
    if (injectSpans) {
      text += '\t\t\tTracer: client.internal.Tracer(),\n';
    }
    text += '\t\t})\n';
  }
  text += '\t\treturn poller, err\n';
  text += '\t} else {\n';

  // creating the poller from resume token branch

  text += '\t\treturn runtime.NewPollerFromResumeToken';
  if (pollerType === 'nil' && !injectSpans) {
    text += pollerTypeParam;
  }
  text += '(options.ResumeToken, client.internal.Pipeline(), ';
  if (pollerType === 'nil' && !injectSpans) {
    text += 'nil)\n';
  } else {
    text += `&runtime.NewPollerFromResumeTokenOptions${pollerTypeParam}{\n`;
    if (pollerType !== 'nil') {
      text += `\t\t\tResponse: ${pollerType},\n`;
    }
    if (injectSpans) {
      text += '\t\t\tTracer: client.internal.Tracer(),\n';
    }
    text  += '\t\t})\n';
  }
  text += '\t}\n';

  text += '}\n\n';
  return text;
}

export function fixUpMethodName(method: go.Method): string {
  if (go.isLROMethod(method)) {
    return `Begin${method.name}`;
  } else if (go.isPageableMethod(method)) {
    return `New${method.name}Pager`;
  }
  return method.name;
}
