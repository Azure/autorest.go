/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { comment } from '@azure-tools/codegen';
import { CodeModel, ComplexSchema, ConstantSchema, DictionarySchema, GroupProperty, ImplementationLocation, ObjectSchema, Language, Schema, SchemaType, Parameter, Property } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { isArraySchema, isObjectSchema, hasAdditionalProperties, hasPolymorphicField, commentLength } from '../common/helpers';
import { contentPreamble, elementByValueForParam, hasDescription, sortAscending, substituteDiscriminator } from './helpers';
import { ImportManager } from './imports';

// Creates the content in models.go
export async function generateModels(session: Session<CodeModel>): Promise<string> {
  // this list of packages to import
  const imports = new ImportManager();
  let text = await contentPreamble(session);

  // we do model generation first as it can add imports to the imports list
  const structs = generateStructs(imports, session.model.schemas.objects);
  const responseEnvelopes = <Array<Schema>>session.model.language.go!.responseEnvelopes;
  for (const respEnv of values(responseEnvelopes)) {
    const respType = generateStruct(imports, respEnv.language.go!.responseType, respEnv.language.go!.properties);
    generateUnmarshallerForResponseEnvelope(respType);
    structs.push(respType);
  }
  const paramGroups = <Array<GroupProperty>>session.model.language.go!.parameterGroups;
  for (const paramGroup of values(paramGroups)) {
    structs.push(generateParamGroupStruct(imports, paramGroup.schema.language.go!, paramGroup.originalParameter));
  }

  // imports
  if (imports.length() > 0) {
    text += imports.text();
  }

  // structs
  let needsJSONPopulate = false;
  let needsJSONUnpopulate = false;
  structs.sort((a: StructDef, b: StructDef) => { return sortAscending(a.Language.name, b.Language.name) });
  for (const struct of values(structs)) {
    text += struct.discriminator();
    text += struct.text();
    struct.Methods.sort((a: StructMethod, b: StructMethod) => { return sortAscending(a.name, b.name) });
    for (const method of values(struct.Methods)) {
      if (method.desc.length > 0) {
        text += `${comment(method.desc, '// ', undefined, commentLength)}\n`;
      }
      text += method.text;
    }
    if (struct.HasJSONMarshaller) {
      needsJSONPopulate = true;
    }
    if (struct.HasJSONUnmarshaller) {
      needsJSONUnpopulate = true;
    }
  }
  if (needsJSONPopulate) {
    text += 'func populate(m map[string]interface{}, k string, v interface{}) {\n';
    text += '\tif v == nil {\n';
    text += '\t\treturn\n';
    text += '\t} else if azcore.IsNullValue(v) {\n';
    text += '\t\tm[k] = nil\n';
    text += '\t} else if !reflect.ValueOf(v).IsNil() {\n';
    text += '\t\tm[k] = v\n';
    text += '\t}\n';
    text += '}\n\n';
  }
  if (needsJSONUnpopulate) {
    text += 'func unpopulate(data json.RawMessage, v interface{}) error {\n';
    text += '\tif data == nil {\n';
    text += '\t\treturn nil\n';
    text += '\t}\n';
    text += '\treturn json.Unmarshal(data, v)\n';
    text += '}\n\n';
  }
  return text;
}

interface StructMethod {
  name: string;
  desc: string;
  text: string;
}

// represents a struct definition
class StructDef {
  readonly Language: Language;
  readonly Properties?: Property[];
  readonly Parameters?: Parameter[];
  readonly Methods: StructMethod[];
  readonly ComposedOf: ObjectSchema[];
  HasJSONMarshaller: boolean;
  HasJSONUnmarshaller: boolean;

  constructor(language: Language, props?: Property[], params?: Parameter[]) {
    this.Language = language;
    this.Properties = props;
    this.Parameters = params;
    if (this.Properties) {
      this.Properties.sort((a: Property, b: Property) => { return sortAscending(a.language.go!.name, b.language.go!.name); });
    }
    if (this.Parameters) {
      this.Parameters.sort((a: Parameter, b: Parameter) => { return sortAscending(a.language.go!.name, b.language.go!.name); });
    }
    this.Methods = new Array<StructMethod>();
    this.ComposedOf = new Array<ObjectSchema>();
    this.HasJSONMarshaller = false;
    this.HasJSONUnmarshaller = false;
  }

