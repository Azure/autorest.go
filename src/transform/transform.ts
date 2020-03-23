/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { camelCase, KnownMediaType, pascalCase, serialize } from '@azure-tools/codegen';
import { Host, startSession, Session } from '@azure-tools/autorest-extension-base';
import { ObjectSchema, ArraySchema, codeModelSchema, CodeModel, DateTimeSchema, HttpHeader, HttpResponse, ImplementationLocation, Language, OperationGroup, SchemaType, NumberSchema, Operation, SchemaResponse, Parameter, Property, Protocols, Response, Schema, DictionarySchema, Protocol, ChoiceSchema, SealedChoiceSchema } from '@azure-tools/codemodel';
import { items, values } from '@azure-tools/linq';
import { aggregateParameters, isPageableOperation, isSchemaResponse, PagerInfo, ParamInfo, paramInfo } from '../generator/helpers';

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
      if (prop.schema.type === SchemaType.DateTime) {
        obj.language.go!.needsDateTimeMarshalling = true;
      }
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
      if (op.requests!.length > 1) {
        throw console.error('multiple requests NYI');
      }
      if (op.requests![0].protocol.http!.headers) {
        for (const header of values(op.requests![0].protocol.http!.headers)) {
          const head = <HttpHeader>header;
          head.schema.language.go!.name = schemaTypeToGoType(session.model, head.schema, false);
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
          if (marshallingFormat === 'xml' && bodyParam.schema.serialization?.xml?.name) {
            // mark that this parameter type will need a custom marshaller to handle the XML name
            bodyParam.schema.language.go!.xmlWrapperName = bodyParam.schema.serialization?.xml?.name;
          }
        }
      }
    }
  }
}

function processOperationResponses(session: Session<CodeModel>) {
  if (session.model.language.go!.responseSchemas === undefined) {
    session.model.language.go!.responseSchemas = new Array<Schema>();
  }
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      // annotate all exception types as errors; this is so we know to generate an Error() method
      for (const ex of values(op.exceptions)) {
        const marshallingFormat = getMarshallingFormat(ex.protocol);
        if (marshallingFormat === 'na') {
          // this is for the case where the 'default' response section
          // doesn't specify a model (legal, mostly in the test server)
          ex.language.go!.genericError = true;
          continue;
        }
        const schemaError = (<SchemaResponse>ex).schema;
        schemaError.language.go!.errorType = true;
        schemaError.language.go!.constructorName = `new${schemaError.language.go!.name}`;
        recursiveAddMarshallingFormat(schemaError, marshallingFormat);
      }
      // recursively add the marshalling format to the responses if applicable
      for (const resp of values(op.responses)) {
        if (isSchemaResponse(resp)) {
          resp.schema.language.go!.name = schemaTypeToGoType(session.model, resp.schema, true);
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
      createResponseType(session.model, group, op);
    }
  }
}

interface HttpHeaderWithDescription extends HttpHeader {
  description: string;
}

