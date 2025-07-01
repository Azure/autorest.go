/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import { CodeModelError } from "./errors.js";

/** Docs contains the values used in doc comment generation. */
export interface Docs {
  /** the high level summary */
  summary?: string;

  /** detailed description */
  description?: string;
}

/** defines types that go across the wire */
export type PossibleType = Any | Constant | EncodedBytes | Interface | Literal | Map | Model | PolymorphicModel | QualifiedType | RawJSON | Scalar | Slice | String | Time;

/** the Go any type */
export interface Any {
  isAny: true;
}

/** a const type definition */
export interface Constant {
  name: string;

  docs: Docs;

  type: ConstantType;

  values: Array<ConstantValue>;

  valuesFuncName: string;
}

/** the underlying type of a const */
export type ConstantType = 'bool' | 'float32' | 'float64' | 'int32' | 'int64' | 'string';

/** a const vale definition */
export interface ConstantValue {
  name: string;

  docs: Docs;

  type: Constant;

  value: ConstantValueType;
}

export type ConstantValueType = boolean | number | string;

/** a byte slice that's base64 encoded */
export interface EncodedBytes {
  /** indicates what kind of base64-encoding to use */
  encoding: BytesEncoding;
}

export type BytesEncoding = 'Std' | 'URL';

// InterfaceType represents the interface type for a polymorphic (discriminated) type
export interface Interface {
  // Name is the name of the interface (e.g. FishClassification)
  name: string;

  docs: Docs;

  // contains possible concrete type instances (e.g. Flounder, Carp)
  possibleTypes: Array<PolymorphicModel>;

  // contains the name of the discriminator field in the JSON (e.g. "fishtype")
  discriminatorField: string;

  // does this polymorphic type have a parent (e.g. SalmonClassification has parent FishClassification)
  parent?: Interface;

  // this is the "root" type in the list of polymorphic types (e.g. Fish for FishClassification)
  rootType: PolymorphicModel;
}

// LiteralValue is a literal value (e.g. "foo").
export interface Literal {
  type: LiteralType;

  /* eslint-disable-next-line @typescript-eslint/no-explicit-any */
  literal: any;
}

export type LiteralType = Constant | EncodedBytes | Scalar | String | Time;

export interface Map {
  valueType: MapValueType;

  valueTypeByValue: boolean;
}

export type MapValueType = PossibleType;

// ModelField describes a field within a model
export interface ModelField extends StructField {
  serializedName: string;

  annotations: ModelFieldAnnotations;

  // the value to send over the wire if one isn't specified
  defaultValue?: Literal;

  xml?: XMLInfo;
}

export interface ModelAnnotations {
  omitSerDeMethods: boolean;

  // indicates the model should be converted into multipart/form data
  multipartFormData: boolean;
}

export interface ModelFieldAnnotations {
  required: boolean;

  readOnly: boolean;

  isAdditionalProperties: boolean;

  isDiscriminator: boolean;
}

// ModelType is a struct that participates in serialization over the wire.
export interface Model extends Struct {
  fields: Array<ModelField>;

  annotations: ModelAnnotations;

  usage: UsageFlags;

  xml?: XMLInfo;
}

// PolymorphicType is a discriminated type
export interface PolymorphicModel extends Model {
  // this denotes the polymorphic interface this type implements
  interface: Interface;

  // the value in the JSON that indicates what type was sent over the wire (e.g. goblin, salmon, shark)
  // note that for "root" types (Fish), there is no discriminatorValue. however, "sub-root" types (e.g. Salmon)
  // will have this populated.
  discriminatorValue?: Literal;
}

/** a Go scalar type */
export interface Scalar {
  typeName: ScalarType;
  encodeAsString: boolean;
}

/** the supported Go scalar types */
export type ScalarType = 'bool' | 'byte' | 'float32' | 'float64' | 'int8' | 'int16' | 'int32' | 'int64' | 'rune' | 'uint8' | 'uint16' | 'uint32' | 'uint64';

