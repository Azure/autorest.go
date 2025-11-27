# Release History

## 4.7.6 (2025-11-26)

### Other Changes

* Remove useless Go version header in generated files.

## 4.7.5 (2025-07-01)

### Other Changes

* Upgrade @autorest/testmodeler for compatible with node20.

## 4.7.4 (2025-04-07)

### Features Added

* Added switch `--factory-gather-all-params` to control the `NewClientFactory` constructor parameters. This switch allows gathering either only common parameters of clients or all parameters of clients. The default value is `True`.

## 4.7.3 (2024-04-22)

### Features Added

* Consolidate to use client factory to initialize clients.

## 4.7.2 (2024-04-08)

### Other Changes

* Rearrange autorest pipeline to add go transform info for example model.

## 4.7.1 (2024-02-20)

### Other Changes

* Update dep of @autorest/go.

## 4.7.0 (2023-11-09)

### Features Added

* Add fake support.

### Other Changes

* Update to latest codegen.

## 4.6.2 (2023-07-28)

### Bugs Fixed

* Fix major version module name and uuid issue.

## 4.6.1 (2023-06-29)

### Other Changes

* Update dependencies.

## 4.6.0 (2023-03-13)

### Features Added

* Change example generation to use `ClientFactory`.

## 4.5.2 (2023-01-30)

### Bugs Fixed

* Fix autorest pipeline issue after go generator upgrade.

## 4.5.1 (2023-01-17)

### Bugs Fixed

* Fix test generation problem of any type.

## 4.5.0 (2023-01-16)

### Other Changes

* Upgrade to @autorest/go_4.0.0-preview.45 and do some corresponding change to test generation.

## 4.4.0 (2022-10-25)

### Features Added

* Support refer usage for all types of variables and enhance support for step variables.
* Refine example generation to provide more useful response info.

### Bugs Fixed

* Fix env and prefix string issue for API scenario test generation.
* Fix parse problem for object param in example file.
* Fix wrong camel and snake method.

## 4.3.0 (2022-08-24)

### Features Added

* Support variable with prefix string type for API scenario.

### Other Changes

* Upgrade to new testmodeler to support `operationId` step in API scenario.

## 4.2.2 (2022-08-19)

### Bugs Fixed

* Fix illegal example function name.

## 4.2.1 (2022-08-04)

### Bugs Fixed

* Fix wrong parse for map key value with variable.
* Fix wrong pointer return for LRO test.
* Remove useless gofmt for testgen lint process.

## 4.2.0 (2022-07-27)

### Features Added

* Support API scenario 1.2.

## 4.1.0 (2022-07-19)

### Features Added

* Generate all the examples from swagger for operations.

## 4.0.2 (2022-06-08)

### Other Changes

* Change test and example filename.

## 4.0.1 (2022-05-23)

### Bugs Fixed

* Fix module import problem when SDK version bigger than v1.

## 4.0.0 (2022-05-16)

### Breaking Changes

* Align test code with GA core lib.

## 3.1.2 (2022-04-25)

### Bugs Fixed

* Fix some generation issue.

### Other Changes

* Use oav@2.12.1.

## 3.1.1 (2022-04-18)

### Bugs Fixed

* Fix wrong log.Fatalf usage.

## 3.1.0 (2022-04-15)

### Other Changes

* Upgrade to latest codegen and change list operation name.

## 3.0.1 (2022-04-11)

### Bugs Fixed

* Fix wrong go version in templates.

## 3.0.0 (2022-04-07)

### Breaking Changes

* Support latest GO codegen with generic feature.

## 2.2.1 (2022-03-29)

### Bugs Fixed

* Client subscription param problem.
* LRO need to get final response type name.

### Other Changes

* Use @autorest/testmodeler@2.2.3.

## 2.2.0 (2022-03-17)

### Features Added

* Add sample generation.
* Consolidate manual-written and auto-generated scenario test code.

### Bugs Fixed

* Operation has no subscriptionID param but client has, need to handle it seperately.

### Other Changes

* Update to latest azcore for mock test.
* Change from go get to go install to prevent warnning.

## 2.1.4 (2022-03-07)

### Bugs Fixed

* Fix wrong generation for output variable with chain invoke.

## 2.1.3 (2022-03-03)

### Other Changes

* Change response usage in examples.

## 2.1.2 (2022-03-03)

### Other Changes

* Upgrade to latest testmodeler.

## 2.1.1 (2022-02-24)

### Bugs Fixed

* Fix param render bug for resource deployment step in api scenario.

## 2.1.0 (2022-02-22)

### Other Changes

* Change output variable value fetch method according to new testmodeler.

## 2.0.0 (2022-02-11)

### Breaking Changes

* Add scenario test generation support.
* Add recording support to scenario test.

## 1.3.0 (2022-01-12)

### Features Added

* Use new api scenario through testmodeler.

## 1.2.0 (2022-01-12)

### Features Added

* Compatible with latest azcore and azidentity.
* Add response check to mock test generation.

### Bugs Fixed

* Fix result check problem for lro operation with pageable config.
* Fix result log problem for multiple response operation.
* Fix wrong param name for pageable operation with custom item name.
* Different conversion for choice and sealedchoice.
* Fix wrong generation of null value for object.
* Fix some generated problems including: polymorphism response type, client param, pager response check.
* Fix multiple time format and any-object default value issue.
* Refine log for mock test and fix array item code generate bug.

### Other Changes

* Upgrade to latest autorest/core and autorest/go.

## 1.1.3 (2021-11-29)

### Other Changes

* Replace incomplete response check with just log temporarily.

## 1.1.2 (2021-11-15)

### Bugs Fixed

* Fix some generation corner case.

## 1.1.1 (2021-11-09)

### Other Changes

* Remove `go mod tidy` process.

## 1.1.0 (2021-11-09)

### Other Changes

* Refactor structure and fix most of generation problem.

## 1.0.0 (2021-11-01)

### Other Changes

* Init public version of autorest extension for GO test generation.

