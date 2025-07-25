// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package scalargroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ScalarDecimalVerifyClient - Decimal type verification
// Don't use this type directly, use [ScalarClient.NewScalarDecimalVerifyClient] instead.
type ScalarDecimalVerifyClient struct {
	internal *azcore.Client
	endpoint string
}

// PrepareVerify -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ScalarDecimalVerifyClientPrepareVerifyOptions contains the optional parameters for the ScalarDecimalVerifyClient.PrepareVerify
//     method.
func (client *ScalarDecimalVerifyClient) PrepareVerify(ctx context.Context, options *ScalarDecimalVerifyClientPrepareVerifyOptions) (ScalarDecimalVerifyClientPrepareVerifyResponse, error) {
	var err error
	const operationName = "ScalarDecimalVerifyClient.PrepareVerify"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.prepareVerifyCreateRequest(ctx, options)
	if err != nil {
		return ScalarDecimalVerifyClientPrepareVerifyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalarDecimalVerifyClientPrepareVerifyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScalarDecimalVerifyClientPrepareVerifyResponse{}, err
	}
	resp, err := client.prepareVerifyHandleResponse(httpResp)
	return resp, err
}

// prepareVerifyCreateRequest creates the PrepareVerify request.
func (client *ScalarDecimalVerifyClient) prepareVerifyCreateRequest(ctx context.Context, _ *ScalarDecimalVerifyClientPrepareVerifyOptions) (*policy.Request, error) {
	urlPath := "/type/scalar/decimal/prepare_verify"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// prepareVerifyHandleResponse handles the PrepareVerify response.
func (client *ScalarDecimalVerifyClient) prepareVerifyHandleResponse(resp *http.Response) (ScalarDecimalVerifyClientPrepareVerifyResponse, error) {
	result := ScalarDecimalVerifyClientPrepareVerifyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Float64Array); err != nil {
		return ScalarDecimalVerifyClientPrepareVerifyResponse{}, err
	}
	return result, nil
}

// Verify -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ScalarDecimalVerifyClientVerifyOptions contains the optional parameters for the ScalarDecimalVerifyClient.Verify
//     method.
func (client *ScalarDecimalVerifyClient) Verify(ctx context.Context, body float64, options *ScalarDecimalVerifyClientVerifyOptions) (ScalarDecimalVerifyClientVerifyResponse, error) {
	var err error
	const operationName = "ScalarDecimalVerifyClient.Verify"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.verifyCreateRequest(ctx, body, options)
	if err != nil {
		return ScalarDecimalVerifyClientVerifyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalarDecimalVerifyClientVerifyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ScalarDecimalVerifyClientVerifyResponse{}, err
	}
	return ScalarDecimalVerifyClientVerifyResponse{}, nil
}

// verifyCreateRequest creates the Verify request.
func (client *ScalarDecimalVerifyClient) verifyCreateRequest(ctx context.Context, body float64, _ *ScalarDecimalVerifyClientVerifyOptions) (*policy.Request, error) {
	urlPath := "/type/scalar/decimal/verify"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
