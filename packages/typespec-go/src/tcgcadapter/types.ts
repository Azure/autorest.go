/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as naming from '../../../naming.go/src/naming.js';
import * as go from '../../../codemodel.go/src/index.js';
import * as tcgc from '@azure-tools/typespec-client-generator-core';
import * as tsp from '@typespec/compiler';
import * as http from '@typespec/http';
import { values } from '@azure-tools/linq';
import { uncapitalize } from '@azure-tools/codegen';

// used to convert SDK types to Go code model types
export class typeAdapter {
  // type adaptation might require enabling marshaling helpers
  readonly codeModel: go.CodeModel;

  // cache of previously created types/constant values
  private types: Map<string, go.PossibleType>;
  private constValues: Map<string, go.ConstantValue>;

  // contains the names of referenced types
  private unreferencedEnums: Set<string>;
  private unreferencedModels: Set<string>;
  
  constructor(codeModel: go.CodeModel) {
    this.codeModel = codeModel;
    this.types = new Map<string, go.PossibleType>();
    this.constValues = new Map<string, go.ConstantValue>();
    this.unreferencedEnums = new Set<string>();
    this.unreferencedModels = new Set<string>();
  }

  // converts all model/enum SDK types to Go code model types
  adaptTypes(sdkContext: tcgc.SdkContext, removeUnreferencedTypes: boolean) {
    if (removeUnreferencedTypes) {
      // this is a superset of flagUnreferencedBaseModels
      this.flagUnreferencedTypes(sdkContext);
    } else {
      this.flagUnreferencedBaseModels(sdkContext);
    }

    for (const enumType of sdkContext.sdkPackage.enums) {
      if (enumType.usage === tcgc.UsageFlags.ApiVersionEnum) {
        // we have a pipeline policy for controlling the api-version
        continue;
      } else if (this.unreferencedEnums.has(enumType.name)) {
        // skip unreferenced type
        continue;
      }
      const constType = this.getConstantType(enumType);
      this.codeModel.constants.push(constType);
    }

    // we must adapt all interface/model types first. this is because models can contain cyclic references
    const modelTypes = new Array<ModelTypeSdkModelType>();
    const ifaceTypes = new Array<InterfaceTypeSdkModelType>();
    for (const modelType of sdkContext.sdkPackage.models) {
      if (this.unreferencedModels.has(modelType.name)) {
        // skip unreferenced type
        continue;
      }
      if (modelType.discriminatedSubtypes) {
        // this is a root discriminated type
        const iface = this.getInterfaceType(modelType);
        this.codeModel.interfaceTypes.push(iface);
        ifaceTypes.push({go: iface, tcgc: modelType});
      }
      // TODO: what's the equivalent of x-ms-external?
      const model = this.getModel(modelType);
      modelTypes.push({go: model, tcgc: modelType});
    }
  
    // add the synthesized models from TCGC for paged results
    const pagedResponses = this.getPagedResponses(sdkContext);
    for (const pagedResponse of pagedResponses) {
      // tsp allows custom paged responses, so we must check both the synthesized list and the models list
      if (values(modelTypes).any(each => { return each.tcgc.name === pagedResponse.name; })) {
        continue;
      }
      const model = this.getModel(pagedResponse);
      modelTypes.push({go: model, tcgc: pagedResponse});
    }


    // now that the interface/model types have been generated, we can populate the rootType and possibleTypes
    for (const ifaceType of ifaceTypes) {
      ifaceType.go.rootType = <go.PolymorphicType>this.getModel(ifaceType.tcgc);
      for (const subType of values(ifaceType.tcgc.discriminatedSubtypes)) {
        const possibleType = <go.PolymorphicType>this.getModel(subType);
        ifaceType.go.possibleTypes.push(possibleType);
      }
    }

    // now adapt model fields
    for (const modelType of modelTypes) {
      const content = aggregateProperties(modelType.tcgc);
      for (const prop of values(content.props)) {
        if (prop.kind === 'header') {
          // the common case here is the @header decorator specifying
          // the content-type for the model. we can just skip it.
          // TODO: follow up with tcgc to see if we can remove the entry.
          continue;
        }
        const field = this.getModelField(prop, modelType.tcgc);
        modelType.go.fields.push(field);
      }
      if (content.addlProps) {
        const annotations = new go.ModelFieldAnnotations(false, false, true, false);
        const addlPropsType = new go.MapType(this.getPossibleType(content.addlProps, false, false), isTypePassedByValue(content.addlProps));
        const addlProps = new go.ModelField('AdditionalProperties', addlPropsType, true, '', annotations);
        modelType.go.fields.push(addlProps);
      }
      this.codeModel.models.push(modelType.go);
    }
  }

