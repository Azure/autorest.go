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

// TODO nil or null?
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

// GetLongCreateRequest creates the GetLong request.
func (Service) GetLongCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/long")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetLongHandleResponse handles the GetLong response.
func (Service) GetLongHandleResponse(resp *azcore.Response) (*GetLongResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetLongResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.LongWrapper)
}

// PutLongCreateRequest creates the PutLong request.
func (Service) PutLongCreateRequest(u url.URL, complexBody *LongWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/long")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutLongHandleResponse handles the PutLong response.
func (Service) PutLongHandleResponse(resp *azcore.Response) (*PutLongResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutLongResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.LongWrapper)
}

// GetFloatCreateRequest creates the GetFloat request.
func (Service) GetFloatCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/float")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetFloatHandleResponse handles the GetFloat response.
func (Service) GetFloatHandleResponse(resp *azcore.Response) (*GetFloatResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetFloatResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.FloatWrapper)
}

// PutFloatCreateRequest creates the PutFloat request.
func (Service) PutFloatCreateRequest(u url.URL, complexBody *FloatWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/float")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutFloatHandleResponse handles the PutFloat response.
func (Service) PutFloatHandleResponse(resp *azcore.Response) (*PutFloatResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutFloatResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.FloatWrapper)
}

// GetDoubleCreateRequest creates the GetDouble request.
func (Service) GetDoubleCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/double")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetDoubleHandleResponse handles the GetDouble response.
func (Service) GetDoubleHandleResponse(resp *azcore.Response) (*GetDoubleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetDoubleResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.DoubleWrapper)
}

// PutDoubleCreateRequest creates the PutDouble request.
func (Service) PutDoubleCreateRequest(u url.URL, complexBody *DoubleWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/double")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutDoubleHandleResponse handles the PutDouble response.
func (Service) PutDoubleHandleResponse(resp *azcore.Response) (*PutDoubleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutDoubleResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.DoubleWrapper)
}

// GetBoolCreateRequest creates the GetBool request.
func (Service) GetBoolCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/bool")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetBoolHandleResponse handles the GetBool response.
func (Service) GetBoolHandleResponse(resp *azcore.Response) (*GetBoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetBoolResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.BooleanWrapper)
}

// PutBoolCreateRequest creates the PutBool request.
func (Service) PutBoolCreateRequest(u url.URL, complexBody *BooleanWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/bool")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutBoolHandleResponse handles the PutBool response.
func (Service) PutBoolHandleResponse(resp *azcore.Response) (*PutBoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutBoolResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.BooleanWrapper)
}

// GetStringCreateRequest creates the GetString request.
func (Service) GetStringCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/string")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetStringHandleResponse handles the GetString response.
func (Service) GetStringHandleResponse(resp *azcore.Response) (*GetStringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetStringResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.StringWrapper)
}

// PutStringCreateRequest creates the PutString request.
func (Service) PutStringCreateRequest(u url.URL, complexBody *StringWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/string")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutStringHandleResponse handles the PutString response.
func (Service) PutStringHandleResponse(resp *azcore.Response) (*PutStringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutStringResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.StringWrapper)
}
