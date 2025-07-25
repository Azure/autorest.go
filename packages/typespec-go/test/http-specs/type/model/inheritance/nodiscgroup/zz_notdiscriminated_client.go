// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package nodiscgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// NotDiscriminatedClient - Illustrates not-discriminated inheritance model.
// Don't use this type directly, use a constructor function instead.
type NotDiscriminatedClient struct {
	internal *azcore.Client
	endpoint string
}

// GetValid -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - NotDiscriminatedClientGetValidOptions contains the optional parameters for the NotDiscriminatedClient.GetValid
//     method.
func (client *NotDiscriminatedClient) GetValid(ctx context.Context, options *NotDiscriminatedClientGetValidOptions) (NotDiscriminatedClientGetValidResponse, error) {
	var err error
	const operationName = "NotDiscriminatedClient.GetValid"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return NotDiscriminatedClientGetValidResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotDiscriminatedClientGetValidResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NotDiscriminatedClientGetValidResponse{}, err
	}
	resp, err := client.getValidHandleResponse(httpResp)
	return resp, err
}

// getValidCreateRequest creates the GetValid request.
func (client *NotDiscriminatedClient) getValidCreateRequest(ctx context.Context, _ *NotDiscriminatedClientGetValidOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/not-discriminated/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *NotDiscriminatedClient) getValidHandleResponse(resp *http.Response) (NotDiscriminatedClientGetValidResponse, error) {
	result := NotDiscriminatedClientGetValidResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Siamese); err != nil {
		return NotDiscriminatedClientGetValidResponse{}, err
	}
	return result, nil
}

// PostValid -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - NotDiscriminatedClientPostValidOptions contains the optional parameters for the NotDiscriminatedClient.PostValid
//     method.
func (client *NotDiscriminatedClient) PostValid(ctx context.Context, input Siamese, options *NotDiscriminatedClientPostValidOptions) (NotDiscriminatedClientPostValidResponse, error) {
	var err error
	const operationName = "NotDiscriminatedClient.PostValid"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postValidCreateRequest(ctx, input, options)
	if err != nil {
		return NotDiscriminatedClientPostValidResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotDiscriminatedClientPostValidResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NotDiscriminatedClientPostValidResponse{}, err
	}
	return NotDiscriminatedClientPostValidResponse{}, nil
}

// postValidCreateRequest creates the PostValid request.
func (client *NotDiscriminatedClient) postValidCreateRequest(ctx context.Context, input Siamese, _ *NotDiscriminatedClientPostValidOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/not-discriminated/valid"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

// PutValid -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - NotDiscriminatedClientPutValidOptions contains the optional parameters for the NotDiscriminatedClient.PutValid
//     method.
func (client *NotDiscriminatedClient) PutValid(ctx context.Context, input Siamese, options *NotDiscriminatedClientPutValidOptions) (NotDiscriminatedClientPutValidResponse, error) {
	var err error
	const operationName = "NotDiscriminatedClient.PutValid"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putValidCreateRequest(ctx, input, options)
	if err != nil {
		return NotDiscriminatedClientPutValidResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotDiscriminatedClientPutValidResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NotDiscriminatedClientPutValidResponse{}, err
	}
	resp, err := client.putValidHandleResponse(httpResp)
	return resp, err
}

// putValidCreateRequest creates the PutValid request.
func (client *NotDiscriminatedClient) putValidCreateRequest(ctx context.Context, input Siamese, _ *NotDiscriminatedClientPutValidOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/not-discriminated/valid"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

// putValidHandleResponse handles the PutValid response.
func (client *NotDiscriminatedClient) putValidHandleResponse(resp *http.Response) (NotDiscriminatedClientPutValidResponse, error) {
	result := NotDiscriminatedClientPutValidResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Siamese); err != nil {
		return NotDiscriminatedClientPutValidResponse{}, err
	}
	return result, nil
}
