/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { camelCase } from '@azure-tools/codegen';
import * as go from '../../codemodel.go/src/index.js';
import { getAllClientParameters } from './clientFactory.js';
import * as helpers from './helpers.js';
import { ImportManager } from './imports.js';
import { fixUpMethodName } from './operations.js';

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
    const imports = new ImportManager();
    // the list of packages to import
    if (client.methods.length > 0) {
      // add standard imports for clients with methods.
      // clients that are purely hierarchical (i.e. having no APIs) won't need them.
      imports.add('context');
      imports.add('log');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azidentity');
      imports.add(codeModel.options.module!.name);
    }

    const allClientParams = getAllClientParameters(codeModel);

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
          const clientFactoryParams: go.ParameterExample[] = [];
          for (const clientParam of allClientParams) {
            const clientFactoryParam = clientParameters.find(p => p.parameter.name === clientParam.name);
            if (clientFactoryParam) {
              clientFactoryParams.push(clientFactoryParam);
            } else {
              clientFactoryParams.push({ parameter: clientParam, value: generateFakeExample(clientParam.type, clientParam.name) });
            }
          }
          exampleText += `\tclientFactory, err := ${codeModel.packageName}.NewClientFactory(${clientFactoryParams.map(p => getExampleValue(codeModel, p.value, '\t', imports, helpers.parameterByValue(p.parameter)).slice(1)).join(', ')}, cred, nil)\n`;
          exampleText += `\tif err != nil {\n`;
          exampleText += `\t\tlog.Fatalf("failed to create client: %v", err)\n`;
          exampleText += `\t}\n`;
          clientRef = `clientFactory.${client.constructors[0]?.name}()`;
        } else {
          exampleText += `\tclient, err := ${codeModel.packageName}.${client.constructors[0]?.name}(${clientParameters.map(p => getExampleValue(codeModel, p.value, '\t', imports, helpers.parameterByValue(p.parameter)).slice(1)).join(', ')}, cred, nil)\n`;
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
          methodOptionalParametersText = `&${method.optionalParamsGroup.groupName}{\n`;
          methodOptionalParametersText += methodOptionalParameters.map(p => `${p.parameter.name}: ${getExampleValue(codeModel, p.value, '\t', imports, helpers.parameterByValue(p.parameter)).slice(1)}`).join(',\n');
          methodOptionalParametersText += `}`;
        }

        if (go.isLROMethod(method)) {
          exampleText += `\tpoller, err := ${clientRef}.${fixUpMethodName(method)}(ctx, ${methodParameters.map(p => getExampleValue(codeModel, p.value, '\t', imports, helpers.parameterByValue(p.parameter)).slice(1)).join(', ')}, ${methodOptionalParametersText.split('\n').join('\n\t')})\n`
          exampleText += `\tif err != nil {\n`;
          exampleText += `\t\tlog.Fatalf("failed to finish the request: %v", err)\n`;
          exampleText += `\t}\n`;

          exampleText += `\t${checkResponse ? 'res' : '_'}, err ${checkResponse ? ':=' : '='} poller.PollUntilDone(ctx, nil)\n`
          exampleText += `\tif err != nil {\n`;
          exampleText += `\t\tlog.Fatalf("failed to pull the result: %v", err)\n`;
          exampleText += `\t}\n`;
        } else if (go.isPageableMethod(method)) {
          exampleText += `\tpager := ${clientRef}.${fixUpMethodName(method)}(${methodParameters.map(p => getExampleValue(codeModel, p.value, '\t', imports, helpers.parameterByValue(p.parameter)).slice(1)).join(', ')}${methodParameters.length > 0 ? ', ' : ''}${methodOptionalParametersText.split('\n').join('\n\t')})\n`
        } else {
          exampleText += `\t${checkResponse ? 'res' : '_'}, err ${checkResponse ? ':=' : '='} ${clientRef}.${fixUpMethodName(method)}(ctx, ${methodParameters.map(p => getExampleValue(codeModel, p.value, '\t', imports, helpers.parameterByValue(p.parameter)).slice(1)).join(', ')}, ${methodOptionalParametersText.split('\n').join('\n\t')})\n`
          exampleText += `\tif err != nil {\n`;
          exampleText += `\t\tlog.Fatalf("failed to finish the request: %v", err)\n`;
          exampleText += `\t}\n`;
        }

        // check response
        if (go.isPageableMethod(method)) {
          let resultName = 'pager';
          if (go.isLROMethod(method)) {
            resultName = 'res';
          }
          const itemType = ((method as go.PageableMethod).responseEnvelope.result as go.ModelResult).modelType.fields.find(f => go.isSliceType(f.type))!;
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
          exampleText += `\t\t// \t${(example.responseEnvelope?.result.type as go.ModelType).name}: ${getExampleValue(codeModel, example.responseEnvelope?.result!, '', undefined, true).split('\n').join('\n\t\t// \t')},\n`;
          exampleText += '\t\t// }\n';
          exampleText += `\t}\n`;
        } else if (checkResponse) {
          // if has fieldName, then the result is not a model type
          const fieldName = (method.responseEnvelope as any).fieldName;
          exampleText += `\t// You could use response here. We use blank identifier for just demo purposes.\n`;
          exampleText += `\t_ = res\n`;

          exampleText += `\t// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.\n`;
          exampleText += `\t// res = ${codeModel.packageName}.${example.responseEnvelope?.response.name}{\n`;
          for (const header of example.responseEnvelope?.headers ?? []) {
            exampleText += `\t// \t${header.header.fieldName}: ${getExampleValue(codeModel, header.value, '', undefined, true).split('\n').join('\n\t// \t')}\n`;
          }
          if (example.responseEnvelope?.result) {
            exampleText += `\t// \t${fieldName ? fieldName : (example.responseEnvelope?.result.type as go.ModelType).name}: ${getExampleValue(codeModel, example.responseEnvelope.result, '').split('\n').join('\n\t// \t')},\n`;
          }
          exampleText += '\t// }\n';
        }
        exampleText += `}\n\n`;
      }
    }

    // if no example, then do not generate example file
    if (exampleText === '') continue;

    // stitch it all together
    let text = helpers.contentPreamble(codeModel, codeModel.packageName + '_test');
    text += imports.text();
    text += exampleText;
    examples.push(new ExampleContent(client.name, text));
  }
  return examples;
}