  // returns the synthesized paged response types
  private getPagedResponses(sdkContext: tcgc.SdkContext): Array<tcgc.SdkModelType> {
    const pagedResponses = new Array<tcgc.SdkModelType>();
    const recursiveWalkClients = function(client: tcgc.SdkClientType<tcgc.SdkHttpOperation>): void {
      for (const sdkMethod of client.methods) {
        if (sdkMethod.kind === 'clientaccessor') {
          recursiveWalkClients(sdkMethod.response);
          continue;
        } else if (sdkMethod.kind !== 'paging') {
          continue;
        }

        for (const httpResp of sdkMethod.operation.responses.values()) {
          if (!httpResp.type || httpResp.type.kind !== 'model') {
            continue;
          }

          if (!values(pagedResponses).any(each => { return each.name === (<tcgc.SdkModelType>httpResp.type).name; })) {
            pagedResponses.push(httpResp.type);
          }
        }
      }
    };

    for (const sdkClient of sdkContext.sdkPackage.clients) {
      recursiveWalkClients(sdkClient);
    }
    return pagedResponses;
  }

  // returns the Go code model type for the specified SDK type.
  // the operation is idempotent, so getting the same type multiple times
  // returns the same instance of the converted type.
  getPossibleType(type: tcgc.SdkType, elementTypeByValue: boolean, substituteDiscriminator: boolean): go.PossibleType {
    switch (type.kind) {
      case 'any':
      case 'boolean':
      case 'bytes':
      case 'decimal':
      case 'decimal128':
      case 'float':
      case 'float32':
      case 'float64':
      case 'int8':
      case 'int16':
      case 'int32':
      case 'int64':
      case 'uint8':
      case 'uint16':
      case 'uint32':
      case 'uint64':
      case 'plainDate':
      case 'plainTime':
      case 'string':
      case 'url':
        return this.getBuiltInType(type);
      case 'array': {
        let elementType = type.valueType;
        let nullable = false;
        if (elementType.kind === 'nullable') {
          // unwrap the nullable type
          elementType = elementType.type;
          nullable = true;
        }
        // prefer elementTypeByValue. if false, then if the array elements have been explicitly marked as nullable then prefer that, else fall back to our usual algorithm
        const myElementTypeByValue = elementTypeByValue ? true : nullable ? false : this.codeModel.options.sliceElementsByval || isTypePassedByValue(elementType);
        const keyName = recursiveKeyName(`array-${myElementTypeByValue}`, elementType, substituteDiscriminator);
        let arrayType = this.types.get(keyName);
        if (arrayType) {
          return arrayType;
        }
        arrayType = new go.SliceType(this.getPossibleType(elementType, elementTypeByValue, substituteDiscriminator), myElementTypeByValue);
        this.types.set(keyName, arrayType);
        return arrayType;
      }
      case 'endpoint': {
        const stringKey = 'string';
        let stringType = this.types.get(stringKey);
        if (stringType) {
          return stringType;
        }
        stringType = new go.PrimitiveType('string');
        this.types.set(stringKey, stringType);
        return stringType;
      }
      case 'enum':
        return this.getConstantType(type);
      case 'constant':
      case 'enumvalue':
        return this.getLiteralValue(type);
      case 'offsetDateTime':
        return this.getTimeType(type.encode, false);
      case 'utcDateTime':
        return this.getTimeType(type.encode, true);
      case 'dict': {
        const valueTypeByValue = isTypePassedByValue(type.valueType);
        const keyName = recursiveKeyName(`dict-${valueTypeByValue}`, type.valueType, substituteDiscriminator);
        let mapType = this.types.get(keyName);
        if (mapType) {
          return mapType;
        }
        mapType = new go.MapType(this.getPossibleType(type.valueType, elementTypeByValue, substituteDiscriminator), valueTypeByValue);
        this.types.set(keyName, mapType);
        return mapType;
      }
      case 'duration': {
        switch (type.wireType.kind) {
          case 'float':
          case 'float32':
          case 'float64':
          case 'int32':
          case 'int64':
          case 'string':
            return this.getBuiltInType(type.wireType);
          default:
            throw new Error(`unhandled duration wireType.kind ${type.wireType.kind}`);
        }
      }
      case 'model':
        if (type.discriminatedSubtypes && substituteDiscriminator) {
          return this.getInterfaceType(type);
        }
        return this.getModel(type);
      case 'nullable':
        return this.getPossibleType(type.type, elementTypeByValue, substituteDiscriminator);
      default:
        throw new Error(`unhandled property kind ${type.kind}`);
    }
  }

