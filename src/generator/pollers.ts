/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel, ObjectSchema, Property } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { PagerInfo, PollerInfo } from '../common/helpers';
import { contentPreamble, discriminatorFinalResponse, getFinalResponseEnvelopeName, getResultFieldName, sortAscending } from './helpers';
import { ImportManager } from './imports';

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
  const pollers = <Array<PollerInfo>>session.model.language.go!.pollerTypes;
  pollers.sort((a: PollerInfo, b: PollerInfo) => { return sortAscending(a.name, b.name) });
  let bodyText = '';
  for (const poller of values(pollers)) {
    // generate the poller type
    bodyText += `// ${poller.name} provides polling facilities until the operation reaches a terminal state.\n`;
    bodyText += `type ${poller.name} struct {\n`;
    bodyText += '\tpt *azcore.Poller\n';
    if (poller.op.language.go!.pageableType) {
      bodyText += `\tclient *${poller.op.language.go!.clientName}\n`;
    }
    bodyText += '}\n\n';
    // poller methods
    bodyText += '// Done returns true if the LRO has reached a terminal state.\n';
    bodyText += `func (p *${poller.name}) Done() bool {\n\treturn p.pt.Done()\n}\n\n`;
    bodyText += '// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.\n';
    bodyText += `// If the LRO has completed successfully, the poller's state is updated and the HTTP\n`;
    bodyText += '// response is returned.\n';
    bodyText += `// If the LRO has completed with failure or was cancelled, the poller's state is\n`;
    bodyText += '// updated and the error is returned.\n';
    bodyText += `// If the LRO has not reached a terminal state, the poller's state is updated and\n`;
    bodyText += '// the latest HTTP response is returned.\n';
    bodyText += `// If Poll fails, the poller's state is unmodified and the error is returned.\n`;
    bodyText += '// Calling Poll on an LRO that has reached a terminal state will return the final\n';
    bodyText += '// HTTP response or error.\n';
    bodyText += `func (p *${poller.name}) Poll(ctx context.Context) (*http.Response, error) {\n\treturn p.pt.Poll(ctx)\n}\n\n`;
    bodyText += finalResp(poller);
    bodyText += '// ResumeToken returns a value representing the poller that can be used to resume\n';
    bodyText += '// the LRO at a later time. ResumeTokens are unique per service operation.\n';
    bodyText += `func (p *${poller.name}) ResumeToken() (string, error) {\n\treturn p.pt.ResumeToken()\n}\n\n`;
  }
  text += imports.text();
  text += bodyText;
  return text;
}

function getResponseType(poller: PollerInfo): string {
  // check for pager must come first
  if (poller.op.language.go!.pageableType) {
    return `*${(<PagerInfo>poller.op.language.go!.pageableType).name}`;
  }
  return getFinalResponseEnvelopeName(poller.op);
}

// generates the FinalResponse method
function finalResp(poller: PollerInfo): string {
  const respType = getResponseType(poller);
  let text = '\t// FinalResponse performs a final GET to the service and returns the final response\n';
  text += '\t// for the polling operation. If there is an error performing the final GET then an error is returned.\n';
  text += `\t// If the final GET succeeded then the final ${respType} will be returned.\n`;
  text += `func (p *${poller.name}) FinalResponse(ctx context.Context) (${respType}, error) {\n`;
  if (poller.op.language.go!.pageableType) {
    // pager-pollers have a slightly different impl
    text += `\trespType := &${(<PagerInfo>poller.op.language.go!.pageableType).name}{client: p.client}\n`;
    text += `\tif _, err := p.pt.FinalResponse(ctx, &respType.current.${getResultFieldName(poller.op)}); err != nil {\n\t\treturn nil, err\n\t}\n`;
    text += '\treturn respType, nil\n';
  } else {
    const finalRespEnv = <ObjectSchema>poller.op.language.go!.finalResponseEnv;
    const resultProp = <Property>finalRespEnv.language.go!.resultProp;
    text += `\trespType := ${finalRespEnv.language.go!.name}{}\n`;
    if (resultProp) {
      text += `\t_, err := p.pt.FinalResponse(ctx, &respType${discriminatorFinalResponse(finalRespEnv)})\n`;
    } else {
      // the operation doesn't return a model
      text += `\t_, err := p.pt.FinalResponse(ctx, nil)\n`;
    }
    text += `\tif err != nil {\n\t\treturn ${finalRespEnv.language.go!.name}{}, err\n\t}\n`;
    text += '\treturn respType, nil\n';
  }
  text += '}\n\n';
  return text;
}
