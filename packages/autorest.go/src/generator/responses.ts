/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { comment } from '@azure-tools/codegen';
import { CodeModel, ObjectSchema, Property } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { commentLength } from '../common/helpers';
import { contentPreamble, hasDescription, sortAscending } from './helpers';
import { ImportManager } from './imports';
import { getStar, getXMLSerialization } from './structs';

// Creates the content in response_types.go
export async function generateResponses(session: Session<CodeModel>): Promise<string> {
  const responseEnvelopes = <Array<ObjectSchema>>session.model.language.go!.responseEnvelopes;
  if (responseEnvelopes.length === 0) {
    return '';
  }

  const imports = new ImportManager();
  let text = await contentPreamble(session);
  let content = '';

  responseEnvelopes.sort((a: ObjectSchema, b: ObjectSchema) => { return sortAscending(a.language.go!.name, b.language.go!.name) });
  for (const respEnv of values(responseEnvelopes)) {
    respEnv.properties?.sort((a: Property, b: Property) => { return sortAscending(a.language.go!.name, b.language.go!.name) });
    content += emit(respEnv, imports);
    content += generateUnmarshallerForResponeEnvelope(respEnv, imports);
  }

  text += imports.text();
  text += content;
  return text;
}

// check if the response envelope requires an unmarshaller
function generateUnmarshallerForResponeEnvelope(respEnv: ObjectSchema, imports: ImportManager): string {
  // if the response envelope contains a discriminated type we need an unmarshaller
  let discriminatorProp: Property | undefined;
  // in addition, if it's an LRO operation that returns a scalar, we will also need one
  let nonEmbeddedProp: Property | undefined;
  for (const prop of values(respEnv.properties)) {
    if (prop.isDiscriminator) {
      discriminatorProp = prop;
      break;
    } else if (respEnv.language.go!.forLRO && !prop.language.go!.embeddedType && !prop.language.go!.fromHeader) {
      nonEmbeddedProp = prop;
    }
  }

  if (!discriminatorProp && !nonEmbeddedProp) {
    // no unmarshaller required
    return '';
  }

  const receiver = respEnv.language.go!.name[0].toLowerCase();
  let unmarshaller = `\t${comment(`UnmarshalJSON implements the json.Unmarshaller interface for type ${respEnv.language.go!.name}.`, '// ', undefined, commentLength)}\n`;
  unmarshaller += `func (${receiver} *${respEnv.language.go!.name}) UnmarshalJSON(data []byte) error {\n`;

  // add a custom unmarshaller to the response envelope
  if (discriminatorProp) {
    const type = discriminatorProp.schema.language.go!.discriminatorInterface;
    unmarshaller += `\tres, err := unmarshal${type}(data)\n`;
    unmarshaller += '\tif err != nil {\n';
    unmarshaller += '\t\treturn err\n';
    unmarshaller += '\t}\n';
    unmarshaller += `\t${receiver}.${type} = res\n`;
    unmarshaller += '\treturn nil\n';
  } else if (nonEmbeddedProp) {
    imports.add('encoding/json');
    unmarshaller += `\treturn json.Unmarshal(data, &${receiver}.${nonEmbeddedProp.language.go!.name})\n`;
  } else {
    throw new Error(`unhandled case for response envelope ${respEnv.language.go!.name}`);
  }
  unmarshaller += '}\n\n';
  return unmarshaller;
}

function emit(respEnv: ObjectSchema, imports: ImportManager): string {
  let text = '';
  if (hasDescription(respEnv.language.go!)) {
    text += `${comment(respEnv.language.go!.description, '// ', undefined, commentLength)}\n`;
  }

  text += `type ${respEnv.language.go!.name} struct {\n`
  if (!respEnv.properties) {
    // this is an empty response envelope
    text += '\t// placeholder for future response values\n';
  } else {
    // used to track when to add an extra \n between fields that have comments
    let first = true;

    // embedded properties go first
    for (const prop of values(respEnv.properties)) {
      if (prop.language.go!.embeddedType) {
        if (hasDescription(prop.language.go!)) {
          if (!first) {
            // add an extra new-line between fields IFF the field
            // has a comment and it's not the very first one.
            text += '\n';
          }
          text += `\t${comment(prop.language.go!.description, '// ', undefined, commentLength)}\n`;
        }
        text += `\t${prop.schema.language.go!.name}\n`;
        first = false;
      }
    }

    for (const prop of values(respEnv.properties)) {
      if (prop.language.go!.embeddedType) {
        continue;
      }
      imports.addImportForSchemaType(prop.schema);

      if (hasDescription(prop.language.go!)) {
        if (!first) {
          // add an extra new-line between fields IFF the field
          // has a comment and it's not the very first one.
          text += '\n';
        }
        text += `\t${comment(prop.language.go!.description, '// ', undefined, commentLength)}\n`;
      }

      let tag = '';
      // only emit tags for XML; JSON uses custom marshallers/unmarshallers
      if (respEnv.language.go!.marshallingFormat === 'xml') {
        tag = ` \`xml:"${getXMLSerialization(prop, respEnv.language.go!)}"\``;
      }

      text += `\t${prop.language.go!.name} ${getStar(prop.language.go!)}${prop.schema.language.go!.name}${tag}\n`;
      first = false;
    }
  }

  text += '}\n\n';
  return text;
}
