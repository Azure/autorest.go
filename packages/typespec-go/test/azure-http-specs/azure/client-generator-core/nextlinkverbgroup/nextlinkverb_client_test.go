// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nextlinkverbgroup_test

import (
	"context"
	"nextlinkverbgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestNewListItemsPager(t *testing.T) {
	client, err := nextlinkverbgroup.NewNextLinkVerbClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewListItemsPager(nil)
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		pageCount++
		switch pageCount {
		case 1:
			require.Len(t, page.Items, 1)
			require.Equal(t, nextlinkverbgroup.Test{ID: to.Ptr("test1")}, page.Items[0])
			require.NotNil(t, page.NextLink)
		case 2:
			require.Len(t, page.Items, 1)
			require.Equal(t, nextlinkverbgroup.Test{ID: to.Ptr("test2")}, page.Items[0])
			require.Nil(t, page.NextLink)
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 2, pageCount)
}
