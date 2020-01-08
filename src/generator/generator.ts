/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Host, startSession, Session } from '@azure-tools/autorest-extension-base';
import { codeModelSchema, CodeModel, Language } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';


export async function generator(host: Host) {
  const debug = await host.GetValue('debug') || false;

  try {
    // get the code model from the core
    const session = await startSession<CodeModel>(host, codeModelSchema);

    const c = await session.getValue(<any>null, null);

    // example: do something here.
    let text = "A source file\n";

    const headerText = await session.getValue("header-text", "NO HEADER TEXT?");
    headerText = "foo";
    text = text + headerText;

    for (const each of values(session.model.schemas.objects)) {
      text = text + `schema: ${each.language.go?.name}\n`;
    }

    // example: output a generated text file
    host.WriteFile('go-sample.txt', text, undefined, 'source-file-go');

  } catch (E) {
    if (debug) {
      console.error(`${__filename} - FAILURE  ${JSON.stringify(E)} ${E.stack}`);
    }
    throw E;
  }
}
