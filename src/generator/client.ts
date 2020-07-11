/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { camelCase } from '@azure-tools/codegen';
import { CodeModel, Parameter, ImplementationLocation } from '@azure-tools/codemodel';
import { values, IterableWithLinq } from '@azure-tools/linq';
import { contentPreamble, formatParameterTypeName, sortParametersByRequired, addParameterizedHostFunctionality, formatParamValue, skipURLEncoding, substituteDiscriminator } from './helpers';
import { schemaTypeToGoType, aggregateParameters } from '../common/helpers';
import { ImportManager } from './imports';

// generates content for client.go
export async function generateClient(session: Session<CodeModel>): Promise<string> {
  // add standard imports
  imports.add('fmt');
  const isARM = session.model.language.go!.openApiType === 'arm';
  if (isARM) {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/armcore');
  }
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  imports.add('strings');
  // initialize variables necessary for adding parameterized host related variables and functionality for the client
  const paramHostInfo = addParameterizedHostFunctionality(session.model.operationGroups);
  let addEndpoint = 'endpoint string, ';
  let passEndpoint = 'endpoint, ';
  if (paramHostInfo.addParamHost) {
    addEndpoint = '';
    passEndpoint = '';
  } else if (!paramHostInfo.addParamHost || paramHostInfo.urlOnClient) {
    imports.add('net/url');
  }

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
  if (isARM) {
    text += '\t// DisableRPRegistration controls if an unregistered resource provider should\n';
    text += '\t// automatically be registered. See https://aka.ms/rps-not-found for more information.\n';
    text += '\t// The default value is false, meaning registration will be attempted.\n';
    text += '\tDisableRPRegistration bool\n';
  }
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
  // initialize clientOnlyParams in order to use them in client parameter declarations and assign values for the defaults later on
  const clientOnlyParams = new Array<Parameter>();
  let clientOnlyParamsFuncSig = '';
  let passClientOnlyParams = '';
  let groupClientParams = new Array<Parameter>();
  // groupClientParams will consolidate a unique set of client only params from each operation group
  for (const group of values(session.model.operationGroups)) {
    if (group.language.go!.clientParams !== undefined) {
      groupClientParams = [...groupClientParams, ...<Array<Parameter>>group.language.go!.clientParams];
    }
  }
  groupClientParams = [...new Set(groupClientParams)];
  // if there are global parameters then check for global params that are only meant to exist on the client
  // and which do not exist on operation groups. The paramters found here will be added onto the client.
  if (session.model.globalParameters) { // TODO move all common params to client and away from operation group
    for (const param of values(session.model.globalParameters)) {
      if (!groupClientParams.includes(param) && param.protocol.http!.in === 'uri') {
        clientOnlyParams.push(param);
        clientOnlyParamsFuncSig += `${param.language.go!.name} ${formatParameterTypeName(param)}, `;
        passClientOnlyParams += `${param.language.go!.name}, `;
      }
    }
  }
  // if the parameterized host functionality is not going to be implemented then wipe client only paramter settings
  if (!paramHostInfo.addParamHost && !paramHostInfo.urlOnClient) {
    clientOnlyParamsFuncSig = '';
    passClientOnlyParams = '';
  }
  text += `type ${client} struct {\n`;
  if (paramHostInfo.addParamHost && !paramHostInfo.urlOnClient) {
    if (clientOnlyParams.length > 0) {
      for (const param of values(clientOnlyParams)) {
        if (param.protocol.http!.in === 'uri') {
          text += `\t${param.language.go!.name} ${schemaTypeToGoType(session.model, param.schema, false)}\n`;
        }
      }
    }
  } else {
    text += `\t${urlVar} *url.URL\n`;
  }
  text += `\t${pipelineVar} azcore.Pipeline\n`;
  text += '}\n\n';

  const endpoint = getDefaultEndpoint(session.model.globalParameters);
  if (endpoint) {
    text += `// ${defaultEndpoint} is the default service endpoint.\n`;
    text += `const ${defaultEndpoint} = "${endpoint}"\n\n`;
  }
  let credParam = 'cred azcore.Credential, ';
  if (!session.model.security.authenticationRequired) {
    credParam = '';
  }
  if (defaultsExist(clientOnlyParams)) {
    text += `// ${newDefaultClient} creates an instance of the ${client} type using the ${defaultEndpoint}.\n`;
    text += `func ${newDefaultClient}(${credParam}options *${clientOptions}) (*${client}, error) {\n`;
    let cred = 'cred, ';
    if (!session.model.security.authenticationRequired) {
      cred = '';
    }
    if (paramHostInfo.addParamHost) {
      text += `\treturn ${newClient}(`;
      for (const param of values(clientOnlyParams)) {
        if (param.clientDefaultValue !== undefined) {
          text += `"${param.clientDefaultValue}", `;
        } else {
          text += '"", ';
        }
      }
      text += `${cred}options)\n`;
    } else {
      // when a parameterized host is not set, the default endpoint is set for the url on the client
      text += `\treturn ${newClient}(${defaultEndpoint}, ${cred}options)\n`;
    }
    text += '}\n\n';
  }

  text += `// ${newClient} creates an instance of the ${client} type with the specified endpoint.\n`;
  text += `func ${newClient}(${clientOnlyParamsFuncSig}${addEndpoint}${credParam}options *${clientOptions}) (*${client}, error) {\n`;
  text += '\tif options == nil {\n';
  text += `\t\to := ${defaultClientOptions}()\n`;
  text += '\t\toptions = &o\n';
  text += '\t}\n';
  const telemetryPolicy = 'azcore.NewTelemetryPolicy(options.telemetryOptions())';
  const reqIDPolicy = 'azcore.NewUniqueRequestIDPolicy()';
  const retryPolicy = 'azcore.NewRetryPolicy(&options.Retry)';
  const credPolicy = 'cred.AuthenticationPolicy(azcore.AuthenticationPolicyOptions{Options: azcore.TokenRequestOptions{Scopes: []string{scope}}})';
  const logPolicy = 'azcore.NewRequestLogPolicy(options.LogOptions))';
  // ARM will optionally inject the RP registration policy into the pipeline
  if (isARM) {
    text += '\tpolicies := []azcore.Policy{\n';
    text += `\t\t${telemetryPolicy},\n`;
    text += `\t\t${reqIDPolicy},\n`;
    text += '\t}\n';
    // RP registration policy must appear before the retry policy
    text += '\tif !options.DisableRPRegistration {\n';
    text += '\t\trpOpts := armcore.DefaultRegistrationOptions()\n';
    text += '\t\trpOpts.HTTPClient = options.HTTPClient\n';
    text += '\t\trpOpts.LogOptions = options.LogOptions\n';
    text += '\t\trpOpts.Retry = options.Retry\n';
    text += '\t\tpolicies = append(policies, armcore.NewRPRegistrationPolicy(cred, &rpOpts))\n';
    text += '\t}\n';
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
  text += `\treturn ${newClientWithPipeline}(${passClientOnlyParams}${passEndpoint}p)\n`;
  text += '}\n\n';

  text += `// ${newClientWithPipeline} creates an instance of the ${client} type with the specified endpoint and pipeline.\n`;
  text += `func ${newClientWithPipeline}(${clientOnlyParamsFuncSig}${addEndpoint}${pipelineVar} azcore.Pipeline) (*${client}, error) {\n`;
  if (!paramHostInfo.addParamHost) {
    text += `\t${urlVar}, err := url.Parse(endpoint)\n`;
    text += '\tif err != nil {\n';
    text += '\t\treturn nil, err\n';
    text += '\t}\n';
    text += `\tif ${urlVar}.Scheme == "" {\n`;
    text += '\t\treturn nil, fmt.Errorf("no scheme detected in endpoint %s", endpoint)\n';
    text += '\t}\n';
    text += `\treturn &${client}{${urlVar}: ${urlVar}, ${pipelineVar}: ${pipelineVar}}, nil\n`;
  } else if (paramHostInfo.urlOnClient) { // TODO I think there's a flaw here if there are client params on the operation groups
    const op = session.model.operationGroups[0].operations[0];
    const group = session.model.operationGroups[0];
    text += `\thostURL := "${op.requests![0].protocol.http!.uri}"\n`;
    for (const pp of values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'uri'; })) {
      imports.add('strings');
      // host references may be found among other variables and therefore call to use the host 
      // specified by the endpoint the client was created with
      if ((pp.implementation === ImplementationLocation.Client && group.language.go!.clientParams === undefined) || (pp.implementation === ImplementationLocation.Client && !(<Array<Parameter>>group.language.go!.clientParams).includes(pp))) {
        text += `\thostURL = strings.ReplaceAll(hostURL, "{${pp.language.go!.serializedName}}", client.Client.${pp.language.go!.serializedName})\n`;
      } else {
        let paramValue = `url.PathEscape(${formatParamValue(pp, imports)})`;
        if (skipURLEncoding(pp)) {
          paramValue = formatParamValue(pp, imports);
        } else {
          imports.add('net/url');
        }
        text += `\thostURL = strings.ReplaceAll(hostURL, "{${pp.language.go!.serializedName}}", ${paramValue})\n`;
      }
    }
    text += `\t${urlVar}, err := url.Parse(hostURL)\n`;
    text += '\tif err != nil {\n';
    text += '\t\treturn nil, err\n';
    text += '\t}\n';
    text += `\tif ${urlVar}.Scheme == "" {\n`;
    text += '\t\treturn nil, fmt.Errorf("no scheme detected in endpoint %s", endpoint)\n';
    text += '\t}\n';
    text += `\treturn &${client}{${urlVar}: ${urlVar}, ${pipelineVar}: ${pipelineVar}}, nil\n`;
  } else if (clientOnlyParams.length > 0) {
    text += `\tclient := &${client}{}\n`;
    text += `\tclient.${pipelineVar} = ${pipelineVar}\n`;
    for (const param of values(clientOnlyParams)) {
      text += `\tclient.${param.language.go!.name} = ${param.language.go!.name}\n`;
    }
    text += '\treturn client, nil\n';
  } else {
    text += `\treturn &${client}{${pipelineVar}: ${pipelineVar}}, nil\n`;
  }

  text += '}\n\n';

  for (const group of values(session.model.operationGroups)) {
    const clientLiterals = [`${client}: client`];
    const methodParams = new Array<string>();
    let defaultValParamComments = '';
    // add client params to the operation group getter method
    if (group.language.go!.clientParams) {
      const clientParams = <Array<Parameter>>group.language.go!.clientParams;
      clientParams.sort(sortParametersByRequired);
      for (const clientParam of values(clientParams)) {
        const param = substituteDiscriminator(clientParam.schema);
        let methodPointer = '';
        let literalPointer = '';
        if (!clientParam.required) {
          methodPointer = '*';
        }
        if (clientParam.clientDefaultValue !== undefined) {
          methodPointer = '*';
          literalPointer = '*';
          defaultValParamComments += `// For ${clientParam.language.go!.name} pass nil to use the default value of "${clientParam.clientDefaultValue}"\n`;
        }
        clientLiterals.push(`${clientParam.language.go!.name}: ${literalPointer}${clientParam.language.go!.name}`);
        methodParams.push(`${clientParam.language.go!.name} ${methodPointer}${param}`);
      }
    }
    text += `// ${group.language.go!.clientName} returns the ${group.language.go!.clientName} associated with this client.\n`;
    if (defaultValParamComments.length > 0) {
      text += defaultValParamComments;
    }
    text += `func (client *${client}) ${group.language.go!.clientName}(${methodParams.join(', ')}) ${group.language.go!.clientName} {\n`;
    for (const mp of values(<Array<Parameter>>group.language.go!.clientParams)) {
      if (mp.clientDefaultValue !== undefined) {
        text += `if ${mp.language.go!.name} == nil {
          *${mp.language.go!.name} = "${mp.clientDefaultValue}"
        }
        `;
      }
    }

    text += `\treturn &${group.operations[0].language.go!.clientName}{${clientLiterals.join(', ')}}\n`;
    text += '}\n\n';
  }

  return text;
}

function defaultsExist(params?: Array<Parameter>): boolean {
  for (const param of values(params)) {
    if (param.clientDefaultValue !== undefined) {
      return true;
    }
  }
  return false;
}

function getDefaultEndpoint(params?: Parameter[]) {
  for (const param of values(params)) {
    if (param.language.go!.name === '$host') {
      return param.clientDefaultValue;
    }
  }
}

// the list of packages to import
const imports = new ImportManager();
const pipelineVar = 'p';
const urlVar = 'u';

// return parameters for client function
export function createParametersSig(clientParams: IterableWithLinq<Parameter>): [string, string] {
  const funcParams = new Array<string>();
  const params = new Array<string>();
  for (const p of values(clientParams)) {
    funcParams.push(`${camelCase(p.language.go!.name)} ${formatParameterTypeName(p)}`);
    params.push(camelCase(p.language.go!.name));
  }
  return [funcParams.join(', '), params.join(', ')];
}
