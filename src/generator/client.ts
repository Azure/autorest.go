/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { camelCase } from '@azure-tools/codegen';
import { CodeModel, Parameter } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { ContentPreamble, formatParamInfoTypeName, ImportManager, ParamInfo, sortParamInfoByRequired } from './helpers';

// generates content for client.go
export async function generateClient(session: Session<CodeModel>): Promise<string> {
  // add standard imports
  imports.add('fmt');
  imports.add('net/url');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');

  let text = await ContentPreamble(session);
  text += imports.text();

  // ClientOptions
  text += 'type ClientOptions struct {\n';
  text += '\t// HTTPClient sets the transport for making HTTP requests.\n';
  text += '\tHTTPClient azcore.Transport\n';
  text += '\t// LogOptions configures the built-in request logging policy behavior.\n';
  text += '\tLogOptions azcore.RequestLogOptions\n';
  text += '\t// Retry configures the built-in retry policy behavior.\n';
  text += '\tRetry azcore.RetryOptions\n';
  text += '\t// Telemetry configures the built-in telemetry policy behavior.\n';
  text += '\tTelemetry azcore.TelemetryOptions\n';
  text += '}\n\n';
  text += '// DefaultClientOptions creates a ClientOptions type initialized with default values.\n';
  text += 'func DefaultClientOptions() ClientOptions {\n';
  text += '\treturn ClientOptions{\n';
  text += '\t\tHTTPClient: azcore.DefaultHTTPClientTransport(),\n';
  text += '\t\tRetry: azcore.DefaultRetryOptions(),\n';
  text += '\t}\n';
  text += '}\n\n';

  // Client
  if (session.model.info.description) {
    text += `// Client - ${session.model.info.description}\n`;
  }
  text += 'type Client struct {\n';
  text += `\t${urlVar} *url.URL\n`;
  text += `\t${pipelineVar} azcore.Pipeline\n`;
  text += '}\n\n';

  const endpoint = getDefaultEndpoint(session.model.globalParameters);
  if (endpoint) {
    text += '// DefaultEndpoint is the default service endpoint.\n';
    text += `const DefaultEndpoint = "${endpoint}"\n\n`;
    text += '// NewDefaultClient creates an instance of the Client type using the DefaultEndpoint.\n';
    text += 'func NewDefaultClient(options *ClientOptions) (*Client, error) {\n';
    text += '\treturn NewClient(DefaultEndpoint, options)\n';
    text += '}\n\n';
  }

  text += '// NewClient creates an instance of the Client type with the specified endpoint.\n';
  text += 'func NewClient(endpoint string, options *ClientOptions) (*Client, error) {\n';
  text += '\tif options == nil {\n';
  text += '\t\to := DefaultClientOptions()\n';
  text += '\t\toptions = &o\n';
  text += '\t}\n';
  text += '\tp := azcore.NewPipeline(options.HTTPClient,\n';
  text += '\t\tazcore.NewTelemetryPolicy(options.Telemetry),\n';
  text += '\t\tazcore.NewUniqueRequestIDPolicy(),\n';
  text += '\t\tazcore.NewRetryPolicy(&options.Retry),\n';
  text += '\t\tazcore.NewRequestLogPolicy(options.LogOptions))\n';
  text += '\treturn NewClientWithPipeline(endpoint, p)\n';
  text += '}\n\n';

  text += '// NewClientWithPipeline creates an instance of the Client type with the specified endpoint and pipeline.\n';
  text += `func NewClientWithPipeline(endpoint string, ${pipelineVar} azcore.Pipeline) (*Client, error) {\n`;
  text += `\t${urlVar}, err := url.Parse(endpoint)\n`;
  text += '\tif err != nil {\n';
  text += '\t\treturn nil, err\n';
  text += '\t}\n';
  text += `\tif ${urlVar}.Scheme == "" {\n`;
  text += '\t\treturn nil, fmt.Errorf("no scheme detected in endpoint %s", endpoint)\n';
  text += '\t}\n';
  text += `\treturn &Client{${urlVar}: ${urlVar}, ${pipelineVar}: ${pipelineVar}}, nil\n`;
  text += '}\n\n';

  for (const group of values(session.model.operationGroups)) {
    const clientParams = ['Client: client'];
    const methodParams = new Array<string>();
    // add global params to the operation group getter method
    if (group.language.go!.globals) {
      const globals = <Array<ParamInfo>>group.language.go!.globals;
      globals.sort(sortParamInfoByRequired);
      globals.forEach((value: ParamInfo, index: Number, obj: ParamInfo[]) => {
        if (value.name !== 'urlParameter') {
          clientParams.push(`${value.name}: ${value.name}`);
          methodParams.push(`${value.name} ${formatParamInfoTypeName(value)}`);
        }
      });
    }
    text += `// ${group.language.go!.clientName} returns the ${group.language.go!.clientName} associated with this client.\n`;
    text += `func (client *Client) ${group.language.go!.clientName}(${methodParams.join(', ')}) ${group.language.go!.clientName} {\n`;
    text += `\treturn &${camelCase(group.language.go!.clientName)}{${clientParams.join(', ')}}\n`;
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
