/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../codemodel.go/src/index.js';
import { lt, toSemver } from '@azure-tools/codegen';
import { CodegenError } from './errors.js';

// Creates the content in go.mod if the --module switch was specified.
// if there's a preexisting go.mod file, update its specified version of azcore as needed.
export async function generateGoModFile(codeModel: go.CodeModel, existingGoMod?: string): Promise<string> {
  if (!codeModel.options.module) {
    // must be containing-module, so we're emitting a package into a module that already has a go.mod
    return '';
  }

  const modName = codeModel.options.module;

  // here we specify the minimum version of azcore as required by the code generator.
  // the version can be overwritten by passing the --azcore-version switch during generation.
  let version = '1.17.1';
  if (codeModel.options.azcoreVersion) {
    // when matching versions, we need to handle beta, non-beta, and pseudo versions
    // 1.2.3-beta.1, 1.2.3, 0.22.1-0.20220315231014-ed309e73db6b
    if (!codeModel.options.azcoreVersion.match(/^\d+\.\d+\.\d+(?:-[a-zA-Z0-9_.-]+)?$/)) {
      throw new CodegenError('InvalidArgument', `azcore version ${version} must be in the format major.minor.patch[-beta.N]`);
    }
    version = codeModel.options.azcoreVersion;
  }

  const azcore = 'github.com/Azure/azure-sdk-for-go/sdk/azcore v' + version;
  if (!existingGoMod) {
    // no preexisting go.mod file, generate the default one
    let text = `module ${modName}\n\n`;
    text += 'go 1.23.0\n\n';
    text += `require ${azcore}\n`;
    return text;
  }

  // check if the module identity needs to be replaced due to a major version change
  if (!existingGoMod.match(`module ${modName}$`)) {
    existingGoMod = existingGoMod.replace(/module \S+/, `module ${modName}`);
  }

  // check if the existing version of azcore is greater than or equal to the specified version.
  // note that some modules (e.g. models-only) might not have a dependency on azcore.
  const match = existingGoMod.match(/github\.com\/Azure\/azure-sdk-for-go\/sdk\/azcore\s+v(\d+\.\d+\.\d+(?:-[a-zA-Z0-9_.-]+)?)/);
  if (match) {
    if (match.length < 2) {
      throw new CodegenError('InternalError', 'returned matches were less than expected');
    }
    const existingVer = toSemver(match[1]);
    const specifiedVer = toSemver(version);
    if (lt(existingVer, specifiedVer)) {
      // the existing version of azcore is less than the specified version so update it
      existingGoMod = existingGoMod.replace(/github\.com\/Azure\/azure-sdk-for-go\/sdk\/azcore\s+v\d+\.\d+\.\d+(?:-[a-zA-Z0-9_.-]+)?/, azcore);
    }
  }
  return existingGoMod;
}
