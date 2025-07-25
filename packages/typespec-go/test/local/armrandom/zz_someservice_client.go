// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armrandom

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// SomeServiceClient contains the methods for the SomeService group.
// Don't use this type directly, use NewSomeServiceClient() instead.
type SomeServiceClient struct {
	internal *arm.Client
}

// NewSomeServiceClient creates a new instance of SomeServiceClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSomeServiceClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*SomeServiceClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SomeServiceClient{
		internal: cl,
	}
	return client, nil
}

// CheckTrialAvailability - Return trial status for subscription by region
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - SomeServiceClientCheckTrialAvailabilityOptions contains the optional parameters for the SomeServiceClient.CheckTrialAvailability
//     method.
func (client *SomeServiceClient) CheckTrialAvailability(ctx context.Context, options *SomeServiceClientCheckTrialAvailabilityOptions) (SomeServiceClientCheckTrialAvailabilityResponse, error) {
	var err error
	const operationName = "SomeServiceClient.CheckTrialAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkTrialAvailabilityCreateRequest(ctx, options)
	if err != nil {
		return SomeServiceClientCheckTrialAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SomeServiceClientCheckTrialAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SomeServiceClientCheckTrialAvailabilityResponse{}, err
	}
	resp, err := client.checkTrialAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkTrialAvailabilityCreateRequest creates the CheckTrialAvailability request.
func (client *SomeServiceClient) checkTrialAvailabilityCreateRequest(ctx context.Context, options *SomeServiceClientCheckTrialAvailabilityOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPost, client.internal.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.SKU != nil {
		req.Raw().Header["Content-Type"] = []string{"application/json"}
		if err := runtime.MarshalAsJSON(req, *options.SKU); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// checkTrialAvailabilityHandleResponse handles the CheckTrialAvailability response.
func (client *SomeServiceClient) checkTrialAvailabilityHandleResponse(resp *http.Response) (SomeServiceClientCheckTrialAvailabilityResponse, error) {
	result := SomeServiceClientCheckTrialAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Trial); err != nil {
		return SomeServiceClientCheckTrialAvailabilityResponse{}, err
	}
	return result, nil
}

// NewListThingsPager - Misc test APIs
//
// Generated from API version 2024-03-01
//   - options - SomeServiceClientListThingsOptions contains the optional parameters for the SomeServiceClient.NewListThingsPager
//     method.
func (client *SomeServiceClient) NewListThingsPager(options *SomeServiceClientListThingsOptions) *runtime.Pager[SomeServiceClientListThingsResponse] {
	return runtime.NewPager(runtime.PagingHandler[SomeServiceClientListThingsResponse]{
		More: func(page SomeServiceClientListThingsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SomeServiceClientListThingsResponse) (SomeServiceClientListThingsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SomeServiceClient.NewListThingsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listThingsCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SomeServiceClientListThingsResponse{}, err
			}
			return client.listThingsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listThingsCreateRequest creates the ListThings request.
func (client *SomeServiceClient) listThingsCreateRequest(ctx context.Context, _ *SomeServiceClientListThingsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Random/listThings"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listThingsHandleResponse handles the ListThings response.
func (client *SomeServiceClient) listThingsHandleResponse(resp *http.Response) (SomeServiceClientListThingsResponse, error) {
	result := SomeServiceClientListThingsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ThingsListResult); err != nil {
		return SomeServiceClientListThingsResponse{}, err
	}
	return result, nil
}
