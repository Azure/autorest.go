// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package unversionedgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewNotVersionedClient(options *azcore.ClientOptions) (*NotVersionedClient, error) {
	internal, err := azcore.NewClient("unversionedgroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &NotVersionedClient{
		internal: internal,
		endpoint: "http://localhost:3000",
	}, nil
}
