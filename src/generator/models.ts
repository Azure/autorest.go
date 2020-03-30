/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment, pascalCase } from '@azure-tools/codegen';
import { CodeModel, ConstantSchema, ImplementationLocation, ObjectSchema, Language, Schema, SchemaType, Parameter, Property } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { contentPreamble, hasDescription, ImportManager, isArraySchema, sortAscending } from '../common/helpers';

// Creates the content in models.go
export async function generateModels(session: Session<CodeModel>): Promise<string> {
  let text = await contentPreamble(session);

  // we do model generation first as it can add imports to the imports list
  const structs = generateStructs(session.model.schemas.objects);
  const responseSchemas = <Array<Schema>>session.model.language.go!.responseSchemas;
  for (const schema of values(responseSchemas)) {
    structs.push(generateStruct(schema.language.go!.responseType, schema.language.go!.properties));
  }
  // add types from requests and responses
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      // add structs from optional operation params
      if (op.requests![0].language.go!.optionalParam) {
        structs.push(generateOptionalParamsStruct(op.requests![0].language.go!.optionalParam, op.requests![0].language.go!.optionalParam.params));
      }
    }
  }

  // imports
  if (imports.length() > 0) {
    text += imports.text();
  }

  // structs
  structs.sort((a: StructDef, b: StructDef) => { return sortAscending(a.Language.name, b.Language.name) });
  for (const struct of values(structs)) {
    text += struct.text();
    text += struct.marshaller();
    text += struct.unmarshaller();
  }
  return text;
}

// this list of packages to import
const imports = new ImportManager();

