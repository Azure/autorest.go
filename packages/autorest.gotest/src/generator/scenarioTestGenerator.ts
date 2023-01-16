/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as _ from 'lodash';
import * as path from 'path';
import { BaseCodeGenerator } from './baseGenerator';
import { Config } from '../common/constant';
import { ExampleParameter, ExampleValue, OutputVariableModelType, StepRestCallModel, TestDefinitionModel, TestScenarioModel } from '@autorest/testmodeler/dist/src/core/model';
import { GoExampleModel, ParameterOutput, VariableOutput } from '../common/model';
import { camelCase, snakeCase, upperFirst, lowerFirst } from 'lodash';
import { GroupProperty, Parameter } from '@autorest/codemodel';
import { Helper } from '@autorest/testmodeler/dist/src/util/helper';
import { MockTestDataRender } from './mockTestGenerator';
import { OavStepType } from '@autorest/testmodeler/dist/src/common/constant';
import { Step, Variable } from 'oav/dist/lib/apiScenario/apiScenarioTypes';

export class ScenarioTestDataRender extends MockTestDataRender {
  packagePrefixForGlobalVariables = 'testsuite.';
  globalVariables: Record<string, string | Variable> = {};
  parentVariables: Record<string, string | Variable> = {};
  currentVariables: Record<string, string | Variable> = {};
  scenarioReferencedVariables: Set<string> = new Set<string>();

  public renderData(): void {
    for (const testDef of this.context.codeModel.testModel.scenarioTests) {
      this.generateScenarioTestData(testDef);
    }
  }

  protected generateScenarioTestData(testDef: TestDefinitionModel): void {
    if (testDef.scope.toLowerCase() === 'resourcegroup') {
      this.context.importManager.add('github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources');
    }

    this.globalVariables = {
      ...testDef.requiredVariablesDefault,
      ...testDef.variables,
    };

    for (const step of testDef.prepareSteps) {
      // inner variable should overwrite outer ones
      this.parentVariables = {
        ...this.globalVariables,
      };
      this.currentVariables = {
        ...step.variables,
      };
      this.genRenderDataForStep(step);
    }
    for (const scenario of <Array<TestScenarioModel>>testDef.scenarios) {
      if (scenario.scenario === undefined) {
        scenario.scenario = scenario.description;
      }
      this.scenarioReferencedVariables = new Set<string>();
      for (const step of scenario.steps) {
        // inner variable should overwrite outer ones
        this.parentVariables = {
          ...this.globalVariables,
          ...scenario.requiredVariablesDefault,
          ...scenario.variables,
        };
        this.currentVariables = {
          ...step.variables,
        };
        this.genRenderDataForStep(step);
      }

      // remove useless variable
      for (const variableName of Object.keys(scenario.variables || {})) {
        if (!this.scenarioReferencedVariables.has(variableName)) {
          delete scenario.variables[variableName];
        }
      }

      // resolve scenario variables
      this.parentVariables = {
        ...this.globalVariables,
      };
      this.currentVariables = {
        ...scenario.requiredVariablesDefault,
        ...scenario.variables,
      };
      scenario['variablesOutput'] = {};
      for (const [key, value] of Object.entries(scenario.variables || {})) {
        scenario['variablesOutput'][key] = this.variableValueToString(key, value);
      }
    }
    for (const step of testDef.cleanUpSteps) {
      // inner variable should overwrite outer ones
      this.parentVariables = {
        ...this.globalVariables,
      };
      this.currentVariables = {
        ...step.variables,
      };
      this.genRenderDataForStep(step);
    }

    // resolve scope variables
    this.parentVariables = {};
    this.currentVariables = {
      ...this.globalVariables,
    };
    testDef['variablesOutput'] = {};
    for (const [key, value] of Object.entries(testDef.variables || {})) {
      testDef['variablesOutput'][key] = this.variableValueToString(key, value, false);
    }
  }

