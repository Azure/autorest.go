/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Client, GoCodeModel, Method, PageableMethod, Parameter, ParameterGroup, isBinaryResult, isConstantType, isHeadAsBooleanResult, isHeaderParameter, isInterfaceType, isModelResult, isMonomorphicResult, isMultipartFormBodyParameter, isPolymorphicResult, isPrimitiveType, isResumeTokenParameter, isStandardType } from '../../gocodemodel/gocodemodel';
import { AnyResult, BinaryResult, BodyParameter, MonomorphicResult, PolymorphicResult, ModelResult, isAnyResult, isBytesType, isTimeType } from '../../gocodemodel/gocodemodel';
import { getTypeDeclaration, isBodyParameter, isFormBodyParameter, isLiteralValue, isLROMethod, isPageableMethod, isPathParameter, isQueryParameter, isSliceType, isURIParameter } from '../../gocodemodel/gocodemodel';
import { capitalize, uncapitalize } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import { contentPreamble, formatLiteralValue, formatParameterTypeName, formatStatusCode, formatStatusCodes, formatValue, getMethodParameters, getParentImport, isParameter, isParameterGroup, isLiteralParameter, isRequiredParameter } from '../helpers';
import { fixUpMethodName } from '../operations';
import { ImportManager } from '../imports';

// represents the generated content for an operation group
export class OperationGroupContent {
  readonly name: string;
  readonly content: string;

  constructor(name: string, content: string) {
    this.name = name;
    this.content = content;
  }
}

export async function generateServers(codeModel: GoCodeModel): Promise<Array<OperationGroupContent>> {
  const operations = new Array<OperationGroupContent>();
  const clientPkg = codeModel.packageName;
  for (const client of values(codeModel.clients)) {
    // the list of packages to import
    const imports = new ImportManager();
    imports.add(getParentImport(codeModel));

    // add standard imports
    imports.add('errors');
    imports.add('fmt');
    imports.add('net/http');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake', 'azfake');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');

    // for the fake server, we use the suffix Server instead of Client
    const serverName = capitalize(client.clientName.replace(/[C|c]lient$/, 'Server'));

    let content: string;
    content = `// ${serverName} is a fake server for instances of the ${clientPkg}.${client.clientName} type.\n`;
    content += `type ${serverName} struct{\n`;

    // we might remove some operations from the list
    const finalMethods = new Array<Method>();
    let countLROs = 0;
    let countPagers = 0;

    for (const method of values(client.methods)) {
      let serverResponse: string;

      if (isLROMethod(method)) {
        let respType = `${clientPkg}.${method.responseEnvelope.name}`;
        if (isPageableMethod(method)) {
          respType = `azfake.PagerResponder[${clientPkg}.${method.responseEnvelope.name}]`;
        }
        serverResponse = `resp azfake.PollerResponder[${respType}], errResp azfake.ErrorResponder`;
      } else if (isPageableMethod(method)) {
        serverResponse = `resp azfake.PagerResponder[${clientPkg}.${method.responseEnvelope.name}]`;
      } else {
        serverResponse = `resp azfake.Responder[${clientPkg}.${method.responseEnvelope.name}], errResp azfake.ErrorResponder`;
      }

      const operationName = fixUpMethodName(method);
      content += `\t// ${operationName} is the fake for method ${client.clientName}.${operationName}\n`;
      const successCodes = new Array<string>();
      if (method.responseEnvelope.result && isAnyResult(method.responseEnvelope.result)) {
        for (const httpStatus of values(method.httpStatusCodes)) {
          const result = method.responseEnvelope.result.httpStatusCodeType[httpStatus];
          if (!result) {
            // the operation contains a mix of schemas and non-schema responses
            successCodes.push(`${formatStatusCode(httpStatus)} (no return type)`);
            continue;
          }
          successCodes.push(`${formatStatusCode(httpStatus)} (returns ${getTypeDeclaration(result, clientPkg)})`);
        }
        content += '\t// HTTP status codes to indicate success:\n';
        for (const successCode of successCodes) {
          content += `\t//   - ${successCode}\n`;
        }
      } else {
        for (const statusCode of values(method.httpStatusCodes)) {
          successCodes.push(`${formatStatusCode(statusCode)}`);
        }
        content += `\t// HTTP status codes to indicate success: ${successCodes.join(', ')}\n`;
      }
      content += `\t${operationName} func(${getAPIParametersSig(method, imports, clientPkg)}) (${serverResponse})\n\n`;
      finalMethods.push(method);
      if (isLROMethod(method)) {
        ++countLROs;
      } else if (isPageableMethod(method)) {
        ++countPagers;
      }
    }

    content += '}\n\n';

    ///////////////////////////////////////////////////////////////////////////

    const serverTransport = `${serverName}Transport`;

    content += `// New${serverTransport} creates a new instance of ${serverTransport} with the provided implementation.\n`;
    content += `// The returned ${serverTransport} instance is connected to an instance of ${clientPkg}.${client.clientName} via the\n`;
    content += '// azcore.ClientOptions.Transporter field in the client\'s constructor parameters.\n';
    content += `func New${serverTransport}(srv *${serverName}) *${serverTransport} {\n`;
    if (countLROs === 0 && countPagers === 0) {
      content += `\treturn &${serverTransport}{srv: srv}\n}\n\n`;
    } else {
      content += `\treturn &${serverTransport}{\n\t\tsrv: srv,\n`;
      for (const method of values(finalMethods)) {
        let respType = `${clientPkg}.${method.responseEnvelope.name}`;
        if (isLROMethod(method)) {
          if (isPageableMethod(method)) {
            respType = `azfake.PagerResponder[${clientPkg}.${method.responseEnvelope.name}]`;
          }
          content += `\t\t${uncapitalize(fixUpMethodName(method))}: newTracker[azfake.PollerResponder[${respType}]](),\n`;
        } else if (isPageableMethod(method)) {
          content += `\t\t${uncapitalize(fixUpMethodName(method))}: newTracker[azfake.PagerResponder[${respType}]](),\n`;
        }
      }
      content += '\t}\n}\n\n';
    }

    content += `// ${serverTransport} connects instances of ${clientPkg}.${client.clientName} to instances of ${serverName}.\n`;
    content += `// Don't use this type directly, use New${serverTransport} instead.\n`;
    content += `type ${serverTransport} struct {\n`;
    content += `\tsrv *${serverName}\n`;
    for (const method of values(finalMethods)) {
      // create state machines for any pager/poller operations
      let respType = `${clientPkg}.${method.responseEnvelope.name}`;
      if (isLROMethod(method)) {
        if (isPageableMethod(method)) {
          respType = `azfake.PagerResponder[${clientPkg}.${method.responseEnvelope.name}]`;
        }
        content +=`\t${uncapitalize(fixUpMethodName(method))} *tracker[azfake.PollerResponder[${respType}]]\n`;
      } else if (isPageableMethod(method)) {
        content += `\t${uncapitalize(fixUpMethodName(method))} *tracker[azfake.PagerResponder[${clientPkg}.${method.responseEnvelope.name}]]\n`;
      }
    }
    content += '}\n\n';

    content += generateServerTransportMethods(clientPkg, serverTransport, client, finalMethods, imports);

    ///////////////////////////////////////////////////////////////////////////

    // stitch everything together
    let text = contentPreamble(codeModel, 'fake');
    text += imports.text();
    text += content;
    operations.push(new OperationGroupContent(serverName, text));
  }
  return operations;
}

