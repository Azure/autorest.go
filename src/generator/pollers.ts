/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel, SchemaType, Schema, ArraySchema } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { internalPagerTypeName, internalPollerTypeName, PollerInfo } from '../common/helpers';
import { contentPreamble, sortAscending } from './helpers';
import { ImportManager } from './imports';
import { ensureNameCase } from '../transform/namer';

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
      bodyText += '\tpipeline azcore.Pipeline\n';
      bodyText += '\tpt armcore.Poller\n';
    } else {
      bodyText += '\tpt *azcore.LROPoller\n';
    }
    if (poller.pager) {
      bodyText += `\terrHandler ${ensureNameCase(poller.pager.respType, true)}HandleError\n`;
      bodyText += `\trespHandler ${ensureNameCase(poller.pager.respType, true)}HandleResponse\n`;
      bodyText += '\tstatusCodes []int\n';
    }
    bodyText += '}\n\n';
    // internal poller methods
    bodyText += `func (p *${pollerName}) Done() bool {\n\treturn p.pt.Done()\n}\n\n`;
    let plParam = '';
    if (isARM) {
      plParam = ', p.pipeline';
    }
    bodyText += `func (p *${pollerName}) Poll(ctx context.Context) (*http.Response, error) {\n\treturn p.pt.Poll(ctx${plParam})\n}\n\n`;
    bodyText += pudFinalResp('FinalResponse', poller, imports, isARM);
    bodyText += `func (p *${pollerName}) ResumeToken() (string, error) {\n\treturn p.pt.ResumeToken()\n}\n\n`;
    bodyText += pudFinalResp('pollUntilDone', poller, imports, isARM);
    if (poller.pager) {
      bodyText += pagerHandleResponse(poller);
    }
  }
  text += imports.text();
  text += bodyText;
  return text;
}

function getResponseType(poller: PollerInfo): string {
  let respType = '*http.Response';
  // check for pager must come first
  if (poller.pager) {
    respType = poller.pager.name;
  } else if (poller.respType) {
    respType = poller.respEnv;
  }
  return respType;
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
function pudFinalResp(op: 'pollUntilDone' | 'FinalResponse', poller: PollerInfo, imports: ImportManager, isARM: boolean): string {
  let durParam = '';
  let freqParam = '';
  let plParam = '';
  if (op === 'pollUntilDone') {
    imports.add('time');
    durParam = ', freq time.Duration';
    freqParam = ', freq';
  }
  if (isARM) {
    plParam = ', p.pipeline';
  }
  let text = `func (p *${internalPollerTypeName(poller)}) ${op}(ctx context.Context${durParam}) (${getResponseType(poller)}, error) {\n`;
  if (poller.pager) {
    // pager-pollers have a slightly different impl
    text += `\trespType := &${ensureNameCase(poller.pager.name, true)}{}\n`;
    text += `\tresp, err := p.pt.${op.capitalize()}(ctx${freqParam}${plParam}, respType)\n`;
    text += `\tif err != nil {\n\t\treturn nil, err\n\t}\n`;
    text += '\treturn p.handleResponse(&azcore.Response{Response: resp})\n';
  } else if (poller.respType) {
    let reference = '';
    let respByRef = '&';
    if (poller.respType.type === SchemaType.Array || poller.respType.type === SchemaType.Dictionary) {
      // arrays and maps are returned by value
      respByRef = '';
      // but we need to pass them by reference to the unmarshaller
      reference = '&';
    }
    if (isScalarType(poller.respType)) {
      text += `\trespType := ${poller.respEnv}{}\n`;
      reference = '&';
    } else if (poller.respType.type === SchemaType.Any || poller.respType.type === SchemaType.AnyObject) {
      text += `\trespType := ${poller.respEnv}{}\n`;
    } else {
      text += `\trespType := ${poller.respEnv}{${poller.respField}: ${respByRef}${poller.respType.language.go!.name}{}}\n`;
    }
    text += `\tresp, err := p.pt.${op.capitalize()}(ctx${freqParam}${plParam}, ${reference}respType.${poller.respField})\n`;
    text += `\tif err != nil {\n\t\treturn ${poller.respEnv}{}, err\n\t}\n`;
    text += '\trespType.RawResponse = resp\n';
    text += '\treturn respType, nil\n';
  } else {
    // poller doesn't return a type
    text += `\treturn p.pt.${op.capitalize()}(ctx${freqParam}${plParam}, nil)\n`;
  }
  text += '}\n\n';
  return text;
}

function pagerHandleResponse(poller: PollerInfo): string {
  let text = `func(p * ${internalPollerTypeName(poller)}) handleResponse(resp * azcore.Response)(${getResponseType(poller)}, error) {\n`;
  text += `\treturn &${internalPagerTypeName(poller.pager!)}{\n`;
  text += `\t\tpipeline: p.pipeline,\n`;
  text += `\t\tresp: resp,\n`;
  text += '\t\terrorer: p.errHandler,\n';
  text += `\t\tresponder: p.respHandler,\n`;
  text += `\t\tadvancer: func(ctx context.Context, resp ${poller.pager!.respEnv}) (*azcore.Request, error) {\n`;
  text += `\t\t\treturn azcore.NewRequest(ctx, http.MethodGet, *resp.${poller.pager!.respField}.${poller.pager!.nextLink})\n`;
  text += '\t\t},\n';
  text += `\t\tstatusCodes: p.statusCodes,\n`;
  text += `\t}, nil`;
  text += '}\n\n';
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
