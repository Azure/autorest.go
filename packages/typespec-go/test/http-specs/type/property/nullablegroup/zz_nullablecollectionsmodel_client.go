// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package nullablegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// NullableCollectionsModelClient contains the methods for the NullableCollectionsModel group.
// Don't use this type directly, use [NullableClient.NewNullableCollectionsModelClient] instead.
type NullableCollectionsModelClient struct {
	internal *azcore.Client
	endpoint string
}

// GetNonNull - Get models that will return all properties in the model
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - NullableCollectionsModelClientGetNonNullOptions contains the optional parameters for the NullableCollectionsModelClient.GetNonNull
//     method.
func (client *NullableCollectionsModelClient) GetNonNull(ctx context.Context, options *NullableCollectionsModelClientGetNonNullOptions) (NullableCollectionsModelClientGetNonNullResponse, error) {
	var err error
	const operationName = "NullableCollectionsModelClient.GetNonNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getNonNullCreateRequest(ctx, options)
	if err != nil {
		return NullableCollectionsModelClientGetNonNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NullableCollectionsModelClientGetNonNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NullableCollectionsModelClientGetNonNullResponse{}, err
	}
	resp, err := client.getNonNullHandleResponse(httpResp)
	return resp, err
}

// getNonNullCreateRequest creates the GetNonNull request.
func (client *NullableCollectionsModelClient) getNonNullCreateRequest(ctx context.Context, _ *NullableCollectionsModelClientGetNonNullOptions) (*policy.Request, error) {
	urlPath := "/type/property/nullable/collections/model/non-null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNonNullHandleResponse handles the GetNonNull response.
func (client *NullableCollectionsModelClient) getNonNullHandleResponse(resp *http.Response) (NullableCollectionsModelClientGetNonNullResponse, error) {
	result := NullableCollectionsModelClientGetNonNullResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CollectionsModelProperty); err != nil {
		return NullableCollectionsModelClientGetNonNullResponse{}, err
	}
	return result, nil
}

// GetNull - Get models that will return the default object
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - NullableCollectionsModelClientGetNullOptions contains the optional parameters for the NullableCollectionsModelClient.GetNull
//     method.
func (client *NullableCollectionsModelClient) GetNull(ctx context.Context, options *NullableCollectionsModelClientGetNullOptions) (NullableCollectionsModelClientGetNullResponse, error) {
	var err error
	const operationName = "NullableCollectionsModelClient.GetNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return NullableCollectionsModelClientGetNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NullableCollectionsModelClientGetNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NullableCollectionsModelClientGetNullResponse{}, err
	}
	resp, err := client.getNullHandleResponse(httpResp)
	return resp, err
}

// getNullCreateRequest creates the GetNull request.
func (client *NullableCollectionsModelClient) getNullCreateRequest(ctx context.Context, _ *NullableCollectionsModelClientGetNullOptions) (*policy.Request, error) {
	urlPath := "/type/property/nullable/collections/model/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *NullableCollectionsModelClient) getNullHandleResponse(resp *http.Response) (NullableCollectionsModelClientGetNullResponse, error) {
	result := NullableCollectionsModelClientGetNullResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CollectionsModelProperty); err != nil {
		return NullableCollectionsModelClientGetNullResponse{}, err
	}
	return result, nil
}

// PatchNonNull - Put a body with all properties present.
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - NullableCollectionsModelClientPatchNonNullOptions contains the optional parameters for the NullableCollectionsModelClient.PatchNonNull
//     method.
func (client *NullableCollectionsModelClient) PatchNonNull(ctx context.Context, body CollectionsModelProperty, options *NullableCollectionsModelClientPatchNonNullOptions) (NullableCollectionsModelClientPatchNonNullResponse, error) {
	var err error
	const operationName = "NullableCollectionsModelClient.PatchNonNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.patchNonNullCreateRequest(ctx, body, options)
	if err != nil {
		return NullableCollectionsModelClientPatchNonNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NullableCollectionsModelClientPatchNonNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NullableCollectionsModelClientPatchNonNullResponse{}, err
	}
	return NullableCollectionsModelClientPatchNonNullResponse{}, nil
}

// patchNonNullCreateRequest creates the PatchNonNull request.
func (client *NullableCollectionsModelClient) patchNonNullCreateRequest(ctx context.Context, body CollectionsModelProperty, _ *NullableCollectionsModelClientPatchNonNullOptions) (*policy.Request, error) {
	urlPath := "/type/property/nullable/collections/model/non-null"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/merge-patch+json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// PatchNull - Put a body with default properties.
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - NullableCollectionsModelClientPatchNullOptions contains the optional parameters for the NullableCollectionsModelClient.PatchNull
//     method.
func (client *NullableCollectionsModelClient) PatchNull(ctx context.Context, body CollectionsModelProperty, options *NullableCollectionsModelClientPatchNullOptions) (NullableCollectionsModelClientPatchNullResponse, error) {
	var err error
	const operationName = "NullableCollectionsModelClient.PatchNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.patchNullCreateRequest(ctx, body, options)
	if err != nil {
		return NullableCollectionsModelClientPatchNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NullableCollectionsModelClientPatchNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NullableCollectionsModelClientPatchNullResponse{}, err
	}
	return NullableCollectionsModelClientPatchNullResponse{}, nil
}

// patchNullCreateRequest creates the PatchNull request.
func (client *NullableCollectionsModelClient) patchNullCreateRequest(ctx context.Context, body CollectionsModelProperty, _ *NullableCollectionsModelClientPatchNullOptions) (*policy.Request, error) {
	urlPath := "/type/property/nullable/collections/model/null"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/merge-patch+json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
