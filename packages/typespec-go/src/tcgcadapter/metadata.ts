/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as tcgc from '@azure-tools/typespec-client-generator-core';
import * as go from '../../../codemodel.go/src/index.js';

/**
 * Constant representing "all" API versions in TCGC metadata.
 * Used to indicate that all versions should be included.
 */
const ALL_API_VERSIONS = 'all';

/**
 * Builds metadata object by collecting API version information from TCGC.
 * The actual formatting logic is handled by generateMetadataFile in codegen.go.
 * 
 * @param ctx the TCGC SDK context containing package and version information
 * @param emitterVersion the version of the emitter
 * @returns metadata object for the code model
 */
export function buildMetadata(ctx: tcgc.SdkContext, emitterVersion: string): go.Metadata {
  const metadata: go.Metadata = {
    emitterVersion
  };
  
  // First check if there's a single package-level API version (backward compatibility)
  const packageApiVersion = ctx.sdkPackage.metadata.apiVersion;
  if (packageApiVersion && packageApiVersion !== ALL_API_VERSIONS) {
    // Single API version case - use the package metadata directly
    metadata.apiVersion = packageApiVersion;
    return metadata;
  }
  
  // Multiple services case: collect API versions from package versions map
  const serviceVersionMap = new Map<string, string>();
  
  // This map contains namespace -> versions mapping for all services
  const packageVersions = ctx.getPackageVersions();
  for (const [namespace, versions] of packageVersions.entries()) {
    if (versions && versions.length > 0) {
      // Use the first (or configured) version for this service
      const version = versions[0];
      if (version && version !== ALL_API_VERSIONS) {
        serviceVersionMap.set(namespace.name, version);
      }
    }
  }

  if (serviceVersionMap.size === 0) {
    // No valid service versions found, return empty metadata
    return metadata;
}   
  
  // If we found service-version mappings, add them to metadata
    const services: Record<string, { apiVersion: string }> = {};
    for (const [serviceName, apiVersion] of serviceVersionMap.entries()) {
      services[serviceName] = { apiVersion };
    }
    metadata.services = services;
  
  return metadata;
}
