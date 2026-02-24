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
	require.EqualValues(t, "test-blob", *resp.Name)
	require.EqualValues(t, int64(1024), *resp.Size)
	require.EqualValues(t, "application/octet-stream", *resp.ContentType)
}

func TestIndividuallyNestedWithQueryClient_DeleteStandalone(t *testing.T) {
	client, err := clientinitindividuallygroup.NewIndividuallyNestedWithQueryClientWithNoCredential("http://localhost:3000", "test-blob", nil)
	require.NoError(t, err)
	resp, err := client.DeleteStandalone(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
