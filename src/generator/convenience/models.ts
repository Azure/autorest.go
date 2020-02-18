/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment } from '@azure-tools/codegen';
import { CodeModel } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { InternalPackage, InternalPackagePath } from './helpers';
import { ContentPreamble, getEnums, HasDescription, SortAscending } from '../common/helpers';

// generates content for models.go
export async function generateModels(session: Session<CodeModel>): Promise<string> {
  let text = await ContentPreamble(session);
  text += `import ${InternalPackage} "${await InternalPackagePath(session)}"\n\n`;

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
