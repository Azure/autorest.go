/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../codemodel.go/src/index.js';

// Creates the content in go.mod if the --module switch was specified.
// if there's a preexisting go.mod file, update its specified version of azcore as needed.
export async function generateMetadataFile(codeModel: go.CodeModel): Promise<string> {
    // Return the formatted JSON string
    return JSON.stringify(codeModel.metadata, null, 2);
}
