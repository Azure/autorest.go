// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreclientlocationmoveexistingsubclientgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMoveToExistingSubClients(t *testing.T) {
	client, err := NewMoveToExistingSubClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)

	adminClient := client.NewMoveToExistingSubAdminOperationsClient()
	_, err = adminClient.GetAdminInfo(context.Background(), nil)
	require.NoError(t, err)

	_, err = adminClient.DeleteUser(context.Background(), nil)
	require.NoError(t, err)

	userClient := client.NewMoveToExistingSubUserOperationsClient()
	_, err = userClient.GetUser(context.Background(), nil)
	require.NoError(t, err)
}
