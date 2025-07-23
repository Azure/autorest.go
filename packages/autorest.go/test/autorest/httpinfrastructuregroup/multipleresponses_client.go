// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewMultipleResponsesClient(endpoint string, options *azcore.ClientOptions) (*MultipleResponsesClient, error) {
	client, err := azcore.NewClient("httpinfrastructuregroup.MultipleResponsesClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &MultipleResponsesClient{internal: client, endpoint: endpoint}, nil
}
