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

// OptionalFloatLiteralClient contains the methods for the OptionalFloatLiteral group.
// Don't use this type directly, use [OptionalClient.NewOptionalFloatLiteralClient] instead.
type OptionalFloatLiteralClient struct {
	internal *azcore.Client
}

// GetAll - Get models that will return all properties in the model
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalFloatLiteralClientGetAllOptions contains the optional parameters for the OptionalFloatLiteralClient.GetAll
//     method.
func (client *OptionalFloatLiteralClient) GetAll(ctx context.Context, options *OptionalFloatLiteralClientGetAllOptions) (OptionalFloatLiteralClientGetAllResponse, error) {
	var err error
	const operationName = "OptionalFloatLiteralClient.GetAll"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAllCreateRequest(ctx, options)
	if err != nil {
		return OptionalFloatLiteralClientGetAllResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalFloatLiteralClientGetAllResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OptionalFloatLiteralClientGetAllResponse{}, err
	}
	resp, err := client.getAllHandleResponse(httpResp)
	return resp, err
}

// getAllCreateRequest creates the GetAll request.
func (client *OptionalFloatLiteralClient) getAllCreateRequest(ctx context.Context, _ *OptionalFloatLiteralClientGetAllOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/float/literal/all"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAllHandleResponse handles the GetAll response.
func (client *OptionalFloatLiteralClient) getAllHandleResponse(resp *http.Response) (OptionalFloatLiteralClientGetAllResponse, error) {
	result := OptionalFloatLiteralClientGetAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FloatLiteralProperty); err != nil {
		return OptionalFloatLiteralClientGetAllResponse{}, err
	}
	return result, nil
}

// GetDefault - Get models that will return the default object
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalFloatLiteralClientGetDefaultOptions contains the optional parameters for the OptionalFloatLiteralClient.GetDefault
//     method.
func (client *OptionalFloatLiteralClient) GetDefault(ctx context.Context, options *OptionalFloatLiteralClientGetDefaultOptions) (OptionalFloatLiteralClientGetDefaultResponse, error) {
	var err error
	const operationName = "OptionalFloatLiteralClient.GetDefault"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDefaultCreateRequest(ctx, options)
	if err != nil {
		return OptionalFloatLiteralClientGetDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalFloatLiteralClientGetDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OptionalFloatLiteralClientGetDefaultResponse{}, err
	}
	resp, err := client.getDefaultHandleResponse(httpResp)
	return resp, err
}

// getDefaultCreateRequest creates the GetDefault request.
func (client *OptionalFloatLiteralClient) getDefaultCreateRequest(ctx context.Context, _ *OptionalFloatLiteralClientGetDefaultOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/float/literal/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDefaultHandleResponse handles the GetDefault response.
func (client *OptionalFloatLiteralClient) getDefaultHandleResponse(resp *http.Response) (OptionalFloatLiteralClientGetDefaultResponse, error) {
	result := OptionalFloatLiteralClientGetDefaultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FloatLiteralProperty); err != nil {
		return OptionalFloatLiteralClientGetDefaultResponse{}, err
	}
	return result, nil
}

// PutAll - Put a body with all properties present.
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalFloatLiteralClientPutAllOptions contains the optional parameters for the OptionalFloatLiteralClient.PutAll
//     method.
func (client *OptionalFloatLiteralClient) PutAll(ctx context.Context, body FloatLiteralProperty, options *OptionalFloatLiteralClientPutAllOptions) (OptionalFloatLiteralClientPutAllResponse, error) {
	var err error
	const operationName = "OptionalFloatLiteralClient.PutAll"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putAllCreateRequest(ctx, body, options)
	if err != nil {
		return OptionalFloatLiteralClientPutAllResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalFloatLiteralClientPutAllResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalFloatLiteralClientPutAllResponse{}, err
	}
	return OptionalFloatLiteralClientPutAllResponse{}, nil
}

// putAllCreateRequest creates the PutAll request.
func (client *OptionalFloatLiteralClient) putAllCreateRequest(ctx context.Context, body FloatLiteralProperty, _ *OptionalFloatLiteralClientPutAllOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/float/literal/all"
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

// PutDefault - Put a body with default properties.
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalFloatLiteralClientPutDefaultOptions contains the optional parameters for the OptionalFloatLiteralClient.PutDefault
//     method.
func (client *OptionalFloatLiteralClient) PutDefault(ctx context.Context, body FloatLiteralProperty, options *OptionalFloatLiteralClientPutDefaultOptions) (OptionalFloatLiteralClientPutDefaultResponse, error) {
	var err error
	const operationName = "OptionalFloatLiteralClient.PutDefault"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putDefaultCreateRequest(ctx, body, options)
	if err != nil {
		return OptionalFloatLiteralClientPutDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalFloatLiteralClientPutDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalFloatLiteralClientPutDefaultResponse{}, err
	}
	return OptionalFloatLiteralClientPutDefaultResponse{}, nil
}

// putDefaultCreateRequest creates the PutDefault request.
func (client *OptionalFloatLiteralClient) putDefaultCreateRequest(ctx context.Context, body FloatLiteralProperty, _ *OptionalFloatLiteralClientPutDefaultOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/float/literal/default"
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