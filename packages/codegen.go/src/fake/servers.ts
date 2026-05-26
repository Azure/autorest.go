/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../../codemodel.go/src/index.js';
import * as naming from '../../../naming.go/src/naming.js';
import * as helpers from '../core/helpers.js';
import { ImportManager } from '../core/imports.js';
import { fixUpMethodName } from '../core/operations.js';
import { generateServerInternal, RequiredHelpers } from './internal.js';
import { CodegenError } from '../core/errors.js';

// contains the generated content for all servers and the required helpers
export class ServerContent {
  readonly servers: Array<OperationGroupContent>;
  readonly internals: string;

  constructor(servers: Array<OperationGroupContent>, internals: string) {
    this.servers = servers;
    this.internals = internals;
  }
}

// represents the generated content for an operation group
export class OperationGroupContent {
  readonly name: string;
  readonly content: string;

  constructor(name: string, content: string) {
    this.name = name;
    this.content = content;
  }
}

// used to track the helpers we need to emit. they're all false by default.
const requiredHelpers = new RequiredHelpers();

export function getServerName(client: go.Client): string {
  // for the fake server, we use the suffix Server instead of Client
  return naming.capitalize(client.name.replace(/[C|c]lient$/, 'Server'));
}

/**
 * Generates the contents for the *_server.go files.
 *
 * @param pkg contains the package content
 * @param target the codegen target for the module
 * @returns the contents to generate or an empty object
 */
export function generateServers(pkg: go.FakePackage, target: go.CodeModelType): ServerContent {
  const operations = new Array<OperationGroupContent>();
  for (const client of pkg.parent.clients) {
    if (client.clientAccessors.length === 0 && helpers.clientHasNoExportedMethods(client)) {
      // client has no client accessors and no exported methods, skip it
      continue;
    }

    // the list of packages to import
    const imports = new ImportManager(pkg);

    // add standard imports
    imports.add('errors');
    imports.add('fmt');
    imports.add('net/http');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');

    const serverName = getServerName(client);

    let content: string;
    content = `// ${serverName} is a fake server for instances of the ${go.getTypeDeclaration(client, pkg)} type.\n`;
    content += `type ${serverName} struct{\n`;

    const indent = new helpers.Indentation();

    // we might remove some operations from the list
    const finalMethods = new Array<go.MethodType>();
    let countLROs = 0;
    let countPagers = 0;

    // add server transports for client accessors
    // we might remove some clients from the list
    const finalSubClients = new Array<go.Client>();
    for (const clientAccessor of client.clientAccessors) {
      if (helpers.clientHasNoExportedMethods(clientAccessor.returns)) {
        // client has no exported methods, skip it
        continue;
      }
      const serverName = getServerName(clientAccessor.returns);
      content += `${indent.get()}// ${serverName} contains the fakes for client ${clientAccessor.returns.name}\n`;
      content += `${indent.get()}${serverName} ${serverName}\n\n`;
      finalSubClients.push(clientAccessor.returns);
    }

    for (const method of client.methods) {
      if (helpers.isMethodInternal(method)) {
        // method isn't exported, don't create a fake for it
        continue;
      }

      let serverResponse: string;

      switch (method.kind) {
        case 'lroMethod':
        case 'lroPageableMethod':
          let respType = go.getTypeDeclaration(method.returns, pkg);
          if (method.kind === 'lroPageableMethod') {
            respType = `azfake.PagerResponder[${respType}]`;
          }
          serverResponse = `resp azfake.PollerResponder[${respType}], errResp azfake.ErrorResponder`;
          break;
        case 'method':
          serverResponse = `resp azfake.Responder[${go.getTypeDeclaration(method.returns, pkg)}], errResp azfake.ErrorResponder`;
          break;
        case 'pageableMethod':
          serverResponse = `resp azfake.PagerResponder[${go.getTypeDeclaration(method.returns, pkg)}]`;
          break;
      }

      const operationName = fixUpMethodName(method);
      content += `${indent.get()}// ${operationName} is the fake for method ${client.name}.${operationName}\n`;
      const successCodes = new Array<string>();
      if (method.returns.result?.kind === 'anyResult') {
        for (const httpStatus of getMethodStatusCodes(method)) {
          const result = method.returns.result.httpStatusCodeType[httpStatus];
          if (!result) {
            // the operation contains a mix of schemas and non-schema responses
            successCodes.push(`${helpers.formatStatusCode(httpStatus)} (no return type)`);
            continue;
          }
          successCodes.push(`${helpers.formatStatusCode(httpStatus)} (returns ${go.getTypeDeclaration(result, pkg)})`);
        }
        content += `${indent.get()}// HTTP status codes to indicate success:\n`;
        for (const successCode of successCodes) {
          content += `${indent.get()}//   - ${successCode}\n`;
        }
      } else {
        for (const statusCode of getMethodStatusCodes(method)) {
          successCodes.push(`${helpers.formatStatusCode(statusCode)}`);
        }
        content += `${indent.get()}// HTTP status codes to indicate success: ${successCodes.join(', ')}\n`;
      }
      content += `${indent.get()}${operationName} func(${getAPIParametersSig(pkg, method, imports)}) (${serverResponse})\n\n`;
      finalMethods.push(method);
      switch (method.kind) {
        case 'lroMethod':
        case 'lroPageableMethod':
          ++countLROs;
          break;
        case 'pageableMethod':
          ++countPagers;
          break;
      }
    }

    content += '}\n\n';

    ///////////////////////////////////////////////////////////////////////////

    const serverTransport = `${serverName}Transport`;

    content += `// New${serverTransport} creates a new instance of ${serverTransport} with the provided implementation.\n`;
    content += `// The returned ${serverTransport} instance is connected to an instance of ${go.getTypeDeclaration(client, pkg)} via the\n`;
    content += "// azcore.ClientOptions.Transporter field in the client's constructor parameters.\n";
    content += `func New${serverTransport}(srv *${serverName}) *${serverTransport} {\n`;
    if (countLROs === 0 && countPagers === 0) {
      content += `${indent.get()}return &${serverTransport}{srv: srv}\n}\n\n`;
    } else {
      content += `${indent.get()}return &${serverTransport}{\n`;
      indent.push();
      content += `${indent.get()}srv: srv,\n`;
      for (const method of finalMethods) {
        let respType = go.getTypeDeclaration(method.returns, pkg);
        switch (method.kind) {
          case 'lroMethod':
          case 'lroPageableMethod':
            if (method.kind === 'lroPageableMethod') {
              respType = `azfake.PagerResponder[${go.getTypeDeclaration(method.returns, pkg)}]`;
            }
            requiredHelpers.tracker = true;
            content += `${indent.get()}${naming.uncapitalize(fixUpMethodName(method))}: newTracker[azfake.PollerResponder[${respType}]](),\n`;
            break;
          case 'pageableMethod':
            requiredHelpers.tracker = true;
            content += `${indent.get()}${naming.uncapitalize(fixUpMethodName(method))}: newTracker[azfake.PagerResponder[${respType}]](),\n`;
            break;
        }
      }
      indent.pop();
      content += `${indent.get()}}\n}\n\n`;
    }

    content += `// ${serverTransport} connects instances of ${go.getTypeDeclaration(client, pkg)} to instances of ${serverName}.\n`;
    content += `// Don't use this type directly, use New${serverTransport} instead.\n`;
    content += `type ${serverTransport} struct {\n`;
    content += `${indent.get()}srv *${serverName}\n`;

    // add server transports for client accessors
    if (finalSubClients.length > 0) {
      requiredHelpers.initServer = true;
      imports.add('sync');
      content += `${indent.get()}trMu sync.Mutex\n`;
      for (const subClient of finalSubClients) {
        const serverName = getServerName(subClient);
        content += `${indent.get()}tr${serverName} *${serverName}Transport\n`;
      }
    }

    for (const method of finalMethods) {
      // create state machines for any pager/poller operations
      let respType = go.getTypeDeclaration(method.returns, pkg);
      switch (method.kind) {
        case 'lroMethod':
        case 'lroPageableMethod':
          if (method.kind === 'lroPageableMethod') {
            respType = `azfake.PagerResponder[${go.getTypeDeclaration(method.returns, pkg)}]`;
          }
          requiredHelpers.tracker = true;
          content += `${indent.get()}${naming.uncapitalize(fixUpMethodName(method))} *tracker[azfake.PollerResponder[${respType}]]\n`;
          break;
        case 'pageableMethod':
          requiredHelpers.tracker = true;
          content += `${indent.get()}${naming.uncapitalize(fixUpMethodName(method))} *tracker[azfake.PagerResponder[${go.getTypeDeclaration(method.returns, pkg)}]]\n`;
          break;
      }
    }
    content += '}\n\n';

    content += generateServerTransportDo(serverTransport, client, finalSubClients, finalMethods, indent);
    content += generateServerTransportClientDispatch(serverTransport, finalSubClients, imports, indent);
    content += generateServerTransportMethodDispatch(serverTransport, client, finalMethods, indent);
    content += generateServerTransportMethods(pkg, serverTransport, finalMethods, imports, indent);

    content += `// set this to conditionally intercept incoming requests to ${serverTransport}\n`;
    content += `var ${getTransportInterceptorVarName(client)} interface {\n`;
    content += `${indent.get()}// Do returns true if the server transport should use the returned response/error\n`;
    content += `${indent.get()}Do(*http.Request) (*http.Response, error, bool)\n}\n`;

    ///////////////////////////////////////////////////////////////////////////

    // stitch everything together
    let text = helpers.contentPreamble(pkg);
    text += imports.text();
    text += content;
    operations.push(new OperationGroupContent(serverName, text));
  }

  if (target === 'azure-arm' && pkg.parent.clients.length > 0) {
    // ARM server factory uses the initServer func
    requiredHelpers.initServer = true;
  }

  return new ServerContent(operations, generateServerInternal(pkg, requiredHelpers));
}

