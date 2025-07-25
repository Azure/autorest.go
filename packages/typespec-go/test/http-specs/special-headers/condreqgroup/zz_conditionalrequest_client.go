// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package condreqgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"time"
)

// ConditionalRequestClient - Illustrates conditional request headers
// Don't use this type directly, use a constructor function instead.
type ConditionalRequestClient struct {
	internal *azcore.Client
	endpoint string
}

// HeadIfModifiedSince - Check when only If-Modified-Since in header is defined.
//   - options - ConditionalRequestClientHeadIfModifiedSinceOptions contains the optional parameters for the ConditionalRequestClient.HeadIfModifiedSince
//     method.
func (client *ConditionalRequestClient) HeadIfModifiedSince(ctx context.Context, options *ConditionalRequestClientHeadIfModifiedSinceOptions) (ConditionalRequestClientHeadIfModifiedSinceResponse, error) {
	var err error
	const operationName = "ConditionalRequestClient.HeadIfModifiedSince"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.headIfModifiedSinceCreateRequest(ctx, options)
	if err != nil {
		return ConditionalRequestClientHeadIfModifiedSinceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConditionalRequestClientHeadIfModifiedSinceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConditionalRequestClientHeadIfModifiedSinceResponse{}, err
	}
	return ConditionalRequestClientHeadIfModifiedSinceResponse{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// headIfModifiedSinceCreateRequest creates the HeadIfModifiedSince request.
func (client *ConditionalRequestClient) headIfModifiedSinceCreateRequest(ctx context.Context, options *ConditionalRequestClientHeadIfModifiedSinceOptions) (*policy.Request, error) {
	urlPath := "/special-headers/conditional-request/if-modified-since"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.IfModifiedSince != nil {
		req.Raw().Header["If-Modified-Since"] = []string{options.IfModifiedSince.Format(time.RFC1123)}
	}
	return req, nil
}

// PostIfMatch - Check when only If-Match in header is defined.
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ConditionalRequestClientPostIfMatchOptions contains the optional parameters for the ConditionalRequestClient.PostIfMatch
//     method.
func (client *ConditionalRequestClient) PostIfMatch(ctx context.Context, options *ConditionalRequestClientPostIfMatchOptions) (ConditionalRequestClientPostIfMatchResponse, error) {
	var err error
	const operationName = "ConditionalRequestClient.PostIfMatch"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postIfMatchCreateRequest(ctx, options)
	if err != nil {
		return ConditionalRequestClientPostIfMatchResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConditionalRequestClientPostIfMatchResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConditionalRequestClientPostIfMatchResponse{}, err
	}
	return ConditionalRequestClientPostIfMatchResponse{}, nil
}

// postIfMatchCreateRequest creates the PostIfMatch request.
func (client *ConditionalRequestClient) postIfMatchCreateRequest(ctx context.Context, options *ConditionalRequestClientPostIfMatchOptions) (*policy.Request, error) {
	urlPath := "/special-headers/conditional-request/if-match"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	return req, nil
}

// PostIfNoneMatch - Check when only If-None-Match in header is defined.
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ConditionalRequestClientPostIfNoneMatchOptions contains the optional parameters for the ConditionalRequestClient.PostIfNoneMatch
//     method.
func (client *ConditionalRequestClient) PostIfNoneMatch(ctx context.Context, options *ConditionalRequestClientPostIfNoneMatchOptions) (ConditionalRequestClientPostIfNoneMatchResponse, error) {
	var err error
	const operationName = "ConditionalRequestClient.PostIfNoneMatch"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postIfNoneMatchCreateRequest(ctx, options)
	if err != nil {
		return ConditionalRequestClientPostIfNoneMatchResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConditionalRequestClientPostIfNoneMatchResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConditionalRequestClientPostIfNoneMatchResponse{}, err
	}
	return ConditionalRequestClientPostIfNoneMatchResponse{}, nil
}

// postIfNoneMatchCreateRequest creates the PostIfNoneMatch request.
func (client *ConditionalRequestClient) postIfNoneMatchCreateRequest(ctx context.Context, options *ConditionalRequestClientPostIfNoneMatchOptions) (*policy.Request, error) {
	urlPath := "/special-headers/conditional-request/if-none-match"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	return req, nil
}

// PostIfUnmodifiedSince - Check when only If-Unmodified-Since in header is defined.
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ConditionalRequestClientPostIfUnmodifiedSinceOptions contains the optional parameters for the ConditionalRequestClient.PostIfUnmodifiedSince
//     method.
func (client *ConditionalRequestClient) PostIfUnmodifiedSince(ctx context.Context, options *ConditionalRequestClientPostIfUnmodifiedSinceOptions) (ConditionalRequestClientPostIfUnmodifiedSinceResponse, error) {
	var err error
	const operationName = "ConditionalRequestClient.PostIfUnmodifiedSince"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postIfUnmodifiedSinceCreateRequest(ctx, options)
	if err != nil {
		return ConditionalRequestClientPostIfUnmodifiedSinceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConditionalRequestClientPostIfUnmodifiedSinceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConditionalRequestClientPostIfUnmodifiedSinceResponse{}, err
	}
	return ConditionalRequestClientPostIfUnmodifiedSinceResponse{}, nil
}

// postIfUnmodifiedSinceCreateRequest creates the PostIfUnmodifiedSince request.
func (client *ConditionalRequestClient) postIfUnmodifiedSinceCreateRequest(ctx context.Context, options *ConditionalRequestClientPostIfUnmodifiedSinceOptions) (*policy.Request, error) {
	urlPath := "/special-headers/conditional-request/if-unmodified-since"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.IfUnmodifiedSince != nil {
		req.Raw().Header["If-Unmodified-Since"] = []string{options.IfUnmodifiedSince.Format(time.RFC1123)}
	}
	return req, nil
}
