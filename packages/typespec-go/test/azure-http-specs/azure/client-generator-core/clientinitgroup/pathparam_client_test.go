// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitgroup_test

import (
	"clientinitgroup"
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestPathParamClient_WithQuery(t *testing.T) {
	client, err := clientinitgroup.NewPathParamClientWithNoCredential("http://localhost:3000", "sample-blob", nil)
	require.NoError(t, err)
	_, err = client.WithQuery(context.Background(), &clientinitgroup.PathParamClientWithQueryOptions{
		Format: to.Ptr("text"),
	})
	require.NoError(t, err)
}

func TestPathParamClient_GetStandalone(t *testing.T) {
	client, err := clientinitgroup.NewPathParamClientWithNoCredential("http://localhost:3000", "sample-blob", nil)
	require.NoError(t, err)
	resp, err := client.GetStandalone(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, "sample-blob", *resp.Name)
	require.Equal(t, int64(42), *resp.Size)
	require.Equal(t, "text/plain", *resp.ContentType)
	require.NotNil(t, resp.CreatedOn)
	require.Equal(t, time.Date(2025, time.April, 1, 12, 0, 0, 0, time.UTC), *resp.CreatedOn)
}

func TestPathParamClient_DeleteStandalone(t *testing.T) {
	client, err := clientinitgroup.NewPathParamClientWithNoCredential("http://localhost:3000", "sample-blob", nil)
	require.NoError(t, err)
	_, err = client.DeleteStandalone(context.Background(), nil)
	require.NoError(t, err)
}
