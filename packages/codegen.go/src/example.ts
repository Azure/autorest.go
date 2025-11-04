/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { camelCase, capitalize } from '@azure-tools/codegen';
import * as go from '../../codemodel.go/src/index.js';
import * as helpers from './helpers.js';
import { ImportManager } from './imports.js';
import { fixUpMethodName } from './operations.js';
import { CodegenError } from './errors.js';

// represents the generated content for an example
export class ExampleContent {
  readonly name: string;
  readonly content: string;

  constructor(name: string, content: string) {
    this.name = name;
    this.content = content;
  }
}

// Creates the content for all <example>.go files
export async function generateExamples(codeModel: go.CodeModel): Promise<Array<ExampleContent>> {
  // generate examples
  const examples = new Array<ExampleContent>();
  if (codeModel.clients.length === 0) {
    return examples;
  }

  const azureARM = codeModel.type === 'azure-arm';

  for (const client of codeModel.clients) {
    // client must be constructable to create a sample
    if (client.instance?.kind != 'constructable') {
      continue;
    }
    const imports = new ImportManager();
    // the list of packages to import
    if (client.methods.length > 0) {
      // add standard imports for clients with methods.
      // clients that are purely hierarchical (i.e. having no APIs) won't need them.
      imports.add('context');
      imports.add('log');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azidentity');
      imports.add(codeModel.options.module!);
    }

    let clientFactoryParams = new Array<go.ClientParameter>();
    if (codeModel.options.factoryGatherAllParams) {
      clientFactoryParams =  helpers.getAllClientParameters(codeModel);
    } else {
      clientFactoryParams = helpers.getCommonClientParameters(codeModel);
    }
    const clientFactoryParamsMap = new Map<string, go.ClientParameter>();
    for (const param of clientFactoryParams) {
      clientFactoryParamsMap.set(param.name, param);
    }

    let exampleText = '';
    for (const method of client.methods) {
      if (method.examples.length === 0) continue;
      for (const example of method.examples) {
        // function signature
        exampleText += `// Generated from example definition: ${example.filePath}\n`;
        const exampleFuncNamePrefix = method.examples.length > 1 ? `_${camelCase(example.name)}` : '';
        exampleText += `func Example${client.name}_${fixUpMethodName(method)}${exampleFuncNamePrefix}() {\n`;

        // create credential
        exampleText += `\tcred, err := azidentity.NewDefaultAzureCredential(nil)\n`;
        exampleText += `\tif err != nil {\n`;
        exampleText += `\t\tlog.Fatalf("failed to obtain a credential: %v", err)\n`;
        exampleText += `\t}\n`;

        // create context
        exampleText += `\tctx := context.Background()\n`;

        // create client
        const clientParameters: go.ParameterExample[] = [];
        for (const param of method.parameters) {
          if (param.location === "client") {
            const clientParam = example.parameters.find(p => p.parameter.name === param.name);
            if (clientParam) {
              clientParameters.push(clientParam);
            }
          }
        }
        // TODO: client optional parameters

        let clientRef = '';
        if (azureARM) {
          // since not all operation has all the client factory required parameters, we need to fake for the missing ones
          const clientFactoryParamsExample: go.ParameterExample[] = [];
          for (const clientParam of clientFactoryParams) {
            const clientFactoryParam = clientParameters.find(p => p.parameter.name === clientParam.name);
            if (clientFactoryParam) {
              clientFactoryParamsExample.push(clientFactoryParam);
            } else {
              clientFactoryParamsExample.push({ parameter: clientParam, value: generateFakeExample(clientParam.type, clientParam.name) });
            }
          }
          exampleText += `\tclientFactory, err := ${codeModel.packageName}.NewClientFactory(${clientFactoryParamsExample.map(p => getExampleValue(codeModel, p.value, '\t', imports, p.parameter.byValue)).join(', ')}, nil)\n`;
          exampleText += `\tif err != nil {\n`;
          exampleText += `\t\tlog.Fatalf("failed to create client: %v", err)\n`;
          exampleText += `\t}\n`;
          clientRef = `clientFactory.${client.instance.constructors[0].name}(`;
          const clientPrivateParameters: go.ParameterExample[] = [];
          for (const clientParam of clientParameters) {
            if (!clientFactoryParamsMap.has(clientParam.parameter.name)) {
              clientPrivateParameters.push(clientParam);
            }
          }
          if (clientPrivateParameters.length > 0) {
            clientRef += `${clientPrivateParameters.map(p => getExampleValue(codeModel, p.value, '\t', imports, p.parameter.byValue).slice(1)).join(', ')}`;
          }
          clientRef += `)`;
        } else {
          exampleText += `\tclient, err := ${codeModel.packageName}.${client.instance.constructors[0].name}(${clientParameters.map(p => getExampleValue(codeModel, p.value, '\t', imports, p.parameter.byValue).slice(1)).join(', ')}, cred, nil)\n`;
          exampleText += `\tif err != nil {\n`;
          exampleText += `\t\tlog.Fatalf("failed to create client: %v", err)\n`;
          exampleText += `\t}\n`;
          clientRef = 'client';
        }

        // call method
        const methodParameters: go.ParameterExample[] = [];
        for (const param of method.parameters) {
          if (param.location === "method") {
            const methodParam = example.parameters.find(p => p.parameter.name === param.name);
            if (methodParam) {
              methodParameters.push(methodParam);
            }
          }
        }

        const methodOptionalParameters = example.optionalParamsGroup.filter(p => p.parameter.location === 'method');
        const checkResponse = example.responseEnvelope !== undefined;

        let methodOptionalParametersText = 'nil';
        if (methodOptionalParameters.length > 0) {
          methodOptionalParametersText = `&${codeModel.packageName}.${method.optionalParamsGroup.groupName}{\n`;
          methodOptionalParametersText += methodOptionalParameters.map(p => `${capitalize(p.parameter.name)}: ${getExampleValue(codeModel, p.value, '\t', imports, p.parameter.byValue).slice(1)}`).join(',\n');
          methodOptionalParametersText += `}`;
        }

        switch (method.kind) {
          case 'lroMethod':
          case 'lroPageableMethod':
            exampleText += `\tpoller, err := ${clientRef}.${fixUpMethodName(method)}(ctx, ${methodParameters.map(p => getExampleValue(codeModel, p.value, '\t', imports, p.parameter.byValue).slice(1)).join(', ')}${methodParameters.length > 0 ? ', ' : ''}${methodOptionalParametersText.split('\n').join('\n\t')})\n`;
            exampleText += `\tif err != nil {\n`;
            exampleText += `\t\tlog.Fatalf("failed to finish the request: %v", err)\n`;
            exampleText += `\t}\n`;

            exampleText += `\t${checkResponse ? 'res' : '_'}, err ${checkResponse ? ':=' : '='} poller.PollUntilDone(ctx, nil)\n`
            exampleText += `\tif err != nil {\n`;
            exampleText += `\t\tlog.Fatalf("failed to pull the result: %v", err)\n`;
            exampleText += `\t}\n`;
            break;
          case 'method':
            exampleText += `\t${checkResponse ? 'res' : '_'}, err ${checkResponse ? ':=' : '='} ${clientRef}.${fixUpMethodName(method)}(ctx, ${methodParameters.map(p => getExampleValue(codeModel, p.value, '\t', imports, p.parameter.byValue).slice(1)).join(', ')}${methodParameters.length > 0 ? ', ' : ''}${methodOptionalParametersText.split('\n').join('\n\t')})\n`;
            exampleText += `\tif err != nil {\n`;
            exampleText += `\t\tlog.Fatalf("failed to finish the request: %v", err)\n`;
            exampleText += `\t}\n`;
            break;
          case 'pageableMethod':
            exampleText += `\tpager := ${clientRef}.${fixUpMethodName(method)}(${methodParameters.map(p => getExampleValue(codeModel, p.value, '\t', imports, p.parameter.byValue).slice(1)).join(', ')}${methodParameters.length > 0 ? ', ' : ''}${methodOptionalParametersText.split('\n').join('\n\t')})\n`;
            break;
          default:
            method satisfies never;
        }

        // check response
        if (method.kind === 'lroPageableMethod' || method.kind === 'pageableMethod') {
          let resultName = 'pager';
          if (method.kind === 'lroPageableMethod') {
            resultName = 'res';
          }
          const itemType = ((method as go.PageableMethod).returns.result as go.ModelResult).modelType.fields.find(f => f.type.kind === 'slice')!;
          exampleText += `\tfor ${resultName}.More() {\n`;
          exampleText += `\t\tpage, err := ${resultName}.NextPage(ctx)\n`;
          exampleText += `\t\tif err != nil {\n`;
          exampleText += `\t\t\tlog.Fatalf("failed to advance page: %v", err)\n`;
          exampleText += `\t\t}\n`;
          exampleText += `\t\tfor _, v := range page.${itemType.name} {\n`;
          exampleText += `\t\t\t// You could use page here. We use blank identifier for just demo purposes.\n`;
          exampleText += `\t\t\t_ = v\n`;
          exampleText += `\t\t}\n`;
          exampleText += `\t\t// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.\n`;
          exampleText += `\t\t// page = ${codeModel.packageName}.${example.responseEnvelope?.response.name}{\n`;
          for (const header of example.responseEnvelope?.headers ?? []) {
            exampleText += `\t\t// \t${header.header.fieldName}: ${getExampleValue(codeModel, header.value, '', undefined, true).split('\n').join('\n\t\t// \t')}\n`;
          }
          exampleText += `\t\t// \t${(example.responseEnvelope?.result.type as go.Model).name}: ${getExampleValue(codeModel, example.responseEnvelope?.result!, '', undefined, true).split('\n').join('\n\t\t// \t')},\n`;
          exampleText += '\t\t// }\n';
          exampleText += `\t}\n`;
        } else if (checkResponse) {
          // if has fieldName, then the result is not a model type
          const fieldName = (method.returns as any).fieldName;
          exampleText += `\t// You could use response here. We use blank identifier for just demo purposes.\n`;
          exampleText += `\t_ = res\n`;

          exampleText += `\t// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.\n`;
          exampleText += `\t// res = ${codeModel.packageName}.${example.responseEnvelope?.response.name}{\n`;
          for (const header of example.responseEnvelope?.headers ?? []) {
            exampleText += `\t// \t${header.header.fieldName}: ${getExampleValue(codeModel, header.value, '', undefined, true).split('\n').join('\n\t// \t')}\n`;
          }
          if (example.responseEnvelope?.result) {
            exampleText += `\t// \t${fieldName ? fieldName : (example.responseEnvelope?.result.type as go.Model).name}: ${getExampleValue(codeModel, example.responseEnvelope.result, '').split('\n').join('\n\t// \t')},\n`;
          }
          exampleText += '\t// }\n';
        }
        exampleText += `}\n\n`;
      }
    }

    // if no example, then do not generate example file
    if (exampleText === '') continue;

    // stitch it all together
    let text = helpers.contentPreamble(codeModel, true, codeModel.packageName + '_test');
    text += imports.text();
    text += exampleText;
    examples.push(new ExampleContent(client.name, text));
  }
  return examples;
}

