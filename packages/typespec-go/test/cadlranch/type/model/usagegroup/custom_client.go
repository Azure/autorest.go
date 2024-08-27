//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package usagegroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewUsageClient(options *azcore.ClientOptions) (*UsageClient, error) {
	internal, err := azcore.NewClient("modelusagegroup", "v0.1.0", runtime.PipelineOptions{}, nil)
	if err != nil {
		return nil, err
	}
	return &UsageClient{
		internal: internal,
	}, nil
}