  text(): string {
    let text = '';
    if (hasDescription(this.Language)) {
      text += `${comment(this.Language.description, '// ', undefined, commentLength)}\n`;
    }
    if (this.Language.errorType) {
      text += '// Implements the error and azcore.HTTPResponse interfaces.\n';
    }
    text += `type ${this.Language.name} struct {\n`;
    // any composed types go first
    for (const comp of values(this.ComposedOf)) {
      text += `\t${comp.language.go!.name}\n`;
    }
    // used to track when to add an extra \n between fields that have comments
    let first = true;
    if (this.Properties === undefined && this.Parameters?.length === 0) {
      // this is an optional params placeholder struct
      text += '\t// placeholder for future optional parameters\n';
    }
    if (this.Language.errorType) {
      text += '\traw string\n';
    }
    for (const prop of values(this.Properties)) {
      if (hasDescription(prop.language.go!)) {
        if (!first) {
          // add an extra new-line between fields IFF the field
          // has a comment and it's not the very first one.
          text += '\n';
        }
        text += `\t${comment(prop.language.go!.description, '// ', undefined, commentLength)}\n`;
      }
      let elemByVal = false;
      if (prop.schema.type === SchemaType.Dictionary && prop.extensions?.['x-ms-header-collection-prefix']) {
        elemByVal = true;
      }
      let typeName = substituteDiscriminator(prop.schema, elemByVal);
      if (prop.schema.type === SchemaType.Constant) {
        // for constants we use the underlying type name
        typeName = (<ConstantSchema>prop.schema).valueType.language.go!.name;
      }
      let serialization = prop.serializedName;
      if (this.Language.marshallingFormat === 'json') {
        serialization += ',omitempty';
      } else if (this.Language.marshallingFormat === 'xml') {
        serialization = getXMLSerialization(prop, this.Language);
      }
      let readOnly = '';
      if (prop.readOnly) {
        readOnly = ` azure:"ro"`;
      }
      let tag = ` \`${this.Language.marshallingFormat}:"${serialization}"${readOnly}\``;
      // if this is a response type then omit the tag IFF the marshalling format is
      // JSON, it's a header or is the RawResponse field.  XML marshalling needs a tag.
      // also omit the tag for additionalProperties
      if ((this.Language.responseType === true && (this.Language.marshallingFormat !== 'xml' || prop.language.go!.name === 'RawResponse')) || prop.language.go!.isAdditionalProperties) {
        tag = '';
      }
      let pointer = '*';
      if (prop.language.go!.byValue === true) {
        pointer = '';
      }
      text += `\t${prop.language.go!.name} ${pointer}${typeName}${tag}\n`;
      first = false;
    }
    for (const param of values(this.Parameters)) {
      // if Parameters is set this is a param group struct
      // none of its fields need to participate in marshalling
      if (param.implementation === ImplementationLocation.Client) {
        // don't add globals to the per-method options struct
        continue;
      }
      if (hasDescription(param.language.go!)) {
        text += `\t${comment(param.language.go!.description, '// ', undefined, commentLength)}\n`;
      }
      let pointer = '*';
      if (param.required || param.language.go!.byValue === true) {
        pointer = '';
      }
      const typeName = substituteDiscriminator(param.schema, elementByValueForParam(param));
      text += `\t${(<string>param.language.go!.name).capitalize()} ${pointer}${typeName}\n`;
    }
    text += '}\n\n';
    return text;
  }

  discriminator(): string {
    if (!this.Language.discriminatorInterface) {
      return '';
    }
    const methodName = `Get${this.Language.name}`;
    let text = `// ${this.Language.discriminatorInterface} provides polymorphic access to related types.\n`;
    text += `// Call the interface's ${methodName}() method to access the common type.\n`;
    text += `// Use a type switch to determine the concrete type.  The possible types are:\n`;
    text += comment((<Array<string>>this.Language.discriminatorTypes).join(', '), '// - ');
    text += `\ntype ${this.Language.discriminatorInterface} interface {\n`;
    if (this.Language.discriminatorParent) {
      text += `\t${this.Language.discriminatorParent}\n`;
    }
    text += `\t// ${methodName} returns the ${this.Language.name} content of the underlying type.\n`;
    text += `\t${methodName}() *${this.Language.name}\n`;
    text += '}\n\n';
    return text;
  }

  receiverName(): string {
    const typeName = this.Language.name;
    return typeName[0].toLowerCase();
  }
}

