// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package valuetypesgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ValueTypesCollectionsModelClient contains the methods for the ValueTypesCollectionsModel group.
// Don't use this type directly, use [ValueTypesClient.NewValueTypesCollectionsModelClient] instead.
type ValueTypesCollectionsModelClient struct {
	internal *azcore.Client
	endpoint string
}

// Get - Get call
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ValueTypesCollectionsModelClientGetOptions contains the optional parameters for the ValueTypesCollectionsModelClient.Get
//     method.
func (client *ValueTypesCollectionsModelClient) Get(ctx context.Context, options *ValueTypesCollectionsModelClientGetOptions) (ValueTypesCollectionsModelClientGetResponse, error) {
	var err error
	const operationName = "ValueTypesCollectionsModelClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ValueTypesCollectionsModelClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesCollectionsModelClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesCollectionsModelClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ValueTypesCollectionsModelClient) getCreateRequest(ctx context.Context, _ *ValueTypesCollectionsModelClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/collections/model"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ValueTypesCollectionsModelClient) getHandleResponse(resp *http.Response) (ValueTypesCollectionsModelClientGetResponse, error) {
	result := ValueTypesCollectionsModelClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CollectionsModelProperty); err != nil {
		return ValueTypesCollectionsModelClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
// If the operation fails it returns an *azcore.ResponseError type.
//   - body - body
//   - options - ValueTypesCollectionsModelClientPutOptions contains the optional parameters for the ValueTypesCollectionsModelClient.Put
//     method.
func (client *ValueTypesCollectionsModelClient) Put(ctx context.Context, body CollectionsModelProperty, options *ValueTypesCollectionsModelClientPutOptions) (ValueTypesCollectionsModelClientPutResponse, error) {
	var err error
	const operationName = "ValueTypesCollectionsModelClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ValueTypesCollectionsModelClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesCollectionsModelClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesCollectionsModelClientPutResponse{}, err
	}
	return ValueTypesCollectionsModelClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ValueTypesCollectionsModelClient) putCreateRequest(ctx context.Context, body CollectionsModelProperty, _ *ValueTypesCollectionsModelClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/collections/model"
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
