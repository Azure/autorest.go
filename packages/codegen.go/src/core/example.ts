/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../../codemodel.go/src/index.js';
import * as naming from '../../../naming.go/src/naming.js';
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

/**
 * Creates the content for all the *_example_test.go files.
 *
 * @param pkg contains the package content
 * @param target the codegen target for the module
 * @param options the emitter options
 * @returns the text for the files or the empty string
 */
export function generateExamples(pkg: go.TestPackage, target: go.CodeModelType, options: go.Options): Array<ExampleContent> {
  // generate examples
  const examples = new Array<ExampleContent>();
  if (pkg.src.clients.length === 0) {
    return examples;
  }

  const azureARM = target === 'azure-arm';

  for (const client of pkg.src.clients) {
    // client must be constructable to create a sample
    if (client.instance?.kind != 'constructable') {
      continue;
    }
    const imports = new ImportManager(pkg);
    // the list of packages to import
    if (client.methods.length > 0) {
      // add standard imports for clients with methods.
      // clients that are purely hierarchical (i.e. having no APIs) won't need them.
      imports.add('context');
      imports.add('log');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azidentity');
      imports.addForPkg(pkg.src);
    }

    let clientFactoryParams = new Array<go.ClientParameter>();
    if (options.factoryGatherAllParams) {
      clientFactoryParams = helpers.getAllClientParameters(pkg.src, target);
    } else {
      clientFactoryParams = helpers.getCommonClientParameters(pkg.src, target);
    }
    const clientFactoryParamsMap = new Map<string, go.ClientParameter>();
    for (const param of clientFactoryParams) {
      clientFactoryParamsMap.set(param.name, param);
    }

    let exampleText = '';
    for (const method of client.methods) {
      for (const example of method.examples) {
        const indent = new helpers.Indentation();
        // function signature
        exampleText += `// Generated from example definition: ${example.filePath}\n`;
        const exampleFuncNamePrefix = method.examples.length > 1 ? `_${helpers.camelCase(example.name)}` : '';
        exampleText += `func Example${client.name}_${fixUpMethodName(method)}${exampleFuncNamePrefix}() {\n`;

        // create credential
        exampleText += `${indent.get()}cred, err := azidentity.NewDefaultAzureCredential(nil)\n`;
        exampleText += `${indent.get()}if err != nil {\n`;
        exampleText += `${indent.push().get()}log.Fatalf("failed to obtain a credential: %v", err)\n`;
        exampleText += `${indent.pop().get()}}\n`;

        // create context
        exampleText += `${indent.get()}ctx := context.Background()\n`;

        // create client
        const clientParameters: go.ParameterExample[] = [];
        for (const param of method.parameters) {
          if (param.location === 'client') {
            if (go.isLiteralParameter(param.style)) {
              continue;
            }
            const clientParam = example.parameters.find((p) => p.parameter.name === param.name);
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
            const clientFactoryParam = clientParameters.find((p) => p.parameter.name === clientParam.name);
            if (clientFactoryParam) {
              clientFactoryParamsExample.push(clientFactoryParam);
            } else {
              clientFactoryParamsExample.push({ parameter: clientParam, value: generateFakeExample(clientParam.type, clientParam.name) });
            }
          }
          exampleText += `${indent.get()}clientFactory, err := ${go.getPackageName(pkg.src)}.NewClientFactory(${clientFactoryParamsExample.map((p) => getExampleValue(pkg, p.value, '\t', imports, p.parameter.byValue)).join(', ')}, nil)\n`;
          exampleText += `${indent.get()}if err != nil {\n`;
          exampleText += `${indent.push().get()}log.Fatalf("failed to create client: %v", err)\n`;
          exampleText += `${indent.pop().get()}}\n`;
          clientRef = `clientFactory.${client.instance.constructors[0].name}(`;
          const clientPrivateParameters: go.ParameterExample[] = [];
          for (const clientParam of clientParameters) {
            if (!clientFactoryParamsMap.has(clientParam.parameter.name)) {
              clientPrivateParameters.push(clientParam);
            }
          }
          if (clientPrivateParameters.length > 0) {
            clientRef += `${clientPrivateParameters.map((p) => getExampleValue(pkg, p.value, '\t', imports, p.parameter.byValue).slice(1)).join(', ')}`;
          }
          clientRef += `)`;
        } else {
          exampleText += `${indent.get()}client, err := ${go.getPackageName(client.instance.constructors[0].pkg)}.${client.instance.constructors[0].name}(${clientParameters.map((p) => getExampleValue(pkg, p.value, '\t', imports, p.parameter.byValue).slice(1)).join(', ')}, cred, nil)\n`;
          exampleText += `${indent.get()}if err != nil {\n`;
          exampleText += `${indent.push().get()}log.Fatalf("failed to create client: %v", err)\n`;
          exampleText += `${indent.pop().get()}}\n`;
          clientRef = 'client';
        }

        // call method
        const methodParameters: go.ParameterExample[] = [];
        for (const param of method.parameters) {
          if (param.location === 'method') {
            if (go.isLiteralParameter(param.style)) {
              continue;
            }
            const methodParam = example.parameters.find((p) => p.parameter.name === param.name);
            if (methodParam) {
              methodParameters.push(methodParam);
            } else if (go.isRequiredParameter(param.style)) {
              // if the parameter is required but lacks example value, generate a fake example
              methodParameters.push({ parameter: param, value: generateFakeExample(param.type, param.name) });
            }
          }
        }

        const methodOptionalParameters = example.optionalParamsGroup.filter((p) => p.parameter.location === 'method');
        const checkResponse = example.responseEnvelope !== undefined;

        let methodOptionalParametersText = 'nil';
        if (methodOptionalParameters.length > 0) {
          methodOptionalParametersText = `&${go.getPackageName(method.optionalParamsGroup.pkg)}.${method.optionalParamsGroup.groupName}{\n`;
          methodOptionalParametersText += methodOptionalParameters
            .map((p) => `${naming.capitalize(p.parameter.name)}: ${getExampleValue(pkg, p.value, '\t', imports, isParamByValue(p)).slice(1)}`)
            .join(',\n');
          methodOptionalParametersText += `}`;
        }

        switch (method.kind) {
          case 'lroMethod':
          case 'lroPageableMethod':
            exampleText += `${indent.get()}poller, err := ${clientRef}.${fixUpMethodName(method)}(ctx, ${methodParameters.map((p) => getExampleValue(pkg, p.value, '\t', imports, isParamByValue(p)).slice(1)).join(', ')}${methodParameters.length > 0 ? ', ' : ''}${methodOptionalParametersText.split('\n').join('\n' + indent.get())})\n`;
            exampleText += `${indent.get()}if err != nil {\n`;
            exampleText += `${indent.push().get()}log.Fatalf("failed to finish the request: %v", err)\n`;
            exampleText += `${indent.pop().get()}}\n`;

            exampleText += `${indent.get()}${checkResponse ? 'res' : '_'}, err ${checkResponse ? ':=' : '='} poller.PollUntilDone(ctx, nil)\n`;
            exampleText += `${indent.get()}if err != nil {\n`;
            exampleText += `${indent.push().get()}log.Fatalf("failed to poll the result: %v", err)\n`;
            exampleText += `${indent.pop().get()}}\n`;
            break;
          case 'method':
            exampleText += `${indent.get()}${checkResponse ? 'res' : '_'}, err ${checkResponse ? ':=' : '='} ${clientRef}.${fixUpMethodName(method)}(ctx, ${methodParameters.map((p) => getExampleValue(pkg, p.value, '\t', imports, isParamByValue(p)).slice(1)).join(', ')}${methodParameters.length > 0 ? ', ' : ''}${methodOptionalParametersText.split('\n').join('\n' + indent.get())})\n`;
            exampleText += `${indent.get()}if err != nil {\n`;
            exampleText += `${indent.push().get()}log.Fatalf("failed to finish the request: %v", err)\n`;
            exampleText += `${indent.pop().get()}}\n`;
            break;
          case 'pageableMethod':
            exampleText += `${indent.get()}pager := ${clientRef}.${fixUpMethodName(method)}(${methodParameters.map((p) => getExampleValue(pkg, p.value, '\t', imports, isParamByValue(p)).slice(1)).join(', ')}${methodParameters.length > 0 ? ', ' : ''}${methodOptionalParametersText.split('\n').join('\n' + indent.get())})\n`;
            break;
          default:
            method satisfies never;
        }

        // check response
        if ((method.kind === 'lroPageableMethod' || method.kind === 'pageableMethod') && example.responseEnvelope) {
          let resultName = 'pager';
          if (method.kind === 'lroPageableMethod') {
            resultName = 'res';
          }
          const itemType = ((method as go.PageableMethod).returns.result as go.ModelResult).modelType.fields.find((f) => f.type.kind === 'slice')!;
          exampleText += `${indent.get()}for ${resultName}.More() {\n`;
          exampleText += `${indent.push().get()}page, err := ${resultName}.NextPage(ctx)\n`;
          exampleText += `${indent.get()}if err != nil {\n`;
          exampleText += `${indent.push().get()}log.Fatalf("failed to advance page: %v", err)\n`;
          exampleText += `${indent.pop().get()}}\n`;
          exampleText += `${indent.get()}for _, v := range page.${itemType.name} {\n`;
          exampleText += `${indent.push().get()}// You could use page here. We use blank identifier for just demo purposes.\n`;
          exampleText += `${indent.get()}_ = v\n`;
          exampleText += `${indent.pop().get()}}\n`;
          exampleText += `${indent.get()}// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.\n`;
          exampleText += `${indent.get()}// page = ${go.getPackageName(example.responseEnvelope.response.method.receiver.type.pkg)}.${example.responseEnvelope.response.name}{\n`;
          for (const header of example.responseEnvelope.headers ?? []) {
            exampleText += `${indent.get()}// \t${header.header.fieldName}: ${getExampleValue(pkg, header.value, '', undefined, (header.header as any).byValue)
              .split('\n')
              .join(`\n${indent.get()}// \t`)},\n`;
          }
          exampleText += `${indent.get()}// \t${(example.responseEnvelope.result.type as go.Model).name}: ${getExampleValue(pkg, example.responseEnvelope.result!, '', undefined, true).split('\n').join(`\n${indent.get()}// \t`)},\n`;
          exampleText += `${indent.get()}// }\n`;
          exampleText += `${indent.pop().get()}}\n`;
        } else if (example.responseEnvelope) {
          // if has fieldName, then the result is not a model type
          const fieldName = (method.returns.result as any)?.fieldName;
          exampleText += `${indent.get()}// You could use response here. We use blank identifier for just demo purposes.\n`;
          exampleText += `${indent.get()}_ = res\n`;

          exampleText += `${indent.get()}// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.\n`;
          exampleText += `${indent.get()}// res = ${go.getPackageName(example.responseEnvelope.response.method.receiver.type.pkg)}.${example.responseEnvelope.response.name}{\n`;
          for (const header of example.responseEnvelope?.headers ?? []) {
            exampleText += `${indent.get()}// \t${header.header.fieldName}: ${getExampleValue(pkg, header.value, '', undefined, (header.header as any).byValue)
              .split('\n')
              .join(`\n${indent.get()}// \t`)},\n`;
          }
          if (example.responseEnvelope?.result) {
            // modelResult and polymorphicResult are anonymously embedded by value in the response struct.
            // monomorphicResult has an explicit byValue property. all other result types default to by value.
            let resultByValue = true;
            let resultFieldName = fieldName ? fieldName : (example.responseEnvelope?.result.type as go.Model).name;
            if (method.returns.result?.kind === 'monomorphicResult') {
              resultByValue = method.returns.result.byValue;
            } else if (method.returns.result?.kind === 'polymorphicResult') {
              resultFieldName = method.returns.result.interface.name;
              resultByValue = false;
            }
            exampleText += `${indent.get()}// \t${resultFieldName}: ${getExampleValue(pkg, example.responseEnvelope.result, '', undefined, resultByValue).split('\n').join(`\n${indent.get()}// \t`)},\n`;
          }
          exampleText += `${indent.get()}// }\n`;
        }
        exampleText += `}\n\n`;
      }
    }

    // if no example, then do not generate example file
    if (exampleText === '') continue;

    // stitch it all together
    let text = helpers.contentPreamble(pkg);
    text += imports.text();
    text += exampleText;
    examples.push(new ExampleContent(client.name, text));
  }
  return examples;
}

