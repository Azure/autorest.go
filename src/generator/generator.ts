/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { serialize } from '@azure-tools/codegen';
import { AutorestExtensionHost, startSession } from '@autorest/extension-base';
import { codeModelSchema, CodeModel } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { generateOperations } from './operations';
import { generateModels } from './models';
import { generateResponses } from './responses';
import { generateConstants } from './constants';
import { generateTimeHelpers } from './time';
import { generatePolymorphicHelpers } from './polymorphics';
import { generateGoModFile } from './gomod';
import { generateXMLAdditionalPropsHelpers } from './xmlAdditionalProps';

// The generator emits Go source code files to disk.
export async function protocolGen(host: AutorestExtensionHost) {
  const debug = await host.getValue('debug') || false;

  try {
    // get the code model from the core
    const session = await startSession<CodeModel>(host, codeModelSchema);

    const operations = await generateOperations(session);
    let filePrefix = await session.getValue('file-prefix', '');
    // if a file prefix was specified, ensure it's properly snaked
    if (filePrefix.length > 0 && filePrefix[filePrefix.length - 1] !== '_') {
      filePrefix += '_';
    }

    // output the model to the pipeline.  this must happen after all model
    // updates are complete and before any source files are written.
    host.writeFile({
      filename: 'code-model-v4.yaml',
      content: serialize(session.model),
      artifactType: 'code-model-v4'
    });

    for (const op of values(operations)) {
      let fileName = op.name.toLowerCase();
      // op.name is the client name, e.g. FooClient.
      // insert a _ before Client, i.e. Foo_Client
      // if the name isn't simply Client.
      if (fileName !== 'client') {
        fileName = fileName.substring(0, fileName.length-6) + '_client';
      }
      host.writeFile({
        filename: `${filePrefix}${fileName}.go`,
        content: op.content,
        artifactType: 'source-file-go'
      });
    }

    const constants = await generateConstants(session);
    host.writeFile({
      filename: `${filePrefix}constants.go`,
      content: constants,
      artifactType: 'source-file-go'
    });

    const models = await generateModels(session);
    host.writeFile({
      filename: `${filePrefix}models.go`,
      content: models.models,
      artifactType: 'source-file-go'
    });
    if (models.serDe.length > 0) {
      host.writeFile({
        filename: `${filePrefix}models_serde.go`,
        content: models.serDe,
        artifactType: 'source-file-go'
      });
    }

    const responses = await generateResponses(session);
    if (responses.length > 0) {
      host.writeFile({
        filename: `${filePrefix}response_types.go`,
        content: responses,
        artifactType: 'source-file-go'
      });
    }

    const timeHelpers = await generateTimeHelpers(session);
    for (const helper of values(timeHelpers)) {
      host.writeFile({
        filename: `${filePrefix}${helper.name.toLowerCase()}.go`,
        content: helper.content,
        artifactType: 'source-file-go'
      });
    }

    const polymorphics = await generatePolymorphicHelpers(session);
    if (polymorphics.length > 0) {
      host.writeFile({
        filename: `${filePrefix}polymorphic_helpers.go`,
        content: polymorphics,
        artifactType: 'source-file-go'
      });
    }

    // don't overwrite an existing go.mod file, update it if required
    const existingGoMod = await host.readFile('go.mod');
    const gomod = await generateGoModFile(session, existingGoMod);
    if (gomod.length > 0) {
      host.writeFile({
        filename: 'go.mod',
        content: gomod,
        artifactType: 'source-file-go'
      });
    }

    const xmlAddlProps = await generateXMLAdditionalPropsHelpers(session);
    if (xmlAddlProps.length > 0) {
      host.writeFile({
        filename: `${filePrefix}xml_helper.go`,
        content: xmlAddlProps,
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
