# Release History

## 4.0.0-preview.67 (2024-07-30)

## Bugs Fixed

* Fixed a rare issue causing some method doc comments to be omitted.
* Fixed bad codegen for slices of raw JSON objects.

### Other Changes

* Emit unused params in helper methods with the `_` name.
* Removed unnecessary `aux` variable for some corner-cases.
* Upgraded default `azcore` version to `v1.13.0`.

## 4.0.0-preview.66 (2024-04-25)

### Bugs Fixed

* Removed references to `__filename`.

## 4.0.0-preview.65 (2024-04-24)

### Bugs Fixed

* Fixed missing dependencies.

## 4.0.0-preview.64 (2024-04-23) - DEPRECATED

### Bugs Fixed

* Fixed hard-coded `Metadata` field in header collection responses.
* Don't error on empty time values during unmarshaling.
* Fixed bad codegen for optional multipart/form parameters.
* Fixed bad codegen for templated host parameters.

### Features Added

* Added option `fix-const-stuttering` to fix stuttering for `const` types and values.

## Other Changes

* Removed Go 1.18 build constraints from generated code.
* Use latest `azcore` in generated `go.mod` files.
* Moved response envelope SerDe methods to their own file.
* Improved support for multipart/form parameter types.

## 4.0.0-preview.63 (2024-02-07)

### Bugs Fixed

* Fixed package installation failure due to improper use of workspaces.

## 4.0.0-preview.62 (2024-02-06) - DEPRECATED

### Features Added

* Added switch `--single-client` when generating content with a single client.

### Bugs Fixed

* Fixed bad codegen for slices of base-64 encoded data.
* Fixed bad codegen for parsing response headers containing Unix time.
* Unmarshalers properly handle JSON `null` values.
* Lenient parsing of RFC3339 as ISO8601 with space as date-time separator character.

### Other Changes

* Renamed generated file `response_types.go` to `responses.go`
* Consume centralized codegen and codemodel projects.
* ARM client factory will share the same `*arm.Client` instance across SDKs.
* Skip generating empty `models.go` files.
* Setting header and query params codegen is now sorted by wire name.
* Use latest `azcore` in generated `go.mod` files.
