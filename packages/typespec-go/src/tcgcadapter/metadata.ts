/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as tcgc from '@azure-tools/typespec-client-generator-core';
import { createRequire } from 'module';

/**
 * Builds the package metadata from the TCGC context.
 *
 * @param sdkPackage the TCGC SDK package
 * @returns the metadata object for the code model
 */
export function buildMetadata(sdkPackage: tcgc.SdkPackage<tcgc.SdkHttpOperation>): Record<string, unknown> {
  const packageJson = createRequire(import.meta.url)('../../../../package.json') as Record<string, never>;

  const metadata: Record<string, unknown> = {};

  // convert apiVersions Map to a plain object for JSON serialization
  const apiVersions = sdkPackage.metadata.apiVersions;
  if (apiVersions && apiVersions.size > 0) {
    metadata.apiVersions = Object.fromEntries(apiVersions);
  }

  metadata.emitterVersion = packageJson['version'];

  return metadata;
}
