import { Operation } from '@autorest/codemodel';
import { isLROOperation, isMultiRespOperation, isPageableOperation, isSchemaResponse } from '@autorest/go/dist/src/common/helpers';
import { formatParameterTypeName, formatStatusCode, getMethodParameters, getResponseEnvelopeName, getStatusCodes } from '@autorest/go/dist/src/generator/helpers';
import { capitalize, uncapitalize } from '@azure-tools/codegen';
import { values } from 'lodash';
import { Config } from '../common/constant';
import { ParameterOutput } from '../common/model';
import { BaseCodeGenerator } from './baseGenerator';
import { MockTestDataRender } from './mockTestGenerator';

export class FakeDataRender extends MockTestDataRender {
    public renderData(): void {
        super.renderData();
        const modName = this.context.codeModel.language.go!.module;
        if (modName !== 'none' && modName) {
            this.context.importManager.add(modName + '/fake');
        }
    }
}

export class FakeTestCodeGenerator extends BaseCodeGenerator {
    public generateCode(extraParam: Record<string, unknown> = {}): void {
      this.renderAndWrite(this.context.codeModel.testModel.mockTest, 'fakeTest.go.njk', `${this.getFilePrefix(Config.testFilePrefix)}fake_test.go`, extraParam, {
        getParamsValue: (params: Array<ParameterOutput>) => {
          return params
            .map((p) => {
              return p.paramOutput;
            })
            .join(', ');
        },
        getHttpCode: (op: Operation) => {
            const successCodes = new Array<string>();
            if (isMultiRespOperation(op)) {
              for (const response of values(op.responses)) {
                if (!isSchemaResponse(response)) {
                  // the operation contains a mix of schemas and non-schema responses
                  successCodes.push(`${formatStatusCode(response.protocol.http!.statusCodes[0])}`);
                  continue;
                }
                successCodes.push(`${formatStatusCode(response.protocol.http!.statusCodes[0])}`);
              }
            } else {
              for (const statusCode of values(getStatusCodes(op))) {
                successCodes.push(`${formatStatusCode(statusCode)}`);
              }
            }
            for (const successCode of successCodes) {
                return successCode;
            }
        },
        funcMethodReturns: (op: Operation, packageName: string): string => {
            let serverResponse = undefined;
            if (isLROOperation(op)) {
                let respType = `${packageName}.${getResponseEnvelopeName(op)}`;
                if (isPageableOperation(op)) {
                  respType = `azfake.PagerResponder[${packageName}.${getResponseEnvelopeName(op)}]`;
                }
                serverResponse = `resp azfake.PollerResponder[${respType}], errResp azfake.ErrorResponder`;
              } else if (isPageableOperation(op)) {
                // if (op.language.go!.paging.isNextOp) {
                //   // we don't generate a public API for the methods used to advance pages, so skip it here
                // //   continue;
                // }
                serverResponse = `resp azfake.PagerResponder[${packageName}.${getResponseEnvelopeName(op)}]`;
              } else {
                serverResponse = `resp azfake.Responder[${packageName}.${getResponseEnvelopeName(op)}], errResp azfake.ErrorResponder`;
              }

              return `func(${getAPIParametersSig(op, packageName)}) (${serverResponse})`;
        },
        getParameterType: (op: Operation, pkgName: string, paramName: string): string => {
            const methodParams = getMethodParameters(op);
            for (const methodParam of values(methodParams)) {
                if (methodParam.language.go!.name === paramName) {
                    return formatParameterTypeName(methodParam, pkgName);
                }
            }
            return;
        },
        capitalize: (name: string): string => {
            return capitalize(name);
        },
        uncapitalize: (name: string): string => {
          return uncapitalize(name);
        },
        cutClientSuffix: (client: string): string => {
          return client.substring(0,client.lastIndexOf('Client'));
        },
      });
    }
  }

function getAPIParametersSig(op: Operation, pkgName?: string): string {
    const methodParams = getMethodParameters(op);
    const params = new Array<string>();
    if (!isPageableOperation(op) || isLROOperation(op)) {
        params.push('ctx context.Context');
    }
    for (const methodParam of values(methodParams)) {
        params.push(`${uncapitalize(methodParam.language.go!.name)} ${formatParameterTypeName(methodParam, pkgName)}`);
    }
    return params.join(', ');
}