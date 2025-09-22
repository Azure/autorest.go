/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../codemodel.go/src/index.js';
import { capitalize, comment } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import * as helpers from './helpers.js';
import { ImportManager } from './imports.js';

export interface ModelsSerDe {
  models: string;
  serDe: string;
}

// Creates the content in models.go
export async function generateModels(codeModel: go.CodeModel): Promise<ModelsSerDe> {
  if (codeModel.models.length === 0) {
    return {
      models: '',
      serDe: ''
    };
  }

  // this list of packages to import
  const modelImports = new ImportManager();
  const serdeImports = new ImportManager();
  let modelText = helpers.contentPreamble(codeModel);

  // we do model generation first as it can add imports to the imports list
  const modelDefs = generateModelDefs(modelImports, serdeImports, codeModel);

  modelText += modelImports.text();

  // structs
  let needsJSONPopulate = false;
  let needsJSONUnpopulate = false;
  let needsJSONPopulateByteArray = false;
  let needsJSONPopulateAny = false;
  let needsJSONPopulateMultipart = false;
  let serdeTextBody = '';
  for (const modelDef of values(modelDefs)) {
    modelText += modelDef.text();

    modelDef.Methods.sort((a: ModelMethod, b: ModelMethod) => { return helpers.sortAscending(a.name, b.name); });
    for (const method of values(modelDef.Methods)) {
      if (method.desc.length > 0) {
        modelText += `${comment(method.desc, '// ', undefined, helpers.commentLength)}\n`;
      }
      modelText += method.text;
    }

    modelDef.SerDe.methods.sort((a: ModelMethod, b: ModelMethod) => { return helpers.sortAscending(a.name, b.name); });
    for (const method of values(modelDef.SerDe.methods)) {
      if (method.desc.length > 0) {
        serdeTextBody += `${comment(method.desc, '// ', undefined, helpers.commentLength)}\n`;
      }
      serdeTextBody += method.text;
    }
    if (modelDef.SerDe.needsJSONPopulate) {
      needsJSONPopulate = true;
    }
    if (modelDef.SerDe.needsJSONUnpopulate) {
      needsJSONUnpopulate = true;
    }
    if (modelDef.SerDe.needsJSONPopulateByteArray) {
      needsJSONPopulateByteArray = true;
    }
    if (modelDef.SerDe.needsJSONPopulateAny) {
      needsJSONPopulateAny = true;
    }
    if (modelDef.SerDe.needsJSONPopulateMultipart) {
      needsJSONPopulateMultipart = true;
    }
  }
  if (needsJSONPopulate) {
    serdeImports.add('reflect');
    serdeImports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
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
    serdeImports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
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
    serdeImports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    serdeImports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
    serdeTextBody += 'func populateByteArray[T any](m map[string]any, k string, b []T, convert func() any) {\n';
    serdeTextBody += '\tif azcore.IsNullValue(b) {\n';
    serdeTextBody += '\t\tm[k] = nil\n';
    serdeTextBody += '\t} else if len(b) == 0 {\n';
    serdeTextBody += '\t\treturn\n';
    serdeTextBody += '\t} else {\n';
    serdeTextBody += '\t\tm[k] = convert()\n';
    serdeTextBody += '\t}\n';
    serdeTextBody += '}\n\n';
  }
  if (needsJSONUnpopulate) {
    serdeImports.add('fmt');
    serdeTextBody += 'func unpopulate(data json.RawMessage, fn string, v any) error {\n';
    serdeTextBody += '\tif data == nil || string(data) == "null" {\n';
    serdeTextBody += '\t\treturn nil\n';
    serdeTextBody += '\t}\n';
    serdeTextBody += '\tif err := json.Unmarshal(data, v); err != nil {\n';
    serdeTextBody += '\t\treturn fmt.Errorf("struct field %s: %v", fn, err)\n';
    serdeTextBody += '\t}\n';
    serdeTextBody += '\treturn nil\n';
    serdeTextBody += '}\n\n';
  }
  if (needsJSONPopulateMultipart) {
    serdeImports.add('encoding/json');
    serdeTextBody += 'func populateMultipartJSON(m map[string]any, k string, v any) error {\n';
    serdeTextBody += '\tdata, err := json.Marshal(v)\n';
    serdeTextBody += '\tif err != nil {\n\t\treturn err\n\t}\n';
    serdeTextBody += '\tm[k] = data\n\treturn nil\n';
    serdeTextBody += '}\n\n';
  }
  let serdeText = '';
  if (serdeTextBody.length > 0) {
    serdeText = helpers.contentPreamble(codeModel);
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
    for (const field of model.fields) {
      const descriptionMods = new Array<string>();
      if (field.annotations.readOnly) {
        descriptionMods.push('READ-ONLY');
      } else if (field.annotations.required && (field.type.kind !== 'literal' || model.usage === go.UsageFlags.Output)) {
        descriptionMods.push('REQUIRED');
      } else if (field.type.kind === 'literal') {
        if (!field.annotations.required) {
          descriptionMods.push('FLAG');
        }
        descriptionMods.push('CONSTANT');
      }
      if (field.type.kind === 'literal' && model.usage !== go.UsageFlags.Output) {
        // add a comment with the const value for const properties that are sent over the wire
        if (field.docs.description) {
          field.docs.description += '\n';
        } else {
          field.docs.description = '';
        }
        field.docs.description += `Field has constant value ${helpers.formatLiteralValue(field.type, false)}, any specified value is ignored.`;
      }
      if (field.docs.description) {
        descriptionMods.push(field.docs.description);
      } else if (field.type.kind === 'rawJSON') {
        // add a basic description if one isn't available
        descriptionMods.push('The contents of this field are raw JSON.');
      }
      field.docs.description = descriptionMods.join('; ');
    }

    const serDeFormat = helpers.getSerDeFormat(model, codeModel);
    const modelDef = new ModelDef(model.name, serDeFormat, model.fields, model.docs);
    for (const field of values(modelDef.Fields)) {
      modelImports.addImportForType(field.type);
    }

    if (model.kind === 'model' && serDeFormat === 'XML' && !model.annotations.omitSerDeMethods) {
      serdeImports.add('encoding/xml');
      let needsDateTimeMarshalling = false;
      let byteArrayFormat = false;
      for (const field of values(model.fields)) {
        serdeImports.addImportForType(field.type);
        if (field.type.kind === 'time') {
          needsDateTimeMarshalling = true;
        } else if (field.type.kind === 'encodedBytes') {
          byteArrayFormat = true;
        }
      }
      // due to differences in XML marshallers/unmarshallers, we use different codegen than for JSON
      if (needsDateTimeMarshalling || model.xml?.wrapper || needsXMLArrayMarshalling(model) || byteArrayFormat) {
        generateXMLMarshaller(model, modelDef, serdeImports);
        if (needsDateTimeMarshalling || byteArrayFormat) {
          generateXMLUnmarshaller(model, modelDef, serdeImports);
        }
      } else if (needsXMLDictionaryHelper(model)) {
        generateXMLMarshaller(model, modelDef, serdeImports);
        generateXMLUnmarshaller(model, modelDef, serdeImports);
      }
      modelDefs.push(modelDef);
      continue;
    }
    if (model.kind === 'polymorphicModel') {
      generateDiscriminatorMarkerMethod(model.interface, modelDef);
      for (let parent = model.interface.parent; parent !== undefined; parent = parent.parent) {
        generateDiscriminatorMarkerMethod(parent, modelDef);
      }
    }
    if (model.annotations.multipartFormData) {
      generateToMultipartForm(modelDef);
      modelDef.SerDe.needsJSONPopulateMultipart = true;
    } else if (!model.annotations.omitSerDeMethods) {
      generateJSONMarshaller(model, modelDef, serdeImports);
      generateJSONUnmarshaller(model, modelDef, serdeImports, codeModel.options);
    }
    modelDefs.push(modelDef);
  }
  return modelDefs;
}

