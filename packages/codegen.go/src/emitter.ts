/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import * as go from '../../codemodel.go/src/index.js';
import { generateClientFactory } from './core/clientFactory.js';
import { generateCloudConfig } from './core/cloudConfig.js';
import { generateConstants } from './core/constants.js';
import { generateExamples } from './core/example.js';
import { generateGoModFile } from './core/gomod.js';
import { setCustomHeaderText } from './core/helpers.js';
import { generateInterfaces } from './core/interfaces.js';
import { generateLicenseTxt } from './core/license.js';
import { generateMetadataFile } from './core/metadata.js';
import { generateModels } from './core/models.js';
import { generateOperations } from './core/operations.js';
import { generateOptions } from './core/options.js';
import { generatePolymorphicHelpers } from './core/polymorphics.js';
import { generateResponses } from './core/responses.js';
import { generateTimeHelpers } from './core/time.js';
import { generateVersionInfo } from './core/version.js';
import { generateXMLAdditionalPropsHelpers } from './core/xmlAdditionalProps.js';
import { generateServers } from './fake/servers.js';
import { generateServerFactory } from './fake/factory.js';

/** abstractions over various file handling facilities */
export interface FsFacilities {
  /** checks if a file exists */
  exists: (name: string) => Promise<boolean>;

  /**
   * reads the contents of a file.
   * this function assumes the file exists.
   */
  read: (name: string) => Promise<string>;

  /** writes contents to a file */
  write: (name: string, content: string) => Promise<void>;
}

/** used to write a Go code model to disk */
export class Emitter {
  private readonly codeModel: go.CodeModel;
  private readonly fs: FsFacilities;
  private readonly filePrefix: string;

  constructor(codeModel: go.CodeModel, fs: FsFacilities, options?: { filePrefix?: string }) {
    this.codeModel = codeModel;
    this.fs = fs;
    if (options?.filePrefix) {
      this.filePrefix = options?.filePrefix;
    } else {
      this.filePrefix = '';
    }
    if (this.codeModel.options.headerText) {
      setCustomHeaderText(this.codeModel.options.headerText);
    }
    sortContent(this.codeModel);
  }

