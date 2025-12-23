/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import { ErrorCode } from "../../../codemodel.go/src/errors.js";

/**
 * CodegenError is thrown when the emitter fails some condition
 * in order to generate a part of the code model.
 */
export class CodegenError extends Error {
  readonly code: ErrorCode;

  constructor(code: ErrorCode, message: string) {
    super(message);
    this.code = code;
  }
}
