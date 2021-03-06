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

// PrivateEndpointsClient contains the methods for the PrivateEndpoints group.
// Don't use this type directly, use NewPrivateEndpointsClient() instead.
type PrivateEndpointsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPrivateEndpointsClient creates a new instance of PrivateEndpointsClient with the specified values.
func NewPrivateEndpointsClient(con *armcore.Connection, subscriptionID string) *PrivateEndpointsClient {
	return &PrivateEndpointsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates an private endpoint in the specified resource group.
// If the operation fails it returns the *Error error type.
func (client *PrivateEndpointsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, parameters PrivateEndpoint, options *PrivateEndpointsBeginCreateOrUpdateOptions) (PrivateEndpointsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, privateEndpointName, parameters, options)
	if err != nil {
		return PrivateEndpointsCreateOrUpdatePollerResponse{}, err
	}
	result := PrivateEndpointsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("PrivateEndpointsClient.CreateOrUpdate", "azure-async-operation", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return PrivateEndpointsCreateOrUpdatePollerResponse{}, err
	}
	poller := &privateEndpointsCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (PrivateEndpointsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new PrivateEndpointsCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to PrivateEndpointsCreateOrUpdatePoller.ResumeToken().
func (client *PrivateEndpointsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (PrivateEndpointsCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("PrivateEndpointsClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return PrivateEndpointsCreateOrUpdatePollerResponse{}, err
	}
	poller := &privateEndpointsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return PrivateEndpointsCreateOrUpdatePollerResponse{}, err
	}
	result := PrivateEndpointsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (PrivateEndpointsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates an private endpoint in the specified resource group.
// If the operation fails it returns the *Error error type.
func (client *PrivateEndpointsClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, parameters PrivateEndpoint, options *PrivateEndpointsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateEndpointName, parameters, options)
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
func (client *PrivateEndpointsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, parameters PrivateEndpoint, options *PrivateEndpointsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateEndpointName == "" {
		return nil, errors.New("parameter privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
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
func (client *PrivateEndpointsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDelete - Deletes the specified private endpoint.
// If the operation fails it returns the *Error error type.
func (client *PrivateEndpointsClient) BeginDelete(ctx context.Context, resourceGroupName string, privateEndpointName string, options *PrivateEndpointsBeginDeleteOptions) (PrivateEndpointsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, privateEndpointName, options)
	if err != nil {
		return PrivateEndpointsDeletePollerResponse{}, err
	}
	result := PrivateEndpointsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("PrivateEndpointsClient.Delete", "location", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return PrivateEndpointsDeletePollerResponse{}, err
	}
	poller := &privateEndpointsDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (PrivateEndpointsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new PrivateEndpointsDeletePoller from the specified resume token.
// token - The value must come from a previous call to PrivateEndpointsDeletePoller.ResumeToken().
func (client *PrivateEndpointsClient) ResumeDelete(ctx context.Context, token string) (PrivateEndpointsDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("PrivateEndpointsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return PrivateEndpointsDeletePollerResponse{}, err
	}
	poller := &privateEndpointsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return PrivateEndpointsDeletePollerResponse{}, err
	}
	result := PrivateEndpointsDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (PrivateEndpointsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the specified private endpoint.
// If the operation fails it returns the *Error error type.
func (client *PrivateEndpointsClient) deleteOperation(ctx context.Context, resourceGroupName string, privateEndpointName string, options *PrivateEndpointsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateEndpointName, options)
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
func (client *PrivateEndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, options *PrivateEndpointsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateEndpointName == "" {
		return nil, errors.New("parameter privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
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
func (client *PrivateEndpointsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the specified private endpoint by resource group.
// If the operation fails it returns the *Error error type.
func (client *PrivateEndpointsClient) Get(ctx context.Context, resourceGroupName string, privateEndpointName string, options *PrivateEndpointsGetOptions) (PrivateEndpointsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateEndpointName, options)
	if err != nil {
		return PrivateEndpointsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PrivateEndpointsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, options *PrivateEndpointsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateEndpointName == "" {
		return nil, errors.New("parameter privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
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
func (client *PrivateEndpointsClient) getHandleResponse(resp *azcore.Response) (PrivateEndpointsGetResponse, error) {
	result := PrivateEndpointsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PrivateEndpoint); err != nil {
		return PrivateEndpointsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PrivateEndpointsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Gets all private endpoints in a resource group.
// If the operation fails it returns the *Error error type.
func (client *PrivateEndpointsClient) List(resourceGroupName string, options *PrivateEndpointsListOptions) PrivateEndpointsListPager {
	return &privateEndpointsListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp PrivateEndpointsListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PrivateEndpointListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PrivateEndpointsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *PrivateEndpointsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints"
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

// listHandleResponse handles the List response.
func (client *PrivateEndpointsClient) listHandleResponse(resp *azcore.Response) (PrivateEndpointsListResponse, error) {
	result := PrivateEndpointsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PrivateEndpointListResult); err != nil {
		return PrivateEndpointsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *PrivateEndpointsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListBySubscription - Gets all private endpoints in a subscription.
// If the operation fails it returns the *Error error type.
func (client *PrivateEndpointsClient) ListBySubscription(options *PrivateEndpointsListBySubscriptionOptions) PrivateEndpointsListBySubscriptionPager {
	return &privateEndpointsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp PrivateEndpointsListBySubscriptionResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PrivateEndpointListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *PrivateEndpointsClient) listBySubscriptionCreateRequest(ctx context.Context, options *PrivateEndpointsListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateEndpoints"
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

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *PrivateEndpointsClient) listBySubscriptionHandleResponse(resp *azcore.Response) (PrivateEndpointsListBySubscriptionResponse, error) {
	result := PrivateEndpointsListBySubscriptionResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PrivateEndpointListResult); err != nil {
		return PrivateEndpointsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *PrivateEndpointsClient) listBySubscriptionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
