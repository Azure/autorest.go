/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { KnownMediaType, pascalCase, serialize } from '@azure-tools/codegen';
import { Host, startSession, Session } from '@azure-tools/autorest-extension-base';
import { ObjectSchema, ArraySchema, codeModelSchema, CodeModel, DateTimeSchema, HttpHeader, HttpResponse, ImplementationLocation, Language, SchemaType, NumberSchema, Operation, SchemaResponse, Parameter, Property, Protocols, Response, Schema, DictionarySchema, Protocol, ChoiceSchema } from '@azure-tools/codemodel';
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
      details.name = `${schemaTypeToGoType(session.model, prop.schema, true)}`;
    }
  }
  // fix up enum types
  for (const choice of values(session.model.schemas.choices)) {
    choice.choiceType.language.go!.name = 'string';
  }
  for (const choice of values(session.model.schemas.sealedChoices)) {
    // TODO need to see how to add sealed-choices that have a different schema
    if (choice.choices.length === 1) {
      continue;
    }
    choice.choiceType.language.go!.name = 'string';
  }
}

function schemaTypeToGoType(codeModel: CodeModel, schema: Schema, inBody: boolean): string {
  switch (schema.type) {
    case SchemaType.Any:
      return 'interface{}';
    case SchemaType.Array:
      const arraySchema = <ArraySchema>schema;
      const arrayElem = <Schema>arraySchema.elementType;
      return `[]${schemaTypeToGoType(codeModel, arrayElem, inBody)}`;
    case SchemaType.Binary:
      return 'azcore.ReadSeekCloser';
    case SchemaType.Boolean:
      return 'bool';
    case SchemaType.ByteArray:
      return '[]byte';
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
    case SchemaType.UnixTime:
      return 'time.Time';
    case SchemaType.Dictionary:
      const dictSchema = <DictionarySchema>schema;
      const dictElem = <Schema>dictSchema.elementType;
      return `map[string]*${schemaTypeToGoType(codeModel, dictElem, inBody)}`;
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

function recursiveAddMarshallingFormat(schema: Schema, marshallingFormat: 'json' | 'xml') {
  // only recurse if the schema isn't a primitive type
  const shouldRecurse = function (schema: Schema): boolean {
    return schema.type === SchemaType.Array || schema.type === SchemaType.Dictionary || schema.type === SchemaType.Object;
  };
  if (schema.language.go!.marshallingFormat) {
    // this schema has already been processed, don't do it again
    return;
  }
  schema.language.go!.marshallingFormat = marshallingFormat;
  switch (schema.type) {
    case SchemaType.Array:
      const arraySchema = <ArraySchema>schema;
      if (shouldRecurse(arraySchema.elementType)) {
        recursiveAddMarshallingFormat(arraySchema.elementType, marshallingFormat);
      }
      break;
    case SchemaType.Dictionary:
      const dictSchema = <DictionarySchema>schema;
      if (shouldRecurse(dictSchema.elementType)) {
        recursiveAddMarshallingFormat(dictSchema.elementType, marshallingFormat);
      }
      break;
    case SchemaType.Object:
      const os = <ObjectSchema>schema;
      for (const prop of values(os.properties)) {
        if (shouldRecurse(prop.schema)) {
          recursiveAddMarshallingFormat(prop.schema, marshallingFormat);
        }
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
      if (op.requests) {
        if (op.requests.length > 1) {
          throw console.error('multiple requests NYI');
        }
        if (op.requests![0].protocol.http!.headers) {
          for (const header of values(op.requests![0].protocol.http!.headers)) {
            const head = <HttpHeader>header;
            head.schema.language.go!.name = schemaTypeToGoType(session.model, head.schema, false);
          }
        }
      }
      for (const param of values(aggregateParameters(op))) {
        // skip the host param as we use our own url.URL instead
        if (param.language.go!.name === 'host' || param.language.go!.name === '$host') {
          continue;
        }
        const inBody = param.protocol.http !== undefined && param.protocol.http!.in === 'body';
        param.schema.language.go!.name = schemaTypeToGoType(session.model, param.schema, inBody);
        if (param.implementation === ImplementationLocation.Client && param.schema.type !== SchemaType.Constant) {
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
      createResponseType(session.model, op);
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
        if (resp.protocol.http!.headers) {
          for (const header of values(resp.protocol.http!.headers)) {
            const head = <HttpHeader>header;
            head.schema.language.go!.name = schemaTypeToGoType(session.model, head.schema, false);
          }
        }
        const marshallingFormat = getMarshallingFormat(resp.protocol);
        if (marshallingFormat !== 'na' && isSchemaResponse(resp)) {
          recursiveAddMarshallingFormat(resp.schema, marshallingFormat);
        }
        // fix up schema types for header responses
        const httpResponse = <HttpResponse>resp.protocol.http;
        for (const header of values(httpResponse.headers)) {
          header.schema.language.go!.name = schemaTypeToGoType(session.model, header.schema, false);
        }
      }
    }
  }
}

// creates the response type to be returned from an operation and updates the operation
function createResponseType(codeModel: CodeModel, op: Operation) {
  // create the `type FooResponse struct` response
  // type with a `RawResponse *http.Response` field
  const firstResp = op.responses![0];
  firstResp.language.go!.responseType = true;
  firstResp.language.go!.properties = [
    newProperty('RawResponse', 'RawResponse contains the underlying HTTP response.', newObject('http.Response', 'TODO'))
  ];
  const len = op.responses!.length;
  // if the response defines a schema then add it as a field to the response type
  if (isSchemaResponse(firstResp)) {
    const marshallingFormat = getMarshallingFormat(firstResp.protocol);
    firstResp.language.go!.marshallingFormat = marshallingFormat;
    // for operations that return scalar types we use a fixed field name 'Value'
    let propName = 'Value';
    if (firstResp.schema.type === SchemaType.Object) {
      // for object types use the type's name as the field name
      propName = firstResp.schema.language.go!.name;
    } else if (firstResp.schema.type === SchemaType.Array) {
      // for array types use the element type's name
      propName = (<ArraySchema>firstResp.schema).elementType.language.go!.name;
    }
    if (firstResp.schema.serialization?.xml && firstResp.schema.serialization.xml.name) {
      // always prefer the XML name
      propName = pascalCase(firstResp.schema.serialization.xml.name);
    }
    firstResp.schema.language.go!.name = schemaTypeToGoType(codeModel, firstResp.schema, true);
    firstResp.schema.language.go!.responseValue = propName;
    (<Array<Property>>firstResp.language.go!.properties).push(newProperty(propName, firstResp.schema.language.go!.description, firstResp.schema));
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
