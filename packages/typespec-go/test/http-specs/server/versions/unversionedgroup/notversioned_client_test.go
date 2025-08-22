// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package unversionedgroup_test

import (
	"context"
	"testing"
	"unversionedgroup"

	"github.com/stretchr/testify/require"
)

func TestNotVersionedClient_WithPathAPIVersion(t *testing.T) {
	client, err := unversionedgroup.NewNotVersionedClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.WithPathAPIVersion(context.Background(), "v1.0", nil)
	require.NoError(t, err)
	require.True(t, resp.Success)
}

func TestNotVersionedClient_WithQueryAPIVersion(t *testing.T) {
	client, err := unversionedgroup.NewNotVersionedClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.WithQueryAPIVersion(context.Background(), "v1.0", nil)
	require.NoError(t, err)
	require.True(t, resp.Success)
}

func TestNotVersionedClient_WithoutAPIVersion(t *testing.T) {
	client, err := unversionedgroup.NewNotVersionedClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.WithoutAPIVersion(context.Background(), nil)
	require.NoError(t, err)
	require.True(t, resp.Success)
}
