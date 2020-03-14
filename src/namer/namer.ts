/*  ---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *  --------------------------------------------------------------------------------------------  */

import { serialize, pascalCase, camelCase } from '@azure-tools/codegen';
import { Host, startSession, Session } from '@azure-tools/autorest-extension-base';
import { codeModelSchema, CodeModel, Language, Parameter, SchemaType, SealedChoiceSchema } from '@azure-tools/codemodel';
import { length, visitor, clone, values } from '@azure-tools/linq';
import { CommonAcronyms, ReservedWords } from './mappings';
import { aggregateParameters, LanguageHeader } from '../generator/common/helpers';

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
      details.name = getEscapedReservedName(removePrefix(capitalizeAcronyms(pascalCase(details.name)), 'XMS'), 'Field');
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
      // track any optional parameters
      const optionalParams = new Array<Parameter>();
      for (const param of values(aggregateParameters(op))) {
        const paramDetails = <Language>param.language.go;
        paramDetails.name = getEscapedReservedName(removePrefix(camelCase(paramDetails.name), 'XMS'), 'Parameter');
        // this is a bit of a weird case and might be due to invalid swagger in the test
        // server.  how can you have an optional parameter that's also a constant?
        if (param.required !== true && param.schema.type !== SchemaType.Constant && !(param.schema.type === SchemaType.SealedChoice && (<SealedChoiceSchema>param.schema).choices.length === 1)) {
          optionalParams.push(param);
        }
      }
      if (optionalParams.length > 0) {
        // create a type named <OperationGroup><Operation>Options
        const name = `${group.language.go!.name}${op.language.go!.name}Options`;
        op.requests![0].language.go!.optionalParam = {
          name: name,
          description: `${name} contains the optional parameters for the ${group.language.go!.name}.${op.language.go!.name} method.`,
          params: optionalParams
        };
      }
      details.protocolNaming = new protocolMethods(details.name);
      // TODO check if we still need to fix up response type name and description
      const firstResp = op.responses![0];
      const name = `${opGroupName}${op.language.go!.name}Response`;
      firstResp.language.go!.name = name;
      firstResp.language.go!.description = `${name} contains the response from method ${group.language.go!.name}.${op.language.go!.name}.`;
      for (const resp of values(op.responses)) {
        // add a field to headers to include a Go compliant name for when it needs to be used as a field in a type
        if (resp.protocol.http!.headers) {
          for (const header of values(resp.protocol.http!.headers)) {
            const head = <LanguageHeader>header;
            head.name = getEscapedReservedName(removePrefix(capitalizeAcronyms(pascalCase(head.header)), 'XMS'), 'Header');
          }
        }
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

function removePrefix(name: string, prefix: string): string {
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