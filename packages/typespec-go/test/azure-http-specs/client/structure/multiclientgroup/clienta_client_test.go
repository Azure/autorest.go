//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package multiclientgroup_test

import (
	"context"
	"multiclientgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientAClient_RenamedFive(t *testing.T) {
	client, err := multiclientgroup.NewClientAClientWithNoCredential("http://localhost:3000", multiclientgroup.ClientTypeMultiClient, nil)
	require.NoError(t, err)
	resp, err := client.RenamedFive(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestClientAClient_RenamedOne(t *testing.T) {
	client, err := multiclientgroup.NewClientAClientWithNoCredential("http://localhost:3000", multiclientgroup.ClientTypeMultiClient, nil)
	require.NoError(t, err)
	resp, err := client.RenamedOne(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestClientAClient_RenamedThree(t *testing.T) {
	client, err := multiclientgroup.NewClientAClientWithNoCredential("http://localhost:3000", multiclientgroup.ClientTypeMultiClient, nil)
	require.NoError(t, err)
	resp, err := client.RenamedThree(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
