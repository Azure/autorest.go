/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import * as type from './type.js';

export type BodyFormat = 'JSON' | 'XML' | 'Text' | 'binary';

// parameter is sent in the HTTP request body
export interface BodyParameter extends Parameter {
  bodyFormat: BodyFormat;

  // "application/text" etc...
  contentType: string;

  xml?: type.XMLInfo;
}

// ClientSideDefault is used to represent parameters that have a default value on the client side
export interface ClientSideDefault {
  defaultValue: type.LiteralValue;
}

export type CollectionFormat = 'csv' | 'ssv' | 'tsv' | 'pipes';

export type ExtendedCollectionFormat = CollectionFormat | 'multi';

export interface FormBodyCollectionParameter extends Parameter {
  formDataName: string;

  type: type.SliceType;

  collectionFormat: ExtendedCollectionFormat;
}

export interface FormBodyParameter extends Parameter {
  formDataName: string;
}

export interface HeaderCollectionParameter extends Parameter {
  headerName: string;

  type: type.SliceType;

  collectionFormat: CollectionFormat;
}

// this is a special type to support x-ms-header-collection-prefix (i.e. storage)
export interface HeaderMapParameter extends Parameter {
  headerName: string;

  type: type.MapType;

  collectionPrefix: string;
}

// parameter is sent via an HTTP header
export interface HeaderParameter extends Parameter {
  headerName: string;

  type: HeaderType;
}

export type HeaderType = type.BytesType | type.ConstantType | type.PrimitiveType | type.TimeType | type.LiteralValue;

export interface MultipartFormBodyParameter extends Parameter {
  multipartForm: true;
}

// Parameter is a parameter for a client method
export interface Parameter {
  name: string;

  docs: type.Docs;

  // NOTE: if the type is a LiteralValue the paramType will either be literal or flag
  type: type.PossibleType;

  // kind will have value literal or flag when type is a LiteralValue (see above comment).
  kind: ParameterKind;

  byValue: boolean;

  group?: ParameterGroup;

  location: ParameterLocation;
}

// required - the parameter is required
// optional - the parameter is optional
// literal - there is no formal parameter, the value is emitted directly in the code (e.g. the Accept header parameter)
// flag - the value is a literal and emitted in the code, but sent IFF the flag param is not nil (i.e. an optional LiteralValue)
// ClientSideDefault - the parameter has an emitted default value that's sent if one isn't specified (implies optional)
export type ParameterKind = 'required' | 'optional' | 'literal' | 'flag' | ClientSideDefault;

export type ParameterLocation = 'client' | 'method';

export interface ParameterGroup {
  // name is the name of the parameter
  name: string;

  docs: type.Docs;

  // groupName is the name of the param group (i.e. the struct name)
  groupName: string;

  required: boolean;

  location: ParameterLocation;

  // the params that belong to this group
  params: Array<Parameter>;
}

// PartialBodyParameter is a field within a struct type sent in the body
export interface PartialBodyParameter extends Parameter {
  // the name of the field over the wire
  serializedName: string;

  format: 'JSON' | 'XML';

  xml?: type.XMLInfo;
}

export interface PathCollectionParameter extends Parameter {
  pathSegment: string;

  type: type.SliceType;

  isEncoded: boolean;

  collectionFormat: CollectionFormat;
}

// parameter is a segment in a path
export interface PathParameter extends Parameter {
  pathSegment: string;

  type: PathParameterType;

  isEncoded: boolean;
}

export type PathParameterType = type.BytesType | type.ConstantType | type.PrimitiveType | type.TimeType | type.LiteralValue;

export interface QueryCollectionParameter extends Parameter {
  queryParameter: string;

  type: type.SliceType;

  isEncoded: boolean;

  collectionFormat: ExtendedCollectionFormat;
}

// parameter is sent via an HTTP query parameter
export interface QueryParameter extends Parameter {
  queryParameter: string;

  type: QueryParameterType;

  isEncoded: boolean;
}

export type QueryParameterType = type.BytesType | type.ConstantType | type.PrimitiveType | type.TimeType | type.LiteralValue;

export interface ResumeTokenParameter extends Parameter {
  isResumeToken: true;
}

// parameter is a segment in the host
export interface URIParameter extends Parameter {
  uriPathSegment: string;

  type: URIParameterType;
}

export type URIParameterType = type.ConstantType | type.PrimitiveType;

export function isBodyParameter(param: Parameter): param is BodyParameter {
  return (<BodyParameter>param).bodyFormat !== undefined;
}

export function isClientSideDefault(kind: ParameterKind): kind is ClientSideDefault {
  return (<ClientSideDefault>kind).defaultValue !== undefined;
}

