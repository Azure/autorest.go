/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { AutorestExtensionHost, startSession } from '@autorest/extension-base';
import { codeModelSchema, CodeModel } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { generateServers } from './servers';
import { generateInternal } from './internal';
import { generateTimeHelpers } from '../generator/time';
import { generatePolymorphicHelpers } from '../generator/polymorphics';

// The generator emits Go source code files to disk.
export async function fakeGen(host: AutorestExtensionHost) {
    const debug = await host.getValue('debug') || false;

    try {
      // get the code model from the core
      const session = await startSession<CodeModel>(host, codeModelSchema);
      const generateFakes = await session.getValue('generate-fakes', false);
      if (!generateFakes) {
        return;
      }
  
      const operations = await generateServers(session);
      let filePrefix = await session.getValue('file-prefix', '');
      // if a file prefix was specified, ensure it's properly snaked
      if (filePrefix.length > 0 && filePrefix[filePrefix.length - 1] !== '_') {
        filePrefix += '_';
      }
  
      for (const op of values(operations)) {
        let fileName = op.name.toLowerCase();
        // op.name is the client name, e.g. FooClient.
        // insert a _ before Client, i.e. Foo_Client
        // if the name isn't simply Client.
        if (fileName !== 'server') {
          fileName = fileName.substring(0, fileName.length-6) + '_server';
        }
        host.writeFile({
          filename: `fake/${filePrefix}${fileName}.go`,
          content: op.content,
          artifactType: 'source-file-go'
        });
      }

      const internal = await generateInternal(session);
      host.writeFile({
        filename: `fake/${filePrefix}internal.go`,
        content: internal,
        artifactType: 'source-file-go'
      });

      const timeHelpers = await generateTimeHelpers(session, 'fake');
      for (const helper of values(timeHelpers)) {
        host.writeFile({
          filename: `fake/${filePrefix}${helper.name.toLowerCase()}.go`,
          content: helper.content,
          artifactType: 'source-file-go'
        });
      }

      const polymorphics = await generatePolymorphicHelpers(session, 'fake');
      if (polymorphics.length > 0) {
        host.writeFile({
          filename: `fake/${filePrefix}polymorphic_helpers.go`,
          content: polymorphics,
          artifactType: 'source-file-go'
        });
      }
    } catch (E) {
      if (debug) {
        console.error(`${__filename} - FAILURE  ${JSON.stringify(E)} ${(<Error>E).stack}`);
      }
      throw E;
    }
}
