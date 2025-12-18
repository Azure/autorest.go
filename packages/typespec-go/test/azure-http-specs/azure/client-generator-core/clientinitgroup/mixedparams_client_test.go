// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitgroup_test

import (
	"clientinitgroup"
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestMixedParamsClient_WithQuery(t *testing.T) {
	client, err := clientinitgroup.NewMixedParamsClientWithNoCredential("http://localhost:3000", "test-name-value", nil)
	require.NoError(t, err)
	_, err = client.WithQuery(context.Background(), "us-west", "test-id", nil)
	require.NoError(t, err)
}

func TestMixedParamsClient_WithBody(t *testing.T) {
	client, err := clientinitgroup.NewMixedParamsClientWithNoCredential("http://localhost:3000", "test-name-value", nil)
	require.NoError(t, err)
	_, err = client.WithBody(context.Background(), "us-west", clientinitgroup.WithBodyRequest{
		Name: to.Ptr("test-name"),
	}, nil)
	require.NoError(t, err)
}
