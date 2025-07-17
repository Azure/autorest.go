//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package arraygroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// NewArrayClient creates a new instance of ArrayClient with the specified values.
func NewArrayClient(endpoint string, options *azcore.ClientOptions) (*ArrayClient, error) {
	cl, err := azcore.NewClient("arraygroup.ArrayClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ArrayClient{
		internal: cl,
		endpoint: endpoint,
	}
	return client, nil
}
