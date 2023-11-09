/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { comment } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import { GoCodeModel, ResponseEnvelope, PolymorphicResult, MonomorphicResult } from '../gocodemodel/gocodemodel';
import { getResultPossibleType, getTypeDeclaration, isLROMethod, isMonomorphicResult, isPolymorphicResult, isModelResult } from '../gocodemodel/gocodemodel';
import { commentLength, contentPreamble, sortAscending } from './helpers';
import { ImportManager } from './imports';
import { getStar } from './models';

// Creates the content in response_types.go
export async function generateResponses(codeModel: GoCodeModel): Promise<string> {
  if (codeModel.responseEnvelopes. length === 0) {
    return '';
  }

  const imports = new ImportManager();
  let text = contentPreamble(codeModel);
  let content = '';

  for (const respEnv of codeModel.responseEnvelopes) {
    content += emit(respEnv, imports);
    if (codeModel.options.generateFakes) {
      content += generateMarshaller(respEnv, imports);
    }
    content += generateUnmarshaller(respEnv, imports);
  }
  
  text += imports.text();
  text += content;
  return text;
}

function generateMarshaller(respEnv: ResponseEnvelope, imports: ImportManager): string {
  let text = '';
  if (isLROMethod(respEnv.method) && respEnv.result && isPolymorphicResult(respEnv.result)) {
    // fakes require a custom marshaller for polymorphics results so that the data is in the correct shape.
    // without it, the response envelope type name is the outer type which is incorrect.
    imports.add('encoding/json');
    const receiver = respEnv.name[0].toLowerCase();
    text += `${comment(`MarshalJSON implements the json.Marshaller interface for type ${respEnv.name}.`, '// ', undefined, commentLength)}\n`;
    text += `func (${receiver} ${respEnv.name}) MarshalJSON() ([]byte, error) {\n`;
    // TODO: this doesn't include any headers. however, LROs with header responses are currently broken :(
    text += `\treturn json.Marshal(${receiver}.${getTypeDeclaration(respEnv.result.interfaceType)})\n}\n\n`;
  }
  return text;
}

// check if the response envelope requires an unmarshaller
function generateUnmarshaller(respEnv: ResponseEnvelope, imports: ImportManager): string {
  // if the response envelope contains a discriminated type we need an unmarshaller
  let polymorphicRes: PolymorphicResult | undefined;
  // in addition, if it's an LRO operation that returns a scalar, we will also need one
  let monomorphicRes: MonomorphicResult | undefined;
  if (respEnv.result && isPolymorphicResult(respEnv.result)) {
    polymorphicRes = respEnv.result;
  } else if (isLROMethod(respEnv.method) && respEnv.result && isMonomorphicResult(respEnv.result)) {
    monomorphicRes = respEnv.result;
  }

  if (!polymorphicRes && !monomorphicRes) {
    // no unmarshaller required
    return '';
  }

  const receiver = respEnv.name[0].toLowerCase();
  let unmarshaller = `${comment(`UnmarshalJSON implements the json.Unmarshaller interface for type ${respEnv.name}.`, '// ', undefined, commentLength)}\n`;
  unmarshaller += `func (${receiver} *${respEnv.name}) UnmarshalJSON(data []byte) error {\n`;

  // add a custom unmarshaller to the response envelope
  if (polymorphicRes) {
    const type = polymorphicRes.interfaceType.name;
    unmarshaller += `\tres, err := unmarshal${type}(data)\n`;
    unmarshaller += '\tif err != nil {\n';
    unmarshaller += '\t\treturn err\n';
    unmarshaller += '\t}\n';
    unmarshaller += `\t${receiver}.${type} = res\n`;
    unmarshaller += '\treturn nil\n';
  } else if (monomorphicRes) {
    imports.add('encoding/json');
    unmarshaller += `\treturn json.Unmarshal(data, &${receiver}.${monomorphicRes.fieldName})\n`;
  } else {
    throw new Error(`unhandled case for response envelope ${respEnv.name}`);
  }
  unmarshaller += '}\n\n';
  return unmarshaller;
}

function emit(respEnv: ResponseEnvelope, imports: ImportManager): string {
  let text = '';
  if (respEnv.description) {
    text += `${comment(respEnv.description, '// ', undefined, commentLength)}\n`;
  }

  text += `type ${respEnv.name} struct {\n`;
  if (!respEnv.result && respEnv.headers.length === 0) {
    // this is an empty response envelope
    text += '\t// placeholder for future response values\n';
  } else {
    // fields will contain the merged headers and response field so they can be sorted together
    const fields = new Array<{desc?: string, field: string}>();

    // used to track when to add an extra \n between fields that have comments
    let first = true;

    if (respEnv.result) {
      if (isModelResult(respEnv.result) || isPolymorphicResult(respEnv.result)) {
        // anonymously embedded type always goes first
        if (respEnv.result.description) {
          text += `\t${comment(respEnv.result.description, '// ', undefined, commentLength)}\n`;
        }
        text += `\t${getTypeDeclaration(getResultPossibleType(respEnv.result))}\n`;
        first = false;
      } else {
        let desc: string | undefined;
        if (respEnv.result.description) {
          desc = `\t${comment(respEnv.result.description, '// ', undefined, commentLength)}\n`;
        }

        const type = getResultPossibleType(respEnv.result);
        imports.addImportForType(type);

        let tag = '';
        if (isMonomorphicResult(respEnv.result) && respEnv.result.format === 'XML') {
          // only emit tags for XML; JSON uses custom marshallers/unmarshallers
          if (respEnv.result.xml?.wraps) {
            tag = ` \`xml:"${respEnv.result.xml.wraps}"\``;
          } else if (respEnv.result.xml?.name) {
            tag = ` \`xml:"${respEnv.result.xml.name}"\``;
          }
        }

        fields.push({desc: desc, field: `\t${respEnv.result.fieldName} ${getStar(respEnv.result.byValue)}${getTypeDeclaration(type)}${tag}\n`});
      }
    }

    for (const header of values(respEnv.headers)) {
      imports.addImportForType(header.type);
      let desc: string | undefined;
      if (header.description) {
        desc = `\t${comment(header.description, '// ', undefined, commentLength)}\n`;
      }
      fields.push({desc: desc, field: `\t${header.fieldName} ${getStar(header.byValue)}${getTypeDeclaration(header.type)}\n`});
    }

    fields.sort((a: {desc?: string, field: string}, b: {desc?: string, field: string}) => { return sortAscending(a.field, b.field); });

    for (const field of fields) {
      if (field.desc) {
        if (!first) {
          // add an extra new-line between fields IFF the field
          // has a comment and it's not the very first one.
          text += '\n';
        }
        text += field.desc;
      }
      text += field.field;
      first = false;
    }
  }

  text += '}\n\n';
  return text;
}
