/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment } from '@azure-tools/codegen';
import { CodeModel, ChoiceValue, ImplementationLocation, Language, Operation, Parameter, Response, Schema, SchemaResponse, Schemas, SchemaType, ArraySchema, DictionarySchema } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';

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

  addImportForSchemaType(schema: Schema) {
    switch (schema.type) {
      case SchemaType.Array:
        this.addImportForSchemaType((<ArraySchema>schema).elementType);
        break;
      case SchemaType.Dictionary:
        this.addImportForSchemaType((<DictionarySchema>schema).elementType);
        break;
      case SchemaType.Date:
      case SchemaType.DateTime:
      case SchemaType.Duration:
      case SchemaType.UnixTime:
        this.add('time');
        break;
      case SchemaType.Uri:
        this.add('net/url');
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
  return (lang.description !== undefined && lang.description.length > 0 && !lang.description.startsWith('MISSING'));
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
  global: boolean;
  required: boolean;
  isHost: boolean;
}

export class paramInfo implements ParamInfo {
  name: string;
  type: string;
  global: boolean;
  required: boolean;
  isHost: boolean;
  constructor(name: string, type: string, global: boolean, required: boolean, isHost: boolean) {
    this.name = name;
    this.type = type;
    this.global = global;
    this.required = required;
    this.isHost = isHost;
  }
}

// returns the type name with possible * prefix
export function formatParamInfoTypeName(param: ParamInfo): string {
  if (param.required) {
    return param.type;
  }
  return `*${param.type}`;
}

// aggregates the Parameter in op.parameters and the first request
export function aggregateParameters(op: Operation): Array<Parameter> {
  if (op.requests!.length > 1) {
    throw console.error('multiple requests NYI');
  }
  let params = new Array<Parameter>();
  if (op.parameters) {
    params = params.concat(op.parameters);
  }
  if (op.requests![0].parameters) {
    params = params.concat(op.requests![0].parameters);
  }
  return params;
}

// creates ParamInfo for the specified operation.
// each entry is tuple of param name/param type
export function generateParameterInfo(op: Operation): ParamInfo[] {
  const params = new Array<ParamInfo>();
  for (const param of values(aggregateParameters(op))) {
    if (param.schema.type === SchemaType.Constant) {
      // don't generate a parameter for a constant
      continue;
    }
    if (param.language.go!.name === 'host' || param.language.go!.name === '$host') {
      continue;
    }
    if (param.implementation === ImplementationLocation.Method && param.required !== true) {
      // omit method-optional params as they're grouped in the optional params type
      continue;
    }
    // include client and method params
    const global = param.implementation === ImplementationLocation.Client;
    let isHost = false;
    if (global) {
      isHost = (param.extensions ? (param.extensions!['x-ms-priority'] === 0 && param.extensions!['x-in'] === 'path') : false);
    }
    params.push(new paramInfo(param.language.go!.name, param.schema.language.go!.name, global, param.required === true, isHost));
  }
  // move global optional params to the end of the slice
  params.sort(sortParamInfoByRequired);
  // if there's a method-optional params struct add it last
  if (op.requests![0].language.go!.optionalParam) {
    params.push(new paramInfo('options', op.requests![0].language.go!.optionalParam.name, false, false, false));
  }

  return params;
}

// sorts ParamInfo objects by their required state, ordering required before optional
export function sortParamInfoByRequired(a: ParamInfo, b: ParamInfo): number {
  if (a.required === b.required) {
    return 0;
  }
  if (a.required && !b.required) {
    return -1;
  }
  return 1;
}

// returns the return signature where each entry is the type name
// e.g. [ '*string', 'error' ]
export function genereateReturnsInfo(op: Operation, forHandler: boolean): string[] {
  // TODO check this implementation, if any additional return information needs to be included for multiple responses
  const firstResp = op.responses![0];
  let returnType = '*http.Response';
  // must check pageable first as all pageable operations are also schema responses
  if (!forHandler && isPageableOperation(op)) {
    returnType = op.language.go!.pageableType.name;
  } else if (isSchemaResponse(firstResp)) {
    returnType = '*' + firstResp.schema.language.go!.responseType.name;
  }
  return [returnType, 'error'];
}

// flattens out ParamInfo to return a complete parameter sig string
// e.g. "i int, s string, b bool"
export function generateParamsSig(paramInfo: ParamInfo[], includeGlobal: boolean): string {
  let params = new Array<string>();
  for (const param of values(paramInfo)) {
    if ((param.global && !includeGlobal) || param.isHost) {
      continue;
    }
    params.push(`${param.name} ${formatParamInfoTypeName(param)}`);
  }
  return params.join(', ');
}

// represents an enum type and its values
export class EnumEntry {
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

// returns a collection containing all enum entries and their values
export function getEnums(schemas: Schemas): EnumEntry[] {
  // group all enum categories into a single array so they can be sorted
  const enums = new Array<EnumEntry>();
  for (const choice of values(schemas.choices)) {
    choice.choices.sort((a: ChoiceValue, b: ChoiceValue) => { return SortAscending(a.language.go!.name, b.language.go!.name); });
    const entry = new EnumEntry(choice.language.go!.name, choice.choiceType.language.go!.name, choice.language.go!.possibleValuesFunc, choice.choices);
    if (HasDescription(choice.language.go!)) {
      entry.desc = choice.language.go!.description;
    }
    enums.push(entry);
  }
  for (const choice of values(schemas.sealedChoices)) {
    if (choice.choices.length === 1) {
      continue;
    }
    const entry = new EnumEntry(choice.language.go!.name, choice.choiceType.language.go!.name, choice.language.go!.possibleValuesFunc, choice.choices);
    if (HasDescription(choice.language.go!)) {
      entry.desc = choice.language.go!.description;
    }
    enums.push(entry);
  }
  enums.sort((a: EnumEntry, b: EnumEntry) => { return SortAscending(a.name, b.name) });
  return enums;
}

// returns ArraySchema type predicate if the schema is an ArraySchema
export function isArraySchema(resp: Schema): resp is ArraySchema {
  return (resp as ArraySchema).elementType !== undefined;
}

// returns SchemaResponse type predicate if the response has a schema
export function isSchemaResponse(resp?: Response): resp is SchemaResponse {
  return (resp as SchemaResponse).schema !== undefined;
}

// returns true if the parameter should not be URL encoded
export function skipURLEncoding(param: Parameter): boolean {
  if (param.extensions) {
    return param.extensions['x-ms-skip-url-encoding'] === true;
  }
  return false;
}

export interface PagerInfo {
  name: string;
  schema: Schema;
  client: string;
  nextLink: string;
}

// returns true if the operation is pageable
export function isPageableOperation(op: Operation): boolean {
  return op.language.go!.paging && op.language.go!.paging.nextLinkName !== null;
}
