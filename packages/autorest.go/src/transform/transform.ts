/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

/* eslint-disable @typescript-eslint/no-unsafe-argument */
/* eslint-disable @typescript-eslint/no-unsafe-assignment */
/* eslint-disable @typescript-eslint/no-unsafe-member-access */

import { capitalize, KnownMediaType, serialize, uncapitalize } from '@azure-tools/codegen';
import { AutorestExtensionHost, startSession, Session } from '@autorest/extension-base';
import * as m4 from '@autorest/codemodel';
import { clone, items, values } from '@azure-tools/linq';
import { createOptionsTypeDescription, createResponseEnvelopeDescription } from '../../../naming.go/src/naming.js';
import * as helpers from './helpers.js';
import { namer, protocolMethods } from './namer.js';
import { fromString } from 'html-to-text';
import showdown from 'showdown';
import { fileURLToPath } from 'url';
const { Converter } = showdown;

// The transformer adds Go-specific information to the code model.
export async function transformM4(host: AutorestExtensionHost) {
  const debug = await host.getValue('debug') || false;

  try {
    const session = await startSession<m4.CodeModel>(host, m4.codeModelSchema);

    // run the namer first, so that any transformations are applied on proper names
    await namer(session);
    await process(session);
    await labelUnreferencedTypes(session);

    // output the model to the pipeline
    host.writeFile({
      filename: 'code-model-v4-transform.yaml',
      content: serialize(session.model),
      artifactType: 'code-model-v4'
    });

  } catch (E) {
    if (debug) {
      console.error(`${fileURLToPath(import.meta.url)} - FAILURE  ${JSON.stringify(E)} ${(<Error>E).stack}`);
    }
    throw E;
  }
}

async function process(session: Session<m4.CodeModel>) {
  await processOperationRequests(session);
  processOperationResponses(session);
  // fix up dictionary element types (additional properties)
  // this must happen before processing objects as we depend on the
  // schema type being an actual Go type.
  for (const dictionary of values(session.model.schemas.dictionaries)) {
    dictionary.language.go!.name = schemaTypeToGoType(session.model, dictionary, 'Property');
    dictionary.language.go!.elementIsPtr = !helpers.isTypePassedByValue(dictionary.elementType);
    if (dictionary.language.go!.description) {
      dictionary.language.go!.description = parseComments(dictionary.language.go!.description);
    }
  }
  // fix up struct field types
  for (const obj of values(session.model.schemas.objects)) {
    if (obj.language.go!.description) {
      obj.language.go!.description = parseComments(obj.language.go!.description);
      if (!obj.language.go!.description.startsWith(obj.language.go!.name)) {
        obj.language.go!.description = `${obj.language.go!.name} - ${obj.language.go!.description}`;
      }
    }
    if (obj.discriminator) {
      // discriminators will contain the root type of each discriminated type hierarchy
      if (!session.model.language.go!.discriminators) {
        session.model.language.go!.discriminators = new Array<m4.ObjectSchema>();
      }
      const defs = <Array<m4.ObjectSchema>>session.model.language.go!.discriminators;
      const rootDiscriminator = getRootDiscriminator(obj);
      if (defs.indexOf(rootDiscriminator) < 0) {
        rootDiscriminator.language.go!.rootDiscriminator = true;
        defs.push(rootDiscriminator);
        // fix up discriminator value to use the enum type if available
        const discriminatorEnums = getDiscriminatorEnums(rootDiscriminator);
        // for each child type in the hierarchy, fix up the discriminator value
        for (const child of values(rootDiscriminator.children?.all)) {
          const asObj = <m4.ObjectSchema>child;
          let discValue = getEnumForDiscriminatorValue(asObj.discriminatorValue!, discriminatorEnums);
          if (!discValue) {
            discValue = quoteString(asObj.discriminatorValue!);
          }
          asObj.discriminatorValue = discValue;
        }
      }
    }

    for (const prop of values(obj.properties)) {
      if (prop.language.go!.description) {
        prop.language.go!.description = parseComments(prop.language.go!.description);
      }
      const details = <m4.Language>prop.schema.language.go;
      details.name = `${schemaTypeToGoType(session.model, prop.schema, 'InBody')}`;
      prop.schema = substitueDiscriminator(prop);
      if (prop.schema.type === m4.SchemaType.Any || prop.schema.type === m4.SchemaType.AnyObject || (helpers.isObjectSchema(prop.schema) && prop.schema.discriminator)) {
        prop.language.go!.byValue = true;
      } else if (prop.schema.type === m4.SchemaType.DateTime) {
        obj.language.go!.needsDateTimeMarshalling = true;
      } else if (prop.schema.type === m4.SchemaType.Date) {
        obj.language.go!.needsDateMarshalling = true;
      } else if (prop.schema.type === m4.SchemaType.UnixTime) {
        obj.language.go!.needsUnixTimeMarshalling = true;
      } else if (prop.schema.type === m4.SchemaType.Dictionary && obj.language.go!.marshallingFormat === 'xml') {
        // mark that we need custom XML unmarshalling for a dictionary
        prop.language.go!.needsXMLDictionaryUnmarshalling = true;
        session.model.language.go!.needsXMLDictionaryUnmarshalling = true;
      } else if (prop.schema.type === m4.SchemaType.ByteArray) {
        prop.language.go!.byValue = true;
        obj.language.go!.byteArrayFormat = (<m4.ByteArraySchema>prop.schema).format;
      }
      if (prop.schema.type === m4.SchemaType.Array || prop.schema.type === m4.SchemaType.Dictionary) {
        obj.language.go!.hasArrayMap = true;
        prop.language.go!.byValue = true;
        if (prop.schema.type !== m4.SchemaType.Dictionary && obj.language.go!.marshallingFormat === 'xml') {
          prop.language.go!.needsXMLArrayMarshalling = true;
        }
      }
    }
    if (!obj.language.go!.marshallingFormat) {
      // TODO: workaround due to https://github.com/Azure/autorest.go/issues/412
      // this type isn't used as a parameter/return value so it has no marshalling format.
      // AutoRest doesn't make the global configuration available at present so hard-code
      // the format to JSON as the vast majority of specs use JSON.
      obj.language.go!.marshallingFormat = 'json';
    }
    const addPropsSchema = helpers.hasAdditionalProperties(obj);
    if (addPropsSchema) {
      // add an 'AdditionalProperties' field to the type
      const addProps = newProperty('AdditionalProperties', 'OPTIONAL; Contains additional key/value pairs not defined in the schema.', addPropsSchema);
      addProps.language.go!.isAdditionalProperties = true;
      addProps.language.go!.byValue = true;
      obj.properties?.push(addProps);
    }
  }
  // fix up enum types
  for (const choice of values(session.model.schemas.choices)) {
    choice.choiceType.language.go!.name = schemaTypeToGoType(session.model, choice.choiceType, 'Property');
    if (choice.language.go!.description) {
      choice.language.go!.description = parseComments(choice.language.go!.description);
    }
  }
  for (const choice of values(session.model.schemas.sealedChoices)) {
    choice.choiceType.language.go!.name = schemaTypeToGoType(session.model, choice.choiceType, 'Property');
    if (choice.language.go!.description) {
      choice.language.go!.description = parseComments(choice.language.go!.description);
    }
  }
}