  /**
   * writes the core code model content.
   * this content is common to both emitters.
   */
  async emit(): Promise<void> {
    const clientFactory = generateClientFactory(this.codeModel);
    if (clientFactory.length > 0) {
      await this.fs.write(`${this.filePrefix}client_factory.go`, clientFactory);
    }

    const constants = generateConstants(this.codeModel);
    if (constants.length > 0) {
      await this.fs.write(`${this.filePrefix}constants.go`, constants);
    }

    const operations = generateOperations(this.codeModel);
    for (const op of operations) {
      const fileName = `${snakeClientFileName(op.name)}.go`;
      await this.fs.write(`${this.filePrefix}${fileName}`, op.content);
    }

    // don't overwrite an existing go.mod file, update it if required
    const goModFile = 'go.mod';
    let existingGoMod: string | undefined;
    if (await this.fs.exists(goModFile)) {
      existingGoMod = await this.fs.read(goModFile);
    }
    const gomod = generateGoModFile(this.codeModel, existingGoMod);
    if (gomod.length > 0) {
      await this.fs.write(goModFile, gomod);
    }

    const interfaces = generateInterfaces(this.codeModel);
    if (interfaces.length > 0) {
      await this.fs.write(`${this.filePrefix}interfaces.go`, interfaces);
    }

    const models = generateModels(this.codeModel);
    if (models.models.length > 0) {
      await this.fs.write(`${this.filePrefix}models.go`, models.models);
    }
    if (models.serDe.length > 0) {
      await this.fs.write(`${this.filePrefix}models_serde.go`, models.serDe);
    }

    const options = generateOptions(this.codeModel);
    if (options.length > 0) {
      await this.fs.write(`${this.filePrefix}options.go`, options);
    }

    const polymorphics = generatePolymorphicHelpers(this.codeModel);
    if (polymorphics.length > 0) {
      await this.fs.write(`${this.filePrefix}polymorphic_helpers.go`, polymorphics);
    }

    const responses = generateResponses(this.codeModel);
    if (responses.responses.length > 0) {
      await this.fs.write(`${this.filePrefix}responses.go`, responses.responses);
    }
    if (responses.serDe.length > 0) {
      await this.fs.write(`${this.filePrefix}responses_serde.go`, responses.serDe);
    }

    const timeHelpers = generateTimeHelpers(this.codeModel);
    for (const helper of timeHelpers) {
      await this.fs.write(`${this.filePrefix}${helper.name.toLowerCase()}.go`, helper.content);
    }

    // don't overwrite an existing version.go file
    const versionGo = generateVersionInfo(this.codeModel);
    const versionGoFileName = `${this.filePrefix}version.go`;
    if (versionGo.length > 0 && !await this.fs.exists(versionGoFileName)) {
      await this.fs.write(versionGoFileName, versionGo);
    }

    const xmlAddlProps = generateXMLAdditionalPropsHelpers(this.codeModel);
    if (xmlAddlProps.length > 0) {
      await this.fs.write(`${this.filePrefix}xml_helper.go`, xmlAddlProps);
    }

    if (this.codeModel.options.generateFakes) {
      const serverContent = generateServers(this.codeModel);
      if (serverContent.servers.length > 0) {
        const fakesDir = 'fake';
        for (const op of serverContent.servers) {
          const fileName = `${snakeClientFileName(op.name, 'server')}.go`;
          await this.fs.write(`${fakesDir}/${this.filePrefix}${fileName}`, op.content);
        }

        const serverFactory = generateServerFactory(this.codeModel);
        if (serverFactory.length > 0) {
          await this.fs.write(`${fakesDir}/${this.filePrefix}server_factory.go`, serverFactory);
        }

        await this.fs.write(`${fakesDir}/${this.filePrefix}internal.go`, serverContent.internals);

        const timeHelpers = generateTimeHelpers(this.codeModel, 'fake');
        for (const helper of timeHelpers) {
          await this.fs.write(`${fakesDir}/${this.filePrefix}${helper.name.toLowerCase()}.go`, helper.content);
        }

        const polymorphics = generatePolymorphicHelpers(this.codeModel, 'fake');
        if (polymorphics.length > 0) {
          await this.fs.write(`${fakesDir}/${this.filePrefix}polymorphic_helpers.go`, polymorphics);
        }
      }
    }
  }

  /** writes the cloud_config.go file */
  async emitCloudConfig(): Promise<void> {
    const cloudConfig = generateCloudConfig(this.codeModel);
    if (cloudConfig.length > 0) {
      await this.fs.write(`${this.filePrefix}cloud_config.go`, cloudConfig);
    }
  }

  /** writes the *_example_test.go files */
  async emitExamples(): Promise<void> {
    if (!this.codeModel.options.generateExamples) {
      return;
    }

    const examples = generateExamples(this.codeModel);
    for (const example of examples) {
      const fileName = `${snakeClientFileName(example.name)}_example_test.go`;
      await this.fs.write(`${this.filePrefix}${fileName}`, example.content);
    }
  }

  /** writes the LICENSE.txt file */
  async emitLicenseFile(): Promise<void> {
    // don't overwrite an existing LICENSE.txt file
    const licenseTxt = generateLicenseTxt(this.codeModel);
    const licenseTxtFileName = 'LICENSE.txt';
    if (licenseTxt && !await this.fs.exists(licenseTxtFileName)) {
      await this.fs.write(licenseTxtFileName, licenseTxt);
    }
  }

  /** writes the emitter metadata file */
  async emitMetadataFile(): Promise<void> {
    const metadata = generateMetadataFile(this.codeModel);
    if (metadata.length > 0) {
      await this.fs.write('testdata/_metadata.json', metadata);
    }
  }
}

