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

func TestOptionalUnionFloatLiteralClient_GetAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalUnionFloatLiteralClient().GetAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, optionalitygroup.UnionFloatLiteralPropertyProperty2375, *resp.Property)
}

func TestOptionalUnionFloatLiteralClient_GetDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalUnionFloatLiteralClient().GetDefault(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalUnionFloatLiteralClient_PutAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalUnionFloatLiteralClient().PutAll(context.Background(), optionalitygroup.UnionFloatLiteralProperty{
		Property: to.Ptr(optionalitygroup.UnionFloatLiteralPropertyProperty2375),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalUnionFloatLiteralClient_PutDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalUnionFloatLiteralClient().PutDefault(context.Background(), optionalitygroup.UnionFloatLiteralProperty{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
