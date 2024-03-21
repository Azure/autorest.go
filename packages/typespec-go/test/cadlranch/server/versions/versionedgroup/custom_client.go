// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package versionedgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewVersionedClient(options *azcore.ClientOptions) (*VersionedClient, error) {
	internal, err := azcore.NewClient("versionedgroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &VersionedClient{
		internal: internal,
		endpoint: "http://localhost:3000",
	}, nil
}
