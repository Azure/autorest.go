/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel, ObjectSchema, Property } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { isArraySchema, isDictionarySchema } from '../common/helpers';
import { contentPreamble, sortAscending } from './helpers';
import { ImportManager } from './imports';

// Creates the content in polymorphic_helpers.go
export async function generatePolymorphicHelpers(session: Session<CodeModel>): Promise<string> {
  if (!session.model.language.go!.discriminators) {
    // no polymorphic types
    return '';
  }
  const discriminators = <Array<ObjectSchema>>session.model.language.go!.discriminators.filter((d: ObjectSchema) => !d.language.go!.omitType);
  if (discriminators.length === 0) {
    // all polymorphic types omitted
    return '';
  }
  let text = await contentPreamble(session);
  const imports = new ImportManager();
  imports.add('encoding/json');
  text += imports.text();
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
  const scalars = new Set<string>();
  const arrays = new Set<string>();
  const maps = new Set<string>();
  const trackDisciminator = function(prop: Property) {
    if (prop.schema.language.go!.discriminatorInterface) {
      scalars.add(prop.schema.language.go!.discriminatorInterface);
    } else if (isArraySchema(prop.schema) && prop.schema.elementType.language.go!.discriminatorInterface) {
      scalars.add(prop.schema.elementType.language.go!.discriminatorInterface);
      arrays.add(prop.schema.elementType.language.go!.discriminatorInterface);
    } else if (isDictionarySchema(prop.schema) && prop.schema.elementType.language.go!.discriminatorInterface) {
      scalars.add(prop.schema.elementType.language.go!.discriminatorInterface);
      maps.add(prop.schema.elementType.language.go!.discriminatorInterface);
    }
  }
  // calculate which discriminator helpers we actually need to generate
  for (const obj of values(session.model.schemas.objects)) {
    for (const prop of values(obj.properties)) {
      trackDisciminator(prop);
    }
  }
  for (const respEnv of values(<Array<ObjectSchema>>session.model.language.go!.responseEnvelopes)) {
    if (respEnv.language.go!.resultProp) {
      const resultProp = <Property>respEnv.language.go!.resultProp;
      if (resultProp.isDiscriminator) {
        trackDisciminator(resultProp);
      }
    }
  }
  if (scalars.size === 0 && arrays.size === 0 && maps.size === 0) {
    // this is a corner-case that can happen when all the discriminated types
    // are error types.  there's a bug in M4 that incorrectly annotates such
    // types as 'output', 'exception' in the usage however it's really just
    // 'exception'.  until this is fixed, we can wind up here.
    return '';
  }
  discriminators.sort((a: ObjectSchema, b: ObjectSchema) => { return sortAscending(a.language.go!.discriminatorInterface, b.language.go!.discriminatorInterface) });
  for (const disc of values(discriminators)) {
    // generate unmarshallers for each discriminator
    const discName = disc.language.go!.discriminatorInterface;
    // scalar unmarshaller
    if (scalars.has(discName)) {
      text += `func unmarshal${discName}(rawMsg json.RawMessage) (${discName}, error) {\n`;
      text += '\tif rawMsg == nil {\n';
      text += '\t\treturn nil, nil\n';
      text += '\t}\n';
      text += '\tvar m map[string]any\n';
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
    }

    // array unmarshaller
    if (arrays.has(discName)) {
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
    }

    // map unmarshaller
    if (maps.has(discName)) {
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
  }
  return text;
}
