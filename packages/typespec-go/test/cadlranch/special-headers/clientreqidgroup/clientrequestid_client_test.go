//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientreqidgroup_test

import (
	"clientreqidgroup"
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/stretchr/testify/require"
)

func TestClientRequestIdClient_Get(t *testing.T) {
	// TODO: https://github.com/Azure/typespec-azure/issues/155 causes ClientRequestID optional param
	client, err := clientreqidgroup.NewClientRequestIdClient(nil)
	require.NoError(t, err)
	var httpResp *http.Response
	resp, err := client.Get(policy.WithCaptureResponse(context.Background(), &httpResp), &clientreqidgroup.ClientRequestIdClientGetOptions{})
	require.NoError(t, err)
	require.Zero(t, resp)
	require.EqualValues(t, httpResp.Header.Get("client-request-id"), "00000000-0000-0000-0000-000000000000")
}
