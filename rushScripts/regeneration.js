// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
var exec = require('child_process').exec;

swaggerDir = "src/node_modules/@microsoft.azure/autorest.testserver/swagger/";

goMappings = {
    'additionalpropertiesgroup':['additionalProperties.json'],
    'arraygroup':['body-array.json'],
    'azurereportgroup':['azure-report.json'],
    'booleangroup':['body-boolean.json'],
    'bytegroup':['body-byte.json'],
    'complexgroup':['body-complex.json'],
    'custombaseurlgroup':['custom-baseUrl.json'],
    'dategroup':['body-date.json'],
    'datetimegroup':['body-datetime.json'],
    'datetimerfc1123group':['body-datetime-rfc1123.json'],
    'dictionarygroup':['body-dictionary.json'],
    'durationgroup':['body-duration.json'],
    'filegroup':['body-file.json'],
    'formdatagroup':['body-formdata.json'],
    'headergroup':['header.json'],
    'httpinfrastructuregroup':['httpInfrastructure.json'],
    'integergroup':['body-integer.json'],
    'lrogroup':['lro.json'],
    'modelflatteninggroup':['model-flattening.json'],
    'morecustombaseurigroup':['custom-baseUrl-more-options.json'],
    'numbergroup':['body-number.json'],
    'optionalgroup':['required-optional.json'],
    'paginggroup':['paging.json'],
    'reportgroup':['report.json'],
    'stringgroup':['body-string.json'],
    'urlgroup':['url.json'],
    'urlmultigroup':['url-multi-collectionFormat.json'],
    'validationgroup':['validation.json'],
    'xmlgroup':['xml-service.json'],
  };

// loop through all of the namespaces in goMappings
for (namespace in goMappings) {
    // loop through each file related to a particular namespace 
    for (swagger in goMappings[namespace]) {
        // for each swagger run the autorest-beta command to generate code based on the swagger for the relevant namespace and output to the /generated directory
        child = exec("autorest --use=. --clear-output-folder --license-header=MICROSOFT_MIT_NO_VERSION --input-file=" + swaggerDir + goMappings[namespace][swagger] + " --namespace=" + namespace + " --output-folder=test/autorest/generated/" + namespace + " --version:3.0.6192 --module-path=generatortests/autorest/generated/" + namespace,
        function (error, stdout, stderr) {
            // print any output or error from the autorest-beta command
            if (stdout !== '') {
                console.log('autorest stdout: ' + stdout);
            }
            if (stderr !== '') {
                console.log('autorest stderr: ' + stderr);
            }
            // print any output resulting from executing the autorest-beta command
            if (error !== null) {
                console.log('autorest exec error: ' + error);
            }
            // print any output or error from go fmt
            fmt = exec("go fmt ./test/autorest/generated/" + namespace,
            function (error, stdout, stderr) {
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
        });
    }

} 
