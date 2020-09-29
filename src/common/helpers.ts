/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { ArraySchema, CodeModel, DictionarySchema, ObjectSchema, Operation, Parameter, Response, Schema, SchemaResponse, SchemaType } from '@azure-tools/codemodel';
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
  hasLRO: boolean;
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

// returns true if client types should be exported
export async function exportClients(session: Session<CodeModel>): Promise<boolean> {
  const specType = await session.getValue('openapi-type', 'not_specified');
  return specType === 'arm';
}
