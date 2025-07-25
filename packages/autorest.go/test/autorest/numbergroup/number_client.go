// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package numbergroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewNumberClient(endpoint string, options *azcore.ClientOptions) (*NumberClient, error) {
	client, err := azcore.NewClient("numbergroup.NumberClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &NumberClient{
		internal: client,
		endpoint: endpoint,
	}, nil
}
