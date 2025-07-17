# Release History

## 0.7.0 (unreleased)

### Breaking Changes

* The `module-version` switch is now used to seed the value for the `moduleVersion` constant. It will _not_ change an existing value.
  * If not specified, it has a default value of `0.1.0`.

### Other Changes

* Updated to the latest tsp toolset.
* The `moduleName` and `moduleVersion` constants have been moved out of `constants.go` and into `version.go`.
  * The `version.go` file is emitted for all SDK flavors.

## 0.6.0 (2025-07-15)

### Breaking Changes

* Fixed some cases where a client name could stutter.
* Force the body paramter to be required for `PATCH` and `PUT` operations.

### Features Added

* Added switch `go-generate` to invoke post-generation scripts.
  * The value is an output-relative path to a `.go` file containing `//go:generate` directives.
  * If Go tools are not on the path, and `go-generate` was specified, then an error is produced.

### Other Changes

* When Go tools are found on the path, the following steps happen after successfully generating code and any `go-generate` script is invoked.
  * Execute `gofmt -w .` followed by `go mod tidy` in the output directory.
  * If Go tools are not found, the above steps are skipped and a warning is displayed.

## 0.5.1 (2025-06-26)

### Other Changes

* Updated to the latest tcgc.

## 0.5.0 (2025-06-23)

### Breaking Changes

* Fixed field names for optional parameters and monomorphic responses to align with `autorest.go` code generator.

## 0.4.12 (2025-06-11)

### Features Add

* Added support for qualified type in example mapping.

### Bugs Fixed

* Fix wrong example mapping for discriminated type.

### Other Changes

* Throw error when paging with re-injection parameters.
* Moved the `_metadata.json` file to the `testdata` subdirectory.
* Updated to the latest tsp toolset.

## 0.4.11 (2025-06-09)

### Other Changes

* Updated to latest tcgc.

## 0.4.10 (2025-05-29)

### Bugs Fixed

* Fix wrong example mapping for any type.

## 0.4.9 (2025-05-29)

### Other Changes

* Updated to the latest tsp toolset.
* Refine `NameCollision` diagnostic.

## 0.4.8 (2025-05-20)

### Other Changes

* Generate metadata file.

## 0.4.7 (2025-05-13)

### Bugs Fixed

* Fix wrong example type mapping for discriminated types.
- Fix wrong nullable response type handling logic.
- Fix example generation logic for optional method parameters.

### Other Changes

* Updated to latest tcgc.

## 0.4.6 (2025-05-07)

### Bugs Fixed

* Fix dependencies

### Other Changes

* Updated to latest tcgc.

## 0.4.5 (2025-05-07)

### Other Changes

* Updated to the latest tsp toolset.

## 0.4.4 (2025-04-25)

### Bugs Fixed

* For fakes, unnecessary time helper files are no longer generated.

### Other Changes

* Length check for `regexp` matches includes full match.
* Updated to latest tcgc.

## 0.4.3 (2025-04-09)

### Bugs Fixed

* Add length check for method prameters to fix generated example test code.

## 0.4.2 (2025-04-03)

### Other Changes

* Updated to the latest tsp toolset.
* Changed the default value of the `--factory-gather-all-params` switch from `false` to `true`.

## 0.4.1 (2025-03-25)

### Bugs Fixed

* Unsupported tsp constructs and other errors are now reported as a diagnostic error instead of an unhandled exception.

### Other Changes

* Updated to the latest tsp toolset.
* Upgraded default `azcore` version to `v1.17.1`.
  * NOTE: this also requires updating the `go` directive in `go.mod` files to version `1.23.0`.

## 0.4.0 (2025-03-12)

### Breaking Changes

* The monomorphic response field will no longer have the name `Value` in some cases. This is to preserve compatibility with the behavior of the `autorest.go` code generator.

## 0.3.11 (2025-03-07)

### Other Changes

* Added switch `generate-samples` to control example code generation.
* Deprecated `generate-examples`, use `generate-samples` instead.

## 0.3.10 (2025-03-06)

### Other Changes

* Updated to the latest tsp toolset.

## 0.3.9 (2025-02-25)

### Other Changes

* Updated to the latest tsp toolset.

## 0.3.8 (2025-02-21)

### Bugs Fixed

* Don't export `NewPager` methods when their access is internal.
* Remove filtering Azure core model since some instances of template model is in `Azure.Core` namespace. Logic of filtering exception model could cover the filtering needs.

### Other Changes

* Updated to the latest tsp toolset.
* Report tcgc diagnostics.
* Added switch `--factory-gather-all-params` to control the `NewClientFactory` constructor parameters. This switch allows gathering either only common parameters of clients or all parameters of clients.

## 0.3.7 (2025-02-11)

### Bugs Fixed

* When dealing with mapping of parameters, get operation's body parameter from method's parameter's property if the method's parameter is a model with property decorated with `@bodyRoot` or `@body`.

## 0.3.6 (2025-01-23)

### Bugs Fixed

* `Operation-Location` LROs will correctly set the `OperationLocationResultPath` when constructing a `Poller[T]`.

### Other Changes

* Updated to the latest tsp toolset.
* Upgraded default `azcore` version to `v1.17.0`.

## 0.3.5 (2024-12-18)

### Other Fixes

* Updated dependencies (accounts for latest tsp toolset).

## 0.3.4 (2024-12-03)

### Other Fixes

* Updated dependencies (fixes missing parameters in some cases).

## 0.3.3 (2024-11-19)

### Other Fixes

* Updated dependencies (fixes incorrectly pruned base models).

## 0.3.2 (2024-11-07)

### Features Added

* Add support for XML payloads.
* Internal fake server request interceptor.

### Bugs Fixed

* Fake servers will honor the caller's context in the `*http.Request`.
* Add missing error check when parsing multipart/form content in fakes.
* Optional request bodies will only set the `Content-Type` header when a body is specified.

### Other Fixes

* Fake pollers will always include `http.StatusOK` as an acceptible status code, and `http.StatusNoContent` for operations that don't return a body.

## 0.3.1 (2024-08-14)

### Bugs Fixed

* Don't prune base models that have been marked as output.

## 0.3.0 (2024-08-06)

### Features Added

* Added example code generation.

### Breaking Changes

* Fixes in TCGC for proper handling of `@clientName`.

### Bump Dependencies

* TCGC 0.44.3
* TypeSpec compiler 0.58.1

## 0.2.0 (2024-07-30)

### Breaking Changes

* For spread params, the optional params are now placed in the options type.

### Other Changes

* Upgraded default azcore version to v1.13.0

### Bump Dependencies

* TCGC 0.44.1
* TypeSpec tools 0.58.0

## 0.1.0 (2024-07-17)

* Initial release
