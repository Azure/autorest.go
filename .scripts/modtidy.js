// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
const execSync = require('child_process').execSync;
const fs = require('fs');

recursiveFindGoMod('./packages/autorest.go/test');

function recursiveFindGoMod(cur) {
    const dir = fs.opendirSync(cur);
    while (true) {
        const dirEnt = dir.readSync()
        if (dirEnt === null) {
            break;
        }
        if (dirEnt.isFile() && dirEnt.name === 'go.mod') {
            console.log('go mod tidy ' + cur);
            execSync('go mod tidy', { cwd: cur });
        } else if (dirEnt.isDirectory()) {
            recursiveFindGoMod(cur + '/' + dirEnt.name);
        }
    }
    dir.close();
}
