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

func TestOptionalStringLiteralClient_GetAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalStringLiteralClient().GetAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, "hello", *resp.Property)
}

func TestOptionalStringLiteralClient_GetDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalStringLiteralClient().GetDefault(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalStringLiteralClient_PutAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalStringLiteralClient().PutAll(context.Background(), optionalitygroup.StringLiteralProperty{
		Property: to.Ptr("hello"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalStringLiteralClient_PutDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalStringLiteralClient().PutDefault(context.Background(), optionalitygroup.StringLiteralProperty{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
