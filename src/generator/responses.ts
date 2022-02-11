/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { comment } from '@azure-tools/codegen';
import { CodeModel, ObjectSchema, Property } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { commentLength, PagerInfo, PollerInfo } from '../common/helpers';
import { contentPreamble, discriminatorFinalResponse, emitPoller, getFinalResponseEnvelopeName, sortAscending } from './helpers';
import { ImportManager } from './imports';
import { generateStruct, StructDef, StructMethod } from './structs';

// Creates the content in response_types.go
export async function generateResponses(session: Session<CodeModel>): Promise<string> {
  let text = await contentPreamble(session);
  const imports = new ImportManager();
  const responseEnvelopes = <Array<ObjectSchema>>session.model.language.go!.responseEnvelopes;
  if (responseEnvelopes.length === 0) {
    return '';
  }
  const structs = new Array<StructDef>();
  for (const respEnv of values(responseEnvelopes)) {
    const respType = generateStruct(imports, respEnv.language.go!, respEnv.properties);
    if (respEnv.language.go!.isLRO) {
      generatePollUntilDoneForResponse(respType, <boolean>session.model.language.go!.azureARM);
      generateResumeForResponse(respType, session.model.language.go!.openApiType === 'arm', imports);
    }
    generateUnmarshallerForResponeEnvelope(respType);
    structs.push(respType);
  }
  text += imports.text();
  structs.sort((a: StructDef, b: StructDef) => { return sortAscending(a.Language.name, b.Language.name) });
  for (const struct of values(structs)) {
    text += struct.discriminator();
    text += struct.text();
    struct.Methods.sort((a: StructMethod, b: StructMethod) => { return sortAscending(a.name, b.name) });
    for (const method of values(struct.Methods)) {
      if (method.desc.length > 0) {
        text += `${comment(method.desc, '// ', undefined, commentLength)}\n`;
      }
      text += method.text;
    }
  }
  return text;
}

// check if the response envelope requires an unmarshaller
function generateUnmarshallerForResponeEnvelope(structDef: StructDef) {
  // if the response envelope contains a discriminated type we need an unmarshaller
  let found = false;
  for (const prop of values(structDef.Properties)) {
    if (prop.isDiscriminator) {
      found = true;
      break;
    }
  }
  if (!found) {
    return;
  }
  const receiver = structDef.receiverName();
  let unmarshaller = `func (${receiver} *${structDef.Language.name}) UnmarshalJSON(data []byte) error {\n`;
  // add a custom unmarshaller to the response envelope
  let type = '';
  for (const prop of values(structDef.Properties)) {
    if (prop.isDiscriminator) {
      type = prop.schema.language.go!.discriminatorInterface;
      break;
    }
  }
  if (type === '') {
    throw new Error(`failed to the discriminated type field for response envelope ${structDef.Language.name}`);
  }
  unmarshaller += `\tres, err := unmarshal${type}(data)\n`;
  unmarshaller += '\tif err != nil {\n';
  unmarshaller += '\t\treturn err\n';
  unmarshaller += '\t}\n';
  unmarshaller += `\t${receiver}.${type} = res\n`;
  unmarshaller += '\treturn nil\n';
  unmarshaller += '}\n\n';
  structDef.Methods.push({ name: 'UnmarshalJSON', desc: `UnmarshalJSON implements the json.Unmarshaller interface for type ${structDef.Language.name}.`, text: unmarshaller });
}

function generatePollUntilDoneForResponse(structDef: StructDef, isAzureARM: boolean) {
  const pagedResponse = (<PollerInfo>structDef.Language.pollerInfo).op.language.go!.pageableType;
  const respType = getResponseType(<PollerInfo>structDef.Language.pollerInfo);
  let pollUntilDone = `func (l ${structDef.Language.name}) PollUntilDone(ctx context.Context, freq time.Duration) (`;
  if (pagedResponse) {
    pollUntilDone += '*';
  }
  pollUntilDone += `${respType}, error) {\n`;
  pollUntilDone += `\trespType := `;
  if (pagedResponse) {
    pollUntilDone += '&';
  }
  pollUntilDone += `${respType}{}\n`;
  const finalRespEnv = <ObjectSchema>structDef.Language.pollerInfo.op.language.go!.finalResponseEnv;
  const resultProp = <Property>finalRespEnv.language.go!.resultProp;
  if (resultProp) {
    let current = '';
    if (pagedResponse) {
      current = '.current';
    }
    pollUntilDone += `\t_, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType${current}${discriminatorFinalResponse(finalRespEnv)})\n`;
  } else {
    // the operation doesn't return a model
    pollUntilDone += `\t_, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)\n`;
  }
  pollUntilDone += '\tif err != nil {\n';
  pollUntilDone += '\t\treturn respType, err\n';
  pollUntilDone += '\t}\n';
  if (pagedResponse) {
    pollUntilDone += '\trespType.client = l.Poller.client\n';
  }
  pollUntilDone += '\treturn respType, nil\n';
  pollUntilDone += '}\n\n';
  let desc = 'PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.\nfreq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.';
  if (isAzureARM) {
    desc += '\nA good starting value is 30 seconds. Note that some resources might benefit from a different value.';
  }
  structDef.Methods.push({
    name: 'PollUntilDone',
    desc: desc,
    text: pollUntilDone });
}

function generateResumeForResponse(structDef: StructDef, isARM: boolean, imports: ImportManager) {
  const pollerInfo = <PollerInfo>structDef.Language.pollerInfo;
  const clientName = pollerInfo.op.language.go!.clientName;
  const apiMethod = pollerInfo.op.language.go!.name;
  let resume = `func (l *${structDef.Language.name}) Resume(ctx context.Context, client *${clientName}, token string) error {\n`;
  if (isARM) {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime', 'armruntime');
    resume += `\tpt, err := armruntime.NewPollerFromResumeToken("${clientName}.${apiMethod}", token, client.pl)\n`;
  } else {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
    resume += `\tpt, err := runtime.NewPollerFromResumeToken("${clientName}.${apiMethod}",token, client.pl)\n`;
  }
  resume += '\tif err != nil {\n';
  resume += `\t\treturn err\n`;
  resume += '\t}\n';
  resume += `\tpoller := ${emitPoller(pollerInfo.op)}`;
  resume += '\t_, err = poller.Poll(ctx)\n';
  resume += '\tif err != nil {\n';
  resume += `\t\treturn err\n`;
  resume += '\t}\n';
  resume += `\tl.Poller = poller\n`;
  resume += `\treturn nil\n`;
  resume += '}\n\n';
  structDef.Methods.push({ name: 'Resume', desc: `Resume rehydrates a ${structDef.Language.name} from the provided client and resume token.`, text: resume });
}

function getResponseType(poller: PollerInfo): string {
  // check for pager must come first
  if (poller.op.language.go!.pageableType) {
    return (<PagerInfo>poller.op.language.go!.pageableType).name;
  }
  return getFinalResponseEnvelopeName(poller.op);
}
