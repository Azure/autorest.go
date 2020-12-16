// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
const exec = require('child_process').exec;
const fs = require('fs');

// limit to 8 concurrent builds
const sem = require('./semaphore')(8);

const swaggerDir = 'src/node_modules/@microsoft.azure/autorest.testserver/swagger/';

const goMappings = {
    'additionalpropsgroup': 'additionalProperties.json',
    'arraygroup': 'body-array.json',
    'azurereportgroup': 'azure-report.json',
    'azurespecialsgroup': 'azure-special-properties.json',
    'booleangroup': 'body-boolean.json',
    'bytegroup': 'body-byte.json',
    'complexgroup': 'body-complex.json',
    'complexmodelgroup': 'complex-model.json',
    'custombaseurlgroup': 'custom-baseUrl.json',
    'dategroup': 'body-date.json',
    'datetimegroup': 'body-datetime.json',
    'datetimerfc1123group': 'body-datetime-rfc1123.json',
    'dictionarygroup': 'body-dictionary.json',
    'durationgroup': 'body-duration.json',
    'errorsgroup': 'xms-error-responses.json',
    'extenumsgroup': 'extensible-enums-swagger.json',
    'filegroup': 'body-file.json',
    //'formdatagroup': 'body-formdata.json',
    'headergroup': 'header.json',
    'headgroup': 'head.json',
    'httpinfrastructuregroup': 'httpInfrastructure.json',
    'integergroup': 'body-integer.json',
    'lrogroup': 'lro.json',
    'mediatypesgroup': 'media_types.json',
    'migroup': 'multiple-inheritance.json',
    //'modelflatteninggroup': 'model-flattening.json',
    'morecustombaseurigroup': 'custom-baseUrl-more-options.json',
    'nonstringenumgroup': 'non-string-enum.json',
    'numbergroup': 'body-number.json',
    'optionalgroup': 'required-optional.json',
    'paginggroup': 'paging.json',
    'paramgroupinggroup': 'azure-parameter-grouping.json',
    'reportgroup': 'report.json',
    'stringgroup': 'body-string.json',
    'urlgroup': 'url.json',
    'urlmultigroup': 'url-multi-collectionFormat.json',
    'validationgroup': 'validation.json',
    'xmlgroup': 'xml-service.json',
};

// loop through all of the namespaces in goMappings
for (namespace in goMappings) {
    // for each swagger run the autorest command to generate code based on the swagger for the relevant namespace and output to the /generated directory
    const inputFile = swaggerDir + goMappings[namespace];
    generate(inputFile, 'test/autorest/' + namespace, '--head-as-boolean=true');
}

const blobStorage = 'https://raw.githubusercontent.com/Azure/azure-rest-api-specs/storage-dataplane-preview/specification/storage/data-plane/Microsoft.BlobStorage/preview/2019-07-07/blob.json';
generate(blobStorage, 'test/storage/2019-07-07/azblob', '--credential-scope="https://storage.azure.com/.default" --module="azstorage" --openapi-type="data-plane"');

const network = 'https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/network/resource-manager/readme.md';
generateFromReadme(network, 'package-2020-03', 'test/network/2020-03-01/armnetwork', '--credential-scope="https://management.azure.com//.default" --armcore-connection=true --module=armnetwork');

const compute = 'https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/compute/resource-manager/readme.md';
generateFromReadme(compute, 'package-2019-12-01', 'test/compute/2019-12-01/armcompute', '--credential-scope="https://management.azure.com//.default" --armcore-connection=true --module=armcompute');

const synapseArtifacts = 'https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/synapse/data-plane/readme.md';
generateFromReadme(synapseArtifacts, 'package-artifacts-2019-06-01-preview', 'test/synapse/2019-06-01/azartifacts', '--credential-scope="https://dev.azuresynapse.net/.default" --module="azartifacts" --openapi-type="data-plane"');

const synapseSpark = 'https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/synapse/data-plane/readme.md';
generateFromReadme(synapseSpark, 'package-spark-2019-11-01-preview', 'test/synapse/2019-06-01/azspark', '--credential-scope="https://dev.azuresynapse.net/.default" --module="azspark" --openapi-type="data-plane"');

// helper to log the package being generated before invocation
function generate(inputFile, outputDir, additionalArgs) {
    sem.take(function() {
        console.log('generating ' + inputFile);
        cleanGeneratedFiles(outputDir);
        exec('autorest --version:3.0.6338 --use=. --file-prefix="zz_generated_" --modelerfour.lenient-model-deduplication --license-header=MICROSOFT_MIT_NO_VERSION --input-file=' + inputFile + ' --output-folder=' + outputDir + ' ' + additionalArgs, autorestCallback(outputDir, inputFile));
    });
}

function generateFromReadme(readme, tag, outputDir, additionalArgs) {
    sem.take(function() {
        console.log('generating ' + readme);
        cleanGeneratedFiles(outputDir);
        exec('autorest  --version:3.0.6338 --use=. ' + readme + ' --tag=' + tag + ' --file-prefix="zz_generated_" --modelerfour.lenient-model-deduplication --license-header=MICROSOFT_MIT_NO_VERSION --output-folder=' + outputDir + ' ' + additionalArgs, autorestCallback(outputDir, readme));
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
            exec('go fmt ./' + outputDir,
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
