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

func TestDecimalTypeClient_RequestBody(t *testing.T) {
	client, err := scalargroup.NewScalarClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewScalarDecimalTypeClient().RequestBody(context.Background(), 0.33333, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestDecimalTypeClient_RequestParameter(t *testing.T) {
	client, err := scalargroup.NewScalarClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewScalarDecimalTypeClient().RequestParameter(context.Background(), 0.33333, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestDecimalTypeClient_ResponseBody(t *testing.T) {
	client, err := scalargroup.NewScalarClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewScalarDecimalTypeClient().ResponseBody(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, 0.33333, *resp.Value)
}
