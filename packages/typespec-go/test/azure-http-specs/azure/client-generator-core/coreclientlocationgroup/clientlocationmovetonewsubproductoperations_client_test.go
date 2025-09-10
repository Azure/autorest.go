// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreclientlocationgroup_test

import (
	"context"
	"coreclientlocationgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientLocationMoveToNewSubProductOperationsClient_ListProducts(t *testing.T) {
	client, err := coreclientlocationgroup.NewClientLocationClient("http://localhost:3000", nil)
	require.NoError(t, err)
	require.NotNil(t, client)
	resp, err := client.NewClientLocationMoveToNewSubClient().NewClientLocationMoveToNewSubProductOperationsClient().ListProducts(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
