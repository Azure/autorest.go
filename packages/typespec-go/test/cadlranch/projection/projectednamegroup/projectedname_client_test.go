//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package projectednamegroup_test

import (
	"context"
	"projectednamegroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProjectedNameClient_ClientName(t *testing.T) {
	client, err := projectednamegroup.NewProjectedNameClient(nil)
	require.NoError(t, err)
	resp, err := client.ClientName(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestProjectedNameClient_Parameter(t *testing.T) {
	client, err := projectednamegroup.NewProjectedNameClient(nil)
	require.NoError(t, err)
	resp, err := client.Parameter(context.Background(), "true", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