// used by getDiscriminatorSchema and substitueDiscriminator
const discriminatorSchemas = new Map<string, m4.Schema>();

// creates/gets an ObjectSchema for a discriminated type.
// NOTE: assumes schema is a discriminator.
function getDiscriminatorSchema(schema: m4.ObjectSchema): m4.Schema {
  const discriminatorInterface = schema.language.go!.discriminatorInterface;
  if (!discriminatorSchemas.has(discriminatorInterface)) {
    const discriminatorSchema = new m4.ObjectSchema(discriminatorInterface, 'discriminated type');
    discriminatorSchema.language.go = discriminatorSchema.language.default;
    discriminatorSchema.language.go.discriminatorInterface = discriminatorInterface;
    // copy over fields from the original
    discriminatorSchema.discriminator = schema.discriminator;
    discriminatorSchema.children = schema.children;
    discriminatorSchema.parents = schema.parents;
    discriminatorSchemas.set(discriminatorInterface, discriminatorSchema);
  }
  return discriminatorSchemas.get(discriminatorInterface)!;
}

// replaces item's schema with the appropriate discriminator schema as required
function substitueDiscriminator(item: m4.Property | m4.Parameter | m4.SchemaResponse): m4.Schema {
  if (item.schema.type === m4.SchemaType.Object && item.schema.language.go!.discriminatorInterface) {
    item.language.go!.byValue = true;
    return getDiscriminatorSchema(<m4.ObjectSchema>item.schema);
  } else if (helpers.isArraySchema(item.schema) || helpers.isDictionarySchema(item.schema)) {
    const leafElementSchema = helpers.recursiveUnwrapArrayDictionary(item.schema);
    if (leafElementSchema.type === m4.SchemaType.Object && leafElementSchema.language.go!.discriminatorInterface) {
      return recursiveSubstitueDiscriminator(item.schema);
    }
  }

  // not a discriminator
  return item.schema;
}

// constructs a new schema for arrays/maps with a leaf element that's a discriminator
// NOTE: assumes that the leaf element type is a discriminated type
// e.g. []map[string]DiscriminatorInterface
function recursiveSubstitueDiscriminator(item: m4.Schema): m4.Schema {
  const strings = recursiveBuildDiscriminatorStrings(item);
  let discriminatorSchema = discriminatorSchemas.get(strings.Name);
  if (discriminatorSchema) {
    return discriminatorSchema;
  }
  if (helpers.isArraySchema(item)) {
    discriminatorSchema = new m4.ArraySchema(strings.Name, strings.Desc, recursiveSubstitueDiscriminator(item.elementType));
    discriminatorSchema.language.go = discriminatorSchema.language.default;
    discriminatorSchema.language.go.elementIsPtr = false;
    discriminatorSchemas.set(strings.Name, discriminatorSchema);
    return discriminatorSchema;
  } else if (helpers.isDictionarySchema(item)) {
    discriminatorSchema = new m4.DictionarySchema(strings.Name, strings.Desc, recursiveSubstitueDiscriminator(item.elementType));
    discriminatorSchema.language.go = discriminatorSchema.language.default;
    discriminatorSchema.language.go.elementIsPtr = false;
    discriminatorSchemas.set(strings.Name, discriminatorSchema);
    return discriminatorSchema;
  }
  return getDiscriminatorSchema(<m4.ObjectSchema>item);
}

interface DiscriminatorStrings {
  Name: string;
  Desc: string;
}

// constructs the name and description for arrays/maps with a leaf element that's a discriminator
// NOTE: assumes that the leaf element type is a discriminated type
// Name: []map[string]DiscriminatorInterface
// Desc: slice of map of discriminators
function recursiveBuildDiscriminatorStrings(item: m4.Schema): DiscriminatorStrings {
  if (helpers.isArraySchema(item)) {
    const strings = recursiveBuildDiscriminatorStrings(item.elementType);
    return {
      Name: `[]${strings.Name}`,
      Desc: `array of ${strings.Desc}`
    };
  } else if (helpers.isDictionarySchema(item)) {
    const strings = recursiveBuildDiscriminatorStrings(item.elementType);
    return {
      Name: `map[string]${strings.Name}`,
      Desc: `map of ${strings.Desc}`
    };
  }
  return {
    Name: getDiscriminatorSchema(<m4.ObjectSchema>item).language.go!.discriminatorInterface,
    Desc: 'discriminators'
  };
}

const dictionaryElementAnySchema = new m4.AnySchema('any schema for maps');
dictionaryElementAnySchema.language.go = dictionaryElementAnySchema.language.default;
dictionaryElementAnySchema.language.go.name = 'any';

