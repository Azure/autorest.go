/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { camelCase, comment, pascalCase } from '@azure-tools/codegen';
import { CodeModel, ConstantSchema, GroupProperty, ImplementationLocation, ObjectSchema, Language, Schema, SchemaType, Parameter, Property } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { isArraySchema, isObjectSchema } from '../common/helpers';
import { contentPreamble, hasDescription, sortAscending, substituteDiscriminator } from './helpers';
import { ImportManager } from './imports';

// Creates the content in models.go
export async function generateModels(session: Session<CodeModel>): Promise<string> {
  let text = await contentPreamble(session);

  // we do model generation first as it can add imports to the imports list
  const structs = generateStructs(session.model.schemas.objects);
  const responseSchemas = <Array<Schema>>session.model.language.go!.responseSchemas;
  for (const schema of values(responseSchemas)) {
    const respType = generateStruct(schema.language.go!.responseType, schema.language.go!.properties);
    generateUnmarshallerForResponseEnvelope(respType);
    structs.push(respType);
  }
  const paramGroups = <Array<GroupProperty>>session.model.language.go!.parameterGroups;
  for (const paramGroup of values(paramGroups)) {
    structs.push(generateParamGroupStruct(paramGroup.language.go!, paramGroup.originalParameter));
  }

  // imports
  if (imports.length() > 0) {
    text += imports.text();
  }

  // structs
  structs.sort((a: StructDef, b: StructDef) => { return sortAscending(a.Language.name, b.Language.name) });
  for (const struct of values(structs)) {
    text += struct.discriminator();
    text += struct.text();
    struct.Methods.sort((a: StructMethod, b: StructMethod) => { return sortAscending(a.name, b.name) });
    for (const method of values(struct.Methods)) {
      text += method.text;
    }
  }
  return text;
}

// this list of packages to import
const imports = new ImportManager();

interface StructMethod {
  name: string;
  text: string;
}

// represents a struct definition
class StructDef {
  readonly Language: Language;
  readonly Properties?: Property[];
  readonly Parameters?: Parameter[];
  readonly Methods: StructMethod[];
  readonly ComposedOf: ObjectSchema[];

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
  }

  text(): string {
    let text = '';
    if (hasDescription(this.Language)) {
      text += `${comment(this.Language.description, '// ')}\n`;
    }
    text += `type ${this.Language.name} struct {\n`;
    // any composed types go first
    for (const comp of values(this.ComposedOf)) {
      text += `\t${comp.language.go!.name}\n`;
    }
    // used to track when to add an extra \n between fields that have comments
    let first = true;
    for (const prop of values(this.Properties)) {
      if (hasDescription(prop.language.go!)) {
        if (!first) {
          // add an extra new-line between fields IFF the field
          // has a comment and it's not the very first one.
          text += '\n';
        }
        text += `\t${comment(prop.language.go!.description, '// ')}\n`;
      }
      let typeName = substituteDiscriminator(prop.schema);
      if (prop.schema.type === SchemaType.Constant) {
        // for constants we use the underlying type name
        typeName = (<ConstantSchema>prop.schema).valueType.language.go!.name;
      }
      let serialization = prop.serializedName;
      if (this.Language.marshallingFormat === 'json') {
        serialization += ',omitempty';
      } else if (this.Language.marshallingFormat === 'xml') {
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
          if (prop.schema.serialization?.xml?.wrapped && this.Language.responseType !== true) {
            serialization += `>${inner}`;
          } else {
            serialization = inner;
          }
        }
      }
      let tag = ` \`${this.Language.marshallingFormat}:"${serialization}"\``;
      // if this is a response type then omit the tag IFF the marshalling format is
      // JSON, it's a header or is the RawResponse field.  XML marshalling needs a tag.
      if (this.Language.responseType === true && (this.Language.marshallingFormat !== 'xml' || prop.language.go!.name === 'RawResponse')) {
        tag = '';
      }
      let pointer = '*';
      if (prop.schema.language.go!.discriminatorInterface || prop.schema.language.go!.lroPointerException) {
        // pointer-to-interface introduces very clunky code
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
        text += `\t${comment(param.language.go!.description, '// ')}\n`;
      }
      let pointer = '*';
      if (param.required || param.schema.language.go!.discriminatorInterface) {
        // pointer-to-interface introduces very clunky code
        pointer = '';
      }
      text += `\t${pascalCase(param.language.go!.name)} ${pointer}${param.schema.language.go!.name}\n`;
    }
    text += '}\n\n';
    return text;
  }

  discriminator(): string {
    if (!this.Language.discriminatorInterface) {
      return '';
    }
    let text = `// ${this.Language.discriminatorInterface} provides polymorphic access to related types.\n`;
    text += `type ${this.Language.discriminatorInterface} interface {\n`;
    if (this.Language.discriminatorParent) {
      text += `\t${this.Language.discriminatorParent}\n`;
    }
    text += `\tGet${this.Language.name}() *${this.Language.name}\n`;
    text += '}\n\n';
    return text;
  }
}

