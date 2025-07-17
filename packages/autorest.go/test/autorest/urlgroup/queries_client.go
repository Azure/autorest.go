// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewQueriesClient(endpoint string, options *azcore.ClientOptions) (*QueriesClient, error) {
	client, err := azcore.NewClient("urlgroup.QueriesClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &QueriesClient{
		internal: client,
		endpoint: endpoint,
	}, nil
}
