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

// ArrayNullableFloatValueClient - Array of nullable float values
// Don't use this type directly, use [ArrayClient.NewArrayNullableFloatValueClient] instead.
type ArrayNullableFloatValueClient struct {
	internal *azcore.Client
	endpoint string
}

// Get -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ArrayNullableFloatValueClientGetOptions contains the optional parameters for the ArrayNullableFloatValueClient.Get
//     method.
func (client *ArrayNullableFloatValueClient) Get(ctx context.Context, options *ArrayNullableFloatValueClientGetOptions) (ArrayNullableFloatValueClientGetResponse, error) {
	var err error
	const operationName = "ArrayNullableFloatValueClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ArrayNullableFloatValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArrayNullableFloatValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ArrayNullableFloatValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ArrayNullableFloatValueClient) getCreateRequest(ctx context.Context, _ *ArrayNullableFloatValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/array/nullable-float"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ArrayNullableFloatValueClient) getHandleResponse(resp *http.Response) (ArrayNullableFloatValueClientGetResponse, error) {
	result := ArrayNullableFloatValueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Float32Array); err != nil {
		return ArrayNullableFloatValueClientGetResponse{}, err
	}
	return result, nil
}

// Put -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ArrayNullableFloatValueClientPutOptions contains the optional parameters for the ArrayNullableFloatValueClient.Put
//     method.
func (client *ArrayNullableFloatValueClient) Put(ctx context.Context, body []*float32, options *ArrayNullableFloatValueClientPutOptions) (ArrayNullableFloatValueClientPutResponse, error) {
	var err error
	const operationName = "ArrayNullableFloatValueClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ArrayNullableFloatValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArrayNullableFloatValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ArrayNullableFloatValueClientPutResponse{}, err
	}
	return ArrayNullableFloatValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ArrayNullableFloatValueClient) putCreateRequest(ctx context.Context, body []*float32, _ *ArrayNullableFloatValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/array/nullable-float"
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
