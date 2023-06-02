/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { ArraySchema, ByteArraySchema, ChoiceSchema, CodeModel, ConstantSchema, DateTimeSchema, GroupProperty, ImplementationLocation, NumberSchema, OperationGroup, Operation, Parameter, SchemaType, SerializationStyle } from '@autorest/codemodel';
import { capitalize, uncapitalize } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import { aggregateParameters, formatConstantValue, getSchemaResponse, isBinaryResponseOperation, isLROOperation, isMultiRespOperation, isPageableOperation, isSchemaResponse,isTypePassedByValue } from '../common/helpers';
import { contentPreamble, formatParameterTypeName, formatStatusCode, formatStatusCodes, formatTypeName, formatValue, getMethodParameters, getParentImport, getResponseEnvelope, getResponseEnvelopeName, getStatusCodes } from '../generator/helpers';
import { fixUpOperationName, getMediaType } from '../generator/operations';
import { ImportManager } from '../generator/imports';

// represents the generated content for an operation group
export class OperationGroupContent {
  readonly name: string;
  readonly content: string;

  constructor(name: string, content: string) {
    this.name = name;
    this.content = content;
  }
}

export async function generateServers(session: Session<CodeModel>): Promise<OperationGroupContent[]> {
  const operations = new Array<OperationGroupContent>();
  const clientPkg = session.model.language.go!.packageName;
  for (const group of values(session.model.operationGroups)) {
    // the list of packages to import
    const imports = new ImportManager();
    imports.add(await getParentImport(session));

    // add standard imports
    imports.add('errors');
    imports.add('fmt');
    imports.add('net/http');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake', 'azfake');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');

    // for the fake server, we use the suffix Server instead of Client
    const serverName = capitalize((<string>group.language.go!.clientName).replace(/[C|c]lient$/, 'Server'));

    let content: string;
    content = `// ${serverName} is a fake server for instances of the ${clientPkg}.${group.language.go!.clientName} type.\n`;
    content += `type ${serverName} struct{\n`;

    // we might remove some operations from the list
    const finalOperations = new Array<Operation>();

    for (const op of values(group.operations)) {
      if (isLROOperation(op)) {
        let respType = `${clientPkg}.${getResponseEnvelopeName(op)}`;
        if (isPageableOperation(op)) {
          respType = `azfake.PagerResponder[${clientPkg}.${getResponseEnvelopeName(op)}]`;
        }
        op.language.go!.serverResponse = `resp azfake.PollerResponder[${respType}], errResp azfake.ErrorResponder`;
      } else if (isPageableOperation(op)) {
        if (op.language.go!.paging.isNextOp) {
          // we don't generate a public API for the methods used to advance pages, so skip it here
          continue;
        }
        op.language.go!.serverResponse = `resp azfake.PagerResponder[${clientPkg}.${getResponseEnvelopeName(op)}]`;
      } else {
        op.language.go!.serverResponse = `resp azfake.Responder[${clientPkg}.${getResponseEnvelopeName(op)}], errResp azfake.ErrorResponder`;
      }

      const operationName = fixUpOperationName(op);
      content += `\t// ${operationName} is the fake for method ${group.language.go!.clientName}.${operationName}\n`;
      const successCodes = new Array<string>();
      if (isMultiRespOperation(op)) {
        for (const response of values(op.responses)) {
          if (!isSchemaResponse(response)) {
            // the operation contains a mix of schemas and non-schema responses
            successCodes.push(`${formatStatusCode(response.protocol.http!.statusCodes[0])} (no return type)`);
            continue;
          }
          successCodes.push(`${formatStatusCode(response.protocol.http!.statusCodes[0])} (returns ${formatTypeName(response.schema, clientPkg)})`);
        }
        content += '\t// HTTP status codes to indicate success:\n';
        for (const successCode of successCodes) {
          content += `\t//   - ${successCode}\n`;
        }
      } else {
        for (const statusCode of values(getStatusCodes(op))) {
          successCodes.push(`${formatStatusCode(statusCode)}`);
        }
        content += `\t// HTTP status codes to indicate success: ${successCodes.join(', ')}\n`;
      }
      content += `\t${operationName} func(${getAPIParametersSig(op, imports, clientPkg)}) (${op.language.go!.serverResponse})\n\n`;
      finalOperations.push(op);
    }

    content += '}\n\n';

    ///////////////////////////////////////////////////////////////////////////

    const serverTransport = `${serverName}Transport`;

    content += `// New${serverTransport} creates a new instance of ${serverTransport} with the provided implementation.\n`;
    content += `// The returned ${serverTransport} instance is connected to an instance of ${clientPkg}.${group.language.go!.clientName} by way of the\n`;
    content += `// ${group.language.go!.clientOptionsType}.Transporter field.\n`;
    content += `func New${serverTransport}(srv *${serverName}) *${serverTransport} {\n`;
    content += `\treturn &${serverTransport}{srv: srv}\n}\n\n`;

    content += `// ${serverTransport} connects instances of ${clientPkg}.${group.language.go!.clientName} to instances of ${serverName}.\n`;
    content += `// Don't use this type directly, use New${serverTransport} instead.\n`;
    content += `type ${serverTransport} struct {\n`;
    content += `\tsrv *${serverName}\n`;
    for (const op of values(finalOperations)) {
      // create state machines for any pager/poller operations
      if (isLROOperation(op)) {
        let respType = `${clientPkg}.${getResponseEnvelopeName(op)}`;
        if (isPageableOperation(op)) {
          respType = `azfake.PagerResponder[${clientPkg}.${getResponseEnvelopeName(op)}]`;
        }
        content +=`\t${uncapitalize(fixUpOperationName(op))} *azfake.PollerResponder[${respType}]\n`;
      } else if (isPageableOperation(op)) {
        content += `\t${uncapitalize(fixUpOperationName(op))} *azfake.PagerResponder[${clientPkg}.${getResponseEnvelopeName(op)}]\n`;
      }
    }
    content += '}\n\n';

    content += generateServerTransportMethods(clientPkg, serverTransport, group, finalOperations, imports);

    ///////////////////////////////////////////////////////////////////////////

    // stitch everything together
    let text = await contentPreamble(session, 'fake');
    text += imports.text();
    text += content;
    operations.push(new OperationGroupContent(serverName, text));
  }
  return operations;
}

