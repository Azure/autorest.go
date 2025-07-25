// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package clientopgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// SecondClient contains the methods for the Second group.
// Don't use this type directly, use a constructor function instead.
type SecondClient struct {
	internal *azcore.Client
	endpoint string
	client   ClientType
}

// NewSecondGroup5Client creates a new instance of [SecondGroup5Client].
func (client *SecondClient) NewSecondGroup5Client() *SecondGroup5Client {
	return &SecondGroup5Client{
		internal: client.internal,
		endpoint: client.endpoint,
		client:   client.client,
	}
}

// Five -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - SecondClientFiveOptions contains the optional parameters for the SecondClient.Five method.
func (client *SecondClient) Five(ctx context.Context, options *SecondClientFiveOptions) (SecondClientFiveResponse, error) {
	var err error
	const operationName = "SecondClient.Five"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.fiveCreateRequest(ctx, options)
	if err != nil {
		return SecondClientFiveResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecondClientFiveResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SecondClientFiveResponse{}, err
	}
	return SecondClientFiveResponse{}, nil
}

// fiveCreateRequest creates the Five request.
func (client *SecondClient) fiveCreateRequest(ctx context.Context, _ *SecondClientFiveOptions) (*policy.Request, error) {
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/five"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
