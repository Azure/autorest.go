/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import * as client from './client.js';
import * as param from './param.js';
import * as result from './result.js';
import * as type from './type.js';

export type ExampleType = AnyExample | ArrayExample | BooleanExample | DictionaryExample | NullExample | NumberExample | QualifiedExample| StringExample | StructExample;

export interface AnyExample {
  kind: 'any';
  value: any;
  type: type.Any;
}

export interface ArrayExample {
  kind: 'array';
  value: Array<ExampleType>;
  type: type.Slice;
}

export interface BooleanExample {
  kind: 'boolean';
  value: boolean;
  type: type.Constant | type.Literal | type.Scalar;
}

export interface DictionaryExample {
  kind: 'dictionary';
  value: Record<string, ExampleType>;
  type: type.Map;
}

// MethodExample is an example for a method. This code model part is for example or test generation.
export interface MethodExample {
  name: string;

  docs: type.Docs;

  filePath: string;

  parameters: Array<ParameterExample>;

  optionalParamsGroup: Array<ParameterExample>;

  responseEnvelope?: ResponseEnvelopeExample;
}

export interface NullExample {
  kind: 'null';
  value: null;
  type: type.WireType;
}

export interface NumberExample {
  kind: 'number';
  value: number;
  type: type.Constant | type.Literal | type.Scalar | type.Time;
}

export interface ParameterExample {
  parameter: client.ClientParameter;
  value: ExampleType;
}

export interface QualifiedExample {
  kind: 'qualified';
  value: any;
  type: type.QualifiedType;
}

export interface ResponseEnvelopeExample {
  response: result.ResponseEnvelope;
  headers: Array<ResponseHeaderExample>;
  result: ExampleType;
}

export interface ResponseHeaderExample {
  header: result.HeaderScalarResponse | result.HeaderMapResponse;
  value: ExampleType;
}

export interface StringExample {
  kind: 'string';
  value: string;
  type: type.Constant | type.EncodedBytes | type.Literal | type.Scalar | type.String | type.Time;
}

export interface StructExample {
  kind: 'model';
  value: Record<string, ExampleType>;
  additionalProperties?: Record<string, ExampleType>;
  type: type.Model | type.PolymorphicModel;
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

export class AnyExample implements AnyExample {
  constructor(value: any) {
    this.kind = 'any';
    this.value = value;
    this.type = new type.Any();
  }
}

export class ArrayExample implements ArrayExample {
  constructor(type: type.Slice) {
    this.kind = 'array';
    this.type = type;
    this.value = [];
  }
}

export class BooleanExample implements BooleanExample {
  constructor(value: boolean, type: type.Constant | type.Literal | type.Scalar) {
    this.kind = 'boolean';
    this.value = value;
    this.type = type;
  }
}

export class DictionaryExample implements DictionaryExample {
  constructor(type: type.Map) {
    this.kind = 'dictionary';
    this.type = type;
    this.value = {};
  }
}

export class MethodExample implements MethodExample {
  constructor(name: string, docs: type.Docs, filePath: string) {
    this.name = name;
    this.docs = docs;
    this.filePath = filePath;
    this.parameters = [];
    this.optionalParamsGroup = [];
  }
}

export class NullExample implements NullExample {
  constructor(type: type.WireType) {
    this.kind = 'null';
    this.type = type;
  }
}

export class NumberExample implements NumberExample {
  constructor(value: number, type: type.Constant | type.Literal | type.Scalar | type.Time) {
    this.kind = 'number';
    this.value = value;
    this.type = type;
  }
}

export class ParameterExample implements ParameterExample {
  constructor(parameter: param.MethodParameter, value: ExampleType) {
    this.parameter = parameter;
    this.value = value;
  }
}

export class QualifiedExample implements QualifiedExample {
  constructor(type: type.QualifiedType, value: any) {
    this.kind = 'qualified';
    this.type = type;
    this.value = value;
  }
}

export class ResponseEnvelopeExample implements ResponseEnvelopeExample {
  constructor(response: result.ResponseEnvelope) {
    this.response = response;
    this.headers = [];
  }
}

export class ResponseHeaderExample implements ResponseHeaderExample {
  constructor(header: result.HeaderScalarResponse | result.HeaderMapResponse, value: ExampleType) {
    this.header = header;
    this.value = value;
  }
}

export class StringExample implements StringExample {
  constructor(value: string, type: type.Constant | type.EncodedBytes | type.Literal | type.Scalar | type.String | type.Time) {
    this.kind = 'string';
    this.value = value;
    this.type = type;
  }
}

export class StructExample implements StructExample {
  constructor(type: type.Model | type.PolymorphicModel) {
    this.kind = 'model';
    this.type = type;
    this.value = {};
  }
}
