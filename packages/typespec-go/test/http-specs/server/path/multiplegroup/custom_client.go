// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package multiplegroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewMultipleClient(options *azcore.ClientOptions) (*MultipleClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &MultipleClient{
		internal:   internal,
		endpoint:   "http://localhost:3000",
		apiVersion: VersionsV10,
	}, nil
}
