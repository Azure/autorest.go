// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewPathItemsClient(globalStringPath string, globalStringQuery *string, options *azcore.ClientOptions) (*PathItemsClient, error) {
	client, err := azcore.NewClient("urlgroup.PathItemsClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &PathItemsClient{
		internal:          client,
		globalStringPath:  globalStringPath,
		globalStringQuery: globalStringQuery,
	}, nil
}
