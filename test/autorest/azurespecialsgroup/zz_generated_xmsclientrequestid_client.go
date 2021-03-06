// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// XMSClientRequestIDClient contains the methods for the XMSClientRequestID group.
// Don't use this type directly, use NewXMSClientRequestIDClient() instead.
type XMSClientRequestIDClient struct {
	con *Connection
}

// NewXMSClientRequestIDClient creates a new instance of XMSClientRequestIDClient with the specified values.
func NewXMSClientRequestIDClient(con *Connection) *XMSClientRequestIDClient {
	return &XMSClientRequestIDClient{con: con}
}

// Get - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
// If the operation fails it returns a generic error.
func (client *XMSClientRequestIDClient) Get(ctx context.Context, options *XMSClientRequestIDGetOptions) (XMSClientRequestIDGetResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return XMSClientRequestIDGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return XMSClientRequestIDGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return XMSClientRequestIDGetResponse{}, client.getHandleError(resp)
	}
	return XMSClientRequestIDGetResponse{RawResponse: resp.Response}, nil
}

// getCreateRequest creates the Get request.
func (client *XMSClientRequestIDClient) getCreateRequest(ctx context.Context, options *XMSClientRequestIDGetOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/overwrite/x-ms-client-request-id/method/"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	return req, nil
}

// getHandleError handles the Get error response.
func (client *XMSClientRequestIDClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ParamGet - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
// If the operation fails it returns the *Error error type.
func (client *XMSClientRequestIDClient) ParamGet(ctx context.Context, xmsClientRequestID string, options *XMSClientRequestIDParamGetOptions) (XMSClientRequestIDParamGetResponse, error) {
	req, err := client.paramGetCreateRequest(ctx, xmsClientRequestID, options)
	if err != nil {
		return XMSClientRequestIDParamGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return XMSClientRequestIDParamGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return XMSClientRequestIDParamGetResponse{}, client.paramGetHandleError(resp)
	}
	return XMSClientRequestIDParamGetResponse{RawResponse: resp.Response}, nil
}

// paramGetCreateRequest creates the ParamGet request.
func (client *XMSClientRequestIDClient) paramGetCreateRequest(ctx context.Context, xmsClientRequestID string, options *XMSClientRequestIDParamGetOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/overwrite/x-ms-client-request-id/via-param/method/"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("x-ms-client-request-id", xmsClientRequestID)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// paramGetHandleError handles the ParamGet error response.
func (client *XMSClientRequestIDClient) paramGetHandleError(resp *azcore.Response) error {
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
