/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment } from '@azure-tools/codegen';
import { ChoiceValue, CodeModel } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { ContentPreamble, HasDescription, SortAscending } from '../common/helpers';

const internalPackage = 'azinternal';

export async function generateModels(session: Session<CodeModel>): Promise<string> {
  let text = await ContentPreamble(session);

  const module = await session.getValue('module-path');
  const namespace = await session.getValue('namespace');

  text += `import ${internalPackage} "${module}/internal/${namespace}"\n\n`;

  // create type aliases for enums

  for (const enm of values(session.model.schemas.choices)) {
    if (HasDescription(enm.language.go!)) {
      text += `${comment(enm.language.go!.name, '// ')} - ${enm.language.go!.description}\n`;
    }
    text += `type ${enm.language.go!.name} = ${internalPackage}.${enm.language.go!.name}\n\n`;
    enm.choices.sort((a: ChoiceValue, b: ChoiceValue) => { return SortAscending(a.language.go!.name, b.language.go!.name); });
    text += 'const (\n'
    for (const val of values(enm.choices)) {
      if (HasDescription(val.language.go!)) {
        text += `\t${comment(val.language.go!.name, '// ')} - ${val.language.go!.description}\n`;
      }
      text += `\t${val.language.go!.name} = ${internalPackage}.${val.language.go!.name}\n`;
    }
    text += ")\n\n"
    text += `func ${enm.language.go!.possibleValuesFunc}() []${enm.language.go!.name} {\n`;
    text += `\treturn ${internalPackage}.${enm.language.go!.possibleValuesFunc}()\n`;
    text += '}\n\n';
  }

  // create type aliases for structs

  // create a sorted list of struct type names/descriptions
  type EntryType = { name: string, desc?: string };
  const structs = new Array<EntryType>();
  for (const obj of values(session.model.schemas.objects)) {
    const entry: EntryType = { name: obj.language.go!.name };
    if (HasDescription(obj.language.go!)) {
      entry.desc = obj.language.go!.description;
    }
    structs.push(entry);
  }
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      if (op.responses) {
        const entry: EntryType = {
          name: op.responses[0].language.go!.name,
          desc: op.responses[0].language.go!.description,
        };
        structs.push(entry);
      }
    }
  }
  structs.sort((a: EntryType, b: EntryType) => { return SortAscending(a.name, b.name) });
  for (const struct of values(structs)) {
    if (struct.desc) {
      text += `${comment(struct.desc)}\n`;
    }
    text += `type ${struct.name} = azinternal.${struct.name}\n\n`;
  }
  return text;
}
