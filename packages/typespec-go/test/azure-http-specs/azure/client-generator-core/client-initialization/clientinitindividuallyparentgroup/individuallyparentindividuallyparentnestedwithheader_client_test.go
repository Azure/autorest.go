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

func TestIndividuallyParentNestedWithHeaderClient_WithQuery(t *testing.T) {
	client, err := clientinitindividuallyparentgroup.NewIndividuallyParentIndividuallyParentNestedWithHeaderClientWithNoCredential("http://localhost:3000", "test-name-value", nil)
	require.NoError(t, err)
	resp, err := client.WithQuery(context.Background(), &clientinitindividuallyparentgroup.IndividuallyParentIndividuallyParentNestedWithHeaderClientWithQueryOptions{
		Format: to.Ptr("text"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestIndividuallyParentNestedWithHeaderClient_GetStandalone(t *testing.T) {
	client, err := clientinitindividuallyparentgroup.NewIndividuallyParentIndividuallyParentNestedWithHeaderClientWithNoCredential("http://localhost:3000", "test-name-value", nil)
	require.NoError(t, err)
	resp, err := client.GetStandalone(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestIndividuallyParentNestedWithHeaderClient_DeleteStandalone(t *testing.T) {
	client, err := clientinitindividuallyparentgroup.NewIndividuallyParentIndividuallyParentNestedWithHeaderClientWithNoCredential("http://localhost:3000", "test-name-value", nil)
	require.NoError(t, err)
	resp, err := client.DeleteStandalone(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

