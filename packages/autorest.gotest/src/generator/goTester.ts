/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as _ from 'lodash';
import { AutorestExtensionHost } from '@autorest/extension-base';
import { Config, configDefaults } from '../common/constant';
import { ExampleCodeGenerator, ExampleDataRender } from './exampleGenerator';
import { GenerateContext } from './generateContext';
import { Helper } from '@autorest/testmodeler/dist/src/util/helper';
import { MockTestCodeGenerator, MockTestDataRender } from './mockTestGenerator';
import { SampleCodeGenerator, SampleDataRender } from './sampleGenerator';
import { ScenarioTestCodeGenerator, ScenarioTestDataRender } from './scenarioTestGenerator';
import { TestCodeModeler } from '@autorest/testmodeler/dist/src/core/model';
import { TestConfig } from '@autorest/testmodeler/dist/src/common/testConfig';

export async function processRequest(host: AutorestExtensionHost): Promise<void> {
  const session = await TestCodeModeler.getSessionFromHost(host);

  const config = new TestConfig(await session.getValue(''), configDefaults);
  if (config.getValue(Config.exportCodemodel)) {
    Helper.addCodeModelDump(session, 'go-tester-pre.yaml', false);
  }

  // try to get commit/tree name from require config
  const rpRegex = /Azure\/azure-rest-api-specs\/(blob\/|tree\/|)(?<swaggerCommit>[^/]+)\//;

  let swaggerCommit = 'main';
  if (session.configuration?.require) {
    for (const config of session.configuration.require) {
      const execResult = rpRegex.exec(config);
      if (execResult?.groups['swaggerCommit']) {
        swaggerCommit = execResult?.groups['swaggerCommit'];
        break;
      }
    }
  }

  const context = new GenerateContext(host, session.model, config, swaggerCommit);
  const mockTestDataRender = new MockTestDataRender(context);
  mockTestDataRender.renderData();

  const extraParam = {
    copyright: await Helper.getCopyright(session),
    sendExampleId: config.getValue(Config.sendExampleId),
    verifyResponse: config.getValue(Config.verifyResponse),
  };
  if (config.getValue(Config.generateMockTest)) {
    const mockTestCodeGenerator = new MockTestCodeGenerator(context);
    mockTestCodeGenerator.generateCode(extraParam);
  }
  if (config.getValue(Config.generateSdkExample)) {
    const exampleDataRender = new ExampleDataRender(context);
    exampleDataRender.renderData();
    const exampleCodeGenerator = new ExampleCodeGenerator(context);
    exampleCodeGenerator.generateCode(extraParam);
  }
  if (config.getValue(Config.generateScenarioTest)) {
    const scenarioTestDataRender = new ScenarioTestDataRender(context);
    scenarioTestDataRender.renderData();
    const scenarioTestCodeGenerator = new ScenarioTestCodeGenerator(context);
    scenarioTestCodeGenerator.generateCode(extraParam);
  }
  if (config.getValue(Config.generateSdkSample)) {
    const sampleDataRender = new SampleDataRender(context);
    sampleDataRender.renderData();
    const sampleCodeGenerator = new SampleCodeGenerator(context);
    sampleCodeGenerator.generateCode(extraParam);
  }
  await Helper.outputToModelerfour(host, session, false);
  if (config.getValue(Config.exportCodemodel)) {
    Helper.addCodeModelDump(session, 'go-tester.yaml', false);
  }
  Helper.dump(host);
}