function getXMLSerialization(prop: Property, lang: Language): string {
  let serialization = prop.serializedName;
  // default to using the serialization name
  if (prop.schema.serialization?.xml?.name) {
    // xml can specifiy its own name, prefer that if available
    serialization = prop.schema.serialization.xml.name;
  }
  if (prop.schema.serialization?.xml?.attribute) {
    // value comes from an xml attribute
    serialization += ',attr';
  } else if (isArraySchema(prop.schema)) {
    // start with the serialized name of the element, preferring xml name if available
    let inner = prop.schema.elementType.language.go!.name;
    if (prop.schema.elementType.serialization?.xml?.name) {
      inner = prop.schema.elementType.serialization.xml.name;
    }
    // arrays can be wrapped or unwrapped.  here's a wrapped example
    // note how the array of apple objects is "wrapped" in GoodApples
    // <AppleBarrel>
    //   <GoodApples>
    //     <Apple>Fuji</Apple>
    //     <Apple>Gala</Apple>
    //   </GoodApples>
    // </AppleBarrel>

    // here's an unwrapped example, the array of slide objects
    // is embedded directly in the object (no "wrapping")
    // <slideshow>
    //   <slide>
    //     <title>Wake up to WonderWidgets!</title>
    //   </slide>
    //   <slide>
    //     <title>Overview</title>
    //   </slide>
    // </slideshow>

    // arrays in the response type are handled slightly different as we
    // unmarshal directly into them so no need to add the unwrapping.
    if (prop.schema.serialization?.xml?.wrapped && lang.responseType !== true) {
      serialization += `>${inner}`;
    } else {
      serialization = inner;
    }
  }
  return serialization;
}

function generateStructs(imports: ImportManager, objects?: ObjectSchema[]): StructDef[] {
  const structTypes = new Array<StructDef>();
  for (const obj of values(objects)) {
    const props = new Array<Property>();
    // add immediate properties
    for (const prop of values(obj.properties)) {
      props.push(prop);
    }
    const structDef = generateStruct(imports, obj.language.go!, props);
    // now add the parent type
    let parentType: ObjectSchema | undefined;
    for (const parent of values(obj.parents?.immediate)) {
      if (isObjectSchema(parent)) {
        parentType = parent;
        structDef.ComposedOf.push(parent);
      }
    }
    structDef.ComposedOf.sort((a: ObjectSchema, b: ObjectSchema) => { return sortAscending(a.language.go!.name, b.language.go!.name); });
    if (obj.language.go!.errorType) {
      // add Error() method
      let text = `func (e ${obj.language.go!.name}) Error() string {\n`;
      text += `\treturn e.raw\n`;
      text += '}\n\n';
      structDef.Methods.push({ name: 'Error', desc: `Error implements the error interface for type ${obj.language.go!.name}.\nThe contents of the error text are not contractual and subject to change.`, text: text });
    }
    if (obj.language.go!.marshallingFormat === 'xml') {
      // due to differences in XML marshallers/unmarshallers, we use different codegen than for JSON
      if (obj.language.go!.needsDateTimeMarshalling || obj.language.go!.xmlWrapperName || needsXMLArrayMarshalling(obj)) {
        generateXMLMarshaller(structDef);
        if (obj.language.go!.needsDateTimeMarshalling) {
          generateXMLUnmarshaller(structDef);
        }
      } else if (needsXMLDictionaryUnmarshalling(obj)) {
        generateXMLUnmarshaller(structDef);
      }
      structTypes.push(structDef);
      continue;
    }
    if (obj.discriminator) {
      // only need to generate interface method and internal marshaller for discriminators (Fish, Salmon, Shark)
      generateDiscriminatorMarkerMethod(obj, structDef);
    }
    const needs = determineMarshallers(obj);
    if (needs.intM) {
      generateInternalMarshaller(obj, structDef, parentType);
    }
    if (needs.intU) {
      generateInternalUnmarshaller(obj, structDef, parentType);
    }
    if (needs.M) {
      imports.add('reflect');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
      structDef.HasJSONMarshaller = true;
      generateJSONMarshaller(imports, obj, structDef, parentType);
    }
    if (needs.U) {
      structDef.HasJSONUnmarshaller = true;
      generateJSONUnmarshaller(imports, obj, structDef, parentType);
    }
    structTypes.push(structDef);
  }
  return structTypes;
}

interface Marshallers {
  intM: boolean
  intU: boolean
  M: boolean
  U: boolean
}

function mergeMarshallers(lhs: Marshallers, rhs: Marshallers): Marshallers {
  return {
    intM: lhs.intM || rhs.intM,
    intU: lhs.intU || rhs.intU,
    M: lhs.M || rhs.M,
    U: lhs.U || rhs.U
  }
}

// determines the marshallers need for the specified object.
// it examines the complete inheritance graph to make the determination.
function determineMarshallers(obj: ObjectSchema): Marshallers {
  // things that require internal marshallers and/or unmarshallers:
  //   inheritance
  //   discriminated types

  const parents = recursiveDetermineMarshallers(obj, true);
  const children = recursiveDetermineMarshallers(obj, false);
  return mergeMarshallers(parents, children);
}

