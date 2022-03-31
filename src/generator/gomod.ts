/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel } from '@autorest/codemodel';

// Creates the content in go.mod if the --module switch was specified
export async function generateGoModFile(session: Session<CodeModel>): Promise<string> {
  const modName = await session.getValue('module', 'none');
  if (modName === 'none') {
    return '';
  }
  let text = `module ${modName}\n\n`;
  text += 'go 1.18\n\n';
  // here we specify the minimum version of azcore as required by the code generator.
  // the version can be overwritten by passing the --azcore-version switch during generation.
  const version = await session.getValue('azcore-version', '0.22.1-0.20220330183203-27ad5f44019c');
  if (!version.match(/^\d+\.\d+\.\d+(?:-[a-zA-Z0-9_.-]+)?$/)) {
    throw new Error(`azcore version ${version} must in the format major.minor.patch[-beta.N]`);
  }
  const azcore = 'github.com/Azure/azure-sdk-for-go/sdk/azcore v' + version;
  text += `require ${azcore}\n`;
  return text;
}
