// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package optionalgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ImplicitClient contains the methods for the Implicit group.
// Don't use this type directly, use NewImplicitClient() instead.
type ImplicitClient struct {
	con                 *Connection
	requiredGlobalPath  string
	requiredGlobalQuery string
	optionalGlobalQuery *int32
}

// NewImplicitClient creates a new instance of ImplicitClient with the specified values.
func NewImplicitClient(con *Connection, requiredGlobalPath string, requiredGlobalQuery string, optionalGlobalQuery *int32) *ImplicitClient {
	return &ImplicitClient{con: con, requiredGlobalPath: requiredGlobalPath, requiredGlobalQuery: requiredGlobalQuery, optionalGlobalQuery: optionalGlobalQuery}
}

// GetOptionalGlobalQuery - Test implicitly optional query parameter
// If the operation fails it returns the *Error error type.
func (client *ImplicitClient) GetOptionalGlobalQuery(ctx context.Context, options *ImplicitGetOptionalGlobalQueryOptions) (ImplicitGetOptionalGlobalQueryResponse, error) {
	req, err := client.getOptionalGlobalQueryCreateRequest(ctx, options)
	if err != nil {
		return ImplicitGetOptionalGlobalQueryResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ImplicitGetOptionalGlobalQueryResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ImplicitGetOptionalGlobalQueryResponse{}, client.getOptionalGlobalQueryHandleError(resp)
	}
	return ImplicitGetOptionalGlobalQueryResponse{RawResponse: resp.Response}, nil
}

