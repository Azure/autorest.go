/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../../codemodel.go/src/index.js';

/**
 * Creates the content in _metadata.json.
 * Handles formatting logic for single vs multiple service scenarios.
 */
export function generateMetadataFile(metadata?: go.Metadata): string {
  
  if (!metadata) {
    return '';
  }
  
  // Build output metadata based on the structure
  let outputMetadata: Record<string, unknown>;
  const serviceCount = metadata.services ? Object.keys(metadata.services).length : 0;
  
  if (metadata.services && serviceCount > 1) {
    // Multiple services - use new format
    outputMetadata = {
      emitterVersion: metadata.emitterVersion,
      services: metadata.services
    };
  } else if (metadata.services && serviceCount === 1) {
    // Single service from services map - extract to flat format for backward compatibility
    const [, serviceInfo] = Object.entries(metadata.services)[0];
    outputMetadata = {
      apiVersion: serviceInfo.apiVersion,
      emitterVersion: metadata.emitterVersion
    };
  } else if (metadata.apiVersion) {
    // Single API version (backward compatible)
    outputMetadata = {
      apiVersion: metadata.apiVersion,
      emitterVersion: metadata.emitterVersion
    };
  } else {
    // No API version information
    outputMetadata = {
      emitterVersion: metadata.emitterVersion
    };
  }
  
  // Return the formatted JSON string
  return JSON.stringify(outputMetadata, null, 2) + '\n';
}
