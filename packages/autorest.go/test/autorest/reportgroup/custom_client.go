//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package reportgroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewAutoRestReportServiceClient(options *azcore.ClientOptions) (*AutoRestReportServiceClient, error) {
	cl, err := azcore.NewClient("reportgroup.AutoRestReportServiceClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AutoRestReportServiceClient{
		internal: cl,
	}
	return client, nil
}
