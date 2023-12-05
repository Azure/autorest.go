/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../codemodel.go/src/gocodemodel.js';
import { capitalize, comment } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import { commentLength, contentPreamble, formatLiteralValue, sortAscending } from './helpers.js';
import { ImportManager } from './imports.js';

export interface ModelsSerDe {
  models: string;
  serDe: string;
}

// Creates the content in models.go
export async function generateModels(codeModel: go.CodeModel): Promise<ModelsSerDe> {
  // this list of packages to import
  const modelImports = new ImportManager();
  const serdeImports = new ImportManager();
  let modelText = contentPreamble(codeModel);

  // we do model generation first as it can add imports to the imports list
  const modelDefs = generateModelDefs(modelImports, serdeImports, codeModel);

  modelText += modelImports.text();

  // structs
  let needsJSONPopulate = false;
  let needsJSONUnpopulate = false;
  let needsJSONPopulateByteArray = false;
  let needsJSONPopulateAny = false;
  let serdeTextBody = '';
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
    serdeText = contentPreamble(codeModel);
    serdeText += serdeImports.text();
    serdeText += serdeTextBody;
  }
  return {
    models: modelText,
    serDe: serdeText
  };
}

function generateModelDefs(modelImports: ImportManager, serdeImports: ImportManager, codeModel: go.CodeModel): Array<ModelDef> {
  const models = codeModel.models;
  const modelDefs = new Array<ModelDef>();
  for (const model of models) {
    const modelDef = new ModelDef(model.name, model.format, model.fields, model.description);
    for (const field of values(modelDef.Fields)) {
      modelImports.addImportForType(field.type);
      if (go.isPrimitiveType(field.type) && field.type.typeName === 'any') {
        modelDef.HasAny = true;
      }
    }

    if (go.isModelType(model) && model.format === 'xml' && !model.annotations.omitSerDeMethods) {
      serdeImports.add('encoding/xml');
      let needsDateTimeMarshalling = false;
      let byteArrayFormat = false;
      for (const field of values(model.fields)) {
        serdeImports.addImportForType(field.type);
        if (go.isTimeType(field.type)) {
          needsDateTimeMarshalling = true;
        } else if (go.isBytesType(field.type)) {
          byteArrayFormat = true;
        }
      }
      // due to differences in XML marshallers/unmarshallers, we use different codegen than for JSON
      if (needsDateTimeMarshalling || model.xml?.wrapper || needsXMLArrayMarshalling(model) || byteArrayFormat) {
        generateXMLMarshaller(model, modelDef, serdeImports);
        if (needsDateTimeMarshalling || byteArrayFormat) {
          generateXMLUnmarshaller(model, modelDef, serdeImports);
        }
      } else if (needsXMLDictionaryUnmarshalling(model)) {
        generateXMLUnmarshaller(model, modelDef, serdeImports);
      }
      modelDefs.push(modelDef);
      continue;
    }
    if (go.isPolymorphicType(model)) {
      generateDiscriminatorMarkerMethod(model.interface, modelDef);
      for (let parent = model.interface.parent; parent !== undefined; parent = parent.parent) {
        generateDiscriminatorMarkerMethod(parent, modelDef);
      }
    }
    serdeImports.add('reflect');
    serdeImports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    for (const field of values(model.fields)) {
      if (go.isBytesType(field.type)) {
        modelDef.HasJSONByteArray = true;
      }
    }
    if (!model.annotations.omitSerDeMethods) {
      generateJSONMarshaller(model, modelDef, serdeImports);
      generateJSONUnmarshaller(model, modelDef, serdeImports, codeModel.options);
    }
    modelDefs.push(modelDef);
  }
  return modelDefs;
}

function needsXMLDictionaryUnmarshalling(modelType: go.ModelType): boolean {
  for (const field of values(modelType.fields)) {
    // additional properties uses an internal wrapper type with its own unmarshaller
    if (go.isMapType(field.type) && !field.annotations.isAdditionalProperties) {
      return true;
    }
  }
  return false;
}

function needsXMLArrayMarshalling(modelType: go.ModelType): boolean {
  for (const prop of values(modelType.fields)) {
    if (go.isSliceType(prop.type)) {
      return true;
    }
  }
  return false;
}

