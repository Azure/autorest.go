// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// VPNLinkConnectionsClient contains the methods for the VPNLinkConnections group.
// Don't use this type directly, use NewVPNLinkConnectionsClient() instead.
type VPNLinkConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVPNLinkConnectionsClient creates a new instance of VPNLinkConnectionsClient with the specified values.
func NewVPNLinkConnectionsClient(con *armcore.Connection, subscriptionID string) *VPNLinkConnectionsClient {
	return &VPNLinkConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// ListByVPNConnection - Retrieves all vpn site link connections for a particular virtual wan vpn gateway vpn connection.
// If the operation fails it returns the *CloudError error type.
func (client *VPNLinkConnectionsClient) ListByVPNConnection(resourceGroupName string, gatewayName string, connectionName string, options *VPNLinkConnectionsListByVPNConnectionOptions) VPNLinkConnectionsListByVPNConnectionPager {
	return &vpnLinkConnectionsListByVPNConnectionPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByVPNConnectionCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
		},
		advancer: func(ctx context.Context, resp VPNLinkConnectionsListByVPNConnectionResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVPNSiteLinkConnectionsResult.NextLink)
		},
	}
}

// listByVPNConnectionCreateRequest creates the ListByVPNConnection request.
func (client *VPNLinkConnectionsClient) listByVPNConnectionCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNLinkConnectionsListByVPNConnectionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByVPNConnectionHandleResponse handles the ListByVPNConnection response.
func (client *VPNLinkConnectionsClient) listByVPNConnectionHandleResponse(resp *azcore.Response) (VPNLinkConnectionsListByVPNConnectionResponse, error) {
	result := VPNLinkConnectionsListByVPNConnectionResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ListVPNSiteLinkConnectionsResult); err != nil {
		return VPNLinkConnectionsListByVPNConnectionResponse{}, err
	}
	return result, nil
}

// listByVPNConnectionHandleError handles the ListByVPNConnection error response.
func (client *VPNLinkConnectionsClient) listByVPNConnectionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
