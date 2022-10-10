/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import {
  ArraySchema,
  ChoiceSchema,
  DateTimeSchema,
  DictionarySchema,
  GroupProperty,
  ImplementationLocation,
  Metadata,
  ObjectSchema,
  Parameter,
  Schema,
  SchemaType,
} from '@autorest/codemodel';
import { BaseCodeGenerator, BaseDataRender } from './baseGenerator';
import { Config } from '../common/constant';
import { ExampleParameter, ExampleValue } from '@autorest/testmodeler/dist/src/core/model';
import { GoExampleModel, GoMockTestDefinitionModel, ParameterOutput } from '../common/model';
import { GoHelper } from '../util/goHelper';
import { Helper } from '@autorest/testmodeler/dist/src/util/helper';
import { elementByValueForParam } from '@autorest/go/dist/generator/helpers';
import { generateReturnsInfo, getAPIParametersSig, getClientParametersSig, getSchemaResponse } from '../util/codegenBridge';
import { isLROOperation, isMultiRespOperation, isPageableOperation } from '@autorest/go/dist/common/helpers';
import _ = require('lodash');
export class MockTestDataRender extends BaseDataRender {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  public skipPropertyFunc = (exampleValue: ExampleValue): boolean => {
    // skip any null value
    if (exampleValue.rawValue === null) {
      return true;
    }
    return false;
  };
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  public replaceValueFunc = (rawValue: any, exampleValue: ExampleValue): any => {
    return rawValue;
  };

  public renderData(): void {
    const mockTest = <GoMockTestDefinitionModel>this.context.codeModel.testModel.mockTest;
    for (const exampleGroup of mockTest.exampleGroups) {
      for (const example of <Array<GoExampleModel>>exampleGroup.examples) {
        this.fillExampleOutput(example);
      }
    }
  }

  protected fillExampleOutput(example: GoExampleModel): void {
    const op = example.operation;
    example.opName = op.language.go.name;
    if (isPageableOperation(op) && !isLROOperation(op)) {
      example.opName = `New${example.opName}Pager`;
    }
    if (isLROOperation(<any>op)) {
      example.opName = 'Begin' + example.opName;
      example.isLRO = true;
      example.pollerType = example.operation.language.go.responseEnv.language.go.name;
    } else {
      example.isLRO = false;
    }
    example.isPageable = isPageableOperation(<any>op);
    this.skipPropertyFunc = (exampleValue: ExampleValue): boolean => {
      // skip any null value
      if (exampleValue.rawValue === null) {
        return true;
      }
      return false;
    };
    this.replaceValueFunc = (rawValue: any): any => {
      return rawValue;
    };
    example.methodParametersOutput = this.toParametersOutput(getAPIParametersSig(op), example.methodParameters);
    example.clientParametersOutput = this.toParametersOutput(getClientParametersSig(example.operationGroup), example.clientParameters, true);
    example.returnInfo = generateReturnsInfo(op, 'op');
    const schemaResponse = getSchemaResponse(<any>op);
    if (example.isPageable) {
      const valueName = op.extensions['x-ms-pageable'].itemName === undefined ? 'value' : op.extensions['x-ms-pageable'].itemName;
      for (const property of schemaResponse.schema['properties']) {
        if (property.serializedName === valueName) {
          example.pageableItemName = property.language.go.name;
          break;
        }
      }
    }

    example.checkResponse =
      schemaResponse !== undefined && schemaResponse.protocol.http.statusCodes[0] === '200' && example.responses[schemaResponse.protocol.http.statusCodes[0]]?.body !== undefined;
    example.isMultiRespOperation = isMultiRespOperation(op);
    if (example.checkResponse && this.context.testConfig.getValue(Config.verifyResponse)) {
      this.context.importManager.add('encoding/json');
      this.context.importManager.add('reflect');
      this.skipPropertyFunc = (exampleValue: ExampleValue): boolean => {
        // mock-test will remove all NextLink param
        // skip any null value
        if (exampleValue.language?.go?.name === 'NextLink' || (exampleValue.rawValue === null && exampleValue.language?.go?.name !== 'ProvisioningState')) {
          return true;
        }
        return false;
      };
      this.replaceValueFunc = (rawValue: any, exampleValue: ExampleValue): any => {
        // mock-test will change all ProvisioningState to Succeeded
        if (exampleValue.language?.go?.name === 'ProvisioningState') {
          if (exampleValue.schema.type !== SchemaType.SealedChoice || (<ChoiceSchema>exampleValue.schema).choices.filter((choice) => choice.value === 'Succeeded').length > 0) {
            return 'Succeeded';
          } else {
            return (<ChoiceSchema>exampleValue.schema).choices[0].value;
          }
        }
        return rawValue;
      };
      example.responseOutput = this.exampleValueToString(example.responses[schemaResponse.protocol.http.statusCodes[0]].body, false);
      if (isMultiRespOperation(op)) {
        example.responseTypePointer = false;
        example.responseType = 'Value';
      } else {
        const responseEnv = op.language.go.responseEnv;

        if (responseEnv.language.go?.resultProp.schema.serialization?.xml?.name) {
          example.responseTypePointer = !responseEnv.language.go?.resultProp.schema.language.go?.byValue;
          example.responseType = responseEnv.language.go?.resultProp.schema.language.go?.name;
          if (responseEnv.language.go?.resultProp.schema.isDiscriminator === true) {
            example.responseIsDiscriminator = true;
            example.responseType = responseEnv.language.go.resultProp.schema.language.go?.discriminatorInterface;
            example.responseOutput = `${this.context.packageName}.${responseEnv.language.go.name}{
                            ${example.responseType}: &${example.responseOutput},
                        }`;
          }
        } else {
          example.responseTypePointer = !responseEnv.language.go?.resultProp.language.go?.byValue;
          example.responseType = responseEnv.language.go?.resultProp.language.go?.name;
          if (responseEnv.language.go?.resultProp.isDiscriminator === true) {
            example.responseIsDiscriminator = true;
            example.responseType = responseEnv.language.go.resultProp.schema.language.go?.discriminatorInterface;
            example.responseOutput = `${this.context.packageName}.${responseEnv.language.go.name}{
                            ${example.responseType}: &${example.responseOutput},
                        }`;
          }
        }
      }
    }
  }

