// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
import { exec, execSync } from 'child_process';
import { existsSync, opendirSync, unlinkSync, readFileSync, writeFileSync } from 'fs';
import { semaphore } from '../../../.scripts/semaphore.js';

// limit to 8 concurrent builds
const sem = semaphore(8);

const pkgRoot = execSync('git rev-parse --show-toplevel').toString().trim() + '/packages/typespec-go/';

const httpSpecs = pkgRoot + 'node_modules/@typespec/http-specs/specs/';
const azureHttpSpecs = pkgRoot + 'node_modules/@azure-tools/azure-http-specs/specs/';

const compiler = pkgRoot + 'node_modules/@typespec/compiler/cmd/tsp.js';

// the format is as follows
// 'moduleName': [ 'input', 'emitter option 1', 'emitter option N...' ]
// if no .tsp file is specified in input, it's assumed to be main.tsp
const httpSpecsGroup = {
  'apikeygroup': ['authentication/api-key'],     // missing tests
  'customgroup': ['authentication/http/custom'], // missing tests
  'oauth2group': ['authentication/oauth2'],      // missing tests
  'unionauthgroup': ['authentication/union'],    // missing tests
  'bytesgroup': ['encode/bytes'],
  'datetimegroup': ['encode/datetime', 'slice-elements-byval=true'],
  'durationgroup': ['encode/duration'],
  'numericgroup': ['encode/numeric'],
  'basicparamsgroup': ['parameters/basic'],
  'bodyoptionalgroup': ['parameters/body-optionality'],
  'collectionfmtgroup': ['parameters/collection-format'],
  'spreadgroup': ['parameters/spread'],
  'contentneggroup': ['payload/content-negotiation'],
  'jmergepatchgroup': ['payload/json-merge-patch'],
  'mediatypegroup': ['payload/media-type'],
  //'multipartgroup': ['payload/multipart'], // TODO: https://github.com/Azure/autorest.go/issues/1445
  'pageablegroup': ['payload/pageable'],
  'xmlgroup': ['payload/xml', 'slice-elements-byval=true'],
  'jsongroup': ['serialization/encoded-name/json'],
  'noendpointgroup': ['server/endpoint/not-defined'],
  'multiplegroup': ['server/path/multiple'],
  'singlegroup': ['server/path/single'],
  'unversionedgroup': ['server/versions/not-versioned'],
  'versionedgroup': ['server/versions/versioned'],
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
  //'addlpropsgroup': ['type/property/additional-properties'], // requires union support
  'nullablegroup': ['type/property/nullable'],
  'optionalitygroup': ['type/property/optionality', 'slice-elements-byval=true'],
  'valuetypesgroup': ['type/property/value-types', 'slice-elements-byval=true'],
  'scalargroup': ['type/scalar', 'slice-elements-byval=true'],
  //'uniongroup': ['type/union'], // requires union support
  //'addedgroup': ['versioning/added'], // requires union support
  'madeoptionalgroup': ['versioning/madeOptional'],
  //'removedgroup': ['versioning/removed'], // requires union support
  //'renamedfromgroup': ['versioning/renamedFrom'], // requires union support
  'rettypechangedfromgroup': ['versioning/returnTypeChangedFrom'],
  'typechangedfromgroup': ['versioning/typeChangedFrom'],
  'jsonlgroup': ['streaming/jsonl']
};

const azureHttpSpecsGroup = {
  'accessgroup': ['azure/client-generator-core/access'],
  'flattengroup': ['azure/client-generator-core/flatten-property'],
  'coreusagegroup': ['azure/client-generator-core/usage'],
  // 'clientinitializationgroup': ['azure/client-generator-core/client-initialization'],
  // 'apiversionheadergroup' : ['azure/client-generator-core/api-version/header'],
  // 'apiversionpathgroup' : ['azure/client-generator-core/api-version/path'],
  // 'apiversionquerygroup' : ['azure/client-generator-core/api-version/query'],
  'basicgroup': ['azure/core/basic'],
  'lrorpcgroup': ['azure/core/lro/rpc'],
  'lrostdgroup': ['azure/core/lro/standard'],
  'azurepagegroup': ['azure/core/page/client.tsp'], // requires paging with re-injection support
  'corescalargroup': ['azure/core/scalar'],
  'coremodelgroup': ['azure/core/model'],
  //'traitsgroup': ['azure/core/traits'], // requires union support
  'encodedurationgroup': ['azure/encode/duration'],
  'examplebasicgroup': ['azure/example/basic'],
  'pageablegroup': ['azure/payload/pageable'],
  'commonpropsgroup': ['azure/resource-manager/common-properties'],
  'resources': ['azure/resource-manager/resources', 'factory-gather-all-params=false'],
  'nonresourcegroup' : ['azure/resource-manager/non-resource'],
  'templatesgroup' : ['azure/resource-manager/operation-templates'],
  'xmsclientreqidgroup': ['azure/special-headers/client-request-id'],
  'naminggroup': ['client/naming'],
  'defaultgroup': ['client/structure/default/client.tsp'],
  'multiclientgroup': ['client/structure/multi-client/client.tsp'],
  'renamedopgroup': ['client/structure/renamed-operation/client.tsp'],
  'clientopgroup': ['client/structure/client-operation-group/client.tsp'],
  'clientnamespacegroup': ['client/namespace'],
  'twoopgroup': ['client/structure/two-operation-group/client.tsp'],
  'srvdrivenoldgroup': ['resiliency/srv-driven/old.tsp'],
  'srvdrivennewgroup': ['resiliency/srv-driven'],
};

