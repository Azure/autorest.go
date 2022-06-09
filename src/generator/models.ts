/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { capitalize, comment } from '@azure-tools/codegen';
import { ByteArraySchema, CodeModel, ConstantSchema, DictionarySchema, GroupProperty, ObjectSchema, Language, SchemaType, Parameter, Property } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { formatConstantValue, isArraySchema, isDictionarySchema, isObjectSchema, commentLength } from '../common/helpers';
import { contentPreamble, sortAscending } from './helpers';
import { ImportManager } from './imports';
import { generateStruct, getXMLSerialization, StructDef, StructMethod } from './structs';

export interface modelsSerDe {
  models: string;
  serDe: string;
}

// Creates the content in models.go
export async function generateModels(session: Session<CodeModel>): Promise<modelsSerDe> {
  // this list of packages to import
  const modelImports = new ImportManager();
  const serdeImports = new ImportManager();
  let modelText = await contentPreamble(session);

  // we do model generation first as it can add imports to the imports list
  const structs = generateStructs(modelImports, serdeImports, session.model.schemas.objects);
  const paramGroups = <Array<GroupProperty>>session.model.language.go!.parameterGroups;
  for (const paramGroup of values(paramGroups)) {
    structs.push(generateParamGroupStruct(modelImports, paramGroup.schema.language.go!, paramGroup.originalParameter));
  }

  modelText += modelImports.text();

  // structs
  let needsJSONPopulate = false;
  let needsJSONUnpopulate = false;
  let needsJSONPopulateByteArray = false;
  let serdeTextBody = '';
  structs.sort((a: StructDef, b: StructDef) => { return sortAscending(a.Language.name, b.Language.name) });
  for (const struct of values(structs)) {
    modelText += struct.discriminator();
    modelText += struct.text();

    struct.Methods.sort((a: StructMethod, b: StructMethod) => { return sortAscending(a.name, b.name) });
    for (const method of values(struct.Methods)) {
      if (method.desc.length > 0) {
        modelText += `${comment(method.desc, '// ', undefined, commentLength)}\n`;
      }
      modelText += method.text;
    }

    struct.SerDeMethods.sort((a: StructMethod, b: StructMethod) => { return sortAscending(a.name, b.name) });
    for (const method of values(struct.SerDeMethods)) {
      if (method.desc.length > 0) {
        serdeTextBody += `${comment(method.desc, '// ', undefined, commentLength)}\n`;
      }
      serdeTextBody += method.text;
    }
    if (struct.SerDeMethods.length > 0) {
      needsJSONPopulate = true;
      needsJSONUnpopulate = true;
    }
    if (struct.HasJSONByteArray) {
      needsJSONPopulateByteArray = true;
    }
  }
  if (needsJSONPopulate) {
    serdeTextBody += 'func populate(m map[string]interface{}, k string, v interface{}) {\n';
    serdeTextBody += '\tif v == nil {\n';
    serdeTextBody += '\t\treturn\n';
    serdeTextBody += '\t} else if azcore.IsNullValue(v) {\n';
    serdeTextBody += '\t\tm[k] = nil\n';
    serdeTextBody += '\t} else if !reflect.ValueOf(v).IsNil() {\n';
    serdeTextBody += '\t\tm[k] = v\n';
    serdeTextBody += '\t}\n';
    serdeTextBody += '}\n\n';
  }
  if (needsJSONPopulateByteArray) {
    serdeImports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
    serdeTextBody += 'func populateByteArray(m map[string]interface{}, k string, b []byte, f runtime.Base64Encoding) {\n';
    serdeTextBody += '\tif azcore.IsNullValue(b) {\n';
    serdeTextBody += '\t\tm[k] = nil\n';
    serdeTextBody += '\t} else if len(b) == 0 {\n';
    serdeTextBody += '\t\treturn\n';
    serdeTextBody += '\t} else {\n';
    serdeTextBody += '\t\tm[k] = runtime.EncodeByteArray(b, f)\n';
    serdeTextBody += '\t}\n';
    serdeTextBody += '}\n\n';
  }
  if (needsJSONUnpopulate) {
    serdeImports.add('fmt');
    serdeTextBody += 'func unpopulate(data json.RawMessage, fn string, v interface{}) error {\n';
    serdeTextBody += '\tif data == nil {\n';
    serdeTextBody += '\t\treturn nil\n';
    serdeTextBody += '\t}\n';
    serdeTextBody += '\tif err := json.Unmarshal(data, v); err != nil {\n';
    serdeTextBody += '\t\treturn fmt.Errorf("struct field %s: %v", fn, err)\n';
    serdeTextBody += '\t}\n';
    serdeTextBody += '\treturn nil\n';
    serdeTextBody += '}\n\n';
  }
  let serdeText = '';
  if (serdeTextBody.length > 0) {
    serdeText = await contentPreamble(session);
    serdeText += serdeImports.text();
    serdeText += serdeTextBody;
  }
  return {
    models: modelText,
    serDe: serdeText
  };
}

