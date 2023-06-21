/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { comment } from '@azure-tools/codegen';
import { CodeModel, ObjectSchema, Schema } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { contentPreamble, sortAscending } from './helpers';

// Creates the content in interfaces.go
export async function generateInterfaces(session: Session<CodeModel>): Promise<string> {
  if (!session.model.language.go!.discriminators) {
    // no polymorphic types
    return '';
  }
  const rootDiscriminators = <Array<ObjectSchema>>session.model.language.go!.discriminators.filter((d: ObjectSchema) => !d.language.go!.omitType);
  if (rootDiscriminators.length === 0) {
    // all polymorphic types omitted
    return '';
  }

  const allDiscriminators = new Array<Schema>();

  // for each root, find any sub hierarchies and add them to the list
  for (const discriminator of rootDiscriminators) {
    allDiscriminators.push(discriminator);
    for (const childDiscriminator of values(discriminator.discriminator?.immediate)) {
      if (childDiscriminator.language.go!.discriminatorInterface) {
        allDiscriminators.push(childDiscriminator);
      }
    }
  }

  allDiscriminators.sort((a: Schema, b: Schema) => { return sortAscending(a.language.go!.name, b.language.go!.name); });

  let text = await contentPreamble(session);

  for (const discriminator of allDiscriminators) {
    const methodName = `Get${discriminator.language.go!.name}`;
    text += `// ${discriminator.language.go!.discriminatorInterface} provides polymorphic access to related types.\n`;
    text += `// Call the interface's ${methodName}() method to access the common type.\n`;
    text += '// Use a type switch to determine the concrete type.  The possible types are:\n';
    text += comment((<Array<string>>discriminator.language.go!.discriminatorTypes).join(', '), '// - ');
    text += `\ntype ${discriminator.language.go!.discriminatorInterface} interface {\n`;
    if (discriminator.language.go!.discriminatorParent) {
      text += `\t${discriminator.language.go!.discriminatorParent}\n`;
    }
    text += `\t// ${methodName} returns the ${discriminator.language.go!.name} content of the underlying type.\n`;
    text += `\t${methodName}() *${discriminator.language.go!.name}\n`;
    text += '}\n\n';
  }
  return text;
}
