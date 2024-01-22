// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
import { exec, execSync } from 'child_process';
import { existsSync, opendirSync, unlinkSync } from 'fs';
import semaphore from '../../../.scripts/semaphore.js';

// limit to 8 concurrent builds
const sem = semaphore(8);

const pkgRoot = execSync('git rev-parse --show-toplevel').toString().trim() + '/packages/typespec-go/';

const tspRoot = pkgRoot + 'node_modules/@azure-tools/cadl-ranch-specs/http/';

const compiler = pkgRoot + 'node_modules/@typespec/compiler/node_modules/.bin/tsp';

// the format is as follows
// 'moduleName': [ 'inputDir', 'additional arg 1', 'additional arg N...' ]
const cadlRanch = {
  'apikeygroup': ['authentication/api-key'],
  'customgroup': ['authentication/http/custom'],
  'oauth2group': ['authentication/oauth2'],
  'unionauthgroup': ['authentication/union'],
  'accessgroup': ['azure/client-generator-core/access'],
  'coreusagegroup': ['azure/client-generator-core/usage'],
  'basicgroup': ['azure/core/basic'],
  //'lrorpcgroup': ['azure/core/lro/rpc'],           // requires lro support
  //'lrolegacygroup': ['azure/core/lro/rpc-legacy'], // requires lro support
  //'lrostdgroup': ['azure/core/lro/standard'],      // requires lro support
  //'traitsgroup': ['azure/core/traits'],            // requires union support
  'defaultgroup': ['client/structure/default'],
  'multiclientgroup': ['client/structure/multi-client'],
  'renamedopgroup': ['client/structure/renamed-operation'],
  'twoopgroup': ['client/structure/two-operation-group'],
  'bytesgroup': ['encode/bytes'],
  'datetimegroup': ['encode/datetime', 'slice-elements-byval=true'],
  'durationgroup': ['encode/duration'],
  'bodyoptionalgroup': ['parameters/body-optionality'],
  'collectionfmtgroup': ['parameters/collection-format'],
  //'spreadgroup': ['parameters/spread'], // needs more investigation
  //'contentneggroup': ['payload/content-negotiation'], // https://github.com/Azure/typespec-azure/issues/107
  'pageablegroup': ['payload/pageable'],
  'projectednamegroup': ['projection/projected-name'],
  'srvdrivengroup': ['resiliency/srv-driven'],
  'multiplegroup': ['server/path/multiple'],
  'singlegroup': ['server/path/single'],
  'clientreqidgroup': ['special-headers/client-request-id'],
  'condreqgroup': ['special-headers/conditional-request'],
  //'repeatabilitygroup': ['special-headers/repeatability'],   // requires union support
  'specialwordsgroup': ['special-words'],
  'arraygroup': ['type/array', 'slice-elements-byval=true'],
  'dictionarygroup': ['type/dictionary'],
  'extensiblegroup': ['type/enum/extensible'],
  'fixedgroup': ['type/enum/fixed'],
  'emptygroup': ['type/model/empty'],
  'enumdiscgroup': ['type/model/inheritance/enum-discriminator'],
  //'nesteddiscgroup': ['type/model/inheritance/nested-discriminator'], // not a real scenario
  'nodiscgroup': ['type/model/inheritance/not-discriminated'],
  'recursivegroup': ['type/model/inheritance/recursive', 'slice-elements-byval=true'],
  'singlediscgroup': ['type/model/inheritance/single-discriminator'],
  'usagegroup': ['type/model/usage'],
  'visibilitygroup': ['type/model/visibility'],
  'addlpropsgroup': ['type/property/additional-properties'],
  'nullablegroup': ['type/property/nullable'],
  //'optionalitygroup': ['type/property/optionality'], // requires union support
  //'valuetypesgroup': ['type/property/value-types'], // requires decimal support
  //'scalargroup': ['type/scalar'],                   // requires decimal support
  //'uniongroup': ['type/union'], // requires union support
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
  if (values.length > 1) {
    additionalArgs = values.slice(1);
  }
  // keep the output directory structure similar to the cadl input directory.
  // remove the last dir from the input path as we'll use the module name instead
  const shortend = values[0].substring(0, values[0].lastIndexOf('/'));
  generate(module, tspRoot + values[0], `test/cadlranch/${shortend}/` + module, additionalArgs);
}

function generate(moduleName, inputDir, outputDir, additionalArgs) {
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
      options.push(`--option="@azure-tools/typespec-go.emitter-output-dir=${fullOutputDir}"`);
      options.push(`--option="@azure-tools/typespec-go.file-prefix=zz_"`);
      if (switches.includes('--debugger')) {
        options.push(`--option="@azure-tools/typespec-go.debugger=true"`);
      }
      const command = `${compiler} compile ${inputDir}/main.tsp --emit=${pkgRoot} ${options.join(' ')} ${additionalArgs.join(' ')}`;
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
      });
    } catch (err) {
      console.error(err.output.toString());
    } finally {
      sem.leave();
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
