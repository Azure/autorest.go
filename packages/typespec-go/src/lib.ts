/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { createTypeSpecLibrary, JSONSchemaType, paramMessage } from '@typespec/compiler';

export interface GoEmitterOptions {
  'azcore-version'?: string;
  'containing-module'?: string;

  // When true, unmarshalers will return an error when an unknown field is encountered in the payload.
  'disallow-unknown-fields'?: boolean;

  // Optional prefix to file names. For example, if you set your file prefix to "zzz_", all generated code files will begin with "zzz_".
  'file-prefix'?: string;

  // When true, enables generation of fake servers. The default is false.
  'generate-fakes'?: boolean;
  
  /**
   * Configures invoking `go generate` after emitting the Go code.
   * - The value is an output-relative path to a `.go` file containing `//go:generate` directives.
   * - If Go tools are not on the path, and `go-generate` was specified, then an error is produced.
   */
  'go-generate'?: string;

  // When true, HEAD requests will return a boolean value based on the HTTP status code. The default is false, but will be set to true if --azure-arm is true.
  'head-as-boolean'?: boolean;
  
  // Enables generation of spans for distributed tracing. The default value is set to the value of --azure-arm.
  'inject-spans'?: boolean;
  
  'module'?: string;

  // When true, properties that are untyped (i.e. raw JSON) are exposed as []byte instead of any or map[string]any. The default is false.
  'rawjson-as-bytes'?: boolean;

  // When true, slice elements will not be pointer-to-type. The default is false.
  'slice-elements-byval'?: boolean;

  // Indicates package has a single client. This will omit the Client prefix from options and response types. If multiple clients are detected, an error is returned.
  'single-client'?: boolean;

  // Uses the specified value to remove stuttering from types and funcs instead of the built-in algorithm.
  'stutter'?: string;

  // When true, fix stuttering for const types and their values.
  'fix-const-stuttering'?: boolean;

  /**
   * @deprecated Use 'generate-samples' instead
   */
  'generate-examples'?: boolean;
  
  // When true, the `NewClientFactory` constructor will gather all parameters of clients. When false, the `NewClientFactory` constructor will only gather common parameters of clients. The default value is true.
  'factory-gather-all-params'?: boolean;

  // When true, generate samples.
  'generate-samples'?: boolean;
}

const EmitterOptionsSchema: JSONSchemaType<GoEmitterOptions> = {
  type: 'object',
  additionalProperties: true,
  properties: {
    'azcore-version': {
      type: 'string',
      nullable: true,
      description: 'Semantic version of azcore without the leading \'v\' to use if different from the default version (e.g. 1.2.3).',
    },
    'containing-module': {
      type: 'string',
      nullable: true,
      description: 'The module into which the package is being emitted. Mutually exclusive with module.',
    },
    'disallow-unknown-fields': {
      type: 'boolean',
      nullable: true,
      description: 'When true, unmarshalers will return an error when an unknown field is encountered in the payload. The default is false.',
    },
    'file-prefix': {
      type: 'string',
      nullable: true,
      description: 'Optional prefix to file names. For example, if you set your file prefix to "zzz_", all generated code files will begin with "zzz_".',
    },
    'generate-fakes': {
      type: 'boolean',
      nullable: true,
      description: 'When true, enables generation of fake servers. The default is false.',
    },
    'go-generate': {
      type: 'string',
      nullable: true,
      description: `Path to a post-generation 'go generate' script. The path is relative to the emitter-output-dir.`,
    },
    'head-as-boolean': {
      type: 'boolean',
      nullable: true,
      description: 'When true, HEAD requests will return a boolean value based on the HTTP status code. The default is false.',
    },
    'inject-spans': {
      type: 'boolean',
      nullable: true,
      description: 'Enables generation of spans for distributed tracing. The default is false.',
    },
    'module': {
      type: 'string',
      nullable: true,
      description: 'The module identity to use in go.mod. Mutually exclusive with containing-module.',
    },
    'rawjson-as-bytes': {
      type: 'boolean',
      nullable: true,
      description: 'When true, properties that are untyped (i.e. raw JSON) are exposed as []byte instead of any or map[string]any. The default is false.',
    },
    'slice-elements-byval': {
      type: 'boolean',
      nullable: true,
      description: 'When true, slice elements will not be pointer-to-type. The default is false.',
    },
    'single-client': {
      type: 'boolean',
      nullable: true,
      description: 'Indicates package has a single client. This will omit the Client prefix from options and response types. If multiple clients are detected, an error is returned. The default is false.',
    },
    'stutter': {
      type: 'string',
      nullable: true,
      description: 'Uses the specified value to remove stuttering from types and funcs instead of the built-in algorithm.',
    },
    'fix-const-stuttering': {
      type: 'boolean',
      nullable: true,
      description: 'When true, fix stuttering for `const` types and values. The default is false.',
    },
    'generate-examples': {
      type: 'boolean',
      nullable: true,
      description: 'Deprecated. Use generate-samples instead.',
    },
    'generate-samples': {
      type: 'boolean',
      nullable: true,
      description: 'When true, generate example tests. The default is false.',
    },
    'factory-gather-all-params': {
      type: 'boolean',
      default: true,
      nullable: true,
      description: 'When true, the `NewClientFactory` constructor gathers all parameters. When false, it only gathers common parameters of clients. The default is true.',
    },
  },
  required: [],
};

const libDef = {
  name: '@azure-tools/typespec-go',
  diagnostics: {
    'InternalError': {
      severity: 'error',
      messages: {
        default: paramMessage`The emitter encountered an internal error during preprocessing. Please open an issue at https://github.com/Azure/autorest.go/issues and include the complete error message.\n${'stack'}`
      }
    },
    'InvalidArgument': {
      severity: 'error',
      messages: {
        default: 'Invalid arguments were passed to the emitter.'
      }
    },
    'NameCollision': {
      severity: 'error',
      messages: {
        default: paramMessage`The emitter automatically renamed one or more types which resulted in a type name collision. Please update the client.tsp to rename the type(s) to avoid the collision.\n${'stack'}`
      }
    },
    'UnsupportedTsp': {
      severity: 'error',
      messages: {
        default: paramMessage`The emitter encountered a TypeSpec definition that is currently not supported.\n${'stack'}`
      }
    }
  },
  emitter: {
    options: <JSONSchemaType<GoEmitterOptions>>EmitterOptionsSchema,
  },
} as const;

export const $lib = createTypeSpecLibrary(libDef);
/* eslint-disable-next-line @typescript-eslint/unbound-method */
export const { reportDiagnostic, createStateSymbol, getTracer } = $lib;
