/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { camelCase, pascalCase } from '@azure-tools/codegen';
import { CodeModel, ObjectSchema } from '@azure-tools/codemodel';
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
  discriminators.sort((a: ObjectSchema, b: ObjectSchema) => { return sortAscending(a.language.go!.discriminator, b.language.go!.discriminator) });
  for (const disc of values(discriminators)) {
    // this is the root discriminator (FishType in the test server)
    const rootDisc = disc.language.go!.discriminator;

    // this is used to track any sub-hierarchies (SalmonType, SharkType in the test server)
    const roots = new Array<ObjectSchema>();
    roots.push(disc);

    // constant definition
    // only generate one set from the root as it contains all possible values
    text += 'const (\n';
    // TODO: sort
    for (const val of values(disc.discriminator!.all)) {
      const objSchema = <ObjectSchema>val;
      text += `\t${objSchema.language.go!.discriminatorEnum} = "${objSchema.discriminatorValue!}"\n`;
      if (objSchema.discriminator) {
        // add sub-hierarchy
        roots.push(objSchema);
      }
    }
    text += ')\n\n';

    // generate unmarshallers for each discriminator
    for (const root of values(roots)) {
      const discName = root.language.go!.discriminator;
      // scalar unmarshaller
      text += `func unmarshal${discName}(body []byte) (${discName}, error) {\n`;
      text += '\tvar m map[string]interface{}\n';
      text += '\tif err := json.Unmarshal(body, &m); err != nil {\n';
      text += '\t\treturn nil, err\n';
      text += '\t}\n';
      text += `\tvar b ${discName}\n`;
      text += `\tswitch m["${root.discriminator!.property.serializedName}"] {\n`;
      for (const val of values(root.discriminator!.all)) {
        const objSchema = <ObjectSchema>val;
        text += `\tcase ${camelCase(rootDisc)}${pascalCase(objSchema.discriminatorValue!)}:\n`;
        text += `\t\tb = &${val.language.go!.name}{}\n`;
      }
      text += '\tdefault:\n';
      text += `\t\tb = &${root.language.go!.name}{}\n`;
      text += '\t}\n';
      text += '\treturn b, json.Unmarshal(body, &b)\n';
      text += '}\n\n';

      // array unmarshaller
      text += `func unmarshal${discName}Array(body []byte) (*[]${discName}, error) {\n`;
      text += '\tvar rawMessages []*json.RawMessage\n';
      text += '\tif err := json.Unmarshal(body, &rawMessages); err != nil {\n';
      text += '\t\treturn nil, err\n';
      text += '\t}\n';
      text += `\tfArray := make([]${discName}, len(rawMessages))\n`;
      text += '\tfor index, rawMessage := range rawMessages {\n';
      text += `\t\tf, err := unmarshal${discName}(*rawMessage)\n`;
      text += '\t\tif err != nil {\n';
      text += '\t\t\treturn nil, err\n';
      text += '\t\t}\n';
      text += '\t\tfArray[index] = f\n';
      text += '\t}\n';
      text += '\treturn &fArray, nil\n';
      text += '}\n\n';
    }
  }
  // helper used in discriminator marshallers
  text += 'func strptr(s string) *string {\n';
  text += '\treturn &s\n';
  text += '}\n\n';
  return text;
}