function generateStructs(objects?: ObjectSchema[]): StructDef[] {
  const structTypes = new Array<StructDef>();
  for (const obj of values(objects)) {
    const props = new Array<Property>();
    // add immediate properties
    for (const prop of values(obj.properties)) {
      props.push(prop);
    }
    const structDef = generateStruct(obj.language.go!, props);
    // now add the parent type
    let parentType: ObjectSchema | undefined;
    for (const parent of values(obj.parents?.immediate)) {
      if (isObjectSchema(parent)) {
        parentType = parent;
        structDef.ComposedOf.push(parent);
      }
    }
    const hasPolymorphicField = values(obj.properties).first((each: Property) => {
      if (isObjectSchema(each.schema)) {
        return each.schema.discriminator !== undefined;
      }
      return false;
    });
    if (obj.language.go!.errorType) {
      // add Error() method
      let text = `func (e ${obj.language.go!.name}) Error() string {\n`;
      text += `\tmsg := ""\n`;
      for (const prop of values(structDef.Properties)) {
        text += `\tif e.${prop.language.go!.name} != nil {\n`;
        text += `\t\tmsg += fmt.Sprintf("${prop.language.go!.name}: %v\\n", *e.${prop.language.go!.name})\n`;
        text += `\t}\n`;
      }
      text += '\tif msg == "" {\n';
      text += '\t\tmsg = "missing error info"\n';
      text += '\t}\n';
      text += '\treturn msg\n';
      text += '}\n\n';
      structDef.Methods.push({ name: 'Error', text: text });
    }
    if (obj.discriminator) {
      // only need to generate interface method and internal marshaller for discriminators (Fish, Salmon, Shark)
      generateDiscriminatorMethods(obj, structDef, parentType!);
      // the root type doesn't get a marshaller as callers don't instantiate instances of it
      if (!obj.language.go!.rootDiscriminator) {
        generateDiscriminatedTypeMarshaller(obj, structDef, parentType!);
      }
      generateDiscriminatedTypeUnmarshaller(obj, structDef, parentType!);
    } else if (obj.discriminatorValue) {
      generateDiscriminatedTypeMarshaller(obj, structDef, parentType!);
      generateDiscriminatedTypeUnmarshaller(obj, structDef, parentType!);
    } else if (hasPolymorphicField) {
      generateDiscriminatedTypeUnmarshaller(obj, structDef, parentType!);
    } else if (obj.language.go!.needsDateTimeMarshalling || obj.language.go!.xmlWrapperName) {
      // TODO: unify marshalling schemes?
      generateMarshaller(structDef);
      if (obj.language.go!.needsDateTimeMarshalling) {
        generateUnmarshaller(structDef);
      }
    }
    structDef.ComposedOf.sort((a: ObjectSchema, b: ObjectSchema) => { return sortAscending(a.language.go!.name, b.language.go!.name); });
    structTypes.push(structDef);
  }
  return structTypes;
}

function generateStruct(lang: Language, props?: Property[]): StructDef {
  if (lang.errorType) {
    imports.add('fmt');
  }
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

function generateParamGroupStruct(lang: Language, params: Parameter[]): StructDef {
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
  const receiver = structDef.Language.name[0].toLowerCase();
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
    throw console.error(`failed to the discriminated type field for response envelope ${structDef.Language.name}`);
  }
  unmarshaller += `\tt, err := unmarshal${type}(data)\n`;
  unmarshaller += '\tif err != nil {\n';
  unmarshaller += '\t\treturn err\n';
  unmarshaller += '\t}\n';
  unmarshaller += `\t${receiver}.${field} = t\n`;
  unmarshaller += '\treturn nil\n';
  unmarshaller += '}\n\n';
  structDef.Methods.push({ name: 'UnmarshalJSON', text: unmarshaller });
}

