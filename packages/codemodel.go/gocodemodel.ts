/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

// CodeModel contains a Go-specific abstraction over an OpenAPI (or other) description of REST endpoints.
export interface CodeModel {
  info: Info;

  host?: string;

  type: CodeModelType;

  packageName: string;

  options: Options;

  // all of the struct model types to generate (models.go file)
  models: Array<ModelType | PolymorphicType>;

  // all of the const types to generate (constants.go file)
  constants: Array<ConstantType>;

  // all of the operation groups (i.e. clients and their methods)
  // no clients indicates a models-only build
  clients: Array<Client>;

  // all of the parameter groups including the options types (options.go file)
  paramGroups: Array<StructType>;

  // all of the response envelopes (responses.go file)
  // no response envelopes indicates a models-only build
  responseEnvelopes: Array<ResponseEnvelope>;

  // all of the interfaces for discriminated types (interfaces.go file)
  interfaceTypes: Array<InterfaceType>;

  marshallingRequirements: MarshallingRequirements;
}

export type CodeModelType = 'azure-arm' | 'data-plane';

// Info contains top-level info about the input source
export interface Info {
  title: string;
}

// Options contains global options set on the CodeModel.
export interface Options {
  headerText: string;

  generateFakes: boolean;

  injectSpans: boolean;

  // disallowUnknownFields indicates whether or not to disallow unknown fields in the JSON unmarshaller.
  // reproduce the behavior of https://pkg.go.dev/encoding/json#Decoder.DisallowUnknownFields
  disallowUnknownFields: boolean;

  // NOTE: containingModule and module are mutually exclusive

  // the module into which the package is being generated
  containingModule?: string;

  // the module being generated
  module?: string;

  moduleVersion?: string;

  azcoreVersion?: string;
}

// MarshallingRequirements contains flags for required marshalling helpers
export interface MarshallingRequirements {
  generateDateHelper: boolean;

  generateDateTimeRFC1123Helper: boolean;

  generateDateTimeRFC3339Helper: boolean;

  generateTimeRFC3339Helper: boolean;

  generateUnixTimeHelper: boolean;

  generateXMLDictionaryUnmarshallingHelper: boolean;
}

// Struct describes a vanilla struct definition (pretty much exclusively used for parameter groups/options bag types)
// UDTs that are sent/received are modeled as ModelType.
export interface StructType {
  name: string;

  description?: string;

  // there are only a few corner-cases where a struct has no fields
  fields: Array<StructField>;
}

// ModelFormat indicates what format a model is sent/received as.
export type ModelFormat = 'json' | 'xml';

// ModelType is a struct that participates in serialization over the wire.
export interface ModelType extends StructType {
  fields: Array<ModelField>;

  // format is propagated to models purely as a convenience when determining
  // what marshaller/unmarshaller to generate. technically, a model could
  // participate in both JSON and XML formats. this hasn't been a problem yet
  format: ModelFormat;

  annotations: ModelAnnotations;

  xml?: XMLInfo;
}

export interface ModelAnnotations {
  omitSerDeMethods: boolean;
}

// PolymorphicType is a discriminated type
export interface PolymorphicType extends StructType {
  fields: Array<ModelField>;

  format: 'json';

  annotations: ModelAnnotations;

  // this denotes the polymorphic interface this type implements
  interface: InterfaceType;

  // the value in the JSON that indicates what type was sent over the wire (e.g. goblin, salmon, shark)
  // note that for "root" types (Fish), there is no discriminatorValue. however, "sub-root" types (e.g. Salmon)
  // will have this populated.
  discriminatorValue?: string;
}

// PossibleType describes what can be modeled e.g. in an OpenAPI specification
export type PossibleType = BytesType | ConstantType | InterfaceType | LiteralValue | MapType | ModelType | PolymorphicType | PrimitiveType | SliceType | StandardType | TimeType;

// StructField describes a field definition within a struct
export interface StructField {
  fieldName: string;

  description?: string;

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

  description?: string;

  type: ConstantTypeTypes;

  values: Array<ConstantValue>;

  valuesFuncName: string;
}

export type ConstantValueValueTypes = boolean | number | string;

// ConstantValue describes a const value definition (e.g. FooTypeValue FooType = "value")
export interface ConstantValue {
  valueName: string;

  description?: string;

  type: ConstantType;

  value: ConstantValueValueTypes;
}

// Client is an SDK client
export interface Client {
  clientName: string;

  // groupName is the name of the operation group the client belongs to (e.g. "groupname" from an operation ID of "groupname_operation")
  groupName: string;

  // the name of the client's constructor func
  ctorName: string;

  // contains only modeled client parameters and it's legal for an operation to not have any modeled client parameters
  parameters: Array<Parameter>;

