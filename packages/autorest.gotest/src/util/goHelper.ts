/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as _ from 'lodash';
import { Helper } from '@autorest/testmodeler/dist/src/util/helper';

export class GoHelper {
  public static addPackage(type: string, packageName: string) {
    let result = '';
    let tmpType = '';
    let pos = 0;
    while (pos < type.length) {
      if (type[pos] === '[' || type[pos] === ']' || type[pos] === '*') {
        if (tmpType !== '') {
          if (tmpType[0] === tmpType[0].toLowerCase()) {
            result += tmpType;
          } else {
            result += packageName + '.' + tmpType;
          }
          tmpType = '';
        }
        result += type[pos];
      } else {
        tmpType += type[pos];
      }
      pos++;
    }
    if (tmpType !== '') {
      if (tmpType[0] === tmpType[0].toLowerCase()) {
        result += tmpType;
      } else {
        result += packageName + '.' + tmpType;
      }
      tmpType = '';
    }
    return result;
  }
}
