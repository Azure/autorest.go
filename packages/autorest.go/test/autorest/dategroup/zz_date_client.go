// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package dategroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"time"
)

// DateClient contains the methods for the Date group.
// Don't use this type directly, use a constructor function instead.
type DateClient struct {
	internal *azcore.Client
	endpoint string
}

// GetInvalidDate - Get invalid date value
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - DateClientGetInvalidDateOptions contains the optional parameters for the DateClient.GetInvalidDate method.
func (client *DateClient) GetInvalidDate(ctx context.Context, options *DateClientGetInvalidDateOptions) (DateClientGetInvalidDateResponse, error) {
	var err error
	const operationName = "DateClient.GetInvalidDate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getInvalidDateCreateRequest(ctx, options)
	if err != nil {
		return DateClientGetInvalidDateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DateClientGetInvalidDateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DateClientGetInvalidDateResponse{}, err
	}
	resp, err := client.getInvalidDateHandleResponse(httpResp)
	return resp, err
}

// getInvalidDateCreateRequest creates the GetInvalidDate request.
func (client *DateClient) getInvalidDateCreateRequest(ctx context.Context, _ *DateClientGetInvalidDateOptions) (*policy.Request, error) {
	urlPath := "/date/invaliddate"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getInvalidDateHandleResponse handles the GetInvalidDate response.
func (client *DateClient) getInvalidDateHandleResponse(resp *http.Response) (DateClientGetInvalidDateResponse, error) {
	result := DateClientGetInvalidDateResponse{}
	var aux *dateType
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return DateClientGetInvalidDateResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// GetMaxDate - Get max date value 9999-12-31
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - DateClientGetMaxDateOptions contains the optional parameters for the DateClient.GetMaxDate method.
func (client *DateClient) GetMaxDate(ctx context.Context, options *DateClientGetMaxDateOptions) (DateClientGetMaxDateResponse, error) {
	var err error
	const operationName = "DateClient.GetMaxDate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getMaxDateCreateRequest(ctx, options)
	if err != nil {
		return DateClientGetMaxDateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DateClientGetMaxDateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DateClientGetMaxDateResponse{}, err
	}
	resp, err := client.getMaxDateHandleResponse(httpResp)
	return resp, err
}

// getMaxDateCreateRequest creates the GetMaxDate request.
func (client *DateClient) getMaxDateCreateRequest(ctx context.Context, _ *DateClientGetMaxDateOptions) (*policy.Request, error) {
	urlPath := "/date/max"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getMaxDateHandleResponse handles the GetMaxDate response.
func (client *DateClient) getMaxDateHandleResponse(resp *http.Response) (DateClientGetMaxDateResponse, error) {
	result := DateClientGetMaxDateResponse{}
	var aux *dateType
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return DateClientGetMaxDateResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// GetMinDate - Get min date value 0000-01-01
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - DateClientGetMinDateOptions contains the optional parameters for the DateClient.GetMinDate method.
func (client *DateClient) GetMinDate(ctx context.Context, options *DateClientGetMinDateOptions) (DateClientGetMinDateResponse, error) {
	var err error
	const operationName = "DateClient.GetMinDate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getMinDateCreateRequest(ctx, options)
	if err != nil {
		return DateClientGetMinDateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DateClientGetMinDateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DateClientGetMinDateResponse{}, err
	}
	resp, err := client.getMinDateHandleResponse(httpResp)
	return resp, err
}

// getMinDateCreateRequest creates the GetMinDate request.
func (client *DateClient) getMinDateCreateRequest(ctx context.Context, _ *DateClientGetMinDateOptions) (*policy.Request, error) {
	urlPath := "/date/min"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getMinDateHandleResponse handles the GetMinDate response.
func (client *DateClient) getMinDateHandleResponse(resp *http.Response) (DateClientGetMinDateResponse, error) {
	result := DateClientGetMinDateResponse{}
	var aux *dateType
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return DateClientGetMinDateResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// GetNull - Get null date value
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - DateClientGetNullOptions contains the optional parameters for the DateClient.GetNull method.
func (client *DateClient) GetNull(ctx context.Context, options *DateClientGetNullOptions) (DateClientGetNullResponse, error) {
	var err error
	const operationName = "DateClient.GetNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return DateClientGetNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DateClientGetNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DateClientGetNullResponse{}, err
	}
	resp, err := client.getNullHandleResponse(httpResp)
	return resp, err
}

// getNullCreateRequest creates the GetNull request.
func (client *DateClient) getNullCreateRequest(ctx context.Context, _ *DateClientGetNullOptions) (*policy.Request, error) {
	urlPath := "/date/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *DateClient) getNullHandleResponse(resp *http.Response) (DateClientGetNullResponse, error) {
	result := DateClientGetNullResponse{}
	var aux *dateType
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return DateClientGetNullResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// GetOverflowDate - Get overflow date value
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - DateClientGetOverflowDateOptions contains the optional parameters for the DateClient.GetOverflowDate method.
func (client *DateClient) GetOverflowDate(ctx context.Context, options *DateClientGetOverflowDateOptions) (DateClientGetOverflowDateResponse, error) {
	var err error
	const operationName = "DateClient.GetOverflowDate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getOverflowDateCreateRequest(ctx, options)
	if err != nil {
		return DateClientGetOverflowDateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DateClientGetOverflowDateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DateClientGetOverflowDateResponse{}, err
	}
	resp, err := client.getOverflowDateHandleResponse(httpResp)
	return resp, err
}

// getOverflowDateCreateRequest creates the GetOverflowDate request.
func (client *DateClient) getOverflowDateCreateRequest(ctx context.Context, _ *DateClientGetOverflowDateOptions) (*policy.Request, error) {
	urlPath := "/date/overflowdate"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getOverflowDateHandleResponse handles the GetOverflowDate response.
func (client *DateClient) getOverflowDateHandleResponse(resp *http.Response) (DateClientGetOverflowDateResponse, error) {
	result := DateClientGetOverflowDateResponse{}
	var aux *dateType
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return DateClientGetOverflowDateResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// GetUnderflowDate - Get underflow date value
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - DateClientGetUnderflowDateOptions contains the optional parameters for the DateClient.GetUnderflowDate method.
func (client *DateClient) GetUnderflowDate(ctx context.Context, options *DateClientGetUnderflowDateOptions) (DateClientGetUnderflowDateResponse, error) {
	var err error
	const operationName = "DateClient.GetUnderflowDate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getUnderflowDateCreateRequest(ctx, options)
	if err != nil {
		return DateClientGetUnderflowDateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DateClientGetUnderflowDateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DateClientGetUnderflowDateResponse{}, err
	}
	resp, err := client.getUnderflowDateHandleResponse(httpResp)
	return resp, err
}

// getUnderflowDateCreateRequest creates the GetUnderflowDate request.
func (client *DateClient) getUnderflowDateCreateRequest(ctx context.Context, _ *DateClientGetUnderflowDateOptions) (*policy.Request, error) {
	urlPath := "/date/underflowdate"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getUnderflowDateHandleResponse handles the GetUnderflowDate response.
func (client *DateClient) getUnderflowDateHandleResponse(resp *http.Response) (DateClientGetUnderflowDateResponse, error) {
	result := DateClientGetUnderflowDateResponse{}
	var aux *dateType
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return DateClientGetUnderflowDateResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// PutMaxDate - Put max date value 9999-12-31
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - dateBody - date body
//   - options - DateClientPutMaxDateOptions contains the optional parameters for the DateClient.PutMaxDate method.
func (client *DateClient) PutMaxDate(ctx context.Context, dateBody time.Time, options *DateClientPutMaxDateOptions) (DateClientPutMaxDateResponse, error) {
	var err error
	const operationName = "DateClient.PutMaxDate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putMaxDateCreateRequest(ctx, dateBody, options)
	if err != nil {
		return DateClientPutMaxDateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DateClientPutMaxDateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DateClientPutMaxDateResponse{}, err
	}
	return DateClientPutMaxDateResponse{}, nil
}

// putMaxDateCreateRequest creates the PutMaxDate request.
func (client *DateClient) putMaxDateCreateRequest(ctx context.Context, dateBody time.Time, _ *DateClientPutMaxDateOptions) (*policy.Request, error) {
	urlPath := "/date/max"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, dateType(dateBody)); err != nil {
		return nil, err
	}
	return req, nil
}

// PutMinDate - Put min date value 0000-01-01
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - dateBody - date body
//   - options - DateClientPutMinDateOptions contains the optional parameters for the DateClient.PutMinDate method.
func (client *DateClient) PutMinDate(ctx context.Context, dateBody time.Time, options *DateClientPutMinDateOptions) (DateClientPutMinDateResponse, error) {
	var err error
	const operationName = "DateClient.PutMinDate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putMinDateCreateRequest(ctx, dateBody, options)
	if err != nil {
		return DateClientPutMinDateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DateClientPutMinDateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DateClientPutMinDateResponse{}, err
	}
	return DateClientPutMinDateResponse{}, nil
}

// putMinDateCreateRequest creates the PutMinDate request.
func (client *DateClient) putMinDateCreateRequest(ctx context.Context, dateBody time.Time, _ *DateClientPutMinDateOptions) (*policy.Request, error) {
	urlPath := "/date/min"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, dateType(dateBody)); err != nil {
		return nil, err
	}
	return req, nil
}
