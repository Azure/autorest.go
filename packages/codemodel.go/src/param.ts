/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import * as type from './type.js';

/** indicates the wire format for request bodies */
export type BodyFormat = 'JSON' | 'XML' | 'Text' | 'binary';

/** the union of all form body parameter types */
export type FormBodyParameter = FormBodyCollectionParameter | FormBodyScalarParameter;

/** the union of all header parameter types */
export type HeaderParameter = HeaderCollectionParameter | HeaderMapParameter | HeaderScalarParameter;

/** the union of all path parameter types */
export type PathParameter = PathCollectionParameter | PathScalarParameter;

/** the union of all query parameter types */
export type QueryParameter = QueryCollectionParameter | QueryScalarParameter;

/** defines the possible method parameter types */
export type MethodParameter = BodyParameter | FormBodyParameter | HeaderParameter | MultipartFormBodyParameter 
  | PartialBodyParameter | PathParameter | QueryParameter | ResumeTokenParameter | URIParameter;

/** parameter is sent in the HTTP request body */
export interface BodyParameter extends ParameterBase {
  kind: 'bodyParam';

  /** the wire format of the request body */
  bodyFormat: BodyFormat;

  /** value used for the Content-Type header */
  contentType: string;

  /** optional XML schema metadata */
  xml?: type.XMLInfo;
}

/** represents parameters that have a default value on the client side */
export interface ClientSideDefault {
  /** the literal used for the client-side default value */
  defaultValue: type.Literal;
}

/** indicates how a collection is formatted on the wire */
export type CollectionFormat = 'csv' | 'ssv' | 'tsv' | 'pipes';

/** includes additional wire formats */
export type ExtendedCollectionFormat = CollectionFormat | 'multi';

/** a collection that's placed in form body data */
export interface FormBodyCollectionParameter extends ParameterBase {
  kind: 'formBodyCollectionParam';

  /** the form data name for this parameter */
  formDataName: string;

  /** the type of the parameter */
  type: type.Slice;

  /** the format of the collection */
  collectionFormat: ExtendedCollectionFormat;
}

/** parameter that's placed in form body data */
export interface FormBodyScalarParameter extends ParameterBase {
  kind: 'formBodyScalarParam';

  /** the form data name for this parameter */
  formDataName: string;
}

/** a collection that goes in a HTTP header */
export interface HeaderCollectionParameter extends ParameterBase {
  kind: 'headerCollectionParam';

  /** the header in the HTTP request */
  headerName: string;

  /** the collection of header param values */
  type: type.Slice;

  /** the format of the collection */
  collectionFormat: CollectionFormat;
}

/**
 * parameter map collection that goes in a HTTP header.
 * NOTE: this is a specialized parameter type to support storage.
 */
export interface HeaderMapParameter extends ParameterBase {
  kind: 'headerMapParam';

  /** the header prefix for each header name in type */
  headerName: string;

  /** the type of the param */
  type: type.Map;
}

/** a value that goes in a HTTP header */
export interface HeaderScalarParameter extends ParameterBase {
  kind: 'headerScalarParam';

  /** the header in the HTTP request */
  headerName: string;

  /** the type of the parameter */
  type: HeaderScalarType;
}

/** defines the possible types for a scalar header */
export type HeaderScalarType = type.Constant | type.EncodedBytes | type.Literal | type.Scalar | type.String | type.Time;

/** parameter goes in multipart/form body */
export interface MultipartFormBodyParameter extends ParameterBase {
  kind: 'multipartFormBodyParam';
}

/**
 * defines the style of a parameter.
 * 
 * required - the parameter is required
 * optional - the parameter is optional
 * literal - there is no formal parameter, the value is emitted directly in the code (e.g. the Accept header parameter)
 * flag - the value is a literal and emitted in the code, but sent IFF the flag param is not nil (i.e. an optional LiteralValue)
 * ClientSideDefault - the parameter has an emitted default value that's sent if one isn't specified (implies optional)
 */
export type ParameterStyle = 'required' | 'optional' | 'literal' | 'flag' | ClientSideDefault;

/** indicates where the value of a parameter originates */
export type ParameterLocation = 'client' | 'method';

/** a parameter that's not used for creating HTTP requests (e.g. a credential parameter) */
export interface Parameter extends ParameterBase {
  kind: 'parameter';
}

/** a struct that contains a grouping of parameters */
export interface ParameterGroup {
  kind: 'paramGroup';

  /** the name of the parameter */
  name: string;

