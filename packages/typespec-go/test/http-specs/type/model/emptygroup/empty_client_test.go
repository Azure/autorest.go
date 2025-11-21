// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package emptygroup_test

import (
	"context"
	"emptygroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyClientGetEmpty(t *testing.T) {
	client, err := emptygroup.NewEmptyClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.GetEmpty(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestEmptyClientPostRoundTripEmpty(t *testing.T) {
	client, err := emptygroup.NewEmptyClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.PostRoundTripEmpty(context.Background(), emptygroup.EmptyInputOutput{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestEmptyClientPutEmpty(t *testing.T) {
	client, err := emptygroup.NewEmptyClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.PutEmpty(context.Background(), emptygroup.EmptyInput{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