function getTransportInterceptorVarName(client: go.Client): string {
  // use naming.uncapitalize directly instead of helpers.camelCase so distinct
  // server types with sequential duplicate words (e.g. AuthorizationServerServer,
  // generated from client AuthorizationServerClient) don't collide with
  // AuthorizationServer (generated from AuthorizationClient) on the
  // interceptor variable name.
  return `${naming.uncapitalize(getServerName(client))}TransportInterceptor`;
}

// method names for fakes dispatching
const dispatchMethodFake = 'dispatchToMethodFake';
const dispatchToClientFake = 'dispatchToClientFake';

function generateServerTransportDo(
  serverTransport: string,
  client: go.Client,
  finalSubClients: Array<go.Client>,
  finalMethods: Array<go.MethodType>,
  indent: helpers.Indentation,
): string {
  const receiverName = serverTransport[0].toLowerCase();
  let content = `// Do implements the policy.Transporter interface for ${serverTransport}.\n`;
  content += `func (${receiverName} *${serverTransport}) Do(req *http.Request) (*http.Response, error) {\n`;
  content += `${indent.get()}rawMethod := req.Context().Value(runtime.CtxAPINameKey{})\n`;
  content += `${indent.get()}method, ok := rawMethod.(string)\n`;
  content += `${indent.get()}if !ok {\n`;
  content += `${indent.push().get()}return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}\n`;
  content += `${indent.pop().get()}}\n\n`;

  if (finalSubClients.length > 0 && finalMethods.length > 0) {
    // client contains client accessors and methods.
    // if the method isn't for this client, dispatch to the correct client
    content += `${indent.get()}if client := method[:strings.Index(method, ".")]; client != "${client.name}" {\n`;
    content += `${indent.push().get()}return ${receiverName}.${dispatchToClientFake}(req, client)\n`;
    content += `${indent.pop().get()}}\n`;
    // else dispatch to our method fakes
    content += `${indent.get()}return ${receiverName}.${dispatchMethodFake}(req, method)\n`;
  } else if (finalSubClients.length > 0) {
    content += `${indent.get()}return ${receiverName}.${dispatchToClientFake}(req, method[:strings.Index(method, ".")])\n`;
  } else {
    content += `${indent.get()}return ${receiverName}.${dispatchMethodFake}(req, method)\n`;
  }
  content += '}\n\n'; // end Do
  return content;
}

function generateServerTransportClientDispatch(serverTransport: string, subClients: Array<go.Client>, imports: ImportManager, indent: helpers.Indentation): string {
  if (subClients.length === 0) {
    return '';
  }

  /** gathers all children, not just immediate children, in breadth first order */
  const getAllSubClientsForCase = function (client: go.Client): Array<go.Client> {
    const result = new Array<go.Client>();
    const visited = new Set<go.Client>();
    const queue = new Array<go.Client>();
    for (const clientAccessor of client.clientAccessors) {
      if (helpers.clientHasNoExportedMethods(clientAccessor.returns)) {
        continue;
      }
      if (!visited.has(clientAccessor.returns)) {
        visited.add(clientAccessor.returns);
        queue.push(clientAccessor.returns);
        result.push(clientAccessor.returns);
      }
    }
    while (queue.length > 0) {
      const current = queue.shift()!;
      for (const clientAccessor of current.clientAccessors) {
        if (helpers.clientHasNoExportedMethods(clientAccessor.returns)) {
          continue;
        }
        if (!visited.has(clientAccessor.returns)) {
          visited.add(clientAccessor.returns);
          queue.push(clientAccessor.returns);
          result.push(clientAccessor.returns);
        }
      }
    }
    return result;
  };

  const receiverName = serverTransport[0].toLowerCase();
  imports.add('strings');
  let content = `func (${receiverName} *${serverTransport}) ${dispatchToClientFake}(req *http.Request, client string) (*http.Response, error) {\n`;
  content += `${indent.get()}var resp *http.Response\n${indent.get()}var err error\n\n`;
  content += `${indent.get()}switch client {\n`;
  for (const subClient of subClients) {
    // we must include all child clients, not just the immediate children
    const subClientsForCase = getAllSubClientsForCase(subClient);
    const allClientNamesForCase = new Array<string>(`"${subClient.name}"`);
    allClientNamesForCase.push(...subClientsForCase.map((each) => `"${each.name}"`));
    content += `${indent.get()}case ${allClientNamesForCase.join(', ')}:\n`;
    const serverName = getServerName(subClient);
    indent.push();
    content += `${indent.get()}initServer(&${receiverName}.trMu, &${receiverName}.tr${serverName}, func() *${serverName}Transport {\n`;
    content += `${indent.get()}return New${serverName}Transport(&${receiverName}.srv.${serverName}) })\n`;
    content += `${indent.get()}resp, err = ${receiverName}.tr${serverName}.Do(req)\n`;
    indent.pop();
  }
  content += `${indent.get()}default:\n`;
  content += `${indent.push().get()}err = fmt.Errorf("unhandled client %s", client)\n`;
  indent.pop();
  content += `${indent.get()}}\n\n`; // end switch
  content += `${indent.get()}return resp, err\n}\n\n`;
  return content;
}

function generateServerTransportMethodDispatch(serverTransport: string, client: go.Client, finalMethods: Array<go.MethodType>, indent: helpers.Indentation): string {
  if (finalMethods.length === 0) {
    return '';
  }

  const receiverName = serverTransport[0].toLowerCase();
  let content = `func (${receiverName} *${serverTransport}) ${dispatchMethodFake}(req *http.Request, method string) (*http.Response, error) {\n`;
  content += `${indent.get()}resultChan := make(chan result, 1)\n`;
  content += `${indent.get()}go func() {\n`;
  indent.push();
  content += `${indent.get()}var intercepted bool\n${indent.get()}var res result\n`;
  const interceptorVarName = getTransportInterceptorVarName(client);
  content += `${indent.get()} if ${interceptorVarName} != nil {\n`;
  content += `${indent.push().get()} res.resp, res.err, intercepted = ${interceptorVarName}.Do(req)\n`;
  content += `${indent.pop().get()}}\n`;
  content += `${indent.get()}if !intercepted {\n`;
  indent.push();
  content += `${indent.get()}switch method {\n`;

  for (const method of finalMethods) {
    const operationName = fixUpMethodName(method);
    content += `${indent.get()}case "${client.name}.${operationName}":\n`;
    content += `${indent.push().get()}res.resp, res.err = ${receiverName}.dispatch${operationName}(req)\n`;
    indent.pop();
  }

  content += `${indent.push().get()}default:\n`;
  content += `${indent.get()}res.err = fmt.Errorf("unhandled API %s", method)\n`;
  content += `${indent.pop().get()}}\n\n`; // end switch
  content += `${indent.pop().get()}}\n`; // end if !intercepted

  content += `${indent.get()}resultChan <- res\n`;
  content += `${indent.pop().get()}}()\n\n`; // end goroutine

  content += `${indent.get()}select {\n`;
  content += `${indent.get()}case <-req.Context().Done():\n`;
  content += `${indent.push().get()}return nil, req.Context().Err()\n`;
  content += `${indent.pop().get()}case res := <-resultChan:\n`;
  content += `${indent.push().get()}return res.resp, res.err\n`;
  content += `${indent.pop().get()}}\n}\n\n`;

  return content;
}

/**
 * generates the server transport methods for a fake server transport
 *
 * @param pkg contains the package content
 * @param serverTransport the name of the server transport type
 * @param finalMethods the array of methods for which to generate the fake transports
 * @param imports the import manager currently in scope
 * @param indent the indentation helper currently in scope
 * @returns the text for the server transport methods
 */
