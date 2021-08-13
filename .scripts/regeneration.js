// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
const exec = require('child_process').exec;
const fs = require('fs');

// limit to 8 concurrent builds
const sem = require('./semaphore')(8);

const swaggerDir = 'src/node_modules/@microsoft.azure/autorest.testserver/swagger/';

const goMappings = {
    'additionalpropsgroup': ['additionalProperties.json'],
    'arraygroup': ['body-array.json'],
    'azurereportgroup': ['azure-report.json'],
    'azurespecialsgroup': ['azure-special-properties.json', '--head-as-boolean'],
    'booleangroup': ['body-boolean.json'],
    'bytegroup': ['body-byte.json'],
    'complexgroup': ['body-complex.json'],
    'complexmodelgroup': ['complex-model.json'],
    'custombaseurlgroup': ['custom-baseUrl.json'],
    'dategroup': ['body-date.json'],
    'datetimegroup': ['body-datetime.json'],
    'datetimerfc1123group': ['body-datetime-rfc1123.json'],
    'dictionarygroup': ['body-dictionary.json'],
    'durationgroup': ['body-duration.json'],
    'errorsgroup': ['xms-error-responses.json'],
    'extenumsgroup': ['extensible-enums-swagger.json'],
    'filegroup': ['body-file.json'],
    'formdatagroup': ['body-formdata.json'],
    'headergroup': ['header.json'],
    'headgroup': ['head.json', '--head-as-boolean'],
    'httpinfrastructuregroup': ['httpInfrastructure.json', '--head-as-boolean'],
    'integergroup': ['body-integer.json'],
    'lrogroup': ['lro.json'],
    'mediatypesgroup': ['media_types.json'],
    'migroup': ['multiple-inheritance.json'],
    //'modelflatteninggroup': ['model-flattening.json'],
    'morecustombaseurigroup': ['custom-baseUrl-more-options.json'],
    'nonstringenumgroup': ['non-string-enum.json'],
    'numbergroup': ['body-number.json'],
    'objectgroup': ['object-type.json'],
    'optionalgroup': ['required-optional.json'],
    'paginggroup': ['paging.json'],
    'paramgroupinggroup': ['azure-parameter-grouping.json'],
    'reportgroup': ['report.json'],
    'stringgroup': ['body-string.json'],
    'urlgroup': ['url.json'],
    'urlmultigroup': ['url-multi-collectionFormat.json'],
    'validationgroup': ['validation.json'],
    'xmlgroup': ['xml-service.json'],
};

const args = process.argv.slice(2);
var filter = undefined;
const switches = [];
for (var i = 0 ; i < args.length; i += 1) {
    switch (args[i]) {
        case '--filter':
        case '-f':
            filter = args[i + 1]
            i += 1
            break;
        case '--verbose':
        case '-v':
            switches.push('--debug');
            break;
        case '--debugger':
        case '-d':
            switches.push('--go.debugger');
            break;
        default:
            break;
    }
}

if (filter !== undefined) {
    console.log("Using filter: " + filter)
}

// loop through all of the namespaces in goMappings
for (namespace in goMappings) {
    // for each swagger run the autorest command to generate code based on the swagger for the relevant namespace and output to the /generated directory
    const entry = goMappings[namespace];
    const inputFile = swaggerDir + entry[0];
    let extraParams = ['--export-clients', '--module-version=0.1'];
    if (entry.length > 1) {
        extraParams = extraParams.concat(entry.slice());
    }
    generate(namespace, inputFile, 'test/autorest/' + namespace, extraParams.join(' '));
}

const blobStorage = 'https://raw.githubusercontent.com/Azure/azure-rest-api-specs/ee9bd6fe35eb7850ff0d1496c59259eb74f0d446/specification/storage/data-plane/Microsoft.BlobStorage/preview/2020-06-12/blob.json';
generate("azstorage", blobStorage, 'test/storage/2020-06-12/azblob', '--security=AzureKey --module="azstorage" --openapi-type="data-plane"');

const network = 'https://raw.githubusercontent.com/Azure/azure-rest-api-specs/228cf296647f6e41182cee7d1a403990e6a8fe3c/specification/network/resource-manager/readme.md';
generateFromReadme("armnetwork", network, 'package-2020-03', 'test/network/2020-03-01/armnetwork', '--module=armnetwork --azure-arm=true');

