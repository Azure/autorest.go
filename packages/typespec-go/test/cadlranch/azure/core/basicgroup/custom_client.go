//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package basicgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewBasicClient(options *azcore.ClientOptions) (*BasicClient, error) {
	internal, err := azcore.NewClient("basicgroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &BasicClient{
		internal:   internal,
		apiVersion: "2022-12-01-preview",
	}, nil
}