function schemaTypeToGoType(codeModel: m4.CodeModel, schema: m4.Schema, type: 'Property' | 'InBody' | 'HeaderParam' | 'PathParam' | 'QueryParam'): string {
  const rawJSONAsBytes = <boolean>codeModel.language.go!.rawJSONAsBytes;
  switch (schema.type) {
    case m4.SchemaType.Any:
      if (rawJSONAsBytes) {
        schema.language.go!.rawJSONAsBytes = rawJSONAsBytes;
        return '[]byte';
      }
      return 'any';
    case m4.SchemaType.AnyObject:
      if (rawJSONAsBytes) {
        schema.language.go!.rawJSONAsBytes = rawJSONAsBytes;
        return '[]byte';
      }
      return 'map[string]any';
    case m4.SchemaType.ArmId:
      return 'string';
    case m4.SchemaType.Array: {
      const arraySchema = <m4.ArraySchema>schema;
      const arrayElem = arraySchema.elementType;
      if (rawJSONAsBytes && (arrayElem.type === m4.SchemaType.Any || arrayElem.type === m4.SchemaType.AnyObject)) {
        schema.language.go!.rawJSONAsBytes = rawJSONAsBytes;
        // propagate the setting to the element type
        arrayElem.language.go!.rawJSONAsBytes = rawJSONAsBytes;
        return '[][]byte';
      }
      arraySchema.language.go!.elementIsPtr = !helpers.isTypePassedByValue(arrayElem) && !<boolean>codeModel.language.go!.sliceElementsByValue;
      // passing nil for array elements in headers, paths, and query params
      // isn't very useful as we'd just skip nil entries.  so disable it.
      if (type !== 'Property' && type !== 'InBody') {
        arraySchema.language.go!.elementIsPtr = false;
      }
      arrayElem.language.go!.name = schemaTypeToGoType(codeModel, arrayElem, type);
      if (<boolean>arraySchema.language.go!.elementIsPtr) {
        return `[]*${arrayElem.language.go!.name}`;
      }
      return `[]${arrayElem.language.go!.name}`;
    }
    case m4.SchemaType.Binary:
      return 'io.ReadSeekCloser';
    case m4.SchemaType.Boolean:
      return 'bool';
    case m4.SchemaType.ByteArray:
      return '[]byte';
    case m4.SchemaType.Char:
      return 'rune';
    case m4.SchemaType.Constant: {
      const constSchema = <m4.ConstantSchema>schema;
      constSchema.valueType.language.go!.name = schemaTypeToGoType(codeModel, constSchema.valueType, type);
      return constSchema.valueType.language.go!.name;
    }
    case m4.SchemaType.DateTime: {
      const dateTime = <m4.DateTimeSchema>schema;
      if (dateTime.format === 'date-time-rfc1123') {
        schema.language.go!.internalTimeType = 'RFC1123';
      } else {
        schema.language.go!.internalTimeType = 'RFC3339';
      }
      if (type === 'InBody') {
        // add a marker to the code model indicating that we need
        // to include support for marshalling/unmarshalling time.
        // header/query param values are parsed separately so they
        // don't need the custom time types.
        if (dateTime.format === 'date-time-rfc1123') {
          codeModel.language.go!.generateDateTimeRFC1123Helper = true;
        } else {
          codeModel.language.go!.generateDateTimeRFC3339Helper = true;
        }
      }
      return 'time.Time';
    }
    case m4.SchemaType.UnixTime:
      // unix time always requires the helper time type
      codeModel.language.go!.generateUnixTimeHelper = true;
      schema.language.go!.internalTimeType = 'Unix';
      return 'time.Time';
    case m4.SchemaType.Dictionary: {
      const dictSchema = <m4.DictionarySchema>schema;
      const dictElem = dictSchema.elementType;
      if (rawJSONAsBytes && (dictElem.type === m4.SchemaType.Any || dictElem.type === m4.SchemaType.AnyObject)) {
        dictSchema.elementType = dictionaryElementAnySchema;
        return `map[string]${dictionaryElementAnySchema.language.go!.name}`;
      }
      dictSchema.language.go!.elementIsPtr = !helpers.isTypePassedByValue(dictSchema.elementType);
      dictElem.language.go!.name = schemaTypeToGoType(codeModel, dictElem, type);
      if (<boolean>dictSchema.language.go!.elementIsPtr) {
        return `map[string]*${dictElem.language.go!.name}`;
      }
      return `map[string]${dictElem.language.go!.name}`;
    }
    case m4.SchemaType.Integer:
      if ((<m4.NumberSchema>schema).precision === 32) {
        return 'int32';
      }
      return 'int64';
    case m4.SchemaType.Number:
      if ((<m4.NumberSchema>schema).precision === 32) {
        return 'float32';
      }
      return 'float64';
    case m4.SchemaType.Credential:
    case m4.SchemaType.Duration:
    case m4.SchemaType.ODataQuery:
    case m4.SchemaType.String:
    case m4.SchemaType.Uuid:
    case m4.SchemaType.Uri:
      return 'string';
    case m4.SchemaType.Date:
      schema.language.go!.internalTimeType = 'PlainDate';
      if (type === 'InBody') {
        codeModel.language.go!.generateDateHelper = true;
      }
      return 'time.Time';
    case m4.SchemaType.Time:
      schema.language.go!.internalTimeType = 'PlainTime';
      if (type === 'InBody') {
        codeModel.language.go!.generateTimeRFC3339Helper = true;
      }
      return 'time.Time';
    case m4.SchemaType.Choice:
    case m4.SchemaType.SealedChoice:
    case m4.SchemaType.Object:
      return schema.language.go!.name;
    default:
      throw new Error(`unhandled schema type ${schema.type}`);
  }
}

