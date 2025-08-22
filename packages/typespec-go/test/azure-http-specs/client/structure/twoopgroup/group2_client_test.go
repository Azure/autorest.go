//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package twoopgroup_test

import (
	"context"
	"testing"
	"twoopgroup"

	"github.com/stretchr/testify/require"
)

func TestGroup2Client_Five(t *testing.T) {
	client, err := twoopgroup.NewTwoOperationGroupClientWithNoCredential("http://localhost:3000", twoopgroup.ClientTypeTwoOperationGroup, nil)
	require.NoError(t, err)
	resp, err := client.NewTwoOperationGroupGroup2Client().Five(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestGroup2Client_Six(t *testing.T) {
	client, err := twoopgroup.NewTwoOperationGroupClientWithNoCredential("http://localhost:3000", twoopgroup.ClientTypeTwoOperationGroup, nil)
	require.NoError(t, err)
	resp, err := client.NewTwoOperationGroupGroup2Client().Six(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestGroup2Client_Two(t *testing.T) {
	client, err := twoopgroup.NewTwoOperationGroupClientWithNoCredential("http://localhost:3000", twoopgroup.ClientTypeTwoOperationGroup, nil)
	require.NoError(t, err)
	resp, err := client.NewTwoOperationGroupGroup2Client().Two(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