  /** any docs for the parameter */
  docs: type.Docs;

  /** the name of the parameter group (i.e. the struct name) */
  groupName: string;

  /** indicates if the parameter is required */
  required: boolean;

  /** the location of the parameter */
  location: ParameterLocation;

  /** the parameters that belong to this group */
  params: Array<MethodParameter>;
}

/** a parameter that's a field within a type passed via the HTTP request body */
export interface PartialBodyParameter extends ParameterBase {
  kind: 'partialBodyParam';

  /** the name of the field within the type sent in the request body */
  serializedName: string;

  /** the wire format of the underlying type */
  format: 'JSON' | 'XML';

  /** optional XML schema metadata */
  xml?: type.XMLInfo;
}

/** a collection of values that go in the HTTP path */
export interface PathCollectionParameter extends ParameterBase {
  kind: 'pathCollectionParam';

  /** the segment name to be replaced with the values */
  pathSegment: string;

  /** the type of the parameter */
  type: type.Slice;

  /** indicates if the values must be URL encoded */
  isEncoded: boolean;

  /** the format of the collection */
  collectionFormat: CollectionFormat;
}

/** a value that goes in the HTTP path */
export interface PathScalarParameter extends ParameterBase {
  kind: 'pathScalarParam';

  /** the segment name to be replaced with the value */
  pathSegment: string;

  /** the type of the parameter */
  type: PathScalarParameterType;

  /** indicates if the values must be URL encoded */
  isEncoded: boolean;
}

/** defines the possible types for a PathScalarParameter */
export type PathScalarParameterType = type.Constant | type.EncodedBytes | type.Literal | type.Scalar | type.String | type.Time;

/** a collection of values that go in the HTTP query string */
export interface QueryCollectionParameter extends ParameterBase {
  kind: 'queryCollectionParam';

  /** the query string's key name */
  queryParameter: string;

  /** the type of the parameter */
  type: type.Slice;

  /** indicates if the values must be URL encoded */
  isEncoded: boolean;

  /** the format of the collection */
  collectionFormat: ExtendedCollectionFormat;
}

/** a scalar value that goes in the HTTP query string */
export interface QueryScalarParameter extends ParameterBase {
  kind: 'queryScalarParam';

  /** the query string's key name */
  queryParameter: string;

   /** the type of the parameter */
  type: QueryScalarParameterType;

  /** indicates if the value must be URL encoded */
  isEncoded: boolean;
}

/** defines the possible types for a QueryScalarParameter */
export type QueryScalarParameterType = type.Constant | type.EncodedBytes | type.Literal | type.Scalar | type.String | type.Time;

/** the synthesized resume token parameter for LROs */
export interface ResumeTokenParameter extends ParameterBase {
  kind: 'resumeTokenParam';
}

/** a segment of the host's URI */
export interface URIParameter extends ParameterBase {
  kind: 'uriParam';

  /** the segment name to be replaced with the value */
  uriPathSegment: string;

  /** the type of the parameter */
  type: URIParameterType;
}

/** defines the possible types for a URIParameter */
export type URIParameterType = type.Constant | type.Scalar | type.String;

/** narrows style to a ClientSideDefault within the conditional block */
export function isClientSideDefault(style: ParameterStyle): style is ClientSideDefault {
  return (<ClientSideDefault>style).defaultValue !== undefined;
}

/** narrows param to a FormBodyParameter within the conditional block */
export function isFormBodyParameter(param: MethodParameter): param is FormBodyParameter {
  return param.kind === 'formBodyCollectionParam' || param.kind === 'formBodyScalarParam';
}

/** narrows param to a HeaderParameter within the conditional block */
export function isHeaderParameter(param: MethodParameter): param is HeaderParameter {
  return param.kind === 'headerCollectionParam' || param.kind === 'headerMapParam' || param.kind === 'headerScalarParam';
}

/** narrows type to a HeaderScalarType within the conditional block */
export function isHeaderScalarType(type: type.WireType): type is HeaderScalarType {
  switch (type.kind) {
    case 'constant':
    case 'encodedBytes':
    case 'literal':
    case 'scalar':
    case 'string':
    case 'time':
      return true;
    default:
      return false;
  }
}

/** narrows param to a PathParameter within the conditional block */
export function isPathParameter(param: MethodParameter): param is PathParameter {
  return param.kind === 'pathCollectionParam' || param.kind === 'pathScalarParam';
}

