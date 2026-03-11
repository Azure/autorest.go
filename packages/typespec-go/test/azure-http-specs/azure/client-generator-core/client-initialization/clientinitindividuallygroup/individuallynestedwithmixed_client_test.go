// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitindividuallygroup_test

import (
	"clientinitindividuallygroup"
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestIndividuallyNestedWithMixedClient_WithQuery(t *testing.T) {
	client, err := clientinitindividuallygroup.NewIndividuallyNestedWithMixedClientWithNoCredential("http://localhost:3000", "test-name-value", nil)
	require.NoError(t, err)
	resp, err := client.WithQuery(context.Background(), "us-west", &clientinitindividuallygroup.IndividuallyNestedWithMixedClientWithQueryOptions{
		Format: to.Ptr("text"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestIndividuallyNestedWithMixedClient_GetStandalone(t *testing.T) {
	client, err := clientinitindividuallygroup.NewIndividuallyNestedWithMixedClientWithNoCredential("http://localhost:3000", "test-name-value", nil)
	require.NoError(t, err)
	resp, err := client.GetStandalone(context.Background(), "us-west", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestIndividuallyNestedWithMixedClient_DeleteStandalone(t *testing.T) {
	client, err := clientinitindividuallygroup.NewIndividuallyNestedWithMixedClientWithNoCredential("http://localhost:3000", "test-name-value", nil)
	require.NoError(t, err)
	resp, err := client.DeleteStandalone(context.Background(), "us-west", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
