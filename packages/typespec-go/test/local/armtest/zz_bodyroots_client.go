// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armtest

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BodyRootsClient contains the methods for the BodyRoots group.
// Don't use this type directly, use NewBodyRootsClient() instead.
type BodyRootsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBodyRootsClient creates a new instance of BodyRootsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBodyRootsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BodyRootsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BodyRootsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Action - Revoke a certificate under a certificate profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - bodyRootName - Body root resource name.
//   - action - The content of the action request
//   - options - BodyRootsClientActionOptions contains the optional parameters for the BodyRootsClient.Action method.
func (client *BodyRootsClient) Action(ctx context.Context, resourceGroupName string, bodyRootName string, action ActionRequest, options *BodyRootsClientActionOptions) (BodyRootsClientActionResponse, error) {
	var err error
	const operationName = "BodyRootsClient.Action"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.actionCreateRequest(ctx, resourceGroupName, bodyRootName, action, options)
	if err != nil {
		return BodyRootsClientActionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BodyRootsClientActionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BodyRootsClientActionResponse{}, err
	}
	return BodyRootsClientActionResponse{}, nil
}

// actionCreateRequest creates the Action request.
func (client *BodyRootsClient) actionCreateRequest(ctx context.Context, resourceGroupName string, bodyRootName string, action ActionRequest, _ *BodyRootsClientActionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Test/bodyRoots/{bodyRootName}/action"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if bodyRootName == "" {
		return nil, errors.New("parameter bodyRootName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bodyRootName}", url.PathEscape(bodyRootName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, action); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get details of a certificate profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - bodyRootName - Body root resource name.
//   - options - BodyRootsClientGetOptions contains the optional parameters for the BodyRootsClient.Get method.
func (client *BodyRootsClient) Get(ctx context.Context, resourceGroupName string, bodyRootName string, options *BodyRootsClientGetOptions) (BodyRootsClientGetResponse, error) {
	var err error
	const operationName = "BodyRootsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, bodyRootName, options)
	if err != nil {
		return BodyRootsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BodyRootsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BodyRootsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BodyRootsClient) getCreateRequest(ctx context.Context, resourceGroupName string, bodyRootName string, _ *BodyRootsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Test/bodyRoots/{bodyRootName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if bodyRootName == "" {
		return nil, errors.New("parameter bodyRootName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bodyRootName}", url.PathEscape(bodyRootName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BodyRootsClient) getHandleResponse(resp *http.Response) (BodyRootsClientGetResponse, error) {
	result := BodyRootsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BodyRoot); err != nil {
		return BodyRootsClientGetResponse{}, err
	}
	return result, nil
}
