// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package dictionarygroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// DictionaryModelValueClient - Dictionary of model values
// Don't use this type directly, use [DictionaryClient.NewDictionaryModelValueClient] instead.
type DictionaryModelValueClient struct {
	internal *azcore.Client
}

// Get -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - DictionaryModelValueClientGetOptions contains the optional parameters for the DictionaryModelValueClient.Get
//     method.
func (client *DictionaryModelValueClient) Get(ctx context.Context, options *DictionaryModelValueClientGetOptions) (DictionaryModelValueClientGetResponse, error) {
	var err error
	const operationName = "DictionaryModelValueClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return DictionaryModelValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DictionaryModelValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DictionaryModelValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DictionaryModelValueClient) getCreateRequest(ctx context.Context, _ *DictionaryModelValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/dictionary/model"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DictionaryModelValueClient) getHandleResponse(resp *http.Response) (DictionaryModelValueClientGetResponse, error) {
	result := DictionaryModelValueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return DictionaryModelValueClientGetResponse{}, err
	}
	return result, nil
}

// Put -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - DictionaryModelValueClientPutOptions contains the optional parameters for the DictionaryModelValueClient.Put
//     method.
func (client *DictionaryModelValueClient) Put(ctx context.Context, body map[string]*InnerModel, options *DictionaryModelValueClientPutOptions) (DictionaryModelValueClientPutResponse, error) {
	var err error
	const operationName = "DictionaryModelValueClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return DictionaryModelValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DictionaryModelValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DictionaryModelValueClientPutResponse{}, err
	}
	return DictionaryModelValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *DictionaryModelValueClient) putCreateRequest(ctx context.Context, body map[string]*InnerModel, _ *DictionaryModelValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/dictionary/model"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}