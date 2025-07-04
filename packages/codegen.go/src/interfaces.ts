/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../codemodel.go/src/index.js';
import { comment } from '@azure-tools/codegen';
import { contentPreamble, sortAscending } from './helpers.js';

// Creates the content in interfaces.go
export async function generateInterfaces(codeModel: go.CodeModel): Promise<string> {
  if (codeModel.interfaces.length === 0) {
    // no polymorphic types
    return '';
  }

  let text = contentPreamble(codeModel);

  for (const iface of codeModel.interfaces) {
    const methodName = `Get${iface.rootType.name}`;
    text += `// ${iface.name} provides polymorphic access to related types.\n`;
    text += `// Call the interface's ${methodName}() method to access the common type.\n`;
    text += '// Use a type switch to determine the concrete type.  The possible types are:\n';
    const possibleTypeNames = new Array<string>();
    possibleTypeNames.push(`*${iface.rootType.name}`);
    for (const possibleType of iface.possibleTypes) {
      possibleTypeNames.push(`*${possibleType.name}`);
    }
    possibleTypeNames.sort(sortAscending);
    text += comment(possibleTypeNames.join(', '), '// - ');
    text += `\ntype ${iface.name} interface {\n`;
    if (iface.parent) {
      text += `\t${iface.parent.name}\n`;
    }
    text += `\t// ${methodName} returns the ${iface.rootType.name} content of the underlying type.\n`;
    text += `\t${methodName}() *${iface.rootType.name}\n`;
    text += '}\n\n';
  }
  return text;
}
