/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { capitalize, comment } from '@azure-tools/codegen';
import { Session } from '@autorest/extension-base';
import { CodeModel, ConstantSchema, GroupProperty, ImplementationLocation, Parameter, SchemaType } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { commentLength } from '../common/helpers';
import { contentPreamble, hasDescription, sortAscending } from './helpers';
import { ImportManager } from './imports';

// Creates the content in options.go
export async function generateOptions(session: Session<CodeModel>): Promise<string> {
  const paramGroups = <Array<GroupProperty>>session.model.language.go!.parameterGroups;
  if (!paramGroups) {
    return '';
  }

  const imports = new ImportManager();
  let optionsText = await contentPreamble(session);
  let content = '';

  paramGroups.sort((a: GroupProperty, b: GroupProperty) => { return sortAscending(a.schema.language.go!.name, b.schema.language.go!.name); });
  for (const paramGroup of values(paramGroups)) {
    paramGroup.originalParameter.sort((a: Parameter, b: Parameter) => { return sortAscending(a.language.go!.name, b.language.go!.name); });
    content += emit(paramGroup, imports);
  }

  optionsText += imports.text();
  optionsText += content;
  return optionsText;
}

function emit(paramGroup: GroupProperty, imports: ImportManager): string {
  let text = '';
  if (hasDescription(paramGroup.schema.language.go!)) {
    text += `${comment(paramGroup.schema.language.go!.description, '// ', undefined, commentLength)}\n`;
  }
  text += `type ${paramGroup.schema.language.go!.name} struct {\n`;

  if (paramGroup.originalParameter.length === 0) {
    // this is an optional params placeholder struct
    text += '\t// placeholder for future optional parameters\n';
  } else {
    // used to track when to add an extra \n between fields that have comments
    let first = true;

    for (const param of values(paramGroup.originalParameter)) {
      if (param.implementation === ImplementationLocation.Client && !paramGroup.language.go!.groupedClientParams) {
        // don't add globals to the per-method options struct
        continue;
      }

      imports.addImportForSchemaType(param.schema);
      if (hasDescription(param.language.go!)) {
        if (!first) {
          // add an extra new-line between fields IFF the field
          // has a comment and it's not the very first one.
          text += '\n';
        }
        text += `\t${comment(param.language.go!.description, '// ', undefined, commentLength)}\n`;
      }

      let typeName = param.schema.language.go!.name;
      if (param.schema.type === SchemaType.Constant) {
        // for constants we use the underlying type name
        typeName = (<ConstantSchema>param.schema).valueType.language.go!.name;
      }

      let pointer = '*';
      if (param.required || param.language.go!.byValue === true) {
        pointer = '';
      }
      text += `\t${capitalize(param.language.go!.name)} ${pointer}${typeName}\n`;
      first = false;
    }
  }

  text += '}\n\n';
  return text;
}