function generateServerTransportMethods(clientPkg: string, serverTransport: string, group: OperationGroup, finalOperations: Array<Operation>, imports: ImportManager): string {
  const receiverName = serverTransport[0].toLowerCase();
  let content = `// Do implements the policy.Transporter interface for ${serverTransport}.\n`;
  content += `func (${receiverName} *${serverTransport}) Do(req *http.Request) (*http.Response, error) {\n`;
  content += '\trawMethod := req.Context().Value(runtime.CtxAPINameKey{})\n';
  content += '\tmethod, ok := rawMethod.(string)\n';
  content += '\tif !ok {\n\t\treturn nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}\n\t}\n\n';
  content += '\tvar resp *http.Response\n';
  content += '\tvar err error\n\n';
  content += '\tswitch method {\n';

  for (const op of values(finalOperations)) {
    const operationName = fixUpOperationName(op);
    content += `\tcase "${group.language.go!.clientName}.${operationName}":\n`;
    content += `\t\tresp, err = ${receiverName}.dispatch${operationName}(req)\n`;
  }
  content += '\tdefault:\n\t\terr = fmt.Errorf("unhandled API %s", method)\n';

  content += '\t}\n\n'; // end switch
  content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n\n';
  content += '\treturn resp, nil\n}\n\n';

  ///////////////////////////////////////////////////////////////////////////

  for (const op of values(finalOperations)) {
    content += `func (${receiverName} *${serverTransport}) dispatch${fixUpOperationName(op)}(req *http.Request) (*http.Response, error) {\n`;
    content += `\tif ${receiverName}.srv.${fixUpOperationName(op)} == nil {\n`;
    content += `\t\treturn nil, &nonRetriableError{errors.New("method ${fixUpOperationName(op)} not implemented")}\n\t}\n`;

    if (isLROOperation(op)) {
      // must check LRO before pager as you can have paged LROs
      content += dispatchForLROBody(clientPkg, receiverName, op, imports);
    } else if (isPageableOperation(op)) {
      content += dispatchForPagerBody(clientPkg, receiverName, op, imports);
    } else {
      content += dispatchForOperationBody(clientPkg, receiverName, op, imports);
      content += '\trespContent := server.GetResponseContent(respr)\n';
      const formattedStatusCodes = formatStatusCodes(getStatusCodes(op));
      content += `\tif !contains([]int{${formattedStatusCodes}}, respContent.HTTPStatus) {\n`;
      content += `\t\treturn nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are ${formattedStatusCodes}", respContent.HTTPStatus)}\n\t}\n`;
      if (isMultiRespOperation(op)) {
        content += `\tresp, err := server.MarshalResponseAs${getMediaTypeForMultiRespOperation(op)}(respContent, server.GetResponse(respr).${getResultFieldName(op)}, req)\n`;
      } else if (isBinaryResponseOperation(op)) {
        content += `\tresp, err := server.NewResponse(respContent, req, &server.ResponseOptions{\n`;
        content += `\t\tBody: server.GetResponse(respr).${getResultFieldName(op)},\n`;
        content += '\t\tContentType: "application/octet-stream",\n';
        content += '\t})\n';
      } else if (getSchemaResponse(op)) {
        const schemaResponse = getSchemaResponse(op);
        if (schemaResponse!.schema.type === SchemaType.ByteArray) {
          let format = 'Std';
          if ((<ByteArraySchema>schemaResponse!.schema).format === 'base64url') {
            format = 'URL';
          }
          content += `\tresp, err := server.MarshalResponseAsByteArray(respContent, server.GetResponse(respr).${getResultFieldName(op)}, runtime.Base64${format}Format, req)\n`;
        } else if (schemaResponse!.schema.language.go!.rawJSONAsBytes) {
          imports.add('bytes');
          imports.add('io');
          content += '\tresp, err := server.NewResponse(respContent, req, &server.ResponseOptions{\n';
          content += '\t\tBody: io.NopCloser(bytes.NewReader(server.GetResponse(respr).RawJSON)),\n';
          content += '\t\tContentType: "application/json",\n\t})\n';
        } else {
          let respField = `.${getResultFieldName(op)}`;
          if (getMediaType(schemaResponse!.protocol) === 'XML' && schemaResponse!.schema.type === SchemaType.Array) {
            // for XML array responses we use the response type directly as it has the necessary XML tag for proper marshalling
            respField = '';
          }
          let responseField = `server.GetResponse(respr)${respField}`;
          if (schemaResponse!.schema.type === SchemaType.DateTime || schemaResponse!.schema.type === SchemaType.UnixTime || schemaResponse!.schema.type === SchemaType.Date) {
            responseField = `(*${schemaResponse!.schema.language.go!.internalTimeType})(${responseField})`;
          }
          content += `\tresp, err := server.MarshalResponseAs${getMediaType(schemaResponse!.protocol)}(respContent, ${responseField}, req)\n`;
        }
      } else {
        content += '\tresp, err := server.NewResponse(respContent, req, nil)\n';
      }

      content += `\tif err != nil {\t\treturn nil, err\n\t}\n`;

      // propagate any header response values into the *http.Response
      const respEnv = getResponseEnvelope(op);
      for (const prop of values(respEnv.properties)) {
        if (prop.language.go!.fromHeader) {
          if (prop.schema.language.go!.headerCollectionPrefix) {
            content += `\tfor k, v := range server.GetResponse(respr).${prop.language.go!.name} {\n`;
            content += `\t\tif v != nil {\n`;
            content += `\t\t\tresp.Header.Set("${prop.schema.language.go!.headerCollectionPrefix}"+k, *v)\n`;
            content += `\t\t}\n`;
            content += `\t}\n`;
          } else {
            content += `\tif val := server.GetResponse(respr).${prop.language.go!.name}; val != nil {\n`;
            content += `\t\tresp.Header.Set("${prop.language.go!.fromHeader}", ${formatValue('val', prop.schema, imports, true)})\n\t}\n`;
          }
        }
      }

      content += '\treturn resp, nil\n';
    }
    content += '}\n\n';
  }

  return content;
}

