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

// ValueTypesBooleanClient contains the methods for the ValueTypesBoolean group.
// Don't use this type directly, use [ValueTypesClient.NewValueTypesBooleanClient] instead.
type ValueTypesBooleanClient struct {
	internal *azcore.Client
	endpoint string
}

// Get - Get call
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ValueTypesBooleanClientGetOptions contains the optional parameters for the ValueTypesBooleanClient.Get method.
func (client *ValueTypesBooleanClient) Get(ctx context.Context, options *ValueTypesBooleanClientGetOptions) (ValueTypesBooleanClientGetResponse, error) {
	var err error
	const operationName = "ValueTypesBooleanClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ValueTypesBooleanClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesBooleanClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesBooleanClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ValueTypesBooleanClient) getCreateRequest(ctx context.Context, _ *ValueTypesBooleanClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/boolean"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ValueTypesBooleanClient) getHandleResponse(resp *http.Response) (ValueTypesBooleanClientGetResponse, error) {
	result := ValueTypesBooleanClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BooleanProperty); err != nil {
		return ValueTypesBooleanClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
// If the operation fails it returns an *azcore.ResponseError type.
//   - body - body
//   - options - ValueTypesBooleanClientPutOptions contains the optional parameters for the ValueTypesBooleanClient.Put method.
func (client *ValueTypesBooleanClient) Put(ctx context.Context, body BooleanProperty, options *ValueTypesBooleanClientPutOptions) (ValueTypesBooleanClientPutResponse, error) {
	var err error
	const operationName = "ValueTypesBooleanClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ValueTypesBooleanClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesBooleanClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesBooleanClientPutResponse{}, err
	}
	return ValueTypesBooleanClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ValueTypesBooleanClient) putCreateRequest(ctx context.Context, body BooleanProperty, _ *ValueTypesBooleanClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/boolean"
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
