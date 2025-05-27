# Getting started

This repo contains multiple packages used to generate [Azure SDK's for Go](https://github.com/Azure/azure-sdk-for-go):

- The `autorest.go` package generates client code using [autorest](https://github.com/Azure/autorest).
- The `typespec-go` package generates client code from [typespec](https://github.com/microsoft/typespec).
- The `codegen.go` package contains code used common by both `autorest.go` and `typespec-go`

This guide outlines the getting started steps to contributing to these generators.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Step 1: Clone the repo](#step-1-clone-the-repo)
- [Step 2: Build the code](#step-2-build-the-code)
- [Step 3: Regenerate tests and samples](#step-3-regenerate-tests-and-samples)
  - [For autorest.go](#for-autorestgo)
  - [For typespec-go](#for-typespec-go)
- [Step 4: Test your changes](#step-4-test-your-changes)
  - [Debug](#debug)
  - [Built in commands](#built-in-commands)
- [Step 5: Make a PR](#step-5-make-a-pr)

## Prerequisites

- Install [Node.js](https://nodejs.org/download/)
- Install [pnpm](https://pnpm.io/installation/)
- Install [Go](https://go.dev/doc/install)

## Step 1: Clone the repo

To set up your local development environment, we recommend forking [this repo then cloning](https://github.com/Azure/azure-sdk/blob/main/docs/policies/repobranching.md).

```terminal
git clone https://github.com/<your-github-username>/autorest.go.git
```

## Step 2: Build the code

Once you have the code locally, you can build it.

First, install all dependencies.

```terminal
pnpm install
```

Then, build the code.

To build a specific package, navigate to that directory and run the command from there.

```terminal
pnpm build
```

## Step 3: Regenerate tests and samples

After making changes, build the code again, then run a regeneration command to see how your change has affected client code generation.

### For autorest.go

From the `autorest.go/packages/autorest.go` directory, run the following command:

```terminal
pnpm regenerate
```

To regenerate a specific test:

```terminal
pnpm regenerate --filter=TestName
```

### For typespec-go

From the `autorest.go/packages/typespec-go` directory, run the following command:

```terminal
pnpm tspcompile
```

To regenerate a specific test:

```terminal
pnpm tspcompile --filter=TestName
```

## Step 4: Test your changes

Verify changes made result in the output you expect.

For typespec-go, you can run tests using spector.

```terminal
pnpm spector --start
~ run tests~
pnpm spector --stop
```

### Debug

To debug the code generator:

1. Set a break point
2. In the TypeScript debug terminal in VSCode, run one of the regeneration commands from step 3

### Built in commands

There are a number of custom pnpm commands to help with development. See the `.scripts` folder for more. Add the -w switch when executing the script from a directory other than the repo root.

NOTE: by default, the script will run against the entire repo. To limit its scope, pass a directory argument. E.g. when working from the /packages/autorest.go directory, to execute modtidy for only go.mod files under /packages/autorest.go, execute pnpm -w modtidy $pwd

To run `go build` and `go vet` on every generated module:

```terminal
pnpm -w buildvet
```

To run `go mod tidy` on every generated module:

```terminal
pnpm -w modtidy
```

## Step 5: Make a PR

Once you're satistied with your changes, it's time to make a PR in the [repo](https://github.com/Azure/autorest.go/pulls).

Before you do, make sure to:

1. Format your code using the Prettier configuration file in the root of the repo
2. Dont't forget to rebuild and regenerate everything before pushing your changes
