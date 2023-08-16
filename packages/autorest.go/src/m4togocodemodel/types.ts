/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { ArraySchema, ByteArraySchema, ChoiceSchema, ChoiceValue, ConstantSchema, DictionarySchema, Language, NumberSchema, ObjectSchema, Property, Schema, SchemaType, SealedChoiceSchema } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { BytesType, ConstantType, ConstantValue, MapType, ModelAnnotations, ModelField, ModelFieldAnnotations, ModelFormat, ModelType, PolymorphicType, PossibleType, PrimitiveType, SliceType, StandardType, TimeType, BytesEncoding, LiteralValue, InterfaceType, XMLInfo, getTypeDeclaration, isLiteralValueType, isConstantType } from '../gocodemodel/gocodemodel';

// returns true if the language contains a description
export function hasDescription(lang: Language): boolean {
  return (lang.description !== undefined && lang.description.length > 0 && !lang.description.startsWith('MISSING'));
}

// cache of previously created types
const types = new Map<string, PossibleType>();
const constValues = new Map<string, ConstantValue>();

export function adaptConstantType(choice: ChoiceSchema | SealedChoiceSchema): ConstantType {
  let constType = types.get(choice.language.go!.name);
  if (constType) {
    return <ConstantType>constType;
  }
  constType = new ConstantType(choice.language.go!.name, adaptPrimitiveType(choice.choiceType.language.go!.name), choice.language.go!.possibleValuesFunc);
  constType.values = adaptConstantValue(constType, choice.choices);
  if (hasDescription(choice.language.go!)) {
    constType.description = choice.language.go!.description;
  }
  types.set(choice.language.go!.name, constType);
  return constType;
}

function adaptConstantValue(type: ConstantType, choices: Array<ChoiceValue>): Array<ConstantValue> {
  const values = new Array<ConstantValue>();
  for (const choice of choices) {
    let value = constValues.get(choice.language.go!.name);
    if (!value) {
      value = new ConstantValue(choice.language.go!.name, type, choice.value);
      if (hasDescription(choice.language.go!)) {
        value.description = choice.language.go!.description;
      }
      constValues.set(choice.language.go!.name, value);
    }
    values.push(value);
  }
  return values;
}

function adaptPrimitiveType(name: string): 'bool' | 'float32' | 'float64' | 'int32' | 'int64' | 'string' {
  switch (name) {
    case 'bool':
    case 'float32':
    case 'float64':
    case 'int32':
    case 'int64':
    case 'string':
      return name;
    default:
      throw new Error(`unhandled primitive: ${name}`);
  }
}

export function adaptInterfaceType(obj: ObjectSchema, parent?: InterfaceType): InterfaceType {
  let iface = types.get(obj.language.go!.discriminatorInterface);
  if (iface) {
    return <InterfaceType>iface;
  }

  iface = new InterfaceType(obj.language.go!.discriminatorInterface, obj.discriminator!.property.serializedName);
  if (parent) {
    iface.parent = parent;
  }

  types.set(obj.language.go!.discriminatorInterface, iface);
  return iface;
}

export function adaptModel(obj: ObjectSchema): ModelType | PolymorphicType {
  let modelType = types.get(obj.language.go!.name);
  if (modelType) {
    return <ModelType | PolymorphicType>modelType;
  }

  const annotations = new ModelAnnotations(obj.language.go!.omitSerDeMethods);
  if (obj.discriminator || obj.discriminatorValue) {
    let ifaceName: string | undefined;
    if (obj.language.go!.discriminatorInterface) {
      // only discriminators define the discriminatorInterface
      ifaceName = obj.language.go!.discriminatorInterface;
    } else {
      // get it from the parent which must be a discriminator.
      // there are cases where a type might have multiple parents
      // so we iterate over them until we find the interface name
      // (e.g. KerberosKeytabCredentials type in machine learning)
      for (const parent of values( obj.parents?.immediate)) {
        if (parent.language.go!.discriminatorInterface) {
          ifaceName = parent.language.go!.discriminatorInterface;
          break;
        }
      }
    }
    if (!ifaceName) {
      throw new Error(`failed to find discriminator interface name for type ${obj.language.go!.name}`);
    }
    const iface = types.get(ifaceName);
    if (!iface) {
      throw new Error(`didn't find InterfaceType for discriminator interface ${ifaceName} on type ${obj.language.go!.name}`);
    }
    modelType = new PolymorphicType(obj.language.go!.name, <InterfaceType>iface, annotations);
    // only non-root and sub-root discriminators will have a discriminatorValue
    if (obj.discriminatorValue) {
      (<PolymorphicType>modelType).discriminatorValue = obj.discriminatorValue;
    }
  } else {
    modelType = new ModelType(obj.language.go!.name, adaptModelFormat(obj), annotations);
    // polymorphic types don't have XMLInfo
    modelType.xml = adaptXMLInfo(obj);
  }
  if (hasDescription(obj.language.go!)) {
    modelType.description = obj.language.go!.description;
  }

  types.set(obj.language.go!.name, modelType);
  return modelType;
}

