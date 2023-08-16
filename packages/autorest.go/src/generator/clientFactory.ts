/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { GoCodeModel, Parameter } from '../gocodemodel/gocodemodel'; 
import { values } from '@azure-tools/linq';
import { contentPreamble, formatCommentAsBulletItem, formatParameterTypeName, sortParametersByRequired } from './helpers';
import { ImportManager } from './imports';


// Creates the content for all <operation>.go files
export async function generateClientFactory(codeModel: GoCodeModel): Promise<string> {
  let result = '';
  // generate client factory only for ARM
  if (codeModel.type === 'azure-arm' && codeModel.clients) {
    // the list of packages to import
    const imports = new ImportManager();
    
    // there should be at most one client level param: subscriptionID for ARM, any exception is always a wrong swagger definition that we should fix
    const allClientParams = new Array<Parameter>();
    for (const clients of codeModel.clients) {
      for (const clientParam of values(clients.parameters)) {
        if (values(allClientParams).where(each => each.paramName === clientParam.paramName).any()) {
          continue;
        }
        allClientParams.push(clientParam);
      }
    }
    allClientParams.sort(sortParametersByRequired);

    // add factory type
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    result += '// ClientFactory is a client factory used to create any client in this module.\n';
    result += '// Don\'t use this type directly, use NewClientFactory instead.\n';
    result += 'type ClientFactory struct {\n';
    for (const clientParam of values(allClientParams)) {
      result += `\t${clientParam.paramName} ${formatParameterTypeName(clientParam)}\n`;
    }
    result += '\tcredential azcore.TokenCredential\n';
    result += '\toptions *arm.ClientOptions\n';
    result += '}\n\n';

    // add factory CTOR
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/arm');
    result += '// NewClientFactory creates a new instance of ClientFactory with the specified values.\n';
    result += '// The parameter values will be propagated to any client created from this factory.\n';
    for (const clientParam of values(allClientParams)) {
      result += `${formatCommentAsBulletItem(`${clientParam.paramName} - ${clientParam.description}`)}\n`;
    }
    result += `${formatCommentAsBulletItem('credential - used to authorize requests. Usually a credential from azidentity.')}\n`;
    result += `${formatCommentAsBulletItem('options - pass nil to accept the default values.')}\n`;

    result += `func NewClientFactory(${allClientParams.map(each => { return `${each.paramName} ${formatParameterTypeName(each)}`; }).join(', ')}${allClientParams.length>0 ? ',' : ''} credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {\n`;
    result += '\t_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)\n';
    result += '\tif err != nil {\n';
    result += '\t\treturn nil, err\n';
    result += '\t}\n';
    result += '\treturn &ClientFactory{\n';
    for (const clientParam of values(allClientParams)) {
      result += `\t\t${clientParam.paramName}: \t${clientParam.paramName},`;
    }
    result += '\t\tcredential: credential,\n';
    result += '\t\toptions: options.Clone(),\n';
    result += '\t}, nil\n';
    result += '}\n\n';

    // add new sub client method for all operation groups
    for (const client of codeModel.clients) {
      result += `func (c *ClientFactory) ${client.ctorName}() *${client.clientName} {\n`;
      if (client.parameters) {
        result += `\tsubClient, _ := ${client.ctorName}(${client.parameters.map(each => { return `c.${each.paramName}`; }).join(', ')}, c.credential, c.options)\n`;
      } else {
        result += `\tsubClient, _ := ${client.ctorName}(c.credential, c.options)\n`;
      }
      
      result += '\treturn subClient\n';
      result += '}\n\n';
    }

    result = contentPreamble(codeModel) + imports.text() + result;
  }
  return result;
}
