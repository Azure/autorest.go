// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package datetimegroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewDatetimeClient(endpoint string, options *azcore.ClientOptions) (*DatetimeClient, error) {
	client, err := azcore.NewClient("datetimegroup.DatetimeClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &DatetimeClient{
		internal: client,
		endpoint: endpoint,
	}, nil
}
