// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// XMSClientRequestIDClient contains the methods for the XMSClientRequestID group.
// Don't use this type directly, use a constructor function instead.
type XMSClientRequestIDClient struct {
	internal *azcore.Client
	endpoint string
}

// Get - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01-preview
//   - options - XMSClientRequestIDClientGetOptions contains the optional parameters for the XMSClientRequestIDClient.Get method.
func (client *XMSClientRequestIDClient) Get(ctx context.Context, options *XMSClientRequestIDClientGetOptions) (XMSClientRequestIDClientGetResponse, error) {
	var err error
	const operationName = "XMSClientRequestIDClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return XMSClientRequestIDClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return XMSClientRequestIDClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return XMSClientRequestIDClientGetResponse{}, err
	}
	return XMSClientRequestIDClientGetResponse{}, nil
}

// getCreateRequest creates the Get request.
func (client *XMSClientRequestIDClient) getCreateRequest(ctx context.Context, _ *XMSClientRequestIDClientGetOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/overwrite/x-ms-client-request-id/method/"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ParamGet - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01-preview
//   - xmsClientRequestID - This should appear as a method parameter, use value '9C4D50EE-2D56-4CD3-8152-34347DC9F2B0'
//   - options - XMSClientRequestIDClientParamGetOptions contains the optional parameters for the XMSClientRequestIDClient.ParamGet
//     method.
func (client *XMSClientRequestIDClient) ParamGet(ctx context.Context, xmsClientRequestID string, options *XMSClientRequestIDClientParamGetOptions) (XMSClientRequestIDClientParamGetResponse, error) {
	var err error
	const operationName = "XMSClientRequestIDClient.ParamGet"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.paramGetCreateRequest(ctx, xmsClientRequestID, options)
	if err != nil {
		return XMSClientRequestIDClientParamGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return XMSClientRequestIDClientParamGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return XMSClientRequestIDClientParamGetResponse{}, err
	}
	return XMSClientRequestIDClientParamGetResponse{}, nil
}

// paramGetCreateRequest creates the ParamGet request.
func (client *XMSClientRequestIDClient) paramGetCreateRequest(ctx context.Context, xmsClientRequestID string, _ *XMSClientRequestIDClientParamGetOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/overwrite/x-ms-client-request-id/via-param/method/"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	return req, nil
}