  // get GO code of all parameters for one operation invoke
  protected toParametersOutput(
    paramsSig: Array<[string, string, Parameter | GroupProperty]>,
    exampleParameters: Array<ExampleParameter>,
    isClient = false,
  ): Array<ParameterOutput> {
    return paramsSig.map(([paramName, typeName, parameter]) => {
      if (paramName === 'ctx') {
        return new ParameterOutput('ctx', 'ctx');
      }
      return new ParameterOutput(paramName, this.genParameterOutput(paramName, typeName, parameter, exampleParameters, isClient));
    });
  }

  // get GO code of single parameter for one operation invoke
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  protected genParameterOutput(paramName: string, paramType: string, parameter: Parameter | GroupProperty, exampleParameters: Array<ExampleParameter>, isClient = false): string {
    // get corresponding example value of a parameter
    const findExampleParameter = (name: string, param: Parameter): string => {
      // isPtr need to consider three situation: 1) param is required 2) param is polymorphism 3) param is byValue
      const isPolymophismValue = param?.schema?.type === SchemaType.Object && (<ObjectSchema>param.schema).discriminator?.property.isDiscriminator === true;
      const isPtr: boolean = isPolymophismValue || !(param.required || param.language.go.byValue === true);
      for (const methodParameter of exampleParameters) {
        if (this.getLanguageName(methodParameter.parameter) === name) {
          // we should judge whether a param or property is ptr or not from outside of exampleValueToString
          return this.exampleValueToString(methodParameter.exampleValue, isPtr, elementByValueForParam(param));
        }
      }

      return this.getDefaultValue(param, isPtr, elementByValueForParam(param));
    };

    if ((<GroupProperty>parameter).originalParameter) {
      const group = <GroupProperty>parameter;
      const ptr = paramType.startsWith('*') ? '&' : '';
      let ret = `${ptr}${this.context.packageName + '.'}${this.getLanguageName(parameter.schema)}{`;
      let hasContent = false;
      for (const insideParameter of group.originalParameter) {
        if (insideParameter.implementation === ImplementationLocation.Client) {
          // don't add globals to the per-method options struct
          continue;
        }
        if (this.getLanguageName(insideParameter) === 'ResumeToken') {
          // ignore resumeToken param in options
          continue;
        }
        const insideOutput = findExampleParameter(this.getLanguageName(insideParameter), insideParameter);
        if (insideOutput) {
          ret += `${this.getLanguageName(insideParameter)}: ${insideOutput},\n`;
          hasContent = true;
        }
      }
      ret += '}';
      if (ptr.length > 0 && !hasContent) {
        ret = 'nil';
      }
      return ret;
    }
    return findExampleParameter(paramName, parameter);
  }