// generates discriminator marker method
function generateDiscriminatorMarkerMethod(type: go.InterfaceType, modelDef: ModelDef) {
  const typeName = type.rootType.name;
  const receiver = modelDef.receiverName();
  const interfaceMethod = `Get${typeName}`;
  let method = `func (${receiver} *${modelDef.Name}) ${interfaceMethod}() *${typeName} {`;
  if (type.rootType.name === modelDef.Name) {
    // the marker method is on the discriminator itself, so just return the receiver
    method += ` return ${receiver} }\n\n`;
  } else {
    // the marker method is on a child type, so return an instance of the parent
    // type by copying the parent values into a new instance.
    method += `\n\treturn &${type.rootType.name}{\n`;
    for (const field of values(type.rootType.fields)) {
      method += `\t\t${field.fieldName}: ${modelDef.receiverName()}.${field.fieldName},\n`;
    }
    method += '\t}\n}\n\n';
  }
  modelDef.Methods.push({ name: interfaceMethod, desc: `${interfaceMethod} implements the ${type.name} interface for type ${modelDef.Name}.`, text: method });
}

function generateJSONMarshaller(modelType: go.ModelType | go.PolymorphicType, modelDef: ModelDef, imports: ImportManager) {
  if (go.isModelType(modelType) && modelType.fields.length === 0) {
    // non-discriminated types without content don't need a custom marshaller.
    // there is a case in network where child is allOf base and child has no properties.
    return;
  }
  imports.add('encoding/json');
  const typeName = modelDef.Name;
  const receiver = modelDef.receiverName();
  let marshaller = `func (${receiver} ${typeName}) MarshalJSON() ([]byte, error) {\n`;
  marshaller += '\tobjectMap := make(map[string]any)\n';
  marshaller += generateJSONMarshallerBody(modelType, receiver, imports);
  marshaller += '\treturn json.Marshal(objectMap)\n';
  marshaller += '}\n\n';
  modelDef.SerDeMethods.push({ name: 'MarshalJSON', desc: `MarshalJSON implements the json.Marshaller interface for type ${typeName}.`, text: marshaller });
}

function generateJSONMarshallerBody(modelType: go.ModelType | go.PolymorphicType, receiver: string, imports: ImportManager): string {
  let marshaller = '';
  let addlProps: go.MapType | undefined;
  for (const field of values(modelType.fields)) {
    if (go.isMapType(field.type) && field.annotations.isAdditionalProperties) {
      addlProps = field.type;
      continue;
    }
    if (field.annotations.isDiscriminator) {
      if (field.defaultValue) {
        marshaller += `\tobjectMap["${field.serializedName}"] = ${formatLiteralValue(field.defaultValue)}\n`;
      } else {
        // if there's no discriminator value (e.g. Fish in test server), use the field's value.
        // this will enable support for custom types that aren't (yet) described in the swagger.
        marshaller += `\tobjectMap["${field.serializedName}"] = ${receiver}.${field.fieldName}\n`;
      }
    } else if (go.isBytesType(field.type)) {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      marshaller += `\tpopulateByteArray(objectMap, "${field.serializedName}", ${receiver}.${field.fieldName}, runtime.Base64${field.type.encoding}Format)\n`;
    } else if (go.isSliceType(field.type) && go.isTimeType(field.type.elementType)) {
      const source = `${receiver}.${field.fieldName}`;
      let elementPtr = '*';
      if (field.type.elementTypeByValue) {
        elementPtr = '';
      }
      marshaller += `\taux := make([]${elementPtr}${field.type.elementType.dateTimeFormat}, len(${source}), len(${source}))\n`;
      marshaller += `\tfor i := 0; i < len(${source}); i++ {\n`;
      marshaller += `\t\taux[i] = (${elementPtr}${field.type.elementType.dateTimeFormat})(${source}[i])\n`;
      marshaller += '\t}\n';
      marshaller += `\tpopulate(objectMap, "${field.serializedName}", aux)\n`;
    } else if (go.isLiteralValue(field.type)) {
      marshaller += `\tobjectMap["${field.serializedName}"] = ${formatLiteralValue(field.type)}\n`;
    } else if (go.isSliceType(field.type) && field.type.rawJSONAsBytes) {
      marshaller += `\tpopulate(objectMap, "${field.serializedName}", json.RawMessage(${receiver}.${field.fieldName}))\n`;
    } else {
      if (field.defaultValue) {
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
        marshaller += `\tif ${receiver}.${field.fieldName} == nil {\n\t\t${receiver}.${field.fieldName} = to.Ptr(${formatLiteralValue(field.defaultValue)})\n\t}\n`;
      }
      let populate = 'populate';
      if (go.isTimeType(field.type)) {
        populate += capitalize(field.type.dateTimeFormat);
      } else if (go.isPrimitiveType(field.type) && field.type.typeName === 'any') {
        populate += 'Any';
      }
      marshaller += `\t${populate}(objectMap, "${field.serializedName}", ${receiver}.${field.fieldName})\n`;
    }
  }
  if (addlProps) {
    marshaller += `\tif ${receiver}.AdditionalProperties != nil {\n`;
    marshaller += `\t\tfor key, val := range ${receiver}.AdditionalProperties {\n`;
    let assignment = 'val';
    if (go.isTimeType(addlProps.valueType)) {
      assignment = `(*${addlProps.valueType.dateTimeFormat})(val)`;
    }
    marshaller += `\t\t\tobjectMap[key] = ${assignment}\n`;
    marshaller += '\t\t}\n';
    marshaller += '\t}\n';
  }
  return marshaller;
}

