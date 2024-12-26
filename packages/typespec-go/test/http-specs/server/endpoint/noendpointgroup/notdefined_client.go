// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package noendpointgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewNotDefinedClient(options *azcore.ClientOptions) (*NotDefinedClient, error) {
	internal, err := azcore.NewClient("noendpointgroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &NotDefinedClient{
		internal: internal,
		endpoint: "http://localhost:3000",
	}, nil
}