function generateServerTransportMethods(clientPkg: string, serverTransport: string, client: Client, finalMethods: Array<Method>, imports: ImportManager): string {
  const receiverName = serverTransport[0].toLowerCase();
  let content = `// Do implements the policy.Transporter interface for ${serverTransport}.\n`;
  content += `func (${receiverName} *${serverTransport}) Do(req *http.Request) (*http.Response, error) {\n`;
  content += '\trawMethod := req.Context().Value(runtime.CtxAPINameKey{})\n';
  content += '\tmethod, ok := rawMethod.(string)\n';
  content += '\tif !ok {\n\t\treturn nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}\n\t}\n\n';
  content += '\tvar resp *http.Response\n';
  content += '\tvar err error\n\n';
  content += '\tswitch method {\n';

  for (const method of values(finalMethods)) {
    const operationName = fixUpMethodName(method);
    content += `\tcase "${client.clientName}.${operationName}":\n`;
    content += `\t\tresp, err = ${receiverName}.dispatch${operationName}(req)\n`;
  }
  content += '\tdefault:\n\t\terr = fmt.Errorf("unhandled API %s", method)\n';

  content += '\t}\n\n'; // end switch
  content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n\n';
  content += '\treturn resp, nil\n}\n\n';

  ///////////////////////////////////////////////////////////////////////////

  for (const method of values(finalMethods)) {
    content += `func (${receiverName} *${serverTransport}) dispatch${fixUpMethodName(method)}(req *http.Request) (*http.Response, error) {\n`;
    content += `\tif ${receiverName}.srv.${fixUpMethodName(method)} == nil {\n`;
    content += `\t\treturn nil, &nonRetriableError{errors.New("fake for method ${fixUpMethodName(method)} not implemented")}\n\t}\n`;

    if (isLROMethod(method)) {
      // must check LRO before pager as you can have paged LROs
      content += dispatchForLROBody(clientPkg, receiverName, method, imports);
    } else if (isPageableMethod(method)) {
      content += dispatchForPagerBody(clientPkg, receiverName, method, imports);
    } else {
      content += dispatchForOperationBody(clientPkg, receiverName, method, imports);
      content += '\trespContent := server.GetResponseContent(respr)\n';
      const formattedStatusCodes = formatStatusCodes(method.httpStatusCodes);
      content += `\tif !contains([]int{${formattedStatusCodes}}, respContent.HTTPStatus) {\n`;
      content += `\t\treturn nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are ${formattedStatusCodes}", respContent.HTTPStatus)}\n\t}\n`;
      if (!method.responseEnvelope.result || isHeadAsBooleanResult(method.responseEnvelope.result)) {
        content += '\tresp, err := server.NewResponse(respContent, req, nil)\n';
      } else if (isAnyResult(method.responseEnvelope.result)) {
        content += `\tresp, err := server.MarshalResponseAs${method.responseEnvelope.result.format}(respContent, server.GetResponse(respr).${getResultFieldName(method.responseEnvelope.result)}, req)\n`;
      } else if (isBinaryResult(method.responseEnvelope.result)) {
        content += '\tresp, err := server.NewResponse(respContent, req, &server.ResponseOptions{\n';
        content += `\t\tBody: server.GetResponse(respr).${getResultFieldName(method.responseEnvelope.result)},\n`;
        content += '\t\tContentType: "application/octet-stream",\n';
        content += '\t})\n';
      } else if (isMonomorphicResult(method.responseEnvelope.result)) {
        if (isBytesType(method.responseEnvelope.result.monomorphicType)) {
          const encoding = method.responseEnvelope.result.monomorphicType.encoding;
          content += `\tresp, err := server.MarshalResponseAsByteArray(respContent, server.GetResponse(respr).${getResultFieldName(method.responseEnvelope.result)}, runtime.Base64${encoding}Format, req)\n`;
        } else if (isSliceType(method.responseEnvelope.result.monomorphicType) && method.responseEnvelope.result.monomorphicType.rawJSONAsBytes) {
          imports.add('bytes');
          imports.add('io');
          content += '\tresp, err := server.NewResponse(respContent, req, &server.ResponseOptions{\n';
          content += '\t\tBody: io.NopCloser(bytes.NewReader(server.GetResponse(respr).RawJSON)),\n';
          content += '\t\tContentType: "application/json",\n\t})\n';
        } else {
          let respField = `.${getResultFieldName(method.responseEnvelope.result)}`;
          if (method.responseEnvelope.result.format === 'XML' && isSliceType(method.responseEnvelope.result.monomorphicType)) {
            // for XML array responses we use the response type directly as it has the necessary XML tag for proper marshalling
            respField = '';
          }
          let responseField = `server.GetResponse(respr)${respField}`;
          if (isTimeType(method.responseEnvelope.result.monomorphicType)) {
            responseField = `(*${method.responseEnvelope.result.monomorphicType.dateTimeFormat})(${responseField})`;
          }
          content += `\tresp, err := server.MarshalResponseAs${method.responseEnvelope.result.format}(respContent, ${responseField}, req)\n`;
        }
      } else if (isModelResult(method.responseEnvelope.result) || isPolymorphicResult(method.responseEnvelope.result)) {
        const respField = `.${getResultFieldName(method.responseEnvelope.result)}`;
        const responseField = `server.GetResponse(respr)${respField}`;
        content += `\tresp, err := server.MarshalResponseAs${method.responseEnvelope.result.format}(respContent, ${responseField}, req)\n`;
      }

      content += '\tif err != nil {\t\treturn nil, err\n\t}\n';

      // propagate any header response values into the *http.Response
      for (const header of values(method.responseEnvelope.headers)) {
        if (header.collectionPrefix) {
          content += `\tfor k, v := range server.GetResponse(respr).${header.fieldName} {\n`;
          content += '\t\tif v != nil {\n';
          content += `\t\t\tresp.Header.Set("${header.collectionPrefix}"+k, *v)\n`;
          content += '\t\t}\n';
          content += '\t}\n';
        } else {
          content += `\tif val := server.GetResponse(respr).${header.fieldName}; val != nil {\n`;
          content += `\t\tresp.Header.Set("${header.headerName}", ${formatValue('val', header.type, imports, true)})\n\t}\n`;
        }
      }

      content += '\treturn resp, nil\n';
    }
    content += '}\n\n';
  }

  return content;
}