  private getTimeType(encode: tsp.DateTimeKnownEncoding, utc: boolean): go.TimeType {
    const encoding = getDateTimeEncoding(encode);
    let datetime = this.types.get(encoding);
    if (datetime) {
      return <go.TimeType>datetime;
    }
    datetime = new go.TimeType(encoding, utc);
    this.types.set(encoding, datetime);
    return <go.TimeType>datetime;
  }

  // returns the Go code model type for an io.ReadSeekCloser
  getReadSeekCloser(sliceOf: boolean): go.PossibleType {
    let keyName = 'io-readseekcloser';
    if (sliceOf) {
      keyName = 'sliceof-' + keyName;
    }
    let rsc = this.types.get(keyName);
    if (!rsc) {
      rsc = new go.QualifiedType('ReadSeekCloser', 'io');
      if (sliceOf) {
        rsc = new go.SliceType(rsc, true);
      }
      this.types.set(keyName, rsc);
    }
    return rsc;
  }

  // returns the Go code model type for streaming.MultipartContent
  getMultipartContent(sliceOf: boolean): go.PossibleType {
    let keyName = 'streaming-multipartcontent';
    if (sliceOf) {
      keyName = 'sliceof-' + keyName;
    }
    let rsc = this.types.get(keyName);
    if (!rsc) {
      rsc = new go.QualifiedType('MultipartContent', 'github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
      if (sliceOf) {
        rsc = new go.SliceType(rsc, true);
      }
      this.types.set(keyName, rsc);
    }
    return rsc;
  }

  private getBuiltInType(type: tcgc.SdkBuiltInType): go.PossibleType {
    switch (type.kind) {
      case 'any': {
        if (this.codeModel.options.rawJSONAsBytes) {
          const anyRawJSONKey = 'any-raw-json';
          let anyRawJSON = this.types.get(anyRawJSONKey);
          if (anyRawJSON) {
            return anyRawJSON;
          }
          anyRawJSON = new go.SliceType(new go.PrimitiveType('byte'), true);
          anyRawJSON.rawJSONAsBytes = true;
          this.types.set(anyRawJSONKey, anyRawJSON);
          return anyRawJSON;
        }
        let anyType = this.types.get('any');
        if (anyType) {
          return anyType;
        }
        anyType = new go.PrimitiveType('any');
        this.types.set('any', anyType);
        return anyType;
      }
      case 'boolean': {
        const boolKey = 'boolean';
        let primitiveBool = this.types.get(boolKey);
        if (primitiveBool) {
          return primitiveBool;
        }
        primitiveBool = new go.PrimitiveType('bool');
        this.types.set(boolKey, primitiveBool);
        return primitiveBool;
      }
      case 'bytes':
        return this.adaptBytesType(type);
      case 'plainDate': {
        const dateKey = 'dateType';
        let date = this.types.get(dateKey);
        if (date) {
          return date;
        }
        date = new go.TimeType('dateType', false);
        this.types.set(dateKey, date);
        return date;
      }
      case 'decimal':
      case 'decimal128': {
        const decimalKey = 'float64';
        let decimalType = this.types.get(decimalKey);
        if (decimalType) {
          return decimalType;
        }
        decimalType = new go.PrimitiveType(decimalKey);
        this.types.set(decimalKey, decimalType);
        return decimalType;
      }
      case 'float': // C# and Java define float as 32 bits so we're following suit
      case 'float32': {
        const float32Key = 'float32';
        let float32 = this.types.get(float32Key);
        if (float32) {
          return float32;
        }
        float32 = new go.PrimitiveType(float32Key);
        this.types.set(float32Key, float32);
        return float32;
      }
      case 'float64': {
        const float64Key = 'float64';
        let float64 = this.types.get(float64Key);
        if (float64) {
          return float64;
        }
        float64 = new go.PrimitiveType(float64Key);
        this.types.set(float64Key, float64);
        return float64;
      }
      case 'int8':
      case 'int16':
      case 'int32':
      case 'int64':
      case 'uint8':
      case 'uint16':
      case 'uint32':
      case 'uint64': {
        const keyName = type.kind;
        let intType = this.types.get(keyName);
        if (intType) {
          return intType;
        }
        intType = new go.PrimitiveType(type.kind);
        this.types.set(keyName, intType);
        return intType;
      }
      case 'string':
      case 'url': {
        if (type.crossLanguageDefinitionId === 'Azure.Core.eTag') {
          const etagKey = 'etag';
          let etag = this.types.get(etagKey);
          if (etag) {
            return etag;
          }
          etag = new go.QualifiedType('ETag', 'github.com/Azure/azure-sdk-for-go/sdk/azcore');
          this.types.set(etagKey, etag);
          return etag;
        }

        const stringKey = 'string';
        let stringType = this.types.get(stringKey);
        if (stringType) {
          return stringType;
        }
        stringType = new go.PrimitiveType('string');
        this.types.set(stringKey, stringType);
        return stringType;
      }
      case 'plainTime': {
        const encoding = 'timeRFC3339';
        let time = this.types.get(encoding);
        if (time) {
          return time;
        }
        time = new go.TimeType(encoding, false);
        this.types.set(encoding, time);
        return time;
      
      }
      default:
        throw new Error(`unhandled property kind ${type.kind}`);
    }
  }

  // converts an SdkEnumType to a go.ConstantType
  private getConstantType(enumType: tcgc.SdkEnumType): go.ConstantType {
    let constTypeName = naming.ensureNameCase(enumType.name);
    if (enumType.access === 'internal') {
      constTypeName = naming.getEscapedReservedName(uncapitalize(constTypeName), 'Type');
    }
    let constType = this.types.get(constTypeName);
    if (constType) {
      return <go.ConstantType>constType;
    }
    constType = new go.ConstantType(constTypeName, getPrimitiveType(enumType.valueType.kind), `Possible${constTypeName}Values`);
    constType.values = this.getConstantValues(constType, enumType.values);
    constType.description = enumType.description;
    this.types.set(constTypeName, constType);
    return constType;
  }

  private getInterfaceType(model: tcgc.SdkModelType, parent?: go.InterfaceType): go.InterfaceType {
    if (model.name.length === 0) {
      throw new Error('unnamed model');
    }
    if (!model.discriminatedSubtypes) {
      throw new Error(`type ${model.name} isn't a discriminator root`);
    }
    let ifaceName = naming.createPolymorphicInterfaceName(naming.ensureNameCase(model.name));
    if (model.access === 'internal') {
      ifaceName = uncapitalize(ifaceName);
    }
    let iface = this.types.get(ifaceName);
    if (iface) {
      return <go.InterfaceType>iface;
    }
    // find the discriminator field
    let discriminatorField: string | undefined;
    for (const prop of model.properties) {
      if (prop.kind === 'property' && prop.discriminator) {
        discriminatorField = prop.serializedName;
        break;
      }
    }
    if (!discriminatorField) {
      throw new Error(`failed to find discriminator field for type ${model.name}`);
    }
    iface = new go.InterfaceType(ifaceName, discriminatorField);
    if (parent) {
      iface.parent = parent;
    }
    this.types.set(ifaceName, iface);
    return iface;
  }

  // converts an SdkModelType to a go.ModelType or go.PolymorphicType if the model is polymorphic
  private getModel(model: tcgc.SdkModelType): go.ModelType | go.PolymorphicType {
    let modelName = model.name;
    if (modelName.length === 0) {
      throw new Error('unnamed model');
    }
    modelName = naming.ensureNameCase(modelName);
    if (model.access === 'internal') {
      modelName = naming.getEscapedReservedName(uncapitalize(modelName), 'Model');
    }
    let modelType = this.types.get(modelName);
    if (modelType) {
      return <go.ModelType | go.PolymorphicType>modelType;
    }
  
    let usage = go.UsageFlags.None;
    if (model.usage & tsp.UsageFlags.Input) {
      usage = go.UsageFlags.Input;
    }
    if (model.usage & tsp.UsageFlags.Output) {
      usage |= go.UsageFlags.Output;
    }

    const annotations = new go.ModelAnnotations(false, model.isFormDataType);
    if (model.discriminatedSubtypes || model.discriminatorValue) {
      let iface: go.InterfaceType | undefined;
      let discriminatorLiteral: go.LiteralValue | undefined;

      if (model.discriminatedSubtypes) {
        // root type, we can get the InterfaceType directly from it
        iface = this.getInterfaceType(model);
      } else {
        // walk the parents until we find the first root type
        let parent = model.baseModel;
        while (parent) {
          if (parent.discriminatedSubtypes) {
            iface = this.getInterfaceType(parent);
            break;
          }
          parent = parent.baseModel;
        }
        if (!iface) {
          throw new Error(`failed to find discriminator interface name for type ${model.name}`);
        }

        // find the discriminator property and create the discriminator literal based on it
        for (const prop of model.properties) {
          if (prop.kind === 'property' && prop.discriminator) {
            discriminatorLiteral = this.getDiscriminatorLiteral(prop);
            break;
          }
        }
      }

      modelType = new go.PolymorphicType(modelName, iface, annotations, usage);
      (<go.PolymorphicType>modelType).discriminatorValue = discriminatorLiteral;
    } else {
      modelType = new go.ModelType(modelName, annotations, usage);
      // polymorphic types don't have XMLInfo
      // TODO: XMLInfo
    }
    if (model.description) {
      modelType.description = model.description;
      if (!modelType.description.startsWith(modelName)) {
        modelType.description = `${modelName} - ${modelType.description}`;
      }
    }
    this.types.set(modelName, modelType);
    return modelType;
  }

  private getDiscriminatorLiteral(sdkProp: tcgc.SdkBodyModelPropertyType): go.LiteralValue {
    switch (sdkProp.type.kind) {
      case 'constant':
      case 'enumvalue':
        return this.getLiteralValue(sdkProp.type);
      default:
        throw new Error(`unhandled kind ${sdkProp.type.kind} for discriminator property ${sdkProp.name}`);
    }
  }

  private getModelField(prop: tcgc.SdkModelPropertyType, modelType: tcgc.SdkModelType): go.ModelField {
    if (prop.kind !== 'path' && prop.kind !== 'property') {
      throw new Error(`unexpected kind ${prop.kind} for property ${prop.name} in model ${modelType.name}`);
    }
    const annotations = new go.ModelFieldAnnotations(prop.optional === false, false, false, false);
    // for multipart/form data containing models, default to fields not being pointer-to-type as we
    // don't have to deal with JSON patch shenanigans. only the optional fields will be pointer-to-type.
    const isMultipartFormData = (modelType.usage & tcgc.UsageFlags.MultipartFormData) === tcgc.UsageFlags.MultipartFormData;
    let fieldByValue = isMultipartFormData ? true : isTypePassedByValue(prop.type);
    if (isMultipartFormData && prop.kind === 'property' && prop.optional) {
      fieldByValue = false;
    }
    let type = this.getPossibleType(prop.type, isMultipartFormData, true);
    if (prop.kind === 'property') {
      if (prop.isMultipartFileInput) {
        type = this.getMultipartContent(prop.type.kind === 'array');
      }
      if (prop.visibility) {
        // the field is read-only IFF the only visibility attribute present is Read.
        // a field can have Read & Create set which means it's required on input and
        // returned on output.
        if (prop.visibility.length === 1 && prop.visibility[0] === http.Visibility.Read) {
          annotations.readOnly = true;
        }
      }
    }
    const field = new go.ModelField(naming.capitalize(naming.ensureNameCase(prop.name)), type, fieldByValue, prop.serializedName, annotations);
    field.description = prop.description;
    if (prop.kind === 'path') {
      // for ARM resources, a property of kind path is usually the model
      // key and will be exposed as a discrete method parameter. this also
      // means that the value is read-only.
      annotations.readOnly = true;
    } else if (prop.discriminator && modelType.discriminatorValue) {
      // the presence of modelType.discriminatorValue tells us that this
      // property is on a model that's not the root discriminator
      annotations.isDiscriminator = true;
      field.defaultValue = this.getDiscriminatorLiteral(prop);
    }
  
    // TODO: XMLInfo
    //field.xml = adaptXMLInfo(prop.schema);
  
    return field;
  }

  private getConstantValues(type: go.ConstantType, valueTypes: Array<tcgc.SdkEnumValueType>): Array<go.ConstantValue> {
    const values = new Array<go.ConstantValue>();
    for (const valueType of valueTypes) {
      let valueTypeName = `${type.name}${naming.ensureNameCase(valueType.name)}`;
      if (valueType.enumType.access === 'internal') {
        valueTypeName = naming.getEscapedReservedName(uncapitalize(valueTypeName), 'Type');
      }
      let value = this.constValues.get(valueTypeName);
      if (!value) {
        value = new go.ConstantValue(valueTypeName, type, valueType.value);
        value.description = valueType.description;
        this.constValues.set(valueTypeName, value);
      }
      values.push(value);
    }
    return values;
  }

  private adaptBytesType(sdkType: tcgc.SdkBuiltInType): go.BytesType {
    let format: go.BytesEncoding = 'Std';
    if (sdkType.encode === 'base64url') {
      format = 'URL';
    }
    const keyName = `bytes-${format}`;
    let bytesType = this.types.get(keyName);
    if (bytesType) {
      return <go.BytesType>bytesType;
    }
    bytesType = new go.BytesType(format);
    this.types.set(keyName, bytesType);
    return bytesType;
  }

  private getLiteralValue(constType: tcgc.SdkConstantType | tcgc.SdkEnumValueType): go.LiteralValue {
    if (constType.kind === 'enumvalue') {
      const valueName = `${naming.ensureNameCase(constType.enumType.name)}${naming.ensureNameCase(constType.name)}`;
      const keyName = `literal-${valueName}`;
      let literalConst = this.types.get(keyName);
      if (literalConst) {
        return <go.LiteralValue>literalConst;
      }
      const constValue = this.constValues.get(valueName);
      if (!constValue) {
        throw new Error(`failed to find const value for ${constType.name} in enum ${constType.enumType.name}`);
      }
      literalConst = new go.LiteralValue(this.getConstantType(constType.enumType), constValue);
      this.types.set(keyName, literalConst);
      return literalConst;
    }

    switch (constType.valueType.kind) {
      case 'boolean': {
        const keyName = `literal-boolean-${constType.value}`;
        let literalBool = this.types.get(keyName);
        if (literalBool) {
          return <go.LiteralValue>literalBool;
        }
        literalBool = new go.LiteralValue(new go.PrimitiveType('bool'), constType.value);
        this.types.set(keyName, literalBool);
        return literalBool;
      }
      case 'bytes': {
        const keyName = `literal-bytes-${constType.value}`;
        let literalByteArray = this.types.get(keyName);
        if (literalByteArray) {
          return <go.LiteralValue>literalByteArray;
        }
        literalByteArray = new go.LiteralValue(this.adaptBytesType(constType.valueType), constType.value);
        this.types.set(keyName, literalByteArray);
        return literalByteArray;
      }
      case 'decimal':
      case 'decimal128': {
        const keyName = `literal-${constType.valueType.kind}-${constType.value}`;
        let literalDecimal = this.types.get(keyName);
        if (literalDecimal) {
          return <go.LiteralValue>literalDecimal;
        }
        literalDecimal = new go.LiteralValue(new go.PrimitiveType('float64'), constType.value);
        this.types.set(keyName, literalDecimal);
        return literalDecimal;
      }
      /*case 'date':
      case 'datetime': {
        // TODO: tcgc doesn't expose the encoding for date/datetime constant types
        const encoding = getDateTimeEncoding(constType.encode);
        const keyName = `literal-${encoding}-${constType.value}`;
        let literalTime = this.types.get(keyName);
        if (literalTime) {
          return <go.LiteralValue>literalTime;
        }
        literalTime = new go.LiteralValue(new go.TimeType(encoding), constType.value);
        this.types.set(keyName, literalTime);
        return literalTime;
      }*/
      case 'int8':
      case 'int16':
      case 'int32':
      case 'int64':
      case 'uint8':
      case 'uint16':
      case 'uint32':
      case 'uint64': {
        const keyName = `literal-${constType.valueType.kind}-${constType.value}`;
        let literalInt = this.types.get(keyName);
        if (literalInt) {
          return <go.LiteralValue>literalInt;
        }
        literalInt = new go.LiteralValue(new go.PrimitiveType(constType.valueType.kind), constType.value);
        this.types.set(keyName, literalInt);
        return literalInt;
      }
      case 'float32':
      case 'float64': {
        const keyName = `literal-${constType.valueType.kind}-${constType.value}`;
        let literalFloat = this.types.get(keyName);
        if (literalFloat) {
          return <go.LiteralValue>literalFloat;
        }
        literalFloat = new go.LiteralValue(new go.PrimitiveType(constType.valueType.kind), constType.value);
        this.types.set(keyName, literalFloat);
        return literalFloat;
      }
      case 'string':
      case 'url': {
        const keyName = `literal-string-${constType.value}`;
        let literalString = this.types.get(keyName);
        if (literalString) {
          return <go.LiteralValue>literalString;
        }
        literalString = new go.LiteralValue(new go.PrimitiveType('string'), constType.value);
        this.types.set(keyName, literalString);
        return literalString;
      }
      default:
        throw new Error(`unsupported kind ${constType.valueType.kind} for LiteralValue`);
    }
  
    // TODO: tcgc doesn't support duration as a literal value
  }

  // updates this.unreferencedEnums and this.unreferencedModels
  private flagUnreferencedTypes(sdkContext: tcgc.SdkContext): void {
    const referencedEnums = new Set<string>();
    const referencedModels = new Set<string>();

    const recursiveAddReferencedType = function(type: tcgc.SdkType): void {
      switch (type.kind) {
        case 'array':
        case 'dict':
          recursiveAddReferencedType(type.valueType);
          break;
        case 'enum':
          if (!referencedEnums.has(type.name)) {
            referencedEnums.add(type.name);
          }
          break;
        case 'enumvalue':
          if (!referencedEnums.has(type.enumType.name)) {
            referencedEnums.add(type.enumType.name);
          }
          break;
        case 'model':
          if (!referencedModels.has(type.name)) {
            referencedModels.add(type.name);
            const aggregateProps = aggregateProperties(type);
            for (const prop of aggregateProps.props) {
              recursiveAddReferencedType(prop.type);
            }
            if (aggregateProps.addlProps) {
              recursiveAddReferencedType(aggregateProps.addlProps);
            }
            if (type.discriminatedSubtypes) {
              for (const subType of values(type.discriminatedSubtypes)) {
                recursiveAddReferencedType(subType);
              }
            }
          }
          break;
        case 'nullable':
          return recursiveAddReferencedType(type.type);
      }
    };

    const recursiveWalkClients = function(client: tcgc.SdkClientType<tcgc.SdkHttpOperation>): void {
      for (const method of client.methods) {
        if (method.kind === 'clientaccessor') {
          recursiveWalkClients(method.response);
          continue;
        }

        for (const param of method.parameters) {
          recursiveAddReferencedType(param.type);
        }

        if (method.response.type) {
          recursiveAddReferencedType(method.response.type);
        }
      }
    };

    // traverse all client initialization params and methods to find the set of referenced enums and models
    for (const client of sdkContext.sdkPackage.clients) {
      for (const param of client.initialization.properties) {
        if (param.kind === 'endpoint') {
          let endpointType = getEndpointType(param);
          for (const templateArg of endpointType.templateArguments) {
            recursiveAddReferencedType(templateArg.type);
          }
        }
      }

      recursiveWalkClients(client);
    }

    const pagedResponses = this.getPagedResponses(sdkContext);
    for (const pagedResponse of pagedResponses) {
      recursiveAddReferencedType(pagedResponse);
    }

    // now that we have the referenced set, update the unreferenced set
    for (const sdkEnum of sdkContext.sdkPackage.enums) {
      if (!referencedEnums.has(sdkEnum.name)) {
        this.unreferencedEnums.add(sdkEnum.name);
      }
    }

    for (const model of sdkContext.sdkPackage.models) {
      if (!referencedModels.has(model.name)) {
        this.unreferencedModels.add(model.name);
      }
    }
  }

  // updates this.unreferencedModels
  private flagUnreferencedBaseModels(sdkContext: tcgc.SdkContext): void {
    const baseModels = new Set<string>();
    const referencedBaseModels = new Set<string>();
    const visitedModels = new Set<string>(); // avoids infinite recursion

    const recursiveAddReferencedBaseModel = function(type: tcgc.SdkType): void {
      switch (type.kind) {
        case 'array':
        case 'dict':
          recursiveAddReferencedBaseModel(type.valueType);
          break;
        case 'model':
          if (baseModels.has(type.name)) {
            if (!referencedBaseModels.has(type.name)) {
              referencedBaseModels.add(type.name);
            }
          } else if (!visitedModels.has(type.name)) {
            visitedModels.add(type.name);
            const aggregateProps = aggregateProperties(type);
            for (const prop of aggregateProps.props) {
              recursiveAddReferencedBaseModel(prop.type);
            }
            if (aggregateProps.addlProps) {
              recursiveAddReferencedBaseModel(aggregateProps.addlProps);
            }
            if (type.discriminatedSubtypes) {
              for (const subType of values(type.discriminatedSubtypes)) {
                recursiveAddReferencedBaseModel(subType);
              }
            }
          }
          break;
        case 'nullable':
          return recursiveAddReferencedBaseModel(type.type);
      }
    };

    // collect all the base model types
    for (const model of sdkContext.sdkPackage.models) {
      let parent = model.baseModel;
      while (parent) {
        // exclude any polymorphic root type from the check
        // as we always need to include the root type even
        // if it's not referenced.
        // also exclude types that have been explicitly annotated
        // as output types.
        if (!parent.discriminatedSubtypes && ((model.usage & tcgc.UsageFlags.Output) !== tcgc.UsageFlags.Output) && !baseModels.has(parent.name)) {
          baseModels.add(parent.name);
        }
        parent = parent.baseModel;
      }
    }

    // traverse all methods to find any references to a base model type.
    // NOTE: it's possible for there to be no base types.
    if (baseModels.size > 0) {
      const recursiveWalkClients = function(client: tcgc.SdkClientType<tcgc.SdkHttpOperation>): void {
        for (const method of client.methods) {
          if (method.kind === 'clientaccessor') {
            recursiveWalkClients(method.response);
            continue;
          }
  
          for (const param of method.parameters) {
            recursiveAddReferencedBaseModel(param.type);
          }
  
          if (method.response.type) {
            recursiveAddReferencedBaseModel(method.response.type);
          }
        }
      };

      for (const client of sdkContext.sdkPackage.clients) {
        recursiveWalkClients(client);
      }

      const pagedResponses = this.getPagedResponses(sdkContext);
      for (const pagedResponse of pagedResponses) {
        recursiveAddReferencedBaseModel(pagedResponse);
      }
    }

    // now that we have the referenced set, update the unreferenced set
    for (const baseModel of baseModels) {
      if (!referencedBaseModels.has(baseModel)) {
        this.unreferencedModels.add(baseModel);
      }
    }
  }
}

export function getEndpointType(param: tcgc.SdkEndpointParameter) {
  // for multiple endpoint, we only generate the first one
  if (param.type.kind === 'endpoint') {
    return param.type;
  } else {
    return param.type.values[0];
  }
}

function getPrimitiveType(kind: tcgc.SdkBuiltInKinds): 'bool' | 'float32' | 'float64' | 'int32' | 'int64' | 'string' {
  switch (kind) {
    case 'boolean':
      return 'bool';
    case 'float32':
    case 'float64':
    case 'int32':
    case 'int64':
    case 'string':
      return kind;
    default:
      throw new Error(`unhandled tcgc.SdkBuiltInKinds: ${kind}`);
  }
}

function getDateTimeEncoding(encoding: tsp.DateTimeKnownEncoding): go.DateTimeFormat {
  switch (encoding) {
    case 'rfc3339':
      return 'dateTimeRFC3339';
    case 'rfc7231':
      return 'dateTimeRFC1123';
    case 'unixTimestamp':
      return 'timeUnix';
  }
}

function recursiveKeyName(root: string, obj: tcgc.SdkType, substituteDiscriminator: boolean): string {
  switch (obj.kind) {
    case 'array':
      return recursiveKeyName(`${root}-array`, obj.valueType, substituteDiscriminator);
    case 'enum':
      return `${root}-${obj.name}`;
    case 'enumvalue':
      return `${root}-${obj.enumType.name}-${obj.value}`;
    case 'dict':
      return recursiveKeyName(`${root}-dict`, obj.valueType, substituteDiscriminator);
    case 'plainDate':
      return `${root}-dateType`;
    case 'utcDateTime':
      return `${root}-${getDateTimeEncoding(obj.encode)}`;
    case 'duration':
      return `${root}-${obj.wireType.kind}`;
    case 'model':
      if (substituteDiscriminator) {
        return `${root}-${naming.createPolymorphicInterfaceName(obj.name)}`;
      }
      return `${root}-${obj.name}`;
    case 'nullable':
      return recursiveKeyName(root, obj.type, substituteDiscriminator);
    case 'plainTime':
      if (obj.encode !== 'rfc3339') {
        throw new Error(`unsupported time encoding ${obj.encode}`);
      }
      return `${root}-timeRFC3339`;
    default:
      return `${root}-${obj.kind}`;
  }
}

export function isTypePassedByValue(type: tcgc.SdkType): boolean {
  if (type.kind === 'nullable') {
    type = type.type;
  }
  return type.kind === 'any' || type.kind === 'array' ||
  type.kind === 'bytes' || type.kind === 'dict' ||
    (type.kind === 'model' && !!type.discriminatedSubtypes);
}

interface ModelTypeSdkModelType {
  go: go.ModelType | go.PolymorphicType;
  tcgc: tcgc.SdkModelType;
}

interface InterfaceTypeSdkModelType {
  go: go.InterfaceType;
  tcgc: tcgc.SdkModelType;
}

// aggregate the properties from the provided type and its parent types.
// this includes any inherited additional properties.
function aggregateProperties(model: tcgc.SdkModelType): {props: Array<tcgc.SdkModelPropertyType>, addlProps?: tcgc.SdkType} {
  const allProps = new Array<tcgc.SdkModelPropertyType>();
  for (const prop of model.properties) {
    allProps.push(prop);
  }

  let addlProps = model.additionalProperties;
  let parent = model.baseModel;
  while (parent) {
    for (const parentProp of parent.properties) {
      const exists = values(allProps).where(p => { return p.name === parentProp.name; }).first();
      if (exists) {
        // don't add the duplicate. the TS compiler has better enforcement than OpenAPI
        // to ensure that duplicate fields with different types aren't added.
        continue;
      }
      allProps.push(parentProp);
    }
    // if we haven't found additional properties yet and the parent has it, use it
    if (!addlProps && parent.additionalProperties) {
      addlProps = parent.additionalProperties;
    }
    parent = parent.baseModel;
  }
  return {props: allProps, addlProps: addlProps};
}
