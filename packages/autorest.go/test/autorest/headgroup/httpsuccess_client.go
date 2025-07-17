// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package headgroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewHTTPSuccessClient(endpoint string, options *azcore.ClientOptions) (*HTTPSuccessClient, error) {
	client, err := azcore.NewClient("headgroup.HTTPSuccessClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &HTTPSuccessClient{
		internal: client,
		endpoint: endpoint,
	}, nil
}
