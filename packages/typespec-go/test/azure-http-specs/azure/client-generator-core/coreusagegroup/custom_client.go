// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreusagegroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewUsageClient(options *azcore.ClientOptions) (*UsageClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &UsageClient{
		internal: internal,
	}, nil
}
