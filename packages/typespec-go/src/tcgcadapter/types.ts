/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as naming from '../../../naming.go/src/naming.js';
import * as go from '../../../codemodel.go/src/gocodemodel.js';
import * as tcgc from '@azure-tools/typespec-client-generator-core';
import * as tsp from '@typespec/compiler';
import { values } from '@azure-tools/linq';

// used to convert SDK types to Go code model types
export class typeAdapter {
  // type adaptation might require enabling marshaling helpers
  readonly codeModel: go.CodeModel;

  // cache of previously created types/constant values
  private types: Map<string, go.PossibleType>;
  private constValues: Map<string, go.ConstantValue>;
  
  constructor(codeModel: go.CodeModel) {
    this.codeModel = codeModel;
    this.types = new Map<string, go.PossibleType>();
    this.constValues = new Map<string, go.ConstantValue>();
  }

  // converts all model/enum SDK types to Go code model types
  adaptTypes(sdkContext: tcgc.SdkContext) {
    for (const enumType of sdkContext.sdkPackage.enums) {
      const constType = this.getConstantType(enumType);
      this.codeModel.constants.push(constType);
    }
  
    // we must adapt all interface/model types first. this is because models can contain cyclic references
    const modelTypes = new Array<ModelTypeSdkModelType>();
    const ifaceTypes = new Array<InterfaceTypeSdkModelType>();
    for (const modelType of sdkContext.sdkPackage.models) {
      if (modelType.discriminatedSubtypes) {
        // this is a root discriminated type
        const iface = this.getInterfaceType(modelType);
        this.codeModel.interfaceTypes.push(iface);
        ifaceTypes.push({go: iface, tcgc: modelType});
      }
      // TODO: what's the equivalent of x-ms-external?
      const model = this.getModel(modelType);
      modelTypes.push({go: model, tcgc: modelType});
      // workaround until https://github.com/Azure/typespec-azure/issues/153 is fixed
      if (modelType.additionalProperties) {
        const leafType = recursiveGetType(modelType.additionalProperties);
        if (leafType.kind === 'model') {
          const model = this.getModel(leafType);
          if (!values(modelTypes).any(e => { return e.tcgc.name === model.name; })) {
            modelTypes.push({go: model, tcgc: leafType});
          }
        }
      }
      // end workaround
    }
  
    // add the synthesized models from TCGC for paged results
    const pagedResponses = new Array<tcgc.SdkModelType>();
    for (const sdkClient of sdkContext.sdkPackage.clients) {
      for (const sdkMethod of sdkClient.methods) {
        if (sdkMethod.kind !== 'paging') {
          continue;
        }

        for (const httpResp of Object.values(sdkMethod.operation.responses)) {
          if (!httpResp.type || httpResp.type.kind !== 'model') {
            continue;
          }

          // tsp allows custom paged responses, so we must check both the synthesized list and the models list
          if (!values(pagedResponses).any(each => { return each.name === (<tcgc.SdkModelType>httpResp.type).name; }) && !values(modelTypes).any(each => { return each.tcgc.name === (<tcgc.SdkModelType>httpResp.type).name; })) {
            pagedResponses.push(httpResp.type);
          }
        }
      }
    }
    for (const pagedResponse of pagedResponses) {
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
      const props = aggregateProperties(modelType.tcgc);
      for (const prop of values(props)) {
        if (prop.kind !== 'property') {
          throw new Error(`unexpected kind ${prop.kind} for property ${prop.nameInClient} in model ${modelType.tcgc.name}`);
        }
        const field = this.getModelField(prop, modelType.tcgc);
        modelType.go.fields.push(field);
      }
      if (modelType.tcgc.additionalProperties) {
        const annotations = new go.ModelFieldAnnotations(false, false, true, false);
        const addlPropsType = new go.MapType(this.getPossibleType(modelType.tcgc.additionalProperties, false, false), isTypePassedByValue(modelType.tcgc.additionalProperties));
        const addlProps = new go.ModelField('AdditionalProperties', addlPropsType, true, '', annotations);
        modelType.go.fields.push(addlProps);
      }
      this.codeModel.models.push(modelType.go);
    }
  }

