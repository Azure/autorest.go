/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { values } from '@azure-tools/linq';
import { comment, camelCase } from '@azure-tools/codegen';
import { ArraySchema, CodeModel, DictionarySchema, Language, Parameter, Schema, SchemaType, Operation, GroupProperty, ImplementationLocation } from '@azure-tools/codemodel';
import { ImportManager } from './imports';
import { OperationNaming } from '../transform/namer';
import { aggregateParameters, PagerInfo } from '../common/helpers';

// returns the common source-file preamble (license comment, package name etc)
export async function contentPreamble(session: Session<CodeModel>): Promise<string> {
  const headerText = comment(await session.getValue('header-text', 'MISSING LICENSE HEADER'), '// ');
  // default namespce to the output folder
  const outputFolder = await session.getValue<string>('output-folder');
  // default namespace to equal the output directory name as they have to match.
  const namespace = outputFolder.substr(outputFolder.lastIndexOf('/') + 1);
  let text = `${headerText}\n\n`;
  text += `package ${namespace}\n\n`;
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
      if (schema.language.go!.discriminator) {
        return schema.language.go!.discriminator;
      }
      return schema.language.go!.name;
    default:
      return schema.language.go!.name;
  }
}

// returns the parameters for the internal request creator method.
// e.g. "i int, s string"
export function getCreateRequestParametersSig(op: Operation): string {
  const methodParams = getMethodParameters(op);
  const params = new Array<string>();
  for (const methodParam of values(methodParams)) {
    params.push(`${camelCase(methodParam.language.go!.name)} ${formatParameterTypeName(methodParam)}`);
  }
  return params.join(', ');
}

// returns the complete collection of method parameters
export function getMethodParameters(op: Operation): Parameter[] {
  const params = new Array<Parameter>();
  const paramGroups = new Array<GroupProperty>();
  for (const param of values(aggregateParameters(op))) {
    if (param.implementation === ImplementationLocation.Client) {
      // client params are passed via the receiver
      continue;
    } else if (param.schema.type === SchemaType.Constant) {
      // don't generate a parameter for a constant
      continue;
    } else if (param.language.go!.paramGroup) {
      // param groups will be added after individual params
      if (!paramGroups.includes(param.language.go!.paramGroup)) {
        paramGroups.push(param.language.go!.paramGroup);
      }
      continue;
    }
    params.push(param);
  }
  // move global optional params to the end of the slice
  params.sort(sortParametersByRequired);
  // add any parameter groups.  optional group goes last
  paramGroups.sort((a: GroupProperty, b: GroupProperty) => {
    if (a.required === b.required) {
      return 0;
    }
    if (a.required && !b.required) {
      return -1;
    }
    return 1;
  })
  for (const paramGroup of values(paramGroups)) {
    let name = camelCase(paramGroup.language.go!.name);
    if (!paramGroup.required) {
      name = 'options';
    }
    params.push(paramGroup);
  }
  return params;
}

export function generatePagerReturnInstance(op: Operation, imports: ImportManager): string {
  let text = '';
  const info = <OperationNaming>op.language.go!;
  // split param list into individual params
  const reqParams = getCreateRequestParametersSig(op).split(',');
  // keep the parameter names from the name/type tuples
  for (let i = 0; i < reqParams.length; ++i) {
    reqParams[i] = reqParams[i].trim().split(' ')[0];
  }
  text += `\treturn &${camelCase(op.language.go!.pageableType.name)}{\n`;
  text += `\t\tclient: client,\n`;
  text += `\t\trequest: req,\n`;
  text += `\t\tresponder: client.${info.protocolNaming.responseMethod},\n`;
  const pager = <PagerInfo>op.language.go!.pageableType;
  if (op.language.go!.paging.member) {
    // find the location of the nextLink param
    const nextLinkOpParams = getMethodParameters(op.language.go!.paging.nextLinkOperation);
    let found = false;
    for (let i = 0; i < nextLinkOpParams.length; ++i) {
      if (nextLinkOpParams[i].schema.type === SchemaType.String && nextLinkOpParams[i].language.go!.name.startsWith('next')) {
        // found it
        reqParams.splice(i, 0, `*resp.${pager.schema.language.go!.name}.${pager.nextLink}`);
        found = true;
        break;
      }
    }
    if (!found) {
      throw console.error(`failed to find nextLink parameter for operation ${op.language.go!.paging.nextLinkOperation.language.go!.name}`);
    }
    text += `\t\tadvancer: func(resp *${pager.schema.language.go!.responseType.name}) (*azcore.Request, error) {\n`;
    text += `\t\t\treturn client.${camelCase(op.language.go!.paging.member)}CreateRequest(${reqParams.join(', ')})\n`;
    text += '\t\t},\n';
  } else {
    imports.add('fmt');
    imports.add('net/url');
    let resultTypeName = pager.schema.language.go!.name;
    if (pager.schema.serialization?.xml?.name) {
      // xml can specifiy its own name, prefer that if available
      resultTypeName = pager.schema.serialization.xml.name;
    }
    text += `\t\tadvancer: func(resp *${pager.schema.language.go!.responseType.name}) (*azcore.Request, error) {\n`;
    text += `\t\t\tu, err := url.Parse(*resp.${resultTypeName}.${pager.nextLink})\n`;
    text += `\t\t\tif err != nil {\n`;
    text += `\t\t\t\treturn nil, fmt.Errorf("invalid ${pager.nextLink}: %w", err)\n`;
    text += `\t\t\t}\n`;
    text += `\t\t\tif u.Scheme == "" {\n`;
    text += `\t\t\t\treturn nil, fmt.Errorf("no scheme detected in ${pager.nextLink} %s", *resp.${resultTypeName}.${pager.nextLink})\n`;
    text += `\t\t\t}\n`;
    text += `\t\t\treturn azcore.NewRequest(http.MethodGet, *u), nil\n`;
    text += `\t\t},\n`;
  }
  text += `\t}, nil\n`;
  return text;

}