// getOptionalGlobalQueryCreateRequest creates the GetOptionalGlobalQuery request.
func (client *ImplicitClient) getOptionalGlobalQueryCreateRequest(ctx context.Context, options *ImplicitGetOptionalGlobalQueryOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/global/optional/query"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if client.optionalGlobalQuery != nil {
		reqQP.Set("optional-global-query", strconv.FormatInt(int64(*client.optionalGlobalQuery), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getOptionalGlobalQueryHandleError handles the GetOptionalGlobalQuery error response.
func (client *ImplicitClient) getOptionalGlobalQueryHandleError(resp *azcore.Response) error {
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

// GetRequiredGlobalPath - Test implicitly required path parameter
// If the operation fails it returns the *Error error type.
func (client *ImplicitClient) GetRequiredGlobalPath(ctx context.Context, options *ImplicitGetRequiredGlobalPathOptions) (ImplicitGetRequiredGlobalPathResponse, error) {
	req, err := client.getRequiredGlobalPathCreateRequest(ctx, options)
	if err != nil {
		return ImplicitGetRequiredGlobalPathResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ImplicitGetRequiredGlobalPathResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ImplicitGetRequiredGlobalPathResponse{}, client.getRequiredGlobalPathHandleError(resp)
	}
	return ImplicitGetRequiredGlobalPathResponse{RawResponse: resp.Response}, nil
}

// getRequiredGlobalPathCreateRequest creates the GetRequiredGlobalPath request.
func (client *ImplicitClient) getRequiredGlobalPathCreateRequest(ctx context.Context, options *ImplicitGetRequiredGlobalPathOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/global/required/path/{required-global-path}"
	if client.requiredGlobalPath == "" {
		return nil, errors.New("parameter client.requiredGlobalPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{required-global-path}", url.PathEscape(client.requiredGlobalPath))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getRequiredGlobalPathHandleError handles the GetRequiredGlobalPath error response.
func (client *ImplicitClient) getRequiredGlobalPathHandleError(resp *azcore.Response) error {
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

// GetRequiredGlobalQuery - Test implicitly required query parameter
// If the operation fails it returns the *Error error type.
func (client *ImplicitClient) GetRequiredGlobalQuery(ctx context.Context, options *ImplicitGetRequiredGlobalQueryOptions) (ImplicitGetRequiredGlobalQueryResponse, error) {
	req, err := client.getRequiredGlobalQueryCreateRequest(ctx, options)
	if err != nil {
		return ImplicitGetRequiredGlobalQueryResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ImplicitGetRequiredGlobalQueryResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ImplicitGetRequiredGlobalQueryResponse{}, client.getRequiredGlobalQueryHandleError(resp)
	}
	return ImplicitGetRequiredGlobalQueryResponse{RawResponse: resp.Response}, nil
}

// getRequiredGlobalQueryCreateRequest creates the GetRequiredGlobalQuery request.
func (client *ImplicitClient) getRequiredGlobalQueryCreateRequest(ctx context.Context, options *ImplicitGetRequiredGlobalQueryOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/global/required/query"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("required-global-query", client.requiredGlobalQuery)
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getRequiredGlobalQueryHandleError handles the GetRequiredGlobalQuery error response.
func (client *ImplicitClient) getRequiredGlobalQueryHandleError(resp *azcore.Response) error {
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

// GetRequiredPath - Test implicitly required path parameter
// If the operation fails it returns the *Error error type.
func (client *ImplicitClient) GetRequiredPath(ctx context.Context, pathParameter string, options *ImplicitGetRequiredPathOptions) (ImplicitGetRequiredPathResponse, error) {
	req, err := client.getRequiredPathCreateRequest(ctx, pathParameter, options)
	if err != nil {
		return ImplicitGetRequiredPathResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ImplicitGetRequiredPathResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ImplicitGetRequiredPathResponse{}, client.getRequiredPathHandleError(resp)
	}
	return ImplicitGetRequiredPathResponse{RawResponse: resp.Response}, nil
}

// getRequiredPathCreateRequest creates the GetRequiredPath request.
func (client *ImplicitClient) getRequiredPathCreateRequest(ctx context.Context, pathParameter string, options *ImplicitGetRequiredPathOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/implicit/required/path/{pathParameter}"
	if pathParameter == "" {
		return nil, errors.New("parameter pathParameter cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pathParameter}", url.PathEscape(pathParameter))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getRequiredPathHandleError handles the GetRequiredPath error response.
func (client *ImplicitClient) getRequiredPathHandleError(resp *azcore.Response) error {
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

// PutOptionalBinaryBody - Test implicitly optional body parameter
// If the operation fails it returns the *Error error type.
func (client *ImplicitClient) PutOptionalBinaryBody(ctx context.Context, options *ImplicitPutOptionalBinaryBodyOptions) (ImplicitPutOptionalBinaryBodyResponse, error) {
	req, err := client.putOptionalBinaryBodyCreateRequest(ctx, options)
	if err != nil {
		return ImplicitPutOptionalBinaryBodyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ImplicitPutOptionalBinaryBodyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ImplicitPutOptionalBinaryBodyResponse{}, client.putOptionalBinaryBodyHandleError(resp)
	}
	return ImplicitPutOptionalBinaryBodyResponse{RawResponse: resp.Response}, nil
}

// putOptionalBinaryBodyCreateRequest creates the PutOptionalBinaryBody request.
func (client *ImplicitClient) putOptionalBinaryBodyCreateRequest(ctx context.Context, options *ImplicitPutOptionalBinaryBodyOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/implicit/optional/binary-body"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil && options.BodyParameter != nil {
		return req, req.SetBody(options.BodyParameter, "application/octet-stream")
	}
	return req, nil
}

// putOptionalBinaryBodyHandleError handles the PutOptionalBinaryBody error response.
func (client *ImplicitClient) putOptionalBinaryBodyHandleError(resp *azcore.Response) error {
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

// PutOptionalBody - Test implicitly optional body parameter
// If the operation fails it returns the *Error error type.
func (client *ImplicitClient) PutOptionalBody(ctx context.Context, options *ImplicitPutOptionalBodyOptions) (ImplicitPutOptionalBodyResponse, error) {
	req, err := client.putOptionalBodyCreateRequest(ctx, options)
	if err != nil {
		return ImplicitPutOptionalBodyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ImplicitPutOptionalBodyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ImplicitPutOptionalBodyResponse{}, client.putOptionalBodyHandleError(resp)
	}
	return ImplicitPutOptionalBodyResponse{RawResponse: resp.Response}, nil
}

// putOptionalBodyCreateRequest creates the PutOptionalBody request.
func (client *ImplicitClient) putOptionalBodyCreateRequest(ctx context.Context, options *ImplicitPutOptionalBodyOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/implicit/optional/body"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil && options.BodyParameter != nil {
		return req, req.MarshalAsJSON(*options.BodyParameter)
	}
	return req, nil
}

// putOptionalBodyHandleError handles the PutOptionalBody error response.
func (client *ImplicitClient) putOptionalBodyHandleError(resp *azcore.Response) error {
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

// PutOptionalHeader - Test implicitly optional header parameter
// If the operation fails it returns the *Error error type.
func (client *ImplicitClient) PutOptionalHeader(ctx context.Context, options *ImplicitPutOptionalHeaderOptions) (ImplicitPutOptionalHeaderResponse, error) {
	req, err := client.putOptionalHeaderCreateRequest(ctx, options)
	if err != nil {
		return ImplicitPutOptionalHeaderResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ImplicitPutOptionalHeaderResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ImplicitPutOptionalHeaderResponse{}, client.putOptionalHeaderHandleError(resp)
	}
	return ImplicitPutOptionalHeaderResponse{RawResponse: resp.Response}, nil
}

// putOptionalHeaderCreateRequest creates the PutOptionalHeader request.
func (client *ImplicitClient) putOptionalHeaderCreateRequest(ctx context.Context, options *ImplicitPutOptionalHeaderOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/implicit/optional/header"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	if options != nil && options.QueryParameter != nil {
		req.Header.Set("queryParameter", *options.QueryParameter)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// putOptionalHeaderHandleError handles the PutOptionalHeader error response.
func (client *ImplicitClient) putOptionalHeaderHandleError(resp *azcore.Response) error {
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

// PutOptionalQuery - Test implicitly optional query parameter
// If the operation fails it returns the *Error error type.
func (client *ImplicitClient) PutOptionalQuery(ctx context.Context, options *ImplicitPutOptionalQueryOptions) (ImplicitPutOptionalQueryResponse, error) {
	req, err := client.putOptionalQueryCreateRequest(ctx, options)
	if err != nil {
		return ImplicitPutOptionalQueryResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ImplicitPutOptionalQueryResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ImplicitPutOptionalQueryResponse{}, client.putOptionalQueryHandleError(resp)
	}
	return ImplicitPutOptionalQueryResponse{RawResponse: resp.Response}, nil
}

// putOptionalQueryCreateRequest creates the PutOptionalQuery request.
func (client *ImplicitClient) putOptionalQueryCreateRequest(ctx context.Context, options *ImplicitPutOptionalQueryOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/implicit/optional/query"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.QueryParameter != nil {
		reqQP.Set("queryParameter", *options.QueryParameter)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// putOptionalQueryHandleError handles the PutOptionalQuery error response.
func (client *ImplicitClient) putOptionalQueryHandleError(resp *azcore.Response) error {
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
