/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { capitalize, comment, KnownMediaType, serialize, uncapitalize } from '@azure-tools/codegen';
import { AutorestExtensionHost, startSession, Session } from '@autorest/extension-base';
import { AnySchema, ObjectSchema, ArraySchema, ByteArraySchema, ChoiceValue, codeModelSchema, CodeModel, DateTimeSchema, GroupProperty, HttpHeader, HttpResponse, ImplementationLocation, Language, OperationGroup, SchemaType, NumberSchema, Operation, SchemaResponse, Parameter, Property, Protocols, Response, Schema, DictionarySchema, Protocol, ChoiceSchema, SealedChoiceSchema, ConstantSchema, Request, BooleanSchema } from '@autorest/codemodel';
import { clone, items, values } from '@azure-tools/linq';
import { aggregateParameters, getSchemaResponse, hasAdditionalProperties, isMultiRespOperation, isTypePassedByValue, isPageableOperation, isObjectSchema, isSchemaResponse, PagerInfo, PollerInfo, isLROOperation } from '../common/helpers';
import { namer, protocolMethods } from './namer';
import { fromString } from 'html-to-text';
import { Converter } from 'showdown';

// The transformer adds Go-specific information to the code model.
export async function transform(host: AutorestExtensionHost) {
  const debug = await host.GetValue('debug') || false;

  try {
    const session = await startSession<CodeModel>(host, codeModelSchema);

    // run the namer first, so that any transformations are applied on proper names
    await namer(session);
    await process(session);

    // output the model to the pipeline
    host.writeFile({
      filename: 'code-model-v4-transform.yaml',
      content: serialize(session.model),
      artifactType: 'code-model-v4'
    });

  } catch (E) {
    if (debug) {
      console.error(`${__filename} - FAILURE  ${JSON.stringify(E)} ${(<Error>E).stack}`);
    }
    throw E;
  }
}

