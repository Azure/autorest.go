// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package madeoptionalgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewMadeOptionalClient(options *azcore.ClientOptions) (*MadeOptionalClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &MadeOptionalClient{
		internal: internal,
		endpoint: "http://localhost:3000",
		version:  VersionsV2,
	}, nil
}
