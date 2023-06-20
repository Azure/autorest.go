/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { capitalize, comment } from '@azure-tools/codegen';
import { ByteArraySchema, CodeModel, ConstantSchema, DictionarySchema, Language, ObjectSchema, Property, SchemaType } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { formatConstantValue, isArraySchema, isDictionarySchema, isObjectSchema, commentLength, aggregateProperties } from '../common/helpers';
import { contentPreamble, hasDescription, getClientDefaultValue, sortAscending } from './helpers';
import { ImportManager } from './imports';

export interface ModelsSerDe {
  models: string;
  serDe: string;
}

// Creates the content in models.go and models_serde.go
export async function generateModels(session: Session<CodeModel>): Promise<ModelsSerDe> {
  // this list of packages to import
  const modelImports = new ImportManager();
  const serdeImports = new ImportManager();
  let modelText = await contentPreamble(session);

  // we do model generation first as it can add imports to the imports list
  const modelDefs = generateModelDefs(modelImports, serdeImports, session.model.schemas.objects);

  modelText += modelImports.text();

  // structs
  let needsJSONPopulate = false;
  let needsJSONUnpopulate = false;
  let needsJSONPopulateByteArray = false;
  let needsJSONPopulateAny = false;
  let serdeTextBody = '';
  modelDefs.sort((a: ModelDef, b: ModelDef) => { return sortAscending(a.Language.name, b.Language.name); });
  for (const modelDef of values(modelDefs)) {
    modelText += modelDef.text();

    modelDef.Methods.sort((a: ModelMethod, b: ModelMethod) => { return sortAscending(a.name, b.name); });
    for (const method of values(modelDef.Methods)) {
      if (method.desc.length > 0) {
        modelText += `${comment(method.desc, '// ', undefined, commentLength)}\n`;
      }
      modelText += method.text;
    }

    modelDef.SerDeMethods.sort((a: ModelMethod, b: ModelMethod) => { return sortAscending(a.name, b.name); });
    for (const method of values(modelDef.SerDeMethods)) {
      if (method.desc.length > 0) {
        serdeTextBody += `${comment(method.desc, '// ', undefined, commentLength)}\n`;
      }
      serdeTextBody += method.text;
    }
    if (modelDef.SerDeMethods.length > 0) {
      needsJSONPopulate = true;
      needsJSONUnpopulate = true;
    }
    if (modelDef.HasJSONByteArray) {
      needsJSONPopulateByteArray = true;
    }
    if (modelDef.HasAny) {
      needsJSONPopulateAny = true;
    }
  }
  if (needsJSONPopulate) {
    serdeTextBody += 'func populate(m map[string]any, k string, v any) {\n';
    serdeTextBody += '\tif v == nil {\n';
    serdeTextBody += '\t\treturn\n';
    serdeTextBody += '\t} else if azcore.IsNullValue(v) {\n';
    serdeTextBody += '\t\tm[k] = nil\n';
    serdeTextBody += '\t} else if !reflect.ValueOf(v).IsNil() {\n';
    serdeTextBody += '\t\tm[k] = v\n';
    serdeTextBody += '\t}\n';
    serdeTextBody += '}\n\n';
  }
  if (needsJSONPopulateAny) {
    serdeTextBody += 'func populateAny(m map[string]any, k string, v any) {\n';
    serdeTextBody += '\tif v == nil {\n';
    serdeTextBody += '\t\treturn\n';
    serdeTextBody += '\t} else if azcore.IsNullValue(v) {\n';
    serdeTextBody += '\t\tm[k] = nil\n';
    serdeTextBody += '\t} else {\n';
    serdeTextBody += '\t\tm[k] = v\n';
    serdeTextBody += '\t}\n';
    serdeTextBody += '}\n\n';
  }
  if (needsJSONPopulateByteArray) {
    serdeImports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
    serdeTextBody += 'func populateByteArray(m map[string]any, k string, b []byte, f runtime.Base64Encoding) {\n';
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
    serdeTextBody += 'func unpopulate(data json.RawMessage, fn string, v any) error {\n';
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

function generateModelDefs(modelImports: ImportManager, serdeImports: ImportManager, objects?: Array<ObjectSchema>): Array<ModelDef> {
  const modelDefs = new Array<ModelDef>();
  for (const obj of values(objects)) {
    if (obj.language.go!.omitType || obj.extensions?.['x-ms-external']) {
      continue;
    }

    if (obj.language.go!.isLRO) {
      modelImports.add('time');
      modelImports.add('context');
    }
    const props = aggregateProperties(obj);
    const modelDef = new ModelDef(obj.language.go!, props);
    for (const prop of values(props)) {
      modelImports.addImportForSchemaType(prop.schema);
      if (prop.schema.type === SchemaType.Any && !prop.schema.language.go!.rawJSONAsBytes) {
        modelDef.HasAny = true;
      }
    }

    if (obj.language.go!.marshallingFormat === 'xml' && !obj.language.go!.omitSerDeMethods) {
      serdeImports.add('encoding/xml');
      if (obj.language.go!.needsDateTimeMarshalling) {
        serdeImports.add('time');
      }
      // due to differences in XML marshallers/unmarshallers, we use different codegen than for JSON
      if (obj.language.go!.needsDateTimeMarshalling || obj.language.go!.xmlWrapperName || needsXMLArrayMarshalling(obj) || obj.language.go!.byteArrayFormat) {
        generateXMLMarshaller(modelDef, serdeImports);
        if (obj.language.go!.needsDateTimeMarshalling || obj.language.go!.byteArrayFormat) {
          generateXMLUnmarshaller(modelDef, serdeImports);
        }
      } else if (needsXMLDictionaryUnmarshalling(obj)) {
        generateXMLUnmarshaller(modelDef, serdeImports);
      }
      modelDefs.push(modelDef);
      continue;
    }
    if (obj.discriminator) {
      generateDiscriminatorMarkerMethod(obj, modelDef);
    }
    for (const parent of values(obj.parents?.all)) {
      if (isObjectSchema(parent) && parent.discriminator) {
        generateDiscriminatorMarkerMethod(parent, modelDef);
      }
    }
    serdeImports.add('reflect');
    serdeImports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    if (obj.language.go!.byteArrayFormat) {
      modelDef.HasJSONByteArray = true;
    }
    if (!obj.language.go!.omitSerDeMethods) {
      generateJSONMarshaller(serdeImports, obj, modelDef);
      generateJSONUnmarshaller(serdeImports, modelDef);
    }
    modelDefs.push(modelDef);
  }
  return modelDefs;
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

// generates discriminator marker method
function generateDiscriminatorMarkerMethod(obj: ObjectSchema, modelDef: ModelDef) {
  const typeName = obj.language.go!.name;
  const receiver = modelDef.receiverName();
  const interfaceMethod = `Get${typeName}`;
  let method = `func (${receiver} *${modelDef.Language.name}) ${interfaceMethod}() *${typeName} {`;
  if (obj.language.go!.name === modelDef.Language.name) {
    // the marker method is on the discriminator itself, so just return the receiver
    method += ` return ${receiver} }\n\n`;
  } else {
    // the marker method is on a child type, so return an instance of the parent
    // type by copying the parent values into a new instance.
    method += `\n\treturn &${obj.language.go!.name}{\n`;
    for (const prop of values(aggregateProperties(obj))) {
      method += `\t\t${prop.language.go!.name}: ${modelDef.receiverName()}.${prop.language.go!.name},\n`;
    }
    method += '\t}\n}\n\n';
  }
  modelDef.Methods.push({ name: interfaceMethod, desc: `${interfaceMethod} implements the ${obj.language.go!.discriminatorInterface} interface for type ${modelDef.Language.name}.`, text: method });
}

function generateJSONMarshaller(imports: ImportManager, obj: ObjectSchema, modelDef: ModelDef) {
  if (!obj.discriminatorValue && (!modelDef.Properties || modelDef.Properties.length === 0)) {
    // non-discriminated types without content don't need a custom marshaller.
    // there is a case in network where child is allOf base and child has no properties.
    return;
  }
  imports.add('encoding/json');
  const typeName = modelDef.Language.name;
  const receiver = modelDef.receiverName();
  let marshaller = `func (${receiver} ${typeName}) MarshalJSON() ([]byte, error) {\n`;
  marshaller += '\tobjectMap := make(map[string]any)\n';
  marshaller += generateJSONMarshallerBody(obj, modelDef, imports);
  marshaller += '\treturn json.Marshal(objectMap)\n';
  marshaller += '}\n\n';
  modelDef.SerDeMethods.push({ name: 'MarshalJSON', desc: `MarshalJSON implements the json.Marshaller interface for type ${typeName}.`, text: marshaller });
}

function generateJSONMarshallerBody(obj: ObjectSchema, modelDef: ModelDef, imports: ImportManager): string {
  const receiver = modelDef.receiverName();
  let marshaller = '';
  let addlProps: DictionarySchema | undefined;
  for (const prop of values(modelDef.Properties)) {
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
    } else if (prop.schema.language.go!.rawJSONAsBytes) {
      marshaller += `\tpopulate(objectMap, "${prop.serializedName}", json.RawMessage(${receiver}.${prop.language.go!.name}))\n`;
    } else {
      if (prop.clientDefaultValue) {
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
        marshaller += `\tif ${receiver}.${prop.language.go!.name} == nil {\n\t\t${receiver}.${prop.language.go!.name} = to.Ptr(${getClientDefaultValue(prop)})\n\t}\n`;
      }
      let populate = 'populate';
      if (prop.schema.language.go!.internalTimeType) {
        populate += capitalize(prop.schema.language.go!.internalTimeType);
      } else if (prop.schema.type === SchemaType.Any) {
        populate += 'Any';
      }
      marshaller += `\t${populate}(objectMap, "${prop.serializedName}", ${receiver}.${prop.language.go!.name})\n`;
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
    marshaller += '\t\t}\n';
    marshaller += '\t}\n';
  }
  return marshaller;
}

function generateJSONUnmarshaller(imports: ImportManager, modelDef: ModelDef) {
  // there's a corner-case where a derived type might not add any new fields (Cookiecuttershark).
  // in this case skip adding the unmarshaller as it's not necessary and doesn't compile.
  if (!modelDef.Properties || modelDef.Properties.length === 0) {
    return;
  }
  imports.add('encoding/json');
  imports.add('fmt');
  const typeName = modelDef.Language.name;
  const receiver = modelDef.receiverName();
  let unmarshaller = `func (${receiver} *${typeName}) UnmarshalJSON(data []byte) error {\n`;
  unmarshaller += '\tvar rawMsg map[string]json.RawMessage\n';
  unmarshaller += '\tif err := json.Unmarshal(data, &rawMsg); err != nil {\n';
  unmarshaller += `\t\treturn fmt.Errorf("unmarshalling type %T: %v", ${receiver}, err)\n`;
  unmarshaller += '\t}\n';
  unmarshaller += generateJSONUnmarshallerBody(modelDef, imports);
  unmarshaller += '}\n\n';
  modelDef.SerDeMethods.push({ name: 'UnmarshalJSON', desc: `UnmarshalJSON implements the json.Unmarshaller interface for type ${typeName}.`, text: unmarshaller });
}

function generateJSONUnmarshallerBody(modelDef: ModelDef, imports: ImportManager): string {
  const receiver = modelDef.receiverName();
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
  };
  let unmarshalBody = '';
  unmarshalBody = '\tfor key, val := range rawMsg {\n';
  unmarshalBody += '\t\tvar err error\n';
  unmarshalBody += '\t\tswitch key {\n';
  let addlProps: DictionarySchema | undefined;
  for (const prop of values(modelDef.Properties)) {
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
    } else if (prop.schema.language.go!.rawJSONAsBytes) {
      unmarshalBody += `\t\t\t${receiver}.${prop.language.go!.name} = val\n`;
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

function generateXMLMarshaller(modelDef: ModelDef, imports: ImportManager) {
  // only needed for types with time.Time or where the XML name doesn't match the type name
  const receiver = modelDef.receiverName();
  const desc = `MarshalXML implements the xml.Marshaller interface for type ${modelDef.Language.name}.`;
  let text = `func (${receiver} ${modelDef.Language.name}) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {\n`;
  if (modelDef.Language.xmlWrapperName) {
    text += `\tstart.Name.Local = "${modelDef.Language.xmlWrapperName}"\n`;
  }
  text += generateAliasType(modelDef, receiver, true);
  for (const prop of values(modelDef.Properties)) {
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
      text += `\tif ${receiver}.${prop.language.go!.name} != nil {\n`;
      text += `\t\tencoded${prop.language.go!.name} := runtime.EncodeByteArray(${receiver}.${prop.language.go!.name}, runtime.Base64${base64Format}Format)\n`;
      text += `\t\taux.${prop.language.go!.name} = &encoded${prop.language.go!.name}\n`;
      text += '\t}\n';
    }
  }
  text += '\treturn enc.EncodeElement(aux, start)\n';
  text += '}\n\n';
  modelDef.SerDeMethods.push({ name: 'MarshalXML', desc: desc, text: text });
}

function generateXMLUnmarshaller(modelDef: ModelDef, imports: ImportManager) {
  // non-polymorphic case, must be something with time.Time
  const receiver = modelDef.receiverName();
  const desc = `UnmarshalXML implements the xml.Unmarshaller interface for type ${modelDef.Language.name}.`;
  let text = `func (${receiver} *${modelDef.Language.name}) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {\n`;
  text += generateAliasType(modelDef, receiver, false);
  text += '\tif err := dec.DecodeElement(aux, &start); err != nil {\n';
  text += '\t\treturn err\n';
  text += '\t}\n';
  for (const prop of values(modelDef.Properties)) {
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
      text += `\tif aux.${prop.language.go!.name} != nil {\n`;
      text += `\t\tif err := runtime.DecodeByteArray(*aux.${prop.language.go!.name}, &${receiver}.${prop.language.go!.name}, runtime.Base64${base64Format}Format); err != nil {\n`;
      text += '\t\t\treturn err\n';
      text += '\t\t}\n';
      text += '\t}\n';
    }
  }
  text += '\treturn nil\n';
  text += '}\n\n';
  modelDef.SerDeMethods.push({ name: 'UnmarshalXML', desc: desc, text: text });
}

// generates an alias type used by custom XML marshaller/unmarshaller
function generateAliasType(modelDef: ModelDef, receiver: string, forMarshal: boolean): string {
  let text = `\ttype alias ${modelDef.Language.name}\n`;
  text += '\taux := &struct {\n';
  text += '\t\t*alias\n';
  for (const prop of values(modelDef.Properties)) {
    const sn = getXMLSerialization(prop, modelDef.Language);
    if (prop.schema.type === SchemaType.DateTime) {
      text += `\t\t${prop.language.go!.name} *${prop.schema.language.go!.internalTimeType} \`${modelDef.Language.marshallingFormat}:"${sn}"\`\n`;
    } else if (prop.language.go!.isAdditionalProperties || prop.language.go!.needsXMLDictionaryUnmarshalling) {
      text += `\t\t${prop.language.go!.name} additionalProperties \`${modelDef.Language.marshallingFormat}:"${sn}"\`\n`;
    } else if (prop.language.go!.needsXMLArrayMarshalling) {
      text += `\t\t${prop.language.go!.name} *${prop.schema.language.go!.name} \`${modelDef.Language.marshallingFormat}:"${sn}"\`\n`;
    } else if (prop.schema.type === SchemaType.ByteArray) {
      text += `\t\t${prop.language.go!.name} *string \`${modelDef.Language.marshallingFormat}:"${sn}"\`\n`;
    }
  }
  text += '\t}{\n';
  let rec = receiver;
  if (forMarshal) {
    rec = '&' + rec;
  }
  text += `\t\talias: (*alias)(${rec}),\n`;
  if (forMarshal) {
    // emit code to initialize time fields
    for (const prop of values(modelDef.Properties)) {
      if (prop.schema.type !== SchemaType.DateTime) {
        continue;
      }
      text += `\t\t${prop.language.go!.name}: (*${prop.schema.language.go!.internalTimeType})(${receiver}.${prop.language.go!.name}),\n`;
    }
  }
  text += '\t}\n';
  return text;
}

// represents a method on a model
interface ModelMethod {
  name: string;
  desc: string;
  text: string;
}

// represents model definition as a Go struct
class ModelDef {
  readonly Language: Language;
  readonly Properties?: Array<Property>;
  readonly SerDeMethods: Array<ModelMethod>;
  readonly Methods: Array<ModelMethod>;
  HasJSONByteArray: boolean;
  HasAny: boolean;

  constructor(language: Language, props?: Array<Property>) {
    this.Language = language;
    this.Properties = props;
    if (this.Properties) {
      this.Properties.sort((a: Property, b: Property) => { return sortAscending(a.language.go!.name, b.language.go!.name); });
    }
    this.SerDeMethods = new Array<ModelMethod>();
    this.Methods = new Array<ModelMethod>();
    this.HasJSONByteArray = false;
    this.HasAny = false;
  }

  text(): string {
    let text = '';
    if (hasDescription(this.Language)) {
      text += `${comment(this.Language.description, '// ', undefined, commentLength)}\n`;
    }
    text += `type ${this.Language.name} struct {\n`;
    // used to track when to add an extra \n between fields that have comments
    let first = true;
    // group fields by required/optional/read-only in that order
    this.Properties?.sort((lhs: Property, rhs: Property): number => {
      if ((lhs.required && !rhs.required) || (!lhs.readOnly && rhs.readOnly)) {
        return -1;
      } else if ((rhs.readOnly && !lhs.readOnly) || (!rhs.readOnly && lhs.readOnly)) {
        return 1;
      } else {
        return 0;
      }
    });
    for (const prop of values(this.Properties)) {
      if (prop.language.go!.embeddedType) {
        continue;
      }
      if (hasDescription(prop.language.go!)) {
        if (!first) {
          // add an extra new-line between fields IFF the field
          // has a comment and it's not the very first one.
          text += '\n';
        }
        text += `\t${comment(prop.language.go!.description, '// ', undefined, commentLength)}\n`;
      }
      let typeName = prop.schema.language.go!.name;
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
        readOnly = ' azure:"ro"';
      }
      let tag = '';
      // only emit tags for XML; JSON uses custom marshallers/unmarshallers
      if (this.Language.marshallingFormat === 'xml' && !prop.language.go!.isAdditionalProperties) {
        tag = ` \`${this.Language.marshallingFormat}:"${serialization}"${readOnly}\``;
      }
      text += `\t${prop.language.go!.name} ${getStar(prop.language.go!)}${typeName}${tag}\n`;
      first = false;
    }
    text += '}\n\n';
    return text;
  }

  receiverName(): string {
    const typeName = this.Language.name;
    return typeName[0].toLowerCase();
  }
}

export function getXMLSerialization(prop: Property, lang: Language): string {
  let serialization = prop.serializedName;
  // default to using the serialization name
  if (prop.schema.serialization?.xml?.name) {
    // xml can specifiy its own name, prefer that if available
    serialization = prop.schema.serialization.xml.name;
  } else if (prop.schema.serialization?.xml?.text) {
    // type has the x-ms-text attribute applied so it should be character data, not a node (https://github.com/Azure/autorest/tree/main/docs/extensions#x-ms-text)
    // see https://pkg.go.dev/encoding/xml#Unmarshal for what ,chardata actually means
    serialization = ',chardata';
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

export function getStar(lang: Language): string {
  // lang is assumed to be go
  if (lang.byValue === true) {
    return '';
  }
  return '*';
}
