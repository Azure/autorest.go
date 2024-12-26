//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package enumdiscgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewEnumDiscriminatorClient(options *azcore.ClientOptions) (*EnumDiscriminatorClient, error) {
	internal, err := azcore.NewClient("enumdiscgroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &EnumDiscriminatorClient{
		internal: internal,
	}, nil
}