// determines the marshallers needed for this specific object.
// it does not look at the object graph or consider inheritance.
function determineMarshallersForObj(obj: ObjectSchema): Marshallers {
  // things that require custom marshalling and/or unmarshalling:
  //   needsDateTimeMarshalling M, U
  //   needsDateMarshalling     M, U
  //   hasAdditionalProperties  M, U
  //   hasPolymorphicField      M, U
  //   discriminatorValue       M, U
  //   hasArrayMap              M
  //   needsPatchMarshaller     M

  let needsM = false, needsU = false;
  if (obj.language.go!.needsDateTimeMarshalling ||
    obj.language.go!.needsDateMarshalling ||
    hasAdditionalProperties(obj) ||
    hasPolymorphicField(obj) ||
    obj.discriminatorValue) {
    needsM = needsU = true;
  } else if (obj.language.go!.hasArrayMap ||
    obj.language.go!.needsPatchMarshaller) {
    needsM = true;
  }
  return {
    intM: false,
    intU: false,
    M: needsM,
    U: needsU,
  }
}

// walks the inheritance graph of obj to determine marshallers.
// when parents is true, the parents are walked, else the children.
function recursiveDetermineMarshallers(obj: ObjectSchema, parents: boolean): Marshallers {
  let cs: ComplexSchema[] | undefined;
  if (parents) {
    cs = obj.parents?.immediate;
  } else {
    cs = obj.children?.immediate;
  }

  // first check ourselves
  let result = determineMarshallersForObj(obj);

  // we must include any siblings in the calculation.  if one sibling
  // requires custom marshallers then they all need them.  consider the
  // following example (taken from body-complex.json)
  //   Pet -> Cat
  //   Pet -> Dog
  // Cat requires a custom marshaller, so one will be added to Pet too.
  // Dog does not require a custom marshaller, however if we don't give
  // it one, it will inherit the one from Pet, causing Dog's fields to
  // be omitted from the payload.
  if (obj.parents) {
    let parent: ObjectSchema | undefined;
    for (const cs of values(<ComplexSchema[]>obj.parents.immediate)) {
      if (!isObjectSchema(cs)) {
        continue;
      }
      parent = cs;
      break;
    }
    for (const sibling of values(parent?.children?.immediate)) {
      if (!isObjectSchema(sibling)) {
        continue;
      }
      const sibres = determineMarshallersForObj(sibling);
      result = mergeMarshallers(result, sibres);
    }
  }

  // now check children/parents
  for (const c of values(cs)) {
    if (!isObjectSchema(c)) {
      continue;
    }
    // if we already know we need all kinds don't bother to keep walking the hierarchy
    if (!result.M || !result.U || !result.intM || !result.intU) {
      const other = recursiveDetermineMarshallers(c, parents);
      result = mergeMarshallers(result, other);
    }
  }

  // finally, understand our place in the hierarchy
  if (obj.children && obj.parents) {
    // parent needs both
    result.intM = result.M;
    result.intU = result.U;
  } else if (obj.children) {
    // root also needs both
    result.intM = result.M;
    result.intU = result.U;
  } else if (obj.parents) {
    // leaf requires no internal marshallers
    result.intM = result.intU = false;
  }
  // the root type doesn't get a marshaller as callers don't instantiate instances of it
  if (obj.language.go!.rootDiscriminator) {
    result.M = false;
  }

  return result;
}

function needsXMLDictionaryUnmarshalling(obj: ObjectSchema): boolean {
  for (const prop of values(obj.properties)) {
    if (prop.language.go!.needsXMLDictionaryUnmarshalling) {
      return true;
    }
  }
  return false;
}

function needsXMLArrayMarshalling(obj: ObjectSchema): boolean {
  for (const prop of values(obj.properties)) {
    if (prop.language.go!.needsXMLArrayMarshalling) {
      return true;
    }
  }
  return false;
}

function generateStruct(imports: ImportManager, lang: Language, props?: Property[]): StructDef {
  if (lang.responseType) {
    imports.add('net/http');
  }
  if (lang.isLRO) {
    imports.add('time');
    imports.add('context');
  }
  if (lang.needsDateTimeMarshalling) {
    imports.add('encoding/' + lang.marshallingFormat);
  }
  const st = new StructDef(lang, props);
  for (const prop of values(props)) {
    imports.addImportForSchemaType(prop.schema);
  }
  return st;
}

function generateParamGroupStruct(imports: ImportManager, lang: Language, params: Parameter[]): StructDef {
  const st = new StructDef(lang, undefined, params);
  for (const param of values(params)) {
    imports.addImportForSchemaType(param.schema);
  }
  return st;
}