function getMediaTypeForMultiRespOperation(op: Operation): string {
  let uberMediaType = '';

  for (const response of values(op.responses)) {
    if (isSchemaResponse(response)) {
      const mediaType = getMediaType(response.protocol);
      if (uberMediaType === '') {
        uberMediaType = mediaType;
      } else if (uberMediaType !== mediaType) {
        throw new Error(`operation ${op.language.go!.name} contains mixed media types which is not supported`);
      }
    }
  }

  return uberMediaType;
}

function dispatchForOperationBody(clientPkg: string, receiverName: string, op: Operation, imports: ImportManager): string {
  const numPathParams = values(op.parameters).where((each: Parameter) => { return each.protocol.http?.in === 'path'; }).count();
  let content = '';
  if (numPathParams > 0) {
    imports.add('regexp');
    content += `\tconst regexStr = "${createPathParamsRegex(op)}"\n`;
    content += '\tregex := regexp.MustCompile(regexStr)\n';
    content += '\tmatches := regex.FindStringSubmatch(req.URL.Path)\n';
    content += `\tif matches == nil || len(matches) < ${numPathParams} {\n`;
    content += `\t\treturn nil, fmt.Errorf("failed to parse path %s", req.URL.Path)\n\t}\n`;
  }
  if (values(op.parameters).where((each: Parameter) => { return each.protocol.http?.in === 'query' && each.implementation === ImplementationLocation.Method && (each.schema.type !== SchemaType.Constant || !each.required); }).any()) {
    content += '\tqp := req.URL.Query()\n';
  }
  const mediaType = getMediaType(op.requests![0].protocol);
  if (mediaType === 'multipart') {
    imports.add('io');
    imports.add('mime');
    imports.add('mime/multipart');
    content += '\t_, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))\n';
    content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    content += '\treader := multipart.NewReader(req.Body, params["boundary"])\n';
    for (const param of values(aggregateParameters(op))) {
      if (param.isPartialBody) {
        let pkgPrefix = '';
        if (param.schema.type === SchemaType.Choice || param.schema.type === SchemaType.SealedChoice) {
          pkgPrefix = clientPkg + '.';
        }
        content += `\tvar ${param.language.go!.name} ${pkgPrefix}${param.schema.language.go!.name}\n`;
      }
    }
    content += '\tfor {\n';
    content += '\t\tvar part *multipart.Part\n';
    content += '\t\tpart, err = reader.NextPart()\n';
    content += '\t\tif err == io.EOF {\n\t\t\tbreak\n';
    content += '\t\t} else if err != nil {\n\t\t\treturn nil, err\n\t\t}\n';
    content += '\t\tvar content []byte\n';
    content += '\t\tswitch fn := part.FormName(); fn {\n';
    for (const param of values(aggregateParameters(op))) {
      if (param.isPartialBody) {
        content += `\t\tcase "${param.language.go!.name}":\n`;
        content += '\t\t\tcontent, err = io.ReadAll(part)\n';
        content += '\t\t\tif err != nil {\n\t\t\t\treturn nil, err\n\t\t\t}\n';
        let assignedValue: string;
        if (param.schema.type === SchemaType.Binary) {
          imports.add('bytes');
          imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
          assignedValue = 'streaming.NopCloser(bytes.NewReader(content))';
        } else if (param.schema.type === SchemaType.Choice || param.schema.type === SchemaType.SealedChoice || param.schema.type === SchemaType.String) {
          assignedValue = 'string(content)';
        } else if (param.schema.type === SchemaType.Array) {
          const asArray = <ArraySchema>param.schema;
          if (asArray.elementType.type === SchemaType.Binary) {
            imports.add('bytes');
            imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
            assignedValue = `append(${param.language.go!.name}, streaming.NopCloser(bytes.NewReader(content)))`;
          } else {
            throw new Error(`uhandled multipart parameter array element type ${asArray.elementType.type}`);
          }
        } else {
          throw new Error(`uhandled multipart parameter type ${param.schema.type}`);
        }
        content += `\t\t\t${param.language.go!.name} = ${assignedValue}\n`;
      }
    }
    content += '\t\tdefault:\n\t\t\treturn nil, fmt.Errorf("unexpected part %s", fn)\n';
    content += '\t\t}\n'; // end switch
    content += '\t}\n'; // end for
  } else if (mediaType === 'form') {
    for (const param of values(aggregateParameters(op))) {
      if (param.isPartialBody) {
        let pkgPrefix = '';
        if (param.schema.type === SchemaType.Choice || param.schema.type === SchemaType.SealedChoice) {
          pkgPrefix = clientPkg + '.';
        }
        content += `\tvar ${param.language.go!.name} ${pkgPrefix}${param.schema.language.go!.name}\n`;
      }
    }
    content += '\tif err := req.ParseForm(); err != nil {\n\t\treturn nil, &nonRetriableError{fmt.Errorf("failed parsing form data: %v", err)}\n\t}\n';
    content += '\tfor key := range req.Form {\n';
    content += '\t\tswitch key {\n';
    for (const param of values(aggregateParameters(op))) {
      if (param.isPartialBody) {
        content += `\t\tcase "${param.language.go!.serializedName}":\n`;
        let assignedValue: string;
        switch (param.schema.type) {
          case SchemaType.Choice:
          case SchemaType.SealedChoice:
            assignedValue = `${clientPkg}.${param.schema.language.go!.name}(req.FormValue(key))`;
            break;
          case SchemaType.String:
            assignedValue = 'req.FormValue(key)';
            break;
          default:
            throw new Error(`uhandled form parameter type ${param.schema.type}`);
        }
        content += `\t\t\t${param.language.go!.name} = ${assignedValue}\n`;
      }
    }
    content += '\t\t}\n'; // end switch
    content += '\t}\n'; // end for
  } else if (mediaType === 'binary') {
    // nothing to do for binary media type
  } else if (mediaType === 'Text') {
    const bodyParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http?.in === 'body'; }).first();
    if (bodyParam && bodyParam.schema.type !== SchemaType.Constant) {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake', 'azfake');
      content += `\tbody, err := server.UnmarshalRequestAsText(req)\n`;
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    }
  } else if (mediaType === 'JSON' || mediaType === 'XML') {
    const bodyParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http?.in === 'body'; }).first();
    if (bodyParam && bodyParam.schema.type !== SchemaType.Constant) {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake', 'azfake');
      if (bodyParam.schema.type === SchemaType.ByteArray) {
        let format = 'Std';
        if ((<ByteArraySchema>bodyParam.schema).format === 'base64url') {
          format = 'URL';
        }
        content += `\tbody, err := server.UnmarshalRequestAsByteArray(req, runtime.Base64${format}Format)\n`;
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      } else if (bodyParam.schema.language.go!.rawJSONAsBytes) {
        content += `\tbody, err := io.ReadAll(req.Body)\n`;
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
        content += '\treq.Body.Close()\n';
      } else if (bodyParam.schema.language.go!.discriminatorInterface) {
        content += '\traw, err := readRequestBody(req)\n';
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
        content += `\tbody, err := unmarshal${bodyParam.schema.language.go!.discriminatorInterface}(raw)\n`;
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      } else {
        let bodyTypeName = formatTypeName(bodyParam.schema, clientPkg);
        if (bodyParam.schema.type === SchemaType.DateTime || bodyParam.schema.type === SchemaType.UnixTime || bodyParam.schema.type === SchemaType.Date) {
          bodyTypeName = bodyParam.schema.language.go!.internalTimeType;
        }
        content += `\tbody, err := server.UnmarshalRequestAs${mediaType}[${bodyTypeName}](req)\n`;
        content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
      }
    }
  } /*else {
    // no body, just headers and/or query params
  }*/
  content += createParamGroupParams(clientPkg, op, imports);
  const apiCall = `:= ${receiverName}.srv.${fixUpOperationName(op)}(${populateApiParams(clientPkg, op, imports)})`;
  if (isPageableOperation(op) && !isLROOperation(op)) {
    content += `resp ${apiCall}\n`;
    return content;
  }
  content += `\trespr, errRespr ${apiCall}\n`;
  content += '\tif respErr := server.GetError(errRespr, req); respErr != nil {\n';
  content += '\t\treturn nil, respErr\n\t}\n';
  return content;
}

