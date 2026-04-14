/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as tcgc from '@azure-tools/typespec-client-generator-core';
import * as tsp from '@typespec/compiler';
import * as go from '../../../codemodel.go/src/index.js';

/** source data used to compute XMLInfo */
export interface XMLSourceInfo {
  /** the Go type name */
  goTypeName: string;

  /**
   * the original type name. can be different from
   * goTypeName in some cases (e.g. stuttering cleanup)
   */
  orTypeName: string;

  /** the model/model field type */
  type: go.WireType;

  /** XML serialization data when available */
  xml?: tcgc.XmlSerializationOptions;
}

/**
 * creates XMLInfo for models and model fields.
 * returns undefined if no XMLInfo is required.
 *
 * @param src the source information to adapt
 * @returns XMLInfo or undefined
 */
export function adaptXMLInfo(src: XMLSourceInfo): go.XMLInfo | undefined {
  const xmlInfo = new go.XMLInfo();
  let returnXMLInfo = false;

  if (src.xml?.name && src.xml.name !== src.goTypeName) {
    xmlInfo.name = src.xml.name;
    returnXMLInfo = true;
  }

  if (src.xml?.attribute) {
    xmlInfo.attribute = true;
    returnXMLInfo = true;
  }
  if (src.type.kind === 'slice') {
    const elementXMLInfo = hasXMLInfo(src.type.elementType);
    if (src.xml?.unwrapped === false) {
      if (src.xml.itemsName) {
        xmlInfo.wraps = src.xml.itemsName;
      } else if (elementXMLInfo?.name) {
        xmlInfo.wraps = elementXMLInfo.name;
      } else if (src.orTypeName !== src.goTypeName) {
        xmlInfo.wraps = src.orTypeName;
      } else {
        xmlInfo.wraps = src.goTypeName;
      }
      returnXMLInfo = true;
    } else if (elementXMLInfo?.name) {
      xmlInfo.name = elementXMLInfo.name;
      returnXMLInfo = true;
    } else if (src.orTypeName !== src.goTypeName) {
      // we can land here if the Go-specific type name was renamed to remove stuttering
      xmlInfo.name = src.orTypeName;
      returnXMLInfo = true;
    }
  } else if (src.xml?.unwrapped && src.type.kind === 'string') {
    // an unwrapped string means it's text
    xmlInfo.text = true;
    // the ",chardata" tag is mutually exclusive
    // with a name tag so clear it if set
    xmlInfo.name = undefined;
    returnXMLInfo = true;
  }

  return returnXMLInfo ? xmlInfo : undefined;
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
