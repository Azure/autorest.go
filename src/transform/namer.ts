/*  ---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *  --------------------------------------------------------------------------------------------  */

import { Session } from '@autorest/extension-base';
import { CodeModel, HttpHeader, Language } from '@azure-tools/codemodel';
import { visitor, clone, values } from '@azure-tools/linq';
import { CommonAcronyms, ReservedWords } from './mappings';
import { aggregateParameters, hasAdditionalProperties } from '../common/helpers';

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

export class protocolMethods implements protocolNaming {
  readonly requestMethod: string;
  readonly responseMethod: string;
  readonly errorMethod: string;

  constructor(name: string) {
    this.requestMethod = ensureNameCase(`${name}${requestMethodSuffix}`, true);
    this.responseMethod = ensureNameCase(`${name}${responseMethodSuffix}`, true);
    this.errorMethod = ensureNameCase(`${name}${errorMethodSuffix}`, true);
  }
}

// The namer creates idiomatic Go names for types, properties, operations etc.
export async function namer(session: Session<CodeModel>) {
  const model = session.model;

  if (model.language.go) {
    // this looks like it already has data for this model.
    // send back an error
    session.error('bad flavor', ['go:1000', 'already-processed'], model.language.go);
    throw new Error('Go Namer Failed');
  }

  // copy all the .language.default data into .language.go
  cloneLanguageInfo(model);
  // default namespce to the output folder
  const outputFolder = await session.getValue<string>('output-folder');
  model.language.go!.packageName = outputFolder.substr(outputFolder.lastIndexOf('/') + 1);

  const specType = await session.getValue('openapi-type');
  model.language.go!.openApiType = specType;
  const azureARM = await session.getValue('azure-arm', false);
  model.language.go!.azureARM = azureARM;
  const exportClients = await session.getValue('export-clients', false);
  model.language.go!.exportClients = exportClients;

  // pascal-case and capitzalize acronym names of objects and their fields
  for (const obj of values(model.schemas.objects)) {
    const details = <Language>obj.language.go;
    details.name = ensureNameCase(details.name);
    if (obj.discriminator) {
      // if this is a discriminator add the interface name
      details.discriminatorInterface = createPolymorphicInterfaceName(details.name);
      details.discriminatorTypes = new Array<string>();
      details.discriminatorTypes.push('*' + details.name);
      for (const child of values(obj.discriminator.all)) {
        details.discriminatorTypes.push('*' + child.language.go!.name);
      }
    }
    for (const prop of values(obj.properties)) {
      const details = <Language>prop.language.go;
      details.name = ensureNameCase(removePrefix(details.name, 'XMS'));
      if (hasAdditionalProperties(obj) && details.name === 'AdditionalProperties') {
        // this is the case where a type contains the generic additional properties
        // and also has a field named additionalProperties.  we rename the field.
        details.name = 'AdditionalProperties1';
      }
    }
  }

  const exportClient = session.model.language.go!.openApiType === 'arm' || exportClients;
  // pascal-case and capitzalize acronym operation groups and their operations
  for (const group of values(model.operationGroups)) {
    const groupDetails = <Language>group.language.go;
    // use the swagger title as the default name for operation groups that don't specify a group name
    if (groupDetails.name.length === 0) {
      groupDetails.name = session.model.info.title;
    }
    groupDetails.name = ensureNameCase(removePrefix(groupDetails.name, 'XMS'));
    groupDetails.clientName = `${groupDetails.name}Client`;
    if (groupDetails.name.endsWith('Client')) {
      // don't generate a name like FooClientClient
      groupDetails.clientName = groupDetails.name;
    }
    groupDetails.clientCtorName = `New${groupDetails.clientName}`;
    if (!exportClient) {
      groupDetails.clientName = ensureNameCase(<string>groupDetails.clientName, true);
      groupDetails.clientCtorName = (<string>groupDetails.clientCtorName).uncapitalize();
    }
    for (const op of values(group.operations)) {
      const details = <OperationNaming>op.language.go;
      // propagate these settings to each operation for ease of access
      details.azureARM = model.language.go!.azureARM
      details.openApiType = model.language.go!.openApiType
      details.name = ensureNameCase(details.name);
      // add the client name to the operation as it's needed all over the place
      details.clientName = groupDetails.clientName;
      for (const param of values(aggregateParameters(op))) {
        if (param.language.go!.name === '$host' || param.language.go!.name.toUpperCase() === 'URL') {
          param.language.go!.name = 'endpoint';
          continue;
        }
        const inParamGroup = param.extensions?.['x-ms-parameter-grouping'] || param.required !== true;
        const paramDetails = <Language>param.language.go;
        // if this is part of a param group struct then don't apply param naming rules to it
        paramDetails.name = ensureNameCase(removePrefix(paramDetails.name, 'XMS'), !inParamGroup);
        // fix up any param group names
        if (param.extensions?.['x-ms-parameter-grouping']) {
          if (param.extensions['x-ms-parameter-grouping'].name) {
            param.extensions['x-ms-parameter-grouping'].name = ensureNameCase(<string>param.extensions['x-ms-parameter-grouping'].name);
          } else if (param.extensions['x-ms-parameter-grouping'].postfix) {
            param.extensions['x-ms-parameter-grouping'].postfix = ensureNameCase(<string>param.extensions['x-ms-parameter-grouping'].postfix);
          }
        } else {
          // only escape the name if it's not in a parameter group struct
          paramDetails.name = getEscapedReservedName(paramDetails.name, 'Param');
        }
      }
      details.protocolNaming = new protocolMethods(details.name);
      if (op.language.go!.paging) {
        if (op.language.go!.paging.nextLinkName !== null) {
          // apply same naming logic as per struct fields
          op.language.go!.paging.nextLinkName = ensureNameCase((<string>op.language.go!.paging.nextLinkName));
        }
        if (op.language.go!.paging.member) {
          op.language.go!.paging.member = (<string>op.language.go!.paging.member).uncapitalize();
        }
      }
      for (const resp of values(op.responses)) {
        for (const header of values(resp.protocol.http!.headers)) {
          const head = <HttpHeader>header;
          head.language.go!.name = ensureNameCase(removePrefix(head.language.go!.name, 'XMS'));
        }
      }
    }
  }

  // fix up enum type and value names and capitzalize acronyms
  for (const enm of values(session.model.schemas.choices)) {
    enm.language.go!.name = ensureNameCase(enm.language.go!.name);
    // add PossibleValues func name
    enm.language.go!.possibleValuesFunc = `Possible${enm.language.go!.name}Values`;
    for (const choice of values(enm.choices)) {
      const details = <Language>choice.language.go;
      details.name = `${enm.language.go?.name}${ensureNameCase(removePrefix(details.name, 'XMS'))}`;
    }
  }
  for (const enm of values(session.model.schemas.sealedChoices)) {
    enm.language.go!.name = ensureNameCase(enm.language.go!.name);
    // add PossibleValues func name
    enm.language.go!.possibleValuesFunc = `Possible${enm.language.go!.name}Values`;
    for (const choice of values(enm.choices)) {
      const details = <Language>choice.language.go;
      details.name = `${enm.language.go?.name}${ensureNameCase(removePrefix(details.name, 'XMS'))}`;
    }
  }

  for (const globalParam of values(session.model.globalParameters)) {
    const details = <Language>globalParam.language.go;
    const inParamGroup = globalParam.extensions?.['x-ms-parameter-grouping'] || globalParam.required !== true;
    // if this is part of a param group struct then don't apply param naming rules to it
    details.name = getEscapedReservedName(ensureNameCase(removePrefix(details.name, 'XMS'), !inParamGroup), 'Param');
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

// make sure that reserved words are escaped
function getEscapedReservedName(name: string, appendValue: string): string {
  if (ReservedWords.includes(name)) {
    name += appendValue;
  }

  return name;
}

// used in ensureNameCase() to track which names have already been transformed.
// this improves efficiency and also fixes some corner-cases where renaming the
// same thing in succession gives the wrong result due to how the regex works.
// e.g. SubscriptionId -> subscriptionID -> subscriptionid
const gRenamed = new Map<string, boolean>();

export function ensureNameCase(name: string, lowerFirst?: boolean): string {
  if (gRenamed.has(name) && gRenamed.get(name) === lowerFirst) {
    return name;
  }
  let reconstructed = '';
  // split the word into multiple words, either on Unicode word boundaries
  // or a defined set of separation characters, *preserving* the existing casing.
  // we cannot use deconstruct() as it is *not* case-preserving.
  // remove any entries that are empty and return the array.
  const words = values(name.split(new RegExp('(\\p{Lu}\\p{Ll}+\\d*)|\\.|_|@|-|\\s|\\$', 'gmu'))).where(s => s !== undefined && s.trim() !== '').toArray();
  for (let i = 0; i < words.length; ++i) {
    let word = words[i];
    // for params, lower-case the first segment
    if (lowerFirst && i === 0) {
      word = word.toLowerCase();
    } else {
      for (const tla of values(CommonAcronyms)) {
        // perform a case-insensitive match against the list of TLAs
        const match = word.match(new RegExp(tla, 'i'));
        if (match) {
          // replace the match with its upper-case version
          word = word.replace(match[0], match[0].toUpperCase());
        }
      }
      word = word.capitalize();
    }
    reconstructed += word;
  }
  gRenamed.set(reconstructed, lowerFirst === true);
  return reconstructed;
}

export function removePrefix(name: string, prefix: string): string {
  // perform case-insensitive comparison
  const nameU = name.toUpperCase();
  const prefixU = prefix.toUpperCase();
  for (let i = 0; i < prefixU.length; i++) {
    if (prefixU[i] !== nameU[i]) {
      return name
    }
  }

  return name.slice(prefix.length);
}

export function createPolymorphicInterfaceName(base: string): string {
  return base + 'Classification';
}
