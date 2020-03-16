/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { AutoRestExtension, } from '@azure-tools/autorest-extension-base';
import { namer } from './namer/namer';
import { transform } from './transform/transform';
import { protocolGen } from './generator/protocol/generator';

require('source-map-support').install();

export async function main() {
  const pluginHost = new AutoRestExtension();
  pluginHost.Add('go-namer', namer);
  pluginHost.Add('go-transform', transform);
  pluginHost.Add('go-protocol', protocolGen);
  await pluginHost.Run();
}

main();
