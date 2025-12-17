// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitgroup_test

import (
	"clientinitgroup"
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestParentChildClient_WithQuery(t *testing.T) {
	client, err := clientinitgroup.NewParentChildClientWithNoCredential("http://localhost:3000", "sample-blob", nil)
	require.NoError(t, err)
	_, err = client.WithQuery(context.Background(), &clientinitgroup.ParentChildClientWithQueryOptions{
		Format: to.Ptr("text"),
	})
	require.NoError(t, err)
}

func TestParentChildClient_GetStandalone(t *testing.T) {
	client, err := clientinitgroup.NewParentChildClientWithNoCredential("http://localhost:3000", "sample-blob", nil)
	require.NoError(t, err)
	resp, err := client.GetStandalone(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.BlobProperties)
	require.Equal(t, "sample-blob", *resp.BlobProperties.Name)
	require.Equal(t, int64(42), *resp.BlobProperties.Size)
	require.Equal(t, "text/plain", *resp.BlobProperties.ContentType)
	require.NotNil(t, resp.BlobProperties.CreatedOn)
	require.Equal(t, time.Date(2025, time.April, 1, 12, 0, 0, 0, time.UTC), *resp.BlobProperties.CreatedOn)
}

func TestParentChildClient_DeleteStandalone(t *testing.T) {
	client, err := clientinitgroup.NewParentChildClientWithNoCredential("http://localhost:3000", "sample-blob", nil)
	require.NoError(t, err)
	_, err = client.DeleteStandalone(context.Background(), nil)
	require.NoError(t, err)
}

func TestParentClient_NewParentChildClient(t *testing.T) {
	parent, err := clientinitgroup.NewParentClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	child := parent.NewParentChildClient("sample-blob")
	require.NotNil(t, child)
	_, err = child.WithQuery(context.Background(), &clientinitgroup.ParentChildClientWithQueryOptions{
		Format: to.Ptr("text"),
	})
	require.NoError(t, err)
}
