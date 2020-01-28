/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { serialize, pascalCase } from '@azure-tools/codegen';
import { Host, startSession, Session } from '@azure-tools/autorest-extension-base';
import { codeModelSchema, CodeModel, Language } from '@azure-tools/codemodel';
import { visitor, clone, values } from '@azure-tools/linq';

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
    details.name = pascalCase(details.name);
    details.name = capitalizeAcronyms(details.name)
    for (const prop of values(obj.properties)) {
      const details = <Language>prop.language.go;
      details.name = pascalCase(details.name);
      details.name = capitalizeAcronyms(details.name)
    }
  }

  // pascal-case and capitzalize acronym operation groups and their operations
  for (const group of values(model.operationGroups)) {
    const details = <Language>group.language.go;
    details.name = pascalCase(details.name);
    details.name = capitalizeAcronyms(details.name)
    for (const op of values(group.operations)) {
      const details = <Language>op.language.go;
      details.name = pascalCase(details.name);
      details.name = capitalizeAcronyms(details.name)
    }
  }

  // fix up enum type and value names and capitzalize acronyms
  for (const enm of values(session.model.schemas.choices)) {
    for (const choice of values(enm.choices)) {
      const details = <Language>choice.language.go;
      details.name = `${enm.language.go?.name}${pascalCase(details.name.toLowerCase())}`;
      details.name = `${enm.language.go?.name}${capitalizeAcronyms(details.name)}`;
    }
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

const acronyms = [
  'Acl',
  'Api',
  'Ascii',
  'Cpu',
  'Css',
  'Dns',
  'Eof',
  'Guid',
  'Html',
  'Http',
  'Https',
  'Id',
  'Ip',
  'Json',
  'Lhs',
  'Qps',
  'Ram',
  'Rfc', // TODO check
  'Rhs',
  'Rpc',
  'Sla',
  'Smtp',
  'Sql',
  'Ssh',
  'Tcp',
  'Tls',
  'Ttl',
  'Udp',
  'Ui',
  'Uid',
  'Uuid',
  'Uri',
  'Url',
  'Utf8',
  'Vm',
  'Xml',
  'Xsrf',
  'Xss'
]

// make sure that common acronyms are capitalized
// NOTE: this function does not perform a case insensitive check considering scenarios where this would cause problems
// for example 'curl' would end up as 'cURL' if we did case insensitive checks
function capitalizeAcronyms(name: string): string {
  for (const word of acronyms) {
    name = name.replace(word, word.toUpperCase())
  }
  return name
}
