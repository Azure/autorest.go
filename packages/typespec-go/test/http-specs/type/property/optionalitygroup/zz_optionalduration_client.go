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

// OptionalDurationClient contains the methods for the OptionalDuration group.
// Don't use this type directly, use [OptionalClient.NewOptionalDurationClient] instead.
type OptionalDurationClient struct {
	internal *azcore.Client
	endpoint string
}

// GetAll - Get models that will return all properties in the model
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalDurationClientGetAllOptions contains the optional parameters for the OptionalDurationClient.GetAll method.
func (client *OptionalDurationClient) GetAll(ctx context.Context, options *OptionalDurationClientGetAllOptions) (OptionalDurationClientGetAllResponse, error) {
	var err error
	const operationName = "OptionalDurationClient.GetAll"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAllCreateRequest(ctx, options)
	if err != nil {
		return OptionalDurationClientGetAllResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalDurationClientGetAllResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OptionalDurationClientGetAllResponse{}, err
	}
	resp, err := client.getAllHandleResponse(httpResp)
	return resp, err
}

// getAllCreateRequest creates the GetAll request.
func (client *OptionalDurationClient) getAllCreateRequest(ctx context.Context, _ *OptionalDurationClientGetAllOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/duration/all"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAllHandleResponse handles the GetAll response.
func (client *OptionalDurationClient) getAllHandleResponse(resp *http.Response) (OptionalDurationClientGetAllResponse, error) {
	result := OptionalDurationClientGetAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DurationProperty); err != nil {
		return OptionalDurationClientGetAllResponse{}, err
	}
	return result, nil
}

// GetDefault - Get models that will return the default object
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalDurationClientGetDefaultOptions contains the optional parameters for the OptionalDurationClient.GetDefault
//     method.
func (client *OptionalDurationClient) GetDefault(ctx context.Context, options *OptionalDurationClientGetDefaultOptions) (OptionalDurationClientGetDefaultResponse, error) {
	var err error
	const operationName = "OptionalDurationClient.GetDefault"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDefaultCreateRequest(ctx, options)
	if err != nil {
		return OptionalDurationClientGetDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalDurationClientGetDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OptionalDurationClientGetDefaultResponse{}, err
	}
	resp, err := client.getDefaultHandleResponse(httpResp)
	return resp, err
}

// getDefaultCreateRequest creates the GetDefault request.
func (client *OptionalDurationClient) getDefaultCreateRequest(ctx context.Context, _ *OptionalDurationClientGetDefaultOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/duration/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDefaultHandleResponse handles the GetDefault response.
func (client *OptionalDurationClient) getDefaultHandleResponse(resp *http.Response) (OptionalDurationClientGetDefaultResponse, error) {
	result := OptionalDurationClientGetDefaultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DurationProperty); err != nil {
		return OptionalDurationClientGetDefaultResponse{}, err
	}
	return result, nil
}

// PutAll - Put a body with all properties present.
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - OptionalDurationClientPutAllOptions contains the optional parameters for the OptionalDurationClient.PutAll method.
func (client *OptionalDurationClient) PutAll(ctx context.Context, body DurationProperty, options *OptionalDurationClientPutAllOptions) (OptionalDurationClientPutAllResponse, error) {
	var err error
	const operationName = "OptionalDurationClient.PutAll"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putAllCreateRequest(ctx, body, options)
	if err != nil {
		return OptionalDurationClientPutAllResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalDurationClientPutAllResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalDurationClientPutAllResponse{}, err
	}
	return OptionalDurationClientPutAllResponse{}, nil
}

// putAllCreateRequest creates the PutAll request.
func (client *OptionalDurationClient) putAllCreateRequest(ctx context.Context, body DurationProperty, _ *OptionalDurationClientPutAllOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/duration/all"
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
//   - options - OptionalDurationClientPutDefaultOptions contains the optional parameters for the OptionalDurationClient.PutDefault
//     method.
func (client *OptionalDurationClient) PutDefault(ctx context.Context, body DurationProperty, options *OptionalDurationClientPutDefaultOptions) (OptionalDurationClientPutDefaultResponse, error) {
	var err error
	const operationName = "OptionalDurationClient.PutDefault"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putDefaultCreateRequest(ctx, body, options)
	if err != nil {
		return OptionalDurationClientPutDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalDurationClientPutDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalDurationClientPutDefaultResponse{}, err
	}
	return OptionalDurationClientPutDefaultResponse{}, nil
}

// putDefaultCreateRequest creates the PutDefault request.
func (client *OptionalDurationClient) putDefaultCreateRequest(ctx context.Context, body DurationProperty, _ *OptionalDurationClientPutDefaultOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/duration/default"
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
