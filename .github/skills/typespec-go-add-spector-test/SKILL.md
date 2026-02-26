---
name: typespec-go-add-spector-test
description: Adds a Spector mock API test for the typespec-go emitter. Use when given a Spector case link (http-specs or azure-http-specs) to generate the Go client, add it to tspcompile.js, write `*_client_test.go` tests, and validate them against the Spector mock server.
---

# Adding a Spector test for typespec-go

## Inputs

You will receive a **Spector case link** pointing to a specific scenario/case under:

- `https://github.com/microsoft/typespec/tree/main/packages/http-specs/specs/...`
- `https://github.com/Azure/typespec-azure/tree/main/packages/azure-http-specs/specs/...`

## Output

- An entry added to `packages/typespec-go/.scripts/tspcompile.js` in the appropriate spec group (`httpSpecsGroup` or `azureHttpSpecsGroup`).
- Generated Go client code under `packages/typespec-go/test/http-specs/` or `packages/typespec-go/test/azure-http-specs/`.
- Hand-written `*_client_test.go` files that validate the generated SDK against the Spector mock API server. One test file per generated client (i.e., per `zz_<name>_client.go` file).

## Workflow (copy as checklist)

- [ ] Ensure prerequisites are met (`pnpm install`, package build)
- [ ] Identify the Spector case link and determine spec type (http-specs vs azure-http-specs)
- [ ] Choose a group name and add the entry to `tspcompile.js`
- [ ] Run `pnpm tspcompile --filter=<groupname>` to generate Go client code
- [ ] Read the `mockapi.ts` file to understand expected request/response
- [ ] Write one `*_client_test.go` test file per generated client, matching the mock API expectations
- [ ] Start the Spector mock server (`pnpm spector --start`)
- [ ] Run the tests and verify they pass
- [ ] Stop the Spector mock server (`pnpm spector --stop`)

## Prerequisites — Environment setup

Before starting, ensure the build environment is ready.

All commands below run from the **repo root** (`autorest.go/`).

1. **Install dependencies**:
   ```bash
   pnpm install
   ```
2. **Build the typespec-go package**:
   ```bash
   cd packages/typespec-go
   pnpm build
   ```

## Step 1 — Identify the Spector case and spec type

### Determine the spec type from the link

- If the link is under `microsoft/typespec/.../packages/http-specs/specs/<path>`:
  - Spec type = **http-specs**
  - The entry goes into `httpSpecsGroup` in `tspcompile.js`
  - The spec input path is `<path>` (relative to the `http-specs/specs/` root)
  - Generated output goes under `packages/typespec-go/test/http-specs/`

- If the link is under `Azure/typespec-azure/.../packages/azure-http-specs/specs/<path>`:
  - Spec type = **azure-http-specs**
  - The entry goes into `azureHttpSpecsGroup` in `tspcompile.js`
  - The spec input path is `<path>` (relative to the `azure-http-specs/specs/` root)
  - Generated output goes under `packages/typespec-go/test/azure-http-specs/`

### Extract the spec path

From the link, extract the path after `specs/`. For example:
- Link: `https://github.com/microsoft/typespec/tree/main/packages/http-specs/specs/type/model/empty`
- Spec path: `type/model/empty`

### Check if the case already exists

Search `tspcompile.js` for the spec path. If it already exists, skip to Step 4 (writing tests).

## Step 2 — Add the entry to tspcompile.js

Open `packages/typespec-go/.scripts/tspcompile.js` and add a new entry to the appropriate group object.

### Entry format

```javascript
'<groupname>': ['<spec-path>', '<optional-emitter-options>...'],
```

### Naming conventions

- **Group name**: lowercase, no hyphens, append `group` suffix. Derived from the last segment of the spec path.
  - `type/model/empty` → `emptygroup`
  - `azure/core/basic` → `basicgroup`
  - `encode/datetime` → `datetimegroup`
  - `azure/client-generator-core/access` → `accessgroup`

- **Spec path**: the path relative to the specs root (e.g., `type/model/empty`, `azure/core/basic`).

### Example additions

For http-specs:
```javascript
const httpSpecsGroup = {
  // ... existing entries ...
  'bytesgroup': ['encode/bytes'],
};
```

For azure-http-specs:
```javascript
const azureHttpSpecsGroup = {
  // ... existing entries ...
  'basicgroup': ['azure/core/basic'],
};
```

### Alphabetical ordering

Insert the new entry in a logical position (typically alphabetical by group name or grouped with related specs).

## Step 3 — Generate the Go client code

From the `packages/typespec-go` directory, run:

```bash
pnpm tspcompile --filter=<groupname>
```

For example:
```bash
pnpm tspcompile --filter=emptygroup
```

This generates Go files (prefixed `zz_`) into the output directory under `test/http-specs/` or `test/azure-http-specs/`.

### Verify generation succeeded

Check that the output directory was created and contains generated files:
- `zz_*_client.go` — client implementation
- `zz_models.go` — model types
- `zz_options.go` — options structs
- `zz_responses.go` — response types
- `go.mod` — Go module file

If the output directory is empty or missing `zz_` files, generation failed. Check the console output for errors.

### Initialize go.sum

After generation, you need to ensure `go.sum` is populated. From the generated module directory:

```bash
cd test/<http-specs|azure-http-specs>/<path>/<groupname>
go mod tidy
```

## Step 4 — Read the mockapi.ts to understand expectations

The `mockapi.ts` file defines the expected HTTP requests and responses for each scenario. It is located at:

- For http-specs: `packages/typespec-go/node_modules/@typespec/http-specs/specs/<spec-path>/mockapi.ts`
- For azure-http-specs: `packages/typespec-go/node_modules/@azure-tools/azure-http-specs/specs/<spec-path>/mockapi.ts`

