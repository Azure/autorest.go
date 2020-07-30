/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { camelCase } from '@azure-tools/codegen';
import { CodeModel, SchemaResponse, SchemaType, Operation } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { PagerInfo, PollerInfo, isSchemaResponse, isPageableOperation } from '../common/helpers';
import { contentPreamble, sortAscending, getCreateRequestParametersSig, getMethodParameters } from './helpers';
import { ImportManager } from './imports';
import { OperationNaming } from '../transform/namer';

function getPutCheck(op: Operation): string {
  const respSchema = <SchemaResponse>op.responses![0];
  let text = 'if p.pt.pollerMethodVerb() == http.MethodPut || p.pt.pollerMethodVerb() == http.MethodPatch {';
  if (!isPageableOperation(op)) {
    text += `
    res, err := p.handleResponse(p.pt.latestResponse())
    if err != nil {
      return nil, err
    }
    `;
    switch (respSchema.schema.type) {
      case SchemaType.Array:
      case SchemaType.Dictionary:
        text += `if res != nil && res.${respSchema.schema.language.go!.responseType.value} != nil {`;
        break;
      case SchemaType.String:
        text += `if res != nil && (*res.${respSchema.schema.language.go!.responseType.value} != "") {`;
        break;
      default:
        text += `if res != nil && (*res.${respSchema.schema.language.go!.responseType.value} != ${respSchema.schema.language.go!.responseType.value}{}) {`;
    }
    text += `
        return res, nil
      }`;
  } else {
    text += 'return p.handleResponse(p.pt.latestResponse())';
  }
  text += '}';
  return text;
}

function generatePagerReturnInstance(op: Operation, imports: ImportManager): string {
  let text = '';
  const info = <OperationNaming>op.language.go!;
  // split param list into individual params
  const reqParams = getCreateRequestParametersSig(op).split(',');
  // keep the parameter names from the name/type tuples
  for (let i = 0; i < reqParams.length; ++i) {
    reqParams[i] = reqParams[i].trim().split(' ')[0];
  }
  text += `\treturn &${camelCase(op.language.go!.pageableType.name)}{\n`;
  text += `\t\tpipeline: p.pipeline,\n`;
  text += `\t\tresp: resp,\n`;
  text += `\t\tresponder: p.respHandler,\n`;
  const pager = <PagerInfo>op.language.go!.pageableType;
  const pagerSchema = <SchemaResponse>pager.op.responses![0];
  if (op.language.go!.paging.member) {
    // find the location of the nextLink param
    const nextLinkOpParams = getMethodParameters(op.language.go!.paging.nextLinkOperation);
    let found = false;
    for (let i = 0; i < nextLinkOpParams.length; ++i) {
      if (nextLinkOpParams[i].schema.type === SchemaType.String && nextLinkOpParams[i].language.go!.name.startsWith('next')) {
        // found it
        reqParams.splice(i, 0, `*resp.${pagerSchema.schema.language.go!.name}.${pager.op.language.go!.paging.nextLinkName}`);
        found = true;
        break;
      }
    }
    if (!found) {
      throw console.error(`failed to find nextLink parameter for operation ${op.language.go!.paging.nextLinkOperation.language.go!.name}`);
    }
    text += `\t\tadvancer: func(resp *${pagerSchema.schema.language.go!.responseType.name}) (*azcore.Request, error) {\n`;
    text += `\t\t\treturn client.${camelCase(op.language.go!.paging.member)}CreateRequest(${reqParams.join(', ')})\n`;
    text += '\t\t},\n';
  } else {
    imports.add('fmt');
    imports.add('net/url');
    let resultTypeName = pagerSchema.schema.language.go!.name;
    if (pagerSchema.schema.serialization?.xml?.name) {
      // xml can specifiy its own name, prefer that if available
      resultTypeName = pagerSchema.schema.serialization.xml.name;
    }
    text += `\t\tadvancer: func(resp *${pagerSchema.schema.language.go!.responseType.name}) (*azcore.Request, error) {\n`;
    text += `\t\t\tu, err := url.Parse(*resp.${resultTypeName}.${pager.op.language.go!.paging.nextLinkName})\n`;
    text += `\t\t\tif err != nil {\n`;
    text += `\t\t\t\treturn nil, fmt.Errorf("invalid ${pager.op.language.go!.paging.nextLinkName}: %w", err)\n`;
    text += `\t\t\t}\n`;
    text += `\t\t\tif u.Scheme == "" {\n`;
    text += `\t\t\t\treturn nil, fmt.Errorf("no scheme detected in ${pager.op.language.go!.paging.nextLinkName} %s", *resp.${resultTypeName}.${pager.op.language.go!.paging.nextLinkName})\n`;
    text += `\t\t\t}\n`;
    text += `\t\t\treturn azcore.NewRequest(http.MethodGet, *u), nil\n`;
    text += `\t\t},\n`;
  }
  text += `\t}, nil`;
  return text;
}

