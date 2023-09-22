import {
  createTypeSpecLibrary,
  JSONSchemaType
} from '@typespec/compiler';

export interface GoEmitterOptions {
    'basic-setup-py'?: boolean;
    'package-version'?: string;
    'package-name'?: string;
    'output-dir'?: string;
    'package-mode'?: string;
    'package-pprint-name'?: string;
    'head-as-boolean'?: boolean;
    'models-mode'?: string;
    'debug'?: boolean;
}

const EmitterOptionsSchema: JSONSchemaType<GoEmitterOptions> = {
  type: 'object',
  additionalProperties: true,
  properties: {
    'basic-setup-py': { type: 'boolean', nullable: true },
    'package-version': { type: 'string', nullable: true },
    'package-name': { type: 'string', nullable: true },
    'output-dir': { type: 'string', nullable: true },
    'package-mode': { type: 'string', nullable: true },
    'package-pprint-name': { type: 'string', nullable: true },
    'head-as-boolean': { type: 'boolean', nullable: true },
    'models-mode': { type: 'string', nullable: true },
    'debug': { type: 'boolean', nullable: true },
  },
  required: [],
};

const libDef = {
  name: '@azure-tools/typespec-go',
  diagnostics: {},
  emitter: {
    options: <JSONSchemaType<GoEmitterOptions>>EmitterOptionsSchema,
  },
} as const;

export const $lib = createTypeSpecLibrary(libDef);
export const { reportDiagnostic, createStateSymbol, getTracer } = $lib;
