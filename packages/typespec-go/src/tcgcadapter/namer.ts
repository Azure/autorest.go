/*  ---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *  --------------------------------------------------------------------------------------------  */

import { CommonAcronyms, ReservedWords } from './mappings.js';

// make sure that reserved words are escaped
export function getEscapedReservedName(name: string, appendValue: string): string {
  if (ReservedWords.includes(name)) {
    name += appendValue;
  }

  return name;
}

// used in ensureNameCase() to track which names have already been transformed.
const gRenamed = new Map<string, boolean>();

export function ensureNameCase(name: string, lowerFirst?: boolean): string {
  if (gRenamed.has(name) && gRenamed.get(name) === lowerFirst) {
    return name;
  }
  // XMS prefix requires special handling due to too many permutations that cause weird splits in the word
  name = name.replace(new RegExp('^(xms)', 'i'), 'XMS');
  let reconstructed = '';
  const words = deconstruct(name);
  for (let i = 0; i < words.length; ++i) {
    let word = words[i];
    // for params, lower-case the first segment
    if (lowerFirst && i === 0) {
      word = word.toLowerCase();
    } else {
      for (const tla of CommonAcronyms) {
        // perform a case-insensitive match against the list of TLAs
        const match = word.match(new RegExp(tla, 'i'));
        if (match) {
          // replace the match with its upper-case version
          let toReplace = match[0];
          if (match.length === 2) {
            // a capture group was specified, use it instead
            toReplace = match[1];
          }
          word = word.replace(toReplace, toReplace.toUpperCase());
        }
      }
      // note that capitalize() will convert the following acronyms to all upper-case
      // 'ip', 'os', 'ms', 'vm'
      word = capitalize(word);
    }
    reconstructed += word;
  }
  gRenamed.set(reconstructed, lowerFirst === true);
  return reconstructed;
}

// case-preserving version of deconstruct() that also splits on more path-separator characters
function deconstruct(identifier: string): Array<string> {
  return `${identifier}`.
    replace(/([a-z]+)([A-Z])/g, '$1 $2').
    replace(/(\d+)([a-z|A-Z]+)/g, '$1 $2').
    replace(/\b([A-Z]+)([A-Z])([a-z])/, '$1 $2$3').
    split(/[\W|_|.|@|-|\s|$]+/);
}

// removes pkg from val based on some heuristics
export function trimPackagePrefix(pkg: string, val: string): string {
  // foo.Foo doesn't stutter.
  if (val.length <= pkg.length) {
    return val;
  }

  // pkg is already upper-case
  if (pkg !== val.substring(0, pkg.length).toUpperCase()) {
    return val;
  }

  // we cannot simply remove pkg from val, consider the following case:
  //   pkg = tables, val = TableServicesClient; we'd end up with ervicesClient
  // we have to ensure that pkg ends on a word-boundary, i.e. the next
  // character is upper-case.
  if (val.charAt(pkg.length) !== val.charAt(pkg.length).toUpperCase()) {
    return val;
  }
  return val.substring(pkg.length);
}

// the following was copied from @azure-tools/codegen to keep compat without dragging in YAD
const acronyms = new Set([
  'ip',
  'os',
  'ms',
  'vm',
]);

export function capitalize(str: string): string {
  if (acronyms.has(str)) {
    return str.toUpperCase();
  }
  return str ? `${str.charAt(0).toUpperCase()}${str.slice(1)}` : str;
}

export function uncapitalize(str: string): string {
  return str ? `${str.charAt(0).toLowerCase()}${str.slice(1)}` : str;
}
