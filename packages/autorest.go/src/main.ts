/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { AutoRestExtension, } from '@autorest/extension-base';
import { transform } from './transform/transform';
import { protocolGen } from './generator/generator';

require('source-map-support').install();

export async function main() {
  const pluginHost = new AutoRestExtension();
  pluginHost.add('go-transform', transform);
  pluginHost.add('go-protocol', protocolGen);
  await pluginHost.run();
}

main();
