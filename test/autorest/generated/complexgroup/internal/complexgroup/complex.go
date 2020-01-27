// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"net/http"
	"net/url"
	"path"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// ComplexClient ..
type ComplexClient struct{}

// GetValidCreateRequest creates the GetValid request.
func (ComplexClient) GetValidCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/valid")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetValidHandleResponse handles the GetValid response.
func (ComplexClient) GetValidHandleResponse(resp *azcore.Response) (*GetValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetValidResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Basic)
}

// PutValidCreateRequest creates the PutValid request.
func (ComplexClient) PutValidCreateRequest(u url.URL, basicBody Basic) (*azcore.Request, error) {
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
func (ComplexClient) PutValidHandleResponse(resp *azcore.Response) (*PutValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PutValidResponse{StatusCode: resp.StatusCode}, nil
}

// GetInvalidCreateRequest creates the GetValid request.
func (ComplexClient) GetInvalidCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/invalid")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetInvalidHandleResponse handles the GetValid response.
func (ComplexClient) GetInvalidHandleResponse(resp *azcore.Response) (*GetInvalidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetInvalidResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Basic)
}

// GetEmptyCreateRequest creates the GetEmpty request.
func (ComplexClient) GetEmptyCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/empty")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetEmptyHandleResponse handles the GetEmpty response.
func (ComplexClient) GetEmptyHandleResponse(resp *azcore.Response) (*GetEmptyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetEmptyResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Basic)
}

// GetNullCreateRequest creates the GetNull request.
func (ComplexClient) GetNullCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/null")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetNullHandleResponse handles the GetNull response.
func (ComplexClient) GetNullHandleResponse(resp *azcore.Response) (*GetNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetNullResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Basic)
}

// GetNotProvidedCreateRequest creates the GetNotProvided request.
func (ComplexClient) GetNotProvidedCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/basic/notprovided")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetNotProvidedHandleResponse handles the GetNotProvided response.
func (ComplexClient) GetNotProvidedHandleResponse(resp *azcore.Response) (*GetNotProvidedResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetNotProvidedResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Basic)
}

// GetIntCreateRequest creates the GetInt request.
func (ComplexClient) GetIntCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/integer")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetIntHandleResponse handles the GetInt response.
func (ComplexClient) GetIntHandleResponse(resp *azcore.Response) (*GetIntResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetIntResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.IntWrapper)
}

// PutIntCreateRequest creates the PutInt request.
func (ComplexClient) PutIntCreateRequest(u url.URL, complexBody IntWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/integer")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutIntHandleResponse handles the PutInt response.
func (ComplexClient) PutIntHandleResponse(resp *azcore.Response) (*PutIntResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutIntResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.IntWrapper)
}

// GetLongCreateRequest creates the GetLong request.
func (ComplexClient) GetLongCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/long")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetLongHandleResponse handles the GetLong response.
func (ComplexClient) GetLongHandleResponse(resp *azcore.Response) (*GetLongResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetLongResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.LongWrapper)
}

// PutLongCreateRequest creates the PutLong request.
func (ComplexClient) PutLongCreateRequest(u url.URL, complexBody LongWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/long")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutLongHandleResponse handles the PutLong response.
func (ComplexClient) PutLongHandleResponse(resp *azcore.Response) (*PutLongResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutLongResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.LongWrapper)
}

// GetFloatCreateRequest creates the GetFloat request.
func (ComplexClient) GetFloatCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/float")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetFloatHandleResponse handles the GetFloat response.
func (ComplexClient) GetFloatHandleResponse(resp *azcore.Response) (*GetFloatResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetFloatResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.FloatWrapper)
}

// PutFloatCreateRequest creates the PutFloat request.
func (ComplexClient) PutFloatCreateRequest(u url.URL, complexBody FloatWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/float")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutFloatHandleResponse handles the PutFloat response.
func (ComplexClient) PutFloatHandleResponse(resp *azcore.Response) (*PutFloatResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutFloatResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.FloatWrapper)
}

