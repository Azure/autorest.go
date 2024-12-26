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

func TestOptionalStringClient_GetAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalStringClient().GetAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, "hello", *resp.Property)
}

func TestOptionalStringClient_GetDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalStringClient().GetDefault(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalStringClient_PutAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalStringClient().PutAll(context.Background(), optionalitygroup.StringProperty{
		Property: to.Ptr("hello"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalStringClient_PutDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalStringClient().PutDefault(context.Background(), optionalitygroup.StringProperty{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
