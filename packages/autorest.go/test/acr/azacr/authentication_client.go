//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azacr

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

type AuthenticationClientOptions struct {
	azcore.ClientOptions
}

func NewAuthenticationClient(endpoint string, options *AuthenticationClientOptions) (*AuthenticationClient, error) {
	if options == nil {
		options = &AuthenticationClientOptions{}
	}
	client, err := azcore.NewClient("azacr.AuthenticationClient", "v0.1.0", runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	return &AuthenticationClient{
		internal: client,
		endpoint: endpoint,
	}, nil
}
