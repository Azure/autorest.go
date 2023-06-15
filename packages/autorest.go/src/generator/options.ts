/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel, GroupProperty } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { contentPreamble, sortAscending } from './helpers';
import { ImportManager } from './imports';
import { StructDef } from './structs';

// Creates the content in options.go
export async function generateOptions(session: Session<CodeModel>): Promise<string> {
  const paramGroups = <Array<GroupProperty>>session.model.language.go!.parameterGroups;
  if (!paramGroups) {
    return '';
  }

  let optionsText = await contentPreamble(session);
  const imports = new ImportManager();
  const structs = new Array<StructDef>();

  for (const paramGroup of values(paramGroups)) {
    const sd = new StructDef(paramGroup.schema.language.go!, undefined, paramGroup.originalParameter);
    for (const param of values(paramGroup.originalParameter)) {
      imports.addImportForSchemaType(param.schema);
    }
    structs.push(sd);
  }

  optionsText += imports.text();

  structs.sort((a: StructDef, b: StructDef) => { return sortAscending(a.Language.name, b.Language.name) });
  for (const struct of values(structs)) {
    optionsText += struct.text();
  }

  return optionsText;
}