function generateDiscriminatorMethods(obj: ObjectSchema, structDef: StructDef, parentType: ObjectSchema) {
  const typeName = obj.language.go!.name;
  const receiver = typeName[0].toLowerCase();
  // generate interface method
  const interfaceMethod = `Get${typeName}`;
  const method = `func (${receiver} *${typeName}) ${interfaceMethod}() *${typeName} { return ${receiver} }\n\n`;
  structDef.Methods.push({ name: interfaceMethod, text: method });
  if (obj.language.go!.errorType || obj.language.go!.childErrorType) {
    // errors don't need custom marshallers
    return;
  }
  // generate internal marshaller method
  const paramType = obj.discriminator!.property.schema.language.go!.name;
  const paramName = 'discValue';
  let marshalInteral = `func (${receiver} ${typeName}) marshalInternal(${paramName} ${paramType}) map[string]interface{} {\n`;
  if (parentType) {
    marshalInteral += `\tobjectMap := ${receiver}.${parentType.language.go!.name}.marshalInternal(${paramName})\n`;
  } else {
    marshalInteral += '\tobjectMap := make(map[string]interface{})\n';
  }
  for (const prop of values(structDef.Properties)) {
    if (prop.isDiscriminator) {
      marshalInteral += `\t${receiver}.${prop.language.go!.name} = &${paramName}\n`;
      marshalInteral += `\tobjectMap["${prop.serializedName}"] = ${receiver}.${prop.language.go!.name}\n`;
    } else {
      marshalInteral += `\tif ${receiver}.${prop.language.go!.name} != nil {\n`;
      if (prop.schema.language.go!.internalTimeType) {
        marshalInteral += `\t\tobjectMap["${prop.serializedName}"] = (*${prop.schema.language.go!.internalTimeType})(${receiver}.${prop.language.go!.name})\n`;
      } else {
        marshalInteral += `\t\tobjectMap["${prop.serializedName}"] = ${receiver}.${prop.language.go!.name}\n`;
      }
      marshalInteral += `\t}\n`;
    }
  }
  marshalInteral += '\treturn objectMap\n';
  marshalInteral += '}\n\n';
  structDef.Methods.push({ name: 'marshalInternal', text: marshalInteral });
}

function generateDiscriminatedTypeMarshaller(obj: ObjectSchema, structDef: StructDef, parentType: ObjectSchema) {
  if (obj.language.go!.errorType || obj.language.go!.childErrorType) {
    // errors don't need custom marshallers
    return;
  }
  imports.add('encoding/json');
  const typeName = structDef.Language.name;
  const receiver = typeName[0].toLowerCase();
  // generate marshaller method
  let marshaller = `func (${receiver} ${typeName}) MarshalJSON() ([]byte, error) {\n`;
  marshaller += `\tobjectMap := ${receiver}.${parentType!.language.go!.name}.marshalInternal(${obj.discriminatorValue})\n`;
  for (const prop of values(structDef.Properties)) {
    marshaller += `\tif ${receiver}.${prop.language.go!.name} != nil {\n`;
    if (prop.schema.language.go!.internalTimeType) {
      marshaller += `\t\tobjectMap["${prop.serializedName}"] = (*${prop.schema.language.go!.internalTimeType})(${receiver}.${prop.language.go!.name})\n`;
    } else {
      marshaller += `\t\tobjectMap["${prop.serializedName}"] = ${receiver}.${prop.language.go!.name}\n`;
    }
    marshaller += `\t}\n`;
  }
  marshaller += '\treturn json.Marshal(objectMap)\n';
  marshaller += '}\n\n';
  structDef.Methods.push({ name: 'MarshalJSON', text: marshaller });
}

function generateDiscriminatedTypeUnmarshaller(obj: ObjectSchema, structDef: StructDef, parentType?: ObjectSchema) {
  // there's a corner-case where a derived type might not add any new fields (Cookiecuttershark).
  // in this case skip adding the unmarshaller as it's not necessary and doesn't compile.
  if (!structDef.Properties || structDef.Properties.length === 0) {
    return;
  }
  imports.add('encoding/json');
  const typeName = structDef.Language.name;
  const receiver = typeName[0].toLowerCase();
  let unmarshaller = `func (${receiver} *${typeName}) UnmarshalJSON(data []byte) error {\n`;
  // polymorphic type, or type containing a polymorphic type
  unmarshaller += '\tvar rawMsg map[string]*json.RawMessage\n';
  unmarshaller += '\tif err := json.Unmarshal(data, &rawMsg); err != nil {\n';
  unmarshaller += '\t\treturn err\n';
  unmarshaller += '\t}\n';
  unmarshaller += '\tfor k, v := range rawMsg {\n';
  unmarshaller += '\t\tvar err error\n';
  unmarshaller += '\t\tswitch k {\n';
  // unmarshal each field one by one
  for (const prop of values(structDef.Properties)) {
    unmarshaller += `\t\tcase "${prop.serializedName}":\n`;
    unmarshaller += '\t\t\tif v != nil {\n';
    if (prop.schema.language.go!.discriminatorInterface) {
      unmarshaller += `\t\t\t\t${receiver}.${prop.language.go!.name}, err = unmarshal${prop.schema.language.go!.discriminatorInterface}(*v)\n`;
    } else if (isArraySchema(prop.schema) && prop.schema.elementType.language.go!.discriminatorInterface) {
      unmarshaller += `\t\t\t\t${receiver}.${prop.language.go!.name}, err = unmarshal${prop.schema.elementType.language.go!.discriminatorInterface}Array(*v)\n`;
    } else if (prop.schema.language.go!.internalTimeType) {
      unmarshaller += `\t\t\t\tvar aux ${prop.schema.language.go!.internalTimeType}\n`;
      unmarshaller += '\t\t\t\terr = json.Unmarshal(*v, &aux)\n';
      unmarshaller += `\t\t\t\t${receiver}.${prop.language.go!.name} = (*time.Time)(&aux)\n`;
    } else {
      unmarshaller += `\t\t\t\terr = json.Unmarshal(*v, &${receiver}.${prop.language.go!.name})\n`;
    }
    unmarshaller += '\t\t\t}\n';
  }
  unmarshaller += '\t\t}\n';
  unmarshaller += '\t\tif err != nil {\n';
  unmarshaller += '\t\t\treturn err\n';
  unmarshaller += '\t\t}\n';
  unmarshaller += '\t}\n';
  if (parentType) {
    unmarshaller += `\treturn json.Unmarshal(data, &${receiver}.${parentType.language.go!.name})\n`;
  } else {
    unmarshaller += '\treturn nil\n';
  }
  unmarshaller += '}\n\n';
  structDef.Methods.push({ name: 'UnmarshalJSON', text: unmarshaller });
}