function dispatchForOperationBody(clientPkg: string, receiverName: string, method: Method, imports: ImportManager): string {
  const numPathParams = values(method.parameters).where((each: Parameter) => { return isPathParameter(each) && !isLiteralParameter(each); }).count();
  let content = '';
  if (numPathParams > 0) {
    imports.add('regexp');
    content += `\tconst regexStr = \`${createPathParamsRegex(method)}\`\n`;
    content += '\tregex := regexp.MustCompile(regexStr)\n';
    content += '\tmatches := regex.FindStringSubmatch(req.URL.EscapedPath())\n';
    content += `\tif matches == nil || len(matches) < ${numPathParams} {\n`;
    content += '\t\treturn nil, fmt.Errorf("failed to parse path %s", req.URL.Path)\n\t}\n';
  }
  if (values(method.parameters).where((each: Parameter) => { return isQueryParameter(each) && each.location === 'method' && !isLiteralParameter(each); }).any()) {
    content += '\tqp := req.URL.Query()\n';
  }
  const bodyParam = <BodyParameter | undefined>values(method.parameters).where((each: Parameter) => { return isBodyParameter(each) || isFormBodyParameter(each) || isMultipartFormBodyParameter(each); }).first();
  if (!bodyParam) {
    // no body, just headers and/or query params
  } else if (isMultipartFormBodyParameter(bodyParam)) {
    imports.add('io');
    imports.add('mime');
    imports.add('mime/multipart');
    content += '\t_, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))\n';
    content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    content += '\treader := multipart.NewReader(req.Body, params["boundary"])\n';
    for (const param of values(method.parameters)) {
      if (isMultipartFormBodyParameter(param)) {
        let pkgPrefix = '';
        if (isConstantType(param.type)) {
          pkgPrefix = clientPkg + '.';
        }
        content += `\tvar ${param.paramName} ${pkgPrefix}${getTypeDeclaration(param.type)}\n`;
      }
    }
    content += '\tfor {\n';
    content += '\t\tvar part *multipart.Part\n';
    content += '\t\tpart, err = reader.NextPart()\n';
    content += '\t\tif err == io.EOF {\n\t\t\tbreak\n';
    content += '\t\t} else if err != nil {\n\t\t\treturn nil, err\n\t\t}\n';
    content += '\t\tvar content []byte\n';
    content += '\t\tswitch fn := part.FormName(); fn {\n';
    for (const param of values(method.parameters)) {
      if (isMultipartFormBodyParameter(param)) {
        content += `\t\tcase "${param.paramName}":\n`;
        content += '\t\t\tcontent, err = io.ReadAll(part)\n';
        content += '\t\t\tif err != nil {\n\t\t\t\treturn nil, err\n\t\t\t}\n';
        let assignedValue: string;
        if (isStandardType(param.type) && param.type.typeName === 'io.ReadSeekCloser') {
          imports.add('bytes');
          imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
          assignedValue = 'streaming.NopCloser(bytes.NewReader(content))';
        } else if (isConstantType(param.type) || (isPrimitiveType(param.type) && param.type.typeName === 'string')) {
          assignedValue = 'string(content)';
        } else if (isSliceType(param.type)) {
          if (isStandardType(param.type.elementType) && param.type.elementType.typeName === 'io.ReadSeekCloser') {
            imports.add('bytes');
            imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
            assignedValue = `append(${param.paramName}, streaming.NopCloser(bytes.NewReader(content)))`;
          } else {
            throw new Error(`uhandled multipart parameter array element type ${getTypeDeclaration(param.type.elementType)}`);
          }
        } else {
          throw new Error(`uhandled multipart parameter type ${getTypeDeclaration(param.type)}`);
        }
        content += `\t\t\t${param.paramName} = ${assignedValue}\n`;
      }
    }
    content += '\t\tdefault:\n\t\t\treturn nil, fmt.Errorf("unexpected part %s", fn)\n';
    content += '\t\t}\n'; // end switch
    content += '\t}\n'; // end for
  } else if (isFormBodyParameter(bodyParam)) {
    for (const param of values(method.parameters)) {
      if (isFormBodyParameter(param)) {
        let pkgPrefix = '';
        if (isConstantType(param.type)) {
          pkgPrefix = clientPkg + '.';
        }
        content += `\tvar ${param.paramName} ${pkgPrefix}${getTypeDeclaration(param.type)}\n`;
      }
    }
    content += '\tif err := req.ParseForm(); err != nil {\n\t\treturn nil, &nonRetriableError{fmt.Errorf("failed parsing form data: %v", err)}\n\t}\n';
    content += '\tfor key := range req.Form {\n';
    content += '\t\tswitch key {\n';
    for (const param of values(method.parameters)) {
      if (isFormBodyParameter(param)) {
        content += `\t\tcase "${param.formDataName}":\n`;
        let assignedValue: string;
        if (isConstantType(param.type)) {
          assignedValue = `${getTypeDeclaration(param.type, clientPkg)}(req.FormValue(key))`;
        } else if (isPrimitiveType(param.type) && param.type.typeName === 'string') {
          assignedValue = 'req.FormValue(key)';
        } else {
          throw new Error(`uhandled form parameter type ${getTypeDeclaration(param.type)}`);
        }
        content += `\t\t\t${param.paramName} = ${assignedValue}\n`;
      }
    }
    content += '\t\t}\n'; // end switch
    content += '\t}\n'; // end for
  } else if (bodyParam.bodyFormat === 'binary') {
    // nothing to do for binary media type
  } else if (bodyParam.bodyFormat === 'Text') {
    if (bodyParam && !isLiteralParameter(bodyParam)) {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake', 'azfake');
      content += '\tbody, err := server.UnmarshalRequestAsText(req)\n';
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    }
  } else if (bodyParam.bodyFormat === 'JSON' || bodyParam.bodyFormat === 'XML') {
    if (bodyParam && !isLiteralParameter(bodyParam)) {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake', 'azfake');
      if (isBytesType(bodyParam.type)) {
        content += `\tbody, err := server.UnmarshalRequestAsByteArray(req, runtime.Base64${bodyParam.type.encoding}Format)\n`;
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      } else if (isSliceType(bodyParam.type) && bodyParam.type.rawJSONAsBytes) {
        content += '\tbody, err := io.ReadAll(req.Body)\n';
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
        content += '\treq.Body.Close()\n';
      } else if (isInterfaceType(bodyParam.type)) {
        content += '\traw, err := readRequestBody(req)\n';
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
        content += `\tbody, err := unmarshal${bodyParam.type.name}(raw)\n`;
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      } else {
        let bodyTypeName = getTypeDeclaration(bodyParam.type, clientPkg);
        if (isTimeType(bodyParam.type)) {
          bodyTypeName = bodyParam.type.dateTimeFormat;
        }
        content += `\tbody, err := server.UnmarshalRequestAs${bodyParam.bodyFormat}[${bodyTypeName}](req)\n`;
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      }
    }
  }
  content += createParamGroupParams(clientPkg, method, imports);
  const apiCall = `:= ${receiverName}.srv.${fixUpMethodName(method)}(${populateApiParams(clientPkg, method, imports)})`;
  if (isPageableMethod(method) && !isLROMethod(method)) {
    content += `resp ${apiCall}\n`;
    return content;
  }
  content += `\trespr, errRespr ${apiCall}\n`;
  content += '\tif respErr := server.GetError(errRespr, req); respErr != nil {\n';
  content += '\t\treturn nil, respErr\n\t}\n';
  return content;
}

