// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// FlattencomplexOperations contains the methods for the Flattencomplex group.
type FlattencomplexOperations interface {
	GetValid(ctx context.Context) (*MyBaseTypeResponse, error)
}

// flattencomplexOperations implements the FlattencomplexOperations interface.
type flattencomplexOperations struct {
	*Client
}

func (client *flattencomplexOperations) GetValid(ctx context.Context) (*MyBaseTypeResponse, error) {
	req, err := client.getValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getValidCreateRequest creates the GetValid request.
func (client *flattencomplexOperations) getValidCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/flatten/valid"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *flattencomplexOperations) getValidHandleResponse(resp *azcore.Response) (*MyBaseTypeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getValidHandleError(resp)
	}
	result := MyBaseTypeResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result)
}

// getValidHandleError handles the GetValid error response.
func (client *flattencomplexOperations) getValidHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}