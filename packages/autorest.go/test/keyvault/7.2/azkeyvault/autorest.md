### AutoRest Configuration

> see https://aka.ms/autorest

``` yaml
require:
- https://github.com/Azure/azure-rest-api-specs/blob/1e2c9f3ec93078da8078389941531359e274f32a/specification/keyvault/data-plane/readme.md
- https://github.com/Azure/azure-rest-api-specs/blob/1e2c9f3ec93078da8078389941531359e274f32a/specification/keyvault/data-plane/readme.go.md
license-header: MICROSOFT_MIT_NO_VERSION
module: azkeyvault
module-version: 0.1.0
security: AADToken
security-scopes: https://vault.azure.net/.default
export-clients: true

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