  methods: Array<Method | LROMethod | PageableMethod | LROPageableMethod>;

  // client has a statically defined host
  host?: string;

  // can be empty if there are no host params
  hostParams: Array<URIParameter>;

  // complexHostParams indicates that the parameters to construct the full host name
  // span the client and the method. see custombaseurlgroup for an example of this.
  complexHostParams: boolean;
}

// Method is a method on a client
export interface Method {
  methodName: string;

  description?: string;

  httpPath: string;

  httpMethod: HTTPMethod;

  // any modeled parameters. the ones we add to the generated code (context.Context etc) aren't included here
  parameters: Array<Parameter>;
  
  optionalParamsGroup: ParameterGroup;

  responseEnvelope: ResponseEnvelope;

  // the complete list of successful HTTP status codes
  httpStatusCodes: Array<number>;

  client: Client;

  naming: MethodNaming;

  apiVersions: Array<string>; // TODO: not sure why this needs to be an array
}

export type HTTPMethod = 'delete' | 'get' | 'head' | 'patch' | 'post' | 'put';

export interface MethodNaming {
  // this is the name of the internal method for consumption by LROs/paging methods.
  internalMethod: string;

  requestMethod: string;

  responseMethod: string;
}

export interface LROMethod extends Method {
  finalStateVia?: 'azure-async-operation' | 'location' | 'operation-location' | 'original-uri';

  isLRO: true;
}

export interface PageableMethod extends Method {
  nextLinkName?: string;

  nextPageMethod?: NextPageMethod;

  isPageable: true;
}

// NextPageMethod is the internal method used for fetching the next page for a PageableMethod.
// It's unique from a regular Method as it's not exported and has no optional params/response envelope.
// thus, it's not included in the array of methods for a client.
export interface NextPageMethod {
  methodName: string;
 
  httpPath: string;

  httpMethod: HTTPMethod;

  // any modeled parameters
  parameters: Array<Parameter>;

  // the complete list of successful HTTP status codes
  httpStatusCodes: Array<number>;

  client: Client;

  apiVersions: Array<string>; // TODO: not sure why this needs to be an array

  isNextPageMethod: true;
}

export interface LROPageableMethod extends LROMethod, PageableMethod {
  // no new fields are added
}

export function isMethod(method: Method | NextPageMethod): method is Method {
  return (<NextPageMethod>method).isNextPageMethod === undefined;
}

export function isLROMethod(method: Method | LROMethod | PageableMethod): method is LROMethod {
  return (<LROMethod>method).isLRO === true;
}

export function isPageableMethod(method: Method | LROMethod | PageableMethod): method is PageableMethod {
  return (<PageableMethod>method).isPageable === true;
}

// required - the parameter is required
// optional - the parameter is optional
// literal - there is no formal parameter, the value is emitted directly in the code (e.g. the Accept header parameter)
// flag - the value is a literal and emitted in the code, but sent IFF the flag param is not nil (i.e. an optional LiteralValue)
// ClientSideDefault - the parameter has an emitted default value that's sent if one isn't specified (implies optional)
export type ParameterType = 'required' | 'optional' | 'literal' | 'flag' | ClientSideDefault;

export interface ClientSideDefault {
  defaultValue: LiteralValue;
}

// Parameter is a parameter for a client method
export interface Parameter {
  paramName: string;

  description?: string;

  // NOTE: if the type is a LiteralValue the paramType will either be literal or flag
  type: PossibleType;

  // paramType will have value literal or flag when type is a LiteralValue (see above comment).
  paramType: ParameterType;

  byValue: boolean;

  group?: ParameterGroup;

  location: ParameterLocation;

  xml?: XMLInfo;
}

