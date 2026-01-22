/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { ClientAdapter } from './clients.js';
import { AdapterError } from './errors.js';
import { TypeAdapter } from './types.js';
import { GoEmitterOptions } from '../lib.js';
import * as go from '../../../codemodel.go/src/index.js';
import * as naming from '../../../naming.go/src/index.js';
import * as tcgc from '@azure-tools/typespec-client-generator-core';
import * as tsp from '@typespec/compiler';
import { createRequire } from 'module';

/**
 * Constant representing "all" API versions in TCGC metadata.
 * Used to indicate that all versions should be included.
 */
const ALL_API_VERSIONS = 'all';

/**
 * ExternalError is thrown when an external component reports a
 * diagnostic error that would prevent the emitter from proceeding.
 */
export class ExternalError extends Error { }

/** Adapter converts the tcgc code model to an instance of the Go code model */
export class Adapter {
  /**
   * Creates an Adapter for the specified EmitContext.
   * 
   * @param context the compiler context from which to create the Adapter
   * @returns a new Adapter for the provided context
   */
  static async create(context: tsp.EmitContext<GoEmitterOptions>): Promise<Adapter> {
    naming.CommonAcronyms.push('^iso\\d+$');
    // @encodedName can be used in XML scenarios, it
    // is effectively the same as TypeSpec.Xml.@name.
    // however, it's filtered out by default so we need
    // to add it to the allow list of decorators
    const ctx = await tcgc.createSdkContext(context, '@azure-tools/typespec-go', {
      additionalDecorators: ['TypeSpec\\.@encodedName'],
      disableUsageAccessPropagationToBase: true,
    });

    context.program.reportDiagnostics(ctx.diagnostics);
    for (const diag of ctx.diagnostics) {
      if (diag.severity === 'error') {
        // there's no point in continuing if tcgc
        // has reported diagnostic errors, so exit.
        // this prevents spurious crashes in the
        // emitter as our input state is invalid.
        throw new ExternalError();
      }
    }

    return new Adapter(ctx, context.options, context.emitterOutputDir);
  }

  private readonly ctx: tcgc.SdkContext;
  private readonly options: GoEmitterOptions;
  private readonly codeModel: go.CodeModel;

  private constructor(ctx: tcgc.SdkContext, options: GoEmitterOptions, emitterOutputDir: string) {
    this.ctx = ctx;
    this.options = options;

    if (this.options['containing-module'] && this.options.module) {
      throw new AdapterError('InvalidArgument', 'module and containing-module are mutually exclusive');
    }

    const goOptions = new go.Options(
      this.options['generate-fakes'] === true,
      this.options['inject-spans'] === true,
      this.options['disallow-unknown-fields'] === true,
      // generate-examples has been deprecated, for compat we still support it.
      this.options['generate-examples'] === true || this.options['generate-samples'] === true
    );
    goOptions.headerText = this.ctx.sdkPackage.licenseInfo?.header;
    goOptions.licenseText = this.ctx.sdkPackage.licenseInfo?.description;
    goOptions.azcoreVersion = this.options['azcore-version'];
    goOptions.omitConstructors = this.options['omit-constructors'] ?? false;

    let root: go.ContainingModule | go.Module;
    if (this.options.module) {
      root = new go.Module(this.options.module);
    } else if (this.options['containing-module']) {
      root = new go.ContainingModule(this.options['containing-module']);
      root.package = new go.Package(naming.packageNameFromOutputFolder(emitterOutputDir), root);
    } else {
      throw new AdapterError('InvalidArgument', 'missing argument module or containing-module');
    }

    const info = new go.Info(this.ctx.sdkPackage.crossLanguagePackageId);
    const codeModelType: go.CodeModelType = this.ctx.arm === true ? 'azure-arm' : 'data-plane';
    this.codeModel = new go.CodeModel(info, codeModelType, goOptions, root);

    // get the emitter version from our package.json
    const packageJson = createRequire(import.meta.url)('../../../../package.json') as Record<string, never>;
    
    // Build metadata with support for multiple API versions
    this.codeModel.metadata = this.buildMetadata(packageJson['version']);

    this.codeModel.options.rawJSONAsBytes = this.options['rawjson-as-bytes'] ?? false;
    this.codeModel.options.sliceElementsByval = this.options['slice-elements-byval'] ?? false;
    this.codeModel.options.factoryGatherAllParams = this.options['factory-gather-all-params'] ?? true;
  }

