// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// HTTPClientFailureClient contains the methods for the HTTPClientFailure group.
// Don't use this type directly, use NewHTTPClientFailureClient() instead.
type HTTPClientFailureClient struct {
	con *Connection
}

// NewHTTPClientFailureClient creates a new instance of HTTPClientFailureClient with the specified values.
func NewHTTPClientFailureClient(con *Connection) *HTTPClientFailureClient {
	return &HTTPClientFailureClient{con: con}
}

// Delete400 - Return 400 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Delete400(ctx context.Context, options *HTTPClientFailureDelete400Options) (*http.Response, error) {
	req, err := client.delete400CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.delete400HandleError(resp)
	}
	return resp.Response, nil
}

// delete400CreateRequest creates the Delete400 request.
func (client *HTTPClientFailureClient) delete400CreateRequest(ctx context.Context, options *HTTPClientFailureDelete400Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/400"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// delete400HandleError handles the Delete400 error response.
func (client *HTTPClientFailureClient) delete400HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Delete407 - Return 407 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Delete407(ctx context.Context, options *HTTPClientFailureDelete407Options) (*http.Response, error) {
	req, err := client.delete407CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.delete407HandleError(resp)
	}
	return resp.Response, nil
}

// delete407CreateRequest creates the Delete407 request.
func (client *HTTPClientFailureClient) delete407CreateRequest(ctx context.Context, options *HTTPClientFailureDelete407Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/407"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// delete407HandleError handles the Delete407 error response.
func (client *HTTPClientFailureClient) delete407HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Delete417 - Return 417 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Delete417(ctx context.Context, options *HTTPClientFailureDelete417Options) (*http.Response, error) {
	req, err := client.delete417CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.delete417HandleError(resp)
	}
	return resp.Response, nil
}

// delete417CreateRequest creates the Delete417 request.
func (client *HTTPClientFailureClient) delete417CreateRequest(ctx context.Context, options *HTTPClientFailureDelete417Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/417"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// delete417HandleError handles the Delete417 error response.
func (client *HTTPClientFailureClient) delete417HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get400 - Return 400 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Get400(ctx context.Context, options *HTTPClientFailureGet400Options) (*http.Response, error) {
	req, err := client.get400CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.get400HandleError(resp)
	}
	return resp.Response, nil
}

// get400CreateRequest creates the Get400 request.
func (client *HTTPClientFailureClient) get400CreateRequest(ctx context.Context, options *HTTPClientFailureGet400Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/400"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// get400HandleError handles the Get400 error response.
func (client *HTTPClientFailureClient) get400HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get402 - Return 402 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Get402(ctx context.Context, options *HTTPClientFailureGet402Options) (*http.Response, error) {
	req, err := client.get402CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.get402HandleError(resp)
	}
	return resp.Response, nil
}

// get402CreateRequest creates the Get402 request.
func (client *HTTPClientFailureClient) get402CreateRequest(ctx context.Context, options *HTTPClientFailureGet402Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/402"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// get402HandleError handles the Get402 error response.
func (client *HTTPClientFailureClient) get402HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get403 - Return 403 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Get403(ctx context.Context, options *HTTPClientFailureGet403Options) (*http.Response, error) {
	req, err := client.get403CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.get403HandleError(resp)
	}
	return resp.Response, nil
}

// get403CreateRequest creates the Get403 request.
func (client *HTTPClientFailureClient) get403CreateRequest(ctx context.Context, options *HTTPClientFailureGet403Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/403"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// get403HandleError handles the Get403 error response.
func (client *HTTPClientFailureClient) get403HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get411 - Return 411 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Get411(ctx context.Context, options *HTTPClientFailureGet411Options) (*http.Response, error) {
	req, err := client.get411CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.get411HandleError(resp)
	}
	return resp.Response, nil
}

// get411CreateRequest creates the Get411 request.
func (client *HTTPClientFailureClient) get411CreateRequest(ctx context.Context, options *HTTPClientFailureGet411Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/411"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// get411HandleError handles the Get411 error response.
func (client *HTTPClientFailureClient) get411HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get412 - Return 412 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Get412(ctx context.Context, options *HTTPClientFailureGet412Options) (*http.Response, error) {
	req, err := client.get412CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.get412HandleError(resp)
	}
	return resp.Response, nil
}

// get412CreateRequest creates the Get412 request.
func (client *HTTPClientFailureClient) get412CreateRequest(ctx context.Context, options *HTTPClientFailureGet412Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/412"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// get412HandleError handles the Get412 error response.
func (client *HTTPClientFailureClient) get412HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get416 - Return 416 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Get416(ctx context.Context, options *HTTPClientFailureGet416Options) (*http.Response, error) {
	req, err := client.get416CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.get416HandleError(resp)
	}
	return resp.Response, nil
}

