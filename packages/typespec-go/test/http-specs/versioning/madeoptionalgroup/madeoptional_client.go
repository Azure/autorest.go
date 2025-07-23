// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package madeoptionalgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewMadeOptionalClient(endpoint string, options *azcore.ClientOptions) (*MadeOptionalClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &MadeOptionalClient{
		internal: internal,
		endpoint: endpoint,
		version:  VersionsV2,
	}, nil
}
