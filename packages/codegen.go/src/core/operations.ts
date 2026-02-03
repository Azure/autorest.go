/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../../codemodel.go/src/index.js';
import * as naming from '../../../naming.go/src/naming.js';
import * as helpers from './helpers.js';
import { ImportManager } from './imports.js';
import { CodegenError } from './errors.js';

// represents the generated content for an operation group
export class OperationGroupContent {
  readonly name: string;
  readonly content: string;

  constructor(name: string, content: string) {
    this.name = name;
    this.content = content;
  }
}

/**
 * Creates the content for all the *_client.go files.
 * 
 * @param pkg contains the package content
 * @param target the codegen target for the module
 * @param options the emitter options
 * @returns the text for the files or the empty string
 */
export function generateOperations(pkg: go.PackageContent, target: go.CodeModelType, options: go.Options): Array<OperationGroupContent> {
  // generate protocol operations
  const operations = new Array<OperationGroupContent>();
  if (pkg.clients.length === 0) {
    return operations;
  }
  const azureARM = target === 'azure-arm';
  for (const client of pkg.clients) {
    // the list of packages to import
    const imports = new ImportManager(pkg);
    if (client.methods.length > 0) {
      // add standard imports for clients with methods.
      // clients that are purely hierarchical (i.e. having no APIs) won't need them.
      imports.add('net/http');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/policy');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
    }

    imports.add(azureARM  ? 'github.com/Azure/azure-sdk-for-go/sdk/azcore/arm' : 'github.com/Azure/azure-sdk-for-go/sdk/azcore');

    // generate client type

    let clientText = helpers.formatDocComment(client.docs);
    clientText += '// Don\'t use this type directly, use ';
    if (client.instance?.kind === 'constructable' && client.instance.constructors.length === 1) {
      clientText += `${client.instance.constructors[0].name}() instead.\n`;
    } else if (client.parent) {
      // find the accessor method
      let accessorMethod: string | undefined;
      for (const clientAccessor of client.parent.clientAccessors) {
        if (clientAccessor.returns === client) {
          accessorMethod = clientAccessor.name;
          break;
        }
      }
      if (!accessorMethod) {
        throw new CodegenError('InternalError', `didn't find accessor method for client ${client.name} on parent client ${client.parent.name}`);
      }
      clientText += `[${client.parent.name}.${accessorMethod}] instead.\n`;
    } else {
      clientText += 'a constructor function instead.\n';
    }
    clientText += `type ${client.name} struct {\n`;
    clientText += `\tinternal *${azureARM ? 'arm' : 'azcore'}.Client\n`;

    // check for any optional host params
    const optionalParams = new Array<go.ClientParameter>();

    const isParamPointer = function(param: go.ClientParameter): boolean {
      // for client params, only optional and flag types are passed by pointer
      return param.style === 'flag' || param.style === 'optional';
    };

    // now emit any client params (non parameterized host params case)
    if (client.parameters.length > 0) {
      const addedGroups = new Set<string>();
      for (const clientParam of client.parameters) {
        if (go.isLiteralParameter(clientParam.style)) {
          continue;
        }
        if (clientParam.group) {
          if (!addedGroups.has(clientParam.group.groupName)) {
            clientText += `\t${naming.uncapitalize(clientParam.group.groupName)} ${!isParamPointer(clientParam) ? '' : '*'}${clientParam.group.groupName}\n`;
            addedGroups.add(clientParam.group.groupName);
          }
          continue;
        }
        clientText += `\t${clientParam.name} `;
        if (!isParamPointer(clientParam)) {
          clientText += `${go.getTypeDeclaration(clientParam.type, client.pkg)}\n`;
        } else {
          clientText += `${helpers.formatParameterTypeName(client.pkg, clientParam)}\n`;
        }
        if (!go.isRequiredParameter(clientParam.style)) {
          optionalParams.push(clientParam);
        }
      }
    }

    // end of client definition
    clientText += '}\n\n';

    clientText += generateConstructors(client, target, imports);

    // generate client accessors and operations
    let opText = '';
    for (const clientAccessor of client.clientAccessors) {
      imports.addForType(clientAccessor.returns);
      const subClientDecl = go.getTypeDeclaration(clientAccessor.returns, pkg);
      opText += helpers.formatDocComment(clientAccessor.docs);
      opText += `func (client *${client.name}) ${clientAccessor.name}(${getAPIParametersSig(clientAccessor, imports)}) *${subClientDecl} {\n`;
      opText += `\treturn &${subClientDecl}{\n`;
      const initFields = new Array<string>('internal: client.internal');
      // propagate all client params
      for (const param of clientAccessor.parameters) {
        // by convention, the client accessor params have the
        // same name as their corresponding client fields.
        initFields.push(`${param.name}: ${param.name}`);
      }

      // accessor params and client fields are mutually exclusive
      // so we don't need to worry about potentials for duplication.
      for (const param of client.parameters) {
        if (go.isLiteralParameter(param.style)) {
          continue;
        }
        initFields.push(`${param.name}: client.${param.name}`);
      }

      initFields.sort();
      for (const initField of initFields) {
        opText += `\t\t${initField},\n`;
      }
      opText += '\t}\n}\n\n';
    }

    const nextPageMethods = new Array<go.NextPageMethod>();
    for (const method of client.methods) {
      // protocol creation can add imports to the list so
      // it must be done before the imports are written out
      if (go.isLROMethod(method)) {
        // generate Begin method
        opText += generateLROBeginMethod(method, imports, options);
      }
      opText += generateOperation(method, imports, options);
      opText += createProtocolRequest(azureARM, method, imports);
      if (method.kind !== 'lroMethod') {
        // LRO responses are handled elsewhere, with the exception of pageable LROs
        opText += createProtocolResponse(method, imports);
      }
      if ((method.kind === 'lroPageableMethod' || method.kind === 'pageableMethod') && method.nextPageMethod && !nextPageMethods.includes(method.nextPageMethod)) {
        // track the next page methods to generate as multiple operations can use the same next page operation
        nextPageMethods.push(method.nextPageMethod);
      }
    }

    for (const method of nextPageMethods) {
      opText += createProtocolRequest(azureARM, method, imports);
    }

    // stitch it all together
    let text = helpers.contentPreamble(pkg);
    text += imports.text();
    text += clientText;
    text += opText;
    operations.push(new OperationGroupContent(client.name, text));
  }
  return operations;
}

/**
 * generates all modeled client constructors and client options types.
 * if there are no client constructors, the empty string is returned.
 * 
 * @param client the client for which to generate constructors and the client options type
 * @param imports the import manager currently in scope
 * @returns the client constructor code or the empty string
 */
