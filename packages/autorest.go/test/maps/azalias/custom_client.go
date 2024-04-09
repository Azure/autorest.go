// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azalias

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

type ClientOptions struct {
	azcore.ClientOptions
	Geography *Geography
}

func NewClient(options *ClientOptions) (*Client, error) {
	if options == nil {
		options = &ClientOptions{}
	}
	client, err := azcore.NewClient("Client", "v0.0.1", runtime.PipelineOptions{}, &options.ClientOptions)
	if err != nil {
		return nil, err
	}
	geography := GeographyUs
	if options.Geography != nil {
		geography = *options.Geography
	}
	return &Client{
		internal:  client,
		geography: geography,
	}, nil
}
