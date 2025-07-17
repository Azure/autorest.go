// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package typechangedfromgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewTypeChangedFromClient(endpoint string, options *azcore.ClientOptions) (*TypeChangedFromClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &TypeChangedFromClient{
		internal: internal,
		endpoint: endpoint,
		version:  VersionsV2,
	}, nil
}