function needsXMLDictionaryHelper(modelType: go.Model): boolean {
  for (const field of values(modelType.fields)) {
    // additional properties uses an internal wrapper type with its own serde impl
    if (field.type.kind === 'map' && !field.annotations.isAdditionalProperties) {
      return true;
    }
  }
  return false;
}

function needsXMLArrayMarshalling(modelType: go.Model): boolean {
  for (const prop of values(modelType.fields)) {
    if (prop.type.kind === 'slice') {
      return true;
    }
  }
  return false;
}

// generates discriminator marker method
function generateDiscriminatorMarkerMethod(type: go.Interface, modelDef: ModelDef) {
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
      method += `\t\t${field.name}: ${modelDef.receiverName()}.${field.name},\n`;
    }
    method += '\t}\n}\n\n';
  }
  modelDef.Methods.push({ name: interfaceMethod, desc: `${interfaceMethod} implements the ${type.name} interface for type ${modelDef.Name}.`, text: method });
}

function generateToMultipartForm(modelDef: ModelDef) {
  const receiver = modelDef.receiverName();
  let method = `func (${receiver} ${modelDef.Name}) toMultipartFormData() (map[string]any, error) {\n`;
  method += '\tobjectMap := make(map[string]any)\n';
  for (const field of modelDef.Fields) {
    const fieldType = helpers.recursiveUnwrapMapSlice(field.type);
    let setField: string;
    let star = '';
    if (!field.byValue) {
      star = '*';
    }
    if (fieldType.kind === 'model' && !fieldType.annotations.multipartFormData) {
      setField = `\tif err := populateMultipartJSON(objectMap, "${field.serializedName}", ${star}${receiver}.${field.name}); err != nil {\n\t\treturn nil, err\n\t}\n`;
    } else {
      setField = `\tobjectMap["${field.serializedName}"] = ${star}${receiver}.${field.name}\n`;
    }
    if (!field.byValue) {
      setField = `\tif ${receiver}.${field.name} != nil {\n\t\t${setField}\t}\n`;
    }
    method += setField;
  }
  method += '\treturn objectMap, nil\n}\n\n';
  modelDef.SerDe.methods.push({ name: 'toMultipartFormData', desc: `toMultipartFormData converts ${modelDef.Name} to multipart/form data.`, text: method });
}

