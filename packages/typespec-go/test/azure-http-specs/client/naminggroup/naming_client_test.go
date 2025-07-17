// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package naminggroup_test

import (
	"context"
	"naminggroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestNamingClient_Client(t *testing.T) {
	client, err := naminggroup.NewNamingClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Client(context.Background(), naminggroup.ClientNameModel{
		to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestNamingClient_ClientName(t *testing.T) {
	client, err := naminggroup.NewNamingClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.ClientName(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestNamingClient_CompatibleWithEncodedName(t *testing.T) {
	client, err := naminggroup.NewNamingClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.CompatibleWithEncodedName(context.Background(), naminggroup.ClientNameAndJSONEncodedNameModel{
		ClientName: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestNamingClient_Language(t *testing.T) {
	client, err := naminggroup.NewNamingClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Language(context.Background(), naminggroup.LanguageClientNameModel{
		GoName: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestNamingClient_Parameter(t *testing.T) {
	client, err := naminggroup.NewNamingClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Parameter(context.Background(), "true", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestNamingClient_Request(t *testing.T) {
	client, err := naminggroup.NewNamingClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Request(context.Background(), "true", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestNamingClient_Response(t *testing.T) {
	client, err := naminggroup.NewNamingClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Response(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.DefaultName)
	require.EqualValues(t, "true", *resp.DefaultName)
}

func TestUnionEnumClient_UnionEnumMemberName(t *testing.T) {
	client, err := naminggroup.NewNamingClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewNamingUnionEnumClient().UnionEnumMemberName(context.Background(), naminggroup.ExtensibleEnumClientEnumValue1, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestUnionEnumClient_UnionEnumName(t *testing.T) {
	client, err := naminggroup.NewNamingClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewNamingUnionEnumClient().UnionEnumName(context.Background(), naminggroup.ClientExtensibleEnumEnumValue1, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
