// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package examplebasicgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewBasicServiceOperationGroupClient(options *azcore.ClientOptions) (*BasicServiceOperationGroupClient, error) {
	internal, err := azcore.NewClient("examplebasicgroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &BasicServiceOperationGroupClient{
		internal: internal,
	}, nil
}
