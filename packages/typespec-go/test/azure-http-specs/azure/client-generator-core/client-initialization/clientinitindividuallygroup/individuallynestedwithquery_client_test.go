// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitindividuallygroup_test

import (
	"clientinitindividuallygroup"
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestIndividuallyNestedWithQueryClient_WithQuery(t *testing.T) {
	client, err := clientinitindividuallygroup.NewIndividuallyNestedWithQueryClientWithNoCredential("http://localhost:3000", "test-blob", nil)
	require.NoError(t, err)
	resp, err := client.WithQuery(context.Background(), &clientinitindividuallygroup.IndividuallyNestedWithQueryClientWithQueryOptions{
		Format: to.Ptr("text"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestIndividuallyNestedWithQueryClient_GetStandalone(t *testing.T) {
	client, err := clientinitindividuallygroup.NewIndividuallyNestedWithQueryClientWithNoCredential("http://localhost:3000", "test-blob", nil)
	require.NoError(t, err)
	resp, err := client.GetStandalone(context.Background(), nil)
	require.NoError(t, err)
	expectedTime, err := time.Parse(time.RFC3339, "2023-01-01T12:00:00Z")
	require.NoError(t, err)
	require.EqualValues(t, clientinitindividuallygroup.BlobProperties{
		Name:        to.Ptr("test-blob"),
		Size:        to.Ptr[int64](1024),
		ContentType: to.Ptr("application/octet-stream"),
		CreatedOn:   to.Ptr(expectedTime),
	}, resp.BlobProperties)
}

func TestIndividuallyNestedWithQueryClient_DeleteStandalone(t *testing.T) {
	client, err := clientinitindividuallygroup.NewIndividuallyNestedWithQueryClientWithNoCredential("http://localhost:3000", "test-blob", nil)
	require.NoError(t, err)
	resp, err := client.DeleteStandalone(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
