/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../../codemodel.go/src/index.js';

/**
 * Metadata structure that can be provided to generateMetadataFile.
 * Supports both single-service and multi-service formats.
 */
export interface MetadataInput {
  /** Emitter version */
  emitterVersion?: string;
  
  /** Single API version (backward compatible format) */
  apiVersion?: string;
  
  /** Multiple services with their API versions */
  services?: Record<string, { apiVersion: string }>;
}

/**
 * Creates the content in _metadata.json.
 * Handles formatting logic for single vs multiple service scenarios.
 */
export function generateMetadataFile(codeModel: go.CodeModel): string {
  const metadata = codeModel.metadata as MetadataInput | undefined;
  
  if (!metadata) {
    return '';
  }
  
  // Build output metadata based on the structure
  let outputMetadata: Record<string, unknown>;
  
  if (metadata.services && Object.keys(metadata.services).length > 1) {
    // Multiple services - use new format
    outputMetadata = {
      emitterVersion: metadata.emitterVersion,
      services: metadata.services
    };
  } else if (metadata.services && Object.keys(metadata.services).length === 1) {
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