function generateServerTransportMethods(
  pkg: go.FakePackage,
  serverTransport: string,
  finalMethods: Array<go.MethodType>,
  imports: ImportManager,
  indent: helpers.Indentation,
): string {
  if (finalMethods.length === 0) {
    return '';
  }

  imports.addForPkg(pkg.parent);
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake', 'azfake');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server');
  imports.add('slices');

  const receiverName = serverTransport[0].toLowerCase();

  let content = '';
  for (const method of finalMethods) {
    content += `func (${receiverName} *${serverTransport}) dispatch${fixUpMethodName(method)}(req *http.Request) (*http.Response, error) {\n`;
    content += `${indent.get()}if ${receiverName}.srv.${fixUpMethodName(method)} == nil {\n`;
    content += `${indent.push().get()}return nil, &nonRetriableError{errors.New("fake for method ${fixUpMethodName(method)} not implemented")}\n`;
    content += `${indent.pop().get()}}\n`;

    switch (method.kind) {
      case 'lroMethod':
      case 'lroPageableMethod':
        // must check LRO before pager as you can have paged LROs
        content += dispatchForLROBody(pkg, receiverName, method, imports, indent);
        break;
      case 'method': {
        content += dispatchForOperationBody(pkg, receiverName, method, imports, indent);
        content += `${indent.get()}respContent := server.GetResponseContent(respr)\n`;
        const formattedStatusCodes = helpers.formatStatusCodes(method.httpStatusCodes);
        content += `${indent.get()}if !slices.Contains([]int{${formattedStatusCodes}}, respContent.HTTPStatus) {\n`;
        content += `${indent.push().get()}return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are ${formattedStatusCodes}", respContent.HTTPStatus)}\n`;
        content += `${indent.pop().get()}}\n`;
        if (!method.returns.result || method.returns.result.kind === 'headAsBooleanResult') {
          content += `${indent.get()}resp, err := server.NewResponse(respContent, req, nil)\n`;
        } else if (method.returns.result.kind === 'anyResult') {
          content += `${indent.get()}resp, err := server.MarshalResponseAs${method.returns.result.format}(respContent, server.GetResponse(respr).${getResultFieldName(method.returns.result)}, req)\n`;
        } else if (method.returns.result.kind === 'binaryResult') {
          content += `${indent.get()}resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{\n`;
          indent.push();
          content += `${indent.get()}Body: server.GetResponse(respr).${getResultFieldName(method.returns.result)},\n`;
          content += `${indent.get()}ContentType: req.Header.Get("Content-Type"),\n`;
          content += `${indent.pop().get()}})\n`;
        } else if (method.returns.result.kind === 'monomorphicResult') {
          if (method.returns.result.monomorphicType.kind === 'encodedBytes') {
            const encoding = method.returns.result.monomorphicType.encoding;
            content += `${indent.get()}resp, err := server.MarshalResponseAsByteArray(respContent, server.GetResponse(respr).${getResultFieldName(method.returns.result)}, runtime.Base64${encoding}Format, req)\n`;
          } else if (method.returns.result.monomorphicType.kind === 'rawJSON') {
            imports.add('bytes');
            imports.add('io');
            content += `${indent.get()}resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{\n`;
            indent.push();
            content += `${indent.get()}Body: io.NopCloser(bytes.NewReader(server.GetResponse(respr).RawJSON)),\n`;
            content += `${indent.get()}ContentType: "application/json",\n`;
            content += `${indent.pop().get()}})\n`;
          } else if (method.returns.result.format === 'Text') {
            let contentToMarshal: string;
            const respField = getResultFieldName(method.returns.result);
            const getResponseField = `server.GetResponse(respr).${respField}`;
            switch (method.returns.result.monomorphicType.kind) {
              case 'scalar': {
                imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
                // create a local var that will hold the string-formatted scalar
                contentToMarshal = `formatted${respField}`;
                content += `${indent.get()}var ${contentToMarshal} *string\n`;
                const localVar = naming.uncapitalize(respField);
                const resultType = method.returns.result.monomorphicType;
                // if value := server.GetResponse(respr).Value; value != nil {...format as string...}
                content += `${indent.get()}${helpers.buildIfBlock(indent, {
                  condition: `${localVar} := ${getResponseField}; ${localVar} != nil`,
                  body: (indent) => `${indent.get()}${contentToMarshal} = to.Ptr(${helpers.formatValue(localVar, resultType, imports, true)})\n`,
                })}\n`;
                break;
              }
              case 'string':
                contentToMarshal = getResponseField;
                break;
              default:
                throw new CodegenError(
                  'UnsupportedTsp',
                  `unsupported text return kind ${method.returns.result.monomorphicType.kind} for method ${method.receiver.type.name}.${method.name}`,
                );
            }
            content += `${indent.get()}resp, err := server.MarshalResponseAsText(respContent, ${contentToMarshal}, req)\n`;
          } else {
            let respField = `.${getResultFieldName(method.returns.result)}`;
            if (method.returns.result.format === 'XML' && method.returns.result.monomorphicType.kind === 'slice') {
              respField = '';
            }
            let responseField = `server.GetResponse(respr)${respField}`;
            if (method.returns.result.monomorphicType.kind === 'time') {
              imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/datetime');
              responseField = `(*datetime.${method.returns.result.monomorphicType.format})(${responseField})`;
            }
            content += `${indent.get()}resp, err := server.MarshalResponseAs${method.returns.result.format}(respContent, ${responseField}, req)\n`;
          }
        } else if (method.returns.result.kind === 'modelResult' || method.returns.result.kind === 'polymorphicResult') {
          const respField = `.${getResultFieldName(method.returns.result)}`;
          const responseField = `server.GetResponse(respr)${respField}`;
          content += `${indent.get()}resp, err := server.MarshalResponseAs${method.returns.result.format}(respContent, ${responseField}, req)\n`;
        }

        content += `${indent.get()}if err != nil {\n`;
        content += `${indent.push().get()}return nil, err\n`;
        content += `${indent.pop().get()}}\n`;

        // propagate any header response values into the *http.Response
        for (const header of method.returns.headers) {
          if (header.kind === 'headerMapResponse') {
            content += `${indent.get()}for k, v := range server.GetResponse(respr).${header.fieldName} {\n`;
            content += `${indent.push().get()}if v != nil {\n`;
            content += `${indent.push().get()}resp.Header.Set("${header.headerName}"+k, *v)\n`;
            content += `${indent.pop().get()}}\n`;
            content += `${indent.pop().get()}}\n`;
          } else {
            content += `${indent.get()}if val := server.GetResponse(respr).${header.fieldName}; val != nil {\n`;
            content += `${indent.push().get()}resp.Header.Set("${header.headerName}", ${helpers.formatValue('val', header.type, imports, true)})\n`;
            content += `${indent.pop().get()}}\n`;
          }
        }

        content += `${indent.get()}return resp, nil\n`;
        break;
      }
      case 'pageableMethod':
        content += dispatchForPagerBody(pkg, receiverName, method, imports, indent);
        break;
      default:
        method satisfies never;
    }
    content += '}\n\n';
  }

  return content;
}

/**
 * generates the core dispatching logic for a server dispatch method.
 * this code is common to all method types.
 *
 * @param pkg contains the package content
 * @param receiverName the name of the receiver for the dispatch method
 * @param method the method for which to emit dispatching logic
 * @param imports the import manager currently in scope
 * @param indent the indentation helper currently in scope
 * @returns the text for dispatching logic
 */
