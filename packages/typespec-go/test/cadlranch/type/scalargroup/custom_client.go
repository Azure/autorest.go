//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package scalargroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewBooleanClient(options *azcore.ClientOptions) (*BooleanClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &BooleanClient{
		internal: internal,
	}, nil
}

func NewDecimal128TypeClient(options *azcore.ClientOptions) (*Decimal128TypeClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &Decimal128TypeClient{
		internal: internal,
	}, nil
}

func NewDecimal128VerifyClient(options *azcore.ClientOptions) (*Decimal128VerifyClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &Decimal128VerifyClient{
		internal: internal,
	}, nil
}

func NewDecimalTypeClient(options *azcore.ClientOptions) (*DecimalTypeClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &DecimalTypeClient{
		internal: internal,
	}, nil
}

func NewDecimalVerifyClient(options *azcore.ClientOptions) (*DecimalVerifyClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &DecimalVerifyClient{
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

func NewUnknownClient(options *azcore.ClientOptions) (*UnknownClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &UnknownClient{
		internal: internal,
	}, nil
}

func newClient(options *azcore.ClientOptions) (*azcore.Client, error) {
	return azcore.NewClient("scalargroup", "v0.1.0", runtime.PipelineOptions{}, options)
}
