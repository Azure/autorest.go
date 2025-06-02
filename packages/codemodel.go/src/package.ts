/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import * as client from './client.js';
import { CodeModelError } from './errors.js';
import * as result from './result.js';
import * as type from './type.js';

// CodeModel contains a Go-specific abstraction over an OpenAPI (or other) description of REST endpoints.
export interface CodeModel {
  info: Info;

  host?: string;

  type: CodeModelType;

  packageName: string;

  options: Options;

  // all of the struct model types to generate (models.go file)
  models: Array<type.ModelType | type.PolymorphicType>;

  // all of the const types to generate (constants.go file)
  constants: Array<type.ConstantType>;

  // all of the operation groups (i.e. clients and their methods)
  // no clients indicates a models-only build
  clients: Array<client.Client>;

  // all of the parameter groups including the options types (options.go file)
  paramGroups: Array<type.StructType>;

  // all of the response envelopes (responses.go file)
  // no response envelopes indicates a models-only build
  responseEnvelopes: Array<result.ResponseEnvelope>;

  // all of the interfaces for discriminated types (interfaces.go file)
  interfaceTypes: Array<type.InterfaceType>;

  // metadata of the package
  metadata?: {};
}

export type CodeModelType = 'azure-arm' | 'data-plane';

// Info contains top-level info about the input source
export interface Info {
  title: string;
}

export interface Module {
  // the full module path excluding any major version suffix
  name: string;
  
  // the semantic version x.y.z[-beta.N]
  version: string;
}

// Options contains global options set on the CodeModel.
export interface Options {
  headerText: string;

  generateFakes: boolean;

  injectSpans: boolean;

  // disallowUnknownFields indicates whether or not to disallow unknown fields in the JSON unmarshaller.
  // reproduce the behavior of https://pkg.go.dev/encoding/json#Decoder.DisallowUnknownFields
  disallowUnknownFields: boolean;

  // NOTE: containingModule and module are mutually exclusive

  // the module into which the package is being generated
  containingModule?: string;

  // module and containingModule are mutually exclusive
  module?: Module;

  azcoreVersion?: string;

  rawJSONAsBytes: boolean;

  sliceElementsByval: boolean;

  generateExamples: boolean;

  // whether or not to gather all client parameters for the client factory.
  factoryGatherAllParams: boolean;
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

export class CodeModel implements CodeModel {
  constructor(info: Info, type: CodeModelType, packageName: string, options: Options) {
    this.clients = new Array<client.Client>();
    this.constants = new Array<type.ConstantType>();
    this.info = info;
    this.interfaceTypes = new Array<type.InterfaceType>();
    this.models = new Array<type.ModelType | type.PolymorphicType>();
    this.options = options;
    this.packageName = packageName;
    this.paramGroups = new Array<type.StructType>();
    this.responseEnvelopes = new Array<result.ResponseEnvelope>();
    this.type = type;
  }

  sortContent() {
    const sortAscending = function(a: string, b: string): number {
      return a < b ? -1 : a > b ? 1 : 0;
    };

    this.constants.sort((a: type.ConstantType, b: type.ConstantType) => { return sortAscending(a.name, b.name); });
    for (const enm of this.constants) {
      enm.values.sort((a: type.ConstantValue, b: type.ConstantValue) => { return sortAscending(a.name, b.name); });
    }
  
    this.interfaceTypes.sort((a: type.InterfaceType, b: type.InterfaceType) => { return sortAscending(a.name, b.name); });
    for (const iface of this.interfaceTypes) {
      // we sort by literal value so that the switch/case statements in polymorphic_helpers.go
      // are ordered by the literal value which can be somewhat different from the model name.
      iface.possibleTypes.sort((a: type.PolymorphicType, b: type.PolymorphicType) => { return sortAscending(a.discriminatorValue!.literal, b.discriminatorValue!.literal); });
    }
  
    this.models.sort((a: type.ModelType | type.PolymorphicType, b: type.ModelType | type.PolymorphicType) => { return sortAscending(a.name, b.name); });
    for (const model of this.models) {
      model.fields.sort((a: type.ModelField, b: type.ModelField) => { return sortAscending(a.name, b.name); });
    }
  
    this.paramGroups.sort((a: type.StructType, b: type.StructType) => { return sortAscending(a.name, b.name); });
    for (const paramGroup of this.paramGroups) {
      paramGroup.fields.sort((a: type.StructField, b: type.StructField) => { return sortAscending(a.name, b.name); });
    }
  
    this.responseEnvelopes.sort((a: result.ResponseEnvelope, b: result.ResponseEnvelope) => { return sortAscending(a.name, b.name); });
    for (const respEnv of this.responseEnvelopes) {
      respEnv.headers.sort((a: result.HeaderResponse | result.HeaderMapResponse, b: result.HeaderResponse | result.HeaderMapResponse) => { return sortAscending(a.fieldName, b.fieldName); });
    }
  
    this.clients.sort((a: client.Client, b: client.Client) => { return sortAscending(a.name, b.name); });
    for (const client of this.clients) {
      client.methods.sort((a: client.Method, b: client.Method) => { return sortAscending(a.name, b.name); });
      client.clientAccessors.sort((a: client.ClientAccessor, b: client.ClientAccessor) => { return sortAscending(a.name, b.name); });
      for (const method of client.methods) {
        method.httpStatusCodes.sort();
      }
    }
  }
}

export class Info implements Info {
  constructor(title: string) {
    this.title = title;
  }
}

export class Module implements Module {
  constructor(name: string, version: string) {
    if (name.match(/\/v\d+$/)) {
      throw new CodeModelError('module name must not contain major version suffix');
    }
    if (!version.match(/^(\d+\.\d+\.\d+(?:-beta\.\d+)?)?$/)) {
      throw new CodeModelError(`module version ${version} must be in the format major.minor.patch[-beta.N]`);
    }
    // if the module's major version is greater than one, add a major version suffix to the module name
    const majorVersion = version.substring(0, version.indexOf('.'));
    if (Number(majorVersion) > 1) {
      name += '/v' + majorVersion;
    }
    this.name = name;
    this.version = version;
  }
}

export class Options implements Options {
  constructor(headerText: string, generateFakes: boolean, injectSpans: boolean, disallowUnknownFields: boolean, generateExamples: boolean) {
    this.headerText = headerText;
    this.generateFakes = generateFakes;
    this.injectSpans = injectSpans;
    this.disallowUnknownFields = disallowUnknownFields;
    this.generateExamples = generateExamples;
  }
}
