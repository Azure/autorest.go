// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlgroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewXMLClient(endpoint string, options *azcore.ClientOptions) (*XMLClient, error) {
	client, err := azcore.NewClient("xmlgroup.XMLClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &XMLClient{internal: client, endpoint: endpoint}, nil
}
