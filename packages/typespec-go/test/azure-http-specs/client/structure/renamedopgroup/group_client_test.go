// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package renamedopgroup_test

import (
	"context"
	"renamedopgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRenamedOperationGroupClient_RenamedFour(t *testing.T) {
	client, err := renamedopgroup.NewRenamedOperationClientWithNoCredential("http://localhost:3000", renamedopgroup.ClientTypeRenamedOperation, nil)
	require.NoError(t, err)
	resp, err := client.NewRenamedOperationGroupClient().RenamedFour(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestRenamedOperationGroupClient_RenamedSix(t *testing.T) {
	client, err := renamedopgroup.NewRenamedOperationClientWithNoCredential("http://localhost:3000", renamedopgroup.ClientTypeRenamedOperation, nil)
	require.NoError(t, err)
	resp, err := client.NewRenamedOperationGroupClient().RenamedSix(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestRenamedOperationGroupClient_RenamedTwo(t *testing.T) {
	client, err := renamedopgroup.NewRenamedOperationClientWithNoCredential("http://localhost:3000", renamedopgroup.ClientTypeRenamedOperation, nil)
	require.NoError(t, err)
	resp, err := client.NewRenamedOperationGroupClient().RenamedTwo(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
