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

func TestNewLinkPager(t *testing.T) {
	client, err := pageablegroup.NewPageableServerDrivenPaginationClient(nil)
	require.NoError(t, err)
	pager := client.NewLinkPager(nil)
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.Pets, 2)
		pageCount++
		switch pageCount {
		case 1:
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
		case 2:
			require.Equal(t, []*pageablegroup.Pet{
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
	}
	require.EqualValues(t, 2, pageCount)
}
