// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
import { exec } from 'child_process';
import { execSync } from 'child_process';
import * as fs from 'fs';
import semaphore from '../../../.scripts/semaphore.js';

// limit to 8 concurrent builds
const sem = semaphore(8);

const swaggerDir = 'packages/autorest.go/node_modules/@microsoft.azure/autorest.testserver/swagger/';

const goMappings = {
  'additionalpropsgroup': ['additionalProperties.json', '--remove-unreferenced-types', '--disallow-unknown-fields'],
  'arraygroup': ['body-array.json', '--remove-unreferenced-types'],
  'azurereportgroup': ['azure-report.json', '--remove-unreferenced-types'],
  'azurespecialsgroup': ['azure-special-properties.json', '--head-as-boolean', '--remove-unreferenced-types'],
  'binarygroup': ['body-binary.json', '--remove-unreferenced-types'],
  'booleangroup': ['body-boolean.json', '--remove-unreferenced-types'],
  'bytegroup': ['body-byte.json', '--remove-unreferenced-types'],
  'complexgroup': ['body-complex.json', '--remove-unreferenced-types', '--rawjson-as-bytes'],
  'complexmodelgroup': ['complex-model.json', '--remove-unreferenced-types'],
  'custombaseurlgroup': ['custom-baseUrl.json', '--remove-unreferenced-types'],
  'dategroup': ['body-date.json', '--remove-unreferenced-types'],
  'datetimegroup': ['body-datetime.json', '--remove-unreferenced-types'],
  'datetimerfc1123group': ['body-datetime-rfc1123.json', '--remove-unreferenced-types'],
  'dictionarygroup': ['body-dictionary.json', '--remove-unreferenced-types'],
  'durationgroup': ['body-duration.json', '--remove-unreferenced-types'],
  'errorsgroup': ['xms-error-responses.json', '--remove-unreferenced-types'],
  'extenumsgroup': ['extensible-enums-swagger.json', '--remove-unreferenced-types'],
  'filegroup': ['body-file.json', '--remove-unreferenced-types'],
  'formdatagroup': ['body-formdata.json', '--remove-unreferenced-types'],
  'headergroup': ['header.json', '--remove-unreferenced-types'],
  'headgroup': ['head.json', '--head-as-boolean', '--remove-unreferenced-types'],
  'httpinfrastructuregroup': ['httpInfrastructure.json', '--head-as-boolean', '--remove-unreferenced-types'],
  'integergroup': ['body-integer.json', '--remove-unreferenced-types'],
  'lrogroup': ['lro.json', '--remove-unreferenced-types'],
  'mediatypesgroup': ['media_types.json', '--remove-unreferenced-types'],
  'mediatypesgroupwithnormailzedoperationname': ['media_types.json', '--remove-unreferenced-types', '--normalize-operation-name=true'],
  'migroup': ['multiple-inheritance.json', '--remove-unreferenced-types'],
  //'modelflatteninggroup': ['model-flattening.json'],
  'morecustombaseurigroup': ['custom-baseUrl-more-options.json', '--remove-unreferenced-types'],
  'nonstringenumgroup': ['non-string-enum.json', '--remove-unreferenced-types'],
  'noopsgroup': ['no-operations.json'],
  'numbergroup': ['body-number.json', '--remove-unreferenced-types'],
  'objectgroup': ['object-type.json', '--remove-unreferenced-types', '--rawjson-as-bytes'],
  'optionalgroup': ['required-optional.json', '--remove-unreferenced-types --honor-body-placement'],
  'paginggroup': ['paging.json', '--remove-unreferenced-types'],
  'paramgroupinggroup': ['azure-parameter-grouping.json', '--remove-unreferenced-types'],
  'reportgroup': ['report.json', '--remove-unreferenced-types'],
  'stringgroup': ['body-string.json', '--remove-unreferenced-types'],
  'urlgroup': ['url.json', '--remove-unreferenced-types'],
  'urlmultigroup': ['url-multi-collectionFormat.json', '--remove-unreferenced-types'],
  'validationgroup': ['validation.json', '--remove-unreferenced-types'],
  'xmlgroup': ['xml-service.json', '--remove-unreferenced-types'],
};

