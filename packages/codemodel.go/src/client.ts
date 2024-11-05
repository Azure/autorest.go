/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import { MethodExample } from './examples.js';
import * as pkg from './package.js';
import * as param from './param.js';
import * as result from './result.js';
import * as type from './type.js';

// Client is an SDK client
export interface Client {
  name: string;

  docs: type.Docs;

  // the client options type. for ARM, this will be a QualifiedType (arm.ClientOptions)
  options: ClientOptions;

  // constructor params that are persisted as fields on the client, can be empty
  parameters: Array<param.Parameter>;

  // all the constructors for this client, can be empty
  constructors: Array<Constructor>;

  // contains client methods. can be empty
  methods: Array<Method | LROMethod | PageableMethod | LROPageableMethod>;

  // contains any client accessor methods. can be empty
  clientAccessors: Array<ClientAccessor>;

  // client has a statically defined or templated host
  host?: string;

  // templatedHost indicates that there's one or more URIParameters
  // required to construct the complete host. the parameters can
  // be solely on the client or span client and method params.
  templatedHost: boolean;

  // the parent client in a hierarchical client
  parent?: Client;
}

export type ClientOptions = param.ParameterGroup | param.Parameter;

// represents a client constructor function
export interface Constructor {
  name: string;

  // the modeled parameters. can be empty
  parameters: Array<param.Parameter>;
}

// ClientAccessor is a client method that returns a sub-client instance.
export interface ClientAccessor {
  // the name of the client accessor method
  name: string;

  // the client returned by the accessor method
  subClient: Client;
}

// Method is a method on a client
export interface Method {
  name: string;

  docs: type.Docs;

  httpPath: string;

  httpMethod: HTTPMethod;

  // any modeled parameters. the ones we add to the generated code (context.Context etc) aren't included here
  parameters: Array<param.Parameter>;

  optionalParamsGroup: param.ParameterGroup;

  responseEnvelope: result.ResponseEnvelope;

  // the complete list of successful HTTP status codes
  httpStatusCodes: Array<number>;

  client: Client;

  naming: MethodNaming;

  apiVersions: Array<string>;

  examples: Array<MethodExample>;
}

export type HTTPMethod = 'delete' | 'get' | 'head' | 'patch' | 'post' | 'put';

export interface MethodNaming {
  // this is the name of the internal method for consumption by LROs/paging methods.
  internalMethod: string;

  requestMethod: string;

  responseMethod: string;
}

export interface LROMethod extends Method {
  finalStateVia?: 'azure-async-operation' | 'location' | 'operation-location' | 'original-uri';

  isLRO: true;
}

export interface PageableMethod extends Method {
  nextLinkName?: string;

  nextPageMethod?: NextPageMethod;

  isPageable: true;
}

// NextPageMethod is the internal method used for fetching the next page for a PageableMethod.
// It's unique from a regular Method as it's not exported and has no optional params/response envelope.
// thus, it's not included in the array of methods for a client.
export interface NextPageMethod {
  name: string;
 
  httpPath: string;

  httpMethod: HTTPMethod;

  // any modeled parameters
  parameters: Array<param.Parameter>;

  // the complete list of successful HTTP status codes
  httpStatusCodes: Array<number>;

  client: Client;

  apiVersions: Array<string>;

  isNextPageMethod: true;
}

export interface LROPageableMethod extends LROMethod, PageableMethod {
  // no new fields are added
}

export function isMethod(method: Method | NextPageMethod): method is Method {
  return (<NextPageMethod>method).isNextPageMethod === undefined;
}

export function isLROMethod(method: Method | LROMethod | PageableMethod): method is LROMethod {
  return (<LROMethod>method).isLRO === true;
}

export function isPageableMethod(method: Method | LROMethod | PageableMethod): method is PageableMethod {
  return (<PageableMethod>method).isPageable === true;
}

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

export class Method implements Method {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    if (statusCodes.length === 0) {
      throw new Error('statusCodes cannot be empty');
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
    this.methods = new Array<Method>();
    this.clientAccessors = new Array<ClientAccessor>();
    this.parameters = new Array<param.Parameter>();
    this.options = options;
  }
}

export class Constructor implements Constructor {
  constructor(name: string) {
    this.name = name;
    this.parameters = new Array<param.Parameter>();
  }
}

export class ClientAccessor implements ClientAccessor {
  constructor(name: string, subClient: Client) {
    this.name = name;
    this.subClient = subClient;
  }
}

export class MethodNaming implements MethodNaming {
  constructor(internalMethod: string, requestMethod: string, responseMethod: string) {
    this.internalMethod = internalMethod;
    this.requestMethod = requestMethod;
    this.responseMethod = responseMethod;
  }
}

export class LROMethod extends Method implements LROMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    super(name, client, httpPath, httpMethod, statusCodes, naming);
    this.isLRO = true;
  }
}

export class PageableMethod extends Method implements PageableMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    super(name, client, httpPath, httpMethod, statusCodes, naming);
    this.isPageable = true;
  }
}

export class LROPageableMethod extends Method implements LROPageableMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>, naming: MethodNaming) {
    super(name, client, httpPath, httpMethod, statusCodes, naming);
    this.isLRO = true;
    this.isPageable = true;
  }
}

export class NextPageMethod implements NextPageMethod {
  constructor(name: string, client: Client, httpPath: string, httpMethod: HTTPMethod, statusCodes: Array<number>) {
    if (statusCodes.length === 0) {
      throw new Error('statusCodes cannot be empty');
    }
    this.apiVersions = new Array<string>();
    this.client = client;
    this.httpMethod = httpMethod;
    this.httpPath = httpPath;
    this.httpStatusCodes = statusCodes;
    this.isNextPageMethod = true;
    this.name = name;
    this.parameters = new Array<param.Parameter>();
  }
}
