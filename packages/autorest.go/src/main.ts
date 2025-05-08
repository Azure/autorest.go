/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { AutoRestExtension } from '@autorest/extension-base';
import { transformM4 } from './transform/transform.js';
import { m4ToGoCodeModel } from './m4togocodemodel/adapter.js';
import { generateCode } from './generator/generator.js';
import 'source-map-support/register.js';

export async function main() {
  const pluginHost = new AutoRestExtension();
  pluginHost.add('go-transform-m4', transformM4);
  pluginHost.add('go-m4-to-gocodemodel', m4ToGoCodeModel);
  pluginHost.add('go-codegen', generateCode);
  await pluginHost.run();
}

await main();
