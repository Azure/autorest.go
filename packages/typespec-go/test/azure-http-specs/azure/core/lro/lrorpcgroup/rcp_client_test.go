// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrorpcgroup_test

import (
	"context"
	"lrorpcgroup"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestRpcClient_BeginLongRunningRPC(t *testing.T) {
	client, err := lrorpcgroup.NewRPCClient(nil)
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
