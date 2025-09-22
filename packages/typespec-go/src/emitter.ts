/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { GoEmitterOptions } from './lib.js';
import { tcgcToGoCodeModel } from './tcgcadapter/adapter.js';
import { AdapterError } from './tcgcadapter/errors.js';
import { generateClientFactory } from '../../codegen.go/src/clientFactory.js';
import { generateCloudConfig } from '../../codegen.go/src/cloudConfig.js';
import { generateConstants } from '../../codegen.go/src/constants.js';
import { generateExamples } from '../../codegen.go/src/example.js';
import { generateGoModFile } from '../../codegen.go/src/gomod.js';
import { generateInterfaces } from '../../codegen.go/src/interfaces.js';
import { generateModels } from '../../codegen.go/src/models.js';
import { generateOperations } from '../../codegen.go/src/operations.js';
import { generateOptions } from '../../codegen.go/src/options.js';
import { generatePolymorphicHelpers } from '../../codegen.go/src/polymorphics.js';
import { generateResponses } from '../../codegen.go/src/responses.js';
import { generateTimeHelpers } from '../../codegen.go/src/time.js';
import { generateServers } from '../../codegen.go/src/fake/servers.js';
import { generateServerFactory } from '../../codegen.go/src/fake/factory.js';
import { generateXMLAdditionalPropsHelpers } from '../../codegen.go/src/xmlAdditionalProps.js';
import { generateMetadataFile } from '../../codegen.go/src/metadata.js';
import { generateVersionInfo } from '../../codegen.go/src/version.js';
import { CodeModelError } from '../../codemodel.go/src/errors.js';
import { existsSync, opendirSync, unlinkSync, readFileSync } from 'fs';
import { mkdir, readFile, writeFile } from 'fs/promises';
import { DiagnosticSeverity, EmitContext, NoTarget } from '@typespec/compiler';
import 'source-map-support/register.js';
import { reportDiagnostic } from './lib.js';
import { CodegenError } from '../../codegen.go/src/errors.js';
import { execSync } from 'child_process';
import { doNotEditRegex } from '../../codegen.go/src/helpers.js';

export async function $onEmit(context: EmitContext<GoEmitterOptions>) {
  try {
    await generate(context);

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
      execSync('gofmt -w .', { cwd: context.emitterOutputDir, encoding: 'ascii' });
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
    } else if (error instanceof CodegenError) {
      reportDiagnostic(context.program, {
        code: error.code,
        target: NoTarget,
        format: {
          stack: error.stack ? truncateStack(error.stack, 'generate(') : 'Stack trace unavailable\n',
        },
      });
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
      if (doNotEditRegex.test(content)) {
        unlinkSync(dir.path + '/' + dirEnt.name);
      }
    }
  }
  dir.closeSync();
  cleanupGeneratedFiles(outputDir + '/fake');
}

