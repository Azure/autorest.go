// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azkeyvault

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// RoleAssignmentsClient contains the methods for the RoleAssignments group.
// Don't use this type directly, use NewRoleAssignmentsClient() instead.
type RoleAssignmentsClient struct {
	con *Connection
}

// NewRoleAssignmentsClient creates a new instance of RoleAssignmentsClient with the specified values.
func NewRoleAssignmentsClient(con *Connection) *RoleAssignmentsClient {
	return &RoleAssignmentsClient{con: con}
}

// Create - Creates a role assignment.
// If the operation fails it returns the *KeyVaultError error type.
func (client *RoleAssignmentsClient) Create(ctx context.Context, vaultBaseURL string, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsCreateOptions) (RoleAssignmentsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, vaultBaseURL, scope, roleAssignmentName, parameters, options)
	if err != nil {
		return RoleAssignmentsCreateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentsCreateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return RoleAssignmentsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *RoleAssignmentsClient) createCreateRequest(ctx context.Context, vaultBaseURL string, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsCreateOptions) (*azcore.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", url.PathEscape(roleAssignmentName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "7.2")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createHandleResponse handles the Create response.
func (client *RoleAssignmentsClient) createHandleResponse(resp *azcore.Response) (RoleAssignmentsCreateResponse, error) {
	result := RoleAssignmentsCreateResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.RoleAssignment); err != nil {
		return RoleAssignmentsCreateResponse{}, err
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *RoleAssignmentsClient) createHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Deletes a role assignment.
// If the operation fails it returns the *KeyVaultError error type.
func (client *RoleAssignmentsClient) Delete(ctx context.Context, vaultBaseURL string, scope string, roleAssignmentName string, options *RoleAssignmentsDeleteOptions) (RoleAssignmentsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, vaultBaseURL, scope, roleAssignmentName, options)
	if err != nil {
		return RoleAssignmentsDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentsDeleteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RoleAssignmentsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *RoleAssignmentsClient) deleteCreateRequest(ctx context.Context, vaultBaseURL string, scope string, roleAssignmentName string, options *RoleAssignmentsDeleteOptions) (*azcore.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", url.PathEscape(roleAssignmentName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "7.2")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *RoleAssignmentsClient) deleteHandleResponse(resp *azcore.Response) (RoleAssignmentsDeleteResponse, error) {
	result := RoleAssignmentsDeleteResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.RoleAssignment); err != nil {
		return RoleAssignmentsDeleteResponse{}, err
	}
	return result, nil
}

// deleteHandleError handles the Delete error response.
func (client *RoleAssignmentsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Get the specified role assignment.
// If the operation fails it returns the *KeyVaultError error type.
func (client *RoleAssignmentsClient) Get(ctx context.Context, vaultBaseURL string, scope string, roleAssignmentName string, options *RoleAssignmentsGetOptions) (RoleAssignmentsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vaultBaseURL, scope, roleAssignmentName, options)
	if err != nil {
		return RoleAssignmentsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RoleAssignmentsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleAssignmentsClient) getCreateRequest(ctx context.Context, vaultBaseURL string, scope string, roleAssignmentName string, options *RoleAssignmentsGetOptions) (*azcore.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", url.PathEscape(roleAssignmentName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "7.2")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleAssignmentsClient) getHandleResponse(resp *azcore.Response) (RoleAssignmentsGetResponse, error) {
	result := RoleAssignmentsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.RoleAssignment); err != nil {
		return RoleAssignmentsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RoleAssignmentsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListForScope - Gets role assignments for a scope.
// If the operation fails it returns the *KeyVaultError error type.
func (client *RoleAssignmentsClient) ListForScope(vaultBaseURL string, scope string, options *RoleAssignmentsListForScopeOptions) RoleAssignmentsListForScopePager {
	return &roleAssignmentsListForScopePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listForScopeCreateRequest(ctx, vaultBaseURL, scope, options)
		},
		advancer: func(ctx context.Context, resp RoleAssignmentsListForScopeResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RoleAssignmentListResult.NextLink)
		},
	}
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleAssignmentsClient) listForScopeCreateRequest(ctx context.Context, vaultBaseURL string, scope string, options *RoleAssignmentsListForScopeOptions) (*azcore.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "7.2")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleAssignmentsClient) listForScopeHandleResponse(resp *azcore.Response) (RoleAssignmentsListForScopeResponse, error) {
	result := RoleAssignmentsListForScopeResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.RoleAssignmentListResult); err != nil {
		return RoleAssignmentsListForScopeResponse{}, err
	}
	return result, nil
}

// listForScopeHandleError handles the ListForScope error response.
func (client *RoleAssignmentsClient) listForScopeHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := KeyVaultError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
