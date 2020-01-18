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
	// TODO check if we need to include content type header
	return req, nil
}

// PutValidHandleResponse handles the PutValid response.
func (Service) PutValidHandleResponse(resp *azcore.Response) (*PutValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PutValidResponse{StatusCode: resp.StatusCode}, nil
}
