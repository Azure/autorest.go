// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bytegroup

import (
	"net/http"
	"net/url"
	"path"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// Service ..
type Service struct{}

// GetEmptyCreateRequest creates the GetEmpty request.
func (Service) GetEmptyCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/byte/empty")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetEmptyHandleResponse handles the GetEmpty response.
func (Service) GetEmptyHandleResponse(resp *azcore.Response) (*GetEmptyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetEmptyResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetInvalidCreateRequest creates the GetInvalid request.
func (Service) GetInvalidCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/byte/invalid")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetInvalidHandleResponse handles the GetInvalid response.
func (Service) GetInvalidHandleResponse(resp *azcore.Response) (*GetInvalidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetInvalidResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetNonASCIICreateRequest creates the GetNonASCII request.
func (Service) GetNonASCIICreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/byte/nonAscii")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetNonASCIIHandleResponse handles the GetNonASCII response.
func (Service) GetNonASCIIHandleResponse(resp *azcore.Response) (*GetNonASCIIResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetNonASCIIResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetNullCreateRequest creates the GetNull request.
func (Service) GetNullCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/byte/null")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetNullHandleResponse handles the GetNull response.
func (Service) GetNullHandleResponse(resp *azcore.Response) (*GetNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetNullResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// PutNonASCIICreateRequest creates the PutNonASCII request.
func (Service) PutNonASCIICreateRequest(u url.URL, byteBody []byte) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/byte/nonAscii")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(byteBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutNonASCIIHandleResponse handles the PutNonASCII response.
func (Service) PutNonASCIIHandleResponse(resp *azcore.Response) (*PutNonASCIIResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PutNonASCIIResponse{StatusCode: resp.StatusCode}, nil
}
