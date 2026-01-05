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

    switch (this.codeModel.root.kind) {
      case 'containingModule':
        sortContent(this.codeModel.root.package);
        break;
      case 'module':
        sortContent(this.codeModel.root);
        break;
      default:
        this.codeModel.root satisfies never;
    }
  }

  /**
   * writes the core code model content.
   * this content is common to both emitters.
   */
  async emit(): Promise<void> {
    if (this.codeModel.root.kind === 'module') {
      // don't overwrite an existing go.mod file, update it if required
      const goModFile = 'go.mod';
      let existingGoMod: string | undefined;
      if (await this.fs.exists(goModFile)) {
        existingGoMod = await this.fs.read(goModFile);
      }
      const gomod = generateGoModFile(this.codeModel.root, this.codeModel.options, existingGoMod);
      if (gomod.length > 0) {
        await this.fs.write(goModFile, gomod);
      }
    }

    await this.recursiveEmit(async (pkg: go.PackageContent, write: (name: string, content: string, subdir?: string) => Promise<void>): Promise<void> => {
      const clientFactory = generateClientFactory(pkg, this.codeModel.type, this.codeModel.options);
      if (clientFactory.length > 0) {
        await write('client_factory.go', clientFactory);
      }

      const constants = generateConstants(pkg);
      if (constants.length > 0) {
        await write('constants.go', constants);
      }

      const interfaces = generateInterfaces(pkg);
      if (interfaces.length > 0) {
        await write('interfaces.go', interfaces);
      }

      const operations = generateOperations(pkg, this.codeModel.type, this.codeModel.options);
      for (const op of operations) {
        await write(`${snakeClientFileName(op.name)}.go`, op.content);
      }

      const models = generateModels(pkg, this.codeModel.options);
      if (models.models.length > 0) {
        await write('models.go', models.models);
      }
      if (models.serDe.length > 0) {
        await write('models_serde.go', models.serDe);
      }

      const options = generateOptions(pkg);
      if (options.length > 0) {
        await write('options.go', options);
      }

      const polymorphics = generatePolymorphicHelpers(pkg);
      if (polymorphics.length > 0) {
        await write('polymorphic_helpers.go', polymorphics);
      }

      const responses = generateResponses(pkg, this.codeModel.options);
      if (responses.responses.length > 0) {
        await write('responses.go', responses.responses);
      }
      if (responses.serDe.length > 0) {
        await write('responses_serde.go', responses.serDe);
      }

      const xmlAddlProps = generateXMLAdditionalPropsHelpers(pkg);
      if (xmlAddlProps.length > 0) {
        await write('xml_helper.go', xmlAddlProps);
      }

      if (this.codeModel.options.generateFakes) {
        const fakePkg = new go.FakePackage(pkg);
        const serverContent = generateServers(fakePkg);
        if (serverContent.servers.length > 0) {
          for (const op of serverContent.servers) {
            const fileName = `${snakeClientFileName(op.name, 'server')}.go`;
            await write(fileName, op.content, fakePkg.kind);
          }

          const serverFactory = generateServerFactory(fakePkg, this.codeModel.type);
          if (serverFactory.length > 0) {
            await write('server_factory.go', serverFactory, fakePkg.kind);
          }

          await write('internal.go', serverContent.internals, fakePkg.kind);

          const polymorphics = generatePolymorphicHelpers(fakePkg);
          if (polymorphics.length > 0) {
            await write('polymorphic_helpers.go', polymorphics, fakePkg.kind);
          }
        }
      }
    });

    // only one version.go file per module
    if (this.codeModel.root.kind === 'module') {
      // don't overwrite an existing version.go file
      const versionGo = generateVersionInfo(this.codeModel.root);
      const versionGoFileName = `${this.filePrefix}version.go`;
      if (versionGo.length > 0 && !await this.fs.exists(versionGoFileName)) {
        await this.fs.write(versionGoFileName, versionGo);
      }
    }
  }

  /** writes the cloud_config.go file */
  async emitCloudConfig(): Promise<void> {
    if (this.codeModel.root.kind !== 'module') {
      return;
    }
    const cloudConfig = generateCloudConfig(this.codeModel.root, this.codeModel.type);
    if (cloudConfig.length > 0) {
      await this.fs.write(`${this.filePrefix}cloud_config.go`, cloudConfig);
    }
  }

  /** writes the *_example_test.go files */
  async emitExamples(): Promise<void> {
    if (!this.codeModel.options.generateExamples) {
      return;
    }

    await this.recursiveEmit(async (pkg: go.PackageContent, write: (name: string, content: string) => Promise<void>): Promise<void> => {
      const examples = generateExamples(new go.TestPackage(pkg), this.codeModel.type, this.codeModel.options);
      for (const example of examples) {
        await write(`${snakeClientFileName(example.name)}_example_test.go`, example.content);
      }
    });
  }

  /** writes the LICENSE.txt file */
  async emitLicenseFile(): Promise<void> {
    if (this.codeModel.root.kind !== 'module') {
      return;
    }

    // don't overwrite an existing LICENSE.txt file
    const licenseTxt = generateLicenseTxt(this.codeModel.options);
    const licenseTxtFileName = 'LICENSE.txt';
    if (licenseTxt && !await this.fs.exists(licenseTxtFileName)) {
      await this.fs.write(licenseTxtFileName, licenseTxt);
    }
  }

  /** writes the emitter metadata file */
  async emitMetadataFile(): Promise<void> {
    if (this.codeModel.root.kind !== 'module') {
      return;
    }

    const metadata = generateMetadataFile(this.codeModel);
    if (metadata.length > 0) {
      await this.fs.write('testdata/_metadata.json', metadata);
    }
  }

  /**
   * recursively emits package contents.
   * 
   * @param emitForPkg the package contents to emit
   */
  private async recursiveEmit(emitForPkg: (pkg: go.PackageContent, write: (name: string, content: string, subdir?: string) => Promise<void>) => Promise<void>): Promise<void> {
    const recursiveEmit = async (pkg: go.PackageContent, dir: string): Promise<void> => {
      await emitForPkg(pkg, async (name: string, content: string, subdir?: string) => {
        return await this.fs.write(`${dir}${subdir ? `${subdir}/` : ''}${this.filePrefix}${name}`, content);
      });

      // recursively emit any sub-packages
      for (const subPkg of pkg.packages) {
        await recursiveEmit(subPkg, `${dir}${subPkg.name}/`);
      }
    };

    switch (this.codeModel.root.kind) {
      case 'containingModule':
        await recursiveEmit(this.codeModel.root.package, '');
        break;
      case 'module':
        // when emitting a module, the root directory is the module directory
        await recursiveEmit(this.codeModel.root, '');
        break;
      default:
        this.codeModel.root satisfies never;
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
 * recursively sorts code model contents by name in alphabetical order.
 * 
 * @param pkg the contents to sort
 */
function sortContent(pkg: go.PackageContent): void {
  const sortAscending = function(a: string, b: string): number {
    return a < b ? -1 : a > b ? 1 : 0;
  };

  pkg.constants.sort((a: go.Constant, b: go.Constant) => { return sortAscending(a.name, b.name); });
  for (const enm of pkg.constants) {
    enm.values.sort((a: go.ConstantValue, b: go.ConstantValue) => { return sortAscending(a.name, b.name); });
  }

  pkg.interfaces.sort((a: go.Interface, b: go.Interface) => { return sortAscending(a.name, b.name); });
  for (const iface of pkg.interfaces) {
    // we sort by literal value so that the switch/case statements in polymorphic_helpers.go
    // are ordered by the literal value which can be somewhat different from the model name.
    iface.possibleTypes.sort((a: go.PolymorphicModel, b: go.PolymorphicModel) => { return sortAscending(a.discriminatorValue!.literal, b.discriminatorValue!.literal); });
  }

  pkg.models.sort((a: go.Model | go.PolymorphicModel, b: go.Model | go.PolymorphicModel) => { return sortAscending(a.name, b.name); });
  for (const model of pkg.models) {
    model.fields.sort((a: go.ModelField, b: go.ModelField) => { return sortAscending(a.name, b.name); });
  }

  pkg.paramGroups.sort((a: go.Struct, b: go.Struct) => { return sortAscending(a.name, b.name); });
  for (const paramGroup of pkg.paramGroups) {
    paramGroup.fields.sort((a: go.StructField, b: go.StructField) => { return sortAscending(a.name, b.name); });
  }

  pkg.responseEnvelopes.sort((a: go.ResponseEnvelope, b: go.ResponseEnvelope) => { return sortAscending(a.name, b.name); });
  for (const respEnv of pkg.responseEnvelopes) {
    respEnv.headers.sort((a: go.HeaderScalarResponse | go.HeaderMapResponse, b: go.HeaderScalarResponse | go.HeaderMapResponse) => { return sortAscending(a.fieldName, b.fieldName); });
  }

  pkg.clients.sort((a: go.Client, b: go.Client) => { return sortAscending(a.name, b.name); });
  for (const client of pkg.clients) {
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

  for (const subPkg of pkg.packages) {
    sortContent(subPkg);
  }
}
