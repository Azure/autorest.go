// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreclientlocationgroup_test

import (
	"context"
	"coreclientlocationgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientLocationArchiveOperationsClient_ArchiveProduct(t *testing.T) {
	factory, err := coreclientlocationgroup.NewClientLocationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	client := factory.NewClientLocationArchiveOperationsClient()
	require.NotNil(t, client)
	resp, err := client.ArchiveProduct(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestClientLocationMoveToRootResourceOperationsClient_GetResource(t *testing.T) {
	factory, err := coreclientlocationgroup.NewClientLocationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	client := factory.NewClientLocationMoveToRootClient().NewClientLocationMoveToRootResourceOperationsClient()
	require.NotNil(t, client)
	resp, err := client.GetResource(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
