/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../../codemodel.go/src/index.js';
import * as helpers from './helpers.js';
import { ImportManager } from './imports.js';

export interface ModelsSerDe {
  models: string;
  serDe: string;
}

/**
 * Creates the content for the models.go file.
 *
 * @param pkg contains the package content
 * @param options the emitter options
 * @returns the text for the files or the empty string
 */
export function generateModels(pkg: go.PackageContent, options: go.Options): ModelsSerDe {
  if (pkg.models.length === 0) {
    return {
      models: '',
      serDe: '',
    };
  }

  // this list of packages to import
  const modelImports = new ImportManager(pkg);
  const serdeImports = new ImportManager(pkg);
  let modelText = helpers.contentPreamble(pkg);

  const indent = new helpers.Indentation();

  // we do model generation first as it can add imports to the imports list
  const modelDefs = generateModelDefs(modelImports, serdeImports, pkg, options, indent);

  modelText += modelImports.text();

  // structs
  let needsJSONPopulate = false;
  let needsJSONPopulateTime = false;
  let needsJSONUnpopulate = false;
  let needsJSONUnpopulateTime = false;
  let needsJSONPopulateByteArray = false;
  let needsJSONPopulateAny = false;
  let needsJSONPopulateMultipart = false;
  let serdeTextBody = '';
  for (const modelDef of modelDefs) {
    modelText += modelDef.text(indent);

    modelDef.Methods.sort((a: ModelMethod, b: ModelMethod) => {
      return helpers.sortAscending(a.name, b.name);
    });
    for (const method of modelDef.Methods) {
      if (method.desc.length > 0) {
        modelText += `${helpers.comment(method.desc, '// ', undefined, helpers.commentLength)}\n`;
      }
      modelText += method.text;
    }

    modelDef.SerDe.methods.sort((a: ModelMethod, b: ModelMethod) => {
      return helpers.sortAscending(a.name, b.name);
    });
    for (const method of modelDef.SerDe.methods) {
      if (method.desc.length > 0) {
        serdeTextBody += `${helpers.comment(method.desc, '// ', undefined, helpers.commentLength)}\n`;
      }
      serdeTextBody += method.text;
    }
    if (modelDef.SerDe.needsJSONPopulate) {
      needsJSONPopulate = true;
    }
    if (modelDef.SerDe.needsJSONPopulateTime) {
      needsJSONPopulateTime = true;
    }
    if (modelDef.SerDe.needsJSONUnpopulate) {
      needsJSONUnpopulate = true;
    }
    if (modelDef.SerDe.needsJSONUnpopulateTime) {
      needsJSONUnpopulateTime = true;
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
    serdeTextBody += `${indent.get()}if v == nil {\n`;
    serdeTextBody += `${indent.push().get()}return\n`;
    serdeTextBody += `${indent.pop().get()}} else if azcore.IsNullValue(v) {\n`;
    serdeTextBody += `${indent.push().get()}m[k] = nil\n`;
    serdeTextBody += `${indent.pop().get()}} else if !reflect.ValueOf(v).IsNil() {\n`;
    serdeTextBody += `${indent.push().get()}m[k] = v\n`;
    serdeTextBody += `${indent.pop().get()}}\n`;
    serdeTextBody += '}\n\n';
  }
  if (needsJSONPopulateTime) {
    serdeImports.add('time');
    serdeImports.add('reflect');
    serdeImports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    serdeTextBody += 'func populateTime[T dateTimeConstraints](m map[string]any, k string, t *time.Time) {\n';
    serdeTextBody += `${indent.get()}if t == nil {\n`;
    serdeTextBody += `${indent.push().get()}return\n`;
    serdeTextBody += `${indent.pop().get()}} else if azcore.IsNullValue(t) {\n`;
    serdeTextBody += `${indent.push().get()}m[k] = nil\n`;
    serdeTextBody += `${indent.pop().get()}} else if !reflect.ValueOf(t).IsNil() {\n`;
    indent.push();
    serdeTextBody += `${indent.get()}newTime := T(*t)\n`;
    serdeTextBody += `${indent.get()}m[k] = (*T)(&newTime)\n`;
    indent.pop();
    serdeTextBody += `${indent.get()}}\n`;
    serdeTextBody += '}\n\n';
  }
  if (needsJSONPopulateAny) {
    serdeImports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    serdeTextBody += 'func populateAny(m map[string]any, k string, v any) {\n';
    serdeTextBody += `${indent.get()}if v == nil {\n`;
    serdeTextBody += `${indent.push().get()}return\n`;
    serdeTextBody += `${indent.pop().get()}} else if azcore.IsNullValue(v) {\n`;
    serdeTextBody += `${indent.push().get()}m[k] = nil\n`;
    serdeTextBody += `${indent.pop().get()}} else {\n`;
    serdeTextBody += `${indent.push().get()}m[k] = v\n`;
    serdeTextBody += `${indent.pop().get()}}\n`;
    serdeTextBody += '}\n\n';
  }
  if (needsJSONPopulateByteArray) {
    serdeImports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    serdeImports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
    serdeTextBody += 'func populateByteArray[T any](m map[string]any, k string, b []T, convert func() any) {\n';
    serdeTextBody += `${indent.get()}if azcore.IsNullValue(b) {\n`;
    serdeTextBody += `${indent.push().get()}m[k] = nil\n`;
    serdeTextBody += `${indent.pop().get()}} else if len(b) == 0 {\n`;
    serdeTextBody += `${indent.push().get()}return\n`;
    serdeTextBody += `${indent.pop().get()}} else {\n`;
    serdeTextBody += `${indent.push().get()}m[k] = convert()\n`;
    serdeTextBody += `${indent.pop().get()}}\n`;
    serdeTextBody += '}\n\n';
  }
  if (needsJSONUnpopulate) {
    serdeImports.add('fmt');
    serdeTextBody += 'func unpopulate(data json.RawMessage, fn string, v any) error {\n';
    serdeTextBody += `${indent.get()}if data == nil || string(data) == "null" {\n`;
    serdeTextBody += `${indent.push().get()}return nil\n`;
    serdeTextBody += `${indent.pop().get()}}\n`;
    serdeTextBody += `${indent.get()}if err := json.Unmarshal(data, v); err != nil {\n`;
    serdeTextBody += `${indent.push().get()}return fmt.Errorf("struct field %s: %v", fn, err)\n`;
    serdeTextBody += `${indent.pop().get()}}\n`;
    serdeTextBody += `${indent.get()}return nil\n`;
    serdeTextBody += '}\n\n';
  }
  if (needsJSONUnpopulateTime) {
    serdeImports.add('fmt');
    serdeImports.add('time');
    serdeTextBody += 'func unpopulateTime[T dateTimeConstraints](data json.RawMessage, fn string, t **time.Time) error {\n';
    serdeTextBody += `${indent.get()}if data == nil || string(data) == "null" {\n`;
    serdeTextBody += `${indent.push().get()}return nil\n`;
    serdeTextBody += `${indent.pop().get()}}\n`;
    serdeTextBody += `${indent.get()}var aux T\n`;
    serdeTextBody += `${indent.get()}if err := json.Unmarshal(data, &aux); err != nil {\n`;
    serdeTextBody += `${indent.push().get()}return fmt.Errorf("struct field %s: %v", fn, err)\n`;
    serdeTextBody += `${indent.pop().get()}}\n`;
    serdeTextBody += `${indent.get()}newTime := time.Time(aux)\n`;
    serdeTextBody += `${indent.get()}*t = &newTime\n`;
    serdeTextBody += `${indent.get()}return nil\n`;
    serdeTextBody += '}\n\n';
  }
  if (needsJSONPopulateMultipart) {
    serdeImports.add('encoding/json');
    serdeTextBody += 'func populateMultipartJSON(m map[string]any, k string, v any) error {\n';
    serdeTextBody += `${indent.get()}data, err := json.Marshal(v)\n`;
    serdeTextBody += `${indent.get()}if err != nil {\n`;
    serdeTextBody += `${indent.push().get()}return err\n`;
    serdeTextBody += `${indent.pop().get()}}\n`;
    serdeTextBody += `${indent.get()}m[k] = data\n`;
    serdeTextBody += `${indent.get()}return nil\n`;
    serdeTextBody += '}\n\n';
  }
  if (needsJSONUnpopulateTime || needsJSONPopulateTime) {
    serdeTextBody += `type dateTimeConstraints interface {\n`;
    serdeTextBody += `${indent.get()}datetime.PlainDate | datetime.PlainTime | datetime.RFC1123 | datetime.RFC3339 | datetime.Unix\n`;
    serdeTextBody += '}\n\n';
  }
  let serdeText = '';
  if (serdeTextBody.length > 0) {
    serdeText = helpers.contentPreamble(pkg);
    serdeText += serdeImports.text();
    serdeText += serdeTextBody;
  }
  return {
    models: modelText,
    serDe: serdeText,
  };
}

/**
 * converts model types to an array of ModelDef types
 *
 * @param modelImports the import manager for the models file
 * @param serdeImports the import manager for the models_serde file
 * @param pkg contains the package content
 * @param options the Go emitter options
 * @returns an array of ModelDefs
 */
function generateModelDefs(modelImports: ImportManager, serdeImports: ImportManager, pkg: go.PackageContent, options: go.Options, indent: helpers.Indentation): Array<ModelDef> {
  const models = pkg.models;
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

    const serDeFormat = helpers.getSerDeFormat(model, pkg);
    const modelDef = new ModelDef(model, serDeFormat);
    for (const field of modelDef.Model.fields) {
      modelImports.addForType(field.type);
    }

    if (model.kind === 'model' && serDeFormat === 'XML' && !model.annotations.omitSerDeMethods) {
      serdeImports.add('encoding/xml');
      let needsDateTimeMarshalling = false;
      let byteArrayFormat = false;
      for (const field of model.fields) {
        if (field.type.kind !== 'etag') {
          // azcore.ETag un/marshals on its own so no need to
          // import azcore as we don't explicitly reference the type
          serdeImports.addForType(field.type);
        }
        if (field.type.kind === 'time') {
          needsDateTimeMarshalling = true;
        } else if (field.type.kind === 'encodedBytes') {
          byteArrayFormat = true;
        }
      }
      // due to differences in XML marshallers/unmarshallers, we use different codegen than for JSON
      const needsXMLDictionaryUnmarshalling = needsXMLDictionaryHelper(model);
      if (needsDateTimeMarshalling || model.xml?.name || needsXMLArrayMarshalling(model) || byteArrayFormat) {
        generateXMLMarshaller(modelDef, serdeImports, indent);
        if (needsDateTimeMarshalling || needsXMLDictionaryUnmarshalling || byteArrayFormat) {
          generateXMLUnmarshaller(modelDef, serdeImports, indent);
        }
      } else if (needsXMLDictionaryUnmarshalling) {
        generateXMLMarshaller(modelDef, serdeImports, indent);
        generateXMLUnmarshaller(modelDef, serdeImports, indent);
      }
      modelDefs.push(modelDef);
      continue;
    }
    if (model.kind === 'polymorphicModel') {
      generateDiscriminatorMarkerMethod(model.interface, modelDef, indent);
      for (let parent = model.interface.parent; parent !== undefined; parent = parent.parent) {
        generateDiscriminatorMarkerMethod(parent, modelDef, indent);
      }
    }
    if (model.annotations.multipartFormData) {
      generateToMultipartForm(modelDef, indent);
      modelDef.SerDe.needsJSONPopulateMultipart = true;
    } else if (!model.annotations.omitSerDeMethods) {
      generateJSONMarshaller(modelDef, serdeImports, indent);
      generateJSONUnmarshaller(modelDef, options, serdeImports, indent);
    }
    modelDefs.push(modelDef);
  }
  return modelDefs;
}

