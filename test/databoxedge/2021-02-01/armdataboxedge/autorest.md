### AutoRest Configuration

> see https://aka.ms/autorest

``` yaml
azure-arm: true
require:
- https://github.com/Azure/azure-rest-api-specs/blob/3865f04d22e82db481be0727b406021d29cd2b70/specification/databoxedge/resource-manager/readme.md
- https://github.com/Azure/azure-rest-api-specs/blob/3865f04d22e82db481be0727b406021d29cd2b70/specification/databoxedge/resource-manager/readme.go.md
license-header: MICROSOFT_MIT_NO_VERSION
module: armdataboxedge
module-version: 0.1.0

# stuttering clean-up causes a name collision so we fix up the name
# https://github.com/Azure/autorest/blob/main/docs/generate/built-in-directives.md#rename-model
directive:
  - rename-model:
      from: 'Sku'
      to: 'SkuType'
```
