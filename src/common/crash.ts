import { serialize } from '@azure-tools/codegen';
import { Host, Session, Channel } from '@autorest/extension-base';
import { CodeModel } from '@autorest/codemodel';

export function writeCodeModelOnCrash (host: Host, session: Session<CodeModel> | null)
{
    try
    {
      if (session !== null) {
        host.Message({
         Channel: Channel.Fatal,
         Text: `autorest.go has crashed.\nPlease file an issue at https://github.com/Azure/autorest.go/issues/new. \n\nAttach the written 'code-model-v4.yaml' or the original swagger along with all output provided so we can reproduce your error.\n`
        })
        host.WriteFile('code-model-v4.yaml', serialize(session.model), undefined, 'source-file-go');
      }
    }
    catch {}
}
