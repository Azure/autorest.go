/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

/** Docs contains the values used in doc comment generation. */
export interface Docs {
  /** the high level summary */
  summary?: string;

  /** detailed description */
  description?: string;
}

/** defines types used in generated code but do not go across the wire */
export type SdkType = ArmClientOptions;

/** defines types that go across the wire */
export type WireType = ArmClientOptions | Any | Constant | EncodedBytes | ETag | Interface | Literal | Map | Model | MultipartContent | PolymorphicModel | RawJSON | ReadCloser | ReadSeekCloser | Scalar | Slice | String | Time;

/** defines a type within the Go type system */
export type Type = SdkType | WireType;

/** the Go any type */
export interface Any {
  kind: 'any';
}

/** an arm.ClientOptions type from azcore */
export interface ArmClientOptions extends QualifiedType {
  kind: 'armClientOptions';
}

/** a const type definition */
export interface Constant {
  kind: 'constant';

  /** the const type name */
  name: string;

  /** any docs for the const type */
  docs: Docs;

  /** the underlying type of the const */
  type: ConstantType;

  /** the possible values for this const */
  values: Array<ConstantValue>;

  /** the name of the func that returns the set of values */
  valuesFuncName: string;
}

/** the underlying type of a const */
export type ConstantType = 'bool' | 'float32' | 'float64' | 'int32' | 'int64' | 'string';

/** a const value definition */
export interface ConstantValue {
  kind: 'constantValue';

  /** the const value name */
  name: string;

  /** any docs for the const value */
  docs: Docs;

  /** the const to which this value belongs */
  type: Constant;

  /** the value for this const */
  value: ConstantValueType;
}

/** the underlying type of a const value */
export type ConstantValueType = boolean | number | string;

/** a byte slice that's base64 encoded */
export interface EncodedBytes {
  kind: 'encodedBytes';

  /** indicates what kind of base64-encoding to use */
  encoding: BytesEncoding;
}

/** the types of base64 encoding */
export type BytesEncoding = 'Std' | 'URL';

/** an azcore.ETag type */
export interface ETag extends QualifiedType {
  kind: 'etag';
}

/** a Go interface type used for discriminated types */
export interface Interface {
  kind: 'interface';

  /** the name of the interface (e.g. FishClassification) */
  name: string;

  /** any docs for the interface */
  docs: Docs;

  /** contains possible concrete type instances (e.g. Flounder, Carp) */
  possibleTypes: Array<PolymorphicModel>;

  /** contains the name of the discriminator field in the JSON (e.g. "fishtype") */
  discriminatorField: string;

  /** does this polymorphic type have a parent (e.g. SalmonClassification has parent FishClassification) */
  parent?: Interface;

  /**  this is the "root" type in the list of polymorphic types (e.g. Fish for FishClassification) */
  rootType: PolymorphicModel;
}

/** a literal value (e.g. "foo", 123, true) */
export interface Literal {
  kind: 'literal';

  /** the literal's underlying type */
  type: LiteralType;

  /* eslint-disable-next-line @typescript-eslint/no-explicit-any */
  /** the value for this literal */
  literal: any;
}

/** the possible types of literals */
export type LiteralType = Constant | EncodedBytes | Scalar | String | Time;

/** a Go map type. note that the key is always a string */
export interface Map {
  kind: 'map';

  /** the type of values in the map */
  valueType: MapValueType;

  /** indicates if the map's value type is pointer-to-type or not */
  valueTypeByValue: boolean;
}

/** the set of map value types */
export type MapValueType = WireType;

/** a field within a model */
export interface ModelField extends StructField {
  /** the name of the field as it's sent/received over the wire */
  serializedName: string;

  /** metadata for this field */
  annotations: ModelFieldAnnotations;

  /** the value to send over the wire if one isn't specified */
  defaultValue?: Literal;

  /** any XML metadata */
  xml?: XMLInfo;
}

