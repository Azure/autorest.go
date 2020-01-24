/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment, joinComma } from '@azure-tools/codegen'
import { CodeModel, ObjectSchema, ChoiceSchema, Schema, SchemaType, StringSchema, Property } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';

// Creates the content in models.go
export async function generateModels(session: Session<CodeModel>): Promise<string> {
  const headerText = comment(await session.getValue("header-text", "MISSING LICENSE HEADER"), "// ");
  const namespace = await session.getValue('namespace');
  let text = `${headerText}\n\n`
  text += `package ${namespace}\n\n`

  // we do generation first as it can add imports to the imports list
  const enums = generateEnums(session.model.schemas.choices);
  const structs = generateStructs(session.model.schemas.objects);
  // add structs from operation responses
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      if (op.responses) {
        structs.push(generateStruct(op.responses[0].language.go!.name, op.responses[0].language.go!.description, op.responses[0].language.go!.properties));
      }
    }
  }

  // imports
  if (imports.length() > 0) {
    const items = imports.items();
    items.sort(sortAscending);
    text += 'import (\n';
    for (const imp of values(items)) {
      text += `\t"${imp}"\n`;
    }
    text += ')\n\n';
  }

  // enums
  enums.sort((a: EnumType, b: EnumType): number => { return sortAscending(a.name, b.name) });
  for (const enm of values(enums)) {
    text += `type ${enm.name} ${enm.type}\n\n`;
    enm.values.sort((a: EnumValue, b: EnumValue): number => { return sortAscending(a.name, b.name) });
    text += 'const (\n'
    for (const val of values(enm.values)) {
      text += `\t${val.name} ${enm.name} = "${val.value}"\n`;
    }
    text += ")\n\n"
    text += `func Possible${enm.name}Values() []${enm.name} {\n`;
    text += `\treturn []${enm.name}{${joinComma(enm.values, (item: EnumValue) => item.name)}}\n`;
    text += '}\n\n';
  }

  // structs
  structs.sort((a: StructType, b: StructType): number => { return sortAscending(a.name, b.name) });
  for (const struct of values(structs)) {
    if (struct.comment) {
      text += `${comment(struct.comment, '// ')}\n`;
    }
    text += `type ${struct.name} struct {\n`;
    for (const field of values(struct.fields)) {
      if (field.comment) {
        text += `\t${comment(field.comment, '// ')}\n`;
      }
      text += `\t${field.name} ${field.type}\n`;
    }
    text += '}\n\n'
  }
  return text;
}

// tracks packages that need to be imported
class ImportManager {
  private imports: string[];

  constructor() {
    this.imports = new Array<string>();
  }

  // adds a package for importing if not already in the list
  add(imp: string) {
    for (const existing of values(this.imports)) {
      if (existing === imp) {
        return;
      }
    }
    this.imports.push(imp);
  }

  // returns the list of packages to import
  items(): string[] {
    return this.imports;
  }

  // returns the number of packages in the list
  length(): number {
    return this.imports.length
  }
}

// this list of packages to import
const imports = new ImportManager();

// represents an enum type definition and its values
class EnumType {
  readonly name: string;
  readonly type: string;
  readonly values: EnumValue[];

  constructor(name: string, type: string) {
    this.name = name;
    this.type = type;
    this.values = new Array<EnumValue>();
  }
}

// represents an enum value
class EnumValue {
  readonly name: string;
  readonly value: string | number | boolean;

  constructor(name: string, value: string | number | boolean) {
    this.name = name;
    this.value = value;
  }
}

// represents a struct type definition
class StructType {
  readonly name: string;
  readonly comment?: string;
  readonly fields: StructField[];

  constructor(name: string, comment: string) {
    this.name = name;
    if (!comment.startsWith('MISSING')) {
      this.comment = comment;
    }
    this.fields = new Array<StructField>();
  }
}

// represents a field in a struct
class StructField {
  readonly name: string;
  readonly comment?: string;
  readonly type: string;

  constructor(name: string, comment: string, type: string) {
    this.name = name;
    if (!comment.startsWith('MISSING')) {
      this.comment = comment;
    }
    this.type = type;
  }
}

function generateEnums(enums?: ChoiceSchema<StringSchema>[]): EnumType[] {
  const enumTypes = new Array<EnumType>();
  for (const enm of values(enums)) {
    const et = new EnumType(enm.language.go!.name, enm.choiceType.language.go!.name);
    for (const val of values(enm.choices)) {
      et.values.push(new EnumValue(val.language.go!.name, val.value));
    }
    enumTypes.push(et);
  }
  return enumTypes;
}

function generateStructs(objects?: ObjectSchema[]): StructType[] {
  const structTypes = new Array<StructType>();
  for (const obj of values(objects)) {
    structTypes.push(generateStruct(obj.language.go!.name, obj.language.go!.description, obj.properties));
  }
  return structTypes;
}

function generateStruct(name: string, comment: string, props?: Property[]): StructType {
  const st = new StructType(name, comment);
  for (const prop of values(props)) {
    addImportForSchemaType(prop.schema);
    st.fields.push(new StructField(prop.language.go!.name, prop.language.go!.description, prop.schema.language.go!.name));
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

function sortAscending(a: string, b: string): number {
  return a < b ? -1 : a > b ? 1 : 0;
}
