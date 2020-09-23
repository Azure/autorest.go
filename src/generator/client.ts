/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { CodeModel, Parameter } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { contentPreamble, formatParameterTypeName } from './helpers';
import { ImportManager } from './imports';
import { exportClients } from '../common/helpers'

// generates content for client.go
export async function generateClient(session: Session<CodeModel>): Promise<string> {
  // add standard imports
  imports.add('fmt');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');

  let text = await contentPreamble(session);
  const exportClient = await exportClients(session);
  // content generation can add to the imports list, so execute it before emitting any text
  const content = generateContent(session, exportClient);
  text += imports.text();
  if (session.model.security.authenticationRequired) {
    const scope = await session.getValue('credential-scope');
    text += `const scope = "${scope}"\n`;
  }
  text += content;
  return text;
}

function generateContent(session: Session<CodeModel>, exportClient: boolean): string {
  const isARM = session.model.language.go!.openApiType === 'arm';
  let text = `const telemetryInfo = "azsdk-go-${session.model.language.go!.packageName}/<version>"\n`;
  let clientOptions = 'ClientOptions';
  let defaultClientOptions = 'DefaultClientOptions';
  if (!exportClient) {
    clientOptions = 'clientOptions';
    defaultClientOptions = 'defaultClientOptions';
  }
  text += `// ${clientOptions} contains configuration settings for the default client's pipeline.\n`;
  text += `type ${clientOptions} struct {\n`;
  text += '\t// HTTPClient sets the transport for making HTTP requests.\n';
  text += '\tHTTPClient azcore.Transport\n';
  text += '\t// Retry configures the built-in retry policy behavior.\n';
  text += '\tRetry azcore.RetryOptions\n';
  text += '\t// Telemetry configures the built-in telemetry policy behavior.\n';
  text += '\tTelemetry azcore.TelemetryOptions\n';
  if (isARM && session.model.security.authenticationRequired) {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/armcore');
    text += '\t// RegisterRPOptions configures the built-in RP registration policy behavior.\n';
    text += '\tRegisterRPOptions armcore.RegistrationOptions\n';
  }
  text += '}\n\n';
  text += `// ${defaultClientOptions} creates a ${clientOptions} type initialized with default values.\n`;
  text += `func ${defaultClientOptions}() ${clientOptions} {\n`;
  text += `\treturn ${clientOptions}{\n`;
  text += '\t\tHTTPClient: azcore.DefaultHTTPClientTransport(),\n';
  text += '\t\tRetry: azcore.DefaultRetryOptions(),\n';
  if (isARM && session.model.security.authenticationRequired) {
    text += '\t\tRegisterRPOptions: armcore.DefaultRegistrationOptions(),\n';
  }
  text += '\t}\n';
  text += '}\n\n';

  text += `func (c *${clientOptions}) telemetryOptions() azcore.TelemetryOptions {\n`;
  text += '\tto := c.Telemetry\n';
  text += '\tif to.Value == "" {\n';
  text += '\t\tto.Value = telemetryInfo\n';
  text += '\t} else {\n';
  text += '\t\tto.Value = fmt.Sprintf("%s %s", telemetryInfo, to.Value)\n';
  text += '\t}\n';
  text += '\treturn to\n';
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
  if (session.model.language.go!.hostParams && session.model.language.go!.complexHostParams) {
    // complex host params means we have to construct and parse the
    // URL in the operation.  place all host params in the client.
    const hostParams = <Array<Parameter>>session.model.language.go!.hostParams;
    for (const param of values(hostParams)) {
      text += `\t${param.language.go!.name} ${param.schema.language.go!.name}\n`;
    }
  } else {
    text += `\tu string\n`;
  }
  text += `\tp azcore.Pipeline\n`;
  text += '}\n\n';

  let credParam = 'cred azcore.Credential, ';
  if (!session.model.security.authenticationRequired) {
    credParam = '';
  }
  const endpoint = getDefaultEndpoint(session.model.globalParameters);
  if (endpoint) {
    text += `// ${defaultEndpoint} is the default service endpoint.\n`;
    text += `const ${defaultEndpoint} = "${endpoint}"\n\n`;
    text += `// ${newDefaultClient} creates an instance of the ${client} type using the ${defaultEndpoint}.\n`;
    text += `func ${newDefaultClient}(${credParam}options *${clientOptions}) *${client} {\n`;
    let cred = 'cred, ';
    if (!session.model.security.authenticationRequired) {
      cred = '';
    }
    text += `\treturn ${newClient}(${defaultEndpoint}, ${cred}options)\n`;
    text += '}\n\n';
  }

  // build the set of ctor params based on swagger host or parameterized host
  var ctorParamsSig: string;
  var ctorParams: string;
  if (session.model.language.go!.hostParams) {
    // parameterized host
    const hostParams = <Array<Parameter>>session.model.language.go!.hostParams;
    const fullParams = new Array<string>();
    const params = new Array<string>();
    for (const param of values(hostParams)) {
      const paramName = param.language.go!.name;
      fullParams.push(`${paramName} ${formatParameterTypeName(param)}`);
      params.push(paramName);
    }
    ctorParamsSig = fullParams.join(', ');
    ctorParams = params.join(', ');
  } else {
    // swagger host
    const hostParam = getHostParam(session.model.globalParameters);
    const hostParamName = hostParam!.language.go!.name;
    ctorParamsSig = `${hostParamName} ${hostParam!.schema.language.go!.name}`;
    ctorParams = hostParamName;
  }

  text += `// ${newClient} creates an instance of the ${client} type with the specified endpoint.\n`;
  text += `func ${newClient}(${ctorParamsSig}, ${credParam}options *${clientOptions}) *${client} {\n`;
  text += '\tif options == nil {\n';
  text += `\t\to := ${defaultClientOptions}()\n`;
  text += '\t\toptions = &o\n';
  text += '\t}\n';
  const telemetryPolicy = 'azcore.NewTelemetryPolicy(options.telemetryOptions())';
  const reqIDPolicy = 'azcore.NewUniqueRequestIDPolicy()';
  const retryPolicy = 'azcore.NewRetryPolicy(&options.Retry)';
  const credPolicy = 'cred.AuthenticationPolicy(azcore.AuthenticationPolicyOptions{Options: azcore.TokenRequestOptions{Scopes: []string{scope}}})';
  const logPolicy = 'azcore.NewRequestLogPolicy(nil))';
  // ARM will optionally inject the RP registration policy into the pipeline
  if (isARM && session.model.security.authenticationRequired) {
    text += '\tpolicies := []azcore.Policy{\n';
    text += `\t\t${telemetryPolicy},\n`;
    text += `\t\t${reqIDPolicy},\n`;
    text += '\t}\n';
    // RP registration policy must appear before the retry policy
    text += '\tpolicies = append(policies, armcore.NewRPRegistrationPolicy(cred, &options.RegisterRPOptions))\n';
    text += '\tpolicies = append(policies,\n';
    text += `\t\t${retryPolicy},\n`;
    // ARM implies authentication is required
    text += `\t\t${credPolicy},\n`;
    text += `\t\t${logPolicy}\n`;
    text += '\tp := azcore.NewPipeline(options.HTTPClient, policies...)\n';
  } else {
    text += '\tp := azcore.NewPipeline(options.HTTPClient,\n';
    text += `\t\t${telemetryPolicy},\n`;
    text += `\t\t${reqIDPolicy},\n`;
    text += `\t\t${retryPolicy},\n`;
    if (session.model.security.authenticationRequired) {
      text += `\t\t${credPolicy},\n`;
    }
    text += `\t\t${logPolicy}\n`;
  }
  text += `\treturn ${newClientWithPipeline}(${ctorParams}, p)\n`;
  text += '}\n\n';

  text += `// ${newClientWithPipeline} creates an instance of the ${client} type with the specified endpoint and pipeline.\n`;
  text += `func ${newClientWithPipeline}(${ctorParamsSig}, p azcore.Pipeline) *${client} {\n`;
  if (!session.model.language.go!.complexHostParams) {
    // simple case, construct the full host here
    var hostURL: string;
    const uriTemplate = <string>session.model.operationGroups[0].operations[0].requests![0].protocol.http!.uri;
    // if the uriTemplate is simply {whatever} then we can skip doing the strings.ReplaceAll thing.
    if (session.model.language.go!.hostParams && !uriTemplate.match(/^\{\w+\}$/)) {
      // parameterized host
      imports.add('strings');
      text += `\thostURL := "${uriTemplate}"\n`;
      const hostParams = <Array<Parameter>>session.model.language.go!.hostParams;
      for (const hostParam of values(hostParams)) {
        // dereference optional params
        let pointer = '';
        if (hostParam.clientDefaultValue) {
          pointer = '*';
          text += `\tif ${hostParam.language.go!.name} == nil {\n`;
          text += `\t\tdefaultValue := "${hostParam.clientDefaultValue}"\n`;
          text += `\t\t${hostParam.language.go!.name} = &defaultValue\n`;
          text += '\t}\n';
        }
        text += `\thostURL = strings.ReplaceAll(hostURL, "{${hostParam.language.go!.serializedName}}", ${pointer}${hostParam.language.go!.name})\n`;
      }
      hostURL = 'hostURL';
    } else {
      // swagger host, the host URL is the only ctor param
      hostURL = ctorParams;
    }
    text += `\treturn &${client}{u: ${hostURL}, p: p}\n`;
  } else {
    // complex case, full URL will be constructed and parsed in operations
    text += `\tclient := &${client}{\n`;
    text += '\t\tp: p,\n';
    const hostParams = <Array<Parameter>>session.model.language.go!.hostParams;
    for (const hostParam of values(hostParams)) {
      let val = hostParam.language.go!.name;
      if (hostParam.clientDefaultValue) {
        val = `"${hostParam.clientDefaultValue}"`;
      }
      text += `\t\t${hostParam.language.go!.name}: ${val},\n`;
    }
    text += '\t}\n';
    // handle optional host params
    for (const hostParam of values(hostParams)) {
      if (hostParam.clientDefaultValue) {
        text += `\tif ${hostParam.language.go!.name} != nil {\n`;
        text += `\t\tclient.${hostParam.language.go!.name} = *${hostParam.language.go!.name}\n`;
        text += '\t}\n';
      }
    }
    text += '\treturn client\n';
  }
  text += '}\n\n';
  return text;
}

function getDefaultEndpoint(params?: Parameter[]) {
  for (const param of values(params)) {
    if (param.language.default.name === '$host') {
      return param.clientDefaultValue;
    }
  }
}

function getHostParam(params?: Parameter[]): Parameter | undefined {
  for (const param of values(params)) {
    if (param.language.default.name === '$host') {
      return param;
    }
  }
  return undefined;
}

// the list of packages to import
const imports = new ImportManager();
