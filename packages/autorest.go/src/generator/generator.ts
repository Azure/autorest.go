/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { serialize } from '@azure-tools/codegen';
import { AutorestExtensionHost, startSession } from '@autorest/extension-base';
import { values } from '@azure-tools/linq';
import { Client, ConstantType, ConstantValue, GoCodeModel, HeaderMapResponse, HeaderResponse, InterfaceType, Method, ModelField, ModelType, PolymorphicType, ResponseEnvelope, StructField, StructType } from '../gocodemodel/gocodemodel';
import { generateClientFactory } from './clientFactory';
import { generateOperations } from './operations';
import { generateModels } from './models';
import { generateOptions } from './options';
import { generateInterfaces } from './interfaces';
import { generateResponses } from './responses';
import { generateConstants } from './constants';
import { generateTimeHelpers } from './time';
import { generatePolymorphicHelpers } from './polymorphics';
import { generateGoModFile } from './gomod';
import { generateXMLAdditionalPropsHelpers } from './xmlAdditionalProps';
import { generateServers } from './fake/servers';
import { generateServerFactory } from './fake/factory';
import { sortAscending } from './helpers';

// The generator emits Go source code files to disk.
export async function generateCode(host: AutorestExtensionHost) {
  const debug = await host.getValue('debug') || false;

  try {
    // get the code model from the core
    const session = await startSession<GoCodeModel>(host);
    sortContent(session.model);

    const operations = await generateOperations(session.model);
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

    const clientFactory = await generateClientFactory(session.model);
    if (clientFactory.length > 0) {
      host.writeFile({
        filename: `${filePrefix}client_factory.go`,
        content: clientFactory,
        artifactType: 'source-file-go'
      });
    }
    
    const constants = await generateConstants(session.model);
    host.writeFile({
      filename: `${filePrefix}constants.go`,
      content: constants,
      artifactType: 'source-file-go'
    });

    const models = await generateModels(session.model);
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

    const options = await generateOptions(session.model);
    if (options.length > 0) {
      host.writeFile({
        filename: `${filePrefix}options.go`,
        content: options,
        artifactType: 'source-file-go'
      });
    }

    const interfaces = await generateInterfaces(session.model);
    if (interfaces.length > 0) {
      host.writeFile({
        filename: `${filePrefix}interfaces.go`,
        content: interfaces,
        artifactType: 'source-file-go'
      });
    }

    const responses = await generateResponses(session.model);
    if (responses.length > 0) {
      host.writeFile({
        filename: `${filePrefix}response_types.go`,
        content: responses,
        artifactType: 'source-file-go'
      });
    }

    const timeHelpers = await generateTimeHelpers(session.model);
    for (const helper of values(timeHelpers)) {
      host.writeFile({
        filename: `${filePrefix}${helper.name.toLowerCase()}.go`,
        content: helper.content,
        artifactType: 'source-file-go'
      });
    }

    const polymorphics = await generatePolymorphicHelpers(session.model);
    if (polymorphics.length > 0) {
      host.writeFile({
        filename: `${filePrefix}polymorphic_helpers.go`,
        content: polymorphics,
        artifactType: 'source-file-go'
      });
    }

    // don't overwrite an existing go.mod file, update it if required
    const existingGoMod = await host.readFile('go.mod');
    const gomod = await generateGoModFile(session.model, existingGoMod);
    if (gomod.length > 0) {
      host.writeFile({
        filename: 'go.mod',
        content: gomod,
        artifactType: 'source-file-go'
      });
    }

    const xmlAddlProps = await generateXMLAdditionalPropsHelpers(session.model);
    if (xmlAddlProps.length > 0) {
      host.writeFile({
        filename: `${filePrefix}xml_helper.go`,
        content: xmlAddlProps,
        artifactType: 'source-file-go'
      });
    }

    if (session.model.options.generateFakes) {
      const serverContent = await generateServers(session.model);
      for (const op of values(serverContent.servers)) {
        let fileName = op.name.toLowerCase();
        // op.name is the server name, e.g. FooServer.
        // insert a _ before Server, i.e. Foo_Server
        // if the name isn't simply Server.
        if (fileName !== 'server') {
          fileName = fileName.substring(0, fileName.length-6) + '_server';
        }
        host.writeFile({
          filename: `fake/${filePrefix}${fileName}.go`,
          content: op.content,
          artifactType: 'source-file-go'
        });
      }

      const serverFactory = generateServerFactory(session.model);
      if (serverFactory !== '') {
        host.writeFile({
          filename: `fake/${filePrefix}server_factory.go`,
          content: serverFactory,
          artifactType: 'source-file-go'
        });
      }

      host.writeFile({
        filename: `fake/${filePrefix}internal.go`,
        content: serverContent.internals,
        artifactType: 'source-file-go'
      });

      const timeHelpers = await generateTimeHelpers(session.model, 'fake');
      for (const helper of values(timeHelpers)) {
        host.writeFile({
          filename: `fake/${filePrefix}${helper.name.toLowerCase()}.go`,
          content: helper.content,
          artifactType: 'source-file-go'
        });
      }

      const polymorphics = await generatePolymorphicHelpers(session.model, 'fake');
      if (polymorphics.length > 0) {
        host.writeFile({
          filename: `fake/${filePrefix}polymorphic_helpers.go`,
          content: polymorphics,
          artifactType: 'source-file-go'
        });
      }
    }
  } catch (E) {
    if (debug) {
      console.error(`${__filename} - FAILURE  ${JSON.stringify(E)} ${(<Error>E).stack}`);
    }
    throw E;
  }
}

function sortContent(codeModel: GoCodeModel) {
  codeModel.constants.sort((a: ConstantType, b: ConstantType) => { return sortAscending(a.name, b.name); });
  for (const enm of values(codeModel.constants)) {
    enm.values.sort((a: ConstantValue, b: ConstantValue) => { return sortAscending(a.valueName, b.valueName); });
  }

  codeModel.interfaceTypes.sort((a: InterfaceType, b: InterfaceType) => { return sortAscending(a.name, b.name); });
  for (const iface of values(codeModel.interfaceTypes)) {
    iface.possibleTypes.sort((a: PolymorphicType, b: PolymorphicType) => { return sortAscending(a.discriminatorValue!, b.discriminatorValue!); });
  }

  codeModel.models.sort((a: ModelType | PolymorphicType, b: ModelType | PolymorphicType) => { return sortAscending(a.name, b.name); });
  for (const model of values(codeModel.models)) {
    model.fields.sort((a: ModelField, b: ModelField) => { return sortAscending(a.fieldName, b.fieldName); });
  }

  codeModel.paramGroups.sort((a: StructType, b: StructType) => { return sortAscending(a.name, b.name); });
  for (const paramGroup of values(codeModel.paramGroups)) {
    paramGroup.fields.sort((a: StructField, b: StructField) => { return sortAscending(a.fieldName, b.fieldName); });
  }

  codeModel.responseEnvelopes.sort((a: ResponseEnvelope, b: ResponseEnvelope) => { return sortAscending(a.name, b.name); });
  for (const respEnv of values(codeModel.responseEnvelopes)) {
    respEnv.headers.sort((a: HeaderResponse | HeaderMapResponse, b: HeaderResponse | HeaderMapResponse) => { return sortAscending(a.fieldName, b.fieldName); });
  }

  codeModel.clients.sort((a: Client, b: Client) => { return sortAscending(a.clientName, b.clientName); });
  for (const client of values(codeModel.clients)) {
    client.methods.sort((a: Method, b: Method) => { return sortAscending(a.methodName, b.methodName); });
  }
}