export function isClientSideDefault(paramType: ParameterType): paramType is ClientSideDefault {
  return (<ClientSideDefault>paramType).defaultValue !== undefined;
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

export type ParameterLocation = 'client' | 'method';

export interface ParameterGroup {
  // paramName is the name of the parameter
  paramName: string;

  description?: string;

  // groupName is the name of the param group (i.e. the struct name)
  groupName: string;

  required: boolean;

  location: ParameterLocation;

  // the params that belong to this group
  params: Array<Parameter>;
}

export type HeaderType = BytesType | ConstantType | PrimitiveType | TimeType | LiteralValue;

// parameter is sent via an HTTP header
export interface HeaderParameter extends Parameter {
  headerName: string;

  type: HeaderType;
}

export type CollectionFormat = 'csv' | 'ssv' | 'tsv' | 'pipes';

export interface HeaderCollectionParameter extends Parameter {
  headerName: string;

  type: SliceType;

  collectionFormat: CollectionFormat;
}

// this is a special type to support x-ms-header-collection-prefix (i.e. storage)
export interface HeaderMapParameter extends Parameter {
  headerName: string;

  type: MapType;

  collectionPrefix: string;
}

export type PathParameterType = BytesType | ConstantType | PrimitiveType | TimeType | LiteralValue;

// parameter is a segment in a path
export interface PathParameter extends Parameter {
  pathSegment: string;

  type: PathParameterType;

  isEncoded: boolean;
}

export interface PathCollectionParameter extends Parameter {
  pathSegment: string;

  type: SliceType;

  isEncoded: boolean;

  collectionFormat: CollectionFormat;
}

export type QueryParameterType = BytesType | ConstantType | PrimitiveType | TimeType | LiteralValue;

// parameter is sent via an HTTP query parameter
export interface QueryParameter extends Parameter {
  queryParameter: string;

  type: QueryParameterType;

  isEncoded: boolean;
}

export type ExtendedCollectionFormat = CollectionFormat | 'multi';

export interface QueryCollectionParameter extends Parameter {
  queryParameter: string;

  type: SliceType;

  isEncoded: boolean;

  collectionFormat: ExtendedCollectionFormat;
}

export type URIParameterType = ConstantType | PrimitiveType;

// parameter is a segment in the host
export interface URIParameter extends Parameter {
  uriPathSegment: string;

  type: URIParameterType;
}

export type BodyFormat = 'JSON' | 'XML' | 'Text' | 'binary';

// parameter is sent in the HTTP request body
export interface BodyParameter extends Parameter {
  bodyFormat: BodyFormat;

  // "application/text" etc...
  contentType: string;
}

export interface FormBodyParameter extends Parameter {
  formDataName: string;
}

export interface FormBodyCollectionParameter extends Parameter {
  formDataName: string;

  type: SliceType;

  collectionFormat: ExtendedCollectionFormat;
}

export interface MultipartFormBodyParameter extends Parameter {
  multipartForm: true;
}

export interface ResumeTokenParameter extends Parameter {
  isResumeToken: true;
}

export function isBodyParameter(param: Parameter): param is BodyParameter {
  return (<BodyParameter>param).bodyFormat !== undefined;
}

export function isFormBodyParameter(param: Parameter): param is FormBodyParameter {
  return (<FormBodyParameter>param).formDataName !== undefined;
}

export function isFormBodyCollectionParameter(param: Parameter): param is FormBodyCollectionParameter {
  return (<FormBodyCollectionParameter>param).formDataName !== undefined && (<FormBodyCollectionParameter>param).collectionFormat !== undefined;
}

export function isMultipartFormBodyParameter(param: Parameter): param is MultipartFormBodyParameter {
  return (<MultipartFormBodyParameter>param).multipartForm !== undefined;
}

export function isHeaderParameter(param: Parameter): param is HeaderParameter {
  return (<HeaderParameter>param).headerName !== undefined;
}

export function isHeaderCollectionParameter(param: Parameter): param is HeaderCollectionParameter {
  return (<HeaderCollectionParameter>param).headerName !== undefined && (<HeaderCollectionParameter>param).collectionFormat !== undefined;
}

export function isHeaderMapParameter(param: Parameter): param is HeaderMapParameter {
  return (<HeaderMapParameter>param).headerName !== undefined && (<HeaderMapParameter>param).collectionPrefix !== undefined;
}

export function isPathParameter(param: Parameter): param is PathParameter {
  return (<PathParameter>param).pathSegment !== undefined;
}

export function isPathCollectionParameter(param: Parameter): param is PathCollectionParameter {
  return (<PathCollectionParameter>param).pathSegment !== undefined && (<PathCollectionParameter>param).collectionFormat !== undefined;
}

export function isQueryParameter(param: Parameter): param is QueryParameter {
  return (<QueryParameter>param).queryParameter !== undefined;
}

export function isQueryCollectionParameter(param: Parameter): param is QueryCollectionParameter {
  return (<QueryCollectionParameter>param).queryParameter !== undefined && (<QueryCollectionParameter>param).collectionFormat !== undefined;
}

export function isURIParameter(param: Parameter): param is URIParameter {
  return (<URIParameter>param).uriPathSegment !== undefined;
}

export function isResumeTokenParameter(param: Parameter): param is ResumeTokenParameter {
  return (<ResumeTokenParameter>param).isResumeToken !== undefined;
}

export type ResultType = AnyResult | BinaryResult | HeadAsBooleanResult | MonomorphicResult | PolymorphicResult | ModelResult;

// ResponseEnvelope is the type returned from a client method
export interface ResponseEnvelope {
  name: string;

  description: string;

  // for operations that return no body (e.g. a 204) this will be undefined.
  result?: ResultType;

  // any modeled response headers
  headers: Array<HeaderResponse | HeaderMapResponse>;

  method: Method | LROMethod | PageableMethod | LROPageableMethod;
}

export type ResultFormat = 'JSON' | 'XML' | 'Text';

// AnyResult is for endpoints that return a different schema based on the HTTP status code.
export interface AnyResult {
  // the name of the field within the response envelope
  fieldName: string;

  description?: string;

  // maps an HTTP status code to a result type.
  // status codes that don't return a schema will be absent.
  httpStatusCodeType: Record<number, PossibleType>;

  // the format in which the result is returned
  format: ResultFormat;

  byValue: true;
}

// TODO: would this ever be anything else?
export type BinaryResultFormat = 'binary';

// BinaryResult is for responses that return the streaming response (i.e. the http.Response.Body)
export interface BinaryResult {
  // the name of the field within the response envelope
  fieldName: string;

  description?: string;

  binaryFormat: BinaryResultFormat;

  byValue: true;
}

// HeadAsBooleanResult is for responses to HTTP HEAD requests that treat the HTTP status code as success/failure
export interface HeadAsBooleanResult {
  // the name of the field within the response envelope
  fieldName: string;

  description?: string;

  headAsBoolean: true;

  byValue: true;
}

export type MonomorphicResultType = BytesType | ConstantType | MapType | PrimitiveType | SliceType | TimeType;

// MonomorphicResult includes scalar results (ints, bools) or maps/slices of scalars/InterfaceTypes/ModelTypes.
// maps/slices can be recursive and/or combinatorial (e.g. map[string][]*sometype)
export interface MonomorphicResult {
  // the name of the field within the response envelope
  fieldName: string;

  description?: string;

  monomorphicType: MonomorphicResultType;

  // the format in which the result is returned
  format: ResultFormat;

  byValue: boolean;

  xml?: XMLInfo;
}

// PolymorphicResult is for discriminated types.
// The type is anonymously embedded in the response envelope.
export interface PolymorphicResult {
  description?: string;

  interfaceType: InterfaceType;

  // the format in which the result is returned
  format: 'JSON';
}

// ModelResult is a standard schema response.
// The type is anonymously embedded in the response envelope.
export interface ModelResult {
  description?: string;

  modelType: ModelType;

  // the format in which the result is returned
  format: ResultFormat;
}

export function isAnyResult(resultType: ResultType): resultType is AnyResult {
  return (<AnyResult>resultType).httpStatusCodeType !== undefined;
}

export function isBinaryResult(resultType: ResultType): resultType is BinaryResult {
  return (<BinaryResult>resultType).binaryFormat !== undefined;
}

export function isHeadAsBooleanResult(resultType: ResultType): resultType is HeadAsBooleanResult {
  return (<HeadAsBooleanResult>resultType).headAsBoolean !== undefined;
}

export function isMonomorphicResult(resultType: ResultType): resultType is MonomorphicResult {
  return (<MonomorphicResult>resultType).monomorphicType !== undefined;
}

export function isPolymorphicResult(resultType: ResultType): resultType is PolymorphicResult {
  return (<PolymorphicResult>resultType).interfaceType !== undefined;
}

export function isModelResult(resultType: ResultType): resultType is ModelResult {
  return (<ModelResult>resultType).modelType !== undefined;
}

export function getResultPossibleType(resultType: ResultType): PossibleType {
  if (isAnyResult(resultType)) {
    return new PrimitiveType('any');
  } else if (isBinaryResult(resultType)) {
    return new StandardType('io.ReadCloser', 'io');
  } else if (isHeadAsBooleanResult(resultType)) {
    return new PrimitiveType('bool');
  } else if (isMonomorphicResult(resultType)) {
    return resultType.monomorphicType;
  } else if (isPolymorphicResult(resultType)) {
    return resultType.interfaceType;
  } else if (isModelResult(resultType)) {
    return resultType.modelType;
  } else {
    throw new Error(`unhandled result type ${resultType}`);
  }
}

export interface HeaderResponse {
  // the name of the field within the response envelope
  fieldName: string;

  description?: string;

  type: HeaderType;

  byValue: boolean;

  // the name of the header sent over the wire
  headerName: string;
}

// this is a special type to support x-ms-header-collection-prefix (i.e. storage)
export interface HeaderMapResponse {
  // the name of the field within the response envelope
  fieldName: string;

  description?: string;

  type: MapType;

  byValue: boolean;

  // the name of the header sent over the wire
  headerName: string;

  collectionPrefix: string;
}

export type PrimitiveTypeName = 'any' | 'bool' | 'byte' | 'float32' | 'float64' | 'int32' | 'int64' | 'rune' | 'string';

export type BytesEncoding = 'Std' | 'URL';

// BytesType is a base-64 encoded sequence of bytes
export interface BytesType {
  encoding: BytesEncoding;
}

// PrimitiveType is a Go integral type
export interface PrimitiveType {
  typeName: PrimitiveTypeName;
}

export type LiteralValueType = BytesType | ConstantType | PrimitiveType | TimeType;

// LiteralValue is a literal value (e.g. "foo").
export interface LiteralValue {
  type: LiteralValueType;

  literal: any;
}

// StandardType is a type from the Go standard library (excluding time.Time)
export interface StandardType {
  // this is the fully-qualified name (e.g. io.Reader)
  typeName: string;

  // the full name of the package to import (e.g. "net/url")
  packageName: string;
}

export type DateTimeFormat = 'dateType' | 'dateTimeRFC1123' | 'dateTimeRFC3339' | 'timeRFC3339' | 'timeUnix';

// TimeType is a time.Time type from the standard library with a format specifier.
export interface TimeType extends StandardType {
  typeName: 'time.Time';

  packageName: 'time';

  dateTimeFormat: DateTimeFormat;
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

  description?: string;

  // contains possible concrete type instances (e.g. Flounder, Carp)
  possibleTypes: Array<PolymorphicType>;

  // contains the name of the discriminator field in the JSON (e.g. "fishtype")
  discriminatorField: string;

  // does this polymorphic type have a parent (e.g. SalmonClassification has parent FishClassification)
  parent?: InterfaceType;

  // this is the "root" type in the list of polymorphic types (e.g. Fish for FishClassification)
  rootType: PolymorphicType;
}

export function isBytesType(type: PossibleType): type is BytesType {
  return (<BytesType>type).encoding !== undefined;
}

export function isConstantType(type: PossibleType): type is ConstantType {
  return (<ConstantType>type).values !== undefined;
}

export function isHeaderMapResponse(resp: HeaderResponse | HeaderMapResponse): resp is HeaderMapResponse {
  return (<HeaderMapResponse>resp).collectionPrefix !== undefined;
}

export function isLiteralValueType(type: PossibleType): type is LiteralValueType {
  return isConstantType(type) || isPrimitiveType(type);
}

export function isPrimitiveType(type: PossibleType): type is PrimitiveType {
  return (<PrimitiveType>type).typeName !== undefined && !isStandardType(type) && !isTimeType(type);
}

export function isStandardType(type: PossibleType): type is StandardType {
  return (<StandardType>type).packageName !== undefined && !isTimeType(type);
}

export function isTimeType(type: PossibleType): type is TimeType {
  return (<TimeType>type).dateTimeFormat !== undefined;
}

export function isMapType(type: PossibleType): type is MapType {
  return (<MapType>type).valueType !== undefined;
}

export function isModelType(type: PossibleType): type is ModelType {
  return (<ModelType>type).format !== undefined;
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
  } else {
    return literal.typeName;
  }
}

