// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package pageablegroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewPageableServerDrivenPaginationClient(options *azcore.ClientOptions) (*PageableServerDrivenPaginationClient, error) {
	client, err := azcore.NewClient("pageablegroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &PageableServerDrivenPaginationClient{
		internal: client,
	}, nil
}
