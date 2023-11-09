// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azalias

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewClient(endpoint string, options *azcore.ClientOptions) (*Client, error) {
	client, err := azcore.NewClient("azalias.Client", "v0.0.1", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &Client{
		internal: client,
		endpoint: endpoint,
	}, nil
}