/** additional settings for a model type */
export interface ModelAnnotations {
  /** when true, serde methods will not be generated */
  omitSerDeMethods: boolean;

  /** indicates the model should be converted into multipart/form data */
  multipartFormData: boolean;
}

/** additional settings for a model field */
export interface ModelFieldAnnotations {
  /** the field is required on input and will always be populated on output */
  required: boolean;

  /** the field is read-only and will be populated on output. any set value on input will be ignored */
  readOnly: boolean;

  /** field is JSON additional properties */
  isAdditionalProperties: boolean;

  /** field is the discriminator for a discriminated type */
  isDiscriminator: boolean;
}

/** a struct that participates in serialization over the wire */
export interface Model extends ModelBase {
  kind: 'model';
}

/** a streaming.MultipartContent type from azcore */
export interface MultipartContent extends QualifiedType {
  kind: 'multipartContent';
}

/** a model that's a discriminated type */
export interface PolymorphicModel extends ModelBase {
  kind: 'polymorphicModel';

  /** the polymorphic interface this type implements */
  interface: Interface;

  /**
   * the value in the JSON that indicates what type was sent over the wire (e.g. goblin, salmon, shark)
   * note that for "root" types (Fish), there is no discriminatorValue. however, "sub-root" types (e.g. Salmon)
   * will have this populated.
   */
  discriminatorValue?: Literal;
}

/** a byte slice containing raw JSON */
export interface RawJSON {
  kind: 'rawJSON';
}

/** a Go scalar type */
export interface Scalar {
  kind: 'scalar';

  /** the type of scalar */
  type: ScalarType;

  /** indicates the value is sent/received as a string */
  encodeAsString: boolean;
}

/** an io.ReadCloser */
export interface ReadCloser extends QualifiedType {
  kind: 'readCloser'
}

/** an io.ReadSeekCloser */
export interface ReadSeekCloser extends QualifiedType {
  kind: 'readSeekCloser'
}

/** the supported Go scalar types */
export type ScalarType = 'bool' | 'byte' | 'float32' | 'float64' | 'int8' | 'int16' | 'int32' | 'int64' | 'rune' | 'uint8' | 'uint16' | 'uint32' | 'uint64';

/** a Go slice */
export interface Slice {
  kind: 'slice';

  /** the element type for this slice */
  elementType: SliceElementType;

  /** indicates if the slice's element type is pointer-to-type or not */
  elementTypeByValue: boolean;
}

/** the set of slice element types */
export type SliceElementType = WireType;

/** a Go string */
export interface String {
  kind: 'string';
}

/** a vanilla struct definition (pretty much exclusively used for parameter groups/options bag types) */
export interface Struct {
  /** the name of the struct */
  name: string;

  /** and docs for this struct */
  docs: Docs;

  /** the fields in this struct. can be empty */
  fields: Array<StructField>;
}

/** a field definition within a struct */
export interface StructField {
  /** the name of the field */
  name: string;

  /** and docs for this field */
  docs: Docs;

  /** the field's underlying type */
  type: WireType;

  /** indicates if the field is pointer-to-type or not */
  byValue: boolean;
}

/** a time.Time type from the standard library with a format specifier */
export interface Time extends QualifiedType {
  kind: 'time';

  /** the serde format used */
  format: TimeFormat;

  /** indicates if the time is always in UTC */
  utc: boolean;
}

/** the set of time serde formats */
export type TimeFormat = 'dateType' | 'dateTimeRFC1123' | 'dateTimeRFC3339' | 'timeRFC3339' | 'timeUnix';

/** bit flags indicating how a model/polymorphic type is used */
export enum UsageFlags {
  /** the type is unreferenced */
  None = 0,

  /** the type is received over the wire */
  Input = 1,

  /** the type is sent over the wire */
  Output = 2
}

/** metadata used for XML serde */
export interface XMLInfo {
  /** element name to use instead of the default name */
  name?: string;

  /** name propagated to the generated wrapper type */
  wrapper? :string;

