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

func TestOptionalUnionStringLiteralClient_GetAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalUnionStringLiteralClient().GetAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, optionalitygroup.UnionStringLiteralPropertyPropertyWorld, *resp.Property)
}

func TestOptionalUnionStringLiteralClient_GetDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalUnionStringLiteralClient().GetDefault(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalUnionStringLiteralClient_PutAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalUnionStringLiteralClient().PutAll(context.Background(), optionalitygroup.UnionStringLiteralProperty{
		Property: to.Ptr(optionalitygroup.UnionStringLiteralPropertyPropertyWorld),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalUnionStringLiteralClient_PutDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalUnionStringLiteralClient().PutDefault(context.Background(), optionalitygroup.UnionStringLiteralProperty{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
