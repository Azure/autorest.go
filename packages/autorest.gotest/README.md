# Autorest Extension for GO Test Generation

Generate \*.md config files in Azure REST API specification:

https://github.com/Azure/azure-rest-api-specs

## How to Generate GO Test Code

```
autorest --version=3.8.2 --use=@autorest/go@latest --use=@autorest/gotest@latest --go --track2 --output-folder=<RP package path> --file-prefix="zz_generated_" --clear-output-folder=false --go.clear-output-folder=false --testmodeler.generate-mock-test --testmodeler.generate-sdk-example --testmodeler.generate-scenario-test --testmodeler.generate-sdk-sample <RP config md file path>
```

## Configurations

Below are options can be used for autorest.gotest

### --generate-sdk

Generate GO SDK code along with test and examples.

### --debug

Generate modeler files in [output-foler]/\_\_debuger for debug purpose.

### --testmodeler.generate-mock-test

Generate mock test.

### --testmodeler.generate-sdk-example

Generate SDK usage examples.

### --testmodeler.generate-scenario-test

Generate scenario test.

### --testmodeler.generate-sdk-sample

Generate SDK sample code.

### --example-file-prefix

File prefix for auto-generated examples.

### --test-file-prefix

File prefix for auto-generated tests.

## Autorest Pipeline Configurations

```yaml $(go)
clear-output-folder: false

include-x-ms-examples-original-file: true
modelerfour:
  include-x-ms-examples-original-file: true

testmodeler:
  scenario:
    codemodel-restcall-only: false

version: 3.8.2

use-extension:
  '@autorest/go': '4.0.0-preview.47'
  '@autorest/testmodeler': '2.3.2'

pipeline:
  test-modeler:
    input:
      - go
    output-artifact: source-file-test-modeler
  testmodeler/emitter:
    input:
      - test-modeler
    scope: scope-testmodeler/emitter
  go-transform:
    input:
      - test-modeler
  go-tester:
    input: go-transform
    output-artifact: source-file-go-tester
  gotest/emitter:
    input:
      - go-tester
    scope: scope-gotest/emitter
  go-linter:
    input:
      - go-tester
      - gotest/emitter

scope-testmodeler/emitter:
  input-artifact:
    - source-file-test-modeler
  output-uri-expr: $key

scope-gotest/emitter:
  input-artifact:
    - source-file-go-tester
  output-uri-expr: $key
```

```yaml $(go) && !$(generate-sdk)
pipeline:
  go/emitter:
    scope: scope-gotest/emitter
```

```yaml $(debug)
testmodeler:
  export-codemodel: true
```