// get416CreateRequest creates the Get416 request.
func (client *HTTPClientFailureClient) get416CreateRequest(ctx context.Context, options *HTTPClientFailureGet416Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/416"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// get416HandleError handles the Get416 error response.
func (client *HTTPClientFailureClient) get416HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Head400 - Return 400 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Head400(ctx context.Context, options *HTTPClientFailureHead400Options) (*http.Response, error) {
	req, err := client.head400CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.head400HandleError(resp)
	}
	return resp.Response, nil
}

// head400CreateRequest creates the Head400 request.
func (client *HTTPClientFailureClient) head400CreateRequest(ctx context.Context, options *HTTPClientFailureHead400Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/400"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// head400HandleError handles the Head400 error response.
func (client *HTTPClientFailureClient) head400HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Head401 - Return 401 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Head401(ctx context.Context, options *HTTPClientFailureHead401Options) (*http.Response, error) {
	req, err := client.head401CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.head401HandleError(resp)
	}
	return resp.Response, nil
}

// head401CreateRequest creates the Head401 request.
func (client *HTTPClientFailureClient) head401CreateRequest(ctx context.Context, options *HTTPClientFailureHead401Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/401"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// head401HandleError handles the Head401 error response.
func (client *HTTPClientFailureClient) head401HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Head410 - Return 410 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Head410(ctx context.Context, options *HTTPClientFailureHead410Options) (*http.Response, error) {
	req, err := client.head410CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.head410HandleError(resp)
	}
	return resp.Response, nil
}

// head410CreateRequest creates the Head410 request.
func (client *HTTPClientFailureClient) head410CreateRequest(ctx context.Context, options *HTTPClientFailureHead410Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/410"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// head410HandleError handles the Head410 error response.
func (client *HTTPClientFailureClient) head410HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Head429 - Return 429 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Head429(ctx context.Context, options *HTTPClientFailureHead429Options) (*http.Response, error) {
	req, err := client.head429CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.head429HandleError(resp)
	}
	return resp.Response, nil
}

// head429CreateRequest creates the Head429 request.
func (client *HTTPClientFailureClient) head429CreateRequest(ctx context.Context, options *HTTPClientFailureHead429Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/429"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// head429HandleError handles the Head429 error response.
func (client *HTTPClientFailureClient) head429HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Options400 - Return 400 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Options400(ctx context.Context, options *HTTPClientFailureOptions400Options) (*http.Response, error) {
	req, err := client.options400CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.options400HandleError(resp)
	}
	return resp.Response, nil
}

