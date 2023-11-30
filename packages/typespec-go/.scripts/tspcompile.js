// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
import { exec, execSync } from 'child_process';

import { existsSync, opendirSync, unlinkSync } from 'fs';

import { semaphore } from './semaphore.js';

// limit to 8 concurrent builds
const sem = semaphore(8);

const pkgRoot = execSync('git rev-parse --show-toplevel').toString().trim() + '/packages/typespec-go/';

const tspRoot = pkgRoot + 'node_modules/@azure-tools/cadl-ranch-specs/http/';

// the format is as follows
// 'moduleName': [ 'moduleVersion', 'inputDir', 'additional arg 1', 'additional arg N...' ]
const cadlRanch = {
  'arraygroup': ['0.1.1', 'type/array', 'slice-elements-byval=true'],
  'dictionarygroup': ['0.1.1', 'type/dictionary'],
  'extensibleenumgroup': ['0.1.1', 'type/enum/extensible'],
  //'singlediscriminatorgroup': ['0.1.1', 'type/model/inheritance/single-discriminator'],
  'visibilitygroup': ['0.1.1', 'type/model/visibility']
};

// any new args must also be added to autorest.go\common\config\rush\command-line.json
const args = process.argv.slice(2);
var filter = undefined;
const switches = [];
for (var i = 0 ; i < args.length; i += 1) {
  switch (args[i]) {
    case '--debugger':
      switches.push('--debugger');
      break;
    case '--filter':
      filter = args[i + 1]
      i += 1
      break;
    case '--verbose':
      switches.push('--verbose');
      break;
    default:
      break;
  }
}

if (filter !== undefined) {
  console.log("Using filter: " + filter)
}

function should_generate(name) {
  if (filter !== undefined) {
    const re = new RegExp(filter);
    return re.test(name)
  }
  return true
}

for (const module in cadlRanch) {
  const values = cadlRanch[module];
  let additionalArgs;
  if (values.length > 2) {
    additionalArgs = values.slice(2);
  }
  generate(module, values[0], tspRoot + values[1], 'test/cadlranch/' + module, additionalArgs);
}

function generate(moduleName, moduleVersion, inputDir, outputDir, additionalArgs) {
  if (!should_generate(moduleName)) {
    return
  }
  if (additionalArgs === undefined) {
    additionalArgs = [];
  } else {
    for (let i = 0; i < additionalArgs.length; ++i) {
      additionalArgs[i] = `--option="@azure-tools/typespec-go.${additionalArgs[i]}"`;
    }
  }
  sem.take(function() {
    console.log('generating ' + inputDir);
    const fullOutputDir = pkgRoot + outputDir;
    try {
      const options = [];
      options.push(`--option="@azure-tools/typespec-go.module=${moduleName}"`);
      options.push(`--option="@azure-tools/typespec-go.module-version=${moduleVersion}"`);
      options.push(`--option="@azure-tools/typespec-go.emitter-output-dir=${fullOutputDir}"`);
      options.push(`--option="@azure-tools/typespec-go.file-prefix=zz_"`);
      if (switches.includes('--debugger')) {
        options.push(`--option="@azure-tools/typespec-go.debugger=true"`);
      }
      const command = `tsp compile ${inputDir}/main.tsp --emit=${pkgRoot} ${options.join(' ')} ${additionalArgs.join(' ')}`;
      if (switches.includes('--verbose')) {
        console.log(command);
      }
      cleanGeneratedFiles(fullOutputDir);
      exec(command, function(error, stdout, stderr) {
        // print any output or error from the tsp compile command
        logResult(error, stdout, stderr);
        // format on success
        if (error === null && stderr === '') {
          execSync('gofmt -w .', { cwd: fullOutputDir});
        }
        sem.leave();
      });
    } catch (err) {
      console.error(err.output.toString());
    }
  });
}

function cleanGeneratedFiles(outputDir) {
  if (!existsSync(outputDir)) {
      return;
  }
  const dir = opendirSync(outputDir);
  while (true) {
      const dirEnt = dir.readSync()
      if (dirEnt === null) {
          break;
      }
      if (dirEnt.isFile() && dirEnt.name.startsWith('zz_')) {
          unlinkSync(dir.path + '/' + dirEnt.name);
      }
  }
  dir.close();
  cleanGeneratedFiles(outputDir + '/fake');
}

function logResult(error, stdout, stderr) {
  if (stdout !== '') {
    console.log('stdout: ' + stdout);
  }
  if (stderr !== '') {
    console.error('\x1b[91m%s\x1b[0m', 'stderr: ' + stderr);
  }
  if (error !== null) {
    console.error('\x1b[91m%s\x1b[0m', 'exec error: ' + error);
  }
}
