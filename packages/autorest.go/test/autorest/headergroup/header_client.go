// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package headergroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewHeaderClient(endpoint string, options *azcore.ClientOptions) (*HeaderClient, error) {
	client, err := azcore.NewClient("headergroup.HeaderClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &HeaderClient{
		internal: client,
		endpoint: endpoint,
	}, nil
}