// any new args must also be added to autorest.go\common\config\rush\command-line.json
const args = process.argv.slice(2);
var filter = undefined;
const switches = [];
for (var i = 0 ; i < args.length; i += 1) {
  switch (args[i]) {
    case '--filter':
      filter = args[i + 1]
      i += 1
      break;
    case '--verbose':
      switches.push('--debug');
      break;
    case '--debugger':
      switches.push('--go.debugger');
      break;
    case '--dump-code-model':
      switches.push('--output-artifact:code-model-v4');
      break;
    default:
      break;
  }
}

if (filter !== undefined) {
  console.log("Using filter: " + filter)
}

// loop through all of the namespaces in goMappings
for (const namespace in goMappings) {
  // for each swagger run the autorest command to generate code based on the swagger for the relevant namespace and output to the /generated directory
  const entry = goMappings[namespace];
  const inputFile = swaggerDir + entry[0];
  let extraParams = ['--export-clients --containing-module=generatortests'];
  if (entry.length > 1) {
    extraParams = extraParams.concat(entry.slice());
  }
  generate(namespace, inputFile, 'test/autorest/' + namespace, extraParams.join(' '));
}

const blobStorage = './swagger/specification/storage/data-plane/Microsoft.BlobStorage/readme.md';
generateFromReadme("azblob", blobStorage, 'package-2021-12', 'test/storage/azblob', '--module-version=0.1.0 --inject-spans');

const network = './swagger/specification/network/resource-manager/readme.md';
generateFromReadme("armnetwork", network, 'package-2022-09', 'test/network/armnetwork', '--module=armnetwork --module-version=0.1.0 --azure-arm=true --remove-unreferenced-types');

const compute = './swagger/specification/compute/resource-manager/readme.md';
generateFromReadme("armcompute", compute, 'package-2021-12-01', 'test/compute/armcompute', '--module=armcompute --module-version=0.1.0 --azure-arm=true --remove-unreferenced-types --slice-elements-byval');

const synapseArtifacts = './swagger/specification/synapse/data-plane/readme.md';
generateFromReadme("azartifacts", synapseArtifacts, 'package-artifacts-composite-v6', 'test/synapse/azartifacts', '--security=AADToken --security-scopes="https://dev.azuresynapse.net/.default" --module="azartifacts" --module-version=0.1.0 --openapi-type="data-plane"');

const synapseSpark = './swagger/specification/synapse/data-plane/readme.md';
generateFromReadme("azspark", synapseSpark, 'package-spark-2020-12-01', 'test/synapse/azspark', '--security=AADToken --security-scopes="https://dev.azuresynapse.net/.default" --module="azspark" --module-version=0.1.0 --openapi-type="data-plane"');

const tables = './swagger/specification/cosmos-db/data-plane/readme.md';
generateFromReadme("aztables", tables, 'package-2019-02', 'test/tables/aztables', '--security=AADToken --security-scopes="https://tables.azure.com/.default" --module=aztables --module-version=0.1.0 --openapi-type="data-plane" --export-clients --azure-validator=false --group-parameters=false --stutter=table --rawjson-as-bytes');

const keyvault = './swagger/specification/keyvault/data-plane/readme.md';
generateFromReadme("azkeyvault", keyvault, 'package-7.2', 'test/keyvault/azkeyvault', '--module=azkeyvault --module-version=0.1.0');

const consumption = './swagger/specification/consumption/resource-manager/readme.md';
generateFromReadme("armconsumption", consumption, 'package-2019-10', 'test/consumption/armconsumption', '--module=armconsumption --module-version=1.0.0 --azure-arm=true --generate-fakes=false --inject-spans=false --remove-unreferenced-types');

const databoxedge = './swagger/specification/databoxedge/resource-manager/readme.md';
generateFromReadme("armdataboxedge", databoxedge, 'package-2021-02-01', 'test/databoxedge/armdataboxedge', '--module=armdataboxedge --module-version=2.0.0 --azure-arm=true --remove-unreferenced-types --inject-spans=false --fix-const-stuttering=true');