function dispatchForLROBody(clientPkg: string, receiverName: string, method: Method, imports: ImportManager): string {
  const operationName = fixUpMethodName(method);
  const localVarName = uncapitalize(operationName);
  const operationStateMachine = `${receiverName}.${uncapitalize(operationName)}`;
  let content = `\t${localVarName} := ${operationStateMachine}.get(req)\n`;
  content += `\tif ${localVarName} == nil {\n`;
  content += dispatchForOperationBody(clientPkg, receiverName, method, imports);
  content += `\t\t${localVarName} = &respr\n`;
  content += `\t\t${operationStateMachine}.add(req, ${localVarName})\n`;
  content += '\t}\n\n';

  content += `\tresp, err := server.PollerResponderNext(${localVarName}, req)\n`;
  content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n\n';

  const formattedStatusCodes = formatStatusCodes(method.httpStatusCodes);
  content += `\tif !contains([]int{${formattedStatusCodes}}, resp.StatusCode) {\n`;
  content += `\t\t${operationStateMachine}.remove(req)\n`;
  content += `\t\treturn nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are ${formattedStatusCodes}", resp.StatusCode)}\n\t}\n`;

  content += `\tif !server.PollerResponderMore(${localVarName}) {\n`;
  content += `\t\t${operationStateMachine}.remove(req)\n\t}\n\n`;
  content += '\treturn resp, nil\n';
  return content;
}