async function generate(context: EmitContext<GoEmitterOptions>) {
  const codeModel = await tcgcToGoCodeModel(context);
  await mkdir(context.emitterOutputDir, { recursive: true });

  // clean up existing generated Go files
  cleanupGeneratedFiles(context.emitterOutputDir);

  // don't overwrite an existing go.mod file, update it if required
  const goModFile = `${context.emitterOutputDir}/go.mod`;
  let existingGoMod: string | undefined;
  if (existsSync(goModFile)) {
    existingGoMod = (await readFile(goModFile)).toString();
  }
  const gomod = await generateGoModFile(codeModel, existingGoMod);
  if (gomod.length > 0) {
    await writeFile(goModFile, gomod);
  }

  const metadata = generateMetadataFile(codeModel);
  if (metadata.length > 0) {
    const metedataDir = context.emitterOutputDir + '/testdata';
    await mkdir(metedataDir, { recursive: true });
    await writeFile(`${metedataDir}/_metadata.json`, metadata);
  }

  let filePrefix = '';
  if (context.options['file-prefix']) {
    filePrefix = context.options['file-prefix'];
    // if a file prefix was specified, ensure it's properly snaked
    if (filePrefix[filePrefix.length - 1] !== '_') {
      filePrefix += '_';
    }
  }

  const clientFactory = await generateClientFactory(codeModel);
  if (clientFactory.length > 0) {
    await writeFile(`${context.emitterOutputDir}/${filePrefix}client_factory.go`, clientFactory);
  }

  const constants = await generateConstants(codeModel);
  if (constants.length > 0) {
    await writeFile(`${context.emitterOutputDir}/${filePrefix}constants.go`, constants);
  }

  const interfaces = await generateInterfaces(codeModel);
  if (interfaces.length > 0) {
    await writeFile(`${context.emitterOutputDir}/${filePrefix}interfaces.go`, interfaces);
  }

  const models = await generateModels(codeModel);
  if (models.models.length > 0) {
    await writeFile(`${context.emitterOutputDir}/${filePrefix}models.go`, models.models);
  }
  if (models.serDe.length > 0) {
    await writeFile(`${context.emitterOutputDir}/${filePrefix}models_serde.go`, models.serDe);
  }

  const operations = await generateOperations(codeModel);
  for (const op of operations) {
    let fileName = op.name.toLowerCase();
    // op.name is the client name, e.g. FooClient.
    // insert a _ before Client, i.e. Foo_Client
    // if the name isn't simply Client.
    if (fileName !== 'client') {
      fileName = fileName.substring(0, fileName.length - 6) + '_client';
    }
    await writeFile(`${context.emitterOutputDir}/${filePrefix}${fileName}.go`, op.content);
  }

  if (codeModel.options.generateExamples) {
    const examples = await generateExamples(codeModel);
    for (const example of examples) {
      let fileName = example.name.toLowerCase();
      // op.name is the client name, e.g. FooClient.
      // insert a _ before Client, i.e. Foo_Client
      // if the name isn't simply Client.
      // and insert _example_test at the end.
      if (fileName !== 'client') {
        fileName = fileName.substring(0, fileName.length - 6) + '_client';
      }
      fileName += '_example_test';
      await writeFile(`${context.emitterOutputDir}/${filePrefix}${fileName}.go`, example.content);
    }
  }

  const options = await generateOptions(codeModel);
  if (options.length > 0) {
    await writeFile(`${context.emitterOutputDir}/${filePrefix}options.go`, options);
  }

  const polymorphics = await generatePolymorphicHelpers(codeModel);
  if (polymorphics.length > 0) {
    await writeFile(`${context.emitterOutputDir}/${filePrefix}polymorphic_helpers.go`, polymorphics);
  }

  const responses = await generateResponses(codeModel);
  if (responses.responses.length > 0) {
    await writeFile(`${context.emitterOutputDir}/${filePrefix}responses.go`, responses.responses);
  }
  if (responses.serDe.length > 0) {
    await writeFile(`${context.emitterOutputDir}/${filePrefix}responses_serde.go`, responses.serDe);
  }

  const timeHelpers = await generateTimeHelpers(codeModel);
  for (const helper of timeHelpers) {
    await writeFile(`${context.emitterOutputDir}/${filePrefix}${helper.name.toLowerCase()}.go`, helper.content);
  }

  // don't overwrite an existing version.go file
  const versionGo = await generateVersionInfo(codeModel);
  const versionGoFileName = `${context.emitterOutputDir}/${filePrefix}version.go`;
  if (versionGo.length > 0 && !existsSync(versionGoFileName)) {
    await writeFile(versionGoFileName, versionGo);
  }

  const xmlAddlProps = await generateXMLAdditionalPropsHelpers(codeModel);
  if (xmlAddlProps.length > 0) {
    await writeFile(`${context.emitterOutputDir}/${filePrefix}xml_helper.go`, xmlAddlProps);
  }

  const cloudConfig = generateCloudConfig(codeModel);
  if (cloudConfig.length > 0) {
    await writeFile(`${context.emitterOutputDir}/${filePrefix}cloud_config.go`, cloudConfig);
  }

  if (codeModel.options.generateFakes) {
    const serverContent = await generateServers(codeModel);
    if (serverContent.servers.length > 0) {
      const fakesDir = context.emitterOutputDir + '/fake';
      await mkdir(fakesDir, { recursive: true });
      for (const op of serverContent.servers) {
        let fileName = op.name.toLowerCase();
        // op.name is the server name, e.g. FooServer.
        // insert a _ before Server, i.e. Foo_Server
        // if the name isn't simply Server.
        if (fileName !== 'server') {
          fileName = fileName.substring(0, fileName.length - 6) + '_server';
        }
        await writeFile(`${fakesDir}/${filePrefix}${fileName}.go`, op.content);
      }

      const serverFactory = generateServerFactory(codeModel);
      if (serverFactory.length > 0) {
        await writeFile(`${fakesDir}/${filePrefix}server_factory.go`, serverFactory);
      }

      await writeFile(`${fakesDir}/${filePrefix}internal.go`, serverContent.internals);

      const timeHelpers = await generateTimeHelpers(codeModel, 'fake');
      for (const helper of timeHelpers) {
        await writeFile(`${fakesDir}/${filePrefix}${helper.name.toLowerCase()}.go`, helper.content);
      }

      const polymorphics = await generatePolymorphicHelpers(codeModel, 'fake');
      if (polymorphics.length > 0) {
        await writeFile(`${fakesDir}/${filePrefix}polymorphic_helpers.go`, polymorphics);
      }
    }
  }
}
