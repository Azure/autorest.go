// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexmodelgroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewComplexModelClient(endpoint string, options *azcore.ClientOptions) (*ComplexModelClient, error) {
	client, err := azcore.NewClient("complexmodelgroup.ComplexModelClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ComplexModelClient{
		internal: client,
		endpoint: endpoint,
	}, nil
}