function generateUnmarshallerForResponseEnvelope(structDef: StructDef) {
  // if the response envelope contains a discriminated type we need an unmarshaller
  let found = false;
  for (const prop of values(structDef.Properties)) {
    if (prop.isDiscriminator) {
      found = true;
      break;
    }
  }
  if (!found) {
    return;
  }
  const receiver = structDef.receiverName();
  let unmarshaller = `func (${receiver} *${structDef.Language.name}) UnmarshalJSON(data []byte) error {\n`;
  // add a custom unmarshaller to the response envelope
  // find the discriminated type field
  let field = '';
  let type = '';
  for (const prop of values(structDef.Properties)) {
    if (prop.isDiscriminator) {
      field = prop.language.go!.name;
      type = prop.schema.language.go!.discriminatorInterface;
      break;
    }
  }
  if (field === '' || type === '') {
    throw new Error(`failed to the discriminated type field for response envelope ${structDef.Language.name}`);
  }
  unmarshaller += `\tres, err := unmarshal${type}(data)\n`;
  unmarshaller += '\tif err != nil {\n';
  unmarshaller += '\t\treturn err\n';
  unmarshaller += '\t}\n';
  unmarshaller += `\t${receiver}.${field} = res\n`;
  unmarshaller += '\treturn nil\n';
  unmarshaller += '}\n\n';
  structDef.Methods.push({ name: 'UnmarshalJSON', desc: `UnmarshalJSON implements the json.Unmarshaller interface for type ${structDef.Language.name}.`, text: unmarshaller });
}

// generates discriminator marker method, internal marshaller and internal unmarshaller
function generateDiscriminatorMarkerMethod(obj: ObjectSchema, structDef: StructDef) {
  const typeName = obj.language.go!.name;
  const receiver = structDef.receiverName();
  const interfaceMethod = `Get${typeName}`;
  const method = `func (${receiver} *${typeName}) ${interfaceMethod}() *${typeName} { return ${receiver} }\n\n`;
  structDef.Methods.push({ name: interfaceMethod, desc: `${interfaceMethod} implements the ${obj.language.go!.discriminatorInterface} interface for type ${typeName}.`, text: method });
}

function generateInternalMarshaller(obj: ObjectSchema, structDef: StructDef, parentType?: ObjectSchema) {
  if (obj.language.go!.errorType || obj.language.go!.inheritedErrorType) {
    // errors don't need custom marshallers
    return;
  }
  const typeName = obj.language.go!.name;
  const receiver = structDef.receiverName();
  // marshalInternal doesn't have any params in the non-discriminated type inheritence case
  let paramType = '';
  let paramName = '';
  if (obj.discriminator) {
    paramType = ' ' + obj.discriminator.property.schema.language.go!.name;
    paramName = 'discValue';
  }
  let marshalInteral = `func (${receiver} ${typeName}) marshalInternal(${paramName}${paramType}) map[string]interface{} {\n`;
  if (parentType) {
    // if the parent isn't a discriminator it won't have a param
    let parentParam = '';
    if (parentType.discriminator) {
      parentParam = paramName;
    }
    marshalInteral += `\tobjectMap := ${receiver}.${parentType.language.go!.name}.marshalInternal(${parentParam})\n`;
  } else {
    marshalInteral += '\tobjectMap := make(map[string]interface{})\n';
  }
  for (const prop of values(structDef.Properties)) {
    if (prop.language.go!.isAdditionalProperties) {
      continue;
    }
    if (prop.isDiscriminator) {
      marshalInteral += `\t${receiver}.${prop.language.go!.name} = &${paramName}\n`;
      marshalInteral += `\tobjectMap["${prop.serializedName}"] = ${receiver}.${prop.language.go!.name}\n`;
    } else {
      let source = `${receiver}.${prop.language.go!.name}`;
      if (prop.schema.language.go!.internalTimeType) {
        source = `(*${prop.schema.language.go!.internalTimeType})(${receiver}.${prop.language.go!.name})`;
      }
      marshalInteral += `\tpopulate(objectMap, "${prop.serializedName}", ${source})\n`;
    }
  }
  if (hasAdditionalProperties(obj)) {
    marshalInteral += `\tif ${receiver}.AdditionalProperties != nil {\n`;
    marshalInteral += `\t\tfor key, val := range ${receiver}.AdditionalProperties {\n`;
    marshalInteral += '\t\t\tobjectMap[key] = val\n';
    marshalInteral += '\t\t}\n';;
    marshalInteral += '\t}\n';
  }
  marshalInteral += '\treturn objectMap\n';
  marshalInteral += '}\n\n';
  structDef.Methods.push({ name: 'marshalInternal', desc: '', text: marshalInteral });
}

