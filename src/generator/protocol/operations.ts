/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment, pascalCase } from '@azure-tools/codegen'
import { CodeModel, Operation, Parameter, Protocols, ImplementationLocation } from '@azure-tools/codemodel';
import { length, values } from '@azure-tools/linq';
import { ContentPreamble, ImportManager, SortAscending } from './helpers';
import { OperationNaming } from '../../namer/namer';

// represents an operation group
export class OperationInfo {
  readonly name: string;
  readonly content: string;

  constructor(name: string, content: string) {
    this.name = name;
    this.content = content;
  }
}

// Creates the content for all <operation>.go files
export async function generateOperations(session: Session<CodeModel>): Promise<OperationInfo[]> {
  // add standard imorts
  imports.add('net/http');
  imports.add('net/url');
  imports.add('path');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');

  // generate protocol operations
  const operations = new Array<OperationInfo>();
  for (const group of values(session.model.operationGroups)) {
    let text = await ContentPreamble(session);
    text += imports.text();

    const clientName = group.language.go!.clientName;
    text += `type ${clientName} struct{}\n\n`;

    group.operations.sort((a: Operation, b: Operation) => { return SortAscending(a.language.go!.name, b.language.go!.name) });
    for (const op of values(group.operations)) {
      text += createProtocolRequest(clientName, op);
      text += createProtocolResponse(clientName, op);
    }
    operations.push(new OperationInfo(group.language.go!.name, text));
  }
  return operations;
}

// this list of packages to import
const imports = new ImportManager();

function createProtocolRequest(client: string, op: Operation): string {
  const info = <OperationNaming>op.language.go!;
  const name = info.protocolNaming.requestMethod;
  const params = ['u url.URL'];
  for (const param of values(op.request.parameters)) {
    if (param.implementation === ImplementationLocation.Method) {
      params.push(`${param.language.go!.name} ${param.schema.language.go!.name}`);
    }
  }
  const returns = ['*azcore.Request', 'error'];
  let text = `${comment(name, '// ')} creates the ${info.name} request.\n`;
  text += `func (${client}) ${name}(${params.join(', ')}) (${returns.join(', ')}) {\n`;
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
  const params = ['resp *azcore.Response'];
  if (length(op.responses) > 1) {
    throw console.error('multiple responses NYI');
  }
  const resp = op.responses![0];
  const returns = [`*${resp.language.go!.name}`, 'error'];

  let text = `${comment(name, '// ')} handles the ${info.name} response.\n`;
  text += `func (${client}) ${name}(${params}) (${returns.join(', ')}) {\n`;
  text += `\tif !resp.HasStatusCode(http.StatusOK) {\n`;
  text += `\t\treturn nil, newError(resp)\n`;
  text += '\t}\n';
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
