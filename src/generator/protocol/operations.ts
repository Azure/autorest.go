/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment, pascalCase } from '@azure-tools/codegen'
import { CodeModel, Language, Operation, Parameter, Protocols } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { ContentPreamble, generateParamsSig, generateParameterInfo, genereateReturnsInfo, ImportManager, MethodSig, ParamInfo, SortAscending } from '../common/helpers';
import { OperationNaming } from '../../namer/namer';

// represents the generated content for an operation group
export class OperationGroupContent {
  readonly name: string;
  readonly content: string;

  constructor(name: string, content: string) {
    this.name = name;
    this.content = content;
  }
}

// Creates the content for all <operation>.go files
export async function generateOperations(session: Session<CodeModel>): Promise<OperationGroupContent[]> {
  // add standard imorts
  imports.add('net/http');
  imports.add('net/url');
  imports.add('path');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');

  // generate protocol operations
  const operations = new Array<OperationGroupContent>();
  for (const group of values(session.model.operationGroups)) {
    let text = await ContentPreamble(session);
    text += imports.text();

    const clientName = group.language.go!.clientName;
    text += `type ${clientName} struct{}\n\n`;

    group.operations.sort((a: Operation, b: Operation) => { return SortAscending(a.language.go!.name, b.language.go!.name) });
    for (const op of values(group.operations)) {
      op.language.go!.protocolSigs = new protocolSigs();
      text += createProtocolRequest(clientName, op);
      text += createProtocolResponse(clientName, op);
    }
    operations.push(new OperationGroupContent(group.language.go!.name, text));
  }
  return operations;
}

// contains method signature information for request and response methods
export interface ProtocolSig extends Language {
  protocolSigs: ProtocolSigs;
}

interface ProtocolSigs {
  requestMethod: MethodSig;
  responseMethod: MethodSig;
}

class protocolSigs implements ProtocolSigs {
  requestMethod: MethodSig;
  responseMethod: MethodSig;
  constructor() {
    this.requestMethod = new methodSig();
    this.responseMethod = new methodSig();
  }
}

class methodSig implements MethodSig {
  params: ParamInfo[];
  returns: string[];
  constructor() {
    this.params = new Array<ParamInfo>();
    this.returns = new Array<string>();
  }
}

// this list of packages to import
const imports = new ImportManager();

function createProtocolRequest(client: string, op: Operation): string {
  const info = <OperationNaming>op.language.go!;
  const name = info.protocolNaming.requestMethod;
  // stick the method signature info into the code model so other generators can access it later
  const sig = <ProtocolSig>op.language.go!;
  sig.protocolSigs.requestMethod.params = [{ name: 'u', type: 'url.URL' }].concat(generateParameterInfo(op));
  sig.protocolSigs.requestMethod.returns = ['*azcore.Request', 'error'];
  let text = `${comment(name, '// ')} creates the ${info.name} request.\n`;
  text += `func (${client}) ${name}(${generateParamsSig(sig.protocolSigs.requestMethod.params)}) (${sig.protocolSigs.requestMethod.returns.join(', ')}) {\n`;
  text += `\tu.Path = path.Join(u.Path, "${op.request.protocol.http!.path}")\n`;
  const reqObj = `azcore.NewRequest(http.Method${pascalCase(op.request.protocol.http!.method)}, u)`;
  if (getMediaType(op.request.protocol) === 'none') {
    // no request body so nothing to marshal
    text += `\treturn ${reqObj}, nil\n`;
  } else {
    const bodyParam = values(op.request.parameters).where((each: Parameter) => { return each.protocol.http!.in === 'body'; }).first();
    text += `\treq := ${reqObj}\n`;
    text += `\terr := req.MarshalAs${getMediaType(op.request.protocol)}(${bodyParam?.language.go!.name})\n`;
    text += `\tif err != nil {\n`;
    text += `\t\treturn nil, err\n`;
    text += `\t}\n`;
    text += `\treturn req, nil\n`;
  }
  text += '}\n\n';
  return text;
}

function createProtocolResponse(client: string, op: Operation): string {
  const info = <OperationNaming>op.language.go!;
  const name = info.protocolNaming.responseMethod;
  // stick the method signature info into the code model so other generators can access it later
  const sig = <ProtocolSig>op.language.go!;
  sig.protocolSigs.responseMethod.params = [{ name: 'resp', type: '*azcore.Response' }];
  sig.protocolSigs.responseMethod.returns = genereateReturnsInfo(op);

  let text = `${comment(name, '// ')} handles the ${info.name} response.\n`;
  text += `func (${client}) ${name}(${generateParamsSig(sig.protocolSigs.responseMethod.params)}) (${sig.protocolSigs.responseMethod.returns.join(', ')}) {\n`;
  text += `\tif !resp.HasStatusCode(http.StatusOK) {\n`;
  text += `\t\treturn nil, newError(resp)\n`;
  text += '\t}\n';

  const resp = op.responses![0];
  const respObj = `${resp.language.go!.name}{StatusCode: resp.StatusCode}`;
  if (getMediaType(resp.protocol) === 'none') {
    // no response body so nothing to unmarshal
    text += `\treturn &${respObj}, nil\n`;
  } else {
    text += `\tresult := ${respObj}\n`;
    text += `\treturn &result, resp.UnmarshalAs${getMediaType(resp.protocol)}(&result.Value)\n`;
  }
  text += '}\n\n';
  return text;
}

function getMediaType(protocol: Protocols): 'JSON' | 'none' {
  if (protocol.http!.knownMediaType === undefined) {
    return 'none';
  }
  return 'JSON';
}