function getExampleValue(codeModel: go.CodeModel, example: go.ExampleType, indent: string, imports?: ImportManager, byValue: boolean = false, inArray: boolean = false): string {
  switch (example.kind) {
    case 'string': {
      let exampleText = `"${escapeString(example.value)}"`;
      if (example.type.kind === 'constant') {
        exampleText = getConstantValue(codeModel, example.type, example.value);
      } else if (example.type.kind === 'time') {
        exampleText = getTimeValue(example.type, example.value, imports);
      } else if (example.type.kind === 'encodedBytes') {
        exampleText = `[]byte("${escapeString(example.value)}")`
      } else if (example.type.kind === 'literal' && example.type.type.kind === 'constant') {
        exampleText = getConstantValue(codeModel, example.type.type, example.type.literal.value);
      } else if (example.type.kind === 'etag') {
        imports?.add(example.type.module);
        return `${indent}${getRef(byValue)}${go.getTypeDeclaration(example.type, codeModel.packageName)}(${jsonToGo(example.value, '')})`;
      }
      return `${indent}${getPointerValue(example.type, exampleText, byValue, imports)}`;
    }
    case 'number': {
      let exampleText = `${example.value}`;
      switch (example.type.kind) {
        case 'constant':
          exampleText = `${indent}${getConstantValue(codeModel, example.type, example.value)}`;
          break;
        case 'time':
          exampleText = getTimeValue(example.type, example.value, imports);
          break;
      }
      return `${indent}${getPointerValue(example.type, exampleText, byValue, imports)}`;
    }
    case 'boolean': {
      let exampleText = `${example.value}`;
      if (example.type.kind === 'constant') {
        exampleText = `${indent}${getConstantValue(codeModel, example.type, example.value)}`;
      }
      return `${indent}${getPointerValue(example.type, exampleText, byValue, imports)}`;
    }
    case 'null':
      return `${indent}nil`;
    case 'any':
      return jsonToGo(example.value, indent);
    case 'array': {
      const isElementByValue = example.type.elementTypeByValue;
      // if polymorphic, need to add type name in array, so inArray will be set to false
      // if other case, no need to add type name in array, so inArray will be set to true
      const isElementPolymorphic = example.type.elementType.kind === 'interface';
      let exampleText = `${indent}${getRef(byValue)}${go.getTypeDeclaration(example.type, codeModel.packageName)}{\n`;
      for (const element of example.value) {
        exampleText += `${getExampleValue(codeModel, element, indent + '\t', imports, isElementByValue && !isElementPolymorphic, !isElementPolymorphic)},\n`;
      }
      exampleText += `${indent}}`;
      return exampleText;
    }
    case 'dictionary': {
      let exampleText = `${indent}${getRef(byValue)}${go.getTypeDeclaration(example.type, codeModel.packageName)}{\n`;
      const isValueByValue = example.type.valueTypeByValue;
      const isValuePolymorphic = example.type.valueType.kind === 'interface';
      for (const key in example.value) {
        exampleText += `${indent}\t"${key}": ${getExampleValue(codeModel, example.value[key], indent + '\t', imports, isValueByValue && !isValuePolymorphic).slice(indent.length + 1)},\n`;
      }
      exampleText += `${indent}}`;
      return exampleText;
    }
    case 'model': {
      const isModelPolymorphic = example.type.kind === 'polymorphicModel';
      let exampleText = `${indent}${getRef(byValue && !isModelPolymorphic)}${go.getTypeDeclaration(example.type, codeModel.packageName)}{\n`;
      if (inArray) {
        exampleText = `${indent}{\n`;
      }
      for (const field in example.value) {
        const goField = example.type.fields.find(f => f.name === field)!;
        const isFieldByValue = goField.byValue ?? false;
        const isFieldPolymorphic = goField.type.kind === 'interface';
        exampleText += `${indent}\t${field}: ${getExampleValue(codeModel, example.value[field], indent + '\t', imports, isFieldByValue && !isFieldPolymorphic).slice(indent.length + 1)},\n`;
      }
      if (example.additionalProperties) {
        const additionalPropertiesField = example.type.fields.find(f => f.annotations.isAdditionalProperties)!;
        if (additionalPropertiesField.type.kind !== 'map') {
          throw new CodegenError('InternalError', `additional properties field type should be map type`);
        }
        const isAdditionalPropertiesFieldByValue = additionalPropertiesField.type.valueTypeByValue ?? false;
        const isAdditionalPropertiesPolymorphic = additionalPropertiesField.type.valueType.kind === 'interface';
        exampleText += `${indent}\t${additionalPropertiesField.name}: ${getRef(additionalPropertiesField.byValue)}${go.getTypeDeclaration(additionalPropertiesField.type, codeModel.packageName)}{\n`;
        for (const key in example.additionalProperties) {
          exampleText += `${indent}\t"${key}": ${getExampleValue(codeModel, example.additionalProperties[key], indent + '\t', imports, isAdditionalPropertiesFieldByValue && !isAdditionalPropertiesPolymorphic).slice(indent.length + 1)},\n`;
        }
        exampleText += `${indent}},\n`;
      }
      exampleText += `${indent}}`;
      return exampleText;
    }
    case 'tokenCredential':
      return example.value;
  }
}