function dispatchForOperationBody(pkg: go.FakePackage, receiverName: string, method: go.MethodType, imports: ImportManager, indent: helpers.Indentation): string {
  const methodParamGroups = helpers.getMethodParamGroups(method);
  const numPathParams = methodParamGroups.pathParams.filter((each: go.PathParameter) => !go.isLiteralParameter(each.style)).length;
  let content = '';
  if (numPathParams > 0) {
    imports.add('regexp');
    content += `${indent.get()}const regexStr = \`${createPathParamsRegex(method, methodParamGroups.pathParams)}\`\n`;
    content += `${indent.get()}regex := regexp.MustCompile(regexStr)\n`;
    content += `${indent.get()}matches := regex.FindStringSubmatch(req.URL.EscapedPath())\n`;
    // the total number of matches is the number of capture groups
    // plus the full match. so we add + 1 to include the full match.
    content += `${indent.get()}if len(matches) < ${numPathParams + 1} {\n`;
    content += `${indent.push().get()}return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)\n${indent.pop().get()}}\n`;
  }

  const allQueryParams = methodParamGroups.encodedQueryParams.concat(methodParamGroups.unencodedQueryParams);
  if (allQueryParams.find((each: go.QueryParameter) => each.location === 'method' && !go.isLiteralParameter(each.style))) {
    content += `${indent.get()}qp := req.URL.Query()\n`;
  }

  // note that these are mutually exclusive
  const bodyParam = methodParamGroups.bodyParam;
  const formBodyParams = methodParamGroups.formBodyParams;
  const multipartBodyParams = methodParamGroups.multipartBodyParams;
  const partialBodyParams = methodParamGroups.partialBodyParams;

  if (bodyParam) {
    switch (bodyParam.bodyFormat) {
      case 'JSON':
      case 'XML':
        if (bodyParam && !go.isLiteralParameter(bodyParam.style)) {
          imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake', 'azfake');
          switch (bodyParam.type.kind) {
            case 'encodedBytes':
              content += `${indent.get()}body, err := server.UnmarshalRequestAsByteArray(req, runtime.Base64${bodyParam.type.encoding}Format)\n`;
              content += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
              break;
            case 'interface':
              requiredHelpers.readRequestBody = true;
              content += `${indent.get()}raw, err := readRequestBody(req)\n`;
              content += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
              content += `${indent.get()}body, err := unmarshal${bodyParam.type.name}(raw)\n`;
              content += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
              break;
            case 'rawJSON':
              imports.add('io');
              content += `${indent.get()}body, err := io.ReadAll(req.Body)\n`;
              content += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
              content += `${indent.get()}req.Body.Close()\n`;
              break;
            default: {
              let bodyTypeName = go.getTypeDeclaration(bodyParam.type, pkg);
              if (bodyParam.type.kind === 'time') {
                imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/datetime');
                bodyTypeName = `datetime.${bodyParam.type.format}`;
              }
              content += `${indent.get()}body, err := server.UnmarshalRequestAs${bodyParam.bodyFormat}[${bodyTypeName}](req)\n`;
              content += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
            }
          }
        }
        break;
      case 'Text':
        if (bodyParam && !go.isLiteralParameter(bodyParam.style)) {
          imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake', 'azfake');
          content += `${indent.get()}body, err := server.UnmarshalRequestAsText(req)\n`;
          content += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
        }
        break;
    }
    // nothing to do for binary media type
  } else if (multipartBodyParams.length > 0) {
    imports.add('io');
    imports.add('mime');
    imports.add('mime/multipart');
    content += `${indent.get()}_, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))\n`;
    content += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
    content += `${indent.get()}reader := multipart.NewReader(req.Body, params["boundary"])\n`;
    for (const param of multipartBodyParams) {
      content += `${indent.get()}var ${param.name} ${go.getTypeDeclaration(param.type, pkg)}\n`;
    }

    content += `${indent.get()}for {\n`;
    indent.push();
    content += `${indent.get()}var part *multipart.Part\n`;
    content += `${indent.get()}part, err = reader.NextPart()\n`;
    content += `${indent.get()}if errors.Is(err, io.EOF) {\n${indent.push().get()}break\n`;
    content += `${indent.pop().get()}} else if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
    content += `${indent.get()}var content []byte\n`;
    content += `${indent.get()}switch fn := part.FormName(); fn {\n`;

    // specify boolTarget if parsing bools happens in place.
    // i.e. the result from the parsing doesn't require further conversion (e.g. casting)
    // otherwise the parsed value is in a local var named parsed.
    const parsePrimitiveType = function (typeName: go.ScalarType, boolTarget?: string): string {
      let parseErr = 'parseErr';
      const parseResults = `parsed, ${parseErr}`;
      let parsingCode = '';
      imports.add('strconv');
      switch (typeName) {
        case 'bool':
          if (boolTarget) {
            // we reuse the err var declared earlier when calling reader.NextPart()
            parsingCode = `${indent.get()}${boolTarget}, err = strconv.ParseBool(string(content))\n`;
            parseErr = 'err';
          } else {
            parsingCode = `${indent.get()}${parseResults} := strconv.ParseBool(string(content))\n`;
          }
          break;
        case 'float32':
        case 'float64':
          parsingCode = `${indent.get()}${parseResults} := strconv.ParseFloat(string(content), ${helpers.getBitSizeForNumber(typeName)})\n`;
          break;
        case 'int8':
        case 'int16':
        case 'int32':
        case 'int64':
          parsingCode = `${indent.get()}${parseResults} := strconv.ParseInt(string(content), 10, ${helpers.getBitSizeForNumber(typeName)})\n`;
          break;
        default:
          throw new CodegenError('InternalError', `unhandled multipart parameter primitive type ${typeName}`);
      }
      parsingCode += `${indent.get()}if ${parseErr} != nil {\n${indent.push().get()}return nil, ${parseErr}\n${indent.pop().get()}}\n`;
      return parsingCode;
    };

    const isModelType = function (type: go.WireType): type is go.Model | go.PolymorphicModel {
      return type.kind === 'model' || type.kind === 'polymorphicModel';
    };

    const emitCase = function (caseValue: string, paramVar: string, type: go.WireType, destIsByValue: boolean): string {
      let caseContent = `${indent.get()}case "${caseValue}":\n`;
      indent.push();
      caseContent += `${indent.get()}content, err = io.ReadAll(part)\n`;
      caseContent += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
      let assignedValue: string | undefined;
      if (isModelType(helpers.recursiveUnwrapMapSlice(type))) {
        imports.add('encoding/json');
        caseContent += `${indent.get()}if err = json.Unmarshal(content, &${paramVar}); err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
      } else if (type.kind === 'readSeekCloser') {
        imports.add('bytes');
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
        assignedValue = 'streaming.NopCloser(bytes.NewReader(content))';
      } else if (type.kind === 'constant') {
        let from: string;
        switch (type.type) {
          case 'bool':
          case 'float32':
          case 'float64':
          case 'int32':
          case 'int64':
            caseContent += parsePrimitiveType(type.type);
            from = 'parsed';
            break;
          case 'string':
            from = 'content';
            break;
        }
        assignedValue = `${go.getTypeDeclaration(type, pkg)}(${from})`;
      } else if (type.kind === 'scalar') {
        switch (type.type) {
          case 'bool':
            imports.add('strconv');
            // ParseBool happens in place, so no need to set assignedValue
            caseContent += parsePrimitiveType(type.type, paramVar);
            break;
          case 'float32':
          case 'float64':
          case 'int8':
          case 'int16':
          case 'int32':
          case 'int64':
            caseContent += parsePrimitiveType(type.type);
            assignedValue = `${type.type}(parsed)`;
            break;
          default:
            throw new CodegenError('InternalError', `unhandled multipart parameter primitive type ${type.type}`);
        }
      } else if (type.kind === 'string') {
        assignedValue = 'string(content)';
      } else if (helpers.recursiveUnwrapMapSlice(type).kind === 'multipartContent') {
        imports.add('bytes');
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
        const bodyContent = 'streaming.NopCloser(bytes.NewReader(content))';
        const contentType = 'part.Header.Get("Content-Type")';
        const filename = 'part.FileName()';
        if (type.kind === 'slice') {
          caseContent += `${indent.get()}${paramVar} = append(${paramVar}, streaming.MultipartContent{\n`;
          indent.push();
          caseContent += `${indent.get()}Body: ${bodyContent},\n`;
          caseContent += `${indent.get()}ContentType: ${contentType},\n`;
          caseContent += `${indent.get()}Filename: ${filename},\n`;
          caseContent += `${indent.pop().get()}})\n`;
        } else {
          caseContent += `${indent.get()}${paramVar}.Body = ${bodyContent}\n`;
          caseContent += `${indent.get()}${paramVar}.ContentType = ${contentType}\n`;
          caseContent += `${indent.get()}${paramVar}.Filename = ${filename}\n`;
        }
      } else if (type.kind === 'slice') {
        if (type.elementType.kind === 'readSeekCloser') {
          imports.add('bytes');
          imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
          assignedValue = `append(${paramVar}, streaming.NopCloser(bytes.NewReader(content)))`;
        } else {
          throw new CodegenError('InternalError', `uhandled multipart parameter array element kind ${type.elementType.kind}`);
        }
      } else if (type.kind === 'encodedBytes') {
        imports.add('encoding/base64');
        caseContent += `${indent.get()}${paramVar}, err = base64.${type.encoding}Encoding.DecodeString(string(content))\n`;
        caseContent += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
      } else {
        throw new CodegenError('InternalError', `uhandled multipart parameter kind ${type.kind}`);
      }

      if (assignedValue) {
        if (!destIsByValue) {
          imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
          assignedValue = `to.Ptr(${assignedValue})`;
        }
        caseContent += `${indent.get()}${paramVar} = ${assignedValue}\n`;
      }

      indent.pop();
      return caseContent;
    };

    for (const param of multipartBodyParams) {
      if (isModelType(param.type)) {
        for (const field of param.type.fields) {
          content += emitCase(field.serializedName, `${param.name}.${field.name}`, field.type, field.byValue);
        }
      } else {
        // for this case we've emitted local vars of the underlying
        // type which is why we pass true for param destIsByValue
        content += emitCase(param.name, param.name, param.type, true);
      }
    }

    content += `${indent.get()}default:\n`;
    content += `${indent.push().get()}return nil, fmt.Errorf("unexpected part %s", fn)\n`;
    content += `${indent.pop().get()}}\n`; // end switch
    content += `${indent.pop().get()}}\n`; // end for
  } else if (formBodyParams.length > 0) {
    for (const param of formBodyParams) {
      content += `${indent.get()}var ${param.name} ${go.getTypeDeclaration(param.type, pkg)}\n`;
    }
    content += `${indent.get()}if err := req.ParseForm(); err != nil {\n${indent.push().get()}return nil, &nonRetriableError{fmt.Errorf("failed parsing form data: %v", err)}\n${indent.pop().get()}}\n`;
    content += `${indent.get()}for key := range req.Form {\n`;
    content += `${indent.push().get()}switch key {\n`;
    for (const param of formBodyParams) {
      content += `${indent.get()}case "${param.formDataName}":\n`;
      let assignedValue: string;
      switch (param.type.kind) {
        case 'constant':
          assignedValue = `${go.getTypeDeclaration(param.type, pkg)}(req.FormValue(key))`;
          break;
        case 'string':
          assignedValue = 'req.FormValue(key)';
          break;
        default:
          throw new CodegenError('InternalError', `uhandled form parameter kind ${param.type.kind}`);
      }
      content += `${indent.push().get()}${param.name} = ${assignedValue}\n`;
      indent.pop();
    }
    content += `${indent.pop().get()}}\n`; // end switch
    content += `${indent.get()}}\n`; // end for
  } else if (partialBodyParams.length > 0) {
    // construct the partial body params type and unmarshal it
    content += `${indent.get()}type partialBodyParams struct {\n`;
    indent.push();
    for (const partialBodyParam of partialBodyParams) {
      content += `${indent.get()}${naming.capitalize(partialBodyParam.name)} ${helpers.star(partialBodyParam.byValue)}${go.getTypeDeclaration(partialBodyParam.type, pkg)} \`json:"${partialBodyParam.serializedName}"\`\n`;
    }
    content += `${indent.pop().get()}}\n`;
    content += `${indent.get()}body, err := server.UnmarshalRequestAs${partialBodyParams[0].format}[partialBodyParams](req)\n`;
    content += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
  }

  const result = parseHeaderPathQueryParams(pkg, method, imports, indent);
  content += result.content;

  // translate each partial body param to its field within the unmarshalled body
  for (const partialBodyParam of partialBodyParams) {
    result.params.set(partialBodyParam.name, `${helpers.star(partialBodyParam.byValue)}body.${naming.capitalize(partialBodyParam.name)}`);
  }

  const apiCall = `:= ${receiverName}.srv.${fixUpMethodName(method)}(${populateApiParams(pkg, method, result.params, imports)})`;
  if (method.kind === 'pageableMethod') {
    content += `resp ${apiCall}\n`;
    return content;
  }
  content += `${indent.get()}respr, errRespr ${apiCall}\n`;
  content += `${indent.get()}if respErr := server.GetError(errRespr, req); respErr != nil {\n`;
  content += `${indent.push().get()}return nil, respErr\n${indent.pop().get()}}\n`;
  return content;
}

