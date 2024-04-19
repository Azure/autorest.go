// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
const execSync = require('child_process').execSync;
const fs = require('fs');
const path = require('path');

recursiveUpdateGoMod(process.env.RUSH_INVOKED_FOLDER);

function recursiveUpdateGoMod(cur) {
    const dir = fs.opendirSync(cur);
    while (true) {
        const dirEnt = dir.readSync()
        if (dirEnt === null) {
            break;
        }
        if (dirEnt.isFile() && dirEnt.name === 'go.mod') {
            console.log('go get -u ' + cur);
            execSync('go get -u', { cwd: cur });
            console.log('go mod tidy ' + cur);
            execSync('go mod tidy', { cwd: cur });
        } else if (dirEnt.isDirectory()) {
            recursiveUpdateGoMod(path.join(cur, dirEnt.name));
        }
    }
    dir.close();
}
