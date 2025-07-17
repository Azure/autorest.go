/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as helpers from './helpers.js';
import * as go from '../../codemodel.go/src/index.js';
import { CodegenError } from './errors.js';

// Creates the content in version.go
export async function generateVersionInfo(codeModel: go.CodeModel): Promise<string> {
  if (codeModel.options.containingModule) {
    // code is being emitted into an existing module
    return '';
  } else if (!codeModel.options.module) {
    throw new CodegenError('InvalidArgument', 'missing --module or --containing-module argument');
  }

  let text = helpers.contentPreamble(codeModel, false);

  text += 'const (\n';
  // strip off any major version suffix. this is for telemetry
  // purposes, so all major versions coalesce into the same bucket
  text += `\tmoduleName = "${codeModel.options.module.name.replace(/\/v\d+$/, '')}"\n`;
  text += `\tmoduleVersion = "v${codeModel.options.module.version}"\n`;
  text += ')\n\n';

  return text;
}
