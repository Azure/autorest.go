// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package mediatypegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// StringBodyClient contains the methods for the Payload.MediaType namespace.
// Don't use this type directly, use [MediaTypeClient.NewStringBodyClient] instead.
type StringBodyClient struct {
	internal *azcore.Client
}

// - options - StringBodyClientGetAsJSONOptions contains the optional parameters for the StringBodyClient.GetAsJSON method.
func (client *StringBodyClient) GetAsJSON(ctx context.Context, options *StringBodyClientGetAsJSONOptions) (StringBodyClientGetAsJSONResponse, error) {
	var err error
	req, err := client.getAsJSONCreateRequest(ctx, options)
	if err != nil {
		return StringBodyClientGetAsJSONResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringBodyClientGetAsJSONResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringBodyClientGetAsJSONResponse{}, err
	}
	resp, err := client.getAsJSONHandleResponse(httpResp)
	return resp, err
}

// getAsJSONCreateRequest creates the GetAsJSON request.
func (client *StringBodyClient) getAsJSONCreateRequest(ctx context.Context, options *StringBodyClientGetAsJSONOptions) (*policy.Request, error) {
	urlPath := "/payload/media-type/string-body/getAsJson"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAsJSONHandleResponse handles the GetAsJSON response.
func (client *StringBodyClient) getAsJSONHandleResponse(resp *http.Response) (StringBodyClientGetAsJSONResponse, error) {
	result := StringBodyClientGetAsJSONResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return StringBodyClientGetAsJSONResponse{}, err
	}
	return result, nil
}

// - options - StringBodyClientGetAsTextOptions contains the optional parameters for the StringBodyClient.GetAsText method.
func (client *StringBodyClient) GetAsText(ctx context.Context, options *StringBodyClientGetAsTextOptions) (StringBodyClientGetAsTextResponse, error) {
	var err error
	req, err := client.getAsTextCreateRequest(ctx, options)
	if err != nil {
		return StringBodyClientGetAsTextResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringBodyClientGetAsTextResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringBodyClientGetAsTextResponse{}, err
	}
	resp, err := client.getAsTextHandleResponse(httpResp)
	return resp, err
}

// getAsTextCreateRequest creates the GetAsText request.
func (client *StringBodyClient) getAsTextCreateRequest(ctx context.Context, options *StringBodyClientGetAsTextOptions) (*policy.Request, error) {
	urlPath := "/payload/media-type/string-body/getAsText"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"text/plain"}
	return req, nil
}

// getAsTextHandleResponse handles the GetAsText response.
func (client *StringBodyClient) getAsTextHandleResponse(resp *http.Response) (StringBodyClientGetAsTextResponse, error) {
	result := StringBodyClientGetAsTextResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return StringBodyClientGetAsTextResponse{}, err
	}
	return result, nil
}

// - options - StringBodyClientSendAsJSONOptions contains the optional parameters for the StringBodyClient.SendAsJSON method.
func (client *StringBodyClient) SendAsJSON(ctx context.Context, textParam string, options *StringBodyClientSendAsJSONOptions) (StringBodyClientSendAsJSONResponse, error) {
	var err error
	req, err := client.sendAsJSONCreateRequest(ctx, textParam, options)
	if err != nil {
		return StringBodyClientSendAsJSONResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringBodyClientSendAsJSONResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringBodyClientSendAsJSONResponse{}, err
	}
	return StringBodyClientSendAsJSONResponse{}, nil
}

// sendAsJSONCreateRequest creates the SendAsJSON request.
func (client *StringBodyClient) sendAsJSONCreateRequest(ctx context.Context, textParam string, options *StringBodyClientSendAsJSONOptions) (*policy.Request, error) {
	urlPath := "/payload/media-type/string-body/sendAsJson"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["content-type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, textParam); err != nil {
		return nil, err
	}
	return req, nil
}

// - options - StringBodyClientSendAsTextOptions contains the optional parameters for the StringBodyClient.SendAsText method.
func (client *StringBodyClient) SendAsText(ctx context.Context, textParam string, options *StringBodyClientSendAsTextOptions) (StringBodyClientSendAsTextResponse, error) {
	var err error
	req, err := client.sendAsTextCreateRequest(ctx, textParam, options)
	if err != nil {
		return StringBodyClientSendAsTextResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringBodyClientSendAsTextResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringBodyClientSendAsTextResponse{}, err
	}
	return StringBodyClientSendAsTextResponse{}, nil
}

// sendAsTextCreateRequest creates the SendAsText request.
func (client *StringBodyClient) sendAsTextCreateRequest(ctx context.Context, textParam string, options *StringBodyClientSendAsTextOptions) (*policy.Request, error) {
	urlPath := "/payload/media-type/string-body/sendAsText"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["content-type"] = []string{"text/plain"}
	if err := runtime.MarshalAsJSON(req, textParam); err != nil {
		return nil, err
	}
	return req, nil
}