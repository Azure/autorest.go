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

func TestGroup1Client_Four(t *testing.T) {
	client, err := twoopgroup.NewTwoOperationGroupClientWithNoCredential("http://localhost:3000", twoopgroup.ClientTypeTwoOperationGroup, nil)
	require.NoError(t, err)
	resp, err := client.NewTwoOperationGroupGroup1Client().Four(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestGroup1Client_One(t *testing.T) {
	client, err := twoopgroup.NewTwoOperationGroupClientWithNoCredential("http://localhost:3000", twoopgroup.ClientTypeTwoOperationGroup, nil)
	require.NoError(t, err)
	resp, err := client.NewTwoOperationGroupGroup1Client().One(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestGroup1Client_Three(t *testing.T) {
	client, err := twoopgroup.NewTwoOperationGroupClientWithNoCredential("http://localhost:3000", twoopgroup.ClientTypeTwoOperationGroup, nil)
	require.NoError(t, err)
	resp, err := client.NewTwoOperationGroupGroup1Client().Three(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
