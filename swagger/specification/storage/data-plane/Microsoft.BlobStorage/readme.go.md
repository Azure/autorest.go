# Code Generation - Azure Blob SDK for Golang

### Settings

```yaml
go: true
clear-output-folder: false
version: "^3.0.0"
license-header: MICROSOFT_MIT_NO_VERSION
input-file: "preview/2021-12-02/blob.json"
credential-scope: "https://storage.azure.com/.default"
openapi-type: "data-plane"
verbose: true
security: AzureKey
honor-body-placement: true
module: azblob
modelerfour:
  group-parameters: false
  lenient-model-deduplication: true
export-clients: true
```

``` yaml
directive:
- from: swagger-document
  where: $.definitions
  transform: >
    $.BlobItemInternal["x-ms-go-omit-serde-methods"] = true;
    $.AccessPolicy["x-ms-go-omit-serde-methods"] = true;
```
