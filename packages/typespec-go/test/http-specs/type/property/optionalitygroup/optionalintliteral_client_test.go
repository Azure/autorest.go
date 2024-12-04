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

func TestOptionalIntLiteralClient_GetAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalIntLiteralClient().GetAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, int32(1), *resp.Property)
}

func TestOptionalIntLiteralClient_GetDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalIntLiteralClient().GetDefault(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalIntLiteralClient_PutAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalIntLiteralClient().PutAll(context.Background(), optionalitygroup.IntLiteralProperty{
		Property: to.Ptr[int32](1),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalIntLiteralClient_PutDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalIntLiteralClient().PutDefault(context.Background(), optionalitygroup.IntLiteralProperty{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