function getExampleValue(pkg: go.TestPackage, example: go.ExampleType, indent: string, imports?: ImportManager, byValue: boolean = false, inArray: boolean = false): string {
  switch (example.kind) {
    case 'string': {
      let exampleText = `"${escapeString(example.value)}"`;
      if (example.type.kind === 'constant') {
        exampleText = getConstantValue(pkg, example.type, example.value);
      } else if (example.type.kind === 'time') {
        exampleText = getTimeValue(example.type, example.value, imports);
      } else if (example.type.kind === 'encodedBytes') {
        exampleText = `[]byte("${escapeString(example.value)}")`;
      } else if (example.type.kind === 'literal' && example.type.type.kind === 'constant') {
        exampleText = getConstantValue(pkg, example.type.type, (<go.ConstantValue>example.type.literal).value);
      } else if (example.type.kind === 'etag') {
        imports?.add(example.type.module);
        exampleText = `${go.getTypeDeclaration(example.type, pkg)}("${escapeString(example.value)}")`;
      } else if (example.type.kind === 'scalar' && example.type.type === 'byte') {
        exampleText = `io.NopCloser(bytes.NewReader([]byte("${escapeString(example.value)}")))`;
      } else if (example.type.kind === 'readSeekCloser') {
        imports?.add('bytes');
        imports?.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
        exampleText = `streaming.NopCloser(bytes.NewReader([]byte("${escapeString(example.value)}")))`;
      }
      return `${indent}${getPointerValue(example.type, exampleText, byValue, imports)}`;
    }
    case 'number': {
      let exampleText = `${example.value}`;
      switch (example.type.kind) {
        case 'constant':
          exampleText = `${indent}${getConstantValue(pkg, example.type, example.value)}`;
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
        exampleText = `${indent}${getConstantValue(pkg, example.type, example.value)}`;
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
      let exampleText = `${indent}${getRef(byValue)}${go.getTypeDeclaration(example.type, pkg)}{\n`;
      for (const element of example.value) {
        exampleText += `${getExampleValue(pkg, element, indent + '\t', imports, isElementByValue && !isElementPolymorphic, !isElementPolymorphic)},\n`;
      }
      exampleText += `${indent}}`;
      return exampleText;
    }
    case 'dictionary': {
      let exampleText = `${indent}${getRef(byValue)}${go.getTypeDeclaration(example.type, pkg)}{\n`;
      const isValueByValue = example.type.valueTypeByValue;
      const isValuePolymorphic = example.type.valueType.kind === 'interface';
      for (const key in example.value) {
        exampleText += `${indent}\t"${key}": ${getExampleValue(pkg, example.value[key], indent + '\t', imports, isValueByValue && !isValuePolymorphic).slice(indent.length + 1)},\n`;
      }
      exampleText += `${indent}}`;
      return exampleText;
    }
    case 'model': {
      let exampleText = `${indent}${getRef(byValue)}${go.getTypeDeclaration(example.type, pkg)}{\n`;
      if (inArray) {
        exampleText = `${indent}{\n`;
      }
      for (const field in example.value) {
        const goField = example.type.fields.find((f) => f.name === field)!;
        const isFieldByValue = goField.byValue ?? false;
        const isFieldPolymorphic = goField.type.kind === 'interface';
        exampleText += `${indent}\t${field}: ${getExampleValue(pkg, example.value[field], indent + '\t', imports, isFieldByValue && !isFieldPolymorphic).slice(indent.length + 1)},\n`;
      }
      if (example.additionalProperties) {
        const additionalPropertiesField = example.type.fields.find((f) => f.annotations.isAdditionalProperties)!;
        if (additionalPropertiesField.type.kind !== 'map') {
          throw new CodegenError('InternalError', `additional properties field type should be map type`);
        }
        const isAdditionalPropertiesFieldByValue = additionalPropertiesField.type.valueTypeByValue ?? false;
        const isAdditionalPropertiesPolymorphic = additionalPropertiesField.type.valueType.kind === 'interface';
        exampleText += `${indent}\t${additionalPropertiesField.name}: ${getRef(additionalPropertiesField.byValue)}${go.getTypeDeclaration(additionalPropertiesField.type, pkg)}{\n`;
        for (const key in example.additionalProperties) {
          exampleText += `${indent}\t"${key}": ${getExampleValue(pkg, example.additionalProperties[key], indent + '\t', imports, isAdditionalPropertiesFieldByValue && !isAdditionalPropertiesPolymorphic).slice(indent.length + 1)},\n`;
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

function getConstantValue(pkg: go.TestPackage, type: go.Constant, value: any): string {
  for (const constantValue of type.values) {
    if (constantValue.value === value) {
      return go.getTypeDeclaration(constantValue, pkg);
    }
  }
  switch (type.type) {
    case 'string':
      return `${go.getTypeDeclaration(type, pkg)}("${value}")`;
    default:
      return `${go.getTypeDeclaration(type, pkg)}(${value})`;
  }
}

function getTimeValue(type: go.Time, value: any, imports?: ImportManager): string {
  const formatMap: Record<string, string> = {
    PlainDate: helpers.plainDateFormat,
    PlainTime: helpers.plainTimeFormat,
    RFC1123: helpers.RFC1123Format,
    RFC3339: helpers.RFC3339Format,
  };

  if (type.format in formatMap) {
    imports?.add('time');
    const format = formatMap[type.format];
    return `func() time.Time { t, _ := time.Parse(${format}, "${value}"); return t}()`;
  } else {
    imports?.add('strconv');
    imports?.add('time');
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
    case 'constantDef':
    case 'etag':
    case 'string':
    case 'time':
      if (imports) imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
      return `to.Ptr(${valueString})`;
    case 'literal':
      // unwrap the literal and delegate to the inner type's case
      return getPointerValue(type.type as go.WireType, valueString, byValue, imports);
    case 'scalar': {
      if (type.type === 'byte') {
        return valueString;
      }
      let prtType = '';
      switch (type.type) {
        case `bool`:
        case `rune`:
          prtType = 'Ptr';
          break;
        default:
          prtType = `Ptr[${type.type}]`;
      }
      if (imports) imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
      return `to.${prtType}(${valueString})`;
    }
    case 'readSeekCloser':
      return valueString;
    default:
      return `&${valueString}`;
  }
}

function jsonToGo(value: any, indent: string): string {
  if (typeof value === 'string') {
    return `${indent}"${escapeString(value)}"`;
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
      return new go.StringExample(<string>goType.literal, goType);
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
    case 'encodedBytes':
      return new go.StringExample(`<${name ?? 'test'}>`, goType);
    case 'time':
      // use a placeholder date value for time types
      return new go.StringExample('2006-01-02T15:04:05Z', goType);
    case 'etag':
      return new go.StringExample(`<${name ?? 'etag'}>`, goType);
    case 'model':
    case 'polymorphicModel':
      // return an empty struct example for model types
      return new go.StructExample(goType);
    case 'slice':
      // return an empty array example for slice types
      return new go.ArrayExample(goType);
    case 'map':
      // return an empty map example for map types
      return new go.DictionaryExample(goType);
    case 'interface':
      // for interface types, use the root type (which is a PolymorphicModel) to create an example
      return new go.StructExample(goType.rootType);
    default:
      throw new CodegenError('InternalError', `unhandled fake example kind ${goType.kind}`);
  }
}

function escapeString(str: string): string {
  return str.split('\\').join('\\\\').split('"').join('\\"').replace(/\n/g, '\\n').replace(/\r/g, '\\r');
}

// interface parameters require a pointer (&) for Go interface satisfaction via pointer receivers,
// even when the parameter style is byValue. This helper computes the effective byValue for example generation.
function isParamByValue(p: go.ParameterExample): boolean {
  return p.parameter.byValue && p.parameter.type.kind !== 'interface';
}
