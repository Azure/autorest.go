// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package multiplegroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MultipleClient contains the methods for the Multiple group.
// Don't use this type directly, use a constructor function instead.
type MultipleClient struct {
	internal   *azcore.Client
	endpoint   string
	apiVersion Versions
}

// NoOperationParams -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - MultipleClientNoOperationParamsOptions contains the optional parameters for the MultipleClient.NoOperationParams
//     method.
func (client *MultipleClient) NoOperationParams(ctx context.Context, options *MultipleClientNoOperationParamsOptions) (MultipleClientNoOperationParamsResponse, error) {
	var err error
	const operationName = "MultipleClient.NoOperationParams"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.noOperationParamsCreateRequest(ctx, options)
	if err != nil {
		return MultipleClientNoOperationParamsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultipleClientNoOperationParamsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MultipleClientNoOperationParamsResponse{}, err
	}
	return MultipleClientNoOperationParamsResponse{}, nil
}

// noOperationParamsCreateRequest creates the NoOperationParams request.
func (client *MultipleClient) noOperationParamsCreateRequest(ctx context.Context, _ *MultipleClientNoOperationParamsOptions) (*policy.Request, error) {
	host := "{endpoint}/server/path/multiple/{apiVersion}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{apiVersion}", string(client.apiVersion))
	req, err := runtime.NewRequest(ctx, http.MethodGet, host)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// WithOperationPathParam -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - MultipleClientWithOperationPathParamOptions contains the optional parameters for the MultipleClient.WithOperationPathParam
//     method.
func (client *MultipleClient) WithOperationPathParam(ctx context.Context, keyword string, options *MultipleClientWithOperationPathParamOptions) (MultipleClientWithOperationPathParamResponse, error) {
	var err error
	const operationName = "MultipleClient.WithOperationPathParam"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.withOperationPathParamCreateRequest(ctx, keyword, options)
	if err != nil {
		return MultipleClientWithOperationPathParamResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultipleClientWithOperationPathParamResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MultipleClientWithOperationPathParamResponse{}, err
	}
	return MultipleClientWithOperationPathParamResponse{}, nil
}

// withOperationPathParamCreateRequest creates the WithOperationPathParam request.
func (client *MultipleClient) withOperationPathParamCreateRequest(ctx context.Context, keyword string, _ *MultipleClientWithOperationPathParamOptions) (*policy.Request, error) {
	host := "{endpoint}/server/path/multiple/{apiVersion}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{apiVersion}", string(client.apiVersion))
	urlPath := "/{keyword}"
	if keyword == "" {
		return nil, errors.New("parameter keyword cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyword}", url.PathEscape(keyword))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
