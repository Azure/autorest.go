/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { ObjectSchema, Parameter, SchemaType } from '@autorest/codemodel';
import { ExampleModel, MockTestDefinitionModel } from '@autorest/testmodeler/dist/src/core/model';
import { camelCase, trimEnd } from 'lodash';
import { Config } from '../common/constant';
import { sortParametersByRequired } from '../common/helpers';
import { ParameterOutput } from '../common/model';
import { BaseCodeGenerator } from './baseGenerator';
import { MockTestDataRender } from './mockTestGenerator';

export class ExampleDataRender extends MockTestDataRender {
  public renderData(): void {
    super.renderData();
    const allClientParams = new Array<Parameter>();
    for (const group of this.context.codeModel.operationGroups) {
      if (group.language.go!.clientParams) {
        const clientParams = <Array<Parameter>>group.language.go!.clientParams;
        for (const clientParam of clientParams) {
          if (allClientParams.filter((cp) => cp.language.go!.name === clientParam.language.go!.name).length > 0) {
            continue;
          }
          allClientParams.push(clientParam);
        }
      }
    }
    allClientParams.sort(sortParametersByRequired);
    const clientFactoryParametersOutput = new Array<ParameterOutput>();
    for (const clientParam of allClientParams) {
      const isPolymophismValue = clientParam?.schema?.type === SchemaType.Object && (<ObjectSchema>clientParam.schema).discriminator?.property.isDiscriminator === true;
      const isPtr: boolean = isPolymophismValue || !(clientParam.required || clientParam.language.go.byValue === true);
      clientFactoryParametersOutput.push(new ParameterOutput(this.getLanguageName(clientParam), this.getDefaultValue(clientParam, isPtr)));
    }
    this.context.codeModel.testModel.mockTest['clientFactoryParametersOutput'] = clientFactoryParametersOutput;
  }
}

export class ExampleCodeGenerator extends BaseCodeGenerator {
  public generateCode(extraParam: Record<string, unknown> = {}): void {
    for (const [_, exampleGroups] of Object.entries(MockTestDefinitionModel.groupByOperationGroup(this.context.codeModel.testModel.mockTest.exampleGroups))) {
      let exampleModel: ExampleModel = null;
      for (const exampleGroup of exampleGroups) {
        if (exampleGroup.examples.length > 0) {
          exampleModel = exampleGroup.examples[0];
          break;
        }
      }
      if (exampleModel === null) {
        continue;
      }

      let fileName = exampleModel.operationGroup.language.go!.clientName.toLowerCase();
      if (fileName !== 'client' && fileName.endsWith('client')) {
        fileName = fileName.substring(0, fileName.length - 6) + '_client';
      }

      this.renderAndWrite(
        {
          clientFactoryParametersOutput: this.context.codeModel.testModel.mockTest['clientFactoryParametersOutput'],
          exampleGroups: exampleGroups,
          swaggerCommit: this.context.swaggerCommit,
        },
        'exampleTest.go.njk',
        `${this.getFilePrefix(Config.exampleFilePrefix)}${fileName}_example_test.go`,
        extraParam,
        {
          getParamsValue: this.getParamsValue,
          getExampleSuffix: (exampleKey: string) => {
            return camelCase(exampleKey);
          },
          getCommentResponseOutput: this.getCommentResponseOutput,
        },
      );
    }
  }

  public getCommentResponseOutput(responseOutput: string): string {
    let result = '';
    const indent = '\t';
    let indentNum = 0;
    let firstLine = true;
    for (const line of responseOutput.split('\n')) {
      if (!firstLine) {
        result += '// ';
      } else {
        firstLine = false;
      }
      for (const ch of line) {
        if (ch === '}') {
          indentNum--;
        } else {
          break;
        }
      }
      result += indent.repeat(indentNum) + line + '\n';
      if (line.endsWith('{')) {
        indentNum++;
      }
    }
    return trimEnd(result, '\n');
  }

  public getParamsValue(params: Array<ParameterOutput>) {
    return params
      .map((p) => {
        return p.paramOutput;
      })
      .join(', ');
  }
}
