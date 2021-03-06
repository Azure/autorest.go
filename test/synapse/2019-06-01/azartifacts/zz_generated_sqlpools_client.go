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
)

type sqlPoolsClient struct {
	con *connection
}

// Get - Get Sql Pool
// If the operation fails it returns the *ErrorContract error type.
func (client *sqlPoolsClient) Get(ctx context.Context, sqlPoolName string, options *SQLPoolsGetOptions) (SQLPoolsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, sqlPoolName, options)
	if err != nil {
		return SQLPoolsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SQLPoolsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SQLPoolsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *sqlPoolsClient) getCreateRequest(ctx context.Context, sqlPoolName string, options *SQLPoolsGetOptions) (*azcore.Request, error) {
	urlPath := "/sqlPools/{sqlPoolName}"
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
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

// getHandleResponse handles the Get response.
func (client *sqlPoolsClient) getHandleResponse(resp *azcore.Response) (SQLPoolsGetResponse, error) {
	result := SQLPoolsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SQLPool); err != nil {
		return SQLPoolsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *sqlPoolsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorContract{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - List Sql Pools
// If the operation fails it returns the *ErrorContract error type.
func (client *sqlPoolsClient) List(ctx context.Context, options *SQLPoolsListOptions) (SQLPoolsListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return SQLPoolsListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SQLPoolsListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SQLPoolsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *sqlPoolsClient) listCreateRequest(ctx context.Context, options *SQLPoolsListOptions) (*azcore.Request, error) {
	urlPath := "/sqlPools"
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

// listHandleResponse handles the List response.
func (client *sqlPoolsClient) listHandleResponse(resp *azcore.Response) (SQLPoolsListResponse, error) {
	result := SQLPoolsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SQLPoolInfoListResult); err != nil {
		return SQLPoolsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *sqlPoolsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorContract{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
