# AutoRest Go Generator 

The AutoRest Go generator is intended to be used from AutoRest. 

> see https://aka.ms/autorest

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

### Autorest plugin configuration
- Please don't edit this section unless you're re-configuring how the Go extension plugs in to AutoRest
AutoRest needs the below config to pick this up as a plug-in - see https://github.com/Azure/autorest/blob/main/docs/developer/writing-an-extension.md

# Pipeline Configuration
``` yaml
version: 3.9.7
use-extension:
  "@autorest/modelerfour" : "4.26.2"

modelerfour:
  treat-type-object-as-anything: true
  resolve-schema-name-collisons: true
  naming:
    preserve-uppercase-max-length: 64

pipeline:
  go:
    pass-thru: true
    input: modelerfour/identity

  # fix up names add Go-specific data to the code model
  go-transform-m4:
    input: go

  go-m4-to-gocodemodel:
    input: go-transform-m4

  # generates code
  go-codegen:
    input: go-m4-to-gocodemodel

  # extensibility: allow text-transforms after the code gen
  go/text-transform:
    input:
      - go-codegen

  # output the files to disk
  go/emitter:
    input: 
      - go-transform-m4  # this allows us to dump out the code model after transformation (add --output-artifact:code-model-v4 on the command line)
      - go/text-transform # this grabs the outputs after the last step.
      
    is-object: false # tell it that we're not putting an object graph out
    output-artifact: source-file-go # the file 'type' that we're outputting.

  #go/emitter/command:
  #  input: emitter
  #  run: 
  #    - node -e "console.log('hi'); process.exit(1);"
  #    - node -e "console.log('hi'); process.exit(0);"
```

#### Help

```yaml
help-content:
  go: # type: Help as defined in autorest-core/help.ts
    activationScope: go
    categoryFriendlyName: Go Generator
    settings:
      - key: module
        type: string
        description: The name of the Go module written to go.mod.  Omit to skip go.mod generation.
      - key: azcore-version
        description: Semantic version of azcore without the leading 'v' to use if different from the default version (e.g. 1.2.3).
        type: string
      - key: file-prefix
        type: string
        description: Optional prefix to file names. For example, if you set your file prefix to "zzz_", all generated code files will begin with "zzz_".
      - key: module-version
        description: When --azure-arm is true, semantic version to include in generated telemetryInfo constant without the leading 'v' (e.g. 1.2.3).
        type: string
      - key: group-parameters
        description: Enables parameter grouping via x-ms-parameter-grouping, defaults to true.
        type: boolean
      - key: stutter
        type: string
        description: Uses the specified value to remove stuttering from types and funcs instead of the built-in algorithm.
      - key: honor-body-placement
        type: boolean
        description: When true, optional body parameters are treated as such for PATCH and PUT operations.
      - key: remove-unreferenced-types
        type: boolean
        description: When true, non-reference schema will be removed from the generated code.
      - key: normalize-operation-name
        type: boolean
        description: When true, add suffix for operation with unstructured body type and keep original name for operation with structured body type. When false, keep original name if only one body type, and add suffix for operation with non-binary body type if more than one body type.
      - key: rawjson-as-bytes
        type: boolean
        description: When true, properties that are untyped (i.e. raw JSON) are exposed as []byte instead of any or map[string]any. The default is false.
      - key: generate-fakes
        type: boolean
        description: When true, enables generation of fake servers. The default is false.
      - key: slice-elements-byval
        type: boolean
        description: When true, slice elements will not be pointer-to-type. The default is false.
      - key: head-as-boolean
        description: When true, HEAD requests will return a boolean value based on the HTTP status code. The default is false, but will be set to true if --azure-arm is true.
      - key: generate-fakes
        description: Enables generation of fake servers. The default value is set to the value of --azure-arm.
      - key: inject-spans
        description: Enables generation of spans for distributed tracing. The default value is set to the value of --azure-arm.
      - key: single-client
        type: boolean
        description: Indicates package has a single client. This will omit the Client prefix from options and response types. If multiple clients are detected, an error is returned.
      - key: disallow-unknown-fields
        type: boolean
        description: When true, unmarshalers will return an error when an unknown field is encountered in the payload.
      - key: fix-const-stuttering
        type: boolean
        description: When true, fix stuttering for const types and their values.
```