// options400CreateRequest creates the Options400 request.
func (client *HTTPClientFailureClient) options400CreateRequest(ctx context.Context, options *HTTPClientFailureOptions400Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/400"
	req, err := azcore.NewRequest(ctx, http.MethodOptions, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// options400HandleError handles the Options400 error response.
func (client *HTTPClientFailureClient) options400HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Options403 - Return 403 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Options403(ctx context.Context, options *HTTPClientFailureOptions403Options) (*http.Response, error) {
	req, err := client.options403CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.options403HandleError(resp)
	}
	return resp.Response, nil
}

// options403CreateRequest creates the Options403 request.
func (client *HTTPClientFailureClient) options403CreateRequest(ctx context.Context, options *HTTPClientFailureOptions403Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/403"
	req, err := azcore.NewRequest(ctx, http.MethodOptions, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// options403HandleError handles the Options403 error response.
func (client *HTTPClientFailureClient) options403HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Options412 - Return 412 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Options412(ctx context.Context, options *HTTPClientFailureOptions412Options) (*http.Response, error) {
	req, err := client.options412CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.options412HandleError(resp)
	}
	return resp.Response, nil
}

// options412CreateRequest creates the Options412 request.
func (client *HTTPClientFailureClient) options412CreateRequest(ctx context.Context, options *HTTPClientFailureOptions412Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/412"
	req, err := azcore.NewRequest(ctx, http.MethodOptions, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// options412HandleError handles the Options412 error response.
func (client *HTTPClientFailureClient) options412HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Patch400 - Return 400 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Patch400(ctx context.Context, options *HTTPClientFailurePatch400Options) (*http.Response, error) {
	req, err := client.patch400CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.patch400HandleError(resp)
	}
	return resp.Response, nil
}

// patch400CreateRequest creates the Patch400 request.
func (client *HTTPClientFailureClient) patch400CreateRequest(ctx context.Context, options *HTTPClientFailurePatch400Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/400"
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// patch400HandleError handles the Patch400 error response.
func (client *HTTPClientFailureClient) patch400HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Patch405 - Return 405 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Patch405(ctx context.Context, options *HTTPClientFailurePatch405Options) (*http.Response, error) {
	req, err := client.patch405CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.patch405HandleError(resp)
	}
	return resp.Response, nil
}

// patch405CreateRequest creates the Patch405 request.
func (client *HTTPClientFailureClient) patch405CreateRequest(ctx context.Context, options *HTTPClientFailurePatch405Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/405"
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// patch405HandleError handles the Patch405 error response.
func (client *HTTPClientFailureClient) patch405HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Patch414 - Return 414 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Patch414(ctx context.Context, options *HTTPClientFailurePatch414Options) (*http.Response, error) {
	req, err := client.patch414CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.patch414HandleError(resp)
	}
	return resp.Response, nil
}

// patch414CreateRequest creates the Patch414 request.
func (client *HTTPClientFailureClient) patch414CreateRequest(ctx context.Context, options *HTTPClientFailurePatch414Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/414"
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// patch414HandleError handles the Patch414 error response.
func (client *HTTPClientFailureClient) patch414HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Post400 - Return 400 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Post400(ctx context.Context, options *HTTPClientFailurePost400Options) (*http.Response, error) {
	req, err := client.post400CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.post400HandleError(resp)
	}
	return resp.Response, nil
}

// post400CreateRequest creates the Post400 request.
func (client *HTTPClientFailureClient) post400CreateRequest(ctx context.Context, options *HTTPClientFailurePost400Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/400"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// post400HandleError handles the Post400 error response.
func (client *HTTPClientFailureClient) post400HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Post406 - Return 406 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Post406(ctx context.Context, options *HTTPClientFailurePost406Options) (*http.Response, error) {
	req, err := client.post406CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.post406HandleError(resp)
	}
	return resp.Response, nil
}

// post406CreateRequest creates the Post406 request.
func (client *HTTPClientFailureClient) post406CreateRequest(ctx context.Context, options *HTTPClientFailurePost406Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/406"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// post406HandleError handles the Post406 error response.
func (client *HTTPClientFailureClient) post406HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Post415 - Return 415 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Post415(ctx context.Context, options *HTTPClientFailurePost415Options) (*http.Response, error) {
	req, err := client.post415CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.post415HandleError(resp)
	}
	return resp.Response, nil
}

// post415CreateRequest creates the Post415 request.
func (client *HTTPClientFailureClient) post415CreateRequest(ctx context.Context, options *HTTPClientFailurePost415Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/415"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// post415HandleError handles the Post415 error response.
func (client *HTTPClientFailureClient) post415HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Put400 - Return 400 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Put400(ctx context.Context, options *HTTPClientFailurePut400Options) (*http.Response, error) {
	req, err := client.put400CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.put400HandleError(resp)
	}
	return resp.Response, nil
}

// put400CreateRequest creates the Put400 request.
func (client *HTTPClientFailureClient) put400CreateRequest(ctx context.Context, options *HTTPClientFailurePut400Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/400"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// put400HandleError handles the Put400 error response.
func (client *HTTPClientFailureClient) put400HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Put404 - Return 404 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Put404(ctx context.Context, options *HTTPClientFailurePut404Options) (*http.Response, error) {
	req, err := client.put404CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.put404HandleError(resp)
	}
	return resp.Response, nil
}

// put404CreateRequest creates the Put404 request.
func (client *HTTPClientFailureClient) put404CreateRequest(ctx context.Context, options *HTTPClientFailurePut404Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/404"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// put404HandleError handles the Put404 error response.
func (client *HTTPClientFailureClient) put404HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Put409 - Return 409 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Put409(ctx context.Context, options *HTTPClientFailurePut409Options) (*http.Response, error) {
	req, err := client.put409CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.put409HandleError(resp)
	}
	return resp.Response, nil
}

// put409CreateRequest creates the Put409 request.
func (client *HTTPClientFailureClient) put409CreateRequest(ctx context.Context, options *HTTPClientFailurePut409Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/409"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// put409HandleError handles the Put409 error response.
func (client *HTTPClientFailureClient) put409HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Put413 - Return 413 status code - should be represented in the client as an error
func (client *HTTPClientFailureClient) Put413(ctx context.Context, options *HTTPClientFailurePut413Options) (*http.Response, error) {
	req, err := client.put413CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.put413HandleError(resp)
	}
	return resp.Response, nil
}

// put413CreateRequest creates the Put413 request.
func (client *HTTPClientFailureClient) put413CreateRequest(ctx context.Context, options *HTTPClientFailurePut413Options) (*azcore.Request, error) {
	urlPath := "/http/failure/client/413"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// put413HandleError handles the Put413 error response.
func (client *HTTPClientFailureClient) put413HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}