function getMethodStatusCodes(method: go.MethodType): Array<number> {
  // NOTE: don't modify the original array!
  const statusCodes = Array.from(method.httpStatusCodes);
  switch (method.kind) {
    case 'lroMethod':
    case 'lroPageableMethod':
      if (!statusCodes.includes(200)) {
        // pollers always include 200 as an acceptible status code so we emulate that here
        statusCodes.unshift(200);
      }
      if (!method.returns.result && !statusCodes.includes(204)) {
        // also include 204 if the LRO doesn't return a body
        statusCodes.push(204);
      }
      break;
  }
  return statusCodes;
}

/**
 * generates the dispatching logic for an LRO server dispatch method
 *
 * @param pkg contains the package contents
 * @param receiverName the name of the receiver for the dispatch method
 * @param method the LRO method for which to emit the dispatch logic
 * @param imports the import manager currently in scope
 * @param indent the indentation helper currently in scope
 * @returns the text for the LRO dispatch logic
 */
function dispatchForLROBody(pkg: go.FakePackage, receiverName: string, method: go.LROMethod | go.LROPageableMethod, imports: ImportManager, indent: helpers.Indentation): string {
  const operationName = fixUpMethodName(method);
  const localVarName = naming.uncapitalize(operationName);
  const operationStateMachine = `${receiverName}.${naming.uncapitalize(operationName)}`;
  let content = `${indent.get()}${localVarName} := ${operationStateMachine}.get(req)\n`;
  content += `${indent.get()}if ${localVarName} == nil {\n`;
  content += dispatchForOperationBody(pkg, receiverName, method, imports, indent);
  indent.push();
  content += `${indent.get()}${localVarName} = &respr\n`;
  content += `${indent.get()}${operationStateMachine}.add(req, ${localVarName})\n`;
  content += `${indent.pop().get()}}\n\n`;

  content += `${indent.get()}resp, err := server.PollerResponderNext(${localVarName}, req)\n`;
  content += `${indent.get()}if err != nil {\n`;
  content += `${indent.push().get()}return nil, err\n`;
  content += `${indent.pop().get()}}\n\n`;

  const formattedStatusCodes = helpers.formatStatusCodes(getMethodStatusCodes(method));
  content += `${indent.get()}if !slices.Contains([]int{${formattedStatusCodes}}, resp.StatusCode) {\n`;
  indent.push();
  content += `${indent.get()}${operationStateMachine}.remove(req)\n`;
  content += `${indent.get()}return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are ${formattedStatusCodes}", resp.StatusCode)}\n`;
  content += `${indent.pop().get()}}\n`;

  content += `${indent.get()}if !server.PollerResponderMore(${localVarName}) {\n`;
  content += `${indent.push().get()}${operationStateMachine}.remove(req)\n`;
  content += `${indent.pop().get()}}\n\n`;
  content += `${indent.get()}return resp, nil\n`;
  return content;
}

/**
 * generates the dispatching logic for a paged server dispatch method
 *
 * @param pkg contains the package contents
 * @param receiverName the name of the receiver for the dispatch method
 * @param method the pageable method for which to emit the dispatch logic
 * @param imports the import manager currently in scope
 * @param indent the indentation helper currently in scope
 * @returns the text for the pageable dispatch logic
 */
function dispatchForPagerBody(pkg: go.FakePackage, receiverName: string, method: go.PageableMethod, imports: ImportManager, indent: helpers.Indentation): string {
  const operationName = fixUpMethodName(method);
  const localVarName = naming.uncapitalize(operationName);
  const operationStateMachine = `${receiverName}.${naming.uncapitalize(operationName)}`;
  let content = `${indent.get()}${localVarName} := ${operationStateMachine}.get(req)\n`;
  content += `${indent.get()}if ${localVarName} == nil {\n`;
  content += dispatchForOperationBody(pkg, receiverName, method, imports, indent);
  indent.push();
  content += `${indent.get()}${localVarName} = &resp\n`;
  content += `${indent.get()}${operationStateMachine}.add(req, ${localVarName})\n`;
  if (method.strategy?.kind === 'nextLink') {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
    content += `${indent.get()}server.PagerResponderInjectNextLinks(${localVarName}, req, func(page *${go.getTypeDeclaration(method.returns, pkg)}, createLink func() string) {\n`;
    content += `${indent.push().get()}page.${helpers.buildNextLinkPath(method.strategy)} = to.Ptr(createLink())\n`;
    content += `${indent.pop().get()}})\n`;
  }
  indent.pop();
  content += `${indent.get()}}\n`; // end if
  content += `${indent.get()}resp, err := server.PagerResponderNext(${localVarName}, req)\n`;
  content += `${indent.get()}if err != nil {\n`;
  content += `${indent.push().get()}return nil, err\n`;
  content += `${indent.pop().get()}}\n`;

  const formattedStatusCodes = helpers.formatStatusCodes(method.httpStatusCodes);
  content += `${indent.get()}if !slices.Contains([]int{${formattedStatusCodes}}, resp.StatusCode) {\n`;
  indent.push();
  content += `${indent.get()}${operationStateMachine}.remove(req)\n`;
  content += `${indent.get()}return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are ${formattedStatusCodes}", resp.StatusCode)}\n`;
  content += `${indent.pop().get()}}\n`;

  content += `${indent.get()}if !server.PagerResponderMore(${localVarName}) {\n`;
  content += `${indent.push().get()}${operationStateMachine}.remove(req)\n`;
  content += `${indent.pop().get()}}\n`;
  content += `${indent.get()}return resp, nil\n`;
  return content;
}

function sanitizeRegexpCaptureGroupName(name: string): string {
  // dash '-' characters are not allowed so replace them with '_'
  return name.replace('-', '_');
}