export function getTypeDeclaration(type: PossibleType, pkgName?: string): string {
  if (isPrimitiveType(type) || isStandardType(type)) {
    return type.typeName;
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
///////////////////////////////////////////////////////////////////////////////////////////////////

export class Info implements Info {
  constructor(title: string) {
    this.title = title;
  }
}

export class Options implements Options {
  constructor(headerText: string, generateFakes: boolean, injectSpans: boolean, disallowUnknownFields: boolean) {
    this.headerText = headerText;
    this.generateFakes = generateFakes;
    this.injectSpans = injectSpans;
    this.disallowUnknownFields = disallowUnknownFields;
  }
}

export class MarshallingRequirements implements MarshallingRequirements {
  constructor() {
    this.generateDateHelper = false;
    this.generateDateTimeRFC1123Helper = false;
    this.generateDateTimeRFC3339Helper = false;
    this.generateTimeRFC3339Helper = false;
    this.generateUnixTimeHelper = false;
    this.generateXMLDictionaryUnmarshallingHelper = false;
  }
}

export class CodeModel implements CodeModel {
  constructor(info: Info, type: CodeModelType, packageName: string, options: Options) {
    this.clients = new Array<Client>();
    this.constants = new Array<ConstantType>();
    this.info = info;
    this.interfaceTypes = new Array<InterfaceType>();
    this.marshallingRequirements = new MarshallingRequirements();
    this.models = new Array<ModelType | PolymorphicType>();
    this.options = options;
    this.packageName = packageName;
    this.paramGroups = new Array<StructType>();
    this.responseEnvelopes = new Array<ResponseEnvelope>();
    this.type = type;
  }

  sortContent() {
    const sortAscending = function(a: string, b: string): number {
      return a < b ? -1 : a > b ? 1 : 0;
    };

    this.constants.sort((a: ConstantType, b: ConstantType) => { return sortAscending(a.name, b.name); });
    for (const enm of this.constants) {
      enm.values.sort((a: ConstantValue, b: ConstantValue) => { return sortAscending(a.valueName, b.valueName); });
    }
  
    this.interfaceTypes.sort((a: InterfaceType, b: InterfaceType) => { return sortAscending(a.name, b.name); });
    for (const iface of this.interfaceTypes) {
      iface.possibleTypes.sort((a: PolymorphicType, b: PolymorphicType) => { return sortAscending(a.discriminatorValue!, b.discriminatorValue!); });
    }
  
    this.models.sort((a: ModelType | PolymorphicType, b: ModelType | PolymorphicType) => { return sortAscending(a.name, b.name); });
    for (const model of this.models) {
      model.fields.sort((a: ModelField, b: ModelField) => { return sortAscending(a.fieldName, b.fieldName); });
    }
  
    this.paramGroups.sort((a: StructType, b: StructType) => { return sortAscending(a.name, b.name); });
    for (const paramGroup of this.paramGroups) {
      paramGroup.fields.sort((a: StructField, b: StructField) => { return sortAscending(a.fieldName, b.fieldName); });
    }
  
    this.responseEnvelopes.sort((a: ResponseEnvelope, b: ResponseEnvelope) => { return sortAscending(a.name, b.name); });
    for (const respEnv of this.responseEnvelopes) {
      respEnv.headers.sort((a: HeaderResponse | HeaderMapResponse, b: HeaderResponse | HeaderMapResponse) => { return sortAscending(a.fieldName, b.fieldName); });
    }
  
    this.clients.sort((a: Client, b: Client) => { return sortAscending(a.clientName, b.clientName); });
    for (const client of this.clients) {
      client.methods.sort((a: Method, b: Method) => { return sortAscending(a.methodName, b.methodName); });
    }
  }
}

export class Client implements Client {
  constructor(name: string, groupName: string, ctorName: string) {
    this.clientName = name;
    this.complexHostParams = false;
    this.ctorName = ctorName;
    this.groupName = groupName;
    this.hostParams = new Array<URIParameter>();
    this.methods = new Array<Method>();
    this.parameters = new Array<Parameter>();
  }
}

export class ConstantType implements ConstantType {
  constructor(name: string, type: ConstantTypeTypes, valuesFuncName: string) {
    this.name = name;
    this.type = type;
    this.values = new Array<ConstantValue>();
    this.valuesFuncName = valuesFuncName;
  }
}

export class ConstantValue implements ConstantValue {
  constructor(valueName: string, type: ConstantType, value: ConstantValueValueTypes) {
    this.valueName = valueName;
    this.type = type;
    this.value = value;
  }
}

export class MethodNaming implements MethodNaming {
  constructor(internalMethod: string, requestMethod: string, responseMethod: string) {
    this.internalMethod = internalMethod;
    this.requestMethod = requestMethod;
    this.responseMethod = responseMethod;
  }
}

export class Method implements Method {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    if (statusCodes.length === 0) {
      throw new Error('statusCodes cannot be empty');
    }
    this.apiVersions = new Array<string>();
    this.client = client;
    this.httpMethod = httpMethod;
    this.httpPath = httpPath;
    this.httpStatusCodes = statusCodes;
    this.methodName = name;
    this.naming = naming;
    this.parameters = new Array<Parameter>();
  }
}

export class LROMethod implements LROMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    if (statusCodes.length === 0) {
      throw new Error('statusCodes cannot be empty');
    }
    this.apiVersions = new Array<string>();
    this.client = client;
    this.httpMethod = httpMethod;
    this.httpPath = httpPath;
    this.httpStatusCodes = statusCodes;
    this.isLRO = true;
    this.methodName = name;
    this.naming = naming;
    this.parameters = new Array<Parameter>();
  }
}

