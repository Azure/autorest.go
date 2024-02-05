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
	client, err := collectionfmtgroup.NewQueryClient(nil)
	require.NoError(t, err)
	resp, err := client.CSV(context.Background(), []string{"blue", "red", "green"}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClient_Multi(t *testing.T) {
	client, err := collectionfmtgroup.NewQueryClient(nil)
	require.NoError(t, err)
	resp, err := client.Multi(context.Background(), []string{"blue", "red", "green"}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClient_Pipes(t *testing.T) {
	client, err := collectionfmtgroup.NewQueryClient(nil)
	require.NoError(t, err)
	resp, err := client.Pipes(context.Background(), []string{"blue", "red", "green"}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClient_Ssv(t *testing.T) {
	client, err := collectionfmtgroup.NewQueryClient(nil)
	require.NoError(t, err)
	resp, err := client.Ssv(context.Background(), []string{"blue", "red", "green"}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClient_Tsv(t *testing.T) {
	client, err := collectionfmtgroup.NewQueryClient(nil)
	require.NoError(t, err)
	resp, err := client.Tsv(context.Background(), []string{"blue", "red", "green"}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
