// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdatabasewatcher

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

// SharedPrivateLinkResourcesClient contains the methods for the SharedPrivateLinkResources group.
// Don't use this type directly, use NewSharedPrivateLinkResourcesClient() instead.
type SharedPrivateLinkResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSharedPrivateLinkResourcesClient creates a new instance of SharedPrivateLinkResourcesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSharedPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SharedPrivateLinkResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SharedPrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create a SharedPrivateLinkResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - watcherName - The database watcher name.
//   - sharedPrivateLinkResourceName - The Shared Private Link resource name.
//   - resource - Resource create parameters.
//   - options - SharedPrivateLinkResourcesClientBeginCreateOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.BeginCreate
//     method.
func (client *SharedPrivateLinkResourcesClient) BeginCreate(ctx context.Context, resourceGroupName string, watcherName string, sharedPrivateLinkResourceName string, resource SharedPrivateLinkResource, options *SharedPrivateLinkResourcesClientBeginCreateOptions) (*runtime.Poller[SharedPrivateLinkResourcesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, watcherName, sharedPrivateLinkResourceName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SharedPrivateLinkResourcesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SharedPrivateLinkResourcesClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a SharedPrivateLinkResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-19-preview
func (client *SharedPrivateLinkResourcesClient) create(ctx context.Context, resourceGroupName string, watcherName string, sharedPrivateLinkResourceName string, resource SharedPrivateLinkResource, options *SharedPrivateLinkResourcesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "SharedPrivateLinkResourcesClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, watcherName, sharedPrivateLinkResourceName, resource, options)
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

// createCreateRequest creates the Create request.
func (client *SharedPrivateLinkResourcesClient) createCreateRequest(ctx context.Context, resourceGroupName string, watcherName string, sharedPrivateLinkResourceName string, resource SharedPrivateLinkResource, _ *SharedPrivateLinkResourcesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DatabaseWatcher/watchers/{watcherName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	if sharedPrivateLinkResourceName == "" {
		return nil, errors.New("parameter sharedPrivateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sharedPrivateLinkResourceName}", url.PathEscape(sharedPrivateLinkResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a SharedPrivateLinkResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - watcherName - The database watcher name.
//   - sharedPrivateLinkResourceName - The Shared Private Link resource name.
//   - options - SharedPrivateLinkResourcesClientBeginDeleteOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.BeginDelete
//     method.
func (client *SharedPrivateLinkResourcesClient) BeginDelete(ctx context.Context, resourceGroupName string, watcherName string, sharedPrivateLinkResourceName string, options *SharedPrivateLinkResourcesClientBeginDeleteOptions) (*runtime.Poller[SharedPrivateLinkResourcesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, watcherName, sharedPrivateLinkResourceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SharedPrivateLinkResourcesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SharedPrivateLinkResourcesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a SharedPrivateLinkResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-19-preview
func (client *SharedPrivateLinkResourcesClient) deleteOperation(ctx context.Context, resourceGroupName string, watcherName string, sharedPrivateLinkResourceName string, options *SharedPrivateLinkResourcesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SharedPrivateLinkResourcesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, watcherName, sharedPrivateLinkResourceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SharedPrivateLinkResourcesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, watcherName string, sharedPrivateLinkResourceName string, _ *SharedPrivateLinkResourcesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DatabaseWatcher/watchers/{watcherName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	if sharedPrivateLinkResourceName == "" {
		return nil, errors.New("parameter sharedPrivateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sharedPrivateLinkResourceName}", url.PathEscape(sharedPrivateLinkResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a SharedPrivateLinkResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - watcherName - The database watcher name.
//   - sharedPrivateLinkResourceName - The Shared Private Link resource name.
//   - options - SharedPrivateLinkResourcesClientGetOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.Get
//     method.
func (client *SharedPrivateLinkResourcesClient) Get(ctx context.Context, resourceGroupName string, watcherName string, sharedPrivateLinkResourceName string, options *SharedPrivateLinkResourcesClientGetOptions) (SharedPrivateLinkResourcesClientGetResponse, error) {
	var err error
	const operationName = "SharedPrivateLinkResourcesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, watcherName, sharedPrivateLinkResourceName, options)
	if err != nil {
		return SharedPrivateLinkResourcesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SharedPrivateLinkResourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SharedPrivateLinkResourcesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SharedPrivateLinkResourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, watcherName string, sharedPrivateLinkResourceName string, _ *SharedPrivateLinkResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DatabaseWatcher/watchers/{watcherName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	if sharedPrivateLinkResourceName == "" {
		return nil, errors.New("parameter sharedPrivateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sharedPrivateLinkResourceName}", url.PathEscape(sharedPrivateLinkResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SharedPrivateLinkResourcesClient) getHandleResponse(resp *http.Response) (SharedPrivateLinkResourcesClientGetResponse, error) {
	result := SharedPrivateLinkResourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedPrivateLinkResource); err != nil {
		return SharedPrivateLinkResourcesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWatcherPager - List SharedPrivateLinkResource resources by Watcher
//
// Generated from API version 2024-07-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - watcherName - The database watcher name.
//   - options - SharedPrivateLinkResourcesClientListByWatcherOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.NewListByWatcherPager
//     method.
func (client *SharedPrivateLinkResourcesClient) NewListByWatcherPager(resourceGroupName string, watcherName string, options *SharedPrivateLinkResourcesClientListByWatcherOptions) *runtime.Pager[SharedPrivateLinkResourcesClientListByWatcherResponse] {
	return runtime.NewPager(runtime.PagingHandler[SharedPrivateLinkResourcesClientListByWatcherResponse]{
		More: func(page SharedPrivateLinkResourcesClientListByWatcherResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SharedPrivateLinkResourcesClientListByWatcherResponse) (SharedPrivateLinkResourcesClientListByWatcherResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SharedPrivateLinkResourcesClient.NewListByWatcherPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByWatcherCreateRequest(ctx, resourceGroupName, watcherName, options)
			}, nil)
			if err != nil {
				return SharedPrivateLinkResourcesClientListByWatcherResponse{}, err
			}
			return client.listByWatcherHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWatcherCreateRequest creates the ListByWatcher request.
func (client *SharedPrivateLinkResourcesClient) listByWatcherCreateRequest(ctx context.Context, resourceGroupName string, watcherName string, _ *SharedPrivateLinkResourcesClientListByWatcherOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DatabaseWatcher/watchers/{watcherName}/sharedPrivateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWatcherHandleResponse handles the ListByWatcher response.
func (client *SharedPrivateLinkResourcesClient) listByWatcherHandleResponse(resp *http.Response) (SharedPrivateLinkResourcesClientListByWatcherResponse, error) {
	result := SharedPrivateLinkResourcesClientListByWatcherResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedPrivateLinkResourceListResult); err != nil {
		return SharedPrivateLinkResourcesClientListByWatcherResponse{}, err
	}
	return result, nil
}
