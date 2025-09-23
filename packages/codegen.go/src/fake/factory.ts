/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { values } from '@azure-tools/linq';
import * as go from '../../../codemodel.go/src/index.js';
import { getServerName } from './servers.js';
import * as helpers from '../helpers.js';
import { ImportManager } from '../imports.js';

export function generateServerFactory(codeModel: go.CodeModel): string {
  // generate server factory only for ARM
  if (codeModel.type !== 'azure-arm' || !codeModel.clients) {
    return '';
  }

  const imports = new ImportManager();
  const indent = new helpers.indentation();
  imports.add('errors');
  imports.add('fmt');
  imports.add('net/http');
  imports.add('strings');
  imports.add('sync');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');

  let text = helpers.contentPreamble(codeModel, true, 'fake');
  text += imports.text();

  text += `// ServerFactory is a fake server for instances of the ${codeModel.packageName}.ClientFactory type.\n`;
  text += 'type ServerFactory struct {\n';

  // add server transports for client accessors
  // we might remove some clients from the list
  const finalSubClients = new Array<go.Client>();
  for (const client of codeModel.clients) {
    if (client.clientAccessors.length === 0 && values(client.methods).all(method => { return helpers.isMethodInternal(method) })) {
      // client has no client accessors and no exported methods, skip it
      continue;
    }
    const serverName = getServerName(client);
    text += `${indent.get()}// ${serverName} contains the fakes for client ${client.name}\n`;
    text += `${indent.get()}${serverName} ${serverName}\n\n`;
    finalSubClients.push(client);
  }
  text += '}\n\n';

  text += '// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.\n';
  text += `// The returned ServerFactoryTransport instance is connected to an instance of ${codeModel.packageName}.ClientFactory via the\n`;
  text += '// azcore.ClientOptions.Transporter field in the client\'s constructor parameters.\n';
  text += 'func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {\n';
  text += `${indent.get()}return &ServerFactoryTransport{\n${indent.push().get()}srv: srv,\n${indent.pop().get()}}\n}\n\n`;

  text += `// ServerFactoryTransport connects instances of ${codeModel.packageName}.ClientFactory to instances of ServerFactory.\n`;
  text += '// Don\'t use this type directly, use NewServerFactoryTransport instead.\n';
  text += 'type ServerFactoryTransport struct {\n';
  text += `${indent.get()}srv *ServerFactory\n`;
  text += `${indent.get()}trMu sync.Mutex\n`;
  for (const client of finalSubClients) {
    const serverName = getServerName(client);
    text += `${indent.get()}tr${serverName} *${serverName}Transport\n`;
  }
  text += '}\n\n';

  text += '// Do implements the policy.Transporter interface for ServerFactoryTransport.\n';
  text += 'func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {\n';
  text += `${indent.get()}rawMethod := req.Context().Value(runtime.CtxAPINameKey{})\n`;
  text += `${indent.get()}method, ok := rawMethod.(string)\n`;
  text += `${indent.get()}if !ok {\n${indent.push().get()}return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}\n${indent.pop().get()}}\n\n`;
  text += `${indent.get()}client := method[:strings.Index(method, ".")]\n`;
  text += `${indent.get()}var resp *http.Response\n${indent.get()}var err error\n\n`;
  text += `${indent.get()}switch client {\n`;
  for (const client of finalSubClients) {
    text += `${indent.get()}case "${client.name}":\n`;
    const serverName = getServerName(client);
    text += `${indent.push().get()}initServer(s, &s.tr${serverName}, func() *${serverName}Transport { return New${serverName}Transport(&s.srv.${serverName}) })\n`;
    text += `${indent.get()}resp, err = s.tr${serverName}.Do(req)\n`;
    indent.pop();
  }
  text += `${indent.get()}default:\n${indent.push().get()}err = fmt.Errorf("unhandled client %s", client)\n`;
  text += `${indent.pop().get()}}\n\n`;
  text += `${indent.get()}${helpers.buildErrCheck(indent, 'err', 'nil, err')}\n\n`;
  text += `${indent.get()}return resp, nil\n}\n\n`;

  text += 'func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {\n';
  text += `${indent.get()}s.trMu.Lock()\n`;
  text += `${indent.get()}${helpers.buildIfBlock(indent, {
    condition: '*dst == nil',
    body: (indent) => `${indent.get()}*dst = src()\n`,
  })}\n`;
  text += `${indent.get()}s.trMu.Unlock()\n}\n`;
  return text;
}