function needsXMLDictionaryHelper(modelType: go.Model): boolean {
  for (const field of modelType.fields) {
    // additional properties uses an internal wrapper type with its own serde impl
    if (field.type.kind === 'map' && !field.annotations.isAdditionalProperties) {
      return true;
    }
  }
  return false;
}

/**
 * returns true if the model contains one or more slices,
 * indicating that a custom marshaler is required.
 * we use a custom marshaler to omit the XML tags for empty
 * slices (the default marshaler will include them).
 *
 * @param modelType the model to check for slice fields
 * @returns true if a custom marshaler is required
 */
function needsXMLArrayMarshalling(modelType: go.Model): boolean {
  for (const prop of modelType.fields) {
    if (prop.type.kind === 'slice') {
      return true;
    }
  }
  return false;
}

// generates discriminator marker method
function generateDiscriminatorMarkerMethod(type: go.Interface, modelDef: ModelDef, indent: helpers.Indentation) {
  const typeName = type.rootType.name;
  const receiver = modelDef.receiverName();
  const interfaceMethod = `Get${typeName}`;
  let method = `func (${receiver} *${modelDef.Model.name}) ${interfaceMethod}() *${typeName} {`;
  if (type.rootType.name === modelDef.Model.name) {
    // the marker method is on the discriminator itself, so just return the receiver
    method += ` return ${receiver} }\n\n`;
  } else {
    // the marker method is on a child type, so return an instance of the parent
    // type by copying the parent values into a new instance.
    method += `\n${indent.get()}return &${type.rootType.name}{\n`;
    indent.push();
    for (const field of type.rootType.fields) {
      method += `${indent.get()}${field.name}: ${modelDef.receiverName()}.${field.name},\n`;
    }
    method += `${indent.pop().get()}}\n}\n\n`;
  }
  modelDef.Methods.push({ name: interfaceMethod, desc: `${interfaceMethod} implements the ${type.name} interface for type ${modelDef.Model.name}.`, text: method });
}

