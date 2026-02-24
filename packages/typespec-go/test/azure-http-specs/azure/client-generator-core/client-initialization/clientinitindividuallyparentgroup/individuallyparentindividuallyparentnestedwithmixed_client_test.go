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

func TestIndividuallyParentNestedWithMixedClient_WithQuery(t *testing.T) {
	client, err := clientinitindividuallyparentgroup.NewIndividuallyParentIndividuallyParentNestedWithMixedClientWithNoCredential("http://localhost:3000", "test-name-value", nil)
	require.NoError(t, err)
	resp, err := client.WithQuery(context.Background(), "us-west", &clientinitindividuallyparentgroup.IndividuallyParentIndividuallyParentNestedWithMixedClientWithQueryOptions{
		Format: to.Ptr("text"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestIndividuallyParentNestedWithMixedClient_GetStandalone(t *testing.T) {
	client, err := clientinitindividuallyparentgroup.NewIndividuallyParentIndividuallyParentNestedWithMixedClientWithNoCredential("http://localhost:3000", "test-name-value", nil)
	require.NoError(t, err)
	resp, err := client.GetStandalone(context.Background(), "us-west", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestIndividuallyParentNestedWithMixedClient_DeleteStandalone(t *testing.T) {
	client, err := clientinitindividuallyparentgroup.NewIndividuallyParentIndividuallyParentNestedWithMixedClientWithNoCredential("http://localhost:3000", "test-name-value", nil)
	require.NoError(t, err)
	resp, err := client.DeleteStandalone(context.Background(), "us-west", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestIndividuallyParentNestedWithMixedClient_ViaParent_WithQuery(t *testing.T) {
	parent, err := clientinitindividuallyparentgroup.NewIndividuallyParentClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	client := parent.NewIndividuallyParentIndividuallyParentNestedWithMixedClient("test-name-value")
	resp, err := client.WithQuery(context.Background(), "us-west", &clientinitindividuallyparentgroup.IndividuallyParentIndividuallyParentNestedWithMixedClientWithQueryOptions{
		Format: to.Ptr("text"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}
