// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrorpcgroup

import (
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewRPCClient(options *azcore.ClientOptions) (*RPCClient, error) {
	const apiVersion = "2022-12-01-preview"
	internal, err := azcore.NewClient("lrorpcgroup", "v0.1.0", runtime.PipelineOptions{
		PerCall: []policy.Policy{&apiVersionPolicy{apiVersion: apiVersion}},
	}, options)
	if err != nil {
		return nil, err
	}
	return &RPCClient{
		internal: internal,
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
