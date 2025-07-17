// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package corescalargroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewScalarClient(options *azcore.ClientOptions) (*ScalarClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, nil)
	if err != nil {
		return nil, err
	}
	return &ScalarClient{
		internal: internal,
	}, nil
}
