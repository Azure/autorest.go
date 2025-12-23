// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreclientlocationmovemethodparametertoclientgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMoveMethodParameterToBlobOperations_GetBlob(t *testing.T) {
	client, err := NewMoveMethodParameterToClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)

	blobClient := client.NewMoveMethodParameterToBlobOperationsClient()
	blobClient.storageAccount = "testaccount"

	resp, err := blobClient.GetBlob(context.Background(), "testcontainer", "testblob.txt", nil)
	require.NoError(t, err)

	require.NotNil(t, resp.ID)
	require.Equal(t, "blob-001", *resp.ID)
	require.NotNil(t, resp.Name)
	require.Equal(t, "testblob.txt", *resp.Name)
	require.NotNil(t, resp.Path)
	require.Equal(t, "/testcontainer/testblob.txt", *resp.Path)
	require.NotNil(t, resp.Size)
	require.Equal(t, int32(1024), *resp.Size)
}
