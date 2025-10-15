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

func TestMultipleParamsClient_WithBody(t *testing.T) {
	client, err := clientinitializationgroup.NewMultipleParamsClientWithNoCredential("http://localhost:3000", "test-name-value", "us-west", nil)
	require.NoError(t, err)
	require.NotNil(t, client)

	body := clientinitializationgroup.Input{
		Name: to.Ptr("test-name"),
	}

	resp, err := client.WithBody(context.Background(), body, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestMultipleParamsClient_WithQuery(t *testing.T) {
	client, err := clientinitializationgroup.NewMultipleParamsClientWithNoCredential("http://localhost:3000", "test-name-value", "us-west", nil)
	require.NoError(t, err)
	require.NotNil(t, client)

	resp, err := client.WithQuery(context.Background(), "test-id", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