  protected genRenderDataForStep(step: Step): void {
    switch (step.type) {
      case OavStepType.restCall: {
        const example = <GoExampleModel>(<StepRestCallModel>step).exampleModel;
        if (example === undefined) {
          break;
        }
        // request and response parse
        this.fillExampleOutput(example);

        // response output variable convert
        if (step.outputVariables && Object.keys(step.outputVariables).length > 0) {
          example.checkResponse = true;
          for (const [variableName, variableConfig] of Object.entries((<StepRestCallModel>step).outputVariablesModel)) {
            let isPtr = false;
            for (let i = 0; i < variableConfig.length; i++) {
              if (variableConfig[i].type === OutputVariableModelType.object) {
                variableConfig[i]['languageName'] = `.${variableConfig[i].languages.go.name}`;
                isPtr = !variableConfig[i].languages.go?.byValue;
              } else if (variableConfig[i].type === OutputVariableModelType.index) {
                variableConfig[i]['languageName'] = `[${variableConfig[i].index}]`;
              } else if (variableConfig[i].type === OutputVariableModelType.key) {
                variableConfig[i]['languageName'] = `["${variableConfig[i].key}"]`;
              }
            }
            step.outputVariables[variableName]['isPtr'] = isPtr;
          }
        } else {
          example.checkResponse = false;
        }
        break;
      }
      case OavStepType.armTemplate: {
        const copyOfPayload = _.cloneDeep(step.armTemplatePayload);
        // template parse
        step['armTemplateOutput'] = this.objectToString(copyOfPayload);

        break;
      }
      default:
    }
  }

  protected toParametersOutput(
    paramsSig: Array<[string, string, Parameter | GroupProperty]>,
    exampleParameters: Array<ExampleParameter>,
    isClient = false,
  ): Array<ParameterOutput> {
    return paramsSig.map(([paramName, typeName, parameter]) => {
      if (paramName === 'ctx') {
        return new ParameterOutput('ctx', this.packagePrefixForGlobalVariables + 'ctx');
      }
      return new ParameterOutput(paramName, this.genParameterOutput(paramName, typeName, parameter, exampleParameters, isClient));
    });
  }

  // For some method which has no subscriptionId param but client has, oav will not do the variable replacement. So we need to specific handle it.
  protected exampleValueToString(exampleValue: ExampleValue, isPtr: boolean, elemByVal = false, inArray = false): string {
    if (exampleValue.language?.default?.name === 'SubscriptionId') {
      return this.packagePrefixForGlobalVariables + 'subscriptionId';
    } else {
      return super.exampleValueToString(exampleValue, isPtr, elemByVal, inArray);
    }
  }

  protected getDefaultValue(param: Parameter | ExampleValue, isPtr: boolean, elemByVal = false): string {
    if (param.language?.default?.name === 'SubscriptionId') {
      return this.packagePrefixForGlobalVariables + 'subscriptionId';
    } else {
      return super.getDefaultValue(param, isPtr, elemByVal);
    }
  }

  protected variableValueToString(key: string, variable: Variable, replaceGlobal = true): VariableOutput {
    let type: string, value: string;
    if (variable.type === 'string' || variable.type === 'secureString') {
      type = 'string';
      if (variable.value !== undefined) {
        value = this.getStringValue(variable.value);
      } else if (variable.prefix !== undefined) {
        type = 'prefix-string';
        value = this.getStringValue(variable.prefix);
      }
    } else if (variable.type === 'bool') {
      type = 'bool';
      if (variable.value !== undefined) {
        value = variable.value ? 'true' : 'false';
      }
    } else if (variable.type === 'int') {
      type = 'int';
      if (variable.value !== undefined) {
        value = variable.value + '';
      }
    } else if (variable.type === 'array') {
      type = '[]any';
      if (variable.value !== undefined) {
        value = this.arrayToString(variable.value);
      }
    } else if (variable.type === 'object' || variable.type === 'secureObject') {
      type = 'map[string]any';
      if (variable.value !== undefined) {
        value = this.objectToString(variable.value);
      }
    }
    if (replaceGlobal && variable.value === undefined && Object.prototype.hasOwnProperty.call(this.globalVariables, key)) {
      value = this.packagePrefixForGlobalVariables + key;
    }
    return new VariableOutput(type, value);
  }

  protected getStringValue(rawValue: any): string {
    if (typeof rawValue === 'string') {
      return this.parseOavVariable(rawValue, { ...this.parentVariables, ...this.currentVariables }).join(' + ');
    } else {
      return Helper.quotedEscapeString(rawValue);
    }
  }

