/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { internalPagerTypeName, isLROOperation, PagerInfo } from '../common/helpers';
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
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
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
    // create pager interface
    text += `type ${pager.name} interface {\n`;
    text += '\tazcore.Pager\n';
    text += `\t// PageResponse returns the current ${respEnv}.\n`;
    text += `\tPageResponse() ${respEnv}\n`;
    text += '}\n\n';
    // now create internal pager type
    const internalPager = internalPagerTypeName(pager);
    text += `type ${internalPager} struct {\n`;
    text += `\tclient *${pager.op.language.go!.clientName}\n`;
    text += `\tcurrent ${respEnv}\n`;
    text += '\terr error\n';
    if (isLROOperation(pager.op)) {
      text += '\tsecond bool\n';
    } else {
      text += '\trequester func(context.Context) (*azcore.Request, error)\n';
      text += `\tadvancer func(context.Context, ${respEnv}) (*azcore.Request, error)\n`;
    }
    text += '}\n\n';
    // internal pager methods
    text += `func (p *${internalPager}) Err() error {\n\treturn p.err\n}\n\n`;
    text += `func (p *${internalPager}) NextPage(ctx context.Context) bool {\n`;
    if (isLROOperation(pager.op)) {
      text += '\tif !p.second {\n';
      text += '\t\tp.second = true\n';
      text += '\t\treturn true\n';
      text += '\t} else ';
    } else {
      // note the trailing tab for the next line
      text += '\tvar req *azcore.Request\n\tvar err error\n\t';
    }
    text += 'if !reflect.ValueOf(p.current).IsZero() {\n';
    const nextLinkField = `${getResultFieldName(pager.op)}.${pager.op.language.go!.paging.nextLinkName}`;
    text += `\t\tif p.current.${nextLinkField} == nil || len(*p.current.${nextLinkField}) == 0 {\n`;
    text += '\t\t\treturn false\n\t\t}\n';
    if (isLROOperation(pager.op)) {
      text += `\t}\n\treq, err := azcore.NewRequest(ctx, http.MethodGet, *p.current.${nextLinkField})\n`;
    } else {
      text += '\t\treq, err = p.advancer(ctx, p.current)\n';
      text += '\t} else {\n';
      text += '\t\treq, err = p.requester(ctx)\n\t}\n';
    }
    text += '\tif err != nil {\n\t\tp.err = err\n\t\treturn false\n\t}\n';
    text += '\tresp, err := p.client.con.Pipeline().Do(req)\n';
    text += '\tif err != nil {\n\t\tp.err = err\n\t\treturn false\n\t}\n';
    let statusCodes: string;
    if (isLROOperation(pager.op)) {
      // 204 no content excluded because why would you get a 204 for paged results?
      statusCodes = 'http.StatusOK, http.StatusCreated, http.StatusAccepted';
    } else {
      statusCodes = formatStatusCodes(getStatusCodes(pager.op));
    }
    text += `\tif !resp.HasStatusCode(${statusCodes}) {\n`;
    text += `\t\tp.err = p.client.${pager.op.language.go!.protocolNaming.errorMethod}(resp)\n\t\treturn false\n\t}\n`;
    text += `\tresult, err := p.client.${pager.op.language.go!.protocolNaming.responseMethod}(resp)\n`;
    text += '\tif err != nil {\n\t\tp.err = err\n\t\treturn false\n\t}\n';
    text += '\tp.current = result\n\treturn true\n';
    text += '}\n\n';
    text += `func (p *${internalPager}) PageResponse() ${respEnv} {\n\treturn p.current\n}\n\n`;
  }
  return text;
}
