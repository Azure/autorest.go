// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package lrorpcgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// RpcClient - Illustrates bodies templated with Azure Core with long-running RPC operation
// Don't use this type directly, use a constructor function instead.
type RpcClient struct {
	internal *azcore.Client
}

// BeginLongRunningRPC - Generate data.
//   - options - RpcClientLongRunningRPCOptions contains the optional parameters for the RpcClient.LongRunningRPC method.
func (client *RpcClient) BeginLongRunningRPC(ctx context.Context, body GenerationOptions, options *RpcClientLongRunningRPCOptions) (*runtime.Poller[RpcClientLongRunningRPCResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.longRunningRPC(ctx, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RpcClientLongRunningRPCResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RpcClientLongRunningRPCResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// LongRunningRPC - Generate data.
func (client *RpcClient) longRunningRPC(ctx context.Context, body GenerationOptions, options *RpcClientLongRunningRPCOptions) (*http.Response, error) {
	var err error
	const operationName = "RpcClient.BeginLongRunningRPC"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.longRunningRPCCreateRequest(ctx, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// longRunningRPCCreateRequest creates the LongRunningRPC request.
func (client *RpcClient) longRunningRPCCreateRequest(ctx context.Context, body GenerationOptions, options *RpcClientLongRunningRPCOptions) (*policy.Request, error) {
	urlPath := "/azure/core/lro/rpc/generations:submit"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
