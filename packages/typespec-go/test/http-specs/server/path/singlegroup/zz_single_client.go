// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package singlegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// SingleClient - Illustrates server with a single path parameter @server
// Don't use this type directly, use a constructor function instead.
type SingleClient struct {
	internal *azcore.Client
	endpoint string
}

// - options - SingleClientMyOpOptions contains the optional parameters for the SingleClient.MyOp method.
func (client *SingleClient) MyOp(ctx context.Context, options *SingleClientMyOpOptions) (SingleClientMyOpResponse, error) {
	var err error
	const operationName = "SingleClient.MyOp"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.myOpCreateRequest(ctx, options)
	if err != nil {
		return SingleClientMyOpResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SingleClientMyOpResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SingleClientMyOpResponse{}, err
	}
	return SingleClientMyOpResponse{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// myOpCreateRequest creates the MyOp request.
func (client *SingleClient) myOpCreateRequest(ctx context.Context, _ *SingleClientMyOpOptions) (*policy.Request, error) {
	urlPath := "/server/path/single/myOp"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
