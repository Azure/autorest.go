//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package visibilitygroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewVisibilityClient(options *azcore.ClientOptions) (*VisibilityClient, error) {
	internal, err := azcore.NewClient("visibilitygroup", "v0.1.1", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &VisibilityClient{
		internal: internal,
		endpoint: "http://localhost:3000",
	}, nil
}
