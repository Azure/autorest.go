// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package srvdrivennewgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewResiliencyServiceDrivenClient(apiVersion string, options *azcore.ClientOptions) (*ResiliencyServiceDrivenClient, error) {
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ResiliencyServiceDrivenClient{
		internal:                 internal,
		endpoint:                 "http://localhost:3000",
		serviceDeploymentVersion: "v2",
		apiVersion:               apiVersion,
	}, nil
}
