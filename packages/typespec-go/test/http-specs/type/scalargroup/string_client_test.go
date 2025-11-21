// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package scalargroup_test

import (
	"context"
	"scalargroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringClient_Get(t *testing.T) {
	client, err := scalargroup.NewScalarClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewScalarStringClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, "test", *resp.Value)
}

func TestStringClient_Put(t *testing.T) {
	client, err := scalargroup.NewScalarClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewScalarStringClient().Put(context.Background(), "test", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