/** narrows type to a PathScalarParameterType within the conditional block */
export function isPathScalarParameterType(type: type.WireType): type is PathScalarParameterType {
  switch (type.kind) {
    case 'constant':
    case 'encodedBytes':
    case 'literal':
    case 'scalar':
    case 'string':
    case 'time':
      return true;
    default:
      return false;
  }
}

/** narrows param to a QueryParameter within the conditional block */
export function isQueryParameter(param: MethodParameter): param is QueryParameter {
  return param.kind === 'queryCollectionParam' || param.kind === 'queryScalarParam';
}

/** narrows type to a QueryScalarParameterType within the conditional block */
export function isQueryScalarParameterType(type: type.WireType): type is QueryScalarParameterType {
  switch (type.kind) {
    case 'constant':
    case 'encodedBytes':
    case 'literal':
    case 'scalar':
    case 'string':
    case 'time':
      return true;
    default:
      return false;
  }
}

/** returns true if the param is required */
export function isRequiredParameter(param: MethodParameter | Parameter): boolean {
  // parameters with a client-side default value are always optional
  if (isClientSideDefault(param.style)) {
    return false;
  }
  return param.style === 'required';
}

/** narrows type to a URIParameterType within the conditional block */
export function isURIParameterType(type: type.WireType): type is URIParameterType {
  switch (type.kind) {
    case 'constant':
    case 'scalar':
    case 'string':
      return true;
    default:
      return false;
  }
}

