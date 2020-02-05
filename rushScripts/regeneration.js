var exec = require('child_process').exec;

swaggerDir = "src/node_modules/@microsoft.azure/autorest.testserver/swagger/";

goMappings = {
    'additionalproperties':['additionalProperties.json'],
    'arraygroup':['body-array.json'],
    'booleangroup':['body-boolean.json'],
    'bytegroup':['body-byte.json'],
    'complexgroup':['body-complex.json'],
    'dategroup':['body-date.json'],
    'datetimerfc1123group':['body-datetime-rfc1123.json'],
    'datetimegroup':['body-datetime.json'],
    'dictionarygroup':['body-dictionary.json'],
    'durationgroup':['body-duration.json'],
    'filegroup':['body-file.json'],
    'formdatagroup':['body-formdata.json'],
    'integergroup':['body-integer.json'],
    'numbergroup':['body-number.json'],
    'stringgroup':['body-string.json'],
    'custombaseurlgroup':['custom-baseUrl.json'],
    'headergroup':['header.json'],
    'httpinfrastructuregroup':['httpInfrastructure.json'],
    'lrogroup':['lro.json'],
    'modelflatteninggroup':['model-flattening.json'],
    'report':['report.json'],
    'optionalgroup':['required-optional.json'],
    'urlgroup':['url.json'],
    'urlmultigroup':['url-multi-collectionFormat.json'],
    'validationgroup':['validation.json'],
    'paginggroup':['paging.json'],
    'morecustombaseurigroup':['custom-baseUrl-more-options.json'],
    'azurereport':['azure-report.json']
  };

// loop through all of the namespaces in goMappings
for (namespace in goMappings) {
    // loop through each file related to a particular namespace 
    for (swagger in goMappings[namespace]) {
        // for each swagger run the autorest-beta command to generate code based on the swagger for the relevant namespace and output to the /generated directory
        child = exec("autorest-beta --use=. --clear-output-folder --license-header=MICROSOFT_MIT_NO_VERSION --input-file=" + swaggerDir + goMappings[namespace][swagger] + " --namespace=" + namespace + " --output-folder=test/autorest/generated/" + namespace + " --version:3.0.6192 --module-path=generatortests/autorest/generated/" + namespace,
        function (error, stdout, stderr) {
            // print any output or error from the autorest-beta command
            if (stdout !== '') {
                console.log('autorest-beta stdout: ' + stdout);
            }
            if (stderr !== '') {
                console.log('autorest-beta stderr: ' + stderr);
            }
            // print any output resulting from executing the autorest-beta command
            if (error !== null) {
                console.log('autorest-beta exec error: ' + error);
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
