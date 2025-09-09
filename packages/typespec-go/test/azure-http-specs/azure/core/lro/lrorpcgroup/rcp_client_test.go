// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrorpcgroup_test

import (
	"context"
	"lrorpcgroup"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

type apiVersionPolicy struct {
	apiVersion string
}

func (a *apiVersionPolicy) Do(req *policy.Request) (*http.Response, error) {
	rawQP := req.Raw().URL.Query()
	rawQP.Set("api-version", a.apiVersion)
	req.Raw().URL.RawQuery = rawQP.Encode()
	return req.Next()
}

func TestRpcClient_BeginLongRunningRPC(t *testing.T) {
	client, err := lrorpcgroup.NewRPCClientWithNoCredential("http://localhost:3000", &lrorpcgroup.RPCClientOptions{
		azcore.ClientOptions{
			PerCallPolicies: []policy.Policy{&apiVersionPolicy{apiVersion: "2022-12-01-preview"}},
		},
	})
	require.NoError(t, err)
	poller, err := client.BeginLongRunningRPC(context.Background(), lrorpcgroup.GenerationOptions{
		Prompt: to.Ptr("text"),
	}, nil)
	require.NoError(t, err)
	resp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{
		Frequency: time.Second,
	})
	require.NoError(t, err)
	require.Equal(t, lrorpcgroup.GenerationResult{
		Data: to.Ptr("text data"),
	}, resp.GenerationResult)
}
