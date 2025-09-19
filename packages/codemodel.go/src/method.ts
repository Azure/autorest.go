/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import * as type from './type.js';

/** a Go method */
export interface Method<TReceiver, TReturns> {
  /** the name of the method */
  name: string;

  /** any docs for the method */
  docs: type.Docs;

  /** contains info about the receiver */
  receiver: Receiver<TReceiver>;

  /** the parameters passed to the method. can be empty */
  parameters: Array<Parameter>;

  /** the method's return type */
  returns?: TReturns;
}

/** a Go function or method parameter */
export interface Parameter {
  /** the name of the parameter */
  name: string;

  /** any docs for the parameter */
  docs: type.Docs;

  /** the parameter's type */
  type: type.Type;

  /** indicates if the param is pointer-to-type or not */
  byValue: boolean;
}

/** a method's receiver parameter */
export interface Receiver<T> {
  /** the receiver var name */
  name: string;

  /** the receiver param's type */
  type: T;

  /** indicates if the receiver is pointer-to-type or not */
  byValue: boolean;
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

export class Method<TReceiver, TReturns> implements Method<TReceiver, TReturns> {
  constructor(name: string, receiver: Receiver<TReceiver>) {
    this.name = name;
    this.receiver = receiver;
    this.parameters = new Array<Parameter>();
    this.docs = {};
  }
}

export class Parameter implements Parameter {
  constructor(name: string, type: type.Type, byValue: boolean) {
    this.name = name;
    this.type = type;
    this.byValue = byValue;
    this.docs = {};
  }
}

export class Receiver<T> implements Receiver<T> {
  constructor(name: string, type: T, byValue: boolean) {
    this.name = name;
    this.type = type;
    this.byValue = byValue;
  }
}