function dispatchForLROBody(clientPkg: string, receiverName: string, op: Operation, imports: ImportManager): string {
  const operationName = fixUpOperationName(op);
  const operationStateMachine = `${receiverName}.${uncapitalize(operationName)}`;
  let content = `\tif ${operationStateMachine} == nil {\n`;
  content += dispatchForOperationBody(clientPkg, receiverName, op, imports);
  content += `\t\t${operationStateMachine} = &respr\n`;
  content += '\t}\n\n';

  content += `\tresp, err := server.PollerResponderNext(${operationStateMachine}, req)\n`;
  content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n\n';

  const formattedStatusCodes = formatStatusCodes(getStatusCodes(op));
  content += `\tif !contains([]int{${formattedStatusCodes}}, resp.StatusCode) {\n`;
  content += `\t\treturn nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are ${formattedStatusCodes}", resp.StatusCode)}\n\t}\n`;

  content += `\tif !server.PollerResponderMore(${operationStateMachine}) {\n`;
  content += `\t\t${operationStateMachine} = nil\n\t}\n\n`;
  content += '\treturn resp, nil\n'
  return content;
}

function dispatchForPagerBody(clientPkg: string, receiverName: string, op: Operation, imports: ImportManager): string {
  const operationName = fixUpOperationName(op);
  const operationStateMachine = `${receiverName}.${uncapitalize(operationName)}`;
  let content = `\tif ${operationStateMachine} == nil {\n`;
  content += dispatchForOperationBody(clientPkg, receiverName, op, imports);
  content += `\t\t${operationStateMachine} = &resp\n`;
  if (op.language.go!.paging.nextLinkName) {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
    content += `\t\tserver.PagerResponderInjectNextLinks(${operationStateMachine}, req, func(page *${clientPkg}.${getResponseEnvelopeName(op)}, createLink func() string) {\n`;
    content += `\t\t\tpage.${op.language.go!.paging.nextLinkName} = to.Ptr(createLink())\n`;
    content += `\t\t})\n`;
  }
  content += '\t}\n'; // end if
  content += `\tresp, err := server.PagerResponderNext(${operationStateMachine}, req)\n`;
  content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';

  const formattedStatusCodes = formatStatusCodes(getStatusCodes(op));
  content += `\tif !contains([]int{${formattedStatusCodes}}, resp.StatusCode) {\n`;
  content += `\t\treturn nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are ${formattedStatusCodes}", resp.StatusCode)}\n\t}\n`;

  content += `\tif !server.PagerResponderMore(${operationStateMachine}) {\n`;
  content += `\t\t${operationStateMachine} = nil\n\t}\n`;
  content += '\treturn resp, nil\n';
  return content;
}

