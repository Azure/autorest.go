// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewLROsClient(options *azcore.ClientOptions) (*LROsClient, error) {
	cl, err := azcore.NewClient("lrogroup.LROsClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &LROsClient{
		internal: cl,
	}
	return client, nil
}
