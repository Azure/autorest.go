//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package durationgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewHeaderClient(options *azcore.ClientOptions) (*HeaderClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &HeaderClient{
		internal: internal,
	}, nil
}

func NewPropertyClient(options *azcore.ClientOptions) (*PropertyClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &PropertyClient{
		internal: internal,
	}, nil
}

func NewQueryClient(options *azcore.ClientOptions) (*QueryClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &QueryClient{
		internal: internal,
	}, nil
}

func newClient(options *azcore.ClientOptions) (*azcore.Client, error) {
	return azcore.NewClient("durationgroup", "v0.1.0", runtime.PipelineOptions{}, options)
}
