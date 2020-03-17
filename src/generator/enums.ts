/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment, joinComma } from '@azure-tools/codegen';
import { CodeModel } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { ContentPreamble, getEnums, HasDescription } from './helpers';

// Creates the content in enums.go
export async function generateEnums(session: Session<CodeModel>): Promise<string> {
  const enums = getEnums(session.model.schemas);
  if (enums.length === 0) {
    // no enums to generate
    return '';
  }
  let text = await ContentPreamble(session);
  for (const enm of values(enums)) {
    if (enm.desc) {
      text += `${comment(enm.name, '// ')} - ${enm.desc}\n`;
    }
    text += `type ${enm.name} ${enm.type}\n\n`;
    const vals = new Array<string>();
    text += 'const (\n'
    for (const val of values(enm.choices)) {
      if (HasDescription(val.language.go!)) {
        text += `\t${comment(val.language.go!.name, '// ')} - ${val.language.go!.description}\n`;
      }
      text += `\t${val.language.go!.name} ${enm.name} = "${val.value}"\n`;
      vals.push(val.language.go!.name);
    }
    text += ")\n\n"
    text += `func ${enm.funcName}() []${enm.name} {\n`;
    text += `\treturn []${enm.name}{${joinComma(vals, (item: string) => item)}}\n`;
    text += '}\n\n';
    text += `func (c ${enm.name}) ToPtr() *${enm.name} {\n`;
    text += '\treturn &c\n';
    text += `}\n\n`;
  }
  return text;
}
