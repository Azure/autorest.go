/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { CodeModel } from '@azure-tools/codemodel';

// the name of the internal protocol package
export const InternalPackage = 'azinternal';

// the import path of the internal protocol package
export async function InternalPackagePath(session: Session<CodeModel>): Promise<string> {
  const module = await session.getValue('module-path');
  const namespace = await session.getValue('namespace');
  return `${module}/internal/${namespace}`;
}
