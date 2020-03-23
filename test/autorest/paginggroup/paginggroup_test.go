// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package paginggrouptest

import (
	"context"
	"generatortests/autorest/generated/paginggroup"
	"generatortests/helpers"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

func getPagingOperations(t *testing.T) paginggroup.PagingOperations {
	options := paginggroup.DefaultClientOptions()
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.HTTPClient = httpClientWithCookieJar()
	client, err := paginggroup.NewClient("http://localhost:3000", &options)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	return client.PagingOperations()
}

// copied from azcore, fix once type is exported
type transportFunc func(context.Context, *http.Request) (*http.Response, error)

func (tf transportFunc) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	return tf(ctx, req)
}

func httpClientWithCookieJar() azcore.Transport {
	j, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Jar = j
	return transportFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return http.DefaultClient.Do(req.WithContext(ctx))
	})
}

// GetMultiplePages - A paging operation that includes a nextLink that has 10 pages
func TestGetMultiplePages(t *testing.T) {
	client := getPagingOperations(t)
	page, err := client.GetMultiplePages(nil)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err = page.Err(); err != nil {
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
	client := getPagingOperations(t)
	page, err := client.GetMultiplePagesFailure()
	if err != nil {
		t.Fatal(err)
	}
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
	client := getPagingOperations(t)
	page, err := client.GetMultiplePagesFailureURI()
	if err != nil {
		t.Fatal(err)
	}
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
	client := getPagingOperations(t)
	page, err := client.GetMultiplePagesFragmentNextLink("1.6", "test_user")
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.OdataProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err = page.Err(); err != nil {
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
	client := getPagingOperations(t)
	page, err := client.GetMultiplePagesFragmentWithGroupingNextLink("1.6", "test_user")
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.OdataProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err = page.Err(); err != nil {
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
	t.Skip("LRO NYI")
}

// GetMultiplePagesRetryFirst - A paging operation that fails on the first call with 500 and then retries and then get a response including a nextLink that has 10 pages
func TestGetMultiplePagesRetryFirst(t *testing.T) {
	client := getPagingOperations(t)
	page, err := client.GetMultiplePagesRetryFirst()
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err = page.Err(); err != nil {
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
	client := getPagingOperations(t)
	page, err := client.GetMultiplePagesRetrySecond()
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err = page.Err(); err != nil {
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
	client := getPagingOperations(t)
	page, err := client.GetMultiplePagesWithOffset(0, nil)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err = page.Err(); err != nil {
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
	client := getPagingOperations(t)
	page, err := client.GetNoItemNamePages()
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResultValue.Value) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err = page.Err(); err != nil {
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
	client := getPagingOperations(t)
	resp, err := client.GetNullNextLinkNamePages(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(*resp.ProductResult.Values) == 0 {
		t.Fatal("missing payload")
	}
}

// GetOdataMultiplePages - A paging operation that includes a nextLink in odata format that has 10 pages
func TestGetOdataMultiplePages(t *testing.T) {
	client := getPagingOperations(t)
	page, err := client.GetOdataMultiplePages(nil)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.OdataProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err = page.Err(); err != nil {
		t.Fatal(err)
	}
	if count != 10 {
		helpers.DeepEqualOrFatal(t, count, 10)
	}
	if page.PageResponse() == nil {
		t.Fatal("unexpected nil payload")
	}
}

// GetSinglePages - A paging operation that finishes on the first call without a nextlink
func TestGetSinglePages(t *testing.T) {
	client := getPagingOperations(t)
	page, err := client.GetSinglePages()
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for page.NextPage(context.Background()) {
		resp := page.PageResponse()
		if len(*resp.ProductResult.Values) == 0 {
			t.Fatal("missing payload")
		}
		count++
	}
	if err = page.Err(); err != nil {
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
	client := getPagingOperations(t)
	page, err := client.GetSinglePagesFailure()
	if err != nil {
		t.Fatal(err)
	}
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