function generateToMultipartForm(modelDef: ModelDef, indent: helpers.Indentation) {
  const receiver = modelDef.receiverName();
  let method = `func (${receiver} ${modelDef.Model.name}) toMultipartFormData() (map[string]any, error) {\n`;
  method += `${indent.get()}objectMap := make(map[string]any)\n`;
  for (const field of modelDef.Model.fields) {
    const fieldType = helpers.recursiveUnwrapMapSlice(field.type);
    let star = '';
    if (!field.byValue) {
      star = '*';
    }
    if (!field.byValue) {
      method += `${indent.get()}if ${receiver}.${field.name} != nil {\n`;
      indent.push();
    }
    if (fieldType.kind === 'model' && !fieldType.annotations.multipartFormData) {
      method += `${indent.get()}if err := populateMultipartJSON(objectMap, "${field.serializedName}", ${star}${receiver}.${field.name}); err != nil {\n`;
      method += `${indent.push().get()}return nil, err\n`;
      method += `${indent.pop().get()}}\n`;
    } else {
      method += `${indent.get()}objectMap["${field.serializedName}"] = ${star}${receiver}.${field.name}\n`;
    }
    if (!field.byValue) {
      indent.pop();
      method += `${indent.get()}}\n`;
    }
  }
  method += `${indent.get()}return objectMap, nil\n}\n\n`;
  modelDef.SerDe.methods.push({ name: 'toMultipartFormData', desc: `toMultipartFormData converts ${modelDef.Model.name} to multipart/form data.`, text: method });
}

/**
 * generates the MarshalJSON method for the provided type.
 * the method impl is added to modelDef.SerDe.methods.
 *
 * @param modelDef the type for which to emit the method
 * @param imports the import manager currently in scope
 * @param indent the indentation helper currently in scope
 */
function generateJSONMarshaller(modelDef: ModelDef, imports: ImportManager, indent: helpers.Indentation): void {
  if (modelDef.Model.kind === 'model' && modelDef.Model.fields.length === 0) {
    // non-discriminated types without content don't need a custom marshaller.
    // there is a case in network where child is allOf base and child has no properties.
    return;
  }
  imports.add('encoding/json');
  const typeName = modelDef.Model.name;
  const receiver = modelDef.receiverName();
  let marshaller = `func (${receiver} ${typeName}) MarshalJSON() ([]byte, error) {\n`;
  marshaller += `${indent.get()}objectMap := make(map[string]any)\n`;
  marshaller += generateJSONMarshallerBody(modelDef, receiver, imports, indent);
  marshaller += `${indent.get()}return json.Marshal(objectMap)\n`;
  marshaller += '}\n\n';
  modelDef.SerDe.methods.push({ name: 'MarshalJSON', desc: `MarshalJSON implements the json.Marshaller interface for type ${typeName}.`, text: marshaller });
}

/**
 * generates the contents of MarshalJSON that encode the target type
 *
 * @param modelDef the type being encoded
 * @param receiver the name of the receiver in the MarshalJSON method
 * @param imports the import manager currently in scope
 * @param indent the indentation helper currently in scope
 * @returns the text for encoding the target type
 */
