// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bytegroup

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// Service ..
type Service struct{}

// GetEmptyCreateRequest creates the GetEmpty request.
func (Service) GetEmptyCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/byte/empty")
	// TODO: this makes two copies
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetEmptyHandleResponse handles the GetEmpty response.
func (Service) GetEmptyHandleResponse(resp *azcore.Response) (*GetEmptyResponse, error) {
	if err := resp.CheckStatusCode(http.StatusOK); err != nil {
		return nil, err
	}
	// TODO: add resp.UnmarshalAsJSON() in azcore
	if len(resp.Payload) == 0 {
		return nil, errors.New("missing payload")
	}
	result := GetEmptyResponse{StatusCode: resp.StatusCode}
	err := json.Unmarshal(resp.Payload, &result.Value)
	if err != nil {
		return nil, errors.New("unmarshalling GetEmptyResponse")
	}
	return &result, nil
}

// GetInvalidRequest creates the GetEmpty request.
func (Service) GetInvalidRequest(u *url.URL) (*azcore.Request, error) {
	if !strings.HasSuffix(u.Path, "/") {
		u.Path = u.Path + "/"
	}
	u.Path = u.Path + "byte/invalid"
	return azcore.NewRequest(http.MethodGet, *u), nil
}

// GetInvalidHandleResponse handles the GetEmpty response.
func (Service) GetInvalidHandleResponse(resp *azcore.Response) (*ByteArray, error) {
	return &ByteArray{Value: &resp.Payload}, nil
}

// GetNonASCIIRequest creates the GetEmpty request.
func (Service) GetNonASCIIRequest(u *url.URL) (*azcore.Request, error) {
	if !strings.HasSuffix(u.Path, "/") {
		u.Path = u.Path + "/"
	}
	u.Path = u.Path + "byte/nonAscii"
	return azcore.NewRequest(http.MethodGet, *u), nil
}

// GetNonASCIIHandleResponse handles the GetEmpty response.
func (Service) GetNonASCIIHandleResponse(resp *azcore.Response) (*ByteArray, error) {
	return &ByteArray{Value: &resp.Payload}, nil
}

// GetNilRequest creates the GetEmpty request.
func (Service) GetNilRequest(u *url.URL) (*azcore.Request, error) {
	if !strings.HasSuffix(u.Path, "/") {
		u.Path = u.Path + "/"
	}
	u.Path = u.Path + "byte/null"
	return azcore.NewRequest(http.MethodGet, *u), nil
}

// GetNilHandleResponse handles the GetEmpty response.
func (Service) GetNilHandleResponse(resp *azcore.Response) (*ByteArray, error) {
	return &ByteArray{Value: &resp.Payload}, nil
}

// PutNonASCIIRequest creates the GetEmpty request.
func (Service) PutNonASCIIRequest(u *url.URL, byteBody []byte) (*azcore.Request, error) {
	if !strings.HasSuffix(u.Path, "/") {
		u.Path = u.Path + "/"
	}
	u.Path = u.Path + "byte/nonAscii"
	req := azcore.NewRequest(http.MethodPut, *u)
	req.Header = http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}
	err := req.SetBody(azcore.NopCloser(bytes.NewReader(byteBody)))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutNonASCIIHandleResponse handles the GetEmpty response.
func (Service) PutNonASCIIHandleResponse(resp *azcore.Response) (*ByteArray, error) {
	return &ByteArray{Value: &resp.Payload}, nil // TODO what does this endpoint actually return?
}
