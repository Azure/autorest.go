// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmsclientreqidgroup

import (
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewXMSClientRequestIDClient(options *azcore.ClientOptions) (*XMSClientRequestIDClient, error) {
	internal, err := azcore.NewClient("xmsclientreqidgroup", "v0.1.0", runtime.PipelineOptions{
		PerCall: []policy.Policy{
			newRequestIDPolicy(),
		},
	}, options)
	if err != nil {
		return nil, err
	}
	return &XMSClientRequestIDClient{
		internal: internal,
	}, nil
}

type requestIDPolicy struct{}

// NewRequestIDPolicy returns a policy that add the x-ms-client-request-id header
func newRequestIDPolicy() policy.Policy {
	return &requestIDPolicy{}
}

func (r *requestIDPolicy) Do(req *policy.Request) (*http.Response, error) {
	const requestID = "x-ms-client-request-id"
	if req.Raw().Header.Get(requestID) == "" {
		req.Raw().Header.Set(requestID, "00000000-0000-0000-0000-000000000000")
	}
	return req.Next()
}
