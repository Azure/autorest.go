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

func TestPropertyClient_Client(t *testing.T) {
	client, err := projectednamegroup.NewProjectedNameClient(nil)
	require.NoError(t, err)
	resp, err := client.NewPropertyClient().Client(context.Background(), projectednamegroup.ClientProjectedNameModel{
		ClientName: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestPropertyClient_JSON(t *testing.T) {
	client, err := projectednamegroup.NewProjectedNameClient(nil)
	require.NoError(t, err)
	resp, err := client.NewPropertyClient().JSON(context.Background(), projectednamegroup.JSONProjectedNameModel{
		DefaultName: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestPropertyClient_JSONAndClient(t *testing.T) {
	client, err := projectednamegroup.NewProjectedNameClient(nil)
	require.NoError(t, err)
	resp, err := client.NewPropertyClient().JSONAndClient(context.Background(), projectednamegroup.JSONAndClientProjectedNameModel{
		ClientName: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestPropertyClient_Language(t *testing.T) {
	client, err := projectednamegroup.NewProjectedNameClient(nil)
	require.NoError(t, err)
	resp, err := client.NewPropertyClient().Language(context.Background(), projectednamegroup.LanguageProjectedNameModel{
		GoName: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
