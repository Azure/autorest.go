/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import * as client from './client.js';
import * as param from './param.js';
import * as type from './type.js';

/** defines the possible method result types within a response envelope */
export type Result = AnyResult | BinaryResult | HeadAsBooleanResult | ModelResult | MonomorphicResult | PolymorphicResult;

/** for endpoints that return a different schema based on the HTTP status code */
export interface AnyResult {
  kind: 'anyResult';

  /** the name of the field within the response envelope */
  fieldName: string;

  /** any docs for the result */
  docs: type.Docs;

  /**
   * maps an HTTP status code to a result type.
   * status codes that don't return a schema will be absent.
   */
  httpStatusCodeType: Record<number, type.PossibleType>;

  /** the format in which the result is returned */
  format: ResultFormat;
}

/** for endpoints that return a streaming response (i.e. the http.Response.Body) */
export interface BinaryResult {
  kind: 'binaryResult';

  /** the name of the field within the response envelope */
  fieldName: string;

  /** any docs for the result */
  docs: type.Docs;
}

/** used for responses to HTTP HEAD requests that treat the HTTP status code as success/failure */
export interface HeadAsBooleanResult {
  kind: 'headAsBooleanResult';

  /** the name of the field within the response envelope */
  fieldName: string;

  /** any docs for the result */
  docs: type.Docs;
}

/**
 * a collection of header responses.
 * NOTE: this is a specialized type to support storage.
 */
export interface HeaderMapResponse {
  kind: 'headerMapResponse';

  /** the name of the field within the response envelope */
  fieldName: string;

  /** any docs for the header */
  docs: type.Docs;

  /** the type of the response header */
  type: type.MapType;

  /** the header prefix for each header name in type */
  headerName: string;
}

/** a typed header returned in a HTTP response */
export interface HeaderScalarResponse {
  kind: 'headerScalarResponse';

  /** the name of the field within the response envelope */
  fieldName: string;

  /** any docs for the header */
  docs: type.Docs;

  /** the type of the response header */
  type: param.HeaderScalarType;

  /** indicates if the header is returned by value or by pointer */
  byValue: boolean;

  /** the name of the header sent over the wire */
  headerName: string;
}

/**
 * used for methods that return a typed payload.
 * the type is anonymously embedded in the response envelope.
 */
export interface ModelResult {
  kind: 'modelResult';

  /** any docs for the result type */
  docs: type.Docs;

  /** the type returned in the response envelope */
  modelType: type.ModelType;

  /** the format in which the result is returned */
  format: ModelResultFormat;
}

/** the wire format for model response bodies */
export type ModelResultFormat = 'JSON' | 'XML';

/**
 * includes scalar results (ints, bools) or maps/slices of scalars/InterfaceTypes/ModelTypes.
 * maps/slices can be recursive and/or combinatorial (e.g. map[string][]*sometype)
 */
export interface MonomorphicResult {
  kind: 'monomorphicResult';

  /** the name of the field within the response envelope */
  fieldName: string;

  /** any docs for the result type */
  docs: type.Docs;

  /** the type returned in the response envelope */
  monomorphicType: MonomorphicResultType;

  /** the format in which the result is returned */
  format: ResultFormat;

  /** indicates if the response type is returned by value or by pointer */
  byValue: boolean;

  /** optional XML schema metadata */
  xml?: type.XMLInfo;
}

/** the possible monomorphic result types */
export type MonomorphicResultType = type.BytesType | type.ConstantType | type.MapType | type.PrimitiveType | type.SliceType | type.TimeType;

/**
 * used for methods that return a discriminated type.
 * the type is anonymously embedded in the response envelope.
 */
export interface PolymorphicResult {
  kind: 'polymorphicResult';

  /** any docs for the result type */
  docs: type.Docs;

  /** the interface type used for the discriminated union of possible types */
  interfaceType: type.InterfaceType;

