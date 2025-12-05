/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as helpers from './helpers.js';
import * as go from '../../../codemodel.go/src/index.js';

/**
 * generates the contents for the cloud_config.go file.
 * if cloud config info isn't required, the empty string is returned.
 * 
 * @param codeModel the code model for which to generate the file
 * @returns the text for the file or the empty string
 */
export function generateCloudConfig(codeModel: go.CodeModel): string {
  if (codeModel.type === 'azure-arm') {
    // this is handled in azcore
    return '';
  }

  // check if this SDK uses token auth
  let tokenCred: go.TokenCredential | undefined;
  for (const client of codeModel.clients) {
    if (client.instance?.kind !== 'constructable') {
      continue;
    }
    for (const constructor of client.instance.constructors) {
      for (const param of constructor.parameters) {
        if (param.kind === 'credentialParam' && param.type.kind === 'tokenCredential') {
          tokenCred = param.type;
          break;
        }
      }
      if (tokenCred) {
        break;
      }
    }
    if (tokenCred) {
      break;
    }
  }

  if (!tokenCred) {
    // cloud config is only required for token auth
    return '';
  }

  let cloudConfig = helpers.contentPreamble(codeModel.packageName);
  cloudConfig += 'import "github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"\n\n';

  cloudConfig += '// ServiceName is the [cloud.ServiceName] for this module, used to identify the respective [cloud.ServiceConfiguration].\n';

  let serviceName: string;
  if (codeModel.options.module) {
    serviceName = codeModel.options.module;
  } else if (codeModel.options.containingModule) {
    serviceName = `${codeModel.options.containingModule}/${codeModel.packageName}`;
  } else {
    throw new Error('unreachable');
  }

  const azureSdkPrefix = 'github.com/Azure/azure-sdk-for-go/sdk/';
  if (serviceName.startsWith(azureSdkPrefix)) {
    serviceName = serviceName.substring(azureSdkPrefix.length);
  }
  cloudConfig += `const ServiceName cloud.ServiceName = "${serviceName}"\n\n`;

  // we omit the Endpoint field as all client constructors take an endpoint parameter
  cloudConfig += `func init() {\n`;
  cloudConfig += 'cloud.AzurePublic.Services[ServiceName] = cloud.ServiceConfiguration{\n';
  // we assume a single scope. this is enforced when adapting the data from tcgc
  cloudConfig += `\t\tAudience: "${helpers.splitScope(tokenCred.scopes[0]).audience}",\n`;
  cloudConfig += '\t}\n';
  cloudConfig += `}\n\n`;

  return cloudConfig;
}
