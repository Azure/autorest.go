// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package paginggroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

func newPagingClient() PagingOperations {
	options := DefaultConnectionOptions()
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.HTTPClient = httpClientWithCookieJar()
	return NewPagingClient(NewDefaultConnection(&options))
}

func httpClientWithCookieJar() azcore.Transport {
	j, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Jar = j
	return azcore.TransportFunc(func(req *http.Request) (*http.Response, error) {
		return http.DefaultClient.Do(req)
	})
}

// GetMultiplePages - A paging operation that includes a nextLink that has 10 pages
func TestGetMultiplePages(t *testing.T) {
	client := newPagingClient()
	page := client.GetMultiplePages(nil)
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err := page.Err(); err != nil {
		t.Fatal(err)
	}
	if count != 10 {
		helpers.DeepEqualOrFatal(t, count, 10)
	}
	if page.PageResponse() == nil {
		t.Fatal("unexpected nil payload")
	}
}

// GetMultiplePagesFailure - A paging operation that receives a 400 on the second call
func TestGetMultiplePagesFailure(t *testing.T) {
	client := newPagingClient()
	page := client.GetMultiplePagesFailure(nil)
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if page.Err() == nil {
		t.Fatal("unexpected nil error")
	}
	if count != 1 {
		helpers.DeepEqualOrFatal(t, count, 1)
	}
	if page.PageResponse() == nil {
		t.Fatal("unexpected nil payload")
	}
}

// GetMultiplePagesFailureURI - A paging operation that receives an invalid nextLink
func TestGetMultiplePagesFailureURI(t *testing.T) {
	client := newPagingClient()
	page := client.GetMultiplePagesFailureURI(nil)
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if page.Err() == nil {
		t.Fatal("unexpected nil error")
	}
	if count != 1 {
		helpers.DeepEqualOrFatal(t, count, 1)
	}
	if page.PageResponse() == nil {
		t.Fatal("unexpected nil payload")
	}
}

// GetMultiplePagesFragmentNextLink - A paging operation that doesn't return a full URL, just a fragment
func TestGetMultiplePagesFragmentNextLink(t *testing.T) {
	client := newPagingClient()
	page := client.GetMultiplePagesFragmentNextLink("1.6", "test_user", nil)
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.OdataProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err := page.Err(); err != nil {
		t.Fatal(err)
	}
	if count != 10 {
		helpers.DeepEqualOrFatal(t, count, 10)
	}
	if page.PageResponse() == nil {
		t.Fatal("unexpected nil payload")
	}
}

// GetMultiplePagesFragmentWithGroupingNextLink - A paging operation that doesn't return a full URL, just a fragment with parameters grouped
func TestGetMultiplePagesFragmentWithGroupingNextLink(t *testing.T) {
	client := newPagingClient()
	page := client.GetMultiplePagesFragmentWithGroupingNextLink(CustomParameterGroup{
		ApiVersion: "1.6",
		Tenant:     "test_user",
	})
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.OdataProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err := page.Err(); err != nil {
		t.Fatal(err)
	}
	if count != 10 {
		helpers.DeepEqualOrFatal(t, count, 10)
	}
	if page.PageResponse() == nil {
		t.Fatal("unexpected nil payload")
	}
}

