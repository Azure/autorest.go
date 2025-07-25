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

// XMLModelWithAttributesValueClient - Operations for the ModelWithAttributes type.
// Don't use this type directly, use [XMLClient.NewXMLModelWithAttributesValueClient] instead.
type XMLModelWithAttributesValueClient struct {
	internal *azcore.Client
	endpoint string
}

// Get -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - XMLModelWithAttributesValueClientGetOptions contains the optional parameters for the XMLModelWithAttributesValueClient.Get
//     method.
func (client *XMLModelWithAttributesValueClient) Get(ctx context.Context, options *XMLModelWithAttributesValueClientGetOptions) (XMLModelWithAttributesValueClientGetResponse, error) {
	var err error
	const operationName = "XMLModelWithAttributesValueClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return XMLModelWithAttributesValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return XMLModelWithAttributesValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return XMLModelWithAttributesValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *XMLModelWithAttributesValueClient) getCreateRequest(ctx context.Context, _ *XMLModelWithAttributesValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/payload/xml/modelWithAttributes"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *XMLModelWithAttributesValueClient) getHandleResponse(resp *http.Response) (XMLModelWithAttributesValueClientGetResponse, error) {
	result := XMLModelWithAttributesValueClientGetResponse{}
	if val := resp.Header.Get("content-type"); val != "" {
		result.ContentType = &val
	}
	if err := runtime.UnmarshalAsXML(resp, &result.ModelWithAttributes); err != nil {
		return XMLModelWithAttributesValueClientGetResponse{}, err
	}
	return result, nil
}

// Put -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - XMLModelWithAttributesValueClientPutOptions contains the optional parameters for the XMLModelWithAttributesValueClient.Put
//     method.
func (client *XMLModelWithAttributesValueClient) Put(ctx context.Context, input ModelWithAttributes, options *XMLModelWithAttributesValueClientPutOptions) (XMLModelWithAttributesValueClientPutResponse, error) {
	var err error
	const operationName = "XMLModelWithAttributesValueClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, input, options)
	if err != nil {
		return XMLModelWithAttributesValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return XMLModelWithAttributesValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return XMLModelWithAttributesValueClientPutResponse{}, err
	}
	return XMLModelWithAttributesValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *XMLModelWithAttributesValueClient) putCreateRequest(ctx context.Context, input ModelWithAttributes, _ *XMLModelWithAttributesValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/payload/xml/modelWithAttributes"
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
