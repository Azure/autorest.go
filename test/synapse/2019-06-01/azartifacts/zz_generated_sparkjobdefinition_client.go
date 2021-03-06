// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type sparkJobDefinitionClient struct {
	con *connection
}

// BeginCreateOrUpdateSparkJobDefinition - Creates or updates a Spark Job Definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) BeginCreateOrUpdateSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, sparkJobDefinition SparkJobDefinitionResource, options *SparkJobDefinitionBeginCreateOrUpdateSparkJobDefinitionOptions) (SparkJobDefinitionCreateOrUpdateSparkJobDefinitionPollerResponse, error) {
	resp, err := client.createOrUpdateSparkJobDefinition(ctx, sparkJobDefinitionName, sparkJobDefinition, options)
	if err != nil {
		return SparkJobDefinitionCreateOrUpdateSparkJobDefinitionPollerResponse{}, err
	}
	result := SparkJobDefinitionCreateOrUpdateSparkJobDefinitionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("sparkJobDefinitionClient.CreateOrUpdateSparkJobDefinition", resp, client.con.Pipeline(), client.createOrUpdateSparkJobDefinitionHandleError)
	if err != nil {
		return SparkJobDefinitionCreateOrUpdateSparkJobDefinitionPollerResponse{}, err
	}
	poller := &sparkJobDefinitionCreateOrUpdateSparkJobDefinitionPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SparkJobDefinitionCreateOrUpdateSparkJobDefinitionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdateSparkJobDefinition creates a new SparkJobDefinitionCreateOrUpdateSparkJobDefinitionPoller from the specified resume token.
// token - The value must come from a previous call to SparkJobDefinitionCreateOrUpdateSparkJobDefinitionPoller.ResumeToken().
func (client *sparkJobDefinitionClient) ResumeCreateOrUpdateSparkJobDefinition(ctx context.Context, token string) (SparkJobDefinitionCreateOrUpdateSparkJobDefinitionPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("sparkJobDefinitionClient.CreateOrUpdateSparkJobDefinition", token, client.con.Pipeline(), client.createOrUpdateSparkJobDefinitionHandleError)
	if err != nil {
		return SparkJobDefinitionCreateOrUpdateSparkJobDefinitionPollerResponse{}, err
	}
	poller := &sparkJobDefinitionCreateOrUpdateSparkJobDefinitionPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SparkJobDefinitionCreateOrUpdateSparkJobDefinitionPollerResponse{}, err
	}
	result := SparkJobDefinitionCreateOrUpdateSparkJobDefinitionPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SparkJobDefinitionCreateOrUpdateSparkJobDefinitionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdateSparkJobDefinition - Creates or updates a Spark Job Definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) createOrUpdateSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, sparkJobDefinition SparkJobDefinitionResource, options *SparkJobDefinitionBeginCreateOrUpdateSparkJobDefinitionOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, sparkJobDefinition, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateSparkJobDefinitionCreateRequest creates the CreateOrUpdateSparkJobDefinition request.