// represents a struct definition
class StructDef {
  readonly Language: Language;
  readonly Properties?: Property[];
  readonly Parameters?: Parameter[];

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
  }

  text(): string {
    let text = '';
    if (hasDescription(this.Language)) {
      text += `${comment(this.Language.description, '// ')}\n`;
    }
    text += `type ${this.Language.name} struct {\n`;
    // used to track when to add an extra \n between fields that have comments
    let first = true;
    for (const prop of values(this.Properties)) {
      // adding the Inner prefix on error types, since errors in Go have an Error() method
      // in order to implement the error interface. This causes errors to not be able to have
      // an Error field as well, since it would cause confusion
      if (this.Language.errorType && prop.language.go!.name === 'Error') {
        prop.language.go!.name = 'Inner' + prop.language.go!.name;
      }
      if (hasDescription(prop.language.go!)) {
        if (!first) {
          // add an extra new-line between fields IFF the field
          // has a comment and it's not the very first one.
          text += '\n';
        }
        text += `\t${comment(prop.language.go!.description, '// ')}\n`;
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
      text += `\t${prop.language.go!.name} *${typeName}${tag}\n`;
      first = false;
    }
    for (const param of values(this.Parameters)) {
      // if Parameters is set this is an optional args struct
      // none of its fields need to participate in marshalling
      if (param.implementation === ImplementationLocation.Client) {
        // don't add globals to the per-method options struct
        continue;
      }
      if (hasDescription(param.language.go!)) {
        text += `\t${comment(param.language.go!.description, '// ')}\n`;
      }
      text += `\t${pascalCase(param.language.go!.name)} *${param.schema.language.go!.name}\n`;
    }
    text += '}\n\n';
    if (this.Language.errorType) {
      text += `func ${this.Language.constructorName}(resp *azcore.Response) error {\n`;
      text += `\terr := ${this.Language.name}{}\n`;
      text += `\tif err := resp.UnmarshalAs${(<string>this.Language.marshallingFormat).toUpperCase()}(&err); err != nil {\n`;
      text += `\t\treturn err\n`;
      text += `\t}\n`;
      text += '\treturn err\n';
      text += '}\n\n';
      text += `func (e ${this.Language.name}) Error() string {\n`;
      text += `\tmsg := ""\n`;
      for (const prop of values(this.Properties)) {
        text += `\tif e.${prop.language.go!.name} != nil {\n`;
        text += `\t\tmsg += fmt.Sprintf("${prop.language.go!.name}: %v\\n", *e.${prop.language.go!.name})\n`;
        text += `\t}\n`;
      }
      text += '\tif msg == "" {\n';
      text += '\t\tmsg = "missing error info"\n';
      text += '\t}\n';
      text += '\treturn msg\n';
      text += '}\n\n';
    }
    return text;
  }

  // creates a custom marshaller for this type
  marshaller(): string {
    // only needed for types with time.Time or where the XML name doesn't match the type name
    if (this.Language.needsDateTimeMarshalling === undefined && this.Language.xmlWrapperName === undefined) {
      return '';
    }
    const receiver = this.Language.name[0].toLowerCase();
    let formatSig = 'JSON() ([]byte, error)';
    if (this.Language.marshallingFormat === 'xml') {
      formatSig = 'XML(e *xml.Encoder, start xml.StartElement) error'
    }
    let text = `func (${receiver} ${this.Language.name}) Marshal${formatSig} {\n`;
    if (this.Language.xmlWrapperName) {
      text += `\tstart.Name.Local = "${this.Language.xmlWrapperName}"\n`;
    }
    text += this.generateAliasType(receiver, true);
    if (this.Language.marshallingFormat === 'json') {
      text += '\treturn json.Marshal(aux)\n';
    } else {
      text += '\treturn e.EncodeElement(aux, start)\n';
    }
    text += '}\n\n';
    return text;
  }

  // creates a custom unmarshaller for this type
  unmarshaller(): string {
    // only needed for types with time.Time
    if (this.Language.needsDateTimeMarshalling === undefined) {
      return '';
    }
    const receiver = this.Language.name[0].toLowerCase();
    let formatSig = 'JSON(data []byte)';
    if (this.Language.marshallingFormat === 'xml') {
      formatSig = 'XML(d *xml.Decoder, start xml.StartElement)';
    }
    let text = `func (${receiver} *${this.Language.name}) Unmarshal${formatSig} error {\n`;
    text += this.generateAliasType(receiver, false);
    if (this.Language.marshallingFormat === 'json') {
      text += '\tif err := json.Unmarshal(data, aux); err != nil {\n';
      text += '\t\treturn err\n';
      text += '\t}\n';
    } else {
      text += '\tif err := d.DecodeElement(aux, &start); err != nil {\n';
      text += '\t\treturn err\n';
      text += '\t}\n';
    }
    for (const prop of values(this.Properties)) {
      if (prop.schema.type !== SchemaType.DateTime) {
        continue;
      }
      text += `\t${receiver}.${prop.language.go!.name} = (*time.Time)(aux.${prop.language.go!.name})\n`;
    }
    text += '\treturn nil\n';
    text += '}\n\n';
    return text;
  }

  // generates an alias type used by custom marshaller/unmarshaller
  private generateAliasType(receiver: string, forMarshal: boolean): string {
    let text = `\ttype alias ${this.Language.name}\n`;
    text += `\taux := &struct {\n`;
    text += `\t\t*alias\n`;
    for (const prop of values(this.Properties)) {
      if (prop.schema.type !== SchemaType.DateTime) {
        continue;
      }
      let sn = prop.serializedName;
      if (prop.schema.serialization?.xml?.name) {
        // xml can specifiy its own name, prefer that if available
        sn = prop.schema.serialization.xml.name;
      }
      text += `\t\t${prop.language.go!.name} *${prop.schema.language.go!.internalTimeType} \`${this.Language.marshallingFormat}:"${sn}"\`\n`;
    }
    text += `\t}{\n`;
    let rec = receiver;
    if (forMarshal) {
      rec = '&' + rec;
    }
    text += `\t\talias: (*alias)(${rec}),\n`;
    if (forMarshal) {
      // emit code to initialize time fields
      for (const prop of values(this.Properties)) {
        if (prop.schema.type !== SchemaType.DateTime) {
          continue;
        }
        text += `\t\t${prop.language.go!.name}: (*${prop.schema.language.go!.internalTimeType})(${receiver}.${prop.language.go!.name}),\n`;
      }
    }
    text += `\t}\n`;
    return text;
  }
}

function generateStructs(objects?: ObjectSchema[]): StructDef[] {
  const structTypes = new Array<StructDef>();
  for (const obj of values(objects)) {
    structTypes.push(generateStruct(obj.language.go!, obj.properties));
  }
  return structTypes;
}

function generateStruct(lang: Language, props?: Property[]): StructDef {
  if (lang.errorType) {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    imports.add('fmt');
  }
  if (lang.responseType) {
    imports.add("net/http");
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

function newProperty(name: string, desc: string, schema: Schema): Property {
  let prop = new Property(name, desc, schema);
  prop.language.go = prop.language.default;
  return prop;
}

function generateOptionalParamsStruct(lang: Language, params: Parameter[]): StructDef {
  const st = new StructDef(lang, undefined, params);
  for (const param of values(params)) {
    imports.addImportForSchemaType(param.schema);
  }
  return st;
}
