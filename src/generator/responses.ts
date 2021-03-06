/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { comment } from '@azure-tools/codegen';
import { CodeModel, ObjectSchema, Property } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { commentLength, isObjectSchema } from '../common/helpers';
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
    structs.push(respType);
    // if the response envelope contains a result envelope, generate that too
    if (respEnv.language.go!.resultEnv) {
      const resultEnv = <Property>respEnv.language.go!.resultEnv;
      if (isObjectSchema(resultEnv.schema)) {
        const resultType = generateStruct(imports, resultEnv.schema.language.go!, resultEnv.schema.properties)
        generateUnmarshallerForResultEnvelope(resultType);
        structs.push(resultType);
      }
    }
  }
  imports.add('net/http');
  text += imports.text();
  structs.sort((a: StructDef, b: StructDef) => { return sortAscending(a.Language.name, b.Language.name) });
  for (const struct of values(structs)) {
    text += struct.discriminator();
    text += struct.text();
    struct.Methods.sort((a: StructMethod, b: StructMethod) => { return sortAscending(a.name, b.name) });
    for (const method of values(struct.Methods)) {
      if (method.desc.length > 0) {
        text += `${comment(method.desc, '// ', undefined, commentLength)}\n`;
      }
      text += method.text;
    }
  }
  return text;
}

// if the response envelope has a result envelope, check if the result envelope requires an unmarshaller
function generateUnmarshallerForResultEnvelope(structDef: StructDef) {
  // if the response envelope contains a discriminated type we need an unmarshaller
  let found = false;
  for (const prop of values(structDef.Properties)) {
    if (prop.isDiscriminator) {
      found = true;
      break;
    }
  }
  if (!found) {
    return;
  }
  const receiver = structDef.receiverName();
  let unmarshaller = `func (${receiver} *${structDef.Language.name}) UnmarshalJSON(data []byte) error {\n`;
  // add a custom unmarshaller to the response envelope
  let type = '';
  for (const prop of values(structDef.Properties)) {
    if (prop.isDiscriminator) {
      type = prop.schema.language.go!.discriminatorInterface;
      break;
    }
  }
  if (type === '') {
    throw new Error(`failed to the discriminated type field for response envelope ${structDef.Language.name}`);
  }
  unmarshaller += `\tres, err := unmarshal${type}(data)\n`;
  unmarshaller += '\tif err != nil {\n';
  unmarshaller += '\t\treturn err\n';
  unmarshaller += '\t}\n';
  unmarshaller += `\t${receiver}.${type} = res\n`;
  unmarshaller += '\treturn nil\n';
  unmarshaller += '}\n\n';
  structDef.Methods.push({ name: 'UnmarshalJSON', desc: `UnmarshalJSON implements the json.Unmarshaller interface for type ${structDef.Language.name}.`, text: unmarshaller });
}
