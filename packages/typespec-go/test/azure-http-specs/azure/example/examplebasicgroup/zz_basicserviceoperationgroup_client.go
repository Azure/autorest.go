// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package examplebasicgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// BasicServiceOperationGroupClient contains the methods for the BasicServiceOperationGroup group.
// Don't use this type directly, use [BasicClient.NewBasicServiceOperationGroupClient] instead.
type BasicServiceOperationGroupClient struct {
	internal *azcore.Client
	endpoint string
}

// Basic -
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - options - BasicServiceOperationGroupClientBasicOptions contains the optional parameters for the BasicServiceOperationGroupClient.Basic
//     method.
func (client *BasicServiceOperationGroupClient) Basic(ctx context.Context, queryParam string, headerParam string, body ActionRequest, options *BasicServiceOperationGroupClientBasicOptions) (BasicServiceOperationGroupClientBasicResponse, error) {
	var err error
	const operationName = "BasicServiceOperationGroupClient.Basic"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.basicCreateRequest(ctx, queryParam, headerParam, body, options)
	if err != nil {
		return BasicServiceOperationGroupClientBasicResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BasicServiceOperationGroupClientBasicResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BasicServiceOperationGroupClientBasicResponse{}, err
	}
	resp, err := client.basicHandleResponse(httpResp)
	return resp, err
}

// basicCreateRequest creates the Basic request.
func (client *BasicServiceOperationGroupClient) basicCreateRequest(ctx context.Context, queryParam string, headerParam string, body ActionRequest, _ *BasicServiceOperationGroupClientBasicOptions) (*policy.Request, error) {
	urlPath := "/azure/example/basic/basic"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	reqQP.Set("query-param", queryParam)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["header-param"] = []string{headerParam}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// basicHandleResponse handles the Basic response.
func (client *BasicServiceOperationGroupClient) basicHandleResponse(resp *http.Response) (BasicServiceOperationGroupClientBasicResponse, error) {
	result := BasicServiceOperationGroupClientBasicResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActionResponse); err != nil {
		return BasicServiceOperationGroupClientBasicResponse{}, err
	}
	return result, nil
}
