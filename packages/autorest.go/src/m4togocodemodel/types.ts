/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

/* eslint-disable @typescript-eslint/no-unsafe-argument */

import * as m4 from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import * as go from '../../../codemodel.go/src/index.js';

// returns true if the language contains a description
export function hasDescription(lang: m4.Language): boolean {
  return (lang.description !== undefined && lang.description.length > 0 && !lang.description.startsWith('MISSING'));
}

// cache of previously created types
const types = new Map<string, go.WireType>();
const constValues = new Map<string, go.ConstantValue>();

export function adaptConstantType(choice: m4.ChoiceSchema | m4.SealedChoiceSchema): go.Constant {
  let constType = types.get(choice.language.go!.name);
  if (constType) {
    return <go.Constant>constType;
  }
  constType = new go.Constant(choice.language.go!.name, adaptPrimitiveType(choice.choiceType.language.go!.name), choice.language.go!.possibleValuesFunc);
  constType.values = adaptConstantValue(constType, choice.choices);
  if (hasDescription(choice.language.go!)) {
    constType.docs.description = choice.language.go!.description;
  }
  types.set(choice.language.go!.name, constType);
  return constType;
}

function adaptConstantValue(type: go.Constant, choices: Array<m4.ChoiceValue>): Array<go.ConstantValue> {
  const values = new Array<go.ConstantValue>();
  for (const choice of choices) {
    let value = constValues.get(choice.language.go!.name);
    if (!value) {
      value = new go.ConstantValue(choice.language.go!.name, type, choice.value);
      if (hasDescription(choice.language.go!)) {
        value.docs.description = choice.language.go!.description;
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

export function adaptInterfaceType(obj: m4.ObjectSchema, parent?: go.Interface): go.Interface {
  let iface = types.get(obj.language.go!.discriminatorInterface);
  if (iface) {
    return <go.Interface>iface;
  }

  iface = new go.Interface(obj.language.go!.discriminatorInterface, obj.discriminator!.property.serializedName);
  if (parent) {
    iface.parent = parent;
  }

  types.set(obj.language.go!.discriminatorInterface, iface);
  return iface;
}

export function adaptModel(obj: m4.ObjectSchema): go.Model | go.PolymorphicModel {
  let modelType = types.get(obj.language.go!.name);
  if (modelType) {
    return <go.Model | go.PolymorphicModel>modelType;
  }

  const annotations = new go.ModelAnnotations(obj.language.go!.omitSerDeMethods, false);
  if (obj.discriminator || obj.discriminatorValue) {
    let ifaceName: string | undefined;
    if (obj.language.go!.discriminatorInterface) {
      // only discriminators define the discriminatorInterface
      ifaceName = <string>obj.language.go!.discriminatorInterface;
    } else {
      // get it from the parent which must be a discriminator.
      // there are cases where a type might have multiple parents
      // so we iterate over them until we find the interface name
      // (e.g. KerberosKeytabCredentials type in machine learning)
      for (const parent of values( obj.parents?.immediate)) {
        if (parent.language.go!.discriminatorInterface) {
          ifaceName = <string>parent.language.go!.discriminatorInterface;
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
    modelType = new go.PolymorphicModel(obj.language.go!.name, <go.Interface>iface, annotations, adaptUsage(obj));
    // only non-root and sub-root discriminators will have a discriminatorValue
    if (obj.discriminatorValue) {
      modelType.discriminatorValue = getDiscriminatorLiteral(obj.discriminatorValue);
    }
  } else {
    modelType = new go.Model(obj.language.go!.name, annotations, adaptUsage(obj));
    // polymorphic types don't have XMLInfo
    modelType.xml = adaptXMLInfo(obj);
  }
  if (hasDescription(obj.language.go!)) {
    modelType.docs.description = obj.language.go!.description;
  }

  types.set(obj.language.go!.name, modelType);
  return modelType;
}

// returns the usage flags for this schema
function adaptUsage(obj: m4.ObjectSchema): go.UsageFlags {
  let flags = go.UsageFlags.None;
  if (values(obj.usage).any((u) => { return u === m4.SchemaContext.Input; })) {
    flags = go.UsageFlags.Input;
  }
  if (values(obj.usage).any((u) => { return u === m4.SchemaContext.Output; })) {
    flags |= go.UsageFlags.Output;
  }
  return flags;
}

function getDiscriminatorLiteral(discriminatorValue: string): go.Literal {
  const createLiteralValue = function(type: go.LiteralType, value: string | go.ConstantValue): go.Literal {
    let valueKey: go.ConstantValueType;
    if (typeof value === 'string') {
      valueKey = value;
    } else {
      valueKey = value.value;
    }
    const keyName = `literal-${go.getTypeDeclaration(type)}-${valueKey}`;
    let literalString = types.get(keyName);
    if (literalString) {
      return <go.Literal>literalString;
    }
    literalString = new go.Literal(type, value);
    types.set(keyName, literalString);
    return literalString;
  };

  let discriminatorLiteral: go.Literal;
  // the discriminatorValue is either a quoted string or a constant (i.e. enum) value
  if (discriminatorValue[0] === '"') {
    discriminatorLiteral = createLiteralValue(new go.String(), discriminatorValue);
  } else {
    // find the corresponding constant value
    const value = constValues.get(discriminatorValue);
    if (!value) {
      throw new Error(`didn't find a constant value for discriminator value ${discriminatorValue}`);
    }
    discriminatorLiteral = createLiteralValue(value.type, value);
  }
  return discriminatorLiteral;
}

export function adaptModelField(prop: m4.Property, obj: m4.ObjectSchema): go.ModelField {
  const fieldType = adaptWireType(prop.schema);
  let required = prop.required === true;
  if (fieldType.kind === 'literal') {
    // for OpenAPI, literal values are always considered required
    required = true;
  }
  const annotations = new go.ModelFieldAnnotations(required, prop.readOnly === true, prop.language.go!.isAdditionalProperties === true, prop.isDiscriminator === true, false);
  const field = new go.ModelField(prop.language.go!.name, fieldType, prop.language.go!.byValue === true, prop.serializedName, annotations);
  if (hasDescription(prop.language.go!)) {
    field.docs.description = prop.language.go!.description;
  }
  if (prop.isDiscriminator && obj.discriminatorValue) {
    field.defaultValue = getDiscriminatorLiteral(obj.discriminatorValue);
  } else if (prop.clientDefaultValue) {
    if (!go.isLiteralValueType(field.type)) {
      throw new Error(`unsupported default value type ${go.getTypeDeclaration(field.type)} for field ${field.name}`);
    }
    if (field.type.kind === 'constant') {
      // find the corresponding ConstantValue
      const constType = types.get(field.type.name);
      if (!constType) {
        throw new Error(`didn't find ConstantType for ${field.type.name}`);
      }
      let found = false;
      for (const val of values((<go.Constant>constType).values)) {
        if (val.value === prop.clientDefaultValue) {
          const keyName = `literal-${val.name}`;
          let literalValue = types.get(keyName);
          if (!literalValue) {
            literalValue = new go.Literal(field.type, val);
            types.set(keyName, literalValue);
          }
          field.defaultValue = <go.Literal>literalValue;
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
        literalValue = new go.Literal(field.type, prop.clientDefaultValue);
        types.set(keyName, literalValue);
      }
      field.defaultValue = <go.Literal>literalValue;
    }
  }

  field.xml = adaptXMLInfo(prop.schema);

  return field;
}

export function adaptXMLInfo(obj: m4.Schema): go.XMLInfo | undefined {
  const xmlInfo = new go.XMLInfo();
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
  if (obj.type === m4.SchemaType.Array) {
    const asArray = <m4.ArraySchema>obj;
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
    xmlInfo.wrapper = <string | undefined>obj.language.go!.xmlWrapperName;
    includeXMLField = true;
  }

  if (includeXMLField) {
    return xmlInfo;
  }

  return undefined;
}

// converts an M4 schema type to a Go code model type
export function adaptWireType(schema: m4.Schema, elementTypeByValue?: boolean): go.WireType {
  const rawJSONAsBytes = <boolean>schema.language.go!.rawJSONAsBytes;
  switch (schema.type) {
    case m4.SchemaType.Any: {
      if (rawJSONAsBytes) {
        const anyRawJSONKey = `${m4.SchemaType.Any}-raw-json`;
        let anyRawJSON = types.get(anyRawJSONKey);
        if (anyRawJSON) {
          return anyRawJSON;
        }
        anyRawJSON = new go.RawJSON();
        types.set(anyRawJSONKey, anyRawJSON);
        return anyRawJSON;
      }
      let anyType = types.get(m4.SchemaType.Any);
      if (anyType) {
        return anyType;
      }
      anyType = new go.Any();
      types.set(m4.SchemaType.Any, anyType);
      return anyType;
    }
    case m4.SchemaType.AnyObject: {
      if (rawJSONAsBytes) {
        const anyObjectRawJSONKey = `${m4.SchemaType.Any}-raw-json`;
        let anyObjectRawJSON = types.get(anyObjectRawJSONKey);
        if (anyObjectRawJSON) {
          return anyObjectRawJSON;
        }
        anyObjectRawJSON = new go.RawJSON();
        types.set(anyObjectRawJSONKey, anyObjectRawJSON);
        return anyObjectRawJSON;
      }
      let anyObject = types.get(m4.SchemaType.AnyObject);
      if (anyObject) {
        return anyObject;
      }
      anyObject = new go.Map(new go.Any(), true);
      types.set(m4.SchemaType.AnyObject, anyObject);
      return anyObject;
    }
    case m4.SchemaType.ArmId: {
      let stringType = types.get(m4.SchemaType.String);
      if (stringType) {
        return stringType;
      }
      stringType = new go.String();
      types.set(m4.SchemaType.ArmId, stringType);
      return stringType;
    }
    case m4.SchemaType.Array: {
      let myElementTypeByValue = !schema.language.go!.elementIsPtr;
      if (elementTypeByValue) {
        myElementTypeByValue = elementTypeByValue;
      }
      const keyName = recursiveKeyName(`${m4.SchemaType.Array}-${myElementTypeByValue}`, (<m4.ArraySchema>schema).elementType);
      let arrayType = types.get(keyName);
      if (arrayType) {
        return arrayType;
      }
      arrayType = new go.Slice(adaptWireType((<m4.ArraySchema>schema).elementType, elementTypeByValue), myElementTypeByValue);
      types.set(keyName, arrayType);
      return arrayType;
    }
    case m4.SchemaType.Binary: {
      let binaryType = types.get(m4.SchemaType.Binary);
      if (binaryType) {
        return binaryType;
      }
      binaryType = new go.QualifiedType('ReadSeekCloser', 'io');
      types.set(m4.SchemaType.Binary, binaryType);
      return binaryType;
    }
    case m4.SchemaType.Boolean: {
      let primitiveBool = types.get(m4.SchemaType.Boolean);
      if (primitiveBool) {
        return primitiveBool;
      }
      primitiveBool = new go.Scalar('bool', false);
      types.set(m4.SchemaType.Boolean, primitiveBool);
      return primitiveBool;
    }
    case m4.SchemaType.ByteArray:
      return adaptBytesType(<m4.ByteArraySchema>schema);
    case m4.SchemaType.Char: {
      let rune = types.get(m4.SchemaType.Char);
      if (rune) {
        return rune;
      }
      rune = new go.Scalar('rune', false);
      types.set(m4.SchemaType.Char, rune);
      return rune;
    }
    case m4.SchemaType.Choice:
      return adaptConstantType(<m4.ChoiceSchema>schema);
    case m4.SchemaType.Constant:
      return adaptLiteralValue(<m4.ConstantSchema>schema);
    case m4.SchemaType.Credential: {
      let credType = types.get(m4.SchemaType.Credential);
      if (credType) {
        return credType;
      }
      credType = new go.String();
      types.set(m4.SchemaType.Credential, credType);
      return credType;
    }
    case m4.SchemaType.Date:
    case m4.SchemaType.DateTime:
    case m4.SchemaType.Time:
    case m4.SchemaType.UnixTime: {
      let time = types.get(schema.language.go!.internalTimeType);
      if (time) {
        return time;
      }
      time = new go.Time(schema.language.go!.internalTimeType, false);
      types.set(schema.language.go!.internalTimeType, time);
      return time;
    }
    case m4.SchemaType.Dictionary: {
      const valueTypeByValue = !schema.language.go!.elementIsPtr;
      const keyName = recursiveKeyName(`${m4.SchemaType.Dictionary}-${valueTypeByValue}`, (<m4.DictionarySchema>schema).elementType);
      let mapType = types.get(keyName);
      if (mapType) {
        return mapType;
      }
      mapType = new go.Map(adaptWireType((<m4.DictionarySchema>schema).elementType, elementTypeByValue), valueTypeByValue);
      types.set(keyName, mapType);
      return mapType;
    }
    case m4.SchemaType.Duration: {
      let duration = types.get(m4.SchemaType.Duration);
      if (duration) {
        return duration;
      }
      duration = new go.String();
      types.set(m4.SchemaType.Duration, duration);
      return duration;
    }
    case m4.SchemaType.Integer: {
      if ((<m4.NumberSchema>schema).precision === 32) {
        const int32Key = 'int32';
        let int32 = types.get(int32Key);
        if (int32) {
          return int32;
        }
        int32 = new go.Scalar(int32Key, false);
        types.set(int32Key, int32);
        return int32;
      }
      const int64Key = 'int64';
      let int64 = types.get(int64Key);
      if (int64) {
        return int64;
      }
      int64 = new go.Scalar(int64Key, false);
      types.set(int64Key, int64);
      return int64;
    }
    case m4.SchemaType.Number: {
      if ((<m4.NumberSchema>schema).precision === 32) {
        const float32Key = 'float32';
        let float32 = types.get(float32Key);
        if (float32) {
          return float32;
        }
        float32 = new go.Scalar(float32Key, false);
        types.set(float32Key, float32);
        return float32;
      }
      const float64Key = 'float64';
      let float64 = types.get(float64Key);
      if (float64) {
        return float64;
      }
      float64 = new go.Scalar(float64Key, false);
      types.set(float64Key, float64);
      return float64;
    }
    case m4.SchemaType.Object:
      return adaptModel(<m4.ObjectSchema>schema);
    case m4.SchemaType.ODataQuery: {
      let odataType = types.get(m4.SchemaType.ODataQuery);
      if (odataType) {
        return odataType;
      }
      odataType = new go.String();
      types.set(m4.SchemaType.ODataQuery, odataType);
      return odataType;
    }
    case m4.SchemaType.SealedChoice:
      return adaptConstantType(<m4.SealedChoiceSchema>schema);
    case m4.SchemaType.String: {
      let stringType = types.get(m4.SchemaType.String);
      if (stringType) {
        return stringType;
      }
      stringType = new go.String();
      types.set(m4.SchemaType.String, stringType);
      return stringType;
    }
    case m4.SchemaType.Uri: {
      let uriType = types.get(m4.SchemaType.Uri);
      if (uriType) {
        return uriType;
      }
      uriType = new go.String();
      types.set(m4.SchemaType.Uri, uriType);
      return uriType;
    }
    case m4.SchemaType.Uuid: {
      let uuid = types.get(m4.SchemaType.Uuid);
      if (uuid) {
        return uuid;
      }
      uuid = new go.String();
      types.set(m4.SchemaType.Uuid, uuid);
      return uuid;
    }
    default:
      throw new Error(`unhandled property schema type ${schema.type}`);
  }
}

function adaptLiteralValue(constSchema: m4.ConstantSchema): go.Literal {
  switch (constSchema.valueType.type) {
    case m4.SchemaType.Boolean: {
      const keyName = `literal-${m4.SchemaType.Boolean}-${constSchema.value.value}`;
      let literalBool = types.get(keyName);
      if (literalBool) {
        return <go.Literal>literalBool;
      }
      literalBool = new go.Literal(new go.Scalar('bool', false), constSchema.value.value);
      types.set(keyName, literalBool);
      return literalBool;
    }
    case m4.SchemaType.ByteArray: {
      const keyName = `literal-${m4.SchemaType.ByteArray}-${constSchema.value.value}`;
      let literalByteArray = types.get(keyName);
      if (literalByteArray) {
        return <go.Literal>literalByteArray;
      }
      literalByteArray = new go.Literal(adaptBytesType(<m4.ByteArraySchema>constSchema.valueType), constSchema.value.value);
      types.set(keyName, literalByteArray);
      return literalByteArray;
    }
    case m4.SchemaType.Choice:
    case m4.SchemaType.SealedChoice: {
      const keyName = `literal-${constSchema.language.go!.name}-${constSchema.value.value}`;
      let literalConst = types.get(keyName);
      if (literalConst) {
        return <go.Literal>literalConst;
      }
      literalConst = new go.Literal(adaptConstantType(<m4.ChoiceSchema>constSchema.valueType), constSchema.value.value);
      types.set(keyName, literalConst);
      return literalConst;
    }
    case m4.SchemaType.Date:
    case m4.SchemaType.DateTime:
    case m4.SchemaType.UnixTime: {
      const keyName = `literal-${constSchema.valueType.language.go!.internalTimeType}-${constSchema.value.value}`;
      let literalTime = types.get(keyName);
      if (literalTime) {
        return <go.Literal>literalTime;
      }
      literalTime = new go.Literal(new go.Time(constSchema.valueType.language.go!.internalTimeType, false), constSchema.value.value);
      types.set(keyName, literalTime);
      return literalTime;
    }
    case m4.SchemaType.Integer: {
      const keyName = `literal-int${(<m4.NumberSchema>constSchema.valueType).precision}-${constSchema.value.value}`;
      let literalInt = types.get(keyName);
      if (literalInt) {
        return <go.Literal>literalInt;
      }
      if ((<m4.NumberSchema>constSchema.valueType).precision === 32) {
        literalInt = new go.Literal(new go.Scalar('int32', false), constSchema.value.value);
      } else {
        literalInt = new go.Literal(new go.Scalar('int64', false), constSchema.value.value);
      }
      types.set(keyName, literalInt);
      return literalInt;
    }
    case m4.SchemaType.Number: {
      const keyName = `literal-float${(<m4.NumberSchema>constSchema.valueType).precision}-${constSchema.value.value}`;
      let literalFloat = types.get(keyName);
      if (literalFloat) {
        return <go.Literal>literalFloat;
      }
      if ((<m4.NumberSchema>constSchema.valueType).precision === 32) {
        literalFloat = new go.Literal(new go.Scalar('float32', false), constSchema.value.value);
      } else {
        literalFloat = new go.Literal(new go.Scalar('float64', false), constSchema.value.value);
      }
      types.set(keyName, literalFloat);
      return literalFloat;
    }
    case m4.SchemaType.String:
    case m4.SchemaType.Duration:
    case m4.SchemaType.Uuid: {
      const keyName = `literal-string-${constSchema.value.value}`;
      let literalString = types.get(keyName);
      if (literalString) {
        return <go.Literal>literalString;
      }
      literalString = new go.Literal(new go.String(), constSchema.value.value);
      types.set(keyName, literalString);
      return literalString;
    }
    default:
      throw new Error(`unsupported scheam type ${constSchema.valueType.type} for LiteralValue`);
  }
}

function adaptBytesType(schema: m4.ByteArraySchema): go.EncodedBytes {
  let format: go.BytesEncoding = 'Std';
  if (schema.format === 'base64url') {
    format = 'URL';
  }
  const keyName = `${m4.SchemaType.ByteArray}-${format}`;
  let bytesType = types.get(keyName);
  if (bytesType) {
    return <go.EncodedBytes>bytesType;
  }
  bytesType = new go.EncodedBytes(format);
  types.set(keyName, bytesType);
  return bytesType;
}

function recursiveKeyName(root: string, obj: m4.Schema): string {
  switch (obj.type) {
    case m4.SchemaType.Array:
      return recursiveKeyName(`${root}-${m4.SchemaType.Array}`, (<m4.ArraySchema>obj).elementType);
    case m4.SchemaType.Dictionary:
      return recursiveKeyName(`${root}-${m4.SchemaType.Dictionary}`, (<m4.DictionarySchema>obj).elementType);
    case m4.SchemaType.Date:
    case m4.SchemaType.DateTime:
    case m4.SchemaType.UnixTime:
      return `${root}-${obj.language.go!.internalTimeType}`;
    default:
      return `${root}-${obj.language.go!.name}`;
  }
}
