// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// NetworkVirtualAppliancesClient contains the methods for the NetworkVirtualAppliances group.
// Don't use this type directly, use NewNetworkVirtualAppliancesClient() instead.
type NetworkVirtualAppliancesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewNetworkVirtualAppliancesClient creates a new instance of NetworkVirtualAppliancesClient with the specified values.
func NewNetworkVirtualAppliancesClient(con *armcore.Connection, subscriptionID string) *NetworkVirtualAppliancesClient {
	return &NetworkVirtualAppliancesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates the specified Network Virtual Appliance.
func (client *NetworkVirtualAppliancesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters NetworkVirtualAppliance, options *NetworkVirtualAppliancesBeginCreateOrUpdateOptions) (NetworkVirtualAppliancePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, networkVirtualApplianceName, parameters, options)
	if err != nil {
		return NetworkVirtualAppliancePollerResponse{}, err
	}
	result := NetworkVirtualAppliancePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkVirtualAppliancesClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return NetworkVirtualAppliancePollerResponse{}, err
	}
	poller := &networkVirtualAppliancePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (NetworkVirtualApplianceResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new NetworkVirtualAppliancePoller from the specified resume token.
// token - The value must come from a previous call to NetworkVirtualAppliancePoller.ResumeToken().
func (client *NetworkVirtualAppliancesClient) ResumeCreateOrUpdate(ctx context.Context, token string) (NetworkVirtualAppliancePollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkVirtualAppliancesClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return NetworkVirtualAppliancePollerResponse{}, err
	}
	poller := &networkVirtualAppliancePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return NetworkVirtualAppliancePollerResponse{}, err
	}
	result := NetworkVirtualAppliancePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (NetworkVirtualApplianceResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the specified Network Virtual Appliance.
func (client *NetworkVirtualAppliancesClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters NetworkVirtualAppliance, options *NetworkVirtualAppliancesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, parameters, options)
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
func (client *NetworkVirtualAppliancesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters NetworkVirtualAppliance, options *NetworkVirtualAppliancesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
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

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *NetworkVirtualAppliancesClient) createOrUpdateHandleResponse(resp *azcore.Response) (NetworkVirtualApplianceResponse, error) {
	var val *NetworkVirtualAppliance
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NetworkVirtualApplianceResponse{}, err
	}
	return NetworkVirtualApplianceResponse{RawResponse: resp.Response, NetworkVirtualAppliance: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *NetworkVirtualAppliancesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified Network Virtual Appliance.
func (client *NetworkVirtualAppliancesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *NetworkVirtualAppliancesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, networkVirtualApplianceName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkVirtualAppliancesClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *NetworkVirtualAppliancesClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkVirtualAppliancesClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the specified Network Virtual Appliance.
func (client *NetworkVirtualAppliancesClient) deleteOperation(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *NetworkVirtualAppliancesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, options)
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
func (client *NetworkVirtualAppliancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *NetworkVirtualAppliancesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
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
func (client *NetworkVirtualAppliancesClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified Network Virtual Appliance.
func (client *NetworkVirtualAppliancesClient) Get(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *NetworkVirtualAppliancesGetOptions) (NetworkVirtualApplianceResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, options)
	if err != nil {
		return NetworkVirtualApplianceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return NetworkVirtualApplianceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return NetworkVirtualApplianceResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NetworkVirtualAppliancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *NetworkVirtualAppliancesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
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
func (client *NetworkVirtualAppliancesClient) getHandleResponse(resp *azcore.Response) (NetworkVirtualApplianceResponse, error) {
	var val *NetworkVirtualAppliance
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NetworkVirtualApplianceResponse{}, err
	}
	return NetworkVirtualApplianceResponse{RawResponse: resp.Response, NetworkVirtualAppliance: val}, nil
}

// getHandleError handles the Get error response.
func (client *NetworkVirtualAppliancesClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all Network Virtual Appliances in a subscription.
func (client *NetworkVirtualAppliancesClient) List(options *NetworkVirtualAppliancesListOptions) NetworkVirtualApplianceListResultPager {
	return &networkVirtualApplianceListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp NetworkVirtualApplianceListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkVirtualApplianceListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *NetworkVirtualAppliancesClient) listCreateRequest(ctx context.Context, options *NetworkVirtualAppliancesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkVirtualAppliances"
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
func (client *NetworkVirtualAppliancesClient) listHandleResponse(resp *azcore.Response) (NetworkVirtualApplianceListResultResponse, error) {
	var val *NetworkVirtualApplianceListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NetworkVirtualApplianceListResultResponse{}, err
	}
	return NetworkVirtualApplianceListResultResponse{RawResponse: resp.Response, NetworkVirtualApplianceListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *NetworkVirtualAppliancesClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Lists all Network Virtual Appliances in a resource group.
func (client *NetworkVirtualAppliancesClient) ListByResourceGroup(resourceGroupName string, options *NetworkVirtualAppliancesListByResourceGroupOptions) NetworkVirtualApplianceListResultPager {
	return &networkVirtualApplianceListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp NetworkVirtualApplianceListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkVirtualApplianceListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *NetworkVirtualAppliancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *NetworkVirtualAppliancesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances"
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
func (client *NetworkVirtualAppliancesClient) listByResourceGroupHandleResponse(resp *azcore.Response) (NetworkVirtualApplianceListResultResponse, error) {
	var val *NetworkVirtualApplianceListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NetworkVirtualApplianceListResultResponse{}, err
	}
	return NetworkVirtualApplianceListResultResponse{RawResponse: resp.Response, NetworkVirtualApplianceListResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *NetworkVirtualAppliancesClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates a Network Virtual Appliance.
func (client *NetworkVirtualAppliancesClient) UpdateTags(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters TagsObject, options *NetworkVirtualAppliancesUpdateTagsOptions) (NetworkVirtualApplianceResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, parameters, options)
	if err != nil {
		return NetworkVirtualApplianceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return NetworkVirtualApplianceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return NetworkVirtualApplianceResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *NetworkVirtualAppliancesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters TagsObject, options *NetworkVirtualAppliancesUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
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
func (client *NetworkVirtualAppliancesClient) updateTagsHandleResponse(resp *azcore.Response) (NetworkVirtualApplianceResponse, error) {
	var val *NetworkVirtualAppliance
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NetworkVirtualApplianceResponse{}, err
	}
	return NetworkVirtualApplianceResponse{RawResponse: resp.Response, NetworkVirtualAppliance: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *NetworkVirtualAppliancesClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}