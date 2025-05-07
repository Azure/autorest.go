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

recursiveFindGoMod(dir);

function recursiveFindGoMod(cur) {
  const dir = opendirSync(cur);
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
      recursiveFindGoMod(join(cur, dirEnt.name));
    }
  }
  dir.close();
}