  /** slices only. this is the name of the wrapped type */
  wraps?: string;

  /** value is an XML attribute */
  attribute: boolean;

  /** value is raw text */
  text: boolean;
}

/**
 * returns the Go type declaration for the specified LiteralType
 * 
 * @param literal the type for which to emit the declaration
 * @returns the Go type declaration
 */
export function getLiteralTypeDeclaration(literal: LiteralType): string {
  switch (literal.kind) {
    case 'constant':
      return literal.name;
    case 'encodedBytes':
      return '[]byte';
    case 'scalar':
      return literal.type;
    case 'string':
      return literal.kind;
    case 'time':
      return 'time.Time';
  }
}

/**
 * returns the Go type declaration for the specified type.
 * any value in pkgName is prefixed to the underlying type name.
 * 
 * @param type the type for which to emit the declaration
 * @param pkgName optional package name prefix for the type
 * @returns the Go type declaration
 */
export function getTypeDeclaration(type: Type, pkgName?: string): string {
  switch (type.kind) {
    case 'any':
    case 'string':
      return type.kind;
    case 'constant':
    case 'interface':
    case 'model':
    case 'polymorphicModel':
      if (pkgName) {
        return `${pkgName}.${type.name}`;
      }
      return type.name;
    case 'encodedBytes':
    case 'rawJSON':
      return '[]byte';
    case 'literal':
      return getTypeDeclaration(type.type, pkgName);
    case 'map':
      return `map[string]${type.valueTypeByValue ? '' : '*'}` + getTypeDeclaration(type.valueType, pkgName);
    case 'scalar':
      return type.type;
    case 'slice':
      return `[]${type.elementTypeByValue ? '' : '*'}` + getTypeDeclaration(type.elementType, pkgName);
    case 'time':
      return 'time.Time';
    case 'armClientOptions':
    case 'etag':
    case 'multipartContent':
    case 'readCloser':
    case 'readSeekCloser': {
      // strip module to just the leaf package as required
      let pkg = type.module;
      const pathChar = pkg.lastIndexOf('/');
      if (pathChar) {
        pkg = pkg.substring(pathChar+1);
      }
      return pkg + '.' + type.name;
    }
  }
}

