/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment } from '@azure-tools/codegen';
import { CodeModel, ConstantSchema, ObjectSchema, ChoiceSchema, Language, Schema, SchemaType, StringSchema, Property, HttpHeader } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { ContentPreamble, HasDescription, ImportManager, LanguageHeader, SortAscending } from '../common/helpers';

// Creates the content in models.go
export async function generateModels(session: Session<CodeModel>): Promise<string> {
  let text = await ContentPreamble(session);

  // we do model generation first as it can add imports to the imports list
  const structs = generateStructs(session.model.schemas.objects);
  // add structs from operation responses
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      if (op.responses![0]) {
        // check if the response has http headers that it will expect information from. 
        if (op.responses![0].protocol.http!.headers) {
          for (const header of values(op.responses![0].protocol.http!.headers)) {
            const head = <LanguageHeader>header;
            // convert each header to a property and append it to the response properties list
            op.responses![0].language.go!.properties.push(newProperty(head.name, `${head.name} contains the information returned from the ${head.name} header response.`, <Schema>head.schema));
          }
        }
        structs.push(generateStruct(op.responses![0].language.go!, op.responses![0].language.go!.properties));
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
      let tag = ` \`${prop.schema.language.go!.marshallingFormat}:"${prop.serializedName},omitempty"\``;
      if (this.Language.responseType) {
        // tags aren't required for response types
        tag = '';
      }
      text += `\t${prop.language.go!.name} *${typeName}${tag}\n\n`;
    }
    text += '}\n\n';
    if (this.Language.errorType) {
      text += `func new${this.Language.name}(resp *azcore.Response) error {\n`;
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
