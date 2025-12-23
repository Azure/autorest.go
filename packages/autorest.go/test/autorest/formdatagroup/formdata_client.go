// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package formdatagroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewFormdataClient(endpoint string, options *azcore.ClientOptions) (*FormdataClient, error) {
	client, err := azcore.NewClient("formdatagroup.FormdataClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &FormdataClient{
		internal: client,
		endpoint: endpoint,
	}, nil
}
