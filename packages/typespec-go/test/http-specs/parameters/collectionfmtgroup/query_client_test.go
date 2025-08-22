//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package collectionfmtgroup_test

import (
	"collectionfmtgroup"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryClient_CSV(t *testing.T) {
	client, err := collectionfmtgroup.NewCollectionFormatClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewCollectionFormatQueryClient().CSV(context.Background(), []string{"blue", "red", "green"}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClient_Multi(t *testing.T) {
	client, err := collectionfmtgroup.NewCollectionFormatClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewCollectionFormatQueryClient().Multi(context.Background(), []string{"blue", "red", "green"}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClient_Pipes(t *testing.T) {
	client, err := collectionfmtgroup.NewCollectionFormatClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewCollectionFormatQueryClient().Pipes(context.Background(), []string{"blue", "red", "green"}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClient_Ssv(t *testing.T) {
	client, err := collectionfmtgroup.NewCollectionFormatClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewCollectionFormatQueryClient().Ssv(context.Background(), []string{"blue", "red", "green"}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
