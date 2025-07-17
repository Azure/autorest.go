// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewPathsClient(endpoint string, options *azcore.ClientOptions) (*PathsClient, error) {
	client, err := azcore.NewClient("urlgroup.PathsClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &PathsClient{
		internal: client,
		endpoint: endpoint,
	}, nil
}
