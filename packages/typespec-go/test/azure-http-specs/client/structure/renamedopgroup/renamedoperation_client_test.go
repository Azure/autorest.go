//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package renamedopgroup_test

import (
	"context"
	"renamedopgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRenamedOperationClient_RenamedFive(t *testing.T) {
	client, err := renamedopgroup.NewRenamedOperationClientWithNoCredential("http://localhost:3000", renamedopgroup.ClientTypeRenamedOperation, nil)
	require.NoError(t, err)
	resp, err := client.RenamedFive(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestRenamedOperationClient_RenamedOne(t *testing.T) {
	client, err := renamedopgroup.NewRenamedOperationClientWithNoCredential("http://localhost:3000", renamedopgroup.ClientTypeRenamedOperation, nil)
	require.NoError(t, err)
	resp, err := client.RenamedOne(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestRenamedOperationClient_RenamedThree(t *testing.T) {
	client, err := renamedopgroup.NewRenamedOperationClientWithNoCredential("http://localhost:3000", renamedopgroup.ClientTypeRenamedOperation, nil)
	require.NoError(t, err)
	resp, err := client.RenamedThree(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