function generateInternalUnmarshaller(obj: ObjectSchema, structDef: StructDef, parentType?: ObjectSchema) {
  const typeName = obj.language.go!.name;
  const receiver = structDef.receiverName();
  let unmarshalInternall = `func (${receiver} *${typeName}) unmarshalInternal(rawMsg map[string]json.RawMessage) error {\n`;
  unmarshalInternall += generateJSONUnmarshallerBody(obj, structDef, parentType);
  unmarshalInternall += '}\n\n';
  structDef.Methods.push({ name: 'unmarshalInternal', desc: '', text: unmarshalInternall });
}

function generateJSONMarshaller(imports: ImportManager, obj: ObjectSchema, structDef: StructDef, parentType?: ObjectSchema) {
  if (obj.language.go!.errorType || obj.language.go!.inheritedErrorType) {
    // errors don't need custom marshallers
    return;
  } else if (!obj.discriminatorValue && (!structDef.Properties || structDef.Properties.length === 0)) {
    // non-discriminated types without content don't need a custom marshaller.
    // there is a case in network where child is allOf base and child has no properties.
    return;
  }
  imports.add('encoding/json');
  const typeName = structDef.Language.name;
  const receiver = structDef.receiverName();
  let marshaller = `func (${receiver} ${typeName}) MarshalJSON() ([]byte, error) {\n`;
  if (obj.discriminator) {
    marshaller += `\tobjectMap := ${receiver}.marshalInternal(${obj.discriminatorValue})\n`;
  } else if (obj.children?.immediate && isObjectSchema(obj.children.immediate[0])) {
    // non-discriminated type inheritence case (no param)
    marshaller += `\tobjectMap := ${receiver}.marshalInternal()\n`;
  } else {
    if (parentType) {
      let internalParamName = '';
      if (obj.discriminatorValue) {
        internalParamName = obj.discriminatorValue;
      }
      marshaller += `\tobjectMap := ${receiver}.${parentType!.language.go!.name}.marshalInternal(${internalParamName})\n`;
    } else {
      marshaller += '\tobjectMap := make(map[string]interface{})\n';
    }
    for (const prop of values(structDef.Properties)) {
      if (prop.language.go!.isAdditionalProperties) {
        continue;
      }
      let source = `${receiver}.${prop.language.go!.name}`;
      if (prop.schema.language.go!.internalTimeType) {
        source = `(*${prop.schema.language.go!.internalTimeType})(${receiver}.${prop.language.go!.name})`;
      }
      marshaller += `\tpopulate(objectMap, "${prop.serializedName}", ${source})\n`;
    }
    if (hasAdditionalProperties(obj)) {
      marshaller += `\tif ${receiver}.AdditionalProperties != nil {\n`;
      marshaller += `\t\tfor key, val := range ${receiver}.AdditionalProperties {\n`;
      marshaller += '\t\t\tobjectMap[key] = val\n';
      marshaller += '\t\t}\n';;
      marshaller += '\t}\n';
    }
  }
  marshaller += '\treturn json.Marshal(objectMap)\n';
  marshaller += '}\n\n';
  structDef.Methods.push({ name: 'MarshalJSON', desc: `MarshalJSON implements the json.Marshaller interface for type ${typeName}.`, text: marshaller });
}

function generateJSONUnmarshaller(imports: ImportManager, obj: ObjectSchema, structDef: StructDef, parentType?: ObjectSchema) {
  // there's a corner-case where a derived type might not add any new fields (Cookiecuttershark).
  // in this case skip adding the unmarshaller as it's not necessary and doesn't compile.
  if (!structDef.Properties || structDef.Properties.length === 0) {
    return;
  }
  imports.add('encoding/json');
  const typeName = structDef.Language.name;
  const receiver = structDef.receiverName();
  let unmarshaller = `func (${receiver} *${typeName}) UnmarshalJSON(data []byte) error {\n`;
  unmarshaller += '\tvar rawMsg map[string]json.RawMessage\n';
  unmarshaller += '\tif err := json.Unmarshal(data, &rawMsg); err != nil {\n';
  unmarshaller += '\t\treturn err\n';
  unmarshaller += '\t}\n';
  // the raw field won't exist on parents of the errorType
  if (obj.language.go!.errorType || obj.language.go!.inheritedErrorType === 'child') {
    unmarshaller += `\t${receiver}.raw = string(data)\n`;
  }
  if (obj.discriminator || obj.children?.immediate && isObjectSchema(obj.children.immediate[0])) {
    unmarshaller += `\treturn ${receiver}.unmarshalInternal(rawMsg)\n`;
  } else {
    unmarshaller += generateJSONUnmarshallerBody(obj, structDef, parentType);
  }
  unmarshaller += '}\n\n';
  structDef.Methods.push({ name: 'UnmarshalJSON', desc: `UnmarshalJSON implements the json.Unmarshaller interface for type ${typeName}.`, text: unmarshaller });
}

