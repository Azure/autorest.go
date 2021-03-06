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

// ExpressRouteCrossConnectionPeeringsClient contains the methods for the ExpressRouteCrossConnectionPeerings group.
// Don't use this type directly, use NewExpressRouteCrossConnectionPeeringsClient() instead.
type ExpressRouteCrossConnectionPeeringsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRouteCrossConnectionPeeringsClient creates a new instance of ExpressRouteCrossConnectionPeeringsClient with the specified values.
func NewExpressRouteCrossConnectionPeeringsClient(con *armcore.Connection, subscriptionID string) *ExpressRouteCrossConnectionPeeringsClient {
	return &ExpressRouteCrossConnectionPeeringsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a peering in the specified ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionPeeringsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsBeginCreateOrUpdateOptions) (ExpressRouteCrossConnectionPeeringsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, crossConnectionName, peeringName, peeringParameters, options)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsCreateOrUpdatePollerResponse{}, err
	}
	result := ExpressRouteCrossConnectionPeeringsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ExpressRouteCrossConnectionPeeringsClient.CreateOrUpdate", "azure-async-operation", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsCreateOrUpdatePollerResponse{}, err
	}
	poller := &expressRouteCrossConnectionPeeringsCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteCrossConnectionPeeringsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ExpressRouteCrossConnectionPeeringsCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCrossConnectionPeeringsCreateOrUpdatePoller.ResumeToken().
func (client *ExpressRouteCrossConnectionPeeringsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (ExpressRouteCrossConnectionPeeringsCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ExpressRouteCrossConnectionPeeringsClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsCreateOrUpdatePollerResponse{}, err
	}
	poller := &expressRouteCrossConnectionPeeringsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsCreateOrUpdatePollerResponse{}, err
	}
	result := ExpressRouteCrossConnectionPeeringsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteCrossConnectionPeeringsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a peering in the specified ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionPeeringsClient) createOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, peeringParameters, options)
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
func (client *ExpressRouteCrossConnectionPeeringsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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
	return req, req.MarshalAsJSON(peeringParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRouteCrossConnectionPeeringsClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - Deletes the specified peering from the ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionPeeringsClient) BeginDelete(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsBeginDeleteOptions) (ExpressRouteCrossConnectionPeeringsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, crossConnectionName, peeringName, options)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsDeletePollerResponse{}, err
	}
	result := ExpressRouteCrossConnectionPeeringsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ExpressRouteCrossConnectionPeeringsClient.Delete", "location", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsDeletePollerResponse{}, err
	}
	poller := &expressRouteCrossConnectionPeeringsDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteCrossConnectionPeeringsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new ExpressRouteCrossConnectionPeeringsDeletePoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCrossConnectionPeeringsDeletePoller.ResumeToken().
func (client *ExpressRouteCrossConnectionPeeringsClient) ResumeDelete(ctx context.Context, token string) (ExpressRouteCrossConnectionPeeringsDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ExpressRouteCrossConnectionPeeringsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsDeletePollerResponse{}, err
	}
	poller := &expressRouteCrossConnectionPeeringsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsDeletePollerResponse{}, err
	}
	result := ExpressRouteCrossConnectionPeeringsDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteCrossConnectionPeeringsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the specified peering from the ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionPeeringsClient) deleteOperation(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, options)
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
func (client *ExpressRouteCrossConnectionPeeringsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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
func (client *ExpressRouteCrossConnectionPeeringsClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Gets the specified peering for the ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionPeeringsClient) Get(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsGetOptions) (ExpressRouteCrossConnectionPeeringsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, options)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExpressRouteCrossConnectionPeeringsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteCrossConnectionPeeringsClient) getCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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
func (client *ExpressRouteCrossConnectionPeeringsClient) getHandleResponse(resp *azcore.Response) (ExpressRouteCrossConnectionPeeringsGetResponse, error) {
	result := ExpressRouteCrossConnectionPeeringsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionPeering); err != nil {
		return ExpressRouteCrossConnectionPeeringsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ExpressRouteCrossConnectionPeeringsClient) getHandleError(resp *azcore.Response) error {
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

// List - Gets all peerings in a specified ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionPeeringsClient) List(resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionPeeringsListOptions) ExpressRouteCrossConnectionPeeringsListPager {
	return &expressRouteCrossConnectionPeeringsListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, crossConnectionName, options)
		},
		advancer: func(ctx context.Context, resp ExpressRouteCrossConnectionPeeringsListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCrossConnectionPeeringList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ExpressRouteCrossConnectionPeeringsClient) listCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionPeeringsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
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
func (client *ExpressRouteCrossConnectionPeeringsClient) listHandleResponse(resp *azcore.Response) (ExpressRouteCrossConnectionPeeringsListResponse, error) {
	result := ExpressRouteCrossConnectionPeeringsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionPeeringList); err != nil {
		return ExpressRouteCrossConnectionPeeringsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ExpressRouteCrossConnectionPeeringsClient) listHandleError(resp *azcore.Response) error {
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