  /**
   * Builds metadata object with support for multiple API versions.
   * If there's a single API version, returns backward-compatible format with single apiVersion field.
   * If there are multiple services with different API versions, returns services object.
   * 
   * @param emitterVersion the version of the emitter
   * @returns metadata object for the code model
   */
  private buildMetadata(emitterVersion: string): Record<string, unknown> {
    // Collect unique service-version pairs from all clients
    const serviceVersionMap = new Map<string, string>();
    
    // First check if there's a single package-level API version (backward compatibility)
    const packageApiVersion = this.ctx.sdkPackage.metadata.apiVersion;
    if (packageApiVersion && packageApiVersion !== ALL_API_VERSIONS) {
      // Single API version case - use the package metadata directly for backward compatibility
      return {
        apiVersion: packageApiVersion,
        emitterVersion
      };
    }
    
    // Multiple services case: collect API versions from package versions map
    // This map contains namespace -> versions mapping for all services
    const packageVersions = this.ctx.getPackageVersions();
    for (const [namespace, versions] of packageVersions.entries()) {
      if (versions && versions.length > 0) {
        // Use the first (or configured) version for this service
        const version = versions[0];
        if (version && version !== ALL_API_VERSIONS) {
          serviceVersionMap.set(namespace.name, version);
        }
      }
    }
    
    // Also check from clients in case package versions is empty
    if (serviceVersionMap.size === 0) {
      for (const clientType of this.ctx.sdkPackage.clients) {
        // Get service namespace from the raw SdkClient
        if (clientType.__raw.kind === 'SdkClient') {
          const services = Array.isArray(clientType.__raw.service) ? clientType.__raw.service : [clientType.__raw.service];
          
          // Get API versions from the client type
          if (clientType.apiVersions && clientType.apiVersions.length > 0) {
            // Use the first API version which is typically the one configured for emission
            const apiVersion = clientType.apiVersions[0];
            
            if (apiVersion && apiVersion !== ALL_API_VERSIONS) {
              // Map each service to its API version
              for (const service of services) {
                const serviceName = service.name;
                // Only add if not already present (first client wins for a given service)
                if (!serviceVersionMap.has(serviceName)) {
                  serviceVersionMap.set(serviceName, apiVersion);
                }
              }
            }
          }
        }
      }
    }
    
    // Build the metadata based on the number of unique service-version pairs
    if (serviceVersionMap.size === 0) {
      // No API versions found
      return {
        emitterVersion
      };
    } else if (serviceVersionMap.size === 1) {
      // Single service - still use single apiVersion format for consistency
      const [, apiVersion] = Array.from(serviceVersionMap.entries())[0];
      return {
        apiVersion,
        emitterVersion
      };
    } else {
      // Multiple services - use new format
      const services: Record<string, { apiVersion: string }> = {};
      for (const [serviceName, apiVersion] of serviceVersionMap.entries()) {
        services[serviceName] = { apiVersion };
      }
      return {
        emitterVersion,
        services
      };
    }
  }

  /** performs all the steps to convert tcgc to the Go code model */
  tcgcToGoCodeModel(): go.CodeModel {
    // TODO: stuttering fix-ups will need some rethinking for namespaces
    const packageName = this.codeModel.root.kind === 'containingModule' ? this.codeModel.root.package.name : naming.packageNameFromOutputFolder(this.ctx.emitContext.emitterOutputDir);
    fixStutteringTypeNames(this.ctx.sdkPackage, packageName, this.options);

    const ta = new TypeAdapter(this.ctx, this.codeModel);
    ta.adaptTypes();

    const ca = new ClientAdapter(ta);
    ca.adaptClients();

    return this.codeModel;
  }
}

/**
 * fixes up names in the tcgc model to avoid stuttering.
 * 
 * @param sdkPackage the tcgc data model
 * @param packageName the package name used to remove stuttering
 * @param options the Go emitter options
 */
function fixStutteringTypeNames(sdkPackage: tcgc.SdkPackage<tcgc.SdkHttpOperation>, packageName: string, options: GoEmitterOptions): void {
  let stutteringPrefix = packageName;

  if (options.stutter) {
    stutteringPrefix = options.stutter;
  } else {
    // if there's a well-known prefix, remove it
    if (stutteringPrefix.startsWith('arm')) {
      stutteringPrefix = stutteringPrefix.substring(3);
    } else if (stutteringPrefix.startsWith('az')) {
      stutteringPrefix = stutteringPrefix.substring(2);
    }
  }
  stutteringPrefix = stutteringPrefix.toUpperCase();

  // ensure that enum, client, and struct type names don't stutter

  const recursiveWalkClients = function(client: tcgc.SdkClientType<tcgc.SdkHttpOperation>): void {
    // NOTE: we MUST do this before calling trimPackagePrefix to properly handle
    // the case where the client name is the same as the package name.
    if (!client.name.match(/Client$/)) {
      client.name += 'Client';
    }
    client.name = naming.trimPackagePrefix(stutteringPrefix, client.name);

    // fix up the synthesized type names for page responses
    if (client.children && client.children.length > 0) {
      for (const child of client.children) {
        recursiveWalkClients(child);
      }
    }
    for (const sdkMethod of client.methods) {
      if (sdkMethod.kind !== 'paging') {
        continue;
      }

      for (const httpResp of sdkMethod.operation.responses.values()) {
        if (!httpResp.type || httpResp.type.kind !== 'model') {
          continue;
        }

        httpResp.type.name = naming.trimPackagePrefix(stutteringPrefix, httpResp.type.name);
      }
    }
  };

  for (const sdkClient of sdkPackage.clients) {
    recursiveWalkClients(sdkClient);
  }

  // check if the name collides with an existing name. we only do
  // this for model and enum types, as clients get a suffix.
  const nameCollision = function(newName: string): boolean {
    for (const modelType of sdkPackage.models) {
      if (modelType.name === newName) {
        return true;
      }
    }
    for (const enumType of sdkPackage.enums) {
      if (enumType.name === newName) {
        return true;
      }
    }
    return false;
  };

  // tracks type name collilsions due to renaming
  const collisions = new Array<string>();

  // trims the stuttering prefix from typeName and returns the new name.
  // if there's a collision, an entry is added to the collision list.
  const renameType = function(typeName: string): string {
    const originalName = typeName;
    const newName = naming.trimPackagePrefix(stutteringPrefix, originalName); 

    // if the type was renamed to remove stuttering, check if it collides with an existing type name
    if (newName !== originalName && nameCollision(newName)) {
      collisions.push(`type ${originalName} was renamed to ${newName} which collides with an existing type name`);
    }
    return newName;
  };

  // to keep compat with autorest.go, this is off by default
  if (options['fix-const-stuttering'] === true) {
    for (const sdkEnum of sdkPackage.enums) {
      sdkEnum.name = renameType(sdkEnum.name);
    }
  }

  for (const modelType of sdkPackage.models) {
    modelType.name = renameType(modelType.name);
  }

  if (collisions.length > 0) {
    throw new AdapterError('NameCollision', collisions.join('\n'));
  }
}
