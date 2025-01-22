# Release History

## 0.3.6 (unreleased)

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