function getExampleValue(codeModel: go.CodeModel, example: go.ExampleType, indent: string, imports?: ImportManager, byValue: boolean = false, inArray: boolean = false): string {
  if (example.kind === 'string') {
    let exampleText = `"${escapeString(example.value)}"`;
    if (go.isConstantType(example.type)) {
      exampleText = getConstantValue(codeModel, example.type, example.value);
    } else if (go.isTimeType(example.type)) {
      exampleText = getTimeValue(example.type, example.value, imports);
    } else if (go.isBytesType(example.type)) {
      exampleText = `[]byte("${escapeString(example.value)}")`
    }
    return `${indent}${getPointerValue(example.type, exampleText, byValue, imports)}`;
  } else if (example.kind === 'number') {
    let exampleText = `${example.value}`;
    if (go.isConstantType(example.type)) {
      exampleText = `${indent}${getConstantValue(codeModel, example.type, example.value)}`;
    } else if (go.isTimeType(example.type)) {
      exampleText = getTimeValue(example.type, example.value, imports);
    }
    return `${indent}${getPointerValue(example.type, exampleText, byValue, imports)}`;
  } else if (example.kind === 'boolean') {
    let exampleText = `${example.value}`;
    if (go.isConstantType(example.type)) {
      exampleText = `${indent}${getConstantValue(codeModel, example.type, example.value)}`;
    }
    return `${indent}${getPointerValue(example.type, exampleText, byValue, imports)}`;
  } else if (example.kind === 'null') {
    return `${indent}nil`;
  } else if (example.kind === 'any') {
    return jsonToGo(example.value, indent);
  } else if (example.kind === 'array') {
    const isElementByValue = example.type.elementTypeByValue;
    // if polymorphic, need to add type name in array, so inArray will be set to false
    // if other case, no need to add type name in array, so inArray will be set to true
    const isElementPolymorphic = go.isInterfaceType(example.type.elementType);
    let exampleText = `${indent}${getRef(byValue)}${go.getTypeDeclaration(example.type, codeModel.packageName)}{\n`;
    for (const element of example.value) {
      exampleText += `${getExampleValue(codeModel, element, indent + '\t', imports, isElementByValue && !isElementPolymorphic, !isElementPolymorphic)},\n`;
    }
    exampleText += `${indent}}`;
    return exampleText;
  } else if (example.kind === 'dictionary') {
    let exampleText = `${indent}${getRef(byValue)}${go.getTypeDeclaration(example.type, codeModel.packageName)}{\n`;
    const isValueByValue = example.type.valueTypeByValue;
    const isValuePolymorphic = go.isInterfaceType(example.type.valueType);
    for (const key in example.value) {
      exampleText += `${indent}\t"${key}": ${getExampleValue(codeModel, example.value[key], indent + '\t', imports, isValueByValue && !isValuePolymorphic).slice(indent.length + 1)},\n`;
    }
    exampleText += `${indent}}`;
    return exampleText;
  } else if (example.kind === 'model') {
    let exampleText = `${indent}${getRef(byValue)}${go.getTypeDeclaration(example.type, codeModel.packageName)}{\n`;
    if (inArray) {
      exampleText = `${indent}{\n`;
    }
    for (const field in example.value) {
      const goField = example.type.fields.find(f => f.name === field)!;
      const isFieldByValue = goField.byValue ?? false;
      const isFieldPolymorphic = go.isInterfaceType(goField.type);
      exampleText += `${indent}\t${field}: ${getExampleValue(codeModel, example.value[field], indent + '\t', imports, isFieldByValue && !isFieldPolymorphic).slice(indent.length + 1)},\n`;
    }
    if (example.additionalProperties) {
      const additionalPropertiesField = example.type.fields.find(f => f.annotations.isAdditionalProperties)!;
      const isAdditionalPropertiesFieldByValue = additionalPropertiesField.byValue ?? false;
      const isAdditionalPropertiesPolymorphic = go.isInterfaceType(additionalPropertiesField.type);
      exampleText += `${indent}\t${additionalPropertiesField.name}: map[string]${getRef(additionalPropertiesField.byValue)}${go.getTypeDeclaration(additionalPropertiesField.type, codeModel.packageName)}{\n`;
      for (const key in example.additionalProperties) {
        exampleText += `${indent}\t"${key}": ${getExampleValue(codeModel, example.additionalProperties[key], indent + '\t', imports, isAdditionalPropertiesFieldByValue && !isAdditionalPropertiesPolymorphic).slice(indent.length + 1)},\n`;
      }
      exampleText += `${indent}}\n`;
    }
    exampleText += `${indent}}`;
    return exampleText;
  }
  return '';
}

