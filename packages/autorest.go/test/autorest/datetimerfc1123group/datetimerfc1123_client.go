// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package datetimerfc1123group

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewDatetimerfc1123Client(endpoint string, options *azcore.ClientOptions) (*Datetimerfc1123Client, error) {
	client, err := azcore.NewClient("datetimerfc1123group.Datetimerfc1123Client", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &Datetimerfc1123Client{
		internal: client,
		endpoint: endpoint,
	}, nil
}
