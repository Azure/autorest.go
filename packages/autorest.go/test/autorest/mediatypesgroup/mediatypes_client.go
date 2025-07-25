//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package mediatypesgroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewMediaTypesClient(endpoint string, options *azcore.ClientOptions) (*MediaTypesClient, error) {
	client, err := azcore.NewClient("mediatypesgroup.MediaTypesClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &MediaTypesClient{
		internal: client,
		endpoint: endpoint,
	}, nil
}
