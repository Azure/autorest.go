// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package extensiblegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ExtensibleStringClient contains the methods for the ExtensibleString group.
// Don't use this type directly, use [ExtensibleClient.NewExtensibleStringClient] instead.
type ExtensibleStringClient struct {
	internal *azcore.Client
}

// GetKnownValue -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ExtensibleStringClientGetKnownValueOptions contains the optional parameters for the ExtensibleStringClient.GetKnownValue
//     method.
func (client *ExtensibleStringClient) GetKnownValue(ctx context.Context, options *ExtensibleStringClientGetKnownValueOptions) (ExtensibleStringClientGetKnownValueResponse, error) {
	var err error
	const operationName = "ExtensibleStringClient.GetKnownValue"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getKnownValueCreateRequest(ctx, options)
	if err != nil {
		return ExtensibleStringClientGetKnownValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtensibleStringClientGetKnownValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExtensibleStringClientGetKnownValueResponse{}, err
	}
	resp, err := client.getKnownValueHandleResponse(httpResp)
	return resp, err
}

// getKnownValueCreateRequest creates the GetKnownValue request.
func (client *ExtensibleStringClient) getKnownValueCreateRequest(ctx context.Context, _ *ExtensibleStringClientGetKnownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/extensible/string/known-value"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getKnownValueHandleResponse handles the GetKnownValue response.
func (client *ExtensibleStringClient) getKnownValueHandleResponse(resp *http.Response) (ExtensibleStringClientGetKnownValueResponse, error) {
	result := ExtensibleStringClientGetKnownValueResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return ExtensibleStringClientGetKnownValueResponse{}, err
	}
	return result, nil
}

// GetUnknownValue -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ExtensibleStringClientGetUnknownValueOptions contains the optional parameters for the ExtensibleStringClient.GetUnknownValue
//     method.
func (client *ExtensibleStringClient) GetUnknownValue(ctx context.Context, options *ExtensibleStringClientGetUnknownValueOptions) (ExtensibleStringClientGetUnknownValueResponse, error) {
	var err error
	const operationName = "ExtensibleStringClient.GetUnknownValue"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getUnknownValueCreateRequest(ctx, options)
	if err != nil {
		return ExtensibleStringClientGetUnknownValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtensibleStringClientGetUnknownValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExtensibleStringClientGetUnknownValueResponse{}, err
	}
	resp, err := client.getUnknownValueHandleResponse(httpResp)
	return resp, err
}

// getUnknownValueCreateRequest creates the GetUnknownValue request.
func (client *ExtensibleStringClient) getUnknownValueCreateRequest(ctx context.Context, _ *ExtensibleStringClientGetUnknownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/extensible/string/unknown-value"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getUnknownValueHandleResponse handles the GetUnknownValue response.
func (client *ExtensibleStringClient) getUnknownValueHandleResponse(resp *http.Response) (ExtensibleStringClientGetUnknownValueResponse, error) {
	result := ExtensibleStringClientGetUnknownValueResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return ExtensibleStringClientGetUnknownValueResponse{}, err
	}
	return result, nil
}

// PutKnownValue -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ExtensibleStringClientPutKnownValueOptions contains the optional parameters for the ExtensibleStringClient.PutKnownValue
//     method.
func (client *ExtensibleStringClient) PutKnownValue(ctx context.Context, body DaysOfWeekExtensibleEnum, options *ExtensibleStringClientPutKnownValueOptions) (ExtensibleStringClientPutKnownValueResponse, error) {
	var err error
	const operationName = "ExtensibleStringClient.PutKnownValue"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putKnownValueCreateRequest(ctx, body, options)
	if err != nil {
		return ExtensibleStringClientPutKnownValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtensibleStringClientPutKnownValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ExtensibleStringClientPutKnownValueResponse{}, err
	}
	return ExtensibleStringClientPutKnownValueResponse{}, nil
}

// putKnownValueCreateRequest creates the PutKnownValue request.
func (client *ExtensibleStringClient) putKnownValueCreateRequest(ctx context.Context, body DaysOfWeekExtensibleEnum, _ *ExtensibleStringClientPutKnownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/extensible/string/known-value"
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

// PutUnknownValue -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ExtensibleStringClientPutUnknownValueOptions contains the optional parameters for the ExtensibleStringClient.PutUnknownValue
//     method.
func (client *ExtensibleStringClient) PutUnknownValue(ctx context.Context, body DaysOfWeekExtensibleEnum, options *ExtensibleStringClientPutUnknownValueOptions) (ExtensibleStringClientPutUnknownValueResponse, error) {
	var err error
	const operationName = "ExtensibleStringClient.PutUnknownValue"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putUnknownValueCreateRequest(ctx, body, options)
	if err != nil {
		return ExtensibleStringClientPutUnknownValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtensibleStringClientPutUnknownValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ExtensibleStringClientPutUnknownValueResponse{}, err
	}
	return ExtensibleStringClientPutUnknownValueResponse{}, nil
}

// putUnknownValueCreateRequest creates the PutUnknownValue request.
func (client *ExtensibleStringClient) putUnknownValueCreateRequest(ctx context.Context, body DaysOfWeekExtensibleEnum, _ *ExtensibleStringClientPutUnknownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/extensible/string/unknown-value"
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