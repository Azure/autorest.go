// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package jmergepatchgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewJSONMergePatchClient(options *azcore.ClientOptions) (*JSONMergePatchClient, error) {
	internal, err := azcore.NewClient("jmergepatchgroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &JSONMergePatchClient{
		internal: internal,
	}, nil
}
