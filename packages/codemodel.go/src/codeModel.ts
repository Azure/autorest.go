/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as module from './module.js';

/** a Go-specific abstraction over REST endpoints */
export interface CodeModel {
  /** the info for this code model */
  info: Info;

  /** the service type for this code model */
  type: CodeModelType;

  /** contains the options for this code model */
  options: Options;

  /** package metadata */
  metadata?: {};

  /** the root container of content to emit */
  root: module.ContainingModule | module.Module;
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
  /**
   * the header text to emit per file. usually contains license and copyright info.
   * the default is the MIT license with a Microsoft copyright.
   */
  headerText?: string;

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
  constructor(info: Info, type: CodeModelType, options: Options, root: module.ContainingModule | module.Module) {
    this.info = info;
    this.options = options;
    this.type = type;
    this.root = root;
  }
}

export class Info implements Info {
  constructor(title: string) {
    this.title = title;
  }
}

export class Options implements Options {
  constructor(generateFakes: boolean, injectSpans: boolean, disallowUnknownFields: boolean, generateExamples: boolean) {
    this.generateFakes = generateFakes;
    this.injectSpans = injectSpans;
    this.omitConstructors = false;
    this.disallowUnknownFields = disallowUnknownFields;
    this.generateExamples = generateExamples;
  }
}
