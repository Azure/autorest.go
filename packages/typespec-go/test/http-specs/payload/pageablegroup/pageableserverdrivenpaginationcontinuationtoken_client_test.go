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

func TestNewRequestHeaderNestedResponseBodyPager(t *testing.T) {
	var token string
	client, err := pageablegroup.NewPageableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestHeaderNestedResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderNestedResponseBodyOptions{Bar: to.Ptr("bar"), Foo: to.Ptr("foo")})
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.NestedItems.Pets, 2)
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
		}, page.NestedItems.Pets)

		token = *page.NestedNext.NextToken
	}
	require.EqualValues(t, 1, pageCount)

	pagerToken := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestHeaderNestedResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderNestedResponseBodyOptions{Bar: to.Ptr("bar"), Foo: to.Ptr("foo"), Token: to.Ptr(token)})
	pageTokenCount := 0
	for pagerToken.More() {
		page, err := pagerToken.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.NestedItems.Pets, 2)
		pageTokenCount++
		require.Equal(t, []*pageablegroup.Pet{
			{
				ID:   to.Ptr("3"),
				Name: to.Ptr("bird"),
			},
			{
				ID:   to.Ptr("4"),
				Name: to.Ptr("fish"),
			},
		}, page.NestedItems.Pets)
		require.Nil(t, page.NestedNext)
	}
	require.EqualValues(t, 1, pageTokenCount)
}

func TestNewRequestHeaderResponseBodyPager(t *testing.T) {
	var token string
	client, err := pageablegroup.NewPageableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestHeaderResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderResponseBodyOptions{Bar: to.Ptr("bar"), Foo: to.Ptr("foo")})
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

		token = *page.NextToken
	}
	require.EqualValues(t, 1, pageCount)

	pagerToken := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestHeaderResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderResponseBodyOptions{Bar: to.Ptr("bar"), Foo: to.Ptr("foo"), Token: to.Ptr(token)})
	pageTokenCount := 0
	for pagerToken.More() {
		page, err := pagerToken.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.Pets, 2)
		pageTokenCount++
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
		require.Nil(t, page.NextToken)
	}
	require.EqualValues(t, 1, pageTokenCount)
}

func TestNewRequestHeaderResponseHeaderPager(t *testing.T) {
	var token string
	client, err := pageablegroup.NewPageableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestHeaderResponseHeaderPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderResponseHeaderOptions{Bar: to.Ptr("bar"), Foo: to.Ptr("foo")})
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

		token = *page.NextToken
	}
	require.EqualValues(t, 1, pageCount)

	pagerToken := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestHeaderResponseHeaderPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderResponseHeaderOptions{Bar: to.Ptr("bar"), Foo: to.Ptr("foo"), Token: to.Ptr(token)})
	pageTokenCount := 0
	for pagerToken.More() {
		page, err := pagerToken.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.Pets, 2)
		pageTokenCount++
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
		require.Nil(t, page.NextToken)
	}
	require.EqualValues(t, 1, pageTokenCount)
}

func TestNewRequestQueryNestedResponseBodyPager(t *testing.T) {
	var token string
	client, err := pageablegroup.NewPageableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestQueryNestedResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestQueryNestedResponseBodyOptions{Bar: to.Ptr("bar"), Foo: to.Ptr("foo")})
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.NestedItems.Pets, 2)
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
		}, page.NestedItems.Pets)

		token = *page.NestedNext.NextToken
	}
	require.EqualValues(t, 1, pageCount)

	pagerToken := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestQueryNestedResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestQueryNestedResponseBodyOptions{Bar: to.Ptr("bar"), Foo: to.Ptr("foo"), Token: to.Ptr(token)})
	pageTokenCount := 0
	for pagerToken.More() {
		page, err := pagerToken.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.NestedItems.Pets, 2)
		pageTokenCount++
		require.Equal(t, []*pageablegroup.Pet{
			{
				ID:   to.Ptr("3"),
				Name: to.Ptr("bird"),
			},
			{
				ID:   to.Ptr("4"),
				Name: to.Ptr("fish"),
			},
		}, page.NestedItems.Pets)
		require.Nil(t, page.NestedNext)
	}
	require.EqualValues(t, 1, pageTokenCount)
}

func TestNewRequestQueryResponseBodyPager(t *testing.T) {
	var token string
	client, err := pageablegroup.NewPageableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestQueryResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestQueryResponseBodyOptions{Bar: to.Ptr("bar"), Foo: to.Ptr("foo")})
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

		token = *page.NextToken
	}
	require.EqualValues(t, 1, pageCount)

	pagerToken := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestQueryResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestQueryResponseBodyOptions{Bar: to.Ptr("bar"), Foo: to.Ptr("foo"), Token: to.Ptr(token)})
	pageTokenCount := 0
	for pagerToken.More() {
		page, err := pagerToken.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.Pets, 2)
		pageTokenCount++
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
		require.Nil(t, page.NextToken)
	}
	require.EqualValues(t, 1, pageTokenCount)
}

func TestNewRequestQueryResponseHeaderPager(t *testing.T) {
	var token string
	client, err := pageablegroup.NewPageableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestQueryResponseHeaderPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestQueryResponseHeaderOptions{Bar: to.Ptr("bar"), Foo: to.Ptr("foo")})
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

		token = *page.NextToken
	}
	require.EqualValues(t, 1, pageCount)

	pagerToken := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestQueryResponseHeaderPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestQueryResponseHeaderOptions{Bar: to.Ptr("bar"), Foo: to.Ptr("foo"), Token: to.Ptr(token)})
	pageTokenCount := 0
	for pagerToken.More() {
		page, err := pagerToken.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.Pets, 2)
		pageTokenCount++
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
		require.Nil(t, page.NextToken)
	}
	require.EqualValues(t, 1, pageTokenCount)
}
