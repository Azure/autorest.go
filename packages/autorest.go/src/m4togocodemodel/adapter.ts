/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

/* eslint-disable @typescript-eslint/no-unsafe-argument */

import * as m4 from '@autorest/codemodel';
import { serialize } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import { AutorestExtensionHost, startSession } from '@autorest/extension-base';
import * as go from '../../../codemodel.go/src/index.js';
import { adaptClients } from './clients.js';
import { adaptConstantType, adaptInterfaceType, adaptModel, adaptModelField } from './types.js';
import { aggregateProperties } from '../transform/helpers.js';
import { fileURLToPath } from 'url';

// converts an M4 code model into a GoCodeModel
export async function m4ToGoCodeModel(host: AutorestExtensionHost) {
  const debug = await host.getValue('debug') || false;

  try {
    const session = await startSession<m4.CodeModel>(host, m4.codeModelSchema);

    const info = new go.Info(session.model.info.title);
    const options = new go.Options(
      await session.getValue('header-text', 'MISSING LICENSE HEADER'), 
      await session.getValue('generate-fakes', session.model.language.go!.azureARM), 
      await session.getValue('inject-spans', session.model.language.go!.azureARM),
      await session.getValue('disallow-unknown-fields', false),
      await session.getValue('generate-sdk-example', false));
    const azcoreVersion = await session.getValue('azcore-version', '');
    if (azcoreVersion !== '') {
      options.azcoreVersion = azcoreVersion;
    }

    let type: go.CodeModelType = 'data-plane';
    if (session.model.language.go!.azureARM) {
      type = 'azure-arm';
    }

    const codeModel = new go.CodeModel(info, type, session.model.language.go!.packageName, options);
    if (session.model.language.go!.module) {
      codeModel.options.module = <string>session.model.language.go!.module;
    } else if (session.model.language.go!.containingModule !== '') {
      codeModel.options.containingModule = <string>session.model.language.go!.containingModule;
    }

    adaptConstantTypes(session.model, codeModel);
    adaptInterfaceTypes(session.model, codeModel);
    adaptModels(session.model, codeModel);
    adaptClients(session.model, codeModel);

    const paramGroups = new Map<string, go.ParameterGroup>();

    for (const client of values(codeModel.clients)) {
      for (const method of client.methods) {
        codeModel.responseEnvelopes.push(method.responseEnvelope);
        for (const param of values(method.parameters)) {
          if (param.group) {
            if (!paramGroups.has(param.group.groupName)) {
              paramGroups.set(param.group.groupName, param.group);
            }
          }
        }
        if (!paramGroups.has(method.optionalParamsGroup.groupName)) {
          // the optional params group wasn't present, that means that it's empty.
          paramGroups.set(method.optionalParamsGroup.groupName, method.optionalParamsGroup);
        }
      }
    }

    if (paramGroups.size > 0) {
      // adapt all of the parameter groups
      for (const groupName of paramGroups.keys()) {
        const paramGroup = paramGroups.get(groupName);
        codeModel.paramGroups.push(adaptParameterGroup(paramGroup!));
      }
    }

    codeModel.sortContent();

    // output the model to the pipeline
    host.writeFile({
      filename: 'go-code-model.yaml',
      content: serialize(codeModel),
      artifactType: 'go-code-model'
    });
  } catch (E) {
    if (debug) {
      console.error(`${fileURLToPath(import.meta.url)} - FAILURE  ${JSON.stringify(E)} ${(<Error>E).stack}`);
    }
    throw E;
  }
}

function adaptConstantTypes(m4CodeModel: m4.CodeModel, goCodeModel: go.CodeModel) {
  // group all enum categories into a single array so they can be sorted
  for (const choice of values(m4CodeModel.schemas.choices)) {
    if (choice.language.go!.omitType) {
      continue;
    }
    const constType = adaptConstantType(choice);
    goCodeModel.constants.push(constType);
  }
  for (const choice of values(m4CodeModel.schemas.sealedChoices)) {
    if (choice.language.go!.omitType || choice.choices.length === 1) {
      continue;
    }
    const constType = adaptConstantType(choice);
    goCodeModel.constants.push(constType);
  }
}

