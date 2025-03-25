//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurepagegroup_test

import (
	"azurepagegroup"
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestPageClient_NewListWithCustomPageModelPager(t *testing.T) {
	client, err := azurepagegroup.NewPageClient(nil)
	require.NoError(t, err)
	pager := client.NewListWithCustomPageModelPager(nil)
	pages := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.EqualValues(t, []*azurepagegroup.User{
			{
				ID:   to.Ptr[int32](1),
				Name: to.Ptr("Madge"),
				Etag: to.Ptr[azcore.ETag]("11bdc430-65e8-45ad-81d9-8ffa60d55b59"),
			},
		}, page.Items)
		pages++
	}
	require.EqualValues(t, 1, pages)
}

func TestPageClient_NewListWithPagePager(t *testing.T) {
	client, err := azurepagegroup.NewPageClient(nil)
	require.NoError(t, err)
	pager := client.NewListWithPagePager(nil)
	pages := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.EqualValues(t, []*azurepagegroup.User{
			{
				ID:   to.Ptr[int32](1),
				Name: to.Ptr("Madge"),
				Etag: to.Ptr[azcore.ETag]("11bdc430-65e8-45ad-81d9-8ffa60d55b59"),
			},
		}, page.Value)
		pages++
	}
	require.EqualValues(t, 1, pages)
}

func TestPageClient_NewListWithParametersPager(t *testing.T) {
	client, err := azurepagegroup.NewPageClient(nil)
	require.NoError(t, err)
	pager := client.NewListWithParametersPager(azurepagegroup.ListItemInputBody{
		InputName: to.Ptr("Madge"),
	}, &azurepagegroup.PageClientListWithParametersOptions{
		Another: to.Ptr(azurepagegroup.ListItemInputExtensibleEnumSecond),
	})
	pages := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.EqualValues(t, []*azurepagegroup.User{
			{
				ID:   to.Ptr[int32](1),
				Name: to.Ptr("Madge"),
				Etag: to.Ptr[azcore.ETag]("11bdc430-65e8-45ad-81d9-8ffa60d55b59"),
			},
		}, page.Value)
		pages++
	}
	require.EqualValues(t, 1, pages)
}
