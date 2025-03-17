/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

/** ErrorCode defines the types of errors */
export type ErrorCode =
  /** the emitter encounted an internal error. this is always a bug in the emitter */
  'InternalError' |

  /** invalid arguments were passed to the emitter */
  'InvalidArgument' |

  /**
   * renaming types resulted in one or more name collisions.
   * this will likely require an update to client.tsp to resolve.
   */
  'NameCollision' |

  /** the emitter does not support the encountered TypeSpec construct */
  'UnsupportedTsp';

/**
 * CodeModelError is thrown when the an invariant in the code model has been violated.
 * This is always a bug in the adapter.
 */
export class CodeModelError extends Error {
  readonly code: ErrorCode = 'InternalError';

  constructor(message: string) {
    super(message);
  }
}
