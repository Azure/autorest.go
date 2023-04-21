//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurereportgroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewAutoRestReportServiceForAzureClient(options *azcore.ClientOptions) (*AutoRestReportServiceForAzureClient, error) {
	cl, err := azcore.NewClient("azurereportgroup.AutoRestReportServiceForAzureClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AutoRestReportServiceForAzureClient{
		internal: cl,
	}
	return client, nil
}
