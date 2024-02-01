# Release History

## 4.0.0-preview.62 (unreleased)

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