func (client *sparkJobDefinitionClient) createOrUpdateSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, sparkJobDefinition SparkJobDefinitionResource, options *SparkJobDefinitionBeginCreateOrUpdateSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}"
	if sparkJobDefinitionName == "" {
		return nil, errors.New("parameter sparkJobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(sparkJobDefinition)
}

// createOrUpdateSparkJobDefinitionHandleError handles the CreateOrUpdateSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) createOrUpdateSparkJobDefinitionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDebugSparkJobDefinition - Debug the spark job definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) BeginDebugSparkJobDefinition(ctx context.Context, sparkJobDefinitionAzureResource SparkJobDefinitionResource, options *SparkJobDefinitionBeginDebugSparkJobDefinitionOptions) (SparkJobDefinitionDebugSparkJobDefinitionPollerResponse, error) {
	resp, err := client.debugSparkJobDefinition(ctx, sparkJobDefinitionAzureResource, options)
	if err != nil {
		return SparkJobDefinitionDebugSparkJobDefinitionPollerResponse{}, err
	}
	result := SparkJobDefinitionDebugSparkJobDefinitionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("sparkJobDefinitionClient.DebugSparkJobDefinition", resp, client.con.Pipeline(), client.debugSparkJobDefinitionHandleError)
	if err != nil {
		return SparkJobDefinitionDebugSparkJobDefinitionPollerResponse{}, err
	}
	poller := &sparkJobDefinitionDebugSparkJobDefinitionPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SparkJobDefinitionDebugSparkJobDefinitionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDebugSparkJobDefinition creates a new SparkJobDefinitionDebugSparkJobDefinitionPoller from the specified resume token.
// token - The value must come from a previous call to SparkJobDefinitionDebugSparkJobDefinitionPoller.ResumeToken().
func (client *sparkJobDefinitionClient) ResumeDebugSparkJobDefinition(ctx context.Context, token string) (SparkJobDefinitionDebugSparkJobDefinitionPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("sparkJobDefinitionClient.DebugSparkJobDefinition", token, client.con.Pipeline(), client.debugSparkJobDefinitionHandleError)
	if err != nil {
		return SparkJobDefinitionDebugSparkJobDefinitionPollerResponse{}, err
	}
	poller := &sparkJobDefinitionDebugSparkJobDefinitionPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SparkJobDefinitionDebugSparkJobDefinitionPollerResponse{}, err
	}
	result := SparkJobDefinitionDebugSparkJobDefinitionPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SparkJobDefinitionDebugSparkJobDefinitionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// DebugSparkJobDefinition - Debug the spark job definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) debugSparkJobDefinition(ctx context.Context, sparkJobDefinitionAzureResource SparkJobDefinitionResource, options *SparkJobDefinitionBeginDebugSparkJobDefinitionOptions) (*azcore.Response, error) {
	req, err := client.debugSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionAzureResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.debugSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// debugSparkJobDefinitionCreateRequest creates the DebugSparkJobDefinition request.
func (client *sparkJobDefinitionClient) debugSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionAzureResource SparkJobDefinitionResource, options *SparkJobDefinitionBeginDebugSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/debugSparkJobDefinition"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(sparkJobDefinitionAzureResource)
}

// debugSparkJobDefinitionHandleError handles the DebugSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) debugSparkJobDefinitionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDeleteSparkJobDefinition - Deletes a Spark Job Definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) BeginDeleteSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginDeleteSparkJobDefinitionOptions) (SparkJobDefinitionDeleteSparkJobDefinitionPollerResponse, error) {
	resp, err := client.deleteSparkJobDefinition(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return SparkJobDefinitionDeleteSparkJobDefinitionPollerResponse{}, err
	}
	result := SparkJobDefinitionDeleteSparkJobDefinitionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("sparkJobDefinitionClient.DeleteSparkJobDefinition", resp, client.con.Pipeline(), client.deleteSparkJobDefinitionHandleError)
	if err != nil {
		return SparkJobDefinitionDeleteSparkJobDefinitionPollerResponse{}, err
	}
	poller := &sparkJobDefinitionDeleteSparkJobDefinitionPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SparkJobDefinitionDeleteSparkJobDefinitionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDeleteSparkJobDefinition creates a new SparkJobDefinitionDeleteSparkJobDefinitionPoller from the specified resume token.
// token - The value must come from a previous call to SparkJobDefinitionDeleteSparkJobDefinitionPoller.ResumeToken().
func (client *sparkJobDefinitionClient) ResumeDeleteSparkJobDefinition(ctx context.Context, token string) (SparkJobDefinitionDeleteSparkJobDefinitionPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("sparkJobDefinitionClient.DeleteSparkJobDefinition", token, client.con.Pipeline(), client.deleteSparkJobDefinitionHandleError)
	if err != nil {
		return SparkJobDefinitionDeleteSparkJobDefinitionPollerResponse{}, err
	}
	poller := &sparkJobDefinitionDeleteSparkJobDefinitionPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SparkJobDefinitionDeleteSparkJobDefinitionPollerResponse{}, err
	}
	result := SparkJobDefinitionDeleteSparkJobDefinitionPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SparkJobDefinitionDeleteSparkJobDefinitionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// DeleteSparkJobDefinition - Deletes a Spark Job Definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) deleteSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginDeleteSparkJobDefinitionOptions) (*azcore.Response, error) {
	req, err := client.deleteSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// deleteSparkJobDefinitionCreateRequest creates the DeleteSparkJobDefinition request.
