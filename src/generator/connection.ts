/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { ChoiceSchema, ChoiceValue, CodeModel, Parameter, SchemaType, SealedChoiceSchema } from '@autorest/codemodel';
import { length, values } from '@azure-tools/linq';
import { contentPreamble, formatParameterTypeName } from './helpers';
import { ImportManager } from './imports';

// generates content for connection.go
export async function generateConnection(session: Session<CodeModel>): Promise<string> {
  // ARM doesn't use a Connection type
  if (length(session.model.operationGroups) === 0 || <boolean>session.model.language.go!.azureARM) {
    return '';
  }
  // the list of packages to import
  const imports = new ImportManager();
  // add standard imports
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');

  let text = await contentPreamble(session);
  // content generation can add to the imports list, so execute it before emitting any text
  const content = generateContent(session, imports);
  text += imports.text();
  text += content;
  return text;
}

function generateContent(session: Session<CodeModel>, imports: ImportManager): string {
  let text = '';
  const forceExports = <boolean>session.model.language.go!.exportClients;

  // Connection
  let connection = 'Connection';
  let defaultEndpoint = 'DefaultEndpoint';
  let newDefaultConnection = 'NewDefaultConnection';
  let newConnection = 'NewConnection';
  if (!forceExports) {
    connection = connection.uncapitalize();
    defaultEndpoint = defaultEndpoint.uncapitalize();
    newDefaultConnection = newDefaultConnection.uncapitalize();
    newConnection = newConnection.uncapitalize();
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
  } else if (session.model.language.go!.hostParams || getHostParam(session.model.globalParameters)) {
    // there's a client host param
    text += `\tu string\n`;
  }
  text += `\tp runtime.Pipeline\n`;
  text += '}\n\n';

  const endpoint = getDefaultEndpoint(session.model.globalParameters);
  if (endpoint) {
    text += `// ${defaultEndpoint} is the default service endpoint.\n`;
    text += `const ${defaultEndpoint} = "${endpoint}"\n\n`;
    text += `// ${newDefaultConnection} creates an instance of the ${connection} type using the ${defaultEndpoint}.\n`;
    text += '// Pass nil to accept the default options; this is the same as passing a zero-value options.\n';
    text += `func ${newDefaultConnection}(options *azcore.ClientOptions) *${connection} {\n`;
    text += `\treturn ${newConnection}(${defaultEndpoint}, options)\n`;
    text += '}\n\n';
  }

  // build the set of ctor params based on swagger host or parameterized host
  var ctorParamsSig: string;
  var ctorParams: string;
  if (session.model.language.go!.hostParams) {
    // client parameterized host
    const hostParams = <Array<Parameter>>session.model.language.go!.hostParams;
    const fullParams = new Array<string>();
    const params = new Array<string>();
    for (const param of values(hostParams)) {
      const paramName = param.language.go!.name;
      fullParams.push(`${paramName} ${formatParameterTypeName(param)}`);
      params.push(paramName);
    }
    ctorParamsSig = `${fullParams.join(', ')}, `;
    ctorParams = params.join(', ');
  } else if (getHostParam(session.model.globalParameters)) {
    // swagger host
    const hostParam = getHostParam(session.model.globalParameters);
    const hostParamName = hostParam!.language.go!.name;
    ctorParamsSig = `${hostParamName} ${hostParam!.schema.language.go!.name}, `;
    ctorParams = hostParamName;
  } else {
    // method parameterized host
    ctorParamsSig = ctorParams = '';
  }

  text += `// ${newConnection} creates an instance of the ${connection} type with the specified endpoint.\n`;
  text += '// Pass nil to accept the default options; this is the same as passing a zero-value options.\n';
  text += `func ${newConnection}(${ctorParamsSig}options *azcore.ClientOptions) *${connection} {\n`;
  text += '\tcp := azcore.ClientOptions{}\n';
  text += '\tif options != nil {\n';
  text += '\t\tcp = *options\n';
  text += '\t}\n';
  const pipeline = 'runtime.NewPipeline(module, version, nil, nil, &cp)';
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
          text += `\t\tdefaultValue := ${getClientDefaultValue(hostParam)}\n`;
          text += `\t\t${hostParam.language.go!.name} = &defaultValue\n`;
          text += '\t}\n';
        }
        text += `\thostURL = strings.ReplaceAll(hostURL, "{${hostParam.language.go!.serializedName}}", `;
        switch (hostParam.schema.type) {
          case SchemaType.Choice:
          case SchemaType.SealedChoice:
            text += `string(${pointer}${hostParam.language.go!.name}))\n`;
            break;
          case SchemaType.String:
            text += `${pointer}${hostParam.language.go!.name})\n`;
            break;
          default:
            imports.add('fmt');
            text += `fmt.Sprint(${pointer}${hostParam.language.go!.name}))\n`;
            break;
        }
      }
      hostURL = 'u: hostURL, ';
    } else if (ctorParams !== '') {
      // swagger host, the host URL is the only ctor param
      hostURL = `u: ${ctorParams}, `;
    } else {
      // method parameterized host
      hostURL = '';
    }
    text += `\treturn &${connection}{${hostURL}p: ${pipeline}}\n`;
    text += '}\n\n';
    text += '// Endpoint returns the connection\'s endpoint.\n';
    text += `func (c *${connection}) Endpoint() string {\n`;
    text += '\treturn c.u\n';
    text += '}\n\n';
  } else {
    // complex case, full URL will be constructed and parsed in operations
    text += `\tclient := &${connection}{\n`;
    text += `\t\tp: ${pipeline},\n`;
    const hostParams = <Array<Parameter>>session.model.language.go!.hostParams;
    for (const hostParam of values(hostParams)) {
      let val = hostParam.language.go!.name;
      if (hostParam.clientDefaultValue) {
        val = getClientDefaultValue(hostParam);
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
  text += `func (c *${connection}) Pipeline() (runtime.Pipeline) {\n`;
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

// returns the clientDefaultValue of the specified param.
// this is usually the value in quotes (i.e. a string) however
// it could also be a constant.
function getClientDefaultValue(param: Parameter): string {
  const getChoiceValue = function (choices: ChoiceValue[]): string {
    // find the corresponding const type name
    for (const choice of values(choices)) {
      if (choice.value === param.clientDefaultValue) {
        return choice.language.go!.name;
      }
    }
    throw new Error(`failed to find matching constant for default value ${param.clientDefaultValue}`);
  }
  switch (param.schema.type) {
    case SchemaType.Choice:
      return getChoiceValue((<ChoiceSchema>param.schema).choices);
    case SchemaType.SealedChoice:
      return getChoiceValue((<SealedChoiceSchema>param.schema).choices);
    case SchemaType.String:
      return `"${param.clientDefaultValue}"`;
    default:
      return param.clientDefaultValue;
  }
}
