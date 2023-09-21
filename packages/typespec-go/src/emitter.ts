import {
  EmitContext
} from '@typespec/compiler';

import {
  mkdir,
  writeFile
} from 'fs/promises';

import {
  GoEmitterOptions
} from './lib.js';

export async function $onEmit(context: EmitContext<GoEmitterOptions>) {
  await mkdir(context.emitterOutputDir, {recursive: true});
  await writeFile(`${context.emitterOutputDir}/main.go`, '// TODO: make code\n');
}
