// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgroup

import (
	"net/http"
	"net/url"
	"path"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// Service ..
type Service struct{}

// GetMBCSCreateRequest creates the GetMBCS request.
func (Service) GetMBCSCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/string/mbcs")
	// TODO: this makes two copies
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetMBCSHandleResponse handles the GetMBCS response.
func (Service) GetMBCSHandleResponse(resp *azcore.Response) (*GetMBCSResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetMBCSResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// PutMBCSCreateRequest creates the PutMBCS request.
func (Service) PutMBCSCreateRequest(u url.URL, stringBody []string) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/string/mbcs")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(stringBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutMBCSHandleResponse handles the PutMBCS response.
func (Service) PutMBCSHandleResponse(resp *azcore.Response) (*PutMBCSResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PutMBCSResponse{StatusCode: resp.StatusCode}, nil
}
