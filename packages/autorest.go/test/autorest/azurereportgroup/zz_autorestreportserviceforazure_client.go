// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurereportgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// AutoRestReportServiceForAzureClient contains the methods for the AutoRestReportServiceForAzure group.
// Don't use this type directly, use a constructor function instead.
type AutoRestReportServiceForAzureClient struct {
	internal *azcore.Client
	endpoint string
}

// GetReport - Get test coverage report
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - AutoRestReportServiceForAzureClientGetReportOptions contains the optional parameters for the AutoRestReportServiceForAzureClient.GetReport
//     method.
func (client *AutoRestReportServiceForAzureClient) GetReport(ctx context.Context, options *AutoRestReportServiceForAzureClientGetReportOptions) (AutoRestReportServiceForAzureClientGetReportResponse, error) {
	var err error
	const operationName = "AutoRestReportServiceForAzureClient.GetReport"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getReportCreateRequest(ctx, options)
	if err != nil {
		return AutoRestReportServiceForAzureClientGetReportResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AutoRestReportServiceForAzureClientGetReportResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AutoRestReportServiceForAzureClientGetReportResponse{}, err
	}
	resp, err := client.getReportHandleResponse(httpResp)
	return resp, err
}

// getReportCreateRequest creates the GetReport request.
func (client *AutoRestReportServiceForAzureClient) getReportCreateRequest(ctx context.Context, options *AutoRestReportServiceForAzureClientGetReportOptions) (*policy.Request, error) {
	urlPath := "/report/azure"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
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
func (client *AutoRestReportServiceForAzureClient) getReportHandleResponse(resp *http.Response) (AutoRestReportServiceForAzureClientGetReportResponse, error) {
	result := AutoRestReportServiceForAzureClientGetReportResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return AutoRestReportServiceForAzureClientGetReportResponse{}, err
	}
	return result, nil
}
