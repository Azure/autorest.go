/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment } from '@azure-tools/codegen';
import { CodeModel, ImplementationLocation, Language, Operation } from '@azure-tools/codemodel';
import { length, values } from '@azure-tools/linq';

type importEntry = { imp: string, alias?: string };

// tracks packages that need to be imported
export class ImportManager {
  private imports: importEntry[];

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
    return this.imports.length
  }

  // returns the import list as Go source code
  text(): string {
    if (this.imports.length === 1) {
      const first = this.imports[0];
      return `import ${this.alias(first)}"${first.imp}"\n\n`;
    }
    this.imports.sort((a: importEntry, b: importEntry) => { return SortAscending(a.imp, b.imp) });
    let text = 'import (\n';
    for (const imp of values(this.imports)) {
      text += `\t${this.alias(imp)}"${imp.imp}"\n`;
    }
    text += ')\n\n';
    return text;
  }

  private alias(entry: importEntry): string {
    if (entry.alias) {
      return `${entry.alias} `;
    }
    return '';
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

// returns true if the language contains a description
export function HasDescription(lang: Language): boolean {
  return (lang.description.length > 0 && !lang.description.startsWith('MISSING'));
}

// describes a method's signature, including parameters and return values
export interface MethodSig {
  params: ParamInfo[];
  returns: string[];
}

// describes a method paramater
export interface ParamInfo {
  name: string;
  type: string;
}

// creates ParamInfo for the specified operation.
// each entry is tuple of param name/param type
export function generateParameterInfo(op: Operation): ParamInfo[] {
  const params = new Array<ParamInfo>();
  for (const param of values(op.request.parameters)) {
    if (param.implementation === ImplementationLocation.Method) {
      params.push({ name: param.language.go!.name, type: param.schema.language.go!.name });
    }
  }
  return params;
}

// returns the return signature where each entry is the type name
// e.g. [ '*string', 'error' ]
export function genereateReturnsInfo(op: Operation): string[] {
  if (length(op.responses) > 1) {
    throw console.error('multiple responses NYI');
  }
  const resp = op.responses![0];
  return [`*${resp.language.go!.name}`, 'error'];
}

// flattens out ParamInfo to return a complete parameter sig string
// e.g. "i int, s string, b bool"
export function generateParamsSig(paramInfo: ParamInfo[]): string {
  let params = new Array<string>();
  for (const param of values(paramInfo)) {
    params.push(`${param.name} ${param.type}`);
  }
  return params.join(', ');
}

// returns an array of just the parameter names
// e.g. [ 'i', 's', 'b' ]
export function extractParamNames(paramInfo: ParamInfo[]): string[] {
  let paramNames = new Array<string>();
  for (const param of values(paramInfo)) {
    paramNames.push(param.name);
  }
  return paramNames;
}