/** a Go string */
export interface String {
  isString: true;
}

// QualifiedType is a type from some package, e.g. the Go standard library (excluding time.Time)
export interface QualifiedType {
  // this is the type name minus any package qualifier (e.g. URL)
  exportName: string;

  // the full name of the package to import (e.g. "net/url")
  packageName: string;
}

/** a byte slice containing raw JSON */
export interface RawJSON {
  rawJSON: true;
}

export interface Slice {
  elementType: SliceElementType;

  elementTypeByValue: boolean;
}

export type SliceElementType = PossibleType;

// Struct describes a vanilla struct definition (pretty much exclusively used for parameter groups/options bag types)
// UDTs that are sent/received are modeled as ModelType.
export interface Struct {
  name: string;

  docs: Docs;

  // there are only a few corner-cases where a struct has no fields
  fields: Array<StructField>;
}

// StructField describes a field definition within a struct
export interface StructField {
  name: string;

  docs: Docs;

  type: PossibleType;

  byValue: boolean;
}

// TimeType is a time.Time type from the standard library with a format specifier.
export interface Time {
  format: TimeFormat;

  utc: boolean;
}

export type TimeFormat = 'dateType' | 'dateTimeRFC1123' | 'dateTimeRFC3339' | 'timeRFC3339' | 'timeUnix';

// UsageFlags are bit flags indicating how a model/polymorphic type is used
export enum UsageFlags {
  // the type is unreferenced
  None = 0,

  // the type is received over the wire
  Input = 1,

