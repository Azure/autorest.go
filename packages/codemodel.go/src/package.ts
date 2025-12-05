/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import * as client from './client.js';
import * as result from './result.js';
import * as type from './type.js';

/** a Go-specific abstraction over REST endpoints */
export interface CodeModel {
  /** the info for this code model */
  info: Info;

  /** the service type for this code model */
  type: CodeModelType;

  /** the Go package name for this code model */
  packageName: string;

  /** contains the options for this code model */
  options: Options;

  /** all of the struct model types to generate (models.go file). can be empty */
  models: Array<type.Model | type.PolymorphicModel>;

  /** all of the const types to generate (constants.go file). can be empty */
  constants: Array<type.Constant>;

  /**
   * all of the operation clients. can be empty (models-only build)
   */
  clients: Array<client.Client>;

  /** all of the parameter groups including the options types (options.go file) */
  paramGroups: Array<type.Struct>;

  /** all of the response envelopes (responses.go file). can be empty */
  responseEnvelopes: Array<result.ResponseEnvelope>;

  /** all of the interfaces for discriminated types (interfaces.go file) */
  interfaces: Array<type.Interface>;

  /** package metadata */
  metadata?: {};
}

/** the service type that the code model represents */
export type CodeModelType = 'azure-arm' | 'data-plane';

/** contains top-level info about the input source */
export interface Info {
  title: string;
}

/**
 * contains global options set on the CodeModel.
 * most of the values come from command-line args.
 */
export interface Options {
  /** the header text to emit per file. usually contains license and copyright info */
  headerText: string;

  /**
   * custom content for the LICENSE.txt file to be emitted.
   * the default is the MIT license with a Microsoft copyright.
   */
  licenseText?: string;

  /** indicates if fakes should be emitted. the default is false */
  generateFakes: boolean;

  /** indicates if tracing spans should be emitted. the default is false */
  injectSpans: boolean;

  /** indicates if client constructors should be omitted. the default is false */
  omitConstructors: boolean;

  /**
   * indicates whether or not to disallow unknown fields in the JSON unmarshaller.
   * reproduce the behavior of https://pkg.go.dev/encoding/json#Decoder.DisallowUnknownFields
   */
  disallowUnknownFields: boolean;

  /**
   * the module into which the package is being generated.
   * this is mutually exclusive with module
   */
  containingModule?: string;

  /**
   * the module identity including any major version suffix.
   * this is mutually exclusive with containingModule.
   */
  module?: string;

  /** custom version of azcore to use instead of the emitter's default value */
  azcoreVersion?: string;

  /** emits Go any types as []byte containing raw JSON. the default value is false */
  rawJSONAsBytes: boolean;

  /** emit slice element types by value (e.g. []string not []*string). the default value is false */
  sliceElementsByval: boolean;

  /** generates example _test.go files. the default value is false */
  generateExamples: boolean;

  /** whether or not to gather all client parameters for the client factory. the default value is true */
  factoryGatherAllParams: boolean;
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

export class CodeModel implements CodeModel {
  constructor(info: Info, type: CodeModelType, packageName: string, options: Options) {
    this.clients = new Array<client.Client>();
    this.constants = new Array<type.Constant>();
    this.info = info;
    this.interfaces = new Array<type.Interface>();
    this.models = new Array<type.Model | type.PolymorphicModel>();
    this.options = options;
    this.packageName = packageName;
    this.paramGroups = new Array<type.Struct>();
    this.responseEnvelopes = new Array<result.ResponseEnvelope>();
    this.type = type;
  }
}

export class Info implements Info {
  constructor(title: string) {
    this.title = title;
  }
}

export class Options implements Options {
  constructor(headerText: string, generateFakes: boolean, injectSpans: boolean, disallowUnknownFields: boolean, generateExamples: boolean) {
    this.headerText = headerText;
    this.generateFakes = generateFakes;
    this.injectSpans = injectSpans;
    this.omitConstructors = false;
    this.disallowUnknownFields = disallowUnknownFields;
    this.generateExamples = generateExamples;
  }
}
