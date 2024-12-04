//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package renamedopgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewRenamedOperationClient(options *azcore.ClientOptions) (*RenamedOperationClient, error) {
	internal, err := azcore.NewClient("renamedopgroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &RenamedOperationClient{
		internal: internal,
		endpoint: "http://localhost:3000",
		client:   ClientTypeRenamedOperation,
	}, nil
}