/** returns true if the param is a literal */
export function isLiteralParameter(param: MethodParameter | Parameter): boolean {
  if (isClientSideDefault(param.style)) {
    return false;
  }
  return param.style === 'literal';
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// base types
///////////////////////////////////////////////////////////////////////////////////////////////////

interface ParameterBase {
  /** the name of the parameter */
  name: string;

  /** any docs for the parameter */
  docs: type.Docs;

  /**
   * the parameter's type.
   * NOTE: if the type is a LiteralValue the style will either be literal or flag
   */
  type: type.WireType;

  /** kind will have value literal or flag when type is a LiteralValue (see above comment) */
  style: ParameterStyle;

  /** indicates if the parameter is passed by value or by pointer */
  byValue: boolean;

  /** indicates if the parameter belongs to a parameter group */
  group?: ParameterGroup;

  /** indicates if the parameter is part of the method signature or a value on the client */
  location: ParameterLocation;
}

class ParameterBase implements ParameterBase {
  constructor(name: string, type: type.WireType, style: ParameterStyle, byValue: boolean, location: ParameterLocation) {
    this.name = name;
    this.type = type;
    this.style = style;
    this.byValue = byValue;
    this.location = location;
    this.docs = {};
  }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

export class BodyParameter extends ParameterBase implements BodyParameter {
  constructor(name: string, bodyFormat: BodyFormat, contentType: string, type: type.WireType, style: ParameterStyle, byValue: boolean) {
    super(name, type, style, byValue, 'method');
    this.kind = 'bodyParam';
    this.bodyFormat = bodyFormat;
    this.contentType = contentType;
  }
}

export class ClientSideDefault implements ClientSideDefault {
  constructor(defaultValue: type.Literal) {
    this.defaultValue = defaultValue;
  }
}

export class FormBodyCollectionParameter extends ParameterBase implements FormBodyCollectionParameter {
  constructor(name: string, formDataName: string, type: type.Slice, collectionFormat: ExtendedCollectionFormat, style: ParameterStyle, byValue: boolean) {
    super(name, type, style, byValue, 'method');
    this.kind = 'formBodyCollectionParam';
    this.formDataName = formDataName;
    this.collectionFormat = collectionFormat;
  }
}

export class FormBodyScalarParameter extends ParameterBase implements FormBodyScalarParameter {
  constructor(name: string, formDataName: string, type: type.WireType, style: ParameterStyle, byValue: boolean) {
    super(name, type, style, byValue, 'method');
    this.kind = 'formBodyScalarParam';
    this.formDataName = formDataName;
  }
}

export class HeaderCollectionParameter extends ParameterBase implements HeaderCollectionParameter {
  constructor(name: string, headerName: string, type: type.Slice, collectionFormat: CollectionFormat, style: ParameterStyle, byValue: boolean, location: ParameterLocation) {
    super(name, type, style, byValue, location);
    this.kind = 'headerCollectionParam';
    this.headerName = headerName;
    this.collectionFormat = collectionFormat;
  }
}

export class HeaderMapParameter extends ParameterBase implements HeaderMapParameter {
  constructor(name: string, headerName: string, type: type.Map, style: ParameterStyle, byValue: boolean, location: ParameterLocation) {
    super(name, type, style, byValue, location);
    this.kind = 'headerMapParam';
    this.headerName = headerName;
  }
}

export class HeaderScalarParameter extends ParameterBase implements HeaderScalarParameter {
  constructor(name: string, headerName: string, type: HeaderScalarType, style: ParameterStyle, byValue: boolean, location: ParameterLocation) {
    super(name, type, style, byValue, location);
    this.kind = 'headerScalarParam';
    this.headerName = headerName;
  }
}

export class MultipartFormBodyParameter extends ParameterBase implements MultipartFormBodyParameter {
  constructor(name: string, type: type.WireType, style: ParameterStyle, byValue: boolean) {
    super(name, type, style, byValue, 'method');
    this.kind = 'multipartFormBodyParam';
  }
}

export class Parameter extends ParameterBase implements Parameter {
  constructor(name: string, type: type.WireType, style: ParameterStyle, byValue: boolean, location: ParameterLocation) {
    super(name, type, style, byValue, location);
    this.kind = 'parameter';
  }
}

export class ParameterGroup implements ParameterGroup {
  constructor(name: string, groupName: string, required: boolean, location: ParameterLocation) {
    this.kind = 'paramGroup';
    this.groupName = groupName;
    this.location = location;
    this.name = name;
    // params is required but must be populated post construction
    this.params = new Array<MethodParameter>();
    this.required = required;
    this.docs = {};
  }
}

export class PartialBodyParameter extends ParameterBase implements PartialBodyParameter{
  constructor(name: string, serializedName: string, format: 'JSON' | 'XML', type: type.WireType, style: ParameterStyle, byValue: boolean) {
    super(name, type, style, byValue, 'method');
    this.kind = 'partialBodyParam';
    this.format = format;
    this.serializedName = serializedName;
  }
}

export class PathCollectionParameter extends ParameterBase implements PathCollectionParameter {
  constructor(name: string, pathSegment: string, isEncoded: boolean, type: type.Slice, collectionFormat: CollectionFormat, style: ParameterStyle, byValue: boolean, location: ParameterLocation) {
    super(name, type, style, byValue, location);
    this.kind = 'pathCollectionParam';
    this.pathSegment = pathSegment;
    this.isEncoded = isEncoded;
    this.collectionFormat = collectionFormat;
  }
}

export class PathScalarParameter extends ParameterBase implements PathScalarParameter {
  constructor(name: string, pathSegment: string, isEncoded: boolean, type: PathScalarParameterType, style: ParameterStyle, byValue: boolean, location: ParameterLocation) {
    super(name, type, style, byValue, location);
    this.kind = 'pathScalarParam';
    this.pathSegment = pathSegment;
    this.isEncoded = isEncoded;
  }
}

export class QueryCollectionParameter extends ParameterBase implements QueryCollectionParameter {
  constructor(name: string, queryParam: string, isEncoded: boolean, type: type.Slice, collectionFormat: ExtendedCollectionFormat, style: ParameterStyle, byValue: boolean, location: ParameterLocation) {
    super(name, type, style, byValue, location);
    this.kind = 'queryCollectionParam';
    this.queryParameter = queryParam;
    this.isEncoded = isEncoded;
    this.collectionFormat = collectionFormat;
  }
}

export class QueryScalarParameter extends ParameterBase implements QueryScalarParameter {
  constructor(name: string, queryParam: string, isEncoded: boolean, type: QueryScalarParameterType, style: ParameterStyle, byValue: boolean, location: ParameterLocation) {
    super(name, type, style, byValue, location);
    this.kind = 'queryScalarParam';
    this.queryParameter = queryParam;
    this.isEncoded = isEncoded;
  }
}

export class ResumeTokenParameter extends ParameterBase implements ResumeTokenParameter {
  constructor() {
    super('ResumeToken', new type.String(), 'optional', true, 'method');
    this.kind = 'resumeTokenParam';
    this.docs.summary = 'Resumes the long-running operation from the provided token.';
  }
}

export class URIParameter extends ParameterBase implements URIParameter {
  constructor(name: string, uriPathSegment: string, type: URIParameterType, style: ParameterStyle, byValue: boolean, location: ParameterLocation) {
    super(name, type, style, byValue, location);
    this.kind = 'uriParam';
    this.uriPathSegment = uriPathSegment;
  }
}