/** narrows type to a LiteralType within the conditional block */
export function isLiteralValueType(type: WireType): type is LiteralType {
  switch (type.kind) {
    case 'constant':
    case 'encodedBytes':
    case 'scalar':
    case 'string':
    case 'time':
      return true;
    default:
      return false;
  }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// exported base types
///////////////////////////////////////////////////////////////////////////////////////////////////

export class StructField implements StructField {
  constructor(name: string, type: WireType, byValue: boolean) {
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

/** used when building types that come from an external package */
export interface QualifiedType {
  /** the type name minus any package qualifier (e.g. URL) */
  name: string;

  /** the full name of the module to import (e.g. "net/url") */
  module: string;
}

export class QualifiedType implements QualifiedType {
  constructor(name: string, module: string) {
    this.name = name;
    this.module = module;
  }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// base types
///////////////////////////////////////////////////////////////////////////////////////////////////

interface ModelBase extends Struct {
  /** the fields in this model. can be empty */
  fields: Array<ModelField>;

  /** any annotations for this model */
  annotations: ModelAnnotations;

  /** usage flags for this model */
  usage: UsageFlags;

  /** any XML metadata */
  xml?: XMLInfo;
}

class ModelBase extends Struct implements ModelBase {
  constructor(name: string, annotations: ModelAnnotations, usage: UsageFlags) {
    super(name);
    this.annotations = annotations;
    this.usage = usage;
    this.fields = new Array<ModelField>();
  }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

export class Any implements Any {
  constructor() {
    this.kind = 'any';
  }
}

export class ArmClientOptions extends QualifiedType implements ArmClientOptions {
  constructor() {
    super('ClientOptions', 'github.com/Azure/azure-sdk-for-go/sdk/azcore/arm');
    this.kind = 'armClientOptions';
  }
}

export class Constant implements Constant {
  constructor(name: string, type: ConstantType, valuesFuncName: string) {
    this.kind = 'constant';
    this.name = name;
    this.type = type;
    this.values = new Array<ConstantValue>();
    this.valuesFuncName = valuesFuncName;
    this.docs = {};
  }
}

export class ConstantValue implements ConstantValue {
  constructor(name: string, type: Constant, value: ConstantValueType) {
    this.kind = 'constantValue';
    this.name = name;
    this.type = type;
    this.value = value;
    this.docs = {};
  }
}

export class EncodedBytes implements EncodedBytes {
  constructor(encoding: BytesEncoding) {
    this.kind = 'encodedBytes';
    this.encoding = encoding;
  }
}

export class ETag extends QualifiedType implements ETag {
  constructor() {
    super('ETag', 'github.com/Azure/azure-sdk-for-go/sdk/azcore');
    this.kind = 'etag';
  }
}

export class Interface implements Interface {
  // WireTypes and rootType are required. however, we have a chicken-and-egg
  // problem as creating a PolymorphicType requires the necessary InterfaceType.
  // so these fields MUST be populated after creating the InterfaceType.
  constructor(name: string, discriminatorField: string) {
    this.kind = 'interface';
    this.name = name;
    this.discriminatorField = discriminatorField;
    this.possibleTypes = new Array<PolymorphicModel>();
    this.docs = {};
  }
}

export class Literal implements Literal {
  /* eslint-disable-next-line @typescript-eslint/no-explicit-any */
  constructor(type: LiteralType, literal: any) {
    this.kind = 'literal';
    this.type = type;
    /* eslint-disable-next-line @typescript-eslint/no-unsafe-assignment */
    this.literal = literal;
  }
}

export class Map implements Map {
  constructor(valueType: MapValueType, valueTypeByValue: boolean) {
    this.kind = 'map';
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
  constructor(name: string, type: WireType, byValue: boolean, serializedName: string, annotations: ModelFieldAnnotations) {
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

export class Model extends ModelBase implements Model {
  constructor(name: string, annotations: ModelAnnotations, usage: UsageFlags) {
    super(name, annotations, usage);
    this.kind = 'model';
    this.fields = new Array<ModelField>();
  }
}

export class MultipartContent extends QualifiedType implements MultipartContent {
  constructor() {
    super('MultipartContent', 'github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
    this.kind = 'multipartContent';
  }
}

export class PolymorphicModel extends ModelBase implements PolymorphicModel {
  constructor(name: string, iface: Interface, annotations: ModelAnnotations, usage: UsageFlags) {
    super(name, annotations, usage);
    this.kind = 'polymorphicModel';
    this.interface = iface;
  }
}

export class RawJSON implements RawJSON {
  constructor() {
    this.kind = 'rawJSON';
  }
}

export class ReadCloser extends QualifiedType implements ReadCloser {
  constructor() {
    super('ReadCloser', 'io');
    this.kind = 'readCloser';
  }
}

export class ReadSeekCloser extends QualifiedType implements ReadSeekCloser {
  constructor() {
    super('ReadSeekCloser', 'io');
    this.kind = 'readSeekCloser';
  }
}

export class Scalar implements Scalar {
  constructor(type: ScalarType, encodeAsString: boolean) {
    this.kind = 'scalar';
    this.type = type;
    this.encodeAsString = encodeAsString;
  }
}

export class Slice implements Slice {
  constructor(elementType: SliceElementType, elementTypeByValue: boolean) {
    this.kind = 'slice';
    this.elementType = elementType;
    this.elementTypeByValue = elementTypeByValue;
  }
}

export class String implements String {
  constructor() {
    this.kind = 'string';
  }
}

export class Time extends QualifiedType implements Time {
  constructor(format: TimeFormat, utc: boolean) {
    super('Time', 'time');
    this.kind = 'time';
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
