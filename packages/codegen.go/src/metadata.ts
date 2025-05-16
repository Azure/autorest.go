/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../codemodel.go/src/index.js';

// Creates the content in _metadata.json
export async function generateMetadataFile(codeModel: go.CodeModel): Promise<string> {
    // Return the formatted JSON string
    return JSON.stringify(codeModel.metadata, null, 2);
}
