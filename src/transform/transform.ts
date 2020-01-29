/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { serialize } from '@azure-tools/codegen';
import { Host, startSession, Session } from '@azure-tools/autorest-extension-base';
import { ArraySchema, codeModelSchema, CodeModel, Language, SchemaType, NumberSchema, Operation, OperationGroup, SchemaResponse, Property, Response, Schema } from '@azure-tools/codemodel';
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
  processOperationResponses(session);
  // fix up struct field types
  for (const obj of values(session.model.schemas.objects)) {
    for (const prop of values(obj.properties)) {
      const details = <Language>prop.schema.language.go;
      details.name = `${schemaTypeToGoType(prop.schema)}`;
    }
  }
}

function schemaTypeToGoType(schema: Schema): string {
  switch (schema.type) {
    case SchemaType.Array:
      return `[]${(<ArraySchema>schema).elementType.language.go!.name}`;
    case SchemaType.Boolean:
      return 'bool';
    case SchemaType.ByteArray:
      return '[]byte';
    case SchemaType.Date:
    case SchemaType.DateTime:
      return 'time.Time';
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
    default:
      return schema.language.go!.name;
  }
}

function processOperationResponses(session: Session<CodeModel>) {
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      createResponseType(group, op);
      // annotate all exception types as errors; this is so we know to generate an Error() method
      for (const ex of values(op.exceptions)) {
        (<SchemaResponse>ex).schema.language.go!.errorType = true;
      }
    }
  }
}

// creates the response type to be returned from an operation and updates the operation
function createResponseType(group: OperationGroup, op: Operation) {
  if (length(op.responses) > 1) {
    throw console.error('multiple responses NYI');
  }
  // create the `type FooResponse struct` response
  // type with a `StatusCode int` field
  const resp = op.responses![0];
  resp.language.go!.properties = [
    newProperty('StatusCode', 'StatusCode contains the HTTP status code.', newNumber('int', 'TODO', SchemaType.Integer, 32))
  ];
  // if the response defines a schema then add it as a field to the response type
  if (isSchemaResponse(resp)) {
    resp.schema.language.go!.name = schemaTypeToGoType(resp.schema);
    (<Array<Property>>resp.language.go!.properties).push(newProperty('Value', resp.schema.language.go!.description, resp.schema));
  }
}

function newNumber(name: string, desc: string, type: SchemaType.Integer | SchemaType.Number, precisioin: number): NumberSchema {
  let num = new NumberSchema(name, desc, type, precisioin);
  num.language.go = num.language.default;
  return num;
}

function newProperty(name: string, desc: string, schema: Schema): Property {
  let prop = new Property(name, desc, schema);
  prop.language.go = prop.language.default;
  return prop;
}

function isSchemaResponse(resp?: Response): resp is SchemaResponse {
  return (resp as SchemaResponse).schema !== undefined;
}
