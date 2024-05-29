// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package rettypechangedfromgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewReturnTypeChangedFromClient(options *azcore.ClientOptions) (*ReturnTypeChangedFromClient, error) {
	internal, err := azcore.NewClient("rettypechangedfromgroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ReturnTypeChangedFromClient{
		internal: internal,
		endpoint: "http://localhost:3000",
		version:  VersionsV2,
	}, nil
}
