/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { camelCase } from '@azure-tools/codegen';
import { CodeModel, Parameter } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { contentPreamble, formatParameterTypeName, sortParametersByRequired } from './helpers';
import { ImportManager } from './imports';

// generates content for client.go
export async function generateClient(session: Session<CodeModel>): Promise<string> {
  // add standard imports
  imports.add('fmt');
  imports.add('net/url');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  imports.add('strings');

  let text = await contentPreamble(session);
  text += imports.text();

  if (session.model.security.authenticationRequired) {
    const scope = await session.getValue('credential-scope');
    text += `const scope = "${scope}"\n`;
  }
  text += `const telemetryInfo = "azsdk-go-${session.model.language.go!.packageName}/<version>"\n`;
  const exportClient = await session.getValue('export-client', true);
  let clientOptions = 'ClientOptions';
  let defaultClientOptions = 'DefaultClientOptions';
  if (!exportClient) {
    clientOptions = 'clientOptions';
    defaultClientOptions = 'defaultClientOptions';
  }
  text += `// ClientOptions contains configuration settings for the default client's pipeline.\n`;
  text += `type ${clientOptions} struct {\n`;
  text += '\t// HTTPClient sets the transport for making HTTP requests.\n';
  text += '\tHTTPClient azcore.Transport\n';
  text += '\t// LogOptions configures the built-in request logging policy behavior.\n';
  text += '\tLogOptions azcore.RequestLogOptions\n';
  text += '\t// Retry configures the built-in retry policy behavior.\n';
  text += '\tRetry azcore.RetryOptions\n';
  text += '\t// Telemetry configures the built-in telemetry policy behavior.\n';
  text += '\tTelemetry azcore.TelemetryOptions\n';
  text += '\t// ApplicationID is an application-specific identification string used in telemetry.\n';
  text += '\t// It has a maximum length of 24 characters and must not contain any spaces.\n';
  text += '\tApplicationID string\n';
  text += '}\n\n';
  text += `// ${defaultClientOptions} creates a ${clientOptions} type initialized with default values.\n`;
  text += `func ${defaultClientOptions}() ${clientOptions} {\n`;
  text += `\treturn ${clientOptions}{\n`;
  text += '\t\tHTTPClient: azcore.DefaultHTTPClientTransport(),\n';
  text += '\t\tRetry: azcore.DefaultRetryOptions(),\n';
  text += '\t}\n';
  text += '}\n\n';

  text += 'func (c *ClientOptions) telemetryOptions() azcore.TelemetryOptions {\n';
  text += '\tt := telemetryInfo\n';
  text += '\tif c.ApplicationID != "" {\n';
  text += '\t\ta := strings.ReplaceAll(c.ApplicationID, " ", "/")\n';
  text += '\t\tif len(a) > 24 {\n';
  text += '\t\t\ta = a[:24]\n';
  text += '\t\t}\n';
  text += '\t\tt = fmt.Sprintf("%s %s", a, telemetryInfo)\n';
  text += '\t}\n';
  text += '\tif c.Telemetry.Value == "" {\n';
  text += '\t\treturn azcore.TelemetryOptions{Value: t}\n';
  text += '\t}\n';
  text += '\treturn azcore.TelemetryOptions{Value: fmt.Sprintf("%s %s", c.Telemetry.Value, t)}\n';
  text += '}\n\n';

  // Client
  let client = 'Client';
  let defaultEndpoint = 'DefaultEndpoint';
  let newDefaultClient = 'NewDefaultClient';
  let newClient = 'NewClient';
  let newClientWithPipeline = 'NewClientWithPipeline';
  if (!exportClient) {
    client = 'client';
    defaultEndpoint = 'defaultEndpoint';
    newDefaultClient = 'newDefaultClient';
    newClient = 'newClient';
    newClientWithPipeline = 'newClientWithPipeline';
  }
  if (session.model.info.description) {
    text += `// ${client} - ${session.model.info.description}\n`;
  }
  text += `type ${client} struct {\n`;
  text += `\t${urlVar} *url.URL\n`;
  text += `\t${pipelineVar} azcore.Pipeline\n`;
  text += '}\n\n';

  const endpoint = getDefaultEndpoint(session.model.globalParameters);
  let credParam = 'cred azcore.Credential, ';
  if (!session.model.security.authenticationRequired) {
    credParam = '';
  }
  if (endpoint) {
    text += `// ${defaultEndpoint} is the default service endpoint.\n`;
    text += `const ${defaultEndpoint} = "${endpoint}"\n\n`;
    text += `// ${newDefaultClient} creates an instance of the ${client} type using the ${defaultEndpoint}.\n`;
    text += `func ${newDefaultClient}(${credParam}options *${clientOptions}) (*${client}, error) {\n`;
    let cred = 'cred, ';
    if (!session.model.security.authenticationRequired) {
      cred = '';
    }
    text += `\treturn ${newClient}(${defaultEndpoint}, ${cred}options)\n`;
    text += '}\n\n';
  }

  text += `// ${newClient} creates an instance of the ${client} type with the specified endpoint.\n`;
  text += `func ${newClient}(endpoint string, ${credParam}options *${clientOptions}) (*${client}, error) {\n`;
  text += '\tif options == nil {\n';
  text += `\t\to := ${defaultClientOptions}()\n`;
  text += '\t\toptions = &o\n';
  text += '\t}\n';
  text += '\tp := azcore.NewPipeline(options.HTTPClient,\n';
  text += '\t\tazcore.NewTelemetryPolicy(options.telemetryOptions()),\n';
  text += '\t\tazcore.NewUniqueRequestIDPolicy(),\n';
  text += '\t\tazcore.NewRetryPolicy(&options.Retry),\n';
  if (session.model.security.authenticationRequired) {
    text += '\t\tcred.AuthenticationPolicy(azcore.AuthenticationPolicyOptions{Options: azcore.TokenRequestOptions{Scopes: []string{scope}}}),\n';
  }
  text += '\t\tazcore.NewRequestLogPolicy(options.LogOptions))\n';
  text += `\treturn ${newClientWithPipeline}(endpoint, p)\n`;
  text += '}\n\n';

  text += `// ${newClientWithPipeline} creates an instance of the ${client} type with the specified endpoint and pipeline.\n`;
  text += `func ${newClientWithPipeline}(endpoint string, ${pipelineVar} azcore.Pipeline) (*${client}, error) {\n`;
  text += `\t${urlVar}, err := url.Parse(endpoint)\n`;
  text += '\tif err != nil {\n';
  text += '\t\treturn nil, err\n';
  text += '\t}\n';
  text += `\tif ${urlVar}.Scheme == "" {\n`;
  text += '\t\treturn nil, fmt.Errorf("no scheme detected in endpoint %s", endpoint)\n';
  text += '\t}\n';
  text += `\treturn &${client}{${urlVar}: ${urlVar}, ${pipelineVar}: ${pipelineVar}}, nil\n`;
  text += '}\n\n';

  for (const group of values(session.model.operationGroups)) {
    const clientLiterals = [`${client}: client`];
    const methodParams = new Array<string>();
    // add client params to the operation group getter method
    if (group.language.go!.clientParams) {
      const clientParams = <Array<Parameter>>group.language.go!.clientParams;
      clientParams.sort(sortParametersByRequired);
      for (const clientParam of values(clientParams)) {
        clientLiterals.push(`${clientParam.language.go!.name}: ${clientParam.language.go!.name}`);
        methodParams.push(`${clientParam.language.go!.name} ${formatParameterTypeName(clientParam)}`);
      }
    }
    text += `// ${group.language.go!.clientName} returns the ${group.language.go!.clientName} associated with this client.\n`;
    text += `func (client *${client}) ${group.language.go!.clientName}(${methodParams.join(', ')}) ${group.language.go!.clientName} {\n`;
    text += `\treturn &${group.operations[0].language.go!.clientName}{${clientLiterals.join(', ')}}\n`;
    text += '}\n\n';
  }

  return text;
}

function getDefaultEndpoint(params?: Parameter[]) {
  for (const param of values(params)) {
    if (param.language.go!.name === '$host' || param.language.go!.name === 'host') {
      return param.clientDefaultValue;
    }
  }
}

// the list of packages to import
const imports = new ImportManager();
const urlVar = 'u';
const pipelineVar = 'p';
