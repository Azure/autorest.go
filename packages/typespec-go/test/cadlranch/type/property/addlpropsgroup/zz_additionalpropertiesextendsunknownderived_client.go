// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package addlpropsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// AdditionalPropertiesExtendsUnknownDerivedClient contains the methods for the Type.Property.AdditionalProperties namespace.
// Don't use this type directly, use [AdditionalPropertiesClient.NewAdditionalPropertiesExtendsUnknownDerivedClient] instead.
type AdditionalPropertiesExtendsUnknownDerivedClient struct {
	internal *azcore.Client
}

// Get - Get call
//   - options - AdditionalPropertiesExtendsUnknownDerivedClientGetOptions contains the optional parameters for the AdditionalPropertiesExtendsUnknownDerivedClient.Get
//     method.
func (client *AdditionalPropertiesExtendsUnknownDerivedClient) Get(ctx context.Context, options *AdditionalPropertiesExtendsUnknownDerivedClientGetOptions) (AdditionalPropertiesExtendsUnknownDerivedClientGetResponse, error) {
	var err error
	const operationName = "AdditionalPropertiesExtendsUnknownDerivedClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return AdditionalPropertiesExtendsUnknownDerivedClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AdditionalPropertiesExtendsUnknownDerivedClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AdditionalPropertiesExtendsUnknownDerivedClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AdditionalPropertiesExtendsUnknownDerivedClient) getCreateRequest(ctx context.Context, options *AdditionalPropertiesExtendsUnknownDerivedClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/extendsRecordUnknownDerived"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AdditionalPropertiesExtendsUnknownDerivedClient) getHandleResponse(resp *http.Response) (AdditionalPropertiesExtendsUnknownDerivedClientGetResponse, error) {
	result := AdditionalPropertiesExtendsUnknownDerivedClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendsUnknownAdditionalPropertiesDerived); err != nil {
		return AdditionalPropertiesExtendsUnknownDerivedClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
//   - body - body
//   - options - AdditionalPropertiesExtendsUnknownDerivedClientPutOptions contains the optional parameters for the AdditionalPropertiesExtendsUnknownDerivedClient.Put
//     method.
func (client *AdditionalPropertiesExtendsUnknownDerivedClient) Put(ctx context.Context, body ExtendsUnknownAdditionalPropertiesDerived, options *AdditionalPropertiesExtendsUnknownDerivedClientPutOptions) (AdditionalPropertiesExtendsUnknownDerivedClientPutResponse, error) {
	var err error
	const operationName = "AdditionalPropertiesExtendsUnknownDerivedClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return AdditionalPropertiesExtendsUnknownDerivedClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AdditionalPropertiesExtendsUnknownDerivedClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AdditionalPropertiesExtendsUnknownDerivedClientPutResponse{}, err
	}
	return AdditionalPropertiesExtendsUnknownDerivedClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *AdditionalPropertiesExtendsUnknownDerivedClient) putCreateRequest(ctx context.Context, body ExtendsUnknownAdditionalPropertiesDerived, options *AdditionalPropertiesExtendsUnknownDerivedClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/extendsRecordUnknownDerived"
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