function generateJSONMarshallerBody(modelDef: ModelDef, receiver: string, imports: ImportManager, indent: helpers.Indentation): string {
  let marshaller = '';
  let addlProps: go.Map | undefined;
  for (const field of modelDef.Model.fields) {
    if (field.type.kind === 'map' && field.annotations.isAdditionalProperties) {
      addlProps = field.type;
      continue;
    }
    if (field.annotations.isDiscriminator) {
      if (field.defaultValue) {
        marshaller += `${indent.get()}objectMap["${field.serializedName}"] = ${helpers.formatLiteralValue(field.defaultValue, true)}\n`;
      } else {
        // if there's no discriminator value (e.g. Fish in test server), use the field's value.
        // this will enable support for custom types that aren't (yet) described in the swagger.
        marshaller += `${indent.get()}objectMap["${field.serializedName}"] = ${receiver}.${field.name}\n`;
      }
    } else if (field.type.kind === 'encodedBytes') {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      marshaller += `${indent.get()}populateByteArray(objectMap, "${field.serializedName}", ${receiver}.${field.name}, func() any {\n`;
      marshaller += `${indent.push().get()}return runtime.EncodeByteArray(${receiver}.${field.name}, runtime.Base64${field.type.encoding}Format)\n`;
      marshaller += `${indent.pop().get()}})\n`;
      modelDef.SerDe.needsJSONPopulateByteArray = true;
    } else if (field.type.kind === 'slice' && field.type.elementType.kind === 'encodedBytes') {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      marshaller += `${indent.get()}populateByteArray(objectMap, "${field.serializedName}", ${receiver}.${field.name}, func() any {\n`;
      marshaller += `${indent.push().get()}encodedValue := make([]string, len(${receiver}.${field.name}))\n`;
      marshaller += `${indent.get()}for i := 0; i < len(${receiver}.${field.name}); i++ {\n`;
      marshaller += `${indent.push().get()}encodedValue[i] = runtime.EncodeByteArray(${receiver}.${field.name}[i], runtime.Base64${field.type.elementType.encoding}Format)\n`;
      marshaller += `${indent.pop().get()}}\n`;
      marshaller += `${indent.get()}return encodedValue\n`;
      marshaller += `${indent.pop().get()}})\n`;
      modelDef.SerDe.needsJSONPopulateByteArray = true;
    } else if (field.type.kind === 'slice' && field.type.elementType.kind === 'time') {
      const source = `${receiver}.${field.name}`;
      let elementPtr = '*';
      if (field.type.elementTypeByValue) {
        elementPtr = '';
      }
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/datetime');
      marshaller += `${indent.get()}aux := make([]${elementPtr}datetime.${field.type.elementType.format}, len(${source}), len(${source}))\n`;
      marshaller += `${indent.get()}for i := 0; i < len(${source}); i++ {\n`;
      marshaller += `${indent.push().get()}aux[i] = (${elementPtr}datetime.${field.type.elementType.format})(${source}[i])\n`;
      marshaller += `${indent.pop().get()}}\n`;
      marshaller += `${indent.get()}populate(objectMap, "${field.serializedName}", aux)\n`;
      modelDef.SerDe.needsJSONPopulate = true;
    } else if (field.type.kind === 'literal') {
      const setter = `objectMap["${field.serializedName}"] = ${helpers.formatLiteralValue(field.type, true)}`;
      if (!field.annotations.required) {
        marshaller += `${indent.get()}if ${receiver}.${field.name} != nil {\n`;
        marshaller += `${indent.push().get()}${setter}\n`;
        marshaller += `${indent.pop().get()}}\n`;
      } else {
        marshaller += `${indent.get()}${setter}\n`;
      }
    } else if (field.type.kind === 'rawJSON') {
      marshaller += `${indent.get()}populate(objectMap, "${field.serializedName}", json.RawMessage(${receiver}.${field.name}))\n`;
      modelDef.SerDe.needsJSONPopulate = true;
    } else {
      if (field.defaultValue) {
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
        marshaller += `${indent.get()}if ${receiver}.${field.name} == nil {\n`;
        marshaller += `${indent.push().get()}${receiver}.${field.name} = to.Ptr(${helpers.formatLiteralValue(field.defaultValue, true)})\n`;
        marshaller += `${indent.pop().get()}}\n`;
      }
      let populate: string;
      if (field.type.kind === 'time') {
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/datetime');
        populate = `populateTime[datetime.${field.type.format}]`;
        modelDef.SerDe.needsJSONPopulateTime = true;
      } else if (field.type.kind === 'any') {
        populate = 'populateAny';
        modelDef.SerDe.needsJSONPopulateAny = true;
      } else {
        populate = 'populate';
        modelDef.SerDe.needsJSONPopulate = true;
      }
      if (field.type.kind === 'scalar' && (field.type.type.startsWith('uint') || field.type.type.startsWith('int')) && field.type.encodeAsString) {
        // TODO: need to handle map and slice type with underlying int as string type
        imports.add('strconv');
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
        if ((field.type.type.startsWith('uint') && field.type.type !== 'uint64') || (field.type.type.startsWith('int') && field.type.type !== 'int64')) {
          marshaller += `${indent.get()}${populate}(objectMap, "${field.serializedName}", to.Ptr(strconv.${field.type.type.startsWith('int') ? 'FormatInt' : 'FormatUint'}(${field.type.type.startsWith('int') ? 'int64' : 'uint64'}(*${receiver}.${field.name}), 10)))\n`;
        } else {
          marshaller += `${indent.get()}${populate}(objectMap, "${field.serializedName}", to.Ptr(strconv.${field.type.type.startsWith('int') ? 'FormatInt' : 'FormatUint'}(*${receiver}.${field.name}, 10)))\n`;
        }
      } else {
        marshaller += `${indent.get()}${populate}(objectMap, "${field.serializedName}", ${receiver}.${field.name})\n`;
      }
    }
  }
  if (addlProps) {
    marshaller += `${indent.get()}if ${receiver}.AdditionalProperties != nil {\n`;
    marshaller += `${indent.push().get()}for key, val := range ${receiver}.AdditionalProperties {\n`;
    let assignment = 'val';
    if (addlProps.valueType.kind === 'time') {
      assignment = `(*${addlProps.valueType.format})(val)`;
    }
    marshaller += `${indent.push().get()}objectMap[key] = ${assignment}\n`;
    marshaller += `${indent.pop().get()}}\n`;
    marshaller += `${indent.pop().get()}}\n`;
  }
  return marshaller;
}

/**
 * generates the UnmarshalJSON method for the provided type.
 * the method impl is added to modelDef.SerDe.methods.
 *
 * @param modelDef the type for which to emit the method
 * @param options the Go emitter options
 * @param imports the import manager currently in scope
 * @param indent the indentation helper currently in scope
 */
function generateJSONUnmarshaller(modelDef: ModelDef, options: go.Options, imports: ImportManager, indent: helpers.Indentation): void {
  // there's a corner-case where a derived type might not add any new fields (Cookiecuttershark).
  // in this case skip adding the unmarshaller as it's not necessary and doesn't compile.
  if (modelDef.Model.fields.length === 0) {
    return;
  }
  imports.add('encoding/json');
  imports.add('fmt');
  const typeName = modelDef.Model.name;
  const receiver = modelDef.receiverName();
  let unmarshaller = `func (${receiver} *${typeName}) UnmarshalJSON(data []byte) error {\n`;
  unmarshaller += `${indent.get()}var rawMsg map[string]json.RawMessage\n`;
  unmarshaller += `${indent.get()}if err := json.Unmarshal(data, &rawMsg); err != nil {\n`;
  unmarshaller += `${indent.push().get()}return fmt.Errorf("unmarshalling type %T: %v", ${receiver}, err)\n`;
  unmarshaller += `${indent.pop().get()}}\n`;
  unmarshaller += generateJSONUnmarshallerBody(modelDef, receiver, options, imports, indent);
  unmarshaller += '}\n\n';
  modelDef.SerDe.methods.push({ name: 'UnmarshalJSON', desc: `UnmarshalJSON implements the json.Unmarshaller interface for type ${typeName}.`, text: unmarshaller });
}

/**
 * generates the contents of UnmarshalJSON that decode the target type
 *
 * @param modelDef the type being decoded
 * @param receiver the receiver for the UnmarshalJSON method
 * @param options the Go emitter options
 * @param imports the import manager currently in scope
 * @param indent the indentation helper currently in scope
 * @returns the text for decoding the target type
 */
