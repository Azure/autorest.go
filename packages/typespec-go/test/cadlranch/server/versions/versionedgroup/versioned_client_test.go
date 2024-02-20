// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package versionedgroup_test

import (
	"context"
	"testing"
	"versionedgroup"

	"github.com/stretchr/testify/require"
)

func TestVersionedClient_WithPathAPIVersion(t *testing.T) {
	client, err := versionedgroup.NewVersionedClient(nil)
	require.NoError(t, err)
	resp, err := client.WithPathAPIVersion(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestVersionedClient_WithQueryAPIVersion(t *testing.T) {
	client, err := versionedgroup.NewVersionedClient(nil)
	require.NoError(t, err)
	resp, err := client.WithQueryAPIVersion(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestVersionedClient_WithoutAPIVersion(t *testing.T) {
	client, err := versionedgroup.NewVersionedClient(nil)
	require.NoError(t, err)
	resp, err := client.WithoutAPIVersion(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
