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

func TestOptionalCollectionsModelClient_GetAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalCollectionsModelClient().GetAll(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, []optionalitygroup.StringProperty{
		{
			Property: to.Ptr("hello"),
		},
		{
			Property: to.Ptr("world"),
		},
	}, resp.Property)
}

func TestOptionalCollectionsModelClient_GetDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalCollectionsModelClient().GetDefault(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalCollectionsModelClient_PutAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalCollectionsModelClient().PutAll(context.Background(), optionalitygroup.CollectionsModelProperty{
		Property: []optionalitygroup.StringProperty{
			{
				Property: to.Ptr("hello"),
			},
			{
				Property: to.Ptr("world"),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalCollectionsModelClient_PutDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalCollectionsModelClient().PutDefault(context.Background(), optionalitygroup.CollectionsModelProperty{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