function generateJSONMarshaller(modelType: go.Model | go.PolymorphicModel, modelDef: ModelDef, imports: ImportManager) {
  if (modelType.kind === 'model' && modelType.fields.length === 0) {
    // non-discriminated types without content don't need a custom marshaller.
    // there is a case in network where child is allOf base and child has no properties.
    return;
  }
  imports.add('encoding/json');
  const typeName = modelDef.Name;
  const receiver = modelDef.receiverName();
  let marshaller = `func (${receiver} ${typeName}) MarshalJSON() ([]byte, error) {\n`;
  marshaller += '\tobjectMap := make(map[string]any)\n';
  marshaller += generateJSONMarshallerBody(modelType, modelDef, receiver, imports);
  marshaller += '\treturn json.Marshal(objectMap)\n';
  marshaller += '}\n\n';
  modelDef.SerDe.methods.push({ name: 'MarshalJSON', desc: `MarshalJSON implements the json.Marshaller interface for type ${typeName}.`, text: marshaller });
}

function generateJSONMarshallerBody(modelType: go.Model | go.PolymorphicModel, modelDef: ModelDef, receiver: string, imports: ImportManager): string {
  let marshaller = '';
  let addlProps: go.Map | undefined;
  for (const field of values(modelType.fields)) {
    if (field.type.kind === 'map' && field.annotations.isAdditionalProperties) {
      addlProps = field.type;
      continue;
    }
    if (field.annotations.isDiscriminator) {
      if (field.defaultValue) {
        marshaller += `\tobjectMap["${field.serializedName}"] = ${helpers.formatLiteralValue(field.defaultValue, true)}\n`;
      } else {
        // if there's no discriminator value (e.g. Fish in test server), use the field's value.
        // this will enable support for custom types that aren't (yet) described in the swagger.
        marshaller += `\tobjectMap["${field.serializedName}"] = ${receiver}.${field.name}\n`;
      }
    } else if (field.type.kind === 'encodedBytes') {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      marshaller += `\tpopulateByteArray(objectMap, "${field.serializedName}", ${receiver}.${field.name}, func() any {\n`;
      marshaller += `\t\treturn runtime.EncodeByteArray(${receiver}.${field.name}, runtime.Base64${field.type.encoding}Format)\n\t})\n`;
      modelDef.SerDe.needsJSONPopulateByteArray = true;
    } else if (field.type.kind === 'slice' && field.type.elementType.kind === 'encodedBytes') {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      marshaller += `\tpopulateByteArray(objectMap, "${field.serializedName}", ${receiver}.${field.name}, func() any {\n`;
      marshaller += `\t\tencodedValue := make([]string, len(${receiver}.${field.name}))\n`;
      marshaller += `\t\tfor i := 0; i < len(${receiver}.${field.name}); i++ {\n`;
      marshaller += `\t\t\tencodedValue[i] = runtime.EncodeByteArray(${receiver}.${field.name}[i], runtime.Base64${field.type.elementType.encoding}Format)\n\t\t}\n`;
      marshaller += '\t\treturn encodedValue\n\t})\n';
      modelDef.SerDe.needsJSONPopulateByteArray = true;
    } else if (field.type.kind === 'slice' && field.type.elementType.kind === 'time') {
      const source = `${receiver}.${field.name}`;
      let elementPtr = '*';
      if (field.type.elementTypeByValue) {
        elementPtr = '';
      }
      marshaller += `\taux := make([]${elementPtr}${field.type.elementType.format}, len(${source}), len(${source}))\n`;
      marshaller += `\tfor i := 0; i < len(${source}); i++ {\n`;
      marshaller += `\t\taux[i] = (${elementPtr}${field.type.elementType.format})(${source}[i])\n`;
      marshaller += '\t}\n';
      marshaller += `\tpopulate(objectMap, "${field.serializedName}", aux)\n`;
      modelDef.SerDe.needsJSONPopulate = true;
    } else if (field.type.kind === 'literal') {
      const setter = `objectMap["${field.serializedName}"] = ${helpers.formatLiteralValue(field.type, true)}`;
      if (!field.annotations.required) {
        marshaller += `\tif ${receiver}.${field.name} != nil {\n\t\t${setter}\n\t}\n`;
      } else {
        marshaller += `\t${setter}\n`;
      }
    } else if (field.type.kind === 'rawJSON') {
      marshaller += `\tpopulate(objectMap, "${field.serializedName}", json.RawMessage(${receiver}.${field.name}))\n`;
      modelDef.SerDe.needsJSONPopulate = true;
    } else {
      if (field.defaultValue) {
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
        marshaller += `\tif ${receiver}.${field.name} == nil {\n\t\t${receiver}.${field.name} = to.Ptr(${helpers.formatLiteralValue(field.defaultValue, true)})\n\t}\n`;
      }
      let populate = 'populate';
      if (field.type.kind === 'time') {
        populate += capitalize(field.type.format);
        modelDef.SerDe.needsJSONPopulate = true;
      } else if (field.type.kind === 'any') {
        populate += 'Any';
        modelDef.SerDe.needsJSONPopulateAny = true;
      } else {
        modelDef.SerDe.needsJSONPopulate = true;
      }
      if (field.type.kind === 'scalar' && (field.type.type.startsWith('uint') || field.type.type.startsWith('int')) && field.type.encodeAsString) {
        // TODO: need to handle map and slice type with underlying int as string type
        imports.add('strconv');
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
        if (field.type.type.startsWith('uint') && field.type.type !== 'uint64' || field.type.type.startsWith('int') && field.type.type !== 'int64') {
          marshaller += `\t${populate}(objectMap, "${field.serializedName}", to.Ptr(strconv.${field.type.type.startsWith('int') ? 'FormatInt' : 'FormatUint'}(${field.type.type.startsWith('int') ? 'int64' : 'uint64'}(*${receiver}.${field.name}), 10)))\n`;
        } else {
          marshaller += `\t${populate}(objectMap, "${field.serializedName}", to.Ptr(strconv.${field.type.type.startsWith('int') ? 'FormatInt' : 'FormatUint'}(*${receiver}.${field.name}, 10)))\n`;
        }
      } else {
        marshaller += `\t${populate}(objectMap, "${field.serializedName}", ${receiver}.${field.name})\n`;
      }
    }
  }
  if (addlProps) {
    marshaller += `\tif ${receiver}.AdditionalProperties != nil {\n`;
    marshaller += `\t\tfor key, val := range ${receiver}.AdditionalProperties {\n`;
    let assignment = 'val';
    if (addlProps.valueType.kind === 'time') {
      assignment = `(*${addlProps.valueType.format})(val)`;
    }
    marshaller += `\t\t\tobjectMap[key] = ${assignment}\n`;
    marshaller += '\t\t}\n';
    marshaller += '\t}\n';
  }
  return marshaller;
}

