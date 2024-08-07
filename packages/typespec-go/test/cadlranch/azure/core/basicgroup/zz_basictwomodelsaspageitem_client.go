// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package basicgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// BasicTwoModelsAsPageItemClient contains the methods for the BasicTwoModelsAsPageItem group.
// Don't use this type directly, use [BasicClient.NewBasicTwoModelsAsPageItemClient] instead.
type BasicTwoModelsAsPageItemClient struct {
	internal *azcore.Client
}

// NewListFirstItemPager - Two operations with two different page item types should be successfully generated. Should generate
// model for FirstItem.
//
// Generated from API version 2022-12-01-preview
//   - options - BasicTwoModelsAsPageItemClientListFirstItemOptions contains the optional parameters for the BasicTwoModelsAsPageItemClient.NewListFirstItemPager
//     method.
func (client *BasicTwoModelsAsPageItemClient) NewListFirstItemPager(options *BasicTwoModelsAsPageItemClientListFirstItemOptions) *runtime.Pager[BasicTwoModelsAsPageItemClientListFirstItemResponse] {
	return runtime.NewPager(runtime.PagingHandler[BasicTwoModelsAsPageItemClientListFirstItemResponse]{
		More: func(page BasicTwoModelsAsPageItemClientListFirstItemResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BasicTwoModelsAsPageItemClientListFirstItemResponse) (BasicTwoModelsAsPageItemClientListFirstItemResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BasicTwoModelsAsPageItemClient.NewListFirstItemPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listFirstItemCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return BasicTwoModelsAsPageItemClientListFirstItemResponse{}, err
			}
			return client.listFirstItemHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listFirstItemCreateRequest creates the ListFirstItem request.
func (client *BasicTwoModelsAsPageItemClient) listFirstItemCreateRequest(ctx context.Context, _ *BasicTwoModelsAsPageItemClientListFirstItemOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/first-item"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listFirstItemHandleResponse handles the ListFirstItem response.
func (client *BasicTwoModelsAsPageItemClient) listFirstItemHandleResponse(resp *http.Response) (BasicTwoModelsAsPageItemClientListFirstItemResponse, error) {
	result := BasicTwoModelsAsPageItemClientListFirstItemResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedFirstItem); err != nil {
		return BasicTwoModelsAsPageItemClientListFirstItemResponse{}, err
	}
	return result, nil
}

// NewListSecondItemPager - Two operations with two different page item types should be successfully generated. Should generate
// model for SecondItem.
//
// Generated from API version 2022-12-01-preview
//   - options - BasicTwoModelsAsPageItemClientListSecondItemOptions contains the optional parameters for the BasicTwoModelsAsPageItemClient.NewListSecondItemPager
//     method.
func (client *BasicTwoModelsAsPageItemClient) NewListSecondItemPager(options *BasicTwoModelsAsPageItemClientListSecondItemOptions) *runtime.Pager[BasicTwoModelsAsPageItemClientListSecondItemResponse] {
	return runtime.NewPager(runtime.PagingHandler[BasicTwoModelsAsPageItemClientListSecondItemResponse]{
		More: func(page BasicTwoModelsAsPageItemClientListSecondItemResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BasicTwoModelsAsPageItemClientListSecondItemResponse) (BasicTwoModelsAsPageItemClientListSecondItemResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BasicTwoModelsAsPageItemClient.NewListSecondItemPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listSecondItemCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return BasicTwoModelsAsPageItemClientListSecondItemResponse{}, err
			}
			return client.listSecondItemHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listSecondItemCreateRequest creates the ListSecondItem request.
func (client *BasicTwoModelsAsPageItemClient) listSecondItemCreateRequest(ctx context.Context, _ *BasicTwoModelsAsPageItemClientListSecondItemOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/second-item"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSecondItemHandleResponse handles the ListSecondItem response.
func (client *BasicTwoModelsAsPageItemClient) listSecondItemHandleResponse(resp *http.Response) (BasicTwoModelsAsPageItemClientListSecondItemResponse, error) {
	result := BasicTwoModelsAsPageItemClientListSecondItemResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedSecondItem); err != nil {
		return BasicTwoModelsAsPageItemClientListSecondItemResponse{}, err
	}
	return result, nil
}
