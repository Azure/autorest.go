// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// LoadBalancerFrontendIPConfigurationsClient contains the methods for the LoadBalancerFrontendIPConfigurations group.
// Don't use this type directly, use NewLoadBalancerFrontendIPConfigurationsClient() instead.
type LoadBalancerFrontendIPConfigurationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewLoadBalancerFrontendIPConfigurationsClient creates a new instance of LoadBalancerFrontendIPConfigurationsClient with the specified values.
func NewLoadBalancerFrontendIPConfigurationsClient(con *armcore.Connection, subscriptionID string) *LoadBalancerFrontendIPConfigurationsClient {
	return &LoadBalancerFrontendIPConfigurationsClient{con: con, subscriptionID: subscriptionID}
}

// Get - Gets load balancer frontend IP configuration.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancerFrontendIPConfigurationsClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, frontendIPConfigurationName string, options *LoadBalancerFrontendIPConfigurationsGetOptions) (LoadBalancerFrontendIPConfigurationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, frontendIPConfigurationName, options)
	if err != nil {
		return LoadBalancerFrontendIPConfigurationsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LoadBalancerFrontendIPConfigurationsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return LoadBalancerFrontendIPConfigurationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LoadBalancerFrontendIPConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, frontendIPConfigurationName string, options *LoadBalancerFrontendIPConfigurationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/frontendIPConfigurations/{frontendIPConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if frontendIPConfigurationName == "" {
		return nil, errors.New("parameter frontendIPConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontendIPConfigurationName}", url.PathEscape(frontendIPConfigurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoadBalancerFrontendIPConfigurationsClient) getHandleResponse(resp *azcore.Response) (LoadBalancerFrontendIPConfigurationsGetResponse, error) {
	result := LoadBalancerFrontendIPConfigurationsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.FrontendIPConfiguration); err != nil {
		return LoadBalancerFrontendIPConfigurationsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *LoadBalancerFrontendIPConfigurationsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Gets all the load balancer frontend IP configurations.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancerFrontendIPConfigurationsClient) List(resourceGroupName string, loadBalancerName string, options *LoadBalancerFrontendIPConfigurationsListOptions) LoadBalancerFrontendIPConfigurationsListPager {
	return &loadBalancerFrontendIPConfigurationsListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
		},
		advancer: func(ctx context.Context, resp LoadBalancerFrontendIPConfigurationsListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerFrontendIPConfigurationListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *LoadBalancerFrontendIPConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancerFrontendIPConfigurationsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/frontendIPConfigurations"
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
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LoadBalancerFrontendIPConfigurationsClient) listHandleResponse(resp *azcore.Response) (LoadBalancerFrontendIPConfigurationsListResponse, error) {
	result := LoadBalancerFrontendIPConfigurationsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.LoadBalancerFrontendIPConfigurationListResult); err != nil {
		return LoadBalancerFrontendIPConfigurationsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *LoadBalancerFrontendIPConfigurationsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
