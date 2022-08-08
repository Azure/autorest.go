/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { BaseCodeGenerator } from './baseGenerator';
import { Config } from '../common/constant';
import { ExampleModel, MockTestDefinitionModel } from '@autorest/testmodeler/dist/src/core/model';
import { MockTestDataRender } from './mockTestGenerator';
import { ParameterOutput } from '../common/model';
import path = require('path');

export class ExampleDataRender extends MockTestDataRender {}

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
      if (fileName !== 'client') {
        fileName = fileName.substring(0, fileName.length-6) + '_client';
      }

      this.renderAndWrite(
        { exampleGroups: exampleGroups },
        'exampleTest.go.njk',
        `${this.getFilePrefix(Config.exampleFilePrefix)}${fileName}_example_test.go`,
        extraParam,
        {
          getParamsValue: (params: Array<ParameterOutput>) => {
            return params
              .map((p) => {
                return p.paramOutput;
              })
              .join(', ');
          },
          getExampleSuffix: (exampleFilename: string) => {
            const name = path.parse(exampleFilename).name;
            return name.charAt(0).toLowerCase() + name.slice(1);
          },
        },
      );
    }
  }
}
