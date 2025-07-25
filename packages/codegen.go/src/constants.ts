/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { values } from '@azure-tools/linq';
import * as helpers from './helpers.js';
import * as go from '../../codemodel.go/src/index.js';

// Creates the content in constants.go
export async function generateConstants(codeModel: go.CodeModel): Promise<string> {
  if (codeModel.constants.length === 0) {
    return '';
  }

  let text = helpers.contentPreamble(codeModel);

  for (const enm of values(codeModel.constants)) {
    text += helpers.formatDocCommentWithPrefix(enm.name, enm.docs);
    text += `type ${enm.name} ${enm.type}\n\n`;
    const vals = new Array<string>();
    text += 'const (\n';
    for (const val of values(enm.values)) {
      text += helpers.formatDocCommentWithPrefix(val.name, val.docs);
      let formatValue = `"${val.value}"`;
      if (enm.type !== 'string') {
        formatValue = `${val.value}`;
      }
      text += `\t${val.name} ${enm.name} = ${formatValue}\n`;
      vals.push(val.name);
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