// creates the response type to be returned from an operation and updates the operation
function createResponseType(codeModel: CodeModel, group: OperationGroup, op: Operation) {
  // create the `type <type>Response struct` response
  // type with a `RawResponse *http.Response` field
  const firstResp = op.responses![0];

  // when receiving multiple possible responses, they might expect the same headers in many cases
  // we use a map to only add unique headers to the response model based on the header name
  const headers = new Map<string, HttpHeaderWithDescription>();
  for (const resp of values(op.responses)) {
    // check if the response is expecting information from headers
    if (resp.protocol.http!.headers) {
      for (const header of values(resp.protocol.http!.headers)) {
        const head = <HttpHeader>header;
        // convert each header to a property and append it to the response properties list
        const name = pascalCase(head.header);
        if (!headers.has(name)) {
          const description = `${name} contains the information returned from the ${head.header} header response.`
          headers.set(name, <HttpHeaderWithDescription>{
            ...head,
            description: description
          });
        }
      }
    }
  }

  // if the response defines a schema then add it as a field to the response type.
  // only do this if the response schema hasn't been processed yet.

  if (!isSchemaResponse(firstResp)) {
    // the response doesn't return a model.  if it returns
    // headers then create a model that contains them.
    if (headers.size > 0) {
      const name = `${group.language.go!.name}${op.language.go!.name}Response`;
      const description = `${name} contains the response from method ${group.language.go!.name}.${op.language.go!.name}.`;
      const object = new ObjectSchema(name, description);
      object.language.go = object.language.default;
      object.language.go!.properties = [
        newProperty('RawResponse', 'RawResponse contains the underlying HTTP response.', newObject('http.Response', 'raw HTTP response'))
      ];
      for (const item of items(headers)) {
        const prop = newProperty(item.key, item.value.description, item.value.schema);
        prop.language.go!.fromHeader = item.value.header;
        (<Array<Property>>object.language.go!.properties).push(prop);
      }
      // mark as a response type
      object.language.go!.responseType = {
        name: name,
        description: description,
        responseType: true,
      }
      // add this response schema to the global list of response
      const responseSchemas = <Array<Schema>>codeModel.language.go!.responseSchemas;
      responseSchemas.push(object);
      // attach it to the response
      (<SchemaResponse>firstResp).schema = object;
    }
  } else if (!responseTypeCreated(codeModel, firstResp.schema)) {
    firstResp.schema.language.go!.responseType = generateResponseTypeName(firstResp.schema);
    firstResp.schema.language.go!.properties = [
      newProperty('RawResponse', 'RawResponse contains the underlying HTTP response.', newObject('http.Response', 'TODO'))
    ];
    const marshallingFormat = getMarshallingFormat(firstResp.protocol);
    firstResp.schema.language.go!.responseType.marshallingFormat = marshallingFormat;
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
    firstResp.schema.language.go!.responseType.value = propName;
    (<Array<Property>>firstResp.schema.language.go!.properties).push(newProperty(propName, firstResp.schema.language.go!.description, firstResp.schema));
    // add any headers to the response type
    for (const item of items(headers)) {
      const prop = newProperty(item.key, item.value.description, item.value.schema);
      prop.language.go!.fromHeader = item.value.header;
      (<Array<Property>>firstResp.schema.language.go!.properties).push(prop);
    }
    // add this response schema to the global list of response
    const responseSchemas = <Array<Schema>>codeModel.language.go!.responseSchemas;
    responseSchemas.push(firstResp.schema);
  }
  // create pageable type info
  if (isPageableOperation(op)) {
    if (codeModel.language.go!.pageableTypes === undefined) {
      codeModel.language.go!.pageableTypes = new Array<PagerInfo>();
    }
    const name = `${(<SchemaResponse>firstResp).schema.language.go!.name}Pager`;
    // check to see if the pager has already been created
    const pagers = <Array<PagerInfo>>codeModel.language.go!.pageableTypes;
    for (const pager of values(pagers)) {
      if (pager.name === name) {
        // found a match, hook it up to the method
        op.language.go!.pageableType = pager;
        return;
      }
    }
    // create a new one, add to global list and assign to method
    const pager = {
      name: name,
      schema: (<SchemaResponse>firstResp).schema,
      client: camelCase(group.language.go!.clientName),
      nextLink: op.language.go!.paging.nextLinkName,
    };
    pagers.push(pager);
    op.language.go!.pageableType = pager;
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

function responseTypeCreated(codeModel: CodeModel, schema: Schema): boolean {
  const responseType = generateResponseTypeName(schema);
  const responseSchemas = <Array<Schema>>codeModel.language.go!.responseSchemas;
  for (const responseSchema of values(responseSchemas)) {
    if (responseSchema.language.go!.responseType.name === responseType.name) {
      // unnamed string enum responses and string responses are different schemas
      // but with identical layouts.  so we have a corner-case where we've already
      // created a response type (i.e. StringResponse) for one of the schemas but
      // not for the other.  so if the response type has already been created and
      // the responseType hasn't been set, copy it over.
      if (schema.language.go!.responseType === undefined) {
        schema.language.go!.responseType = responseSchema.language.go!.responseType;
      }
      return true;
    }
  }
  return false;
}

function generateResponseTypeName(schema: Schema): Language {
  let name = '';
  switch (schema.type) {
    case SchemaType.Any:
      name = 'InterfaceResponse';
      break;
    case SchemaType.Array:
      const arraySchema = <ArraySchema>schema;
      const arrayElem = <Schema>arraySchema.elementType;
      name = `${pascalCase(arrayElem.language.go!.name)}ArrayResponse`;
      break;
    case SchemaType.Boolean:
      name = 'BoolResponse';
      break;
    case SchemaType.ByteArray:
      name = 'ByteArrayResponse';
      break;
    case SchemaType.Choice:
      const choiceSchema = <ChoiceSchema>schema;
      name = `${choiceSchema.language.go!.name}Response`;
      break;
    case SchemaType.SealedChoice:
      const sealedChoiceSchema = <SealedChoiceSchema>schema;
      name = `${sealedChoiceSchema.language.go!.name}Response`;
      break;
    case SchemaType.Date:
    case SchemaType.DateTime:
    case SchemaType.UnixTime:
      name = 'TimeResponse';
      break;
    case SchemaType.Dictionary:
      const dictSchema = <DictionarySchema>schema;
      const dictElem = <Schema>dictSchema.elementType;
      name = `MapOf${pascalCase(dictElem.language.go!.name)}Response`;
      break;
    case SchemaType.Duration:
      name = 'DurationResponse';
      break;
    case SchemaType.Integer:
      if ((<NumberSchema>schema).precision === 32) {
        name = 'Int32Response';
        break;
      }
      name = 'Int64Response';
      break;
    case SchemaType.Number:
      if ((<NumberSchema>schema).precision === 32) {
        name = 'Float32Response';
        break;
      }
      name = 'Float64Response';
      break;
    case SchemaType.Object:
      name = `${schema.language.go!.name}Response`;
      break;
    case SchemaType.String:
    case SchemaType.Uuid:
      name = 'StringResponse';
      break;
    default:
      throw console.error(`unhandled response schema type ${schema.type}`);
  }
  return {
    name: name,
    description: `${name} is the response envelope for operations that return a ${schema.language.go!.name} type.`,
    responseType: true,
  }
}
