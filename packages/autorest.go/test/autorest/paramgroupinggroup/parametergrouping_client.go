// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package paramgroupinggroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewParameterGroupingClient(endpoint string, options *azcore.ClientOptions) (*ParameterGroupingClient, error) {
	client, err := azcore.NewClient("paramgroupinggroup.ParameterGroupingClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ParameterGroupingClient{internal: client, endpoint: endpoint}, nil
}
