/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { AutorestExtensionHost } from '@autorest/extension-base';
import { getModuleNameWithMajorVersion } from '../common/helpers';
import { ImportManager } from './imports';
import { TestCodeModel } from '@autorest/testmodeler';
import { TestConfig } from '@autorest/testmodeler';
export class GenerateContext {
  public packageName: string;
  public importManager: ImportManager;

  public constructor(public host: AutorestExtensionHost, public codeModel: TestCodeModel, public testConfig: TestConfig, public swaggerCommit = 'main') {
    this.packageName = this.codeModel?.language?.go?.packageName;
    this.importManager = new ImportManager();
    if (this.packageName) {
      const modName = getModuleNameWithMajorVersion(this.codeModel);
      if (modName) {
        this.importManager.add(modName);
      }
    }
  }
}
