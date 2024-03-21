// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrolegacygroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewLegacyClient(options *azcore.ClientOptions) (*LegacyClient, error) {
	internal, err := azcore.NewClient("lrolegacygroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &LegacyClient{
		internal: internal,
	}, nil
}
