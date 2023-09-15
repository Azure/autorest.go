import { ArraySchema, DictionarySchema, Operation, Parameter, Schema, SchemaType } from '@autorest/codemodel';
import { isLROOperation, isMultiRespOperation, isPageableOperation, isSchemaResponse } from '@autorest/go/dist/src/transform/helpers';
import { capitalize, uncapitalize } from '@azure-tools/codegen';
import { values } from 'lodash';
import { Config } from '../common/constant';
import { ParameterOutput } from '../common/model';
import { getMethodParameters, getResponseEnvelopeName } from '../util/codegenBridge';
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
        getExampleParams: (params: Array<ParameterOutput>) => {
          return params
            .map((p) => {
              if (p.paramName === 'ctx' || p.paramName === 'options') {
                return p.paramOutput;
              }
              return 'example'+ capitalize(p.paramName);
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
        capitalize: capitalize,
        uncapitalize: uncapitalize,
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

function formatParameterTypeName(param: Parameter, pkgName?: string): string {
  const typeName = formatTypeName(param.schema, pkgName);
  if (param.required) {
    return typeName;
  }
  return `*${typeName}`;
}

function formatTypeName(schema: Schema, pkgName?: string): string {
  const typeName = schema.language.go!.name;
  if (!pkgName) {
    return typeName;
  }

  // if not an array/dictionary, just prepend the package name
  if (schema.type !== SchemaType.Array && schema.type !== SchemaType.Dictionary) {
    if (schema.type === SchemaType.Choice || schema.type === SchemaType.SealedChoice || schema.type === SchemaType.Object) {
      return `${pkgName}.${typeName}`;
    }
    return typeName;
  }

  // for array/dictionary, we need to splice the package name into the correct location
  const elementType = unwrapSchemaType(schema);
  if (elementType.type === SchemaType.Choice || elementType.type === SchemaType.SealedChoice || elementType.type === SchemaType.Object) {
    // don't use a * prefix on the type name as --slice-elements-byval will remove it for slices
    return typeName.replace(`${elementType.language.go!.name}`, `${pkgName}.${elementType.language.go!.name}`);
  }

  return typeName;
}

// recursively gets the element schema from an array/dictionary
function unwrapSchemaType(schema: Schema): Schema {
  if (schema.type === SchemaType.Array) {
    return unwrapSchemaType((<ArraySchema>schema).elementType);
  } else if (schema.type === SchemaType.Dictionary) {
    return unwrapSchemaType((<DictionarySchema>schema).elementType);
  }
  return schema;
}

function formatStatusCode(statusCode: string): string {
  switch (statusCode) {
    // 1xx
    case '100':
      return 'http.StatusContinue';
    case '101':
      return 'http.StatusSwitchingProtocols';
    case '102':
      return 'http.StatusProcessing';
    case '103':
      return 'http.StatusEarlyHints';
    // 2xx
    case '200':
      return 'http.StatusOK';
    case '201':
      return 'http.StatusCreated';
    case '202':
      return 'http.StatusAccepted';
    case '203':
      return 'http.StatusNonAuthoritativeInfo';
    case '204':
      return 'http.StatusNoContent';
    case '205':
      return 'http.StatusResetContent';
    case '206':
      return 'http.StatusPartialContent';
    case '207':
      return 'http.StatusMultiStatus';
    case '208':
      return 'http.StatusAlreadyReported';
    case '226':
      return 'http.StatusIMUsed';
    // 3xx
    case '300':
      return 'http.StatusMultipleChoices';
    case '301':
      return 'http.StatusMovedPermanently';
    case '302':
      return 'http.StatusFound';
    case '303':
      return 'http.StatusSeeOther';
    case '304':
      return 'http.StatusNotModified';
    case '305':
      return 'http.StatusUseProxy';
    case '307':
      return 'http.StatusTemporaryRedirect';
    // 4xx
    case '400':
      return 'http.StatusBadRequest';
    case '401':
      return 'http.StatusUnauthorized';
    case '402':
      return 'http.StatusPaymentRequired';
    case '403':
      return 'http.StatusForbidden';
    case '404':
      return 'http.StatusNotFound';
    case '405':
      return 'http.StatusMethodNotAllowed';
    case '406':
      return 'http.StatusNotAcceptable';
    case '407':
      return 'http.StatusProxyAuthRequired';
    case '408':
      return 'http.StatusRequestTimeout';
    case '409':
      return 'http.StatusConflict';
    case '410':
      return 'http.StatusGone';
    case '411':
      return 'http.StatusLengthRequired';
    case '412':
      return 'http.StatusPreconditionFailed';
    case '413':
      return 'http.StatusRequestEntityTooLarge';
    case '414':
      return 'http.StatusRequestURITooLong';
    case '415':
      return 'http.StatusUnsupportedMediaType';
    case '416':
      return 'http.StatusRequestedRangeNotSatisfiable';
    case '417':
      return 'http.StatusExpectationFailed';
    case '418':
      return 'http.StatusTeapot';
    case '421':
      return 'http.StatusMisdirectedRequest';
    case '422':
      return 'http.StatusUnprocessableEntity';
    case '423':
      return 'http.StatusLocked';
    case '424':
      return 'http.StatusFailedDependency';
    case '425':
      return 'http.StatusTooEarly';
    case '426':
      return 'http.StatusUpgradeRequired';
    case '428':
      return 'http.StatusPreconditionRequired';
    case '429':
      return 'http.StatusTooManyRequests';
    case '431':
      return 'http.StatusRequestHeaderFieldsTooLarge';
    case '451':
      return 'http.StatusUnavailableForLegalReasons';
    // 5xx
    case '500':
      return 'http.StatusInternalServerError';
    case '501':
      return 'http.StatusNotImplemented';
    case '502':
      return 'http.StatusBadGateway';
    case '503':
      return 'http.StatusServiceUnavailable';
    case '504':
      return 'http.StatusGatewayTimeout ';
    case '505':
      return 'http.StatusHTTPVersionNotSupported';
    case '506':
      return 'http.StatusVariantAlsoNegotiates';
    case '507':
      return 'http.StatusInsufficientStorage';
    case '508':
      return 'http.StatusLoopDetected';
    case '510':
      return 'http.StatusNotExtended';
    case '511':
      return 'http.StatusNetworkAuthenticationRequired';
    default:
      throw new Error(`unhandled status code ${statusCode}`);
  }
}

function getStatusCodes(op: Operation): Array<string> {
  // concat all status codes that return the same schema into one array.
  // this is to support operations that specify multiple response codes
  // that return the same schema (or no schema).
  let statusCodes = new Array<string>();
  for (const resp of values(op.responses)) {
    statusCodes = statusCodes.concat(resp.protocol.http?.statusCodes);
  }
  if (statusCodes.length === 0) {
    // if the operation defines no status codes (which is non-conformant)
    // then add 200, 201, 202, and 204 to the list.  this is to accomodate
    // some quirky tests in the test server.
    // TODO: https://github.com/Azure/autorest.go/issues/659
    statusCodes = ['200', '201', '202', '204'];
  }
  return statusCodes;
}