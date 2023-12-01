/*  ---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *  --------------------------------------------------------------------------------------------  */

import { Session } from '@autorest/extension-base';
import { CodeModel, HttpHeader, HttpMethod, Language } from '@autorest/codemodel';
import { visitor, clone, values } from '@azure-tools/linq';
import { ensureNameCase, getEscapedReservedName, packageNameFromOutputFolder, trimPackagePrefix, uncapitalize } from '../../../naming.go/src/naming.js';
import { aggregateParameters, hasAdditionalProperties } from './helpers.js';

const requestMethodSuffix = 'CreateRequest';
const responseMethodSuffix = 'HandleResponse';

// contains extended naming information for operations
export interface OperationNaming extends Language {
  protocolNaming: protocolNaming
}

interface protocolNaming {
  internalMethod: string;
  requestMethod: string;
  responseMethod: string;
}

export class protocolMethods implements protocolNaming {
  readonly internalMethod: string;
  readonly requestMethod: string;
  readonly responseMethod: string;

  constructor(name: string) {
    // uncapitalizing runs the risk of reserved name collision, e.g. Import -> import
    this.internalMethod = getEscapedReservedName(uncapitalize(name), 'Operation');
    this.requestMethod = ensureNameCase(`${name}${requestMethodSuffix}`, true);
    this.responseMethod = ensureNameCase(`${name}${responseMethodSuffix}`, true);
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
  model.language.go!.packageName = packageNameFromOutputFolder(outputFolder);

  // default to the package name
  let stutteringPrefix = <string>model.language.go!.packageName;
  // if there's a well-known prefix, remove it
  if (stutteringPrefix.startsWith('arm')) {
    stutteringPrefix = stutteringPrefix.substring(3);
  } else if (stutteringPrefix.startsWith('az')) {
    stutteringPrefix = stutteringPrefix.substring(2);
  }
  // use the user-specified value if available
  stutteringPrefix = await session.getValue<string>('stutter', stutteringPrefix);
  stutteringPrefix = stutteringPrefix.toUpperCase();

  const specType = await session.getValue('openapi-type');
  model.language.go!.openApiType = specType;
  const azureARM = await session.getValue('azure-arm', false);
  model.language.go!.azureARM = azureARM;
  const headAsBoolean = await session.getValue('head-as-boolean', false);
  model.language.go!.headAsBoolean = headAsBoolean;
  const groupParameters = await session.getValue('group-parameters', true);
  model.language.go!.groupParameters = groupParameters;
  const honorBodyPlacement = await session.getValue('honor-body-placement', false);
  const rawJSONAsBytes = await session.getValue('rawjson-as-bytes', false);
  model.language.go!.rawJSONAsBytes = rawJSONAsBytes;
  const sliceElementsByValue = await session.getValue('slice-elements-byval', false);
  model.language.go!.sliceElementsByValue = sliceElementsByValue;
  const moduleVersion = await session.getValue('module-version', '');
  if (moduleVersion !== '' && !moduleVersion.match(/^(\d+\.\d+\.\d+(?:-beta\.\d+)?)?$/)) {
    throw new Error(`module version ${moduleVersion} must in the format major.minor.patch[-beta.N]`);
  }
  model.language.go!.moduleVersion = moduleVersion;

  let module = await session.getValue('module', 'none');
  if (module !== 'none') {
    if (module.match(/\/v\d+$/)) {
      throw new Error('module name must not contain major version suffix');
    }
    if (moduleVersion !== '') {
      // if the modules major version is greater than one, add a major version suffix to the module name
      const majorVersion = moduleVersion.substring(0, moduleVersion.indexOf('.'));
      if (Number(majorVersion) > 1) {
        module += '/v' + majorVersion;
      }
    }
  }
  model.language.go!.module = module;

  const containingModule = await session.getValue('containing-module', 'none');
  if (containingModule !== 'none' && module !== 'none') {
    throw new Error('--module and --containing-module are mutually exclusive');
  }
  model.language.go!.containingModule = containingModule;

  // fix up type names
  const structNames = new Set<string>();
  for (const obj of values(model.schemas.objects)) {
    obj.language.go!.name = ensureNameCase(obj.language.go!.name);
    structNames.add(obj.language.go!.name);
  }

  // fix stuttering type names
  const collisions = new Array<string>();
  for (const obj of values(model.schemas.objects)) {
    const details = <Language>obj.language.go;
    const originalName = details.name;
    details.name = trimPackagePrefix(stutteringPrefix, originalName);
    // if the type was renamed to remove stuttering, check if it collides with an existing type name
    if (details.name !== originalName && structNames.has(details.name)) {
      collisions.push(`type ${originalName} was renamed to ${details.name} which collides with an existing type name`);
    }
  }
  if (collisions.length > 0) {
    throw new Error(collisions.join('\n'));
  }

  // fix property names and other bits
  for (const obj of values(model.schemas.objects)) {
    const details = <Language>obj.language.go;
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
    // adding this extension to a type will skip generatings its serde (marshalling/unmarshalling) methods
    if (obj.extensions?.['x-ms-go-omit-serde-methods']) {
      obj.language.go!.omitSerDeMethods = true;
    }
  }

  // fix up operation group names
  const groupNames = new Set<string>();
  for (const group of values(model.operationGroups)) {
    if (group.language.go!.name.length > 0) {
      group.language.go!.name = ensureNameCase(group.language.go!.name);
      groupNames.add(group.language.go!.name);
    }
  }

  // fix up any missing operation group names and add client names
  const fallbackGroupName = ensureNameCase(session.model.info.title);
  const clientNames = new Set<string>();
  for (const group of values(model.operationGroups)) {
    // use the swagger title as the default name for operation groups that don't specify a group name
    if (group.language.go!.name.length === 0) {
      if (groupNames.has(fallbackGroupName)) {
        throw new Error(`the fallback operation group name ${fallbackGroupName} collides with an existing group name`);
      }
      group.language.go!.name = fallbackGroupName;
    }
    group.language.go!.clientName = group.language.go!.name;
    // don't generate a name like FooClientClient
    if (!group.language.go!.clientName.endsWith('Client')) {
      group.language.go!.clientName = `${group.language.go!.name}Client`;
    }
    clientNames.add(group.language.go!.clientName);
  }

  // fix up stuttering client names and operation names
  for (const group of values(model.operationGroups)) {
    const groupDetails = <Language>group.language.go;
    const originalName = groupDetails.clientName;
    groupDetails.clientName = trimPackagePrefix(stutteringPrefix, originalName);
    // if the client was renamed to remove stuttering, check if it collides with an existing client
    if (groupDetails.clientName !== originalName && clientNames.has(groupDetails.clientName)) {
      throw new Error(`client ${originalName} was renamed to ${groupDetails.clientName} which collides with an existing client name`);
    }
    groupDetails.clientCtorName = `New${groupDetails.clientName}`;
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
        if (!honorBodyPlacement) {
          const opMethod = op.requests![0].protocol.http!.method;
          if (param.protocol.http?.in === 'body' && (opMethod === HttpMethod.Patch || opMethod === HttpMethod.Put)) {
            // we enforce PATCH/PUT body parameters to be required.  do this before fixing up the parameter name
            param.required = true;
          }
        }
        const inParamGroup = (param.extensions?.['x-ms-parameter-grouping'] && groupParameters) || param.required !== true;
        const paramDetails = <Language>param.language.go;
        // if this is part of a param group struct then don't apply param naming rules to it
        paramDetails.name = ensureNameCase(paramDetails.name, !inParamGroup);
        // fix up any param group names
        if (param.extensions?.['x-ms-parameter-grouping'] && groupParameters) {
          if (param.extensions['x-ms-parameter-grouping'].name) {
            param.extensions['x-ms-parameter-grouping'].name = ensureNameCase(param.extensions['x-ms-parameter-grouping'].name);
          } else if (param.extensions['x-ms-parameter-grouping'].postfix) {
            param.extensions['x-ms-parameter-grouping'].postfix = ensureNameCase(param.extensions['x-ms-parameter-grouping'].postfix);
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
          op.language.go!.paging.nextLinkName = ensureNameCase(op.language.go!.paging.nextLinkName);
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

function createPolymorphicInterfaceName(base: string): string {
  return base + 'Classification';
}