function getRef(byValue: boolean): string {
  return byValue ? '' : '&';

}

function getConstantValue(codeModel: go.CodeModel, type: go.Constant, value: any): string {
  for (const constantValue of type.values) {
    if (constantValue.value === value) {
      return `${codeModel.packageName}.${constantValue.name}`
    }
  }
  switch (type.type) {
    case 'string':
      return `${codeModel.packageName}.${type.name}("${value}")`;
    default:
      return `${codeModel.packageName}.${type.name}(${value})`;
  }
}

function getTimeValue(type: go.Time, value: any, imports?: ImportManager): string {
  if (type.format === 'dateType' || type.format === 'timeRFC3339') {
    if (imports) imports.add('time');
    let format = helpers.dateFormat;
    if (type.format === 'timeRFC3339') {
      format = helpers.timeRFC3339Format;
    }
    return `func() time.Time { t, _ := time.Parse("${format}", "${value}"); return t}()`;
  } else if (type.format === 'dateTimeRFC1123' || type.format === 'dateTimeRFC3339') {
    if (imports) imports.add('time');
    let format = helpers.datetimeRFC3339Format;
    if (type.format === 'dateTimeRFC1123') {
      format = helpers.datetimeRFC1123Format;
    }
    return `func() time.Time { t, _ := time.Parse(${format}, "${value}"); return t}()`;
  } else {
    if (imports) imports.add('strconv');
    if (imports) imports.add('time');
    return `func() time.Time { t, _ := strconv.ParseInt(${value}, 10, 64); return time.Unix(t, 0).UTC()}()`;
  }
}

