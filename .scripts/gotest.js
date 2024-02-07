// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
const execSync = require('child_process').execSync;
const fs = require('fs');
const path = require('path');

recursiveFindGoMod(process.env.RUSH_INVOKED_FOLDER);

function recursiveFindGoMod(cur) {
  const dir = fs.opendirSync(cur);
  while (true) {
    const dirEnt = dir.readSync()
    if (dirEnt === null) {
      break;
    }
    if (dirEnt.isFile() && dirEnt.name === 'go.mod') {
      console.log('go test ' + cur);
      try {
        execSync('go test ./...', { cwd: cur, encoding: 'ascii' });
      } catch (err) {
        console.error(err);
      }
    } else if (dirEnt.isDirectory()) {
      recursiveFindGoMod(path.join(cur, dirEnt.name));
    }
  }
  dir.close();
}