const compute = 'https://raw.githubusercontent.com/Azure/azure-rest-api-specs/228cf296647f6e41182cee7d1a403990e6a8fe3c/specification/compute/resource-manager/readme.md';
generateFromReadme("armcompute", compute, 'package-2019-12-01', 'test/compute/2019-12-01/armcompute', '--module=armcompute --azure-arm=true');

const synapseArtifacts = 'https://raw.githubusercontent.com/Azure/azure-rest-api-specs/228cf296647f6e41182cee7d1a403990e6a8fe3c/specification/synapse/data-plane/readme.md';
generateFromReadme("azartifacts", synapseArtifacts, 'package-artifacts-2019-06-01-preview', 'test/synapse/2019-06-01/azartifacts', '--security=AADToken --security-scopes="https://dev.azuresynapse.net/.default" --module="azartifacts" --openapi-type="data-plane"');

const synapseSpark = 'https://raw.githubusercontent.com/Azure/azure-rest-api-specs/228cf296647f6e41182cee7d1a403990e6a8fe3c/specification/synapse/data-plane/readme.md';
generateFromReadme("azspark", synapseSpark, 'package-spark-2019-11-01-preview', 'test/synapse/2019-06-01/azspark', '--security=AADToken --security-scopes="https://dev.azuresynapse.net/.default" --module="azspark" --openapi-type="data-plane"');

const tables = 'https://raw.githubusercontent.com/Azure/azure-rest-api-specs/228cf296647f6e41182cee7d1a403990e6a8fe3c/specification/cosmos-db/data-plane/readme.md';
generateFromReadme("aztables", tables, 'package-2019-02', 'test/tables/2019-02-02/aztables', '--security=AADToken --security-scopes="https://tables.azure.com/.default" --module=aztables --openapi-type="data-plane" --export-clients --azure-validator=false --group-parameters=false');

const keyvault = 'https://raw.githubusercontent.com/Azure/azure-rest-api-specs/1e2c9f3ec93078da8078389941531359e274f32a/specification/keyvault/data-plane/readme.md';
generateFromReadme("azkeyvault", keyvault, 'package-7.2', 'test/keyvault/7.2/azkeyvault', '--security=AADToken --security-scopes="https://vault.azure.net/.default" --module=azkeyvault --openapi-type="data-plane" --export-clients');

generate("azalias", 'test/swagger/alias.json', 'test/maps/azalias', '--security=AzureKey --module="azalias" --openapi-type="data-plane"');

function should_generate(name) {
    if (filter !== undefined) {
        let re = new RegExp(filter);
        return re.test(name)
    }
    return true
}

// helper to log the package being generated before invocation
function generate(name, inputFile, outputDir, additionalArgs) {
    if (!should_generate(name)) {
        return
    }
    sem.take(function() {
        console.log('generating ' + inputFile);
        cleanGeneratedFiles(outputDir);
        exec('autorest --use=. --file-prefix="zz_generated_" --modelerfour.lenient-model-deduplication --license-header=MICROSOFT_MIT_NO_VERSION --module-version=0.1 --input-file=' + inputFile + ' --output-folder=' + outputDir + ' ' + additionalArgs + ' ' + switches.join(' '), autorestCallback(outputDir, inputFile));
    });
}

function generateFromReadme(name, readme, tag, outputDir, additionalArgs) {
    if (!should_generate(name)) {
        return
    }
    sem.take(function() {
        console.log('generating ' + readme);
        cleanGeneratedFiles(outputDir);
        exec('autorest --use=. ' + readme + ' --tag=' + tag + ' --file-prefix="zz_generated_" --modelerfour.lenient-model-deduplication --license-header=MICROSOFT_MIT_NO_VERSION --module-version=0.1 --output-folder=' + outputDir + ' ' + additionalArgs + ' ' + switches.join(' '), autorestCallback(outputDir, readme));
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
        if (dirEnt.isFile() && dirEnt.name.startsWith('zz_generated_')) {
            fs.unlinkSync(dir.path + '/' + dirEnt.name);
        }
    }
    dir.close();
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
