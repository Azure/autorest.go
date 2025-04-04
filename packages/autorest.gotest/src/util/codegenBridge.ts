/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { GroupProperty, ImplementationLocation, Operation, OperationGroup, Parameter, SchemaResponse, SchemaType } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { aggregateParameters, isLROOperation, isPageableOperation, isSchemaResponse, sortParametersByRequired } from '../common/helpers';

// homo structureed with getAPIParametersSig() in autorest.go
export function getAPIParametersSig(op: Operation): Array<[string, string, Parameter]> {
  const methodParams = getMethodParameters(op);
  const params = new Array<[string, string, Parameter]>();
  if (!isPageableOperation(op) || isLROOperation(op)) {
    params.push(['ctx', 'context.Context', undefined]);
  }
  for (const methodParam of values(methodParams)) {
    params.push([methodParam.language.go.name, formatParameterTypeName(methodParam), methodParam]);
  }
  return params;
}

export function getClientParametersSig(group: OperationGroup): Array<[string, string, Parameter]> {
  const params = [];

  for (const parameter of values(<Array<Parameter>>(group.language.go?.clientParams || []))) {
    params.push([parameter.language.go.name, formatParameterTypeName(parameter), parameter]);
  }
  return params;
}

export function getParametersSig(params: Array<Parameter>): Array<[string, string, Parameter]> {
  const result = [];

  for (const parameter of values(<Array<Parameter>>(params || []))) {
    result.push([parameter.language.go.name, formatParameterTypeName(parameter), parameter]);
  }
  return result;
}

// homo structured with generateReturnsInfo() in autorest.go
export function generateReturnsInfo(op: Operation, apiType: 'api' | 'op' | 'handler'): Array<string> {
  let returnType = getResponseEnvelopeName(op);
  if (isLROOperation(op)) {
    switch (apiType) {
      case 'api':
        if (isPageableOperation(op)) {
          returnType = `*runtime.Poller[*runtime.Pager[${getResponseEnvelopeName(op)}]]`;
        } else {
          returnType = `*runtime.Poller[${getResponseEnvelopeName(op)}]`;
        }
        break;
      case 'handler':
        // we only have a handler for operations that return a schema
        if (isPageableOperation(op)) {
          // we need to consult the final response type name
          returnType = getResponseEnvelopeName(op);
        } else {
          throw new Error(`handler being generated for non-pageable LRO ${op.language.go.name} which is unexpected`);
        }
        break;
      case 'op':
        // change to get final response type
        // TODO need to check pageable LRO
        returnType = getResponseEnvelopeName(op);
    }
  } else if (isPageableOperation(op)) {
    switch (apiType) {
      case 'api':
      case 'op':
        // pager operations don't return an error
        return [`*runtime.Pager[${returnType}]`];
    }
  }
  return [returnType, 'error'];
}

// returns the schema response for this operation.
// if multi-response operations return first 20x.
export function getSchemaResponse(op: Operation): SchemaResponse | undefined {
  if (!op.responses) {
    return undefined;
  }
  // get the list and count of distinct schema responses
  const schemaResponses = new Array<SchemaResponse>();
  for (const response of values(op.responses)) {
    // perform the comparison by name as some responses have different objects for the same underlying response type
    if (
      isSchemaResponse(response) &&
      !values(schemaResponses)
        .where((sr) => sr.schema.language.go!.name === response.schema.language.go!.name)
        .any()
    ) {
      schemaResponses.push(response);
    }
  }
  if (schemaResponses.length === 0) {
    return undefined;
  } else if (schemaResponses.length === 1) {
    return schemaResponses[0];
  }

  // multiple schema responses, for LROs find the best fit.

  // for LROs, there are a couple of corner-cases we need to handle WRT response types.
  // 1. 200 Foo / 20x Bar - we take Foo and display a warning
  // 2. 201 Foo / 202 Bar - this is a hard error
  // 3. 200 void / 20x Bar - we take Bar
  // since we always assume responses[0] has the return type we need to fix up
  // the list of responses so that it points to the schema we select.

  // multiple schemas, find the one for 200 status code
  // note that case #3 was handled earlier
  let with200: SchemaResponse | undefined;
  for (const response of values(schemaResponses)) {
    if ((<Array<string>>response.protocol.http!.statusCodes).indexOf('200') > -1) {
      with200 = response;
      break;
    }
  }
  if (with200 === undefined) {
    throw new Error(`LRO ${op.language.go!.clientName}.${op.language.go!.name} contains multiple response types which is not supported`);
  }
  return with200;
}

// returns the type name with possible * prefix
export function formatParameterTypeName(param: Parameter): string {
  const typeName = param.schema.language.go!.name;
  // client params with default values are treated as optional
  if (param.required && !(param.implementation === ImplementationLocation.Client && param.clientDefaultValue)) {
    return typeName;
  }
  return `*${typeName}`;
}

// returns the complete collection of method parameters
export function getMethodParameters(op: Operation): Array<Parameter> {
  const params = new Array<Parameter>();
  const paramGroups = new Array<GroupProperty>();
  for (const param of values(aggregateParameters(op))) {
    if (param.implementation === ImplementationLocation.Client) {
      // client params are passed via the receiver
      continue;
    } else if (param.language.go!.paramGroup) {
      // param groups will be added after individual params
      if (!paramGroups.includes(param.language.go!.paramGroup)) {
        paramGroups.push(param.language.go!.paramGroup);
      }
      continue;
    } else if (param.schema.type === SchemaType.Constant) {
      // don't generate a parameter for a constant
      // NOTE: this check must come last as non-required optional constants
      // in header/query params get dumped into the optional params group
      continue;
    }
    params.push(param);
  }
  // move global optional params to the end of the slice
  params.sort(sortParametersByRequired);
  // add any parameter groups.  optional groups go last
  paramGroups.sort(sortParametersByRequired);
  // add the optional param group last if it's not already in the list.
  // all operations should have an optional params type.  the only exception
  // is the next link operation for pageable operations.
  if (
    op.language.go!.optionalParamGroup &&
    !values(paramGroups).any((gp) => {
      return gp.language.go!.name === op.language.go!.optionalParamGroup.language.go!.name;
    })
  ) {
    paramGroups.push(op.language.go!.optionalParamGroup);
  }
  for (const paramGroup of values(paramGroups)) {
    params.push(paramGroup);
  }
  return params;
}

// returns the response envelope type name
export function getResponseEnvelopeName(op: Operation): string {
  return op.language.go!.responseEnv.language.go!.name;
}
