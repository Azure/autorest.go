# Change Log - @autorest/gotest

This log was last generated on Tue, 01 Jul 2025 05:17:10 GMT and should not be manually modified.

## 4.7.5
Tue, 01 Jul 2025 05:17:10 GMT

### Patches

- upgrade @autorest/testmodeler for compatible with node20

## 4.7.4
Mon, 07 Apr 2025 02:41:10 GMT

### Patches

- Added switch `--factory-gather-all-params` to control the `NewClientFactory` constructor parameters. This switch allows gathering either only common parameters of clients or all parameters of clients. The default value is `True`.

## 4.7.3
Mon, 22 Apr 2024 09:31:17 GMT

### Patches

- Consolidate to use client factory to initialize clients

## 4.7.2
Mon, 08 Apr 2024 06:34:25 GMT

### Patches

- Rearrange autorest pipeline to add go transform info for example model

## 4.7.1
Tue, 20 Feb 2024 09:23:59 GMT

### Patches

- Update dep of @autorest/go

## 4.7.0
Thu, 09 Nov 2023 03:29:34 GMT

### Minor changes

- Update to latest codegen and add fake support

## 4.6.2
Fri, 28 Jul 2023 07:01:15 GMT

### Patches

- Fix major version module name and uuid issue.

## 4.6.1
Thu, 29 Jun 2023 09:09:50 GMT

### Patches

- Update dependencies.

## 4.6.0
Mon, 13 Mar 2023 09:28:18 GMT

### Minor changes

- Change example generation to use `ClientFactory`.

## 4.5.2
Mon, 30 Jan 2023 08:26:19 GMT

### Patches

- Fix autorest pipeline issue after go generator upgrade.

## 4.5.1
Tue, 17 Jan 2023 05:21:21 GMT

### Patches

- Fix test generation problem of any type.

## 4.5.0
Mon, 16 Jan 2023 06:25:05 GMT

### Minor changes

- Upgrade to @autorest/go_4.0.0-preview.45 and do some corresponding change to test generation.

## 4.4.0
Tue, 25 Oct 2022 03:01:04 GMT

### Minor changes

- Support refer usage for all types of variables and enhance support for step variables.
- Refine example generation to provide more useful response info.

### Patches

- Fix env and prefix string issue for API scenario test generation.
- Fix parse problem for object param in example file.
- Fix wrong camel and snake method.

## 4.3.0
Wed, 24 Aug 2022 06:43:47 GMT

### Minor changes

- Support variable with prefix string type for API scenario.

### Patches

- Upgrade to new testmodeler to support `operationId` step in API scenario.

## 4.2.2
Fri, 19 Aug 2022 02:39:29 GMT

### Patches

- Fix illegal example funtion name.

## 4.2.1
Thu, 04 Aug 2022 09:20:24 GMT

### Patches

- Fix wrong parse for map key value with variable.
- Fix wrong pointer return for LRO test.
- Remove useless gofmt for testgen lint process.

## 4.2.0
Wed, 27 Jul 2022 02:15:06 GMT

### Minor changes

- Support API scenario 1.2.

## 4.1.0
Tue, 19 Jul 2022 09:34:05 GMT

### Minor changes

- Generate all the examples from swagger for operations.

## 4.0.2
Wed, 08 Jun 2022 07:31:03 GMT

### Patches

- Change test and example filename.

## 4.0.1
Mon, 23 May 2022 07:05:43 GMT

### Patches

- Fix module import problem when SDK version bigger than v1.

## 4.0.0
Mon, 16 May 2022 01:46:35 GMT

### Breaking changes

- Align test code with GA core lib.

## 3.1.2
Mon, 25 Apr 2022 08:06:55 GMT

### Patches

- Use oav@2.12.1
- Fix some generation issue.

## 3.1.1
Mon, 18 Apr 2022 06:25:24 GMT

### Patches

- Fix wrong log.Fatalf usage.

## 3.1.0
Fri, 15 Apr 2022 03:08:44 GMT

### Minor changes

- Upgrade to latest codegen and change list operation name.

## 3.0.1
Mon, 11 Apr 2022 09:15:12 GMT

### Patches

- Fix wrong go version in templates.

## 3.0.0
Thu, 07 Apr 2022 12:22:32 GMT

### Breaking changes

- Support latest GO codegen with generic feature.

## 2.2.1
Tue, 29 Mar 2022 01:56:58 GMT

### Patches

- use @autorest/testmodeler@2.2.3
- Client subscription param problem.
- LRO need to get final response type name.

## 2.2.0
Thu, 17 Mar 2022 07:43:40 GMT

### Minor changes

- Add sample generation.
- Update to latest azcore for mock test.
- Consolidate manual-written and auto-generated scenario test code.

### Patches

- Change from go get to go install to prevent warnning.
- Operation has no subscriptionID param but client has, need to handle it seperately.

## 2.1.4
Mon, 07 Mar 2022 02:56:30 GMT

### Patches

- Fix wrong generation for output variable with chain invoke.

## 2.1.3
Thu, 03 Mar 2022 05:50:36 GMT

### Patches

- Change response usage in examples.

## 2.1.2
Thu, 03 Mar 2022 02:23:21 GMT

### Patches

- Upgrade to latest testmodeler.

## 2.1.1
Thu, 24 Feb 2022 05:54:42 GMT

### Patches

- Fix param render bug for resource deployment step in api scenario.

## 2.1.0
Tue, 22 Feb 2022 10:58:11 GMT

### Minor changes

- Change output variable value fetch method according to new testmodeler.

## 2.0.0
Fri, 11 Feb 2022 09:47:39 GMT

### Breaking changes

- Add scenario test generation support.
- Add recording support to scenario test.

## 1.3.0
Wed, 12 Jan 2022 09:10:46 GMT

### Minor changes

- use new api scenario through testmodeler

## 1.2.0
Wed, 12 Jan 2022 02:19:25 GMT

### Minor changes

- Compatible with latest azcore and azidentity.
- Add response check to mock test generation.

### Patches

- Fix result check problem for lro operation with pageable config.
- Fix result log problem for multiple response operation.
- Fix wrong param name for pageable opeation with custom item name.
- Different conversion for choice and sealedchoice.
- Fix wrong generation of null value for object.
- Fix some generated problems including: polymorphism response type, client param, pager response check.
- Fix multiple time format and any-object default value issue.
- Refine log for mock test and fix array item code generate bug.
- Upgrade to latest autorest/core and autorest/go.

## 1.1.3
Mon, 29 Nov 2021 06:10:09 GMT

### Patches

- Replace incomplete response check with just log temporarily.

## 1.1.2
Mon, 15 Nov 2021 09:39:03 GMT

### Patches

- Fix some generation corner case.

## 1.1.1
Tue, 09 Nov 2021 10:20:51 GMT

### Patches

- Remove `go mod tidy` process.

## 1.1.0
Tue, 09 Nov 2021 09:17:24 GMT

### Minor changes

- Refactor structure and fix most of generation problem.

## 1.0.0
Mon, 01 Nov 2021 09:01:05 GMT

### Breaking changes

- Init public version of autorest extension for GO test generation

