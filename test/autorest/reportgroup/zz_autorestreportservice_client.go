//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package reportgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// AutoRestReportServiceClient contains the methods for the AutoRestReportService group.
// Don't use this type directly, use NewAutoRestReportServiceClient() instead.
type AutoRestReportServiceClient struct {
	pl runtime.Pipeline
}

// NewAutoRestReportServiceClient creates a new instance of AutoRestReportServiceClient with the specified values.
//   - pl - the pipeline used for sending requests and handling responses.
func NewAutoRestReportServiceClient(pl runtime.Pipeline) *AutoRestReportServiceClient {
	client := &AutoRestReportServiceClient{
		pl: pl,
	}
	return client
}

// GetOptionalReport - Get optional test coverage report
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - AutoRestReportServiceClientGetOptionalReportOptions contains the optional parameters for the AutoRestReportServiceClient.GetOptionalReport
//     method.
func (client *AutoRestReportServiceClient) GetOptionalReport(ctx context.Context, options *AutoRestReportServiceClientGetOptionalReportOptions) (AutoRestReportServiceClientGetOptionalReportResponse, error) {
	req, err := client.getOptionalReportCreateRequest(ctx, options)
	if err != nil {
		return AutoRestReportServiceClientGetOptionalReportResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AutoRestReportServiceClientGetOptionalReportResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AutoRestReportServiceClientGetOptionalReportResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOptionalReportHandleResponse(resp)
}

// getOptionalReportCreateRequest creates the GetOptionalReport request.
func (client *AutoRestReportServiceClient) getOptionalReportCreateRequest(ctx context.Context, options *AutoRestReportServiceClientGetOptionalReportOptions) (*policy.Request, error) {
	urlPath := "/report/optional"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Qualifier != nil {
		reqQP.Set("qualifier", *options.Qualifier)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getOptionalReportHandleResponse handles the GetOptionalReport response.
func (client *AutoRestReportServiceClient) getOptionalReportHandleResponse(resp *http.Response) (AutoRestReportServiceClientGetOptionalReportResponse, error) {
	result := AutoRestReportServiceClientGetOptionalReportResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return AutoRestReportServiceClientGetOptionalReportResponse{}, err
	}
	return result, nil
}

// GetReport - Get test coverage report
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - AutoRestReportServiceClientGetReportOptions contains the optional parameters for the AutoRestReportServiceClient.GetReport
//     method.
func (client *AutoRestReportServiceClient) GetReport(ctx context.Context, options *AutoRestReportServiceClientGetReportOptions) (AutoRestReportServiceClientGetReportResponse, error) {
	req, err := client.getReportCreateRequest(ctx, options)
	if err != nil {
		return AutoRestReportServiceClientGetReportResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AutoRestReportServiceClientGetReportResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AutoRestReportServiceClientGetReportResponse{}, runtime.NewResponseError(resp)
	}
	return client.getReportHandleResponse(resp)
}

// getReportCreateRequest creates the GetReport request.
func (client *AutoRestReportServiceClient) getReportCreateRequest(ctx context.Context, options *AutoRestReportServiceClientGetReportOptions) (*policy.Request, error) {
	urlPath := "/report"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Qualifier != nil {
		reqQP.Set("qualifier", *options.Qualifier)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getReportHandleResponse handles the GetReport response.
func (client *AutoRestReportServiceClient) getReportHandleResponse(resp *http.Response) (AutoRestReportServiceClientGetReportResponse, error) {
	result := AutoRestReportServiceClientGetReportResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return AutoRestReportServiceClientGetReportResponse{}, err
	}
	return result, nil
}