function adaptParameterGroup(paramGroup: go.ParameterGroup): go.Struct {
  const structType = new go.Struct(paramGroup.groupName);
  structType.docs = paramGroup.docs;
  if (paramGroup.params.length > 0) {
    for (const param of values(paramGroup.params)) {
      if (param.style === 'literal') {
        continue;
      }
      let byValue = param.style === 'required' || (param.location === 'client' && go.isClientSideDefault(param.style));
      // if the param isn't required, check if it should be passed by value or not.
      // optional params that are implicitly nil-able shouldn't be pointer-to-type.
      if (!byValue) {
        byValue = param.byValue;
      }
      const field = new go.StructField(param.name, param.type, byValue);
      field.docs = param.docs;
      structType.fields.push(field);
    }
  }
  return structType;
}

interface InterfaceTypeObjectSchema {
  iface: go.Interface;
  obj: m4.ObjectSchema;
}

function adaptInterfaceTypes(m4CodeModel: m4.CodeModel, goCodeModel: go.CodeModel) {
  if (!m4CodeModel.language.go!.discriminators) {
    return;
  }

  const ifaceObjs = new Array<InterfaceTypeObjectSchema>();
  const discriminators = <Array<m4.ObjectSchema>>m4CodeModel.language.go!.discriminators;

  // discriminators contains all of the root discriminated types but *not* any sub-roots (e.g. Salmon).
  for (const discriminator of values(discriminators)) {
    if (discriminator.language.go!.omitType || discriminator.extensions?.['x-ms-external']) {
      continue;
    }
    // we must adapt all InterfaceTypes first. this is because ModelTypes/PolymorphicTypes can
    // contain references to InterfaceTypes and/or cyclic references
    recursiveAdaptInterfaceType(discriminator, goCodeModel.interfaces, ifaceObjs);
  }

  // now that the InterfaceTypes have been created, we can populate the rootType and possibleTypes
  for (const ifaceObj of values(ifaceObjs)) {
    ifaceObj.iface.rootType = <go.PolymorphicModel>adaptModel(ifaceObj.obj);
    ifaceObj.iface.possibleTypes = new Array<go.PolymorphicModel>();
    for (const disc of values(ifaceObj.obj.discriminator!.all)) {
      const possibleType = adaptModel(<m4.ObjectSchema>disc);
      ifaceObj.iface.possibleTypes.push(<go.PolymorphicModel>possibleType);
    }
  }
}

function recursiveAdaptInterfaceType(obj: m4.ObjectSchema, ifaces: Array<go.Interface>, ifaceObjs: Array<InterfaceTypeObjectSchema>, parent?: go.Interface) {
  const iface = adaptInterfaceType(obj, parent);
  if (ifaces.includes(iface)) {
    return;
  }
  ifaces.push(iface);
  ifaceObjs.push({iface, obj});

  for (const val of values(obj.discriminator!.immediate)) {
    const asObj = <m4.ObjectSchema>val;
    if (asObj.discriminator) {
      recursiveAdaptInterfaceType(asObj, ifaces, ifaceObjs, iface);
    }
  }
}

interface ModelTypeObjectSchema {
  type: go.Model | go.PolymorphicModel;
  obj: m4.ObjectSchema;
}

function adaptModels(m4CodeModel: m4.CodeModel, goCodeModel: go.CodeModel) {
  const modelObjs = new Array<ModelTypeObjectSchema>();
  for (const obj of values(m4CodeModel.schemas.objects)) {
    if (obj.language.go!.omitType || obj.extensions?.['x-ms-external']) {
      continue;
    }
    // we must adapt all model types first. this is because models can contain cyclic references
    const modelType = adaptModel(obj);
    goCodeModel.models.push(modelType);
    modelObjs.push({type: modelType, obj: obj});
  }

  for (const modelObj of values(modelObjs)) {
    const props = aggregateProperties(modelObj.obj);
    for (const prop of values(props)) {
      const field = adaptModelField(prop, modelObj.obj);
      modelObj.type.fields.push(field);
    }
  }
}
