//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package customgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// CustomClient contains the methods for the Authentication.Http.Custom group.
// Don't use this type directly, use a constructor function instead.
type CustomClient struct {
	internal *azcore.Client
}

// Invalid - Check whether client is authenticated.
func (client *CustomClient) Invalid(ctx context.Context, options *CustomClientInvalidOptions) (CustomClientInvalidResponse, error) {
	var err error
	req, err := client.invalidCreateRequest(ctx, options)
	if err != nil {
		return CustomClientInvalidResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomClientInvalidResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent, http.StatusForbidden) {
		err = runtime.NewResponseError(httpResp)
		return CustomClientInvalidResponse{}, err
	}
	resp, err := client.invalidHandleResponse(httpResp)
	return resp, err
}

// invalidCreateRequest creates the Invalid request.
func (client *CustomClient) invalidCreateRequest(ctx context.Context, options *CustomClientInvalidOptions) (*policy.Request, error) {
	urlPath := "/authentication/http/custom/invalid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// invalidHandleResponse handles the Invalid response.
func (client *CustomClient) invalidHandleResponse(resp *http.Response) (CustomClientInvalidResponse, error) {
	result := CustomClientInvalidResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InvalidAuth); err != nil {
		return CustomClientInvalidResponse{}, err
	}
	return result, nil
}

// Valid - Check whether client is authenticated
func (client *CustomClient) Valid(ctx context.Context, options *CustomClientValidOptions) (CustomClientValidResponse, error) {
	var err error
	req, err := client.validCreateRequest(ctx, options)
	if err != nil {
		return CustomClientValidResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomClientValidResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return CustomClientValidResponse{}, err
	}
	return CustomClientValidResponse{}, nil
}

// validCreateRequest creates the Valid request.
func (client *CustomClient) validCreateRequest(ctx context.Context, options *CustomClientValidOptions) (*policy.Request, error) {
	urlPath := "/authentication/http/custom/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}