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

// IPAllocationsClient contains the methods for the IPAllocations group.
// Don't use this type directly, use NewIPAllocationsClient() instead.
type IPAllocationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewIPAllocationsClient creates a new instance of IPAllocationsClient with the specified values.
func NewIPAllocationsClient(con *armcore.Connection, subscriptionID string) *IPAllocationsClient {
	return &IPAllocationsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates an IpAllocation in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters IPAllocation, options *IPAllocationsBeginCreateOrUpdateOptions) (IPAllocationsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, ipAllocationName, parameters, options)
	if err != nil {
		return IPAllocationsCreateOrUpdatePollerResponse{}, err
	}
	result := IPAllocationsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("IPAllocationsClient.CreateOrUpdate", "azure-async-operation", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return IPAllocationsCreateOrUpdatePollerResponse{}, err
	}
	poller := &ipAllocationsCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (IPAllocationsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new IPAllocationsCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to IPAllocationsCreateOrUpdatePoller.ResumeToken().
func (client *IPAllocationsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (IPAllocationsCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("IPAllocationsClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return IPAllocationsCreateOrUpdatePollerResponse{}, err
	}
	poller := &ipAllocationsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return IPAllocationsCreateOrUpdatePollerResponse{}, err
	}
	result := IPAllocationsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (IPAllocationsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates an IpAllocation in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters IPAllocation, options *IPAllocationsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ipAllocationName, parameters, options)
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
func (client *IPAllocationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters IPAllocation, options *IPAllocationsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipAllocationName == "" {
		return nil, errors.New("parameter ipAllocationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
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
func (client *IPAllocationsClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - Deletes the specified IpAllocation.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) BeginDelete(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsBeginDeleteOptions) (IPAllocationsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, ipAllocationName, options)
	if err != nil {
		return IPAllocationsDeletePollerResponse{}, err
	}
	result := IPAllocationsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("IPAllocationsClient.Delete", "location", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return IPAllocationsDeletePollerResponse{}, err
	}
	poller := &ipAllocationsDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (IPAllocationsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new IPAllocationsDeletePoller from the specified resume token.
// token - The value must come from a previous call to IPAllocationsDeletePoller.ResumeToken().
func (client *IPAllocationsClient) ResumeDelete(ctx context.Context, token string) (IPAllocationsDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("IPAllocationsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return IPAllocationsDeletePollerResponse{}, err
	}
	poller := &ipAllocationsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return IPAllocationsDeletePollerResponse{}, err
	}
	result := IPAllocationsDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (IPAllocationsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the specified IpAllocation.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) deleteOperation(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ipAllocationName, options)
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
func (client *IPAllocationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipAllocationName == "" {
		return nil, errors.New("parameter ipAllocationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
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
func (client *IPAllocationsClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Gets the specified IpAllocation by resource group.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) Get(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsGetOptions) (IPAllocationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ipAllocationName, options)
	if err != nil {
		return IPAllocationsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return IPAllocationsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return IPAllocationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IPAllocationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipAllocationName == "" {
		return nil, errors.New("parameter ipAllocationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IPAllocationsClient) getHandleResponse(resp *azcore.Response) (IPAllocationsGetResponse, error) {
	result := IPAllocationsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.IPAllocation); err != nil {
		return IPAllocationsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *IPAllocationsClient) getHandleError(resp *azcore.Response) error {
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

// List - Gets all IpAllocations in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) List(options *IPAllocationsListOptions) IPAllocationsListPager {
	return &ipAllocationsListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp IPAllocationsListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.IPAllocationListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *IPAllocationsClient) listCreateRequest(ctx context.Context, options *IPAllocationsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/IpAllocations"
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
func (client *IPAllocationsClient) listHandleResponse(resp *azcore.Response) (IPAllocationsListResponse, error) {
	result := IPAllocationsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.IPAllocationListResult); err != nil {
		return IPAllocationsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *IPAllocationsClient) listHandleError(resp *azcore.Response) error {
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

// ListByResourceGroup - Gets all IpAllocations in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) ListByResourceGroup(resourceGroupName string, options *IPAllocationsListByResourceGroupOptions) IPAllocationsListByResourceGroupPager {
	return &ipAllocationsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp IPAllocationsListByResourceGroupResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.IPAllocationListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *IPAllocationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *IPAllocationsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *IPAllocationsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (IPAllocationsListByResourceGroupResponse, error) {
	result := IPAllocationsListByResourceGroupResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.IPAllocationListResult); err != nil {
		return IPAllocationsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *IPAllocationsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
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

// UpdateTags - Updates a IpAllocation tags.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) UpdateTags(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters TagsObject, options *IPAllocationsUpdateTagsOptions) (IPAllocationsUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, ipAllocationName, parameters, options)
	if err != nil {
		return IPAllocationsUpdateTagsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return IPAllocationsUpdateTagsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return IPAllocationsUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *IPAllocationsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters TagsObject, options *IPAllocationsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipAllocationName == "" {
		return nil, errors.New("parameter ipAllocationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// updateTagsHandleResponse handles the UpdateTags response.
func (client *IPAllocationsClient) updateTagsHandleResponse(resp *azcore.Response) (IPAllocationsUpdateTagsResponse, error) {
	result := IPAllocationsUpdateTagsResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.IPAllocation); err != nil {
		return IPAllocationsUpdateTagsResponse{}, err
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *IPAllocationsClient) updateTagsHandleError(resp *azcore.Response) error {
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
