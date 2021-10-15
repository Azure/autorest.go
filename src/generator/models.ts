/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { comment } from '@azure-tools/codegen';
import { ByteArraySchema, CodeModel, ComplexSchema, DictionarySchema, GroupProperty, ObjectSchema, Language, SchemaType, Parameter, Property } from '@autorest/codemodel';
import { length, values } from '@azure-tools/linq';
import { isArraySchema, isObjectSchema, hasAdditionalProperties, hasPolymorphicField, commentLength } from '../common/helpers';
import { contentPreamble, sortAscending } from './helpers';
import { ImportManager } from './imports';
import { generateStruct, getXMLSerialization, StructDef, StructMethod } from './structs';

// Creates the content in models.go
export async function generateModels(session: Session<CodeModel>): Promise<string> {
  // this list of packages to import
  const imports = new ImportManager();
  let text = await contentPreamble(session);

  // we do model generation first as it can add imports to the imports list
  const structs = generateStructs(imports, session.model.schemas.objects);
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
  let needsJSONPopulateByteArray = false;
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
    if (struct.HasJSONByteArray) {
      needsJSONPopulateByteArray = true;
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
  if (needsJSONPopulateByteArray) {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
    text += 'func populateByteArray(m map[string]interface{}, k string, b []byte, f runtime.Base64Encoding) {\n';
    text += '\tif azcore.IsNullValue(b) {\n';
    text += '\t\tm[k] = nil\n';
    text += '\t} else if len(b) == 0 {\n';
    text += '\t\treturn\n';
    text += '\t} else {\n';
    text += '\t\tm[k] = runtime.EncodeByteArray(b, f)\n';
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

function generateStructs(imports: ImportManager, objects?: ObjectSchema[]): StructDef[] {
  const structTypes = new Array<StructDef>();
  for (const obj of values(objects)) {
    const structDef = generateStruct(imports, obj.language.go!, obj.properties);
    // now add the parent type
    for (const parent of values(obj.parents?.immediate)) {
      if (isObjectSchema(parent)) {
        structDef.ComposedOf.push(parent.language.go!.name);
      }
    }
    structDef.ComposedOf.sort();
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
      generateInternalMarshaller(obj, structDef, imports);
    }
    if (needs.intU) {
      generateInternalUnmarshaller(obj, structDef, imports);
    }
    if (needs.M) {
      imports.add('reflect');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
      structDef.HasJSONMarshaller = true;
      if (obj.language.go!.byteArrayFormat) {
        structDef.HasJSONByteArray = true;
      }
      generateJSONMarshaller(imports, obj, structDef);
    }
    if (needs.U) {
      structDef.HasJSONUnmarshaller = true;
      generateJSONUnmarshaller(imports, obj, structDef);
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

  return recursiveDetermineMarshallers(obj, new Array<string>());
}

// determines the marshallers needed for this specific object.
// it does not look at the object graph or consider inheritance.
function determineMarshallersForObj(obj: ObjectSchema): Marshallers {
  // things that require custom marshalling and/or unmarshalling:
  //   needsDateTimeMarshalling M, U
  //   needsDateMarshalling     M, U
  //   needsUnixTimeMarshalling M, U
  //   hasAdditionalProperties  M, U
  //   hasPolymorphicField      M, U
  //   discriminatorValue       M, U
  //   byteArrayFormat          M, U
  //   hasArrayMap              M
  //   needsPatchMarshaller     M

  let needsM = false, needsU = false;
  if (obj.language.go!.needsDateTimeMarshalling ||
    obj.language.go!.needsDateMarshalling ||
    obj.language.go!.needsUnixTimeMarshalling ||
    hasAdditionalProperties(obj) ||
    hasPolymorphicField(obj) ||
    obj.discriminatorValue ||
    obj.language.go!.byteArrayFormat) {
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
// visited tracks the name of the nodes that have been visited.
function recursiveDetermineMarshallers(obj: ObjectSchema, visited: Array<string>): Marshallers {
  // first check ourselves
  let result = determineMarshallersForObj(obj);
  visited.push(obj.language.go!.name);

  const visit = function (immediate: ComplexSchema[] | undefined) {
    for (const c of values(immediate)) {
      if (!isObjectSchema(c)) {
        continue;
      }
      if (visited.includes(c.language.go!.name)) {
        // prevent infinite recursion
        continue;
      }
      // if we already know we need all kinds don't bother to keep walking the hierarchy
      if (!result.M || !result.U || !result.intM || !result.intU) {
        const other = recursiveDetermineMarshallers(c, visited);
        result = mergeMarshallers(result, other);
      }
    }
  }

  // now check children/parents
  visit(obj.parents?.immediate);
  visit(obj.children?.immediate);

  // finally, understand our place in the hierarchy
  if (obj.children && obj.parents) {
    // parent needs both
    result.intM = result.M;
    result.intU = result.U;
  } else if (obj.children) {
    // root also needs both
    if (!result.intM) {
      // there's a corner-case here where result.M will be false for
      // the root discriminator, however we already know we need an
      // internal marshaller.  don't stomp over result.intM when true.
      result.intM = result.M;
    }
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

function generateParamGroupStruct(imports: ImportManager, lang: Language, params: Parameter[]): StructDef {
  const st = new StructDef(lang, undefined, params);
  for (const param of values(params)) {
    imports.addImportForSchemaType(param.schema);
  }
  return st;
}

// generates discriminator marker method, internal marshaller and internal unmarshaller
function generateDiscriminatorMarkerMethod(obj: ObjectSchema, structDef: StructDef) {
  const typeName = obj.language.go!.name;
  const receiver = structDef.receiverName();
  const interfaceMethod = `Get${typeName}`;
  const method = `func (${receiver} *${typeName}) ${interfaceMethod}() *${typeName} { return ${receiver} }\n\n`;
  structDef.Methods.push({ name: interfaceMethod, desc: `${interfaceMethod} implements the ${obj.language.go!.discriminatorInterface} interface for type ${typeName}.`, text: method });
}

function generateInternalMarshaller(obj: ObjectSchema, structDef: StructDef, imports: ImportManager) {
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
    paramName = ', discValue';
  }
  let marshalInternal = `func (${receiver} ${typeName}) marshalInternal(objectMap map[string]interface{}${paramName}${paramType}) {\n`;
  for (const parent of values(obj.parents?.immediate)) {
    let parentParam = '';
    if (isObjectSchema(parent)) {
      // if we have a discriminator value and the parent is a discriminator then we
      // know we need to call the marshalInternal() method with the discriminator value
      if (obj.discriminatorValue && parent.discriminator) {
        parentParam = paramName;
      }
      marshalInternal += `\t${receiver}.${parent.language.go!.name}.marshalInternal(objectMap${parentParam})\n`;
    }
  }
  marshalInternal += generateJSONMarshallerBody(obj, structDef, imports);
  marshalInternal += '}\n\n';
  structDef.Methods.push({ name: 'marshalInternal', desc: '', text: marshalInternal });
}

function generateInternalUnmarshaller(obj: ObjectSchema, structDef: StructDef, imports: ImportManager) {
  const typeName = obj.language.go!.name;
  const receiver = structDef.receiverName();
  let unmarshalInternall = `func (${receiver} *${typeName}) unmarshalInternal(rawMsg map[string]json.RawMessage) error {\n`;
  unmarshalInternall += generateJSONUnmarshallerBody(obj, structDef, imports);
  unmarshalInternall += '}\n\n';
  structDef.Methods.push({ name: 'unmarshalInternal', desc: '', text: unmarshalInternall });
}

function generateJSONMarshaller(imports: ImportManager, obj: ObjectSchema, structDef: StructDef) {
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
  marshaller += '\tobjectMap := make(map[string]interface{})\n';
  if (obj.discriminator) {
    marshaller += `\t${receiver}.marshalInternal(objectMap, ${obj.discriminatorValue})\n`;
  } else if (obj.children?.immediate && isObjectSchema(obj.children.immediate[0])) {
    // non-discriminated type inheritence case (no param)
    marshaller += `\t${receiver}.marshalInternal(objectMap)\n`;
  } else {
    for (const parent of values(obj.parents?.immediate)) {
      if (isObjectSchema(parent)) {
        let internalParamName = '';
        // if we have a discriminator value and the parent is a discriminator then we
        // know we need to call the marshalInternal() method with the discriminator value
        if (obj.discriminatorValue && parent.discriminator) {
          internalParamName = ', ' + obj.discriminatorValue;
        }
        marshaller += `\t${receiver}.${parent.language.go!.name}.marshalInternal(objectMap${internalParamName})\n`;
      }
    }
    marshaller += generateJSONMarshallerBody(obj, structDef, imports);
  }
  marshaller += '\treturn json.Marshal(objectMap)\n';
  marshaller += '}\n\n';
  structDef.Methods.push({ name: 'MarshalJSON', desc: `MarshalJSON implements the json.Marshaller interface for type ${typeName}.`, text: marshaller });
}

function generateJSONMarshallerBody(obj: ObjectSchema, structDef: StructDef, imports: ImportManager): string {
  const receiver = structDef.receiverName();
  let marshaller = '';
  for (const prop of values(structDef.Properties)) {
    if (prop.language.go!.isAdditionalProperties) {
      continue;
    }
    if (prop.isDiscriminator) {
      marshaller += `\t${receiver}.${prop.language.go!.name} = &discValue\n`;
      marshaller += `\tobjectMap["${prop.serializedName}"] = ${receiver}.${prop.language.go!.name}\n`;
    } else if (prop.schema.type === SchemaType.ByteArray) {
      let base64Format = 'Std';
      if ((<ByteArraySchema>prop.schema).format === 'base64url') {
        base64Format = 'URL';
      }
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      marshaller += `\tpopulateByteArray(objectMap, "${prop.serializedName}", ${receiver}.${prop.language.go!.name}, runtime.Base64${base64Format}Format)\n`;
    } else if (isArraySchema(prop.schema) && prop.schema.elementType.language.go!.internalTimeType) {
      const source = `${receiver}.${prop.language.go!.name}`;
      marshaller += `\taux := make([]*${prop.schema.elementType.language.go!.internalTimeType}, len(${source}), len(${source}))\n`;
      marshaller += `\tfor i := 0; i < len(${source}); i++ {\n`;
      marshaller += `\t\taux[i] = (*${prop.schema.elementType.language.go!.internalTimeType})(${source}[i])\n`;
      marshaller += '\t}\n';
      marshaller += `\tpopulate(objectMap, "${prop.serializedName}", aux)\n`;
    } else {
      let source = `${receiver}.${prop.language.go!.name}`;
      if (prop.schema.language.go!.internalTimeType) {
        source = `(*${prop.schema.language.go!.internalTimeType})(${receiver}.${prop.language.go!.name})`;
      }
      marshaller += `\tpopulate(objectMap, "${prop.serializedName}", ${source})\n`;
    }
  }
  const addlProps = hasAdditionalProperties(obj);
  if (addlProps) {
    marshaller += `\tif ${receiver}.AdditionalProperties != nil {\n`;
    marshaller += `\t\tfor key, val := range ${receiver}.AdditionalProperties {\n`;
    let assignment = 'val';
    if (addlProps.elementType.language.go!.internalTimeType) {
      assignment = `(*${addlProps.elementType.language.go!.internalTimeType})(val)`;
    }
    marshaller += `\t\t\tobjectMap[key] = ${assignment}\n`;
    marshaller += '\t\t}\n';;
    marshaller += '\t}\n';
  }
  return marshaller;
}

function generateJSONUnmarshaller(imports: ImportManager, obj: ObjectSchema, structDef: StructDef) {
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
  // checking obj.discriminator isn't enough, also check that there are actual child types
  if ((obj.discriminator && length(obj.discriminator.all) > 0) || obj.children?.immediate && isObjectSchema(obj.children.immediate[0])) {
    unmarshaller += `\treturn ${receiver}.unmarshalInternal(rawMsg)\n`;
  } else {
    unmarshaller += generateJSONUnmarshallerBody(obj, structDef, imports);
  }
  unmarshaller += '}\n\n';
  structDef.Methods.push({ name: 'UnmarshalJSON', desc: `UnmarshalJSON implements the json.Unmarshaller interface for type ${typeName}.`, text: unmarshaller });
}

function generateJSONUnmarshallerBody(obj: ObjectSchema, structDef: StructDef, imports: ImportManager): string {
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
    let auxType = addlProps.elementType.language.go!.name;
    let assignment = `${ref}aux`;
    if (addlProps.elementType.language.go!.internalTimeType) {
      auxType = addlProps.elementType.language.go!.internalTimeType;
      assignment = `(*time.Time)(${assignment})`;
    }
    addlPropsText += `${tab}\t\t\tvar aux ${auxType}\n`;
    addlPropsText += `${tab}\t\t\terr = json.Unmarshal(val, &aux)\n`;
    addlPropsText += `${tab}\t\t\t${receiver}.AdditionalProperties[key] = ${assignment}\n`;
    addlPropsText += `${tab}\t\t}\n`;
    addlPropsText += `${tab}\t\tdelete(rawMsg, key)\n`;
    return addlPropsText;
  }
  let hasParentType = false;
  for (const parent of values(obj.parents?.immediate)) {
    if (isObjectSchema(parent)) {
      hasParentType = true;
      break;
    }
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
      } else if (isDictionarySchema(prop.schema) && prop.schema.elementType.language.go!.discriminatorInterface) {
        unmarshalBody += `\t\t\t\t${receiver}.${prop.language.go!.name}, err = unmarshal${prop.schema.elementType.language.go!.discriminatorInterface}Map(val)\n`;
      } else if (prop.schema.language.go!.internalTimeType) {
        unmarshalBody += `\t\t\t\tvar aux ${prop.schema.language.go!.internalTimeType}\n`;
        unmarshalBody += '\t\t\t\terr = unpopulate(val, &aux)\n';
        unmarshalBody += `\t\t\t\t${receiver}.${prop.language.go!.name} = (*time.Time)(&aux)\n`;
      } else if (isArraySchema(prop.schema) && prop.schema.elementType.language.go!.internalTimeType) {
        unmarshalBody += `\t\t\tvar aux []*${prop.schema.elementType.language.go!.internalTimeType}\n`;
        unmarshalBody += '\t\t\terr = unpopulate(val, &aux)\n';
        unmarshalBody += '\t\t\tfor _, au := range aux {\n';
        unmarshalBody += `\t\t\t\t${receiver}.${prop.language.go!.name} = append(${receiver}.${prop.language.go!.name}, (*time.Time)(au))\n`;
        unmarshalBody += '\t\t\t}\n';
      } else if (prop.schema.type === SchemaType.ByteArray) {
        let base64Format = 'Std';
        if ((<ByteArraySchema>prop.schema).format === 'base64url') {
          base64Format = 'URL';
        }
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
        unmarshalBody += `\t\t\terr = runtime.DecodeByteArray(string(val), &${receiver}.${prop.language.go!.name}, runtime.Base64${base64Format}Format)\n`;
      } else {
        unmarshalBody += `\t\t\t\terr = unpopulate(val, &${receiver}.${prop.language.go!.name})\n`;
      }
      unmarshalBody += '\t\t\t\tdelete(rawMsg, key)\n';
    }
    // if there's no parent type it's safe to unmarshal additional properties right here
    if (addlProps && !hasParentType) {
      unmarshalBody += '\t\tdefault:\n';
      unmarshalBody += emitAddlProps('\t', addlProps);
    }
    unmarshalBody += '\t\t}\n';
    unmarshalBody += '\t\tif err != nil {\n';
    unmarshalBody += '\t\t\treturn err\n';
    unmarshalBody += '\t\t}\n';
    unmarshalBody += '\t}\n'; // end for key, val := range rawMsg
  }
  if (hasParentType) {
    for (const parent of values(obj.parents?.immediate)) {
      if (isObjectSchema(parent)) {
        unmarshalBody += `\tif err := ${receiver}.${parent.language.go!.name}.unmarshalInternal(rawMsg); err != nil {\n`;
        unmarshalBody += '\t\treturn err\n';
        unmarshalBody += '\t}\n';
      }
    }
    if (addlProps) {
      // now unmarshal additional properties
      unmarshalBody += '\tfor key, val := range rawMsg {\n';
      unmarshalBody += '\tvar err error\n';
      unmarshalBody += emitAddlProps('', addlProps);
      unmarshalBody += '\t\tif err != nil {\n';
      unmarshalBody += '\t\t\treturn err\n';
      unmarshalBody += '\t\t}\n';
      unmarshalBody += '\t}\n'; // end for key, val := range rawMsg
    }
  }
  unmarshalBody += '\treturn nil\n';
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