function dispatchForPagerBody(clientPkg: string, receiverName: string, method: PageableMethod, imports: ImportManager): string {
  const operationName = fixUpMethodName(method);
  const localVarName = uncapitalize(operationName);
  const operationStateMachine = `${receiverName}.${uncapitalize(operationName)}`;
  let content = `\t${localVarName} := ${operationStateMachine}.get(req)\n`;
  content += `\tif ${localVarName} == nil {\n`;
  content += dispatchForOperationBody(clientPkg, receiverName, method, imports);
  content += `\t\t${localVarName} = &resp\n`;
  content += `\t\t${operationStateMachine}.add(req, ${localVarName})\n`;
  if (method.nextLinkName) {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
    content += `\t\tserver.PagerResponderInjectNextLinks(${localVarName}, req, func(page *${clientPkg}.${method.responseEnvelope.name}, createLink func() string) {\n`;
    content += `\t\t\tpage.${method.nextLinkName} = to.Ptr(createLink())\n`;
    content += '\t\t})\n';
  }
  content += '\t}\n'; // end if
  content += `\tresp, err := server.PagerResponderNext(${localVarName}, req)\n`;
  content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';

  const formattedStatusCodes = formatStatusCodes(method.httpStatusCodes);
  content += `\tif !contains([]int{${formattedStatusCodes}}, resp.StatusCode) {\n`;
  content += `\t\t${operationStateMachine}.remove(req)\n`;
  content += `\t\treturn nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are ${formattedStatusCodes}", resp.StatusCode)}\n\t}\n`;

  content += `\tif !server.PagerResponderMore(${localVarName}) {\n`;
  content += `\t\t${operationStateMachine}.remove(req)\n\t}\n`;
  content += '\treturn resp, nil\n';
  return content;
}

function sanitizeRegexpCaptureGroupName(name: string): string {
  // dash '-' characters are not allowed so replace them with '_'
  return name.replace('-', '_');
}

function createPathParamsRegex(method: Method): string {
  // "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}"
  // each path param will replaced with a regex capture.
  // note that some path params are optional.
  let urlPath = method.httpPath;
  for (const param of values(method.parameters)) {
    if (!isPathParameter(param)) {
      continue;
    }
    const toReplace = `{${param.pathSegment}}`;
    let replaceWith = `(?P<${sanitizeRegexpCaptureGroupName(param.pathSegment)}>[!#&$-;=?-\\[\\]_a-zA-Z0-9~%@]+)`;
    if (param.paramType === 'optional' || param.paramType === 'flag') {
      replaceWith += '?';
    }
    urlPath = urlPath.replace(toReplace, replaceWith);
  }
  return urlPath;
}

