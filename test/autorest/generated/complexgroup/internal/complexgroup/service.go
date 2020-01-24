// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"net/http"
	"net/url"
	"path"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// Service ..
type Service struct{}

// GetValidCreateRequest creates the GetValid request.
func (Service) GetValidCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/valid")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetValidHandleResponse handles the GetValid response.
func (Service) GetValidHandleResponse(resp *azcore.Response) (*GetValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetValidResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Basic)
}

// PutValidCreateRequest creates the PutValid request.
func (Service) PutValidCreateRequest(u url.URL, basicBody Basic) (*azcore.Request, error) {
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
func (Service) PutValidHandleResponse(resp *azcore.Response) (*PutValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PutValidResponse{StatusCode: resp.StatusCode}, nil
}

// GetInvalidCreateRequest creates the GetValid request.
func (Service) GetInvalidCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/invalid")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetInvalidHandleResponse handles the GetValid response.
func (Service) GetInvalidHandleResponse(resp *azcore.Response) (*GetInvalidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetInvalidResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Basic)
}

// GetEmptyCreateRequest creates the GetEmpty request.
func (Service) GetEmptyCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/empty")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetEmptyHandleResponse handles the GetEmpty response.
func (Service) GetEmptyHandleResponse(resp *azcore.Response) (*GetEmptyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetEmptyResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Basic)
}

// GetNullCreateRequest creates the GetNull request.
func (Service) GetNullCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/null")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetNullHandleResponse handles the GetNull response.
func (Service) GetNullHandleResponse(resp *azcore.Response) (*GetNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetNullResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Basic)
}

// GetNotProvidedCreateRequest creates the GetNotProvided request.
func (Service) GetNotProvidedCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/notprovided")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetNotProvidedHandleResponse handles the GetNotProvided response.
func (Service) GetNotProvidedHandleResponse(resp *azcore.Response) (*GetNotProvidedResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetNotProvidedResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Basic)
}

// GetIntCreateRequest creates the GetInt request.
func (Service) GetIntCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/integer")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetIntHandleResponse handles the GetInt response.
func (Service) GetIntHandleResponse(resp *azcore.Response) (*GetIntResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetIntResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.IntWrapper)
}

// PutIntCreateRequest creates the PutInt request.
func (Service) PutIntCreateRequest(u url.URL, complexBody *IntWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/integer")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutIntHandleResponse handles the PutInt response.
func (Service) PutIntHandleResponse(resp *azcore.Response) (*PutIntResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutIntResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.IntWrapper)
}
