// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package pageablegroup_test

import (
	"context"
	"net/http"
	"pageablegroup"
	"pageablegroup/fake"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestNewRequestHeaderNestedResponseBodyPager(t *testing.T) {
	client, err := pageablegroup.NewPageableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestHeaderNestedResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderNestedResponseBodyOptions{
		Bar: to.Ptr("bar"),
		Foo: to.Ptr("foo"),
	})
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
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
			}, page.NestedItems.Pets)
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
			}, page.NestedItems.Pets)
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 2, pageCount)

	pager = client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestHeaderNestedResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderNestedResponseBodyOptions{
		Bar:   to.Ptr("bar"),
		Foo:   to.Ptr("foo"),
		Token: to.Ptr("page2"),
	})
	pageCount = 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		pageCount++
		switch pageCount {
		case 1:
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
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 1, pageCount)
}

func TestNewRequestHeaderResponseBodyPager(t *testing.T) {
	client, err := pageablegroup.NewPageableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestHeaderResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderResponseBodyOptions{
		Bar: to.Ptr("bar"),
		Foo: to.Ptr("foo"),
	})
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
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
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 2, pageCount)

	pager = client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestHeaderResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderResponseBodyOptions{
		Bar:   to.Ptr("bar"),
		Foo:   to.Ptr("foo"),
		Token: to.Ptr("page2"),
	})
	pageCount = 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		pageCount++
		switch pageCount {
		case 1:
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
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 1, pageCount)
}

func TestNewRequestHeaderResponseHeaderPager(t *testing.T) {
	client, err := pageablegroup.NewPageableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestHeaderResponseHeaderPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderResponseHeaderOptions{
		Bar: to.Ptr("bar"),
		Foo: to.Ptr("foo"),
	})
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
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
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 2, pageCount)

	pager = client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestHeaderResponseHeaderPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderResponseHeaderOptions{
		Bar:   to.Ptr("bar"),
		Foo:   to.Ptr("foo"),
		Token: to.Ptr("page2"),
	})
	pageCount = 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		pageCount++
		switch pageCount {
		case 1:
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
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 1, pageCount)
}

func TestNewRequestQueryNestedResponseBodyPager(t *testing.T) {
	client, err := pageablegroup.NewPageableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestQueryNestedResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestQueryNestedResponseBodyOptions{
		Bar: to.Ptr("bar"),
		Foo: to.Ptr("foo"),
	})
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
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
			}, page.NestedItems.Pets)
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
			}, page.NestedItems.Pets)
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 2, pageCount)

	pager = client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestQueryNestedResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestQueryNestedResponseBodyOptions{
		Bar:   to.Ptr("bar"),
		Foo:   to.Ptr("foo"),
		Token: to.Ptr("page2"),
	})
	pageCount = 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		pageCount++
		switch pageCount {
		case 1:
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
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 1, pageCount)
}

func TestNewRequestQueryResponseBodyPager(t *testing.T) {
	client, err := pageablegroup.NewPageableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestQueryResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestQueryResponseBodyOptions{
		Bar: to.Ptr("bar"),
		Foo: to.Ptr("foo"),
	})
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
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
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 2, pageCount)

	pager = client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestQueryResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestQueryResponseBodyOptions{
		Bar:   to.Ptr("bar"),
		Foo:   to.Ptr("foo"),
		Token: to.Ptr("page2"),
	})
	pageCount = 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		pageCount++
		switch pageCount {
		case 1:
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
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 1, pageCount)
}

func TestNewRequestQueryResponseHeaderPager(t *testing.T) {
	client, err := pageablegroup.NewPageableClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestQueryResponseHeaderPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestQueryResponseHeaderOptions{
		Bar: to.Ptr("bar"),
		Foo: to.Ptr("foo"),
	})
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
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
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 2, pageCount)

	pager = client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestQueryResponseHeaderPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestQueryResponseHeaderOptions{
		Bar:   to.Ptr("bar"),
		Foo:   to.Ptr("foo"),
		Token: to.Ptr("page2"),
	})
	pageCount = 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		pageCount++
		switch pageCount {
		case 1:
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
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 1, pageCount)
}

func TestNewRequestHeaderNestedResponseBodyPager_Fake(t *testing.T) {
	srv := fake.PageableServerDrivenPaginationContinuationTokenServer{
		NewRequestHeaderNestedResponseBodyPager: func(options *pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderNestedResponseBodyOptions) azfake.PagerResponder[pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderNestedResponseBodyResponse] {
			require.NotNil(t, options)
			require.NotNil(t, options.Bar)
			require.Equal(t, "bar", *options.Bar)
			require.NotNil(t, options.Foo)
			require.Equal(t, "foo", *options.Foo)
			pager := azfake.PagerResponder[pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderNestedResponseBodyResponse]{}
			pager.AddPage(http.StatusOK, pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderNestedResponseBodyResponse{
				RequestHeaderNestedResponseBodyResponse: pageablegroup.RequestHeaderNestedResponseBodyResponse{
					NestedItems: &pageablegroup.RequestHeaderNestedResponseBodyResponseNestedItems{
						Pets: []*pageablegroup.Pet{
							{ID: to.Ptr("1"), Name: to.Ptr("dog")},
							{ID: to.Ptr("2"), Name: to.Ptr("cat")},
						},
					},
					NestedNext: &pageablegroup.RequestHeaderNestedResponseBodyResponseNestedNext{
						NextToken: to.Ptr("page2"),
					},
				},
			}, nil)
			pager.AddPage(http.StatusOK, pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderNestedResponseBodyResponse{
				RequestHeaderNestedResponseBodyResponse: pageablegroup.RequestHeaderNestedResponseBodyResponse{
					NestedItems: &pageablegroup.RequestHeaderNestedResponseBodyResponseNestedItems{
						Pets: []*pageablegroup.Pet{
							{ID: to.Ptr("3"), Name: to.Ptr("bird")},
							{ID: to.Ptr("4"), Name: to.Ptr("fish")},
						},
					},
				},
			}, nil)
			return pager
		},
	}

	client, err := pageablegroup.NewPageableClientWithNoCredential("https://fake.endpoint", &pageablegroup.PageableClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewPageableServerDrivenPaginationContinuationTokenServerTransport(&srv),
		},
	})
	require.NoError(t, err)

	pager := client.NewPageableServerDrivenPaginationClient().NewPageableServerDrivenPaginationContinuationTokenClient().NewRequestHeaderNestedResponseBodyPager(&pageablegroup.PageableServerDrivenPaginationContinuationTokenClientRequestHeaderNestedResponseBodyOptions{
		Bar: to.Ptr("bar"),
		Foo: to.Ptr("foo"),
	})
	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		pageCount++
		switch pageCount {
		case 1:
			require.Equal(t, []*pageablegroup.Pet{
				{ID: to.Ptr("1"), Name: to.Ptr("dog")},
				{ID: to.Ptr("2"), Name: to.Ptr("cat")},
			}, page.NestedItems.Pets)
		case 2:
			require.Equal(t, []*pageablegroup.Pet{
				{ID: to.Ptr("3"), Name: to.Ptr("bird")},
				{ID: to.Ptr("4"), Name: to.Ptr("fish")},
			}, page.NestedItems.Pets)
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 2, pageCount)
}