  protected getVariableValue(rawValue: any): string {
    let ret = '';
    const variable = rawValue.slice(2, -1);
    if (Object.prototype.hasOwnProperty.call(this.currentVariables, variable)) {
      const variableValue = this.currentVariables[variable];
      if (typeof variableValue === 'string') {
        return this.getStringValue(variableValue);
      } else {
        if (variableValue.type === 'string') {
          return this.getStringValue(variableValue.value);
        } else if (variableValue.type === 'int') {
          return this.getNumberValue(variableValue.value);
        } else if (variableValue.type === 'bool') {
          return this.getBoolValue(variableValue.value);
        }
      }
    } else {
      if (Object.prototype.hasOwnProperty.call(this.globalVariables, variable)) {
        ret = this.packagePrefixForGlobalVariables + variable;
      } else {
        ret = variable;
      }
      this.scenarioReferencedVariables.add(variable);
    }
    return ret;
  }

  protected getNumberValue(rawValue: any): string {
    if (typeof rawValue === 'string' && rawValue.startsWith('$')) {
      return this.getVariableValue(rawValue);
    } else {
      return `${Number(rawValue)}`;
    }
  }

  protected getBoolValue(rawValue: any): string {
    if (typeof rawValue === 'string' && rawValue.startsWith('$')) {
      return this.getVariableValue(rawValue);
    } else {
      return rawValue.toString();
    }
  }

  // Pick out $(...) variables from normal string
  // For example: "a string with ${var} inside" => ['"a string with "', 'var', '" inside"']
  protected parseOavVariable(s: string, definedVariables: Record<string, string | Variable>): Array<string> {
    if (!s) {
      return ['""'];
    }

    const re = /\$\([^)]+\)/g;
    const ret: Array<string> = [];
    const m = s.match(re);
    let placeHolders = [];

    if (m) {
      placeHolders = m.map((x) => x.toString());
    }
    for (const placeHolder of placeHolders.filter((x) => Object.prototype.hasOwnProperty.call(definedVariables, x.slice(2, -1)))) {
      const p = s.indexOf(placeHolder);
      if (p > 0) {
        ret.push(Helper.quotedEscapeString(s.substring(0, p)));
      }
      ret.push(this.getVariableValue(placeHolder));
      s = s.substring(p + placeHolder.length);
    }
    if (s.length > 0) {
      ret.push(Helper.quotedEscapeString(s));
    }

    return ret;
  }
}

export class ScenarioTestCodeGenerator extends BaseCodeGenerator {
  public generateCode(extraParam: Record<string, unknown> = {}): void {
    if (this.context.codeModel.testModel.scenarioTests.length > 0) {
      for (const testDef of this.context.codeModel.testModel.scenarioTests) {
        const file = path.basename(testDef._filePath);
        const filename = file.split('.').slice(0, -1).join('.');
        const rpRegex = /(\/|\\)(?<rpName>[^/\\]+)(\/|\\)(arm[^/\\]+)/;
        const execResult = rpRegex.exec(this.context.testConfig.getValue(Config.outputFolder));
        extraParam['rpName'] = execResult?.groups ? execResult.groups['rpName'] : '';
        extraParam['globalVariables'] = Object.keys({
          ...testDef.requiredVariablesDefault,
          ...testDef.variables,
        });
        this.renderAndWrite(
          { ...testDef, testCaseName: upperFirst(camelCase(filename)) },
          'scenarioTest.go.njk',
          `${this.getFilePrefix(Config.testFilePrefix)}${filename.toLowerCase()}_live_test.go`,
          extraParam,
          {
            snakeCase: snakeCase,
            lowerFirst: lowerFirst,
            camelCase: camelCase,
            upperFirst: upperFirst,
            quotedEscapeString: Helper.quotedEscapeString,
            genVariableName: (typeName: string) => {
              // This function generate variable instance name from variable type name
              // For instance:
              //   1) VirtualMachineResponse  --> virtualMachineResponse
              //   2) armCompute.VirtualMachineResponse  --> virtualMachineResponse   // remove package name
              //   3) *VirtualMachineResponse  --> virtualMachineResponse  // remove char of pointer.
              return lowerFirst(typeName.split('.').join('*').split('*').pop());
            },
            getParamsValue: (params: Array<ParameterOutput>) => {
              return params
                .map((p) => {
                  return p.paramOutput;
                })
                .join(', ');
            },
          },
        );
      }
    }
  }
}
