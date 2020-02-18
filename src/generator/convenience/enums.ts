/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment } from '@azure-tools/codegen';
import { CodeModel } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { InternalPackage, InternalPackagePath } from './helpers';
import { ContentPreamble, getEnums, HasDescription } from '../common/helpers';

// generates content for enums.go
export async function generateEnums(session: Session<CodeModel>): Promise<string> {
  const enums = getEnums(session.model.schemas);
  if (enums.length === 0) {
    // no enums to generate
    return '';
  }
  let text = await ContentPreamble(session);
  text += `import ${InternalPackage} "${await InternalPackagePath(session)}"\n\n`;

  // create type aliases for enums
  for (const enm of values(enums)) {
    if (enm.desc) {
      text += `${comment(enm.name, '// ')} - ${enm.desc}\n`;
    }
    text += `type ${enm.name} = ${InternalPackage}.${enm.name}\n\n`;
    text += 'const (\n'
    for (const val of values(enm.choices)) {
      if (HasDescription(val.language.go!)) {
        text += `\t${comment(val.language.go!.name, '// ')} - ${val.language.go!.description}\n`;
      }
      text += `\t${val.language.go!.name} = ${InternalPackage}.${val.language.go!.name}\n`;
    }
    text += ")\n\n"
    text += `func ${enm.funcName}() []${enm.name} {\n`;
    text += `\treturn ${InternalPackage}.${enm.funcName}()\n`;
    text += '}\n\n';
  }
  return text;
}
