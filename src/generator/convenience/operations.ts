/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { camelCase } from '@azure-tools/codegen';
import { CodeModel, Operation } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { InternalPackage, InternalPackagePath } from './helpers';
import { ContentPreamble, extractParamNames, generateParamsSig, generateParameterInfo, genereateReturnsInfo, HasDescription, ImportManager } from '../common/helpers';
import { OperationNaming } from '../../namer/namer';
import { ProtocolSig } from '../protocol/operations';

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
  imports.add('context');
  imports.add(await InternalPackagePath(session), InternalPackage);

  // generate protocol operations
  const operations = new Array<OperationInfo>();
  for (const group of values(session.model.operationGroups)) {
    let text = await ContentPreamble(session);
    text += imports.text();

    // interface definition
    text += `// ${group.language.go!.clientName} contains the methods for the ${group.language.go!.name} group.\n`;
    text += `type ${group.language.go!.clientName} interface {\n`;
    for (const op of values(group.operations)) {
      if (HasDescription(op.language.go!)) {
        text += `\t// ${op.language.go!.name} - ${op.language.go!.description} \n`;
      }
      const params = [{ name: 'ctx', type: 'context.Context' }].concat(generateParameterInfo(op));
      const returns = genereateReturnsInfo(op);
      text += `\t${op.language.go!.name}(${generateParamsSig(params)}) (${returns.join(', ')})\n`;
    }
    text += '}\n\n';

    // internal client type
    const clientName = camelCase(group.language.go!.clientName);
    text += `type ${clientName} struct {\n`;
    text += '\t*Client\n';
    text += `\t${InternalPackage}.${group.language.go!.clientName}\n`;
    text += '}\n\n';

    for (const op of values(group.operations)) {
      text += generateOperation(clientName, op);
    }

    text += `var _ ${group.language.go!.clientName} = (*${clientName})(nil)\n`;
    operations.push(new OperationInfo(group.language.go!.name, text));
  }
  return operations;
}

// this list of packages to import
const imports = new ImportManager();

function generateOperation(clientName: string, op: Operation): string {
  const info = <OperationNaming>op.language.go!;
  const params = [{ name: 'ctx', type: 'context.Context' }].concat(generateParameterInfo(op));
  const returns = genereateReturnsInfo(op);
  const protocol = <ProtocolSig>op.language.go!;
  let text = '';
  if (HasDescription(op.language.go!)) {
    text += `// ${op.language.go!.name} - ${op.language.go!.description} \n`;
  }
  text += `func (client *${clientName}) ${op.language.go!.name}(${generateParamsSig(params)}) (${returns.join(', ')}) {\n`;
  // slice off the first param returned from extractParamNames as we know it's the URL (cheating a bit...)
  const protocolReqParams = ['*client.u'].concat(extractParamNames(protocol.protocolSigs.requestMethod.params).slice(1));
  text += `\treq, err := client.${info.protocolNaming.requestMethod}(${protocolReqParams.join(', ')})\n`;
  text += `\tif err != nil {\n`;
  text += `\t\treturn nil, err\n`;
  text += `\t}\n`;
  text += `\tresp, err := client.p.Do(ctx, req)\n`;
  text += `\tif err != nil {\n`;
  text += `\t\treturn nil, err\n`;
  text += `\t}\n`;
  // also cheating here as at present the only param to the responder is an azcore.Response
  text += `\tresult, err := client.${info.protocolNaming.responseMethod}(resp)\n`;
  text += `\tif err != nil {\n`;
  text += `\t\treturn nil, err\n`;
  text += `\t}\n`;
  text += `\treturn result, nil\n`;
  text += '}\n\n';
  return text;
}
