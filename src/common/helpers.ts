/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { ArraySchema, CodeModel, ConstantSchema, DateTimeSchema, DictionarySchema, NumberSchema, ObjectSchema, Operation, Parameter, Response, Schema, SchemaResponse, SchemaType } from '@azure-tools/codemodel';
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
  return (resp as ArraySchema).elementType !== undefined;
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
  return (obj as ObjectSchema).properties !== undefined;
}

// returns the additional properties schema if the ObjectSchema defines 'additionalProperties'
export function hasAdditionalProperties(obj: ObjectSchema): Schema | undefined {
  for (const parent of values(obj.parents?.immediate)) {
    if (parent.type === SchemaType.Dictionary) {
      return parent;
    }
  }
  return undefined;
}

export function schemaTypeToGoType(codeModel: CodeModel, schema: Schema, inBody: boolean): string {
  switch (schema.type) {
    case SchemaType.Any:
      return 'interface{}';
    case SchemaType.Array:
      const arraySchema = <ArraySchema>schema;
      const arrayElem = <Schema>arraySchema.elementType;
      arrayElem.language.go!.name = schemaTypeToGoType(codeModel, arrayElem, inBody);
      return `[]${arrayElem.language.go!.name}`;
    case SchemaType.Binary:
      return 'azcore.ReadSeekCloser';
    case SchemaType.Boolean:
      return 'bool';
    case SchemaType.ByteArray:
      return '[]byte';
    case SchemaType.Constant:
      let constSchema = <ConstantSchema>schema;
      constSchema.valueType.language.go!.name = schemaTypeToGoType(codeModel, constSchema.valueType, inBody);
      return constSchema.valueType.language.go!.name;
    case SchemaType.DateTime:
      // header/query param values are parsed separately so they don't need custom types
      if (inBody) {
        // add a marker to the code model indicating that we need
        // to include support for marshalling/unmarshalling time.
        const dateTime = <DateTimeSchema>schema;
        if (dateTime.format === 'date-time-rfc1123') {
          codeModel.language.go!.hasTimeRFC1123 = true;
          schema.language.go!.internalTimeType = 'timeRFC1123';
        } else {
          codeModel.language.go!.hasTimeRFC3339 = true;
          schema.language.go!.internalTimeType = 'timeRFC3339';
        }
      }
    case SchemaType.Date:
      return 'time.Time';
    case SchemaType.UnixTime:
      codeModel.language.go!.hasUnixTime = true;
      if (inBody) {
        schema.language.go!.internalTimeType = 'timeUnix';
      }
      return 'time.Time';
    case SchemaType.Dictionary:
      const dictSchema = <DictionarySchema>schema;
      const dictElem = <Schema>dictSchema.elementType;
      dictElem.language.go!.name = schemaTypeToGoType(codeModel, dictElem, inBody);
      return `map[string]${dictElem.language.go!.name}`;
    case SchemaType.Duration:
      return 'time.Duration';
    case SchemaType.Integer:
      if ((<NumberSchema>schema).precision === 32) {
        return 'int32';
      }
      return 'int64';
    case SchemaType.Number:
      if ((<NumberSchema>schema).precision === 32) {
        return 'float32';
      }
      return 'float64';
    case SchemaType.String:
    case SchemaType.Uuid:
      return 'string';
    case SchemaType.Uri:
      return 'url.URL';
    default:
      return schema.language.go!.name;
  }
}