export function adaptModelField(prop: Property, obj: ObjectSchema): ModelField {
  const annotations = new ModelFieldAnnotations(prop.required === true, prop.readOnly === true, prop.language.go!.isAdditionalProperties === true, prop.isDiscriminator === true);
  const field = new ModelField(prop.language.go!.name, adaptPossibleType(prop.schema), prop.language.go!.byValue === true, prop.serializedName, annotations);
  if (hasDescription(prop.language.go!)) {
    field.description = prop.language.go!.description;
  }
  if (prop.isDiscriminator && obj.discriminatorValue) {
    const keyName = `discriminator-value-${obj.discriminatorValue}`;
    let discriminatorLiteral = <LiteralValue>types.get(keyName);
    if (!discriminatorLiteral) {
      // the discriminatorValue is either a quoted string or a constant (i.e. enum) value
      if (obj.discriminatorValue[0] === '"') {
        discriminatorLiteral = new LiteralValue(new PrimitiveType('string'), obj.discriminatorValue);
      } else {
        // find the corresponding constant value
        const value = constValues.get(obj.discriminatorValue);
        if (!value) {
          throw new Error(`didn't find a constant value for discriminator value ${obj.discriminatorValue}`);
        }
        discriminatorLiteral = new LiteralValue(value.type, value);
      }
    }
    types.set(keyName, discriminatorLiteral);
    field.defaultValue = discriminatorLiteral;
  } else if (prop.clientDefaultValue) {
    if (!isLiteralValueType(field.type)) {
      throw new Error(`unsupported default value type ${getTypeDeclaration(field.type)} for field ${field.fieldName}`);
    }
    if (isConstantType(field.type)) {
      // find the corresponding ConstantValue
      const constType = types.get(field.type.name);
      if (!constType) {
        throw new Error(`didn't find ConstantType for ${field.type.name}`);
      }
      let found = false;
      for (const val of values((<ConstantType>constType).values)) {
        if (val.value === prop.clientDefaultValue) {
          const keyName = `literal-${val.valueName}`;
          let literalValue = types.get(keyName);
          if (!literalValue) {
            literalValue = new LiteralValue(field.type, val);
            types.set(keyName, literalValue);
          }
          field.defaultValue = <LiteralValue>literalValue;
          found = true;
          break;
        }
      }
      if (!found) {
        throw new Error(`didn't find ConstantValue for ${prop.clientDefaultValue}`);
      }
    } else {
      const keyName = `literal-${getTypeDeclaration(field.type)}-${prop.clientDefaultValue}`;
      let literalValue = types.get(keyName);
      if (!literalValue) {
        literalValue = new LiteralValue(field.type, prop.clientDefaultValue);
        types.set(keyName, literalValue);
      }
      field.defaultValue = <LiteralValue>literalValue;
    }
  }

  field.xml = adaptXMLInfo(prop.schema);

  return field;
}

function adaptModelFormat(obj: ObjectSchema): ModelFormat {
  if (obj.language.go!.marshallingFormat === 'json') {
    return 'json';
  } else if (obj.language.go!.marshallingFormat === 'xml') {
    return 'xml';
  } else {
    throw new Error(`unsupported marshalling format ${obj.language.go!.marshallingFormat}`);
  }
}

