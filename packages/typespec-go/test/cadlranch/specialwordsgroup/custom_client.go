//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package specialwordsgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewModelPropertiesClient(options *azcore.ClientOptions) (*ModelPropertiesClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &ModelPropertiesClient{
		internal: internal,
	}, nil
}

func NewModelsClient(options *azcore.ClientOptions) (*ModelsClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &ModelsClient{
		internal: internal,
	}, nil
}

func NewOperationsClient(options *azcore.ClientOptions) (*OperationsClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &OperationsClient{
		internal: internal,
	}, nil
}

func NewParametersClient(options *azcore.ClientOptions) (*ParametersClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &ParametersClient{
		internal: internal,
	}, nil
}

func newClient(options *azcore.ClientOptions) (*azcore.Client, error) {
	return azcore.NewClient("specialwordsgroup", "v0.1.0", runtime.PipelineOptions{}, options)
}
