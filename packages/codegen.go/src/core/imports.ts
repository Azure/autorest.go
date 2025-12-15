/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../../codemodel.go/src/index.js';
import { values } from '@azure-tools/linq';
import * as helpers from './helpers.js';

type importEntry = { imp: string, alias?: string };

// tracks packages that need to be imported
export class ImportManager {
  private readonly imports: Array<importEntry>;
  private readonly pkg: go.FakePackage | go.PackageContent | go.TestPackage;

  /**
   * creates a new instance of ImportManager for the specified package
   * 
   * @param pkg the package that contains the import statements to emit
   */
  constructor(pkg: go.FakePackage | go.PackageContent | go.TestPackage) {
    this.imports = new Array<importEntry>();
    this.pkg = pkg;
  }

  /**
   * adds a package for importing if not already in the list
   * accepts an optional package alias.
   * 
   * @param imp the package name to import
   * @param alias optional alias for the import
   */
  add(imp: string, alias?: string): void {
    for (const existing of values(this.imports)) {
      if (existing.imp === imp) {
        return;
      }
    }
    this.imports.push({ imp: imp, alias: alias });
  }

  /**
   * adds the specified package for importing if not already in the list.
   * 
   * @param pkg the package to import
   * @param alias optional package alias
   */
  addForPkg(pkg: go.PackageContent, alias?: string): void {
    let pkgPath: string;
    switch (pkg.kind) {
      case 'module':
        pkgPath = pkg.identity;
        break;
      case 'package': {
        pkgPath = buildImportPath(pkg);
        break;
      }
    }

    this.add(pkgPath, alias);
  }

  /**
   * returns the number of packages in the list
   * 
   * @returns the import count
   */
  length(): number {
    return this.imports.length;
  }

  /**
   * returns the import list as Go source code
   * 
   * @returns the text for the import statement
   */
  text(): string {
    if (this.imports.length === 0) {
      return '';
    } else if (this.imports.length === 1) {
      const first = this.imports[0];
      return `import ${this.alias(first)}"${first.imp}"\n\n`;
    }
    this.imports.sort((a: importEntry, b: importEntry) => { return helpers.sortAscending(a.imp, b.imp); });
    let text = 'import (\n';
    for (const imp of values(this.imports)) {
      text += `\t${this.alias(imp)}"${imp.imp}"\n`;
    }
    text += ')\n\n';
    return text;
  }

  /**
   * adds an import statement for the specified type
   * as required if not already in the list.
   * 
   * @param type the type for which to add the import
   */
  addForType(type: go.Type): void {
    switch (type.kind) {
      case 'map':
        this.addForType(type.valueType);
        break;
      case 'slice':
        this.addForType(type.elementType);
        break;
      case 'constant':
      case 'interface':
      case 'model':
      case 'polymorphicModel':
        if (go.getPackageName(type.pkg) !== go.getPackageName(this.pkg)) {
          this.add(buildImportPath(type.pkg));
        }
    }

    // generic fallback for qualified types
    if ((<go.QualifiedType>type).name !== undefined && (<go.QualifiedType>type).module !== undefined) {
      this.add((<go.QualifiedType>type).module);
    }
  }

  /**
   * returns the import alias or the empty string
   * 
   * @param entry the entry to check for an alias
   * @returns the import alias or the empty string
   */
  private alias(entry: importEntry): string {
    if (entry.alias) {
      return `${entry.alias} `;
    }
    return '';
  }
}

/**
 * builds the complete package import path for the provided package
 * 
 * @param pkg the package for which to build the import path
 * @returns the fully qualified package path
 */
function buildImportPath(pkg: go.ContainingModule | go.Module | go.Package): string {
  const pkgs = new Array<string>();
  let cur: go.ContainingModule | go.Module | go.Package = pkg;
  while (cur.kind === 'package') {
    pkgs.unshift(cur.name);
    cur = cur.parent;
  }
  pkgs.unshift(cur.identity);
  return pkgs.join('/');
}
