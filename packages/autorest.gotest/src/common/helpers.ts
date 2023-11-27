/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Operation, Parameter, Response, SchemaResponse } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';

// aggregates the Parameter in op.parameters and the first request
export function aggregateParameters(op: Operation): Array<Parameter> {
  let params = new Array<Parameter>();
  if (op.parameters) {
    params = params.concat(op.parameters);
  }
  // Loop through each request in an operation to account for all parameters in the initial naming transform.
  // After the transform stage, operations will only have one request and the loop will always traverse only
  // one request per operation.
  for (const req of values(op.requests)) {
    if (req.parameters) {
      params = params.concat(req.parameters);
    }
  }
  return params;
}

// returns true if the operation is a long-running operation
export function isLROOperation(op: Operation): boolean {
  return op.extensions?.['x-ms-long-running-operation'] === true;
}

// returns true if the operation is pageable
export function isPageableOperation(op: Operation): boolean {
  return op.language.go!.paging;
}

// returns SchemaResponse type predicate if the response has a schema
export function isSchemaResponse(resp: Response): resp is SchemaResponse {
  return (<SchemaResponse>resp).schema !== undefined;
}

// used to sort strings in ascending order
export function sortAscending(a: string, b: string): number {
  return a < b ? -1 : a > b ? 1 : 0;
}

// sorts parameters by their required state, ordering required before optional
export function sortParametersByRequired(a: Parameter, b: Parameter): number {
  if (a.required === b.required) {
    return 0;
  }
  if (a.required && !b.required) {
    return -1;
  }
  return 1;
}

// returns true if the operation returns multiple response types
export function isMultiRespOperation(op: Operation): boolean {
  // treat LROs as single-response ops
  if (!op.responses || op.responses.length === 1 || isLROOperation(op)) {
    return false;
  }
  // count the number of distinct schemas returned by this operation
  const schemaResponses = new Array<SchemaResponse>();
  for (const response of values(op.responses)) {
    // perform the comparison by name as some responses have different objects for the same underlying response type
    if (isSchemaResponse(response) && !values(schemaResponses).where(sr => sr.schema.language.go!.name === response.schema.language.go!.name).any()) {
      schemaResponses.push(response);
    }
  }
  return schemaResponses.length > 1;
}