const acr = './swagger/specification/containerregistry/data-plane/Azure.ContainerRegistry/stable/2021-07-01/containerregistry.json';
generate("azacr", acr, 'test/acr/azacr', '--module="azacr" --module-version=0.1.0 --openapi-type="data-plane" --rawjson-as-bytes --generate-fakes');

const machineLearning = './swagger/specification/machinelearningservices/resource-manager';
generateFromReadme("armmachinelearning", machineLearning, 'package-2022-02-01-preview', 'test/machinelearning/armmachinelearning', '--module=armmachinelearning --module-version=1.0.0 --azure-arm=true --generate-fakes=false --inject-spans=false --remove-unreferenced-types');

generate("azalias", 'packages/autorest.go/test/swagger/alias.json', 'test/maps/azalias', '--security=AzureKey --module="azalias" --module-version=0.1.0 --openapi-type="data-plane" --generate-fakes --inject-spans --slice-elements-byval --disallow-unknown-fields --single-client');

function should_generate(name) {
  if (filter !== undefined) {
    let re = new RegExp(filter);
    return re.test(name)
  }
  return true
}

function fullPath(outputDir) {
  const root = execSync('git rev-parse --show-toplevel').toString().trim();
  return root + '/packages/autorest.go/' + outputDir;
}

// helper to log the package being generated before invocation
function generate(name, inputFile, outputDir, additionalArgs) {
  if (!should_generate(name)) {
    return
  }
  if (additionalArgs === undefined) {
    additionalArgs = '';
  }
  sem.take(function() {
    console.log('generating ' + inputFile);
    outputDir = fullPath(outputDir);
    cleanGeneratedFiles(outputDir);
    exec('autorest --use=./packages/autorest.go --file-prefix="zz_" --modelerfour.lenient-model-deduplication --license-header=MICROSOFT_MIT_NO_VERSION --generate-fakes --inject-spans --input-file=' + inputFile + ' --output-folder=' + outputDir + ' ' + additionalArgs + ' ' + switches.join(' '), autorestCallback(outputDir, inputFile));
  });
}

function generateFromReadme(name, readme, tag, outputDir, additionalArgs) {
  if (!should_generate(name)) {
    return
  }
  if (additionalArgs === undefined) {
    additionalArgs = '';
  }
  sem.take(function() {
    console.log('generating ' + readme);
    outputDir = fullPath(outputDir);
    cleanGeneratedFiles(outputDir);
    exec('autorest --use=./packages/autorest.go ' + readme + ' --go --tag=' + tag + ' --file-prefix="zz_" --modelerfour.lenient-model-deduplication --license-header=MICROSOFT_MIT_NO_VERSION --output-folder=' + outputDir + ' ' + additionalArgs + ' ' + switches.join(' '), autorestCallback(outputDir, readme));
  });
}

function cleanGeneratedFiles(outputDir) {
  if (!fs.existsSync(outputDir)) {
    return;
  }
  const dir = fs.opendirSync(outputDir);
  while (true) {
    const dirEnt = dir.readSync()
    if (dirEnt === null) {
      break;
    }
    if (dirEnt.isFile() && dirEnt.name.startsWith('zz_')) {
      fs.unlinkSync(dir.path + '/' + dirEnt.name);
    }
  }
  dir.close();
  cleanGeneratedFiles(outputDir + '/fake');
}

// use a function factory to create the closure so that the values of namespace and inputFile are captured on each iteration
function autorestCallback(outputDir, inputFile) {
  return function (error, stdout, stderr) {
    // print any output or error from the autorest command
    logResult(error, stdout, stderr);
    console.log('done generating ' + inputFile);
    // format the output on success
    // print any output or error from go fmt
    if (stderr === '' && error === null) {
      exec('gofmt -w .',
      { cwd: outputDir },
      function (error, stdout, stderr) {
        console.log('formatting ' + outputDir);
        logResult(error, stdout, stderr);
      });
    }
    sem.leave();
  };
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