func (client *sparkJobDefinitionClient) deleteSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginDeleteSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}"
	if sparkJobDefinitionName == "" {
		return nil, errors.New("parameter sparkJobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteSparkJobDefinitionHandleError handles the DeleteSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) deleteSparkJobDefinitionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginExecuteSparkJobDefinition - Executes the spark job definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) BeginExecuteSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginExecuteSparkJobDefinitionOptions) (SparkJobDefinitionExecuteSparkJobDefinitionPollerResponse, error) {
	resp, err := client.executeSparkJobDefinition(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return SparkJobDefinitionExecuteSparkJobDefinitionPollerResponse{}, err
	}
	result := SparkJobDefinitionExecuteSparkJobDefinitionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("sparkJobDefinitionClient.ExecuteSparkJobDefinition", resp, client.con.Pipeline(), client.executeSparkJobDefinitionHandleError)
	if err != nil {
		return SparkJobDefinitionExecuteSparkJobDefinitionPollerResponse{}, err
	}
	poller := &sparkJobDefinitionExecuteSparkJobDefinitionPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SparkJobDefinitionExecuteSparkJobDefinitionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeExecuteSparkJobDefinition creates a new SparkJobDefinitionExecuteSparkJobDefinitionPoller from the specified resume token.
// token - The value must come from a previous call to SparkJobDefinitionExecuteSparkJobDefinitionPoller.ResumeToken().
func (client *sparkJobDefinitionClient) ResumeExecuteSparkJobDefinition(ctx context.Context, token string) (SparkJobDefinitionExecuteSparkJobDefinitionPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("sparkJobDefinitionClient.ExecuteSparkJobDefinition", token, client.con.Pipeline(), client.executeSparkJobDefinitionHandleError)
	if err != nil {
		return SparkJobDefinitionExecuteSparkJobDefinitionPollerResponse{}, err
	}
	poller := &sparkJobDefinitionExecuteSparkJobDefinitionPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SparkJobDefinitionExecuteSparkJobDefinitionPollerResponse{}, err
	}
	result := SparkJobDefinitionExecuteSparkJobDefinitionPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SparkJobDefinitionExecuteSparkJobDefinitionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ExecuteSparkJobDefinition - Executes the spark job definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) executeSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginExecuteSparkJobDefinitionOptions) (*azcore.Response, error) {
	req, err := client.executeSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.executeSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// executeSparkJobDefinitionCreateRequest creates the ExecuteSparkJobDefinition request.
func (client *sparkJobDefinitionClient) executeSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginExecuteSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}/execute"
	if sparkJobDefinitionName == "" {
		return nil, errors.New("parameter sparkJobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// executeSparkJobDefinitionHandleError handles the ExecuteSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) executeSparkJobDefinitionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetSparkJobDefinition - Gets a Spark Job Definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) GetSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionGetSparkJobDefinitionOptions) (SparkJobDefinitionGetSparkJobDefinitionResponse, error) {
	req, err := client.getSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return SparkJobDefinitionGetSparkJobDefinitionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SparkJobDefinitionGetSparkJobDefinitionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return SparkJobDefinitionGetSparkJobDefinitionResponse{}, client.getSparkJobDefinitionHandleError(resp)
	}
	return client.getSparkJobDefinitionHandleResponse(resp)
}

// getSparkJobDefinitionCreateRequest creates the GetSparkJobDefinition request.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionGetSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}"
	if sparkJobDefinitionName == "" {
		return nil, errors.New("parameter sparkJobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkJobDefinitionHandleResponse handles the GetSparkJobDefinition response.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionHandleResponse(resp *azcore.Response) (SparkJobDefinitionGetSparkJobDefinitionResponse, error) {
	result := SparkJobDefinitionGetSparkJobDefinitionResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SparkJobDefinitionResource); err != nil {
		return SparkJobDefinitionGetSparkJobDefinitionResponse{}, err
	}
	return result, nil
}

