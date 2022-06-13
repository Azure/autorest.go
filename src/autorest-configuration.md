# AutoRest Go

The Go plugin is used to generate Go source code.

### Autorest plugin configuration
- Please don't edit this section unless you're re-configuring how the Go extension plugs in to AutoRest
AutoRest needs the below config to pick this up as a plug-in - see https://github.com/Azure/autorest/blob/master/docs/developer/architecture/AutoRest-extension.md

# Pipeline Configuration
``` yaml
version: 3.8.2
use-extension:
  "@autorest/modelerfour" : "4.23.1"

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
  go-transform:
    input: go

  # generates code for the protocol layer
  go-protocol:
    input: go-transform

  # extensibility: allow text-transforms after the code gen
  go/text-transform:
    input: go-protocol

  # output the files to disk
  go/emitter:
    input: 
      - go-transform  # this allows us to dump out the code model after the namer (add --output-artifact:code-model-v4 on the command line)
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
      - key: export-clients
        type: boolean
        description: Indicates if generated clients are to be exported.  Default to true for ARM, false for data-plane.
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
```