function createParamGroupParams(clientPkg: string, method: Method, imports: ImportManager): string {
  let content = '';

  // create any param groups and populate their values
  const paramGroups = new Map<ParameterGroup, Array<Parameter>>();
  for (const param of values(consolidateHostParams(method.parameters))) {
    if (param.location === 'client' || isLiteralParameter(param)) {
      // client params and parameter literals aren't passed to APIs
      continue;
    }
    if (param.group) {
      if (isResumeTokenParameter(param)) {
        // skip the ResumeToken param as we don't send that back to the caller
        continue;
      }
      if (!paramGroups.has(param.group)) {
        paramGroups.set(param.group, new Array<Parameter>());
      }
      const params = paramGroups.get(param.group);
      params!.push(param);
    }
    if (isBodyParameter(param) || isFormBodyParameter(param) || isMultipartFormBodyParameter(param)) {
      // body params will be unmarshalled, no need for parsing.
      continue;
    }
    
    let paramValue = getParamValue(param);
    if (isPathParameter(param) || isQueryParameter(param)) {
      // path/query params might be escaped, so we need to unescape them first
      imports.add('net/url');
      const paramVar = createLocalVariableName(param, 'Unescaped');
      let where: string;
      if (isPathParameter(param)) {
        where = 'Path';
      } else {
        where = 'Query';
      }
      content += `\t${paramVar}, err := url.${where}Unescape(${paramValue})\n`;
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      paramValue = paramVar;
    }

    // parse params as required
    if (isSliceType(param.type)) {
      if (isConstantType(param.type.elementType) || (isPrimitiveType(param.type.elementType) && param.type.elementType.typeName !== 'string')) {
        const asArray = param.type;
        imports.add('strings');
        content += `\telements := strings.Split(${paramValue}, "${getArraySeparator(param)}")\n`;
        const localVar = createLocalVariableName(param, 'Param');
        const toType = `${clientPkg}.${getTypeDeclaration(asArray.elementType)}`;
        content += `\t${localVar} := make([]${toType}, len(elements))\n`;
        content += '\tfor i := 0; i < len(elements); i++ {\n';
        let fromVar: string;
        if (isConstantType(asArray.elementType)) {
          if (asArray.elementType.type === 'int32' || asArray.elementType.type === 'int64') {
            imports.add('strconv');
            fromVar = 'parsedInt';
            content += `\t\tvar ${fromVar} int64\n`;
            content += `\t\t${fromVar}, err = strconv.ParseInt(elements[i], 10, 32)\n`;
            content += '\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n';
          } else if (asArray.elementType.type === 'float32' || asArray.elementType.type === 'float64') {
            imports.add('strconv');
            fromVar = 'parsedNum';
            content += `\t\tvar ${fromVar} float64\n`;
            content += `\t\t${fromVar}, err = strconv.ParseFloat(elements[i], 32)\n`;
            content += '\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n';
          } else if (asArray.elementType.type === 'string') {
            fromVar = 'elements[i]';
          } else {
            throw new Error(`unhandled array element choice type ${asArray.elementType.type}`);
          }
        } else {
          throw new Error(`unhandled array element type ${getTypeDeclaration(asArray.elementType)}`);
        }
        content += `\t\t${localVar}[i] = ${toType}(${fromVar})\n\t}\n`;
      } else if (param.group) {
        // for slices of strings that aren't in a parameter group, the call to strings.Split(...) is inlined
        // into the invocation of the fake e.g. srv.FakeFunc(strings.Split...). but if it's grouped, then we
        // need to create a local first which will later be copied into the param group.
        imports.add('strings');
        content += `\t${createLocalVariableName(param, 'Param')} := strings.Split(${paramValue}, "${getArraySeparator(param)}")\n`;
      }
    } else if (isPrimitiveType(param.type) && param.type.typeName === 'bool') {
      imports.add('strconv');
      let from = `strconv.ParseBool(${paramValue})`;
      if (!isRequiredParameter(param)) {
        from = `parseOptional(${paramValue}, strconv.ParseBool)`;
      }
      content += `\t${createLocalVariableName(param, 'Param')}, err := ${from}\n`;
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    } else if (isBytesType(param.type)) {
      imports.add('encoding/base64');
      content += `\t${createLocalVariableName(param, 'Param')}, err := base64.StdEncoding.DecodeString(${paramValue})\n`;
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    } else if (param.group && (isConstantType(param.type) || isLiteralValue(param.type) || (isPrimitiveType(param.type) && param.type.typeName === 'string'))) {
      // for params that don't require parsing, the value is inlined into the invocation of the fake.
      // but if it's grouped, then we need to create a local first which will later be copied into the param group.
      content += `\t${createLocalVariableName(param, 'Param')} := `;
      let paramValue = getParamValueWithCast(clientPkg, param, imports);
      if (!isRequiredParameter(param)) {
        paramValue = `getOptional(${paramValue})`;
      }
      content += `${paramValue}\n`;
    } else if (isTimeType(param.type)) {
      if (param.type.dateTimeFormat === 'dateType') {
        imports.add('time');
        let from = `time.Parse("2006-01-02", ${paramValue})`;
        if (!isRequiredParameter(param)) {
          from = `parseOptional(${paramValue}, func(v string) (time.Time, error) { return time.Parse("2006-01-02", v) })`;
        }
        content += `\t${createLocalVariableName(param, 'Param')}, err := ${from}\n`;
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      } else if (param.type.dateTimeFormat === 'timeRFC1123' || param.type.dateTimeFormat === 'timeRFC3339') {
        imports.add('time');
        let format = 'time.RFC3339Nano';
        if (param.type.dateTimeFormat === 'timeRFC1123') {
          format = 'time.RFC1123';
        }
        let from = `time.Parse(${format}, ${paramValue})`;
        if (!isRequiredParameter(param)) {
          from = `parseOptional(${paramValue}, func(v string) (time.Time, error) { return time.Parse(${format}, v) })`;
        }
        content += `\t${createLocalVariableName(param, 'Param')}, err := ${from}\n`;
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      } else {
        imports.add('strconv');
        let from = `strconv.ParseInt(${paramValue}, 10, 64)`;
        if (!isRequiredParameter(param)) {
          from = `parseOptional(${paramValue}, func(v string) (int64, error) { return strconv.ParseInt(v, 10, 64) })`;
        }
        content += `\t${createLocalVariableName(param, 'Param')}, err := ${from}\n`;
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      }
    } else if (isPrimitiveType(param.type) && (param.type.typeName === 'float32' || param.type.typeName === 'float64' || param.type.typeName === 'int32' || param.type.typeName === 'int64')) {
      imports.add('strconv');
      let precision: '32' | '64' = '32';
      if (param.type.typeName === 'float64' || param.type.typeName === 'int64') {
        precision = '64';
      }
      let parseType: 'Int' | 'Float' = 'Int';
      let base = '10, ';
      if (param.type.typeName === 'float32' || param.type.typeName === 'float64') {
        parseType = 'Float';
        base = '';
      }
      let parser = 'parseWithCast';
      if (!isRequiredParameter(param)) {
        parser = 'parseOptional';
      }
      if (precision === '32' || !isRequiredParameter(param)) {
        content += `\t${createLocalVariableName(param, 'Param')}, err := ${parser}(${paramValue}, func(v string) (${parseType.toLowerCase()}${precision}, error) {\n`;
        content += `\t\tp, parseErr := strconv.Parse${parseType}(v, ${base}${precision})\n`;
        content += '\t\tif parseErr != nil {\n\t\t\treturn 0, parseErr\n\t\t}\n';
        let result = 'p';
        if (precision === '32') {
          result = `${parseType.toLowerCase()}${precision}(${result})`;
        }
        content += `\t\treturn ${result}, nil\n\t})\n`;
      } else {
        content += `\t${createLocalVariableName(param, 'Param')}, err := strconv.Parse${parseType}(${paramValue}, ${base}64)\n`;
      }
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    } else if (isHeaderParameter(param) && param.collectionPrefix) {
      imports.add('strings');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
      const localVar = createLocalVariableName(param, 'Param');
      content += `\tvar ${localVar} map[string]*string\n`;
      content += '\tfor hh := range req.Header {\n';
      const headerPrefix = param.collectionPrefix;
      content += `\t\tif len(hh) > len("${headerPrefix}") && strings.EqualFold(hh[:len("x-ms-meta-")], "${headerPrefix}") {\n`;
      content += `\t\t\tif ${localVar} == nil {\n\t\t\t\t${localVar} = map[string]*string{}\n\t\t\t}\n`;
      content += `\t\t\t${localVar}[hh[len("${headerPrefix}"):]] = to.Ptr(getHeaderValue(req.Header, hh))\n`;
      content += '\t\t}\n\t}\n';
    }
  }

  for (const paramGroup of values(paramGroups.keys())) {
    if (paramGroup.required === true) {
      content += `\t${uncapitalize(paramGroup.paramName)} := ${clientPkg}.${paramGroup.groupName}{\n`;
      for (const param of values(paramGroups.get(paramGroup))) {
        content += `\t\t${capitalize(param.paramName)}: ${createLocalVariableName(param, 'Param')},\n`;
      }
      content += '\t}\n';
    } else {
      content += `\tvar ${uncapitalize(paramGroup.paramName)} *${clientPkg}.${paramGroup.groupName}\n`;
      const params = paramGroups.get(paramGroup);
      const paramNilCheck = new Array<string>();
      for (const param of values(params)) {
        // check array before body in case the body is just an array
        if (isSliceType(param.type)) {
          paramNilCheck.push(`len(${createLocalVariableName(param, 'Param')}) > 0`);
        } else if (isBodyParameter(param)) {
          if (param.bodyFormat === 'binary') {
            imports.add('io');
            paramNilCheck.push('req.Body != nil');
          } else {
            imports.add('reflect');
            paramNilCheck.push('!reflect.ValueOf(body).IsZero()');
          }
        } else if (isFormBodyParameter(param) || isMultipartFormBodyParameter(param)) {
          imports.add('reflect');
          paramNilCheck.push(`!reflect.ValueOf(${param.paramName}).IsZero()`);
        } else {
          paramNilCheck.push(`${createLocalVariableName(param, 'Param')} != nil`);
        }
      }
      content += `\tif ${paramNilCheck.join(' || ')} {\n`;
      content += `\t\t${uncapitalize(paramGroup.paramName)} = &${clientPkg}.${paramGroup.groupName}{\n`;
      for (const param of values(params)) {
        let byRef = '&';
        if (param.byValue || (!isRequiredParameter(param) && !isBodyParameter(param) && !isFormBodyParameter(param) && !isMultipartFormBodyParameter(param))) {
          byRef = '';
        }
        content += `\t\t\t${capitalize(param.paramName)}: ${byRef}${createLocalVariableName(param, 'Param')},\n`;
      }
      content += '\t\t}\n';
      content += '\t}\n';
    }
  }

  return content;
}

