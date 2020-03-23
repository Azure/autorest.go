/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { camelCase } from '@azure-tools/codegen';
import { CodeModel } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { ContentPreamble, ImportManager, PagerInfo, SortAscending } from './helpers';

// Creates the content in pagers.go
export async function generatePagers(session: Session<CodeModel>): Promise<string> {
  if (session.model.language.go!.pageableTypes === undefined) {
    return '';
  }
  let text = await ContentPreamble(session);

  // add standard imports
  const imports = new ImportManager();
  imports.add('context');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  text += imports.text();

  const pagers = <Array<PagerInfo>>session.model.language.go!.pageableTypes;
  pagers.sort((a: PagerInfo, b: PagerInfo) => { return SortAscending(a.name, b.name) });
  for (const pager of values(pagers)) {
    const pagerType = camelCase(pager.name);
    const responseType = pager.schema.language.go!.responseType.name;
    const resultType = pager.schema.language.go!.name;
    const responderType = `${camelCase(resultType)}HandleResponse`;
    const advanceType = `${camelCase(resultType)}AdvancePage`;
    text += `// ${pager.name} provides iteration over ${resultType} pages.
type ${pager.name} interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current ${responseType}.
	PageResponse() *${responseType}

	// Err returns the last error encountered while paging.
	Err() error
}

type ${responderType} func(*azcore.Response) (*${responseType}, error)

type ${advanceType} func(*${responseType}) (*azcore.Request, error)

type ${pagerType} struct {
	cli *${pager.client}
	req *azcore.Request
	hnd ${responderType}
	adv ${advanceType}
	cur *${responseType}
	err error
}

func (p *${pagerType}) Err() error {
	return p.err
}

func (p *${pagerType}) NextPage(ctx context.Context) bool {
	if p.cur != nil {
		if p.cur.${resultType}.${pager.nextLink} == nil || len(*p.cur.${resultType}.${pager.nextLink}) == 0 {
			return false
		}
		req, err := p.adv(p.cur)
		if err != nil {
			p.err = err
			return false
		}
		p.req = req
	}
	resp, err := p.cli.p.Do(ctx, p.req)
	if err != nil {
		p.err = err
		return false
	}
	result, err := p.hnd(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.cur = result
	return true
}

func (p *${pagerType}) PageResponse() *${responseType} {
	return p.cur
}

`;
  }
  return text;
}
