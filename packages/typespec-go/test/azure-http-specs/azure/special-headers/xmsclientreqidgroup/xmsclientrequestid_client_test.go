// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmsclientreqidgroup_test

import (
	"context"
	"net/http"
	"testing"
	"xmsclientreqidgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/stretchr/testify/require"
)

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

func TestXMSClientRequestIDClient_Get(t *testing.T) {
	// TODO: https://github.com/Azure/typespec-azure/issues/155 causes ClientRequestID optional param
	_ = xmsclientreqidgroup.XMSClientRequestIDClientGetOptions{
		ClientRequestID: nil, // this should evaporate
	}
	client, err := xmsclientreqidgroup.NewXMSClientRequestIDClientWithNoCredential("http://localhost:3000", &xmsclientreqidgroup.XMSClientRequestIDClientOptions{
		azcore.ClientOptions{
			PerCallPolicies: []policy.Policy{
				newRequestIDPolicy(),
			},
		},
	})
	require.NoError(t, err)
	var httpResp *http.Response
	resp, err := client.Get(policy.WithCaptureResponse(context.Background(), &httpResp), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
	require.EqualValues(t, httpResp.Header.Get("x-ms-client-request-id"), "00000000-0000-0000-0000-000000000000")
}
