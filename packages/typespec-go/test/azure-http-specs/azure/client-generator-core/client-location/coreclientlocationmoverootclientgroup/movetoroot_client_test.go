// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreclientlocationmoverootclientgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMoveToRootClients(t *testing.T) {
	client, err := NewMoveToRootClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)

	_, err = client.GetHealthStatus(context.Background(), nil)
	require.NoError(t, err)

	resourceClient := client.NewMoveToRootResourceOperationsClient()
	_, err = resourceClient.GetResource(context.Background(), nil)
	require.NoError(t, err)
}
