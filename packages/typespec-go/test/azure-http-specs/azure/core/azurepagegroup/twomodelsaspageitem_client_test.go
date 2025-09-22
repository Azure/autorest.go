//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurepagegroup_test

import (
	"azurepagegroup"
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestTwoModelsAsPageItemClient_NewListFirstItemPager(t *testing.T) {
	client, err := azurepagegroup.NewPageClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewPageTwoModelsAsPageItemClient().NewListFirstItemPager(nil)
	pages := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.EqualValues(t, []*azurepagegroup.FirstItem{
			{
				ID: to.Ptr[int32](1),
			},
		}, page.Value)
		pages++
	}
	require.EqualValues(t, 1, pages)
}

func TestTwoModelsAsPageItemClient_NewListSecondItemPager(t *testing.T) {
	client, err := azurepagegroup.NewPageClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewPageTwoModelsAsPageItemClient().NewListSecondItemPager(nil)
	pages := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.EqualValues(t, []*azurepagegroup.SecondItem{
			{
				Name: to.Ptr("Madge"),
			},
		}, page.Value)
		pages++
	}
	require.EqualValues(t, 1, pages)
}
