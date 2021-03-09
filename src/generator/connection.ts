/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel, Parameter } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { contentPreamble, formatParameterTypeName } from './helpers';
import { ImportManager } from './imports';

// generates content for connection.go
export async function generateConnection(session: Session<CodeModel>): Promise<string> {
  if (!<boolean>session.model.language.go!.azureARM) {
    // add standard imports
    imports.add('fmt');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  }

  let text = await contentPreamble(session);
  // content generation can add to the imports list, so execute it before emitting any text
  const content = generateContent(session);
  text += imports.text();
  if (session.model.security.authenticationRequired && !<boolean>session.model.language.go!.azureARM) {
    const scope = await session.getValue('credential-scope');
    text += `const scope = "${scope}"\n`;
  }
  text += content;
  return text;
}

function generateContent(session: Session<CodeModel>): string {
  let text = `const telemetryInfo = "azsdk-go-${session.model.language.go!.packageName}/<version>"\n`;
  if (<boolean>session.model.language.go!.azureARM) {
    // use the Connection type in armcore instead of generating one
    return text;
  }
  const forceExports = <boolean>session.model.language.go!.exportClients;
  const isARM = session.model.language.go!.openApiType === 'arm';
  let connectionOptions = 'ConnectionOptions';
  if (!isARM && !forceExports) {
    connectionOptions = connectionOptions.uncapitalize();
  }
  text += `// ${connectionOptions} contains configuration settings for the connection's pipeline.\n`;
  text += '// All zero-value fields will be initialized with their default values.\n';
  text += `type ${connectionOptions} struct {\n`;
  text += '\t// HTTPClient sets the transport for making HTTP requests.\n';
  text += '\tHTTPClient azcore.Transport\n';
  text += '\t// Retry configures the built-in retry policy behavior.\n';
  text += '\tRetry azcore.RetryOptions\n';
  text += '\t// Telemetry configures the built-in telemetry policy behavior.\n';
  text += '\tTelemetry azcore.TelemetryOptions\n';
  text += '\t// Logging configures the built-in logging policy behavior.\n';
  text += '\tLogging azcore.LogOptions\n';
  text += '}\n\n';

  text += `func (c *${connectionOptions}) telemetryOptions() *azcore.TelemetryOptions {\n`;
  text += '\tto := c.Telemetry\n';
  text += '\tif to.Value == "" {\n';
  text += '\t\tto.Value = telemetryInfo\n';
  text += '\t} else {\n';
  text += '\t\tto.Value = fmt.Sprintf("%s %s", telemetryInfo, to.Value)\n';
  text += '\t}\n';
  text += '\treturn &to\n';
  text += '}\n\n';

  // Connection
  let connection = 'Connection';
  let defaultEndpoint = 'DefaultEndpoint';
  let newDefaultConnection = 'NewDefaultConnection';
  let newConnection = 'NewConnection';
  let newConnectionWithPipeline = 'NewConnectionWithPipeline';
  if (!isARM && !forceExports) {
    connection = connection.uncapitalize();
    defaultEndpoint = defaultEndpoint.uncapitalize();
    newDefaultConnection = newDefaultConnection.uncapitalize();
    newConnection = newConnection.uncapitalize();
    newConnectionWithPipeline = newConnectionWithPipeline.uncapitalize();
  }
  if (session.model.info.description) {
    text += `// ${connection} - ${session.model.info.description}\n`;
  }

  text += `type ${connection} struct {\n`;
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
    text += `// ${newDefaultConnection} creates an instance of the ${connection} type using the ${defaultEndpoint}.\n`;
    text += '// Pass nil to accept the default options; this is the same as passing a zero-value options.\n';
    text += `func ${newDefaultConnection}(${credParam}options *${connectionOptions}) *${connection} {\n`;
    let cred = 'cred, ';
    if (!session.model.security.authenticationRequired) {
      cred = '';
    }
    text += `\treturn ${newConnection}(${defaultEndpoint}, ${cred}options)\n`;
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

  text += `// ${newConnection} creates an instance of the ${connection} type with the specified endpoint.\n`;
  text += '// Pass nil to accept the default options; this is the same as passing a zero-value options.\n';
  text += `func ${newConnection}(${ctorParamsSig}, ${credParam}options *${connectionOptions}) *${connection} {\n`;
  text += '\tif options == nil {\n';
  text += `\t\toptions = &${connectionOptions}{}\n`;
  text += '\t}\n';
  const telemetryPolicy = 'azcore.NewTelemetryPolicy(options.telemetryOptions())';
  const retryPolicy = 'azcore.NewRetryPolicy(&options.Retry)';
  const credPolicy = 'cred.AuthenticationPolicy(azcore.AuthenticationPolicyOptions{Options: azcore.TokenRequestOptions{Scopes: []string{scope}}})';
  const logPolicy = 'azcore.NewLogPolicy(&options.Logging))';
  text += '\tp := azcore.NewPipeline(options.HTTPClient,\n';
  text += `\t\t${telemetryPolicy},\n`;
  text += `\t\t${retryPolicy},\n`;
  if (session.model.security.authenticationRequired) {
    text += `\t\t${credPolicy},\n`;
  }
  text += `\t\t${logPolicy}\n`;
  text += `\treturn ${newConnectionWithPipeline}(${ctorParams}, p)\n`;
  text += '}\n\n';

  text += `// ${newConnectionWithPipeline} creates an instance of the ${connection} type with the specified endpoint and pipeline.\n`;
  text += `func ${newConnectionWithPipeline}(${ctorParamsSig}, p azcore.Pipeline) *${connection} {\n`;
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
    text += `\treturn &${connection}{u: ${hostURL}, p: p}\n`;
    text += '}\n\n';
    text += '// Endpoint returns the connection\'s endpoint.\n';
    text += `func (c *${connection}) Endpoint() string {\n`;
    text += '\treturn c.u\n';
    text += '}\n\n';
  } else {
    // complex case, full URL will be constructed and parsed in operations
    text += `\tclient := &${connection}{\n`;
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
    text += '}\n\n';
    for (const hostParam of values(hostParams)) {
      const hostParamFunc = (<string>hostParam.language.go!.name).capitalize();
      text += `// ${hostParamFunc} returns part of the parameterized host.\n`;
      text += `func (c *${connection}) ${hostParamFunc}() string {\n`;
      text += `\treturn c.${hostParam.language.go!.name}\n`;
      text += '}\n\n';
    }
  }
  text += '// Pipeline returns the connection\'s pipeline.\n';
  text += `func (c *${connection}) Pipeline() (azcore.Pipeline) {\n`;
  text += '\treturn c.p\n';
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
