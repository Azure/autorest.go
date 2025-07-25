// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package optionalitygroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// OptionalUnionIntLiteralClient contains the methods for the OptionalUnionIntLiteral group.
// Don't use this type directly, use [OptionalClient.NewOptionalUnionIntLiteralClient] instead.
type OptionalUnionIntLiteralClient struct {
	internal *azcore.Client
	endpoint string
}

// GetAll - Get models that will return all properties in the model
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalUnionIntLiteralClientGetAllOptions contains the optional parameters for the OptionalUnionIntLiteralClient.GetAll
//     method.
func (client *OptionalUnionIntLiteralClient) GetAll(ctx context.Context, options *OptionalUnionIntLiteralClientGetAllOptions) (OptionalUnionIntLiteralClientGetAllResponse, error) {
	var err error
	const operationName = "OptionalUnionIntLiteralClient.GetAll"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAllCreateRequest(ctx, options)
	if err != nil {
		return OptionalUnionIntLiteralClientGetAllResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalUnionIntLiteralClientGetAllResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OptionalUnionIntLiteralClientGetAllResponse{}, err
	}
	resp, err := client.getAllHandleResponse(httpResp)
	return resp, err
}

// getAllCreateRequest creates the GetAll request.
func (client *OptionalUnionIntLiteralClient) getAllCreateRequest(ctx context.Context, _ *OptionalUnionIntLiteralClientGetAllOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/union/int/literal/all"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAllHandleResponse handles the GetAll response.
func (client *OptionalUnionIntLiteralClient) getAllHandleResponse(resp *http.Response) (OptionalUnionIntLiteralClientGetAllResponse, error) {
	result := OptionalUnionIntLiteralClientGetAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UnionIntLiteralProperty); err != nil {
		return OptionalUnionIntLiteralClientGetAllResponse{}, err
	}
	return result, nil
}

// GetDefault - Get models that will return the default object
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalUnionIntLiteralClientGetDefaultOptions contains the optional parameters for the OptionalUnionIntLiteralClient.GetDefault
//     method.
func (client *OptionalUnionIntLiteralClient) GetDefault(ctx context.Context, options *OptionalUnionIntLiteralClientGetDefaultOptions) (OptionalUnionIntLiteralClientGetDefaultResponse, error) {
	var err error
	const operationName = "OptionalUnionIntLiteralClient.GetDefault"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDefaultCreateRequest(ctx, options)
	if err != nil {
		return OptionalUnionIntLiteralClientGetDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalUnionIntLiteralClientGetDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OptionalUnionIntLiteralClientGetDefaultResponse{}, err
	}
	resp, err := client.getDefaultHandleResponse(httpResp)
	return resp, err
}

// getDefaultCreateRequest creates the GetDefault request.
func (client *OptionalUnionIntLiteralClient) getDefaultCreateRequest(ctx context.Context, _ *OptionalUnionIntLiteralClientGetDefaultOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/union/int/literal/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDefaultHandleResponse handles the GetDefault response.
func (client *OptionalUnionIntLiteralClient) getDefaultHandleResponse(resp *http.Response) (OptionalUnionIntLiteralClientGetDefaultResponse, error) {
	result := OptionalUnionIntLiteralClientGetDefaultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UnionIntLiteralProperty); err != nil {
		return OptionalUnionIntLiteralClientGetDefaultResponse{}, err
	}
	return result, nil
}

// PutAll - Put a body with all properties present.
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalUnionIntLiteralClientPutAllOptions contains the optional parameters for the OptionalUnionIntLiteralClient.PutAll
//     method.
func (client *OptionalUnionIntLiteralClient) PutAll(ctx context.Context, body UnionIntLiteralProperty, options *OptionalUnionIntLiteralClientPutAllOptions) (OptionalUnionIntLiteralClientPutAllResponse, error) {
	var err error
	const operationName = "OptionalUnionIntLiteralClient.PutAll"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putAllCreateRequest(ctx, body, options)
	if err != nil {
		return OptionalUnionIntLiteralClientPutAllResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalUnionIntLiteralClientPutAllResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalUnionIntLiteralClientPutAllResponse{}, err
	}
	return OptionalUnionIntLiteralClientPutAllResponse{}, nil
}

// putAllCreateRequest creates the PutAll request.
func (client *OptionalUnionIntLiteralClient) putAllCreateRequest(ctx context.Context, body UnionIntLiteralProperty, _ *OptionalUnionIntLiteralClientPutAllOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/union/int/literal/all"
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

// PutDefault - Put a body with default properties.
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalUnionIntLiteralClientPutDefaultOptions contains the optional parameters for the OptionalUnionIntLiteralClient.PutDefault
//     method.
func (client *OptionalUnionIntLiteralClient) PutDefault(ctx context.Context, body UnionIntLiteralProperty, options *OptionalUnionIntLiteralClientPutDefaultOptions) (OptionalUnionIntLiteralClientPutDefaultResponse, error) {
	var err error
	const operationName = "OptionalUnionIntLiteralClient.PutDefault"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putDefaultCreateRequest(ctx, body, options)
	if err != nil {
		return OptionalUnionIntLiteralClientPutDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalUnionIntLiteralClientPutDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalUnionIntLiteralClientPutDefaultResponse{}, err
	}
	return OptionalUnionIntLiteralClientPutDefaultResponse{}, nil
}

// putDefaultCreateRequest creates the PutDefault request.
func (client *OptionalUnionIntLiteralClient) putDefaultCreateRequest(ctx context.Context, body UnionIntLiteralProperty, _ *OptionalUnionIntLiteralClientPutDefaultOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/union/int/literal/default"
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
