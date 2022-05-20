// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package paginggroup

import (
	"context"
	"generatortests"
	"net/http"
	"net/http/cookiejar"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newPagingClient() *PagingClient {
	options := azcore.ClientOptions{}
	options.Retry.RetryDelay = time.Second
	options.Transport = httpClientWithCookieJar()
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &options)
	return NewPagingClient(pl)
}

func httpClientWithCookieJar() policy.Transporter {
	j, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Jar = j
	return http.DefaultClient
}

// GetMultiplePages - A paging operation that includes a nextLink that has 10 pages
func TestGetMultiplePages(t *testing.T) {
	client := newPagingClient()
	pager := client.NewGetMultiplePagesPager(nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.NotZero(t, page)
		require.NotEmpty(t, page.Values)
		count++
	}
	if r := cmp.Diff(count, 10); r != "" {
		t.Fatal(r)
	}
	if _, err := pager.NextPage(context.Background()); err == nil {
		t.Fatal("unexpected nil error")
	}
}

// GetMultiplePagesFailure - A paging operation that receives a 400 on the second call
func TestGetMultiplePagesFailure(t *testing.T) {
	client := newPagingClient()
	pager := client.NewGetMultiplePagesFailurePager(nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		if err != nil {
			if !reflect.ValueOf(page).IsZero() {
				t.Fatal("expected empty payload")
			}
			break
		}
		if len(page.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if r := cmp.Diff(count, 1); r != "" {
		t.Fatal(r)
	}
}

// GetMultiplePagesFailureURI - A paging operation that receives an invalid nextLink
func TestGetMultiplePagesFailureURI(t *testing.T) {
	client := newPagingClient()
	pager := client.NewGetMultiplePagesFailureURIPager(nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		if err != nil {
			if !reflect.ValueOf(page).IsZero() {
				t.Fatal("expected empty payload")
			}
			break
		}
		count++
	}
	if r := cmp.Diff(count, 1); r != "" {
		t.Fatal(r)
	}
}

// GetMultiplePagesFragmentNextLink - A paging operation that doesn't return a full URL, just a fragment
func TestGetMultiplePagesFragmentNextLink(t *testing.T) {
	client := newPagingClient()
	pager := client.NewGetMultiplePagesFragmentNextLinkPager("1.6", "test_user", nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.NotZero(t, page)
		require.NotEmpty(t, page.Values)
		count++
	}
	if r := cmp.Diff(count, 10); r != "" {
		t.Fatal(r)
	}
}

// GetMultiplePagesFragmentWithGroupingNextLink - A paging operation that doesn't return a full URL, just a fragment with parameters grouped
func TestGetMultiplePagesFragmentWithGroupingNextLink(t *testing.T) {
	client := newPagingClient()
	pager := client.NewGetMultiplePagesFragmentWithGroupingNextLinkPager(CustomParameterGroup{
		APIVersion: "1.6",
		Tenant:     "test_user",
	}, nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.NotZero(t, page)
		require.NotEmpty(t, page.Values)
		count++
	}
	if r := cmp.Diff(count, 10); r != "" {
		t.Fatal(r)
	}
}

// GetMultiplePagesLro - A long-running paging operation that includes a nextLink that has 10 pages
func TestGetMultiplePagesLro(t *testing.T) {
	client := newPagingClient()
	poller, err := client.BeginGetMultiplePagesLRO(context.Background(), nil)
	require.NoError(t, err)
	rt, err := poller.ResumeToken()
	require.NoError(t, err)
	poller, err = client.BeginGetMultiplePagesLRO(context.Background(), &PagingClientBeginGetMultiplePagesLROOptions{
		ResumeToken: rt,
	})
	require.NoError(t, err)
	pager, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.NotZero(t, page)
		require.NotEmpty(t, page.Values)
		count++
	}
	const pageCount = 10
	if r := cmp.Diff(count, pageCount); r != "" {
		t.Fatal(r)
	}
}

// GetMultiplePagesRetryFirst - A paging operation that fails on the first call with 500 and then retries and then get a response including a nextLink that has 10 pages
func TestGetMultiplePagesRetryFirst(t *testing.T) {
	client := newPagingClient()
	pager := client.NewGetMultiplePagesRetryFirstPager(nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.NotZero(t, page)
		require.NotEmpty(t, page.Values)
		count++
	}
	if r := cmp.Diff(count, 10); r != "" {
		t.Fatal(r)
	}
}

// GetMultiplePagesRetrySecond - A paging operation that includes a nextLink that has 10 pages, of which the 2nd call fails first with 500. The client should retry and finish all 10 pages eventually.
func TestGetMultiplePagesRetrySecond(t *testing.T) {
	client := newPagingClient()
	pager := client.NewGetMultiplePagesRetrySecondPager(nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.NotZero(t, page)
		require.NotEmpty(t, page.Values)
		count++
	}
	if r := cmp.Diff(count, 10); r != "" {
		t.Fatal(r)
	}
}

// GetMultiplePagesWithOffset - A paging operation that includes a nextLink that has 10 pages
func TestGetMultiplePagesWithOffset(t *testing.T) {
	client := newPagingClient()
	pager := client.NewGetMultiplePagesWithOffsetPager(PagingClientGetMultiplePagesWithOffsetOptions{})
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.NotZero(t, page)
		require.NotEmpty(t, page.Values)
		count++
	}
	if r := cmp.Diff(count, 10); r != "" {
		t.Fatal(r)
	}
}

