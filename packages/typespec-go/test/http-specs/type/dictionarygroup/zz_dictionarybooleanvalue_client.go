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

// DictionaryBooleanValueClient - Dictionary of boolean values
// Don't use this type directly, use [DictionaryClient.NewDictionaryBooleanValueClient] instead.
type DictionaryBooleanValueClient struct {
	internal *azcore.Client
	endpoint string
}

// Get -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - DictionaryBooleanValueClientGetOptions contains the optional parameters for the DictionaryBooleanValueClient.Get
//     method.
func (client *DictionaryBooleanValueClient) Get(ctx context.Context, options *DictionaryBooleanValueClientGetOptions) (DictionaryBooleanValueClientGetResponse, error) {
	var err error
	const operationName = "DictionaryBooleanValueClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return DictionaryBooleanValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DictionaryBooleanValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DictionaryBooleanValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DictionaryBooleanValueClient) getCreateRequest(ctx context.Context, _ *DictionaryBooleanValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/dictionary/boolean"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DictionaryBooleanValueClient) getHandleResponse(resp *http.Response) (DictionaryBooleanValueClientGetResponse, error) {
	result := DictionaryBooleanValueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return DictionaryBooleanValueClientGetResponse{}, err
	}
	return result, nil
}

// Put -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - DictionaryBooleanValueClientPutOptions contains the optional parameters for the DictionaryBooleanValueClient.Put
//     method.
func (client *DictionaryBooleanValueClient) Put(ctx context.Context, body map[string]*bool, options *DictionaryBooleanValueClientPutOptions) (DictionaryBooleanValueClientPutResponse, error) {
	var err error
	const operationName = "DictionaryBooleanValueClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return DictionaryBooleanValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DictionaryBooleanValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DictionaryBooleanValueClientPutResponse{}, err
	}
	return DictionaryBooleanValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *DictionaryBooleanValueClient) putCreateRequest(ctx context.Context, body map[string]*bool, _ *DictionaryBooleanValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/dictionary/boolean"
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
