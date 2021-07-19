/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel, ObjectSchema, Property } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { internalPagerTypeName, internalPollerTypeName, PagerInfo, PollerInfo } from '../common/helpers';
import { contentPreamble, getFinalResponseEnvelopeName, getResultFieldName, sortAscending } from './helpers';
import { ImportManager } from './imports';

// Creates the content in pollers.go
export async function generatePollers(session: Session<CodeModel>): Promise<string> {
  if (session.model.language.go!.pollerTypes === undefined) {
    return '';
  }
  let text = await contentPreamble(session);
  const isARM = session.model.language.go!.openApiType === 'arm';
  // add standard imports
  const imports = new ImportManager();
  imports.add('context');
  if (isARM) {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/armcore');
  }
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  imports.add('net/http');
  const pollers = <Array<PollerInfo>>session.model.language.go!.pollerTypes;
  pollers.sort((a: PollerInfo, b: PollerInfo) => { return sortAscending(a.name, b.name) });
  let bodyText = '';
  for (const poller of values(pollers)) {
    // generate the poller interface definition
    bodyText += `// ${poller.name} provides polling facilities until the operation reaches a terminal state.\n`;
    bodyText += `type ${poller.name} interface {\n`;
    bodyText += '\tazcore.Poller\n';
    bodyText += finalResponseDecl(poller);
    bodyText += '}\n\n';
    // now generate the internal poller type
    const pollerName = internalPollerTypeName(poller);
    bodyText += `type ${pollerName} struct {\n`;
    if (isARM) {
      bodyText += '\tpt *armcore.LROPoller\n';
    } else {
      bodyText += '\tpt *azcore.LROPoller\n';
    }
    if (poller.op.language.go!.pageableType) {
      bodyText += `\tclient *${poller.op.language.go!.clientName}\n`;
    }
    bodyText += '}\n\n';
    // internal poller methods
    bodyText += `func (p *${pollerName}) Done() bool {\n\treturn p.pt.Done()\n}\n\n`;
    bodyText += `func (p *${pollerName}) Poll(ctx context.Context) (*http.Response, error) {\n\treturn p.pt.Poll(ctx)\n}\n\n`;
    bodyText += pudFinalResp('FinalResponse', poller, imports);
    bodyText += `func (p *${pollerName}) ResumeToken() (string, error) {\n\treturn p.pt.ResumeToken()\n}\n\n`;
    bodyText += pudFinalResp('pollUntilDone', poller, imports);
  }
  text += imports.text();
  text += bodyText;
  return text;
}

function getResponseType(poller: PollerInfo): string {
  // check for pager must come first
  if (poller.op.language.go!.pageableType) {
    return (<PagerInfo>poller.op.language.go!.pageableType).name;
  }
  return getFinalResponseEnvelopeName(poller.op);
}

function finalResponseDecl(poller: PollerInfo): string {
  const respType = getResponseType(poller);
  let text = '\t// FinalResponse performs a final GET to the service and returns the final response\n';
  text += '\t// for the polling operation. If there is an error performing the final GET then an error is returned.\n';
  text += `\t// If the final GET succeeded then the final ${respType} will be returned.\n`;
  text += `\tFinalResponse(ctx context.Context) (${respType}, error)\n`;
  return text;
}

// generates the pollUntilDone and FinalResponse methods.
// the implementations are almost identical, just a few different params.
function pudFinalResp(op: 'pollUntilDone' | 'FinalResponse', poller: PollerInfo, imports: ImportManager): string {
  let durParam = '';
  let freqParam = '';
  if (op === 'pollUntilDone') {
    imports.add('time');
    durParam = ', freq time.Duration';
    freqParam = ', freq';
  }
  let text = `func (p *${internalPollerTypeName(poller)}) ${op}(ctx context.Context${durParam}) (${getResponseType(poller)}, error) {\n`;
  if (poller.op.language.go!.pageableType) {
    // pager-pollers have a slightly different impl
    text += `\trespType := &${internalPagerTypeName(poller.op.language.go!.pageableType)}{client: p.client}\n`;
    text += `\tif _, err := p.pt.${op.capitalize()}(ctx${freqParam}, &respType.current.${getResultFieldName(poller.op)}); err != nil {\n\t\treturn nil, err\n\t}\n`;
    text += '\treturn respType, nil\n';
  } else {
    const finalRespEnv = <ObjectSchema>poller.op.language.go!.finalResponseEnv;
    const resultEnv = <Property>finalRespEnv.language.go!.resultEnv;
    text += `\trespType := ${finalRespEnv.language.go!.name}{}\n`;
    if (resultEnv) {
      // the operation returns a model of some sort, probe further
      const resultProp = <Property>resultEnv.language.go!.resultField;
      text += `\tresp, err := p.pt.${op.capitalize()}(ctx${freqParam}, &respType.${resultProp.language.go!.name})\n`;
    } else {
      // the operation doesn't return a model
      text += `\tresp, err := p.pt.${op.capitalize()}(ctx${freqParam}, nil)\n`;
    }
    text += `\tif err != nil {\n\t\treturn ${finalRespEnv.language.go!.name}{}, err\n\t}\n`;
    text += '\trespType.RawResponse = resp\n';
    text += '\treturn respType, nil\n';
  }
  text += '}\n\n';
  return text;
}
