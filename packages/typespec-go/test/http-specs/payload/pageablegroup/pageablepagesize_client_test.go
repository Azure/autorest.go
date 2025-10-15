// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package pageablegroup_test

import (
	"context"
	"pageablegroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestPageablePageSizeClientNewListWithoutContinuationPager(t *testing.T) {
	pageableClient, err := pageablegroup.NewPageableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pageableSizeClient := pageableClient.NewPageablePageSizeClient()
	pager := pageableSizeClient.NewListWithoutContinuationPager(nil)
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.Pets, 4)
		pageCount++
		require.Equal(t, []*pageablegroup.Pet{
			{
				ID:   to.Ptr("1"),
				Name: to.Ptr("dog"),
			},
			{
				ID:   to.Ptr("2"),
				Name: to.Ptr("cat"),
			},
			{
				ID:   to.Ptr("3"),
				Name: to.Ptr("bird"),
			},
			{
				ID:   to.Ptr("4"),
				Name: to.Ptr("fish"),
			},
		}, page.Pets)
	}
	require.EqualValues(t, 1, pageCount)
}

func TestPageablePageSizeClientNewListWithPageSizePager(t *testing.T) {
	pageableClient, err := pageablegroup.NewPageableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pageableSizeClient := pageableClient.NewPageablePageSizeClient()

	// Test with pageSize=2
	pager := pageableSizeClient.NewListWithPageSizePager(&pageablegroup.PageablePageSizeClientListWithPageSizeOptions{
		PageSize: to.Ptr[int32](2),
	})
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.Pets, 2)
		pageCount++
		require.Equal(t, []*pageablegroup.Pet{
			{
				ID:   to.Ptr("1"),
				Name: to.Ptr("dog"),
			},
			{
				ID:   to.Ptr("2"),
				Name: to.Ptr("cat"),
			},
		}, page.Pets)
	}
	require.EqualValues(t, 1, pageCount)

	// Test with pageSize=4
	pager = pageableSizeClient.NewListWithPageSizePager(&pageablegroup.PageablePageSizeClientListWithPageSizeOptions{
		PageSize: to.Ptr[int32](4),
	})
	pageCount = 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.Pets, 4)
		pageCount++
		require.Equal(t, []*pageablegroup.Pet{
			{
				ID:   to.Ptr("1"),
				Name: to.Ptr("dog"),
			},
			{
				ID:   to.Ptr("2"),
				Name: to.Ptr("cat"),
			},
			{
				ID:   to.Ptr("3"),
				Name: to.Ptr("bird"),
			},
			{
				ID:   to.Ptr("4"),
				Name: to.Ptr("fish"),
			},
		}, page.Pets)
	}
	require.EqualValues(t, 1, pageCount)
}
