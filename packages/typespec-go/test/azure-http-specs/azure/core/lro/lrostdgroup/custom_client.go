// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrostdgroup

import (
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewStandardClient(endpoint string, options *azcore.ClientOptions) (*StandardClient, error) {
	const apiVersion = "2022-12-01-preview"
	internal, err := azcore.NewClient(moduleName, moduleVersion, runtime.PipelineOptions{
		PerCall: []policy.Policy{&apiVersionPolicy{apiVersion: apiVersion}},
	}, options)
	if err != nil {
		return nil, err
	}
	return &StandardClient{
		internal: internal,
		endpoint: endpoint,
	}, nil
}

type apiVersionPolicy struct {
	apiVersion string
}

func (a *apiVersionPolicy) Do(req *policy.Request) (*http.Response, error) {
	rawQP := req.Raw().URL.Query()
	rawQP.Set("api-version", a.apiVersion)
	req.Raw().URL.RawQuery = rawQP.Encode()
	return req.Next()
}