  protected getDefaultValue(param: Parameter | ExampleValue, isPtr: boolean, elemByVal = false): string {
    if (isPtr) {
      return 'nil';
    } else {
      switch (param.schema.type) {
        case SchemaType.Char:
        case SchemaType.String:
        case SchemaType.Constant:
          return '"<' + Helper.toKebabCase(this.getLanguageName(param)) + '>"';
        case SchemaType.Array: {
          const elementIsPtr = param.schema.language.go.elementIsPtr && !elemByVal;
          const elementPtr = elementIsPtr ? '*' : '';
          let elementTypeName = this.getLanguageName((<ArraySchema>param.schema).elementType);
          const polymophismName = (<ArraySchema>param.schema).elementType.language.go.discriminatorInterface;
          if (polymophismName) {
            elementTypeName = polymophismName;
          }
          return `[]${elementPtr}${GoHelper.addPackage(elementTypeName, this.context.packageName)}{}`;
        }
        case SchemaType.Dictionary: {
          const elementPtr = param.schema.language.go.elementIsPtr ? '*' : '';
          const elementTypeName = this.getLanguageName((<DictionarySchema>param.schema).elementType);
          return `map[string]${elementPtr}${GoHelper.addPackage(elementTypeName, this.context.packageName)}{}`;
        }
        case SchemaType.Boolean:
          return 'false';
        case SchemaType.Integer:
        case SchemaType.Number:
          return '0';
        case SchemaType.Object:
          if (isPtr) {
            return `&${this.context.packageName + '.'}${this.getLanguageName(param.schema)}{}`;
          } else {
            return `${this.context.packageName + '.'}${this.getLanguageName(param.schema)}{}`;
          }
        case SchemaType.AnyObject:
          return 'nil';
        case SchemaType.Any:
          return 'nil';
        default:
          return '';
      }
    }
  }

  protected exampleValueToString(exampleValue: ExampleValue, isPtr: boolean, elemByVal = false, inArray = false): string {
    if (exampleValue === null || exampleValue === undefined || exampleValue.isNull) {
      return 'nil';
    }
    const ptr = isPtr ? '&' : '';
    if (exampleValue.schema?.type === SchemaType.Array) {
      const elementIsPtr = exampleValue.schema.language.go.elementIsPtr && !elemByVal;
      const elementPtr = elementIsPtr ? '*' : '';
      const schema = <ArraySchema>exampleValue.schema;
      const elementIsPolymophism = schema.elementType.language.go.discriminatorInterface !== undefined;
      let elementTypeName = this.getLanguageName(schema.elementType);
      if (elementIsPolymophism) {
        elementTypeName = schema.elementType.language.go.discriminatorInterface;
      }
      if (exampleValue.elements === undefined) {
        const result = `${ptr}[]${elementPtr}${GoHelper.addPackage(elementTypeName, this.context.packageName)}{}`;
        return result;
      } else {
        // for polymorphism element, need to add type name, so pass false for inArray
        const result =
          `${ptr}[]${elementPtr}${GoHelper.addPackage(elementTypeName, this.context.packageName)}{\n` +
          exampleValue.elements.map((x) => this.exampleValueToString(x, elementIsPolymophism || elementIsPtr, false, elementIsPolymophism ? false : true)).join(',\n') +
          '}';
        return result;
      }
    } else if (exampleValue.schema?.type === SchemaType.Object) {
      if (exampleValue.rawValue === null) {
        return 'nil';
      }
      let output: string;
      if (inArray) {
        output = '{\n';
      } else {
        output = `${ptr}${this.context.packageName + '.'}${this.getLanguageName(exampleValue.schema)}{\n`;
      }

      // object parents' properties will be aggregated to the child
      const parentsProps: Array<ExampleValue> = [];
      const additionalProps: Array<ExampleValue> = [];
      this.aggregateParentsProps(exampleValue.parentsValue, parentsProps, additionalProps);
      for (const parentsProp of parentsProps) {
        const isPolymophismValue =
          parentsProp?.schema?.type === SchemaType.Object &&
          ((<ObjectSchema>parentsProp.schema).discriminatorValue !== undefined || (<ObjectSchema>parentsProp.schema).discriminator?.property.isDiscriminator === true);
        output += `${this.getLanguageName(parentsProp)}: ${this.exampleValueToString(parentsProp, isPolymophismValue || !parentsProp.language.go?.byValue === true)},\n`;
      }
      // TODO: handle multiple additionalProps
      for (const additionalProp of additionalProps) {
        output += `AdditionalProperties: ${this.exampleValueToString(additionalProp, false)},\n`;
      }
      for (const [_, value] of Object.entries(exampleValue.properties || {})) {
        if (this.skipPropertyFunc(value)) {
          continue;
        }
        const isPolymophismValue =
          value?.schema?.type === SchemaType.Object &&
          ((<ObjectSchema>value.schema).discriminatorValue !== undefined || (<ObjectSchema>value.schema).discriminator?.property.isDiscriminator === true);
        output += `${this.getLanguageName(value)}: ${this.exampleValueToString(value, isPolymophismValue || !value.language.go?.byValue === true)},\n`;
      }
      output += '}';
      return output;
    } else if (exampleValue.schema?.type === SchemaType.Dictionary) {
      const elementPtr = exampleValue.schema.language.go.elementIsPtr && !elemByVal ? '*' : '';
      const elementIsPolymorphism = (<DictionarySchema>exampleValue.schema).elementType.language.go.discriminatorInterface !== undefined;
      let elementTypeName = this.getLanguageName((<DictionarySchema>exampleValue.schema).elementType);
      if (elementIsPolymorphism) {
        elementTypeName = (<DictionarySchema>exampleValue.schema).elementType.language.go.discriminatorInterface;
      }
      let output = `${ptr}map[string]${elementPtr}${GoHelper.addPackage(elementTypeName, this.context.packageName)}{\n`;
      for (const [key, value] of Object.entries(exampleValue.properties || {})) {
        // for polymorphism map value, value should be a pointer
        output += `${this.getStringValue(key)}: ${this.exampleValueToString(value, exampleValue.schema.language.go.elementIsPtr || elementIsPolymorphism)},\n`;
      }
      output += '}';
      return output;
    }

    const rawValue = this.getRawValue(exampleValue);
    if (rawValue === null) {
      return this.getDefaultValue(exampleValue, isPtr);
    }
    return this.rawValueToString(rawValue, exampleValue.schema, isPtr);
  }

