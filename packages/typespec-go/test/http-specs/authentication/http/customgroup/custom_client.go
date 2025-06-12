//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package customgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewCustomClient(options *azcore.ClientOptions) (*CustomClient, error) {
	internal, err := azcore.NewClient("customgroup", "v0.1.0", runtime.PipelineOptions{
		APIVersion: runtime.APIVersionOptions{
			Location: runtime.APIVersionLocationQueryParam,
			Name:     "api-version",
		},
		AllowedHeaders: []string{"authorization"},
		Tracing: runtime.TracingOptions{
			Namespace: "Microsoft.KeyVault",
		},
	}, options)
	if err != nil {
		return nil, err
	}
	return &CustomClient{
		internal: internal,
	}, nil
}
