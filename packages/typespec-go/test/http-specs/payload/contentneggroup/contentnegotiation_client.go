// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package contentneggroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewContentNegotiationClient(options *azcore.ClientOptions) (*ContentNegotiationClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ContentNegotiationClient{
		internal: internal,
	}, nil
}