  // the type is sent over the wire
  Output = 2
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

export function isAnyType(type: PossibleType): type is Any {
  return (<Any>type).isAny !== undefined;
}

export function isBytesType(type: PossibleType): type is EncodedBytes {
  return (<EncodedBytes>type).encoding !== undefined;
}

export function isConstantType(type: PossibleType): type is Constant {
  return (<Constant>type).values !== undefined;
}

export function isLiteralValueType(type: PossibleType): type is LiteralType {
  return isConstantType(type) || isPrimitiveType(type) || isStringType(type);
}

export function isPrimitiveType(type: PossibleType): type is Scalar {
  return (<Scalar>type).typeName !== undefined;
}

export function isQualifiedType(type: PossibleType): type is QualifiedType {
  return (<QualifiedType>type).exportName !== undefined;
}

export function isRawJSON(type: PossibleType): type is RawJSON {
  return (<RawJSON>type).rawJSON !== undefined;
}

export function isStringType(type: PossibleType): type is String {
  return (<String>type).isString !== undefined;
}

export function isTimeType(type: PossibleType): type is Time {
  return (<Time>type).format !== undefined;
}

export function isMapType(type: PossibleType): type is Map {
  return (<Map>type).valueType !== undefined;
}

export function isModelType(type: PossibleType): type is Model {
  return (<Model>type).fields !== undefined;
}

export function isPolymorphicType(type: PossibleType): type is PolymorphicModel {
  return (<PolymorphicModel>type).interface !== undefined;
}

export function isSliceType(type: PossibleType): type is Slice {
  return (<Slice>type).elementType !== undefined;
}

export function isInterfaceType(type: PossibleType): type is Interface {
  return (<Interface>type).possibleTypes !== undefined;
}

export function isLiteralValue(type: PossibleType): type is Literal {
  return (<Literal>type).literal !== undefined;
}

export function getLiteralValueTypeName(literal: LiteralType): string {
  if (isBytesType(literal)) {
    return '[]byte';
  } else if (isConstantType(literal)) {
    return literal.name;
  } else if (isPrimitiveType(literal)) {
    return literal.typeName;
  } else if (isStringType(literal)) {
    return 'string';
  } else if (isTimeType(literal)) {
    return 'time.Time';
  } else {
    throw new CodeModelError(`unhandled LiteralValueType ${getTypeDeclaration(literal)}`);
  }
}

export function getTypeDeclaration(type: PossibleType, pkgName?: string): string {
  if (isAnyType(type)) {
    return 'any';
  } else if (isPrimitiveType(type)) {
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
  } else if (isBytesType(type) || isRawJSON(type)) {
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
  } else if (isStringType(type)) {
    return 'string';
  } else if (isTimeType(type)) {
    return 'time.Time';
  } else {
    throw new CodeModelError(`unhandled type ${typeof(type)}`);
  }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

// Base classes first (StructField and StructType are base classes)
export class StructField implements StructField {
  constructor(name: string, type: PossibleType, byValue: boolean) {
    this.name = name;
    this.type = type;
    this.byValue = byValue;
    this.docs = {};
  }
}

export class Struct implements Struct {
  constructor(name: string) {
    this.fields = new Array<StructField>();
    this.name = name;
    this.docs = {};
  }
}

export class Any implements Any {
  constructor() {
    this.isAny = true;
  }
}

export class Constant implements Constant {
  constructor(name: string, type: ConstantType, valuesFuncName: string) {
    this.name = name;
    this.type = type;
    this.values = new Array<ConstantValue>();
    this.valuesFuncName = valuesFuncName;
    this.docs = {};
  }
}

export class ConstantValue implements ConstantValue {
  constructor(name: string, type: Constant, value: ConstantValueType) {
    this.name = name;
    this.type = type;
    this.value = value;
    this.docs = {};
  }
}

export class EncodedBytes implements EncodedBytes {
  constructor(encoding: BytesEncoding) {
    this.encoding = encoding;
  }
}

export class Interface implements Interface {
  // possibleTypes and rootType are required. however, we have a chicken-and-egg
  // problem as creating a PolymorphicType requires the necessary InterfaceType.
  // so these fields MUST be populated after creating the InterfaceType.
  constructor(name: string, discriminatorField: string) {
    this.name = name;
    this.discriminatorField = discriminatorField;
    this.possibleTypes = new Array<PolymorphicModel>();
    this.docs = {};
  }
}

export class Literal implements Literal {
  /* eslint-disable-next-line @typescript-eslint/no-explicit-any */
  constructor(type: LiteralType, literal: any) {
    this.type = type;
    /* eslint-disable-next-line @typescript-eslint/no-unsafe-assignment */
    this.literal = literal;
  }
}

export class Map implements Map {
  constructor(valueType: MapValueType, valueTypeByValue: boolean) {
    this.valueType = valueType;
    this.valueTypeByValue = valueTypeByValue;
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

export class Model extends Struct implements Model {
  constructor(name: string, annotations: ModelAnnotations, usage: UsageFlags) {
    super(name);
    this.annotations = annotations;
    this.usage = usage;
    this.fields = new Array<ModelField>();
  }
}

export class PolymorphicModel extends Model implements PolymorphicModel {
  constructor(name: string, iface: Interface, annotations: ModelAnnotations, usage: UsageFlags) {
    super(name, annotations, usage);
    this.interface = iface;
  }
}

export class RawJSON implements RawJSON {
  constructor() {
    this.rawJSON = true;
  }
}

export class Scalar implements Scalar {
  constructor(typeName: ScalarType, encodeAsString?: boolean) {
    this.typeName = typeName;
    this.encodeAsString = encodeAsString ?? false;
  }
}

export class String implements String {
  constructor() {
    this.isString = true;
  }
}

export class QualifiedType implements QualifiedType {
  constructor(exportName: string, packageName: string) {
    this.exportName = exportName;
    this.packageName = packageName;
  }
}

export class Slice implements Slice {
  constructor(elementType: SliceElementType, elementTypeByValue: boolean) {
    this.elementType = elementType;
    this.elementTypeByValue = elementTypeByValue;
  }
}

export class Time implements Time {
  constructor(format: TimeFormat, utc: boolean) {
    this.format = format;
    this.utc = utc;
  }
}

export class XMLInfo implements XMLInfo {
  constructor() {
    this.attribute = false;
    this.text = false;
  }
}