// GetMultiplePagesLro - A long-running paging operation that includes a nextLink that has 10 pages
func TestGetMultiplePagesLro(t *testing.T) {
	client := newPagingClient()
	resp, err := client.BeginGetMultiplePagesLro(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = client.ResumeGetMultiplePagesLro(rt)
	if err != nil {
		t.Fatal(err)
	}
	pager, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for pager.NextPage(context.Background()) {
		resp := pager.PageResponse()
		if len(*resp.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err = pager.Err(); err != nil {
		t.Fatal(err)
	}
	const pageCount = 10
	if count != pageCount {
		helpers.DeepEqualOrFatal(t, count, pageCount)
	}
}

// GetMultiplePagesRetryFirst - A paging operation that fails on the first call with 500 and then retries and then get a response including a nextLink that has 10 pages
func TestGetMultiplePagesRetryFirst(t *testing.T) {
	client := newPagingClient()
	page := client.GetMultiplePagesRetryFirst(nil)
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err := page.Err(); err != nil {
		t.Fatal(err)
	}
	if count != 10 {
		helpers.DeepEqualOrFatal(t, count, 10)
	}
	if page.PageResponse() == nil {
		t.Fatal("unexpected nil payload")
	}
}

// GetMultiplePagesRetrySecond - A paging operation that includes a nextLink that has 10 pages, of which the 2nd call fails first with 500. The client should retry and finish all 10 pages eventually.
func TestGetMultiplePagesRetrySecond(t *testing.T) {
	client := newPagingClient()
	page := client.GetMultiplePagesRetrySecond(nil)
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err := page.Err(); err != nil {
		t.Fatal(err)
	}
	if count != 10 {
		helpers.DeepEqualOrFatal(t, count, 10)
	}
	if page.PageResponse() == nil {
		t.Fatal("unexpected nil payload")
	}
}

// GetMultiplePagesWithOffset - A paging operation that includes a nextLink that has 10 pages
func TestGetMultiplePagesWithOffset(t *testing.T) {
	client := newPagingClient()
	page := client.GetMultiplePagesWithOffset(PagingGetMultiplePagesWithOffsetOptions{})
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err := page.Err(); err != nil {
		t.Fatal(err)
	}
	if count != 10 {
		helpers.DeepEqualOrFatal(t, count, 10)
	}
	if page.PageResponse() == nil {
		t.Fatal("unexpected nil payload")
	}
}

// GetNoItemNamePages - A paging operation that must return result of the default 'value' node.
func TestGetNoItemNamePages(t *testing.T) {
	client := newPagingClient()
	page := client.GetNoItemNamePages(nil)
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResultValue.Value) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err := page.Err(); err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		helpers.DeepEqualOrFatal(t, count, 1)
	}
	if page.PageResponse() == nil {
		t.Fatal("unexpected nil payload")
	}
}

// GetNullNextLinkNamePages - A paging operation that must ignore any kind of nextLink, and stop after page 1.
func TestGetNullNextLinkNamePages(t *testing.T) {
	client := newPagingClient()
	resp, err := client.GetNullNextLinkNamePages(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(*resp.ProductResult.Values) == 0 {
		t.Fatal("missing payload")
	}
}

// GetOdataMultiplePages - A paging operation that includes a nextLink in odata format that has 10 pages
func TestGetOdataMultiplePages(t *testing.T) {
	client := newPagingClient()
	page := client.GetOdataMultiplePages(nil)
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.OdataProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err := page.Err(); err != nil {
		t.Fatal(err)
	}
	if count != 10 {
		helpers.DeepEqualOrFatal(t, count, 10)
	}
	if page.PageResponse() == nil {
		t.Fatal("unexpected nil payload")
	}
}

// GetPagingModelWithItemNameWithXmsClientName - A paging operation that returns a paging model whose item name is is overriden by x-ms-client-name 'indexes'.
func TestGetPagingModelWithItemNameWithXmsClientName(t *testing.T) {
	client := newPagingClient()
	page := client.GetPagingModelWithItemNameWithXmsClientName(nil)
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResultValueWithXmsClientName.Indexes) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err := page.Err(); err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		helpers.DeepEqualOrFatal(t, count, 1)
	}
	if page.PageResponse() == nil {
		t.Fatal("unexpected nil payload")
	}
}

// GetSinglePages - A paging operation that finishes on the first call without a nextlink
func TestGetSinglePages(t *testing.T) {
	client := newPagingClient()
	page := client.GetSinglePages(nil)
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err := page.Err(); err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		helpers.DeepEqualOrFatal(t, count, 1)
	}
	if page.PageResponse() == nil {
		t.Fatal("unexpected nil payload")
	}
}

// GetSinglePagesFailure - A paging operation that receives a 400 on the first call
func TestGetSinglePagesFailure(t *testing.T) {
	client := newPagingClient()
	page := client.GetSinglePagesFailure(nil)
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if page.Err() == nil {
		t.Fatal("unexpected nil error")
	}
	if count != 0 {
		helpers.DeepEqualOrFatal(t, count, 0)
	}
	if page.PageResponse() != nil {
		t.Fatal("expected nil payload")
	}
}

// GetWithQueryParams - A paging operation that includes a next operation. It has a different query parameter from it's next operation nextOperationWithQueryParams. Returns a ProductResult
func TestGetWithQueryParams(t *testing.T) {
	client := newPagingClient()
	page := client.GetWithQueryParams(100, nil)
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err := page.Err(); err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		helpers.DeepEqualOrFatal(t, count, 2)
	}
	if page.PageResponse() == nil {
		t.Fatal("unexpected nil payload")
	}
}
