# AutoRest Go

The Go plugin is used to generate Go source code.

### Autorest plugin configuration
- Please don't edit this section unless you're re-configuring how the Go extension plugs in to AutoRest
AutoRest needs the below config to pick this up as a plug-in - see https://github.com/Azure/autorest/blob/master/docs/developer/architecture/AutoRest-extension.md

> if the modeler is loaded already, use that one, otherwise grab it.

``` yaml !isLoaded('@autorest/remodeler') 
use-extension:
  "@autorest/modelerfour" : "4.15.440" 

modelerfour:
  resolve-schema-name-collisons: true
  naming:
    preserve-uppercase-max-length: 64
```


> Multi-Api Mode
``` yaml
pipeline-model: v3
```

# Pipeline Configuration
``` yaml
pipeline:
  # fix up names add Go-specific data to the code model
  go-transform:
    input: modelerfour/identity

  # generates code for the protocol layer
  go-protocol:
    input: go-transform

  # extensibility: allow text-transforms after the code gen
  go/text-transform:
    input: go-protocol

  # output the files to disk
  go/emitter:
    input: 
      - go-protocol
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