export function adaptXMLInfo(obj: Schema): XMLInfo | undefined {
  const xmlInfo = new XMLInfo();
  let includeXMLField = false;
  if (obj.serialization?.xml?.name) {
    xmlInfo.name = obj.serialization?.xml?.name;
    includeXMLField = true;
  }
  if (obj.serialization?.xml?.text) {
    xmlInfo.text = true;
    includeXMLField = true;
  }
  if (obj.serialization?.xml?.attribute) {
    xmlInfo.attribute = true;
    includeXMLField = true;
  }
  if (obj.type === SchemaType.Array) {
    const asArray = <ArraySchema>obj;
    if (obj.serialization?.xml?.wrapped) {
      if (asArray.elementType.serialization?.xml?.name) {
        xmlInfo.wraps = asArray.elementType.serialization.xml.name;
      } else {
        xmlInfo.wraps = asArray.elementType.language.go!.name;
      }
      includeXMLField = true;
    } else if (asArray.elementType.serialization?.xml?.name) {
      xmlInfo.name = asArray.elementType.serialization.xml.name;
      includeXMLField = true;
    }
  }
  if (obj.language.go!.xmlWrapperName) {
    xmlInfo.wrapper = obj.language.go!.xmlWrapperName;
    includeXMLField = true;
  }

  if (includeXMLField) {
    return xmlInfo;
  }

  return undefined;
}

