//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package defaultgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewServiceClient(client ClientType, options *azcore.ClientOptions) (*ServiceClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &ServiceClient{
		internal: internal,
		endpoint: "http://localhost:3000",
		client:   client,
	}, nil
}

func newClient(options *azcore.ClientOptions) (*azcore.Client, error) {
	return azcore.NewClient("defaultgroup", "v0.1.0", runtime.PipelineOptions{}, options)
}