// Creates the content in pollers.go
export async function generatePollers(session: Session<CodeModel>): Promise<string> {
  if (session.model.language.go!.pollerTypes === undefined) {
    return '';
  }
  let text = await contentPreamble(session);

  // add standard imports
  const imports = new ImportManager();
  imports.add('context');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  imports.add('net/http');
  imports.add('errors');
  imports.add('encoding/json');
  imports.add('net/url');
  imports.add('time');
  let bodyText = '';
  const pollers = <Array<PollerInfo>>session.model.language.go!.pollerTypes;
  pollers.sort((a: PollerInfo, b: PollerInfo) => { return sortAscending(a.name, b.name) });
  for (const poller of values(pollers)) {
    const pollerInterface = poller.name;
    const pollerName = camelCase(poller.name);
    let responseType = 'HTTPResponse';
    // HTTP Pollers do not need to perform the final get request since they do not return a model
    let finalResponseDeclaration = 'FinalResponse() *http.Response';
    let finalResponse = `${finalResponseDeclaration} {
        return p.pt.latestResponse().Response;
      }`;
    let pollUntilDoneResponse = '(*http.Response, error)';
    let pollUntilDoneReturn = 'p.FinalResponse(), nil';
    let handleResponse = '';
    const schemaResponse = <SchemaResponse>poller.op.responses![0];
    let unmarshalResponse = 'nil';
    let pagerFields = '';
    let finalResponseCheckNeeded = false;
    if (isPageableOperation(poller.op)) {
      responseType = poller.op.language.go!.pageableType.name;
      pollUntilDoneResponse = `(${responseType}, error)`;
      pollUntilDoneReturn = 'p.FinalResponse(ctx)';
      // for operations that do return a model add a final response method that handles the final get URL scenario
      finalResponseDeclaration = `FinalResponse(ctx context.Context) (${responseType}, error)`;
      pagerFields = `
      respHandler ${camelCase(poller.op.language.go!.pageableType.op.responses![0].schema.language.go!.name)}HandleResponse`;
      handleResponse = `
      func (p *${pollerName}) handleResponse(resp *azcore.Response) (${responseType}, error) {
        ${generatePagerReturnInstance(poller.op, imports)}
      }
      `;
      finalResponse = `${finalResponseDeclaration} {`;
      finalResponseCheckNeeded = true;
    } else if (isSchemaResponse(schemaResponse) && schemaResponse.schema.language.go!.responseType.name !== undefined) {
      responseType = schemaResponse.schema.language.go!.responseType.name;
      pollUntilDoneResponse = `(*${responseType}, error)`;
      pollUntilDoneReturn = 'p.FinalResponse(ctx)';
      unmarshalResponse = `resp.UnmarshalAsJSON(&result.${schemaResponse.schema.language.go!.responseType.value})`;
      // for operations that do return a model add a final response method that handles the final get URL scenario
      finalResponseDeclaration = `FinalResponse(ctx context.Context) (*${responseType}, error)`;
      handleResponse = `
      func (p *${pollerName}) handleResponse(resp *azcore.Response) (*${responseType}, error) {
        return nil, nil
      }
      `;
      finalResponse = `FinalResponse(ctx context.Context) (*${responseType}, error) {`;
      finalResponseCheckNeeded = true;
    }
    if (finalResponseCheckNeeded) {
      finalResponse += `
       return nil, nil
      }`;
    }
    bodyText += `// ${pollerInterface} provides polling facilities until the operation completes
      type ${pollerInterface} interface {
        Done() bool
        Poll(ctx context.Context) (*http.Response, error)
        ${finalResponseDeclaration}
        ResumeToken() (string, error)
      }

      type ${pollerName} struct {
        // the client for making the request
        pipeline azcore.Pipeline${pagerFields}
        pt pollingTracker
      }

      // Done returns true if there was an error or polling has reached a terminal state
      func (p *${pollerName}) Done() bool {
        return false
      }

      // Poll will send poll the service endpoint and return an http.Response or error received from the service
      func (p *${pollerName}) Poll(ctx context.Context) (*http.Response, error) {
        return nil, nil
      }

      func (p *${pollerName}) ${finalResponse}

      // ResumeToken generates the string token that can be used with the Resume${pollerInterface} method
      // on the client to create a new poller from the data held in the current poller type
      func (p *${pollerName}) ResumeToken() (string, error) {
        return "", nil
      }
  `;
  }
  text += '// TODO replace this code with data-plane specific definitions\n';
  text += imports.text();
  text += bodyText;
  return text;
}


// Creates the content in pollers_helper.go
export async function generatePollersHelper(session: Session<CodeModel>): Promise<string> {
  if (session.model.language.go!.pollerTypes === undefined) {
    return '';
  }
  let text = await contentPreamble(session);
  const pollers = <Array<PollerInfo>>session.model.language.go!.pollerTypes;
  pollers.sort((a: PollerInfo, b: PollerInfo) => { return sortAscending(a.name, b.name) });
  text += '// TODO replace this code with data-plane specific definitions\n';
  return text;
}

