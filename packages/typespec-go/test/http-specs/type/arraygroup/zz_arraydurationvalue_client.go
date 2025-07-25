// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package arraygroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ArrayDurationValueClient - Array of duration values
// Don't use this type directly, use [ArrayClient.NewArrayDurationValueClient] instead.
type ArrayDurationValueClient struct {
	internal *azcore.Client
	endpoint string
}

// Get -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ArrayDurationValueClientGetOptions contains the optional parameters for the ArrayDurationValueClient.Get method.
func (client *ArrayDurationValueClient) Get(ctx context.Context, options *ArrayDurationValueClientGetOptions) (ArrayDurationValueClientGetResponse, error) {
	var err error
	const operationName = "ArrayDurationValueClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ArrayDurationValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArrayDurationValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ArrayDurationValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ArrayDurationValueClient) getCreateRequest(ctx context.Context, _ *ArrayDurationValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/array/duration"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ArrayDurationValueClient) getHandleResponse(resp *http.Response) (ArrayDurationValueClientGetResponse, error) {
	result := ArrayDurationValueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StringArray); err != nil {
		return ArrayDurationValueClientGetResponse{}, err
	}
	return result, nil
}

// Put -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ArrayDurationValueClientPutOptions contains the optional parameters for the ArrayDurationValueClient.Put method.
func (client *ArrayDurationValueClient) Put(ctx context.Context, body []string, options *ArrayDurationValueClientPutOptions) (ArrayDurationValueClientPutResponse, error) {
	var err error
	const operationName = "ArrayDurationValueClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ArrayDurationValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArrayDurationValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ArrayDurationValueClientPutResponse{}, err
	}
	return ArrayDurationValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ArrayDurationValueClient) putCreateRequest(ctx context.Context, body []string, _ *ArrayDurationValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/array/duration"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
