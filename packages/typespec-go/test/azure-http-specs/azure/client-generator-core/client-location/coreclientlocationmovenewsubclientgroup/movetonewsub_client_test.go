// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreclientlocationmovenewsubclientgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMoveToNewSubClients(t *testing.T) {
	client, err := NewMoveToNewSubClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)

	archiveClient := client.NewMoveToNewSubArchiveOperationsClient()
	_, err = archiveClient.ArchiveProduct(context.Background(), nil)
	require.NoError(t, err)

	productClient := client.NewMoveToNewSubProductOperationsClient()
	_, err = productClient.ListProducts(context.Background(), nil)
	require.NoError(t, err)
}