function sanitizeRegexpCaptureGroupName(name: string): string {
  // dash '-' characters are not allowed so replace them with '_'
  return name.replace('-', '_');
}

function createPathParamsRegex(op: Operation): string {
  // "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}"
  // each path param will replaced with a regex capture.
  // note that some path params are optional.
  let urlPath = <string>op.requests![0].protocol.http!.path;
  for (const param of values(op.parameters)) {
    if (param.protocol.http?.in !== 'path') {
      continue;
    }
    const toReplace = `{${param.language.go!.serializedName}}`;
    let replaceWith = `(?P<${sanitizeRegexpCaptureGroupName(param.language.go!.serializedName)}>[a-zA-Z0-9-_]+)`;
    if (param.required === false) {
      replaceWith += '?';
    }
    urlPath = urlPath.replace(toReplace, replaceWith);
  }
  return urlPath;
}

function createParamGroupParams(clientPkg: string, op: Operation, imports: ImportManager): string {
  let content = '';

  // create any param groups and populate their values
  const paramGroups = new Map<GroupProperty, Parameter[]>();
  for (const param of values(consolidateHostParams(aggregateParameters(op)))) {
    if (param.implementation === ImplementationLocation.Client || (param.schema.type === SchemaType.Constant && param.required)) {
      // client params and required constants aren't passed to APIs
      continue;
    }
    if (param.language.go!.paramGroup) {
      if (!paramGroups.has(param.language.go!.paramGroup)) {
        paramGroups.set(param.language.go!.paramGroup, new Array<Parameter>());
      }
      const params = paramGroups.get(param.language.go!.paramGroup);
      params!.push(param);
    }
    if (param.protocol.http?.in === 'body') {
      // body params will be unmarshalled, no need for parsing.
      continue;
    }

    // parse params as required
    if (param.schema.type === SchemaType.Array) {
      if ((<ArraySchema>param.schema).elementType.type !== SchemaType.String) {
        const asArray = <ArraySchema>param.schema;
        imports.add('strings');
        content += `\telements := strings.Split(${getParamValue(param)}, "${getArraySeparator(param)}")\n`;
        const localVar = createLocalVariableName(param, 'Param');
        const toType = `${clientPkg}.${asArray.elementType.language.go!.name}`;
        content += `\t${localVar} := make([]${toType}, len(elements))\n`;
        content += '\tfor i := 0; i < len(elements); i++ {\n';
        let fromVar: string;
        switch (asArray.elementType.type) {
          case SchemaType.Choice:
          case SchemaType.SealedChoice:
            const asChoice = <ChoiceSchema>asArray.elementType;
            // we do support integers/floats as choices but apparently M4 doesn't recognize this.
            switch (<string>asChoice.choiceType.type) {
              case SchemaType.Integer:
                imports.add('strconv');
                fromVar = 'parsedInt';
                content += `\t\t${fromVar}, err := strconv.ParseInt(elements[i], 10, 32)\n`;
                content += '\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n';
                break;
              case SchemaType.Number:
                imports.add('strconv');
                fromVar = 'parsedNum';
                content += `\t\t${fromVar}, err := strconv.ParseFloat(elements[i], 32)\n`;
                content += '\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n';
                break;
              case SchemaType.String:
                fromVar = 'elements[i]';
                break;
              default:
                throw new Error(`unhandled array element choice type ${asChoice.choiceType.type}`);
            }
            break;
          default:
            throw new Error(`unhandled array element type ${asArray.elementType.type}`);
        }
        content += `\t\t${localVar}[i] = ${toType}(${fromVar})\n\t}\n`;
      } else if (param.language.go!.paramGroup) {
        imports.add('strings');
        content += `\t${createLocalVariableName(param, 'Param')} := strings.Split(${getParamValue(param)}, "${getArraySeparator(param)}")\n`;
      }
    } else if (param.schema.type === SchemaType.Boolean) {
      imports.add('strconv');
      let from = `strconv.ParseBool(${getParamValue(param)})`;
      if (!param.required) {
        from = `parseOptional(${getParamValue(param)}, strconv.ParseBool)`;
      }
      content += `\t${createLocalVariableName(param, 'Param')}, err := ${from}\n`;
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    }  else if (param.schema.type === SchemaType.ByteArray) {
      imports.add('encoding/base64');
      content += `\t${createLocalVariableName(param, 'Param')}, err := base64.StdEncoding.DecodeString(${getParamValue(param)})\n`;
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    } else if (param.language.go!.paramGroup && (param.schema.type === SchemaType.Choice || param.schema.type === SchemaType.Constant || param.schema.type === SchemaType.SealedChoice || param.schema.type === SchemaType.String)) {
      content += `\t${createLocalVariableName(param, 'Param')} := `;
      let paramValue = getParamValueWithCast(clientPkg, param, imports);
      if (!param.required) {
        paramValue = `getOptional(${paramValue})`;
      }
      content += `${paramValue}\n`;
    } else if (param.schema.type === SchemaType.Date) {
      imports.add('time');
      let from = `time.Parse("2006-01-02", ${getParamValue(param)})`;
      if (!param.required) {
        from = `parseOptional(${getParamValue(param)}, func(v string) (time.Time, error) { return time.Parse("2006-01-02", v) })`;
      }
      content += `\t${createLocalVariableName(param, 'Param')}, err := ${from}\n`;
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    } else if (param.schema.type === SchemaType.DateTime) {
      imports.add('time');
      const dateTime = <DateTimeSchema>param.schema;
      let format = 'time.RFC3339Nano';
      if (dateTime.format === 'date-time-rfc1123') {
        format = 'time.RFC1123'
      }
      let from = `time.Parse(${format}, ${getParamValue(param)})`;
      if (!param.required) {
        from = `parseOptional(${getParamValue(param)}, func(v string) (time.Time, error) { return time.Parse(${format}, v) })`;
      }
      content += `\t${createLocalVariableName(param, 'Param')}, err := ${from}\n`;
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    } else if (param.schema.type === SchemaType.Integer || param.schema.type === SchemaType.Number) {
      imports.add('strconv');
      const numSchema = <NumberSchema>param.schema;
      let precision = '32';
      if (numSchema.precision !== 32) {
        precision = '64';
      }
      let parseType = 'Int';
      let base = '10, ';
      if (param.schema.type !== SchemaType.Integer) {
        parseType = 'Float';
        base = '';
      }
      let parser = 'parseWithCast';
      if (!param.required) {
        parser = 'parseOptional';
      }
      if (numSchema.precision === 32 || !param.required) {
        content += `\t${createLocalVariableName(param, 'Param')}, err := ${parser}(${getParamValue(param)}, func(v string) (${parseType.toLowerCase()}${precision}, error) {\n`;
        content += `\t\tp, parseErr := strconv.Parse${parseType}(v, ${base}${precision})\n`;
        content += '\t\tif parseErr != nil {\n\t\t\treturn 0, parseErr\n\t\t}\n';
        let result = 'p';
        if (precision === '32') {
          result = `${parseType.toLowerCase()}${precision}(${result})`;
        }
        content += `\t\treturn ${result}, nil\n\t})\n`;
      } else {
        content += `\t${createLocalVariableName(param, 'Param')}, err := strconv.Parse${parseType}(${getParamValue(param)}, ${base}64)\n`;
      }
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    } else if (param.schema.type === SchemaType.UnixTime) {
      imports.add('strconv');
      let from = `strconv.ParseInt(${getParamValue(param)}, 10, 64)`;
      if (!param.required) {
        from = `parseOptional(${getParamValue(param)}, func(v string) (int64, error) { return strconv.ParseInt(v, 10, 64) })`;
      }
      content += `\t${createLocalVariableName(param, 'Param')}, err := ${from}\n`;
      content += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    } else if (param.schema.language.go!.headerCollectionPrefix) {
      imports.add('strings');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
      const localVar = createLocalVariableName(param, 'Param');
      content += `\tvar ${localVar} map[string]*string\n`;
      content += `\tfor hh := range req.Header {\n`;
      const headerPrefix = param.schema.language.go!.headerCollectionPrefix;
      content += `\t\tif len(hh) > len("${headerPrefix}") && strings.EqualFold(hh[:len("x-ms-meta-")], "${headerPrefix}") {\n`;
      content += `\t\t\tif ${localVar} == nil {\n\t\t\t\t${localVar} = map[string]*string{}\n\t\t\t}\n`;
      content += `\t\t\t${localVar}[hh[len("${headerPrefix}"):]] = to.Ptr(getHeaderValue(req.Header, hh))\n`;
      content += `\t\t}\n\t}\n`;
    }
  }

  for (const paramGroup of values(paramGroups.keys())) {
    if (paramGroup.required === true) {
      content += `\t${uncapitalize(paramGroup.language.go!.name)} := ${clientPkg}.${paramGroup.serializedName}{\n`;
      for (const param of values(paramGroups.get(paramGroup))) {
        content += `\t\t${capitalize(param.language.go!.name)}: ${createLocalVariableName(param, 'Param')},\n`;
      }
      content += '\t}\n';
    } else {
      content += `\tvar ${uncapitalize(paramGroup.language.go!.name)} *${clientPkg}.${paramGroup.serializedName}\n`;
      const params = paramGroups.get(paramGroup);
      const paramNilCheck = new Array<string>();
      for (const param of values(params)) {
        // check array before body in case the body is just an array
        if (param.schema.type === SchemaType.Array) {
          paramNilCheck.push(`len(${createLocalVariableName(param, 'Param')}) > 0`);
        } else if (param.protocol.http?.in === 'body') {
          if (param.schema.type === SchemaType.Binary) {
            paramNilCheck.push('req.Body != nil');
          } else {
            imports.add('reflect');
            let bodyParamName = 'body';
            if (param.isPartialBody) {
              bodyParamName = param.language.go!.name;
            }
            paramNilCheck.push(`!reflect.ValueOf(${bodyParamName}).IsZero()`);
          }
        } else {
          paramNilCheck.push(`${createLocalVariableName(param, 'Param')} != nil`);
        }
      }
      content += `\tif ${paramNilCheck.join(' || ')} {\n`;
      content += `\t\t${uncapitalize(paramGroup.language.go!.name)} = &${clientPkg}.${paramGroup.serializedName}{\n`;
      for (const param of values(params)) {
        let byRef = '&';
        if (isTypePassedByValue(param.schema) || (!param.required && param.protocol.http?.in !== 'body')) {
          byRef = '';
        }
        content += `\t\t\t${capitalize(param.language.go!.name)}: ${byRef}${createLocalVariableName(param, 'Param')},\n`;
      }
      content += '\t\t}\n';
      content += '\t}\n';
    }
  }

  return content;
}

