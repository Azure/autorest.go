// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitdefaultgroup_test

import (
	"clientinitdefaultgroup"
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestPathParamClient_WithQuery(t *testing.T) {
	client, err := clientinitdefaultgroup.NewPathParamClientWithNoCredential("http://localhost:3000", "sample-blob", nil)
	require.NoError(t, err)
	resp, err := client.WithQuery(context.Background(), &clientinitdefaultgroup.PathParamClientWithQueryOptions{
		Format: to.Ptr("text"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestPathParamClient_GetStandalone(t *testing.T) {
	client, err := clientinitdefaultgroup.NewPathParamClientWithNoCredential("http://localhost:3000", "sample-blob", nil)
	require.NoError(t, err)
	resp, err := client.GetStandalone(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, "sample-blob", *resp.Name)
	require.EqualValues(t, int64(42), *resp.Size)
	require.EqualValues(t, "text/plain", *resp.ContentType)
}

func TestPathParamClient_DeleteStandalone(t *testing.T) {
	client, err := clientinitdefaultgroup.NewPathParamClientWithNoCredential("http://localhost:3000", "sample-blob", nil)
	require.NoError(t, err)
	resp, err := client.DeleteStandalone(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
