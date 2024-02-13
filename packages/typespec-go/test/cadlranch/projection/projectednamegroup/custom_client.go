//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package projectednamegroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewProjectedNameClient(options *azcore.ClientOptions) (*ProjectedNameClient, error) {
	internal, err := azcore.NewClient("projectednamegroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ProjectedNameClient{
		internal: internal,
	}, nil
}
