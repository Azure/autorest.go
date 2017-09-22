
# Issues
If you're reading this, you've likely noticed that this repository hasn't enabled issues. That isn't because we don't want your feedback! **Please file issues for the repository in one of the following repositories** as appropriate:

  - [Azure/azure-sdk-for-go](https://github.com/Azure/azure-sdk-for-go/issues)
  - [Azure/autorest](https://github.com/Azure/autorest)
  - [Azure/azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs)

# Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.microsoft.com.

When you submit a pull request, a CLA-bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., label, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

# AutoRest extension configuration

``` yaml
use-extension:
  "@microsoft.azure/autorest.modeler": "1.9.7"

pipeline:
  go/modeler:
    input: swagger-document/identity
    output-artifact: code-model-v1
    scope: go
  go/commonmarker:
    input: modeler
    output-artifact: code-model-v1
  go/cm/transform:
    input: commonmarker
    output-artifact: code-model-v1
  go/cm/emitter:
    input: transform
    scope: scope-cm/emitter
  go/generate:
    plugin: go
    input: cm/transform
    output-artifact: source-file-go
  go/transform:
    input: generate
    output-artifact: source-file-go
    scope: scope-transform-string
  go/emitter:
    input: transform
    scope: scope-go/emitter

scope-go/emitter:
  input-artifact: source-file-go
  output-uri-expr: $key

output-artifact:
- source-file-go
```
