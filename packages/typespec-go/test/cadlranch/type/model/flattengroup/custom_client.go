// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package flattengroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewFlattenClient(options *azcore.ClientOptions) (*FlattenClient, error) {
	internal, err := azcore.NewClient("flattengroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &FlattenClient{
		internal: internal,
		endpoint: "http://localhost:3000",
	}, nil
}
