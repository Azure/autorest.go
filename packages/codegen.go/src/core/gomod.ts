/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../../codemodel.go/src/index.js';
import { CodegenError } from './errors.js';
import * as semver from 'semver';

/**
 * Creates the content for the go.mod file.
 * If there's a preexisting go.mod file, update its specified version of azcore as needed.
 * 
 * @param module the module for which to generate version.go
 * @param options the emitter options
 * @param existingGoMod preexisting go.mod file content
 * @returns the contents for the go.mod file
 */
export function generateGoModFile(module: go.Module, options: go.Options, existingGoMod?: string): string {
  const modName = module.identity;

  // here we specify the minimum version of azcore as required by the code generator.
  // the version can be overwritten by passing the --azcore-version switch during generation.
  let version = '1.21.0';
  if (options.azcoreVersion) {
    // when matching versions, we need to handle beta, non-beta, and pseudo versions
    // 1.2.3-beta.1, 1.2.3, 0.22.1-0.20220315231014-ed309e73db6b
    if (!options.azcoreVersion.match(/^\d+\.\d+\.\d+(?:-[a-zA-Z0-9_.-]+)?$/)) {
      throw new CodegenError('InvalidArgument', `azcore version ${version} must be in the format major.minor.patch[-beta.N]`);
    }
    version = options.azcoreVersion;
  }

  const azcore = 'github.com/Azure/azure-sdk-for-go/sdk/azcore v' + version;
  if (!existingGoMod) {
    // no preexisting go.mod file, generate the default one
    let text = `module ${modName}\n\n`;
    text += 'go 1.24.0\n\n';
    text += `require ${azcore}\n`;
    return text;
  }

  // check if the module identity needs to be replaced due to a major version change
  if (!existingGoMod.match(`module ${modName}$`)) {
    existingGoMod = existingGoMod.replace(/module \S+/, `module ${modName}`);
  }

  // check if the existing version of azcore is greater than or equal to the specified version.
  // note that some modules (e.g. models-only) might not have a dependency on azcore.
  const match = existingGoMod.match(/github\.com\/Azure\/azure-sdk-for-go\/sdk\/azcore\s+v(\d+\.\d+\.\d+(?:-[a-zA-Z0-9_.-]+)?)/);
  if (match) {
    if (match.length < 2) {
      throw new CodegenError('InternalError', 'returned matches were less than expected');
    }
    const existingVer = toSemver(match[1]);
    const specifiedVer = toSemver(version);
    if (lt(existingVer, specifiedVer)) {
      // the existing version of azcore is less than the specified version so update it
      existingGoMod = existingGoMod.replace(/github\.com\/Azure\/azure-sdk-for-go\/sdk\/azcore\s+v\d+\.\d+\.\d+(?:-[a-zA-Z0-9_.-]+)?/, azcore);
    }
  }
  return existingGoMod;
}

// the following was copied from @azure-tools/codegen as it's being deprecated
function toSemver(apiversion: string) {
  // strip off leading "v" or "=" character
  apiversion = apiversion.replace(/^v|^=/gi, "");
  // eslint-disable-next-line no-useless-escape
  const versionedDateRegex = new RegExp(/(^\d{4}\-\d{2}\-\d{2})(\.\d+\.\d+$)/gi);
  if (apiversion.match(versionedDateRegex)) {
    // convert yyyy-mm-dd.x1.x2      --->     (miliseconds since 1970-01-01).x1.x2
    const date = apiversion.replace(versionedDateRegex, "$1");
    const miliseconds = new Date(date).getTime();
    const lastNumbers = apiversion.replace(versionedDateRegex, "$2");
    return `${miliseconds}${lastNumbers}`;
  }
  const [major, minor, revision, tag] =
    /^(\d+)-(\d+)(?:-(\d+))?(.*)/.exec(apiversion) ||
    /(\d*)\.(\d*)\.(\d*)(.*)/.exec(apiversion) ||
    /(\d*)\.(\d*)()(.*)/.exec(apiversion) ||
    /(\d*)()()(.*)/.exec(apiversion) ||
    [];
  return `${Number.parseInt(major || "0") || 0}.${Number.parseInt(minor || "0") || 0}.${
    Number.parseInt(revision || "0") || 0
  }${tag?.startsWith("-") ? tag : ""}`;
}

function lt(apiVersion1: string, apiVersion2: string) {
  const v1 = toSemver(apiVersion1);
  const v2 = toSemver(apiVersion2);
  return semver.lt(v1, v2);
}
// end ports from @azure-tools/codegen
