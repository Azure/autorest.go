// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
import { execSync } from 'child_process';
import * as fs from 'fs';

recursiveFindGoMod('./packages/autorest.go/test');

function recursiveFindGoMod(cur) {
  const dir = fs.opendirSync(cur);
  while (true) {
    const dirEnt = dir.readSync()
    if (dirEnt === null) {
      break;
    }
    if (dirEnt.isFile() && dirEnt.name === 'go.mod') {
      console.log('go build && go vet ' + cur);
      try {
        execSync('go build ./...', { cwd: cur });
        execSync('go vet ./...', { cwd: cur });
      } catch (err) {
        console.error(err);
      }
    } else if (dirEnt.isDirectory()) {
      recursiveFindGoMod(cur + '/' + dirEnt.name);
    }
  }
  dir.close();
}
