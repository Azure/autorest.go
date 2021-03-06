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
	"time"
)

// FirewallPolicyRuleGroupsClient contains the methods for the FirewallPolicyRuleGroups group.
// Don't use this type directly, use NewFirewallPolicyRuleGroupsClient() instead.
type FirewallPolicyRuleGroupsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewFirewallPolicyRuleGroupsClient creates a new instance of FirewallPolicyRuleGroupsClient with the specified values.
func NewFirewallPolicyRuleGroupsClient(con *armcore.Connection, subscriptionID string) *FirewallPolicyRuleGroupsClient {
	return &FirewallPolicyRuleGroupsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates the specified FirewallPolicyRuleGroup.
// If the operation fails it returns the *CloudError error type.
func (client *FirewallPolicyRuleGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, parameters FirewallPolicyRuleGroup, options *FirewallPolicyRuleGroupsBeginCreateOrUpdateOptions) (FirewallPolicyRuleGroupsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, parameters, options)
	if err != nil {
		return FirewallPolicyRuleGroupsCreateOrUpdatePollerResponse{}, err
	}
	result := FirewallPolicyRuleGroupsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("FirewallPolicyRuleGroupsClient.CreateOrUpdate", "azure-async-operation", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return FirewallPolicyRuleGroupsCreateOrUpdatePollerResponse{}, err
	}
	poller := &firewallPolicyRuleGroupsCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (FirewallPolicyRuleGroupsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new FirewallPolicyRuleGroupsCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to FirewallPolicyRuleGroupsCreateOrUpdatePoller.ResumeToken().
func (client *FirewallPolicyRuleGroupsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (FirewallPolicyRuleGroupsCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("FirewallPolicyRuleGroupsClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return FirewallPolicyRuleGroupsCreateOrUpdatePollerResponse{}, err
	}
	poller := &firewallPolicyRuleGroupsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return FirewallPolicyRuleGroupsCreateOrUpdatePollerResponse{}, err
	}
	result := FirewallPolicyRuleGroupsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (FirewallPolicyRuleGroupsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the specified FirewallPolicyRuleGroup.
// If the operation fails it returns the *CloudError error type.
func (client *FirewallPolicyRuleGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, parameters FirewallPolicyRuleGroup, options *FirewallPolicyRuleGroupsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FirewallPolicyRuleGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, parameters FirewallPolicyRuleGroup, options *FirewallPolicyRuleGroupsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if ruleGroupName == "" {
		return nil, errors.New("parameter ruleGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *FirewallPolicyRuleGroupsClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - Deletes the specified FirewallPolicyRuleGroup.
// If the operation fails it returns the *CloudError error type.
func (client *FirewallPolicyRuleGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsBeginDeleteOptions) (FirewallPolicyRuleGroupsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, options)
	if err != nil {
		return FirewallPolicyRuleGroupsDeletePollerResponse{}, err
	}
	result := FirewallPolicyRuleGroupsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("FirewallPolicyRuleGroupsClient.Delete", "location", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return FirewallPolicyRuleGroupsDeletePollerResponse{}, err
	}
	poller := &firewallPolicyRuleGroupsDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (FirewallPolicyRuleGroupsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new FirewallPolicyRuleGroupsDeletePoller from the specified resume token.
// token - The value must come from a previous call to FirewallPolicyRuleGroupsDeletePoller.ResumeToken().
func (client *FirewallPolicyRuleGroupsClient) ResumeDelete(ctx context.Context, token string) (FirewallPolicyRuleGroupsDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("FirewallPolicyRuleGroupsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return FirewallPolicyRuleGroupsDeletePollerResponse{}, err
	}
	poller := &firewallPolicyRuleGroupsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return FirewallPolicyRuleGroupsDeletePollerResponse{}, err
	}
	result := FirewallPolicyRuleGroupsDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (FirewallPolicyRuleGroupsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the specified FirewallPolicyRuleGroup.
// If the operation fails it returns the *CloudError error type.
func (client *FirewallPolicyRuleGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FirewallPolicyRuleGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if ruleGroupName == "" {
		return nil, errors.New("parameter ruleGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// deleteHandleError handles the Delete error response.
func (client *FirewallPolicyRuleGroupsClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Gets the specified FirewallPolicyRuleGroup.
// If the operation fails it returns the *CloudError error type.
func (client *FirewallPolicyRuleGroupsClient) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsGetOptions) (FirewallPolicyRuleGroupsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, options)
	if err != nil {
		return FirewallPolicyRuleGroupsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FirewallPolicyRuleGroupsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return FirewallPolicyRuleGroupsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FirewallPolicyRuleGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if ruleGroupName == "" {
		return nil, errors.New("parameter ruleGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
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
func (client *FirewallPolicyRuleGroupsClient) getHandleResponse(resp *azcore.Response) (FirewallPolicyRuleGroupsGetResponse, error) {
	result := FirewallPolicyRuleGroupsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.FirewallPolicyRuleGroup); err != nil {
		return FirewallPolicyRuleGroupsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *FirewallPolicyRuleGroupsClient) getHandleError(resp *azcore.Response) error {
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

// List - Lists all FirewallPolicyRuleGroups in a FirewallPolicy resource.
// If the operation fails it returns the *CloudError error type.
func (client *FirewallPolicyRuleGroupsClient) List(resourceGroupName string, firewallPolicyName string, options *FirewallPolicyRuleGroupsListOptions) FirewallPolicyRuleGroupsListPager {
	return &firewallPolicyRuleGroupsListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, firewallPolicyName, options)
		},
		advancer: func(ctx context.Context, resp FirewallPolicyRuleGroupsListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.FirewallPolicyRuleGroupListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *FirewallPolicyRuleGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPolicyRuleGroupsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
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
func (client *FirewallPolicyRuleGroupsClient) listHandleResponse(resp *azcore.Response) (FirewallPolicyRuleGroupsListResponse, error) {
	result := FirewallPolicyRuleGroupsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.FirewallPolicyRuleGroupListResult); err != nil {
		return FirewallPolicyRuleGroupsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *FirewallPolicyRuleGroupsClient) listHandleError(resp *azcore.Response) error {
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