function getRef(byValue: boolean): string {
  return byValue ? '' : '&';

}

function getConstantValue(codeModel: go.CodeModel, type: go.ConstantType, value: any): string {
  for (const constantValue of type.values) {
    if (constantValue.value === value) {
      return `${codeModel.packageName}.${constantValue.name}`
    }
  }
  return `${value}`;
}

function getTimeValue(type: go.TimeType, value: any, imports?: ImportManager): string {
  if (type.dateTimeFormat === 'dateType' || type.dateTimeFormat === 'timeRFC3339') {
    if (imports) imports.add('time');
    let format = helpers.dateFormat;
    if (type.dateTimeFormat === 'timeRFC3339') {
      format = helpers.timeRFC3339Format;
    }
    return `func() time.Time { t, _ := time.Parse("${format}", "${value}"); return t}()`;
  } else if (type.dateTimeFormat === 'dateTimeRFC1123' || type.dateTimeFormat === 'dateTimeRFC3339') {
    if (imports) imports.add('time');
    let format = helpers.datetimeRFC3339Format;
    if (type.dateTimeFormat === 'dateTimeRFC1123') {
      format = helpers.datetimeRFC1123Format;
    }
    return `func() time.Time { t, _ := time.Parse(${format}, "${value}"); return t}()`;
  } else {
    if (imports) imports.add('strconv');
    if (imports) imports.add('time');
    return `func() time.Time { t, _ := strconv.ParseInt(${value}, 10, 64); return time.Unix(t, 0).UTC()}()`;
  }
}

function getPointerValue(type: go.PossibleType, valueString: string, byValue: boolean, imports?: ImportManager): string {
  if (byValue) {
    return valueString;
  }

  if (go.isPrimitiveType(type)) {
    let prtType = '';
    switch (type.typeName) {
      case 'any':
      case `bool`:
      case `byte`:
      case `rune`:
      case `string`:
        prtType = 'Ptr';
        break;
      default:
        prtType = `Ptr[${type.typeName}]`;
    }
    if (imports) imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
    return `to.${prtType}(${valueString})`;
  } else if (go.isConstantType(type) || go.isTimeType(type)) {
    if (imports) imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
    return `to.Ptr(${valueString})`;
  } else if (go.isLiteralValue(type)) {
    if (imports) imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
    return `to.Ptr(${valueString})`;
  } else {
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
      let result = `${indent}any{\n`;
      for (const item of value) {
        result += `${jsonToGo(item, indent + '\t')},\n`;
      }
      result += `${indent}}`;
      return result;
    } else {
      let result = `map[string]any{\n`;
      for (const key in value) {
        result += `${indent}\t"${key}": ${jsonToGo(value[key], indent + '\t').slice(indent.length + 1)},\n`;
      }
      result += `${indent}}`;
      return result;
    }
  }
  return '';
}

function generateFakeExample(goType: go.PossibleType, name?: string): go.ExampleType {
  if (go.isPrimitiveType(goType)) {
    switch (goType.typeName) {
      case 'any':
        return new go.NullExample(goType);
      case 'bool':
        return new go.BooleanExample(false, goType);
      case 'byte':
      case 'string':
      case 'rune':
        return new go.StringExample(`<${name ?? 'test'}>`, goType);
      default:
        return new go.NumberExample(0, goType);
    }
  } else if (go.isLiteralValue(goType)) {
    return new go.StringExample(goType.literal, goType);
  } else if (go.isConstantType(goType)) {
    switch (goType.type) {
      case 'bool':
        return new go.BooleanExample(goType.values[0].value as boolean, goType);
      case 'string':
        return new go.StringExample(goType.values[0].value as string, goType);
      default:
        return new go.NumberExample(goType.values[0].value as number, goType);
    }
  }
  throw new Error(`do not support to fake example for none primitive type: ${goType}`);
}

function escapeString(str: string): string {
  return str.split('\\').join('\\\\').split('"').join('\\"').replace(/\n/g, '\\n').replace(/\r/g, '\\r');
}