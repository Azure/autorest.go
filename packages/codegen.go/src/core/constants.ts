/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as helpers from './helpers.js';
import * as go from '../../../codemodel.go/src/index.js';

/**
 * Creates the content for the constants.go file.
 *
 * @param pkg contains the package content
 * @returns the text for the file or the empty string
 */
export function generateConstants(pkg: go.PackageContent): string {
  if (pkg.constants.length === 0) {
    return '';
  }

  const indent = new helpers.Indentation();

  let text = helpers.contentPreamble(pkg);

  for (const enm of pkg.constants) {
    text += helpers.formatDocCommentWithPrefix(enm.name, enm.docs);
    text += `type ${enm.name} ${enm.type}\n\n`;
    const vals = new Array<string>();
    text += 'const (\n';
    for (const val of enm.values) {
      text += helpers.formatDocCommentWithPrefix(val.name, val.docs);
      let formatValue = `"${val.value}"`;
      if (enm.type !== 'string') {
        formatValue = `${val.value}`;
      }
      text += `${indent.get()}${val.name} ${enm.name} = ${formatValue}\n`;
      vals.push(val.name);
    }
    text += ')\n\n';
    text += `// ${enm.valuesFuncName} returns the possible values for the ${enm.name} const type.\n`;
    text += `func ${enm.valuesFuncName}() []${enm.name} {\n`;
    text += `${indent.get()}return []${enm.name}{\n`;
    indent.push();
    for (const val of vals) {
      text += `${indent.get()}${val},\n`;
    }
    text += `${indent.pop().get()}}\n`;
    text += '}\n\n';
  }
  return text;
}
