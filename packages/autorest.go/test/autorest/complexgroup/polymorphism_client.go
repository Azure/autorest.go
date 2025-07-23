// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewPolymorphismClient(endpoint string, options *azcore.ClientOptions) (*PolymorphismClient, error) {
	client, err := azcore.NewClient("complexgroup.PolymorphismClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &PolymorphismClient{internal: client, endpoint: endpoint}, nil
}