function generateJSONUnmarshaller(modelType: go.ModelType | go.PolymorphicType, modelDef: ModelDef, imports: ImportManager, options: go.Options) {
  // there's a corner-case where a derived type might not add any new fields (Cookiecuttershark).
  // in this case skip adding the unmarshaller as it's not necessary and doesn't compile.
  if (modelDef.Fields.length === 0) {
    return;
  }
  imports.add('encoding/json');
  imports.add('fmt');
  const typeName = modelDef.Name;
  const receiver = modelDef.receiverName();
  let unmarshaller = `func (${receiver} *${typeName}) UnmarshalJSON(data []byte) error {\n`;
  unmarshaller += '\tvar rawMsg map[string]json.RawMessage\n';
  unmarshaller += '\tif err := json.Unmarshal(data, &rawMsg); err != nil {\n';
  unmarshaller += `\t\treturn fmt.Errorf("unmarshalling type %T: %v", ${receiver}, err)\n`;
  unmarshaller += '\t}\n';
  unmarshaller += generateJSONUnmarshallerBody(modelType, receiver, imports, options);
  unmarshaller += '}\n\n';
  modelDef.SerDeMethods.push({ name: 'UnmarshalJSON', desc: `UnmarshalJSON implements the json.Unmarshaller interface for type ${typeName}.`, text: unmarshaller });
}