function getPointerValue(type: go.WireType, valueString: string, byValue: boolean, imports?: ImportManager): string {
  if (byValue) {
    return valueString;
  }

  switch (type.kind) {
    case 'any':
    case 'constant':
    case 'literal':
    case 'string':
    case 'time':
      if (imports) imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
      return `to.Ptr(${valueString})`;
    case 'scalar': {
      let prtType = '';
      switch (type.type) {
        case `bool`:
        case `byte`:
        case `rune`:
          prtType = 'Ptr';
          break;
        default:
          prtType = `Ptr[${type.type}]`;
      }
      if (imports) imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
      return `to.${prtType}(${valueString})`;
    }
    default:
      return `&${valueString}`;
  }
}

function jsonToGo(value: any, indent: string): string {
  if (typeof value === 'string') {
    return `${indent}"${value}"`;
  } else if (typeof value === 'number' || typeof value === 'bigint') {
    return `${indent}${value}`;
  } else if (typeof value === 'boolean') {
    return `${indent}${value}`;
  } else if (typeof value === 'undefined') {
    return `${indent}nil`;
  } else if (typeof value === 'object') {
    if (value === null) {
      return `${indent}nil`;
    } else if (Array.isArray(value)) {
      let result = `${indent}[]any{\n`;
      for (const item of value) {
        result += `${jsonToGo(item, indent + '\t')},\n`;
      }
      result += `${indent}}`;
      return result;
    } else {
      let result = `${indent}map[string]any{\n`;
      for (const key in value) {
        result += `${indent}\t"${key}": ${jsonToGo(value[key], indent + '\t').slice(indent.length + 1)},\n`;
      }
      result += `${indent}}`;
      return result;
    }
  }
  return '';
}

