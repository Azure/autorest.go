//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package projectednamegroup_test

import (
	"context"
	"projectednamegroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestModelClient_Client(t *testing.T) {
	client, err := projectednamegroup.NewProjectedNameClient(nil)
	require.NoError(t, err)
	resp, err := client.NewModelClient().Client(context.Background(), projectednamegroup.ClientModel{
		DefaultName: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelClient_Language(t *testing.T) {
	client, err := projectednamegroup.NewProjectedNameClient(nil)
	require.NoError(t, err)
	resp, err := client.NewModelClient().Language(context.Background(), projectednamegroup.GoModel{
		DefaultName: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
