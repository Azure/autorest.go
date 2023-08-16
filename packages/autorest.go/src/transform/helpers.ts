/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { ArraySchema, BinaryResponse, ConstantSchema, DictionarySchema, ObjectSchema, Operation, Parameter, Property, Response, Schema, SchemaContext, SchemaResponse, SchemaType } from '@autorest/codemodel';
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

// returns DictionarySchema type predicate if the schema is a DictionarySchema
export function isDictionarySchema(resp: Schema): resp is DictionarySchema {
  return resp.type === SchemaType.Dictionary;
}

// returns SchemaResponse type predicate if the response has a schema
export function isSchemaResponse(resp: Response): resp is SchemaResponse {
  return (<SchemaResponse>resp).schema !== undefined;
}

// returns BinaryResponse type predicate if the response is a binary response
export function isBinaryResponse(resp: Response): resp is BinaryResponse {
  return (<BinaryResponse>resp).binary !== undefined;
}

// returns true if the operation is pageable
export function isPageableOperation(op: Operation): boolean {
  return op.language.go!.paging;
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
  for (const prop of values(obj.properties)) {
    if (isObjectSchema(prop.schema)) {
      if (prop.schema.discriminator !== undefined) {
        return true;
      }
    } else if (isArraySchema(prop.schema) && isObjectSchema(prop.schema.elementType)) {
      if (prop.schema.elementType.discriminator !== undefined) {
        return true;
      }
    } else if (isDictionarySchema(prop.schema) && isObjectSchema(prop.schema.elementType)) {
      if (prop.schema.elementType.discriminator !== undefined) {
        return true;
      }
    }
  }
  return false;
}

// returns the schema response for this operation.
// calling this on multi-response operations will result in an error.
export function getSchemaResponse(op: Operation): SchemaResponse | undefined {
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
    throw new Error('getSchemaResponse() called for multi-response operation');
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

// returns the BinaryResponse if the operation returns a binary response or undefined.
// throws an error if called on a multi-response operation.
export function isBinaryResponseOperation(op: Operation): BinaryResponse | undefined {
  if (!op.responses) {
    return undefined;
  } else if (isMultiRespOperation(op)) {
    throw new Error('isBinaryResponseOperation() called for multi-response operation');
  }
  if (isBinaryResponse(op.responses[0])) {
    return op.responses[0];
  }
  return undefined;
}

// returns true if the type is implicitly passed by value (map, slice, etc)
export function isTypePassedByValue(schema: Schema): boolean {
  if (schema.type === SchemaType.Any ||
    schema.type === SchemaType.AnyObject ||
    schema.type === SchemaType.Array ||
    schema.type === SchemaType.ByteArray ||
    schema.type === SchemaType.Binary ||
    schema.type === SchemaType.Dictionary ||
    (isObjectSchema(schema) && schema.discriminator)) {
    return true;
  }
  return false;
}

// returns a Go-formatted constant value
export function formatConstantValue(schema: ConstantSchema): string {
  // null check must come before any type checks
  if (schema.value.value === null) {
    return 'nil';
  } else if (schema.valueType.type === SchemaType.String) {
    return `"${schema.value.value}"`;
  }
  return schema.value.value;
}

//  returns true if the object is used for output only
export function isOutputOnly(obj: ObjectSchema): boolean {
  return !values(obj.usage).any((u) => { return u === SchemaContext.Input; });
}

// aggregate the properties from the provided type and its parent types
export function aggregateProperties(obj: ObjectSchema): Array<Property> {
  const allProps = new Array<Property>();
  for (const prop of values(obj.properties)) {
    allProps.push(prop);
  }
  for (const parent of values(obj.parents?.all)) {
    if (isObjectSchema(parent)) {
      for (const parentProp of values(parent.properties)) {
        // ensure that the parent doesn't contain any properties with the same name but different type
        const exists = values(allProps).where(p => { return p.language.go!.name === parentProp.language.go!.name; }).first();
        if (exists) {
          if (exists.schema.language.go!.name !== parentProp.schema.language.go!.name) {
            const msg = `type ${obj.language.go!.name} contains duplicate property ${exists.language.go!.name} with mismatched types`;
            throw new Error(msg);
          }
          // don't add the duplicate
          continue;
        }
        allProps.push(parentProp);
      }
    }
  }
  return allProps;
}

// returns the leaf element type for an array or dictionary
export function recursiveUnwrapArrayDictionary(item: ArraySchema | DictionarySchema): Schema {
  if (isArraySchema(item.elementType) || isDictionarySchema(item.elementType)) {
    return recursiveUnwrapArrayDictionary(item.elementType);
  }
  return item.elementType;
}
