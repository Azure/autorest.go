/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { CodeModel } from '@azure-tools/codemodel';

// Creates the content in go.mod if the --module switch was specified
export async function generateGoModFile(session: Session<CodeModel>): Promise<string> {
  const modName = await session.getValue('module', 'none');
  if (modName === 'none') {
    return '';
  }
  let text = `module ${modName}\n\n`;
  text += 'go 1.13\n\n';
  // here we specify the minimum version of armcore/azcore as required by the code generator
  // TODO: come up with a way to get the latest minor/patch versions.
  const azcore = 'github.com/Azure/azure-sdk-for-go/sdk/azcore v0.11.0';
  if (session.model.language.go!.openApiType === 'arm') {
    text += 'require (\n';
    text += '\tgithub.com/Azure/azure-sdk-for-go/sdk/armcore v0.3.2\n';
    text += `\t${azcore}\n`;
    text += ')\n'
  } else {
    text += `require ${azcore}\n`;
  }
  return text;
}
