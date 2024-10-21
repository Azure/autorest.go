/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

// Docs contains the values used in doc comment generation.
export interface Docs {
  // the high level summary
  summary?: string;

  // detailed description
  description?: string;
}

// Struct describes a vanilla struct definition (pretty much exclusively used for parameter groups/options bag types)
// UDTs that are sent/received are modeled as ModelType.
export interface StructType {
  name: string;

  docs: Docs;

  // there are only a few corner-cases where a struct has no fields
  fields: Array<StructField>;
}

// ModelType is a struct that participates in serialization over the wire.
export interface ModelType extends StructType {
  fields: Array<ModelField>;

  annotations: ModelAnnotations;

  usage: UsageFlags;

  xml?: XMLInfo;
}

export interface ModelAnnotations {
  omitSerDeMethods: boolean;

  // indicates the model should be converted into multipart/form data
  multipartFormData: boolean;
}

// UsageFlags are bit flags indicating how a model/polymorphic type is used
export enum UsageFlags {
  // the type is unreferenced
  None = 0,

  // the type is received over the wire
  Input = 1,

  // the type is sent over the wire
  Output = 2
}

// PolymorphicType is a discriminated type
export interface PolymorphicType extends StructType {
  fields: Array<ModelField>;

  annotations: ModelAnnotations;

  usage: UsageFlags;

  // this denotes the polymorphic interface this type implements
  interface: InterfaceType;

  // the value in the JSON that indicates what type was sent over the wire (e.g. goblin, salmon, shark)
  // note that for "root" types (Fish), there is no discriminatorValue. however, "sub-root" types (e.g. Salmon)
  // will have this populated.
  discriminatorValue?: LiteralValue;
}

// PossibleType describes what can be modeled e.g. in an OpenAPI specification
export type PossibleType = BytesType | ConstantType | InterfaceType | LiteralValue | MapType | ModelType | PolymorphicType | PrimitiveType | SliceType | QualifiedType | TimeType;

// StructField describes a field definition within a struct
export interface StructField {
  name: string;

  docs: Docs;

  type: PossibleType;

  byValue: boolean;
}

// ModelField describes a field within a model
export interface ModelField extends StructField {
  serializedName: string;

  annotations: ModelFieldAnnotations;

  // the value to send over the wire if one isn't specified
  defaultValue?: LiteralValue;

  xml?: XMLInfo;
}

export interface ModelFieldAnnotations {
  required: boolean;

  readOnly: boolean;

  isAdditionalProperties: boolean;

  isDiscriminator: boolean;
}

// ConstantTypeTypes contains the possible underlying type of a const
export type ConstantTypeTypes = 'bool' | 'float32' | 'float64' | 'int32' | 'int64' | 'string';

// Constant describes a const type definition (e.g. type FooType string, i.e. our fake enums)
export interface ConstantType {
  name: string;

  docs: Docs;

  type: ConstantTypeTypes;

  values: Array<ConstantValue>;

  valuesFuncName: string;
}

export type ConstantValueValueTypes = boolean | number | string;

// ConstantValue describes a const value definition (e.g. FooTypeValue FooType = "value")
export interface ConstantValue {
  name: string;

  docs: Docs;

  type: ConstantType;

  value: ConstantValueValueTypes;
}

export type PrimitiveTypeName = 'any' | 'bool' | 'byte' | 'float32' | 'float64' | 'int8' | 'int16' | 'int32' | 'int64' | 'rune' | 'string' | 'uint8' | 'uint16' | 'uint32' | 'uint64';

export type BytesEncoding = 'Std' | 'URL';

// BytesType is a base-64 encoded sequence of bytes
export interface BytesType {
  encoding: BytesEncoding;
}

// PrimitiveType is a Go integral type
export interface PrimitiveType {
  typeName: PrimitiveTypeName;
  encodeAsString: boolean;
}

export type LiteralValueType = BytesType | ConstantType | PrimitiveType | TimeType;

// LiteralValue is a literal value (e.g. "foo").
export interface LiteralValue {
  type: LiteralValueType;

  literal: any;
}

// QualifiedType is a type from some package, e.g. the Go standard library (excluding time.Time)
export interface QualifiedType {
  // this is the type name minus any package qualifier (e.g. URL)
  exportName: string;

  // the full name of the package to import (e.g. "net/url")
  packageName: string;
}

export type DateTimeFormat = 'dateType' | 'dateTimeRFC1123' | 'dateTimeRFC3339' | 'timeRFC3339' | 'timeUnix';

