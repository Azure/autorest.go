/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { serialize } from '@azure-tools/codegen';
import { AutorestExtensionHost, startSession } from '@autorest/extension-base';
import * as go from '../../../codemodel.go/src/index.js';
import * as codegen from '../../../codegen.go/src/index.js';
import { fileURLToPath } from 'url';

// The generator emits Go source code files to disk.
export async function generateCode(host: AutorestExtensionHost) {
  const debug = await host.getValue('debug') || false;

  try {
    // get the code model from the core
    const session = await startSession<go.CodeModel>(host);

    // output the model to the pipeline.  this must happen after all model
    // updates are complete and before any source files are written.
    host.writeFile({
      filename: 'code-model-v4.yaml',
      content: serialize(session.model),
      artifactType: 'code-model-v4'
    });

    let filePrefix = await session.getValue('file-prefix', '');
    // if a file prefix was specified, ensure it's properly snaked
    if (filePrefix.length > 0 && filePrefix[filePrefix.length - 1] !== '_') {
      filePrefix += '_';
    }

    const emitter = new codegen.Emitter(session.model, {
      exists: async (name: string) => {
        const content = await host.readFile(name);
        return Promise.resolve(content !== null);
      },
      read: (name: string) => {
        return host.readFile(name);
      },
      write: (name: string, content: string) => {
        return Promise.resolve(host.writeFile({
          filename: name,
          content: content,
          artifactType: 'source-file-go'
        }));
      }
    }, { filePrefix });
    await emitter.emit();
  } catch (E) {
    if (debug) {
      console.error(`${fileURLToPath(import.meta.url)} - FAILURE  ${JSON.stringify(E)} ${(<Error>E).stack}`);
    }
    throw E;
  }
}
