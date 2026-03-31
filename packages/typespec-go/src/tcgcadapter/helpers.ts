/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as tcgc from '@azure-tools/typespec-client-generator-core';
import * as tsp from '@typespec/compiler';
import * as go from '../../../codemodel.go/src/index.js';

// called for models and model fields. for the former, the field param will be undefined
export function adaptXMLInfo(pkg: go.PackageContent, decorators: Array<tcgc.DecoratorInfo>, field?: go.ModelField): go.XMLInfo | undefined {
  // if there are no decorators and this isn't a slice
  // type in a model field then do nothing
  if (decorators.length === 0 && (!field || field.type.kind !== 'slice')) {
    return undefined;
  }

  const xmlInfo = new go.XMLInfo();
  if (field && field.type.kind === 'slice') {
    // for tsp, arrays are wrapped by default
    xmlInfo.wraps = go.getTypeDeclaration(field.type.elementType, pkg);
  }

  const handleName = (decorator: tcgc.DecoratorInfo): void => {
    if (field) {
      xmlInfo.name = <string>decorator.arguments['name'];
    } else {
      // when applied to a model, it means the model's XML element
      // node has a different name than the model.
      xmlInfo.wrapper = <string>decorator.arguments['name'];
    }
  };

  for (const decorator of decorators) {
    switch (decorator.name) {
      case 'TypeSpec.@encodedName':
        if (decorator.arguments['mimeType'] === 'application/xml') {
          handleName(decorator);
        }
        break;
      case 'TypeSpec.Xml.@attribute':
        xmlInfo.attribute = true;
        break;
      case 'TypeSpec.Xml.@name':
        handleName(decorator);
        break;
      case 'TypeSpec.Xml.@unwrapped':
        // unwrapped can only be applied to fields
        if (field) {
          switch (field.type.kind) {
            case 'slice':
              // unwrapped slice. default to using the serialized name
              xmlInfo.wraps = undefined;
              xmlInfo.name = field.serializedName;
              break;
            case 'string':
              // an unwrapped string means it's text
              xmlInfo.text = true;
              break;
          }
        }
        break;
    }
  }

  return xmlInfo;
}

/**
 * returns any XMLInfo available for the provided type or undefined
 *
 * @param type the type to inspect for XMLInfo
 * @returns the XMLInfo or undefined
 */
export function hasXMLInfo(type: go.WireType): go.XMLInfo | undefined {
  if ('xml' in type) {
    return type.xml;
  }
  return undefined;
}

/**
 * returns true if model is a polymorphic root type.
 *
 * @param model the model to inspect
 * @returns true if the model is a polymorphic root
 */
export function isPolymorphicRoot(model: tcgc.SdkModelType): boolean {
  if (model.discriminatedSubtypes) {
    // when there are sub-types we know for sure it's a polymorphic root
    return true;
  } else if (model.discriminatorProperty && !model.discriminatorValue) {
    // we can land here if it's a root but has no child types
    return true;
  } else {
    return false;
  }
}

/**
 * returns true if the specified type doesn't need to be pointer-to-type
 * because it's implicitly nil-able.
 *
 * @param type the type to inspect
 * @returns true if the type is implicitly nil-able
 */
export function isTypePassedByValue(type: tcgc.SdkType): boolean {
  if (type.kind === 'nullable') {
    type = type.type;
  }
  return type.kind === 'unknown' || type.kind === 'array' || type.kind === 'bytes' || type.kind === 'dict' || (type.kind === 'model' && isPolymorphicRoot(type));
}

/** contains the set of client options */
const clientOptionKinds = ['monomorphicResponseFieldName', 'omitSerdeMethods', 'responseEnvelopeName'] as const;
export type ClientOptionKind = (typeof clientOptionKinds)[number];

/**
 * returns the value of the specified client option if found in the provided decorators, otherwise returns undefined.
 * reports a warning diagnostic for any client options found with an unrecognized name
 *
 * @param option the client option to search for
 * @param src the source to search for the client option
 * @param program the tsp Program currently in scope, used to report warning diagnostics for unhandled key/value pairs
 * @returns the value of the client option if found, otherwise undefined
 */
export function getClientOption(option: ClientOptionKind, src: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation> | tcgc.SdkModelType, program: tsp.Program): string | undefined {
  const clientOptions = src.decorators.filter((decorator) => decorator.name === 'Azure.ClientGenerator.Core.@clientOption');
  for (const clientOption of clientOptions) {
    const optionName = <string>clientOption.arguments['name'];
    if (!clientOptionKinds.includes(optionName as ClientOptionKind)) {
      program.reportDiagnostic({
        code: 'InvalidClientOption',
        severity: 'warning',
        message: `invalid client option ${optionName}`,
        target: src.__raw?.node ?? tsp.NoTarget,
      });
      continue;
    }

    // ensure that the option is valid for the provided source
    let valid = false;
    switch (src.kind) {
      case 'model':
        valid = optionName === 'omitSerdeMethods';
        break;
      default:
        // this branch is currently all method types
        valid = optionName === 'monomorphicResponseFieldName' || optionName === 'responseEnvelopeName';
    }

    if (!valid) {
      program.reportDiagnostic({
        code: 'InvalidClientOption',
        severity: 'warning',
        message: `inapplicable client option ${optionName}`,
        target: src.__raw?.node ?? tsp.NoTarget,
      });
      continue;
    }

    const optionValue = <string>clientOption.arguments['value'];
    if (optionName === option) {
      return optionValue;
    }
  }
  return undefined;
}
