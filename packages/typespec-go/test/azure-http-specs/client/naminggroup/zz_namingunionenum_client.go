// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package naminggroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// NamingUnionEnumClient contains the methods for the NamingUnionEnum group.
// Don't use this type directly, use [NamingClient.NewNamingUnionEnumClient] instead.
type NamingUnionEnumClient struct {
	internal *azcore.Client
}

// UnionEnumMemberName -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - NamingUnionEnumClientUnionEnumMemberNameOptions contains the optional parameters for the NamingUnionEnumClient.UnionEnumMemberName
//     method.
func (client *NamingUnionEnumClient) UnionEnumMemberName(ctx context.Context, body ExtensibleEnum, options *NamingUnionEnumClientUnionEnumMemberNameOptions) (NamingUnionEnumClientUnionEnumMemberNameResponse, error) {
	var err error
	const operationName = "NamingUnionEnumClient.UnionEnumMemberName"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.unionEnumMemberNameCreateRequest(ctx, body, options)
	if err != nil {
		return NamingUnionEnumClientUnionEnumMemberNameResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NamingUnionEnumClientUnionEnumMemberNameResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NamingUnionEnumClientUnionEnumMemberNameResponse{}, err
	}
	return NamingUnionEnumClientUnionEnumMemberNameResponse{}, nil
}

// unionEnumMemberNameCreateRequest creates the UnionEnumMemberName request.
func (client *NamingUnionEnumClient) unionEnumMemberNameCreateRequest(ctx context.Context, body ExtensibleEnum, _ *NamingUnionEnumClientUnionEnumMemberNameOptions) (*policy.Request, error) {
	urlPath := "/client/naming/union-enum/union-enum-member-name"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// UnionEnumName -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - NamingUnionEnumClientUnionEnumNameOptions contains the optional parameters for the NamingUnionEnumClient.UnionEnumName
//     method.
func (client *NamingUnionEnumClient) UnionEnumName(ctx context.Context, body ClientExtensibleEnum, options *NamingUnionEnumClientUnionEnumNameOptions) (NamingUnionEnumClientUnionEnumNameResponse, error) {
	var err error
	const operationName = "NamingUnionEnumClient.UnionEnumName"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.unionEnumNameCreateRequest(ctx, body, options)
	if err != nil {
		return NamingUnionEnumClientUnionEnumNameResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NamingUnionEnumClientUnionEnumNameResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NamingUnionEnumClientUnionEnumNameResponse{}, err
	}
	return NamingUnionEnumClientUnionEnumNameResponse{}, nil
}

// unionEnumNameCreateRequest creates the UnionEnumName request.
func (client *NamingUnionEnumClient) unionEnumNameCreateRequest(ctx context.Context, body ClientExtensibleEnum, _ *NamingUnionEnumClientUnionEnumNameOptions) (*policy.Request, error) {
	urlPath := "/client/naming/union-enum/union-enum-name"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}