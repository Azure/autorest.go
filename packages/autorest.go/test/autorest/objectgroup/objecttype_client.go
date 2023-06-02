// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package objectgroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewObjectTypeClient(options *azcore.ClientOptions) (*ObjectTypeClient, error) {
	client, err := azcore.NewClient("objectgroup.ObjectTypeClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ObjectTypeClient{internal: client}, nil
}
