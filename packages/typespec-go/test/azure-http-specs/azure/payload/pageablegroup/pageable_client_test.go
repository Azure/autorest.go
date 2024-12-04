//go:build go1.18
// +build go1.18

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

func TestPageableClientNewListPager(t *testing.T) {
	client, err := pageablegroup.NewPageableClient(nil)
	require.NoError(t, err)
	pager := client.NewListPager(&pageablegroup.PageableClientListOptions{
		Maxpagesize: to.Ptr[int32](3),
	})
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		pageCount++
		switch pageCount {
		case 1:
			require.Len(t, page.Value, 3)
			require.NotNil(t, page.NextLink)
		case 2:
			require.Len(t, page.Value, 1)
			require.Nil(t, page.NextLink)
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 2, pageCount)
}
