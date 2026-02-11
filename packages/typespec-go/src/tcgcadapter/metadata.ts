/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { createRequire } from 'module';

/**
 * Builds the package metadata from the TCGC context.
 *
 * @param metadata the TCGC SDK package metadata
 * @returns the metadata object for the code model
 */
export function buildMetadata(metadata: any): Record<string, unknown> {
  const packageJson = createRequire(import.meta.url)('../../../../package.json') as Record<string, never>;

  const result: Record<string, unknown> = {};

  // convert apiVersions Map to a plain object for JSON serialization
  const apiVersions = metadata.apiVersions;
  if (apiVersions && apiVersions.size > 0) {
    result.apiVersions = Object.fromEntries(apiVersions);
  }

  result.emitterVersion = packageJson['version'];

  return result;
}