function generateStructs(modelImports: ImportManager, serdeImports: ImportManager, objects?: ObjectSchema[]): StructDef[] {
  const structTypes = new Array<StructDef>();
  for (const obj of values(objects)) {
    if (obj.language.go!.omitType) {
      continue;
    }
    const structDef = generateStruct(modelImports, obj.language.go!, aggregateProperties(obj));
    if (obj.language.go!.marshallingFormat === 'xml') {
      serdeImports.add('encoding/xml');
      if (obj.language.go!.needsDateTimeMarshalling) {
        serdeImports.add('time');
      }
      // due to differences in XML marshallers/unmarshallers, we use different codegen than for JSON
      if (obj.language.go!.needsDateTimeMarshalling || obj.language.go!.xmlWrapperName || needsXMLArrayMarshalling(obj) || obj.language.go!.byteArrayFormat) {
        generateXMLMarshaller(structDef, serdeImports);
        if (obj.language.go!.needsDateTimeMarshalling || obj.language.go!.byteArrayFormat) {
          generateXMLUnmarshaller(structDef, serdeImports);
        }
      } else if (needsXMLDictionaryUnmarshalling(obj)) {
        generateXMLUnmarshaller(structDef, serdeImports);
      }
      structTypes.push(structDef);
      continue;
    }
    if (obj.discriminator) {
      generateDiscriminatorMarkerMethod(obj, structDef);
    }
    for (const parent of values(obj.parents?.all)) {
      if (isObjectSchema(parent) && parent.discriminator) {
        generateDiscriminatorMarkerMethod(parent, structDef);
      }
    }
    serdeImports.add('reflect');
    serdeImports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    if (obj.language.go!.byteArrayFormat) {
      structDef.HasJSONByteArray = true;
    }
    generateJSONMarshaller(serdeImports, obj, structDef);
    generateJSONUnmarshaller(serdeImports, structDef);
    structTypes.push(structDef);
  }
  return structTypes;
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

// aggregate the properties from the provided type and its parent types
function aggregateProperties(obj: ObjectSchema): Array<Property> {
  const allProps = new Array<Property>();
  for (const prop of values(obj.properties)) {
    allProps.push(prop);
  }
  for (const parent of values(obj.parents?.all)) {
    if (isObjectSchema(parent)) {
      for (const parentProp of values(parent.properties)) {
        // ensure that the parent doesn't contain any properties with the same name but different type
        const exists = values(allProps).where(p => { return p.language.go!.name === parentProp.language.go!.name}).first();
        if (exists) {
          if (exists.schema.language.go!.name !== parentProp.schema.language.go!.name) {
            const msg = `type ${obj.language.go!.name} contains duplicate property ${exists.language.go!.name} with mismatched types`;
            throw new Error(msg);
          }
          // don't add the duplicate
          continue;
        }
        allProps.push(parentProp);
      }
    }
  }
  return allProps;
}

// generates discriminator marker method
function generateDiscriminatorMarkerMethod(obj: ObjectSchema, structDef: StructDef) {
  const typeName = obj.language.go!.name;
  const receiver = structDef.receiverName();
  const interfaceMethod = `Get${typeName}`;
  let method = `func (${receiver} *${structDef.Language.name}) ${interfaceMethod}() *${typeName} {`;
  if (obj.language.go!.name === structDef.Language.name) {
    // the marker method is on the discriminator itself, so just return the receiver
    method += ` return ${receiver} }\n\n`;
  } else {
    // the marker method is on a child type, so return an instance of the parent
    // type by copying the parent values into a new instance.
    method += `\n\treturn &${obj.language.go!.name}{\n`;
    for (const prop of values(aggregateProperties(obj))) {
      method += `\t\t${prop.language.go!.name}: ${structDef.receiverName()}.${prop.language.go!.name},\n`;
    }
    method += '\t}\n}\n\n';
  }
  structDef.Methods.push({ name: interfaceMethod, desc: `${interfaceMethod} implements the ${obj.language.go!.discriminatorInterface} interface for type ${structDef.Language.name}.`, text: method });
}

function generateJSONMarshaller(imports: ImportManager, obj: ObjectSchema, structDef: StructDef) {
if (!obj.discriminatorValue && (!structDef.Properties || structDef.Properties.length === 0)) {
    // non-discriminated types without content don't need a custom marshaller.
    // there is a case in network where child is allOf base and child has no properties.
    return;
  }
  imports.add('encoding/json');
  const typeName = structDef.Language.name;
  const receiver = structDef.receiverName();
  let marshaller = `func (${receiver} ${typeName}) MarshalJSON() ([]byte, error) {\n`;
  marshaller += '\tobjectMap := make(map[string]interface{})\n';
  marshaller += generateJSONMarshallerBody(obj, structDef, imports);
  marshaller += '\treturn json.Marshal(objectMap)\n';
  marshaller += '}\n\n';
  structDef.SerDeMethods.push({ name: 'MarshalJSON', desc: `MarshalJSON implements the json.Marshaller interface for type ${typeName}.`, text: marshaller });
}

function generateJSONMarshallerBody(obj: ObjectSchema, structDef: StructDef, imports: ImportManager): string {
  const receiver = structDef.receiverName();
  let marshaller = '';
  let addlProps: DictionarySchema | undefined;
  for (const prop of values(structDef.Properties)) {
    if (prop.language.go!.isAdditionalProperties) {
      addlProps = <DictionarySchema>prop.schema;
      continue;
    }
    if (prop.isDiscriminator) {
      if (obj.discriminatorValue) {
        marshaller += `\tobjectMap["${prop.serializedName}"] = ${obj.discriminatorValue}\n`;
      } else {
        // if there's no discriminator value (e.g. Fish in test server), use the field's value.
        // this will enable support for custom types that aren't (yet) described in the swagger.
        marshaller += `\tobjectMap["${prop.serializedName}"] = ${receiver}.${prop.language.go!.name}\n`;
      }
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
    } else if (prop.schema.type === SchemaType.Constant) {
      marshaller += `\tobjectMap["${prop.serializedName}"] = ${formatConstantValue(<ConstantSchema>prop.schema)}\n`;
    } else {
      let populate = 'populate';
      let addr = '';
      if (prop.schema.language.go!.internalTimeType) {
        populate += capitalize(prop.schema.language.go!.internalTimeType);
      } else if (prop.schema.type === SchemaType.Any) {
        // for fields that are interface{} we pass their address so populate() IsNil() doesn't panic
        addr = '&';
      }
      marshaller += `\t${populate}(objectMap, "${prop.serializedName}", ${addr}${receiver}.${prop.language.go!.name})\n`;
    }
  }
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

function generateJSONUnmarshaller(imports: ImportManager, structDef: StructDef) {
  // there's a corner-case where a derived type might not add any new fields (Cookiecuttershark).
  // in this case skip adding the unmarshaller as it's not necessary and doesn't compile.
  if (!structDef.Properties || structDef.Properties.length === 0) {
    return;
  }
  imports.add('encoding/json');
  imports.add('fmt');
  const typeName = structDef.Language.name;
  const receiver = structDef.receiverName();
  let unmarshaller = `func (${receiver} *${typeName}) UnmarshalJSON(data []byte) error {\n`;
  unmarshaller += '\tvar rawMsg map[string]json.RawMessage\n';
  unmarshaller += '\tif err := json.Unmarshal(data, &rawMsg); err != nil {\n';
  unmarshaller += `\t\treturn fmt.Errorf("unmarshalling type %T: %v", ${receiver}, err)\n`;
  unmarshaller += '\t}\n';
  unmarshaller += generateJSONUnmarshallerBody(structDef, imports);
  unmarshaller += '}\n\n';
  structDef.SerDeMethods.push({ name: 'UnmarshalJSON', desc: `UnmarshalJSON implements the json.Unmarshaller interface for type ${typeName}.`, text: unmarshaller });
}

function generateJSONUnmarshallerBody(structDef: StructDef, imports: ImportManager): string {
  const receiver = structDef.receiverName();
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
      imports.add('time');
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
  let unmarshalBody = '';
  unmarshalBody = '\tfor key, val := range rawMsg {\n';
  unmarshalBody += '\t\tvar err error\n';
  unmarshalBody += '\t\tswitch key {\n';
  let addlProps: DictionarySchema | undefined;
  for (const prop of values(structDef.Properties)) {
    if (prop.language.go!.isAdditionalProperties) {
      addlProps = <DictionarySchema>prop.schema;
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
      unmarshalBody += `\t\t\t\terr = unpopulate${capitalize(prop.schema.language.go!.internalTimeType)}(val, "${prop.language.go!.name}", &${receiver}.${prop.language.go!.name})\n`;
    } else if (isArraySchema(prop.schema) && prop.schema.elementType.language.go!.internalTimeType) {
      imports.add('time');
      unmarshalBody += `\t\t\tvar aux []*${prop.schema.elementType.language.go!.internalTimeType}\n`;
      unmarshalBody += `\t\t\terr = unpopulate(val, "${prop.language.go!.name}", &aux)\n`;
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
      unmarshalBody += `\t\t\t\terr = unpopulate(val, "${prop.language.go!.name}", &${receiver}.${prop.language.go!.name})\n`;
    }
    unmarshalBody += '\t\t\t\tdelete(rawMsg, key)\n';
  }
  if (addlProps) {
    unmarshalBody += '\t\tdefault:\n';
    unmarshalBody += emitAddlProps('\t', addlProps);
  }
  unmarshalBody += '\t\t}\n';
  unmarshalBody += '\t\tif err != nil {\n';
  unmarshalBody += `\t\t\treturn fmt.Errorf("unmarshalling type %T: %v", ${receiver}, err)\n`;
  unmarshalBody += '\t\t}\n';
  unmarshalBody += '\t}\n'; // end for key, val := range rawMsg
  unmarshalBody += '\treturn nil\n';
  return unmarshalBody;
}

function generateXMLMarshaller(structDef: StructDef, imports: ImportManager) {
  // only needed for types with time.Time or where the XML name doesn't match the type name
  const receiver = structDef.receiverName();
  const desc = `MarshalXML implements the xml.Marshaller interface for type ${structDef.Language.name}.`;
  let text = `func (${receiver} ${structDef.Language.name}) MarshalXML(e *xml.Encoder, start xml.StartElement) error {\n`;
  if (structDef.Language.xmlWrapperName) {
    text += `\tstart.Name.Local = "${structDef.Language.xmlWrapperName}"\n`;
  }
  text += generateAliasType(structDef, receiver, true);
  for (const prop of values(structDef.Properties)) {
    if (prop.language.go!.needsXMLArrayMarshalling) {
      text += `\tif ${receiver}.${prop.language.go!.name} != nil {\n`;
      text += `\t\taux.${prop.language.go!.name} = &${receiver}.${prop.language.go!.name}\n`;
      text += '\t}\n';
    } else if (prop.schema.type === SchemaType.ByteArray) {
      let base64Format = 'Std';
      if ((<ByteArraySchema>prop.schema).format === 'base64url') {
        base64Format = 'URL';
      }
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      text += `\tif ${receiver}.${prop.language.go!.name} != nil {\n`
      text += `\t\tencoded${prop.language.go!.name} := runtime.EncodeByteArray(${receiver}.${prop.language.go!.name}, runtime.Base64${base64Format}Format)\n`;
      text += `\t\taux.${prop.language.go!.name} = &encoded${prop.language.go!.name}\n`;
      text += '\t}\n';
    }
  }
  text += '\treturn e.EncodeElement(aux, start)\n';
  text += '}\n\n';
  structDef.SerDeMethods.push({ name: 'MarshalXML', desc: desc, text: text });
}

function generateXMLUnmarshaller(structDef: StructDef, imports: ImportManager) {
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
    } else if (prop.schema.type === SchemaType.ByteArray) {
      let base64Format = 'Std';
      if ((<ByteArraySchema>prop.schema).format === 'base64url') {
        base64Format = 'URL';
      }
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      text += `\tif aux.${prop.language.go!.name} != nil {\n`
      text += `\t\tif err := runtime.DecodeByteArray(*aux.${prop.language.go!.name}, &${receiver}.${prop.language.go!.name}, runtime.Base64${base64Format}Format); err != nil {\n`;
      text += '\t\t\treturn err\n';
      text += '\t\t}\n';
      text += '\t}\n';
    }
  }
  text += '\treturn nil\n';
  text += '}\n\n';
  structDef.SerDeMethods.push({ name: 'UnmarshalXML', desc: desc, text: text });
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
    } else if (prop.schema.type === SchemaType.ByteArray) {
      text += `\t\t${prop.language.go!.name} *string \`${structDef.Language.marshallingFormat}:"${sn}"\`\n`;
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
