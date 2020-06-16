/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment } from '@azure-tools/codegen';
import { ArraySchema, CodeModel, DictionarySchema, Language, Parameter, Schema, SchemaType } from '@azure-tools/codemodel';

// returns the common source-file preamble (license comment, package name etc)
export async function contentPreamble(session: Session<CodeModel>): Promise<string> {
  const headerText = comment(await session.getValue('header-text', 'MISSING LICENSE HEADER'), '// ');
  let text = `${headerText}\n\n`;
  text += `package ${session.model.language.go!.packageName}\n\n`;
  return text;
}

// returns true if the language contains a description
export function hasDescription(lang: Language): boolean {
  return (lang.description !== undefined && lang.description.length > 0 && !lang.description.startsWith('MISSING'));
}

// used to sort strings in ascending order
export function sortAscending(a: string, b: string): number {
  return a < b ? -1 : a > b ? 1 : 0;
}

// returns the type name with possible * prefix
export function formatParameterTypeName(param: Parameter): string {
  const typeName = substituteDiscriminator(param.schema);
  if (param.required) {
    return typeName;
  }
  return `*${typeName}`;
}

// returns true if the parameter should not be URL encoded
export function skipURLEncoding(param: Parameter): boolean {
  if (param.extensions) {
    return param.extensions['x-ms-skip-url-encoding'] === true;
  }
  return false;
}

// sorts parameters by their required state, ordering required before optional
export function sortParametersByRequired(a: Parameter, b: Parameter): number {
  if (a.required === b.required) {
    return 0;
  }
  if (a.required && !b.required) {
    return -1;
  }
  return 1;
}

// if a field is a discriminator use the interface type instead
export function substituteDiscriminator(schema: Schema): string {
  switch (schema.type) {
    case SchemaType.Array:
      const arraySchema = <ArraySchema>schema;
      const arrayElem = <Schema>arraySchema.elementType;
      return `[]${substituteDiscriminator(arrayElem)}`;
    case SchemaType.Dictionary:
      const dictSchema = <DictionarySchema>schema;
      const dictElem = <Schema>dictSchema.elementType;
      return `map[string]${substituteDiscriminator(dictElem)}`;
    case SchemaType.Object:
      if (schema.language.go!.discriminatorInterface) {
        return schema.language.go!.discriminatorInterface;
      }
      return schema.language.go!.name;
    default:
      return schema.language.go!.name;
  }
}