// GetDoubleCreateRequest creates the GetDouble request.
func (ComplexClient) GetDoubleCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/double")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetDoubleHandleResponse handles the GetDouble response.
func (ComplexClient) GetDoubleHandleResponse(resp *azcore.Response) (*GetDoubleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetDoubleResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.DoubleWrapper)
}

// PutDoubleCreateRequest creates the PutDouble request.
func (ComplexClient) PutDoubleCreateRequest(u url.URL, complexBody DoubleWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/double")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutDoubleHandleResponse handles the PutDouble response.
func (ComplexClient) PutDoubleHandleResponse(resp *azcore.Response) (*PutDoubleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutDoubleResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.DoubleWrapper)
}

// GetBoolCreateRequest creates the GetBool request.
func (ComplexClient) GetBoolCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/bool")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetBoolHandleResponse handles the GetBool response.
func (ComplexClient) GetBoolHandleResponse(resp *azcore.Response) (*GetBoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetBoolResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.BooleanWrapper)
}

// PutBoolCreateRequest creates the PutBool request.
func (ComplexClient) PutBoolCreateRequest(u url.URL, complexBody BooleanWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/bool")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutBoolHandleResponse handles the PutBool response.
func (ComplexClient) PutBoolHandleResponse(resp *azcore.Response) (*PutBoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutBoolResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.BooleanWrapper)
}

// GetStringCreateRequest creates the GetString request.
func (ComplexClient) GetStringCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/string")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetStringHandleResponse handles the GetString response.
func (ComplexClient) GetStringHandleResponse(resp *azcore.Response) (*GetStringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetStringResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.StringWrapper)
}

// PutStringCreateRequest creates the PutString request.
func (ComplexClient) PutStringCreateRequest(u url.URL, complexBody StringWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/string")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutStringHandleResponse handles the PutString response.
func (ComplexClient) PutStringHandleResponse(resp *azcore.Response) (*PutStringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutStringResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.StringWrapper)
}

// // GetDateCreateRequest creates the GetDate request.
// func (ComplexClient) GetDateCreateRequest(u url.URL) (*azcore.Request, error) {
// 	u.Path = path.Join(u.Path, "/complex/primitive/date")
// 	return azcore.NewRequest(http.MethodGet, u), nil
// }

// // GetDateHandleResponse handles the GetDate response.
// func (ComplexClient) GetDateHandleResponse(resp *azcore.Response) (*GetDateResponse, error) {
// 	if !resp.HasStatusCode(http.StatusOK) {
// 		return nil, newError(resp)
// 	}
// 	result := GetDateResponse{StatusCode: resp.StatusCode}
// 	return &result, resp.UnmarshalAsJSON(&result.DateWrapper)
// }

// // PutDateCreateRequest creates the PutDate request.
// func (ComplexClient) PutDateCreateRequest(u url.URL, complexBody DateWrapper) (*azcore.Request, error) {
// 	u.Path = path.Join(u.Path, "/complex/primitive/date")
// 	req := azcore.NewRequest(http.MethodPut, u)
// 	err := req.MarshalAsJSON(complexBody)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// TODO delete this
// 	// r, e := ioutil.ReadAll(req.Body)
// 	// if e != nil {
// 	// 	fmt.Println(e)
// 	// }
// 	// fmt.Println(string(r))
// 	return req, nil
// }

// // PutDateHandleResponse handles the PutDate response.
// func (ComplexClient) PutDateHandleResponse(resp *azcore.Response) (*PutDateResponse, error) {
// 	if !resp.HasStatusCode(http.StatusOK) {
// 		return nil, newError(resp)
// 	}
// 	result := PutDateResponse{StatusCode: resp.StatusCode}
// 	return &result, resp.UnmarshalAsJSON(&result.DateWrapper)
// }
