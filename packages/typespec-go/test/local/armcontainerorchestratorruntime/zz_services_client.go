// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcontainerorchestratorruntime

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

// ServicesClient contains the methods for the Services group.
// Don't use this type directly, use NewServicesClient() instead.
type ServicesClient struct {
	internal *arm.Client
}

// NewServicesClient creates a new instance of ServicesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServicesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ServicesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServicesClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Create a ServiceResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - serviceName - The name of the the service
//   - resource - Resource create parameters.
//   - options - ServicesClientCreateOrUpdateOptions contains the optional parameters for the ServicesClient.CreateOrUpdate method.
func (client *ServicesClient) CreateOrUpdate(ctx context.Context, resourceURI string, serviceName string, resource ServiceResource, options *ServicesClientCreateOrUpdateOptions) (ServicesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ServicesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceURI, serviceName, resource, options)
	if err != nil {
		return ServicesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ServicesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceURI string, serviceName string, resource ServiceResource, _ *ServicesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.KubernetesRuntime/services/{serviceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ServicesClient) createOrUpdateHandleResponse(resp *http.Response) (ServicesClientCreateOrUpdateResponse, error) {
	result := ServicesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResource); err != nil {
		return ServicesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a ServiceResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - serviceName - The name of the the service
//   - options - ServicesClientDeleteOptions contains the optional parameters for the ServicesClient.Delete method.
func (client *ServicesClient) Delete(ctx context.Context, resourceURI string, serviceName string, options *ServicesClientDeleteOptions) (ServicesClientDeleteResponse, error) {
	var err error
	const operationName = "ServicesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceURI, serviceName, options)
	if err != nil {
		return ServicesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServicesClientDeleteResponse{}, err
	}
	return ServicesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServicesClient) deleteCreateRequest(ctx context.Context, resourceURI string, serviceName string, _ *ServicesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.KubernetesRuntime/services/{serviceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a ServiceResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - serviceName - The name of the the service
//   - options - ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
func (client *ServicesClient) Get(ctx context.Context, resourceURI string, serviceName string, options *ServicesClientGetOptions) (ServicesClientGetResponse, error) {
	var err error
	const operationName = "ServicesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceURI, serviceName, options)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServicesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServicesClient) getCreateRequest(ctx context.Context, resourceURI string, serviceName string, _ *ServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.KubernetesRuntime/services/{serviceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServicesClient) getHandleResponse(resp *http.Response) (ServicesClientGetResponse, error) {
	result := ServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResource); err != nil {
		return ServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List ServiceResource resources by parent
//
// Generated from API version 2024-03-01
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - options - ServicesClientListOptions contains the optional parameters for the ServicesClient.NewListPager method.
func (client *ServicesClient) NewListPager(resourceURI string, options *ServicesClientListOptions) *runtime.Pager[ServicesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServicesClientListResponse]{
		More: func(page ServicesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServicesClientListResponse) (ServicesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServicesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceURI, options)
			}, nil)
			if err != nil {
				return ServicesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ServicesClient) listCreateRequest(ctx context.Context, resourceURI string, _ *ServicesClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.KubernetesRuntime/services"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServicesClient) listHandleResponse(resp *http.Response) (ServicesClientListResponse, error) {
	result := ServicesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResourceListResult); err != nil {
		return ServicesClientListResponse{}, err
	}
	return result, nil
}