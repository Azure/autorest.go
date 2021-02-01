/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { camelCase } from '@azure-tools/codegen';
import { CodeModel } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { PagerInfo } from '../common/helpers';
import { contentPreamble, sortAscending } from './helpers';
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
  imports.add('reflect');
  text += imports.text();

  const pagers = <Array<PagerInfo>>session.model.language.go!.pageableTypes;
  pagers.sort((a: PagerInfo, b: PagerInfo) => { return sortAscending(a.name, b.name) });
  for (const pager of values(pagers)) {
    const pagerType = camelCase(pager.name);
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
    const requesterType = `${camelCase(pager.respType)}CreateRequest`;
    const errorerType = `${camelCase(pager.respType)}HandleError`;
    const responderType = `${camelCase(pager.respType)}HandleResponse`;
    const advanceType = `${camelCase(pager.respType)}AdvancePage`;
    text += `// ${pager.name} provides iteration over ${pager.respType} pages.
type ${pager.name} interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current ${pager.respEnv}.
	PageResponse() ${pager.respEnv}

	// Err returns the last error encountered while paging.
	Err() error
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
