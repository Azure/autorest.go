// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package singlegroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewSingleClient(options *azcore.ClientOptions) (*SingleClient, error) {
	internal, err := azcore.NewClient("singlegroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &SingleClient{
		internal: internal,
		endpoint: "http://localhost:3000",
	}, nil
}
