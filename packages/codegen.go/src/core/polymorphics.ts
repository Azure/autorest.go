/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../../codemodel.go/src/index.js';
import * as helpers from './helpers.js';
import { ImportManager } from './imports.js';

/**
 * Creates the content for the polymorphic_helpers.go file.
 *
 * @param pkg contains the package content
 * @returns the text for the file or the empty string
 */
export function generatePolymorphicHelpers(pkg: go.FakePackage | go.PackageContent): string {
  const content = pkg.kind === 'fake' ? pkg.parent : pkg;
  if (content.interfaces.length === 0) {
    // no polymorphic types
    return '';
  }

  let text = helpers.contentPreamble(pkg);
  const imports = new ImportManager(pkg);
  imports.add('encoding/json');
  if (pkg.kind === 'fake') {
    // content is being generated into a separate package, add the necessary import
    imports.addForPkg(pkg.parent);
  }

  text += imports.text();

  const scalars = new Set<string>();
  const arrays = new Set<string>();
  const maps = new Set<string>();

  // we know there are polymorphic types but we don't know how they're used.
  // i.e. are they vanilla fields, elements in a slice, or values in a map.
  // polymorphic types within maps/slices will also need the scalar helpers.
  const trackDisciminator = function (type: go.WireType) {
    switch (type.kind) {
      case 'interface':
        scalars.add(type.name);
        break;
      case 'map': {
        const leafType = helpers.recursiveUnwrapMapSlice(type);
        if (leafType.kind === 'interface') {
          scalars.add(leafType.name);
          maps.add(leafType.name);
        }
        break;
      }
      case 'slice': {
        const leafType = helpers.recursiveUnwrapMapSlice(type);
        if (leafType.kind === 'interface') {
          scalars.add(leafType.name);
          arrays.add(leafType.name);
        }
        break;
      }
    }
  };

  // calculate which discriminator helpers we actually need to generate

  if (pkg.kind === 'fake') {
    // when generating for the fakes server, we must look at operation parameters instead of return values
    for (const client of pkg.parent.clients) {
      for (const method of client.methods) {
        for (const param of method.parameters) {
          trackDisciminator(param.type);
        }
      }
    }
  } else {
    for (const model of pkg.models) {
      for (const field of model.fields) {
        trackDisciminator(field.type);
      }
    }

    for (const respEnv of pkg.responseEnvelopes) {
      switch (respEnv.result?.kind) {
        case 'monomorphicResult':
          switch (respEnv.result.monomorphicType.kind) {
            case 'map':
              trackDisciminator(respEnv.result.monomorphicType.valueType);
              break;
            case 'slice':
              trackDisciminator(respEnv.result.monomorphicType.elementType);
              break;
          }
          break;
        case 'polymorphicResult':
          trackDisciminator(respEnv.result.interface);
          break;
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

  let prefix = '';
  if (pkg.kind === 'fake') {
    // content is being generated into a separate package, set the type name prefix
    prefix = `${go.getPackageName(pkg.parent)}.`;
  }

  const indent = new helpers.Indentation();

  for (const interfaceType of content.interfaces) {
    // generate unmarshallers for each discriminator

    // scalar unmarshaller
    if (scalars.has(interfaceType.name)) {
      text += `func unmarshal${interfaceType.name}(rawMsg json.RawMessage) (${prefix}${interfaceType.name}, error) {\n`;
      text += `${indent.get()}if rawMsg == nil || string(rawMsg) == "null" {\n`;
      text += `${indent.push().get()}return nil, nil\n`;
      text += `${indent.pop().get()}}\n`;
      text += `${indent.get()}var m map[string]any\n`;
      text += `${indent.get()}if err := json.Unmarshal(rawMsg, &m); err != nil {\n`;
      text += `${indent.push().get()}return nil, err\n`;
      text += `${indent.pop().get()}}\n`;
      text += `${indent.get()}var b ${prefix}${interfaceType.name}\n`;
      text += `${indent.get()}switch m["${interfaceType.discriminatorField}"] {\n`;
      for (const possibleType of interfaceType.possibleTypes) {
        let disc = helpers.formatLiteralValue(possibleType.discriminatorValue!, true);
        // when the discriminator value is an enum, cast the const as a string
        if (possibleType.discriminatorValue!.type.kind === 'constant') {
          disc = `string(${prefix}${disc})`;
        }
        text += `${indent.get()}case ${disc}:\n`;
        indent.push();
        text += `${indent.get()}b = &${prefix}${possibleType.name}{}\n`;
        indent.pop();
      }
      text += `${indent.get()}default:\n`;
      text += `${indent.push().get()}b = &${prefix}${interfaceType.rootType.name}{}\n`;
      text += `${indent.pop().get()}}\n`;
      text += `${indent.get()}if err := json.Unmarshal(rawMsg, b); err != nil {\n`;
      text += `${indent.push().get()}return nil, err\n`;
      text += `${indent.pop().get()}}\n`;
      text += `${indent.get()}return b, nil\n`;
      text += '}\n\n';
    }

    // array unmarshaller
    if (arrays.has(interfaceType.name)) {
      text += `func unmarshal${interfaceType.name}Array(rawMsg json.RawMessage) ([]${prefix}${interfaceType.name}, error) {\n`;
      text += `${indent.get()}if rawMsg == nil || string(rawMsg) == "null" {\n`;
      text += `${indent.push().get()}return nil, nil\n`;
      text += `${indent.pop().get()}}\n`;
      text += `${indent.get()}var rawMessages []json.RawMessage\n`;
      text += `${indent.get()}if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {\n`;
      text += `${indent.push().get()}return nil, err\n`;
      text += `${indent.pop().get()}}\n`;
      text += `${indent.get()}fArray := make([]${prefix}${interfaceType.name}, len(rawMessages))\n`;
      text += `${indent.get()}for index, rawMessage := range rawMessages {\n`;
      indent.push();
      text += `${indent.get()}f, err := unmarshal${interfaceType.name}(rawMessage)\n`;
      text += `${indent.get()}if err != nil {\n`;
      text += `${indent.push().get()}return nil, err\n`;
      text += `${indent.pop().get()}}\n`;
      text += `${indent.get()}fArray[index] = f\n`;
      text += `${indent.pop().get()}}\n`;
      text += `${indent.get()}return fArray, nil\n`;
      text += '}\n\n';
    }

    // map unmarshaller
    if (maps.has(interfaceType.name)) {
      text += `func unmarshal${interfaceType.name}Map(rawMsg json.RawMessage) (map[string]${prefix}${interfaceType.name}, error) {\n`;
      text += `${indent.get()}if rawMsg == nil || string(rawMsg) == "null" {\n`;
      text += `${indent.push().get()}return nil, nil\n`;
      text += `${indent.pop().get()}}\n`;
      text += `${indent.get()}var rawMessages map[string]json.RawMessage\n`;
      text += `${indent.get()}if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {\n`;
      text += `${indent.push().get()}return nil, err\n`;
      text += `${indent.pop().get()}}\n`;
      text += `${indent.get()}fMap := make(map[string]${prefix}${interfaceType.name}, len(rawMessages))\n`;
      text += `${indent.get()}for key, rawMessage := range rawMessages {\n`;
      indent.push();
      text += `${indent.get()}f, err := unmarshal${interfaceType.name}(rawMessage)\n`;
      text += `${indent.get()}if err != nil {\n`;
      text += `${indent.push().get()}return nil, err\n`;
      text += `${indent.pop().get()}}\n`;
      text += `${indent.get()}fMap[key] = f\n`;
      text += `${indent.pop().get()}}\n`;
      text += `${indent.get()}return fMap, nil\n`;
      text += '}\n\n';
    }
  }
  return text;
}
