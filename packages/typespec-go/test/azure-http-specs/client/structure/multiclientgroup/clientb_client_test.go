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

func TestClientBClient_RenamedFour(t *testing.T) {
	client, err := multiclientgroup.NewClientBClientWithNoCredential("http://localhost:3000", multiclientgroup.ClientTypeMultiClient, nil)
	require.NoError(t, err)
	resp, err := client.RenamedFour(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestClientBClient_RenamedSix(t *testing.T) {
	client, err := multiclientgroup.NewClientBClientWithNoCredential("http://localhost:3000", multiclientgroup.ClientTypeMultiClient, nil)
	require.NoError(t, err)
	resp, err := client.RenamedSix(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestClientBClient_RenamedTwo(t *testing.T) {
	client, err := multiclientgroup.NewClientBClientWithNoCredential("http://localhost:3000", multiclientgroup.ClientTypeMultiClient, nil)
	require.NoError(t, err)
	resp, err := client.RenamedTwo(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
