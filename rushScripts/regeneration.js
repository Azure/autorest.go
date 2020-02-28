// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
var exec = require('child_process').exec;

swaggerDir = 'src/node_modules/@microsoft.azure/autorest.testserver/swagger/';

goMappings = {
    'additionalpropertiesgroup': 'additionalProperties.json',
    //'arraygroup': 'body-array.json',
    //'azurereportgroup': 'azure-report.json',
    'booleangroup': 'body-boolean.json',
    'bytegroup': 'body-byte.json',
    //'complexgroup': 'body-complex.json',
    'custombaseurlgroup': 'custom-baseUrl.json',
    //'dategroup': 'body-date.json',
    //'datetimegroup': 'body-datetime.json',
    //'datetimerfc1123group': 'body-datetime-rfc1123.json',
    //'dictionarygroup': 'body-dictionary.json',
    //'durationgroup': 'body-duration.json',
    'filegroup': 'body-file.json',
    //'formdatagroup': 'body-formdata.json',
    'headergroup': 'header.json',
    //'httpinfrastructuregroup': 'httpInfrastructure.json',
    //'integergroup': 'body-integer.json',
    //'lrogroup': 'lro.json',
    //'modelflatteninggroup': 'model-flattening.json',
    'morecustombaseurigroup': 'custom-baseUrl-more-options.json',
    'numbergroup': 'body-number.json',
    //'optionalgroup': 'required-optional.json',
    //'paginggroup': 'paging.json',
    //'reportgroup': 'report.json',
    'stringgroup': 'body-string.json',
    'urlgroup': 'url.json',
    //'urlmultigroup': 'url-multi-collectionFormat.json',
    //'validationgroup': 'validation.json',
    //'xmlgroup': 'xml-service.json',
  };

// loop through all of the namespaces in goMappings
for (namespace in goMappings) {
    // for each swagger run the autorest command to generate code based on the swagger for the relevant namespace and output to the /generated directory
    let inputFile = swaggerDir + goMappings[namespace];
    exec('autorest --use=. --clear-output-folder --license-header=MICROSOFT_MIT_NO_VERSION --input-file=' + inputFile + ' --namespace=' + namespace + ' --output-folder=test/autorest/generated/' + namespace + ' --module-path=generatortests/autorest/generated/' + namespace, autorestCallback(namespace, inputFile));
} 

// use a function factory to create the closure so that the values of namespace and inputFile are captured on each iteration
function autorestCallback(namespace, inputFile) {
    return function (error, stdout, stderr) {
        console.log('generating ' + inputFile);
        // print any output or error from the autorest command
        if (stdout !== '') {
            console.log('autorest stdout: ' + stdout);
        }
        if (stderr !== '') {
            console.log('autorest stderr: ' + stderr);
        }
        // print any output resulting from executing the autorest command
        if (error !== null) {
            console.log('autorest exec error: ' + error);
        }
        // format the output
        // print any output or error from go fmt
        let formatDir = './test/autorest/generated/' + namespace + '/...';
        exec('go fmt ' + formatDir,
        function (error, stdout, stderr) {
            console.log('formatting ' + formatDir);
            if (stdout !== '') {
                console.log('fmt stdout: ' + stdout);
            }
            if (stderr !== '') {
                console.log('fmt stderr: ' + stderr);
            }
            // print any output resulting from a failure to execute go fmt
            if (error !== null) {
                console.log('fmt exec error: ' + error);
            }
        });
    };
}
