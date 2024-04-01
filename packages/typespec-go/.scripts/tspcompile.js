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
// 'moduleName': [ 'input', 'additional arg 1', 'additional arg N...' ]
// if no .tsp file is specified in input, it's assumed to be main.tsp
const cadlRanch = {
  'apikeygroup': ['authentication/api-key'],     // missing tests
  'customgroup': ['authentication/http/custom'], // missing tests
  'oauth2group': ['authentication/oauth2'],      // missing tests
  'unionauthgroup': ['authentication/union'],    // missing tests
  'accessgroup': ['azure/client-generator-core/access'],
  'flattengroup': ['azure/client-generator-core/flatten'],
  'coreusagegroup': ['azure/client-generator-core/usage'],
  'basicgroup': ['azure/core/basic'],
  'lrorpcgroup': ['azure/core/lro/rpc'],
  'lrolegacygroup': ['azure/core/lro/rpc-legacy'],
  'lrostdgroup': ['azure/core/lro/standard'],
  'corescalargroup': ['azure/core/scalar'],
  //'traitsgroup': ['azure/core/traits'], // requires union support
  'naminggroup': ['client/naming'],
  'defaultgroup': ['client/structure/default/client.tsp'],
  'multiclientgroup': ['client/structure/multi-client/client.tsp'],
  'renamedopgroup': ['client/structure/renamed-operation/client.tsp'],
  'twoopgroup': ['client/structure/two-operation-group/client.tsp'],
  'bytesgroup': ['encode/bytes'],
  'datetimegroup': ['encode/datetime', 'slice-elements-byval=true'],
  'durationgroup': ['encode/duration'],
  'bodyoptionalgroup': ['parameters/body-optionality'],
  'collectionfmtgroup': ['parameters/collection-format'],
  //'spreadgroup': ['parameters/spread'], // needs more investigation
  'contentneggroup': ['payload/content-negotiation'],
  'jmergepatchgroup': ['payload/json-merge-patch'],
  'mediatypegroup': ['payload/media-type'],
  'multipartgroup': ['payload/multipart'],
  'pageablegroup': ['payload/pageable'],
  'projectednamegroup': ['projection/projected-name'],
  'srvdrivengroup': ['resiliency/srv-driven'], // missing tests
  'jsongroup': ['serialization/encoded-name/json'],
  'multiplegroup': ['server/path/multiple'],
  'singlegroup': ['server/path/single'],
  'unversionedgroup': ['server/versions/not-versioned'],
  'versionedgroup': ['server/versions/versioned'],
  'clientreqidgroup': ['special-headers/client-request-id'],
  'condreqgroup': ['special-headers/conditional-request'],
  //'repeatabilitygroup': ['special-headers/repeatability'],   // requires union support
  'specialwordsgroup': ['special-words'],
  'arraygroup': ['type/array', 'slice-elements-byval=true'],
  'dictionarygroup': ['type/dictionary'],
  'extensiblegroup': ['type/enum/extensible'],
  'fixedgroup': ['type/enum/fixed'],
  'emptygroup': ['type/model/empty', 'single-client=true'],
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
  //'valuetypesgroup': ['type/property/value-types'], // requires union support
  'scalargroup': ['type/scalar', 'slice-elements-byval=true'],
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

const armcodesigning = pkgRoot + 'test/tsp/CodeSigning.Management';
generate('armcodesigning', armcodesigning, 'test/armcodesigning');

const armapicenter = pkgRoot +  'test/tsp/ApiCenter.Management';
generate('armapicenter', armapicenter, 'test/armapicenter');

const armlargeinstance = pkgRoot + 'test/tsp/AzureLargeInstance.Management';
generate('armlargeinstance', armlargeinstance, 'test/armlargeinstance', ['stutter=AzureLargeInstance']);

for (const module in cadlRanch) {
  const values = cadlRanch[module];
  let additionalArgs;
  if (values.length > 1) {
    additionalArgs = values.slice(1);
  }
  // keep the output directory structure similar to the cadl input directory.
  // remove the last dir from the input path as we'll use the module name instead.
  // if the input specifies a .tsp file, remove that first.
  let outDir = values[0];
  if (outDir.lastIndexOf('.tsp') > -1) {
    outDir = outDir.substring(0, outDir.lastIndexOf('/'));
  }
  outDir = outDir.substring(0, outDir.lastIndexOf('/'));
  generate(module, tspRoot + values[0], `test/cadlranch/${outDir}/` + module, additionalArgs);
}

function generate(moduleName, input, outputDir, additionalArgs) {
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
    // default to main.tsp if a .tsp file isn't specified in the input
    if (input.lastIndexOf('.tsp') === -1) {
      input += '/main.tsp';
    }
    console.log('generating ' + input);
    const fullOutputDir = pkgRoot + outputDir;
    try {
      const options = [];
      options.push(`--option="@azure-tools/typespec-go.module=${moduleName}"`);
      options.push(`--option="@azure-tools/typespec-go.module-version=0.1.0"`);
      options.push(`--option="@azure-tools/typespec-go.emitter-output-dir=${fullOutputDir}"`);
      options.push(`--option="@azure-tools/typespec-go.file-prefix=zz_"`);
      options.push(`--option="@azure-tools/typespec-go.generate-fakes=true"`);
      options.push(`--option="@azure-tools/typespec-go.inject-spans=true"`);
      if (switches.includes('--debugger')) {
        options.push(`--option="@azure-tools/typespec-go.debugger=true"`);
      }
      const command = `${compiler} compile ${input} --emit=${pkgRoot} ${options.join(' ')} ${additionalArgs.join(' ')}`;
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
