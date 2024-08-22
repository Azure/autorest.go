// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreusagegroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewUsageClient(options *azcore.ClientOptions) (*UsageClient, error) {
	internal, err := azcore.NewClient("coreusagegroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &UsageClient{
		internal: internal,
		endpoint: "http://localhost:3000",
	}, nil
}
