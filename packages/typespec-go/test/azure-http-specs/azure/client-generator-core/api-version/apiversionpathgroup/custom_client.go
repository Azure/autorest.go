// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package apiversionpathgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewPathClient(options *azcore.ClientOptions) (*PathClient, error) {
	internal, err := azcore.NewClient("apiversionpathgroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &PathClient{
		internal: internal,
	}, nil
}