// TimeType is a time.Time type from the standard library with a format specifier.
export interface TimeType {
  dateTimeFormat: DateTimeFormat;

  utc: boolean;
}

export type MapValueType = PossibleType;

export interface MapType {
  valueType: MapValueType;

  valueTypeByValue: boolean;
}

export type SliceElementType = PossibleType;

export interface SliceType {
  elementType: SliceElementType;

  elementTypeByValue: boolean;

  // this slice is bytes of raw JSON
  rawJSONAsBytes: boolean;
}

// InterfaceType represents the interface type for a polymorphic (discriminated) type
export interface InterfaceType {
  // Name is the name of the interface (e.g. FishClassification)
  name: string;

  docs: Docs;

  // contains possible concrete type instances (e.g. Flounder, Carp)
  possibleTypes: Array<PolymorphicType>;

  // contains the name of the discriminator field in the JSON (e.g. "fishtype")
  discriminatorField: string;

  // does this polymorphic type have a parent (e.g. SalmonClassification has parent FishClassification)
  parent?: InterfaceType;

  // this is the "root" type in the list of polymorphic types (e.g. Fish for FishClassification)
  rootType: PolymorphicType;
}

export interface XMLInfo {
  name?: string;

  // name propagated to the generated wrapper type
  wrapper? :string;

  // slices only. this is the name of the wrapped type
  wraps?: string;

  attribute: boolean;

  text: boolean;
}

export function isBytesType(type: PossibleType): type is BytesType {
  return (<BytesType>type).encoding !== undefined;
}

export function isConstantType(type: PossibleType): type is ConstantType {
  return (<ConstantType>type).values !== undefined;
}

export function isLiteralValueType(type: PossibleType): type is LiteralValueType {
  return isConstantType(type) || isPrimitiveType(type);
}

export function isPrimitiveType(type: PossibleType): type is PrimitiveType {
  return (<PrimitiveType>type).typeName !== undefined;
}

export function isQualifiedType(type: PossibleType): type is QualifiedType {
  return (<QualifiedType>type).exportName !== undefined;
}

export function isTimeType(type: PossibleType): type is TimeType {
  return (<TimeType>type).dateTimeFormat !== undefined;
}

export function isMapType(type: PossibleType): type is MapType {
  return (<MapType>type).valueType !== undefined;
}

export function isModelType(type: PossibleType): type is ModelType {
  return (<ModelType>type).fields !== undefined;
}

export function isPolymorphicType(type: PossibleType): type is PolymorphicType {
  return (<PolymorphicType>type).interface !== undefined;
}

export function isSliceType(type: PossibleType): type is SliceType {
  return (<SliceType>type).elementType !== undefined;
}

export function isInterfaceType(type: PossibleType): type is InterfaceType {
  return (<InterfaceType>type).possibleTypes !== undefined;
}

export function isLiteralValue(type: PossibleType): type is LiteralValue {
  return (<LiteralValue>type).literal !== undefined;
}

export function getLiteralValueTypeName(literal: LiteralValueType): string {
  if (isBytesType(literal)) {
    return '[]byte';
  } else if (isConstantType(literal)) {
    return literal.name;
  } else if (isPrimitiveType(literal)) {
    return literal.typeName;
  } else if (isTimeType(literal)) {
    return 'time.Time';
  } else {
    throw new Error(`unhandled LiteralValueType ${getTypeDeclaration(literal)}`);
  }
}

