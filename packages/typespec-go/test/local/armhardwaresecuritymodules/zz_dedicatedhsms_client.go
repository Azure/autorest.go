// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armhardwaresecuritymodules

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DedicatedHsmsClient contains the methods for the DedicatedHsms group.
// Don't use this type directly, use NewDedicatedHsmsClient() instead.
type DedicatedHsmsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDedicatedHsmsClient creates a new instance of DedicatedHsmsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDedicatedHsmsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DedicatedHsmsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DedicatedHsmsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or Update a dedicated HSM in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - name - Name of the dedicated Hsm
//   - parameters - Parameters to create or update the dedicated hsm
//   - options - DedicatedHsmsClientBeginCreateOrUpdateOptions contains the optional parameters for the DedicatedHsmsClient.BeginCreateOrUpdate
//     method.
func (client *DedicatedHsmsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsm, options *DedicatedHsmsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DedicatedHsmsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, name, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DedicatedHsmsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DedicatedHsmsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or Update a dedicated HSM in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-30-preview
func (client *DedicatedHsmsClient) createOrUpdate(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsm, options *DedicatedHsmsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DedicatedHsmsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, name, parameters, options)
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
func (client *DedicatedHsmsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsm, _ *DedicatedHsmsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified Azure Dedicated HSM.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - name - Name of the dedicated Hsm
//   - options - DedicatedHsmsClientBeginDeleteOptions contains the optional parameters for the DedicatedHsmsClient.BeginDelete
//     method.
func (client *DedicatedHsmsClient) BeginDelete(ctx context.Context, resourceGroupName string, name string, options *DedicatedHsmsClientBeginDeleteOptions) (*runtime.Poller[DedicatedHsmsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, name, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DedicatedHsmsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DedicatedHsmsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified Azure Dedicated HSM.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-30-preview
func (client *DedicatedHsmsClient) deleteOperation(ctx context.Context, resourceGroupName string, name string, options *DedicatedHsmsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DedicatedHsmsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, options)
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
func (client *DedicatedHsmsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, _ *DedicatedHsmsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified Azure dedicated HSM.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - name - Name of the dedicated Hsm
//   - options - DedicatedHsmsClientGetOptions contains the optional parameters for the DedicatedHsmsClient.Get method.
func (client *DedicatedHsmsClient) Get(ctx context.Context, resourceGroupName string, name string, options *DedicatedHsmsClientGetOptions) (DedicatedHsmsClientGetResponse, error) {
	var err error
	const operationName = "DedicatedHsmsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return DedicatedHsmsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DedicatedHsmsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DedicatedHsmsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DedicatedHsmsClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, _ *DedicatedHsmsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DedicatedHsmsClient) getHandleResponse(resp *http.Response) (DedicatedHsmsClientGetResponse, error) {
	result := DedicatedHsmsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHsm); err != nil {
		return DedicatedHsmsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - The List operation gets information about the dedicated hsms associated with the subscription
// and within the specified resource group.
//
// Generated from API version 2024-06-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - DedicatedHsmsClientListByResourceGroupOptions contains the optional parameters for the DedicatedHsmsClient.NewListByResourceGroupPager
//     method.
func (client *DedicatedHsmsClient) NewListByResourceGroupPager(resourceGroupName string, options *DedicatedHsmsClientListByResourceGroupOptions) *runtime.Pager[DedicatedHsmsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DedicatedHsmsClientListByResourceGroupResponse]{
		More: func(page DedicatedHsmsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DedicatedHsmsClientListByResourceGroupResponse) (DedicatedHsmsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DedicatedHsmsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return DedicatedHsmsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DedicatedHsmsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DedicatedHsmsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DedicatedHsmsClient) listByResourceGroupHandleResponse(resp *http.Response) (DedicatedHsmsClientListByResourceGroupResponse, error) {
	result := DedicatedHsmsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHsmListResult); err != nil {
		return DedicatedHsmsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - The List operation gets information about the dedicated HSMs associated with the subscription.
//
// Generated from API version 2024-06-30-preview
//   - options - DedicatedHsmsClientListBySubscriptionOptions contains the optional parameters for the DedicatedHsmsClient.NewListBySubscriptionPager
//     method.
func (client *DedicatedHsmsClient) NewListBySubscriptionPager(options *DedicatedHsmsClientListBySubscriptionOptions) *runtime.Pager[DedicatedHsmsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[DedicatedHsmsClientListBySubscriptionResponse]{
		More: func(page DedicatedHsmsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DedicatedHsmsClientListBySubscriptionResponse) (DedicatedHsmsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DedicatedHsmsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DedicatedHsmsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DedicatedHsmsClient) listBySubscriptionCreateRequest(ctx context.Context, options *DedicatedHsmsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DedicatedHsmsClient) listBySubscriptionHandleResponse(resp *http.Response) (DedicatedHsmsClientListBySubscriptionResponse, error) {
	result := DedicatedHsmsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHsmListResult); err != nil {
		return DedicatedHsmsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// NewListOutboundNetworkDependenciesEndpointsPager - Gets a list of egress endpoints (network endpoints of all outbound dependencies)
// in the specified dedicated hsm resource. The operation returns properties of each egress endpoint.
//
// Generated from API version 2024-06-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - name - Name of the dedicated Hsm
//   - options - DedicatedHsmsClientListOutboundNetworkDependenciesEndpointsOptions contains the optional parameters for the DedicatedHsmsClient.NewListOutboundNetworkDependenciesEndpointsPager
//     method.
func (client *DedicatedHsmsClient) NewListOutboundNetworkDependenciesEndpointsPager(resourceGroupName string, name string, options *DedicatedHsmsClientListOutboundNetworkDependenciesEndpointsOptions) *runtime.Pager[DedicatedHsmsClientListOutboundNetworkDependenciesEndpointsResponse] {
	return runtime.NewPager(runtime.PagingHandler[DedicatedHsmsClientListOutboundNetworkDependenciesEndpointsResponse]{
		More: func(page DedicatedHsmsClientListOutboundNetworkDependenciesEndpointsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DedicatedHsmsClientListOutboundNetworkDependenciesEndpointsResponse) (DedicatedHsmsClientListOutboundNetworkDependenciesEndpointsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DedicatedHsmsClient.NewListOutboundNetworkDependenciesEndpointsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listOutboundNetworkDependenciesEndpointsCreateRequest(ctx, resourceGroupName, name, options)
			}, nil)
			if err != nil {
				return DedicatedHsmsClientListOutboundNetworkDependenciesEndpointsResponse{}, err
			}
			return client.listOutboundNetworkDependenciesEndpointsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listOutboundNetworkDependenciesEndpointsCreateRequest creates the ListOutboundNetworkDependenciesEndpoints request.
func (client *DedicatedHsmsClient) listOutboundNetworkDependenciesEndpointsCreateRequest(ctx context.Context, resourceGroupName string, name string, _ *DedicatedHsmsClientListOutboundNetworkDependenciesEndpointsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}/outboundNetworkDependenciesEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listOutboundNetworkDependenciesEndpointsHandleResponse handles the ListOutboundNetworkDependenciesEndpoints response.
func (client *DedicatedHsmsClient) listOutboundNetworkDependenciesEndpointsHandleResponse(resp *http.Response) (DedicatedHsmsClientListOutboundNetworkDependenciesEndpointsResponse, error) {
	result := DedicatedHsmsClientListOutboundNetworkDependenciesEndpointsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundEnvironmentEndpointCollection); err != nil {
		return DedicatedHsmsClientListOutboundNetworkDependenciesEndpointsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a dedicated HSM in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - name - Name of the dedicated Hsm
//   - parameters - Parameters to patch the dedicated HSM
//   - options - DedicatedHsmsClientBeginUpdateOptions contains the optional parameters for the DedicatedHsmsClient.BeginUpdate
//     method.
func (client *DedicatedHsmsClient) BeginUpdate(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsmPatchParameters, options *DedicatedHsmsClientBeginUpdateOptions) (*runtime.Poller[DedicatedHsmsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, name, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DedicatedHsmsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DedicatedHsmsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a dedicated HSM in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-30-preview
func (client *DedicatedHsmsClient) update(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsmPatchParameters, options *DedicatedHsmsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DedicatedHsmsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, name, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *DedicatedHsmsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsmPatchParameters, _ *DedicatedHsmsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
