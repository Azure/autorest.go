/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel, OperationGroup, Parameter} from '@autorest/codemodel';
import { length, values } from '@azure-tools/linq';
import { contentPreamble, formatCommentAsBulletItem, formatParameterTypeName, sortAscending, sortParametersByRequired } from './helpers';
import { ImportManager } from './imports';


// Creates the content for all <operation>.go files
export async function generateClientFactory(session: Session<CodeModel>): Promise<string> {
  const azureARM = <boolean>session.model.language.go!.azureARM;
  let result = '';
  // generate client factory only for ARM
  if (azureARM && length(session.model.operationGroups) > 0) {
    session.model.operationGroups.sort((a: OperationGroup, b: OperationGroup) => { return sortAscending(a.language.go!.clientName, b.language.go!.clientName); });

    // the list of packages to import
    const imports = new ImportManager();
    
    // there should be at most one client level param: subscriptionID for ARM, any exception is always a wrong swagger definition that we should fix
    const allClientParams = new Array<Parameter>();
    for (const group of values(session.model.operationGroups)) {
      if (group.language.go!.clientParams) {
        const clientParams = <Array<Parameter>>group.language.go!.clientParams;
        for (const clientParam of values(clientParams)) {
          if (values(allClientParams).where(cp => cp.language.go!.name === clientParam.language.go!.name).any()) {
            continue;
          }
          allClientParams.push(clientParam);
        }
      }
    }
    allClientParams.sort(sortParametersByRequired);

    // add factory type
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
    result += '// ClientFactory is a client factory used to create any client in this module.\n';
    result += '// Don\'t use this type directly, use NewClientFactory instead.\n';
    result += 'type ClientFactory struct {\n';
    for (const clientParam of values(allClientParams)) {
      result += `\t${clientParam.language.go!.name} ${formatParameterTypeName(clientParam)}\n`;
    }
    result += '\tcredential azcore.TokenCredential\n';
    result += '\toptions *arm.ClientOptions\n';
    result += '}\n\n';

    // add factory CTOR
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/arm');
    result += '// NewClientFactory creates a new instance of ClientFactory with the specified values.\n';
    result += '// The parameter values will be propagated to any client created from this factory.\n';
    for (const clientParam of values(allClientParams)) {
      result += `${formatCommentAsBulletItem(`${clientParam.language.go!.name} - ${clientParam.language.go!.description}`)}\n`;
    }
    result += `${formatCommentAsBulletItem('credential - used to authorize requests. Usually a credential from azidentity.')}\n`;
    result += `${formatCommentAsBulletItem('options - pass nil to accept the default values.')}\n`;

    result += `func NewClientFactory(${allClientParams.map(p => { return `${p.language.go!.name} ${formatParameterTypeName(p)}`; }).join(', ')}${allClientParams.length>0 ? ',' : ''} credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {\n`;
    result += '\t_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)\n';
    result += '\tif err != nil {\n';
    result += '\t\treturn nil, err\n';
    result += '\t}\n';
    result += '\treturn &ClientFactory{\n';
    for (const clientParam of values(allClientParams)) {
      result += `\t\t${clientParam.language.go!.name}: \t${clientParam.language.go!.name},`;
    }
    result += '\t\tcredential: credential,\n';
    result += '\t\toptions: options.Clone(),\n';
    result += '\t}, nil\n';
    result += '}\n\n';

    // add new sub client method for all operation groups
    for (const group of values(session.model.operationGroups)) {
      result += `func (c *ClientFactory) ${group.language.go!.clientCtorName}() *${group.language.go!.clientName} {\n`;
      const clientParams = <Array<Parameter>>group.language.go!.clientParams;
      if (clientParams) {
        clientParams.sort(sortParametersByRequired);
        result += `\tsubClient, _ := ${group.language.go!.clientCtorName}(${clientParams.map(p => { return `c.${p.language.go!.name}`; }).join(', ')}, c.credential, c.options)\n`;
      } else {
        result += `\tsubClient, _ := ${group.language.go!.clientCtorName}(c.credential, c.options)\n`;
      }
      
      result += '\treturn subClient\n';
      result += '}\n\n';
    }

    result = await contentPreamble(session) + imports.text() + result;
  }
  return result;
}