function generateConstructors(client: go.Client, type: go.CodeModelType, imports: ImportManager): string {
  if (client.instance?.kind !== 'constructable') {
    return '';
  }

  const clientOptions = client.instance.options;

  let ctorText = '';

  if (clientOptions.kind === 'clientOptions') {
    // for non-ARM, the options type will always be a parameter group
    ctorText += `// ${clientOptions.name} contains the optional values for creating a [${client.name}].\n`;
    ctorText += `type ${clientOptions.name} struct {\n\tazcore.ClientOptions\n`;
    for (const param of clientOptions.parameters) {
      if (go.isAPIVersionParameter(param)) {
        // we use azcore.ClientOptions.APIVersion
        continue;
      }
      ctorText += helpers.formatDocCommentWithPrefix(naming.ensureNameCase(param.name), param.docs);
      if (go.isClientSideDefault(param.style)) {
        if (!param.docs.description && !param.docs.summary) {
          ctorText += '\n';
        }
        ctorText += `\t${helpers.comment(`The default value is ${helpers.formatLiteralValue(param.style.defaultValue, false)}`, '// ')}.\n`;
      }
      ctorText += `\t${naming.ensureNameCase(param.name)} *${go.getTypeDeclaration(param.type, client.pkg)}\n`;
    }
    ctorText += '}\n\n';
  }

  for (const constructor of client.instance.constructors) {
    const ctorParams = new Array<string>();
    const paramDocs = new Array<string>();

    // ctor params can also be present in the supplemental endpoint parameters
    const consolidatedCtorParams = new Array<go.ClientParameter>();
    if (client.instance.endpoint) {
      consolidatedCtorParams.push(client.instance.endpoint.parameter);
      if (client.instance.endpoint.supplemental) {
        consolidatedCtorParams.push(...client.instance.endpoint.supplemental.parameters);
      }
    }

    for (const param of helpers.sortClientParameters(constructor.parameters, type)) {
      if (!consolidatedCtorParams.includes(param)) {
        consolidatedCtorParams.push(param);
      }
    }

    for (const ctorParam of consolidatedCtorParams) {
      if (!go.isRequiredParameter(ctorParam.style)) {
        // param is part of the options group
        continue;
      }
      imports.addForType(ctorParam.type);
      ctorParams.push(`${ctorParam.name} ${helpers.formatParameterTypeName(client.pkg, ctorParam)}`);
      if (ctorParam.docs.summary || ctorParam.docs.description) {
        paramDocs.push(helpers.formatCommentAsBulletItem(ctorParam.name, ctorParam.docs));
      }
    }

    const emitProlog = function(optionsTypeName: string, tokenAuth: boolean, plOpts?: string): string {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
      let bodyText = `\tif options == nil {\n\t\toptions = &${optionsTypeName}{}\n\t}\n`;
      let apiVersionConfig = '';
      // check if there's an api version parameter
      let apiVersionParam: go.HeaderScalarParameter | go.PathScalarParameter | go.QueryScalarParameter | go.URIParameter | undefined;
      for (const param of consolidatedCtorParams) {
        switch (param.kind) {
          case 'headerScalarParam':
          case 'pathScalarParam':
          case 'queryScalarParam':
          case 'uriParam':
            if (param.isApiVersion) {
              apiVersionParam = param;
            }
        }
      }

      if (tokenAuth) {
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud');
        imports.add('fmt');
        imports.add('reflect');
        bodyText += '\tif reflect.ValueOf(options.Cloud).IsZero() {\n';
        bodyText += '\t\toptions.Cloud = cloud.AzurePublic\n\t}\n';
        bodyText += '\tc, ok := options.Cloud.Services[ServiceName]\n';
        bodyText += '\tif !ok {\n';
        bodyText += '\t\treturn nil, fmt.Errorf("provided Cloud field is missing configuration for %s", ServiceName)\n';
        bodyText += '\t} else if c.Audience == "" {\n';
        bodyText += '\t\treturn nil, fmt.Errorf("provided Cloud field is missing Audience for %s", ServiceName)\n\t}\n';
      }

      if (apiVersionParam) {
        let location: string;
        let name: string | undefined;
        switch (apiVersionParam.kind) {
          case 'headerScalarParam':
            location = 'Header';
            name = apiVersionParam.headerName;
            break;
          case 'pathScalarParam':
          case 'uriParam':
            location = 'Path';
            // name isn't used for the path case
            break;
          case 'queryScalarParam':
            location = 'QueryParam';
            name = apiVersionParam.queryParameter;
            break;
        }

        if (name) {
          name = `\n\t\t\tName: "${name}",`;
        } else {
          name = '';
        }

        apiVersionConfig = `\n\t\tAPIVersion: runtime.APIVersionOptions{${name}\n\t\t\tLocation: runtime.APIVersionLocation${location},\n\t\t},`;
        if (!plOpts) {
          apiVersionConfig += '\n';
        }
      }
      bodyText += `\tcl, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{${apiVersionConfig}${plOpts ?? ''}}, &options.ClientOptions)\n`;
      return bodyText;
    };

    // check if there's a credential parameter
    let credentialParam: go.ClientCredentialParameter | undefined;
    for (const param of constructor.parameters) {
      if (param.kind === 'credentialParam') {
        credentialParam = param;
        break;
      }
    }

    let prolog: string;
    if (credentialParam) {
      switch (credentialParam.type.kind) {
        case 'tokenCredential':
          imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
          paramDocs.push(helpers.formatCommentAsBulletItem('credential', { summary: 'used to authorize requests. Usually a credential from azidentity.' }));
          switch (clientOptions.kind) {
            case 'clientOptions': {
              imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/policy');
              const tokenPolicyOpts = '&policy.BearerTokenOptions{\n\t\t\tInsecureAllowCredentialWithHTTP: options.InsecureAllowCredentialWithHTTP,\n\t\t}';
              // we assume a single scope. this is enforced when adapting the data from tcgc
              const tokenPolicy = `\n\t\tPerCall: []policy.Policy{\n\t\truntime.NewBearerTokenPolicy(credential, []string{c.Audience + "${helpers.splitScope(credentialParam.type.scopes[0]).scope}"}, ${tokenPolicyOpts}),\n\t\t},\n`;
              prolog = emitProlog(go.getTypeDeclaration(clientOptions, client.pkg), true, tokenPolicy);
              break;
            }
            case 'armClientOptions':
              // this is the ARM case
              imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/arm');
              prolog = '\tcl, err := arm.NewClient(moduleName, moduleVersion, credential, options)\n';
              break;
          }
          break;
      }
    } else {
      prolog = emitProlog(go.getTypeDeclaration(clientOptions, client.pkg), false);
    }

    // add client options last
    ctorParams.push(`options ${helpers.formatParameterTypeName(client.pkg, clientOptions)}`);
    paramDocs.push(helpers.formatCommentAsBulletItem('options', { summary: 'Contains optional client configuration. Pass nil to accept the default values.' }));

    ctorText += `// ${constructor.name} creates a new instance of ${client.name} with the specified values.\n`;
    for (const doc of paramDocs) {
      ctorText += doc;
    }

    ctorText += `func ${constructor.name}(${ctorParams.join(', ')}) (*${client.name}, error) {\n`;
    ctorText += prolog;
    ctorText += '\tif err != nil {\n';
    ctorText += '\t\treturn nil, err\n';
    ctorText += '\t}\n';

    // handle any client-side defaults
    if (clientOptions.kind === 'clientOptions') {
      for (const param of clientOptions.parameters) {
        if (go.isClientSideDefault(param.style)) {
          let name: string;
          if (go.isAPIVersionParameter(param)) {
            name = 'APIVersion';
          } else {
            name = naming.ensureNameCase(param.name);
          }
          ctorText += `\t${param.name} := ${helpers.formatLiteralValue(param.style.defaultValue, false)}\n`;
          ctorText += `\tif options.${name} != ${helpers.zeroValue(param)} {\n\t\t${param.name} = ${helpers.star(param.byValue)}options.${name}\n\t}\n`;
        }
      }
    }

    // construct the supplemental path and join it to the endpoint
    if (client.instance.endpoint?.supplemental) {
      // the endpoint param is always the first ctor param
      const endpointParam = client.instance.constructors[0].parameters[0];
      if (client.instance.endpoint.supplemental.parameters.length > 0) {
        imports.add('strings');
        ctorText += `\thost := "${client.instance.endpoint.supplemental.path}"\n`;
        for (const param of client.instance.endpoint.supplemental.parameters) {
          ctorText += `\thost = strings.ReplaceAll(host, "{${param.uriPathSegment}}", ${helpers.formatValue(param.name, param.type, imports)})\n`;
        }
        ctorText += `\t${endpointParam.name} = runtime.JoinPaths(${endpointParam.name}, host)\n`;
      } else {
        // there are no params for the supplemental host, so just append it
        ctorText += `\t${endpointParam.name} = runtime.JoinPaths(${endpointParam.name}, "${client.instance.endpoint.supplemental.path}")\n`;
      }
    }

    // construct client literal
    let clientVar = 'client';
    // ensure clientVar doesn't collide with any params
    for (const param of consolidatedCtorParams) {
      if (param.name === clientVar) {
        clientVar = naming.ensureNameCase(client.name, true);
        break;
      }
    }

    ctorText += `\t${clientVar} := &${client.name}{\n`;
    // NOTE: we don't enumerate consolidatedCtorParams here
    // as any supplemental endpoint params are ephemeral and
    // consumed during client construction.
    for (const parameter of client.parameters) {
      if (go.isLiteralParameter(parameter.style)) {
        continue;
      }
      // each client field will have a matching parameter with the same name
      ctorText += `\t\t${parameter.name}: ${parameter.name},\n`;
    }
    ctorText += '\tinternal: cl,\n';
    ctorText += '\t}\n';
    ctorText += `\treturn ${clientVar}, nil\n`;
    ctorText += '}\n\n';
  }

  return ctorText;
}

