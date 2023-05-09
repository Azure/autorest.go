// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package paginggroup_test

import (
	"context"
	"generatortests/paginggroup"
	"generatortests/paginggroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFakeNewGetMultiplePagesPager(t *testing.T) {
	page1 := []*paginggroup.Product{
		{
			Properties: &paginggroup.ProductProperties{
				ID: to.Ptr[int32](1),
			},
		},
		{
			Properties: &paginggroup.ProductProperties{
				ID: to.Ptr[int32](2),
			},
		},
	}
	page2 := []*paginggroup.Product{
		{
			Properties: &paginggroup.ProductProperties{
				ID: to.Ptr[int32](3),
			},
		},
		{
			Properties: &paginggroup.ProductProperties{
				ID: to.Ptr[int32](4),
			},
		},
	}
	page3 := []*paginggroup.Product{
		{
			Properties: &paginggroup.ProductProperties{
				ID: to.Ptr[int32](5),
			},
		},
		{
			Properties: &paginggroup.ProductProperties{
				ID: to.Ptr[int32](6),
			},
		},
	}
	server := fake.PagingServer{
		NewGetMultiplePagesPager: func(options *paginggroup.PagingClientGetMultiplePagesOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesResponse]) {
			resp.AddPage(http.StatusOK, paginggroup.PagingClientGetMultiplePagesResponse{
				ProductResult: paginggroup.ProductResult{
					Values: page1,
				},
			}, nil)
			resp.AddPage(http.StatusOK, paginggroup.PagingClientGetMultiplePagesResponse{
				ProductResult: paginggroup.ProductResult{
					Values: page2,
				},
			}, nil)
			resp.AddPage(http.StatusOK, paginggroup.PagingClientGetMultiplePagesResponse{
				ProductResult: paginggroup.ProductResult{
					Values: page3,
				},
			}, nil)
			return
		},
	}
	client, err := paginggroup.NewPagingClient(&azcore.ClientOptions{
		Transport: fake.NewPagingServerTransport(&server),
	})
	require.NoError(t, err)
	pager := client.NewGetMultiplePagesPager(nil)
	pageCount := 0
	for pager.More() {
		pageCount++
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.Len(t, page.Values, 2)
		switch pageCount {
		case 1:
			require.EqualValues(t, page1, page.Values)
		case 2:
			require.EqualValues(t, page2, page.Values)
		case 3:
			require.EqualValues(t, page3, page.Values)
		default:
			t.Fatalf("unexpected page number %d", pageCount)
		}
	}
	require.EqualValues(t, 3, pageCount)
}

func TestFakeNewGetMultiplePagesFailurePager(t *testing.T) {
	page1 := []*paginggroup.Product{
		{
			Properties: &paginggroup.ProductProperties{
				ID: to.Ptr[int32](1),
			},
		},
		{
			Properties: &paginggroup.ProductProperties{
				ID: to.Ptr[int32](2),
			},
		},
	}
	server := fake.PagingServer{
		NewGetMultiplePagesFailurePager: func(options *paginggroup.PagingClientGetMultiplePagesFailureOptions) (resp azfake.PagerResponder[paginggroup.PagingClientGetMultiplePagesFailureResponse]) {
			resp.AddPage(http.StatusOK, paginggroup.PagingClientGetMultiplePagesFailureResponse{
				ProductResult: paginggroup.ProductResult{
					Values: page1,
				},
			}, nil)
			resp.AddResponseError(http.StatusInternalServerError, "InternalServerError")
			return
		},
	}
	client, err := paginggroup.NewPagingClient(&azcore.ClientOptions{
		Transport: fake.NewPagingServerTransport(&server),
	})
	require.NoError(t, err)
	pager := client.NewGetMultiplePagesFailurePager(nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		if err != nil {
			var respErr *azcore.ResponseError
			require.ErrorAs(t, err, &respErr)
			require.EqualValues(t, "InternalServerError", respErr.ErrorCode)
			require.EqualValues(t, http.StatusInternalServerError, respErr.StatusCode)
			require.Zero(t, page)
			break
		}
		require.EqualValues(t, page1, page.Values)
		count++
	}
	require.EqualValues(t, 1, count)
}
