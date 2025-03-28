/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../../codemodel.go/src/index.js';
import { getServerName } from './servers.js';
import { contentPreamble } from '../helpers.js';
import { ImportManager } from '../imports.js';

export function generateServerFactory(codeModel: go.CodeModel): string {
  // generate server factory only for ARM
  if (codeModel.type !== 'azure-arm' || !codeModel.clients) {
    return '';
  }

  const imports = new ImportManager();
  imports.add('errors');
  imports.add('fmt');
  imports.add('net/http');
  imports.add('strings');
  imports.add('sync');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');

  let text = contentPreamble(codeModel, 'fake');
  text += imports.text();

  text += `// ServerFactory is a fake server for instances of the ${codeModel.packageName}.ClientFactory type.\n`;
  text += 'type ServerFactory struct {\n';
  for (const client of codeModel.clients) {
    const serverName = getServerName(client);
    text += `\t// ${serverName} contains the fakes for client ${client.name}\n`;
    text += `\t${serverName} ${serverName}\n\n`;
  }
  text += '}\n\n';

  text += '// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.\n';
  text += `// The returned ServerFactoryTransport instance is connected to an instance of ${codeModel.packageName}.ClientFactory via the\n`;
  text += '// azcore.ClientOptions.Transporter field in the client\'s constructor parameters.\n';
  text += 'func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {\n';
  text += '\treturn &ServerFactoryTransport{\n\t\tsrv: srv,\n\t}\n}\n\n';

  text += `// ServerFactoryTransport connects instances of ${codeModel.packageName}.ClientFactory to instances of ServerFactory.\n`;
  text += '// Don\'t use this type directly, use NewServerFactoryTransport instead.\n';
  text += 'type ServerFactoryTransport struct {\n';
  text += '\tsrv *ServerFactory\n';
  text += '\ttrMu sync.Mutex\n';
  for (const client of codeModel.clients) {
    const serverName = getServerName(client);
    text += `\ttr${serverName} *${serverName}Transport\n`;
  }
  text += '}\n\n';

  text += '// Do implements the policy.Transporter interface for ServerFactoryTransport.\n';
  text += 'func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {\n';
  text += '\trawMethod := req.Context().Value(runtime.CtxAPINameKey{})\n';
  text += '\tmethod, ok := rawMethod.(string)\n';
  text += '\tif !ok {\n\t\treturn nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}\n\t}\n\n';
  text += '\tclient := method[:strings.Index(method, ".")]\n';
  text += '\tvar resp *http.Response\n\tvar err error\n\n';
  text += '\tswitch client {\n';
  for (const client of codeModel.clients) {
    text += `\tcase "${client.name}":\n`;
    const serverName = getServerName(client);
    text += `\t\tinitServer(s, &s.tr${serverName}, func() *${serverName}Transport { return New${serverName}Transport(&s.srv.${serverName}) })\n`;
    text += `\t\tresp, err = s.tr${serverName}.Do(req)\n`;
  }
  text += '\tdefault:\n\t\terr = fmt.Errorf("unhandled client %s", client)\n';
  text += '\t}\n\n';
  text += '\tif err != nil {\n\t\treturn nil, err\n\t}\n\n';
  text += '\treturn resp, nil\n}\n\n';

  text += 'func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {\n';
  text += '\ts.trMu.Lock()\n';
  text += '\tif *dst == nil {\n\t\t*dst = src()\n\t}\n';
  text += '\ts.trMu.Unlock()\n}\n';
  return text;
}
