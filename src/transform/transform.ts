/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { KnownMediaType, pascalCase, serialize } from '@azure-tools/codegen';
import { Host, startSession, Session } from '@azure-tools/autorest-extension-base';
import { ObjectSchema, ArraySchema, codeModelSchema, CodeModel, ImplementationLocation, Language, SchemaType, NumberSchema, Operation, SchemaResponse, Parameter, Property, Protocols, Response, Schema, DictionarySchema, Protocol } from '@azure-tools/codemodel';
import { length, values } from '@azure-tools/linq';
import { aggregateParameters, ParamInfo, paramInfo } from '../generator/common/helpers';

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
  // fix up enum types
  for (const choice of values(session.model.schemas.choices)) {
    choice.choiceType.language.go!.name = 'string';
  }
  for (const choice of values(session.model.schemas.sealedChoices)) {
    choice.choiceType.language.go!.name = 'string';
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
      for (const param of values(aggregateParameters(op))) {
        // skip the host param as we use our own url.URL instead
        if (param.language.go!.name === 'host' || param.language.go!.name === '$host') {
          continue;
        }
        param.schema.language.go!.name = schemaTypeToGoType(param.schema);
        if (param.implementation === ImplementationLocation.Client) {
          // add global param info to the operation group
          if (group.language.go!.globals === undefined) {
            group.language.go!.globals = new Array<ParamInfo>();
          }
          const globals = <Array<ParamInfo>>group.language.go!.globals;
          // check if this global param has already been added
          const index = globals.findIndex((value: ParamInfo, index: Number, obj: ParamInfo[]) => {
            if (value.name === param.language.go!.name) {
              return true;
            }
            return false;
          });
          if (index === -1) {
            globals.push(new paramInfo(param.language.go!.name, param.schema.language.go!.name, true, param.required === true));
          }
        }
      }
      // recursively add the marshalling format to the body param if applicable
      const marshallingFormat = getMarshallingFormat(op.requests![0].protocol);
      if (marshallingFormat !== 'na') {
        const bodyParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http!.in === 'body'; }).first();
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
    newProperty('RawResponse', 'RawResponse contains the underlying HTTP response.', newObject('http.Response', 'TODO'))
  ];
  // if the response defines a schema then add it as a field to the response type
  if (isSchemaResponse(resp)) {
    // for operations that return scalar types we use a fixed field name 'Value'
    let propName = 'Value';
    if (resp.schema.type === SchemaType.Object) {
      // for object types use the type's name as the field name
      propName = resp.schema.language.go!.name;
    } else if (resp.schema.type === SchemaType.Array) {
      // for array types use the element type's name
      propName = (<ArraySchema>resp.schema).elementType.language.go!.name;
    }
    if (resp.schema.serialization?.xml && resp.schema.serialization.xml.name) {
      // always prefer the XML name
      propName = pascalCase(resp.schema.serialization.xml.name);
    }
    resp.schema.language.go!.name = schemaTypeToGoType(resp.schema);
    resp.schema.language.go!.responseValue = propName;
    (<Array<Property>>resp.language.go!.properties).push(newProperty(propName, resp.schema.language.go!.description, resp.schema));
  }
}

function newObject(name: string, desc: string): ObjectSchema {
  let obj = new ObjectSchema(name, desc);
  obj.language.go = obj.language.default;
  return obj;
}

function newProperty(name: string, desc: string, schema: Schema): Property {
  let prop = new Property(name, desc, schema);
  prop.language.go = prop.language.default;
  return prop;
}

function isSchemaResponse(resp?: Response): resp is SchemaResponse {
  return (resp as SchemaResponse).schema !== undefined;
}

// returns the format used for marshallling/unmarshalling.
// if the media type isn't applicable then 'na' is returned.
function getMarshallingFormat(protocol: Protocols): 'json' | 'xml' | 'na' {
  switch ((<Protocol>protocol).http.knownMediaType) {
    case KnownMediaType.Json:
      return 'json';
    case KnownMediaType.Xml:
      return 'xml';
    default:
      return 'na';
  }
}
