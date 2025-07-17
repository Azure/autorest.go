// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package durationgroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewDurationClient(endpoint string, options *azcore.ClientOptions) (*DurationClient, error) {
	client, err := azcore.NewClient("durationgroup.DurationClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &DurationClient{
		internal: client,
		endpoint: endpoint,
	}, nil
}