export class PageableMethod implements PageableMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    if (statusCodes.length === 0) {
      throw new Error('statusCodes cannot be empty');
    }
    this.apiVersions = new Array<string>();
    this.client = client;
    this.httpMethod = httpMethod;
    this.httpPath = httpPath;
    this.httpStatusCodes = statusCodes;
    this.isPageable = true;
    this.methodName = name;
    this.naming = naming;
    this.parameters = new Array<Parameter>();
  }
}

export class LROPageableMethod implements LROPageableMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    if (statusCodes.length === 0) {
      throw new Error('statusCodes cannot be empty');
    }
    this.apiVersions = new Array<string>();
    this.client = client;
    this.httpMethod = httpMethod;
    this.httpPath = httpPath;
    this.httpStatusCodes = statusCodes;
    this.isLRO = true;
    this.isPageable = true;
    this.methodName = name;
    this.naming = naming;
    this.parameters = new Array<Parameter>();
  }
}

export class NextPageMethod implements NextPageMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>) {
    if (statusCodes.length === 0) {
      throw new Error('statusCodes cannot be empty');
    }
    this.apiVersions = new Array<string>();
    this.client = client;
    this.httpMethod = httpMethod;
    this.httpPath = httpPath;
    this.httpStatusCodes = statusCodes;
    this.isNextPageMethod = true;
    this.methodName = name;
    this.parameters = new Array<Parameter>();
  }
}