async function process(session: Session<CodeModel>) {
  processOperationRequests(session);
  processOperationResponses(session);
  // fix up dictionary element types (additional properties)
  // this must happen before processing objects as we depend on the
  // schema type being an actual Go type.
  for (const dictionary of values(session.model.schemas.dictionaries)) {
    dictionary.elementType.language.go!.name = schemaTypeToGoType(session.model, dictionary.elementType, true);
    dictionary.language.go!.elementIsPtr = !isTypePassedByValue(dictionary.elementType);
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
        session.model.language.go!.discriminators = new Array<ObjectSchema>();
      }
      const defs = <Array<ObjectSchema>>session.model.language.go!.discriminators;
      const rootDiscriminator = getRootDiscriminator(obj);
      if (defs.indexOf(rootDiscriminator) < 0) {
        rootDiscriminator.language.go!.rootDiscriminator = true;
        defs.push(rootDiscriminator);
        // fix up discriminator value to use the enum type if available
        const discriminatorEnums = getDiscriminatorEnums(rootDiscriminator);
        // for each child type in the hierarchy, fix up the discriminator value
        for (const child of values(rootDiscriminator.children?.all)) {
          const asObj = <ObjectSchema>child;
          let discValue = getEnumForDiscriminatorValue(asObj.discriminatorValue!, discriminatorEnums);
          if (!discValue) {
            // this happens when there are intermediate types in the hierarchy that aren't discriminators.
            // e.g. BaseType -> GenericProps -> DerivedType
            // here, GenericProps is just a collection of standard properties common to various objects.
            // M4 gives it a discriminatorValue of "GenericProps" that won't exist in the enum defined
            // in BaseType.  unclear if this is desirable, but for now we just embed that value directly.
            discValue = quoteString(asObj.discriminatorValue!);
          }
          asObj.discriminatorValue = discValue;
        }
        // add the error interface as the parent interface for discriminated types that are also errors
        if (rootDiscriminator.language.go!.errorType) {
          rootDiscriminator.language.go!.discriminatorParent = 'error';
        }
      }
    }
    for (const prop of values(obj.properties)) {
      const descriptionMods = new Array<string>();
      if (prop.readOnly) {
        descriptionMods.push('READ-ONLY');
      } else if (prop.required) {
        descriptionMods.push('REQUIRED');
      }
      if (prop.language.go!.description) {
        descriptionMods.push(parseComments(prop.language.go!.description));
      }
      prop.language.go!.description = descriptionMods.join('; ');
      const details = <Language>prop.schema.language.go;
      details.name = `${schemaTypeToGoType(session.model, prop.schema, true)}`;
      if (prop.schema.type === SchemaType.Any || prop.schema.type === SchemaType.AnyObject || (isObjectSchema(prop.schema) && prop.schema.discriminator)) {
        prop.language.go!.byValue = true;
      } else if (prop.schema.type === SchemaType.DateTime) {
        obj.language.go!.needsDateTimeMarshalling = true;
      } else if (prop.schema.type === SchemaType.Date) {
        obj.language.go!.needsDateMarshalling = true;
      } else if (prop.schema.type === SchemaType.UnixTime) {
        obj.language.go!.needsUnixTimeMarshalling = true;
      } else if (prop.schema.type === SchemaType.Dictionary && obj.language.go!.marshallingFormat === 'xml') {
        // mark that we need custom XML unmarshalling for a dictionary
        prop.language.go!.needsXMLDictionaryUnmarshalling = true;
        session.model.language.go!.needsXMLDictionaryUnmarshalling = true;
      } else if (prop.schema.type === SchemaType.ByteArray) {
        obj.language.go!.byteArrayFormat = (<ByteArraySchema>prop.schema).format;
      }
      if (prop.schema.type === SchemaType.Array || prop.schema.type === SchemaType.ByteArray || prop.schema.type === SchemaType.Dictionary) {
        obj.language.go!.hasArrayMap = true;
        prop.language.go!.byValue = true;
        if (prop.schema.type !== SchemaType.Dictionary && obj.language.go!.marshallingFormat === 'xml') {
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
    const addPropsSchema = hasAdditionalProperties(obj);
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
    choice.choiceType.language.go!.name = schemaTypeToGoType(session.model, choice.choiceType, false);
    if (choice.language.go!.description) {
      choice.language.go!.description = parseComments(choice.language.go!.description);
    }
  }
  for (const choice of values(session.model.schemas.sealedChoices)) {
    choice.choiceType.language.go!.name = schemaTypeToGoType(session.model, choice.choiceType, false);
    if (choice.language.go!.description) {
      choice.language.go!.description = parseComments(choice.language.go!.description);
    }
  }
}

function schemaTypeToGoType(codeModel: CodeModel, schema: Schema, inBody: boolean): string {
  switch (schema.type) {
    case SchemaType.Any:
      return 'interface{}';
    case SchemaType.AnyObject:
      return 'map[string]interface{}';
    case SchemaType.Array:
      const arraySchema = <ArraySchema>schema;
      arraySchema.language.go!.elementIsPtr = !isTypePassedByValue(arraySchema.elementType);
      const arrayElem = <Schema>arraySchema.elementType;
      arrayElem.language.go!.name = schemaTypeToGoType(codeModel, arrayElem, inBody);
      if (<boolean>arraySchema.language.go!.elementIsPtr) {
        return `[]*${arrayElem.language.go!.name}`;
      }
      return `[]${arrayElem.language.go!.name}`;
    case SchemaType.Binary:
      return 'io.ReadSeekCloser';
    case SchemaType.Boolean:
      return 'bool';
    case SchemaType.ByteArray:
      return '[]byte';
    case SchemaType.Char:
      return 'rune';
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
      return 'time.Time';
    case SchemaType.UnixTime:
      codeModel.language.go!.hasUnixTime = true;
      if (inBody) {
        schema.language.go!.internalTimeType = 'timeUnix';
      }
      return 'time.Time';
    case SchemaType.Dictionary:
      const dictSchema = <DictionarySchema>schema;
      dictSchema.language.go!.elementIsPtr = !isTypePassedByValue(dictSchema.elementType);
      const dictElem = <Schema>dictSchema.elementType;
      dictElem.language.go!.name = schemaTypeToGoType(codeModel, dictElem, inBody);
      if (<boolean>dictSchema.language.go!.elementIsPtr) {
        return `map[string]*${dictElem.language.go!.name}`;
      }
      return `map[string]${dictElem.language.go!.name}`;
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
    case SchemaType.Credential:
    case SchemaType.Duration:
    case SchemaType.ODataQuery:
    case SchemaType.String:
    case SchemaType.Uuid:
    case SchemaType.Uri:
      return 'string';
    case SchemaType.Date:
      if (inBody) {
        codeModel.language.go!.hasDate = true;
        schema.language.go!.internalTimeType = 'dateType';
      }
      return 'time.Time';
    case SchemaType.Time:
      return 'time.Time';
    case SchemaType.Choice:
    case SchemaType.SealedChoice:
    case SchemaType.Object:
      return schema.language.go!.name;
    default:
      throw new Error(`unhandled schema type ${schema.type}`);
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
  // pre-process multi-request operations as it can add operations to the operations
  // collection, and iterating over a modified collection yeilds incorrect results
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations).toArray()) {
      if (op.language.go!.description) {
        op.language.go!.description = parseComments(op.language.go!.description);
      }
      if (op.requests!.length > 1) {
        for (const req of values(op.requests)) {
          const newOp = clone(op);
          newOp.requests = (<Array<Request>>op.requests).filter(r => r === req);
          let name = op.language.go!.name;
          // for the non-binary media types we create a new method with the
          // media type name as a suffix, e.g. FooAPIWithJSON()
          if (req.protocol.http!.knownMediaType !== KnownMediaType.Binary) {
            let suffix: string;
            switch (req.protocol.http!.knownMediaType) {
              case KnownMediaType.Json:
                suffix = 'JSON';
                break;
              case KnownMediaType.Xml:
                suffix = 'XML';
                break;
              default:
                suffix = capitalize(<string>req.protocol.http!.knownMediaType);
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
    }
  }
  // track any client-level parameterized host params.
  const hostParams = new Array<Parameter>();
  // track any parameter groups and/or optional parameters
  const paramGroups = new Map<string, GroupProperty>();
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      if (op.language.go!.description) {
        op.language.go!.description = parseComments(op.language.go!.description);
      }
      if (op.requests![0].protocol.http!.headers) {
        for (const header of values(op.requests![0].protocol.http!.headers)) {
          const head = <HttpHeader>header;
          head.schema.language.go!.name = schemaTypeToGoType(session.model, head.schema, false);
        }
      }
      for (const param of values(aggregateParameters(op))) {
        if (param.language.go!.description) {
          param.language.go!.description = parseComments(param.language.go!.description);
        }
        if (!param.required && param.schema.type === SchemaType.Constant && !param.language.go!.amendedDesc) {
          if (param.language.go!.description) {
            param.language.go!.description += '. ';
          }
          param.language.go!.description += `Specifying any value will set the value to ${(<ConstantSchema>param.schema).value.value}.`;
          param.language.go!.amendedDesc = true;
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
          for (const internalOp of values(group.operations)) {
            if (internalOp.language.go!.name === dupeOp.language.default.name) {
              internalOp.language.go!.paging.isNextOp = true;
              break;
            }
          }
        }
        const inBody = param.protocol.http !== undefined && param.protocol.http!.in === 'body';
        param.schema.language.go!.name = schemaTypeToGoType(session.model, param.schema, inBody);
        if (isTypePassedByValue(param.schema)) {
          param.language.go!.byValue = true;
        }
        // check if this is a header collection
        if (param.extensions?.['x-ms-header-collection-prefix']) {
          param.schema.language.go!.headerCollectionPrefix = param.extensions['x-ms-header-collection-prefix'];
        }
        if (param.implementation === ImplementationLocation.Client && (param.schema.type !== SchemaType.Constant || !param.required) && param.language.default.name !== '$host') {
          if (param.protocol.http!.in === 'uri') {
            // this is a parameterized host param.
            // use the param name to avoid reference equality checks.
            if (!values(hostParams).where(p => p.language.go!.name === param.language.go!.name).any()) {
              hostParams.push(param);
            }
            continue;
          }
          // add global param info to the operation group
          if (group.language.go!.clientParams === undefined) {
            group.language.go!.clientParams = new Array<Parameter>();
          }
          const clientParams = <Array<Parameter>>group.language.go!.clientParams;
          // check if this global param has already been added
          if (values(clientParams).where(cp => cp.language.go!.name === param.language.go!.name).any()) {
            continue;
          }
          clientParams.push(param);
        } else if (param.implementation === ImplementationLocation.Method && param.protocol.http!.in === 'uri') {
          // at least one method contains a parameterized host param, bye-bye simple case
          group.language.go!.complexHostParams = true;
        }
        let opName = op.language.go!.name;
        if (isLROOperation(op)) {
          opName = 'Begin' + opName;
        }
        // check for grouping
        if (param.extensions?.['x-ms-parameter-grouping'] && <boolean>session.model.language.go!.groupParameters) {
          // this param belongs to a param group, init name with default
          let paramGroupName = `${group.language.go!.name}${opName}Parameters`;
          if (param.extensions['x-ms-parameter-grouping'].name) {
            // use the specified name
            paramGroupName = <string>param.extensions['x-ms-parameter-grouping'].name;
          } else if (param.extensions['x-ms-parameter-grouping'].postfix) {
            // use the suffix
            paramGroupName = `${group.language.go!.name}${opName}${<string>param.extensions['x-ms-parameter-grouping'].postfix}`;
          }
          // create group entry and add the param to it
          if (!paramGroups.has(paramGroupName)) {
            const desc = `${paramGroupName} contains a group of parameters for the ${group.language.go!.name}.${opName} method.`;
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
            throw new Error(`parameter group ${paramGroupName} contains overlapping parameters with different schemas`);
          }
          continue;
        }
        // create an optional params struct even if the operation contains no optional params.
        // this provides version resiliency in case optional params are added in the future.
        // don't do this for paging next link operation as this isn't part of the public API
        if (op.language.go!.paging && op.language.go!.paging.isNextOp) {
          continue;
        }
        // create a type named <OperationGroup><Operation>Options
        const paramGroupName = `${group.language.go!.name}${opName}Options`;
        if (!paramGroups.has(paramGroupName)) {
          const desc = `${paramGroupName} contains the optional parameters for the ${group.language.go!.name}.${opName} method.`;
          const gp = createGroupProperty(paramGroupName, desc);
          gp.required = false;
          paramGroups.set(paramGroupName, gp);
          // associate the param group with the operation
          op.language.go!.optionalParamGroup = gp;
        }
        // include non-required constants that aren't body params in the optional values struct
        if (param.required !== true && !(param.schema.type === SchemaType.Constant && param.protocol.http!.in === 'body')) {
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
      session.model.language.go!.parameterGroups = new Array<GroupProperty>();
    }
    const pg = <Array<GroupProperty>>session.model.language.go!.parameterGroups;
    for (const items of paramGroups.entries()) {
      pg.push(items[1]);
    }
  }  
}

function createGroupProperty(name: string, description: string): GroupProperty {
  const schema = new ObjectSchema(name, description);
  schema.language.go = schema.language.default;
  const gp = new GroupProperty(name, description, schema);
  gp.language.go = gp.language.default;
  return gp;
}

function processOperationResponses(session: Session<CodeModel>) {
  if (session.model.language.go!.responseEnvelopes === undefined) {
    session.model.language.go!.responseEnvelopes = new Array<Schema>();
  }
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      // annotate all exception types as errors; this is so we know to generate an Error() method
      const exTypeNames = new Array<string>();
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
            if (prop.extensions?.['x-ms-client-flatten'] === true) {
              schemaError.language.go!.flattenedErr = prop.language.go!.name;
            }
          }
          // annotate all child and parent error types.  note that errorType has
          // special significance which is why we use inheritedErrorType instead.
          for (const child of values(schemaError.children?.all)) {
            if (isObjectSchema(child)) {
              child.language.go!.inheritedErrorType = 'child';
            }
          }
          for (const parent of values(schemaError.parents?.all)) {
            if (isObjectSchema(parent)) {
              parent.language.go!.inheritedErrorType = 'parent';
            }
          }
          if (schemaError.discriminator) {
            // if the error is a discriminator we need to create an internal wrapper type
            schemaError.language.go!.internalErrorType = uncapitalize(<string>schemaError.language.go!.name);
          }
          if (schemaError.discriminator) {
            for (const dt of values(<Array<string>>schemaError.language.go!.discriminatorTypes)) {
              if (!exTypeNames.includes(dt)) {
                exTypeNames.push(dt);
              }
            }
          } else {
            exTypeNames.push('*' + schemaError.language.go!.name);
          }
        } else {
          schemaError.language.go!.name = schemaTypeToGoType(session.model, schemaError, true);
          if (schemaError.type === SchemaType.Any || schemaError.type === SchemaType.AnyObject) {
            ex.language.go!.genericError = true;
            continue;
          }
        }
        schemaError.language.go!.errorType = true;
        recursiveAddMarshallingFormat(schemaError, marshallingFormat);
      }
      op.language.go!.description += '\nIf the operation fails it returns ';
      if (exTypeNames.length > 1) {
        exTypeNames.sort();
        op.language.go!.description += `one of the following error types.\n`;
        op.language.go!.description += comment(exTypeNames.join(', '), ' - ');
      } else if (exTypeNames.length === 1) {
        op.language.go!.description += `the ${exTypeNames[0]} error type.`;
      } else {
        op.language.go!.description += 'a generic error.';
      }
      // recursively add the marshalling format to the responses if applicable.
      // also remove any HTTP redirects from the list of responses.
      const filtered = new Array<Response>();
      for (const resp of values(op.responses)) {
        if (skipRedirectStatusCode(<string>op.requests![0].protocol.http!.method, resp)) {
          // redirects are transient status codes, they aren't actually returned
          continue;
        }
        if (isSchemaResponse(resp)) {
          if (resp.schema.type === SchemaType.Binary) {
            // don't create response envelopes for binary responses.
            // callers read directly from the *http.Response.Body
            continue;
          }
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
      // create pageable type info before poller as the pager-poller depends on it
      if (isPageableOperation(op)) {
        if (session.model.language.go!.pageableTypes === undefined) {
          session.model.language.go!.pageableTypes = new Array<PagerInfo>();
        }
        const name = `${group.language.go!.name}${op.language.go!.name}Pager`;
        // create a new one, add to global list and assign to method
        const pager = {
          name: name,
          op: op,
        };
        const pagers = <Array<PagerInfo>>session.model.language.go!.pageableTypes;
        pagers.push(pager);
        op.language.go!.pageableType = pager;
      }
      if (isLROOperation(op)) {
        if (session.model.language.go!.pollerTypes === undefined) {
          session.model.language.go!.pollerTypes = new Array<PollerInfo>();
        }
        createLROResponseEnvelope(session.model, group, op);
      } else {
        createResponseEnvelope(session.model, group, op);
      }
    }
  }
}

