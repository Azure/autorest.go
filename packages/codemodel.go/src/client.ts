/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import { CodeModelError } from './errors.js';
import { MethodExample } from './examples.js';
import * as pkg from './package.js';
import * as param from './param.js';
import * as result from './result.js';
import * as type from './type.js';

/** an SDK client */
export interface Client {
  /** the name of the client */
  name: string;

  /** any docs for the client */
  docs: type.Docs;

  /** the client options type. for ARM, this will be a QualifiedType (arm.ClientOptions) */
  options: ClientOptions;

  /** constructor params that are persisted as fields on the client, can be empty */
  parameters: Array<param.Parameter>;

  /** all the constructors for this client, can be empty */
  constructors: Array<Constructor>;

  /** contains client methods. can be empty */
  methods: Array<MethodType>;

  /** contains any client accessor methods. can be empty */
  clientAccessors: Array<ClientAccessor>;

  /** client has a statically defined or templated host */
  host?: string;

  /**
   * templatedHost indicates that there's one or more URIParameters
   * required to construct the complete host. the parameters can
   * be solely on the client or span client and method params.
   */
  templatedHost: boolean;

  /** the parent client in a hierarchical client */
  parent?: Client;
}

/** the possible types used for the client options type */
export type ClientOptions = param.ParameterGroup | param.Parameter;

/** a client method that returns a sub-client instance */
export interface ClientAccessor {
  /** the name of the client accessor method */
  name: string;

  /** the client returned by the accessor method */
  subClient: Client;
}

/** a client constructor function */
export interface Constructor {
  /** the name of the constructor function */
  name: string;

  /** the modeled parameters. can be empty */
  parameters: Array<param.Parameter>;
}

/** the possible values defining the "final state via" behavior for LROs */
export type FinalStateVia = 'azure-async-operation' | 'location' | 'operation-location' | 'original-uri';

/** the supported HTTP verbs */
export type HTTPMethod = 'delete' | 'get' | 'head' | 'patch' | 'post' | 'put';

/** the possible method types */
export type MethodType = LROMethod | LROPageableMethod | Method | PageableMethod;

/** a long-running operation method */
export interface LROMethod extends LROMethodBase {
  kind: 'lroMethod';
}

/** a long-running operation method that returns pages of responses */
export interface LROPageableMethod extends LROMethodBase, PageableMethodBase {
  kind: 'lroPageableMethod';
}

/** a synchronous method */
export interface Method extends MethodBase {
  kind: 'method';
}

/** contains the names of the helper methods used to create a complete method implementation */
export interface MethodNaming {
  /** the name of the internal method for consumption by LROs/paging methods */
  internalMethod: string;

  /** the name of the internal method that creates the HTTP request */
  requestMethod: string;

  /** the name of the internal method that handles the HTTP response */
  responseMethod: string;
}

/**
 * the internal method used for fetching the next page for a PageableMethod.
 * It's unique from a regular Method as it's not exported and has no optional params/response envelope.
 * thus, it's not included in the array of methods for a client.
 */
export interface NextPageMethod {
  kind: 'nextPageMethod';

  /** the name of the next page method */
  name: string;
 
  /** the HTTP path used when creating the request */
  httpPath: string;

  /** the HTTP verb used when creating the request */
  httpMethod: HTTPMethod;

  /** any modeled parameters */
  parameters: Array<param.Parameter>;

  /** the complete list of successful HTTP status codes */
  httpStatusCodes: Array<number>;

  /** the client to which the method belongs */
  client: Client;

  apiVersions: Array<string>;
}

/** a synchronous method that returns pages of responses */
export interface PageableMethod extends PageableMethodBase {
  kind: 'pageableMethod';
}

/** narrows method to a LRO method type within the conditional block */
export function isLROMethod(method: MethodType): method is LROMethod | LROPageableMethod {
  return method.kind === 'lroMethod' || method.kind === 'lroPageableMethod';
}

/** narrows method to a pageable method type within the conditional block */
export function isPageableMethod(method: MethodType): method is LROPageableMethod | PageableMethod {
  return method.kind === 'lroPageableMethod' || method.kind === 'pageableMethod';
}

