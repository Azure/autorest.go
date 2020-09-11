/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { camelCase } from '@azure-tools/codegen';
import { CodeModel, SchemaResponse } from '@azure-tools/codemodel';
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
  text += imports.text();

  const pagers = <Array<PagerInfo>>session.model.language.go!.pageableTypes;
  pagers.sort((a: PagerInfo, b: PagerInfo) => { return sortAscending(a.name, b.name) });
  for (const pager of values(pagers)) {
    const pagerType = camelCase(pager.name);
    const schemaResponse = <SchemaResponse>pager.op.responses![0];
    const responseType = schemaResponse.schema.language.go!.responseType.name;
    const resultType = schemaResponse.schema.language.go!.name;
    let resultTypeName = resultType;
    let pollerRespField = '';
    let respFieldCheck = '\tresp, err := p.pipeline.Do(req)';
    let requesterCondition = '';
    if (pager.respField) {
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
    if (schemaResponse.schema.serialization?.xml?.name) {
      // xml can specifiy its own name, prefer that if available
      resultTypeName = schemaResponse.schema.serialization.xml.name;
    }
    const requesterType = `${camelCase(resultType)}CreateRequest`;
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

type ${requesterType} func(context.Context) (*azcore.Request, error)

type ${responderType} func(*azcore.Response) (*${responseType}, error)

type ${advanceType} func(context.Context, *${responseType}) (*azcore.Request, error)

type ${pagerType} struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester ${requesterType}
	// callback for handling the HTTP response
	responder ${responderType}
	// callback for advancing to the next page
	advancer ${advanceType}
	// contains the current response
	current *${responseType}
	// any error encountered
	err error${pollerRespField}
}

func (p *${pagerType}) Err() error {
	return p.err
}

func (p *${pagerType}) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if p.current != nil {
		if p.current.${resultTypeName}.${pager.op.language.go!.paging.nextLinkName} == nil || len(*p.current.${resultTypeName}.${pager.op.language.go!.paging.nextLinkName}) == 0 {
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
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *${pagerType}) PageResponse() *${responseType} {
	return p.current
}

`;
  }
  return text;
}
