/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment, pascalCase } from '@azure-tools/codegen';
import { CodeModel, ConstantSchema, ImplementationLocation, ObjectSchema, Language, SchemaType, Parameter, Property } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { ContentPreamble, HasDescription, ImportManager, SortAscending } from '../common/helpers';

// Creates the content in models.go
export async function generateModels(session: Session<CodeModel>): Promise<string> {
  let text = await ContentPreamble(session);

  // we do model generation first as it can add imports to the imports list
  const structs = generateStructs(session.model.schemas.objects);
  // add types from requests and responses
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      // add structs from optional operation params
      if (op.request.language.go!.optionalParam) {
        structs.push(generateOptionalParamsStruct(op.request.language.go!.optionalParam, op.request.language.go!.optionalParam.params));
      }
      // add structs from operation responses
      if (op.responses) {
        structs.push(generateStruct(op.responses[0].language.go!, op.responses[0].language.go!.properties));
      }
    }
  }

  // imports
  if (imports.length() > 0) {
    text += imports.text();
  }

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
  readonly Parameters?: Parameter[];

  constructor(language: Language, props?: Property[], params?: Parameter[]) {
    this.Language = language;
    this.Properties = props;
    this.Parameters = params;
    if (this.Properties) {
      this.Properties.sort((a: Property, b: Property) => { return SortAscending(a.language.go!.name, b.language.go!.name); });
    }
    if (this.Parameters) {
      this.Parameters.sort((a: Parameter, b: Parameter) => { return SortAscending(a.language.go!.name, b.language.go!.name); });
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
      text += `\t${prop.language.go!.name} *${typeName}${tag}\n`;
    }
    for (const param of values(this.Parameters)) {
      // if Parameters is set this is an optional args struct
      // none of its fields need to participate in marshalling
      if (param.implementation === ImplementationLocation.Client) {
        // don't add globals to the per-method options struct
        continue;
      }
      if (HasDescription(param.language.go!)) {
        text += `\t${comment(param.language.go!.description, '// ')}\n`;
      }
      text += `\t${pascalCase(param.language.go!.name)} *${param.schema.language.go!.name}\n`;
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

function generateOptionalParamsStruct(lang: Language, params: Parameter[]): StructDef {
  const st = new StructDef(lang, undefined, params);
  for (const param of values(params)) {
    imports.addImportForSchemaType(param.schema);
  }
  return st;
}
