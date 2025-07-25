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

// XMLModelWithEncodedNamesValueClient - Operations for the ModelWithEncodedNames type.
// Don't use this type directly, use [XMLClient.NewXMLModelWithEncodedNamesValueClient] instead.
type XMLModelWithEncodedNamesValueClient struct {
	internal *azcore.Client
	endpoint string
}

// Get -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - XMLModelWithEncodedNamesValueClientGetOptions contains the optional parameters for the XMLModelWithEncodedNamesValueClient.Get
//     method.
func (client *XMLModelWithEncodedNamesValueClient) Get(ctx context.Context, options *XMLModelWithEncodedNamesValueClientGetOptions) (XMLModelWithEncodedNamesValueClientGetResponse, error) {
	var err error
	const operationName = "XMLModelWithEncodedNamesValueClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return XMLModelWithEncodedNamesValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return XMLModelWithEncodedNamesValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return XMLModelWithEncodedNamesValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *XMLModelWithEncodedNamesValueClient) getCreateRequest(ctx context.Context, _ *XMLModelWithEncodedNamesValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/payload/xml/modelWithEncodedNames"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *XMLModelWithEncodedNamesValueClient) getHandleResponse(resp *http.Response) (XMLModelWithEncodedNamesValueClientGetResponse, error) {
	result := XMLModelWithEncodedNamesValueClientGetResponse{}
	if val := resp.Header.Get("content-type"); val != "" {
		result.ContentType = &val
	}
	if err := runtime.UnmarshalAsXML(resp, &result.ModelWithEncodedNames); err != nil {
		return XMLModelWithEncodedNamesValueClientGetResponse{}, err
	}
	return result, nil
}

// Put -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - XMLModelWithEncodedNamesValueClientPutOptions contains the optional parameters for the XMLModelWithEncodedNamesValueClient.Put
//     method.
func (client *XMLModelWithEncodedNamesValueClient) Put(ctx context.Context, input ModelWithEncodedNames, options *XMLModelWithEncodedNamesValueClientPutOptions) (XMLModelWithEncodedNamesValueClientPutResponse, error) {
	var err error
	const operationName = "XMLModelWithEncodedNamesValueClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, input, options)
	if err != nil {
		return XMLModelWithEncodedNamesValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return XMLModelWithEncodedNamesValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return XMLModelWithEncodedNamesValueClientPutResponse{}, err
	}
	return XMLModelWithEncodedNamesValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *XMLModelWithEncodedNamesValueClient) putCreateRequest(ctx context.Context, input ModelWithEncodedNames, _ *XMLModelWithEncodedNamesValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/payload/xml/modelWithEncodedNames"
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