function generateJSONUnmarshallerBody(modelType: go.ModelType | go.PolymorphicType, receiver: string, imports: ImportManager, options: go.Options): string {
  const emitAddlProps = function (tab: string, addlProps: go.MapType): string {
    let addlPropsText = `${tab}\t\tif ${receiver}.AdditionalProperties == nil {\n`;
    let ref = '';
    if (!addlProps.valueTypeByValue) {
      ref = '&';
    }
    addlPropsText += `${tab}\t\t\t${receiver}.AdditionalProperties = ${go.getTypeDeclaration(addlProps)}{}\n`;
    addlPropsText += `${tab}\t\t}\n`;
    addlPropsText += `${tab}\t\tif val != nil {\n`;
    let auxType = go.getTypeDeclaration(addlProps.valueType);
    let assignment = `${ref}aux`;
    if (go.isTimeType(addlProps.valueType)) {
      imports.add('time');
      auxType = addlProps.valueType.dateTimeFormat!;
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
  let addlProps: go.MapType | undefined;
  for (const field of values(modelType.fields)) {
    if (go.isMapType(field.type) && field.annotations.isAdditionalProperties) {
      addlProps = field.type;
      continue;
    }
    unmarshalBody += `\t\tcase "${field.serializedName}":\n`;
    if (hasDiscriminatorInterface(field.type)) {
      unmarshalBody += generateDiscriminatorUnmarshaller(field, receiver);
    } else if (go.isTimeType(field.type)) {
      unmarshalBody += `\t\t\t\terr = unpopulate${capitalize(field.type.dateTimeFormat)}(val, "${field.fieldName}", &${receiver}.${field.fieldName})\n`;
    } else if (go.isSliceType(field.type) && go.isTimeType(field.type.elementType)) {
      imports.add('time');
      let elementPtr = '*';
      if (field.type.elementTypeByValue) {
        elementPtr = '';
      }
      unmarshalBody += `\t\t\tvar aux []${elementPtr}${field.type.elementType.dateTimeFormat}\n`;
      unmarshalBody += `\t\t\terr = unpopulate(val, "${field.fieldName}", &aux)\n`;
      unmarshalBody += '\t\t\tfor _, au := range aux {\n';
      unmarshalBody += `\t\t\t\t${receiver}.${field.fieldName} = append(${receiver}.${field.fieldName}, (${elementPtr}time.Time)(au))\n`;
      unmarshalBody += '\t\t\t}\n';
    } else if (go.isBytesType(field.type)) {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      unmarshalBody += `\t\t\terr = runtime.DecodeByteArray(string(val), &${receiver}.${field.fieldName}, runtime.Base64${field.type.encoding}Format)\n`;
    } else if (go.isSliceType(field.type) && field.type.rawJSONAsBytes) {
      unmarshalBody += `\t\t\t${receiver}.${field.fieldName} = val\n`;
    } else {
      unmarshalBody += `\t\t\t\terr = unpopulate(val, "${field.fieldName}", &${receiver}.${field.fieldName})\n`;
    }
    unmarshalBody += '\t\t\tdelete(rawMsg, key)\n';
  }
  if (addlProps) {
    unmarshalBody += '\t\tdefault:\n';
    unmarshalBody += emitAddlProps('\t', addlProps);
  } else if (options.disallowUnknownFields) {
    unmarshalBody += '\t\tdefault:\n';
    unmarshalBody += `\t\t\terr = fmt.Errorf("unmarshalling type %T, unknown field %q", ${receiver}, key)\n`;
  }
  unmarshalBody += '\t\t}\n';
  unmarshalBody += '\t\tif err != nil {\n';
  unmarshalBody += `\t\t\treturn fmt.Errorf("unmarshalling type %T: %v", ${receiver}, err)\n`;
  unmarshalBody += '\t\t}\n';
  unmarshalBody += '\t}\n'; // end for key, val := range rawMsg
  unmarshalBody += '\treturn nil\n';
  return unmarshalBody;
}

// returns true if item has a discriminator interface.
// recursively called for arrays and dictionaries.
function hasDiscriminatorInterface(item: go.PossibleType): boolean {
  if (go.isInterfaceType(item)) {
    return true;
  } else if (go.isMapType(item)) {
    return hasDiscriminatorInterface(item.valueType);
  } else if (go.isSliceType(item)) {
    return hasDiscriminatorInterface(item.elementType);
  }
  return false;
}

// returns the text for unmarshalling a discriminated type
function generateDiscriminatorUnmarshaller(field: go.ModelField, receiver: string): string {
  const startingIndentation = '\t\t\t';
  const propertyName = field.fieldName;

  // these are the simple, non-nested cases (e.g. IterfaceType, []InterfaceType, map[string]InterfaceType)
  if (go.isInterfaceType(field.type)) {
    return `${startingIndentation}${receiver}.${propertyName}, err = unmarshal${field.type.name}(val)\n`;
  } else if (go.isSliceType(field.type) && go.isInterfaceType(field.type.elementType)) {
    return `${startingIndentation}${receiver}.${propertyName}, err = unmarshal${field.type.elementType.name}Array(val)\n`;
  } else if (go.isMapType(field.type) && go.isInterfaceType(field.type.valueType)) {
    return `${startingIndentation}${receiver}.${propertyName}, err = unmarshal${field.type.valueType.name}Map(val)\n`;
  }

  // nested case (e.g. [][]InterfaceType, map[string]map[string]InterfaceType etc)
  // first, unmarshal the raw data
  const rawTargetVar = `${field.serializedName}Raw`;
  let text = `${startingIndentation}var ${rawTargetVar} ${recursiveGetDiscriminatorTypeName(field.type, true)}\n`;
  text += `${startingIndentation}if err = json.Unmarshal(val, &${rawTargetVar}); err != nil {\n`;
  text += `${startingIndentation}\treturn err\n${startingIndentation}}\n`;

  // create a local instantiation of the final type
  const finalTargetVar = field.serializedName;
  let finalTargetCtor = recursiveGetDiscriminatorTypeName(field.type, false);
  if (go.isSliceType(field.type)) {
    finalTargetCtor = `make(${finalTargetCtor}, len(${rawTargetVar}))`;
  } else {
    // must be a dictionary
    finalTargetCtor = `${finalTargetCtor}{}`;
  }
  text += `${startingIndentation}${finalTargetVar} := ${finalTargetCtor}\n`;

  // now populate the final type
  text += recursivePopulateDiscriminator(field.type, receiver, rawTargetVar, finalTargetVar, startingIndentation, 1);

  // finally, assign the final target to the property
  text += `${startingIndentation}${receiver}.${propertyName} = ${finalTargetVar}\n`;
  return text;
}

// constructs the type name for a nested discriminated type
// raw e.g. map[string]json.RawMessage, []json.RawMessage etc
// !raw e.g. map[string]map[string]InterfaceType, [][]InterfaceType etc
function recursiveGetDiscriminatorTypeName(item: go.PossibleType, raw: boolean): string {
  // when raw is true, stop recursing at the level before the leaf schema
  if (go.isSliceType(item)) {
    if (!raw || !go.isInterfaceType(item.elementType)) {
      return `[]${recursiveGetDiscriminatorTypeName(item.elementType, raw)}`;
    }
  } else if (go.isMapType(item)) {
    if (!raw || !go.isInterfaceType(item.valueType)) {
      return `map[string]${recursiveGetDiscriminatorTypeName(item.valueType, raw)}`;
    }
  }
  if (raw) {
    return 'json.RawMessage';
  }
  return go.getTypeDeclaration(item);
}

// recursively constructs the text to populate a nested discriminator
function recursivePopulateDiscriminator(item: go.PossibleType, receiver: string, rawSrc: string, dest: string, indent: string, nesting: number): string {
  let text = '';
  let interfaceName = '';
  let targetType = '';

  if (go.isSliceType(item)) {
    if (!go.isInterfaceType(item.elementType)) {
      if (nesting > 1) {
        // at nestling level 1, the destination var was already created in generateDiscriminatorUnmarshaller()
        text += `${indent}${dest} = make(${recursiveGetDiscriminatorTypeName(item, false)}, len(${rawSrc}))\n`;
      }

      text += `${indent}for i${nesting} := range ${rawSrc} {\n`;
      rawSrc = `${rawSrc}[i${nesting}]`; // source becomes each element in the source slice
      dest = `${dest}[i${nesting}]`; // update destination to each element in the destination slice
      text += recursivePopulateDiscriminator(item.elementType, receiver, rawSrc, dest, indent+'\t', nesting+1);
      text += `${indent}}\n`;
      return text;
    }

    // we're at leaf node - 1, so get the interface from the element's type
    interfaceName = go.getTypeDeclaration(item.elementType);
    targetType = 'Array';
  } else if (go.isMapType(item)) {
    if (!go.isInterfaceType(item.valueType)) {
      if (nesting > 1) {
        // at nestling level 1, the destination var was already created in generateDiscriminatorUnmarshaller()
        text += `${indent}${dest} = ${recursiveGetDiscriminatorTypeName(item, false)}{}\n`;
      }

      text += `${indent}for k${nesting}, v${nesting} := range ${rawSrc} {\n`;
      rawSrc = `v${nesting}`; // source becomes the current value in the source map
      dest = `${dest}[k${nesting}]`; // update destination to the destination map's value for the current key
      text += recursivePopulateDiscriminator(item.valueType, receiver, rawSrc, dest, indent+'\t', nesting+1);
      text += `${indent}}\n`;
      return text;
    }

    // we're at leaf node - 1, so get the interface from the element's type
    interfaceName = go.getTypeDeclaration(item.valueType);
    targetType = 'Map';
  }

  text += `${indent}${dest}, err = unmarshal${interfaceName}${targetType}(${rawSrc})\n`;
  text += `${indent}if err != nil {\n${indent}\treturn fmt.Errorf("unmarshalling type %T: %v", ${receiver}, err)\n${indent}}\n`;

  return text;
}

function generateXMLMarshaller(modelType: go.ModelType, modelDef: ModelDef, imports: ImportManager) {
  // only needed for types with time.Time or where the XML name doesn't match the type name
  const receiver = modelDef.receiverName();
  const desc = `MarshalXML implements the xml.Marshaller interface for type ${modelDef.Name}.`;
  let text = `func (${receiver} ${modelDef.Name}) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {\n`;
  if (modelType.xml?.wrapper) {
    text += `\tstart.Name.Local = "${modelType.xml.wrapper}"\n`;
  }
  text += generateAliasType(modelType, receiver, true);
  for (const field of values(modelDef.Fields)) {
    if (go.isSliceType(field.type)) {
      text += `\tif ${receiver}.${field.fieldName} != nil {\n`;
      text += `\t\taux.${field.fieldName} = &${receiver}.${field.fieldName}\n`;
      text += '\t}\n';
    } else if (go.isBytesType(field.type)) {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      text += `\tif ${receiver}.${field.fieldName} != nil {\n`;
      text += `\t\tencoded${field.fieldName} := runtime.EncodeByteArray(${receiver}.${field.fieldName}, runtime.Base64${field.type.encoding}Format)\n`;
      text += `\t\taux.${field.fieldName} = &encoded${field.fieldName}\n`;
      text += '\t}\n';
    }
  }
  text += '\treturn enc.EncodeElement(aux, start)\n';
  text += '}\n\n';
  modelDef.SerDeMethods.push({ name: 'MarshalXML', desc: desc, text: text });
}

function generateXMLUnmarshaller(modelType: go.ModelType, modelDef: ModelDef, imports: ImportManager) {
  // non-polymorphic case, must be something with time.Time
  const receiver = modelDef.receiverName();
  const desc = `UnmarshalXML implements the xml.Unmarshaller interface for type ${modelDef.Name}.`;
  let text = `func (${receiver} *${modelDef.Name}) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {\n`;
  text += generateAliasType(modelType, receiver, false);
  text += '\tif err := dec.DecodeElement(aux, &start); err != nil {\n';
  text += '\t\treturn err\n';
  text += '\t}\n';
  for (const field of values(modelDef.Fields)) {
    if (go.isTimeType(field.type)) {
      text += `\t${receiver}.${field.fieldName} = (*time.Time)(aux.${field.fieldName})\n`;
    } else if (field.annotations.isAdditionalProperties || go.isMapType(field.type)) {
      text += `\t${receiver}.${field.fieldName} = (map[string]*string)(aux.${field.fieldName})\n`;
    } else if (go.isBytesType(field.type)) {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      text += `\tif aux.${field.fieldName} != nil {\n`;
      text += `\t\tif err := runtime.DecodeByteArray(*aux.${field.fieldName}, &${receiver}.${field.fieldName}, runtime.Base64${field.type.encoding}Format); err != nil {\n`;
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
function generateAliasType(modelType: go.ModelType, receiver: string, forMarshal: boolean): string {
  let text = `\ttype alias ${modelType.name}\n`;
  text += '\taux := &struct {\n';
  text += '\t\t*alias\n';
  for (const field of values(modelType.fields)) {
    const sn = getXMLSerialization(field, false);
    if (go.isTimeType(field.type)) {
      text += `\t\t${field.fieldName} *${field.type.dateTimeFormat} \`${modelType.format}:"${sn}"\`\n`;
    } else if (field.annotations.isAdditionalProperties || go.isMapType(field.type)) {
      text += `\t\t${field.fieldName} additionalProperties \`${modelType.format}:"${sn}"\`\n`;
    } else if (go.isSliceType(field.type)) {
      text += `\t\t${field.fieldName} *${go.getTypeDeclaration(field.type)} \`${modelType.format}:"${sn}"\`\n`;
    } else if (go.isBytesType(field.type)) {
      text += `\t\t${field.fieldName} *string \`${modelType.format}:"${sn}"\`\n`;
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
    for (const field of modelType.fields) {
      if (!go.isTimeType(field.type)) {
        continue;
      }
      text += `\t\t${field.fieldName}: (*${field.type.dateTimeFormat})(${receiver}.${field.fieldName}),\n`;
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
  readonly Name: string;
  readonly Format: go.ModelFormat;
  readonly Description?: string;
  readonly Fields: Array<go.ModelField>;
  readonly SerDeMethods: Array<ModelMethod>;
  readonly Methods: Array<ModelMethod>;
  HasJSONByteArray: boolean;
  HasAny: boolean;

  constructor(name: string, format: go.ModelFormat, fields: Array<go.ModelField>, description?: string) {
    this.Name = name;
    this.Format = format;
    this.Description = description;
    this.Fields = fields;
    this.SerDeMethods = new Array<ModelMethod>();
    this.Methods = new Array<ModelMethod>();
    this.HasJSONByteArray = false;
    this.HasAny = false;
  }

  text(): string {
    let text = '';
    if (this.Description) {
      text += `${comment(this.Description, '// ', undefined, commentLength)}\n`;
    }
    text += `type ${this.Name} struct {\n`;

    // group fields by required/optional/read-only in that order
    this.Fields?.sort((lhs: go.ModelField, rhs: go.ModelField): number => {
      if ((lhs.annotations.required && !rhs.annotations.required) || (!lhs.annotations.readOnly && rhs.annotations.readOnly)) {
        return -1;
      } else if ((rhs.annotations.readOnly && !lhs.annotations.readOnly) || (!rhs.annotations.readOnly && lhs.annotations.readOnly)) {
        return 1;
      } else {
        return 0;
      }
    });

    // used to track when to add an extra \n between fields that have comments
    let first = true;

    for (const field of values(this.Fields)) {
      if (field.description) {
        if (!first) {
          // add an extra new-line between fields IFF the field
          // has a comment and it's not the very first one.
          text += '\n';
        }
        text += `\t${comment(field.description, '// ', undefined, commentLength)}\n`;
      }
      let typeName = go.getTypeDeclaration(field.type);
      if (go.isLiteralValue(field.type)) {
        // for constants we use the underlying type name
        typeName = go.getLiteralValueTypeName(field.type.type);
      }
      let serialization = field.serializedName;
      if (this.Format === 'json') {
        serialization += ',omitempty';
      } else if (this.Format === 'xml') {
        serialization = getXMLSerialization(field, false);
      }
      let tag = '';
      // only emit tags for XML; JSON uses custom marshallers/unmarshallers
      if (this.Format === 'xml' && !field.annotations.isAdditionalProperties) {
        tag = ` \`${this.Format}:"${serialization}"\``;
      }
      text += `\t${field.fieldName} ${getStar(field.byValue)}${typeName}${tag}\n`;
      first = false;
    }

    text += '}\n\n';
    return text;
  }

  receiverName(): string {
    const typeName = this.Name;
    return typeName[0].toLowerCase();
  }
}

export function getXMLSerialization(field: go.ModelField, isResponseEnvelope: boolean): string {
  let serialization = field.serializedName;
  // default to using the serialization name
  if (field.xml?.name) {
    // xml can specifiy its own name, prefer that if available
    serialization = field.xml.name;
  } else if (field.xml?.text) {
    // type has the x-ms-text attribute applied so it should be character data, not a node (https://github.com/Azure/autorest/tree/main/docs/extensions#x-ms-text)
    // see https://pkg.go.dev/encoding/xml#Unmarshal for what ,chardata actually means
    serialization = ',chardata';
  }
  if (field.xml?.attribute) {
    // value comes from an xml attribute
    serialization += ',attr';
  } else if (go.isSliceType(field.type)) {
    // start with the serialized name of the element, preferring xml name if available
    let inner = go.getTypeDeclaration(field.type.elementType);
    if (field.xml?.name) {
      inner = field.xml.name;
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
    if (field.xml?.wraps && !isResponseEnvelope) {
      serialization += `>${field.xml.wraps}`;
    } else {
      serialization = inner;
    }
  }
  return serialization;
}

export function getStar(byValue: boolean): string {
  if (byValue === true) {
    return '';
  }
  return '*';
}