function populateApiParams(clientPkg: string, op: Operation, imports: ImportManager): string {
  // FooOperation(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], qp.Get("api-version"), nil)
  // this assumes that our caller has created matches and qp as required
  const params = new Array<string>();

  // for non-paged APIs, first param is always the context. use the one
  // from the HTTP request. be careful to properly handle paged LROs
  if (isLROOperation(op) || !isPageableOperation(op)) {
    params.push('req.Context()');
  }

  // now create the API call sig
  for (const param of values(consolidateHostParams(getMethodParameters(op)))) {
    if (param === op.language.go!.optionalParamGroup) {
      // this is the optional params type. in some cases we just pass nil
      const countParams = values((<GroupProperty>op.language.go!.optionalParamGroup).originalParameter).where((each: Parameter) => { return each.implementation === ImplementationLocation.Method; }).count();
      if (countParams === 0 || (isLROOperation(op) && countParams === 1)) {
        // if the options param is empty or only contains the resume token param just pass nil
        params.push('nil');
        continue;
      }
    }
    imports.addImportForSchemaType(param.schema);
    params.push(`${getParamValueWithCast(clientPkg, param, imports)}`);
  }

  return params.join(', ');
}

function getParamValue(param: Parameter): string {
  if (param.isPartialBody) {
    // multipart form data values have been read and assigned
    // to local params with the same name. must check this first
    // as it's a superset of other cases that follow.
    return param.language.go!.name;
  } else if (param.extensions?.['x-ms-header-collection-prefix']) {
    // must check before http?.in === 'header'
    return createLocalVariableName(param, 'Param');
  } else if (param.schema.type === SchemaType.Constant && param.required) {
    // required constants have their values embedded in the generated code
    return formatConstantValue(<ConstantSchema>param.schema);
  } else if (param.protocol.http?.in === 'path') {
    // path params are in the matches slice
    return `matches[regex.SubexpIndex("${sanitizeRegexpCaptureGroupName(param.language.go!.serializedName)}")]`;
  } else if (param.protocol.http?.in === 'query') {
    // use qp
    return `qp.Get("${param.language.go!.serializedName}")`;
  } else if (param.protocol.http?.in === 'header') {
    // use req
    return `getHeaderValue(req.Header, "${param.language.go!.serializedName}")`;
  } else if (param.protocol.http?.in === 'body') {
    if (param.schema.type === SchemaType.Binary) {
      return 'req.Body.(io.ReadSeekCloser)';
    }
    // JSON/XML bodies have been unmarshalled into a local named body
    return 'body';
  } else if (param.protocol.http?.in === 'uri') {
    return 'req.URL.Host';
  } else if ((<GroupProperty>param).originalParameter !== null) {
    // this is a parameter group param
    return uncapitalize(param.language.go!.name);
  } else {
    throw new Error(`unhandled parameter ${param.language.go!.name} location ${param.protocol.http?.in}`);
  }
}

