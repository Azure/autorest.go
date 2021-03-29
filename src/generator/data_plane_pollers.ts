/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { PollerInfo } from '../common/helpers';
import { contentPreamble, sortAscending } from './helpers';
import { ImportManager } from './imports';
import { generateARMPollers } from './arm_pollers';

// Creates the content in pollers.go
export async function generatePollers(session: Session<CodeModel>): Promise<string> {
  // get the openapi-type value specified. Default to ARM behavior, unless "data-plane" is specified
  const isARM = session.model.language.go!.openApiType === 'arm';
  if (isARM) {
    return generateARMPollers(session);
  }

  if (session.model.language.go!.pollerTypes === undefined) {
    return '';
  }
  let text = await contentPreamble(session);

  // add standard imports
  const imports = new ImportManager();
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  let bodyText = '';
  const pollers = <Array<PollerInfo>>session.model.language.go!.pollerTypes;
  pollers.sort((a: PollerInfo, b: PollerInfo) => { return sortAscending(a.name, b.name) });
  for (const poller of values(pollers)) {
    const pollerInterface = poller.name;
    bodyText += `// ${pollerInterface} provides polling facilities until the operation completes
      type ${pollerInterface} interface {
        azcore.Poller
      }
  `;
  }
  text += imports.text();
  text += bodyText;
  return text;
}
