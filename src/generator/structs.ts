/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { capitalize, comment } from '@azure-tools/codegen';
import { ConstantSchema, ImplementationLocation, Language, SchemaType, Parameter, Property } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { commentLength, isArraySchema } from '../common/helpers';
import { elementByValueForParam, hasDescription, sortAscending, substituteDiscriminator } from './helpers';
import { ImportManager } from './imports';

// represents a struct method
export interface StructMethod {
  name: string;
  desc: string;
  text: string;
}

// represents a struct definition
export class StructDef {
  readonly Language: Language;
  readonly Properties?: Property[];
  readonly Parameters?: Parameter[];
  readonly SerDeMethods: StructMethod[];
  readonly Methods: StructMethod[];
  readonly ComposedOf: string[];
  HasJSONByteArray: boolean;

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
    this.SerDeMethods = new Array<StructMethod>();
    this.Methods = new Array<StructMethod>();
    this.ComposedOf = new Array<string>();
    this.HasJSONByteArray = false;
  }

  text(): string {
    let text = '';
    if (hasDescription(this.Language)) {
      text += `${comment(this.Language.description, '// ', undefined, commentLength)}\n`;
    }
    text += `type ${this.Language.name} struct {\n`;
    // any composed types go first
    for (const comp of values(this.ComposedOf)) {
      text += `\t${comp}\n`;
    }
    // used to track when to add an extra \n between fields that have comments
    let first = true;
    if (this.Properties === undefined && this.Parameters?.length === 0) {
      // this is an optional params placeholder struct
      text += '\t// placeholder for future optional parameters\n';
    } else if (this.Properties === undefined && this.Parameters === undefined) {
      // this is an empty response envelope
      text += '\t// placeholder for future response values\n';
    }
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
      let elemByVal = false;
      if (prop.schema.type === SchemaType.Dictionary && prop.extensions?.['x-ms-header-collection-prefix']) {
        elemByVal = true;
      }
      let typeName = substituteDiscriminator(prop.schema, elemByVal);
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
        readOnly = ` azure:"ro"`;
      }
      let tag = ` \`${this.Language.marshallingFormat}:"${serialization}"${readOnly}\``;
      // if this is a response type then omit the tag IFF the marshalling format is
      // JSON, or it's a header.  XML marshalling needs a tag.
      // also omit the tag for additionalProperties
      if ((this.Language.responseType === true && (this.Language.marshallingFormat !== 'xml')) || prop.language.go!.isAdditionalProperties) {
        tag = '';
      }
      let pointer = '*';
      if (prop.language.go!.byValue === true) {
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
        text += `\t${comment(param.language.go!.description, '// ', undefined, commentLength)}\n`;
      }
      let pointer = '*';
      if (param.required || param.language.go!.byValue === true) {
        pointer = '';
      }
      const typeName = substituteDiscriminator(param.schema, elementByValueForParam(param));
      text += `\t${capitalize(param.language.go!.name)} ${pointer}${typeName}\n`;
    }
    text += '}\n\n';
    return text;
  }

  discriminator(): string {
    if (!this.Language.discriminatorInterface) {
      return '';
    }
    const methodName = `Get${this.Language.name}`;
    let text = `// ${this.Language.discriminatorInterface} provides polymorphic access to related types.\n`;
    text += `// Call the interface's ${methodName}() method to access the common type.\n`;
    text += `// Use a type switch to determine the concrete type.  The possible types are:\n`;
    text += comment((<Array<string>>this.Language.discriminatorTypes).join(', '), '// - ');
    text += `\ntype ${this.Language.discriminatorInterface} interface {\n`;
    if (this.Language.discriminatorParent) {
      text += `\t${this.Language.discriminatorParent}\n`;
    }
    text += `\t// ${methodName} returns the ${this.Language.name} content of the underlying type.\n`;
    text += `\t${methodName}() *${this.Language.name}\n`;
    text += '}\n\n';
    return text;
  }

  receiverName(): string {
    const typeName = this.Language.name;
    return typeName[0].toLowerCase();
  }
}

export function generateStruct(imports: ImportManager, lang: Language, props?: Property[]): StructDef {
  if (lang.isLRO) {
    imports.add('time');
    imports.add('context');
  }
  const st = new StructDef(lang, props);
  for (const prop of values(props)) {
    imports.addImportForSchemaType(prop.schema);
    if (prop.language.go!.embeddedType) {
      st.ComposedOf.push(substituteDiscriminator(prop.schema, true));
    }
  }
  return st;
}

export function getXMLSerialization(prop: Property, lang: Language): string {
  let serialization = prop.serializedName;
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
    if (prop.schema.serialization?.xml?.wrapped && lang.responseType !== true) {
      serialization += `>${inner}`;
    } else {
      serialization = inner;
    }
  }
  return serialization;
}
