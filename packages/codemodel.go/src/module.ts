/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import * as client from './client.js';
import * as result from './result.js';
import * as type from './type.js';

/**
 * represents an external Go module. this is used
 * by the containing-module switch when emitting
 * packages into an existing module.
 */
export interface ContainingModule {
  kind: 'containingModule';

  /** the containing module's identity */
  identity: string;

  /** the subpackage to emit */
  package: Package;
}

/** represents a Go module */
export interface Module extends PackageBase {
  kind: 'module';

  /** the module's identity */
  identity: string;
}

/** represents a Go package */
export interface Package extends PackageBase {
  kind: 'package';

  /** the name of the package */
  name: string;

  /** the container for this package */
  parent: ContainingModule | Module | Package;
}

/** provides access to module and package contents */
export type PackageContent = Module | Package;

///////////////////////////////////////////////////////////////////////////////////////////////////
// base types
///////////////////////////////////////////////////////////////////////////////////////////////////

interface PackageBase {
  /** all of the struct model types to generate (models.go file). can be empty */
  models: Array<type.Model | type.PolymorphicModel>;

  /** all of the const types to generate (constants.go file). can be empty */
  constants: Array<type.Constant>;

  /** all of the operation clients. can be empty */
  clients: Array<client.Client>;

  /** all of the parameter groups including the options types (options.go file) */
  paramGroups: Array<type.Struct>;

  /** all of the response envelopes (responses.go file). can be empty */
  responseEnvelopes: Array<result.ResponseEnvelope>;

  /** all of the interfaces for discriminated types (interfaces.go file) */
  interfaces: Array<type.Interface>;

  /** any subpackages within this package. can be empty */
  packages: Array<Package>;
}

class PackageBase implements PackageBase {
  constructor() {
    this.clients = new Array<client.Client>();
    this.constants = new Array<type.Constant>();
    this.interfaces = new Array<type.Interface>();
    this.models = new Array<type.Model | type.PolymorphicModel>();
    this.packages = new Array<Package>();
    this.paramGroups = new Array<type.Struct>();
    this.responseEnvelopes = new Array<result.ResponseEnvelope>();
  }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

export class ContainingModule implements ContainingModule {
  constructor(identity: string) {
    this.kind = 'containingModule';
    this.identity = identity;
  }
}

export class Module extends PackageBase implements Module {
  constructor(identity: string) {
    super();
    this.kind = 'module';
    this.identity = identity;
  }
}

export class Package extends PackageBase implements Package {
  constructor(name: string, parent: ContainingModule | Module | Package) {
    super();
    this.kind = 'package';
    this.name = name;
    this.parent = parent;
  }
}