  protected aggregateParentsProps(exampleValues: Record<string, ExampleValue>, parentsProps: Array<ExampleValue>, additionalProps: Array<ExampleValue>): void {
    for (const [_, value] of Object.entries(exampleValues || {})) {
      if (value.schema?.type === SchemaType.Object) {
        this.aggregateParentsProps(value.parentsValue, parentsProps, additionalProps);
        for (const [_, property] of Object.entries(value.properties)) {
          if (this.skipPropertyFunc(property)) {
            continue;
          }
          if (
            parentsProps.filter((p) => {
              return p.language.go.name === property.language.go.name;
            }).length > 0
          ) {
            continue;
          }
          parentsProps.push(property);
        }
      } else if (value.schema?.type === SchemaType.Dictionary) {
        additionalProps.push(value);
      } else {
        parentsProps.push(value);
      }
    }
  }

  protected getRawValue(exampleValue: ExampleValue): void {
    exampleValue.rawValue = this.replaceValueFunc(exampleValue.rawValue, exampleValue);
    return exampleValue.rawValue;
  }

  protected getStringValue(rawValue: string): string {
    return Helper.quotedEscapeString(rawValue);
  }

  protected rawValueToString(rawValue: any, schema: Schema, isPtr: boolean): string {
    let ret = JSON.stringify(rawValue);
    const goType = this.getLanguageName(schema);
    if (schema.type === SchemaType.Choice) {
      if ((<ChoiceSchema>schema).choiceType.type === SchemaType.String) {
        try {
          const choiceValue = Helper.findChoiceValue(<ChoiceSchema>schema, rawValue);
          ret = this.context.packageName + '.' + this.getLanguageName(choiceValue);
        } catch (error) {
          ret = `${this.context.packageName}.${this.getLanguageName(schema)}("${rawValue}")`;
        }
      } else {
        ret = `${this.context.packageName}.${this.getLanguageName(schema)}(${rawValue})`;
      }
    } else if (schema.type === SchemaType.SealedChoice) {
      const choiceValue = Helper.findChoiceValue(<ChoiceSchema>schema, rawValue);
      ret = this.context.packageName + '.' + this.getLanguageName(choiceValue);
    } else if (goType === 'string') {
      ret = this.getStringValue(rawValue);
    } else if (schema.type === SchemaType.ByteArray) {
      ret = `[]byte(${this.getStringValue(rawValue)})`;
    } else if (['int32', 'int64', 'float32', 'float64'].indexOf(goType) >= 0) {
      ret = `${Number(rawValue)}`;
    } else if (goType === 'time.Time') {
      if (schema.type === SchemaType.UnixTime) {
        this.context.importManager.add('time');
        ret = `time.Unix(${rawValue}, 0)`;
      } else if (schema.type === SchemaType.Date) {
        this.context.importManager.add('time');
        ret = `func() time.Time { t, _ := time.Parse("2006-01-02", "${rawValue}"); return t}()`;
      } else {
        this.context.importManager.add('time');
        const timeFormat = (<DateTimeSchema>schema).format === 'date-time-rfc1123' ? 'time.RFC1123' : 'time.RFC3339Nano';
        ret = `func() time.Time { t, _ := time.Parse(${timeFormat}, "${rawValue}"); return t}()`;
      }
    } else if (goType === 'map[string]interface{}') {
      ret = this.objectToString(rawValue);
    } else if (goType === 'interface{}' && Array.isArray(rawValue)) {
      ret = this.arrayToString(rawValue);
    } else if (goType === 'interface{}' && typeof rawValue === 'object') {
      ret = this.objectToString(rawValue);
    } else if (goType === 'interface{}' && _.isNumber(rawValue)) {
      ret = `float64(${rawValue})`;
    } else if (goType === 'interface{}' && _.isString(rawValue)) {
      ret = this.getStringValue(rawValue);
    } else if (goType === 'bool') {
      ret = rawValue.toString();
    }

    if (isPtr) {
      const ptrConverts = {
        string: 'Ptr',
        bool: 'Ptr',
        'time.Time': 'Ptr',
        int32: 'Ptr[int32]',
        int64: 'Ptr[int64]',
        float32: 'Ptr[float32]',
        float64: 'Ptr[float64]',
      };

      if ([SchemaType.Choice, SchemaType.SealedChoice].indexOf(schema.type) >= 0) {
        ret = `to.Ptr(${ret})`;
      } else if (Object.prototype.hasOwnProperty.call(ptrConverts, goType)) {
        ret = `to.${ptrConverts[goType]}(${ret})`;
        this.context.importManager.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
      } else {
        ret = '&' + ret;
      }
    }

    return ret;
  }