  // returns the Go code model type for the specified SDK type.
  // the operation is idempotent, so getting the same type multiple times
  // returns the same instance of the converted type.
  getPossibleType(type: tcgc.SdkType, elementTypeByValue: boolean, substituteDiscriminator: boolean): go.PossibleType {
    switch (type.kind) {
      case 'any':
      case 'armId':
      case 'boolean':
      case 'bytes':
      case 'date':
      case 'decimal':
      case 'decimal128':
      case 'etag':
      case 'float32':
      case 'float64':
      case 'guid':
      case 'int32':
      case 'int64':
      case 'string':
      case 'time':
      case 'url':
      case 'uuid':
        return this.getBuiltInType(type);
      case 'array': {
        // prefer elementTypeByValue. if false, then if the array elements have been explicitly marked as nullable then prefer that, else fall back to our usual algorithm
        const myElementTypeByValue = elementTypeByValue ? true : type.valueType.nullable ? false : this.codeModel.options.sliceElementsByval || isTypePassedByValue(type.valueType);
        const keyName = recursiveKeyName(`array-${myElementTypeByValue}`, type.valueType, substituteDiscriminator);
        let arrayType = this.types.get(keyName);
        if (arrayType) {
          return arrayType;
        }
        arrayType = new go.SliceType(this.getPossibleType(type.valueType, elementTypeByValue, substituteDiscriminator), myElementTypeByValue);
        this.types.set(keyName, arrayType);
        return arrayType;
      }
      /*case m4.SchemaType.Binary: {
        let binaryType = types.get(m4.SchemaType.Binary);
        if (binaryType) {
          return binaryType;
        }
        binaryType = new go.StandardType('io.ReadSeekCloser', 'io');
        types.set(m4.SchemaType.Binary, binaryType);
        return binaryType;
      }*/
      /*case m4.SchemaType.Char: {
        let rune = types.get(m4.SchemaType.Char);
        if (rune) {
          return rune;
        }
        rune = new go.PrimitiveType('rune');
        types.set(m4.SchemaType.Char, rune);
        return rune;
      }*/
      case 'enum':
        return this.getConstantType(type);
      case 'constant':
      case 'enumvalue':
        return this.getLiteralValue(type);
      /*case m4.SchemaType.Credential: {
        let credType = types.get(m4.SchemaType.Credential);
        if (credType) {
          return credType;
        }
        credType = new go.PrimitiveType('string');
        types.set(m4.SchemaType.Credential, credType);
        return credType;
      }*/
      case 'datetime': {
        const encoding = getDateTimeEncoding(type.encode);
        let datetime = this.types.get(encoding);
        if (datetime) {
          return datetime;
        }
        datetime = new go.TimeType(encoding);
        this.types.set(encoding, datetime);
        switch (encoding) {
          case 'dateTimeRFC1123':
            this.codeModel.marshallingRequirements.generateDateTimeRFC1123Helper = true;
            break;
          case 'dateTimeRFC3339':
            this.codeModel.marshallingRequirements.generateDateTimeRFC3339Helper = true;
            break;
          case 'timeUnix':
            this.codeModel.marshallingRequirements.generateUnixTimeHelper = true;
            break;
          default:
            throw new Error(`unhandled datetime encoding ${encoding}`);
        }
        return datetime;
      }
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
      default:
        throw new Error(`unhandled property kind ${type.kind}`);
    }
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
      case 'date': {
        if (type.encode !== 'rfc3339') {
          throw new Error(`unsupported date encoding ${type.encode}`);
        }
        const dateKey = `date-${type.encode}`;
        let date = this.types.get(dateKey);
        if (date) {
          return date;
        }
        date = new go.TimeType('dateType');
        this.types.set(dateKey, date);
        this.codeModel.marshallingRequirements.generateDateHelper = true;
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
        this. types.set(decimalKey, decimalType);
        return decimalType;
      }
      case 'etag': {
        const etagKey = 'etag';
        let etag = this.types.get(etagKey);
        if (etag) {
          return etag;
        }
        etag = new go.QualifiedType('ETag', 'github.com/Azure/azure-sdk-for-go/sdk/azcore');
        this.types.set(etagKey, etag);
        return etag;
      }
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
        this. types.set(float64Key, float64);
        return float64;
      }
      case 'int32': {
        const int32Key = 'int32';
        let int32 = this.types.get(int32Key);
        if (int32) {
          return int32;
        }
        int32 = new go.PrimitiveType(int32Key);
        this.types.set(int32Key, int32);
        return int32;
      }
      case 'int64': {
        const int64Key = 'int64';
        let int64 = this.types.get(int64Key);
        if (int64) {
          return int64;
        }
        int64 = new go.PrimitiveType(int64Key);
        this.types.set(int64Key, int64);
        return int64;
      }
      case 'armId':
      case 'string': {
        const stringKey = 'string';
        let stringType = this.types.get(stringKey);
        if (stringType) {
          return stringType;
        }
        stringType = new go.PrimitiveType('string');
        this.types.set(stringKey, stringType);
        return stringType;
      }
      case 'time': {
        if (type.encode !== 'rfc3339') {
          throw new Error(`unsupported time encoding ${type.encode}`);
        }
        const encoding = 'timeRFC3339';
        let time = this.types.get(encoding);
        if (time) {
          return time;
        }
        time = new go.TimeType(encoding);
        this.types.set(encoding, time);
        this.codeModel.marshallingRequirements.generateTimeRFC3339Helper = true;
        return time;
      }
      case 'url': {
        const urlKey = 'url';
        let uriType = this.types.get(urlKey);
        if (uriType) {
          return uriType;
        }
        uriType = new go.PrimitiveType('string');
        this.types.set(urlKey, uriType);
        return uriType;
      }
      case 'uuid':
      case 'guid': {
        const uuidKey = 'uuid';
        let uuid = this.types.get(uuidKey);
        if (uuid) {
          return uuid;
        }
        uuid = new go.PrimitiveType('string');
        this.types.set(uuidKey, uuid);
        return uuid;
      }
      default:
        throw new Error(`unhandled property kind ${type.kind}`);
    }
  }

  // converts an SdkEnumType to a go.ConstantType
  private getConstantType(enumType: tcgc.SdkEnumType): go.ConstantType {
    const constTypeName = naming.ensureNameCase(enumType.name);
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
    if (!model.discriminatedSubtypes) {
      throw new Error(`type ${model.name} isn't a discriminator root`);
    }
    const ifaceName = naming.createPolymorphicInterfaceName(naming.ensureNameCase(model.name));
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
    const modelName = naming.ensureNameCase(model.name);
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
    // TODO: what's the extension equivalent in TS?
    const annotations = new go.ModelAnnotations(false);
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
      // TODO: hard-coded format
      modelType = new go.ModelType(modelName, 'json', annotations, usage);
      // polymorphic types don't have XMLInfo
      // TODO: XMLInfo
    }
    modelType.description = model.description;
    this.types.set(modelName, modelType);
    return modelType;
  }

  private getDiscriminatorLiteral(sdkProp: tcgc.SdkBodyModelPropertyType): go.LiteralValue {
    switch (sdkProp.type.kind) {
      case 'constant':
      case 'enumvalue':
        return this.getLiteralValue(sdkProp.type);
      default:
        throw new Error(`unhandled kind ${sdkProp.type.kind} for discriminator property ${sdkProp.nameInClient}`);
    }
  }

  private getModelField(prop: tcgc.SdkBodyModelPropertyType, modelType: tcgc.SdkModelType): go.ModelField {
    // TODO: hard-coded values
    const annotations = new go.ModelFieldAnnotations(prop.optional == false, false, false, false);
    const field = new go.ModelField(naming.capitalize(naming.ensureNameCase(prop.nameInClient)), this.getPossibleType(prop.type, false, true), isTypePassedByValue(prop.type), prop.serializedName, annotations);
    field.description = prop.description;
    // the presence of modelType.discriminatorValue tells us that this
    // property is on a model that's not the root discriminator
    if (prop.discriminator && modelType.discriminatorValue) {
      annotations.isDiscriminator = true;
      field.defaultValue = this.getDiscriminatorLiteral(prop);
    } /*else if (prop.clientDefaultValue) {
      if (!go.isLiteralValueType(field.type)) {
        throw new Error(`unsupported default value type ${go.getTypeDeclaration(field.type)} for field ${field.fieldName}`);
      }
      if (go.isConstantType(field.type)) {
        // find the corresponding ConstantValue
        const constType = types.get(field.type.name);
        if (!constType) {
          throw new Error(`didn't find ConstantType for ${field.type.name}`);
        }
        let found = false;
        for (const val of values((<go.ConstantType>constType).values)) {
          if (val.value === prop.clientDefaultValue) {
            const keyName = `literal-${val.valueName}`;
            let literalValue = types.get(keyName);
            if (!literalValue) {
              literalValue = new go.LiteralValue(field.type, val);
              types.set(keyName, literalValue);
            }
            field.defaultValue = <go.LiteralValue>literalValue;
            found = true;
            break;
          }
        }
        if (!found) {
          throw new Error(`didn't find ConstantValue for ${prop.clientDefaultValue}`);
        }
      } else {
        const keyName = `literal-${go.getTypeDeclaration(field.type)}-${prop.clientDefaultValue}`;
        let literalValue = types.get(keyName);
        if (!literalValue) {
          literalValue = new go.LiteralValue(field.type, prop.clientDefaultValue);
          types.set(keyName, literalValue);
        }
        field.defaultValue = <go.LiteralValue>literalValue;
      }
    }*/
  
    // TODO: XMLInfo
    //field.xml = adaptXMLInfo(prop.schema);
  
    return field;
  }

  private getConstantValues(type: go.ConstantType, valueTypes: Array<tcgc.SdkEnumValueType>): Array<go.ConstantValue> {
    const values = new Array<go.ConstantValue>();
    for (const valueType of valueTypes) {
      const valueTypeName = `${type.name}${naming.ensureNameCase(valueType.name)}`;
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
      /*case m4.SchemaType.ByteArray: {
        const keyName = `literal-${m4.SchemaType.ByteArray}-${constType.value.value}`;
        let literalByteArray = types.get(keyName);
        if (literalByteArray) {
          return <go.LiteralValue>literalByteArray;
        }
        literalByteArray = new go.LiteralValue(adaptBytesType(<m4.ByteArraySchema>constType.valueType), constType.value.value);
        types.set(keyName, literalByteArray);
        return literalByteArray;
      }*/
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
      case 'int32':
      case 'int64': {
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
      case 'guid':
      case 'uuid': {
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

function recursiveGetType(sdkType: tcgc.SdkType): tcgc.SdkType {
  if (sdkType.kind !== 'array' && sdkType.kind !== 'dict') {
    return sdkType;
  }
  return recursiveGetType(sdkType.valueType);
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
    case 'date':
      if (obj.encode !== 'rfc3339') {
        throw new Error(`unsupported date encoding ${obj.encode}`);
      }
      return `${root}-dateRFC3339`;
    case 'datetime':
      return `${root}-${getDateTimeEncoding(obj.encode)}`;
    case 'duration':
      return `${root}-${obj.wireType.kind}`;
    case 'model':
      if (substituteDiscriminator) {
        return `${root}-${naming.createPolymorphicInterfaceName(obj.name)}`;
      }
      return `${root}-${obj.name}`;
    case 'time':
      if (obj.encode !== 'rfc3339') {
        throw new Error(`unsupported time encoding ${obj.encode}`);
      }
      return `${root}-timeRFC3339`;
    default:
      return `${root}-${obj.kind}`;
  }
}

export function isTypePassedByValue(type: tcgc.SdkType): boolean {
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

// aggregate the properties from the provided type and its parent types
function aggregateProperties(model: tcgc.SdkModelType): Array<tcgc.SdkModelPropertyType> {
  const allProps = new Array<tcgc.SdkModelPropertyType>();
  for (const prop of model.properties) {
    allProps.push(prop);
  }

  let parent = model.baseModel;
  while (parent) {
    for (const parentProp of parent.properties) {
      const exists = values(allProps).where(p => { return p.nameInClient === parentProp.nameInClient; }).first();
      if (exists) {
        // don't add the duplicate. the TS compiler has better enforcement than OpenAPI
        // to ensure that duplicate fields with different types aren't added.
        continue;
      }
      allProps.push(parentProp);
    }
    parent = parent.baseModel;
  }
  return allProps;
}
