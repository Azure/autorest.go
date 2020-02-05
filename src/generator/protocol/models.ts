/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment, joinComma } from '@azure-tools/codegen';
import { CodeModel, ConstantSchema, ObjectSchema, ChoiceSchema, Language, Schema, Schemas, SchemaType, StringSchema, Property } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { ContentPreamble, getEnums, HasDescription, ImportManager, SortAscending } from '../common/helpers';

// Creates the content in models.go
export async function generateModels(session: Session<CodeModel>): Promise<string> {
  let text = await ContentPreamble(session);

  // we do model generation first as it can add imports to the imports list
  const structs = generateStructs(session.model.schemas.objects);
  // add structs from operation responses
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      if (op.responses) {
        structs.push(generateStruct(op.responses[0].language.go!, op.responses[0].language.go!.properties));
      }
    }
  }

  // imports
  if (imports.length() > 0) {
    text += imports.text();
  }

  // enums
  session.model.schemas.choices?.sort(
    (a: ChoiceSchema<StringSchema>, b: ChoiceSchema<StringSchema>) => { return SortAscending(a.language.go!.name, b.language.go!.name); }
  );
  text += generateEnums(session.model.schemas);

  // structs
  structs.sort((a: StructDef, b: StructDef) => { return SortAscending(a.Language.name, b.Language.name) });
  for (const struct of values(structs)) {
    text += struct.text();
  }
  return text;
}

// this list of packages to import
const imports = new ImportManager();

// represents a struct definition
class StructDef {
  readonly Language: Language;
  readonly Properties?: Property[];

  constructor(language: Language, props?: Property[]) {
    this.Language = language;
    this.Properties = props;
    if (this.Properties) {
      this.Properties.sort((a: Property, b: Property) => { return SortAscending(a.language.go!.name, b.language.go!.name); });
    }
  }

  text(): string {
    let text = '';
    if (HasDescription(this.Language)) {
      text += `${comment(this.Language.description, '// ')}\n`;
    }
    text += `type ${this.Language.name} struct {\n`;
    for (const prop of values(this.Properties)) {
      if (HasDescription(prop.language.go!)) {
        text += `\t${comment(prop.language.go!.description, '// ')}\n`;
      }
      let typeName = prop.schema.language.go!.name;
      if (prop.schema.type === SchemaType.Constant) {
        // for constants we use the underlying type name
        typeName = (<ConstantSchema>prop.schema).valueType.language.go!.name;
      }
      text += `\t${prop.language.go!.name} ${this.byref(prop)}${typeName}\n`;
    }
    text += '}\n\n';
    if (this.Language.errorType) {
      text += `func new${this.Language.name}(resp *azcore.Response) error {\n`;
      text += `\terr := ${this.Language.name}{}\n`;
      text += `\tif err := resp.UnmarshalAsJSON(&err); err != nil {\n`;
      text += `\t\treturn err\n`;
      text += `\t}\n`;
      text += '\treturn err\n';
      text += '}\n\n';
      text += `func (e ${this.Language.name}) Error() string {\n`;
      text += `\treturn "TODO"\n`;
      text += '}\n\n';
    }
    return text;
  }

  private byref(prop: Property): string {
    if (prop.language.go!.noByRef === true) {
      return '';
    }
    return '*';
  }
}

function generateEnums(schemas: Schemas): string {
  const enums = getEnums(schemas);
  let text = '';
  for (const enm of values(enums)) {
    if (enm.desc) {
      text += `${comment(enm.name, '// ')} - ${enm.desc}\n`;
    }
    text += `type ${enm.name} ${enm.type}\n\n`;
    const vals = new Array<string>();
    text += 'const (\n'
    for (const val of values(enm.choices)) {
      if (HasDescription(val.language.go!)) {
        text += `\t${comment(val.language.go!.name, '// ')} - ${val.language.go!.description}\n`;
      }
      text += `\t${val.language.go!.name} ${enm.name} = "${val.value}"\n`;
      vals.push(val.language.go!.name);
    }
    text += ")\n\n"
    text += `func ${enm.funcName}() []${enm.name} {\n`;
    text += `\treturn []${enm.name}{${joinComma(vals, (item: string) => item)}}\n`;
    text += '}\n\n';
  }
  return text;
}

function generateStructs(objects?: ObjectSchema[]): StructDef[] {
  const structTypes = new Array<StructDef>();
  for (const obj of values(objects)) {
    structTypes.push(generateStruct(obj.language.go!, obj.properties));
  }
  return structTypes;
}

function generateStruct(lang: Language, props?: Property[]): StructDef {
  const st = new StructDef(lang, props);
  for (const prop of values(props)) {
    if (lang.errorType) {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    }
    addImportForSchemaType(prop.schema);
  }
  return st;
}

function addImportForSchemaType(schema: Schema) {
  switch (schema.type) {
    case SchemaType.Date:
    case SchemaType.DateTime:
      imports.add('time');
  }
}