// use this to generate the code that will help process values returned in response headers
function formatHeaderResponseValue(headerResp: go.HeaderScalarResponse | go.HeaderMapResponse, imports: ImportManager, respObj: string, zeroResp: string): string {
  // dictionaries are handled slightly different so we do that first
  if (headerResp.kind === 'headerMapResponse') {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
    imports.add('strings');
    const headerPrefix = headerResp.headerName;
    let text = '\tfor hh := range resp.Header {\n';
    text += `\t\tif len(hh) > len("${headerPrefix}") && strings.EqualFold(hh[:len("${headerPrefix}")], "${headerPrefix}") {\n`;
    text += `\t\t\tif ${respObj}.${headerResp.fieldName} == nil {\n`;
    text += `\t\t\t\t${respObj}.${headerResp.fieldName} = map[string]*string{}\n`;
    text += '\t\t\t}\n';
    text += `\t\t\t${respObj}.${headerResp.fieldName}[hh[len("${headerPrefix}"):]] = to.Ptr(resp.Header.Get(hh))\n`;
    text += '\t\t}\n';
    text += '\t}\n';
    return text;
  }

  let text = `\tif val := resp.Header.Get("${headerResp.headerName}"); val != "" {\n`;
  let name = naming.uncapitalize(headerResp.fieldName);
  let byRef = '&';
  switch (headerResp.type.kind) {
    case 'constant':
      text += `\t\t${respObj}.${headerResp.fieldName} = (*${headerResp.type.name})(&val)\n`;
      text += '\t}\n';
      return text;
    case 'encodedBytes':
      // a base-64 encoded value in string format
      imports.add('encoding/base64');
      text += `\t\t${name}, err := base64.${helpers.formatBytesEncoding(headerResp.type.encoding)}Encoding.DecodeString(val)\n`;
      byRef = '';
      break;
    case 'literal':
      text += `\t\t${respObj}.${headerResp.fieldName} = &val\n`;
      text += '\t}\n';
      return text;
    case 'scalar':
      imports.add('strconv');
      switch (headerResp.type.type) {
        case 'bool':
          text += `\t\t${name}, err := strconv.ParseBool(val)\n`;
          break;
        case 'float32':
          text += `\t\t${name}32, err := strconv.ParseFloat(val, 32)\n`;
          text += `\t\t${name} := float32(${name}32)\n`;
          break;
        case 'float64':
          text += `\t\t${name}, err := strconv.ParseFloat(val, 64)\n`;
          break;
        case 'int32':
          text += `\t\t${name}32, err := strconv.ParseInt(val, 10, 32)\n`;
          text += `\t\t${name} := int32(${name}32)\n`;
          break;
        case 'int64':
          text += `\t\t${name}, err := strconv.ParseInt(val, 10, 64)\n`;
          break;
        default:
          throw new CodegenError('InternalError', `unhandled scalar type ${headerResp.type.type}`);
      }
      break;
    case 'string':
      text += `\t\t${respObj}.${headerResp.fieldName} = &val\n`;
      text += '\t}\n';
      return text;
    case 'time':
      imports.add('time');
      switch (headerResp.type.format) {
        case 'RFC1123':
        case 'RFC3339':
          text += `\t\t${name}, err := time.Parse(${headerResp.type.format === 'RFC1123' ? helpers.RFC1123Format : helpers.RFC3339Format}, val)\n`;
          break;
        case 'PlainDate':
          text += `\t\t${name}, err := time.Parse(${helpers.plainDateFormat}, val)\n`;
          break;
        case 'PlainTime':
          text += `\t\t${name}, err := time.Parse(${helpers.plainTimeFormat}, val)\n`;
          break;
        case 'Unix':
          imports.add('strconv');
          imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/to');
          text += '\t\tsec, err := strconv.ParseInt(val, 10, 64)\n';
          name = 'to.Ptr(time.Unix(sec, 0))';
          byRef = '';
          break;
      }
  }

  // NOTE: only cases that required parsing will fall through to here
  text += '\t\tif err != nil {\n';
  text += `\t\t\treturn ${zeroResp}, err\n`;
  text += '\t\t}\n';
  text += `\t\t${respObj}.${headerResp.fieldName} = ${byRef}${name}\n`;
  text += '\t}\n';
  return text;
}

function getZeroReturnValue(method: go.MethodType, apiType: 'api' | 'op' | 'handler'): string {
  let returnType = `${method.returns.name}{}`;
  if (go.isLROMethod(method)) {
    if (apiType === 'api' || apiType === 'op') {
      // the api returns a *Poller[T]
      // the operation returns an *http.Response
      returnType = 'nil';
    }
  }
  return returnType;
}

// Helper function to generate nil checks for a dotted path
function generateNilChecks(path: string, prefix: string = 'page'): string {
  const segments = path.split('.');
  const checks: string[] = [];
  
  for (let i = 0; i < segments.length; i++) {
    const currentPath = [prefix, ...segments.slice(0, i + 1)].join('.');
    checks.push(`${currentPath} != nil`);
  }
  
  return checks.join(' && ');
}

function emitPagerDefinition(method: go.LROPageableMethod | go.PageableMethod, imports: ImportManager, options: go.Options): string {
  imports.add('context');
  let text = `runtime.NewPager(runtime.PagingHandler[${method.returns.name}]{\n`;
  text += `\t\tMore: func(page ${method.returns.name}) bool {\n`;
  // there is no advancer for single-page pagers
  if (method.nextLinkName) {
    const nilChecks = generateNilChecks(method.nextLinkName);
    text += `\t\t\treturn ${nilChecks} && len(*page.${method.nextLinkName}) > 0\n`;
    text += '\t\t},\n';
  } else {
    text += '\t\t\treturn false\n';
    text += '\t\t},\n';
  }
  text += `\t\tFetcher: func(ctx context.Context, page *${method.returns.name}) (${method.returns.name}, error) {\n`;
  const reqParams = helpers.getCreateRequestParameters(method);
  if (options.generateFakes) {
    text += `\t\tctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "${method.receiver.type.name}.${fixUpMethodName(method)}")\n`;
  }
  if (method.nextLinkName) {
    let nextLinkVar: string;
    if (method.kind === 'pageableMethod') {
      text += '\t\t\tnextLink := ""\n';
      nextLinkVar = 'nextLink';
      text += '\t\t\tif page != nil {\n';
      text += `\t\t\t\tnextLink = *page.${method.nextLinkName}\n\t\t\t}\n`;
    } else {
      nextLinkVar = `*page.${method.nextLinkName}`;
    }
    text += `\t\t\tresp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), ${nextLinkVar}, func(ctx context.Context) (*policy.Request, error) {\n`;
    text += `\t\t\t\treturn client.${method.naming.requestMethod}(${reqParams})\n\t\t\t}, `;
    // nextPageMethod might be absent in some cases, see https://github.com/Azure/autorest/issues/4393
    if (method.nextPageMethod) {
      const nextOpParams = helpers.getCreateRequestParametersSig(method.nextPageMethod).split(',');
      // keep the parameter names from the name/type tuples and find nextLink param
      for (let i = 0; i < nextOpParams.length; ++i) {
        const paramName = nextOpParams[i].trim().split(' ')[0];
        const paramType = nextOpParams[i].trim().split(' ')[1];
        if (paramName.startsWith('next') && paramType === 'string') {
          nextOpParams[i] = 'encodedNextLink';
        } else {
          nextOpParams[i] = paramName;
        }
      }
      // add a definition for the nextReq func that uses the nextLinkOperation
      text += '&runtime.FetcherForNextLinkOptions{\n\t\t\t\tNextReq: func(ctx context.Context, encodedNextLink string) (*policy.Request, error) {\n';
      text += `\t\t\t\t\treturn client.${method.nextPageMethod.name}(${nextOpParams.join(', ')})\n\t\t\t\t},\n\t\t\t})\n`;
    } else if (method.nextLinkVerb !== 'get') {
      text += `&runtime.FetcherForNextLinkOptions{\n\t\t\t\tHTTPVerb: http.Method${naming.capitalize(method.nextLinkVerb)},\n\t\t\t})\n`;
    } else {
      text += 'nil)\n';
    }
    text += `\t\t\tif err != nil {\n\t\t\t\treturn ${method.returns.name}{}, err\n\t\t\t}\n`;
    text += `\t\t\treturn client.${method.naming.responseMethod}(resp)\n`;
    text += '\t\t\t},\n';
  } else {
    // this is the singular page case, no fetcher helper required
    text += `\t\t\treq, err := client.${method.naming.requestMethod}(${reqParams})\n`;
    text += '\t\t\tif err != nil {\n';
    text += `\t\t\t\treturn ${method.returns.name}{}, err\n`;
    text += '\t\t\t}\n';
    text += '\t\t\tresp, err := client.internal.Pipeline().Do(req)\n';
    text += '\t\t\tif err != nil {\n';
    text += `\t\t\t\treturn ${method.returns.name}{}, err\n`;
    text += '\t\t\t}\n';
    text += '\t\t\tif !runtime.HasStatusCode(resp, http.StatusOK) {\n';
    text += `\t\t\t\treturn ${method.returns.name}{}, runtime.NewResponseError(resp)\n`;
    text += '\t\t\t}\n';
    text += `\t\t\treturn client.${method.naming.responseMethod}(resp)\n`;
    text += '\t\t},\n';
  }
  if (options.injectSpans) {
    text += '\t\tTracer: client.internal.Tracer(),\n';
  }
  text += '\t})\n';
  return text;
}

function genApiVersionDoc(apiVersions: Array<string>): string {
  if (apiVersions.length === 0) {
    return '';
  }
  return `//\n// Generated from API version ${apiVersions.join(', ')}\n`;
}

function genRespErrorDoc(method: go.MethodType): string {
  if (!(method.returns.result?.kind === 'headAsBooleanResult') && !go.isPageableMethod(method)) {
    // when head-as-boolean is enabled, no error is returned for 4xx status codes.
    // pager constructors don't return an error
    return '// If the operation fails it returns an *azcore.ResponseError type.\n';
  }
  return '';
}