function recursiveAddMarshallingFormat(schema: m4.Schema, marshallingFormat: 'json' | 'xml') {
  // only recurse if the schema isn't a primitive type
  const shouldRecurse = function (schema: m4.Schema): boolean {
    return schema.type === m4.SchemaType.Array || schema.type === m4.SchemaType.Dictionary || schema.type === m4.SchemaType.Object;
  };
  if (schema.language.go!.marshallingFormat) {
    // this schema has already been processed, don't do it again
    return;
  }
  schema.language.go!.marshallingFormat = marshallingFormat;
  switch (schema.type) {
    case m4.SchemaType.Array: {
      const arraySchema = <m4.ArraySchema>schema;
      if (shouldRecurse(arraySchema.elementType)) {
        recursiveAddMarshallingFormat(arraySchema.elementType, marshallingFormat);
      }
      break;
    }
    case m4.SchemaType.Dictionary: {
      const dictSchema = <m4.DictionarySchema>schema;
      if (shouldRecurse(dictSchema.elementType)) {
        recursiveAddMarshallingFormat(dictSchema.elementType, marshallingFormat);
      }
      break;
    }
    case m4.SchemaType.Object: {
      const os = <m4.ObjectSchema>schema;
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
}

// we will transform operation request parameter schema types to Go types
async function processOperationRequests(session: Session<m4.CodeModel>) {
  // pre-process multi-request operations as it can add operations to the operations
  // collection, and iterating over a modified collection yeilds incorrect results
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations).toArray()) {
      if (op.language.go!.description) {
        op.language.go!.description = parseComments(op.language.go!.description);
      }

      const normalizeOperationName = await session.getValue('normalize-operation-name', false);

      // previous operation naming logic: keep original name if only one body type, and add suffix for operation with non-binary body type if more than one body type
      // new normalized operation naming logic: add suffix for operation with unstructured body type and keep original name for operation with structured body type 
      if (!normalizeOperationName){
        if (op.requests!.length > 1) {
          // for the non-binary media types we create a new method with the
          // media type name as a suffix, e.g. FooAPIWithJSON()
          separateOperationByRequestsProtocol(group, op, [KnownMediaType.Binary]);
        }
      } else {
        // add suffix to binary/text, suppose there will be only one structured media type
        separateOperationByRequestsProtocol(group, op, [KnownMediaType.Json, KnownMediaType.Xml, KnownMediaType.Form, KnownMediaType.Multipart]);
      }

      if (!group.language.go!.host) {
        // all operations/groups have the same host, so just grab the first one
        group.language.go!.host = op.requests?.[0].protocol.http?.uri;
      }
    }
  }
  // track any client-level parameterized host params.
  const hostParams = new Array<m4.Parameter>();
  // track any parameter groups and/or optional parameters
  const paramGroups = new Map<string, m4.GroupProperty>();
  const singleClient = <boolean>session.model.language.go!.singleClient;
  if (singleClient && session.model.operationGroups.length > 1) {
    throw new Error('single-client cannot be enabled when there are multiple clients');
  }
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      if (op.language.go!.description) {
        op.language.go!.description = parseComments(op.language.go!.description);
      }
      if (op.requests![0].protocol.http!.headers) {
        for (const header of values(op.requests![0].protocol.http!.headers)) {
          const head = <m4.HttpHeader>header;
          head.schema.language.go!.name = schemaTypeToGoType(session.model, head.schema, 'Property');
        }
      }
      const opName = helpers.isLROOperation(op) ? 'Begin' + op.language.go!.name : op.language.go!.name;
      const params = values(helpers.aggregateParameters(op));
      // create an optional params struct even if the operation contains no optional params.
      // this provides version resiliency in case optional params are added in the future.
      // don't do this for paging next link operation as this isn't part of the public API
      if (!op.language.go!.paging || !op.language.go!.paging.isNextOp) {
        // create a type named <OperationGroup><Operation>Options
        // if single-client is enabled, omit the <OperationGroup> prefix
        let clientPrefix = capitalize(group.language.go!.clientName);
        if (singleClient) {
          clientPrefix = '';
        }
        const optionalParamsGroupName = `${clientPrefix}${opName}Options`;
        const desc = createOptionsTypeDescription(optionalParamsGroupName, `${group.language.go!.clientName}.${helpers.isPageableOperation(op) && !helpers.isLROOperation(op) ? `New${opName}Pager` : opName}`);
        const gp = createGroupProperty(optionalParamsGroupName, desc, false);
        gp.language.go!.name = 'options';
        // if there's an existing parameter with the name options then pick something else
        for (const param of params) {
          if (param.language.go!.name === gp.language.go!.name) {
            gp.language.go!.name = 'opts';
            break;
          }
        }
        gp.required = false;
        paramGroups.set(optionalParamsGroupName, gp);
        // associate the param group with the operation
        op.language.go!.optionalParamGroup = gp;
      }
      for (const param of params) {
        if (param.language.go!.description) {
          param.language.go!.description = parseComments(param.language.go!.description);
        }
        if (param.clientDefaultValue && param.implementation === m4.ImplementationLocation.Method) {
          // we treat method params with a client-side default as optional
          // since if you don't specify a value, a default is sent and the
          // zero-value is ambiguous.
          // NOTE: the assumption for client parameters with client-side defaults is that
          // the proper default values are being set in the client (hand-written) constructor.
          param.required = false;
        }
        if (!param.required && param.schema.type === m4.SchemaType.Constant && !param.language.go!.amendedDesc) {
          if (param.language.go!.description) {
            param.language.go!.description += '. ';
          }
          param.language.go!.description += `Specifying any value will set the value to ${(<m4.ConstantSchema>param.schema).value.value}.`;
          param.language.go!.amendedDesc = true;
        }
        // this is to work around M4 bug #202
        // replace the duplicate operation entry in nextLinkOperation with
        // the one from our operation group so that things like parameter
        // groups/types etc are consistent.
        if (op.language.go!.paging && op.language.go!.paging.nextLinkOperation) {
          const dupeOp = <m4.Operation>op.language.go!.paging.nextLinkOperation;
          for (const internalOp of values(group.operations)) {
            if (internalOp.language.default.name === dupeOp.language.default.name) {
              op.language.go!.paging.nextLinkOperation = internalOp;
              break;
            }
          }
          for (const internalOp of values(group.operations)) {
            if (internalOp.language.go!.name === dupeOp.language.default.name) {
              if (!internalOp.language.go!.paging.isNextOp) {
                internalOp.language.go!.paging.isNextOp = true;
                internalOp.language.go!.name = `${uncapitalize(internalOp.language.go!.name)}CreateRequest`;
              }
              break;
            }
          }
        }
        let paramType: 'Property' | 'InBody' | 'HeaderParam' | 'PathParam' | 'QueryParam';
        switch (param.protocol.http?.in) {
          case 'body':
            paramType = 'InBody';
            break;
          case 'header':
            paramType = 'HeaderParam';
            break;
          case 'path':
            paramType = 'PathParam';
            break;
          case 'query':
            paramType = 'QueryParam';
            break;
          default:
            paramType = 'Property';
        }
        param.schema.language.go!.name = schemaTypeToGoType(session.model, param.schema, paramType);
        if (helpers.isTypePassedByValue(param.schema)) {
          param.language.go!.byValue = true;
        }
        param.schema = substitueDiscriminator(param);
        // check if this is a header collection
        if (param.extensions?.['x-ms-header-collection-prefix']) {
          param.schema.language.go!.headerCollectionPrefix = param.extensions['x-ms-header-collection-prefix'];
        }
        if (param.implementation === m4.ImplementationLocation.Client && (param.schema.type !== m4.SchemaType.Constant || !param.required)) {
          if (param.language.default.name === '$host' && <boolean>session.model.language.go!.azureARM) {
            // for ARM, the host is handled by azcore/arm.Client so we can skip it
            continue;
          }
          if (param.protocol.http!.in === 'uri') {
            // this is a parameterized host param.
            // use the param name to avoid reference equality checks.
            if (!values(hostParams).where(p => p.language.go!.name === param.language.go!.name).any()) {
              hostParams.push(param);
            }
            // we special-case fully templated host param, e.g. {endpoint}
            // as there's no need to do a find/replace in this case, we'd
            // just directly use the endpoint param value.
            if (param.language.default.name !== '$host' && !(<string>group.language.go!.host).match(/^\{\w+\}$/)) {
              group.language.go!.complexHostParams = true;
            }
            continue;
          }
          // add global param info to the operation group
          if (group.language.go!.clientParams === undefined) {
            group.language.go!.clientParams = new Array<m4.Parameter>();
          }
          const clientParams = <Array<m4.Parameter>>group.language.go!.clientParams;
          // check if this global param has already been added
          if (values(clientParams).where(cp => cp.language.go!.name === param.language.go!.name).any()) {
            continue;
          }
          clientParams.push(param);
        } else if (param.implementation === m4.ImplementationLocation.Method && param.protocol.http!.in === 'uri') {
          // at least one method contains a parameterized host param, bye-bye simple case
          group.language.go!.complexHostParams = true;
        }
        // check for grouping
        if (param.extensions?.['x-ms-parameter-grouping'] && <boolean>session.model.language.go!.groupParameters) {
          // this param belongs to a param group, init name with default
          // if single-client is enabled, omit the <OperationGroup> prefix
          let clientPrefix = capitalize(group.language.go!.clientName);
          if (singleClient) {
            clientPrefix = '';
          }
          let paramGroupName = `${clientPrefix}${opName}Parameters`;
          if (param.extensions['x-ms-parameter-grouping'].name) {
            // use the specified name
            paramGroupName = <string>param.extensions['x-ms-parameter-grouping'].name;
          } else if (param.extensions['x-ms-parameter-grouping'].postfix) {
            // use the suffix
            paramGroupName = `${clientPrefix}${opName}${<string>param.extensions['x-ms-parameter-grouping'].postfix}`;
          }
          // create group entry and add the param to it
          if (!paramGroups.has(paramGroupName)) {
            let subtext = `.${opName} method`;
            let groupedClientParams = false;
            if (param.implementation === m4.ImplementationLocation.Client) {
              subtext = ' client';
              groupedClientParams = true;
            }
            const desc = `${paramGroupName} contains a group of parameters for the ${group.language.go!.clientName}${subtext}.`;
            const paramGroup = createGroupProperty(paramGroupName, desc, groupedClientParams);
            paramGroups.set(paramGroupName, paramGroup);
          }
          // associate the group with the param
          const paramGroup = paramGroups.get(paramGroupName);
          param.language.go!.paramGroup = paramGroup;
          // check for a duplicate, if it has the same schema then skip it
          const dupe = values(paramGroup!.originalParameter).first((each: m4.Parameter) => { return each.language.go!.name === param.language.go!.name; });
          if (!dupe) {
            paramGroup!.originalParameter.push(param);
            if (param.required) {
              // mark the group as required if at least one param in the group is required
              paramGroup!.required = true;
            }
          } else if (dupe.schema !== param.schema) {
            throw new Error(`parameter group ${paramGroupName} contains overlapping parameters with different schemas`);
          }
          // check for groupings of client and method params
          for (const otherParam of values(paramGroup?.originalParameter)) {
            if (otherParam.implementation !== param.implementation) {
              throw new Error(`parameter group ${paramGroupName} contains client and method parameters`);
            }
          }
        } else if (param.implementation === m4.ImplementationLocation.Method && param.required !== true) {
          // include all non-required method params in the optional values struct.
          (<m4.GroupProperty>op.language.go!.optionalParamGroup).originalParameter.push(param);
          // associate the group with the param
          param.language.go!.paramGroup = op.language.go!.optionalParamGroup;
        }
      }
      if (helpers.isLROOperation(op)) {
        // add the ResumeToken to the optional params type
        const tokenParam = newParameter('ResumeToken', '', newString('string', ''));
        tokenParam.language.go!.byValue = true;
        tokenParam.language.go!.isResumeToken = true;
        tokenParam.required = false;
        op.parameters?.push(tokenParam);
        tokenParam.language.go!.paramGroup = op.language.go!.optionalParamGroup;
        (<m4.GroupProperty>op.language.go!.optionalParamGroup).originalParameter.push(tokenParam);
      }
      // recursively add the marshalling format to the body param if applicable
      const marshallingFormat = getMarshallingFormat(op.requests![0].protocol);
      if (marshallingFormat !== 'na') {
        const bodyParam = values(helpers.aggregateParameters(op)).where((each: m4.Parameter) => { return each.protocol.http?.in === 'body'; }).first();
        if (bodyParam) {
          recursiveAddMarshallingFormat(bodyParam.schema, marshallingFormat);
          if (marshallingFormat === 'xml' && bodyParam.schema.serialization?.xml?.name) {
            // mark that this parameter type will need a custom marshaller to handle the XML name
            bodyParam.schema.language.go!.xmlWrapperName = bodyParam.schema.serialization?.xml?.name;
          } else if (marshallingFormat === 'json' && op.requests![0].protocol.http!.method === 'patch') {
            // mark that this type will need a custom marshaller to handle JSON nulls
            bodyParam.schema.language.go!.needsPatchMarshaller = true;
          }
        }
      }
    }
    if (hostParams.length > 0) {
      // attach host params to the operation group
      group.language.go!.hostParams = hostParams;
    } else if (!<boolean>session.model.language.go!.azureARM) {
      // if there are no host params and this isn't Azure ARM, check for a swagger-provided host (e.g. test server)
      for (const param of values(session.model.globalParameters)) {
        if (param.language.default.name === '$host') {
          group.language.go!.host = param.clientDefaultValue;
          session.model.language.go!.host = group.language.go!.host;
          break;
        }
      }
    }
  }
  // emit any param groups
  if (paramGroups.size > 0) {
    if (!session.model.language.go!.parameterGroups) {
      session.model.language.go!.parameterGroups = new Array<m4.GroupProperty>();
    }
    const pg = <Array<m4.GroupProperty>>session.model.language.go!.parameterGroups;
    for (const items of paramGroups.entries()) {
      pg.push(items[1]);
    }
  }  
}

