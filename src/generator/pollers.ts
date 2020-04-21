/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { pascalCase } from '@azure-tools/codegen';
import { CodeModel } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { PollerInfo } from '../common/helpers';
import { contentPreamble, sortAscending } from './helpers';
import { ImportManager } from './imports';

// Creates the content in pagers.go
export async function generatePollers(session: Session<CodeModel>): Promise<string> {
  if (session.model.language.go!.pollerTypes === undefined) {
    return '';
  }
  let text = await contentPreamble(session);

  // add standard imports
  const imports = new ImportManager();
  imports.add('context');
  // imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  imports.add('net/http');
  imports.add('time');
  text += imports.text();

  const pollers = <Array<PollerInfo>>session.model.language.go!.pollerTypes;
  pollers.sort((a: PollerInfo, b: PollerInfo) => { return sortAscending(a.name, b.name) });
  for (const poller of values(pollers)) {
    const pollerInterface = pascalCase(poller.name);
    let responseType = '';
    if (poller.schema === undefined) {
      responseType = 'http.Response';
    } else {
      responseType = poller.schema.language.go!.responseType.name;
    }
    text += `// ${pollerInterface} provides polling facilities until the operation completes
type ${pollerInterface} interface {
	Done() bool
	ID() string
	Poll(context.Context) (*${responseType}, error)
	Wait(ctx context.Context, pollingInterval time.Duration) (*${responseType}, error)
}

type ${poller.name} struct {
	// the client for making the request
	client *${poller.client}
	// polling tracker
	// pt *pollingTracker${pascalCase(poller.pollingMethod)}
}

func (p *${poller.name}) Done() bool {
	return false
}

func (p *${poller.name}) ID() string {
	return "NYI"
}

func (p *${poller.name}) Poll(ctx context.Context) (*${responseType}, error) {
	return nil, nil
}

func (p *${poller.name}) Wait(ctx context.Context, pollingInterval time.Duration) (*${responseType}, error) {
	return nil, nil
}
`;
  }
  return text;
}
