/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel, ObjectSchema } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { contentPreamble, sortAscending } from './helpers';
import { ImportManager } from './imports';

// Creates the content in polymorphic_helpers.go
export async function generatePolymorphicHelpers(session: Session<CodeModel>): Promise<string> {
  if (!session.model.language.go!.discriminators) {
    // no polymorphic types
    return '';
  }
  let text = await contentPreamble(session);
  const imports = new ImportManager();
  imports.add('encoding/json');
  text += imports.text();
  const discriminators = <Array<ObjectSchema>>session.model.language.go!.discriminators;
  // add any sub-hierarchies (SalmonType, SharkType in the test server) to the list
  for (const disc of values(discriminators)) {
    for (const val of values(disc.discriminator!.all)) {
      const objSchema = <ObjectSchema>val;
      // some hierarchies can overlap, so conditionally add
      if (objSchema.discriminator && !discriminators.includes(objSchema)) {
        discriminators.push(objSchema);
      }
    }
  }
  discriminators.sort((a: ObjectSchema, b: ObjectSchema) => { return sortAscending(a.language.go!.discriminatorInterface, b.language.go!.discriminatorInterface) });
  for (const disc of values(discriminators)) {
    // generate unmarshallers for each discriminator
    const discName = disc.language.go!.discriminatorInterface;
    if (disc.language.go!.internalErrorType) {
      text += `type ${disc.language.go!.internalErrorType} struct {\n`;
      text += `\twrapped ${discName}\n`;
      text += '}\n\n';
      const receiver = <string>disc.language.go!.internalErrorType[0];
      text += `func (${receiver} *${disc.language.go!.internalErrorType}) UnmarshalJSON(data []byte) (err error) {\n`;
      text += `\t${receiver}.wrapped, err = unmarshal${discName}(data)\n`;
      text += '\treturn\n';
      text += '}\n\n';
    }
    // scalar unmarshaller
    text += `func unmarshal${discName}(rawMsg json.RawMessage) (${discName}, error) {\n`;
    text += '\tif rawMsg == nil {\n';
    text += '\t\treturn nil, nil\n';
    text += '\t}\n';
    text += '\tvar m map[string]interface{}\n';
    text += '\tif err := json.Unmarshal(rawMsg, &m); err != nil {\n';
    text += '\t\treturn nil, err\n';
    text += '\t}\n';
    text += `\tvar b ${discName}\n`;
    text += `\tswitch m["${disc.discriminator!.property.serializedName}"] {\n`;
    for (const val of values(disc.discriminator!.all)) {
      const objSchema = <ObjectSchema>val;
      let disc = objSchema.discriminatorValue;
      // when the discriminator value is an enum, cast the const as a string
      if (disc![0] !== '"') {
        disc = `string(${disc})`;
      }
      text += `\tcase ${disc}:\n`;
      text += `\t\tb = &${val.language.go!.name}{}\n`;
    }
    text += '\tdefault:\n';
    text += `\t\tb = &${disc.language.go!.name}{}\n`;
    text += '\t}\n';
    text += '\treturn b, json.Unmarshal(rawMsg, b)\n';
    text += '}\n\n';

    // array unmarshaller
    text += `func unmarshal${discName}Array(rawMsg json.RawMessage) ([]${discName}, error) {\n`;
    text += '\tif rawMsg == nil {\n';
    text += '\t\treturn nil, nil\n';
    text += '\t}\n';
    text += '\tvar rawMessages []json.RawMessage\n';
    text += '\tif err := json.Unmarshal(rawMsg, &rawMessages); err != nil {\n';
    text += '\t\treturn nil, err\n';
    text += '\t}\n';
    text += `\tfArray := make([]${discName}, len(rawMessages))\n`;
    text += '\tfor index, rawMessage := range rawMessages {\n';
    text += `\t\tf, err := unmarshal${discName}(rawMessage)\n`;
    text += '\t\tif err != nil {\n';
    text += '\t\t\treturn nil, err\n';
    text += '\t\t}\n';
    text += '\t\tfArray[index] = f\n';
    text += '\t}\n';
    text += '\treturn fArray, nil\n';
    text += '}\n\n';

    // map unmarshaller
    text += `func unmarshal${discName}Map(rawMsg json.RawMessage) (map[string]${discName}, error) {\n`;
    text += '\tif rawMsg == nil {\n';
    text += '\t\treturn nil, nil\n';
    text += '\t}\n';
    text += '\tvar rawMessages map[string]json.RawMessage\n';
    text += '\tif err := json.Unmarshal(rawMsg, &rawMessages); err != nil {\n';
    text += '\t\treturn nil, err\n';
    text += '\t}\n';
    text += `\tfMap := make(map[string]${discName}, len(rawMessages))\n`;
    text += '\tfor key, rawMessage := range rawMessages {\n';
    text += `\t\tf, err := unmarshal${discName}(rawMessage)\n`;
    text += '\t\tif err != nil {\n';
    text += '\t\t\treturn nil, err\n';
    text += '\t\t}\n';
    text += '\t\tfMap[key] = f\n';
    text += '\t}\n';
    text += '\treturn fMap, nil\n';
    text += '}\n\n';
  }
  return text;
}
