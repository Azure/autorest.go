// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package rawjson

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func Parent(options *azcore.ClientOptions) (*azcore.Client, error) {
	return azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
}
