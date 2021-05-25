/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { internalPagerTypeName, PagerInfo } from '../common/helpers';
import { contentPreamble, sortAscending } from './helpers';
import { ImportManager } from './imports';
import { ensureNameCase } from '../transform/namer';

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
  imports.add('reflect');
  text += imports.text();

  const pagers = <Array<PagerInfo>>session.model.language.go!.pageableTypes;
  pagers.sort((a: PagerInfo, b: PagerInfo) => { return sortAscending(a.name, b.name) });
  for (const pager of values(pagers)) {
    const pagerType = internalPagerTypeName(pager);
    let pollerRespField = '';
    let respFieldCheck = '\tresp, err := p.pipeline.Do(req)';
    let requesterCondition = '';
    if (pager.hasLRO) {
      pollerRespField = `
      // previous response from the endpoint (LRO case)
      resp *azcore.Response`;
      respFieldCheck =
        `resp := p.resp
	if resp == nil {
		resp, err = p.pipeline.Do(req)
	} else {
		p.resp = nil
  }`;
      requesterCondition = ' if p.resp == nil';
    }
    const requesterType = ensureNameCase(`${pager.respType}CreateRequest`, true);
    const errorerType = ensureNameCase(`${pager.respType}HandleError`, true);
    const responderType = ensureNameCase(`${pager.respType}HandleResponse`, true);
    const advanceType = ensureNameCase(`${pager.respType}AdvancePage`, true);
    text += `// ${pager.name} provides iteration over ${pager.respType} pages.
type ${pager.name} interface {
	azcore.Pager

	// PageResponse returns the current ${pager.respEnv}.
	PageResponse() ${pager.respEnv}
}

type ${requesterType} func(context.Context) (*azcore.Request, error)

type ${errorerType} func(*azcore.Response) error

type ${responderType} func(*azcore.Response) (${pager.respEnv}, error)

type ${advanceType} func(context.Context, ${pager.respEnv}) (*azcore.Request, error)

type ${pagerType} struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester ${requesterType}
	// callback for handling response errors
	errorer ${errorerType}
	// callback for handling the HTTP response
	responder ${responderType}
	// callback for advancing to the next page
	advancer ${advanceType}
	// contains the current response
	current ${pager.respEnv}
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error${pollerRespField}
}

func (p *${pagerType}) Err() error {
	return p.err
}

func (p *${pagerType}) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.${pager.respField}.${pager.nextLink} == nil || len(*p.current.${pager.respField}.${pager.nextLink}) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
  } else${requesterCondition} {
		req, err = p.requester(ctx)
  }
	if err != nil {
		p.err = err
		return false
	}
  ${respFieldCheck}
	if err != nil {
		p.err = err
		return false
	}
`;
    text += `\tif !resp.HasStatusCode(p.statusCodes...) {\n`;
    text += `\tp.err = p.errorer(resp)\n`
    text += `\t\treturn false\n`;
    text += '\t}\n';
    text += `	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *${pagerType}) PageResponse() ${pager.respEnv} {
	return p.current
}

`;
  }
  return text;
}
