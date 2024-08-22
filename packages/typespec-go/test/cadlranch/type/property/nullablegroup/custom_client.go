//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nullablegroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewNullableClient(options *azcore.ClientOptions) (*NullableClient, error) {
	internal, err := azcore.NewClient("nullablegroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &NullableClient{
		internal: internal,
		endpoint: "http://localhost:3000",
	}, nil
}