function createPathParamsRegex(method: go.MethodType, pathParams: Array<go.PathParameter>): string {
  // "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}"
  // each path param will replaced with a regex capture.
  // note that some path params are optional.
  let urlPath = method.httpPath;
  // escape any characters in the path that could be interpreted as regex tokens
  // per RFC3986, these are the pchars that also double as regex tokens
  // . $ * + ()
  urlPath = urlPath.replace(/([.$*+()])/g, '\\$1');
  for (const param of pathParams) {
    const toReplace = `{${param.pathSegment}}`;
    let replaceWith = `(?P<${sanitizeRegexpCaptureGroupName(param.pathSegment)}>[!#&$-;=?-\\[\\]_a-zA-Z0-9~%@]+)`;
    if (param.style === 'optional' || param.style === 'flag') {
      replaceWith += '?';
    }
    urlPath = urlPath.replace(toReplace, replaceWith);
  }
  return urlPath;
}

interface parseResult {
  // contains the param parsing code
  content: string;

  // maps a param name to the var containing the "final" value.
  // only params that required parsing/casting will have an entry.
  params: Map<string, string>;
}

/**
 * parses header/path/query params as required
 *
 * @param pkg contains the package contents
 * @param method the method for which to emit parameter parsing logic
 * @param imports the import manager currently in scope
 * @param indent the indentation helper currently in scope
 * @returns the parsing code and the params that contain the parsed values
 */
