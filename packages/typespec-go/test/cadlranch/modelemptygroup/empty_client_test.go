//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package modelemptygroup_test

import (
	"context"
	"modelemptygroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyClientGetEmpty(t *testing.T) {
	client, err := modelemptygroup.NewEmptyClient(nil)
	require.NoError(t, err)
	resp, err := client.GetEmpty(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestEmptyClientPostRoundTripEmpty(t *testing.T) {
	client, err := modelemptygroup.NewEmptyClient(nil)
	require.NoError(t, err)
	resp, err := client.PostRoundTripEmpty(context.Background(), modelemptygroup.EmptyInputOutput{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestEmptyClientPutEmpty(t *testing.T) {
	client, err := modelemptygroup.NewEmptyClient(nil)
	require.NoError(t, err)
	resp, err := client.PutEmpty(context.Background(), modelemptygroup.EmptyInput{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