function generateMarshaller(structDef: StructDef) {
  // only needed for types with time.Time or where the XML name doesn't match the type name
  const receiver = structDef.Language.name[0].toLowerCase();
  let formatSig = 'JSON() ([]byte, error)';
  let methodName = 'MarshalJSON';
  if (structDef.Language.marshallingFormat === 'xml') {
    formatSig = 'XML(e *xml.Encoder, start xml.StartElement) error';
    methodName = 'MarshalXML';
  }
  let text = `func (${receiver} ${structDef.Language.name}) Marshal${formatSig} {\n`;
  if (structDef.Language.xmlWrapperName) {
    text += `\tstart.Name.Local = "${structDef.Language.xmlWrapperName}"\n`;
  }
  text += generateAliasType(structDef, receiver, true);
  if (structDef.Language.marshallingFormat === 'json') {
    text += '\treturn json.Marshal(aux)\n';
  } else {
    text += '\treturn e.EncodeElement(aux, start)\n';
  }
  text += '}\n\n';
  structDef.Methods.push({ name: methodName, text: text });
}

function generateUnmarshaller(structDef: StructDef) {
  // non-polymorphic case, must be something with time.Time
  const receiver = structDef.Language.name[0].toLowerCase();
  let formatSig = 'JSON(data []byte)';
  let methodName = 'UnmarshalJSON';
  if (structDef.Language.marshallingFormat === 'xml') {
    formatSig = 'XML(d *xml.Decoder, start xml.StartElement)';
    methodName = 'UnmarshalXML';
  }
  let text = `func (${receiver} *${structDef.Language.name}) Unmarshal${formatSig} error {\n`;
  text += generateAliasType(structDef, receiver, false);
  if (structDef.Language.marshallingFormat === 'json') {
    text += '\tif err := json.Unmarshal(data, aux); err != nil {\n';
    text += '\t\treturn err\n';
    text += '\t}\n';
  } else {
    text += '\tif err := d.DecodeElement(aux, &start); err != nil {\n';
    text += '\t\treturn err\n';
    text += '\t}\n';
  }
  for (const prop of values(structDef.Properties)) {
    if (prop.schema.type !== SchemaType.DateTime) {
      continue;
    }
    text += `\t${receiver}.${prop.language.go!.name} = (*time.Time)(aux.${prop.language.go!.name})\n`;
  }
  text += '\treturn nil\n';
  text += '}\n\n';
  structDef.Methods.push({ name: methodName, text: text });
}

// generates an alias type used by custom marshaller/unmarshaller
function generateAliasType(structDef: StructDef, receiver: string, forMarshal: boolean): string {
  let text = `\ttype alias ${structDef.Language.name}\n`;
  text += `\taux := &struct {\n`;
  text += `\t\t*alias\n`;
  for (const prop of values(structDef.Properties)) {
    if (prop.schema.type !== SchemaType.DateTime) {
      continue;
    }
    let sn = prop.serializedName;
    if (prop.schema.serialization?.xml?.name) {
      // xml can specifiy its own name, prefer that if available
      sn = prop.schema.serialization.xml.name;
    }
    text += `\t\t${prop.language.go!.name} *${prop.schema.language.go!.internalTimeType} \`${structDef.Language.marshallingFormat}:"${sn}"\`\n`;
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