export function isPartialBodyParameter(param: Parameter): param is PartialBodyParameter {
  return (<PartialBodyParameter>param).serializedName !== undefined;
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

export function isRequiredParameter(param: Parameter): boolean {
  // parameters with a client-side default value are always optional
  if (isClientSideDefault(param.kind)) {
    return false;
  }
  return param.kind === 'required';
}

export function isLiteralParameter(param: Parameter): boolean {
  if (isClientSideDefault(param.kind)) {
    return false;
  }
  return param.kind === 'literal';
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// base types
///////////////////////////////////////////////////////////////////////////////////////////////////

export class Parameter implements Parameter {
  constructor(name: string, type: type.PossibleType, kind: ParameterKind, byValue: boolean, location: ParameterLocation) {
    this.name = name;
    this.type = type;
    this.kind = kind;
    this.byValue = byValue;
    this.location = location;
    this.docs = {};
  }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

export class BodyParameter extends Parameter implements BodyParameter {
  constructor(name: string, bodyFormat: BodyFormat, contentType: string, type: type.PossibleType, kind: ParameterKind, byValue: boolean) {
    super(name, type, kind, byValue, 'method');
    this.bodyFormat = bodyFormat;
    this.contentType = contentType;
  }
}

export class ClientSideDefault implements ClientSideDefault {
  constructor(defaultValue: type.LiteralValue) {
    this.defaultValue = defaultValue;
  }
}

export class FormBodyCollectionParameter extends Parameter implements FormBodyCollectionParameter {
  constructor(name: string, formDataName: string, type: type.SliceType, collectionFormat: ExtendedCollectionFormat, kind: ParameterKind, byValue: boolean) {
    super(name, type, kind, byValue, 'method');
    this.formDataName = formDataName;
    this.collectionFormat = collectionFormat;
  }
}

export class FormBodyParameter extends Parameter implements FormBodyParameter {
  constructor(name: string, formDataName: string, type: type.PossibleType, kind: ParameterKind, byValue: boolean) {
    super(name, type, kind, byValue, 'method');
    this.formDataName = formDataName;
  }
}

export class HeaderCollectionParameter extends Parameter implements HeaderCollectionParameter {
  constructor(name: string, headerName: string, type: type.SliceType, collectionFormat: CollectionFormat, kind: ParameterKind, byValue: boolean, location: ParameterLocation) {
    super(name, type, kind, byValue, location);
    this.headerName = headerName;
    this.collectionFormat = collectionFormat;
  }
}

export class HeaderMapParameter extends Parameter implements HeaderMapParameter {
  constructor(name: string, headerName: string, type: type.MapType, collectionPrefix: string, kind: ParameterKind, byValue: boolean, location: ParameterLocation) {
    super(name, type, kind, byValue, location);
    this.headerName = headerName;
    this.collectionPrefix = collectionPrefix;
  }
}

export class HeaderParameter extends Parameter implements HeaderParameter {
  constructor(name: string, headerName: string, type: HeaderType, kind: ParameterKind, byValue: boolean, location: ParameterLocation) {
    super(name, type, kind, byValue, location);
    this.headerName = headerName;
  }
}

export class MultipartFormBodyParameter extends Parameter implements MultipartFormBodyParameter {
  constructor(name: string, type: type.PossibleType, kind: ParameterKind, byValue: boolean) {
    super(name, type, kind, byValue, 'method');
    this.multipartForm = true;
  }
}

export class ParameterGroup implements ParameterGroup {
  constructor(name: string, groupName: string, required: boolean, location: ParameterLocation) {
    this.groupName = groupName;
    this.location = location;
    this.name = name;
    // params is required but must be populated post construction
    this.params = new Array<Parameter>();
    this.required = required;
    this.docs = {};
  }
}

export class PartialBodyParameter extends Parameter implements PartialBodyParameter{
  constructor(name: string, serializedName: string, format: 'JSON' | 'XML', type: type.PossibleType, kind: ParameterKind, byValue: boolean) {
    super(name, type, kind, byValue, 'method');
    this.format = format;
    this.serializedName = serializedName;
  }
}

export class PathCollectionParameter extends Parameter implements PathCollectionParameter {
  constructor(name: string, pathSegment: string, isEncoded: boolean, type: type.SliceType, collectionFormat: CollectionFormat, kind: ParameterKind, byValue: boolean, location: ParameterLocation) {
    super(name, type, kind, byValue, location);
    this.pathSegment = pathSegment;
    this.isEncoded = isEncoded;
    this.collectionFormat = collectionFormat;
  }
}

export class PathParameter extends Parameter implements PathParameter {
  constructor(name: string, pathSegment: string, isEncoded: boolean, type: PathParameterType, kind: ParameterKind, byValue: boolean, location: ParameterLocation) {
    super(name, type, kind, byValue, location);
    this.pathSegment = pathSegment;
    this.isEncoded = isEncoded;
  }
}

export class QueryCollectionParameter extends Parameter implements QueryCollectionParameter {
  constructor(name: string, queryParam: string, isEncoded: boolean, type: type.SliceType, collectionFormat: ExtendedCollectionFormat, kind: ParameterKind, byValue: boolean, location: ParameterLocation) {
    super(name, type, kind, byValue, location);
    this.queryParameter = queryParam;
    this.isEncoded = isEncoded;
    this.collectionFormat = collectionFormat;
  }
}

export class QueryParameter extends Parameter implements QueryParameter {
  constructor(name: string, queryParam: string, isEncoded: boolean, type: QueryParameterType, kind: ParameterKind, byValue: boolean, location: ParameterLocation) {
    super(name, type, kind, byValue, location);
    this.queryParameter = queryParam;
    this.isEncoded = isEncoded;
  }
}

export class ResumeTokenParameter extends Parameter implements ResumeTokenParameter {
  constructor() {
    super('ResumeToken', new type.PrimitiveType('string'), 'optional', true, 'method');
    this.isResumeToken = true;
    this.docs.summary = 'Resumes the long-running operation from the provided token.';
  }
}

export class URIParameter extends Parameter implements URIParameter {
  constructor(name: string, uriPathSegment: string, type: type.ConstantType | type.PrimitiveType, kind: ParameterKind, byValue: boolean, location: ParameterLocation) {
    super(name, type, kind, byValue, location);
    this.uriPathSegment = uriPathSegment;
  }
}
