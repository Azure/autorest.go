/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as tsp from '@typespec/compiler';
import { ErrorCode } from '../../../codemodel.go/src/errors.js';

/**
 * AdapterError is thrown when the emitter fails to convert part of the tcgc code
 * model to the emitter code model. this could be due to the emitter not supporting
 * some tsp construct.
 */
export class AdapterError extends Error {
  readonly code: ErrorCode;
  readonly target: tsp.DiagnosticTarget | typeof tsp.NoTarget;

  constructor(code: ErrorCode, message: string, target?: tsp.DiagnosticTarget) {
    super(message);
    this.code = code;
    this.target = target ?? tsp.NoTarget;
  }
}