function createGroupProperty(name: string, description: string, groupedClientParams: boolean): m4.GroupProperty {
  const schema = new m4.ObjectSchema(name, description);
  schema.language.go = schema.language.default;
  const gp = new m4.GroupProperty(name, description, schema);
  gp.language.go = gp.language.default;
  if (groupedClientParams) {
    gp.language.go.groupedClientParams = true;
  }
  return gp;
}

function processOperationResponses(session: Session<m4.CodeModel>) {
  if (session.model.language.go!.responseEnvelopes === undefined) {
    session.model.language.go!.responseEnvelopes = new Array<m4.Schema>();
  }
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      // recursively add the marshalling format to the responses if applicable.
      // also remove any HTTP redirects from the list of responses.
      const filtered = new Array<m4.Response>();
      for (const resp of values(op.responses)) {
        if (skipRedirectStatusCode(<string>op.requests![0].protocol.http!.method, resp)) {
          // redirects are transient status codes, they aren't actually returned
          continue;
        }
        if (helpers.isSchemaResponse(resp)) {
          if (resp.schema.type === m4.SchemaType.Binary) {
            // don't create response envelopes for binary responses.
            // callers read directly from the *http.Response.Body
            continue;
          }
          resp.schema.language.go!.name = schemaTypeToGoType(session.model, resp.schema, 'InBody');
        }
        const marshallingFormat = getMarshallingFormat(resp.protocol);
        if (marshallingFormat !== 'na' && helpers.isSchemaResponse(resp)) {
          recursiveAddMarshallingFormat(resp.schema, marshallingFormat);
        }
        // fix up schema types for header responses
        const httpResponse = <m4.HttpResponse>resp.protocol.http;
        for (const header of values(httpResponse.headers)) {
          header.schema.language.go!.name = schemaTypeToGoType(session.model, header.schema, 'Property');
          // check if this is a header collection
          if (header.extensions?.['x-ms-header-collection-prefix']) {
            header.schema.language.go!.headerCollectionPrefix = header.extensions['x-ms-header-collection-prefix'];
          }
        }
        filtered.push(resp);
        if (resp.language.go!.description) {
          resp.language.go!.description = parseComments(resp.language.go!.description);
        }
      }
      // replace with the filtered list if applicable
      if (filtered.length === 0) {
        // handling of operations with no responses expects an undefined list, not an empty one
        op.responses = undefined;
      } else if (op.responses?.length !== filtered.length) {
        op.responses = filtered;
      }
      createResponseEnvelope(session.model, group, op);
    }
  }
}

