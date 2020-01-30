// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"net/http"
	"net/url"
	"path"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// BasicOperations ..
type BasicOperations struct{}

// GetValidCreateRequest creates the GetValid request.
func (BasicOperations) GetValidCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/valid")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetValidHandleResponse handles the GetValid response.
func (BasicOperations) GetValidHandleResponse(resp *azcore.Response) (*GetValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetValidResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Basic)
}

// PutValidCreateRequest creates the PutValid request.
func (BasicOperations) PutValidCreateRequest(u url.URL, basicBody Basic) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/valid")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(basicBody)
	if err != nil {
		return nil, err
	}
	req.SetQueryParam("api-version", "2016-02-29")
	return req, nil
}

// PutValidHandleResponse handles the PutValid response.
func (BasicOperations) PutValidHandleResponse(resp *azcore.Response) (*PutValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PutValidResponse{StatusCode: resp.StatusCode}, nil
}

// GetInvalidCreateRequest creates the GetValid request.
func (BasicOperations) GetInvalidCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/invalid")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetInvalidHandleResponse handles the GetValid response.
func (BasicOperations) GetInvalidHandleResponse(resp *azcore.Response) (*GetInvalidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetInvalidResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Basic)
}

// GetEmptyCreateRequest creates the GetEmpty request.
func (BasicOperations) GetEmptyCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/empty")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetEmptyHandleResponse handles the GetEmpty response.
func (BasicOperations) GetEmptyHandleResponse(resp *azcore.Response) (*GetEmptyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetEmptyResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Basic)
}

// GetNullCreateRequest creates the GetNull request.
func (BasicOperations) GetNullCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/null")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetNullHandleResponse handles the GetNull response.
func (BasicOperations) GetNullHandleResponse(resp *azcore.Response) (*GetNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetNullResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Basic)
}

// GetNotProvidedCreateRequest creates the GetNotProvided request.
func (BasicOperations) GetNotProvidedCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/notprovided")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetNotProvidedHandleResponse handles the GetNotProvided response.
func (BasicOperations) GetNotProvidedHandleResponse(resp *azcore.Response) (*GetNotProvidedResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetNotProvidedResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Basic)
}
