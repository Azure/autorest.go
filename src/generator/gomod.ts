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
  text += 'go 1.16\n\n';
  // here we specify the minimum version of azcore as required by the code generator
  // TODO: come up with a way to get the latest minor/patch versions.
  const version = 'v0.21.1';
  const azcore = 'github.com/Azure/azure-sdk-for-go/sdk/azcore ' + version;
  text += `require ${azcore}\n`;
  return text;
}
