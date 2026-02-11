/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as tcgc from '@azure-tools/typespec-client-generator-core';
import { createRequire } from 'module';

/**
 * Build the package metadata from the TCGC SDK package metadata.
 *
 * @param metadata the TCGC SDK package metadata
 * @returns the metadata object for the code model
 */
export function buildMetadata(metadata: tcgc.SdkPackage<tcgc.SdkHttpOperation>['metadata']): Record<string, unknown> {
  const result: Record<string, unknown> = {};
  if (metadata.apiVersions && metadata.apiVersions.size > 0) {
    result.apiVersions = Object.fromEntries(metadata.apiVersions);
  }

  const packageJson = createRequire(import.meta.url)('../../../../package.json') as Record<string, never>;
  result.emitterVersion = packageJson['version'];

  return result;
}
