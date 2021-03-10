/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { camelCase } from '@azure-tools/codegen';
import { CodeModel, SchemaType, Schema, ArraySchema } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { internalPagerTypeName, internalPollerTypeName, PollerInfo, PagerInfo } from '../common/helpers';
import { contentPreamble, sortAscending } from './helpers';
import { ImportManager } from './imports';

function generatePagerReturnInstance(pager: PagerInfo): string {
  let text = '';
  text += `\treturn &${internalPagerTypeName(pager)}{\n`;
  text += `\t\tpipeline: p.pipeline,\n`;
  text += `\t\tresp: resp,\n`;
  text += '\t\terrorer: p.errHandler,\n';
  text += `\t\tresponder: p.respHandler,\n`;
  text += `\t\tadvancer: func(ctx context.Context, resp ${pager.respEnv}) (*azcore.Request, error) {\n`;
  text += `\t\t\treturn azcore.NewRequest(ctx, http.MethodGet, *resp.${pager.respField}.${pager.nextLink})\n`;
  text += '\t\t},\n';
  text += `\t\tstatusCodes: p.statusCodes,\n`;
  text += `\t}, nil`;
  return text;
}

// Creates the content in pollers.go
export async function generateARMPollers(session: Session<CodeModel>): Promise<string> {
  if (session.model.language.go!.pollerTypes === undefined) {
    return '';
  }
  let text = await contentPreamble(session);

  // add standard imports
  const imports = new ImportManager();
  imports.add('context');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/armcore');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  imports.add('net/http');
  imports.add('time');
  let bodyText = '';
  const pollers = <Array<PollerInfo>>session.model.language.go!.pollerTypes;
  pollers.sort((a: PollerInfo, b: PollerInfo) => { return sortAscending(a.name, b.name) });
  for (const poller of values(pollers)) {
    const pollerInterface = poller.name;
    const pollerName = internalPollerTypeName(poller);
    let responseType = 'HTTPResponse';
    // HTTP Pollers do not need to perform the final get request since they do not return a model
    let finalResponseDeclaration = 'FinalResponse(ctx context.Context) (*http.Response, error)';
    let finalResponse = `${finalResponseDeclaration} {
        return p.pt.FinalResponse(ctx, p.pipeline, nil)
      }`;
    let pollUntilDoneResponse = '(*http.Response, error)';
    let pollUntilDone = `return p.pt.PollUntilDone(ctx, frequency, p.pipeline, nil)`;
    let handleResponse = '';
    let pagerFields = '';
    if (poller.pager) {
      function finalPagerProcessing(name: string, params: string): string {
        return `respType := &${camelCase(responseType)}{}
                resp, err := p.pt.${name}(${params})
                if err != nil {
                  return nil, err
                }
                return p.handleResponse(&azcore.Response{Response: resp})`;
      }
      responseType = poller.pager.name;
      pollUntilDoneResponse = `(${responseType}, error)`;
      // for operations that do return a model add a final response method that handles the final get URL scenario
      finalResponseDeclaration = `FinalResponse(ctx context.Context) (${responseType}, error)`;
      pagerFields = `
      errHandler  ${camelCase(poller.pager.respType)}HandleError
      respHandler ${camelCase(poller.pager.respType)}HandleResponse
      statusCodes []int`;
      handleResponse = `
      func (p *${pollerName}) handleResponse(resp *azcore.Response) (${responseType}, error) {
        ${generatePagerReturnInstance(poller.pager)}
      }
      `;
      finalResponse = `${finalResponseDeclaration} {
      ${finalPagerProcessing('FinalResponse', 'ctx, p.pipeline, respType')}
	  }
      `;
      pollUntilDone = finalPagerProcessing('PollUntilDone', 'ctx, frequency, p.pipeline, respType');
    } else if (poller.respType) {
      responseType = poller.respEnv;
      pollUntilDoneResponse = `(${responseType}, error)`;
      // for operations that do return a model add a final response method that handles the final get URL scenario
      finalResponseDeclaration = `FinalResponse(ctx context.Context) (${responseType}, error)`;
      finalResponse = `FinalResponse(ctx context.Context) (${responseType}, error) {`;
      let reference = '';
      let respByRef = '&';
      if (poller.respType.type === SchemaType.Array || poller.respType.type === SchemaType.Dictionary) {
        // arrays and maps are returned by value
        respByRef = '';
        // but we need to pass them by reference to the unmarshaller
        reference = '&';
      }
      let respType = `respType := ${responseType}{${poller.respField}: ${respByRef}${poller.respType.language.go!.name}{}}`;
      const isScalar = isScalarType(poller.respType);
      if (isScalar) {
        respType = `respType := ${responseType}{}\n`;
        reference = '&';
      }
      pollUntilDone = `${respType}
		resp, err := p.pt.PollUntilDone(ctx, frequency, p.pipeline, ${reference}respType.${poller.respField})
		if err != nil {
			return ${responseType}{}, err
    }
    respType.RawResponse = resp
    return respType, nil`;
      finalResponse += `
      ${respType}
		resp, err := p.pt.FinalResponse(ctx, p.pipeline, ${reference}respType.${poller.respField})
		if err != nil {
			return ${responseType}{}, err
    }
    respType.RawResponse = resp
		return respType, nil
	  }
      `;
    }
    bodyText += `// ${pollerInterface} provides polling facilities until the operation completes
      type ${pollerInterface} interface {
        azcore.Poller

        // FinalResponse performs a final GET to the service and returns the final response
        // for the polling operation. If there is an error performing the final GET then an error is returned.
        // If the final GET succeeded then the final ${responseType} will be returned.
        ${finalResponseDeclaration}
      }

      type ${pollerName} struct {
        // the client for making the request
        pipeline azcore.Pipeline${pagerFields}
        pt armcore.Poller
      }

      // Done returns true if there was an error or polling has reached a terminal state
      func (p *${pollerName}) Done() bool {
        return p.pt.Done()
      }

      // Poll will send poll the service endpoint and return an http.Response or error received from the service
      func (p *${pollerName}) Poll(ctx context.Context) (*http.Response, error) {
        return p.pt.Poll(ctx, p.pipeline)
      }

      func (p *${pollerName}) ${finalResponse}

      // ResumeToken generates the string token that can be used with the Resume${pollerInterface} method
      // on the client to create a new poller from the data held in the current poller type
      func (p *${pollerName}) ResumeToken() (string, error) {
        return p.pt.ResumeToken()
      }

      func (p *${pollerName}) pollUntilDone(ctx context.Context, frequency time.Duration) ${pollUntilDoneResponse} {
      ${pollUntilDone}
      }
  ${handleResponse}
  `;
  }
  text += imports.text();
  text += bodyText;
  return text;
}

function isScalarType(schema: Schema): boolean {
  switch (schema.type) {
    case SchemaType.Array:
      return isScalarType((<ArraySchema>schema).elementType);
    case SchemaType.Boolean:
    case SchemaType.ByteArray:
    case SchemaType.Choice:
    case SchemaType.Duration:
    case SchemaType.Integer:
    case SchemaType.Number:
    case SchemaType.SealedChoice:
    case SchemaType.String:
    case SchemaType.Time:
    case SchemaType.UnixTime:
    case SchemaType.Uri:
    case SchemaType.Uuid:
      return true;
    default:
      return false;
  }
}
