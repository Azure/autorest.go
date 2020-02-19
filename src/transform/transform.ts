/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { KnownMediaType, serialize } from '@azure-tools/codegen';
import { Host, startSession, Session } from '@azure-tools/autorest-extension-base';
import { ObjectSchema, ArraySchema, codeModelSchema, CodeModel, Language, SchemaType, NumberSchema, Operation, SchemaResponse, Parameter, Property, Protocols, Response, Schema, DictionarySchema } from '@azure-tools/codemodel';
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
      const details = <Language>prop.schema.language.go;
      details.name = `${schemaTypeToGoType(prop.schema)}`;
    }
  }
}

function schemaTypeToGoType(schema: Schema): string {
  switch (schema.type) {
    case SchemaType.Any:
      return 'interface{}';
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
    case SchemaType.String:
    case SchemaType.Uuid:
      return 'string';
    default:
      return schema.language.go!.name;
  }
}

function recursiveAddMarshallingFormat(schema: Schema, marshallingFormat: 'json' | 'xml') {
  if (schema.language.go!.marshallingFormat) {
    // this schema has already been processed, don't do it again
    return;
  }
  schema.language.go!.marshallingFormat = marshallingFormat;
  switch (schema.type) {
    case SchemaType.Array:
      const arraySchema = <ArraySchema>schema;
      recursiveAddMarshallingFormat(arraySchema.elementType, marshallingFormat);
      break;
    case SchemaType.Dictionary:
      const dictSchema = <DictionarySchema>schema;
      recursiveAddMarshallingFormat(dictSchema.elementType, marshallingFormat);
      break;
    case SchemaType.Object:
      const os = <ObjectSchema>schema;
      for (const prop of values(os.properties)) {
        recursiveAddMarshallingFormat(prop.schema, marshallingFormat);
      }
      // if this is a discriminated type, update children and parents
      for (const child of values(os.children?.all)) {
        recursiveAddMarshallingFormat(child, marshallingFormat);
      }
      for (const parent of values(os.parents?.all)) {
        recursiveAddMarshallingFormat(parent, marshallingFormat);
      }
      break;
  }
}

// we will transform operation request parameter schema types to Go types
function processOperationRequests(session: Session<CodeModel>) {
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      for (const param of values(op.request.parameters)) {
        param.schema.language.go!.name = schemaTypeToGoType(param.schema);
      }
      // recursively add the marshalling format to the body param if applicable
      const marshallingFormat = getMarshallingFormat(op.request.protocol);
      if (marshallingFormat !== 'na') {
        const bodyParam = values(op.request.parameters).where((each: Parameter) => { return each.protocol.http!.in === 'body'; }).first();
        if (bodyParam) {
          recursiveAddMarshallingFormat(bodyParam.schema, marshallingFormat);
        }
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
        const marshallingFormat = getMarshallingFormat(ex.protocol);
        if (marshallingFormat === 'na') {
          throw console.error(`unexpected media type none for ${ex.language.go!.name} error type`);
        }
        (<SchemaResponse>ex).schema.language.go!.errorType = true;
        recursiveAddMarshallingFormat((<SchemaResponse>ex).schema, marshallingFormat);
      }
      // recursively add the marshalling format to the responses if applicable
      for (const resp of values(op.responses)) {
        const marshallingFormat = getMarshallingFormat(resp.protocol);
        if (marshallingFormat !== 'na' && isSchemaResponse(resp)) {
          recursiveAddMarshallingFormat(resp.schema, marshallingFormat);
        }
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
  resp.language.go!.responseType = true;
  resp.language.go!.properties = [
    newProperty('StatusCode', 'StatusCode contains the HTTP status code.', newNumber('int', 'TODO', SchemaType.Integer, 32), true)
  ];
  // if the response defines a schema then add it as a field to the response type
  if (isSchemaResponse(resp)) {
    // for operations that return scalar types we use a fixed field name 'Value'
    let propName = 'Value';
    if (resp.schema.type === SchemaType.Object) {
      // for object types use the type's name as the field name
      propName = resp.schema.language.go!.name;
    }
    resp.schema.language.go!.name = schemaTypeToGoType(resp.schema);
    resp.schema.language.go!.responseValue = propName;
    (<Array<Property>>resp.language.go!.properties).push(newProperty(propName, resp.schema.language.go!.description, resp.schema, false));
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

// returns the format used for marshallling/unmarshalling.
// if the media type isn't applicable then 'na' is returned.
function getMarshallingFormat(protocol: Protocols): 'json' | 'xml' | 'na' {
  switch (protocol.http!.knownMediaType) {
    case KnownMediaType.Json:
      return 'json';
    case KnownMediaType.Xml:
      return 'xml';
    default:
      return 'na';
  }
}
