// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package srvdrivenoldgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewResiliencyServiceDrivenClient(serviceVersion string, options *azcore.ClientOptions) (*ResiliencyServiceDrivenClient, error) {
	internal, err := azcore.NewClient("srvdrivengroup", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ResiliencyServiceDrivenClient{
		internal:                 internal,
		endpoint:                 "http://localhost:3000",
		serviceDeploymentVersion: serviceVersion,
		apiVersion:               "v1",
	}, nil
}
