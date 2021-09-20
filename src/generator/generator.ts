/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { serialize } from '@azure-tools/codegen';
import { Host, startSession, Session } from '@autorest/extension-base';
import { codeModelSchema, CodeModel } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { generateOperations } from './operations';
import { generateModels } from './models';
import { generateResponses } from './responses';
import { generateConstants } from './constants';
import { generateConnection } from './connection';
import { generateTimeHelpers } from './time';
import { generatePagers } from './pagers';
import { generatePollers } from './pollers';
import { generatePolymorphicHelpers } from './polymorphics';
import { generateGoModFile } from './gomod';
import { generateXMLAdditionalPropsHelpers } from './xmlAdditionalProps';

async function getModuleVersion(session: Session<CodeModel>): Promise<string> {
  const version = await session.getValue('module-version', '');
  if (version === '') {
    throw new Error('--module-version is a required parameter');
  } else if (!version.match(/^\d+\.\d+\.\d+$/) && !version.match(/^\d+\.\d+\.\d+-beta\.\d+$/)) {
    throw new Error(`module version ${version} must in the format major.minor.patch[-beta.N]`);
  }
  return version;
}

// The generator emits Go source code files to disk.
export async function protocolGen(host: Host) {
  const debug = await host.GetValue('debug') || false;

  try {
    // get the code model from the core
    const session = await startSession<CodeModel>(host, codeModelSchema);
    const version = await getModuleVersion(session);

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
      host.WriteFile(`${filePrefix}${op.name.toLowerCase()}_client.go`, op.content, undefined, 'source-file-go');
    }

    const constants = await generateConstants(session, version);
    host.WriteFile(`${filePrefix}constants.go`, constants, undefined, 'source-file-go');

    const models = await generateModels(session);
    host.WriteFile(`${filePrefix}models.go`, models, undefined, 'source-file-go');

    const responses = await generateResponses(session);
    if (responses.length > 0) {
      host.WriteFile(`${filePrefix}response_types.go`, responses, undefined, 'source-file-go');
    }

    const connection = await generateConnection(session);
    if (connection.length > 0) {
      host.WriteFile(`${filePrefix}connection.go`, connection, undefined, 'source-file-go');
    }

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
      console.error(`${__filename} - FAILURE  ${JSON.stringify(E)} ${(<Error>E).stack}`);
    }
    throw E;
  }
}
