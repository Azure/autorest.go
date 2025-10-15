// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitializationgroup_test

import (
	"context"
	"testing"

	"clientinitializationgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestParentChildClient_GetStandalone(t *testing.T) {
	client, err := clientinitializationgroup.NewParentChildClientWithNoCredential("sample-blob", "http://localhost:3000", nil)
	require.NoError(t, err)
	require.NotNil(t, client)

	resp, err := client.GetStandalone(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestParentChildClient_DeleteStandalone(t *testing.T) {
	client, err := clientinitializationgroup.NewParentChildClientWithNoCredential("sample-blob", "http://localhost:3000", nil)
	require.NoError(t, err)
	require.NotNil(t, client)

	resp, err := client.DeleteStandalone(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestParentChildClient_WithQuery(t *testing.T) {
	client, err := clientinitializationgroup.NewParentChildClientWithNoCredential("sample-blob", "http://localhost:3000", nil)
	require.NoError(t, err)
	require.NotNil(t, client)

	opts := &clientinitializationgroup.ParentChildClientWithQueryOptions{
		Format: to.Ptr("text"),
	}
	resp, err := client.WithQuery(context.Background(), opts)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
