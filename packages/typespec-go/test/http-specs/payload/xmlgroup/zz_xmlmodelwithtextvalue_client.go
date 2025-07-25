// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package xmlgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// XMLModelWithTextValueClient - Operations for the ModelWithText type.
// Don't use this type directly, use [XMLClient.NewXMLModelWithTextValueClient] instead.
type XMLModelWithTextValueClient struct {
	internal *azcore.Client
	endpoint string
}

// Get -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - XMLModelWithTextValueClientGetOptions contains the optional parameters for the XMLModelWithTextValueClient.Get
//     method.
func (client *XMLModelWithTextValueClient) Get(ctx context.Context, options *XMLModelWithTextValueClientGetOptions) (XMLModelWithTextValueClientGetResponse, error) {
	var err error
	const operationName = "XMLModelWithTextValueClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return XMLModelWithTextValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return XMLModelWithTextValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return XMLModelWithTextValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *XMLModelWithTextValueClient) getCreateRequest(ctx context.Context, _ *XMLModelWithTextValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/payload/xml/modelWithText"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *XMLModelWithTextValueClient) getHandleResponse(resp *http.Response) (XMLModelWithTextValueClientGetResponse, error) {
	result := XMLModelWithTextValueClientGetResponse{}
	if val := resp.Header.Get("content-type"); val != "" {
		result.ContentType = &val
	}
	if err := runtime.UnmarshalAsXML(resp, &result.ModelWithText); err != nil {
		return XMLModelWithTextValueClientGetResponse{}, err
	}
	return result, nil
}

// Put -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - XMLModelWithTextValueClientPutOptions contains the optional parameters for the XMLModelWithTextValueClient.Put
//     method.
func (client *XMLModelWithTextValueClient) Put(ctx context.Context, input ModelWithText, options *XMLModelWithTextValueClientPutOptions) (XMLModelWithTextValueClientPutResponse, error) {
	var err error
	const operationName = "XMLModelWithTextValueClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, input, options)
	if err != nil {
		return XMLModelWithTextValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return XMLModelWithTextValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return XMLModelWithTextValueClientPutResponse{}, err
	}
	return XMLModelWithTextValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *XMLModelWithTextValueClient) putCreateRequest(ctx context.Context, input ModelWithText, _ *XMLModelWithTextValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/payload/xml/modelWithText"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/xml"}
	if err := runtime.MarshalAsXML(req, input); err != nil {
		return nil, err
	}
	return req, nil
}