function generateJSONUnmarshallerBody(obj: ObjectSchema, structDef: StructDef, parentType?: ObjectSchema): string {
  const receiver = structDef.receiverName();
  const addlProps = hasAdditionalProperties(obj);
  const emitAddlProps = function (tab: string, addlProps: DictionarySchema): string {
    let addlPropsText = `${tab}\t\tif ${receiver}.AdditionalProperties == nil {\n`;
    let ptr = '', ref = '';
    if (<boolean>addlProps.language.go!.elementIsPtr) {
      ptr = '*';
      ref = '&';
    }
    addlPropsText += `${tab}\t\t\t${receiver}.AdditionalProperties = map[string]${ptr}${addlProps.elementType.language.go!.name}{}\n`;
    addlPropsText += `${tab}\t\t}\n`;
    addlPropsText += `${tab}\t\tif val != nil {\n`;
    addlPropsText += `${tab}\t\t\tvar aux ${addlProps.elementType.language.go!.name}\n`;
    addlPropsText += `${tab}\t\t\terr = json.Unmarshal(val, &aux)\n`;
    addlPropsText += `${tab}\t\t\t${receiver}.AdditionalProperties[key] = ${ref}aux\n`;
    addlPropsText += `${tab}\t\t}\n`;
    addlPropsText += `${tab}\t\tdelete(rawMsg, key)\n`;
    return addlPropsText;
  }
  let unmarshalBody = '';
  // handle the case where the type in the hierarchy doesn't contain any fields.
  // e.g. parent->intermediate->child and intermediate has no fields
  if (addlProps || (structDef.Properties && structDef.Properties.length > 0)) {
    unmarshalBody = '\tfor key, val := range rawMsg {\n';
    unmarshalBody += '\t\tvar err error\n';
    unmarshalBody += '\t\tswitch key {\n';
    // unmarshal content for the current type
    for (const prop of values(structDef.Properties)) {
      if (prop.language.go!.isAdditionalProperties) {
        continue;
      }
      unmarshalBody += `\t\tcase "${prop.serializedName}":\n`;
      if (prop.schema.language.go!.discriminatorInterface) {
        unmarshalBody += `\t\t\t\t${receiver}.${prop.language.go!.name}, err = unmarshal${prop.schema.language.go!.discriminatorInterface}(val)\n`;
      } else if (isArraySchema(prop.schema) && prop.schema.elementType.language.go!.discriminatorInterface) {
        unmarshalBody += `\t\t\t\t${receiver}.${prop.language.go!.name}, err = unmarshal${prop.schema.elementType.language.go!.discriminatorInterface}Array(val)\n`;
      } else if (prop.schema.language.go!.internalTimeType) {
        unmarshalBody += `\t\t\t\tvar aux ${prop.schema.language.go!.internalTimeType}\n`;
        unmarshalBody += '\t\t\t\terr = unpopulate(val, &aux)\n';
        unmarshalBody += `\t\t\t\t${receiver}.${prop.language.go!.name} = (*time.Time)(&aux)\n`;
      } else {
        unmarshalBody += `\t\t\t\terr = unpopulate(val, &${receiver}.${prop.language.go!.name})\n`;
      }
      unmarshalBody += '\t\t\t\tdelete(rawMsg, key)\n';
    }
    // if there's no parent type it's safe to unmarshal additional properties right here
    if (addlProps && !parentType) {
      unmarshalBody += '\t\tdefault:\n';
      unmarshalBody += emitAddlProps('\t', addlProps);
    }
    unmarshalBody += '\t\t}\n';
    unmarshalBody += '\t\tif err != nil {\n';
    unmarshalBody += '\t\t\treturn err\n';
    unmarshalBody += '\t\t}\n';
    unmarshalBody += '\t}\n'; // end for key, val := range rawMsg
  }
  if (parentType) {
    if (!addlProps) {
      unmarshalBody += `\treturn ${receiver}.${parentType.language.go!.name}.unmarshalInternal(rawMsg)\n`;
    } else {
      // unmarshal parent content first
      unmarshalBody += `\tif err := ${receiver}.${parentType.language.go!.name}.unmarshalInternal(rawMsg); err != nil {\n`;
      unmarshalBody += '\t\treturn err\n';
      unmarshalBody += '\t}\n';
      // now unmarshal additional properties
      unmarshalBody += '\tfor key, val := range rawMsg {\n';
      unmarshalBody += '\tvar err error\n';
      unmarshalBody += emitAddlProps('', addlProps);
      unmarshalBody += '\t\tif err != nil {\n';
      unmarshalBody += '\t\t\treturn err\n';
      unmarshalBody += '\t\t}\n';
      unmarshalBody += '\t}\n'; // end for key, val := range rawMsg
      unmarshalBody += '\treturn nil\n';
    }
  } else {
    // nothing left to unmarshal
    unmarshalBody += '\treturn nil\n';
  }
  return unmarshalBody;
}

