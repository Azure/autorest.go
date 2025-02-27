/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { configDefaults as baseConfigDefaults } from '@autorest/testmodeler/dist/src/common/constant';

export enum Config {
  exportCodemodel = 'testmodeler.export-codemodel',
  generateMockTest = 'testmodeler.generate-mock-test',
  generateSdkExample = 'testmodeler.generate-sdk-example',
  generateScenarioTest = 'testmodeler.generate-scenario-test',
  generateSdkSample = 'testmodeler.generate-sdk-sample',
  generateFakeTest = 'testmodeler.generate-fake-test',
  parents = '__parents',
  outputFolder = 'output-folder',
  module = 'module',
  moduleVersion = 'module-version',
  filePrefix = 'file-prefix',
  exampleFilePrefix = 'example-file-prefix',
  testFilePrefix = 'test-file-prefix',
  sendExampleId = 'testmodeler.mock.send-example-id',
  verifyResponse = 'testmodeler.mock.verify-response',
  skipLint = 'gotest.skip-lint',
  factoryGatherAllParams = 'factory-gather-all-params',
}

export const configDefaults = {
  ...baseConfigDefaults,
  [Config.exportCodemodel]: false,
  [Config.generateMockTest]: false,
  [Config.generateSdkExample]: false,
  [Config.generateScenarioTest]: false,
  [Config.generateSdkSample]: false,
  [Config.generateFakeTest]: false,
  [Config.filePrefix]: '',
  [Config.exampleFilePrefix]: '',
  [Config.testFilePrefix]: '',
  [Config.skipLint]: false,
  [Config.factoryGatherAllParams]: false,
};