// returns true if the specified status code is an automatic HTTP redirect.
// certain redirects are automatically handled by the HTTP stack and thus are
// transient so they are never actually returned to the caller.  we skip them
// so they aren't included in the potential result set of an operation.
function skipRedirectStatusCode(verb: string, resp: m4.Response): boolean {
  const statusCodes = <Array<string>>resp.protocol.http!.statusCodes;
  if (statusCodes.length > 1) {
    return false;
  }
  // taken from src/net/http/client.go in the gostdlib
  switch (statusCodes[0]) {
    case '301':
    case '302':
    case '303':
      if (verb === 'get' || verb === 'head') {
        return true;
      }
      break;
    case '307':
    case '308':
      return true;
  }
  return false;
}

interface HttpHeaderWithDescription extends m4.HttpHeader {
  description: string;
}

// the name of the struct field for scalar responses (int, string, etc)
const scalarResponsePropName = 'Value';

// creates the response envelope type to be returned from an operation and updates the operation.
function createResponseEnvelope(codeModel: m4.CodeModel, group: m4.OperationGroup, op: m4.Operation) {
  // create the `type <type>Response struct` response

  // aggregate headers from all responses as all of them will go into the same result envelope
  const headers = new Map<string, HttpHeaderWithDescription>();
  // skip adding headers for LROs for now, this is to avoid adding
  // any LRO-specific polling headers.
  // TODO: maybe switch to skip specific polling ones in the future?
  if (!helpers.isLROOperation(op)) {
    for (const resp of values(op.responses)) {
      // check if the response is expecting information from headers
      for (const header of values(resp.protocol.http!.headers)) {
        const head = <m4.HttpHeader>header;
        // convert each header to a property and append it to the response properties list
        const name = head.language.go!.name;
        if (!headers.has(name)) {
          const description = `${name} contains the information returned from the ${head.header} header response.`;
          headers.set(name, <HttpHeaderWithDescription>{
            ...head,
            description: description
          });
        }
      }
    }
  }

  // contains all the response envelopes
  const responseEnvelopes = <Array<m4.ObjectSchema>>codeModel.language.go!.responseEnvelopes;
  // first create the response envelope, each operation gets one
  // if single-client is enabled, omit the <OperationGroup> prefix
  let clientPrefix = capitalize(group.language.go!.clientName);
  if (codeModel.language.go!.singleClient) {
    clientPrefix = '';
  }
  const respEnvName = ensureUniqueModelName(codeModel, `${clientPrefix}${op.language.go!.name}Response`, 'Envelope');
  const opName = helpers.isLROOperation(op) ? 'Begin' + op.language.go!.name : op.language.go!.name;
  const respEnv = newObject(respEnvName, createResponseEnvelopeDescription(respEnvName, `${group.language.go!.clientName}.${helpers.isPageableOperation(op) && !helpers.isLROOperation(op) ? `New${opName}Pager` : opName}`));
  respEnv.language.go!.responseType = true;
  respEnv.properties = new Array<m4.Property>();
  responseEnvelopes.push(respEnv);
  op.language.go!.responseEnv = respEnv;
  if (helpers.isLROOperation(op)) {
    respEnv.language.go!.forLRO = true;
  }

  // add any headers to the response
  for (const item of items(headers)) {
    const prop = newRespProperty(item.key, item.value.description, item.value.schema, false);
    // propagate any extensions so we can access them through the property
    prop.extensions = item.value.extensions;
    prop.language.go!.fromHeader = item.value.header;
    respEnv.properties.push(prop);
  }

  // now create the result field

  if (codeModel.language.go!.headAsBoolean && op.requests![0].protocol.http!.method === 'head') {
    op.language.go!.headAsBoolean = true;
    const successProp = newProperty('Success', 'Success indicates if the operation succeeded or failed.', newBoolean('bool', 'bool response'));
    successProp.language.go!.byValue = true;
    respEnv.properties.push(successProp);
    respEnv.language.go!.resultProp = successProp;
    return;
  }
  if (helpers.isMultiRespOperation(op)) {
    const resultTypes = new Array<string>();
    for (const response of values(op.responses)) {
      // the operation might contain a mix of schemas and non-schema responses.
      // we only care about the ones that return a schema.
      if (helpers.isSchemaResponse(response)) {
        resultTypes.push(response.schema.language.go!.name);
      }
    }
    const resultProp = newRespProperty('Value', `Possible types are ${resultTypes.join(', ')}\n`, newAny('multi-response value'), true);
    respEnv.properties.push(resultProp);
    respEnv.language.go!.resultProp = resultProp;
    return;
  }
  const response = helpers.getSchemaResponse(op);
  // if the response defines a schema then add it to the response envelope
  if (response) {
    const rawJSONAsBytes = <boolean>codeModel.language.go!.rawJSONAsBytes;
    // propagate marshalling format to the response envelope
    respEnv.language.go!.marshallingFormat = response.schema.language.go!.marshallingFormat;
    // for operations that return scalar types we use a fixed field name
    let propName = scalarResponsePropName;
    if (response.schema.type === m4.SchemaType.Object) {
      // for object types use the type's name as the field name
      propName = response.schema.language.go!.name;
    } else if (response.schema.type === m4.SchemaType.Array) {
      // for array types use the element type's name
      propName = recursiveTypeName(response.schema);
    } else if (rawJSONAsBytes && (response.schema.type === m4.SchemaType.Any || response.schema.type === m4.SchemaType.AnyObject)) {
      propName = 'RawJSON';
    } else if (response.schema.type === m4.SchemaType.Any) {
      propName = 'Interface';
    } else if (response.schema.type === m4.SchemaType.AnyObject) {
      propName = 'Object';
    }
    if (response.schema.serialization?.xml && response.schema.serialization.xml.name) {
      // always prefer the XML name
      propName = capitalize(response.schema.serialization.xml.name);
    }
    // we want to pass integral types byref to maintain parity with struct fields
    const byValue = helpers.isTypePassedByValue(response.schema) || response.schema.type === m4.SchemaType.Object;
    const resultSchema = substitueDiscriminator(response);
    const resultProp = newRespProperty(propName, response.schema.language.go!.description, resultSchema, byValue);
    if (resultSchema.language.go!.discriminatorInterface) {
      // if the schema is a discriminator we need to flag this on the property itself.
      // this is so the correct unmarshaller is created for the response envelope.
      resultProp.isDiscriminator = true;
    }
    respEnv.properties.push(resultProp);
    respEnv.language.go!.resultProp = resultProp;
  } else if (helpers.isBinaryResponseOperation(op)) {
    const binaryProp = newProperty('Body', 'Body contains the streaming response.', newBinary('binary response'));
    binaryProp.language.go!.byValue = true;
    respEnv.properties.push(binaryProp);
    respEnv.language.go!.resultProp = binaryProp;
  }
  if (respEnv.properties.length === 0) {
    // if we get here it means the operation doesn't return anything. we set
    // this to undefined to simplify detection of an empty response envelope
    respEnv.properties = undefined;
  }
}

