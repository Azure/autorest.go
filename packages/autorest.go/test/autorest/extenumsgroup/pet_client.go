// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package extenumsgroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewPetClient(endpoint string, options *azcore.ClientOptions) (*PetClient, error) {
	client, err := azcore.NewClient("extenumsgroup.PetClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &PetClient{
		internal: client,
		endpoint: endpoint,
	}, nil
}
