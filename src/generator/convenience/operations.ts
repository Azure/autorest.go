/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { camelCase } from '@azure-tools/codegen';
import { CodeModel, ImplementationLocation, Operation } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { InternalPackage, InternalPackagePath } from './helpers';
import { ContentPreamble, generateParamsSig, generateParameterInfo, genereateReturnsInfo, HasDescription, ImportManager, ParamInfo } from '../common/helpers';
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
  // generate protocol operations
  const operations = new Array<OperationInfo>();
  for (const group of values(session.model.operationGroups)) {
    // this list of packages to import
    const imports = new ImportManager();
    // add standard imorts
    imports.add('context');
    imports.add(await InternalPackagePath(session), InternalPackage);

    let interfaceText = '';
    // interface definition
    // this can add imports to the list so it must
    // be done before the imports are written out
    interfaceText += `// ${group.language.go!.clientName} contains the methods for the ${group.language.go!.name} group.\n`;
    interfaceText += `type ${group.language.go!.clientName} interface {\n`;
    for (const op of values(group.operations)) {
      for (const param of values(op.request.parameters)) {
        if (param.implementation !== ImplementationLocation.Method) {
          continue;
        }
        imports.addImportForSchemaType(param.schema);
      }
      if (HasDescription(op.language.go!)) {
        interfaceText += `\t// ${op.language.go!.name} - ${op.language.go!.description} \n`;
      }
      const params = [{ name: 'ctx', type: 'context.Context', global: false }].concat(generateParameterInfo(op));
      const returns = genereateReturnsInfo(op);
      interfaceText += `\t${op.language.go!.name}(${generateParamsSig(params, false)}) (${returns.join(', ')})\n`;
    }
    interfaceText += '}\n\n';

    let text = await ContentPreamble(session);
    text += imports.text();
    text += interfaceText;

    // internal client type
    const clientName = camelCase(group.language.go!.clientName);
    text += `type ${clientName} struct {\n`;
    text += '\t*Client\n';
    text += `\t${InternalPackage}.${group.language.go!.clientName}\n`;
    if (group.language.go!.globals) {
      const globals = <Array<ParamInfo>>group.language.go!.globals;
      globals.forEach((value: ParamInfo, index: Number, obj: ParamInfo[]) => {
        text += `\t${value.name} ${value.type}\n`;
      })
    }
    text += '}\n\n';

    for (const op of values(group.operations)) {
      text += generateOperation(clientName, op);
    }

    text += `var _ ${group.language.go!.clientName} = (*${clientName})(nil)\n`;
    operations.push(new OperationInfo(group.language.go!.name, text));
  }
  return operations;
}

function generateOperation(clientName: string, op: Operation): string {
  const info = <OperationNaming>op.language.go!;
  const params = [{ name: 'ctx', type: 'context.Context', global: false }].concat(generateParameterInfo(op));
  const returns = genereateReturnsInfo(op);
  const protocol = <ProtocolSig>op.language.go!;
  let text = '';
  if (HasDescription(op.language.go!)) {
    text += `// ${op.language.go!.name} - ${op.language.go!.description} \n`;
  }
  text += `func (client *${clientName}) ${op.language.go!.name}(${generateParamsSig(params, false)}) (${returns.join(', ')}) {\n`;
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

// returns an array of just the parameter names
// e.g. [ 'i', 's', 'b' ]
function extractParamNames(paramInfo: ParamInfo[]): string[] {
  let paramNames = new Array<string>();
  for (const param of values(paramInfo)) {
    let name = param.name;
    if (param.global) {
      name = `client.${name}`;
    }
    paramNames.push(name);
  }
  return paramNames;
}
