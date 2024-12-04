// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
import { execSync } from 'child_process';

const nodeModulesRoot = execSync('git rev-parse --show-toplevel').toString().trim() + '/packages/typespec-go/node_modules/';
const spector = nodeModulesRoot + '@typespec/spector/node_modules/.bin/tsp-spector';
const httpSpecs = nodeModulesRoot + '@typespec/http-specs/http';
const azureHttpSpecs = nodeModulesRoot + '@azure-tools/azure-http-specs/http';

const switches = [];
let execSyncOptions;

let spec = httpSpecs;
if (process.argv[3] === '--azure') {
  spec = azureHttpSpecs;
}

switch (process.argv[2]) {
  case '--serve':
    switches.push('serve');
    switches.push(spec);
    execSyncOptions = {stdio: 'inherit'};
    break;
  case '--start':
    switches.push('server');
    switches.push('start');
    switches.push(spec);
    break;
  case '--stop':
    switches.push('server');
    switches.push('stop');
    break;
}

if (switches.length === 0) {
  throw new Error('missing arg: [--start] [--stop]');
}

const cmdLine = spector + ' ' + switches.join(' ');
console.log(cmdLine);
execSync(cmdLine, execSyncOptions);
