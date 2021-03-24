/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { ArraySchema, DictionarySchema, ObjectSchema, Operation, Parameter, Response, Schema, SchemaResponse, SchemaType } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { ensureNameCase } from '../transform/namer';

// variable to be used to determine comment length when calling comment from @azure-tools
export const commentLength = 150;

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

// returns ArraySchema type predicate if the schema is an ArraySchema
export function isArraySchema(resp: Schema): resp is ArraySchema {
  return resp.type === SchemaType.Array;
}

// returns DictionarySchema type predicate if the schema is a DictionarySchema
export function isDictionarySchema(resp: Schema): resp is DictionarySchema {
  return resp.type === SchemaType.Dictionary;
}

// returns SchemaResponse type predicate if the response has a schema
export function isSchemaResponse(resp?: Response): resp is SchemaResponse {
  return (resp as SchemaResponse).schema !== undefined;
}

export interface PagerInfo {
  name: string;
  respEnv: string;   // name of the response envelope
  respType: string;  // name of the response type
  respField: string; // name of the response type field within the response envelope
  nextLink: string;  // name of the next link field within the response envelope
  hasLRO: boolean;   // true if this pager is used with an LRO
}

// returns the type name of the internal pager type
export function internalPagerTypeName(pi: PagerInfo): string {
  return ensureNameCase(pi.name, true);
}

// returns true if the operation is pageable
export function isPageableOperation(op: Operation): boolean {
  return op.language.go!.paging && op.language.go!.paging.nextLinkName !== null;
}

export interface PollerInfo {
  name: string;
  respEnv: string;    // name of the response envelope
  respField?: string; // name of the response type field within the response envelope
  respType?: Schema;  // schema of the response type
  pager?: PagerInfo;
}

// returns the type name of the internal poller type
export function internalPollerTypeName(pi: PollerInfo): string {
  return ensureNameCase(pi.name, true);
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

// returns true if the object contains a property that's a discriminated type
export function hasPolymorphicField(obj: ObjectSchema): boolean {
  let found = false;
  for (const prop of values(obj.properties)) {
    if (isObjectSchema(prop.schema)) {
      found = prop.schema.discriminator !== undefined;
    } else if (isArraySchema(prop.schema) && isObjectSchema(prop.schema.elementType)) {
      found = prop.schema.elementType.discriminator !== undefined;
    }
  }
  return found;
}

// returns the object's position in an inheritence hierarchy
export function getRelationship(obj: ObjectSchema): 'none' | 'root' | 'parent' | 'leaf' {
  let hasParent = false;
  for (const parent of values(obj.parents?.immediate)) {
    if (isObjectSchema(parent)) {
      hasParent = true;
      break;
    }
  }
  let hasChild = false;
  for (const child of values(obj.children?.immediate)) {
    if (isObjectSchema(child)) {
      hasChild = true;
      break;
    }
  }
  if (!hasParent && !hasChild) {
    return 'none';
  } else if (!hasChild) {
    return 'leaf';
  } else if (!hasParent) {
    return 'root';
  } else {
    return 'parent';
  }
}

// returns the schema response for this operation.
// calling this on multi-response operations will result in an error.
export function getResponse(op: Operation): SchemaResponse | undefined {
  if (!op.responses) {
    return undefined;
  }
  // get the list and count of distinct schema responses
  const schemaResponses = new Array<SchemaResponse>();
  for (const response of values(op.responses)) {
    // perform the comparison by name as some responses have different objects for the same underlying response type
    if (isSchemaResponse(response) && !values(schemaResponses).where(sr => sr.schema.language.go!.name === response.schema.language.go!.name).any()) {
      schemaResponses.push(response);
    }
  }
  if (schemaResponses.length === 0) {
    return undefined;
  } else if (schemaResponses.length === 1) {
    return schemaResponses[0];
  }
  // multiple schema responses, for LROs find the best fit.
  if (!isLROOperation(op)) {
    throw new Error('getResponse() called for multi-response operation');
  }
  // for LROs, there are a couple of corner-cases we need to handle WRT response types.
  // 1. 200 Foo / 20x Bar - we take Foo and display a warning
  // 2. 201 Foo / 202 Bar - this is a hard error
  // 3. 200 void / 20x Bar - we take Bar
  // since we always assume responses[0] has the return type we need to fix up
  // the list of responses so that it points to the schema we select.

  // multiple schemas, find the one for 200 status code
  // note that case #3 was handled earlier
  let with200: SchemaResponse | undefined;
  for (const response of values(schemaResponses)) {
    if ((<Array<string>>response.protocol.http!.statusCodes).indexOf('200') > -1) {
      with200 = response;
      break;
    }
  }
  if (with200 === undefined) {
    // case #2
    throw new Error(`LRO ${op.language.go!.clientName}.${op.language.go!.name} contains multiple response types which is not supported`);
  }
  // case #1
  // TODO: log warning
  return with200;
}
