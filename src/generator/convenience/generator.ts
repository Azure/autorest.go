/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { serialize } from '@azure-tools/codegen';
import { Host, startSession } from '@azure-tools/autorest-extension-base';
import { codeModelSchema, CodeModel } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { generateClient } from './client';
import { generateModels } from './models';
import { generateOperations } from './operations';
import { generateEnums } from './enums';

// The generator emits Go source code files to disk.
export async function convenienceGen(host: Host) {
  const debug = await host.GetValue('debug') || false;

  try {

    // get the code model from the core
    const session = await startSession<CodeModel>(host, codeModelSchema);

    // output the model to the pipeline.  this must happen after all model
    // updates are complete and before any source files are written.
    host.WriteFile('code-model-v4.yaml', serialize(session.model), undefined, 'code-model-v4');

    const client = await generateClient(session);
    host.WriteFile('client.go', client, undefined, 'source-file-go');

    const operations = await generateOperations(session);
    for (const op of values(operations)) {
      host.WriteFile(`${op.name.toLowerCase()}.go`, op.content, undefined, 'source-file-go');
    }

    const enums = await generateEnums(session);
    if (enums.length > 0) {
      host.WriteFile('enums.go', enums, undefined, 'source-file-go');
    }

    const models = await generateModels(session);
    host.WriteFile('models.go', models, undefined, 'source-file-go');

  } catch (E) {
    if (debug) {
      console.error(`${__filename} - FAILURE  ${JSON.stringify(E)} ${E.stack}`);
    }
    throw E;
  }
}