// appends suffix to name if name is an existing model type
function ensureUniqueModelName(codeModel: m4.CodeModel, name: string, suffix: string): string {
  for (const obj of values(codeModel.schemas.objects)) {
    if (obj.language.go!.name === name) {
      return name + suffix;
    }
  }
  return name;
}

function newObject(name: string, desc: string): m4.ObjectSchema {
  const obj = new m4.ObjectSchema(name, desc);
  obj.language.go = obj.language.default;
  return obj;
}

function newAny(desc: string): m4.AnySchema {
  const any = new m4.AnySchema(desc);
  any.language.go = any.language.default;
  any.language.go.name = 'any';
  return any;
}

function newBoolean(name: string, desc: string): m4.BooleanSchema {
  const bool = new m4.BooleanSchema(name, desc);
  bool.language.go = bool.language.default;
  bool.language.go.name = 'bool';
  return bool;
}

function newBinary(desc: string): m4.BinarySchema {
  const binary = new m4.BinarySchema(desc);
  binary.language.go = binary.language.default;
  binary.language.go.name = 'io.ReadCloser';
  return binary;
}

function newString(name: string, desc: string): m4.StringSchema {
  const string = new m4.StringSchema(name, desc);
  string.language.go = string.language.default;
  string.language.go.name = 'string';
  return string;
}

function newProperty(name: string, desc: string, schema: m4.Schema): m4.Property {
  const prop = new m4.Property(name, desc, schema);
  if (helpers.isObjectSchema(schema) && schema.discriminator) {
    prop.isDiscriminator = true;
  }
  prop.language.go = prop.language.default;
  return prop;
}

function newParameter(name: string, desc: string, schema: m4.Schema): m4.Parameter {
  const param = new m4.Parameter(name, desc, schema);
  param.language.go = param.language.default;
  param.implementation = m4.ImplementationLocation.Method;
  return param;
}

function newRespProperty(name: string, desc: string, schema: m4.Schema, byValue: boolean): m4.Property {
  const prop = newProperty(name, desc, schema);
  if (helpers.isTypePassedByValue(schema)) {
    byValue = true;
  }
  prop.language.go!.byValue = byValue;
  if (schema.type === m4.SchemaType.Object) {
    // indicates we should anonymously embed this type into the containing type
    prop.language.go!.embeddedType = true;
  }
  return prop;
}

// returns the format used for marshallling/unmarshalling.
// if the media type isn't applicable then 'na' is returned.
function getMarshallingFormat(protocol: m4.Protocols): 'json' | 'xml' | 'na' {
  switch ((<m4.Protocol>protocol).http.knownMediaType) {
    case KnownMediaType.Json:
      return 'json';
    case KnownMediaType.Xml:
      return 'xml';
    default:
      return 'na';
  }
}

function recursiveTypeName(schema: m4.Schema): string {
  const rawJSON = 'RawJSON';
  switch (schema.type) {
    case m4.SchemaType.Any:
      if (schema.language.go!.rawJSONAsBytes) {
        return rawJSON;
      }
      return 'Interface';
    case m4.SchemaType.AnyObject:
      if (schema.language.go!.rawJSONAsBytes) {
        return rawJSON;
      }
      return 'Object';
    case m4.SchemaType.Array: {
      const arraySchema = <m4.ArraySchema>schema;
      const arrayElem = arraySchema.elementType;
      return `${recursiveTypeName(arrayElem)}Array`;
    }
    case m4.SchemaType.Boolean:
      return 'Bool';
    case m4.SchemaType.ByteArray:
      return 'ByteArray';
    case m4.SchemaType.Choice:
      return (<m4.ChoiceSchema>schema).language.go!.name;
    case m4.SchemaType.SealedChoice:
      return (<m4.SealedChoiceSchema>schema).language.go!.name;
    case m4.SchemaType.Date:
    case m4.SchemaType.DateTime:
    case m4.SchemaType.UnixTime:
      return 'Time';
    case m4.SchemaType.Dictionary: {
      const dictSchema = <m4.DictionarySchema>schema;
      const dictElem = dictSchema.elementType;
      return `MapOf${recursiveTypeName(dictElem)}`;
    }
    case m4.SchemaType.Integer:
      if ((<m4.NumberSchema>schema).precision === 32) {
        return 'Int32';
      }
      return 'Int64';
    case m4.SchemaType.Number:
      if ((<m4.NumberSchema>schema).precision === 32) {
        return 'Float32';
      }
      return 'Float64';
    case m4.SchemaType.Object:
      return schema.language.go!.name;
    case m4.SchemaType.Duration:
    case m4.SchemaType.String:
    case m4.SchemaType.Uuid:
      return 'String';
    default:
      throw new Error(`unhandled response schema type ${schema.type}`);
  }
}

