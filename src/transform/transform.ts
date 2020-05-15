/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { camelCase, KnownMediaType, pascalCase, serialize } from '@azure-tools/codegen';
import { Host, startSession, Session } from '@azure-tools/autorest-extension-base';
import { ObjectSchema, ArraySchema, codeModelSchema, ChoiceValue, CodeModel, DateTimeSchema, GroupProperty, HttpHeader, HttpResponse, ImplementationLocation, Language, OperationGroup, SchemaType, NumberSchema, Operation, SchemaResponse, Parameter, Property, Protocols, Schema, DictionarySchema, Protocol, ChoiceSchema, SealedChoiceSchema, ConstantSchema } from '@azure-tools/codemodel';
import { items, values } from '@azure-tools/linq';
import { aggregateParameters, isPageableOperation, isObjectSchema, isSchemaResponse, PagerInfo, isLROOperation, PollerInfo } from '../common/helpers';
import { createPolymorphicInterfaceName, namer, removePrefix } from './namer';

// The transformer adds Go-specific information to the code model.
export async function transform(host: Host) {
  const debug = await host.GetValue('debug') || false;

  try {
    const session = await startSession<CodeModel>(host, {}, codeModelSchema);

    // run the namer first, so that any transformations are applied on proper names
    await namer(session);
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
    if (obj.discriminator) {
      const discriminator = annotateDiscriminatedTypes(obj);
      if (discriminator) {
        // discriminators will contain the root type of each discriminated type hierarchy
        if (!session.model.language.go!.discriminators) {
          session.model.language.go!.discriminators = new Array<ObjectSchema>();
        }
        const defs = <Array<ObjectSchema>>session.model.language.go!.discriminators;
        defs.push(discriminator);
      }
    }
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
    case SchemaType.UnixTime:
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
  // track any parameter groups and/or optional parameters
  const paramGroups = new Map<string, GroupProperty>();
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
        // skip the host param as it's a field on the client
        if (isHostParameter(param)) {
          continue;
        }
        // this is to work around M4 bug #202
        // replace the duplicate operation entry in nextLinkOperation with
        // the one from our operation group so that things like parameter
        // groups/types etc are consistent.
        if (op.language.go!.paging && op.language.go!.paging.nextLinkOperation) {
          const dupeOp = <Operation>op.language.go!.paging.nextLinkOperation;
          for (const internalOp of values(group.operations)) {
            if (internalOp.language.default.name === dupeOp.language.default.name) {
              op.language.go!.paging.nextLinkOperation = internalOp;
              break;
            }
          }
        }
        const inBody = param.protocol.http !== undefined && param.protocol.http!.in === 'body';
        param.schema.language.go!.name = schemaTypeToGoType(session.model, param.schema, inBody);
        if (param.implementation === ImplementationLocation.Client && param.schema.type !== SchemaType.Constant) {
          // add global param info to the operation group
          if (group.language.go!.clientParams === undefined) {
            group.language.go!.clientParams = new Array<Parameter>();
          }
          const clientParams = <Array<Parameter>>group.language.go!.clientParams;
          // check if this global param has already been added
          if (clientParams.includes(param)) {
            continue;
          }
          clientParams.push(param);
        }
        // check for grouping
        if (param.extensions?.['x-ms-parameter-grouping']) {
          // this param belongs to a param group, init name with default
          let paramGroupName = `${group.language.go!.name}${op.language.go!.name}Parameters`;
          if (param.extensions['x-ms-parameter-grouping'].name) {
            // use the specified name
            paramGroupName = pascalCase(param.extensions['x-ms-parameter-grouping'].name);
          } else if (param.extensions['x-ms-parameter-grouping'].postfix) {
            // use the suffix
            paramGroupName = `${group.language.go!.name}${op.language.go!.name}${pascalCase(param.extensions['x-ms-parameter-grouping'].postfix)}`;
          }
          // create group entry and add the param to it
          if (!paramGroups.has(paramGroupName)) {
            const desc = `${paramGroupName} contains a group of parameters for the ${group.language.go!.name}.${op.language.go!.name} method.`;
            paramGroups.set(paramGroupName, createGroupProperty(paramGroupName, desc));
          }
          // associate the group with the param
          const paramGroup = paramGroups.get(paramGroupName);
          param.language.go!.paramGroup = paramGroup;
          // check for a duplicate, if it has the same schema then skip it
          const dupe = values(paramGroup!.originalParameter).first((each: Parameter) => { return each.language.go!.name === param.language.go!.name; });
          if (!dupe) {
            paramGroup!.originalParameter.push(param);
            if (param.required) {
              // mark the group as required if at least one param in the group is required
              paramGroup!.required = true;
            }
          } else if (dupe.schema !== param.schema) {
            throw console.error(`parameter group ${paramGroupName} contains overlapping parameters with different schemas`);
          }
          continue;
        }
        // this is a bit of a weird case and might be due to invalid swagger in the test
        // server.  how can you have an optional parameter that's also a constant?
        // TODO once non-required constants are fixed
        if (param.required !== true && param.schema.type !== SchemaType.Constant && !(param.schema.type === SchemaType.SealedChoice && (<SealedChoiceSchema>param.schema).choices.length === 1)) {
          // create a type named <OperationGroup><Operation>Options
          const paramGroupName = `${group.language.go!.name}${op.language.go!.name}Options`;
          // create group entry and add the param to it
          if (!paramGroups.has(paramGroupName)) {
            const desc = `${paramGroupName} contains the optional parameters for the ${group.language.go!.name}.${op.language.go!.name} method.`;
            paramGroups.set(paramGroupName, createGroupProperty(paramGroupName, desc));
          }
          // associate the group with the param
          param.language.go!.paramGroup = paramGroups.get(paramGroupName);
          paramGroups.get(paramGroupName)!.originalParameter.push(param);
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
  // emit any param groups
  if (paramGroups.size > 0) {
    if (!session.model.language.go!.parameterGroups) {
      session.model.language.go!.parameterGroups = new Array<GroupProperty>();
    }
    const pg = <Array<GroupProperty>>session.model.language.go!.parameterGroups;
    for (const items of paramGroups.entries()) {
      pg.push(items[1]);
    }
  }
}

function isHostParameter(param: Parameter): boolean {
  if (param.language.go!.name === 'host' || param.language.go!.name === '$host') {
    return true;
  }
  return param.extensions?.['x-ms-priority'] === 0 && param.extensions?.['x-in'] === 'path';
}

function createGroupProperty(name: string, description: string): GroupProperty {
  const schema = new ObjectSchema(name, description);
  schema.language.go = schema.language.default;
  const gp = new GroupProperty(name, description, schema);
  gp.language.go = gp.language.default;
  return gp;
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
        if (isObjectSchema(schemaError)) {
          for (const prop of values(schemaError.properties)) {
            // adding the Inner prefix on error types, since errors in Go have an Error() method
            // in order to implement the error interface. This causes errors to not be able to have
            // an Error field as well, since it would cause confusion
            if (prop.language.go!.name === 'Error') {
              prop.language.go!.name = 'Inner' + prop.language.go!.name;
            }
          }
        }
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
  if (!op.responses) {
    return;
  }
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
        const name = removePrefix(pascalCase(head.header), 'XMS');
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
    if (isLROOperation(op)) {
      const name = 'HttpResponse';
      const description = `${name} contains the HTTP response from the call to the service endpoint`;
      const object = new ObjectSchema(name, description);
      object.language.go = object.language.default;
      const pollUntilDone = newProperty('PollUntilDone', 'PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received', newObject('func(ctx context.Context, frequency time.Duration) (*HttpResponse, error)', 'TODO'));
      const getPoller = newProperty('GetPoller', 'GetPoller will return an initialized poller', newObject('func() HttpPoller', 'TODO'));
      pollUntilDone.schema.language.go!.funcType = true;
      getPoller.schema.language.go!.funcType = true;
      object.language.go!.properties = [
        newProperty('RawResponse', 'RawResponse contains the underlying HTTP response.', newObject('http.Response', 'raw HTTP response')),
        pollUntilDone,
        getPoller
      ];
      // mark as a response type
      object.language.go!.responseType = {
        name: name,
        description: description,
        responseType: true,
      };
      if (!responseExists(codeModel, object.language.go!.responseType.name)) {
        // add this response schema to the global list of response
        const responseSchemas = <Array<Schema>>codeModel.language.go!.responseSchemas;
        responseSchemas.push(object);
        // attach it to the response
        (<SchemaResponse>firstResp).schema = object;
      }
    } else if (headers.size > 0) {
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
      if (!responseExists(codeModel, object.language.go!.responseType.name)) {
        // add this response schema to the global list of response
        const responseSchemas = <Array<Schema>>codeModel.language.go!.responseSchemas;
        responseSchemas.push(object);
        // attach it to the response
        (<SchemaResponse>firstResp).schema = object;
      }
    }
  } else if (!responseTypeCreated(codeModel, firstResp.schema)) {
    firstResp.schema.language.go!.responseType = generateResponseTypeName(firstResp.schema);
    if (isLROOperation(op)) {
      firstResp.schema.language.go!.responseType.name = `${firstResp.schema.language.go!.responseType.name}`;
    }
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
      propName = recursiveTypeName(firstResp.schema);
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
    if (isLROOperation(op)) {
      firstResp.schema.language.go!.needsTimeAndContext = true;
      let prop = newProperty('PollUntilDone', 'PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received', newObject(`func(ctx context.Context, frequency time.Duration) (*${firstResp.schema.language.go!.responseType.name}, error)`, 'TODO'));
      prop.schema.language.go!.funcType = true;
      (<Array<Property>>firstResp.schema.language.go!.properties).push(prop);
      prop = newProperty('GetPoller', 'GetPoller will return an initialized poller', newObject(`func() ${firstResp.schema.language.go!.responseType.value}Poller`, 'TODO'));
      prop.schema.language.go!.funcType = true;
      (<Array<Property>>firstResp.schema.language.go!.properties).push(prop);
    }
    if (!responseExists(codeModel, firstResp.schema.language.go!.name)) {
      // add this response schema to the global list of response
      const responseSchemas = <Array<Schema>>codeModel.language.go!.responseSchemas;
      responseSchemas.push(firstResp.schema);
    }
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
      op: op,
    };
    pagers.push(pager);
    op.language.go!.pageableType = pager;
  }
  // create poller type info
  if (isLROOperation(op)) {
    if (codeModel.language.go!.pollerTypes === undefined) {
      codeModel.language.go!.pollerTypes = new Array<PollerInfo>();
    }
    let type = 'Http';
    if (isSchemaResponse(firstResp)) {
      type = firstResp.schema.language.go!.responseType.value;
    }
    if (type == undefined) {
      type = 'Http';
    }
    const name = `${camelCase(type)}Poller`;
    const pollers = <Array<PollerInfo>>codeModel.language.go!.pollerTypes;
    for (const poller of values(pollers)) {
      if (poller.name === name) {
        // found a match, hook it up to the method
        const tempPoller = {
          name: poller.name,
          responseType: poller.responseType,
          declareResume: false,
          op: op,
        };
        op.language.go!.pollerType = tempPoller;
        return;
      }
    }
    // Adding the operation group name to the poller name for polling operations that need to be unique to that operation group
    // const name = `${camelCase(group.language.go!.name)}${op.language.go!.name}Poller`;
    // create a new one, add to global list and assign to method
    const poller = {
      name: name,
      responseType: type,
      declareResume: true,
      op: op,
    };
    pollers.push(poller);
    op.language.go!.pollerType = poller;
  }
}

