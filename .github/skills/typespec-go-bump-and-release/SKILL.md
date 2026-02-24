---
name: typespec-go-bump-and-release
description: Upgrade tsp toolset dependencies for the @azure-tools/typespec-go package in Azure/autorest.go repo. Use when user wants to bump, update, or upgrade the TypeSpec toolset dependencies (e.g., @typespec/compiler, @typespec/http, @azure-tools/typespec-client-generator-core, @azure-tools/azure-http-specs), create a release PR, or publish a new version of typespec-go.
---

# TypeSpec Go Bump and Release

Upgrade the TypeSpec toolset dependencies for `@azure-tools/typespec-go` and create a release PR for the Azure/autorest.go repository.

## Context Variables

- **CURRENT_DATE**: Current date in YYYY-MM-DD format (e.g., "2025-01-01")

## Workflow

### Step 1: Check and Upgrade TypeSpec Dependencies

1. Change work directory to `packages/typespec-go` folder.
2. Run `ncu -u` on `package.json` to check if there are any TypeSpec-related packages that need upgrading. Focus on packages matching:
   - `@typespec/*`
   - `@azure-tools/typespec-*`
   - `@azure-tools/azure-http-specs`
3. Review the proposed upgrades and apply them.
4. Check and update the corresponding versions in `peerDependencies` of `package.json` to match the upgraded versions. For each upgraded package that also appears in `peerDependencies`, update its version range accordingly.

### Step 2: Install and Build

Run `pnpm install` from the repo root, and then run `pnpm build` from the `packages/typespec-go` directory:

```bash
cd <repo-root>
pnpm install
cd packages/typespec-go
pnpm build
```

This ensures `pnpm-lock.yaml` is updated and the project builds correctly.

### Step 3: Run TypeSpec Compile

From the `packages/typespec-go` directory, run:

```bash
pnpm tspcompile
```

**Notes:**
- This command needs several minutes to finish; wait until it completes.
- All changed files after `pnpm tspcompile` must be committed.

### Step 4: Update CHANGELOG.md

Update `packages/typespec-go/CHANGELOG.md` with a new entry. The format should be:

```markdown
## x.x.x ({{CURRENT_DATE}})

### Other Changes

* Updated to the latest tsp toolset.
```

**Version rules:**
1. Run `git diff` to inspect the changes.
2. If any changes contain new features or breaking changes, bump the version accordingly (major for breaking, minor for features, patch for fixes/other).
3. If there is already an unreleased version entry at the top, append the changelog entry into that existing version instead of creating a new one.
4. Update the version in `packages/typespec-go/package.json` to match.

### Step 5: Stage, Commit and Push

```bash
git add -A && git commit -m "upgrade tsp toolset" && git push origin HEAD
```

### Step 6: Create PR

If no existing PR exists for the current branch, create a new pull request targeting the main branch.