function generateJSONUnmarshaller(modelType: go.Model | go.PolymorphicModel, modelDef: ModelDef, imports: ImportManager, options: go.Options) {
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
  unmarshaller += generateJSONUnmarshallerBody(modelType, modelDef, receiver, imports, options);
  unmarshaller += '}\n\n';
  modelDef.SerDe.methods.push({ name: 'UnmarshalJSON', desc: `UnmarshalJSON implements the json.Unmarshaller interface for type ${typeName}.`, text: unmarshaller });
}

function generateJSONUnmarshallerBody(modelType: go.Model | go.PolymorphicModel, modelDef: ModelDef, receiver: string, imports: ImportManager, options: go.Options): string {
  // we almost always need to have an error check when unmarshaling the values.
  // however, fields that are raw JSON don't require any unmarshaling. so, if all
  // of the fields in a type are raw JSON, then the error check isn't necessary
  // and can be elided (the linter complains about it otherwise).
  let needsErrCheck = false;

  const emitAddlProps = function (tab: string, addlProps: go.Map): string {
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
    if (addlProps.valueType.kind === 'time') {
      imports.add('time');
      auxType = addlProps.valueType.format;
      assignment = `(*time.Time)(${assignment})`;
    }
    addlPropsText += `${tab}\t\t\tvar aux ${auxType}\n`;
    addlPropsText += `${tab}\t\t\terr = json.Unmarshal(val, &aux)\n`;
    addlPropsText += `${tab}\t\t\t${receiver}.AdditionalProperties[key] = ${assignment}\n`;
    addlPropsText += `${tab}\t\t}\n`;
    addlPropsText += `${tab}\t\tdelete(rawMsg, key)\n`;
    needsErrCheck = true;
    return addlPropsText;
  };

  const emitSwitchCase = function(): string {
    let unmarshalBody = '';
    let addlProps: go.Map | undefined;
    unmarshalBody += '\t\tswitch key {\n';
    for (const field of values(modelType.fields)) {
      if (field.type.kind === 'map' && field.annotations.isAdditionalProperties) {
        addlProps = field.type;
        continue;
      }
      unmarshalBody += `\t\tcase "${field.serializedName}":\n`;
      if (hasDiscriminatorInterface(field.type)) {
        unmarshalBody += generateDiscriminatorUnmarshaller(field, receiver);
        needsErrCheck = true;
      } else if (field.type.kind === 'time') {
        unmarshalBody += `\t\t\t\terr = unpopulate${capitalize(field.type.format)}(val, "${field.name}", &${receiver}.${field.name})\n`;
        modelDef.SerDe.needsJSONUnpopulate = true;
        needsErrCheck = true;
      } else if (field.type.kind === 'slice' && field.type.elementType.kind === 'time') {
        imports.add('time');
        let elementPtr = '*';
        if (field.type.elementTypeByValue) {
          elementPtr = '';
        }
        unmarshalBody += `\t\t\tvar aux []${elementPtr}${field.type.elementType.format}\n`;
        unmarshalBody += `\t\t\terr = unpopulate(val, "${field.name}", &aux)\n`;
        unmarshalBody += '\t\t\tfor _, au := range aux {\n';
        unmarshalBody += `\t\t\t\t${receiver}.${field.name} = append(${receiver}.${field.name}, (${elementPtr}time.Time)(au))\n`;
        unmarshalBody += '\t\t\t}\n';
        modelDef.SerDe.needsJSONUnpopulate = true;
        needsErrCheck = true;
      } else if (field.type.kind === 'encodedBytes') {
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
        unmarshalBody += '\t\tif val != nil && string(val) != "null" {\n';
        unmarshalBody += `\t\t\t\terr = runtime.DecodeByteArray(string(val), &${receiver}.${field.name}, runtime.Base64${field.type.encoding}Format)\n\t\t}\n`;
        needsErrCheck = true;
      } else if (field.type.kind === 'slice' && field.type.elementType.kind === 'encodedBytes') {
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
        unmarshalBody += '\t\t\tvar encodedValue []string\n';
        unmarshalBody += `\t\t\terr = unpopulate(val, "${field.name}", &encodedValue)\n`;
        unmarshalBody += '\t\t\tif err == nil && len(encodedValue) > 0 {\n';
        unmarshalBody += `\t\t\t\t${receiver}.${field.name} = make([][]byte, len(encodedValue))\n`;
        unmarshalBody += '\t\t\t\tfor i := 0; i < len(encodedValue) && err == nil; i++ {\n';
        unmarshalBody += `\t\t\t\t\terr = runtime.DecodeByteArray(encodedValue[i], &${receiver}.${field.name}[i], runtime.Base64${field.type.elementType.encoding}Format)\n`;
        unmarshalBody += '\t\t\t\t}\n\t\t\t}\n';
        modelDef.SerDe.needsJSONUnpopulate = true;
        needsErrCheck = true;
      } else if (field.type.kind === 'rawJSON') {
        unmarshalBody += '\t\t\tif string(val) != "null" {\n';
        unmarshalBody += `\t\t\t\t${receiver}.${field.name} = val\n\t\t\t}\n`;
      } else if (field.type.kind === 'scalar' && (field.type.type.startsWith('uint') || field.type.type.startsWith('int')) && field.type.encodeAsString) {
        // TODO: need to handle map and slice type with underlying int as string type
        imports.add('strconv');
        unmarshalBody += `\t\t\t\tvar aux string\n`;
        unmarshalBody += `\t\t\t\terr = unpopulate(val, "${field.name}", &aux)\n`;
        unmarshalBody += `\t\t\t\tif err == nil {\n`;
        unmarshalBody += `\t\t\t\t\tvar v ${field.type.type.startsWith('int') ? 'int64' : 'uint64'}\n`;
        unmarshalBody += `\t\t\t\t\tv, err = strconv.${field.type.type.startsWith('int') ? 'ParseInt' : 'ParseUint'}(aux, 10, 0)\n`;
        unmarshalBody += `\t\t\t\t\tif err == nil {\n`;
        if (field.type.type.startsWith('uint') && field.type.type !== 'uint64' || field.type.type.startsWith('int') && field.type.type !== 'int64') {
          unmarshalBody += `\t\t\t\t\t\t${receiver}.${field.name} = to.Ptr(${field.type.type}(v))\n`;
        } else {
          unmarshalBody += `\t\t\t\t\t\t${receiver}.${field.name} = to.Ptr(v)\n`;
        }
        unmarshalBody += '\t\t\t\t\t}\n';
        unmarshalBody += '\t\t\t\t}\n';
        modelDef.SerDe.needsJSONUnpopulate = true;
        needsErrCheck = true;
      } else {
        unmarshalBody += `\t\t\t\terr = unpopulate(val, "${field.name}", &${receiver}.${field.name})\n`;
        modelDef.SerDe.needsJSONUnpopulate = true;
        needsErrCheck = true;
      }
      unmarshalBody += '\t\t\tdelete(rawMsg, key)\n';
    }
    if (addlProps) {
      unmarshalBody += '\t\tdefault:\n';
      unmarshalBody += emitAddlProps('\t', addlProps);
    } else if (options.disallowUnknownFields) {
      unmarshalBody += '\t\tdefault:\n';
      unmarshalBody += `\t\t\terr = fmt.Errorf("unmarshalling type %T, unknown field %q", ${receiver}, key)\n`;
      needsErrCheck = true;
    }
    unmarshalBody += '\t\t}\n';
    return unmarshalBody;
  };

  let unmarshalBody = '\tfor key, val := range rawMsg {\n';

  // emitSwitchCase sets needsErrCheck so we must call it first
  const switchCaseBody = emitSwitchCase();

  if (needsErrCheck) {
    unmarshalBody += '\t\tvar err error\n';
  }
  unmarshalBody += switchCaseBody;
  if (needsErrCheck) {
    unmarshalBody += '\t\tif err != nil {\n';
    unmarshalBody += `\t\t\treturn fmt.Errorf("unmarshalling type %T: %v", ${receiver}, err)\n`;
    unmarshalBody += '\t\t}\n';
  }
  unmarshalBody += '\t}\n'; // end for key, val := range rawMsg
  unmarshalBody += '\treturn nil\n';
  return unmarshalBody;
}

