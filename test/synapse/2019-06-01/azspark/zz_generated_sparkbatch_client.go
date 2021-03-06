// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azspark

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type sparkBatchClient struct {
	con *connection
}

// CancelSparkBatchJob - Cancels a running spark batch job.
// If the operation fails it returns a generic error.
func (client *sparkBatchClient) CancelSparkBatchJob(ctx context.Context, batchID int32, options *SparkBatchCancelSparkBatchJobOptions) (SparkBatchCancelSparkBatchJobResponse, error) {
	req, err := client.cancelSparkBatchJobCreateRequest(ctx, batchID, options)
	if err != nil {
		return SparkBatchCancelSparkBatchJobResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SparkBatchCancelSparkBatchJobResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SparkBatchCancelSparkBatchJobResponse{}, client.cancelSparkBatchJobHandleError(resp)
	}
	return SparkBatchCancelSparkBatchJobResponse{RawResponse: resp.Response}, nil
}

// cancelSparkBatchJobCreateRequest creates the CancelSparkBatchJob request.
func (client *sparkBatchClient) cancelSparkBatchJobCreateRequest(ctx context.Context, batchID int32, options *SparkBatchCancelSparkBatchJobOptions) (*azcore.Request, error) {
	urlPath := "/batches/{batchId}"
	urlPath = strings.ReplaceAll(urlPath, "{batchId}", url.PathEscape(strconv.FormatInt(int64(batchID), 10)))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	return req, nil
}

// cancelSparkBatchJobHandleError handles the CancelSparkBatchJob error response.
func (client *sparkBatchClient) cancelSparkBatchJobHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// CreateSparkBatchJob - Create new spark batch job.
// If the operation fails it returns a generic error.
func (client *sparkBatchClient) CreateSparkBatchJob(ctx context.Context, sparkBatchJobOptions SparkBatchJobOptions, options *SparkBatchCreateSparkBatchJobOptions) (SparkBatchCreateSparkBatchJobResponse, error) {
	req, err := client.createSparkBatchJobCreateRequest(ctx, sparkBatchJobOptions, options)
	if err != nil {
		return SparkBatchCreateSparkBatchJobResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SparkBatchCreateSparkBatchJobResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SparkBatchCreateSparkBatchJobResponse{}, client.createSparkBatchJobHandleError(resp)
	}
	return client.createSparkBatchJobHandleResponse(resp)
}

// createSparkBatchJobCreateRequest creates the CreateSparkBatchJob request.
func (client *sparkBatchClient) createSparkBatchJobCreateRequest(ctx context.Context, sparkBatchJobOptions SparkBatchJobOptions, options *SparkBatchCreateSparkBatchJobOptions) (*azcore.Request, error) {
	urlPath := "/batches"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Detailed != nil {
		reqQP.Set("detailed", strconv.FormatBool(*options.Detailed))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(sparkBatchJobOptions)
}

// createSparkBatchJobHandleResponse handles the CreateSparkBatchJob response.
func (client *sparkBatchClient) createSparkBatchJobHandleResponse(resp *azcore.Response) (SparkBatchCreateSparkBatchJobResponse, error) {
	result := SparkBatchCreateSparkBatchJobResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SparkBatchJob); err != nil {
		return SparkBatchCreateSparkBatchJobResponse{}, err
	}
	return result, nil
}

// createSparkBatchJobHandleError handles the CreateSparkBatchJob error response.
func (client *sparkBatchClient) createSparkBatchJobHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GetSparkBatchJob - Gets a single spark batch job.
// If the operation fails it returns a generic error.
func (client *sparkBatchClient) GetSparkBatchJob(ctx context.Context, batchID int32, options *SparkBatchGetSparkBatchJobOptions) (SparkBatchGetSparkBatchJobResponse, error) {
	req, err := client.getSparkBatchJobCreateRequest(ctx, batchID, options)
	if err != nil {
		return SparkBatchGetSparkBatchJobResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SparkBatchGetSparkBatchJobResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SparkBatchGetSparkBatchJobResponse{}, client.getSparkBatchJobHandleError(resp)
	}
	return client.getSparkBatchJobHandleResponse(resp)
}

// getSparkBatchJobCreateRequest creates the GetSparkBatchJob request.
func (client *sparkBatchClient) getSparkBatchJobCreateRequest(ctx context.Context, batchID int32, options *SparkBatchGetSparkBatchJobOptions) (*azcore.Request, error) {
	urlPath := "/batches/{batchId}"
	urlPath = strings.ReplaceAll(urlPath, "{batchId}", url.PathEscape(strconv.FormatInt(int64(batchID), 10)))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Detailed != nil {
		reqQP.Set("detailed", strconv.FormatBool(*options.Detailed))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkBatchJobHandleResponse handles the GetSparkBatchJob response.
func (client *sparkBatchClient) getSparkBatchJobHandleResponse(resp *azcore.Response) (SparkBatchGetSparkBatchJobResponse, error) {
	result := SparkBatchGetSparkBatchJobResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SparkBatchJob); err != nil {
		return SparkBatchGetSparkBatchJobResponse{}, err
	}
	return result, nil
}

// getSparkBatchJobHandleError handles the GetSparkBatchJob error response.
func (client *sparkBatchClient) getSparkBatchJobHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GetSparkBatchJobs - List all spark batch jobs which are running under a particular spark pool.
// If the operation fails it returns a generic error.
func (client *sparkBatchClient) GetSparkBatchJobs(ctx context.Context, options *SparkBatchGetSparkBatchJobsOptions) (SparkBatchGetSparkBatchJobsResponse, error) {
	req, err := client.getSparkBatchJobsCreateRequest(ctx, options)
	if err != nil {
		return SparkBatchGetSparkBatchJobsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SparkBatchGetSparkBatchJobsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SparkBatchGetSparkBatchJobsResponse{}, client.getSparkBatchJobsHandleError(resp)
	}
	return client.getSparkBatchJobsHandleResponse(resp)
}

// getSparkBatchJobsCreateRequest creates the GetSparkBatchJobs request.
func (client *sparkBatchClient) getSparkBatchJobsCreateRequest(ctx context.Context, options *SparkBatchGetSparkBatchJobsOptions) (*azcore.Request, error) {
	urlPath := "/batches"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.From != nil {
		reqQP.Set("from", strconv.FormatInt(int64(*options.From), 10))
	}
	if options != nil && options.Size != nil {
		reqQP.Set("size", strconv.FormatInt(int64(*options.Size), 10))
	}
	if options != nil && options.Detailed != nil {
		reqQP.Set("detailed", strconv.FormatBool(*options.Detailed))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkBatchJobsHandleResponse handles the GetSparkBatchJobs response.
func (client *sparkBatchClient) getSparkBatchJobsHandleResponse(resp *azcore.Response) (SparkBatchGetSparkBatchJobsResponse, error) {
	result := SparkBatchGetSparkBatchJobsResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SparkBatchJobCollection); err != nil {
		return SparkBatchGetSparkBatchJobsResponse{}, err
	}
	return result, nil
}

// getSparkBatchJobsHandleError handles the GetSparkBatchJobs error response.
func (client *sparkBatchClient) getSparkBatchJobsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