// returns true if the specified status code is an automatic HTTP redirect.
// certain redirects are automatically handled by the HTTP stack and thus are
// transient so they are never actually returned to the caller.  we skip them
// so they aren't included in the potential result set of an operation.
function skipRedirectStatusCode(verb: string, resp: Response): boolean {
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

interface HttpHeaderWithDescription extends HttpHeader {
  description: string;
}

// the name of the struct field for scalar responses (int, string, etc)
const scalarResponsePropName = 'Value';

// creates the response/result envelope types to be returned from an operation and updates the operation.
// for LROs, this is also called to create the final response envelope.
// the response envelope consists of two parts, the outer response envelope and the inner result envelope.
// each operation gets its own response/result envelope pair
//
// type GetWidgetResponse struct { <== this is the response envelope, it groups the result with the raw HTTP response (to be replaced by Result[T any])
//   GetWidgetResult
//   RawResponse *http.Response
//}
//
// type GetWidgetResult struct { <== this is the result envelope, it groups the schema result along with any per-operation header values
//   Widget <== this is the result field, i.e. the schema result
//   Header1 *string <== modeled header response
//   Header2 *int    <== modeled header response
//}
function createResponseEnvelope(codeModel: CodeModel, group: OperationGroup, op: Operation) {
  // create the `type <type>Response struct` response
  // type with a `RawResponse *http.Response` field

  // aggregate headers from all responses as all of them will go into the same result envelope
  const headers = new Map<string, HttpHeaderWithDescription>();
  // skip adding headers for LROs for now, this is to avoid adding
  // any LRO-specific polling headers.
  // TODO: maybe switch to skip specific polling ones in the future?
  if (!isLROOperation(op)) {
    for (const resp of values(op.responses)) {
      // check if the response is expecting information from headers
      for (const header of values(resp.protocol.http!.headers)) {
        const head = <HttpHeader>header;
        // convert each header to a property and append it to the response properties list
        const name = head.language.go!.name;
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
  const addHeadersToSchema = function (resultEnv: ObjectSchema) {
    if (headers.size === 0) {
      return;
    }
    if (!resultEnv.properties) {
      resultEnv.properties = new Array<Property>();
    }
    for (const item of items(headers)) {
      const prop = newRespProperty(item.key, item.value.description, item.value.schema, false);
      // propagate any extensions so we can access them through the property
      prop.extensions = item.value.extensions;
      prop.language.go!.fromHeader = item.value.header;
      resultEnv.properties.push(prop);
    }
  }
  // contains all the response envelopes
  const responseEnvelopes = <Array<ObjectSchema>>codeModel.language.go!.responseEnvelopes;
  // first create the response envelope, each operation gets one
  const respEnvName = ensureUniqueModelName(codeModel, `${group.language.go!.name}${op.language.go!.name}Response`, 'Envelope');
  const respEnv = newObject(respEnvName, `${respEnvName} contains the response from method ${group.language.go!.name}.${op.language.go!.name}.`);
  respEnv.language.go!.responseType = true;
  respEnv.properties = new Array<Property>();
  respEnv.properties.push(createRawResponseProp());
  responseEnvelopes.push(respEnv);
  op.language.go!.responseEnv = respEnv;

  // now create the appropriate result envelope

  if (codeModel.language.go!.headAsBoolean && op.requests![0].protocol.http!.method === 'head') {
    op.language.go!.headAsBoolean = true;
    const name = ensureUniqueModelName(codeModel, `${group.language.go!.name}${op.language.go!.name}Result`, 'Envelope');
    const resultEnv = newObject(name, `${name} contains the result from method ${group.language.go!.name}.${op.language.go!.name}.`);
    resultEnv.language.go!.responseType = true;
    const successProp = newProperty('Success', 'Success indicates if the operation succeeded or failed.', newBoolean('bool', 'bool response'));
    successProp.language.go!.byValue = true;
    resultEnv.properties = [
      successProp,
    ];
    addHeadersToSchema(resultEnv);
    // now add the result envelope to the response envelope
    const resultProp = newRespProperty(name, 'Contains the result of the operation.', resultEnv, true);
    respEnv.properties.push(resultProp);
    respEnv.language.go!.resultEnv = resultProp;
    op.language.go!.responseEnv = respEnv;
    return;
  }
  if (isMultiRespOperation(op)) {
    const resultTypes = new Array<string>();
    for (const response of values(op.responses)) {
      // the operation might contain a mix of schemas and non-schema responses.
      // we only care about the ones that return a schema.
      if (isSchemaResponse(response)) {
        resultTypes.push(response.schema.language.go!.name);
      }
    }
    // for multi-response operations, we don't create result envelopes.
    // instead, we add the header responses to the response envelope.
    addHeadersToSchema(respEnv);
    const resultProp = newRespProperty('Value', `// Possible types are ${resultTypes.join(', ')}\n`, newAny('multi-response value'), true);
    respEnv.properties.push(resultProp);
    respEnv.language.go!.resultEnv = resultProp;
    return;
  }
  const response = getSchemaResponse(op);
  // if the response defines a schema then create a result envelope and add it to the response envelope.
  if (response) {
    const name = ensureUniqueModelName(codeModel, `${group.language.go!.name}${op.language.go!.name}Result`, 'Envelope');
    const resultEnv = newObject(name, `${name} contains the result from method ${group.language.go!.name}.${op.language.go!.name}.`);
    resultEnv.language.go!.responseType = true;
    // propagate marshalling format to the result envelope
    resultEnv.language.go!.marshallingFormat = response.schema.language.go!.marshallingFormat;
    resultEnv.properties = new Array<Property>();
    // for operations that return scalar types we use a fixed field name
    let propName = scalarResponsePropName;
    if (response.schema.type === SchemaType.Object) {
      // for object types use the type's name as the field name
      propName = response.schema.language.go!.name;
    } else if (response.schema.type === SchemaType.Array) {
      // for array types use the element type's name
      propName = recursiveTypeName(response.schema);
    } else if (response.schema.type === SchemaType.Any) {
      propName = 'Interface';
    } else if (response.schema.type === SchemaType.AnyObject) {
      propName = 'Object';
    }
    if (response.schema.serialization?.xml && response.schema.serialization.xml.name) {
      // always prefer the XML name
      propName = capitalize(<string>response.schema.serialization.xml.name);
    }
    // add any headers to the response type
    addHeadersToSchema(resultEnv);
    // we want to pass integral types byref to maintain parity with struct fields
    const byValue = isTypePassedByValue(response.schema) || response.schema.type === SchemaType.Object;
    const respProp = newRespProperty(propName, response.schema.language.go!.description, response.schema, byValue);
    resultEnv.properties.push(respProp);
    // now add the result type to the response envelope
    const resultEnvProp = newRespProperty(name, 'Contains the result of the operation.', resultEnv, true);
    respEnv.properties.push(resultEnvProp);
    respEnv.language.go!.resultEnv = resultEnvProp;
    // shortcut to aid finding the result property
    resultEnvProp.language.go!.resultField = respProp;
  } else if (headers.size > 0) {
    // the response doesn't return a model.  if it returns
    // headers then create a result envelope that contains them.
    const name = ensureUniqueModelName(codeModel, `${group.language.go!.name}${op.language.go!.name}Result`, 'Envelope');
    const resultEnv = newObject(name, `${name} contains the result from method ${group.language.go!.name}.${op.language.go!.name}.`);
    resultEnv.language.go!.responseType = true;
    resultEnv.properties = new Array<Property>();
    addHeadersToSchema(resultEnv);
    const resultEnvProp = newRespProperty(name, 'Contains the result of the operation.', resultEnv, true);
    respEnv.properties.push(resultEnvProp);
    respEnv.language.go!.resultEnv = resultEnvProp;
  }
}

// appends suffix to name if name is an existing model type
function ensureUniqueModelName(codeModel: CodeModel, name: string, suffix: string): string {
  for (const obj of values(codeModel.schemas.objects)) {
    if (obj.language.go!.name === name) {
      return name + suffix;
    }
  }
  return name;
}

function createRawResponseProp(): Property {
  return newProperty('RawResponse', 'RawResponse contains the underlying HTTP response.', newObject('http.Response', 'raw HTTP response'));
}

function newObject(name: string, desc: string): ObjectSchema {
  const obj = new ObjectSchema(name, desc);
  obj.language.go = obj.language.default;
  return obj;
}

function newAny(desc: string): AnySchema {
  const any = new AnySchema(desc);
  any.language.go = any.language.default;
  any.language.go!.name = 'interface{}';
  return any;
}

function newBoolean(name: string, desc: string): BooleanSchema {
  const bool = new BooleanSchema(name, desc);
  bool.language.go = bool.language.default;
  bool.language.go!.name = 'bool';
  return bool;
}

function newProperty(name: string, desc: string, schema: Schema): Property {
  let prop = new Property(name, desc, schema);
  if (isObjectSchema(schema) && schema.discriminator) {
    prop.isDiscriminator = true;
  }
  prop.language.go = prop.language.default;
  return prop;
}

function newRespProperty(name: string, desc: string, schema: Schema, byValue: boolean): Property {
  const prop = newProperty(name, desc, schema);
  if (isTypePassedByValue(schema)) {
    byValue = true;
  }
  prop.language.go!.byValue = byValue;
  if (schema.type === SchemaType.Object) {
    // indicates we should anonymously embed this type into the containing type
    prop.language.go!.embeddedType = true;
  }
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

function recursiveTypeName(schema: Schema): string {
  switch (schema.type) {
    case SchemaType.Any:
      return 'Interface';
    case SchemaType.AnyObject:
      return 'Object';
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
    case SchemaType.Duration:
    case SchemaType.String:
    case SchemaType.Uuid:
      return 'String';
    default:
      throw new Error(`unhandled response schema type ${schema.type}`);
  }
}

// creates the LRO response envelope for the specified operation/response
function createLROResponseEnvelope(codeModel: CodeModel, group: OperationGroup, op: Operation) {
  if (op.language.go!.paging && op.language.go!.paging.member) {
    // implementing support for this is very complicated, and since
    // no services at present use this pattern avoid it for now
    throw new Error(`${op.language.go!.name}: unsupported pager-poller that uses next page operation`);
  }
  // LROs have two response envelopes.
  // the outer is the response envelope returned by the Begin* and Resume* methods, it depends on the poller.
  // the inner is the response/result envelope returned by the PollUntilDone and FinalResponse methods, the poller depends on it.
  // so we create them in the following order: inner, poller, outer

  // contains all the response envelopes
  const responseEnvelopes = <Array<ObjectSchema>>codeModel.language.go!.responseEnvelopes;
  // create the "inner" response envelope
  createResponseEnvelope(codeModel, group, op);
  // now create the poller
  if (op.language.go!.paging && op.language.go!.paging.member) {
    // implementing support for this is very complicated, and since
    // no services at present use this pattern avoid it for now
    throw new Error(`${op.language.go!.name}: unsupported pager-poller that uses next page operation`);
  }
  if (codeModel.language.go!.pollerTypes === undefined) {
    codeModel.language.go!.pollerTypes = new Array<PollerInfo>();
  }
  const pollerName = `${group.language.go!.name}${op.language.go!.name}Poller`;
  const pollers = <Array<PollerInfo>>codeModel.language.go!.pollerTypes;
  const poller = {
    name: pollerName,
    op: op,
  };
  pollers.push(poller);
  op.language.go!.pollerType = poller;

  // finally create the outer response envelope
  const outerRespEnvName = `${group.language.go!.name}${op.language.go!.name}PollerResponse`;
  const outerRespEnv = newObject(outerRespEnvName, `${outerRespEnvName} contains the response from method ${group.language.go!.name}.${op.language.go!.name}.`);
  outerRespEnv.language.go!.responseType = true;
  outerRespEnv.language.go!.isLRO = true;
  outerRespEnv.language.go!.pollerInfo = poller;
  outerRespEnv.properties = new Array<Property>();
  outerRespEnv.properties.push(createRawResponseProp());
  // create Poller
  const pollerType = newObject(pollerName, 'poller');
  const pollerProp = newProperty('Poller', 'Poller contains an initialized poller.', pollerType);
  outerRespEnv.properties.push(pollerProp);
  responseEnvelopes.push(outerRespEnv);
  // move the inner response envelope created earlier
  op.language.go!.finalResponseEnv = op.language.go!.responseEnv;
  op.language.go!.responseEnv = outerRespEnv;
}

function getRootDiscriminator(obj: ObjectSchema): ObjectSchema {
  // discriminators can be a root or an "intermediate" root (Salmon in the test server)

  // walk to the root
  let root = obj;
  while (true) {
    if (!root.parents) {
      // simple case, already at the root
      break;
    }
    for (const parent of values(root.parents?.immediate)) {
      // there can be parents that aren't part of the hierarchy.
      // e.g. if type Foo is in a DictionaryOfFoo, then one of
      // Foo's parents will be DictionaryOfFoo which we ignore.
      if (isObjectSchema(parent) && parent.discriminator) {
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
function getDiscriminatorEnums(obj: ObjectSchema): Array<ChoiceValue> | undefined {
  if (obj.discriminator?.property.schema.type === SchemaType.Choice) {
    return (<ChoiceSchema>obj.discriminator!.property.schema).choices;
  } else if (obj.discriminator?.property.schema.type === SchemaType.SealedChoice) {
    return (<SealedChoiceSchema>obj.discriminator!.property.schema).choices;
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
function getEnumForDiscriminatorValue(discValue: string, enums: Array<ChoiceValue> | undefined): string | undefined {
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
  let converter = new Converter();
  converter.setOption('tables', true);
  let html = converter.makeHtml(comment);
  return fromString(html, {
    wordwrap: 200,
    tables: true,
  });
}
