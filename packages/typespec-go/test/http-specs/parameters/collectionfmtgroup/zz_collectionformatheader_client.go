// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package collectionfmtgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// CollectionFormatHeaderClient contains the methods for the CollectionFormatHeader group.
// Don't use this type directly, use [CollectionFormatClient.NewCollectionFormatHeaderClient] instead.
type CollectionFormatHeaderClient struct {
	internal *azcore.Client
	endpoint string
}

// CSV -
// If the operation fails it returns an *azcore.ResponseError type.
//   - colors - Possible values for colors are [blue,red,green]
//   - options - CollectionFormatHeaderClientCSVOptions contains the optional parameters for the CollectionFormatHeaderClient.CSV
//     method.
func (client *CollectionFormatHeaderClient) CSV(ctx context.Context, colors []string, options *CollectionFormatHeaderClientCSVOptions) (CollectionFormatHeaderClientCSVResponse, error) {
	var err error
	const operationName = "CollectionFormatHeaderClient.CSV"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.csvCreateRequest(ctx, colors, options)
	if err != nil {
		return CollectionFormatHeaderClientCSVResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CollectionFormatHeaderClientCSVResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return CollectionFormatHeaderClientCSVResponse{}, err
	}
	return CollectionFormatHeaderClientCSVResponse{}, nil
}

// csvCreateRequest creates the CSV request.
func (client *CollectionFormatHeaderClient) csvCreateRequest(ctx context.Context, colors []string, _ *CollectionFormatHeaderClientCSVOptions) (*policy.Request, error) {
	urlPath := "/parameters/collection-format/header/csv"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["colors"] = []string{strings.Join(colors, ",")}
	return req, nil
}
