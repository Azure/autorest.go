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
const clientOptionKinds = ['monomorphicResponseFieldName', 'omitSerdeMethods', 'preserveContentTypeHeader', 'responseEnvelopeName'] as const;
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
export function getClientOption<T extends boolean | string>(
  option: ClientOptionKind,
  src: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation> | tcgc.SdkModelType,
  program: tsp.Program,
): T | undefined {
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
        valid = optionName === 'monomorphicResponseFieldName' || optionName === 'preserveContentTypeHeader' || optionName === 'responseEnvelopeName';
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

    const optionValue = <T>clientOption.arguments['value'];
    if (optionName === option) {
      return optionValue;
    }
  }
  return undefined;
}

/**
 * returns true if the response header should be omitted from the response envelope.
 *
 * currently, this is the case for `Content-Type` response headers whose value is
 * a literal/constant, unless the method opts in via the `preserveContentTypeHeader`
 * client option. callers (response envelope construction and example mapping)
 * must keep the two sites in sync; use this helper from both.
 *
 * @param httpHeader the tcgc response header to inspect
 * @param sdkMethod the tcgc service method that owns the response
 * @param program the tsp Program currently in scope
 */
export function isOmittedResponseHeader(
  httpHeader: tcgc.SdkServiceResponseHeader,
  sdkMethod: tcgc.SdkServiceMethod<tcgc.SdkHttpOperation>,
  program: tsp.Program,
): boolean {
  if (!httpHeader.serializedName.match(/^content-type$/i)) {
    return false;
  }
  // literal/constant header values are folded into the request/response shape and
  // are therefore not surfaced as fields on the response envelope.
  if (httpHeader.type.kind !== 'constant' && httpHeader.type.kind !== 'enumvalue') {
    return false;
  }
  const preserveHeader = getClientOption<boolean>('preserveContentTypeHeader', sdkMethod, program);
  return preserveHeader !== true;
}

/**
 * returns the record for the specified decorator's arguments if it exists
 *
 * @param name the name of the decorator to find
 * @param decorators the array of decorators to search
 * @returns the decorator's record of arguments or undefined
 */
/* eslint-disable-next-line @typescript-eslint/no-explicit-any */
export function hasDecorator(name: string, decorators: Array<tcgc.DecoratorInfo>): Record<string, any> | undefined {
  for (const decorator of decorators) {
    if (decorator.name.endsWith(name)) {
      return decorator.arguments;
    }
  }
  return undefined;
}

/**
 * returns a doc comment for a literal that's used as a client-side default.
 *
 * @param literal the literal for which to create the doc comment
 * @returns the doc comment
 */
export function getClientDefaultValueDoc(literal: go.Literal): string {
  let value = <string>literal.literal;
  switch (literal.type.kind) {
    case 'constant':
      if ((<go.ConstantValue>literal.literal)?.kind === 'constantValue') {
        value = (<go.ConstantValue>literal.literal).name;
      } else if (literal.type.type === 'string') {
        value = `${literal.type.name}("${<string>literal.literal}")`;
      } else {
        value = `${literal.type.name}(${<string>literal.literal})`;
      }
      break;
    case 'string':
      value = `"${value}"`;
      break;
  }
  return `The default value is ${value}.`;
}

/**
 * returns true if the type is an extensible enum (i.e. an enum or union-as-enum
 * with isFixed === false). nullable wrappers are unwrapped before the check.
 */
export function isExtensibleEnum(type: tcgc.SdkType): boolean {
  if (type.kind === 'nullable') {
    return isExtensibleEnum(type.type);
  }
  return type.kind === 'enum' && !type.isFixed;
}

/** returns true if model is a TypeSpec.Http.File type */
export function isHttpFileType(model: tcgc.SdkModelType): boolean {
  // we use startsWith('file') as there might be multiple
  // instances with different suffixes (e.g. File1, File2, etc)
  return model.namespace.toLowerCase() === 'typespec.http' && model.name.toLowerCase().startsWith('file');
}