function parseHeaderPathQueryParams(pkg: go.FakePackage, method: go.MethodType, imports: ImportManager, indent: helpers.Indentation): parseResult {
  let content = '';
  const paramValues = new Map<string, string>();

  const createLocalVariableName = function (param: go.MethodParameter, suffix: string): string {
    const paramName = `${naming.uncapitalize(param.name)}${suffix}`;
    paramValues.set(param.name, paramName);
    return paramName;
  };

  const emitNumericConversion = function (src: string, type: 'float32' | 'float64' | 'int32' | 'int64'): string {
    imports.add('strconv');
    let precision: '32' | '64' = '32';
    if (type === 'float64' || type === 'int64') {
      precision = '64';
    }
    let parseType: 'Int' | 'Float' = 'Int';
    let base = '10, ';
    if (type === 'float32' || type === 'float64') {
      parseType = 'Float';
      base = '';
    }
    return `strconv.Parse${parseType}(${src}, ${base}${precision})`;
  };

  // track the param groups that need to be instantiated/populated.
  // we track the params separately as it might be a subset of ParameterGroup.params
  const paramGroups = new Map<go.ParameterGroup, Array<go.MethodParameter>>();

  for (const param of consolidateHostParams(method.parameters)) {
    if (param.location === 'client' || go.isLiteralParameter(param.style)) {
      // client params and parameter literals aren't passed to APIs
      continue;
    }
    if (param.kind === 'resumeTokenParam') {
      // skip the ResumeToken param as we don't send that back to the caller
      continue;
    }

    // NOTE: param group check must happen before skipping body params.
    // this is to handle the case where the body param is grouped/optional
    if (param.group) {
      let params = paramGroups.get(param.group);
      if (!params) {
        params = new Array<go.MethodParameter>();
        paramGroups.set(param.group, params);
      }
      params.push(param);
    }

    switch (param.kind) {
      case 'bodyParam':
      case 'formBodyCollectionParam':
      case 'formBodyScalarParam':
      case 'multipartFormBodyParam':
      case 'partialBodyParam':
        // body params will be unmarshalled, no need for parsing.
        continue;
    }

    // paramValue is initialized with the "raw" source value.
    // e.g. getHeaderValue(...), qp.Get("foo") etc
    // since path/query params need to be unescaped, the value
    // of paramValue will be updated with the var name that
    // contains the unescaped value.
    let paramValue = getRawParamValue(param);

    // path params are escaped, so we need to unescape them first.
    if (go.isPathParameter(param)) {
      imports.add('net/url');
      let paramVar = createLocalVariableName(param, 'Unescaped');
      if (go.isRequiredParameter(param.style) && param.type.kind === 'constant' && param.type.type === 'string') {
        // for string-based enums, we perform the conversion as part of unescaping
        requiredHelpers.parseWithCast = true;
        paramVar = createLocalVariableName(param, 'Param');
        content += `${indent.get()}${paramVar}, err := parseWithCast(${paramValue}, func (v string) (${go.getTypeDeclaration(param.type, pkg)}, error) {\n`;
        content += `${indent.push().get()}p, unescapeErr := url.PathUnescape(v)\n`;
        content += `${indent.get()}if unescapeErr != nil {\n${indent.push().get()}return "", unescapeErr\n${indent.pop().get()}}\n`;
        content += `${indent.get()}return ${go.getTypeDeclaration(param.type, pkg)}(p), nil\n${indent.pop().get()}})\n`;
      } else {
        if (go.isRequiredParameter(param.style) && (param.type.kind === 'string' || (param.type.kind === 'slice' && param.type.elementType.kind === 'string'))) {
          // by convention, if the value is in its "final form" (i.e. no parsing required)
          // then its var is to have the "Param" suffix. the only case is string, everything
          // else requires some amount of parsing/conversion.
          paramVar = createLocalVariableName(param, 'Param');
        }
        content += `${indent.get()}${paramVar}, err := url.PathUnescape(${paramValue})\n`;
      }
      content += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
      paramValue = paramVar;
    }

    // parse params as required
    if (param.kind === 'headerCollectionParam' || param.kind === 'pathCollectionParam' || param.kind === 'queryCollectionParam') {
      // any element type other than string will require some form of conversion/parsing
      if (param.type.elementType.kind !== 'string') {
        if (param.collectionFormat !== 'multi') {
          requiredHelpers.splitHelper = true;
          const elementsParam = createLocalVariableName(param, 'Elements');
          content += `${indent.get()}${elementsParam} := splitHelper(${paramValue}, "${helpers.getDelimiterForCollectionFormat(param.collectionFormat)}")\n`;
          paramValue = elementsParam;
        }

        const paramVar = createLocalVariableName(param, 'Param');
        let elementFormat: go.ScalarType | go.TimeFormat | go.BytesEncoding | 'string';
        switch (param.type.elementType.kind) {
          case 'constant':
          case 'scalar':
            elementFormat = param.type.elementType.type;
            break;
          case 'encodedBytes':
            elementFormat = param.type.elementType.encoding;
            break;
          case 'time':
            elementFormat = param.type.elementType.format;
            break;
          default:
            throw new CodegenError('InternalError', `unhandled element kind ${param.type.elementType.kind}`);
        }

        const toType = go.getTypeDeclaration(param.type.elementType, pkg);
        content += `${indent.get()}${paramVar} := make([]${toType}, len(${paramValue}))\n`;
        content += `${indent.get()}for i := 0; i < len(${paramValue}); i++ {\n`;
        indent.push();
        let fromVar: string;

        // TODO: consolidate with non-collection parsing code
        if (elementFormat === 'bool') {
          imports.add('strconv');
          fromVar = 'parsedBool';
          content += `${indent.get()}${fromVar}, parseErr := strconv.ParseBool(${paramValue}[i])\n`;
          content += `${indent.get()}if parseErr != nil {\n${indent.push().get()}return nil, parseErr\n${indent.pop().get()}}\n`;
        } else if (elementFormat === 'float32' || elementFormat === 'float64' || elementFormat === 'int32' || elementFormat === 'int64') {
          fromVar = `parsed${naming.capitalize(elementFormat)}`;
          content += `${indent.get()}${fromVar}, parseErr := ${emitNumericConversion(`${paramValue}[i]`, elementFormat)}\n`;
          content += `${indent.get()}if parseErr != nil {\n${indent.push().get()}return nil, parseErr\n${indent.pop().get()}}\n`;
        } else if (elementFormat === 'string') {
          // we're casting an enum string value to its const type
          // TODO: what about enums that aren't strings?
          fromVar = `${paramValue}[i]`;
        } else if (elementFormat === 'Std' || elementFormat === 'URL') {
          imports.add('encoding/base64');
          fromVar = `parsed${naming.capitalize(elementFormat)}`;
          content += `${indent.get()}${fromVar}, parseErr := base64.${elementFormat}Encoding.DecodeString(${paramValue}[i])\n`;
          content += `${indent.get()}if parseErr != nil {\n${indent.push().get()}return nil, parseErr\n${indent.pop().get()}}\n`;
        } else if (elementFormat === 'RFC1123' || elementFormat === 'RFC3339' || elementFormat === 'Unix') {
          imports.add('time');
          fromVar = `parsed${naming.capitalize(elementFormat)}`;
          if (elementFormat === 'Unix') {
            imports.add('strconv');
            content += `${indent.get()}p, parseErr := strconv.ParseInt(${paramValue}[i], 10, 64)\n`;
            content += `${indent.get()}if parseErr != nil {\n${indent.push().get()}return nil, parseErr\n${indent.pop().get()}}\n`;
            content += `${indent.get()}${fromVar} := time.Unix(p, 0).UTC()\n`;
          } else {
            let format = 'time.RFC3339Nano';
            if (elementFormat === 'RFC1123') {
              format = 'time.RFC1123';
            }
            content += `${indent.get()}${fromVar}, parseErr := time.Parse(${format}, ${paramValue}[i])\n`;
            content += `${indent.get()}if parseErr != nil {\n${indent.push().get()}return nil, parseErr\n${indent.pop().get()}}\n`;
          }
        } else {
          throw new CodegenError('InternalError', `unhandled element format ${elementFormat}`);
        }
        // TODO: remove cast in some cases
        content += `${indent.get()}${paramVar}[i] = ${toType}(${fromVar})\n${indent.pop().get()}}\n`;
      } else if (!go.isRequiredParameter(param.style) && param.collectionFormat !== 'multi') {
        // for slices of strings that are required, the call to splitHelper(...) is inlined into
        // the invocation of the fake e.g. srv.FakeFunc(splitHelper...). but if it's optional, we
        // need to create a local first which will later be copied into the optional param group.
        requiredHelpers.splitHelper = true;
        content += `${indent.get()}${createLocalVariableName(param, 'Param')} := splitHelper(${paramValue}, "${helpers.getDelimiterForCollectionFormat(param.collectionFormat)}")\n`;
      }
    } else if (param.type.kind === 'scalar' && param.type.type === 'bool') {
      imports.add('strconv');
      let from = `strconv.ParseBool(${paramValue})`;
      if (!go.isRequiredParameter(param.style)) {
        requiredHelpers.parseOptional = true;
        from = `parseOptional(${paramValue}, strconv.ParseBool)`;
      }
      content += `${indent.get()}${createLocalVariableName(param, 'Param')}, err := ${from}\n`;
      content += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
    } else if (param.type.kind === 'encodedBytes') {
      imports.add('encoding/base64');
      content += `${indent.get()}${createLocalVariableName(param, 'Param')}, err := base64.${param.type.encoding}Encoding.DecodeString(${paramValue})\n`;
      content += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
    } else if (param.type.kind === 'time') {
      const formatMap: Record<string, string> = {
        PlainDate: helpers.plainDateFormat,
        PlainTime: helpers.plainTimeFormat,
        RFC1123: helpers.RFC1123Format,
        RFC3339: helpers.RFC3339Format,
      };
      imports.add('time');
      if (param.type.format in formatMap) {
        const format = formatMap[param.type.format];
        let from = `time.Parse(${format}, ${paramValue})`;
        if (!go.isRequiredParameter(param.style)) {
          requiredHelpers.parseOptional = true;
          from = `parseOptional(${paramValue}, func(v string) (time.Time, error) { return time.Parse(${format}, v) })`;
        }
        content += `${indent.get()}${createLocalVariableName(param, 'Param')}, err := ${from}\n`;
        content += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
      } else {
        imports.add('strconv');
        let parser: string;
        if (!go.isRequiredParameter(param.style)) {
          requiredHelpers.parseOptional = true;
          parser = 'parseOptional';
        } else {
          requiredHelpers.parseWithCast = true;
          parser = 'parseWithCast';
        }
        content += `${indent.get()}${createLocalVariableName(param, 'Param')}, err := ${parser}(${paramValue}, func (v string) (time.Time, error) {\n`;
        content += `${indent.push().get()}p, parseErr := strconv.ParseInt(v, 10, 64)\n`;
        content += `${indent.get()}if parseErr != nil {\n${indent.push().get()}return time.Time{}, parseErr\n${indent.pop().get()}}\n`;
        content += `${indent.get()}return time.Unix(p, 0).UTC(), nil\n${indent.pop().get()}})\n`;
        content += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
      }
    } else if (param.type.kind === 'scalar' && (param.type.type === 'float32' || param.type.type === 'float64' || param.type.type === 'int32' || param.type.type === 'int64')) {
      let parser: string;
      if (!go.isRequiredParameter(param.style)) {
        requiredHelpers.parseOptional = true;
        parser = 'parseOptional';
      } else {
        requiredHelpers.parseWithCast = true;
        parser = 'parseWithCast';
      }
      if (param.type.type === 'float32' || param.type.type === 'int32' || !go.isRequiredParameter(param.style)) {
        content += `${indent.get()}${createLocalVariableName(param, 'Param')}, err := ${parser}(${paramValue}, func(v string) (${param.type.type}, error) {\n`;
        content += `${indent.push().get()}p, parseErr := ${emitNumericConversion('v', param.type.type)}\n`;
        content += `${indent.get()}if parseErr != nil {\n${indent.push().get()}return 0, parseErr\n${indent.pop().get()}}\n`;
        let result = 'p';
        if (param.type.type === 'float32' || param.type.type === 'int32') {
          result = `${param.type.type}(${result})`;
        }
        content += `${indent.get()}return ${result}, nil\n${indent.pop().get()}})\n`;
      } else {
        content += `${indent.get()}${createLocalVariableName(param, 'Param')}, err := ${emitNumericConversion(paramValue, param.type.type)}\n`;
      }
      content += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
    } else if (param.kind === 'headerMapParam') {
      imports.add('strings');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
      const localVar = createLocalVariableName(param, 'Param');
      content += `${indent.get()}var ${localVar} map[string]*string\n`;
      content += `${indent.get()}for hh := range ${paramValue} {\n`;
      const headerPrefix = param.headerName;
      requiredHelpers.getHeaderValue = true;
      content += `${indent.push().get()}if len(hh) > len("${headerPrefix}") && strings.EqualFold(hh[:len("x-ms-meta-")], "${headerPrefix}") {\n`;
      content += `${indent.push().get()}if ${localVar} == nil {\n${indent.push().get()}${localVar} = map[string]*string{}\n${indent.pop().get()}}\n`;
      content += `${indent.get()}${localVar}[hh[len("${headerPrefix}"):]] = to.Ptr(getHeaderValue(req.Header, hh))\n`;
      content += `${indent.pop().get()}}\n${indent.pop().get()}}\n`;
    } else if (param.type.kind === 'constant' && param.type.type !== 'string') {
      let parseHelper: string;
      if (!go.isRequiredParameter(param.style)) {
        requiredHelpers.parseOptional = true;
        parseHelper = 'parseOptional';
      } else {
        requiredHelpers.parseWithCast = true;
        parseHelper = 'parseWithCast';
      }
      let parse: string;
      let zeroValue: string;
      if (param.type.type === 'bool') {
        imports.add('strconv');
        parse = 'strconv.ParseBool(v)';
        zeroValue = 'false';
      } else {
        // emitNumericConversion adds the necessary import of strconv
        parse = emitNumericConversion('v', param.type.type);
        zeroValue = '0';
      }
      const toConstType = go.getTypeDeclaration(param.type, pkg);
      content += `${indent.get()}${createLocalVariableName(param, 'Param')}, err := ${parseHelper}(${paramValue}, func(v string) (${toConstType}, error) {\n`;
      content += `${indent.push().get()}p, parseErr := ${parse}\n`;
      content += `${indent.get()}if parseErr != nil {\n${indent.push().get()}return ${zeroValue}, parseErr\n${indent.pop().get()}}\n`;
      content += `${indent.get()}return ${toConstType}(p), nil\n${indent.pop().get()}})\n`;
      content += `${indent.get()}if err != nil {\n${indent.push().get()}return nil, err\n${indent.pop().get()}}\n`;
    } else if (!go.isRequiredParameter(param.style)) {
      // we check this last as it's a superset of the previous conditions
      requiredHelpers.getOptional = true;
      if (param.type.kind === 'constant' || param.type.kind === 'etag') {
        imports.addForType(param.type);
        paramValue = `${go.getTypeDeclaration(param.type, pkg)}(${paramValue})`;
      }
      content += `${indent.get()}${createLocalVariableName(param, 'Param')} := getOptional(${paramValue})\n`;
    }
  }

  // create the param groups and populate their values
  for (const paramGroup of paramGroups.keys()) {
    if (paramGroup.required) {
      content += `${indent.get()}${naming.uncapitalize(paramGroup.name)} := ${go.getTypeDeclaration(paramGroup, pkg)}{\n`;
      const params = paramGroups.get(paramGroup);
      if (params) {
        indent.push();
        for (const param of params) {
          content += `${indent.get()}${naming.capitalize(param.name)}: ${getFinalParamValue(pkg, param, paramValues)},\n`;
        }
        indent.pop();
      }
      content += `${indent.get()}}\n`;
    } else {
      content += `${indent.get()}var ${naming.uncapitalize(paramGroup.name)} *${go.getTypeDeclaration(paramGroup, pkg)}\n`;
      const params = paramGroups.get(paramGroup);
      const paramNilCheck = new Array<string>();
      if (params) {
        for (const param of params) {
          // check array before body in case the body is just an array
          if (param.type.kind === 'slice') {
            paramNilCheck.push(`len(${getFinalParamValue(pkg, param, paramValues)}) > 0`);
          } else if (param.kind === 'bodyParam') {
            if (param.bodyFormat === 'binary') {
              imports.add('io');
              paramNilCheck.push('req.Body != nil');
            } else {
              imports.add('reflect');
              paramNilCheck.push('!reflect.ValueOf(body).IsZero()');
            }
          } else if (go.isFormBodyParameter(param) || param.kind === 'multipartFormBodyParam') {
            imports.add('reflect');
            paramNilCheck.push(`!reflect.ValueOf(${param.name}).IsZero()`);
          } else {
            paramNilCheck.push(`${getFinalParamValue(pkg, param, paramValues)} != nil`);
          }
        }
      }
      content += `${indent.get()}if ${paramNilCheck.join(' || ')} {\n`;
      content += `${indent.push().get()}${naming.uncapitalize(paramGroup.name)} = &${go.getTypeDeclaration(paramGroup, pkg)}{\n`;
      if (params) {
        indent.push();
        for (const param of params) {
          let byRef = '&';
          if (param.byValue || (!go.isRequiredParameter(param.style) && param.kind !== 'bodyParam' && !go.isFormBodyParameter(param) && param.kind !== 'multipartFormBodyParam')) {
            byRef = '';
          }
          content += `${indent.get()}${naming.capitalize(param.name)}: ${byRef}${getFinalParamValue(pkg, param, paramValues)},\n`;
        }
        indent.pop();
      }
      content += `${indent.get()}}\n`;
      content += `${indent.pop().get()}}\n`;
    }
  }

  return {
    content: content,
    params: paramValues,
  };
}