function generateXMLMarshaller(structDef: StructDef) {
  // only needed for types with time.Time or where the XML name doesn't match the type name
  const receiver = structDef.receiverName();
  const desc = `MarshalXML implements the xml.Marshaller interface for type ${structDef.Language.name}.`;
  let text = `func (${receiver} ${structDef.Language.name}) MarshalXML(e *xml.Encoder, start xml.StartElement) error {\n`;
  if (structDef.Language.xmlWrapperName) {
    text += `\tstart.Name.Local = "${structDef.Language.xmlWrapperName}"\n`;
  }
  text += generateAliasType(structDef, receiver, true);
  // check for fields that require array marshalling
  const arrays = new Array<Property>();
  for (const prop of values(structDef.Properties)) {
    if (prop.language.go!.needsXMLArrayMarshalling) {
      arrays.push(prop);
    }
  }
  for (const array of values(arrays)) {
    text += `\tif ${receiver}.${array.language.go!.name} != nil {\n`;
    text += `\t\taux.${array.language.go!.name} = &${receiver}.${array.language.go!.name}\n`;
    text += '\t}\n';
  }
  text += '\treturn e.EncodeElement(aux, start)\n';
  text += '}\n\n';
  structDef.Methods.push({ name: 'MarshalXML', desc: desc, text: text });
}

function generateXMLUnmarshaller(structDef: StructDef) {
  // non-polymorphic case, must be something with time.Time
  const receiver = structDef.receiverName();
  const desc = `UnmarshalXML implements the xml.Unmarshaller interface for type ${structDef.Language.name}.`;
  let text = `func (${receiver} *${structDef.Language.name}) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {\n`;
  text += generateAliasType(structDef, receiver, false);
  text += '\tif err := d.DecodeElement(aux, &start); err != nil {\n';
  text += '\t\treturn err\n';
  text += '\t}\n';
  for (const prop of values(structDef.Properties)) {
    if (prop.schema.type === SchemaType.DateTime) {
      text += `\t${receiver}.${prop.language.go!.name} = (*time.Time)(aux.${prop.language.go!.name})\n`;
    } else if (prop.language.go!.isAdditionalProperties || prop.language.go!.needsXMLDictionaryUnmarshalling) {
      text += `\t${receiver}.${prop.language.go!.name} = (map[string]*string)(aux.${prop.language.go!.name})\n`;
    }
  }
  text += '\treturn nil\n';
  text += '}\n\n';
  structDef.Methods.push({ name: 'UnmarshalXML', desc: desc, text: text });
}

// generates an alias type used by custom XML marshaller/unmarshaller
function generateAliasType(structDef: StructDef, receiver: string, forMarshal: boolean): string {
  let text = `\ttype alias ${structDef.Language.name}\n`;
  text += `\taux := &struct {\n`;
  text += `\t\t*alias\n`;
  for (const prop of values(structDef.Properties)) {
    let sn = getXMLSerialization(prop, structDef.Language);
    if (prop.schema.type === SchemaType.DateTime) {
      text += `\t\t${prop.language.go!.name} *${prop.schema.language.go!.internalTimeType} \`${structDef.Language.marshallingFormat}:"${sn}"\`\n`;
    } else if (prop.language.go!.isAdditionalProperties || prop.language.go!.needsXMLDictionaryUnmarshalling) {
      text += `\t\t${prop.language.go!.name} additionalProperties \`${structDef.Language.marshallingFormat}:"${sn}"\`\n`;
    } else if (prop.language.go!.needsXMLArrayMarshalling) {
      text += `\t\t${prop.language.go!.name} *${prop.schema.language.go!.name} \`${structDef.Language.marshallingFormat}:"${sn}"\`\n`;
    }
  }
  text += `\t}{\n`;
  let rec = receiver;
  if (forMarshal) {
    rec = '&' + rec;
  }
  text += `\t\talias: (*alias)(${rec}),\n`;
  if (forMarshal) {
    // emit code to initialize time fields
    for (const prop of values(structDef.Properties)) {
      if (prop.schema.type !== SchemaType.DateTime) {
        continue;
      }
      text += `\t\t${prop.language.go!.name}: (*${prop.schema.language.go!.internalTimeType})(${receiver}.${prop.language.go!.name}),\n`;
    }
  }
  text += `\t}\n`;
  return text;
}
