/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Host, startSession } from '@azure-tools/autorest-extension-base';
import { codeModelSchema, CodeModel } from '@azure-tools/codemodel';
import { generateModels } from './models'

// The generator emits Go source code files to disk.
export async function generator(host: Host) {
  const debug = await host.GetValue('debug') || false;

  try {

    // get the code model from the core
    const session = await startSession<CodeModel>(host, codeModelSchema);
    const namespace = await session.getValue('namespace');

    const models = await generateModels(session);

    host.WriteFile(`internal/${namespace}/models.go`, models, undefined, 'source-file-go');

  } catch (E) {
    if (debug) {
      console.error(`${__filename} - FAILURE  ${JSON.stringify(E)} ${E.stack}`);
    }
    throw E;
  }
}
