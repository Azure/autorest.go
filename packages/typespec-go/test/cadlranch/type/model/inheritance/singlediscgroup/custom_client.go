//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package singlediscgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewSingleDiscriminatorClient(options *azcore.ClientOptions) (*SingleDiscriminatorClient, error) {
	internal, err := azcore.NewClient("singlediscgroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &SingleDiscriminatorClient{
		internal: internal,
		endpoint: "http://localhost:3000",
	}, nil
}
