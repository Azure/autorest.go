// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package scalargroup_test

import (
	"context"
	"scalargroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnknownClient_Get(t *testing.T) {
	client, err := scalargroup.NewScalarClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewScalarUnknownClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, "test", resp.Interface)
}

func TestUnknownClient_Put(t *testing.T) {
	client, err := scalargroup.NewScalarClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewScalarUnknownClient().Put(context.Background(), "test", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
