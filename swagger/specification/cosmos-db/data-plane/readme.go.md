## Go

These settings apply only when `--go` is specified on the command line.

``` yaml $(go)
go:
  license-header: MICROSOFT_MIT_NO_VERSION
  namespace: storagetables
  clear-output-folder: false
```

### Go multi-api

``` yaml $(go) && $(multiapi)
batch:
  - tag: package-2019-02
```

### Tag:  package-2019-02 and go

These settings apply only when `--tag=package-2019-02 --go` is specified on the command line.
Please also specify `--go-sdk-folder=<path to the root directory of your azure-sdk-for-go clone>`.

``` yaml $(tag) == 'package-2019-02' && $(go)
output-folder: $(go-sdk-folder)/services/preview/storage/tables/2019-02-02-preview/$(namespace)
```