### Key elements in mockapi.ts

Each scenario entry contains:
- `uri` — the HTTP endpoint path
- `method` — HTTP method (get, post, put, patch, delete)
- `request` — expected request details:
  - `body` — request body (wrapped in `json(...)`)
  - `headers` — expected request headers
  - `query` — expected query parameters
  - `pathParams` — expected path parameters
- `response` — what the mock server returns:
  - `status` — HTTP status code
  - `body` — response body (wrapped in `json(...)`)
  - `headers` — response headers

### Map scenarios to SDK methods

The scenario names follow a pattern like `Type_Model_Empty_getEmpty` which maps to the generated client method `GetEmpty`. Use the generated `zz_*_client.go` files to find the exact method signatures.

## Step 5 — Write the test files

Create one `*_client_test.go` file per generated client. Each generated `zz_<name>_client.go` should have a corresponding `<name>_client_test.go` test file.

### File structure and conventions

```go
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package <groupname>_test

import (
	"context"
	"<groupname>"
	"testing"

	"github.com/stretchr/testify/require"
)
```

### Client construction

All generated clients have a `New<ClientName>WithNoCredential` constructor. Always use `http://localhost:3000` as the endpoint:

```go
client, err := <groupname>.New<ClientName>WithNoCredential("http://localhost:3000", nil)
require.NoError(t, err)
```

### Test function naming

- Follow the pattern `Test<ClientName>_<MethodName>` or `Test<ClientName><MethodName>`:
  - `TestEmptyClientGetEmpty`
  - `TestBasicClient_CreateOrReplace`
  - `TestHeaderClientDefault`

### Sub-clients

If the generated code has sub-clients (e.g., `NewBytesHeaderClient()`), access them via the parent client:

```go
client, err := bytesgroup.NewBytesClientWithNoCredential("http://localhost:3000", nil)
require.NoError(t, err)
resp, err := client.NewBytesHeaderClient().Base64(context.Background(), []byte("test"), nil)
```

### Assertion patterns

- **No response body (204)**: Use `require.Zero(t, resp)`
- **Response with body**: Use `require.EqualValues(t, expected, resp.<Field>)`
- **Pointer fields**: Use `to.Ptr(value)` from `github.com/Azure/azure-sdk-for-go/sdk/azcore/to`
- **Time values**: Use `require.WithinDuration(t, expected, actual, 0)`
- **Pager results**: Iterate with `pager.More()` / `pager.NextPage(context.Background())`
- **LRO (Long-Running Operations)**: Use `poller.PollUntilDone(context.Background(), nil)`

### Common imports

```go
import (
	"context"
	"<groupname>"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"      // for azcore.ETag, etc.
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"    // for to.Ptr()
	"github.com/stretchr/testify/require"                 // for assertions
)
```

### Complete example

```go
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package emptygroup_test

import (
	"context"
	"emptygroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyClientGetEmpty(t *testing.T) {
	client, err := emptygroup.NewEmptyClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.GetEmpty(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestEmptyClientPutEmpty(t *testing.T) {
	client, err := emptygroup.NewEmptyClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.PutEmpty(context.Background(), emptygroup.EmptyInput{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
```

## Step 6 — Start the Spector mock server and run tests

### Start the Spector server

From the `packages/typespec-go` directory:

```bash
pnpm spector --start
```

This starts the Spector mock server in the background on `http://localhost:3000`. It serves both `http-specs` and `azure-http-specs` scenarios.

### Run the tests

Navigate to the generated module directory and run the tests:

```bash
cd test/<http-specs|azure-http-specs>/<path>/<groupname>
go test -v ./...
```

For example:
```bash
cd test/http-specs/type/model/emptygroup
go test -v ./...
```

### Verify all tests pass

All tests should pass. If a test fails:

1. **Check the mock server is running** — `pnpm spector --start` should have been run first.
2. **Check the endpoint URL** — must be `http://localhost:3000`.
3. **Check the mockapi.ts** — ensure the test matches the expected request/response exactly.
4. **Check generated code** — read `zz_*_client.go` to verify method signatures and parameter types.
5. **Fix and re-run** — iterate until all tests pass.

### Stop the Spector server

After tests pass, stop the server:

```bash
pnpm spector --stop
```

## Notes

- Each test group directory is a **standalone Go module** with its own `go.mod`.
- Generated files are prefixed with `zz_` and should NOT be manually edited.
- Test files (`*_client_test.go`) are hand-written and ARE committed to the repo.
- The `go.mod`, `go.sum`, and `LICENSE.txt` files in the test directory are auto-generated.
- The `testdata/_metadata.json` file tracks emitter version (normalized to `"0.0.0"` by the build script).
- Do NOT edit `zz_version.go` — it is preserved across regenerations for v2+ major version scenarios.
- The mock server runs on `http://localhost:3000` — all test clients must use this as the endpoint.
- Use `require` from `github.com/stretchr/testify` for all assertions (never `assert`).
- There are no shared test utility packages — each test module is fully self-contained.
- Create one test file per generated client: for each `zz_<name>_client.go`, create a corresponding `<name>_client_test.go`. Do NOT consolidate all tests into a single file.
- Only commit: test files (`*_client_test.go`) and changes to `tspcompile.js`. Generated `zz_*` files, `go.mod`, `go.sum`, `LICENSE.txt`, and `testdata/` are also committed (they are NOT gitignored for typespec-go tests).
- When adding options to `tspcompile.js`, fixed options (`module`, `emitter-output-dir`, `file-prefix`) cannot be overridden. Default options (`generate-fakes`, `inject-spans`, `head-as-boolean`, `fix-const-stuttering`) can be overridden per test.
