// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package statuscoderangegroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewStatusCodeRangeClient(endpoint string, options *azcore.ClientOptions) (*StatusCodeRangeClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &StatusCodeRangeClient{
		internal: internal,
		endpoint: endpoint,
	}, nil
}
