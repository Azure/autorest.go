/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { serialize } from '@azure-tools/codegen';
import { Host, startSession } from '@azure-tools/autorest-extension-base';
import { codeModelSchema, CodeModel } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { generateOperations } from './operations'
import { generateModels } from './models'
import { generateEnums } from './enums'
import { generateTimeHelpers } from './time'

// The generator emits Go source code files to disk.
export async function protocolGen(host: Host) {
  const debug = await host.GetValue('debug') || false;

  try {

    // get the code model from the core
    const session = await startSession<CodeModel>(host, codeModelSchema);
    const namespace = await session.getValue('namespace');

    const operations = await generateOperations(session);

    // output the model to the pipeline.  this must happen after all model
    // updates are complete and before any source files are written.
    host.WriteFile('code-model-v4.yaml', serialize(session.model), undefined, 'code-model-v4');

    for (const op of values(operations)) {
      host.WriteFile(`internal/${namespace}/${op.name.toLowerCase()}.go`, op.content, undefined, 'source-file-go');
    }

    const enums = await generateEnums(session);
    if (enums.length > 0) {
      host.WriteFile(`internal/${namespace}/enums.go`, enums, undefined, 'source-file-go');
    }

    const models = await generateModels(session);
    host.WriteFile(`internal/${namespace}/models.go`, models, undefined, 'source-file-go');

    const timeHelpers = await generateTimeHelpers(session);
    for (const helper of values(timeHelpers)) {
      host.WriteFile(`internal/${namespace}/${helper.name.toLowerCase()}.go`, helper.content, undefined, 'source-file-go');
    }
  } catch (E) {
    if (debug) {
      console.error(`${__filename} - FAILURE  ${JSON.stringify(E)} ${E.stack}`);
    }
    throw E;
  }
}
