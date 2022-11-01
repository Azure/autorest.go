//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LoadBalancerProbesClient contains the methods for the LoadBalancerProbes group.
// Don't use this type directly, use NewLoadBalancerProbesClient() instead.
type LoadBalancerProbesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLoadBalancerProbesClient creates a new instance of LoadBalancerProbesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLoadBalancerProbesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LoadBalancerProbesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &LoadBalancerProbesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets load balancer probe.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - probeName - The name of the probe.
//   - options - LoadBalancerProbesClientGetOptions contains the optional parameters for the LoadBalancerProbesClient.Get method.
func (client *LoadBalancerProbesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, probeName string, options *LoadBalancerProbesClientGetOptions) (LoadBalancerProbesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, probeName, options)
	if err != nil {
		return LoadBalancerProbesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LoadBalancerProbesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LoadBalancerProbesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LoadBalancerProbesClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, probeName string, options *LoadBalancerProbesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes/{probeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if probeName == "" {
		return nil, errors.New("parameter probeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{probeName}", url.PathEscape(probeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoadBalancerProbesClient) getHandleResponse(resp *http.Response) (LoadBalancerProbesClientGetResponse, error) {
	result := LoadBalancerProbesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Probe); err != nil {
		return LoadBalancerProbesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all the load balancer probes.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - options - LoadBalancerProbesClientListOptions contains the optional parameters for the LoadBalancerProbesClient.List method.
func (client *LoadBalancerProbesClient) NewListPager(resourceGroupName string, loadBalancerName string, options *LoadBalancerProbesClientListOptions) *runtime.Pager[LoadBalancerProbesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LoadBalancerProbesClientListResponse]{
		More: func(page LoadBalancerProbesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LoadBalancerProbesClientListResponse) (LoadBalancerProbesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LoadBalancerProbesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LoadBalancerProbesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LoadBalancerProbesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LoadBalancerProbesClient) listCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancerProbesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LoadBalancerProbesClient) listHandleResponse(resp *http.Response) (LoadBalancerProbesClientListResponse, error) {
	result := LoadBalancerProbesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancerProbeListResult); err != nil {
		return LoadBalancerProbesClientListResponse{}, err
	}
	return result, nil
}