function populateApiParams(clientPkg: string, method: Method, imports: ImportManager): string {
  // FooOperation(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], qp.Get("api-version"), nil)
  // this assumes that our caller has created matches and qp as required
  const params = new Array<string>();

  // for non-paged APIs, first param is always the context. use the one
  // from the HTTP request. be careful to properly handle paged LROs
  if (isLROMethod(method) || !isPageableMethod(method)) {
    params.push('req.Context()');
  }

  // now create the API call sig
  for (const param of values(getMethodParameters(method, consolidateHostParams))) {
    if (isParameterGroup(param)) {
      if (param.groupName === method.optionalParamsGroup.groupName) {
        // this is the optional params type. in some cases we just pass nil
        const countParams = values(param.params).where((each: Parameter) => { return each.location === 'method' && !isResumeTokenParameter(each); }).count();
        if (countParams === 0) {
          // if the options param is empty or only contains the resume token param just pass nil
          params.push('nil');
          continue;
        }
      }
      params.push(uncapitalize(param.paramName));
      continue;
    }
    imports.addImportForType(param.type);
    params.push(`${getParamValueWithCast(clientPkg, param, imports)}`);
  }

  return params.join(', ');
}

function getParamValue(param: Parameter): string {
  if (isFormBodyParameter(param) || isMultipartFormBodyParameter(param)) {
    // multipart form data values have been read and assigned
    // to local params with the same name. must check this first
    // as it's a superset of other cases that follow.
    return param.paramName;
  } else if (isHeaderParameter(param) && param.collectionPrefix) {
    // must check before http?.in === 'header'
    return createLocalVariableName(param, 'Param');
  } else if (isLiteralValue(param.type) && isLiteralParameter(param)) {
    // required constants have their values embedded in the generated code
    return formatLiteralValue(param.type);
  } else if (isPathParameter(param)) {
    // path params are in the matches slice
    return `matches[regex.SubexpIndex("${sanitizeRegexpCaptureGroupName(param.pathSegment)}")]`;
  } else if (isQueryParameter(param)) {
    // use qp
    return `qp.Get("${param.queryParameter}")`;
  } else if (isHeaderParameter(param)) {
    // use req
    return `getHeaderValue(req.Header, "${param.headerName}")`;
  } else if (isBodyParameter(param)) {
    if (param.bodyFormat === 'binary') {
      return 'req.Body.(io.ReadSeekCloser)';
    }
    // JSON/XML bodies have been unmarshalled into a local named body
    return 'body';
  } else if (isURIParameter(param)) {
    return 'req.URL.Host';
  } else if (param.group) {
    // this is a parameter group param
    return uncapitalize(param.paramName);
  } else {
    throw new Error(`unhandled parameter ${param.paramName}`);
  }
}

