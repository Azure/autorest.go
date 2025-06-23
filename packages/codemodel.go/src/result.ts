/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import * as client from './client.js';
import { CodeModelError } from './errors.js';
import * as param from './param.js';
import * as type from './type.js';

export type ResultType = AnyResult | BinaryResult | HeadAsBooleanResult | MonomorphicResult | PolymorphicResult | ModelResult;

// AnyResult is for endpoints that return a different schema based on the HTTP status code.
export interface AnyResult {
  // the name of the field within the response envelope
  fieldName: string;

  docs: type.Docs;

  // maps an HTTP status code to a result type.
  // status codes that don't return a schema will be absent.
  httpStatusCodeType: Record<number, type.PossibleType>;

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

  docs: type.Docs;

  binaryFormat: BinaryResultFormat;

  byValue: true;
}

// HeadAsBooleanResult is for responses to HTTP HEAD requests that treat the HTTP status code as success/failure
export interface HeadAsBooleanResult {
  // the name of the field within the response envelope
  fieldName: string;

  docs: type.Docs;

  headAsBoolean: true;

  byValue: true;
}

// this is a special type to support x-ms-header-collection-prefix (i.e. storage)
export interface HeaderMapResponse {
  // the name of the field within the response envelope
  fieldName: string;

  docs: type.Docs;

  type: type.MapType;

  byValue: boolean;

  // the name of the header sent over the wire
  headerName: string;

  collectionPrefix: string;
}

export interface HeaderResponse {
  // the name of the field within the response envelope
  fieldName: string;

  docs: type.Docs;

  type: param.HeaderScalarType;

  byValue: boolean;

  // the name of the header sent over the wire
  headerName: string;
}

// ModelResult is a standard schema response.
// The type is anonymously embedded in the response envelope.
export interface ModelResult {
  docs: type.Docs;

  modelType: type.ModelType;

  // the format in which the result is returned
  format: ModelResultFormat;
}

export type ModelResultFormat = 'JSON' | 'XML';

// MonomorphicResult includes scalar results (ints, bools) or maps/slices of scalars/InterfaceTypes/ModelTypes.
// maps/slices can be recursive and/or combinatorial (e.g. map[string][]*sometype)
export interface MonomorphicResult {
  // the name of the field within the response envelope
  fieldName: string;

  docs: type.Docs;

  monomorphicType: MonomorphicResultType;

  // the format in which the result is returned
  format: ResultFormat;

  byValue: boolean;

  xml?: type.XMLInfo;
}

export type MonomorphicResultType = type.BytesType | type.ConstantType | type.MapType | type.PrimitiveType | type.SliceType | type.TimeType;

// PolymorphicResult is for discriminated types.
// The type is anonymously embedded in the response envelope.
export interface PolymorphicResult {
  docs: type.Docs;

  interfaceType: type.InterfaceType;

  // the format in which the result is returned.
  // only JSON is supported for polymorphic types.
  format: 'JSON';
}

// ResponseEnvelope is the type returned from a client method
export interface ResponseEnvelope {
  name: string;

  docs: type.Docs;

  // for operations that return no body (e.g. a 204) this will be undefined.
  result?: ResultType;

  // any modeled response headers
  headers: Array<HeaderResponse | HeaderMapResponse>;

  method: client.MethodType;
}

export type ResultFormat = 'JSON' | 'XML' | 'Text';

export function isAnyResult(resultType: ResultType): resultType is AnyResult {
  return (<AnyResult>resultType).httpStatusCodeType !== undefined;
}

export function isBinaryResult(resultType: ResultType): resultType is BinaryResult {
  return (<BinaryResult>resultType).binaryFormat !== undefined;
}

export function isHeadAsBooleanResult(resultType: ResultType): resultType is HeadAsBooleanResult {
  return (<HeadAsBooleanResult>resultType).headAsBoolean !== undefined;
}

export function isHeaderMapResponse(resp: HeaderResponse | HeaderMapResponse): resp is HeaderMapResponse {
  return (<HeaderMapResponse>resp).collectionPrefix !== undefined;
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

export function getResultPossibleType(resultType: ResultType): type.PossibleType {
  if (isAnyResult(resultType)) {
    return new type.PrimitiveType('any');
  } else if (isBinaryResult(resultType)) {
    return new type.QualifiedType('ReadCloser', 'io');
  } else if (isHeadAsBooleanResult(resultType)) {
    return new type.PrimitiveType('bool');
  } else if (isMonomorphicResult(resultType)) {
    return resultType.monomorphicType;
  } else if (isPolymorphicResult(resultType)) {
    return resultType.interfaceType;
  } else if (isModelResult(resultType)) {
    return resultType.modelType;
  } else {
    throw new CodeModelError(`unhandled result type ${resultType}`);
  }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

export class AnyResult implements AnyResult {
  constructor(fieldName: string, format: ResultFormat, resultTypes: Record<number, type.PossibleType>) {
    this.fieldName = fieldName;
    this.format = format;
    this.httpStatusCodeType = resultTypes;
    this.byValue = true;
    this.docs = {};
  }
}

export class BinaryResult implements BinaryResult {
  constructor(fieldName: string, format: BinaryResultFormat) {
    this.fieldName = fieldName;
    this.binaryFormat = format;
    this.byValue = true;
    this.docs = {};
  }
}

export class HeadAsBooleanResult implements HeadAsBooleanResult {
  constructor(fieldName: string) {
    this.fieldName = fieldName;
    this.headAsBoolean = true;
    this.byValue = true;
    this.docs = {};
  }
}

export class HeaderMapResponse implements HeaderMapResponse {
  constructor(fieldName: string, type: type.MapType, collectionPrefix: string, headerName: string, byValue: boolean) {
    this.fieldName = fieldName;
    this.type = type;
    this.collectionPrefix = collectionPrefix;
    this.byValue = byValue;
    this.headerName = headerName;
    this.docs = {};
  }
}

export class HeaderResponse implements HeaderResponse {
  constructor(fieldName: string, type: param.HeaderScalarType, headerName: string, byValue: boolean) {
    this.fieldName = fieldName;
    this.type = type;
    this.byValue = byValue;
    this.headerName = headerName;
    this.docs = {};
  }
}

export class ModelResult implements ModelResult {
  constructor(type: type.ModelType, format: ModelResultFormat) {
    this.modelType = type;
    this.format = format;
    this.docs = {};
  }
}

export class MonomorphicResult implements MonomorphicResult {
  constructor(fieldName: string, format: ResultFormat, type: MonomorphicResultType, byValue: boolean) {
    this.fieldName = fieldName;
    this.format = format;
    this.monomorphicType = type;
    this.byValue = byValue;
    this.docs = {};
  }
}

export class PolymorphicResult implements PolymorphicResult {
  constructor(type: type.InterfaceType) {
    this.interfaceType = type;
    this.format = 'JSON';
    this.docs = {};
  }
}

export class ResponseEnvelope implements ResponseEnvelope {
  constructor(name: string, docs: type.Docs, forMethod: client.MethodType) {
    this.docs = docs;
    this.headers = new Array<HeaderResponse | HeaderMapResponse>();
    this.method = forMethod;
    this.name = name;
  }
}
