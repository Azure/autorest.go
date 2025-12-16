/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { comment } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import * as go from '../../../codemodel.go/src/index.js';
import * as helpers from './helpers.js';
import { ImportManager } from './imports.js';
import { CodegenError } from './errors.js';

export interface ResponsesSerDe {
  responses: string;
  serDe: string;
}

/**
 * Creates the content for the responses.go file.
 * 
 * @param pkg contains the package content
 * @param options the emitter options
 * @returns the text for the file or the empty string
 */
export function generateResponses(pkg: go.PackageContent, options: go.Options): ResponsesSerDe {
  if (pkg.responseEnvelopes.length === 0) {
    return {
      responses: '',
      serDe: ''
    };
  }

  const imports = new ImportManager();
  const serdeImports = new ImportManager();
  let responses = helpers.contentPreamble(pkg);
  let serDe = '';
  let respContent = '';
  let serdeContent = '';

  for (const respEnv of pkg.responseEnvelopes) {
    respContent += emit(respEnv, imports);
    if (options.generateFakes) {
      serdeContent += generateMarshaller(respEnv, serdeImports);
    }
    serdeContent += generateUnmarshaller(respEnv, serdeImports);
  }

  responses += imports.text();
  responses += respContent;

  if (serdeContent.length > 0) {
    serDe = helpers.contentPreamble(pkg);
    serDe += serdeImports.text();
    serDe += serdeContent;
  }

  return {
    responses: responses,
    serDe: serDe
  };
}

function generateMarshaller(respEnv: go.ResponseEnvelope, imports: ImportManager): string {
  let text = '';
  if (go.isLROMethod(respEnv.method) && respEnv.result?.kind === 'polymorphicResult') {
    // fakes require a custom marshaller for polymorphics results so that the data is in the correct shape.
    // without it, the response envelope type name is the outer type which is incorrect.
    imports.add('encoding/json');
    const receiver = respEnv.name[0].toLowerCase();
    text += `${comment(`MarshalJSON implements the json.Marshaller interface for type ${respEnv.name}.`, '// ', undefined, helpers.commentLength)}\n`;
    text += `func (${receiver} ${respEnv.name}) MarshalJSON() ([]byte, error) {\n`;
    // TODO: this doesn't include any headers. however, LROs with header responses are currently broken :(
    text += `\treturn json.Marshal(${receiver}.${go.getTypeDeclaration(respEnv.result.interface)})\n}\n\n`;
  }
  return text;
}

// check if the response envelope requires an unmarshaller
function generateUnmarshaller(respEnv: go.ResponseEnvelope, imports: ImportManager): string {
  // if the response envelope contains a discriminated type we need an unmarshaller
  let polymorphicRes: go.PolymorphicResult | undefined;
  // in addition, if it's an LRO operation that returns a scalar, we will also need one
  let monomorphicRes: go.MonomorphicResult | undefined;
  if (respEnv.result?.kind === 'polymorphicResult') {
    polymorphicRes = respEnv.result;
  } else if (go.isLROMethod(respEnv.method) && respEnv.result?.kind === 'monomorphicResult') {
    monomorphicRes = respEnv.result;
  }

  if (!polymorphicRes && !monomorphicRes) {
    // no unmarshaller required
    return '';
  }

  const receiver = respEnv.name[0].toLowerCase();
  let unmarshaller = `${comment(`UnmarshalJSON implements the json.Unmarshaller interface for type ${respEnv.name}.`, '// ', undefined, helpers.commentLength)}\n`;
  unmarshaller += `func (${receiver} *${respEnv.name}) UnmarshalJSON(data []byte) error {\n`;

  // add a custom unmarshaller to the response envelope
  if (polymorphicRes) {
    const type = polymorphicRes.interface.name;
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
    throw new CodegenError('InternalError', `unhandled case for response envelope ${respEnv.name}`);
  }
  unmarshaller += '}\n\n';
  return unmarshaller;
}

function emit(respEnv: go.ResponseEnvelope, imports: ImportManager): string {
  let text = helpers.formatDocComment(respEnv.docs);

  text += `type ${respEnv.name} struct {\n`;
  if (!respEnv.result && respEnv.headers.length === 0) {
    // this is an empty response envelope
    text += '\t// placeholder for future response values\n';
  } else {
    // fields will contain the merged headers and response field so they can be sorted together
    const fields = new Array<{docs: go.Docs, field: string}>();

    // used to track when to add an extra \n between fields that have comments
    let first = true;

    if (respEnv.result) {
      if (respEnv.result.kind === 'modelResult' || respEnv.result.kind === 'polymorphicResult') {
        // anonymously embedded type always goes first
        text += helpers.formatDocComment(respEnv.result.docs);
        text += `\t${go.getTypeDeclaration(go.getResultType(respEnv.result))}\n`;
        first = false;
      } else {
        const type = go.getResultType(respEnv.result);
        imports.addImportForType(type);

        let tag = '';
        if (respEnv.result.kind === 'monomorphicResult' && respEnv.result.format === 'XML') {
          // only emit tags for XML; JSON uses custom marshallers/unmarshallers
          if (respEnv.result.xml?.wraps) {
            tag = ` \`xml:"${respEnv.result.xml.wraps}"\``;
          } else if (respEnv.result.xml?.name) {
            tag = ` \`xml:"${respEnv.result.xml.name}"\``;
          }
        }

        let byValue = true;
        if (respEnv.result.kind === 'monomorphicResult') {
          byValue = respEnv.result.byValue;
        }

        fields.push({docs: respEnv.result.docs, field: `\t${respEnv.result.fieldName} ${helpers.star(byValue)}${go.getTypeDeclaration(type)}${tag}\n`});
      }
    }

    for (const header of values(respEnv.headers)) {
      imports.addImportForType(header.type);
      let byValue = true;
      if (header.kind === 'headerScalarResponse') {
        byValue = header.byValue;
      }
      fields.push({docs: header.docs, field: `\t${header.fieldName} ${helpers.star(byValue)}${go.getTypeDeclaration(header.type)}\n`});
    }

    fields.sort((a: {desc?: string, field: string}, b: {desc?: string, field: string}) => { return helpers.sortAscending(a.field, b.field); });

    for (const field of fields) {
      if (field.docs.summary || field.docs.description) {
        if (!first) {
          // add an extra new-line between fields IFF the field
          // has a comment and it's not the very first one.
          text += '\n';
        }
        text += helpers.formatDocComment(field.docs);
      }
      text += field.field;
      first = false;
    }
  }

  text += '}\n\n';
  return text;
}
