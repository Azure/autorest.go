/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { ArraySchema, DictionarySchema, ObjectSchema, Operation, Parameter, Response, Schema, SchemaResponse, SchemaType } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';

// aggregates the Parameter in op.parameters and the first request
export function aggregateParameters(op: Operation): Array<Parameter> {
  if (op.requests!.length > 1) {
    throw console.error('multiple requests NYI');
  }
  let params = new Array<Parameter>();
  if (op.parameters) {
    params = params.concat(op.parameters);
  }
  if (op.requests![0].parameters) {
    params = params.concat(op.requests![0].parameters);
  }
  return params;
}

// returns ArraySchema type predicate if the schema is an ArraySchema
export function isArraySchema(resp: Schema): resp is ArraySchema {
  return resp.type === SchemaType.Array;
}

// returns SchemaResponse type predicate if the response has a schema
export function isSchemaResponse(resp?: Response): resp is SchemaResponse {
  return (resp as SchemaResponse).schema !== undefined;
}

export interface PagerInfo {
  name: string;
  op: Operation;
  respField: boolean;
}

// returns true if the operation is pageable
export function isPageableOperation(op: Operation): boolean {
  return op.language.go!.paging && op.language.go!.paging.nextLinkName !== null;
}

export interface PollerInfo {
  name: string;
  op: Operation;
}

// returns true if the operation is a long-running operation
export function isLROOperation(op: Operation): boolean {
  return op.extensions?.['x-ms-long-running-operation'] === true;
}

// returns ObjectSchema type predicate if the schema is an ObjectSchema
export function isObjectSchema(obj: Schema): obj is ObjectSchema {
  return obj.type === SchemaType.Object;
}

// returns the additional properties schema if the ObjectSchema defines 'additionalProperties'
export function hasAdditionalProperties(obj: ObjectSchema): DictionarySchema | undefined {
  for (const parent of values(obj.parents?.immediate)) {
    if (parent.type === SchemaType.Dictionary) {
      return <DictionarySchema>parent;
    }
  }
  return undefined;
}

export function isScalar(schema: Schema): boolean {
  switch (schema.type) {
    case SchemaType.Array:
    case SchemaType.Boolean:
    case SchemaType.ByteArray:
    case SchemaType.Integer:
    case SchemaType.Number:
    case SchemaType.String:
      return true;
    default:
      return false;
  }
}