function getParamValueWithCast(clientPkg: string, param: Parameter, imports: ImportManager): string {
  const value = getParamValue(param);
  if (param.protocol.http?.in === 'body') {
    // if the param is in the body, it's already in the correct format
    if (param.schema.language.go!.internalTimeType) {
      return `time.Time(${value})`;
    }
    return value;
  }
  switch (param.schema.type) {
    case SchemaType.Array:
      const asArray = <ArraySchema>param.schema;
      switch (asArray.elementType.type) {
        case SchemaType.Choice:
        case SchemaType.SealedChoice:
          return createLocalVariableName(param, 'Param');
        case SchemaType.String:
          imports.add('strings');
          return `strings.Split(${value}, "${getArraySeparator(param)}")`;
        default:
          throw new Error(`unhandled array element type ${asArray.elementType.type}`);
      }
    case SchemaType.Boolean:
    case SchemaType.ByteArray:
    case SchemaType.Date:
    case SchemaType.DateTime:
      return createLocalVariableName(param, 'Param');
    case SchemaType.Choice:
    case SchemaType.SealedChoice:
      return `${clientPkg}.${param.schema.language.go!.name}(${value})`;
    case SchemaType.Integer:
    case SchemaType.Number:
      const numSchema = <NumberSchema>param.schema;
      // optional params have been parsed/converted into their required bitness
      if (numSchema.precision === 32 && param.required) {
        let prefix = 'int32';
        if (param.schema.type === SchemaType.Number) {
          prefix = 'float32';
        }
        return `${prefix}(${createLocalVariableName(param, 'Param')})`;
      }
      // parsed ints/numbers have a default type of int64
      return createLocalVariableName(param, 'Param');
    case SchemaType.UnixTime:
      return `time.Unix(${createLocalVariableName(param, 'Param')}, 0)`;
    default:
      return value;
  }
}

