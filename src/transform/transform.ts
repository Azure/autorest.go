/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { serialize } from '@azure-tools/codegen';
import { Host, startSession, Session } from '@azure-tools/autorest-extension-base';
import { ArraySchema, codeModelSchema, CodeModel, Language, SchemaType, NumberSchema, Operation, SchemaResponse, Property, Response, Schema, DictionarySchema, ConstantSchema, ObjectSchema } from '@azure-tools/codemodel';
import { length, values } from '@azure-tools/linq';

// The transformer adds Go-specific information to the code model.
export async function transform(host: Host) {
  const debug = await host.GetValue('debug') || false;

  try {
    const session = await startSession<CodeModel>(host, {}, codeModelSchema);

    await process(session);

    // output the model to the pipeline
    host.WriteFile('code-model-v4.yaml', serialize(session.model), undefined, 'code-model-v4');

  } catch (E) {
    if (debug) {
      console.error(`${__filename} - FAILURE  ${JSON.stringify(E)} ${E.stack}`);
    }
    throw E;
  }
}

async function process(session: Session<CodeModel>) {
  processOperationRequests(session);
  processOperationResponses(session);
  // fix up struct field types
  for (const obj of values(session.model.schemas.objects)) {
    for (const prop of values(obj.properties)) {
      if (noByRef(prop.schema.type)) {
        prop.language.go!.noByRef = true;
      }
      const details = <Language>prop.schema.language.go;
      details.name = `${schemaTypeToGoType(prop.schema)}`;
    }
  }
}

function schemaTypeToGoType(schema: Schema): string {
  switch (schema.type) {
    case SchemaType.Array:
      const arraySchema = <ArraySchema>schema;
      const arrayElem = <Schema>arraySchema.elementType;
      return `[]${schemaTypeToGoType(arrayElem)}`;
    case SchemaType.Boolean:
      return 'bool';
    case SchemaType.ByteArray:
      return '[]byte';
    case SchemaType.Date:
    case SchemaType.DateTime:
    case SchemaType.UnixTime:
      return 'time.Time';
    case SchemaType.Dictionary:
      const dictSchema = <DictionarySchema>schema;
      const dictElem = <Schema>dictSchema.elementType;
      return `map[string]*${schemaTypeToGoType(dictElem)}`;
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
    case SchemaType.Any:
    case SchemaType.String:
    case SchemaType.Uuid:
      return 'string';
    default:
      return schema.language.go!.name;
  }
}

// we will transform operation request parameter schema types to Go types
function processOperationRequests(session: Session<CodeModel>) {
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      for (const param of values(op.request.parameters)) {
        param.schema.language.go!.name = schemaTypeToGoType(param.schema);
      }
    }
  }
}

function processOperationResponses(session: Session<CodeModel>) {
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      createResponseType(op);
      // annotate all exception types as errors; this is so we know to generate an Error() method
      for (const ex of values(op.exceptions)) {
        (<SchemaResponse>ex).schema.language.go!.errorType = true;
      }
    }
  }
}

// creates the response type to be returned from an operation and updates the operation
function createResponseType(op: Operation) {
  if (length(op.responses) > 1) {
    throw console.error('multiple responses NYI');
  }
  // create the `type FooResponse struct` response
  // type with a `StatusCode int` field
  const resp = op.responses![0];
  resp.language.go!.properties = [
    newProperty('StatusCode', 'StatusCode contains the HTTP status code.', newNumber('int', 'TODO', SchemaType.Integer, 32), true)
  ];
  // if the response defines a schema then add it as a field to the response type
  if (isSchemaResponse(resp)) {
    resp.schema.language.go!.name = schemaTypeToGoType(resp.schema);
    (<Array<Property>>resp.language.go!.properties).push(newProperty('Value', resp.schema.language.go!.description, resp.schema, noByRef(resp.schema.type)));
  }
}

function newNumber(name: string, desc: string, type: SchemaType.Integer | SchemaType.Number, precisioin: number): NumberSchema {
  let num = new NumberSchema(name, desc, type, precisioin);
  num.language.go = num.language.default;
  return num;
}

function newProperty(name: string, desc: string, schema: Schema, noByRef?: boolean): Property {
  let prop = new Property(name, desc, schema);
  prop.language.go = prop.language.default;
  prop.language.go.noByRef = noByRef;
  return prop;
}

function isSchemaResponse(resp?: Response): resp is SchemaResponse {
  return (resp as SchemaResponse).schema !== undefined;
}

// returns true if the type should not be a pointer-to-type
function noByRef(type: SchemaType): boolean {
  return type === SchemaType.Array || type === SchemaType.ByteArray || type === SchemaType.Dictionary;
}
