/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { camelCase } from '@azure-tools/codegen';
import { CodeModel, SchemaResponse } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { PollerInfo, isSchemaResponse, isPageableOperation } from '../common/helpers';
import { contentPreamble, sortAscending } from './helpers';
import { ImportManager } from './imports';
import { generateARMPollers } from './arm_pollers';

// Creates the content in pollers.go
export async function generatePollers(session: Session<CodeModel>): Promise<string> {
  // get the openapi-type value specified. Default to ARM behavior, unless "data-plane" is specified
  let openapiType = await session.getValue('openapi-type', '');
  if (openapiType !== 'data-plane') {
    return generateARMPollers(session);
  }

  if (session.model.language.go!.pollerTypes === undefined) {
    return '';
  }
  let text = await contentPreamble(session);

  // add standard imports
  const imports = new ImportManager();
  imports.add('context');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  imports.add('net/http');
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
        return nil
      }`;
    let pollUntilDoneResponse = '(*http.Response, error)';
    let pollUntilDoneReturn = 'p.FinalResponse(), nil';
    let handleResponse = '';
    const schemaResponse = <SchemaResponse>poller.op.responses![0];
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
        return nil, nil
      }
      `;
      finalResponse = `${finalResponseDeclaration} {`;
      finalResponseCheckNeeded = true;
    } else if (isSchemaResponse(schemaResponse) && schemaResponse.schema.language.go!.responseType.name !== undefined) {
      responseType = schemaResponse.schema.language.go!.responseType.name;
      pollUntilDoneResponse = `(*${responseType}, error)`;
      pollUntilDoneReturn = 'p.FinalResponse(ctx)';
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