// default to using the locally built emitter
let emitter = pkgRoot;
const args = process.argv.slice(2);
var filter = undefined;
const switches = [];
for (var i = 0 ; i < args.length; i += 1) {
  const filterArg = args[i].match(/--filter=(?<filter>\w+)/);
  if (filterArg) {
    filter = filterArg.groups['filter'];
    continue;
  }

  switch (args[i]) {
    case '--verbose':
      switches.push('--verbose');
      break;
    case '--emitter-installed':
      // the emitter has been installed so use that one instead
      emitter = '@azure-tools/typespec-go';
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
generate('armcodesigning', armcodesigning, 'test/local/armcodesigning', [`examples-directory=${armcodesigning}/examples`, 'generate-samples=true']);

const armapicenter = pkgRoot +  'test/tsp/ApiCenter.Management';
generate('armapicenter', armapicenter, 'test/local/armapicenter', [`examples-directory=${armapicenter}/examples`, 'generate-samples=true']);

const armlargeinstance = pkgRoot + 'test/tsp/AzureLargeInstance.Management';
generate('armlargeinstance', armlargeinstance, 'test/local/armlargeinstance', ['stutter=AzureLargeInstance', `examples-directory=${armlargeinstance}/examples`, 'generate-samples=true']);

const armdatabasewatcher = pkgRoot + 'test/tsp/DatabaseWatcher.Management';
generate('armdatabasewatcher', armdatabasewatcher, 'test/local/armdatabasewatcher', ['fix-const-stuttering=false', `examples-directory=${armdatabasewatcher}/examples`, 'generate-samples=true']);

const armloadtestservice = pkgRoot + 'test/tsp/LoadTestService.Management';
generate('armloadtestservice', armloadtestservice, 'test/local/armloadtestservice', [`examples-directory=${armloadtestservice}/examples`, 'generate-samples=true', 'factory-gather-all-params=false']);

const armdevopsinfrastructure = pkgRoot + 'test/tsp/Microsoft.DevOpsInfrastructure';
generate('armdevopsinfrastructure', armdevopsinfrastructure, 'test/local/armdevopsinfrastructure', [`examples-directory=${armdevopsinfrastructure}/examples`, 'generate-samples=true']);

const armrandom = pkgRoot + 'test/tsp/Random.Management';
generate('armrandom', armrandom, 'test/local/armrandom');

const armcommunitymanagement = pkgRoot + 'test/tsp/Community.Management';
generate('armcommunitymanagement', armcommunitymanagement, 'test/local/armcommunitymanagement', [`examples-directory=${armcommunitymanagement}/examples`, 'generate-samples=true']);

const armmongocluster = pkgRoot + 'test/tsp/MongoCluster.Management';
generate('armmongocluster', armmongocluster, 'test/local/armmongocluster', [`examples-directory=${armmongocluster}/examples`, 'generate-samples=true']);

const armcontainerorchestratorruntime = pkgRoot + 'test/tsp/KubernetesRuntime.Management';
generate('armcontainerorchestratorruntime', armcontainerorchestratorruntime, 'test/local/armcontainerorchestratorruntime', [`examples-directory=${armcontainerorchestratorruntime}/examples`, 'generate-samples=true']);

const azmodelsonly = pkgRoot + 'test/tsp/ModelsOnlyWithBaseTypes';
generate('azmodelsonly', azmodelsonly, 'test/local/azmodelsonly');

const azkeys = pkgRoot + 'test/tsp/KeyVault.Keys/client.tsp';
generate('azkeys', azkeys, 'test/local/azkeys', ['single-client=true']);

const armtest = pkgRoot + 'test/tsp/Test.Management';
generate('armtest', armtest, 'test/local/armtest');

const internalpager = pkgRoot + 'test/tsp/Internal.Pager';
generate('internalpager', internalpager, 'test/local/internalpager', ['generate-fakes=false']);

const armoracledatabase = pkgRoot + 'test/tsp/Oracle.Database.Management';
generate('armoracledatabase', armoracledatabase, 'test/local/armoracledatabase', [`examples-directory=${armoracledatabase}/examples`, 'generate-samples=true', 'module-version=2.0.0']);

const armhealthbot = pkgRoot + 'test/tsp/Healthbot.Management';
generate('armhealthbot', armhealthbot, 'test/local/armhealthbot', [`examples-directory=${armhealthbot}/examples`, 'generate-samples=true', 'module-version=1.0.0']);

const armhardwaresecuritymodules = pkgRoot + 'test/tsp/HardwareSecurityModules.Management';
generate('armhardwaresecuritymodules', armhardwaresecuritymodules, 'test/local/armhardwaresecuritymodules', [`examples-directory=${armhardwaresecuritymodules}/examples`, 'generate-samples=true']);

const armcomputeschedule = pkgRoot + 'test/tsp/ComputeSchedule.Management';
generate('armcomputeschedule', armcomputeschedule, 'test/local/armcomputeschedule', [`examples-directory=${armcomputeschedule}/examples`, 'generate-samples=true']);

const armbillingbenefits = pkgRoot + 'test/tsp/BillingBenefits.Management';
generate('armbillingbenefits', armbillingbenefits, 'test/local/armbillingbenefits', [`examples-directory=${armbillingbenefits}/examples`, 'generate-samples=true']);

const nooptionalbody = pkgRoot + 'test/tsp/NoOptionalBody';
generate('nooptionalbody', nooptionalbody, 'test/local/nooptionalbody', ['generate-fakes=false', 'go-generate=after_generate.go', 'no-optional-body=true']);

loopSpec(httpSpecsGroup, httpSpecs, 'test/http-specs')
loopSpec(azureHttpSpecsGroup, azureHttpSpecs, 'test/azure-http-specs')

function loopSpec(group, root, prefix) {
  for (const module in group) {
    const values = group[module];
    let perTestOptions;
    if (values.length > 1) {
      perTestOptions = values.slice(1);
    }
    // keep the output directory structure similar to the cadl input directory.
    // remove the last dir from the input path as we'll use the module name instead.
    // if the input specifies a .tsp file, remove that first.
    let outDir = values[0];
    if (outDir.lastIndexOf('.tsp') > -1) {
      outDir = outDir.substring(0, outDir.lastIndexOf('/'));
    }
    outDir = outDir.substring(0, outDir.lastIndexOf('/'));
    generate(module, root + values[0], `${prefix}/${outDir}/` + module, perTestOptions);
  }
}

function generate(moduleName, input, outputDir, perTestOptions) {
  if (!should_generate(moduleName)) {
    return
  }
  if (perTestOptions === undefined) {
    perTestOptions = [];
  }

  const fullOutputDir = pkgRoot + outputDir;

  // these options can't be changed per test
  const fixedOptions = [
    `module=${moduleName}`,
    `emitter-output-dir=${fullOutputDir}`,
    'file-prefix=zz_',
  ];

  // these options _can_ be changed per test
  // TODO: disabled examples by default https://github.com/Azure/autorest.go/issues/1441
  const defaultOptions = [
    'generate-fakes=true',
    'inject-spans=true',
    'head-as-boolean=true',
    'fix-const-stuttering=true',
  ];

  let allOptions = fixedOptions;

  // merge in any per-test options.
  // if a per-test option overlaps with a default option, use the per-test one.
  for (const perTestOption of perTestOptions) {
    // perTestOption === 'option=value', grab the option name
    const optionName = perTestOption.match(/^([a-zA-Z0-9_-]+)/)[0];
    const index = defaultOptions.findIndex((value, index, obj) => {
      return value.startsWith(optionName);
    });
    if (index > -1) {
      // found a match, replace the default option with the per-test one
      defaultOptions[index] = perTestOption;
    } else {
      allOptions.push(perTestOption);
    }
  }

  allOptions = allOptions.concat(defaultOptions);

  sem.take(function() {
    // default to main.tsp if a .tsp file isn't specified in the input
    if (input.lastIndexOf('.tsp') === -1) {
      input += '/main.tsp';
    }
    console.log('generating ' + input);
    try {
      const options = [];
      for (const option of allOptions) {
        options.push(`--option="@azure-tools/typespec-go.${option}"`);
      }
      if (switches.includes('--debugger')) {
        options.push(`--option="@azure-tools/typespec-go.debugger=true"`);
      }
      const command = `node ${compiler} compile ${input} --emit=${emitter} ${options.join(' ')}`;
      if (switches.includes('--verbose')) {
        console.log(command);
      }
      cleanGeneratedFiles(fullOutputDir);
      exec(command, function(error, stdout, stderr) {
        // print any output or error from the tsp compile command
        logResult(error, stdout, stderr);
        // format on success
        if (error === null) {
          // Force emitter version to a constant in _metadata.json to avoid unnecessary version drift in committed files
          const metadataPath = `${fullOutputDir}/testdata/_metadata.json`;
          if (existsSync(metadataPath)) {
            const metadata = JSON.parse(readFileSync(metadataPath, 'utf8'));
            metadata.emitterVersion = '0.0.0';
            writeFileSync(metadataPath, JSON.stringify(metadata, null, 2));
          }
        }
      });
    } catch (err) {
      console.error('An error occurred:');  
      if (err.message) {
        console.error('Message:', err.message);  
      }
      if (err.stack) {
        console.error('Stack:', err.stack);  
      }
      if (err.output) {
        console.error('Output:', err.output.toString());  
      }
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
  // typespec compiler prints compiler progress to stderr
  // but it's not an error, so we use console.log
  // to print it out.
  if (stderr !== '') {
    console.log('stderr: ' + stderr);
  }
  if (error !== null) {
    console.error('\x1b[91m%s\x1b[0m', 'exec error: ' + error);
  }
}