function getParamValueWithCast(clientPkg: string, param: Parameter, imports: ImportManager): string {
  let value = getParamValue(param);
  if (isBodyParameter(param) || isFormBodyParameter(param) || isMultipartFormBodyParameter(param)) {
    // if the param is in the body, it's already in the correct format
    if (isTimeType(param.type)) {
      return `time.Time(${value})`;
    }
    return value;
  } else if (isPathParameter(param) || isQueryParameter(param)) {
    // path/query params were unescaped earlier
    value = createLocalVariableName(param, 'Unescaped');
  }

  if (isSliceType(param.type)) {
    const asArray = param.type;
    if (isConstantType(asArray.elementType)) {
      return createLocalVariableName(param, 'Param');
    } else if (isPrimitiveType(asArray.elementType) && asArray.elementType.typeName === 'string') {
      imports.add('strings');
      return `strings.Split(${value}, "${getArraySeparator(param)}")`;
    } else {
      throw new Error(`unhandled array element type ${getTypeDeclaration(asArray.elementType)}`);
    }
  } else if (isConstantType(param.type)) {
    return `${clientPkg}.${getTypeDeclaration(param.type)}(${value})`;
  } else if ((isPrimitiveType(param.type) && param.type.typeName === 'bool') || isBytesType(param.type) || (isTimeType(param.type) && param.type.dateTimeFormat !== 'timeUnix')) {
    return createLocalVariableName(param, 'Param');
  } else if (isPrimitiveType(param.type) && (param.type.typeName === 'float32' || param.type.typeName === 'float64' || param.type.typeName === 'int32' || param.type.typeName === 'int64')) {
    const numSchema = param.type;
    // optional params have been parsed/converted into their required bitness
    if ((numSchema.typeName === 'float32' || numSchema.typeName === 'int32') && isRequiredParameter(param)) {
      return `${numSchema.typeName}(${createLocalVariableName(param, 'Param')})`;
    }
    // parsed ints/numbers have a default type of int64
    return createLocalVariableName(param, 'Param');
  } else if (isTimeType(param.type) && param.type.dateTimeFormat === 'timeUnix') {
    return `time.Unix(${createLocalVariableName(param, 'Param')}, 0)`;
  } else {
    return value;
  }
}

function createLocalVariableName(param: Parameter, suffix: string): string {
  if (isFormBodyParameter(param) || isMultipartFormBodyParameter(param)) {
    return param.paramName;
  } else if (isBodyParameter(param)) {
    if (param.bodyFormat === 'binary') {
      return 'req.Body.(io.ReadSeekCloser)';
    }
    return 'body';
  }
  return `${uncapitalize(param.paramName)}${suffix}`;
}

// call sites of this function make the assumption that if a param is a slice
// then it has a delimiter. while this is probably true considering the context
// where this is invoked, it's a bit of a hack.
function getArraySeparator(param: Parameter): string {
  if ((isFormBodyParameter(param) || isHeaderParameter(param) || isPathParameter(param) || isQueryParameter(param)) && param.delimiter) {
    return param.delimiter;
  }
  return ',';
}

// takes multiple host parameters and consolidates them into a single "host" parameter.
// this is necessary as there's no way to rehydrate multiple host parameters.
// e.g. host := "{vault}{secret}{dnsSuffix}" becomes http://contososecret.com
// there's no way to reliably split the host back up into its constituent parameters.
// so we just pass the full value as a single host parameter.
function consolidateHostParams(params?: Array<Parameter>): Array<Parameter> | undefined {
  if (!values(params).where((each: Parameter) => { return isURIParameter(each); }).any()) {
    // no host params
    return params;
  }

  // consolidate multiple host params into a single "host" param
  const consolidatedParams = new Array<Parameter>();
  let hostParamAdded = false;
  for (const param of values(params)) {
    if (!isURIParameter(param)) {
      consolidatedParams.push(param);
    } else if (!hostParamAdded) {
      consolidatedParams.push(param);
      hostParamAdded = true;
    }
  }

  return consolidatedParams;
}

// copied from generator/operations.ts but with a slight tweak to consolidate host parameters
function getAPIParametersSig(method: Method, imports: ImportManager, pkgName?: string): string {
  const methodParams = getMethodParameters(method, consolidateHostParams);
  const params = new Array<string>();
  if (!isPageableMethod(method) || isLROMethod(method)) {
    imports.add('context');
    params.push('ctx context.Context');
  }
  for (const methodParam of values(methodParams)) {
    let paramName = uncapitalize(methodParam.paramName);
    if (isParameter(methodParam) && isURIParameter(methodParam)) {
      paramName = 'host';
    }
    params.push(`${paramName} ${formatParameterTypeName(methodParam, pkgName)}`);
  }
  return params.join(', ');
}

// copied from generator/helpers.ts but without the XML-specific stuff
function getResultFieldName(result: AnyResult | BinaryResult | MonomorphicResult | PolymorphicResult | ModelResult): string {
  if (isAnyResult(result)) {
    return result.fieldName;
  } else if (isModelResult(result)) {
    return result.modelType.name;
  } else if (isPolymorphicResult(result)) {
    return result.interfaceType.name;
  }
  return result.fieldName;
}
