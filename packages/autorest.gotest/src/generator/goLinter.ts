/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { AutorestExtensionHost } from '@autorest/extension-base';
import { TestConfig } from '@autorest/testmodeler/dist/src/common/testConfig';
import { Helper } from '@autorest/testmodeler/dist/src/util/helper';
import * as path from 'path';
import { Config, configDefaults } from '../common/constant';

export async function processRequest(host: AutorestExtensionHost): Promise<void> {
  const testConfig = new TestConfig(await host.GetValue(''), configDefaults);
  if (!testConfig.getValue(Config.skipLint)) {
    const files = await host.listInputs();
    Helper.execSync('go install golang.org/x/tools/cmd/goimports@latest');
    for (const outputFile of files) {
      if (outputFile.endsWith('.go')) {
        const pathName = path.join(testConfig.getValue(Config.outputFolder), outputFile);
        Helper.execSync(`goimports -w ${pathName}`);
      }
    }
  }
}
