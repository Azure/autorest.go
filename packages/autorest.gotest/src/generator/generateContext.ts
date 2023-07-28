/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { AutorestExtensionHost } from '@autorest/extension-base';
import { Config } from '../common/constant';
import { ImportManager } from '@autorest/go/dist/src/generator/imports';
import { TestCodeModel } from '@autorest/testmodeler/dist/src/core/model';
import { TestConfig } from '@autorest/testmodeler/dist/src/common/testConfig';
export class GenerateContext {
  public packageName: string;
  public importManager: ImportManager;

  public constructor(public host: AutorestExtensionHost, public codeModel: TestCodeModel, public testConfig: TestConfig, public swaggerCommit = 'main') {
    this.packageName = this.codeModel?.language?.go?.packageName;
    this.importManager = new ImportManager();
    if (this.packageName) {
      const modName = this.codeModel.language.go!.module
      if (modName !== 'none' && modName) {
        this.importManager.add(modName);
      }
    }
  }
}