function generateFakeExample(goType: go.Type, name?: string): go.ExampleType {
  switch (goType.kind) {
    case 'any':
      return new go.NullExample(goType);
    case 'constant':
      switch (goType.type) {
        case 'bool':
          return new go.BooleanExample(goType.values[0].value as boolean, goType);
        case 'string':
          return new go.StringExample(goType.values[0].value as string, goType);
        default:
          return new go.NumberExample(goType.values[0].value as number, goType);
      }
    case 'literal':
      return new go.StringExample(goType.literal, goType);
    case 'scalar':
      switch (goType.type) {
        case 'bool':
          return new go.BooleanExample(false, goType);
        case 'byte':
        case 'rune':
          return new go.StringExample(`<${name ?? 'test'}>`, goType);
        default:
          return new go.NumberExample(0, goType);
      }
    case 'string':
      return new go.StringExample(`<${name ?? 'test'}>`, goType);
    case 'tokenCredential':
      // we hard code the credential var name to cred
      return new go.TokenCredentialExample('cred');
    default:
      throw new CodegenError('InternalError', `unhandled fake example kind ${goType.kind}`);
  }
}

function escapeString(str: string): string {
  return str.split('\\').join('\\\\').split('"').join('\\"').replace(/\n/g, '\\n').replace(/\r/g, '\\r');
}