/**
 * returns the receiver definition for a client
 * 
 * @param receiver the receiver for which to emit the definition
 * @returns the receiver definition
 */
function getClientReceiverDefinition(receiver: go.Receiver<go.Client>): string {
  return `(${receiver.name} ${receiver.byValue ? '' : '*'}${receiver.type.name})`;
}

function generateOperation(method: go.MethodType, imports: ImportManager, options: go.Options): string {
  const params = getAPIParametersSig(method, imports);
  const returns = generateReturnsInfo(method, 'op');
  let methodName = method.name;
  if(method.kind === 'pageableMethod') {
    methodName = fixUpMethodName(method);
  }
  let text = '';
  const respErrDoc = genRespErrorDoc(method);
  const apiVerDoc = genApiVersionDoc(method.apiVersions);
  if (method.docs.summary || method.docs.description) {
    text += helpers.formatDocCommentWithPrefix(methodName, method.docs);
  } else if (respErrDoc.length > 0 || apiVerDoc.length > 0) {
    // if the method has no doc comment but we're adding other
    // doc comments, add an empty method name comment. this preserves
    // existing behavior and makes the docs look better overall.
    text += `// ${methodName} -\n`;
  }
  text += respErrDoc;
  text += apiVerDoc;
  if (go.isLROMethod(method)) {
    methodName = method.naming.internalMethod;
  } else {
    for (const param of helpers.getMethodParameters(method)) {
      text += helpers.formatCommentAsBulletItem(param.name, param.docs);
    }
  }
  text += `func ${getClientReceiverDefinition(method.receiver)} ${methodName}(${params}) (${returns.join(', ')}) {\n`;
  const reqParams = helpers.getCreateRequestParameters(method);
  if (method.kind === 'pageableMethod') {
    text += '\treturn ';
    text += emitPagerDefinition(method, imports, options);
    text += '}\n\n';
    return text;
  }
  text += '\tvar err error\n';
  let operationName = `"${method.receiver.type.name}.${fixUpMethodName(method)}"`;
  if (options.generateFakes && options.injectSpans) {
    text += `\tconst operationName = ${operationName}\n`;
    operationName = 'operationName';
  }
  if (options.generateFakes) {
    text += `\tctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, ${operationName})\n`;
  }
  if (options.injectSpans) {
    text += `\tctx, endSpan := runtime.StartSpan(ctx, ${operationName}, client.internal.Tracer(), nil)\n`;
    text += '\tdefer func() { endSpan(err) }()\n';
  }
  const zeroResp = getZeroReturnValue(method, 'op');
  text += `\treq, err := client.${method.naming.requestMethod}(${reqParams})\n`;
  text += '\tif err != nil {\n';
  text += `\t\treturn ${zeroResp}, err\n`;
  text += '\t}\n';
  text += '\thttpResp, err := client.internal.Pipeline().Do(req)\n';
  text += '\tif err != nil {\n';
  text += `\t\treturn ${zeroResp}, err\n`;
  text += '\t}\n';
  text += `\tif !runtime.HasStatusCode(httpResp, ${helpers.formatStatusCodes(method.httpStatusCodes)}) {\n`;
  text += '\t\terr = runtime.NewResponseError(httpResp)\n';
  text += `\t\treturn ${zeroResp}, err\n`;
  text += '\t}\n';
  // HAB with headers response is handled in protocol responder
  if (method.returns.result?.kind === 'headAsBooleanResult' && method.returns.headers.length === 0) {
    text += `\treturn ${method.returns.name}{${method.returns.result.fieldName}: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil\n`;
  } else {
    if (go.isLROMethod(method)) {
      text += '\treturn httpResp, nil\n';
    } else if (needsResponseHandler(method)) {
      // also cheating here as at present the only param to the responder is an http.Response
      text += `\tresp, err := client.${method.naming.responseMethod}(httpResp)\n`;
      text += '\treturn resp, err\n';
    } else if (method.returns.result?.kind === 'binaryResult') {
      text += `\treturn ${method.returns.name}{${method.returns.result.fieldName}: httpResp.Body}, nil\n`;
    } else {
      text += `\treturn ${method.returns.name}{}, nil\n`;
    }
  }
  text += '}\n\n';
  return text;
}