  /**
   * the format in which the result is returned.
   * only JSON is supported for polymorphic types.
   */
  format: 'JSON';
}

/**
 * the type returned from a client method.
 * this combines response headers with any response body.
 */
export interface ResponseEnvelope {
  /** the name of the type */
  name: string;

  /** any docs for the type */
  docs: type.Docs;

  /**
   * contains the body result type.
   * for operations that return no body (e.g. a 204) this will be undefined.
   */
  result?: Result;

  /** any modeled response headers. can be empty */
  headers: Array<HeaderScalarResponse | HeaderMapResponse>;

  /** the method that returns this type */
  method: client.MethodType;
}

/** indicates the wire format for response bodies */
export type ResultFormat = 'JSON' | 'XML' | 'Text';

/** returns the underlying type used for the specified result type */
export function getResultType(result: Result): type.InterfaceType | type.ModelType | MonomorphicResultType | type.PrimitiveType | type.QualifiedType {
  switch (result.kind) {
    case 'anyResult':
      return new type.PrimitiveType('any');
    case 'binaryResult':
      return new type.QualifiedType('ReadCloser', 'io');
    case 'headAsBooleanResult':
      return new type.PrimitiveType('bool');
    case 'modelResult':
      return result.modelType;
    case 'monomorphicResult':
      return result.monomorphicType;
    case 'polymorphicResult':
      return result.interfaceType;
  }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

export class AnyResult implements AnyResult {
  constructor(fieldName: string, format: ResultFormat, resultTypes: Record<number, type.PossibleType>) {
    this.kind = 'anyResult';
    this.fieldName = fieldName;
    this.format = format;
    this.httpStatusCodeType = resultTypes;
    this.docs = {};
  }
}

export class BinaryResult implements BinaryResult {
  constructor(fieldName: string) {
    this.kind = 'binaryResult';
    this.fieldName = fieldName;
    this.docs = {};
  }
}

export class HeadAsBooleanResult implements HeadAsBooleanResult {
  constructor(fieldName: string) {
    this.kind = 'headAsBooleanResult';
    this.fieldName = fieldName;
    this.docs = {};
  }
}

export class HeaderMapResponse implements HeaderMapResponse {
  constructor(fieldName: string, type: type.MapType, headerName: string) {
    this.kind = 'headerMapResponse';
    this.fieldName = fieldName;
    this.type = type;
    this.headerName = headerName;
    this.docs = {};
  }
}

export class HeaderScalarResponse implements HeaderScalarResponse {
  constructor(fieldName: string, type: param.HeaderScalarType, headerName: string, byValue: boolean) {
    this.kind = 'headerScalarResponse';
    this.fieldName = fieldName;
    this.type = type;
    this.byValue = byValue;
    this.headerName = headerName;
    this.docs = {};
  }
}

export class ModelResult implements ModelResult {
  constructor(type: type.ModelType, format: ModelResultFormat) {
    this.kind = 'modelResult';
    this.modelType = type;
    this.format = format;
    this.docs = {};
  }
}

export class MonomorphicResult implements MonomorphicResult {
  constructor(fieldName: string, format: ResultFormat, type: MonomorphicResultType, byValue: boolean) {
    this.kind = 'monomorphicResult';
    this.fieldName = fieldName;
    this.format = format;
    this.monomorphicType = type;
    this.byValue = byValue;
    this.docs = {};
  }
}

export class PolymorphicResult implements PolymorphicResult {
  constructor(type: type.InterfaceType) {
    this.kind = 'polymorphicResult';
    this.interfaceType = type;
    this.format = 'JSON';
    this.docs = {};
  }
}

export class ResponseEnvelope implements ResponseEnvelope {
  constructor(name: string, docs: type.Docs, forMethod: client.MethodType) {
    this.docs = docs;
    this.headers = new Array<HeaderScalarResponse | HeaderMapResponse>();
    this.method = forMethod;
    this.name = name;
  }
}
