/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { comment } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import { commentLength, contentPreamble } from './helpers';
import { GoCodeModel } from '../gocodemodel/gocodemodel';

// Creates the content in constants.go
export async function generateConstants(codeModel: GoCodeModel): Promise<string> {
  // lack of operation groups indicates model-only mode.
  if (!codeModel.clients || (codeModel.constants.length === 0 && !codeModel.host && codeModel.type !== 'azure-arm')) {
    return '';
  }
  let text = contentPreamble(codeModel);
  if (codeModel.host) {
    text += `const host = "${codeModel.host}"\n\n`;
  }
  // data-plane clients must manage their own constants for these values
  if (codeModel.type === 'azure-arm') {
    if (!codeModel.options.moduleVersion) {
      throw new Error('--module-version is a required parameter when --azure-arm is set');
    }
    text += 'const (\n';
    text += `\tmoduleName = "${codeModel.options.module}"\n`;
    text += `\tmoduleVersion = "v${codeModel.options.moduleVersion}"\n`;
    text += ')\n\n';
  }
  for (const enm of values(codeModel.constants)) {
    if (enm.description) {
      text += `${comment(`${enm.name} - ${enm.description}`, '// ', undefined, commentLength)}\n`;
    }
    text += `type ${enm.name} ${enm.type}\n\n`;
    const vals = new Array<string>();
    text += 'const (\n';
    for (const val of values(enm.values)) {
      if (val.description) {
        text += `\t${comment(`${val.valueName} - ${val.description}`, '//', undefined, commentLength)}\n`;
      }
      let formatValue = `"${val.value}"`;
      if (enm.type !== 'string') {
        formatValue = `${val.value}`;
      }
      text += `\t${val.valueName} ${enm.name} = ${formatValue}\n`;
      vals.push(val.valueName);
    }
    text += ')\n\n';
    text += `// ${enm.valuesFuncName} returns the possible values for the ${enm.name} const type.\n`;
    text += `func ${enm.valuesFuncName}() []${enm.name} {\n`;
    text += `\treturn []${enm.name}{\t\n`;
    for (const val of values(vals)) {
      text += `\t\t${val},\n`;
    }
    text += '\t}\n';
    text += '}\n\n';
  }
  return text;
}