/** creates the ClientOptions type from the specified input */
export function newClientOptions(modelType: pkg.CodeModelType, clientName: string): ClientOptions {
  let options: ClientOptions;
  if (modelType === 'azure-arm') {
    options = new param.Parameter('options', new type.QualifiedType('ClientOptions', 'github.com/Azure/azure-sdk-for-go/sdk/azcore/arm'), 'optional', false, 'client');
    options.docs.summary = 'pass nil to accept the default values.';
  } else {
    const optionsTypeName = `${clientName}Options`;
    options = new param.ParameterGroup('options', optionsTypeName, false, 'client');
    options.docs.summary = `${optionsTypeName} contains the optional values for creating a [${clientName}]`;
  }
  return options;
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// base types
///////////////////////////////////////////////////////////////////////////////////////////////////

interface LROMethodBase extends MethodBase {
  finalStateVia?: FinalStateVia;

  operationLocationResultPath?: string;
}

interface MethodBase {
  /** the name of the method */
  name: string;

  /** any docs for the method */
  docs: type.Docs;

  /** the HTTP path used when creating the request */
  httpPath: string;

  /** the HTTP verb used when creating the request */
  httpMethod: HTTPMethod;

  /** any modeled parameters. the ones we add to the generated code (context.Context etc) aren't included here */
  parameters: Array<param.Parameter>;

  /** the method options type for this methoid */
  optionalParamsGroup: param.ParameterGroup;

  /** the response type for this method */
  responseEnvelope: result.ResponseEnvelope;

  /** the complete list of successful HTTP status codes */
  httpStatusCodes: Array<number>;

  /** the client to which the method belongs */
  client: Client;

  /** naming info for the internal hepler methods for which this method depends */
  naming: MethodNaming;

  apiVersions: Array<string>;

  /** any examples for this method */
  examples: Array<MethodExample>;
}

interface PageableMethodBase extends MethodBase {
  nextLinkName?: string;

  nextPageMethod?: NextPageMethod;
}

class MethodBase implements MethodBase {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    if (statusCodes.length === 0) {
      throw new CodeModelError('statusCodes cannot be empty');
    }
    this.apiVersions = new Array<string>();
    this.client = client;
    this.httpMethod = httpMethod;
    this.httpPath = httpPath;
    this.httpStatusCodes = statusCodes;
    this.name = name;
    this.naming = naming;
    this.parameters = new Array<param.Parameter>();
    this.examples = [];
    this.docs = {};
  }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

export class Client implements Client {
  constructor(name: string, docs: type.Docs, options: ClientOptions) {
    this.name = name;
    this.templatedHost = false;
    this.constructors = new Array<Constructor>();
    this.docs = docs;
    this.methods = new Array<MethodType>();
    this.clientAccessors = new Array<ClientAccessor>();
    this.parameters = new Array<param.Parameter>();
    this.options = options;
  }
}

export class ClientAccessor implements ClientAccessor {
  constructor(name: string, subClient: Client) {
    this.name = name;
    this.subClient = subClient;
  }
}

export class Constructor implements Constructor {
  constructor(name: string) {
    this.name = name;
    this.parameters = new Array<param.Parameter>();
  }
}

export class LROMethod extends MethodBase implements LROMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    super(name, client, httpPath, httpMethod, statusCodes, naming);
    this.kind = 'lroMethod';
  }
}

export class LROPageableMethod extends MethodBase implements LROPageableMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    super(name, client, httpPath, httpMethod, statusCodes, naming);
    this.kind = 'lroPageableMethod';
  }
}

export class Method extends MethodBase implements Method {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    super(name, client, httpPath, httpMethod, statusCodes, naming);
    this.kind = 'method';
  }
}

export class MethodNaming implements MethodNaming {
  constructor(internalMethod: string, requestMethod: string, responseMethod: string) {
    this.internalMethod = internalMethod;
    this.requestMethod = requestMethod;
    this.responseMethod = responseMethod;
  }
}

export class NextPageMethod implements NextPageMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>) {
    if (statusCodes.length === 0) {
      throw new CodeModelError('statusCodes cannot be empty');
    }
    this.kind = 'nextPageMethod';
    this.apiVersions = new Array<string>();
    this.client = client;
    this.httpMethod = httpMethod;
    this.httpPath = httpPath;
    this.httpStatusCodes = statusCodes;
    this.name = name;
    this.parameters = new Array<param.Parameter>();
  }
}

export class PageableMethod extends MethodBase implements PageableMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    super(name, client, httpPath, httpMethod, statusCodes, naming);
    this.kind = 'pageableMethod';
  }
}
