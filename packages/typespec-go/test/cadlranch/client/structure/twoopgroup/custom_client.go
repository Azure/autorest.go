//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package twoopgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewTwoOperationGroupClient(options *azcore.ClientOptions) (*TwoOperationGroupClient, error) {
	internal, err := azcore.NewClient("twoopgroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &TwoOperationGroupClient{
		internal: internal,
		endpoint: "http://localhost:3000",
		client:   ClientTypeTwoOperationGroup,
	}, nil
}
