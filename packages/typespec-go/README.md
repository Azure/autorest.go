# TypeSpec Go Generator 

The TypeSpec Go generator is intended to be used with the TypeSpec compiler.

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

#### Help

```yaml
options:
    - key: azcore-version
        type: string
        description: Semantic version of azcore without the leading \'v\' to use if different from the default version (e.g. 1.2.3).
    - key: disallow-unknown-fields
        type: boolean
        description: When true, unmarshalers will return an error when an unknown field is encountered in the payload. The default is false.
    - key: file-prefix
        type: string
        description: Optional prefix to file names. For example, if you set your file prefix to "zzz_", all generated code files will begin with "zzz_".
    - key: generate-fakes
        type: boolean
        description: When true, enables generation of fake servers. The default is false.
    - key: head-as-boolean
        type: boolean
        description: When true, HEAD requests will return a boolean value based on the HTTP status code. The default is false.
    - key: inject-spans
        type: boolean
        description: Enables generation of spans for distributed tracing. The default is false.
    - key: module
        type: string
        nullable: true
        description: The name of the Go module written to go.mod. Omit to skip go.mod generation. When module is specified, module-version must also be specified.
    - key: module-version
        type: string
        nullable: true
        description: Semantic version of the Go module without the leading \'v\' written to constants.go. (e.g. 1.2.3). When module-version is specified, module must also be specified.
    - key: rawjson-as-bytes
        type: boolean
        nullable: true
        description: When true, properties that are untyped (i.e. raw JSON) are exposed as []byte instead of any or map[string]any. The default is false.
    - key: slice-elements-byval
        type: boolean
        nullable: true
        description: When true, slice elements will not be pointer-to-type. The default is false.
    - key: single-client
        type: boolean
        nullable: true
        description: Indicates package has a single client. This will omit the Client prefix from options and response types. If multiple clients are detected, an error is returned. The default is false.
    - key: stutter
        type: string
        nullable: true
        description: Uses the specified value to remove stuttering from types and funcs instead of the built-in algorithm.
    - key: fix-const-stuttering
        type: boolean
        nullable: true
        description: When true, fix stuttering for `const` types and values. The default is false.
    - key: generate-examples
        type: boolean
        nullable: true
        description: When true, generate example tests. The default is false. It will be deprecated in the future, please use generate-samples.
    - key: generate-samples
        type: boolean
        nullable: true
        description: When true, generate example tests. The default is false.
    - key: factory-gather-all-params
        type: boolean
        nullable: true
        description: When true, the NewClientFactory constructor gathers all parameters or only common parameters of clients. The default is false.
```
