// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package resources

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

// ExtensionsResourcesClient - The interface of extensions resources,
// it contains 4 kinds of scopes (resource, resource group, subscription and tenant)
// Don't use this type directly, use NewExtensionsResourcesClient() instead.
type ExtensionsResourcesClient struct {
	internal *arm.Client
}

// NewExtensionsResourcesClient creates a new instance of ExtensionsResourcesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExtensionsResourcesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ExtensionsResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExtensionsResourcesClient{
		internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a ExtensionsResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - extensionsResourceName - The name of the ExtensionsResource
//   - resource - Resource create parameters.
//   - options - ExtensionsResourcesClientBeginCreateOrUpdateOptions contains the optional parameters for the ExtensionsResourcesClient.BeginCreateOrUpdate
//     method.
func (client *ExtensionsResourcesClient) BeginCreateOrUpdate(ctx context.Context, resourceURI string, extensionsResourceName string, resource ExtensionsResource, options *ExtensionsResourcesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ExtensionsResourcesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceURI, extensionsResourceName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExtensionsResourcesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ExtensionsResourcesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a ExtensionsResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
func (client *ExtensionsResourcesClient) createOrUpdate(ctx context.Context, resourceURI string, extensionsResourceName string, resource ExtensionsResource, options *ExtensionsResourcesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ExtensionsResourcesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceURI, extensionsResourceName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExtensionsResourcesClient) createOrUpdateCreateRequest(ctx context.Context, resourceURI string, extensionsResourceName string, resource ExtensionsResource, _ *ExtensionsResourcesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Azure.ResourceManager.Resources/extensionsResources/{extensionsResourceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if extensionsResourceName == "" {
		return nil, errors.New("parameter extensionsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionsResourceName}", url.PathEscape(extensionsResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Delete a ExtensionsResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - extensionsResourceName - The name of the ExtensionsResource
//   - options - ExtensionsResourcesClientDeleteOptions contains the optional parameters for the ExtensionsResourcesClient.Delete
//     method.
func (client *ExtensionsResourcesClient) Delete(ctx context.Context, resourceURI string, extensionsResourceName string, options *ExtensionsResourcesClientDeleteOptions) (ExtensionsResourcesClientDeleteResponse, error) {
	var err error
	const operationName = "ExtensionsResourcesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceURI, extensionsResourceName, options)
	if err != nil {
		return ExtensionsResourcesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtensionsResourcesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ExtensionsResourcesClientDeleteResponse{}, err
	}
	return ExtensionsResourcesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExtensionsResourcesClient) deleteCreateRequest(ctx context.Context, resourceURI string, extensionsResourceName string, _ *ExtensionsResourcesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Azure.ResourceManager.Resources/extensionsResources/{extensionsResourceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if extensionsResourceName == "" {
		return nil, errors.New("parameter extensionsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionsResourceName}", url.PathEscape(extensionsResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a ExtensionsResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - extensionsResourceName - The name of the ExtensionsResource
//   - options - ExtensionsResourcesClientGetOptions contains the optional parameters for the ExtensionsResourcesClient.Get method.
func (client *ExtensionsResourcesClient) Get(ctx context.Context, resourceURI string, extensionsResourceName string, options *ExtensionsResourcesClientGetOptions) (ExtensionsResourcesClientGetResponse, error) {
	var err error
	const operationName = "ExtensionsResourcesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceURI, extensionsResourceName, options)
	if err != nil {
		return ExtensionsResourcesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtensionsResourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExtensionsResourcesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ExtensionsResourcesClient) getCreateRequest(ctx context.Context, resourceURI string, extensionsResourceName string, _ *ExtensionsResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Azure.ResourceManager.Resources/extensionsResources/{extensionsResourceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if extensionsResourceName == "" {
		return nil, errors.New("parameter extensionsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionsResourceName}", url.PathEscape(extensionsResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtensionsResourcesClient) getHandleResponse(resp *http.Response) (ExtensionsResourcesClientGetResponse, error) {
	result := ExtensionsResourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionsResource); err != nil {
		return ExtensionsResourcesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByScopePager - List ExtensionsResource resources by parent
//
// Generated from API version 2023-12-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - options - ExtensionsResourcesClientListByScopeOptions contains the optional parameters for the ExtensionsResourcesClient.NewListByScopePager
//     method.
func (client *ExtensionsResourcesClient) NewListByScopePager(resourceURI string, options *ExtensionsResourcesClientListByScopeOptions) *runtime.Pager[ExtensionsResourcesClientListByScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExtensionsResourcesClientListByScopeResponse]{
		More: func(page ExtensionsResourcesClientListByScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExtensionsResourcesClientListByScopeResponse) (ExtensionsResourcesClientListByScopeResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExtensionsResourcesClient.NewListByScopePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByScopeCreateRequest(ctx, resourceURI, options)
			}, nil)
			if err != nil {
				return ExtensionsResourcesClientListByScopeResponse{}, err
			}
			return client.listByScopeHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByScopeCreateRequest creates the ListByScope request.
func (client *ExtensionsResourcesClient) listByScopeCreateRequest(ctx context.Context, resourceURI string, _ *ExtensionsResourcesClientListByScopeOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Azure.ResourceManager.Resources/extensionsResources"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByScopeHandleResponse handles the ListByScope response.
func (client *ExtensionsResourcesClient) listByScopeHandleResponse(resp *http.Response) (ExtensionsResourcesClientListByScopeResponse, error) {
	result := ExtensionsResourcesClientListByScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionsResourceListResult); err != nil {
		return ExtensionsResourcesClientListByScopeResponse{}, err
	}
	return result, nil
}

// Update - Update a ExtensionsResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - extensionsResourceName - The name of the ExtensionsResource
//   - properties - The resource properties to be updated.
//   - options - ExtensionsResourcesClientUpdateOptions contains the optional parameters for the ExtensionsResourcesClient.Update
//     method.
func (client *ExtensionsResourcesClient) Update(ctx context.Context, resourceURI string, extensionsResourceName string, properties ExtensionsResource, options *ExtensionsResourcesClientUpdateOptions) (ExtensionsResourcesClientUpdateResponse, error) {
	var err error
	const operationName = "ExtensionsResourcesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceURI, extensionsResourceName, properties, options)
	if err != nil {
		return ExtensionsResourcesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtensionsResourcesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExtensionsResourcesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ExtensionsResourcesClient) updateCreateRequest(ctx context.Context, resourceURI string, extensionsResourceName string, properties ExtensionsResource, _ *ExtensionsResourcesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Azure.ResourceManager.Resources/extensionsResources/{extensionsResourceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if extensionsResourceName == "" {
		return nil, errors.New("parameter extensionsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionsResourceName}", url.PathEscape(extensionsResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ExtensionsResourcesClient) updateHandleResponse(resp *http.Response) (ExtensionsResourcesClientUpdateResponse, error) {
	result := ExtensionsResourcesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionsResource); err != nil {
		return ExtensionsResourcesClientUpdateResponse{}, err
	}
	return result, nil
}
