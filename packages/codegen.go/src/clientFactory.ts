/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { values } from '@azure-tools/linq';
import * as go from '../../codemodel.go/src/index.js';
import { contentPreamble, formatCommentAsBulletItem, formatParameterTypeName, getAllClientParameters } from './helpers.js';
import { ImportManager } from './imports.js';


// Creates the content for client_factory.go (ARM only)
export async function generateClientFactory(codeModel: go.CodeModel): Promise<string> {
  // generate client factory only for ARM
  if (codeModel.type !== 'azure-arm' || codeModel.clients.length === 0) {
    return '';
  }

  let result = '';
  // the list of packages to import
  const imports = new ImportManager();

  const allClientParams = getAllClientParameters(codeModel);

  // add factory type
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  result += '// ClientFactory is a client factory used to create any client in this module.\n';
  result += '// Don\'t use this type directly, use NewClientFactory instead.\n';
  result += 'type ClientFactory struct {\n';
  for (const clientParam of values(allClientParams)) {
    result += `\t${clientParam.name} ${formatParameterTypeName(clientParam)}\n`;
  }
  result += '\tinternal *arm.Client\n';
  result += '}\n\n';

  // add factory CTOR
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/arm');
  result += '// NewClientFactory creates a new instance of ClientFactory with the specified values.\n';
  result += '// The parameter values will be propagated to any client created from this factory.\n';
  for (const clientParam of values(allClientParams)) {
    result += `${formatCommentAsBulletItem(`${clientParam.name} - ${clientParam.description}`)}\n`;
  }
  result += `${formatCommentAsBulletItem('credential - used to authorize requests. Usually a credential from azidentity.')}\n`;
  result += `${formatCommentAsBulletItem('options - pass nil to accept the default values.')}\n`;

  result += `func NewClientFactory(${allClientParams.map(param => { return `${param.name} ${formatParameterTypeName(param)}`; }).join(', ')}${allClientParams.length > 0 ? ',' : ''} credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {\n`;
  result += '\tinternal, err := arm.NewClient(moduleName, moduleVersion, credential, options)\n';
  result += '\tif err != nil {\n';
  result += '\t\treturn nil, err\n';
  result += '\t}\n';
  result += '\treturn &ClientFactory{\n';
  for (const clientParam of values(allClientParams)) {
    result += `\t\t${clientParam.name}: ${clientParam.name},\n`;
  }
  result += '\t\tinternal: internal,\n';
  result += '\t}, nil\n';
  result += '}\n\n';

  // add new sub client method for all operation groups
  for (const client of codeModel.clients) {
    const ctorName = `New${client.name}`;
    result += `// ${ctorName} creates a new instance of ${client.name}.\n`;
    result += `func (c *ClientFactory) ${ctorName}() *${client.name} {\n`;
    result += `\treturn &${client.name}{\n`;

    // some clients (e.g. operations client) don't utilize the client params
    if (client.parameters.length > 0) {
      for (const clientParam of values(allClientParams)) {
        result += `\t\t${clientParam.name}: c.${clientParam.name},\n`;
      }
    }

    result += '\t\tinternal: c.internal,\n';
    result += '\t}\n';
    result += '}\n\n';
  }

  result = contentPreamble(codeModel) + imports.text() + result;
  return result;
}