function createLocalVariableName(param: Parameter, suffix: string): string {
  if (param.isPartialBody) {
    return param.language.go!.name;
  } else if (param.protocol.http?.in === 'body') {
    if (param.schema.type === SchemaType.Binary) {
      return 'req.Body.(io.ReadSeekCloser)';
    }
    return 'body';
  }
  return `${uncapitalize(param.language.go!.name)}${suffix}`;
}

function getArraySeparator(param: Parameter): string {
  switch (param.protocol.http?.style) {
    case SerializationStyle.PipeDelimited:
      return '|';
    case SerializationStyle.SpaceDelimited:
      return ' ';
    case SerializationStyle.TabDelimited:
      return '\\t';
    default:
      return ',';
  }
}

function consolidateHostParams(params: Parameter[]): Parameter[] {
  if (!values(params).where((each: Parameter) => { return each.protocol.http?.in === 'uri'; }).any()) {
    // no host params
    return params;
  }

  // consolidate multiple host params into a single "host" param
  const consolidatedParams = new Array<Parameter>();
  let hostParamAdded = false;
  for (const param of values(params)) {
    if (param.protocol.http?.in !== 'uri') {
      consolidatedParams.push(param);
    } else if (!hostParamAdded) {
      consolidatedParams.push(param);
      hostParamAdded = true;
    }
  }

  return consolidatedParams;
}

// copied from generator/operations.ts but with a slight tweak to consolidate host parameters
function getAPIParametersSig(op: Operation, imports: ImportManager, pkgName?: string): string {
  const methodParams = consolidateHostParams(getMethodParameters(op));
  const params = new Array<string>();
  if (!isPageableOperation(op) || isLROOperation(op)) {
    imports.add('context');
    params.push('ctx context.Context');
  }
  for (const methodParam of values(methodParams)) {
    let paramName = uncapitalize(methodParam.language.go!.name);
    if (methodParam.protocol.http?.in === 'uri') {
      paramName = 'host';
    }
    params.push(`${paramName} ${formatParameterTypeName(methodParam, pkgName)}`);
  }
  return params.join(', ');
}

// copied from generator/helpers.ts but without the XML-specific stuff
function getResultFieldName(op: Operation): string {
  if (isMultiRespOperation(op)) {
    return 'Value';
  }
  if (op.language.go!.responseEnv.language.go!.resultProp.language.go!.embeddedType) {
    // this is usually the same value as the name of the property except for a few corner-cases
    return op.language.go!.responseEnv.language.go!.resultProp.schema.language.go!.name;
  }
  return op.language.go!.responseEnv.language.go!.resultProp.language.go!.name;
}
