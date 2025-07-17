// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package optionalitygroup_test

import (
	"context"
	"optionalitygroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestOptionalUnionIntLiteralClient_GetAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalUnionIntLiteralClient().GetAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, optionalitygroup.UnionIntLiteralPropertyProperty2, *resp.Property)
}

func TestOptionalUnionIntLiteralClient_GetDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalUnionIntLiteralClient().GetDefault(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalUnionIntLiteralClient_PutAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalUnionIntLiteralClient().PutAll(context.Background(), optionalitygroup.UnionIntLiteralProperty{
		Property: to.Ptr(optionalitygroup.UnionIntLiteralPropertyProperty2),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalUnionIntLiteralClient_PutDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalUnionIntLiteralClient().PutDefault(context.Background(), optionalitygroup.UnionIntLiteralProperty{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
