// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
import { execSync } from 'child_process';
import { opendirSync } from 'fs';
import { join } from 'path';

// default to the root of the repo
let dir = process.cwd();

// check if a directory was specified
let args = process.argv.slice(2);
if (args.length > 0) {
  dir = args[0];
}

recursiveUpdateGoMod(dir);

function recursiveUpdateGoMod(cur) {
    const dir = opendirSync(cur);
    while (true) {
        const dirEnt = dir.readSync()
        if (dirEnt === null) {
            break;
        }
        if (dirEnt.isFile() && dirEnt.name === 'go.mod') {
            console.log('go get -u all toolchain@none ' + cur);
            execSync('go get -u all toolchain@none', { cwd: cur });
            console.log('go mod tidy ' + cur);
            execSync('go mod tidy', { cwd: cur });
        } else if (dirEnt.isDirectory()) {
            recursiveUpdateGoMod(join(cur, dirEnt.name));
        }
    }
    dir.close();
}