  protected getLanguageName(meta: any): string {
    return (<Metadata>meta).language.go.name;
  }

  protected objectToString(rawValue: any): string {
    let ret = 'map[string]interface{}{\n';
    for (const [key, value] of Object.entries(rawValue)) {
      if (_.isArray(value)) {
        ret += `"${key}":`;
        ret += this.arrayToString(value);
        ret += ',\n';
      } else if (_.isObject(value)) {
        ret += `"${key}":`;
        ret += this.objectToString(value);
        ret += ',\n';
      } else if (_.isString(value)) {
        ret += `"${key}": ${this.getStringValue(value)},\n`;
      } else if (value === null) {
        ret += `"${key}": nil,\n`;
      } else if (_.isNumber(value)) {
        ret += `"${key}": float64(${value}),\n`;
      } else {
        ret += `"${key}": ${value},\n`;
      }
    }
    ret += '}';
    return ret;
  }

  protected arrayToString(rawValue: any): string {
    let ret = '[]interface{}{\n';
    for (const item of rawValue) {
      if (_.isArray(item)) {
        ret += this.arrayToString(item);
        ret += ',\n';
      } else if (_.isObject(item)) {
        ret += this.objectToString(item);
        ret += ',\n';
      } else if (_.isString(item)) {
        ret += `${Helper.quotedEscapeString(item)},\n`;
      } else if (item === null) {
        ret += 'nil,\n';
      } else if (_.isNumber(item)) {
        ret += `float64(${item}),\n`;
      } else {
        ret += `${item},\n`;
      }
    }
    ret += '}';
    return ret;
  }
}

export class MockTestCodeGenerator extends BaseCodeGenerator {
  public generateCode(extraParam: Record<string, unknown> = {}): void {
    this.renderAndWrite(this.context.codeModel.testModel.mockTest, 'mockTest.go.njk', `${this.getFilePrefix(Config.testFilePrefix)}mock_test.go`, extraParam, {
      getParamsValue: (params: Array<ParameterOutput>) => {
        return params
          .map((p) => {
            return p.paramOutput;
          })
          .join(', ');
      },
    });
  }
}
