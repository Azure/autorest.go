// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armmongocluster

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PrivateLinksClient contains the methods for the Microsoft.DocumentDB namespace.
// Don't use this type directly, use NewPrivateLinksClient() instead.
type PrivateLinksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateLinksClient creates a new instance of PrivateLinksClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateLinksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinksClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByMongoClusterPager - list private links on the given resource
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mongoClusterName - The name of the mongo cluster.
//   - options - PrivateLinksClientListByMongoClusterOptions contains the optional parameters for the PrivateLinksClient.NewListByMongoClusterPager
//     method.
func (client *PrivateLinksClient) NewListByMongoClusterPager(resourceGroupName string, mongoClusterName string, options *PrivateLinksClientListByMongoClusterOptions) *runtime.Pager[PrivateLinksClientListByMongoClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinksClientListByMongoClusterResponse]{
		More: func(page PrivateLinksClientListByMongoClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinksClientListByMongoClusterResponse) (PrivateLinksClientListByMongoClusterResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateLinksClient.NewListByMongoClusterPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByMongoClusterCreateRequest(ctx, resourceGroupName, mongoClusterName, options)
			}, nil)
			if err != nil {
				return PrivateLinksClientListByMongoClusterResponse{}, err
			}
			return client.listByMongoClusterHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByMongoClusterCreateRequest creates the ListByMongoCluster request.
func (client *PrivateLinksClient) listByMongoClusterCreateRequest(ctx context.Context, resourceGroupName string, mongoClusterName string, _ *PrivateLinksClientListByMongoClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/mongoClusters/{mongoClusterName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mongoClusterName == "" {
		return nil, errors.New("parameter mongoClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mongoClusterName}", url.PathEscape(mongoClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByMongoClusterHandleResponse handles the ListByMongoCluster response.
func (client *PrivateLinksClient) listByMongoClusterHandleResponse(resp *http.Response) (PrivateLinksClientListByMongoClusterResponse, error) {
	result := PrivateLinksClientListByMongoClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinksClientListByMongoClusterResponse{}, err
	}
	return result, nil
}