/**
 * creates a snake cased client file name from the provided client name.
 * the returned name will not contain the .go suffix.
 * 
 * @param clientName the name of the client (e.g. FooClient)
 * @param suffix the client name suffix. the default value is 'client'
 * @returns a snaked client file name prefix (e.g. foo_client)
 */
function snakeClientFileName(clientName: string, suffix: string = 'client'): string {
  clientName = clientName.toLowerCase();
  // fileName is the client name, e.g. FooClient.
  // insert a _ before Client, i.e. Foo_Client
  // if the name isn't simply Client.
  if (clientName !== suffix) {
    clientName = `${clientName.substring(0, clientName.length - suffix.length)}_${suffix}`;
  }
  return clientName;
}

/**
 * sorts code model contents by name in alphabetical order.
 * 
 * @param codeModel the contents to sort
 */
function sortContent(codeModel: go.CodeModel): void {
  const sortAscending = function(a: string, b: string): number {
    return a < b ? -1 : a > b ? 1 : 0;
  };

  codeModel.constants.sort((a: go.Constant, b: go.Constant) => { return sortAscending(a.name, b.name); });
  for (const enm of codeModel.constants) {
    enm.values.sort((a: go.ConstantValue, b: go.ConstantValue) => { return sortAscending(a.name, b.name); });
  }

  codeModel.interfaces.sort((a: go.Interface, b: go.Interface) => { return sortAscending(a.name, b.name); });
  for (const iface of codeModel.interfaces) {
    // we sort by literal value so that the switch/case statements in polymorphic_helpers.go
    // are ordered by the literal value which can be somewhat different from the model name.
    iface.possibleTypes.sort((a: go.PolymorphicModel, b: go.PolymorphicModel) => { return sortAscending(a.discriminatorValue!.literal, b.discriminatorValue!.literal); });
  }

  codeModel.models.sort((a: go.Model | go.PolymorphicModel, b: go.Model | go.PolymorphicModel) => { return sortAscending(a.name, b.name); });
  for (const model of codeModel.models) {
    model.fields.sort((a: go.ModelField, b: go.ModelField) => { return sortAscending(a.name, b.name); });
  }

  codeModel.paramGroups.sort((a: go.Struct, b: go.Struct) => { return sortAscending(a.name, b.name); });
  for (const paramGroup of codeModel.paramGroups) {
    paramGroup.fields.sort((a: go.StructField, b: go.StructField) => { return sortAscending(a.name, b.name); });
  }

  codeModel.responseEnvelopes.sort((a: go.ResponseEnvelope, b: go.ResponseEnvelope) => { return sortAscending(a.name, b.name); });
  for (const respEnv of codeModel.responseEnvelopes) {
    respEnv.headers.sort((a: go.HeaderScalarResponse | go.HeaderMapResponse, b: go.HeaderScalarResponse | go.HeaderMapResponse) => { return sortAscending(a.fieldName, b.fieldName); });
  }

  codeModel.clients.sort((a: go.Client, b: go.Client) => { return sortAscending(a.name, b.name); });
  for (const client of codeModel.clients) {
    if (client.instance?.kind === 'constructable') {
      client.instance.constructors.sort((a: go.Constructor, b: go.Constructor) => sortAscending(a.name, b.name));
      if (client.instance.options.kind === 'clientOptions') {
        client.instance.options.parameters.sort((a: go.ClientParameter, b: go.ClientParameter) => sortAscending(a.name, b.name));
      }
    }
    client.parameters.sort((a: go.ClientParameter, b: go.ClientParameter) => sortAscending(a.name, b.name));
    client.methods.sort((a: go.MethodType, b: go.MethodType) => { return sortAscending(a.name, b.name); });
    client.clientAccessors.sort((a: go.ClientAccessor, b: go.ClientAccessor) => { return sortAscending(a.name, b.name); });
    for (const method of client.methods) {
      method.httpStatusCodes.sort();
    }
  }
}
