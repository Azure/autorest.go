// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package apiversionheadergroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewHeaderClient(options *azcore.ClientOptions) (*HeaderClient, error) {
	internal, err := azcore.NewClient("apiversionheadergroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &HeaderClient{
		internal: internal,
	}, nil
}
