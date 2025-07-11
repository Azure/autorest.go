/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { ExampleModel, MockTestDefinitionModel } from '@autorest/testmodeler';

interface GoFileData {
  packageName: string;
  packagePath?: string;
  imports: string;
}

export class GoMockTestDefinitionModel extends MockTestDefinitionModel implements GoFileData {
  packageName: string;
  imports: string;
}

export class GoExampleModel extends ExampleModel {
  opName: string;
  isLRO: boolean;
  isPageable: boolean;
  isMultiRespOperation: boolean;
  methodParametersOutput: Array<ParameterOutput>;
  clientParametersOutput: Array<ParameterOutput>;
  factoryClientParametersOutput: Array<ParameterOutput>;
  returnInfo: Array<string>;
  checkResponse: boolean;
  responseOutput: string;
  responseType: string;
  responseIsDiscriminator: boolean;
  responseTypePointer: boolean;
  pollerType: string;
  pageableItemName: string;
}

export class ParameterOutput {
  public constructor(public paramName: string, public paramOutput: string) {}
}

export class VariableOutput {
  public constructor(public type: string, public value: string | undefined = undefined) {}
}
