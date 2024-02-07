//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package addlpropsgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewExtendsModelClient(options *azcore.ClientOptions) (*ExtendsModelClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &ExtendsModelClient{
		internal: internal,
	}, nil
}

func NewExtendsModelArrayClient(options *azcore.ClientOptions) (*ExtendsModelArrayClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &ExtendsModelArrayClient{
		internal: internal,
	}, nil
}

func NewExtendsStringClient(options *azcore.ClientOptions) (*ExtendsStringClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &ExtendsStringClient{
		internal: internal,
	}, nil
}

func NewExtendsUnknownClient(options *azcore.ClientOptions) (*ExtendsUnknownClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &ExtendsUnknownClient{
		internal: internal,
	}, nil
}

func NewIsFloatClient(options *azcore.ClientOptions) (*IsFloatClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &IsFloatClient{
		internal: internal,
	}, nil
}

func NewIsModelClient(options *azcore.ClientOptions) (*IsModelClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &IsModelClient{
		internal: internal,
	}, nil
}

func NewIsModelArrayClient(options *azcore.ClientOptions) (*IsModelArrayClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &IsModelArrayClient{
		internal: internal,
	}, nil
}

func NewIsStringClient(options *azcore.ClientOptions) (*IsStringClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &IsStringClient{
		internal: internal,
	}, nil
}

func NewIsUnknownClient(options *azcore.ClientOptions) (*IsUnknownClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &IsUnknownClient{
		internal: internal,
	}, nil
}

func NewExtendsFloatClient(options *azcore.ClientOptions) (*ExtendsFloatClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &ExtendsFloatClient{
		internal: internal,
	}, nil
}

func newClient(options *azcore.ClientOptions) (*azcore.Client, error) {
	return azcore.NewClient("addlpropsgroup", "v0.1.0", runtime.PipelineOptions{}, options)
}
