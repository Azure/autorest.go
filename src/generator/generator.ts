/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { serialize } from '@azure-tools/codegen';
import { Host, startSession } from '@azure-tools/autorest-extension-base';
import { codeModelSchema, CodeModel } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { generateOperations } from './operations';
import { generateModels } from './models';
import { generateEnums } from './enums';
import { generateClient } from './client';
import { generateTimeHelpers } from './time';
import { generatePagers } from './pagers';
import { generatePollers, generatePollersHelper } from './data_plane_pollers';
import { generatePolymorphicHelpers } from './polymorphics';
import { generateGoModFile } from './gomod';
import { generateXMLAdditionalPropsHelpers } from './xmlAdditionalProps';

// The generator emits Go source code files to disk.
export async function protocolGen(host: Host) {
  const debug = await host.GetValue('debug') || false;

  try {
    // get the code model from the core
    const session = await startSession<CodeModel>(host, codeModelSchema);
    const operations = await generateOperations(session);
    let filePrefix = await session.getValue('file-prefix', '');
    // if a file prefix was specified, ensure it's properly snaked
    if (filePrefix.length > 0 && filePrefix[filePrefix.length - 1] !== '_') {
      filePrefix += '_';
    }

    // output the model to the pipeline.  this must happen after all model
    // updates are complete and before any source files are written.
    host.WriteFile('code-model-v4.yaml', serialize(session.model), undefined, 'code-model-v4');

    for (const op of values(operations)) {
      host.WriteFile(`${filePrefix}${op.name.toLowerCase()}.go`, op.content, undefined, 'source-file-go');
    }

    const enums = await generateEnums(session);
    if (enums.length > 0) {
      host.WriteFile(`${filePrefix}enums.go`, enums, undefined, 'source-file-go');
    }

    const models = await generateModels(session);
    host.WriteFile(`${filePrefix}models.go`, models, undefined, 'source-file-go');

    const client = await generateClient(session);
    host.WriteFile(`${filePrefix}client.go`, client, undefined, 'source-file-go');

    const timeHelpers = await generateTimeHelpers(session);
    for (const helper of values(timeHelpers)) {
      host.WriteFile(`${filePrefix}${helper.name.toLowerCase()}.go`, helper.content, undefined, 'source-file-go');
    }

    const pagers = await generatePagers(session);
    if (pagers.length > 0) {
      host.WriteFile(`${filePrefix}pagers.go`, pagers, undefined, 'source-file-go');
    }
    const pollers = await generatePollers(session);
    if (pollers.length > 0) {
      const pollingHelper = await generatePollersHelper(session);
      host.WriteFile(`${filePrefix}pollers_helper.go`, pollingHelper, undefined, 'source-file-go');
      host.WriteFile(`${filePrefix}pollers.go`, pollers, undefined, 'source-file-go');
    }
    const polymorphics = await generatePolymorphicHelpers(session);
    if (polymorphics.length > 0) {
      host.WriteFile(`${filePrefix}polymorphic_helpers.go`, polymorphics, undefined, 'source-file-go');
    }
    const gomod = await generateGoModFile(session);
    if (gomod.length > 0) {
      host.WriteFile('go.mod', gomod, undefined, 'source-file-go');
    }
    const xmlAddlProps = await generateXMLAdditionalPropsHelpers(session);
    if (xmlAddlProps.length > 0) {
      host.WriteFile(`${filePrefix}xml_helper.go`, xmlAddlProps, undefined, 'source-file-go');
    }
  } catch (E) {
    if (debug) {
      console.error(`${__filename} - FAILURE  ${JSON.stringify(E)} ${E.stack}`);
    }
    throw E;
  }
}
