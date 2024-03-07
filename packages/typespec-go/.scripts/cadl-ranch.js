// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
import { execSync } from 'child_process';

const toolsModRoot = execSync('git rev-parse --show-toplevel').toString().trim() + '/packages/typespec-go/node_modules/@azure-tools/';

const switches = [];
switch (process.argv[2]) {
  case '--start':
    switches.push('server');
    switches.push('start');
    switches.push(toolsModRoot + 'cadl-ranch-specs/http');
    break;
  case '--stop':
    switches.push('server');
    switches.push('stop');
    break;
}

if (switches.length === 0) {
  throw new Error('missing arg: [--start] [--stop]');
}

const cmdLine = toolsModRoot + 'cadl-ranch/node_modules/.bin/cadl-ranch ' + switches.join(' ');
console.log(cmdLine);
execSync(cmdLine);
