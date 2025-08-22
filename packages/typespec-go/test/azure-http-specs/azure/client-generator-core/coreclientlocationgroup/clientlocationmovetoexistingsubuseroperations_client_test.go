// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreclientlocationgroup_test

import (
	"context"
	"coreclientlocationgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientLocationMoveToExistingSubUserOperationsClient_GetUser(t *testing.T) {
	factory, err := coreclientlocationgroup.NewClientLocationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	client := factory.NewClientLocationMoveToExistingSubClient().NewClientLocationMoveToExistingSubUserOperationsClient()
	require.NotNil(t, client)
	resp, err := client.GetUser(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