/**
 * generates the code to populate the method parameters that get passed to the fake
 *
 * @param pkg contains the package contents
 * @param method the method to be called with the parsed parameters
 * @param paramValues maps a parameter name to the value to be passed to the fake
 * @param imports the import manager currently in scope
 * @returns the text for the parameters to be passed to the fake
 */
function populateApiParams(pkg: go.FakePackage, method: go.MethodType, paramValues: Map<string, string>, imports: ImportManager): string {
  // FooOperation(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], qp.Get("api-version"), nil)
  // this assumes that our caller has created matches and qp as required
  const params = new Array<string>();

  // for non-paged APIs, first param is always the context. use the one
  // from the HTTP request. be careful to properly handle paged LROs
  if (method.kind !== 'pageableMethod') {
    params.push('req.Context()');
  }

  // now create the API call sig
  for (const param of helpers.getMethodParameters(method, consolidateHostParams)) {
    if (param.kind === 'paramGroup') {
      if (param.groupName === method.optionalParamsGroup.groupName) {
        // this is the optional params type. in some cases we just pass nil
        const countParams = param.params.filter((each: go.MethodParameter) => each.kind !== 'resumeTokenParam').length;
        if (countParams === 0) {
          // if the options param is empty or only contains the resume token param just pass nil
          params.push('nil');
          continue;
        }
      }
      // by convention, for param groups, the param parsing code
      // creates a local var with the name of the param
      params.push(naming.uncapitalize(param.name));
      continue;
    }
    imports.addForType(param.type);
    params.push(getFinalParamValue(pkg, param, paramValues));
  }

  return params.join(', ');
}

// getRawParamValue returns the "raw" value for the specified parameter.
// depending on the type, the value might require parsing before it can be passed to the fake.
function getRawParamValue(param: go.MethodParameter): string {
  switch (param.kind) {
    case 'bodyParam':
      if (param.bodyFormat === 'binary') {
        return 'req.Body.(io.ReadSeekCloser)';
      }
      // JSON/XML/text bodies have been deserialized into a local named body
      return 'body';
    case 'formBodyCollectionParam':
    case 'formBodyScalarParam':
    case 'multipartFormBodyParam':
    case 'partialBodyParam':
      // multipart form data values have been read and assigned
      // to local params with the same name.
      return param.name;
    case 'headerCollectionParam':
    case 'headerScalarParam':
      // use req
      requiredHelpers.getHeaderValue = true;
      return `getHeaderValue(req.Header, "${param.headerName}")`;
    case 'headerMapParam':
      return 'req.Header';
    case 'pathCollectionParam':
    case 'pathScalarParam':
      // path params are in the matches slice
      return `matches[regex.SubexpIndex("${sanitizeRegexpCaptureGroupName(param.pathSegment)}")]`;
    case 'queryCollectionParam':
    case 'queryScalarParam':
      // use qp
      if (param.kind === 'queryCollectionParam' && param.collectionFormat === 'multi') {
        return `qp["${param.queryParameter}"]`;
      }
      return `qp.Get("${param.queryParameter}")`;
    case 'uriParam':
      return 'req.URL.Host';
    default:
      throw new CodegenError('InternalError', `unhandled parameter ${param.name}`);
  }
}

/**
 * returns the final value of param to be passed to the fake.
 * this is usually the value in paramValues but can be slightly
 * different for some cases.
 *
 * @param pkg the contents of the package
 * @param param the parameter being evaluated
 * @param paramValues maps a parameter name to the value to be passed to the fake
 * @returns the value to pass for the provided parameter
 */
function getFinalParamValue(pkg: go.FakePackage, param: go.MethodParameter, paramValues: Map<string, string>): string {
  let paramValue = paramValues.get(param.name);
  if (!paramValue) {
    // the param didn't require parsing so the "raw" value can be used
    paramValue = getRawParamValue(param);
  }

  // there are a few corner-cases that require some fix-ups

  if ((param.kind === 'bodyParam' || go.isFormBodyParameter(param) || param.kind === 'multipartFormBodyParam') && param.type.kind === 'time') {
    // time types in the body have been unmarshalled into our time helpers thus require a cast to time.Time
    return `time.Time(${paramValue})`;
  } else if (go.isRequiredParameter(param.style)) {
    // optional params are always in their "final" form
    if (param.kind === 'headerCollectionParam' || param.kind === 'pathCollectionParam' || param.kind === 'queryCollectionParam') {
      // for required params that are collections of strings, we split them inline.
      // not necessary for optional params as they're already in slice format.
      if (param.collectionFormat !== 'multi' && param.type.elementType.kind === 'string') {
        requiredHelpers.splitHelper = true;
        return `splitHelper(${paramValue}, "${helpers.getDelimiterForCollectionFormat(param.collectionFormat)}")`;
      }
    } else if ((go.isHeaderParameter(param) || go.isQueryParameter(param)) && param.type.kind === 'constant' && param.type.type === 'string') {
      // query params from req.URL.Query() are already decoded, so like headers we cast required, string-based enums inline
      return `${go.getTypeDeclaration(param.type, pkg)}(${paramValue})`;
    }
  } else if (param.kind === 'partialBodyParam') {
    // use the value from the unmarshaled, intermediate struct type
    return `body.${naming.capitalize(param.name)}`;
  }

  return paramValue;
}

// takes multiple host parameters and consolidates them into a single "host" parameter.
// this is necessary as there's no way to rehydrate multiple host parameters.
// e.g. host := "{vault}{secret}{dnsSuffix}" becomes http://contososecret.com
// there's no way to reliably split the host back up into its constituent parameters.
// so we just pass the full value as a single host parameter.
function consolidateHostParams(params: Array<go.MethodParameter>): Array<go.MethodParameter> {
  if (!params.find((each: go.MethodParameter) => each.kind === 'uriParam')) {
    // no host params
    return params;
  }

  // consolidate multiple host params into a single "host" param
  const consolidatedParams = new Array<go.MethodParameter>();
  let hostParamAdded = false;
  for (const param of params) {
    if (param.kind !== 'uriParam') {
      consolidatedParams.push(param);
    } else if (!hostParamAdded) {
      consolidatedParams.push(param);
      hostParamAdded = true;
    }
  }

  return consolidatedParams;
}

/**
 * copied from generator/operations.ts but with a slight tweak to consolidate host parameters
 *
 * @param pkg the contents of the package
 * @param method the method for which to generate the parameter signature
 * @param imports the import manager currently in scope
 * @returns the text for the method's parameter signature
 */
function getAPIParametersSig(pkg: go.FakePackage, method: go.MethodType, imports: ImportManager): string {
  const methodParams = helpers.getMethodParameters(method, consolidateHostParams);
  const params = new Array<string>();
  if (method.kind !== 'pageableMethod') {
    imports.add('context');
    params.push('ctx context.Context');
  }
  for (const methodParam of methodParams) {
    let paramName = naming.uncapitalize(methodParam.name);
    if (methodParam.kind === 'uriParam') {
      paramName = 'host';
    }
    params.push(`${paramName} ${helpers.formatParameterTypeName(pkg, methodParam)}`);
  }
  return params.join(', ');
}

// copied from generator/helpers.ts but without the XML-specific stuff
function getResultFieldName(result: go.AnyResult | go.BinaryResult | go.MonomorphicResult | go.PolymorphicResult | go.ModelResult): string {
  switch (result.kind) {
    case 'anyResult':
      return result.fieldName;
    case 'modelResult':
      return result.modelType.name;
    case 'polymorphicResult':
      return result.interface.name;
    default:
      return result.fieldName;
  }
}
