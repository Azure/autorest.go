/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import * as param from './param.js';
import * as result from './result.js';
import { BytesType, ConstantType, Docs, LiteralValue, MapType, ModelType, PolymorphicType, PossibleType, PrimitiveType, QualifiedType, SliceType, TimeType } from './type.js';

export type ExampleType = AnyExample | ArrayExample | BooleanExample | DictionaryExample | NullExample | NumberExample | QualifiedExample| StringExample | StructExample;

export interface AnyExample {
  kind: 'any';
  value: any;
  type: PossibleType;
}

export interface ArrayExample {
  kind: 'array';
  value: Array<ExampleType>;
  type: SliceType;
}

export interface BooleanExample {
  kind: 'boolean';
  value: boolean;
  type: ConstantType | LiteralValue | PrimitiveType;
}

export interface DictionaryExample {
  kind: 'dictionary';
  value: Record<string, ExampleType>;
  type: MapType;
}

// MethodExample is an example for a method. This code model part is for example or test generation.
export interface MethodExample {
  name: string;

  docs: Docs;

  filePath: string;

  parameters: Array<ParameterExample>;

  optionalParamsGroup: Array<ParameterExample>;

  responseEnvelope?: ResponseEnvelopeExample;
}

export interface NullExample {
  kind: 'null';
  value: null;
  type: PossibleType;
}

export interface NumberExample {
  kind: 'number';
  value: number;
  type: ConstantType | LiteralValue | TimeType | PrimitiveType;
}

export interface ParameterExample {
  parameter: param.Parameter;
  value: ExampleType;
}

export interface QualifiedExample {
  kind: 'qualified';
  value: any;
  type: QualifiedType;
}

export interface ResponseEnvelopeExample {
  response: result.ResponseEnvelope;
  headers: Array<ResponseHeaderExample>;
  result: ExampleType;
}

export interface ResponseHeaderExample {
  header: result.HeaderResponse | result.HeaderMapResponse;
  value: ExampleType;
}

export interface StringExample {
  kind: 'string';
  value: string;
  type: ConstantType | BytesType | LiteralValue | TimeType | PrimitiveType;
}

export interface StructExample {
  kind: 'model';
  value: Record<string, ExampleType>;
  additionalProperties?: Record<string, ExampleType>;
  type: ModelType | PolymorphicType;
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

export class AnyExample implements AnyExample {
  constructor(value: any) {
    this.kind = 'any';
    this.value = value;
    this.type = new PrimitiveType('any');
  }
}

export class ArrayExample implements ArrayExample {
  constructor(type: SliceType) {
    this.kind = 'array';
    this.type = type;
    this.value = [];
  }
}

export class BooleanExample implements BooleanExample {
  constructor(value: boolean, type: ConstantType | LiteralValue | PrimitiveType) {
    this.kind = 'boolean';
    this.value = value;
    this.type = type;
  }
}

export class DictionaryExample implements DictionaryExample {
  constructor(type: MapType) {
    this.kind = 'dictionary';
    this.type = type;
    this.value = {};
  }
}

export class MethodExample implements MethodExample {
  constructor(name: string, docs: Docs, filePath: string) {
    this.name = name;
    this.docs = docs;
    this.filePath = filePath;
    this.parameters = [];
    this.optionalParamsGroup = [];
  }
}

export class NullExample implements NullExample {
  constructor(type: PossibleType) {
    this.kind = 'null';
    this.type = type;
  }
}

export class NumberExample implements NumberExample {
  constructor(value: number, type: ConstantType | LiteralValue | TimeType | PrimitiveType) {
    this.kind = 'number';
    this.value = value;
    this.type = type;
  }
}

export class ParameterExample implements ParameterExample {
  constructor(parameter: param.Parameter, value: ExampleType) {
    this.parameter = parameter;
    this.value = value;
  }
}

export class QualifiedExample implements QualifiedExample {
  constructor(type: QualifiedType, value: any) {
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
  constructor(header: result.HeaderResponse | result.HeaderMapResponse, value: ExampleType) {
    this.header = header;
    this.value = value;
  }
}

export class StringExample implements StringExample {
  constructor(value: string, type: ConstantType | BytesType | LiteralValue | TimeType | PrimitiveType) {
    this.kind = 'string';
    this.value = value;
    this.type = type;
  }
}

export class StructExample implements StructExample {
  constructor(type: ModelType | PolymorphicType) {
    this.kind = 'model';
    this.type = type;
    this.value = {};
  }
}
