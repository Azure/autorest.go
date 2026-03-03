/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as helpers from './helpers.js';
import * as go from '../../../codemodel.go/src/index.js';

/**
 * Creates the content in version.go.
 * if version.go doesn't need to be created, the empty string is returned.
 *
 * @param module the module for which to generate version.go
 * @returns the contents of version.go or the empty string
 */
export function generateVersionInfo(module: go.Module): string {
  let text = helpers.contentPreamble(module, false);

  text += 'const (\n';
  // strip off any major version suffix. this is for telemetry
  // purposes, so all major versions coalesce into the same bucket
  text += `\tmoduleName = "${module.identity.replace(/\/v\d+$/, '')}"\n`;

  // for new modules, we seed the moduleVersion with a prerelease version
  text += '\tmoduleVersion = "v0.1.0"\n';
  text += ')\n\n';

  return text;
}
