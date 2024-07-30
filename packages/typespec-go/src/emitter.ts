/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { EmitContext } from '@typespec/compiler';
import { existsSync } from 'fs';
import { mkdir, readFile, writeFile } from 'fs/promises';
import 'source-map-support/register.js';
import { generateClientFactory } from '../../codegen.go/src/clientFactory.js';
import { generateConstants } from '../../codegen.go/src/constants.js';
import { generateExamples } from '../../codegen.go/src/example.js';
import { generateServerFactory } from '../../codegen.go/src/fake/factory.js';
import { generateServers } from '../../codegen.go/src/fake/servers.js';
import { generateGoModFile } from '../../codegen.go/src/gomod.js';
import { generateInterfaces } from '../../codegen.go/src/interfaces.js';
import { generateModels } from '../../codegen.go/src/models.js';
import { generateOperations } from '../../codegen.go/src/operations.js';
import { generateOptions } from '../../codegen.go/src/options.js';
import { generatePolymorphicHelpers } from '../../codegen.go/src/polymorphics.js';
import { generateResponses } from '../../codegen.go/src/responses.js';
import { generateTimeHelpers } from '../../codegen.go/src/time.js';
import { GoEmitterOptions } from './lib.js';
import { tcgcToGoCodeModel } from './tcgcadapter/adapter.js';

export async function $onEmit(context: EmitContext<GoEmitterOptions>) {
  const codeModel = await tcgcToGoCodeModel(context);
  await mkdir(context.emitterOutputDir, {recursive: true});

  // don't overwrite an existing go.mod file, update it if required
  const goModFile = `${context.emitterOutputDir}/go.mod`;
  let existingGoMod: string | undefined;
  if (existsSync(goModFile)) {
    existingGoMod = (await readFile(goModFile)).toString();
  }
  const gomod = await generateGoModFile(codeModel, existingGoMod!);
  if (gomod.length > 0) {
    writeFile(goModFile, gomod);
  }

  let filePrefix = '';
  if (context.options['file-prefix']) {
    filePrefix = context.options['file-prefix'];
    // if a file prefix was specified, ensure it's properly snaked
    if (filePrefix[filePrefix.length - 1] !== '_') {
      filePrefix += '_';
    }
  }

  const clientFactory = await generateClientFactory(codeModel);
  if (clientFactory.length > 0) {
    writeFile(`${context.emitterOutputDir}/${filePrefix}client_factory.go`, clientFactory);
  }

  const constants = await generateConstants(codeModel);
  if (constants.length > 0) {
    writeFile(`${context.emitterOutputDir}/${filePrefix}constants.go`, constants);
  }

  const interfaces = await generateInterfaces(codeModel);
  if (interfaces.length > 0) {
    writeFile(`${context.emitterOutputDir}/${filePrefix}interfaces.go`, interfaces);
  }

  const models = await generateModels(codeModel);
  if (models.models.length > 0) {
    writeFile(`${context.emitterOutputDir}/${filePrefix}models.go`, models.models);
  }
  if (models.serDe.length > 0) {
    writeFile(`${context.emitterOutputDir}/${filePrefix}models_serde.go`, models.serDe);
  }

  const operations = await generateOperations(codeModel);
  for (const op of operations) {
    let fileName = op.name.toLowerCase();
    // op.name is the client name, e.g. FooClient.
    // insert a _ before Client, i.e. Foo_Client
    // if the name isn't simply Client.
    if (fileName !== 'client') {
      fileName = fileName.substring(0, fileName.length - 6) + '_client';
    }
    writeFile(`${context.emitterOutputDir}/${filePrefix}${fileName}.go`, op.content);
  }

  if (codeModel.options.generateExamples) {
    const examples = await generateExamples(codeModel);
    for (const example of examples) {
      let fileName = example.name.toLowerCase();
      // op.name is the client name, e.g. FooClient.
      // insert a _ before Client, i.e. Foo_Client
      // if the name isn't simply Client.
      // and insert _example_test at the end.
      if (fileName !== 'client') {
        fileName = fileName.substring(0, fileName.length - 6) + '_client';
      }
      fileName += '_example_test';
      writeFile(`${context.emitterOutputDir}/${filePrefix}${fileName}.go`, example.content);
    }
  }

  const options = await generateOptions(codeModel);
  if (options.length > 0) {
    writeFile(`${context.emitterOutputDir}/${filePrefix}options.go`, options);
  }

  const polymorphics = await generatePolymorphicHelpers(codeModel);
  if (polymorphics.length > 0) {
    writeFile(`${context.emitterOutputDir}/${filePrefix}polymorphic_helpers.go`, polymorphics);
  }

  const responses = await generateResponses(codeModel);
  if (responses.responses.length > 0) {
    writeFile(`${context.emitterOutputDir}/${filePrefix}responses.go`, responses.responses);
  }
  if (responses.serDe.length > 0) {
    writeFile(`${context.emitterOutputDir}/${filePrefix}responses_serde.go`, responses.serDe);
  }

  const timeHelpers = await generateTimeHelpers(codeModel);
  for (const helper of timeHelpers) {
    writeFile(`${context.emitterOutputDir}/${filePrefix}${helper.name.toLowerCase()}.go`, helper.content);
  }

  if (codeModel.options.generateFakes) {
    const serverContent = await generateServers(codeModel);
    if (serverContent.servers.length > 0) {
      const fakesDir = context.emitterOutputDir + '/fake';
      await mkdir(fakesDir, { recursive: true });
      for (const op of serverContent.servers) {
        let fileName = op.name.toLowerCase();
        // op.name is the server name, e.g. FooServer.
        // insert a _ before Server, i.e. Foo_Server
        // if the name isn't simply Server.
        if (fileName !== 'server') {
          fileName = fileName.substring(0, fileName.length - 6) + '_server';
        }
        writeFile(`${fakesDir}/${filePrefix}${fileName}.go`, op.content);
      }

      const serverFactory = generateServerFactory(codeModel);
      if (serverFactory.length > 0) {
        writeFile(`${fakesDir}/${filePrefix}server_factory.go`, serverFactory);
      }

      writeFile(`${fakesDir}/${filePrefix}internal.go`, serverContent.internals);

      const timeHelpers = await generateTimeHelpers(codeModel, 'fake');
      for (const helper of timeHelpers) {
        writeFile(`${fakesDir}/${filePrefix}${helper.name.toLowerCase()}.go`, helper.content);
      }

      const polymorphics = await generatePolymorphicHelpers(codeModel, 'fake');
      if (polymorphics.length > 0) {
        writeFile(`${fakesDir}/${filePrefix}polymorphic_helpers.go`, polymorphics);
      }
    }
  }
}
