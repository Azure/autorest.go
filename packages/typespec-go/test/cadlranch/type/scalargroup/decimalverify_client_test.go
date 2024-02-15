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

func TestDecimalVerifyClient_PrepareVerify(t *testing.T) {
	client, err := scalargroup.NewScalarClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDecimalVerifyClient().PrepareVerify(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, []float64{0.1, 0.1, 0.1}, resp.Value)
}

func TestDecimalVerifyClient_Verify(t *testing.T) {
	client, err := scalargroup.NewScalarClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDecimalVerifyClient().Verify(context.Background(), 0.3, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
