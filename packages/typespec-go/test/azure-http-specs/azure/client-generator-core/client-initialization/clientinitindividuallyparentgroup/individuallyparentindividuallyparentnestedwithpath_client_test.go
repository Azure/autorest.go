// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitindividuallyparentgroup_test

import (
	"clientinitindividuallyparentgroup"
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestIndividuallyParentNestedWithPathClient_WithQuery(t *testing.T) {
	parentClient, err := clientinitindividuallyparentgroup.NewIndividuallyParentClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	client := parentClient.NewIndividuallyParentIndividuallyParentNestedWithPathClient("test-blob")
	resp, err := client.WithQuery(context.Background(), &clientinitindividuallyparentgroup.IndividuallyParentIndividuallyParentNestedWithPathClientWithQueryOptions{
		Format: to.Ptr("text"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestIndividuallyParentNestedWithPathClient_GetStandalone(t *testing.T) {
	parentClient, err := clientinitindividuallyparentgroup.NewIndividuallyParentClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	client := parentClient.NewIndividuallyParentIndividuallyParentNestedWithPathClient("test-blob")
	resp, err := client.GetStandalone(context.Background(), nil)
	require.NoError(t, err)
	expectedTime, err := time.Parse(time.RFC3339, "2023-01-01T12:00:00Z")
	require.NoError(t, err)
	require.EqualValues(t, clientinitindividuallyparentgroup.BlobProperties{
		Name:        to.Ptr("test-blob"),
		Size:        to.Ptr[int64](1024),
		ContentType: to.Ptr("application/octet-stream"),
		CreatedOn:   to.Ptr(expectedTime),
	}, resp.BlobProperties)
}

func TestIndividuallyParentNestedWithPathClient_DeleteStandalone(t *testing.T) {
	parentClient, err := clientinitindividuallyparentgroup.NewIndividuallyParentClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	client := parentClient.NewIndividuallyParentIndividuallyParentNestedWithPathClient("test-blob")
	resp, err := client.DeleteStandalone(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
