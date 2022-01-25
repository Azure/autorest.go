/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { isLROOperation, PagerInfo } from '../common/helpers';
import { contentPreamble, getResponseEnvelopeName, getResultFieldName, getStatusCodes, formatStatusCodes, sortAscending, getFinalResponseEnvelopeName } from './helpers';
import { ImportManager } from './imports';

// Creates the content in pagers.go
export async function generatePagers(session: Session<CodeModel>): Promise<string> {
  if (session.model.language.go!.pageableTypes === undefined) {
    return '';
  }
  let text = await contentPreamble(session);

  // add standard imports
  const imports = new ImportManager();
  imports.add('context');
  imports.add('errors');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/policy');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
  imports.add('net/http');
  imports.add('reflect');
  text += imports.text();

  const pagers = <Array<PagerInfo>>session.model.language.go!.pageableTypes;
  pagers.sort((a: PagerInfo, b: PagerInfo) => { return sortAscending(a.name, b.name) });
  for (const pager of values(pagers)) {
    let respEnv = getResponseEnvelopeName(pager.op);
    if (isLROOperation(pager.op)) {
      respEnv = getFinalResponseEnvelopeName(pager.op);
    }
    // create pager type
    text += `// ${pager.name} provides operations for iterating over paged responses.\n`;
    text += `type ${pager.name} struct {\n`;
    text += `\tclient *${pager.op.language.go!.clientName}\n`;
    text += `\tcurrent ${respEnv}\n`;
    if (isLROOperation(pager.op)) {
      text += '\tsecond bool\n';
    } else {
      text += '\trequester func(context.Context) (*policy.Request, error)\n';
      text += `\tadvancer func(context.Context, ${respEnv}) (*policy.Request, error)\n`;
    }
    text += '}\n\n';
    // pager methods
    const nextLinkField = `${getResultFieldName(pager.op)}.${pager.op.language.go!.paging.nextLinkName}`;
    text += '// More returns true if there are more pages to retrieve.\n';
    text += `func (p *${pager.name}) More() bool {\n`;
    text += '\tif !reflect.ValueOf(p.current).IsZero() {\n';
    text += `\t\tif p.current.${nextLinkField} == nil || len(*p.current.${nextLinkField}) == 0 {\n`;
    text += '\t\t\treturn false\n\t\t}\n';
    text += '\t}\n\treturn true\n';
    text += '}\n\n';

    text += '// NextPage advances the pager to the next page.\n'
    text += `func (p *${pager.name}) NextPage(ctx context.Context) (${respEnv}, error) {\n`;
    if (isLROOperation(pager.op)) {
      text += '\tif !p.second {\n';
      text += '\t\tp.second = true\n';
      text += '\t\treturn p.current, nil\n';
      text += '\t} else ';
    } else {
      // note the trailing tab for the next line
      text += '\tvar req *policy.Request\n\tvar err error\n\t';
    }
    text += 'if !reflect.ValueOf(p.current).IsZero() {\n';
    text += `\t\tif !p.More() {\n`;
    text += `\t\t\treturn ${respEnv}{}, errors.New("no more pages")\n\t\t}\n`;
    if (isLROOperation(pager.op)) {
      text += `\t}\n\treq, err := runtime.NewRequest(ctx, http.MethodGet, *p.current.${nextLinkField})\n`;
    } else {
      text += '\t\treq, err = p.advancer(ctx, p.current)\n';
      text += '\t} else {\n';
      text += '\t\treq, err = p.requester(ctx)\n\t}\n';
    }
    text += `\tif err != nil {\n\t\treturn ${respEnv}{}, err\n\t}\n`;
    text += `\tresp, err := p.client.pl.Do(req)\n`;
    text += `\tif err != nil {\n\t\treturn ${respEnv}{}, err\n\t}\n`;
    let statusCodes: string;
    if (isLROOperation(pager.op)) {
      // 204 no content excluded because why would you get a 204 for paged results?
      statusCodes = 'http.StatusOK, http.StatusCreated, http.StatusAccepted';
    } else {
      statusCodes = formatStatusCodes(getStatusCodes(pager.op));
    }
    text += `\tif !runtime.HasStatusCode(resp, ${statusCodes}) {\n`;
    text += `\n\t\treturn ${respEnv}{}, runtime.NewResponseError(resp)\n\t}\n`;
    text += `\tresult, err := p.client.${pager.op.language.go!.protocolNaming.responseMethod}(resp)\n`;
    text += `\tif err != nil {\n\t\treturn ${respEnv}{}, err\n\t}\n`;
    text += '\tp.current = result\n\treturn p.current, nil\n';
    text += '}\n\n';
  }
  return text;
}