function getRootDiscriminator(obj: m4.ObjectSchema): m4.ObjectSchema {
  // discriminators can be a root or an "intermediate" root (Salmon in the test server)

  // walk to the root
  let root = obj;
  for (;;) {
    if (!root.parents) {
      // simple case, already at the root
      break;
    }
    for (const parent of values(root.parents?.immediate)) {
      // there can be parents that aren't part of the hierarchy.
      // e.g. if type Foo is in a DictionaryOfFoo, then one of
      // Foo's parents will be DictionaryOfFoo which we ignore.
      if (helpers.isObjectSchema(parent) && parent.discriminator) {
        root.language.go!.discriminatorParent = parent.language.go!.discriminatorInterface;
        root = parent;
        // update obj with the new root.  this enables us to detect the case
        // where the parent of the new root isn't a discriminator.  consider
        // the case of child <- parent <- non-discriminator parent.  without
        // updating obj we can't detect the parent's non-discriminator as
        // root would always !== obj so the below check is always false.
        obj = root;
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
  return root;
}

// returns the set of enum values used for discriminators
function getDiscriminatorEnums(obj: m4.ObjectSchema): Array<m4.ChoiceValue> | undefined {
  if (obj.discriminator?.property.schema.type === m4.SchemaType.Choice) {
    return (<m4.ChoiceSchema>obj.discriminator.property.schema).choices;
  } else if (obj.discriminator?.property.schema.type === m4.SchemaType.SealedChoice) {
    return (<m4.SealedChoiceSchema>obj.discriminator.property.schema).choices;
  }
  return undefined;
}

// returns s in quotes if not already quoted
function quoteString(s: string): string {
  if (s[0] === '"') {
    return s;
  }
  return `"${s}"`;
}

// returns the enum name for the specified discriminator value, or undefined if missing.
function getEnumForDiscriminatorValue(discValue: string, enums: Array<m4.ChoiceValue> | undefined): string | undefined {
  if (!enums) {
    // some discriminator values are already quoted
    return quoteString(discValue);
  }
  // find the choice value that matches the current type's discriminator
  for (const enm of values(enums)) {
    if (enm.value === discValue) {
      return enm.language.go!.name;
    }
  }
  return undefined;
}

// convert comments that are in Markdown to html and then to plain text
function parseComments(comment: string): string {
  const converter = new Converter();
  converter.setOption('tables', true);
  const html = converter.makeHtml(comment);
  return fromString(html, {
    wordwrap: 200,
    tables: true,
  });
}


// This function try to label all unreferenced types before doing transform.
async function labelUnreferencedTypes(session: Session<m4.CodeModel>) {
  const model = session.model;

  const isRemoveUnreferencedTypes = await session.getValue('remove-unreferenced-types', false);

  if (!isRemoveUnreferencedTypes) return;

  const referencedTypes = new Set<m4.Schema>();

  iterateOperations(model, referencedTypes);

  labelOmitTypes(model.schemas.choices, referencedTypes);
  labelOmitTypes(model.schemas.sealedChoices, referencedTypes);
  labelOmitTypes(model.schemas.objects, referencedTypes);
}

// This function help to label omit types according to the referencedTypes set.
function labelOmitTypes<Type extends m4.Schema>(modelSchemas: Array<Type> | undefined, referencedSet: Set<Type>) {
  if (modelSchemas){
    for (const schema of modelSchemas){
      if (!referencedSet.has(schema)){
        schema.language.go!.omitType = true;
      }
    }
  }
}

// This function iterate all the operations in all operation groups and try to find all the referenced types (including choices, sealedChoices and objects).
// For each operation, we will aggregate all the parameters' type of the operation and put them into the referencedTypes set.
// Also, all responses body types and header types will be put into the referencedTypes set.
// Some special cases: exceptions response types, x-ms-odata types.
function iterateOperations(model: m4.CodeModel, referencedTypes: Set<m4.Schema>) {
  for (const group of values(model.operationGroups)) {
    for (const op of values(group.operations)) {
      for (const param of values(helpers.aggregateParameters(op))) {
        dfsSchema(param.schema, referencedTypes);
      }
      // We do not use error responses' definition to unmarshal response. So exceptions will be ignored.
      // const responses = values(op.responses).concat(values(op.exceptions));
      const responses = values(op.responses);
      for (const resp of values(responses)) {
        if (helpers.isSchemaResponse(resp)) {
          dfsSchema(resp.schema, referencedTypes);
        }
        const httpResponse = <m4.HttpResponse>resp.protocol.http;
        for (const header of values(httpResponse.headers)) {
          dfsSchema(header.schema, referencedTypes);
        }
      }
      // Such odata definition is not used directly in operation. But it seems the model is a detailed definition for some string param. Reserve for now.
      if (op.extensions?.['x-ms-odata']) {
        const schemaParts = (<string>op.extensions['x-ms-odata']).split('/');
        const schema = values(model.schemas.objects).where((o) => o.language.default.name === schemaParts[schemaParts.length - 1]).first();
        if (schema) {
          dfsSchema(schema, referencedTypes);
        }
      }
    }
  }
}

// This function will do a depth first search for the root types.
// All visited types will be put into referencedTypes set.
// Objects children/parents will also be searched.
function dfsSchema(schema: m4.Schema, referencedTypes: Set<m4.Schema>) {
  if (referencedTypes.has(schema)) return;
  referencedTypes.add(schema);
  if (helpers.isObjectSchema(schema)) {
    const allProps = helpers.aggregateProperties(schema);
    for (const prop of allProps){
      dfsSchema(prop.schema, referencedTypes);
    }
    // If schema is a discriminator, then reserve all the children
    if (schema.discriminator) {
      for (const child of values(schema.children?.all)) {
        dfsSchema(child, referencedTypes);
      }
    }
    // If schema is has allOf, then reserve all the parents
    for (const parent of values(schema.parents?.immediate)) {
      if (helpers.isObjectSchema(parent) && parent.discriminator) {
        dfsSchema(parent, referencedTypes);
      }
    }
  } else if (helpers.isArraySchema(schema)) {
    dfsSchema(schema.elementType, referencedTypes);
  } else if (helpers.isDictionarySchema(schema)) {
    dfsSchema(schema.elementType, referencedTypes);
  }
}

function separateOperationByRequestsProtocol(group: m4.OperationGroup, op: m4.Operation, defaultTypes: Array<KnownMediaType>) {
  for (const req of values(op.requests)) {
    const newOp = <m4.Operation>{...op};
    newOp.language = clone(op.language);
    newOp.requests = (<Array<m4.Request>>op.requests).filter(r => r === req);
    let name = op.language.go!.name;
    if (req.protocol.http!.knownMediaType && !defaultTypes.includes(req.protocol.http!.knownMediaType)) {
      let suffix: string;
      switch (req.protocol.http!.knownMediaType) {
        case KnownMediaType.Json:
          suffix = 'JSON';
          break;
        case KnownMediaType.Xml:
          suffix = 'XML';
          break;
        default:
          suffix = capitalize(req.protocol.http!.knownMediaType);
      }
      name = name + 'With' + suffix;
    }
    newOp.language.go!.name = name;
    newOp.language.go!.protocolNaming = new protocolMethods(newOp.language.go!.name);
    group.addOperation(newOp);
    if (req.language.go!.description) {
      req.language.go!.description = parseComments(req.language.go!.description);
    }
  }
  group.operations.splice(group.operations.indexOf(op), 1);
}