// getSparkJobDefinitionHandleError handles the GetSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetSparkJobDefinitionsByWorkspace - Lists spark job definitions.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspace(options *SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceOptions) SparkJobDefinitionGetSparkJobDefinitionsByWorkspacePager {
	return &sparkJobDefinitionGetSparkJobDefinitionsByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getSparkJobDefinitionsByWorkspaceCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SparkJobDefinitionsListResponse.NextLink)
		},
	}
}

// getSparkJobDefinitionsByWorkspaceCreateRequest creates the GetSparkJobDefinitionsByWorkspace request.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionsByWorkspaceCreateRequest(ctx context.Context, options *SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkJobDefinitionsByWorkspaceHandleResponse handles the GetSparkJobDefinitionsByWorkspace response.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionsByWorkspaceHandleResponse(resp *azcore.Response) (SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceResponse, error) {
	result := SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SparkJobDefinitionsListResponse); err != nil {
		return SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceResponse{}, err
	}
	return result, nil
}

// getSparkJobDefinitionsByWorkspaceHandleError handles the GetSparkJobDefinitionsByWorkspace error response.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionsByWorkspaceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginRenameSparkJobDefinition - Renames a sparkJobDefinition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) BeginRenameSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, request ArtifactRenameRequest, options *SparkJobDefinitionBeginRenameSparkJobDefinitionOptions) (SparkJobDefinitionRenameSparkJobDefinitionPollerResponse, error) {
	resp, err := client.renameSparkJobDefinition(ctx, sparkJobDefinitionName, request, options)
	if err != nil {
		return SparkJobDefinitionRenameSparkJobDefinitionPollerResponse{}, err
	}
	result := SparkJobDefinitionRenameSparkJobDefinitionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("sparkJobDefinitionClient.RenameSparkJobDefinition", resp, client.con.Pipeline(), client.renameSparkJobDefinitionHandleError)
	if err != nil {
		return SparkJobDefinitionRenameSparkJobDefinitionPollerResponse{}, err
	}
	poller := &sparkJobDefinitionRenameSparkJobDefinitionPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SparkJobDefinitionRenameSparkJobDefinitionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeRenameSparkJobDefinition creates a new SparkJobDefinitionRenameSparkJobDefinitionPoller from the specified resume token.
// token - The value must come from a previous call to SparkJobDefinitionRenameSparkJobDefinitionPoller.ResumeToken().
func (client *sparkJobDefinitionClient) ResumeRenameSparkJobDefinition(ctx context.Context, token string) (SparkJobDefinitionRenameSparkJobDefinitionPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("sparkJobDefinitionClient.RenameSparkJobDefinition", token, client.con.Pipeline(), client.renameSparkJobDefinitionHandleError)
	if err != nil {
		return SparkJobDefinitionRenameSparkJobDefinitionPollerResponse{}, err
	}
	poller := &sparkJobDefinitionRenameSparkJobDefinitionPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SparkJobDefinitionRenameSparkJobDefinitionPollerResponse{}, err
	}
	result := SparkJobDefinitionRenameSparkJobDefinitionPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SparkJobDefinitionRenameSparkJobDefinitionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// RenameSparkJobDefinition - Renames a sparkJobDefinition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) renameSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, request ArtifactRenameRequest, options *SparkJobDefinitionBeginRenameSparkJobDefinitionOptions) (*azcore.Response, error) {
	req, err := client.renameSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.renameSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// renameSparkJobDefinitionCreateRequest creates the RenameSparkJobDefinition request.
func (client *sparkJobDefinitionClient) renameSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, request ArtifactRenameRequest, options *SparkJobDefinitionBeginRenameSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}/rename"
	if sparkJobDefinitionName == "" {
		return nil, errors.New("parameter sparkJobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// renameSparkJobDefinitionHandleError handles the RenameSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) renameSparkJobDefinitionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
