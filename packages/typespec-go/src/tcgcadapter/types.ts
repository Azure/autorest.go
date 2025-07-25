/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { uncapitalize } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import * as tcgc from '@azure-tools/typespec-client-generator-core';
import * as tsp from '@typespec/compiler';
import * as http from '@typespec/http';
import * as go from '../../../codemodel.go/src/index.js';
import * as naming from '../../../naming.go/src/naming.js';
import { AdapterError } from '../tcgcadapter/errors.js';

// used to convert SDK types to Go code model types
export class typeAdapter {
  // type adaptation might require enabling marshaling helpers
  readonly codeModel: go.CodeModel;

  // cache of previously created types/constant values
  private types: Map<string, go.WireType>;
  private constValues: Map<string, go.ConstantValue>;
  
  constructor(codeModel: go.CodeModel) {
    this.codeModel = codeModel;
    this.types = new Map<string, go.WireType>();
    this.constValues = new Map<string, go.ConstantValue>();
  }

  // converts all model/enum SDK types to Go code model types
  adaptTypes(sdkContext: tcgc.SdkContext) {
    for (const enumType of sdkContext.sdkPackage.enums) {
      if (enumType.usage === tcgc.UsageFlags.ApiVersionEnum) {
        // we have a pipeline policy for controlling the api-version
        continue;
      } else if ((enumType.usage & tcgc.UsageFlags.Input) === 0 && (enumType.usage & tcgc.UsageFlags.Output) === 0) {
        // skip types without input and output usage
        continue;
      }
      const constType = this.getConstantType(enumType);
      this.codeModel.constants.push(constType);
    }

    // we must adapt all interface/model types first. this is because models can contain cyclic references
    const modelTypes = new Array<ModelTypeSdkModelType>();
    const ifaceTypes = new Array<InterfaceTypeSdkModelType>();
    for (const modelType of sdkContext.sdkPackage.models) {
      if (modelType.name.length === 0) {
        // tcgc creates some unnamed models for spread params.
        // we don't use these so just skip them.
        continue;
      } else if (modelType.access === 'internal' && <tcgc.UsageFlags>(modelType.usage & tcgc.UsageFlags.Spread) === tcgc.UsageFlags.Spread) {
        // we don't use the internal models for spread params
        continue;
      } else if ((modelType.usage & tcgc.UsageFlags.Input) === 0 && (modelType.usage & tcgc.UsageFlags.Output) === 0) {
        // skip types without input and output usage
        continue;
      }

      if (modelType.discriminatedSubtypes) {
        // this is a root discriminated type
        const iface = this.getInterfaceType(modelType);
        this.codeModel.interfaces.push(iface);
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
      ifaceType.go.rootType = <go.PolymorphicModel>this.getModel(ifaceType.tcgc);
      for (const subType of values(ifaceType.tcgc.discriminatedSubtypes)) {
        const possibleType = <go.PolymorphicModel>this.getModel(subType);
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
        if (prop.kind === 'query') {
          // skip query params for now, wait for confirming visibility behavior
          continue;
        }
        const field = this.getModelField(prop, modelType.tcgc);
        modelType.go.fields.push(field);
      }
      if (content.addlProps) {
        const annotations = new go.ModelFieldAnnotations(false, false, true, false);
        const addlPropsType = new go.Map(this.getWireType(content.addlProps, false, false), isTypePassedByValue(content.addlProps));
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
      if (client.children && client.children.length > 0) {
        for (const child of client.children) {
          recursiveWalkClients(child);
        }
      }
      for (const sdkMethod of client.methods) {
        if (sdkMethod.kind !== 'paging') {
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
  getWireType(type: tcgc.SdkType, elementTypeByValue: boolean, substituteDiscriminator: boolean): go.WireType {
    switch (type.kind) {
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
      case 'unknown':
      case 'safeint':
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
        arrayType = new go.Slice(this.getWireType(elementType, elementTypeByValue, substituteDiscriminator), myElementTypeByValue);
        this.types.set(keyName, arrayType);
        return arrayType;
      }
      case 'endpoint': {
        const stringKey = 'string';
        let stringType = this.types.get(stringKey);
        if (stringType) {
          return stringType;
        }
        stringType = new go.String();
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
        mapType = new go.Map(this.getWireType(type.valueType, elementTypeByValue, substituteDiscriminator), valueTypeByValue);
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
            throw new AdapterError('UnsupportedTsp', `unsupported duration wire format kind ${type.wireType.kind}`, type.wireType.__raw?.node ?? tsp.NoTarget);
        }
      }
      case 'model':
        if (type.discriminatedSubtypes && substituteDiscriminator) {
          return this.getInterfaceType(type);
        }
        return this.getModel(type);
      case 'nullable':
        return this.getWireType(type.type, elementTypeByValue, substituteDiscriminator);
      default:
        throw new AdapterError('UnsupportedTsp', `unsupported type kind ${type.kind}`, type.__raw?.node ?? tsp.NoTarget);
    }
  }

  private getTimeType(encode: tsp.DateTimeKnownEncoding, utc: boolean): go.Time {
    const encoding = getDateTimeEncoding(encode);
    let datetime = this.types.get(encoding);
    if (datetime) {
      return <go.Time>datetime;
    }
    datetime = new go.Time(encoding, utc);
    this.types.set(encoding, datetime);
    return datetime;
  }

  // returns the Go code model type for an io.ReadSeekCloser
  getReadSeekCloser(sliceOf: boolean): go.WireType {
    let keyName = 'io-readseekcloser';
    if (sliceOf) {
      keyName = 'sliceof-' + keyName;
    }
    let rsc = this.types.get(keyName);
    if (!rsc) {
      rsc = new go.QualifiedType('ReadSeekCloser', 'io');
      if (sliceOf) {
        rsc = new go.Slice(rsc, true);
      }
      this.types.set(keyName, rsc);
    }
    return rsc;
  }

  // returns the Go code model type for streaming.MultipartContent
  getMultipartContent(sliceOf: boolean): go.WireType {
    let keyName = 'streaming-multipartcontent';
    if (sliceOf) {
      keyName = 'sliceof-' + keyName;
    }
    let rsc = this.types.get(keyName);
    if (!rsc) {
      rsc = new go.QualifiedType('MultipartContent', 'github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
      if (sliceOf) {
        rsc = new go.Slice(rsc, true);
      }
      this.types.set(keyName, rsc);
    }
    return rsc;
  }

  private getBuiltInType(type: tcgc.SdkBuiltInType): go.WireType {
    switch (type.kind) {
      case 'unknown': {
        if (this.codeModel.options.rawJSONAsBytes) {
          const anyRawJSONKey = 'any-raw-json';
          let anyRawJSON = this.types.get(anyRawJSONKey);
          if (anyRawJSON) {
            return anyRawJSON;
          }
          anyRawJSON = new go.RawJSON();
          this.types.set(anyRawJSONKey, anyRawJSON);
          return anyRawJSON;
        }
        let anyType = this.types.get('any');
        if (anyType) {
          return anyType;
        }
        anyType = new go.Any();
        this.types.set('any', anyType);
        return anyType;
      }
      case 'boolean': {
        const boolKey = 'boolean';
        let primitiveBool = this.types.get(boolKey);
        if (primitiveBool) {
          return primitiveBool;
        }
        primitiveBool = new go.Scalar('bool', type.encode === 'string');
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
        date = new go.Time('dateType', false);
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
        decimalType = new go.Scalar(decimalKey, type.encode === 'string');
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
        float32 = new go.Scalar(float32Key, type.encode === 'string');
        this.types.set(float32Key, float32);
        return float32;
      }
      case 'float64': {
        const float64Key = 'float64';
        let float64 = this.types.get(float64Key);
        if (float64) {
          return float64;
        }
        float64 = new go.Scalar(float64Key, type.encode === 'string');
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
        const keyName = type.encode === 'string' ? `${type.kind}-string` : type.kind;
        let intType = this.types.get(keyName);
        if (intType) {
          return intType;
        }
        intType = new go.Scalar(type.kind, type.encode === 'string');
        this.types.set(keyName, intType);
        return intType;
      }
      case 'safeint': {
        const safeintkey = type.encode === 'string' ? 'int64-string' : 'int64';
        let int64 = this.types.get(safeintkey);
        if (int64) {
          return int64;
        }
        int64 = new go.Scalar('int64', type.encode === 'string');
        this.types.set(safeintkey, int64);
        return int64;
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
        stringType = new go.String();
        this.types.set(stringKey, stringType);
        return stringType;
      }
      case 'plainTime': {
        const encoding = 'timeRFC3339';
        let time = this.types.get(encoding);
        if (time) {
          return time;
        }
        time = new go.Time(encoding, false);
        this.types.set(encoding, time);
        return time;
      
      }
      default:
        throw new AdapterError('UnsupportedTsp', `unsupported type kind ${type.kind}`, type.__raw?.node ?? tsp.NoTarget);
    }
  }

  // converts an SdkEnumType to a go.ConstantType
  private getConstantType(enumType: tcgc.SdkEnumType): go.Constant {
    let constTypeName = naming.ensureNameCase(enumType.name);
    if (enumType.access === 'internal') {
      constTypeName = naming.getEscapedReservedName(uncapitalize(constTypeName), 'Type');
    }
    let constType = this.types.get(constTypeName);
    if (constType) {
      return <go.Constant>constType;
    }
    const accessPrefix = enumType.access === 'internal' ? 'p' : 'P';
    constType = new go.Constant(constTypeName, getPrimitiveType(enumType.valueType), `${accessPrefix}ossible${constTypeName}Values`);
    constType.values = this.getConstantValues(constType, enumType.values);
    constType.docs.summary = enumType.summary;
    constType.docs.description = enumType.doc;
    this.types.set(constTypeName, constType);
    return constType;
  }

  private getInterfaceType(model: tcgc.SdkModelType, parent?: go.Interface): go.Interface {
    if (model.name.length === 0) {
      throw new AdapterError('InternalError', 'unnamed model', tsp.NoTarget);
    }
    if (!model.discriminatedSubtypes) {
      throw new AdapterError('InternalError', `type ${model.name} isn't a discriminator root`, model.__raw?.node ?? tsp.NoTarget);
    }
    let ifaceName = naming.createPolymorphicInterfaceName(naming.ensureNameCase(model.name));
    if (model.access === 'internal') {
      ifaceName = uncapitalize(ifaceName);
    }
    let iface = this.types.get(ifaceName);
    if (iface) {
      return <go.Interface>iface;
    }
    // find the discriminator field
    let discriminatorField: string | undefined;
    for (const prop of model.properties) {
      if (prop.kind === 'property' && prop.discriminator) {
        // only json support discriminator type
        discriminatorField = prop.serializationOptions.json?.name;
        break;
      }
    }
    if (!discriminatorField) {
      throw new AdapterError('InternalError', `failed to find discriminator field for type ${model.name}`, tsp.NoTarget);
    }
    iface = new go.Interface(ifaceName, discriminatorField);
    if (parent) {
      iface.parent = parent;
    }
    this.types.set(ifaceName, iface);
    return iface;
  }

  // converts an SdkModelType to a go.ModelType or go.PolymorphicType if the model is polymorphic
  private getModel(model: tcgc.SdkModelType): go.Model | go.PolymorphicModel {
    let modelName = naming.ensureNameCase(model.name);
    if (model.access === 'internal') {
      modelName = naming.getEscapedReservedName(uncapitalize(modelName), 'Model');
    }
    let modelType = this.types.get(modelName);
    if (modelType) {
      return <go.Model | go.PolymorphicModel>modelType;
    }
  
    let usage = go.UsageFlags.None;
    if (model.usage & tsp.UsageFlags.Input) {
      usage = go.UsageFlags.Input;
    }
    if (model.usage & tsp.UsageFlags.Output) {
      usage |= go.UsageFlags.Output;
    }

    const annotations = new go.ModelAnnotations(false, <tcgc.UsageFlags>(model.usage & tcgc.UsageFlags.MultipartFormData) === tcgc.UsageFlags.MultipartFormData);
    if (model.discriminatedSubtypes || model.discriminatorValue) {
      let iface: go.Interface | undefined;
      let discriminatorLiteral: go.Literal | undefined;

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
          throw new AdapterError('InternalError', `failed to find discriminator interface name for type ${model.name}`, tsp.NoTarget);
        }

        // find the discriminator property and create the discriminator literal based on it
        for (const prop of model.properties) {
          if (prop.kind === 'property' && prop.discriminator) {
            discriminatorLiteral = this.getDiscriminatorLiteral(prop);
            break;
          }
        }
      }

      modelType = new go.PolymorphicModel(modelName, iface, annotations, usage);
      modelType.discriminatorValue = discriminatorLiteral;
    } else {
      modelType = new go.Model(modelName, annotations, usage);
      // polymorphic types don't have XMLInfo
      modelType.xml = adaptXMLInfo(model.decorators);
    }

    modelType.docs.summary = model.summary;
    modelType.docs.description = model.doc;
    if (modelType.docs.summary) {
      if (!modelType.docs.summary.startsWith(modelName)) {
        modelType.docs.summary = `${modelName} - ${modelType.docs.summary}`;
      }
    } else if (modelType.docs.description) {
      if (!modelType.docs.description.startsWith(modelName)) {
        modelType.docs.description = `${modelName} - ${modelType.docs.description}`;
      }
    }

    this.types.set(modelName, modelType);
    return modelType;
  }

  private getDiscriminatorLiteral(sdkProp: tcgc.SdkBodyModelPropertyType): go.Literal {
    switch (sdkProp.type.kind) {
      case 'constant':
      case 'enumvalue':
        return this.getLiteralValue(sdkProp.type);
      default:
        throw new AdapterError('UnsupportedTsp', `unsupported type kind ${sdkProp.type.kind} for discriminator property ${sdkProp.name}`, sdkProp.__raw?.node ?? tsp.NoTarget);
    }
  }

  private getModelField(prop: tcgc.SdkModelPropertyType, modelType: tcgc.SdkModelType): go.ModelField {
    if (prop.kind !== 'path' && prop.kind !== 'property') {
      throw new AdapterError('UnsupportedTsp', `unsupported kind ${prop.kind} for property ${prop.name} in model ${modelType.name}`, prop.__raw?.node ?? tsp.NoTarget);
    }
    const annotations = new go.ModelFieldAnnotations(prop.optional === false, false, false, false);
    // for multipart/form data containing models, default to fields not being pointer-to-type as we
    // don't have to deal with JSON patch shenanigans. only the optional fields will be pointer-to-type.
    const isMultipartFormData = <tcgc.UsageFlags>(modelType.usage & tcgc.UsageFlags.MultipartFormData) === tcgc.UsageFlags.MultipartFormData;
    let fieldByValue = isMultipartFormData ? true : isTypePassedByValue(prop.type);
    if (isMultipartFormData && prop.kind === 'property' && prop.optional) {
      fieldByValue = false;
    }
    let type = this.getWireType(prop.type, isMultipartFormData, true);
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
    field.docs.summary = prop.summary;
    field.docs.description = prop.doc;
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
  
    field.xml = adaptXMLInfo(prop.decorators, field);
  
    return field;
  }

  private getConstantValues(type: go.Constant, valueTypes: Array<tcgc.SdkEnumValueType>): Array<go.ConstantValue> {
    const values = new Array<go.ConstantValue>();
    for (const valueType of valueTypes) {
      let valueTypeName = `${type.name}${naming.ensureNameCase(valueType.name)}`;
      if (valueType.enumType.access === 'internal') {
        valueTypeName = naming.getEscapedReservedName(uncapitalize(valueTypeName), 'Type');
      }
      let value = this.constValues.get(valueTypeName);
      if (!value) {
        value = new go.ConstantValue(valueTypeName, type, valueType.value);
        value.docs.summary = valueType.summary;
        value.docs.description = valueType.doc;
        this.constValues.set(valueTypeName, value);
      }
      values.push(value);
    }
    return values;
  }

  private adaptBytesType(sdkType: tcgc.SdkBuiltInType): go.EncodedBytes {
    let format: go.BytesEncoding = 'Std';
    if (sdkType.encode === 'base64url') {
      format = 'URL';
    }
    const keyName = `bytes-${format}`;
    let bytesType = this.types.get(keyName);
    if (bytesType) {
      return <go.EncodedBytes>bytesType;
    }
    bytesType = new go.EncodedBytes(format);
    this.types.set(keyName, bytesType);
    return bytesType;
  }

  private getLiteralValue(constType: tcgc.SdkConstantType | tcgc.SdkEnumValueType): go.Literal {
    if (constType.kind === 'enumvalue') {
      const valueName = `${naming.ensureNameCase(constType.enumType.name)}${naming.ensureNameCase(constType.name)}`;
      const keyName = `literal-${valueName}`;
      let literalConst = this.types.get(keyName);
      if (literalConst) {
        return <go.Literal>literalConst;
      }
      const constValue = this.constValues.get(valueName);
      if (!constValue) {
        throw new AdapterError('InternalError', `failed to find const value for ${constType.name} in enum ${constType.enumType.name}`, constType.__raw?.node ?? tsp.NoTarget);
      }
      literalConst = new go.Literal(this.getConstantType(constType.enumType), constValue);
      this.types.set(keyName, literalConst);
      return literalConst;
    }

    switch (constType.valueType.kind) {
      case 'boolean': {
        const keyName = `literal-boolean-${constType.value}`;
        let literalBool = this.types.get(keyName);
        if (literalBool) {
          return <go.Literal>literalBool;
        }
        literalBool = new go.Literal(new go.Scalar('bool', false), constType.value);
        this.types.set(keyName, literalBool);
        return literalBool;
      }
      case 'bytes': {
        const keyName = `literal-bytes-${constType.value}`;
        let literalByteArray = this.types.get(keyName);
        if (literalByteArray) {
          return <go.Literal>literalByteArray;
        }
        literalByteArray = new go.Literal(this.adaptBytesType(constType.valueType), constType.value);
        this.types.set(keyName, literalByteArray);
        return literalByteArray;
      }
      case 'decimal':
      case 'decimal128': {
        const keyName = `literal-${constType.valueType.kind}-${constType.value}`;
        let literalDecimal = this.types.get(keyName);
        if (literalDecimal) {
          return <go.Literal>literalDecimal;
        }
        literalDecimal = new go.Literal(new go.Scalar('float64', false), constType.value);
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
          return <go.Literal>literalInt;
        }
        literalInt = new go.Literal(new go.Scalar(constType.valueType.kind, false), constType.value);
        this.types.set(keyName, literalInt);
        return literalInt;
      }
      case 'float32':
      case 'float64': {
        const keyName = `literal-${constType.valueType.kind}-${constType.value}`;
        let literalFloat = this.types.get(keyName);
        if (literalFloat) {
          return <go.Literal>literalFloat;
        }
        literalFloat = new go.Literal(new go.Scalar(constType.valueType.kind, false), constType.value);
        this.types.set(keyName, literalFloat);
        return literalFloat;
      }
      case 'string':
      case 'url': {
        const keyName = `literal-string-${constType.value}`;
        let literalString = this.types.get(keyName);
        if (literalString) {
          return <go.Literal>literalString;
        }
        literalString = new go.Literal(new go.String(), constType.value);
        this.types.set(keyName, literalString);
        return literalString;
      }
      default:
        throw new AdapterError('UnsupportedTsp', `unsupported kind ${constType.valueType.kind} for LiteralValue`, constType.valueType.__raw?.node ?? tsp.NoTarget);
    }

    // TODO: tcgc doesn't support duration as a literal value
  }
}

function getPrimitiveType(type: tcgc.SdkBuiltInType): 'bool' | 'float32' | 'float64' | 'int32' | 'int64' | 'string' {
  switch (type.kind) {
    case 'boolean':
      return 'bool';
    case 'float32':
    case 'float64':
    case 'int32':
    case 'int64':
    case 'string':
      return type.kind;
    default:
      throw new AdapterError('UnsupportedTsp', `unhandled tcgc.SdkBuiltInKinds: ${type.kind}`, type.__raw?.node ?? tsp.NoTarget);
  }
}

function getDateTimeEncoding(encoding: tsp.DateTimeKnownEncoding): go.TimeFormat {
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
        throw new AdapterError('UnsupportedTsp', `unsupported time encoding ${obj.encode}`, obj.__raw?.node ?? tsp.NoTarget);
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
  return type.kind === 'unknown' || type.kind === 'array' ||
  type.kind === 'bytes' || type.kind === 'dict' ||
    (type.kind === 'model' && !!type.discriminatedSubtypes);
}

interface ModelTypeSdkModelType {
  go: go.Model | go.PolymorphicModel;
  tcgc: tcgc.SdkModelType;
}

interface InterfaceTypeSdkModelType {
  go: go.Interface;
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

// called for models and model fields. for the former, the field param will be undefined
export function adaptXMLInfo(decorators: Array<tcgc.DecoratorInfo>, field?: go.ModelField): go.XMLInfo | undefined {
  // if there are no decorators and this isn't a slice
  // type in a model field then do nothing
  if (decorators.length === 0 && (!field || (field.type.kind !== 'slice'))) {
    return undefined;
  }

  const xmlInfo = new go.XMLInfo();
  if (field && field.type.kind === 'slice') {
    // for tsp, arrays are wrapped by default
    xmlInfo.wraps = go.getTypeDeclaration(field.type.elementType);
  }

  const handleName = (decorator: tcgc.DecoratorInfo): void => {
    if (field) {
      xmlInfo.name = <string>decorator.arguments['name'];
    } else {
      // when applied to a model, it means the model's XML element
      // node has a different name than the model.
      xmlInfo.wrapper = <string>decorator.arguments['name'];
    }
  };

  for (const decorator of decorators) {
    switch (decorator.name) {
      case 'TypeSpec.@encodedName':
        if (decorator.arguments['mimeType'] === 'application/xml') {
          handleName(decorator);
        }
        break;
      case 'TypeSpec.Xml.@attribute':
        xmlInfo.attribute = true;
        break;
      case 'TypeSpec.Xml.@name':
        handleName(decorator);
        break;
      case 'TypeSpec.Xml.@unwrapped':
        // unwrapped can only be applied to fields
        if (field) {
          switch (field.type.kind) {
            case 'slice':
              // unwrapped slice. default to using the serialized name
              xmlInfo.wraps = undefined;
              xmlInfo.name = field.serializedName;
              break;
            case 'string':
              // an unwrapped string means it's text
              xmlInfo.text = true;  
              break;
          }
        }
        break;
    }
  }

  return xmlInfo;
}
