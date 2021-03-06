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

// LoadBalancerOutboundRulesClient contains the methods for the LoadBalancerOutboundRules group.
// Don't use this type directly, use NewLoadBalancerOutboundRulesClient() instead.
type LoadBalancerOutboundRulesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewLoadBalancerOutboundRulesClient creates a new instance of LoadBalancerOutboundRulesClient with the specified values.
func NewLoadBalancerOutboundRulesClient(con *armcore.Connection, subscriptionID string) *LoadBalancerOutboundRulesClient {
	return &LoadBalancerOutboundRulesClient{con: con, subscriptionID: subscriptionID}
}

// Get - Gets the specified load balancer outbound rule.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancerOutboundRulesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, outboundRuleName string, options *LoadBalancerOutboundRulesGetOptions) (LoadBalancerOutboundRulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, outboundRuleName, options)
	if err != nil {
		return LoadBalancerOutboundRulesGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LoadBalancerOutboundRulesGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return LoadBalancerOutboundRulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LoadBalancerOutboundRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, outboundRuleName string, options *LoadBalancerOutboundRulesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules/{outboundRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if outboundRuleName == "" {
		return nil, errors.New("parameter outboundRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outboundRuleName}", url.PathEscape(outboundRuleName))
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
func (client *LoadBalancerOutboundRulesClient) getHandleResponse(resp *azcore.Response) (LoadBalancerOutboundRulesGetResponse, error) {
	result := LoadBalancerOutboundRulesGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.OutboundRule); err != nil {
		return LoadBalancerOutboundRulesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *LoadBalancerOutboundRulesClient) getHandleError(resp *azcore.Response) error {
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

// List - Gets all the outbound rules in a load balancer.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancerOutboundRulesClient) List(resourceGroupName string, loadBalancerName string, options *LoadBalancerOutboundRulesListOptions) LoadBalancerOutboundRulesListPager {
	return &loadBalancerOutboundRulesListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
		},
		advancer: func(ctx context.Context, resp LoadBalancerOutboundRulesListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerOutboundRuleListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *LoadBalancerOutboundRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancerOutboundRulesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules"
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
func (client *LoadBalancerOutboundRulesClient) listHandleResponse(resp *azcore.Response) (LoadBalancerOutboundRulesListResponse, error) {
	result := LoadBalancerOutboundRulesListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.LoadBalancerOutboundRuleListResult); err != nil {
		return LoadBalancerOutboundRulesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *LoadBalancerOutboundRulesClient) listHandleError(resp *azcore.Response) error {
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