export function getTypeDeclaration(type: PossibleType, pkgName?: string): string {
  if (isPrimitiveType(type)) {
    return type.typeName;
  } else if (isQualifiedType(type)) {
    let pkg = type.packageName;
    const pathChar = pkg.lastIndexOf('/');
    if (pathChar) {
      pkg = pkg.substring(pathChar+1);
    }
    return pkg + '.' + type.exportName;
  } else if (isConstantType(type) || isInterfaceType(type) || isModelType(type) || isPolymorphicType(type)) {
    if (pkgName) {
      return `${pkgName}.${type.name}`;
    }
    return type.name;
  } else if (isBytesType(type)) {
    return '[]byte';
  } else if (isLiteralValue(type)) {
    return getTypeDeclaration(type.type, pkgName);
  } else if (isMapType(type)) {
    let pointer = '*';
    if (type.valueTypeByValue) {
      pointer = '';
    }
    return `map[string]${pointer}` + getTypeDeclaration(type.valueType, pkgName);
  } else if (isSliceType(type)) {
    let pointer = '*';
    if (type.elementTypeByValue) {
      pointer = '';
    }
    return `[]${pointer}` + getTypeDeclaration(type.elementType, pkgName);
  } else if (isTimeType(type)) {
    return 'time.Time';
  } else {
    throw new Error(`unhandled type ${typeof(type)}`);
  }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// base types
///////////////////////////////////////////////////////////////////////////////////////////////////

export class StructField implements StructField {
  constructor(name: string, type: PossibleType, byValue: boolean) {
    this.name = name;
    this.type = type;
    this.byValue = byValue;
    this.docs = {};
  }
}

export class StructType implements StructType {
  constructor(name: string) {
    this.fields = new Array<StructField>();
    this.name = name;
    this.docs = {};
  }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

export class ConstantType implements ConstantType {
  constructor(name: string, type: ConstantTypeTypes, valuesFuncName: string) {
    this.name = name;
    this.type = type;
    this.values = new Array<ConstantValue>();
    this.valuesFuncName = valuesFuncName;
    this.docs = {};
  }
}

export class ConstantValue implements ConstantValue {
  constructor(name: string, type: ConstantType, value: ConstantValueValueTypes) {
    this.name = name;
    this.type = type;
    this.value = value;
    this.docs = {};
  }
}

export class LiteralValue implements LiteralValue {
  constructor(type: LiteralValueType, literal: any) {
    this.type = type;
    this.literal = literal;
  }
}

export class ModelType extends StructType implements ModelType {
  constructor(name: string, annotations: ModelAnnotations, usage: UsageFlags) {
    super(name);
    this.annotations = annotations;
    this.usage = usage;
    this.fields = new Array<ModelField>();
  }
}

export class ModelAnnotations implements ModelAnnotations {
  constructor(omitSerDe: boolean, multipartForm: boolean) {
    this.omitSerDeMethods = omitSerDe;
    this.multipartFormData = multipartForm;
  }
}

export class ModelField extends StructField implements ModelField {
  constructor(name: string, type: PossibleType, byValue: boolean, serializedName: string, annotations: ModelFieldAnnotations) {
    super(name, type, byValue);
    this.serializedName = serializedName;
    this.annotations = annotations;
  }
}

export class ModelFieldAnnotations implements ModelFieldAnnotations {
  constructor(required: boolean, readOnly: boolean, isAddlProps: boolean, isDiscriminator: boolean) {
    this.required = required;
    this.readOnly = readOnly;
    this.isAdditionalProperties = isAddlProps;
    this.isDiscriminator = isDiscriminator;
  }
}

export class PolymorphicType extends StructType implements PolymorphicType {
  constructor(name: string, iface: InterfaceType, annotations: ModelAnnotations, usage: UsageFlags) {
    super(name);
    this.interface = iface;
    this.annotations = annotations;
    this.usage = usage;
    this.fields = new Array<ModelField>();
  }
}

export class InterfaceType implements InterfaceType {
  // possibleTypes and rootType are required. however, we have a chicken-and-egg
  // problem as creating a PolymorphicType requires the necessary InterfaceType.
  // so these fields MUST be populated after creating the InterfaceType.
  constructor(name: string, discriminatorField: string) {
    this.name = name;
    this.discriminatorField = discriminatorField;
    this.possibleTypes = new Array<PolymorphicType>();
    this.docs = {};
  }
}

export class BytesType implements BytesType {
  constructor(encoding: BytesEncoding) {
    this.encoding = encoding;
  }
}

export class PrimitiveType implements PrimitiveType {
  constructor(typeName: PrimitiveTypeName, encodeAsString?: boolean) {
    this.typeName = typeName;
    this.encodeAsString = encodeAsString ?? false;
  }
}

export class QualifiedType implements QualifiedType {
  constructor(exportName: string, packageName: string) {
    this.exportName = exportName;
    this.packageName = packageName;
  }
}

export class MapType implements MapType {
  constructor(valueType: MapValueType, valueTypeByValue: boolean) {
    this.valueType = valueType;
    this.valueTypeByValue = valueTypeByValue;
  }
}

export class SliceType implements SliceType {
  constructor(elementType: SliceElementType, elementTypeByValue: boolean) {
    this.elementType = elementType;
    this.elementTypeByValue = elementTypeByValue;
  }
}

export class TimeType implements TimeType {
  constructor(format: DateTimeFormat, utc: boolean) {
    this.dateTimeFormat = format;
    this.utc = utc;
  }
}

export class XMLInfo implements XMLInfo {
  constructor() {
    this.attribute = false;
    this.text = false;
  }
}
