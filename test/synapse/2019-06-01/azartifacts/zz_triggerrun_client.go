//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type triggerRunClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newTriggerRunClient creates a new instance of triggerRunClient with the specified values.
//   - endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
//   - pl - the pipeline used for sending requests and handling responses.
func newTriggerRunClient(endpoint string, pl runtime.Pipeline) *triggerRunClient {
	client := &triggerRunClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// CancelTriggerInstance - Cancel single trigger instance by runId.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - triggerName - The trigger name.
//   - runID - The pipeline run identifier.
//   - options - TriggerRunClientCancelTriggerInstanceOptions contains the optional parameters for the triggerRunClient.CancelTriggerInstance
//     method.
func (client *triggerRunClient) CancelTriggerInstance(ctx context.Context, triggerName string, runID string, options *TriggerRunClientCancelTriggerInstanceOptions) (TriggerRunClientCancelTriggerInstanceResponse, error) {
	req, err := client.cancelTriggerInstanceCreateRequest(ctx, triggerName, runID, options)
	if err != nil {
		return TriggerRunClientCancelTriggerInstanceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TriggerRunClientCancelTriggerInstanceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TriggerRunClientCancelTriggerInstanceResponse{}, runtime.NewResponseError(resp)
	}
	return TriggerRunClientCancelTriggerInstanceResponse{}, nil
}

// cancelTriggerInstanceCreateRequest creates the CancelTriggerInstance request.
func (client *triggerRunClient) cancelTriggerInstanceCreateRequest(ctx context.Context, triggerName string, runID string, options *TriggerRunClientCancelTriggerInstanceOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/triggerRuns/{runId}/cancel"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// QueryTriggerRunsByWorkspace - Query trigger runs.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - filterParameters - Parameters to filter the pipeline run.
//   - options - TriggerRunClientQueryTriggerRunsByWorkspaceOptions contains the optional parameters for the triggerRunClient.QueryTriggerRunsByWorkspace
//     method.
func (client *triggerRunClient) QueryTriggerRunsByWorkspace(ctx context.Context, filterParameters RunFilterParameters, options *TriggerRunClientQueryTriggerRunsByWorkspaceOptions) (TriggerRunClientQueryTriggerRunsByWorkspaceResponse, error) {
	req, err := client.queryTriggerRunsByWorkspaceCreateRequest(ctx, filterParameters, options)
	if err != nil {
		return TriggerRunClientQueryTriggerRunsByWorkspaceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TriggerRunClientQueryTriggerRunsByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TriggerRunClientQueryTriggerRunsByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	return client.queryTriggerRunsByWorkspaceHandleResponse(resp)
}

// queryTriggerRunsByWorkspaceCreateRequest creates the QueryTriggerRunsByWorkspace request.
func (client *triggerRunClient) queryTriggerRunsByWorkspaceCreateRequest(ctx context.Context, filterParameters RunFilterParameters, options *TriggerRunClientQueryTriggerRunsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/queryTriggerRuns"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, filterParameters)
}

// queryTriggerRunsByWorkspaceHandleResponse handles the QueryTriggerRunsByWorkspace response.
func (client *triggerRunClient) queryTriggerRunsByWorkspaceHandleResponse(resp *http.Response) (TriggerRunClientQueryTriggerRunsByWorkspaceResponse, error) {
	result := TriggerRunClientQueryTriggerRunsByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerRunsQueryResponse); err != nil {
		return TriggerRunClientQueryTriggerRunsByWorkspaceResponse{}, err
	}
	return result, nil
}

// RerunTriggerInstance - Rerun single trigger instance by runId.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - triggerName - The trigger name.
//   - runID - The pipeline run identifier.
//   - options - TriggerRunClientRerunTriggerInstanceOptions contains the optional parameters for the triggerRunClient.RerunTriggerInstance
//     method.
func (client *triggerRunClient) RerunTriggerInstance(ctx context.Context, triggerName string, runID string, options *TriggerRunClientRerunTriggerInstanceOptions) (TriggerRunClientRerunTriggerInstanceResponse, error) {
	req, err := client.rerunTriggerInstanceCreateRequest(ctx, triggerName, runID, options)
	if err != nil {
		return TriggerRunClientRerunTriggerInstanceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TriggerRunClientRerunTriggerInstanceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TriggerRunClientRerunTriggerInstanceResponse{}, runtime.NewResponseError(resp)
	}
	return TriggerRunClientRerunTriggerInstanceResponse{}, nil
}

// rerunTriggerInstanceCreateRequest creates the RerunTriggerInstance request.
func (client *triggerRunClient) rerunTriggerInstanceCreateRequest(ctx context.Context, triggerName string, runID string, options *TriggerRunClientRerunTriggerInstanceOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/triggerRuns/{runId}/rerun"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}