// GetNoItemNamePages - A paging operation that must return result of the default 'value' node.
func TestGetNoItemNamePages(t *testing.T) {
	client := newPagingClient()
	pager := client.NewGetNoItemNamePagesPager(nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.NotZero(t, page)
		require.NotEmpty(t, page.Value)
		count++
	}
	if r := cmp.Diff(count, 1); r != "" {
		t.Fatal(r)
	}
}

// GetNullNextLinkNamePages - A paging operation that must ignore any kind of nextLink, and stop after page 1.
func TestGetNullNextLinkNamePages(t *testing.T) {
	client := newPagingClient()
	pager := client.NewGetNullNextLinkNamePagesPager(nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.NotZero(t, page)
		require.NotEmpty(t, page.Values)
		count++
	}
	if r := cmp.Diff(count, 1); r != "" {
		t.Fatal(r)
	}
	if _, err := pager.NextPage(context.Background()); err == nil {
		t.Fatal("unexpected nil error")
	}
}

// GetOdataMultiplePages - A paging operation that includes a nextLink in odata format that has 10 pages
func TestGetOdataMultiplePages(t *testing.T) {
	client := newPagingClient()
	pager := client.NewGetODataMultiplePagesPager(nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.NotZero(t, page)
		require.NotEmpty(t, page.Values)
		count++
	}
	if r := cmp.Diff(count, 10); r != "" {
		t.Fatal(r)
	}
}

// GetPagingModelWithItemNameWithXMSClientName - A paging operation that returns a paging model whose item name is is overriden by x-ms-client-name 'indexes'.
func TestGetPagingModelWithItemNameWithXMSClientName(t *testing.T) {
	client := newPagingClient()
	pager := client.NewGetPagingModelWithItemNameWithXMSClientNamePager(nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.NotZero(t, page)
		require.NotEmpty(t, page.Indexes)
		count++
	}
	if r := cmp.Diff(count, 1); r != "" {
		t.Fatal(r)
	}
}

// GetSinglePages - A paging operation that finishes on the first call without a nextlink
func TestGetSinglePages(t *testing.T) {
	client := newPagingClient()
	pager := client.NewGetSinglePagesPager(nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.NotZero(t, page)
		require.NotEmpty(t, page.Values)
		count++
	}
	if r := cmp.Diff(count, 1); r != "" {
		t.Fatal(r)
	}
}

// GetSinglePagesFailure - A paging operation that receives a 400 on the first call
func TestGetSinglePagesFailure(t *testing.T) {
	client := newPagingClient()
	pager := client.NewGetSinglePagesFailurePager(nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		if err != nil {
			if !reflect.ValueOf(page).IsZero() {
				t.Fatal("expected empty payload")
			}
			break
		}
		count++
	}
	if r := cmp.Diff(count, 0); r != "" {
		t.Fatal(r)
	}
}

// GetWithQueryParams - A paging operation that includes a next operation. It has a different query parameter from it's next operation nextOperationWithQueryParams. Returns a ProductResult
func TestGetWithQueryParams(t *testing.T) {
	client := newPagingClient()
	pager := client.NewGetWithQueryParamsPager(100, nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.NotZero(t, page)
		require.NotEmpty(t, page.ProductResult.Values)
		count++
	}
	if r := cmp.Diff(count, 2); r != "" {
		t.Fatal(r)
	}
}