function createProtocolRequest(azureARM: boolean, method: go.MethodType | go.NextPageMethod, imports: ImportManager): string {
  let name = method.name;
  if (method.kind !== 'nextPageMethod') {
    name = method.naming.requestMethod;
  }

  for (const param of method.parameters) {
    if (param.location !== 'method' || !go.isRequiredParameter(param.style)) {
      continue;
    }
    imports.addForType(param.type);
  }

  const returns = ['*policy.Request', 'error'];
  let text = `${helpers.comment(name, '// ')} creates the ${method.name} request.\n`;
  text += `func ${getClientReceiverDefinition(method.receiver)} ${name}(${helpers.getCreateRequestParametersSig(method)}) (${returns.join(', ')}) {\n`;

  const hostParams = new Array<go.URIParameter>();
  for (const parameter of method.receiver.type.parameters) {
    if (parameter.kind === 'uriParam') {
      hostParams.push(parameter);
    }
  }

  let hostParam: string;
  if (azureARM) {
    hostParam = 'client.internal.Endpoint()';
  } else if (method.receiver.type.instance?.kind === 'templatedHost') {
    imports.add('strings');
    // we have a templated host
    text += `\thost := "${method.receiver.type.instance.path}"\n`;
    // get all the host params on the client
    for (const hostParam of hostParams) {
      text += `\thost = strings.ReplaceAll(host, "{${hostParam.uriPathSegment}}", ${helpers.formatValue(`client.${hostParam.name}`, hostParam.type, imports)})\n`;
    }
    // check for any method local host params
    for (const param of method.parameters) {
      if (param.location === 'method' && param.kind === 'uriParam') {
        text += `\thost = strings.ReplaceAll(host, "{${param.uriPathSegment}}", ${helpers.formatValue(helpers.getParamName(param), param.type, imports)})\n`;
      }
    }
    hostParam = 'host';
  } else if (hostParams.length === 1) {
    // simple parameterized host case
    hostParam = 'client.' + hostParams[0].name;
  } else {
    throw new CodegenError('InternalError', `no host or endpoint defined for method ${method.receiver.type.name}.${method.name}`);
  }

  const methodParamGroups = helpers.getMethodParamGroups(method);
  const hasPathParams = methodParamGroups.pathParams.length > 0;

  // storage needs the client.u to be the source-of-truth for the full path.
  // however, swagger requires that all operations specify a path, which is at odds with storage.
  // to work around this, storage specifies x-ms-path paths with path params but doesn't
  // actually reference the path params (i.e. no params with which to replace the tokens).
  // so, if a path contains tokens but there are no path params, skip emitting the path.
  const pathStr = method.httpPath;
  const pathContainsParms = pathStr.includes('{');
  if (hasPathParams || (!pathContainsParms && pathStr.length > 1)) {
    // there are path params, or the path doesn't contain tokens and is not "/" so emit it
    text += `\turlPath := "${method.httpPath}"\n`;
    hostParam = `runtime.JoinPaths(${hostParam}, urlPath)`;
  }

  // helper to build nil checks for param groups
  const emitParamGroupCheck = function (param: go.MethodParameter): string {
    if (!param.group) {
      throw new CodegenError('InternalError', `emitParamGroupCheck called for ungrouped parameter ${param.name}`);
    }
    let client = '';
    if (param.location === 'client') {
      client = 'client.';
    }
    const paramGroupName = naming.uncapitalize(param.group.name);
    let optionalParamGroupCheck = `${client}${paramGroupName} != nil && `;
    if (param.group.required) {
      optionalParamGroupCheck = '';
    }
    return `\tif ${optionalParamGroupCheck}${client}${paramGroupName}.${naming.capitalize(param.name)} != nil {\n`;
  };

  if (hasPathParams) {
    // swagger defines path params, emit path and replace tokens
    imports.add('strings');
    // replace path parameters
    for (const pp of methodParamGroups.pathParams) {
      let paramValue: string;
      let optionalPathSep = false;
      if (pp.style !== 'optional') {
        // emit check to ensure path param isn't an empty string
        if (pp.kind === 'pathScalarParam') {
          const choiceIsString = function (type: go.PathScalarParameterType): boolean {
            return type.kind === 'constant' && type.type === 'string';
          };
          // we only need to do this for params that have an underlying type of string
          if ((pp.type.kind === 'string' || choiceIsString(pp.type)) && !pp.omitEmptyStringCheck) {
            const paramName = helpers.getParamName(pp);
            imports.add('errors');
            text += `\tif ${paramName} == "" {\n`;
            text += `\t\treturn nil, errors.New("parameter ${paramName} cannot be empty")\n`;
            text += '\t}\n';
          }
        }

        paramValue = helpers.formatParamValue(pp, imports);

        // for collection-based path params, we emit the empty check
        // after calling helpers.formatParamValue as that will have the
        // var name that contains the slice.
        if (pp.kind === 'pathCollectionParam') {
          const paramName = helpers.getParamName(pp);
          const joinedParamName = `${paramName}Param`;
          text += `\t${joinedParamName} := ${paramValue}\n`;
          imports.add('errors');
          text += `\tif len(${joinedParamName}) == 0 {\n`;
          text += `\t\treturn nil, errors.New("parameter ${paramName} cannot be empty")\n`;
          text += '\t}\n';
          paramValue = joinedParamName;
        }
      } else {
        // param isn't required, so emit a local var with
        // the correct default value, then populate it with
        // the optional value when set.
        paramValue = `optional${naming.capitalize(pp.name)}`;
        text += `\t${paramValue} := ""\n`;
        text += emitParamGroupCheck(pp);
        text += `\t${paramValue} = ${helpers.formatParamValue(pp, imports)}\n\t}\n`;

        // there are two cases for optional path params.
        //  - /foo/bar/{optional}
        //  - /foo/bar{/optional}
        // for the second case, we need to include a forward slash
        if (method.httpPath[method.httpPath.indexOf(`{${pp.pathSegment}}`) - 1] !== '/') {
          optionalPathSep = true;
        }
      }

      const emitPathEscape = function(): string {
        if (pp.isEncoded) {
          imports.add('net/url');
          return `url.PathEscape(${paramValue})`;
        }
        return paramValue;
      };

      if (optionalPathSep) {
        text += `\tif len(${paramValue}) > 0 {\n`;
        text += `\t\t${paramValue} = "/"+${emitPathEscape()}\n`;
        text += '\t}\n';
      } else {
        paramValue = emitPathEscape();
      }

      text += `\turlPath = strings.ReplaceAll(urlPath, "{${pp.pathSegment}}", ${paramValue})\n`;
    }
  }

  text += `\treq, err := runtime.NewRequest(ctx, http.Method${naming.capitalize(method.httpMethod)}, ${hostParam})\n`;
  text += '\tif err != nil {\n';
  text += '\t\treturn nil, err\n';
  text += '\t}\n';

  // add query parameters
  const encodedParams = methodParamGroups.encodedQueryParams;
  const unencodedParams = methodParamGroups.unencodedQueryParams;

  const emitQueryParam = function (qp: go.QueryParameter, setter: string): string {
    let qpText = '';
    if (qp.location === 'method' && go.isClientSideDefault(qp.style)) {
      qpText = emitClientSideDefault(qp, qp.style, (name, val) => { return `\treqQP.Set(${name}, ${val})`; }, imports);
    } else if (go.isRequiredParameter(qp.style) || go.isLiteralParameter(qp.style) || (qp.location === 'client' && go.isClientSideDefault(qp.style))) {
      qpText = `\t${setter}\n`;
    } else if (qp.location === 'client' && !qp.group) {
      // global optional param
      qpText = `\tif client.${qp.name} != nil {\n`;
      qpText += `\t\t${setter}\n`;
      qpText += '\t}\n';
    } else {
      qpText = emitParamGroupCheck(qp);
      qpText += `\t\t${setter}\n`;
      qpText += '\t}\n';
    }
    return qpText;
  };

  // emit encoded params first
  if (encodedParams.length > 0) {
    text += '\treqQP := req.Raw().URL.Query()\n';
    for (const qp of encodedParams.sort((a: go.QueryParameter, b: go.QueryParameter) => { return helpers.sortAscending(a.queryParameter, b.queryParameter); })) {
      let setter: string;
      if (qp.kind === 'queryCollectionParam' && qp.collectionFormat === 'multi') {
        setter = `\tfor _, qv := range ${helpers.getParamName(qp)} {\n`;

        // emit a type conversion for the qv based on the array's element type
        let queryVal: string;
        const arrayQP = qp.type;
        switch (arrayQP.elementType.kind) {
          case 'constant':
            switch (arrayQP.elementType.type) {
              case 'string':
                queryVal = 'string(qv)';
                break;
              default:
                imports.add('fmt');
                queryVal = 'fmt.Sprintf("%d", qv)';
            }
            break;
          case 'string':
            queryVal = 'qv';
            break;
          default:
            imports.add('fmt');
            queryVal = 'fmt.Sprintf("%v", qv)';
        }

        setter += `\t\treqQP.Add("${qp.queryParameter}", ${queryVal})\n`;
        setter += '\t}';
      } else {
        // cannot initialize setter to this value as helpers.formatParamValue() can change imports
        setter = `reqQP.Set("${qp.queryParameter}", ${helpers.formatParamValue(qp, imports)})`;
      }
      text += emitQueryParam(qp, setter);
    }
    text += '\treq.Raw().URL.RawQuery = reqQP.Encode()\n';
  }

  // tack on any unencoded params to the end
  if (unencodedParams.length > 0) {
    if (encodedParams.length > 0) {
      text += '\tunencodedParams := []string{req.Raw().URL.RawQuery}\n';
    } else {
      text += '\tunencodedParams := []string{}\n';
    }
    for (const qp of unencodedParams.sort((a: go.QueryParameter, b: go.QueryParameter) => { return helpers.sortAscending(a.queryParameter, b.queryParameter); })) {
      let setter: string;
      if (qp.kind === 'queryCollectionParam' && qp.collectionFormat === 'multi') {
        setter = `\tfor _, qv := range ${helpers.getParamName(qp)} {\n`;
        setter += `\t\tunencodedParams = append(unencodedParams, "${qp.queryParameter}="+qv)\n`;
        setter += '\t}';
      } else {
        setter = `unencodedParams = append(unencodedParams, "${qp.queryParameter}="+${helpers.formatParamValue(qp, imports)})`;
      }
      text += emitQueryParam(qp, setter);
    }
    imports.add('strings');
    text += '\treq.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")\n';
  }

  if (method.kind !== 'nextPageMethod' && method.returns.result?.kind === 'binaryResult') {
    // skip auto-body downloading for binary stream responses
    text += '\truntime.SkipBodyDownload(req)\n';
  }

  // add specific request headers
  const emitHeaderSet = function (headerParam: go.HeaderParameter, prefix: string): string {
    if (headerParam.kind === 'headerMapParam') {
      let headerText = `${prefix}for k, v := range ${helpers.getParamName(headerParam)} {\n`;
      headerText += `${prefix}\tif v != nil {\n`;
      headerText += `${prefix}\t\treq.Raw().Header["${headerParam.headerName}"+k] = []string{*v}\n`;
      headerText += `${prefix}}\n`;
      headerText += `${prefix}}\n`;
      return headerText;
    } else if (headerParam.location === 'method' && go.isClientSideDefault(headerParam.style)) {
      return emitClientSideDefault(headerParam, headerParam.style, (name, val) => {
        return `${prefix}req.Raw().Header[${name}] = []string{${val}}`;
      }, imports);
    } else {
      return `${prefix}req.Raw().Header["${headerParam.headerName}"] = []string{${helpers.formatParamValue(headerParam, imports)}}\n`;
    }
  };

  let contentType: string | undefined;
  for (const param of methodParamGroups.headerParams.sort((a: go.HeaderParameter, b: go.HeaderParameter) => { return helpers.sortAscending(a.headerName, b.headerName);})) {
    if (param.headerName.match(/^content-type$/)) {
      // canonicalize content-type as req.SetBody checks for it via its canonicalized name :(
      param.headerName = 'Content-Type';
    }

    if (param.headerName === 'Content-Type' && param.style === 'literal') {
      // the content-type header will be set as part of emitSetBodyWithErrCheck
      // to handle cases where the body param is optional. we don't want to set
      // the content-type if the body is nil.
      // we do it like this as tsp specifies content-type while swagger does not.
      contentType = helpers.formatParamValue(param, imports);
    } else if (go.isRequiredParameter(param.style) || go.isLiteralParameter(param.style) || go.isClientSideDefault(param.style)) {
      text += emitHeaderSet(param, '\t');
    } else if (param.location === 'client' && !param.group) {
      // global optional param
      text += `\tif client.${param.name} != nil {\n`;
      text += emitHeaderSet(param, '\t');
      text += '\t}\n';
    } else {
      text += emitParamGroupCheck(param);
      text += emitHeaderSet(param, '\t\t');
      text += '\t}\n';
    }
  }

  // note that these are mutually exclusive
  const bodyParam = methodParamGroups.bodyParam;
  const formBodyParams = methodParamGroups.formBodyParams;
  const multipartBodyParams = methodParamGroups.multipartBodyParams;
  const partialBodyParams = methodParamGroups.partialBodyParams;

  const emitSetBodyWithErrCheck = function(setBodyParam: string, contentType?: string): string {
    let content = `if err := ${setBodyParam}; err != nil {\n\treturn nil, err\n}\n;`;
    if (contentType) {
      content = `req.Raw().Header["Content-Type"] = []string{${contentType}}\n` + content;
    }
    return content;
  };

  if (bodyParam) {
    if (bodyParam.bodyFormat === 'JSON' || bodyParam.bodyFormat === 'XML') {
      // default to the body param name
      let body = helpers.getParamName(bodyParam);
      if (bodyParam.type.kind === 'literal') {
        // if the value is constant, embed it directly
        body = helpers.formatLiteralValue(bodyParam.type, true);
      } else if (bodyParam.bodyFormat === 'XML' && bodyParam.type.kind === 'slice') {
        // for XML payloads, create a wrapper type if the payload is an array
        imports.add('encoding/xml');
        text += '\ttype wrapper struct {\n';
        let tagName = go.getTypeDeclaration(bodyParam.type, method.receiver.type.pkg);
        if (bodyParam.xml?.name) {
          tagName = bodyParam.xml.name;
        }
        text += `\t\tXMLName xml.Name \`xml:"${tagName}"\`\n`;
        const fieldName = naming.capitalize(bodyParam.name);
        let tag = go.getTypeDeclaration(bodyParam.type.elementType, method.receiver.type.pkg);
        if (bodyParam.type.elementType.kind === 'model' && bodyParam.type.elementType.xml?.name) {
          tag = bodyParam.type.elementType.xml.name;
        }
        text += `\t\t${fieldName} *${go.getTypeDeclaration(bodyParam.type, method.receiver.type.pkg)} \`xml:"${tag}"\`\n`;
        text += '\t}\n';
        let addr = '&';
        if (!go.isRequiredParameter(bodyParam.style) && !bodyParam.byValue) {
          addr = '';
        }
        body = `wrapper{${fieldName}: ${addr}${body}}`;
      } else if (bodyParam.type.kind === 'time' && bodyParam.type.format !== 'RFC3339') {
        // wrap the body in the internal time type
        // no need for RFC3339 as the JSON marshaler defaults to that.
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/datetime');
        body = `datetime.${bodyParam.type.format}(${body})`;
      } else if (isArrayOfDateTimeForMarshalling(bodyParam.type)) {
        const timeInfo = isArrayOfDateTimeForMarshalling(bodyParam.type);
        let elementPtr = '*';
        if (timeInfo?.elemByVal) {
          elementPtr = '';
        }
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/datetime');
        text += `\taux := make([]${elementPtr}datetime.${timeInfo?.format}, len(${body}))\n`;
        text += `\tfor i := 0; i < len(${body}); i++ {\n`;
        text += `\t\taux[i] = (${elementPtr}datetime.${timeInfo?.format})(${body}[i])\n`;
        text += '\t}\n';
        body = 'aux';
      } else if (isMapOfDateTime(bodyParam.type)) {
        let timeType = isMapOfDateTime(bodyParam.type);
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/datetime');
        text += `\taux := map[string]*datetime.${timeType}{}\n`;
        text += `\tfor k, v := range ${body} {\n`;
        text += `\t\taux[k] = (*datetime.${timeType})(v)\n`;
        text += '\t}\n';
        body = 'aux';
      }
      let setBody = `runtime.MarshalAs${getMediaFormat(bodyParam.type, bodyParam.bodyFormat, `req, ${body}`)}`;
      if (bodyParam.type.kind === 'rawJSON') {
        imports.add('bytes');
        imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
        setBody = `req.SetBody(streaming.NopCloser(bytes.NewReader(${body})), "application/${bodyParam.bodyFormat.toLowerCase()}")`;
      }
      if (go.isRequiredParameter(bodyParam.style) || go.isLiteralParameter(bodyParam.style)) {
        text += `\t${emitSetBodyWithErrCheck(setBody, contentType)}`;
        text += '\treturn req, nil\n';
      } else {
        text += emitParamGroupCheck(bodyParam);
        text += `\t${emitSetBodyWithErrCheck(setBody, contentType)}`;
        text += '\t\treturn req, nil\n';
        text += '\t}\n';
        text += '\treturn req, nil\n';
      }
    } else if (bodyParam.bodyFormat === 'binary') {
      if (go.isRequiredParameter(bodyParam.style)) {
        text += `\t${emitSetBodyWithErrCheck(`req.SetBody(${bodyParam.name}, ${bodyParam.contentType})`, contentType)}`;
        text += '\treturn req, nil\n';
      } else {
        text += emitParamGroupCheck(bodyParam);
        text += `\t${emitSetBodyWithErrCheck(`req.SetBody(${helpers.getParamName(bodyParam)}, ${bodyParam.contentType})`, contentType)}`;
        text += '\treturn req, nil\n';
        text += '\t}\n';
        text += '\treturn req, nil\n';
      }
    } else if (bodyParam.bodyFormat === 'Text') {
      imports.add('strings');
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
      if (go.isRequiredParameter(bodyParam.style)) {
        text += `\tbody := streaming.NopCloser(strings.NewReader(${bodyParam.name}))\n`;
        text += `\t${emitSetBodyWithErrCheck(`req.SetBody(body, ${bodyParam.contentType})`, contentType)}`;
        text += '\treturn req, nil\n';
      } else {
        text += emitParamGroupCheck(bodyParam);
        text += `\tbody := streaming.NopCloser(strings.NewReader(${helpers.getParamName(bodyParam)}))\n`;
        text += `\t${emitSetBodyWithErrCheck(`req.SetBody(body, ${bodyParam.contentType})`, contentType)}`;
        text += '\treturn req, nil\n';
        text += '\t}\n';
        text += '\treturn req, nil\n';
      }
    }
  } else if (partialBodyParams.length > 0) {
    // partial body params are discrete params that are all fields within an internal struct.
    // define and instantiate an instance of the wire type, using the values from each param.
    text += '\tbody := struct {\n';
    for (const partialBodyParam of partialBodyParams) {
      text += `\t\t${naming.capitalize(partialBodyParam.serializedName)} ${helpers.star(partialBodyParam.byValue)}${go.getTypeDeclaration(partialBodyParam.type, method.receiver.type.pkg)} \`${partialBodyParam.format.toLowerCase()}:"${partialBodyParam.serializedName}"\`\n`;
    }
    text += '\t}{\n';
    // required params are emitted as initializers in the struct literal
    for (const partialBodyParam of partialBodyParams) {
      if (go.isRequiredParameter(partialBodyParam.style)) {
        text += `\t\t${naming.capitalize(partialBodyParam.serializedName)}: ${naming.uncapitalize(partialBodyParam.name)},\n`;
      }
    }
    text += '\t}\n';
    // now populate any optional params from the options type
    for (const partialBodyParam of partialBodyParams) {
      if (!go.isRequiredParameter(partialBodyParam.style)) {
        text += emitParamGroupCheck(partialBodyParam);
        text += `\t\tbody.${naming.capitalize(partialBodyParam.serializedName)} = options.${naming.capitalize(partialBodyParam.name)}\n\t}\n`;
      }
    }
    // TODO: spread params are JSON only https://github.com/Azure/autorest.go/issues/1455
    text += '\treq.Raw().Header["Content-Type"] = []string{"application/json"}\n';
    text += '\tif err := runtime.MarshalAsJSON(req, body); err != nil {\n\t\treturn nil, err\n\t}\n';
    text += '\treturn req, nil\n';
  } else if (multipartBodyParams.length > 0) {
    if (multipartBodyParams.length === 1 && multipartBodyParams[0].type.kind === 'model' && multipartBodyParams[0].type.annotations.multipartFormData) {
      text += `\tformData, err := ${multipartBodyParams[0].name}.toMultipartFormData()\n`;
      text += '\tif err != nil {\n\t\treturn nil, err\n\t}\n';
    } else {
      text += '\tformData := map[string]any{}\n';
      for (const param of multipartBodyParams) {
        const setter = `formData["${param.name}"] = ${helpers.getParamName(param)}`;
        if (go.isRequiredParameter(param.style)) {
          text += `\t${setter}\n`;
        } else {
          text += emitParamGroupCheck(param);
          text += `\t${setter}\n\t}\n`;
        }
      }
    }
    text += '\tif err := runtime.SetMultipartFormData(req, formData); err != nil {\n\t\treturn nil, err\n\t}\n';
    text += '\treturn req, nil\n';
  } else if (formBodyParams.length > 0) {
    const emitFormData = function (param: go.FormBodyParameter, setter: string): string {
      let formDataText = '';
      if (go.isRequiredParameter(param.style)) {
        formDataText = `\t${setter}\n`;
      } else {
        formDataText = emitParamGroupCheck(param);
        formDataText += `\t\t${setter}\n`;
        formDataText += '\t}\n';
      }
      return formDataText;
    };
    imports.add('net/url');
    imports.add('strings');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
    text += '\tformData := url.Values{}\n';
    // find all the form body params
    for (const param of formBodyParams) {
      const setter = `formData.Set("${param.formDataName}", ${helpers.formatParamValue(param, imports)})`;
      text += emitFormData(param, setter);
    }
    text += '\tbody := streaming.NopCloser(strings.NewReader(formData.Encode()))\n';
    text += `\t${emitSetBodyWithErrCheck('req.SetBody(body, "application/x-www-form-urlencoded")')}`;
    text += '\treturn req, nil\n';
  } else {
    text += '\treturn req, nil\n';
  }
  text += '}\n\n';
  return text;
}

