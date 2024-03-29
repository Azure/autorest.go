//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package scalargroup_test

import (
	"context"
	"scalargroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBooleanClient_Get(t *testing.T) {
	client, err := scalargroup.NewScalarClient(nil)
	require.NoError(t, err)
	resp, err := client.NewScalarBooleanClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, *resp.Value, true)
}

func TestBooleanClient_Put(t *testing.T) {
	client, err := scalargroup.NewScalarClient(nil)
	require.NoError(t, err)
	resp, err := client.NewScalarBooleanClient().Put(context.Background(), true, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
