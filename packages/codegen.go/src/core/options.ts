/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { capitalize } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import * as go from '../../../codemodel.go/src/index.js';
import * as helpers from './helpers.js';
import { ImportManager } from './imports.js';

/**
 * Creates the content for the options.go file.
 * 
 * @param pkg contains the package content
 * @returns the text for the file or the empty string
 */
export function generateOptions(pkg: go.PackageContent): string {
  if (pkg.paramGroups.length === 0) {
    return '';
  }

  const imports = new ImportManager();
  let optionsText = helpers.contentPreamble(helpers.getPackageName(pkg));
  let content = '';

  for (const paramGroup of pkg.paramGroups) {
    content += emit(paramGroup, imports);
  }

  optionsText += imports.text();
  optionsText += content;
  return optionsText;
}

function emit(struct: go.Struct, imports: ImportManager): string {
  let text = helpers.formatDocComment(struct.docs);
  text += `type ${struct.name} struct {\n`;

  if (struct.fields.length === 0) {
    // this is an optional params placeholder struct
    text += '\t// placeholder for future optional parameters\n';
  } else {
    // used to track when to add an extra \n between fields that have comments
    let first = true;

    for (const field of values(struct.fields)) {
      imports.addImportForType(field.type);
      if (field.docs.summary || field.docs.description) {
        if (!first) {
          // add an extra new-line between fields IFF the field
          // has a comment and it's not the very first one.
          text += '\n';
        }
        text += helpers.formatDocComment(field.docs);
      }

      let typeName = go.getTypeDeclaration(field.type);
      if (field.type.kind === 'literal') {
        // for constants we use the underlying type name
        typeName = go.getLiteralTypeDeclaration(field.type.type);
      }

      let pointer = '*';
      if (field.byValue) {
        pointer = '';
      }
      text += `\t${capitalize(field.name)} ${pointer}${typeName}\n`;
      first = false;
    }
  }

  text += '}\n\n';
  return text;
}