function generateJSONUnmarshallerBody(modelDef: ModelDef, receiver: string, options: go.Options, imports: ImportManager, indent: helpers.Indentation): string {
  // we almost always need to have an error check when unmarshaling the values.
  // however, fields that are raw JSON don't require any unmarshaling. so, if all
  // of the fields in a type are raw JSON, then the error check isn't necessary
  // and can be elided (the linter complains about it otherwise).
  let needsErrCheck = false;

  const emitAddlProps = function (addlProps: go.Map): string {
    // indent is at the case body level when called
    let addlPropsText = `${indent.get()}if ${receiver}.AdditionalProperties == nil {\n`;
    let ref = '';
    if (!addlProps.valueTypeByValue) {
      ref = '&';
    }
    addlPropsText += `${indent.push().get()}${receiver}.AdditionalProperties = ${go.getTypeDeclaration(addlProps, modelDef.Model.pkg)}{}\n`;
    addlPropsText += `${indent.pop().get()}}\n`;
    addlPropsText += `${indent.get()}if val != nil {\n`;
    let auxType = go.getTypeDeclaration(addlProps.valueType, modelDef.Model.pkg);
    let assignment = `${ref}aux`;
    if (addlProps.valueType.kind === 'time') {
      imports.add('time');
      auxType = addlProps.valueType.format;
      assignment = `(*time.Time)(${assignment})`;
    }
    addlPropsText += `${indent.push().get()}var aux ${auxType}\n`;
    addlPropsText += `${indent.get()}err = json.Unmarshal(val, &aux)\n`;
    addlPropsText += `${indent.get()}${receiver}.AdditionalProperties[key] = ${assignment}\n`;
    addlPropsText += `${indent.pop().get()}}\n`;
    addlPropsText += `${indent.get()}delete(rawMsg, key)\n`;
    needsErrCheck = true;
    return addlPropsText;
  };

  const emitSwitchCase = function (): string {
    // indent is at level 2 (for loop body) when called
    let unmarshalBody = '';
    let addlProps: go.Map | undefined;
    unmarshalBody += `${indent.get()}switch key {\n`;
    for (const field of modelDef.Model.fields) {
      if (field.type.kind === 'map' && field.annotations.isAdditionalProperties) {
        addlProps = field.type;
        continue;
      }
      unmarshalBody += `${indent.get()}case "${field.serializedName}":\n`;
      indent.push(); // case body level
      if (hasDiscriminatorInterface(field.type)) {
        unmarshalBody += generateDiscriminatorUnmarshaller(modelDef.Model, field, receiver, indent);
        needsErrCheck = true;
      } else if (field.type.kind === 'time') {
        unmarshalBody += `${indent.get()}err = unpopulateTime[datetime.${field.type.format}](val, "${field.name}", &${receiver}.${field.name})\n`;
        modelDef.SerDe.needsJSONUnpopulateTime = true;
        needsErrCheck = true;
      } else if (field.type.kind === 'slice' && field.type.elementType.kind === 'time') {
        imports.add('time');
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/datetime');
        let elementPtr = '*';
        if (field.type.elementTypeByValue) {
          elementPtr = '';
        }
        unmarshalBody += `${indent.get()}var aux []${elementPtr}datetime.${field.type.elementType.format}\n`;
        unmarshalBody += `${indent.get()}err = unpopulate(val, "${field.name}", &aux)\n`;
        unmarshalBody += `${indent.get()}for _, au := range aux {\n`;
        unmarshalBody += `${indent.push().get()}${receiver}.${field.name} = append(${receiver}.${field.name}, (${elementPtr}time.Time)(au))\n`;
        unmarshalBody += `${indent.pop().get()}}\n`;
        modelDef.SerDe.needsJSONUnpopulate = true;
        needsErrCheck = true;
      } else if (field.type.kind === 'encodedBytes') {
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
        unmarshalBody += `${indent.get()}if val != nil && string(val) != "null" {\n`;
        unmarshalBody += `${indent.push().get()}err = runtime.DecodeByteArray(string(val), &${receiver}.${field.name}, runtime.Base64${field.type.encoding}Format)\n`;
        unmarshalBody += `${indent.pop().get()}}\n`;
        needsErrCheck = true;
      } else if (field.type.kind === 'slice' && field.type.elementType.kind === 'encodedBytes') {
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
        unmarshalBody += `${indent.get()}var encodedValue []string\n`;
        unmarshalBody += `${indent.get()}err = unpopulate(val, "${field.name}", &encodedValue)\n`;
        unmarshalBody += `${indent.get()}if err == nil && len(encodedValue) > 0 {\n`;
        indent.push();
        unmarshalBody += `${indent.get()}${receiver}.${field.name} = make([][]byte, len(encodedValue))\n`;
        unmarshalBody += `${indent.get()}for i := 0; i < len(encodedValue) && err == nil; i++ {\n`;
        unmarshalBody += `${indent.push().get()}err = runtime.DecodeByteArray(encodedValue[i], &${receiver}.${field.name}[i], runtime.Base64${field.type.elementType.encoding}Format)\n`;
        unmarshalBody += `${indent.pop().get()}}\n`;
        indent.pop();
        unmarshalBody += `${indent.get()}}\n`;
        modelDef.SerDe.needsJSONUnpopulate = true;
        needsErrCheck = true;
      } else if (field.type.kind === 'rawJSON') {
        unmarshalBody += `${indent.get()}if string(val) != "null" {\n`;
        unmarshalBody += `${indent.push().get()}${receiver}.${field.name} = val\n`;
        unmarshalBody += `${indent.pop().get()}}\n`;
      } else if (field.type.kind === 'scalar' && (field.type.type.startsWith('uint') || field.type.type.startsWith('int')) && field.type.encodeAsString) {
        // TODO: need to handle map and slice type with underlying int as string type
        imports.add('strconv');
        unmarshalBody += `${indent.get()}var aux string\n`;
        unmarshalBody += `${indent.get()}err = unpopulate(val, "${field.name}", &aux)\n`;
        unmarshalBody += `${indent.get()}if err == nil {\n`;
        indent.push();
        unmarshalBody += `${indent.get()}var v ${field.type.type.startsWith('int') ? 'int64' : 'uint64'}\n`;
        unmarshalBody += `${indent.get()}v, err = strconv.${field.type.type.startsWith('int') ? 'ParseInt' : 'ParseUint'}(aux, 10, 0)\n`;
        unmarshalBody += `${indent.get()}if err == nil {\n`;
        if ((field.type.type.startsWith('uint') && field.type.type !== 'uint64') || (field.type.type.startsWith('int') && field.type.type !== 'int64')) {
          unmarshalBody += `${indent.push().get()}${receiver}.${field.name} = to.Ptr(${field.type.type}(v))\n`;
        } else {
          unmarshalBody += `${indent.push().get()}${receiver}.${field.name} = to.Ptr(v)\n`;
        }
        unmarshalBody += `${indent.pop().get()}}\n`;
        indent.pop();
        unmarshalBody += `${indent.get()}}\n`;
        modelDef.SerDe.needsJSONUnpopulate = true;
        needsErrCheck = true;
      } else {
        const unpopulateField = `${indent.get()}err = unpopulate(val, "${field.name}", &${receiver}.${field.name})\n`;
        if (field.type.kind === 'string' && field.annotations.unmarshalEmptyStringAsNil) {
          unmarshalBody += `${indent.get()}${helpers.buildIfBlock(indent, {
            condition: `string(val) != \`""\``,
            body: (indent) => `${indent.get()}${unpopulateField}`,
          })}\n`;
        } else {
          unmarshalBody += unpopulateField;
        }
        modelDef.SerDe.needsJSONUnpopulate = true;
        needsErrCheck = true;
      }
      unmarshalBody += `${indent.get()}delete(rawMsg, key)\n`;
      indent.pop(); // back to switch level
    }
    if (addlProps) {
      unmarshalBody += `${indent.get()}default:\n`;
      indent.push(); // case body level
      unmarshalBody += emitAddlProps(addlProps);
      indent.pop();
    } else if (options.disallowUnknownFields) {
      unmarshalBody += `${indent.get()}default:\n`;
      unmarshalBody += `${indent.push().get()}err = fmt.Errorf("unmarshalling type %T, unknown field %q", ${receiver}, key)\n`;
      indent.pop();
      needsErrCheck = true;
    }
    unmarshalBody += `${indent.get()}}\n`;
    return unmarshalBody;
  };

  let unmarshalBody = `${indent.get()}for key, val := range rawMsg {\n`;
  indent.push(); // level 2 (for loop body)

  // emitSwitchCase sets needsErrCheck so we must call it first
  const switchCaseBody = emitSwitchCase();

  if (needsErrCheck) {
    unmarshalBody += `${indent.get()}var err error\n`;
  }
  unmarshalBody += switchCaseBody;
  if (needsErrCheck) {
    unmarshalBody += `${indent.get()}if err != nil {\n`;
    unmarshalBody += `${indent.push().get()}return fmt.Errorf("unmarshalling type %T: %v", ${receiver}, err)\n`;
    unmarshalBody += `${indent.pop().get()}}\n`;
  }
  indent.pop(); // level 1
  unmarshalBody += `${indent.get()}}\n`; // end for key, val := range rawMsg
  unmarshalBody += `${indent.get()}return nil\n`;
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
function generateDiscriminatorUnmarshaller(modelType: go.Model | go.PolymorphicModel, field: go.ModelField, receiver: string, indent: helpers.Indentation): string {
  const propertyName = field.name;

  // these are the simple, non-nested cases (e.g. IterfaceType, []InterfaceType, map[string]InterfaceType)
  if (field.type.kind === 'interface') {
    return `${indent.get()}${receiver}.${propertyName}, err = unmarshal${field.type.name}(val)\n`;
  } else if (field.type.kind === 'slice' && field.type.elementType.kind === 'interface') {
    return `${indent.get()}${receiver}.${propertyName}, err = unmarshal${field.type.elementType.name}Array(val)\n`;
  } else if (field.type.kind === 'map' && field.type.valueType.kind === 'interface') {
    return `${indent.get()}${receiver}.${propertyName}, err = unmarshal${field.type.valueType.name}Map(val)\n`;
  }

  // nested case (e.g. [][]InterfaceType, map[string]map[string]InterfaceType etc)
  // first, unmarshal the raw data
  const rawTargetVar = `${field.serializedName}Raw`;
  let text = `${indent.get()}var ${rawTargetVar} ${recursiveGetDiscriminatorTypeName(modelType, field.type, true)}\n`;
  text += `${indent.get()}if err = json.Unmarshal(val, &${rawTargetVar}); err != nil {\n`;
  text += `${indent.push().get()}return err\n`;
  text += `${indent.pop().get()}}\n`;

  // create a local instantiation of the final type
  const finalTargetVar = field.serializedName;
  let finalTargetCtor = recursiveGetDiscriminatorTypeName(modelType, field.type, false);
  if (field.type.kind === 'slice') {
    finalTargetCtor = `make(${finalTargetCtor}, len(${rawTargetVar}))`;
  } else {
    // must be a dictionary
    finalTargetCtor = `${finalTargetCtor}{}`;
  }
  text += `${indent.get()}${finalTargetVar} := ${finalTargetCtor}\n`;

  // now populate the final type
  text += recursivePopulateDiscriminator(modelType, field.type, receiver, rawTargetVar, finalTargetVar, indent, 1);

  // finally, assign the final target to the property
  text += `${indent.get()}${receiver}.${propertyName} = ${finalTargetVar}\n`;
  return text;
}

// constructs the type name for a nested discriminated type
// raw e.g. map[string]json.RawMessage, []json.RawMessage etc
// !raw e.g. map[string]map[string]InterfaceType, [][]InterfaceType etc
function recursiveGetDiscriminatorTypeName(modelType: go.Model | go.PolymorphicModel, item: go.WireType, raw: boolean): string {
  // when raw is true, stop recursing at the level before the leaf schema
  if (item.kind === 'slice') {
    if (!raw || item.elementType.kind !== 'interface') {
      return `[]${recursiveGetDiscriminatorTypeName(modelType, item.elementType, raw)}`;
    }
  } else if (item.kind === 'map') {
    if (!raw || item.valueType.kind !== 'interface') {
      return `map[string]${recursiveGetDiscriminatorTypeName(modelType, item.valueType, raw)}`;
    }
  }
  if (raw) {
    return 'json.RawMessage';
  }
  return go.getTypeDeclaration(item, modelType.pkg);
}

/**
 * recursively constructs the text to populate a nested discriminator
 *
 * @param modelType the type that contains item
 * @param item the type for which to create the population
 * @param receiver the name of the receiver for the method to contain the expression
 * @param rawSrc contains the raw unmarshaled JSON source
 * @param dest the variable that will contain the result of the expression
 * @param indent the current indentation helper (level increases with each recursive call)
 * @param nesting the current level of nesting (increments with each recursive call)
 * @returns the text populating the discriminator
 */
function recursivePopulateDiscriminator(
  modelType: go.Model | go.PolymorphicModel,
  item: go.WireType,
  receiver: string,
  rawSrc: string,
  dest: string,
  indent: helpers.Indentation,
  nesting: number,
): string {
  let text = '';
  let interfaceName = '';
  let targetType = '';

  if (item.kind === 'slice') {
    if (item.elementType.kind !== 'interface') {
      if (nesting > 1) {
        // at nestling level 1, the destination var was already created in generateDiscriminatorUnmarshaller()
        text += `${indent.get()}${dest} = make(${recursiveGetDiscriminatorTypeName(modelType, item, false)}, len(${rawSrc}))\n`;
      }

      text += `${indent.get()}for i${nesting} := range ${rawSrc} {\n`;
      rawSrc = `${rawSrc}[i${nesting}]`; // source becomes each element in the source slice
      dest = `${dest}[i${nesting}]`; // update destination to each element in the destination slice
      indent.push();
      text += recursivePopulateDiscriminator(modelType, item.elementType, receiver, rawSrc, dest, indent, nesting + 1);
      indent.pop();
      text += `${indent.get()}}\n`;
      return text;
    }

    // we're at leaf node - 1, so get the interface from the element's type
    interfaceName = go.getTypeDeclaration(item.elementType, modelType.pkg);
    targetType = 'Array';
  } else if (item.kind === 'map') {
    if (item.valueType.kind !== 'interface') {
      if (nesting > 1) {
        // at nestling level 1, the destination var was already created in generateDiscriminatorUnmarshaller()
        text += `${indent.get()}${dest} = ${recursiveGetDiscriminatorTypeName(modelType, item, false)}{}\n`;
      }

      text += `${indent.get()}for k${nesting}, v${nesting} := range ${rawSrc} {\n`;
      rawSrc = `v${nesting}`; // source becomes the current value in the source map
      dest = `${dest}[k${nesting}]`; // update destination to the destination map's value for the current key
      indent.push();
      text += recursivePopulateDiscriminator(modelType, item.valueType, receiver, rawSrc, dest, indent, nesting + 1);
      indent.pop();
      text += `${indent.get()}}\n`;
      return text;
    }

    // we're at leaf node - 1, so get the interface from the element's type
    interfaceName = go.getTypeDeclaration(item.valueType, modelType.pkg);
    targetType = 'Map';
  }

  text += `${indent.get()}${dest}, err = unmarshal${interfaceName}${targetType}(${rawSrc})\n`;
  text += `${indent.get()}if err != nil {\n`;
  text += `${indent.push().get()}return fmt.Errorf("unmarshalling type %T: %v", ${receiver}, err)\n`;
  text += `${indent.pop().get()}}\n`;

  return text;
}

/**
 * generates an implementation of MarshalXML for the provided type.
 * the method impl is added to modelDef.SerDe.methods.
 *
 * @param modelDef the type for which to implement MarshalXML
 * @param imports the import manager currently in scope
 */
function generateXMLMarshaller(modelDef: ModelDef, imports: ImportManager, indent: helpers.Indentation): void {
  // only needed for types with time.Time, maps, or where the XML name doesn't match the type name
  const receiver = modelDef.receiverName();
  const desc = `MarshalXML implements the xml.Marshaller interface for type ${modelDef.Model.name}.`;
  let text = `func (${receiver} ${modelDef.Model.name}) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {\n`;
  if (modelDef.Model.xml?.name) {
    text += `${indent.get()}start.Name.Local = "${modelDef.Model.xml.name}"\n`;
  }
  text += generateAliasType(modelDef.Model, receiver, true, imports, indent);
  for (const field of modelDef.Model.fields) {
    if (field.type.kind === 'slice') {
      text += `${indent.get()}if ${receiver}.${field.name} != nil {\n`;
      text += `${indent.push().get()}aux.${field.name} = &${receiver}.${field.name}\n`;
      text += `${indent.pop().get()}}\n`;
    } else if (field.annotations.isAdditionalProperties || field.type.kind === 'map') {
      text += `${indent.get()}aux.${field.name} = (additionalProperties)(${receiver}.${field.name})\n`;
    } else if (field.type.kind === 'encodedBytes') {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      text += `${indent.get()}if ${receiver}.${field.name} != nil {\n`;
      indent.push();
      text += `${indent.get()}encoded${field.name} := runtime.EncodeByteArray(${receiver}.${field.name}, runtime.Base64${field.type.encoding}Format)\n`;
      text += `${indent.get()}aux.${field.name} = &encoded${field.name}\n`;
      indent.pop();
      text += `${indent.get()}}\n`;
    }
  }
  text += `${indent.get()}return enc.EncodeElement(aux, start)\n`;
  text += '}\n\n';
  modelDef.SerDe.methods.push({ name: 'MarshalXML', desc: desc, text: text });
}

/**
 * generates an implementation of UnmarshalXML for the provided type.
 * the method impl is added to modelDef.SerDe.methods.
 *
 * @param modelDef the type for which to implement UnmarshalXML
 * @param imports the import manager currently in scope
 */
function generateXMLUnmarshaller(modelDef: ModelDef, imports: ImportManager, indent: helpers.Indentation): void {
  // non-polymorphic case, must be something with time.Time
  const receiver = modelDef.receiverName();
  const desc = `UnmarshalXML implements the xml.Unmarshaller interface for type ${modelDef.Model.name}.`;
  let text = `func (${receiver} *${modelDef.Model.name}) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {\n`;
  text += generateAliasType(modelDef.Model, receiver, false, imports, indent);
  text += `${indent.get()}if err := dec.DecodeElement(aux, &start); err != nil {\n`;
  text += `${indent.push().get()}return err\n`;
  text += `${indent.pop().get()}}\n`;
  for (const field of modelDef.Model.fields) {
    if (field.type.kind === 'time') {
      text += `${indent.get()}if aux.${field.name} != nil && !(*time.Time)(aux.${field.name}).IsZero() {\n`;
      text += `${indent.push().get()}${receiver}.${field.name} = (*time.Time)(aux.${field.name})\n`;
      text += `${indent.pop().get()}}\n`;
    } else if (field.annotations.isAdditionalProperties || field.type.kind === 'map') {
      text += `${indent.get()}${receiver}.${field.name} = (map[string]*string)(aux.${field.name})\n`;
    } else if (field.type.kind === 'encodedBytes') {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      text += `${indent.get()}if aux.${field.name} != nil {\n`;
      indent.push();
      text += `${indent.get()}if err := runtime.DecodeByteArray(*aux.${field.name}, &${receiver}.${field.name}, runtime.Base64${field.type.encoding}Format); err != nil {\n`;
      text += `${indent.push().get()}return err\n`;
      text += `${indent.pop().get()}}\n`;
      indent.pop();
      text += `${indent.get()}}\n`;
    }
  }
  text += `${indent.get()}return nil\n`;
  text += '}\n\n';
  modelDef.SerDe.methods.push({ name: 'UnmarshalXML', desc: desc, text: text });
}

/**
 * generates an alias type used by custom XML marshaller/unmarshaller
 *
 * @param modelType the type for which to create the alias
 * @param receiver the name of the receiver for the type's serde method
 * @param forMarshal when true, indicates type is to be used in a marshaller (else an unmarshaller)
 * @returns the text for an initialized type alias
 */
function generateAliasType(modelType: go.Model | go.PolymorphicModel, receiver: string, forMarshal: boolean, imports: ImportManager, indent: helpers.Indentation): string {
  let text = `${indent.get()}type alias ${modelType.name}\n`;
  text += `${indent.get()}aux := &struct {\n`;
  text += `${indent.push().get()}*alias\n`;
  for (const field of modelType.fields) {
    const sn = getXMLSerialization(field);
    if (field.type.kind === 'time') {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/datetime');
      text += `${indent.get()}${field.name} *datetime.${field.type.format} \`xml:"${sn}"\`\n`;
    } else if (field.annotations.isAdditionalProperties || field.type.kind === 'map') {
      text += `${indent.get()}${field.name} additionalProperties \`xml:"${sn}"\`\n`;
    } else if (field.type.kind === 'slice') {
      text += `${indent.get()}${field.name} *${go.getTypeDeclaration(field.type, modelType.pkg)} \`xml:"${sn}"\`\n`;
    } else if (field.type.kind === 'encodedBytes') {
      text += `${indent.get()}${field.name} *string \`xml:"${sn}"\`\n`;
    }
  }
  text += `${indent.pop().get()}}{\n`;
  let rec = receiver;
  if (forMarshal) {
    rec = '&' + rec;
  }
  text += `${indent.push().get()}alias: (*alias)(${rec}),\n`;
  if (forMarshal) {
    // emit code to initialize time fields
    for (const field of modelType.fields) {
      if (field.type.kind !== 'time') {
        continue;
      }
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/datetime');
      text += `${indent.get()}${field.name}: (*datetime.${field.type.format})(${receiver}.${field.name}),\n`;
    }
  }
  text += `${indent.pop().get()}}\n`;
  return text;
}

/** represents a method on a model */
interface ModelMethod {
  name: string;
  desc: string;
  text: string;
}

/** used to track which serde helpers and methods are required for a struct */
class SerDeInfo {
  methods: Array<ModelMethod>;
  needsJSONPopulate: boolean;
  needsJSONPopulateTime: boolean;
  needsJSONUnpopulate: boolean;
  needsJSONUnpopulateTime: boolean;
  needsJSONPopulateByteArray: boolean;
  needsJSONPopulateAny: boolean;
  needsJSONPopulateMultipart: boolean;

  constructor() {
    this.methods = new Array<ModelMethod>();
    this.needsJSONPopulate = false;
    this.needsJSONPopulateTime = false;
    this.needsJSONUnpopulate = false;
    this.needsJSONUnpopulateTime = false;
    this.needsJSONPopulateByteArray = false;
    this.needsJSONPopulateAny = false;
    this.needsJSONPopulateMultipart = false;
  }
}

/** represents model definition as a Go struct */
class ModelDef {
  readonly Model: go.Model | go.PolymorphicModel;
  readonly Format: helpers.SerDeFormat;
  readonly SerDe: SerDeInfo;
  readonly Methods: Array<ModelMethod>;

  constructor(model: go.Model | go.PolymorphicModel, format: helpers.SerDeFormat) {
    this.Model = model;
    this.Format = format;
    this.SerDe = new SerDeInfo();
    this.Methods = new Array<ModelMethod>();
  }

  /** returns the text for the struct definition */
  text(indent: helpers.Indentation): string {
    let text = helpers.formatDocComment(this.Model.docs);
    text += `type ${this.Model.name} struct {\n`;

    // group fields by required/optional/read-only in that order
    this.Model.fields.sort((lhs: go.ModelField, rhs: go.ModelField): number => {
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

    for (const field of this.Model.fields) {
      if (field.docs.summary || field.docs.description) {
        if (!first) {
          // add an extra new-line between fields IFF the field
          // has a comment and it's not the very first one.
          text += '\n';
        }
        text += helpers.formatDocComment(field.docs);
      }
      let typeName = go.getTypeDeclaration(field.type, this.Model.pkg);
      if (field.type.kind === 'literal') {
        // for constants we use the underlying type name
        typeName = go.getLiteralTypeDeclaration(field.type.type);
      }
      let serialization = field.serializedName;
      if (this.Format === 'JSON') {
        serialization += ',omitempty';
      } else if (this.Format === 'XML') {
        serialization = getXMLSerialization(field);
      }
      let tag = '';
      // only emit tags for XML; JSON uses custom marshallers/unmarshallers
      if (this.Format === 'XML' && !field.annotations.isAdditionalProperties) {
        tag = ` \`xml:"${serialization}"\``;
      }
      text += `${indent.get()}${field.name} ${helpers.star(field.byValue)}${typeName}${tag}\n`;
      first = false;
    }

    text += '}\n\n';
    return text;
  }

  /** returns the name to use for method receivers on this struct */
  receiverName(): string {
    const typeName = this.Model.name;
    return typeName[0].toLowerCase();
  }
}

/**
 * returns the serialization options to use in the XML tag on a model field
 *
 * @param field the field for which to construct the tag's contents
 * @returns the contents for the XML tag
 */
function getXMLSerialization(field: go.ModelField): string {
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
    let inner = field.serializedName;
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
    if (field.xml?.wraps) {
      serialization += `>${field.xml.wraps}`;
    } else {
      serialization = inner;
    }
  }
  return serialization;
}
