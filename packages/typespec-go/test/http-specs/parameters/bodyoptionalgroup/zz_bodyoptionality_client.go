// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package bodyoptionalgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// BodyOptionalityClient - Test describing optionality of the request body.
// Don't use this type directly, use a constructor function instead.
type BodyOptionalityClient struct {
	internal *azcore.Client
	endpoint string
}

// NewBodyOptionalityOptionalExplicitClient creates a new instance of [BodyOptionalityOptionalExplicitClient].
func (client *BodyOptionalityClient) NewBodyOptionalityOptionalExplicitClient() *BodyOptionalityOptionalExplicitClient {
	return &BodyOptionalityOptionalExplicitClient{
		internal: client.internal,
		endpoint: client.endpoint,
	}
}

// RequiredExplicit -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - BodyOptionalityClientRequiredExplicitOptions contains the optional parameters for the BodyOptionalityClient.RequiredExplicit
//     method.
func (client *BodyOptionalityClient) RequiredExplicit(ctx context.Context, body BodyModel, options *BodyOptionalityClientRequiredExplicitOptions) (BodyOptionalityClientRequiredExplicitResponse, error) {
	var err error
	const operationName = "BodyOptionalityClient.RequiredExplicit"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.requiredExplicitCreateRequest(ctx, body, options)
	if err != nil {
		return BodyOptionalityClientRequiredExplicitResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BodyOptionalityClientRequiredExplicitResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BodyOptionalityClientRequiredExplicitResponse{}, err
	}
	return BodyOptionalityClientRequiredExplicitResponse{}, nil
}

// requiredExplicitCreateRequest creates the RequiredExplicit request.
func (client *BodyOptionalityClient) requiredExplicitCreateRequest(ctx context.Context, body BodyModel, _ *BodyOptionalityClientRequiredExplicitOptions) (*policy.Request, error) {
	urlPath := "/parameters/body-optionality/required-explicit"
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

// RequiredImplicit -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - BodyOptionalityClientRequiredImplicitOptions contains the optional parameters for the BodyOptionalityClient.RequiredImplicit
//     method.
func (client *BodyOptionalityClient) RequiredImplicit(ctx context.Context, name string, options *BodyOptionalityClientRequiredImplicitOptions) (BodyOptionalityClientRequiredImplicitResponse, error) {
	var err error
	const operationName = "BodyOptionalityClient.RequiredImplicit"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.requiredImplicitCreateRequest(ctx, name, options)
	if err != nil {
		return BodyOptionalityClientRequiredImplicitResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BodyOptionalityClientRequiredImplicitResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BodyOptionalityClientRequiredImplicitResponse{}, err
	}
	return BodyOptionalityClientRequiredImplicitResponse{}, nil
}

// requiredImplicitCreateRequest creates the RequiredImplicit request.
func (client *BodyOptionalityClient) requiredImplicitCreateRequest(ctx context.Context, name string, _ *BodyOptionalityClientRequiredImplicitOptions) (*policy.Request, error) {
	urlPath := "/parameters/body-optionality/required-implicit"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	body := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
