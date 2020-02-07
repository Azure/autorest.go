/*  ---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *  --------------------------------------------------------------------------------------------  */

import { serialize, pascalCase, camelCase } from '@azure-tools/codegen'
import { Host, startSession, Session } from '@azure-tools/autorest-extension-base'
import { codeModelSchema, CodeModel, Language, ObjectSchema, SchemaType, Schema } from '@azure-tools/codemodel'
import { length, visitor, clone, values } from '@azure-tools/linq'
import { CommonAcronyms, ReservedWords } from './mappings'

// The namer creates idiomatic Go names for types, properties, operations etc.
export async function namer(host: Host) {
  const debug = await host.GetValue('debug') || false;

  try {
    const session = await startSession<CodeModel>(host, {}, codeModelSchema);

    await process(session);

    // output the model to the pipeline
    host.WriteFile('code-model-v4.yaml', serialize(session.model), undefined, 'code-model-v4');

  } catch (E) {
    if (debug) {
      console.error(`${__filename} - FAILURE  ${JSON.stringify(E)} ${E.stack}`);
    }
    throw E;
  }
}

const requestMethodSuffix = 'CreateRequest';
const responseMethodSuffix = 'HandleResponse';

// contains extended naming information for operations
export interface OperationNaming extends Language {
  protocolNaming: protocolNaming
}

interface protocolNaming {
  requestMethod: string;
  responseMethod: string;
}

class protocolMethods implements protocolNaming {
  readonly requestMethod: string;
  readonly responseMethod: string;

  constructor(name: string) {
    this.requestMethod = `${name}${requestMethodSuffix}`;
    this.responseMethod = `${name}${responseMethodSuffix}`;
  }
}

async function process(session: Session<CodeModel>) {
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
    for (const prop of values(obj.properties)) {
      const details = <Language>prop.language.go;
      details.name = getEscapedReservedName(capitalizeAcronyms(pascalCase(details.name)), 'Field');
    }
  }

  // pascal-case and capitzalize acronym operation groups and their operations
  for (const group of values(model.operationGroups)) {
    const details = <Language>group.language.go;
    const opGroupName = capitalizeAcronyms(pascalCase(group.$key));
    details.name = capitalizeAcronyms(pascalCase(details.name));
    // we don't call GetEscapedReservedName here since any operation group that uses a reserved word will have 'Operations' attached to it
    details.clientName = `${details.name}Operations`;
    // this sets Operations as the default name for operation groups that don't specify a group name
    if (length(details.name) === 0) {
      details.name = `Operations`;
    }
    for (const op of values(group.operations)) {
      const details = <OperationNaming>op.language.go;
      details.name = getEscapedReservedName(capitalizeAcronyms(pascalCase(details.name)), 'Method');
      for (const param of values(op.request.parameters)) {
        const paramDetails = <Language>param.language.go;
        paramDetails.name = getEscapedReservedName(camelCase(paramDetails.name), 'Parameter');
      }
      details.protocolNaming = new protocolMethods(details.name);
      // fix up response type name and description
      if (length(op.responses) > 1) {
        throw console.error('multiple responses NYI');
      }
      const resp = op.responses![0];
      const name = `${opGroupName}${op.language.go!.name}Response`;
      resp.language.go!.name = name;
      resp.language.go!.description = `${name} contains the response from method ${group.language.go!.name}.${op.language.go!.name}.`;
    }
  }

  // fix up enum type and value names and capitzalize acronyms
  for (const enm of values(session.model.schemas.choices)) {
    // add PossibleValues func name
    enm.language.go!.possibleValuesFunc = `Possible${enm.language.go!.name}Values`;
    for (const choice of values(enm.choices)) {
      const details = <Language>choice.language.go;
      details.name = `${enm.language.go?.name}${capitalizeAcronyms(pascalCase(details.name.toLowerCase()))}`;
    }
  }
  for (const enm of values(session.model.schemas.sealedChoices)) {
    // add PossibleValues func name
    enm.language.go!.possibleValuesFunc = `Possible${enm.language.go!.name}Values`;
    for (const choice of values(enm.choices)) {
      const details = <Language>choice.language.go;
      details.name = `${enm.language.go?.name}${capitalizeAcronyms(pascalCase(details.name.toLowerCase()))}`;
    }
  }

  for (const globalParam of values(session.model.globalParameters)) {
    const details = <Language>globalParam.language.go;
    details.name = capitalizeAcronyms(details.name);
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