// converts an M4 schema type to a Go code model type
export function adaptPossibleType(schema: Schema, elementTypeByValue?: boolean): PossibleType {
  const rawJSONAsBytes = <boolean>schema.language.go!.rawJSONAsBytes;
  switch (schema.type) {
    case SchemaType.Any: {
      if (rawJSONAsBytes) {
        const anyRawJSONKey = `${SchemaType.Any}-raw-json`;
        let anyRawJSON = types.get(anyRawJSONKey);
        if (anyRawJSON) {
          return anyRawJSON;
        }
        anyRawJSON = new SliceType(new PrimitiveType('byte'), true);
        anyRawJSON.rawJSONAsBytes = true;
        types.set(anyRawJSONKey, anyRawJSON);
        return anyRawJSON;
      }
      let anyType = types.get(SchemaType.Any);
      if (anyType) {
        return anyType;
      }
      anyType = new PrimitiveType('any');
      types.set(SchemaType.Any, anyType);
      return anyType;
    }
    case SchemaType.AnyObject: {
      if (rawJSONAsBytes) {
        const anyObjectRawJSONKey = `${SchemaType.Any}-raw-json`;
        let anyObjectRawJSON = types.get(anyObjectRawJSONKey);
        if (anyObjectRawJSON) {
          return anyObjectRawJSON;
        }
        anyObjectRawJSON = new SliceType(new PrimitiveType('byte'), true);
        anyObjectRawJSON.rawJSONAsBytes = true;
        types.set(anyObjectRawJSONKey, anyObjectRawJSON);
        return anyObjectRawJSON;
      }
      let anyObject = types.get(SchemaType.AnyObject);
      if (anyObject) {
        return anyObject;
      }
      anyObject = new MapType(new PrimitiveType('any'), true);
      types.set(SchemaType.AnyObject, anyObject);
      return anyObject;
    }
    case SchemaType.ArmId: {
      let stringType = types.get(SchemaType.String);
      if (stringType) {
        return stringType;
      }
      stringType = new PrimitiveType('string');
      types.set(SchemaType.ArmId, stringType);
      return stringType;
    }
    case SchemaType.Array: {
      let myElementTypeByValue = !schema.language.go!.elementIsPtr;
      if (elementTypeByValue) {
        myElementTypeByValue = elementTypeByValue;
      }
      const keyName = recursiveKeyName(`${SchemaType.Array}-${myElementTypeByValue}`, (<ArraySchema>schema).elementType);
      let arrayType = types.get(keyName);
      if (arrayType) {
        return arrayType;
      }
      arrayType = new SliceType(adaptPossibleType((<ArraySchema>schema).elementType, elementTypeByValue), myElementTypeByValue);
      types.set(keyName, arrayType);
      return arrayType;
    }
    case SchemaType.Boolean: {
      let primitiveBool = types.get(SchemaType.Boolean);
      if (primitiveBool) {
        return primitiveBool;
      }
      primitiveBool = new PrimitiveType('bool');
      types.set(SchemaType.Boolean, primitiveBool);
      return primitiveBool;
    }
    case SchemaType.Binary: {
      let binaryType = types.get(SchemaType.Binary);
      if (binaryType) {
        return binaryType;
      }
      binaryType = new StandardType('io.ReadSeekCloser', 'io');
      types.set(SchemaType.Binary, binaryType);
      return binaryType;
    }
    case SchemaType.ByteArray:
      return adaptBytesType(<ByteArraySchema>schema);
    case SchemaType.Char: {
      let rune = types.get(SchemaType.Char);
      if (rune) {
        return rune;
      }
      rune = new PrimitiveType('rune');
      types.set(SchemaType.Char, rune);
      return rune;
    }
    case SchemaType.Choice:
      return adaptConstantType(<ChoiceSchema>schema);
    case SchemaType.Constant:
      return adaptLiteralValue(<ConstantSchema>schema);
    case SchemaType.Date:
    case SchemaType.DateTime:
    case SchemaType.Time:
    case SchemaType.UnixTime: {
      let time = types.get(schema.language.go!.internalTimeType);
      if (time) {
        return time;
      }
      time = new TimeType(schema.language.go!.internalTimeType);
      types.set(schema.language.go!.internalTimeType, time);
      return time;
    }
    case SchemaType.Dictionary: {
      const valueTypeByValue = !schema.language.go!.elementIsPtr;
      const keyName = recursiveKeyName(`${SchemaType.Dictionary}-${valueTypeByValue}`, (<DictionarySchema>schema).elementType);
      let mapType = types.get(keyName);
      if (mapType) {
        return mapType;
      }
      mapType = new MapType(adaptPossibleType((<DictionarySchema>schema).elementType, elementTypeByValue), valueTypeByValue);
      types.set(keyName, mapType);
      return mapType;
    }
    case SchemaType.Duration: {
      let duration = types.get(SchemaType.Duration);
      if (duration) {
        return duration;
      }
      duration = new PrimitiveType('string');
      types.set(SchemaType.Duration, duration);
      return duration;
    }
    case SchemaType.Integer: {
      if ((<NumberSchema>schema).precision === 32) {
        const int32Key = 'int32';
        let int32 = types.get(int32Key);
        if (int32) {
          return int32;
        }
        int32 = new PrimitiveType(int32Key);
        types.set(int32Key, int32);
        return int32;
      }
      const int64Key = 'int64';
      let int64 = types.get(int64Key);
      if (int64) {
        return int64;
      }
      int64 = new PrimitiveType(int64Key);
      types.set(int64Key, int64);
      return int64;
    }
    case SchemaType.Number: {
      if ((<NumberSchema>schema).precision === 32) {
        const float32Key = 'float32';
        let float32 = types.get(float32Key);
        if (float32) {
          return float32;
        }
        float32 = new PrimitiveType(float32Key);
        types.set(float32Key, float32);
        return float32;
      }
      const float64Key = 'float64';
      let float64 = types.get(float64Key);
      if (float64) {
        return float64;
      }
      float64 = new PrimitiveType(float64Key);
      types.set(float64Key, float64);
      return float64;
    }
    case SchemaType.Object:
      return adaptModel(<ObjectSchema>schema);
    case SchemaType.SealedChoice:
      return adaptConstantType(<SealedChoiceSchema>schema);
    case SchemaType.String: {
      let stringType = types.get(SchemaType.String);
      if (stringType) {
        return stringType;
      }
      stringType = new PrimitiveType('string');
      types.set(SchemaType.String, stringType);
      return stringType;
    }
    case SchemaType.Uri: {
      let uriType = types.get(SchemaType.Uri);
      if (uriType) {
        return uriType;
      }
      uriType = new PrimitiveType('string');
      types.set(SchemaType.Uri, uriType);
      return uriType;
    }
    case SchemaType.Uuid: {
      let uuid = types.get(SchemaType.Uuid);
      if (uuid) {
        return uuid;
      }
      uuid = new PrimitiveType('string');
      types.set(SchemaType.Uuid, uuid);
      return uuid;
    }
    default:
      throw new Error(`unhandled property schema type ${schema.type}`);
  }
}