function emitClientSideDefault(param: go.HeaderCollectionParameter | go.HeaderScalarParameter | go.QueryParameter, csd: go.ClientSideDefault, setterFormat: (name: string, val: string) => string, imports: ImportManager): string {
  const defaultVar = naming.uncapitalize(param.name) + 'Default';
  let text = `\t${defaultVar} := ${helpers.formatLiteralValue(csd.defaultValue, true)}\n`;
  text += `\tif options != nil && options.${naming.capitalize(param.name)} != nil {\n`;
  text += `\t\t${defaultVar} = *options.${naming.capitalize(param.name)}\n`;
  text += '}\n';

  let serializedName: string;
  switch (param.kind) {
    case 'headerCollectionParam':
    case 'headerScalarParam':
      serializedName = param.headerName;
      break;
    case 'queryCollectionParam':
    case 'queryScalarParam':
      serializedName = param.queryParameter;
      break;
  }

  text += setterFormat(`"${serializedName}"`, helpers.formatValue(defaultVar, param.type, imports)) + '\n';
  return text;
}

function getMediaFormat(type: go.WireType, mediaType: 'JSON' | 'XML', param: string): string {
  let marshaller: 'JSON' | 'XML' | 'ByteArray' = mediaType;
  let format = '';
  if (type.kind === 'encodedBytes') {
    marshaller = 'ByteArray';
    format = `, runtime.Base64${type.encoding}Format`;
  }
  return `${marshaller}(${param}${format})`;
}

