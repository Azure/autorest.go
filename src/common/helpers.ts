/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { ArraySchema, ObjectSchema, Operation, Parameter, Response, Schema, SchemaResponse } from '@azure-tools/codemodel';

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
  return (resp as ArraySchema).elementType !== undefined;
}

// returns SchemaResponse type predicate if the response has a schema
export function isSchemaResponse(resp?: Response): resp is SchemaResponse {
  return (resp as SchemaResponse).schema !== undefined;
}

export interface PagerInfo {
  name: string;
  schema: Schema;
  client: string;
  nextLink: string;
}

// returns true if the operation is pageable
export function isPageableOperation(op: Operation): boolean {
  return op.language.go!.paging && op.language.go!.paging.nextLinkName !== null;
}

export interface PollerInfo {
  name: string;
  schema: Schema;
  client: string;
  pollingMethod: string; // adding method to determine polling tracker type
}

// returns true if the operation is a long-running operation
export function isLROOperation(op: Operation): boolean {
  return op.extensions?.['x-ms-long-running-operation'];
}

// returns ObjectSchema type predicate if the schema is an ObjectSchema
export function isObjectSchema(obj: Schema): obj is ObjectSchema {
  return (obj as ObjectSchema).properties !== undefined;
}
