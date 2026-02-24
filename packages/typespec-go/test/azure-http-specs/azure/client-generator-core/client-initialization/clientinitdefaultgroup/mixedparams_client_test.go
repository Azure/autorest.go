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

func TestMixedParamsClient_WithQuery(t *testing.T) {
	client, err := clientinitdefaultgroup.NewMixedParamsClientWithNoCredential("http://localhost:3000", "test-name-value", nil)
	require.NoError(t, err)
	resp, err := client.WithQuery(context.Background(), "us-west", "test-id", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestMixedParamsClient_WithBody(t *testing.T) {
	client, err := clientinitdefaultgroup.NewMixedParamsClientWithNoCredential("http://localhost:3000", "test-name-value", nil)
	require.NoError(t, err)
	resp, err := client.WithBody(context.Background(), "us-west", clientinitdefaultgroup.WithBodyRequest{
		Name: to.Ptr("test-name"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
