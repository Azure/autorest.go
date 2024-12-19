//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package multiclientgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewClientAClient(options *azcore.ClientOptions) (*ClientAClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &ClientAClient{
		internal: internal,
		endpoint: "http://localhost:3000",
		client:   ClientTypeMultiClient,
	}, err
}

func NewClientBClient(options *azcore.ClientOptions) (*ClientBClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &ClientBClient{
		internal: internal,
		endpoint: "http://localhost:3000",
		client:   ClientTypeMultiClient,
	}, err
}

func newClient(options *azcore.ClientOptions) (*azcore.Client, error) {
	return azcore.NewClient("multiclientgroup", "v0.1.0", runtime.PipelineOptions{}, options)
}