function isArrayOfDateTimeForMarshalling(paramType: go.WireType): { format: go.TimeFormat, elemByVal: boolean } | undefined {
  if (paramType.kind !== 'slice') {
    return undefined;
  }
  if (paramType.elementType.kind !== 'time') {
    return undefined;
  }
  switch (paramType.elementType.format) {
    case 'PlainDate':
    case 'RFC1123':
    case 'PlainTime':
    case 'Unix':
      return {
        format: paramType.elementType.format,
        elemByVal: paramType.elementTypeByValue
      };
    default:
      // RFC3339 uses the default marshaller
      return undefined;
  }
}

// returns true if the method requires a response handler.
// this is used to unmarshal the response body, parse response headers, or both.
function needsResponseHandler(method: go.MethodType): boolean {
  return helpers.hasSchemaResponse(method) || method.returns.headers.length > 0;
}

function generateResponseUnmarshaller(method: go.MethodType, type: go.WireType, format: go.ResultFormat, unmarshalTarget: string, imports: ImportManager): string {
  let unmarshallerText = '';
  const zeroValue = getZeroReturnValue(method, 'handler');
  if (type.kind === 'time') {
    // use the designated time type for unmarshalling
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/datetime');
    unmarshallerText += `\tvar aux *datetime.${type.format}\n`;
    unmarshallerText += `\tif err := runtime.UnmarshalAs${format}(resp, &aux); err != nil {\n`;
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += `\tresult.${helpers.getResultFieldName(method)} = (*time.Time)(aux)\n`;
    return unmarshallerText;
  } else if (isArrayOfDateTime(type)) {
    // unmarshalling arrays of date/time is a little more involved
    const timeInfo = isArrayOfDateTime(type);
    let elementPtr = '*';
    if (timeInfo?.elemByVal) {
      elementPtr = '';
    }
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/datetime');
    unmarshallerText += `\tvar aux []${elementPtr}datetime.${timeInfo?.format}\n`;
    unmarshallerText += `\tif err := runtime.UnmarshalAs${format}(resp, &aux); err != nil {\n`;
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += `\tcp := make([]${elementPtr}time.Time, len(aux))\n`;
    unmarshallerText += '\tfor i := 0; i < len(aux); i++ {\n';
    unmarshallerText += `\t\tcp[i] = (${elementPtr}time.Time)(aux[i])\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += `\tresult.${helpers.getResultFieldName(method)} = cp\n`;
    return unmarshallerText;
  } else if (isMapOfDateTime(type)) {
    let timeType = isMapOfDateTime(type);
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/datetime');
    unmarshallerText += `\taux := map[string]*datetime.${timeType}{}\n`;
    unmarshallerText += `\tif err := runtime.UnmarshalAs${format}(resp, &aux); err != nil {\n`;
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += '\tcp := map[string]*time.Time{}\n';
    unmarshallerText += '\tfor k, v := range aux {\n';
    unmarshallerText += '\t\tcp[k] = (*time.Time)(v)\n';
    unmarshallerText += '\t}\n';
    unmarshallerText += `\tresult.${helpers.getResultFieldName(method)} = cp\n`;
    return unmarshallerText;
  }
  if (format === 'JSON' || format === 'XML') {
    if (type.kind === 'rawJSON') {
      unmarshallerText += '\tbody, err := runtime.Payload(resp)\n';
      unmarshallerText += '\tif err != nil {\n';
      unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
      unmarshallerText += '\t}\n';
      unmarshallerText += `\t${unmarshalTarget} = body\n`;
    } else {
      unmarshallerText += `\tif err := runtime.UnmarshalAs${getMediaFormat(type, format, `resp, &${unmarshalTarget}`)}; err != nil {\n`;
      unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
      unmarshallerText += '\t}\n';
    }
  } else if (format === 'Text') {
    unmarshallerText += '\tbody, err := runtime.Payload(resp)\n';
    unmarshallerText += '\tif err != nil {\n';
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += '\ttxt := string(body)\n';
    unmarshallerText += `\t${unmarshalTarget} = &txt\n`;
  } else {
    // the remaining formats should have been handled elsewhere
    throw new CodegenError('InternalError', `unhandled format ${format} for operation ${method.receiver.type.name}.${method.name}`);
  }
  return unmarshallerText;
}

function createProtocolResponse(method: go.SyncMethod | go.LROPageableMethod | go.PageableMethod, imports: ImportManager): string {
  if (!needsResponseHandler(method)) {
    return '';
  }
  const name = method.naming.responseMethod;
  let text = `${helpers.comment(name, '// ')} handles the ${method.name} response.\n`;
  text += `func ${getClientReceiverDefinition(method.receiver)} ${name}(resp *http.Response) (${generateReturnsInfo(method, 'handler').join(', ')}) {\n`;

  const addHeaders = function (headers: Array<go.HeaderScalarResponse | go.HeaderMapResponse>) {
    for (const header of headers) {
      text += formatHeaderResponseValue(header, imports, 'result', `${method.returns.name}{}`);
    }
  };

  const result = method.returns.result;
  if (!result) {
    // only headers
    text += `\tresult := ${method.returns.name}{}\n`;
    addHeaders(method.returns.headers);
  } else {
    switch (result.kind) {
      case 'anyResult':
        imports.add('fmt');
        text += `\tresult := ${method.returns.name}{}\n`;
        addHeaders(method.returns.headers);
        text += '\tswitch resp.StatusCode {\n';
        for (const statusCode of method.httpStatusCodes) {
          text += `\tcase ${helpers.formatStatusCodes([statusCode])}:\n`;
          const resultType = result.httpStatusCodeType[statusCode];
          if (!resultType) {
            // the operation contains a mix of schemas and non-schema responses
            continue;
          }
          text += `\tvar val ${go.getTypeDeclaration(resultType, method.receiver.type.pkg)}\n`;
          text += generateResponseUnmarshaller(method, resultType, result.format, 'val', imports);
          text += '\tresult.Value = val\n';
        }
        text += '\tdefault:\n';
        text += `\t\treturn ${getZeroReturnValue(method, 'handler')}, fmt.Errorf("unhandled HTTP status code %d", resp.StatusCode)\n`;
        text += '\t}\n';
        break;
      case 'binaryResult':
        text += `\tresult := ${method.returns.name}{${result.fieldName}: resp.Body}\n`;
        addHeaders(method.returns.headers);
        break;
      case 'headAsBooleanResult':
        text += `\tresult := ${method.returns.name}{${result.fieldName}: resp.StatusCode >= 200 && resp.StatusCode < 300}\n`;
        addHeaders(method.returns.headers);
        break;
      case 'modelResult':
        text += `\tresult := ${method.returns.name}{}\n`;
        addHeaders(method.returns.headers);
        text += generateResponseUnmarshaller(method, result.modelType, result.format, `result.${helpers.getResultFieldName(method)}`, imports);
        break;
      case 'monomorphicResult':
        text += `\tresult := ${method.returns.name}{}\n`;
        addHeaders(method.returns.headers);
        let target = `result.${helpers.getResultFieldName(method)}`;
        // when unmarshalling a wrapped XML array, unmarshal into the response envelope
        if (result.format === 'XML' && result.monomorphicType.kind === 'slice') {
          target = 'result';
        }
        text += generateResponseUnmarshaller(method, result.monomorphicType, result.format, target, imports);
        break;
      case 'polymorphicResult':
        text += `\tresult := ${method.returns.name}{}\n`;
        addHeaders(method.returns.headers);
        text += generateResponseUnmarshaller(method, result.interface, result.format, 'result', imports);
        break;
      default:
        result satisfies never;
    }
  }

  text += '\treturn result, nil\n';
  text += '}\n\n';
  return text;
}

function isArrayOfDateTime(paramType: go.WireType): { format: go.TimeFormat, elemByVal: boolean } | undefined {
  if (paramType.kind !== 'slice') {
    return undefined;
  }
  if (paramType.elementType.kind !== 'time') {
    return undefined;
  }
  return {
    format: paramType.elementType.format,
    elemByVal: paramType.elementTypeByValue
  };
}

function isMapOfDateTime(paramType: go.WireType): go.TimeFormat | undefined {
  if (paramType.kind !== 'map') {
    return undefined;
  }
  if (paramType.valueType.kind !== 'time') {
    return undefined;
  }
  return paramType.valueType.format;
}

/**
 * returns the parameters for the public API
 * e.g. "ctx context.Context, i int, s string"
 * 
 * @param method the method for which to emit the parameters
 * @param imports the import manager currently in scope
 * @returns the text for the method parameters
 */
function getAPIParametersSig(method: go.ClientAccessor | go.MethodType, imports: ImportManager): string {
  const params = new Array<string>();
  if (method.kind === 'clientAccessor') {
    // client accessor params don't have a concept
    // of optionality nor do they contain literals
    for (const param of method.parameters) {
      imports.addForType(param.type);
      params.push(`${param.name} ${go.getTypeDeclaration(param.type, method.receiver.type.pkg)}`);
    }
  } else {
    const methodParams = helpers.getMethodParameters(method);
    if (method.kind !== 'pageableMethod') {
      imports.add('context');
      params.push('ctx context.Context');
    }
    for (const methodParam of methodParams) {
      if (methodParam.kind !== 'paramGroup') {
        imports.addForType(methodParam.type);
      }
      params.push(`${naming.uncapitalize(methodParam.name)} ${helpers.formatParameterTypeName(method.receiver.type.pkg, methodParam)}`);
    }
  }
  return params.join(', ');
}

// returns the return signature where each entry is the type name
// e.g. [ '*string', 'error' ]
// apiType describes where the return sig is used.
//   api - for the API definition
//    op - for the operation
// handler - for the response handler
function generateReturnsInfo(method: go.MethodType, apiType: 'api' | 'op' | 'handler'): Array<string> {
  let returnType = method.returns.name;
  switch (method.kind) {
    case 'lroMethod':
    case 'lroPageableMethod':
      switch (apiType) {
        case 'api':
          if (method.kind === 'lroPageableMethod') {
            returnType = `*runtime.Poller[*runtime.Pager[${returnType}]]`;
          } else {
            returnType = `*runtime.Poller[${returnType}]`;
          }
          break;
        case 'handler':
          // we only have a handler for operations that return a schema
          if (method.kind !== 'lroPageableMethod') {
            throw new CodegenError('InternalError', `handler being generated for non-pageable LRO ${method.name} which is unexpected`);
          }
          break;
        case 'op':
          returnType = '*http.Response';
          break;
      }
      break;
    case 'pageableMethod':
      switch (apiType) {
        case 'api':
        case 'op':
          // pager operations don't return an error
          return [`*runtime.Pager[${returnType}]`];
      }
      break;
  }
  return [returnType, 'error'];
}

function generateLROBeginMethod(method: go.LROMethod | go.LROPageableMethod, imports: ImportManager, options: go.Options): string {
  const params = getAPIParametersSig(method, imports);
  const returns = generateReturnsInfo(method, 'api');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
  let text = '';
  if (method.docs.summary || method.docs.description) {
    text += helpers.formatDocCommentWithPrefix(fixUpMethodName(method), method.docs);
    text += genRespErrorDoc(method);
    text += genApiVersionDoc(method.apiVersions);
  }
  const zeroResp = getZeroReturnValue(method, 'api');
  const methodParams = helpers.getMethodParameters(method);
  for (const param of methodParams) {
    text += helpers.formatCommentAsBulletItem(param.name, param.docs);
  }
  text += `func ${getClientReceiverDefinition(method.receiver)} ${fixUpMethodName(method)}(${params}) (${returns.join(', ')}) {\n`;
  let pollerType = 'nil';
  let pollerTypeParam = `[${method.returns.name}]`;
  if (method.kind === 'lroPageableMethod') {
    // for paged LROs, we construct a pager and pass it to the LRO ctor.
    pollerTypeParam = `[*runtime.Pager${pollerTypeParam}]`;
    pollerType = '&pager';
    text += '\tpager := ';
    text += emitPagerDefinition(method, imports, options);
  }

  text += '\tif options == nil || options.ResumeToken == "" {\n';

  // creating the poller from response branch

  const opName = method.naming.internalMethod;
  text += `\t\tresp, err := client.${opName}(${helpers.getCreateRequestParameters(method)})\n`;
  text += '\t\tif err != nil {\n';
  text += `\t\t\treturn ${zeroResp}, err\n`;
  text += '\t\t}\n';

  let finalStateVia = '';
  // LRO operation might have a special configuration set in x-ms-long-running-operation-options
  // which indicates a specific url to perform the final Get operation on
  if (method.finalStateVia) {
    switch (method.finalStateVia) {
      case 'azure-async-operation':
        finalStateVia = 'runtime.FinalStateViaAzureAsyncOp';
        break;
      case 'location':
        finalStateVia = 'runtime.FinalStateViaLocation';
        break;
      case 'original-uri':
        finalStateVia = 'runtime.FinalStateViaOriginalURI';
        break;
      case 'operation-location':
        finalStateVia = 'runtime.FinalStateViaOpLocation';
        break;
      default:
        throw new CodegenError('InternalError', `unhandled final-state-via value ${finalStateVia}`);
    }
  }

  text += '\t\tpoller, err := runtime.NewPoller';
  if (finalStateVia === '' && pollerType === 'nil' && !options.injectSpans) {
    // the generic type param is redundant when it's also specified in the
    // options struct so we only include it when there's no options.
    text += pollerTypeParam;
  }
  text += '(resp, client.internal.Pipeline(), ';
  if (finalStateVia === '' && pollerType === 'nil' && !options.injectSpans && !method.operationLocationResultPath) {
    // no options
    text += 'nil)\n';
  } else {
    // at least one option
    text += `&runtime.NewPollerOptions${pollerTypeParam}{\n`;
    if (finalStateVia !== '') {
      text += `\t\t\tFinalStateVia: ${finalStateVia},\n`;  
    }
    if (method.operationLocationResultPath) {
      text += `\t\t\tOperationLocationResultPath: "${method.operationLocationResultPath}",\n`;
    }
    if (pollerType !== 'nil') {
      text += `\t\t\tResponse: ${pollerType},\n`;
    }
    if (options.injectSpans) {
      text += '\t\t\tTracer: client.internal.Tracer(),\n';
    }
    text += '\t\t})\n';
  }
  text += '\t\treturn poller, err\n';
  text += '\t} else {\n';

  // creating the poller from resume token branch

  text += '\t\treturn runtime.NewPollerFromResumeToken';
  if (pollerType === 'nil' && !options.injectSpans) {
    text += pollerTypeParam;
  }
  text += '(options.ResumeToken, client.internal.Pipeline(), ';
  if (pollerType === 'nil' && !options.injectSpans) {
    text += 'nil)\n';
  } else {
    text += `&runtime.NewPollerFromResumeTokenOptions${pollerTypeParam}{\n`;
    if (pollerType !== 'nil') {
      text += `\t\t\tResponse: ${pollerType},\n`;
    }
    if (options.injectSpans) {
      text += '\t\t\tTracer: client.internal.Tracer(),\n';
    }
    text  += '\t\t})\n';
  }
  text += '\t}\n';

  text += '}\n\n';
  return text;
}

export function fixUpMethodName(method: go.MethodType): string {
  switch (method.kind) {
    case 'lroMethod':
    case 'lroPageableMethod':
      return `Begin${method.name}`;
    case 'pageableMethod': {
      let N = 'N';
      let name = method.name;
      if (method.name[0] !== method.name[0].toUpperCase()) {
        // the method isn't exported; don't export the pager ctor
        N = 'n';
        // ensure correct casing of the emitted function name e.g.,
        // "listThings" -> "newListThingsPager"
        name = name[0].toUpperCase() + name.substring(1);
      }
      return `${N}ew${name}Pager`;
    }
    case 'method':
      return method.name;
  }
}
