/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { capitalize } from '../../../naming.go/src/naming.js';
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
  
    // we must adapt all model types first. this is because models can contain cyclic references
    const modelObjs = new Array<ModelTypeSdkModelType>();
    for (const modelType of sdkContext.sdkPackage.models) {
      // TODO: what's the equivalent of x-ms-external?
      const model = this.getModel(modelType);
      modelObjs.push({type: model, obj: modelType});
    }
  
    // now adapt model fields
    for (const modelObj of modelObjs) {
      const props = aggregateProperties(modelObj.obj);
      for (const prop of values(props)) {
        if (prop.kind !== 'property') {
          throw new Error(`unexpected kind ${prop.kind} for property ${prop.nameInClient} in model ${modelObj.obj.name}`);
        }
        const field = this.getModelField(prop, modelObj.obj);
        modelObj.type.fields.push(field);
      }
      this.codeModel.models.push(modelObj.type);
    }
  }

  // returns the Go code model type for the specified SDK type.
  // the operation is idempotent, so getting the same type multiple times
  // returns the same instance of the converted type.
  getPossibleType(type: tcgc.SdkType, elementTypeByValue: boolean): go.PossibleType {
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
      case 'array': {
        // prefer elementTypeByValue. if false, then if the array elements have been explicitly marked as nullable then prefer that, else fall back to our usual algorithm
        const myElementTypeByValue = elementTypeByValue ? true : type.valueType.nullable ? false : this.codeModel.options.sliceElementsByval || isTypePassedByValue(type.valueType);
        const keyName = recursiveKeyName(`array-${myElementTypeByValue}`, type.valueType);
        let arrayType = this.types.get(keyName);
        if (arrayType) {
          return arrayType;
        }
        arrayType = new go.SliceType(this.getPossibleType(type.valueType, elementTypeByValue), myElementTypeByValue);
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
      case 'dict': {
        const valueTypeByValue = isTypePassedByValue(type.valueType);
        const keyName = recursiveKeyName(`dict-${valueTypeByValue}`, type.valueType);
        let mapType = this.types.get(keyName);
        if (mapType) {
          return mapType;
        }
        mapType = new go.MapType(this.getPossibleType(type.valueType, elementTypeByValue), valueTypeByValue);
        this.types.set(keyName, mapType);
        return mapType;
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
      case 'model':
        return this.getModel(type);
      case 'armId':
      case 'duration':
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
    let constType = this.types.get(enumType.name);
    if (constType) {
      return <go.ConstantType>constType;
    }
    constType = new go.ConstantType(enumType.name, getPrimitiveType(enumType.valueType.kind), `Possible${enumType.name}Values`);
    constType.values = this.getConstantValues(constType, enumType.values);
    constType.description = enumType.description;
    this.types.set(enumType.name, constType);
    return constType;
  }

  // converts an SdkModelType to a go.ModelType or go.PolymorphicType if the model is polymorphic
  private getModel(model: tcgc.SdkModelType): go.ModelType | go.PolymorphicType {
    let modelType = this.types.get(model.name);
    if (modelType) {
      return <go.ModelType | go.PolymorphicType>modelType;
    }
  
    // TODO: what's the extension equivalent in TS?
    const annotations = new go.ModelAnnotations(false);
    if (model.discriminatedSubtypes || model.discriminatorValue) {
      throw new Error('discriminators nyi');
    } else {
      // TODO: hard-coded format
      modelType = new go.ModelType(model.name, 'json', annotations);
      // polymorphic types don't have XMLInfo
      // TODO: XMLInfo
    }
    modelType.description = model.description;
    this.types.set(model.name, modelType);
    return modelType;
  }

  private getModelField(prop: tcgc.SdkBodyModelPropertyType, obj: tcgc.SdkModelType): go.ModelField {
    // TODO: hard-coded values
    const annotations = new go.ModelFieldAnnotations(prop.optional == false, false, false, false);
    const field = new go.ModelField(capitalize(prop.nameInClient), this.getPossibleType(prop.type, false), isTypePassedByValue(prop.type), prop.serializedName, annotations);
    field.description = prop.description;
    if (prop.discriminator && obj.discriminatorValue) {
      const keyName = `discriminator-value-${obj.discriminatorValue}`;
      let discriminatorLiteral = <go.LiteralValue>this.types.get(keyName);
      if (!discriminatorLiteral) {
        // the discriminatorValue is either a quoted string or a constant (i.e. enum) value
        if (obj.discriminatorValue[0] === '"') {
          discriminatorLiteral = new go.LiteralValue(new go.PrimitiveType('string'), obj.discriminatorValue);
        } else {
          // find the corresponding constant value
          const value = this.constValues.get(obj.discriminatorValue);
          if (!value) {
            throw new Error(`didn't find a constant value for discriminator value ${obj.discriminatorValue}`);
          }
          discriminatorLiteral = new go.LiteralValue(value.type, value);
        }
      }
      this.types.set(keyName, discriminatorLiteral);
      field.defaultValue = discriminatorLiteral;
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
      let value = this.constValues.get(valueType.name);
      if (!value) {
        value = new go.ConstantValue(`${type.name}${valueType.name}`, type, valueType.value);
        value.description = valueType.description;
        this.constValues.set(valueType.name, value);
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

  private getLiteralValue(constType: tcgc.SdkConstantType): go.LiteralValue {
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
      }
      case m4.SchemaType.Choice:
      case m4.SchemaType.SealedChoice: {
        const keyName = `literal-choice-${constType.value.value}`;
        let literalConst = types.get(keyName);
        if (literalConst) {
          return <go.LiteralValue>literalConst;
        }
        literalConst = new go.LiteralValue(adaptConstantType(<m4.ChoiceSchema>constType.valueType), constType.value.value);
        types.set(keyName, literalConst);
        return literalConst;
      }
      case m4.SchemaType.Date:
      case m4.SchemaType.DateTime:
      case m4.SchemaType.UnixTime: {
        const keyName = `literal-${constType.valueType.language.go!.internalTimeType}-${constType.value.value}`;
        let literalTime = types.get(keyName);
        if (literalTime) {
          return <go.LiteralValue>literalTime;
        }
        literalTime = new go.LiteralValue(new go.TimeType(constType.valueType.language.go!.internalTimeType), constType.value.value);
        types.set(keyName, literalTime);
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

function recursiveKeyName(root: string, obj: tcgc.SdkType): string {
  switch (obj.kind) {
    case 'array':
      return recursiveKeyName(`${root}-array`, obj.valueType);
    case 'dict':
      return recursiveKeyName(`${root}-dict`, obj.valueType);
    case 'date':
      if (obj.encode !== 'rfc3339') {
        throw new Error(`unsupported date encoding ${obj.encode}`);
      }
      return `${root}-dateRFC3339`;
    case 'datetime':
      return `${root}-${getDateTimeEncoding(obj.encode)}`;
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
  type: go.ModelType | go.PolymorphicType;
  obj: tcgc.SdkModelType;
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
      // ensure that the parent doesn't contain any properties with the same name but different type
      const exists = values(allProps).where(p => { return p.nameInClient === parentProp.nameInClient; }).first();
      if (exists) {
        if (exists.type !== parentProp.type) {
          const msg = `type ${model.name} contains duplicate property ${exists.nameInClient} with mismatched types`;
          throw new Error(msg);
        }
        // don't add the duplicate
        continue;
      }
      allProps.push(parentProp);
    }
    parent = parent.baseModel;
  }
  return allProps;
}
