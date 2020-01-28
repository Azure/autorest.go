/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment } from '@azure-tools/codegen'
import { CodeModel } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';

// tracks packages that need to be imported
export class ImportManager {
  private imports: string[];

  constructor() {
    this.imports = new Array<string>();
  }

  // adds a package for importing if not already in the list
  add(imp: string) {
    for (const existing of values(this.imports)) {
      if (existing === imp) {
        return;
      }
    }
    this.imports.push(imp);
  }

  // returns the number of packages in the list
  length(): number {
    return this.imports.length
  }

  text(): string {
    this.imports.sort(SortAscending);
    let text = 'import (\n';
    for (const imp of values(this.imports)) {
      text += `\t"${imp}"\n`;
    }
    text += ')\n\n';
    return text;
  }
}

// returns the common source-file preamble (license comment, package name etc)
export async function ContentPreamble(session: Session<CodeModel>): Promise<string> {
  const headerText = comment(await session.getValue("header-text", "MISSING LICENSE HEADER"), "// ");
  const namespace = await session.getValue('namespace');
  let text = `${headerText}\n\n`;
  text += `package ${namespace}\n\n`;
  return text;
}

// used to sort strings in ascending order
export function SortAscending(a: string, b: string): number {
  return a < b ? -1 : a > b ? 1 : 0;
}
