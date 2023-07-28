/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { CodeModel as M4CodeModel, ObjectSchema } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { Session } from '@autorest/extension-base';
import { GoCodeModel, CodeModelType, ConstantType, Info, isClientSideDefault, ModelField, ModelType, Options, ParameterGroup, PolymorphicType, ResponseEnvelope, StructField, StructType, InterfaceType } from '../../gocodemodel/gocodemodel';
import { adaptClients } from './clients';
import { adaptConstantType, adaptInterfaceType, adaptModel, adaptModelField } from './types';
import { aggregateProperties } from '../helpers';

export async function m4ToGoCodeModel(session: Session<M4CodeModel>): Promise<GoCodeModel> {
  const info = new Info(session.model.info.title);
  const options = new Options(await session.getValue('header-text', 'MISSING LICENSE HEADER'), await session.getValue('generate-fakes', false), await session.getValue('inject-spans', false));

  let type: CodeModelType = 'data-plane';
  if (await session.getValue('azure-arm', false)) {
    type = 'azure-arm';
  }
  
  const codeModel = new GoCodeModel(info, type, session.model.language.go!.packageName, options);
  if (session.model.language.go!.host) {
    codeModel.host = session.model.language.go!.host;
  }
  if (session.model.language.go!.generateTimeRFC1123Helper) {
    codeModel.marshallingRequirements.generateTimeRFC1123Helper = true;
  }
  if (session.model.language.go!.generateTimeRFC3339Helper) {
    codeModel.marshallingRequirements.generateTimeRFC3339Helper = true;
  }
  if (session.model.language.go!.generateUnixTimeHelper) {
    codeModel.marshallingRequirements.generateUnixTimeHelper = true;
  }
  if (session.model.language.go!.generateDateHelper) {
    codeModel.marshallingRequirements.generateDateHelper = true;
  }
  if (session.model.language.go!.needsXMLDictionaryUnmarshalling) {
    codeModel.marshallingRequirements.generateXMLDictionaryUnmarshallingHelper = true;
  }
  if (session.model.language.go!.moduleVersion !== '') {
    codeModel.options.moduleVersion = session.model.language.go!.moduleVersion;
  }
  if (session.model.language.go!.module !== 'none') {
    codeModel.options.module = session.model.language.go!.module;
  } else if (session.model.language.go!.containingModule !== 'none') {
    codeModel.options.containingModule = session.model.language.go!.containingModule;
  }
  codeModel.constants = adaptConstantTypes(session);
  codeModel.interfaceTypes = adaptInterfaceTypes(session);
  codeModel.models = adaptModels(session);
  codeModel.clients = adaptClients(session, codeModel);

  const paramGroups = new Map<string, ParameterGroup>();

  for (const client of values(codeModel.clients)) {
    for (const method of client.methods) {
      if (!codeModel.responseEnvelopes) {
        codeModel.responseEnvelopes = new Array<ResponseEnvelope>();
      }
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
    codeModel.paramGroups = new Array<StructType>();
    for (const groupName of paramGroups.keys()) {
      const paramGroup = paramGroups.get(groupName);
      codeModel.paramGroups.push(adaptParameterGroup(paramGroup!));
    }
  }

  return codeModel;
}

function adaptConstantTypes(session: Session<M4CodeModel>): Array<ConstantType> | undefined {
  // group all enum categories into a single array so they can be sorted
  const constTypes = new Array<ConstantType>();
  for (const choice of values(session.model.schemas.choices)) {
    if (choice.language.go!.omitType) {
      continue;
    }
    const constType = adaptConstantType(choice);
    constTypes.push(constType);
  }
  for (const choice of values(session.model.schemas.sealedChoices)) {
    if (choice.language.go!.omitType || choice.choices.length === 1) {
      continue;
    }
    const constType = adaptConstantType(choice);
    constTypes.push(constType);
  }

  if (constTypes.length === 0) {
    return undefined;
  }
  return constTypes;
}

function adaptParameterGroup(paramGroup: ParameterGroup): StructType {
  const structType = new StructType(paramGroup.groupName);
  structType.description = paramGroup.description;
  if (paramGroup.params) {
    structType.fields = new Array<StructField>();
    for (const param of values(paramGroup.params)) {
      if (param.paramType === 'literal') {
        continue;
      }
      let byValue = param.paramType === 'required' || (param.location === 'client' && isClientSideDefault(param.paramType));
      // if the param isn't required, check if it should be passed by value or not.
      // optional params that are implicitly nil-able shouldn't be pointer-to-type.
      if (!byValue) {
        byValue = param.byValue;
      }
      const field = new StructField(param.paramName, param.type, byValue);
      field.description = param.description;
      structType.fields.push(field);
    }
  }
  return structType;
}

interface InterfaceTypeObjectSchema {
  iface: InterfaceType;
  obj: ObjectSchema;
}

function adaptInterfaceTypes(session: Session<M4CodeModel>): Array<InterfaceType> | undefined {
  if (!session.model.language.go!.discriminators) {
    return undefined;
  }

  const ifaces = new Array<InterfaceType>();
  const ifaceObjs = new Array<InterfaceTypeObjectSchema>();
  const discriminators = <Array<ObjectSchema>>session.model.language.go!.discriminators;

  // discriminators contains all of the root discriminated types but *not* any sub-roots (e.g. Salmon).
  for (const discriminator of values(discriminators)) {
    if (discriminator.language.go!.omitType || discriminator.extensions?.['x-ms-external']) {
      continue;
    }
    // we must adapt all InterfaceTypes first. this is because ModelTypes/PolymorphicTypes can
    // contain references to InterfaceTypes and/or cyclic references
    recursiveAdaptInterfaceType(discriminator, ifaces, ifaceObjs);
  }

  if (ifaces.length === 0) {
    return undefined;
  }
  
  // now that the InterfaceTypes have been created, we can populate the rootType and possibleTypes
  for (const ifaceObj of values(ifaceObjs)) {
    ifaceObj.iface.rootType = <PolymorphicType>adaptModel(ifaceObj.obj);
    ifaceObj.iface.possibleTypes = new Array<PolymorphicType>();
    for (const disc of values(ifaceObj.obj.discriminator!.all)) {
      const possibleType = adaptModel(<ObjectSchema>disc);
      ifaceObj.iface.possibleTypes.push(<PolymorphicType>possibleType);
    }
  }

  return ifaces;
}

function recursiveAdaptInterfaceType(obj: ObjectSchema, ifaces: Array<InterfaceType>, ifaceObjs: Array<InterfaceTypeObjectSchema>, parent?: InterfaceType) {
  const iface = adaptInterfaceType(obj, parent);
  ifaces.push(iface);
  ifaceObjs.push({iface, obj});

  for (const val of values(obj.discriminator!.immediate)) {
    const asObj = <ObjectSchema>val;
    if (asObj.discriminator) {
      recursiveAdaptInterfaceType(asObj, ifaces, ifaceObjs, iface);
    }
  }
}

interface ModelTypeObjectSchema {
  type: ModelType | PolymorphicType;
  obj: ObjectSchema;
}

function adaptModels(session: Session<M4CodeModel>): Array<ModelType | PolymorphicType> {
  const modelTypes = new Array<ModelType | PolymorphicType>();
  const modelObjs = new Array<ModelTypeObjectSchema>();
  for (const obj of values(session.model.schemas.objects)) {
    if (obj.language.go!.omitType || obj.extensions?.['x-ms-external']) {
      continue;
    }
    // we must adapt all model types first. this is because models can contain cyclic references
    const modelType = adaptModel(obj);
    modelTypes.push(modelType);
    modelObjs.push({type: modelType, obj: obj});
  }

  for (const modelObj of values(modelObjs)) {
    modelObj.type.fields = new Array<ModelField>();
    const props = aggregateProperties(modelObj.obj);
    for (const prop of values(props)) {
      const field = adaptModelField(prop, modelObj.obj);
      modelObj.type.fields.push(field);
    }
  }

  return modelTypes;
}
