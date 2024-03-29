## Go

These settings apply only when `--go` is specified on the command line.

``` yaml $(go)
go:
  license-header: MICROSOFT_MIT_NO_VERSION
  namespace: keyvault
  clear-output-folder: false
```

### Go multi-api

``` yaml $(go) && $(multiapi)
batch:
  - tag: package-7.2
  - tag: package-7.1
  - tag: package-7.0
  - tag: package-2016-10
  - tag: package-2015-06
```

``` yaml $(go)
# stuttering clean-up causes a name collision so we fix up the name
# we can't use the in-box rename-model directive due to cross-file references
# so we copy it and make the necessary modifications
directive:
  - from: swagger-document
    where: $.definitions
    transform: >
      if ($.Error) { $.ErrorInfo = $.Error; delete $.Error; }

  - from: swagger-document
    where: $..['$ref']
    transform: |
      $ = $ === "common.json#/definitions/Error" 
        ? "common.json#/definitions/ErrorInfo" 
        : $

  - from: swagger-document
    where: $..['$ref']
    transform: |
      $ = $ === "#/definitions/Error" 
        ? "#/definitions/ErrorInfo" 
        : $
```

### Tag: package-7.2 and go

These settings apply only when `--tag=package-7.2 --go` is specified on the command line.
Please also specify `--go-sdk-folder=<path to the root directory of your azure-sdk-for-go clone>`.

``` yaml $(tag) == 'package-7.2' && $(go)
output-folder: $(go-sdk-folder)/services/$(namespace)/v7.2/$(namespace)
```

### Tag: package-7.1 and go

These settings apply only when `--tag=package-7.1 --go` is specified on the command line.
Please also specify `--go-sdk-folder=<path to the root directory of your azure-sdk-for-go clone>`.

``` yaml $(tag) == 'package-7.1' && $(go)
output-folder: $(go-sdk-folder)/services/$(namespace)/v7.1/$(namespace)
```

### Tag: package-7.0 and go

These settings apply only when `--tag=package-7.0 --go` is specified on the command line.
Please also specify `--go-sdk-folder=<path to the root directory of your azure-sdk-for-go clone>`.

``` yaml $(tag) == 'package-7.0' && $(go)
output-folder: $(go-sdk-folder)/services/$(namespace)/v7.0/$(namespace)
```

### Tag: package-2016-10 and go

These settings apply only when `--tag=package-2016-10 --go` is specified on the command line.
Please also specify `--go-sdk-folder=<path to the root directory of your azure-sdk-for-go clone>`.

``` yaml $(tag) == 'package-2016-10' && $(go)
output-folder: $(go-sdk-folder)/services/$(namespace)/2016-10-01/$(namespace)
```

### Tag: package-2015-06 and go

These settings apply only when `--tag=package-2015-06 --go` is specified on the command line.
Please also specify `--go-sdk-folder=<path to the root directory of your azure-sdk-for-go clone>`.

``` yaml $(tag) == 'package-2015-06' && $(go)
output-folder: $(go-sdk-folder)/services/$(namespace)/2015-06-01/$(namespace)
```