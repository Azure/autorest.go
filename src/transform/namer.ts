/*  ---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *  --------------------------------------------------------------------------------------------  */

import { Session } from '@autorest/extension-base';
import { capitalize, uncapitalize } from '@azure-tools/codegen';
import { CodeModel, HttpHeader, Language } from '@autorest/codemodel';
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
  internalMethod: string;
  requestMethod: string;
  responseMethod: string;
  errorMethod: string;
}

export class protocolMethods implements protocolNaming {
  readonly internalMethod: string;
  readonly requestMethod: string;
  readonly responseMethod: string;
  readonly errorMethod: string;

  constructor(name: string) {
    // uncapitalizing runs the risk of reserved name collision, e.g. Import -> import
    this.internalMethod = getEscapedReservedName(uncapitalize(name), 'Operation');
    this.requestMethod = ensureNameCase(`${name}${requestMethodSuffix}`, true);
    this.responseMethod = ensureNameCase(`${name}${responseMethodSuffix}`, true);
    this.errorMethod = ensureNameCase(`${name}${errorMethodSuffix}`, true);
  }
}

// returns the leaf folder name from the provided folder
function packageNameFromOutputFolder(folder: string): string {
  for (let i = folder.length - 1; i > -1; --i) {
    if (folder[i] === '/' || folder[i] === '\\') {
      return folder.substr(i + 1);
    }
  }
  // no path separator
  return folder;
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
  model.language.go!.packageName = packageNameFromOutputFolder(outputFolder);

  const specType = await session.getValue('openapi-type');
  model.language.go!.openApiType = specType;
  const azureARM = await session.getValue('azure-arm', false);
  model.language.go!.azureARM = azureARM;
  const exportClients = await session.getValue('export-clients', false);
  model.language.go!.exportClients = exportClients;
  const headAsBoolean = await session.getValue('head-as-boolean', false);
  model.language.go!.headAsBoolean = headAsBoolean;
  const groupParameters = await session.getValue('group-parameters', true);
  model.language.go!.groupParameters = groupParameters;

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
      (<Array<string>>details.discriminatorTypes).sort();
    }
    for (const prop of values(obj.properties)) {
      const details = <Language>prop.language.go;
      details.name = ensureNameCase(details.name);
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
    groupDetails.name = ensureNameCase(groupDetails.name);
    groupDetails.clientName = `${groupDetails.name}Client`;
    if (groupDetails.name.endsWith('Client')) {
      // don't generate a name like FooClientClient
      groupDetails.clientName = groupDetails.name;
    }
    groupDetails.clientCtorName = `New${groupDetails.clientName}`;
    if (!exportClient) {
      groupDetails.clientName = ensureNameCase(<string>groupDetails.clientName, true);
      groupDetails.clientCtorName = uncapitalize(groupDetails.clientCtorName);
    }
    for (const op of values(group.operations)) {
      const details = <OperationNaming>op.language.go;
      // propagate these settings to each operation for ease of access
      details.azureARM = model.language.go!.azureARM;
      details.openApiType = model.language.go!.openApiType;
      details.name = ensureNameCase(details.name);
      // add the client name to the operation as it's needed all over the place
      details.clientName = groupDetails.clientName;
      for (const param of values(aggregateParameters(op))) {
        if (param.language.go!.name === '$host' || param.language.go!.name.toUpperCase() === 'URL') {
          param.language.go!.name = 'endpoint';
          continue;
        }
        const inParamGroup = (param.extensions?.['x-ms-parameter-grouping'] && groupParameters) || param.required !== true;
        const paramDetails = <Language>param.language.go;
        // if this is part of a param group struct then don't apply param naming rules to it
        paramDetails.name = ensureNameCase(paramDetails.name, !inParamGroup);
        // fix up any param group names
        if (param.extensions?.['x-ms-parameter-grouping'] && groupParameters) {
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
        if (op.language.go!.paging.nextLinkName === '') {
          // fix up broken swaggers that incorrectly specify no next link name
          op.language.go!.paging.nextLinkName = null;
        }
        if (op.language.go!.paging.nextLinkName !== null) {
          // apply same naming logic as per struct fields
          op.language.go!.paging.nextLinkName = ensureNameCase((<string>op.language.go!.paging.nextLinkName));
        }
        if (op.language.go!.paging.member) {
          op.language.go!.paging.member = uncapitalize(op.language.go!.paging.member);
        }
      }
      for (const resp of values(op.responses)) {
        for (const header of values(resp.protocol.http!.headers)) {
          const head = <HttpHeader>header;
          head.language.go!.name = ensureNameCase(head.language.go!.name);
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
      details.name = `${enm.language.go?.name}${ensureNameCase(details.name)}`;
    }
  }
  for (const enm of values(session.model.schemas.sealedChoices)) {
    enm.language.go!.name = ensureNameCase(enm.language.go!.name);
    // add PossibleValues func name
    enm.language.go!.possibleValuesFunc = `Possible${enm.language.go!.name}Values`;
    for (const choice of values(enm.choices)) {
      const details = <Language>choice.language.go;
      details.name = `${enm.language.go?.name}${ensureNameCase(details.name)}`;
    }
  }

  for (const globalParam of values(session.model.globalParameters)) {
    const details = <Language>globalParam.language.go;
    const inParamGroup = globalParam.extensions?.['x-ms-parameter-grouping'] && groupParameters;
    // if this is part of a param group struct then don't apply param naming rules to it
    details.name = getEscapedReservedName(ensureNameCase(details.name, !inParamGroup), 'Param');
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
const gRenamed = new Map<string, boolean>();

// case-preserving version of deconstruct() that also splits on more path-separator characters
function deconstruct(identifier: string): string[] {
  return `${identifier}`.
    replace(/([a-z]+)([A-Z])/g, '$1 $2').
    replace(/(\d+)([a-z|A-Z]+)/g, '$1 $2').
    replace(/\b([A-Z]+)([A-Z])([a-z])/, '$1 $2$3').
    split(/[\W|_|\.|@|-|\s|\$]+/);
}

export function ensureNameCase(name: string, lowerFirst?: boolean): string {
  if (gRenamed.has(name) && gRenamed.get(name) === lowerFirst) {
    return name;
  }
  // XMS prefix requires special handling due to too many permutations that cause weird splits in the word
  name = name.replace(new RegExp('^(xms)', 'i'), 'XMS');
  let reconstructed = '';
  const words = deconstruct(name);
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
          let toReplace = match[0];
          if (match.length === 2) {
            // a capture group was specified, use it instead
            toReplace = match[1];
          }
          word = word.replace(toReplace, toReplace.toUpperCase());
        }
      }
      // note that capitalize() will convert the following acronyms to all upper-case
      // 'ip', 'os', 'ms', 'vm'
      word = capitalize(word);
    }
    reconstructed += word;
  }
  gRenamed.set(reconstructed, lowerFirst === true);
  return reconstructed;
}

function createPolymorphicInterfaceName(base: string): string {
  return base + 'Classification';
}
