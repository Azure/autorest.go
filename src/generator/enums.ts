/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment, joinComma } from '@azure-tools/codegen';
import { CodeModel, ChoiceValue, Schemas } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { contentPreamble, hasDescription, sortAscending } from './helpers';

// Creates the content in enums.go
export async function generateEnums(session: Session<CodeModel>): Promise<string> {
  const enums = getEnums(session.model.schemas);
  if (enums.length === 0) {
    // no enums to generate
    return '';
  }
  let text = await contentPreamble(session);
  for (const enm of values(enums)) {
    if (enm.desc) {
      text += `${comment(enm.name, '// ')} - ${enm.desc}\n`;
    }
    text += `type ${enm.name} ${enm.type}\n\n`;
    const vals = new Array<string>();
    text += 'const (\n';
    for (const val of values(enm.choices)) {
      if (hasDescription(val.language.go!)) {
        text += `\t${comment(val.language.go!.name, '// ')} - ${val.language.go!.description}\n`;
      }
      let formatValue = `"${val.value}"`;
      if (enm.type !== 'string') {
        formatValue = `${val.value}`;
      }
      text += `\t${val.language.go!.name} ${enm.name} = ${formatValue}\n`;
      vals.push(val.language.go!.name);
    }
    text += ')\n\n';
    text += `func ${enm.funcName}() []${enm.name} {\n`;
    text += `\treturn []${enm.name}{\t\n`;
    for (const val of values(vals)) {
      text += `\t\t${val},\n`;
    }
    text += '\t}\n';
    text += '}\n\n';
    text += `func (c ${enm.name}) ToPtr() *${enm.name} {\n`;
    text += '\treturn &c\n';
    text += `}\n\n`;
  }
  return text;
}

// returns a collection containing all enum entries and their values
function getEnums(schemas: Schemas): EnumEntry[] {
  // group all enum categories into a single array so they can be sorted
  const enums = new Array<EnumEntry>();
  for (const choice of values(schemas.choices)) {
    choice.choices.sort((a: ChoiceValue, b: ChoiceValue) => { return sortAscending(a.language.go!.name, b.language.go!.name); });
    const entry = new EnumEntry(choice.language.go!.name, choice.choiceType.language.go!.name, choice.language.go!.possibleValuesFunc, choice.choices);
    if (hasDescription(choice.language.go!)) {
      entry.desc = choice.language.go!.description;
    }
    enums.push(entry);
  }
  for (const choice of values(schemas.sealedChoices)) {
    if (choice.choices.length === 1) {
      continue;
    }
    const entry = new EnumEntry(choice.language.go!.name, choice.choiceType.language.go!.name, choice.language.go!.possibleValuesFunc, choice.choices);
    if (hasDescription(choice.language.go!)) {
      entry.desc = choice.language.go!.description;
    }
    enums.push(entry);
  }
  enums.sort((a: EnumEntry, b: EnumEntry) => { return sortAscending(a.name, b.name) });
  return enums;
}

// represents an enum type and its values
class EnumEntry {
  name: string;
  type: string;
  funcName: string;
  desc?: string;
  choices: ChoiceValue[];
  constructor(name: string, type: string, funcName: string, choices: ChoiceValue[]) {
    this.name = name;
    this.type = type;
    this.funcName = funcName;
    this.choices = choices;
  }
}
