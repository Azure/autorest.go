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

func TestIndividuallyNestedWithMultipleClient_WithQuery(t *testing.T) {
	client, err := clientinitindividuallygroup.NewIndividuallyNestedWithMultipleClientWithNoCredential("http://localhost:3000", "test-name-value", "us-west", nil)
	require.NoError(t, err)
	resp, err := client.WithQuery(context.Background(), &clientinitindividuallygroup.IndividuallyNestedWithMultipleClientWithQueryOptions{
		Format: to.Ptr("text"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestIndividuallyNestedWithMultipleClient_GetStandalone(t *testing.T) {
	client, err := clientinitindividuallygroup.NewIndividuallyNestedWithMultipleClientWithNoCredential("http://localhost:3000", "test-name-value", "us-west", nil)
	require.NoError(t, err)
	resp, err := client.GetStandalone(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestIndividuallyNestedWithMultipleClient_DeleteStandalone(t *testing.T) {
	client, err := clientinitindividuallygroup.NewIndividuallyNestedWithMultipleClientWithNoCredential("http://localhost:3000", "test-name-value", "us-west", nil)
	require.NoError(t, err)
	resp, err := client.DeleteStandalone(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
