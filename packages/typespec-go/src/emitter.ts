/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { GoEmitterOptions } from './lib.js';
import { Adapter, ExternalError } from './tcgcadapter/adapter.js';
import { AdapterError } from './tcgcadapter/errors.js';
import { CodeModelError } from '../../codemodel.go/src/errors.js';
import * as codegen from '../../codegen.go/src/index.js';
import { existsSync, opendirSync, unlinkSync, readFileSync } from 'fs';
import { mkdir, readFile, writeFile } from 'fs/promises';
import * as path from 'path';
import { DiagnosticSeverity, EmitContext, NoTarget } from '@typespec/compiler';
import 'source-map-support/register.js';
import { reportDiagnostic } from './lib.js';
import { execSync } from 'child_process';

export async function $onEmit(context: EmitContext<GoEmitterOptions>) {
  try {
    const adapter = await Adapter.create(context);
    const codeModel = adapter.tcgcToGoCodeModel();

    await mkdir(context.emitterOutputDir, { recursive: true });

    // clean up existing generated Go files
    cleanupGeneratedFiles(context.emitterOutputDir);

    let filePrefix: string | undefined;
    if (context.options['file-prefix']) {
      filePrefix = context.options['file-prefix'];
      // if a file prefix was specified, ensure it's properly snaked
      if (filePrefix[filePrefix.length - 1] !== '_') {
        filePrefix += '_';
      }
    }

    const emitter = new codegen.Emitter(codeModel, {
      exists: (name: string) => {
        return Promise.resolve(existsSync(`${context.emitterOutputDir}/${name}`));
      },
      read: (name: string) => readFile(`${context.emitterOutputDir}/${name}`, 'utf8'),
      write: async (name: string, content: string) => {
        await mkdir(path.dirname(`${context.emitterOutputDir}/${name}`), { recursive: true });
        return writeFile(`${context.emitterOutputDir}/${name}`, content);
      }
    }, { filePrefix });
    await emitter.emit();
    await emitter.emitCloudConfig();
    await emitter.emitExamples();
    await emitter.emitLicenseFile();
    await emitter.emitMetadataFile();

    const goGenerateFile = context.options['go-generate'];
    const goGenerateFileExists = goGenerateFile ? existsSync(`${context.emitterOutputDir}/${goGenerateFile}`) : false;

    if (goGenerateFile && !goGenerateFileExists) {
      // go-generate was specified but we didn't find the file, so error and exit
      context.program.reportDiagnostic({
        code: 'gogenerate',
        severity: 'error',
        message: `the go-generate file wasn't found. the complete path is ${context.emitterOutputDir}/${goGenerateFile}`,
        target: NoTarget,
      });

      // don't continue so the state of the SDK can be inspected without any additional changes
      return;
    }

    // probe to see if Go tools are on the path
    try {
      execSync('go version', { stdio: ['ignore', 'ignore', 'ignore'] });
    } catch {
      // if the transforms file exists and we don't have Go
      // on the path then make this an error as it means we
      // expect to transform the generated code but were unable
      // to do so.
      let severity: DiagnosticSeverity = 'warning';
      let message = 'skip executing post emitter steps (is go on the path?)';
      if (goGenerateFileExists) {
        severity = 'error';
        message = 'unable to execute post emitter transformations due to missing go tool (is go on the path?)';
      }

      context.program.reportDiagnostic({
        code: 'GoVersion',
        severity: severity,
        message: message,
        target: NoTarget,
      });

      // no Go tools available so exit
      return;
    }

    // if we have a post-generation transforms file then "go generate" it
    if (goGenerateFileExists) {
      try {
        execSync(`go generate ${goGenerateFile}`, { cwd: context.emitterOutputDir, encoding: 'ascii' });
      } catch (err) {
        context.program.reportDiagnostic({
          code: 'gogenerate',
          severity: 'error',
          message: (<Error>err).message,
          target: NoTarget,
        });

        // don't continue so the state of the SDK can be inspected without any additional changes
        return;
      }
    }

    // format after transforms in case any formatting gets munged
    try {
      execSync('gofmt -s -w .', { cwd: context.emitterOutputDir, encoding: 'ascii' });
    } catch (err) {
      context.program.reportDiagnostic({
        code: 'gofmt',
        severity: 'error',
        message: (<Error>err).message,
        target: NoTarget,
      });

      return;
    }

    // now go mod tidy
    try {
      execSync('go mod tidy', { cwd: context.emitterOutputDir, encoding: 'ascii' });
    } catch (err) {
      context.program.reportDiagnostic({
        code: 'gomodtidy',
        severity: 'error',
        message: (<Error>err).message,
        target: NoTarget,
      });

      return;
    }
  } catch (error) {
    if (error instanceof AdapterError) {
      reportDiagnostic(context.program, {
        code: error.code,
        target: error.target,
        format: {
          stack: error.stack ? truncateStack(error.stack, 'tcgcToGoCodeModel') : 'Stack trace unavailable\n',
        },
      });
    } else if (error instanceof CodeModelError) {
      reportDiagnostic(context.program, {
        code: error.code,
        target: NoTarget,
        format: {
          stack: error.stack ? truncateStack(error.stack, 'tcgcToGoCodeModel') : 'Stack trace unavailable\n',
        },
      });
    } else if (error instanceof codegen.CodegenError) {
      reportDiagnostic(context.program, {
        code: error.code,
        target: NoTarget,
        format: {
          stack: error.stack ? truncateStack(error.stack, 'generate(') : 'Stack trace unavailable\n',
        },
      });
    } else if (error instanceof ExternalError) {
      // we don't want to throw in this case as that will
      // make it appear as if the emitter crashed. just
      // exit so the diagnostic error isn't lost in the noise
      return;
    } else {
      throw error;
    }
  }
}

/**
 * drop frames after the specified frame.
 *
 * @param stack the stack to truncate
 * @returns the truncated stack
 */
function truncateStack(stack: string, finalFrame: string): string {
  const lines = stack.split('\n');
  stack = '';
  for (const line of lines) {
    stack += `${line}\n`;
    if (line.includes(finalFrame)) {
      break;
    }
  }
  return stack;
}

/**
 * Clean up existing generated Go files in the output directory.
 * Removes any .go files that contain the Microsoft code generator comment.
 * 
 * exported for testing purposes only.
 *
 * @param outputDir the directory to clean up
 */
export function cleanupGeneratedFiles(outputDir: string) {
  if (!existsSync(outputDir)) {
    return;
  }
  const dir = opendirSync(outputDir);
  while (true) {
    const dirEnt = dir.readSync();
    if (dirEnt === null) {
      break;
    }
    if (dirEnt.isFile() && dirEnt.name.endsWith('.go')) {
      const content = readFileSync(dir.path + '/' + dirEnt.name, 'utf8');
      if (codegen.doNotEditRegex.test(content)) {
        unlinkSync(dir.path + '/' + dirEnt.name);
      }
    }
  }
  dir.closeSync();
  cleanupGeneratedFiles(outputDir + '/fake');
}
