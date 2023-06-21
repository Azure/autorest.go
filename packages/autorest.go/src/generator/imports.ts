/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { ArraySchema, DictionarySchema, Schema, SchemaType } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { sortAscending } from './helpers';

type importEntry = { imp: string, alias?: string };

// tracks packages that need to be imported
export class ImportManager {
  private imports: Array<importEntry>;

  constructor() {
    this.imports = new Array<importEntry>();
  }

  // adds a package for importing if not already in the list
  // accepts an optional package alias.
  add(imp: string, alias?: string) {
    for (const existing of values(this.imports)) {
      if (existing.imp === imp) {
        return;
      }
    }
    this.imports.push({ imp: imp, alias: alias });
  }

  // returns the number of packages in the list
  length(): number {
    return this.imports.length;
  }

  // returns the import list as Go source code
  text(): string {
    if (this.imports.length === 0) {
      return '';
    } else if (this.imports.length === 1) {
      const first = this.imports[0];
      return `import ${this.alias(first)}"${first.imp}"\n\n`;
    }
    this.imports.sort((a: importEntry, b: importEntry) => { return sortAscending(a.imp, b.imp); });
    let text = 'import (\n';
    for (const imp of values(this.imports)) {
      text += `\t${this.alias(imp)}"${imp.imp}"\n`;
    }
    text += ')\n\n';
    return text;
  }

  addImportForSchemaType(schema: Schema) {
    switch (schema.type) {
      case SchemaType.Array:
        this.addImportForSchemaType((<ArraySchema>schema).elementType);
        break;
      case SchemaType.Binary:
        this.add('io');
        break;
      case SchemaType.Dictionary:
        this.addImportForSchemaType((<DictionarySchema>schema).elementType);
        break;
      case SchemaType.Date:
      case SchemaType.DateTime:
      case SchemaType.UnixTime:
        this.add('time');
        break;
    }
  }

  private alias(entry: importEntry): string {
    if (entry.alias) {
      return `${entry.alias} `;
    }
    return '';
  }
}
