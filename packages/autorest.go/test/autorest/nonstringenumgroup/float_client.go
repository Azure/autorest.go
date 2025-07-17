// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nonstringenumgroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewFloatClient(endpoint string, options *azcore.ClientOptions) (*FloatClient, error) {
	client, err := azcore.NewClient("nonstringenumgroup.FloatClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &FloatClient{
		internal: client,
		endpoint: endpoint,
	}, nil
}
