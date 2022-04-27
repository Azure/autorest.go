/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { comment } from '@azure-tools/codegen';
import { CodeModel, ObjectSchema, Property } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { commentLength } from '../common/helpers';
import { contentPreamble, sortAscending } from './helpers';
import { ImportManager } from './imports';
import { generateStruct, StructDef, StructMethod } from './structs';

// Creates the content in response_types.go
export async function generateResponses(session: Session<CodeModel>): Promise<string> {
  let text = await contentPreamble(session);
  const imports = new ImportManager();
  const responseEnvelopes = <Array<ObjectSchema>>session.model.language.go!.responseEnvelopes;
  if (responseEnvelopes.length === 0) {
    return '';
  }
  const structs = new Array<StructDef>();
  for (const respEnv of values(responseEnvelopes)) {
    const respType = generateStruct(imports, respEnv.language.go!, respEnv.properties);
    generateUnmarshallerForResponeEnvelope(respType, imports);
    structs.push(respType);
  }
  text += imports.text();
  structs.sort((a: StructDef, b: StructDef) => { return sortAscending(a.Language.name, b.Language.name) });
  for (const struct of values(structs)) {
    text += struct.discriminator();
    text += struct.text();
    struct.SerDeMethods.sort((a: StructMethod, b: StructMethod) => { return sortAscending(a.name, b.name) });
    for (const method of values(struct.SerDeMethods)) {
      if (method.desc.length > 0) {
        text += `${comment(method.desc, '// ', undefined, commentLength)}\n`;
      }
      text += method.text;
    }
  }
  return text;
}

// check if the response envelope requires an unmarshaller
function generateUnmarshallerForResponeEnvelope(structDef: StructDef, imports: ImportManager) {
  // if the response envelope contains a discriminated type we need an unmarshaller
  let discriminatorProp: Property | undefined;
  // in addition, if it's an LRO operation that returns a scalar, we will also need one
  let nonEmbeddedProp: Property | undefined;
  for (const prop of values(structDef.Properties)) {
    if (prop.isDiscriminator) {
      discriminatorProp = prop;
      break;
    } else if (structDef.Language.forLRO && !prop.language.go!.embeddedType && !prop.language.go!.fromHeader) {
      nonEmbeddedProp = prop;
    }
  }
  if (!discriminatorProp && !nonEmbeddedProp) {
    return;
  }
  const receiver = structDef.receiverName();
  let unmarshaller = `func (${receiver} *${structDef.Language.name}) UnmarshalJSON(data []byte) error {\n`;
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
    throw new Error(`unhandled case for response envelope ${structDef.Language.name}`);
  }
  unmarshaller += '}\n\n';
  structDef.SerDeMethods.push({ name: 'UnmarshalJSON', desc: `UnmarshalJSON implements the json.Unmarshaller interface for type ${structDef.Language.name}.`, text: unmarshaller });
}
