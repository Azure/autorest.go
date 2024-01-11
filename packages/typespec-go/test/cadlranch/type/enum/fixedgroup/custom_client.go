//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package fixedgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewStringClient(options *azcore.ClientOptions) (*StringClient, error) {
	internal, err := azcore.NewClient("fixedenumgroup", "v0.1.1", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &StringClient{
		internal: internal,
	}, nil
}
