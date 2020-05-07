/*  ---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *  --------------------------------------------------------------------------------------------  */

import { pascalCase, camelCase } from '@azure-tools/codegen';
import { Session } from '@azure-tools/autorest-extension-base';
import { CodeModel, Language } from '@azure-tools/codemodel';
import { length, visitor, clone, values } from '@azure-tools/linq';
import { CommonAcronyms, ReservedWords } from './mappings';
import { aggregateParameters, isLROOperation } from '../common/helpers';

const requestMethodSuffix = 'CreateRequest';
const responseMethodSuffix = 'HandleResponse';
const errorMethodSuffix = 'HandleError';

// contains extended naming information for operations
export interface OperationNaming extends Language {
  protocolNaming: protocolNaming
}

interface protocolNaming {
  requestMethod: string;
  responseMethod: string;
  errorMethod: string;
}

class protocolMethods implements protocolNaming {
  readonly requestMethod: string;
  readonly responseMethod: string;
  readonly errorMethod: string;

  constructor(name: string) {
    this.requestMethod = `${name}${requestMethodSuffix}`;
    this.responseMethod = `${name}${responseMethodSuffix}`;
    this.errorMethod = `${name}${errorMethodSuffix}`;
  }
}

// The namer creates idiomatic Go names for types, properties, operations etc.
export async function namer(session: Session<CodeModel>) {
  const model = session.model;

  if (model.language.go) {
    // this looks like it already has data for this model.
    // send back an error
    session.error('bad flavor', ['go:1000', 'already-processed'], model.language.go);
    throw Error('Go Namer Failed');
  }

  // copy all the .language.default data into .language.go
  cloneLanguageInfo(model);

  // pascal-case and capitzalize acronym names of objects and their fields
  for (const obj of values(model.schemas.objects)) {
    const details = <Language>obj.language.go;
    details.name = getEscapedReservedName(capitalizeAcronyms(pascalCase(details.name)), 'Model');
    if (obj.discriminator) {
      // if this is a discriminator add the interface name
      details.discriminator = `${details.name}Type`;
    }
    for (const prop of values(obj.properties)) {
      const details = <Language>prop.language.go;
      details.name = getEscapedReservedName(removePrefix(capitalizeAcronyms(pascalCase(details.name)), 'XMS'), 'Field');
    }
  }

  // pascal-case and capitzalize acronym operation groups and their operations
  for (const group of values(model.operationGroups)) {
    const groupDetails = <Language>group.language.go;
    // use the swagger title as the default name for operation groups that don't specify a group name
    if (groupDetails.name.length === 0) {
      groupDetails.name = pascalCase(session.model.info.title);
    }
    groupDetails.name = capitalizeAcronyms(pascalCase(groupDetails.name));
    groupDetails.clientName = `${groupDetails.name}Operations`;
    if (groupDetails.name === 'Operations') {
      // if the group name is 'Operations' don't name it 'OperationsOperations'
      groupDetails.clientName = groupDetails.clientName;
    }
    for (const op of values(group.operations)) {
      const details = <OperationNaming>op.language.go;
      details.name = getEscapedReservedName(capitalizeAcronyms(pascalCase(details.name)), 'Method');
      // add the client name to the operation as it's needed all over the place
      details.clientName = camelCase(groupDetails.clientName);
      if (isLROOperation(op)) {
        op.language.go!.methodPrefix = 'Begin';
      }
      for (const param of values(aggregateParameters(op))) {
        const paramDetails = <Language>param.language.go;
        paramDetails.name = getEscapedReservedName(removePrefix(camelCase(paramDetails.name), 'XMS'), 'Parameter');
      }
      details.protocolNaming = new protocolMethods(camelCase(details.name));
      if (op.language.go!.paging && op.language.go!.paging.nextLinkName !== null) {
        op.language.go!.paging.nextLinkName = pascalCase(op.language.go!.paging.nextLinkName);
      }
    }
  }

  // fix up enum type and value names and capitzalize acronyms
  for (const enm of values(session.model.schemas.choices)) {
    // add PossibleValues func name
    enm.language.go!.possibleValuesFunc = `Possible${enm.language.go!.name}Values`;
    for (const choice of values(enm.choices)) {
      const details = <Language>choice.language.go;
      details.name = `${enm.language.go?.name}${removePrefix(capitalizeAcronyms(pascalCase(details.name)), 'XMS')}`;
    }
  }
  for (const enm of values(session.model.schemas.sealedChoices)) {
    // add PossibleValues func name
    enm.language.go!.possibleValuesFunc = `Possible${enm.language.go!.name}Values`;
    for (const choice of values(enm.choices)) {
      const details = <Language>choice.language.go;
      details.name = `${enm.language.go?.name}${removePrefix(capitalizeAcronyms(pascalCase(details.name)), 'XMS')}`;
    }
  }

  for (const globalParam of values(session.model.globalParameters)) {
    const details = <Language>globalParam.language.go;
    details.name = removePrefix(capitalizeAcronyms(details.name), 'XMS');
  }
  return session;
}

function cloneLanguageInfo(graph: any) {
  // make sure recursively that every language field has Go language info
  for (const { index, instance } of visitor(graph)) {
    if (index === 'language' && instance.default && !instance.go) {
      instance.go = clone(instance.default, false, undefined, undefined, ['schema', 'origin']);
    }
  }
}

// make sure that common acronyms are capitalized
// NOTE: this function does not perform a case insensitive check considering scenarios where this would cause problems
// for example 'curl' would end up as 'cURL' if we did case insensitive checks
function capitalizeAcronyms(name: string): string {
  for (const word of CommonAcronyms) {
    name = name.replace(word, word.toUpperCase());
  }
  return name;
}

// make sure that reserved words are escaped
function getEscapedReservedName(name: string, appendValue: string): string {
  if (name === null) {
    throw new Error('GetEscapedReservedName: Cannot pass in a null value for "name" parameter');
  }
  if (appendValue === null) {
    throw new Error('GetEscapedReservedName: Cannot pass in a null value for "appendValue" parameter');
  }

  if (ReservedWords.includes(name)) {
    name += appendValue;
  }

  return name;
}

export function removePrefix(name: string, prefix: string): string {
  if (name === null) {
    throw new Error('removePrefix: Cannot pass in a null value for "name" parameter');
  }
  if (prefix === null) {
    throw new Error('removePrefix: Cannot pass in a null value for "prefix" parameter');
  }

  for (var i = 0; i < prefix.length; i++) {
    if (prefix[i] != name[i]) {
      return name
    }
  }

  return name.slice(prefix.length);
}