export class StructType implements StructType {
  constructor(name: string) {
    this.fields = new Array<StructField>();
    this.name = name;
  }
}

export class StructField implements StructField {
  constructor(fieldName: string, type: PossibleType, byValue: boolean) {
    this.fieldName = fieldName;
    this.type = type;
    this.byValue = byValue;
  }
}

export class ModelType implements ModelType {
  constructor(name: string, format: ModelFormat, annotations: ModelAnnotations) {
    this.name = name;
    this.format = format;
    this.annotations = annotations;
  }
}

export class ModelAnnotations implements ModelAnnotations {
  constructor(omitSerDe: boolean) {
    this.omitSerDeMethods = omitSerDe;
  }
}

export class ModelField implements ModelField {
  constructor(name: string, type: PossibleType, byValue: boolean, serializedName: string, annotations: ModelFieldAnnotations) {
    this.fieldName = name;
    this.type = type;
    this.byValue = byValue;
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

export class PolymorphicType implements PolymorphicType {
  constructor(name: string, iface: InterfaceType, annotations: ModelAnnotations) {
    this.name = name;
    this.interface = iface;
    this.annotations = annotations;
  }
}

export class InterfaceType implements InterfaceType {
  // possibleTypes and rootType are required. however, we have a chicken-and-egg
  // problem as creating a PolymorphicType requires the necessary InterfaceType.
  // so these fields MUST be populated after creating the InterfaceType.
  constructor(name: string, discriminatorField: string) {
    this.name = name;
    this.discriminatorField = discriminatorField;
    //this.possibleTypes = possibleTypes;
    //this.rootType = rootType;
  }
}

export class BytesType implements BytesType {
  constructor(encoding: BytesEncoding) {
    this.encoding = encoding;
  }
}

export class PrimitiveType implements PrimitiveType {
  constructor(typeName: PrimitiveTypeName) {
    this.typeName = typeName;
  }
}

export class StandardType implements StandardType {
  constructor(typeName: string, packageName: string) {
    this.typeName = typeName;
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
  constructor(format: DateTimeFormat) {
    this.dateTimeFormat = format;
  }
}

export class Parameter implements Parameter {
  constructor(paramName: string, type: PossibleType, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.paramName = paramName;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class BodyParameter implements BodyParameter {
  constructor(paramName: string, bodyFormat: BodyFormat, contentType: string, type: PossibleType, paramType: ParameterType, byValue: boolean) {
    this.paramName = paramName;
    this.bodyFormat = bodyFormat;
    this.contentType = contentType;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = 'method';
  }
}

export class FormBodyParameter implements FormBodyParameter {
  constructor(paramName: string, formDataName: string, type: PossibleType, paramType: ParameterType, byValue: boolean) {
    this.paramName = paramName;
    this.formDataName = formDataName;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = 'method';
  }
}

export class FormBodyCollectionParameter implements FormBodyCollectionParameter {
  constructor(paramName: string, formDataName: string, type: SliceType, collectionFormat: ExtendedCollectionFormat, paramType: ParameterType, byValue: boolean) {
    this.paramName = paramName;
    this.formDataName = formDataName;
    this.type = type;
    this.collectionFormat = collectionFormat;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = 'method';
  }
}

export class MultipartFormBodyParameter implements MultipartFormBodyParameter {
  constructor(paramName: string, type: PossibleType, paramType: ParameterType, byValue: boolean) {
    this.paramName = paramName;
    this.multipartForm = true;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = 'method';
  }
}

export class HeaderParameter implements HeaderParameter {
  constructor(paramName: string, headerName: string, type: HeaderType, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.paramName = paramName;
    this.headerName = headerName;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class HeaderCollectionParameter implements HeaderCollectionParameter {
  constructor(paramName: string, headerName: string, type: SliceType, collectionFormat: CollectionFormat, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.paramName = paramName;
    this.headerName = headerName;
    this.type = type;
    this.collectionFormat = collectionFormat;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class HeaderMapParameter implements HeaderMapParameter {
  constructor(paramName: string, headerName: string, type: MapType, collectionPrefix: string, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.paramName = paramName;
    this.headerName = headerName;
    this.type = type;
    this.collectionPrefix = collectionPrefix;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class PathParameter implements PathParameter {
  constructor(paramName: string, pathSegment: string, isEncoded: boolean, type: PathParameterType, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.paramName = paramName;
    this.pathSegment = pathSegment;
    this.isEncoded = isEncoded;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class PathCollectionParameter implements PathCollectionParameter {
  constructor(paramName: string, pathSegment: string, isEncoded: boolean, type: SliceType, collectionFormat: CollectionFormat, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.paramName = paramName;
    this.pathSegment = pathSegment;
    this.isEncoded = isEncoded;
    this.type = type;
    this.collectionFormat = collectionFormat;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class QueryParameter implements QueryParameter {
  constructor(paramName: string, queryParam: string, isEncoded: boolean, type: QueryParameterType, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.paramName = paramName;
    this.queryParameter = queryParam;
    this.isEncoded = isEncoded;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class QueryCollectionParameter implements QueryCollectionParameter {
  constructor(paramName: string, queryParam: string, isEncoded: boolean, type: SliceType, collectionFormat: ExtendedCollectionFormat, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.paramName = paramName;
    this.queryParameter = queryParam;
    this.isEncoded = isEncoded;
    this.type = type;
    this.collectionFormat = collectionFormat;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class URIParameter implements URIParameter {
  constructor(paramName: string, uriPathSegment: string, type: ConstantType | PrimitiveType, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.paramName = paramName;
    this.uriPathSegment = uriPathSegment;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class ResumeTokenParameter implements ResumeTokenParameter {
  constructor(paramName: string) {
    this.isResumeToken = true;
    this.paramName = paramName;
    this.type = new PrimitiveType('string');
    this.paramType = 'optional';
    this.byValue = true;
    this.location = 'method';
  }
}

export class ClientSideDefault implements ClientSideDefault {
  constructor(defaultValue: LiteralValue) {
    this.defaultValue = defaultValue;
  }
}

export class ParameterGroup implements ParameterGroup {
  constructor(paramName: string, groupName: string, required: boolean, location: ParameterLocation) {
    this.groupName = groupName;
    this.location = location;
    this.paramName = paramName;
    // params is required but must be populated post construction
    this.params = new Array<Parameter>();
    this.required = required;
  }
}

export class LiteralValue implements LiteralValue {
  constructor(type: LiteralValueType, literal: any) {
    this.type = type;
    this.literal = literal;
  }
}

export class ResponseEnvelope implements ResponseEnvelope {
  constructor(name: string, description: string, forMethod: Method) {
    this.description = description;
    this.headers = new Array<HeaderResponse | HeaderMapResponse>();
    this.method = forMethod;
    this.name = name;
  }
}

export class HeaderResponse implements HeaderResponse {
  constructor(fieldName: string, type: HeaderType, headerName: string, byValue: boolean) {
    this.fieldName = fieldName;
    this.type = type;
    this.byValue = byValue;
    this.headerName = headerName;
  }
}

export class HeaderMapResponse implements HeaderMapResponse {
  constructor(fieldName: string, type: MapType, collectionPrefix: string, headerName: string, byValue: boolean) {
    this.fieldName = fieldName;
    this.type = type;
    this.collectionPrefix = collectionPrefix;
    this.byValue = byValue;
    this.headerName = headerName;
  }
}

export class AnyResult implements AnyResult {
  constructor(fieldName: string, format: ResultFormat, resultTypes: Record<number, PossibleType>) {
    this.fieldName = fieldName;
    this.format = format;
    this.httpStatusCodeType = resultTypes;
    this.byValue = true;
  }
}

export class BinaryResult implements BinaryResult {
  constructor(fieldName: string, format: BinaryResultFormat) {
    this.fieldName = fieldName;
    this.binaryFormat = format;
    this.byValue = true;
  }
}

export class HeadAsBooleanResult implements HeadAsBooleanResult {
  constructor(fieldName: string) {
    this.fieldName = fieldName;
    this.headAsBoolean = true;
    this.byValue = true;
  }
}

export class MonomorphicResult implements MonomorphicResult {
  constructor(fieldName: string, format: ResultFormat, type: MonomorphicResultType, byValue: boolean) {
    this.fieldName = fieldName;
    this.format = format;
    this.monomorphicType = type;
    this.byValue = byValue;
  }
}

export class PolymorphicResult implements PolymorphicResult {
  constructor(type: InterfaceType) {
    this.interfaceType = type;
    this.format = 'JSON';
  }
}

export class ModelResult implements ModelResult {
  constructor(type: ModelType, format: ResultFormat) {
    this.modelType = type;
    this.format = format;
  }
}

export class XMLInfo implements XMLInfo {
  constructor() {
    this.attribute = false;
    this.text = false;
  }
}