function responseExists(codeModel: CodeModel, name: string): boolean {
  for (const resp of codeModel.language.go!.responseSchemas) {
    if (resp.language.go!.name === name) {
      return true;
    }
  }
  return false;
}

function newObject(name: string, desc: string): ObjectSchema {
  let obj = new ObjectSchema(name, desc);
  obj.language.go = obj.language.default;
  return obj;
}

function newProperty(name: string, desc: string, schema: Schema): Property {
  let prop = new Property(name, desc, schema);
  if (isObjectSchema(schema) && schema.discriminator) {
    prop.isDiscriminator = true;
  }
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

function recursiveTypeName(schema: Schema): string {
  switch (schema.type) {
    case SchemaType.Any:
      return 'Interface';
    case SchemaType.Array:
      const arraySchema = <ArraySchema>schema;
      const arrayElem = <Schema>arraySchema.elementType;
      return `${recursiveTypeName(arrayElem)}Array`;
    case SchemaType.Boolean:
      return 'Bool';
    case SchemaType.ByteArray:
      return 'ByteArray';
    case SchemaType.Choice:
      const choiceSchema = <ChoiceSchema>schema;
      return choiceSchema.language.go!.name;
    case SchemaType.SealedChoice:
      const sealedChoiceSchema = <SealedChoiceSchema>schema;
      return sealedChoiceSchema.language.go!.name;
    case SchemaType.Date:
    case SchemaType.DateTime:
    case SchemaType.UnixTime:
      return 'Time';
    case SchemaType.Dictionary:
      const dictSchema = <DictionarySchema>schema;
      const dictElem = <Schema>dictSchema.elementType;
      return `MapOf${recursiveTypeName(dictElem)}`;
    case SchemaType.Duration:
      return 'Duration';
    case SchemaType.Integer:
      if ((<NumberSchema>schema).precision === 32) {
        return 'Int32';
      }
      return 'Int64';
    case SchemaType.Number:
      if ((<NumberSchema>schema).precision === 32) {
        return 'Float32';
      }
      return 'Float64';
    case SchemaType.Object:
      return schema.language.go!.name;
    case SchemaType.String:
    case SchemaType.Uuid:
      return 'String';
    default:
      throw console.error(`unhandled response schema type ${schema.type}`);
  }
}

function generateResponseTypeName(schema: Schema): Language {
  const name = `${recursiveTypeName(schema)}Response`;
  return {
    name: name,
    description: `${name} is the response envelope for operations that return a ${schema.language.go!.name} type.`,
    responseType: true,
  }
}

function annotateDiscriminatedTypes(obj: ObjectSchema): ObjectSchema | undefined {
  if (obj.language.go!.polymorphicInterfaces !== undefined) {
    // this hierarchy of discriminated types has already been processed
    return;
  }
  // we have a type in the hierarchy of polymorphic types, it can be one of three things
  // 1. root - no parent types, only child types
  // 2. intermediate root - has a parent and also has children (salmon in the test server)
  // 3. child - has parent and no children
  //
  // for cases #1 and #2 we need to generate an interface type, and for
  // case #2 the generated interface must also contain the parent interface
  // for case #3 all that's required is to generate the marker method on
  // the child type(s) for the applicable interface.

  // walk to the root
  let root = obj;
  while (true) {
    if (!root.parents) {
      // simple case, no parent types
      break;
    }
    for (const parent of values(root.parents?.immediate)) {
      // there can be parents that aren't part of the hierarchy.
      // e.g. if type Foo is in a DictionaryOfFoo, then one of
      // Foo's parents will be DictionaryOfFoo which we ignore.
      if (isObjectSchema(parent) && parent.discriminator) {
        root = parent;
      }
    }
    if (root === obj) {
      // we hit this case if the parent isn't a discriminator.
      // consider the case of BaseDiscriminatedType that includes
      // a parent BaseProperties (in xms-error-responses.json),
      // or the DictionaryOfFoo case above
      break;
    }
  }
  // create the interface type name based on the current root
  const rootType = root.language.go!.discriminator;
  // use pre-defined enum values if available
  const choices = getChoices(root);
  if (!choices) {
    // mark that we need to generate our own enum type
    root.language.go!.discriminatorEnumNeeded = true;
  }
  recursiveAnnotateDiscriminatedTypes(root, rootType, rootType, choices);
  return root;
}

function recursiveAnnotateDiscriminatedTypes(obj: ObjectSchema, rootInterface: string, currentInterface: string, choices: Array<ChoiceValue> | undefined) {
  if (!obj.language.go!.polymorphicInterfaces) {
    obj.language.go!.polymorphicInterfaces = new Array<string>();
  }
  const interfaces = <Array<string>>obj.language.go!.polymorphicInterfaces;
  interfaces.push(currentInterface);
  // now walk all the children, annotating them with the interface
  for (const child of values(obj.discriminator?.immediate)) {
    const childSchema = <ObjectSchema>child;
    if (!childSchema.language.go!.polymorphicInterfaces) {
      // copy parent's interfaces
      childSchema.language.go!.polymorphicInterfaces = [...<Array<string>>obj.language.go!.polymorphicInterfaces];
    }
    if (childSchema.discriminator && childSchema.discriminator.all) {
      // case #2 - intermediate root
      childSchema.language.go!.discriminatorParent = currentInterface;
      recursiveAnnotateDiscriminatedTypes(childSchema, rootInterface, createPolymorphicInterfaceName(childSchema.language.go!.name), choices);
    }
    if (choices) {
      // find the choice value that matches the current type's discriminator
      let found = false;
      for (const choice of values(choices)) {
        if (choice.value === childSchema.discriminatorValue) {
          childSchema.language.go!.discriminatorEnum = choice.language.go!.name;
          childSchema.language.go!.discriminatorRealEnum = true;
          found = true;
          break;
        }
      }
      if (!found) {
        throw console.error(`failed to find discriminator choice value for type ${childSchema.language.go!.name}`);
      }
    } else {
      // add the internal enum name for this sub-type
      childSchema.language.go!.discriminatorEnum = `${camelCase(rootInterface)}${pascalCase(childSchema.discriminatorValue!)}`;
    }
  }
}

function getChoices(obj: ObjectSchema): Array<ChoiceValue> | undefined {
  if (obj.discriminator?.property.schema.type === SchemaType.Choice) {
    return (<ChoiceSchema>obj.discriminator!.property.schema).choices;
  } else if (obj.discriminator?.property.schema.type === SchemaType.SealedChoice) {
    return (<SealedChoiceSchema>obj.discriminator!.property.schema).choices;
  }
  return undefined;
}