// returns true if item has a discriminator interface.
// recursively called for arrays and dictionaries.
function hasDiscriminatorInterface(item: go.WireType): boolean {
  switch (item.kind) {
    case 'interface':
      return true;
    case 'map':
      return hasDiscriminatorInterface(item.valueType);
    case 'slice':
      return hasDiscriminatorInterface(item.elementType);
    default:
      return false;
  }
}

// returns the text for unmarshalling a discriminated type
function generateDiscriminatorUnmarshaller(field: go.ModelField, receiver: string): string {
  const startingIndentation = '\t\t\t';
  const propertyName = field.name;

  // these are the simple, non-nested cases (e.g. IterfaceType, []InterfaceType, map[string]InterfaceType)
  if (field.type.kind === 'interface') {
    return `${startingIndentation}${receiver}.${propertyName}, err = unmarshal${field.type.name}(val)\n`;
  } else if (field.type.kind === 'slice' && field.type.elementType.kind === 'interface') {
    return `${startingIndentation}${receiver}.${propertyName}, err = unmarshal${field.type.elementType.name}Array(val)\n`;
  } else if (field.type.kind === 'map' && field.type.valueType.kind === 'interface') {
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
  if (field.type.kind === 'slice') {
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
function recursiveGetDiscriminatorTypeName(item: go.WireType, raw: boolean): string {
  // when raw is true, stop recursing at the level before the leaf schema
  if (item.kind === 'slice') {
    if (!raw || item.elementType.kind !== 'interface') {
      return `[]${recursiveGetDiscriminatorTypeName(item.elementType, raw)}`;
    }
  } else if (item.kind === 'map') {
    if (!raw || item.valueType.kind !== 'interface') {
      return `map[string]${recursiveGetDiscriminatorTypeName(item.valueType, raw)}`;
    }
  }
  if (raw) {
    return 'json.RawMessage';
  }
  return go.getTypeDeclaration(item);
}

// recursively constructs the text to populate a nested discriminator
function recursivePopulateDiscriminator(item: go.WireType, receiver: string, rawSrc: string, dest: string, indent: string, nesting: number): string {
  let text = '';
  let interfaceName = '';
  let targetType = '';

  if (item.kind === 'slice') {
    if (item.elementType.kind !== 'interface') {
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
  } else if (item.kind === 'map') {
    if (item.valueType.kind !== 'interface') {
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

function generateXMLMarshaller(modelType: go.Model, modelDef: ModelDef, imports: ImportManager) {
  // only needed for types with time.Time, maps, or where the XML name doesn't match the type name
  const receiver = modelDef.receiverName();
  const desc = `MarshalXML implements the xml.Marshaller interface for type ${modelDef.Name}.`;
  let text = `func (${receiver} ${modelDef.Name}) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {\n`;
  if (modelType.xml?.wrapper) {
    text += `\tstart.Name.Local = "${modelType.xml.wrapper}"\n`;
  }
  text += generateAliasType(modelType, receiver, true);
  for (const field of values(modelDef.Fields)) {
    if (field.type.kind === 'slice') {
      text += `\tif ${receiver}.${field.name} != nil {\n`;
      text += `\t\taux.${field.name} = &${receiver}.${field.name}\n`;
      text += '\t}\n';
    } else if (field.annotations.isAdditionalProperties || field.type.kind === 'map') {
      text += `\taux.${field.name} = (additionalProperties)(${receiver}.${field.name})\n`;
    } else if (field.type.kind === 'encodedBytes') {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      text += `\tif ${receiver}.${field.name} != nil {\n`;
      text += `\t\tencoded${field.name} := runtime.EncodeByteArray(${receiver}.${field.name}, runtime.Base64${field.type.encoding}Format)\n`;
      text += `\t\taux.${field.name} = &encoded${field.name}\n`;
      text += '\t}\n';
    }
  }
  text += '\treturn enc.EncodeElement(aux, start)\n';
  text += '}\n\n';
  modelDef.SerDe.methods.push({ name: 'MarshalXML', desc: desc, text: text });
}

function generateXMLUnmarshaller(modelType: go.Model, modelDef: ModelDef, imports: ImportManager) {
  // non-polymorphic case, must be something with time.Time
  const receiver = modelDef.receiverName();
  const desc = `UnmarshalXML implements the xml.Unmarshaller interface for type ${modelDef.Name}.`;
  let text = `func (${receiver} *${modelDef.Name}) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {\n`;
  text += generateAliasType(modelType, receiver, false);
  text += '\tif err := dec.DecodeElement(aux, &start); err != nil {\n';
  text += '\t\treturn err\n';
  text += '\t}\n';
  for (const field of values(modelDef.Fields)) {
    if (field.type.kind === 'time') {
      text += `\tif aux.${field.name} != nil && !(*time.Time)(aux.${field.name}).IsZero() {\n`;
      text += `\t\t${receiver}.${field.name} = (*time.Time)(aux.${field.name})\n\t}\n`;
    } else if (field.annotations.isAdditionalProperties || field.type.kind === 'map') {
      text += `\t${receiver}.${field.name} = (map[string]*string)(aux.${field.name})\n`;
    } else if (field.type.kind === 'encodedBytes') {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      text += `\tif aux.${field.name} != nil {\n`;
      text += `\t\tif err := runtime.DecodeByteArray(*aux.${field.name}, &${receiver}.${field.name}, runtime.Base64${field.type.encoding}Format); err != nil {\n`;
      text += '\t\t\treturn err\n';
      text += '\t\t}\n';
      text += '\t}\n';
    }
  }
  text += '\treturn nil\n';
  text += '}\n\n';
  modelDef.SerDe.methods.push({ name: 'UnmarshalXML', desc: desc, text: text });
}

// generates an alias type used by custom XML marshaller/unmarshaller
function generateAliasType(modelType: go.Model, receiver: string, forMarshal: boolean): string {
  let text = `\ttype alias ${modelType.name}\n`;
  text += '\taux := &struct {\n';
  text += '\t\t*alias\n';
  for (const field of values(modelType.fields)) {
    const sn = getXMLSerialization(field, false);
    if (field.type.kind === 'time') {
      text += `\t\t${field.name} *${field.type.format} \`xml:"${sn}"\`\n`;
    } else if (field.annotations.isAdditionalProperties || field.type.kind === 'map') {
      text += `\t\t${field.name} additionalProperties \`xml:"${sn}"\`\n`;
    } else if (field.type.kind === 'slice') {
      text += `\t\t${field.name} *${go.getTypeDeclaration(field.type)} \`xml:"${sn}"\`\n`;
    } else if (field.type.kind === 'encodedBytes') {
      text += `\t\t${field.name} *string \`xml:"${sn}"\`\n`;
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
      if (field.type.kind !== 'time') {
        continue;
      }
      text += `\t\t${field.name}: (*${field.type.format})(${receiver}.${field.name}),\n`;
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

class SerDeInfo {
  methods: Array<ModelMethod>;
  needsJSONPopulate: boolean;
  needsJSONUnpopulate: boolean;
  needsJSONPopulateByteArray: boolean;
  needsJSONPopulateAny: boolean;
  needsJSONPopulateMultipart: boolean;

  constructor() {
    this.methods = new Array<ModelMethod>();
    this.needsJSONPopulate = false;
    this.needsJSONUnpopulate = false;
    this.needsJSONPopulateByteArray = false;
    this.needsJSONPopulateAny = false;
    this.needsJSONPopulateMultipart = false;
  }
}

// represents model definition as a Go struct
class ModelDef {
  readonly Name: string;
  readonly Format: helpers.SerDeFormat;
  readonly Docs: go.Docs;
  readonly Fields: Array<go.ModelField>;
  readonly SerDe: SerDeInfo;
  readonly Methods: Array<ModelMethod>;

  constructor(name: string, format: helpers.SerDeFormat, fields: Array<go.ModelField>, docs: go.Docs) {
    this.Name = name;
    this.Format = format;
    this.Docs = docs;
    this.Fields = fields;
    this.SerDe = new SerDeInfo();
    this.Methods = new Array<ModelMethod>();
  }

  text(): string {
    let text = helpers.formatDocComment(this.Docs);
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
      if (field.docs.summary || field.docs.description) {
        if (!first) {
          // add an extra new-line between fields IFF the field
          // has a comment and it's not the very first one.
          text += '\n';
        }
        text += helpers.formatDocComment(field.docs);
      }
      let typeName = go.getTypeDeclaration(field.type);
      if (field.type.kind === 'literal') {
        // for constants we use the underlying type name
        typeName = go.getLiteralTypeDeclaration(field.type.type);
      }
      let serialization = field.serializedName;
      if (this.Format === 'JSON') {
        serialization += ',omitempty';
      } else if (this.Format === 'XML') {
        serialization = getXMLSerialization(field, false);
      }
      let tag = '';
      // only emit tags for XML; JSON uses custom marshallers/unmarshallers
      if (this.Format === 'XML' && !field.annotations.isAdditionalProperties) {
        tag = ` \`xml:"${serialization}"\``;
      }
      text += `\t${field.name} ${helpers.star(field.byValue)}${typeName}${tag}\n`;
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
  } else if (field.type.kind === 'slice') {
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
