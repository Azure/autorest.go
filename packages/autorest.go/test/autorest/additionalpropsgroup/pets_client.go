//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package additionalpropsgroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// NewPetsClient creates a new instance of PetsClient with the specified values.
func NewPetsClient(options *azcore.ClientOptions) (*PetsClient, error) {
	cl, err := azcore.NewClient("additionalpropsgroup.PetsClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &PetsClient{
		internal: cl,
	}
	return client, nil
}
