/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel, Property } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { PagerInfo, PollerInfo } from '../common/helpers';
import { contentPreamble, discriminatorFinalResponse, getResponseEnvelopeName, sortAscending } from './helpers';
import { ImportManager } from './imports';

// Creates the content in pollers.go
export async function generatePollers(session: Session<CodeModel>): Promise<string> {
  if (session.model.language.go!.pollerTypes === undefined) {
    return '';
  }
  let text = await contentPreamble(session);
  // add standard imports
  const imports = new ImportManager();
  const isARM = session.model.language.go!.openApiType === 'arm';
  imports.add('context');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  imports.add('net/http');
  imports.add('time');
  if (isARM) {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime', 'armruntime');
  } else {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
  }
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

    bodyText += '// Poll fetches the latest state of the LRO.\n';
    bodyText += `// If the LRO has completed successfully, the poller's state is updated and the\n`;
    bodyText += '// response is returned.\n';
    bodyText += `// If the LRO has completed with failure or was cancelled, the poller's state is\n`;
    bodyText += '// updated and the error is returned.\n';
    bodyText += `// If the LRO has not reached a terminal state, the poller's state is updated and\n`;
    bodyText += '// the response is returned.\n';
    bodyText += `// If Poll fails, the poller's state is unmodified and the error is returned.\n`;
    bodyText += '// Calling Poll on an LRO that has reached a terminal state will return the final\n';
    bodyText += '// response or error.\n';
    bodyText += `func (p *${poller.name}) Poll(ctx context.Context) (*http.Response, error) {\n`;
    bodyText += '\treturn p.pt.Poll(ctx)\n}\n\n';

    bodyText += '// Result returns the result of the LRO and is meant to be used in conjunction with Poll and Done.\n';
    bodyText += '// Depending on the operation, calls to Result might perform an additional HTTP GET to fetch the result.\n';
    bodyText += `func (p *${poller.name}) Result(ctx context.Context) (resp ${prefixedResponseType(poller, '*')}, err error) {\n`;
    bodyText += '\t_, err = p.pt.FinalResponse(ctx, &resp)\n';
    bodyText += '\treturn\n}\n\n';

    bodyText += '// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.\n';
    bodyText += '// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.\n'
    if (session.model.language.go!.azureARM) {
      bodyText += '// A good starting value is 30 seconds. Note that some resources might benefit from a different value.\n';
    }
    bodyText += emitPollUntilDone(poller);

    bodyText += '// ResumeToken returns a value representing the poller that can be used to resume\n';
    bodyText += '// the LRO at a later time. ResumeTokens are unique per service operation.\n';
    bodyText += '// Returns an error if the poller is in a terminal state.\n'
    bodyText += `func (p *${poller.name}) ResumeToken() (string, error) {\n\treturn p.pt.ResumeToken()\n}\n\n`;

    bodyText += `// Resume rehydrates a ${poller.name} from the provided client and resume token.\n`;
    bodyText += `// Returns an error if the token is isn't applicable to this poller type.\n`;
    bodyText += emitResume(poller, isARM);
  }
  text += imports.text();
  text += bodyText;
  return text;
}

function emitPollUntilDone(poller: PollerInfo): string {
  let pollUntilDone = `func (p *${poller.name}) PollUntilDone(ctx context.Context, freq time.Duration) (${prefixedResponseType(poller, '*')}, error) {\n`;
  pollUntilDone += `\tresult := ${prefixedResponseType(poller, '&')}{`;
  if (poller.op.language.go!.pageableType) {
    pollUntilDone += 'client: p.client';
  }
  pollUntilDone += '}\n';
  pollUntilDone += `\t_, err := p.pt.PollUntilDone(ctx, freq, ${getResponseField(poller)})\n`;
  if (poller.op.language.go!.pageableType) {
    pollUntilDone += '\tif err != nil {\n\treturn nil, err\n\t}\n';
  }
  pollUntilDone += '\treturn result, err\n';
  pollUntilDone += '}\n\n';
  return pollUntilDone;
}

function emitResume(poller: PollerInfo, isARM: boolean): string {
  const clientName = poller.op.language.go!.clientName;
  const apiMethod = poller.op.language.go!.name;
  let resume = `func (p *${poller.name}) Resume(token string, client *${clientName}) (err error) {\n`;
  resume += '\tp.pt, err = ';
  if (isARM) {
    resume += `armruntime.`;
  } else {
    resume += `runtime.`;
  }
  resume += `NewPollerFromResumeToken("${clientName}.${apiMethod}", token, client.pl)\n`;
  if (poller.op.language.go!.pageableType) {
    resume += '\tp.client = client\n';
  }
  resume += '\treturn\n';
  resume += '}\n\n';
  return resume;
}

function getResponseType(poller: PollerInfo): string {
  // check for pager must come first
  if (poller.op.language.go!.pageableType) {
    return (<PagerInfo>poller.op.language.go!.pageableType).name;
  }
  return getResponseEnvelopeName(poller.op);
}

function getResponseField(poller: PollerInfo): string {
  const pagedResponse = poller.op.language.go!.pageableType;
  const finalRespEnv = poller.op.language.go!.responseEnv;
  const resultProp = <Property>finalRespEnv.language.go!.resultProp;
  if (resultProp) {
    let current = '';
    if (pagedResponse) {
      current = '.current';
    }
    return `&result${current}${discriminatorFinalResponse(finalRespEnv)}`;
  } else {
    // the operation doesn't return a model
    return 'nil';
  }
}

// returned the response type with the specified prefix if it's a pager
function prefixedResponseType(poller: PollerInfo, prefix: string): string {
  const pagedResponse = poller.op.language.go!.pageableType;
  const respType = getResponseType(poller);
  if (!pagedResponse) {
    return respType;
  }
  return prefix + respType;
}
