/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel } from '@autorest/codemodel';
import { lt, toSemver } from '@azure-tools/codegen';

// Creates the content in go.mod if the --module switch was specified.
// if there's a preexisting go.mod file, update its specified version of azcore as needed.
export async function generateGoModFile(session: Session<CodeModel>, existingGoMod: string): Promise<string> {
  // here we specify the minimum version of azcore as required by the code generator.
  // the version can be overwritten by passing the --azcore-version switch during generation.
  const version = await session.getValue('azcore-version', '1.0.0');
  if (!version.match(/^\d+\.\d+\.\d+(?:-[a-zA-Z0-9_.-]+)?$/)) {
    throw new Error(`azcore version ${version} must in the format major.minor.patch[-beta.N]`);
  }

  const azcore = 'github.com/Azure/azure-sdk-for-go/sdk/azcore v' + version;
  if (existingGoMod === null) {
    // no preexisting go.mod file, generate the default one
    const modName = await session.getValue('module', 'none');
    if (modName === 'none') {
      return '';
    }
    let text = `module ${modName}\n\n`;
    text += 'go 1.18\n\n';
    text += `require ${azcore}\n`;
    return text;
  }

  // check if the existing version of azcore is greater than or equal to the specified version
  const match = existingGoMod.match(/github\.com\/Azure\/azure-sdk-for-go\/sdk\/azcore\s+v(\d+\.\d+\.\d+)/);
  if (!match) {
    throw new Error('preexisting go.mod is missing dependency on azcore');
  }
  if (match.length < 2) {
    throw new Error('returned matches were less than expected');
  }
  const existingVer = toSemver(match[1]);
  const specifiedVer = toSemver(version);
  if (lt(existingVer, specifiedVer)) {
    // the existing version of azcore is less than the specified version so update it
    existingGoMod = existingGoMod.replace(/github\.com\/Azure\/azure-sdk-for-go\/sdk\/azcore\s+v\d+\.\d+\.\d+/, azcore);
  }
  return existingGoMod;
}
