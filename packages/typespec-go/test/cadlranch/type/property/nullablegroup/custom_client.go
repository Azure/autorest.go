//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nullablegroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewBytesClient(options *azcore.ClientOptions) (*BytesClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &BytesClient{
		internal: internal,
	}, nil
}

func NewCollectionsByteClient(options *azcore.ClientOptions) (*CollectionsByteClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &CollectionsByteClient{
		internal: internal,
	}, nil
}

func NewCollectionsModelClient(options *azcore.ClientOptions) (*CollectionsModelClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &CollectionsModelClient{
		internal: internal,
	}, nil
}

func NewDatetimeClient(options *azcore.ClientOptions) (*DatetimeClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &DatetimeClient{
		internal: internal,
	}, nil
}

func NewDurationClient(options *azcore.ClientOptions) (*DurationClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &DurationClient{
		internal: internal,
	}, nil
}

func NewStringClient(options *azcore.ClientOptions) (*StringClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &StringClient{
		internal: internal,
	}, nil
}

func newClient(options *azcore.ClientOptions) (*azcore.Client, error) {
	return azcore.NewClient("nullablegroup", "v0.1.0", runtime.PipelineOptions{}, options)
}
