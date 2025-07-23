// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package versionedgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewVersionedClient(endpoint string, options *azcore.ClientOptions) (*VersionedClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{
		APIVersion: runtime.APIVersionOptions{
			Name: "api-version",
		},
	}, options)
	if err != nil {
		return nil, err
	}
	return &VersionedClient{
		internal: internal,
		endpoint: endpoint,
	}, nil
}
