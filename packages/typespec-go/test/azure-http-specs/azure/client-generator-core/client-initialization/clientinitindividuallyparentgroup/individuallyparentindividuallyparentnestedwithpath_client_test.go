// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitindividuallyparentgroup_test

import (
	"clientinitindividuallyparentgroup"
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestIndividuallyParentNestedWithPathClient_WithQuery(t *testing.T) {
	client, err := clientinitindividuallyparentgroup.NewIndividuallyParentIndividuallyParentNestedWithPathClientWithNoCredential("http://localhost:3000", "test-blob", nil)
	require.NoError(t, err)
	resp, err := client.WithQuery(context.Background(), &clientinitindividuallyparentgroup.IndividuallyParentIndividuallyParentNestedWithPathClientWithQueryOptions{
		Format: to.Ptr("text"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestIndividuallyParentNestedWithPathClient_GetStandalone(t *testing.T) {
	client, err := clientinitindividuallyparentgroup.NewIndividuallyParentIndividuallyParentNestedWithPathClientWithNoCredential("http://localhost:3000", "test-blob", nil)
	require.NoError(t, err)
	resp, err := client.GetStandalone(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, "test-blob", *resp.Name)
	require.EqualValues(t, int64(1024), *resp.Size)
	require.EqualValues(t, "application/octet-stream", *resp.ContentType)
}

func TestIndividuallyParentNestedWithPathClient_DeleteStandalone(t *testing.T) {
	client, err := clientinitindividuallyparentgroup.NewIndividuallyParentIndividuallyParentNestedWithPathClientWithNoCredential("http://localhost:3000", "test-blob", nil)
	require.NoError(t, err)
	resp, err := client.DeleteStandalone(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestIndividuallyParentNestedWithPathClient_ViaParent_WithQuery(t *testing.T) {
	parent, err := clientinitindividuallyparentgroup.NewIndividuallyParentClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	client := parent.NewIndividuallyParentIndividuallyParentNestedWithPathClient("test-blob")
	resp, err := client.WithQuery(context.Background(), &clientinitindividuallyparentgroup.IndividuallyParentIndividuallyParentNestedWithPathClientWithQueryOptions{
		Format: to.Ptr("text"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}