function adaptLiteralValue(constSchema: ConstantSchema): LiteralValue {
  switch (constSchema.valueType.type) {
    case SchemaType.Boolean: {
      const keyName = `literal-${SchemaType.Boolean}-${constSchema.value.value}`;
      let literalBool = types.get(keyName);
      if (literalBool) {
        return <LiteralValue>literalBool;
      }
      literalBool = new LiteralValue(new PrimitiveType('bool'), constSchema.value.value);
      types.set(keyName, literalBool);
      return literalBool;
    }
    case SchemaType.ByteArray: {
      const keyName = `literal-${SchemaType.ByteArray}-${constSchema.value.value}`;
      let literalByteArray = types.get(keyName);
      if (literalByteArray) {
        return <LiteralValue>literalByteArray;
      }
      literalByteArray = new LiteralValue(adaptBytesType(<ByteArraySchema>constSchema.valueType), constSchema.value.value);
      types.set(keyName, literalByteArray);
      return literalByteArray;
    }
    case SchemaType.Choice:
    case SchemaType.SealedChoice: {
      const keyName = `literal-choice-${constSchema.value.value}`;
      let literalConst = types.get(keyName);
      if (literalConst) {
        return <LiteralValue>literalConst;
      }
      literalConst = new LiteralValue(adaptConstantType(<ChoiceSchema>constSchema.valueType), constSchema.value.value);
      types.set(keyName, literalConst);
      return literalConst;
    }
    case SchemaType.Date:
    case SchemaType.DateTime:
    case SchemaType.UnixTime: {
      const keyName = `literal-${constSchema.valueType.language.go!.internalTimeType}-${constSchema.value.value}`;
      let literalTime = types.get(keyName);
      if (literalTime) {
        return <LiteralValue>literalTime;
      }
      literalTime = new LiteralValue(new TimeType(constSchema.valueType.language.go!.internalTimeType), constSchema.value.value);
      types.set(keyName, literalTime);
      return literalTime;
    }
    case SchemaType.Integer: {
      const keyName = `literal-int${(<NumberSchema>constSchema.valueType).precision}-${constSchema.value.value}`;
      let literalInt = types.get(keyName);
      if (literalInt) {
        return <LiteralValue>literalInt;
      }
      if ((<NumberSchema>constSchema.valueType).precision === 32) {
        literalInt = new LiteralValue(new PrimitiveType('int32'), constSchema.value.value);
      } else {
        literalInt = new LiteralValue(new PrimitiveType('int64'), constSchema.value.value);
      }
      types.set(keyName, literalInt);
      return literalInt;
    }
    case SchemaType.Number: {
      const keyName = `literal-float${(<NumberSchema>constSchema.valueType).precision}-${constSchema.value.value}`;
      let literalFloat = types.get(keyName);
      if (literalFloat) {
        return <LiteralValue>literalFloat;
      }
      if ((<NumberSchema>constSchema.valueType).precision === 32) {
        literalFloat = new LiteralValue(new PrimitiveType('float32'), constSchema.value.value);
      } else {
        literalFloat = new LiteralValue(new PrimitiveType('float64'), constSchema.value.value);
      }
      types.set(keyName, literalFloat);
      return literalFloat;
    }
    case SchemaType.String:
    case SchemaType.Duration:
    case SchemaType.Uuid: {
      const keyName = `literal-string-${constSchema.value.value}`;
      let literalString = types.get(keyName);
      if (literalString) {
        return <LiteralValue>literalString;
      }
      literalString = new LiteralValue(new PrimitiveType('string'), constSchema.value.value);
      types.set(keyName, literalString);
      return literalString;
    }
    default:
      throw new Error(`unsupported scheam type ${constSchema.valueType.type} for LiteralValue`);
  }
}

function adaptBytesType(schema: ByteArraySchema): BytesType {
  let format: BytesEncoding = 'Std';
  if (schema.format === 'base64url') {
    format = 'URL';
  }
  const keyName = `${SchemaType.ByteArray}-${format}`;
  let bytesType = types.get(keyName);
  if (bytesType) {
    return <BytesType>bytesType;
  }
  bytesType = new BytesType(format);
  types.set(keyName, bytesType);
  return bytesType;
}

function recursiveKeyName(root: string, obj: Schema): string {
  switch (obj.type) {
    case SchemaType.Array:
      return recursiveKeyName(`${root}-${SchemaType.Array}`, (<ArraySchema>obj).elementType);
    case SchemaType.Dictionary:
      return recursiveKeyName(`${root}-${SchemaType.Dictionary}`, (<DictionarySchema>obj).elementType);
    case SchemaType.Date:
    case SchemaType.DateTime:
    case SchemaType.UnixTime:
      return obj.language.go!.internalTimeType;
    default:
      return `${root}-${obj.language.go!.name}`;
  }
}
