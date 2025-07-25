// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// HTTPRedirectsClient contains the methods for the HTTPRedirects group.
// Don't use this type directly, use a constructor function instead.
type HTTPRedirectsClient struct {
	internal *azcore.Client
	endpoint string
}

// Delete307 - Delete redirected with 307, resulting in a 200 after redirect
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientDelete307Options contains the optional parameters for the HTTPRedirectsClient.Delete307 method.
func (client *HTTPRedirectsClient) Delete307(ctx context.Context, options *HTTPRedirectsClientDelete307Options) (HTTPRedirectsClientDelete307Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Delete307"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.delete307CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientDelete307Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientDelete307Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientDelete307Response{}, err
	}
	return HTTPRedirectsClientDelete307Response{}, nil
}

// delete307CreateRequest creates the Delete307 request.
func (client *HTTPRedirectsClient) delete307CreateRequest(ctx context.Context, options *HTTPRedirectsClientDelete307Options) (*policy.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.BooleanValue != nil {
		if err := runtime.MarshalAsJSON(req, true); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// Get300 - Return 300 status code and redirect to /http/success/200
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientGet300Options contains the optional parameters for the HTTPRedirectsClient.Get300 method.
func (client *HTTPRedirectsClient) Get300(ctx context.Context, options *HTTPRedirectsClientGet300Options) (HTTPRedirectsClientGet300Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Get300"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.get300CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientGet300Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientGet300Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusMultipleChoices) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientGet300Response{}, err
	}
	resp, err := client.get300HandleResponse(httpResp)
	return resp, err
}

// get300CreateRequest creates the Get300 request.
func (client *HTTPRedirectsClient) get300CreateRequest(ctx context.Context, _ *HTTPRedirectsClientGet300Options) (*policy.Request, error) {
	urlPath := "/http/redirect/300"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// get300HandleResponse handles the Get300 response.
func (client *HTTPRedirectsClient) get300HandleResponse(resp *http.Response) (HTTPRedirectsClientGet300Response, error) {
	result := HTTPRedirectsClientGet300Response{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.StringArray); err != nil {
		return HTTPRedirectsClientGet300Response{}, err
	}
	return result, nil
}

// Get301 - Return 301 status code and redirect to /http/success/200
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientGet301Options contains the optional parameters for the HTTPRedirectsClient.Get301 method.
func (client *HTTPRedirectsClient) Get301(ctx context.Context, options *HTTPRedirectsClientGet301Options) (HTTPRedirectsClientGet301Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Get301"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.get301CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientGet301Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientGet301Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientGet301Response{}, err
	}
	return HTTPRedirectsClientGet301Response{}, nil
}

// get301CreateRequest creates the Get301 request.
func (client *HTTPRedirectsClient) get301CreateRequest(ctx context.Context, _ *HTTPRedirectsClientGet301Options) (*policy.Request, error) {
	urlPath := "/http/redirect/301"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get302 - Return 302 status code and redirect to /http/success/200
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientGet302Options contains the optional parameters for the HTTPRedirectsClient.Get302 method.
func (client *HTTPRedirectsClient) Get302(ctx context.Context, options *HTTPRedirectsClientGet302Options) (HTTPRedirectsClientGet302Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Get302"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.get302CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientGet302Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientGet302Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientGet302Response{}, err
	}
	return HTTPRedirectsClientGet302Response{}, nil
}

// get302CreateRequest creates the Get302 request.
func (client *HTTPRedirectsClient) get302CreateRequest(ctx context.Context, _ *HTTPRedirectsClientGet302Options) (*policy.Request, error) {
	urlPath := "/http/redirect/302"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get307 - Redirect get with 307, resulting in a 200 success
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientGet307Options contains the optional parameters for the HTTPRedirectsClient.Get307 method.
func (client *HTTPRedirectsClient) Get307(ctx context.Context, options *HTTPRedirectsClientGet307Options) (HTTPRedirectsClientGet307Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Get307"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.get307CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientGet307Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientGet307Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientGet307Response{}, err
	}
	return HTTPRedirectsClientGet307Response{}, nil
}

// get307CreateRequest creates the Get307 request.
func (client *HTTPRedirectsClient) get307CreateRequest(ctx context.Context, _ *HTTPRedirectsClientGet307Options) (*policy.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Head300 - Return 300 status code and redirect to /http/success/200
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientHead300Options contains the optional parameters for the HTTPRedirectsClient.Head300 method.
func (client *HTTPRedirectsClient) Head300(ctx context.Context, options *HTTPRedirectsClientHead300Options) (HTTPRedirectsClientHead300Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Head300"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.head300CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientHead300Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientHead300Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusMultipleChoices) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientHead300Response{}, err
	}
	resp, err := client.head300HandleResponse(httpResp)
	return resp, err
}

// head300CreateRequest creates the Head300 request.
func (client *HTTPRedirectsClient) head300CreateRequest(ctx context.Context, _ *HTTPRedirectsClientHead300Options) (*policy.Request, error) {
	urlPath := "/http/redirect/300"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// head300HandleResponse handles the Head300 response.
func (client *HTTPRedirectsClient) head300HandleResponse(resp *http.Response) (HTTPRedirectsClientHead300Response, error) {
	result := HTTPRedirectsClientHead300Response{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	return result, nil
}

// Head301 - Return 301 status code and redirect to /http/success/200
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientHead301Options contains the optional parameters for the HTTPRedirectsClient.Head301 method.
func (client *HTTPRedirectsClient) Head301(ctx context.Context, options *HTTPRedirectsClientHead301Options) (HTTPRedirectsClientHead301Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Head301"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.head301CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientHead301Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientHead301Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientHead301Response{}, err
	}
	return HTTPRedirectsClientHead301Response{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// head301CreateRequest creates the Head301 request.
func (client *HTTPRedirectsClient) head301CreateRequest(ctx context.Context, _ *HTTPRedirectsClientHead301Options) (*policy.Request, error) {
	urlPath := "/http/redirect/301"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Head302 - Return 302 status code and redirect to /http/success/200
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientHead302Options contains the optional parameters for the HTTPRedirectsClient.Head302 method.
func (client *HTTPRedirectsClient) Head302(ctx context.Context, options *HTTPRedirectsClientHead302Options) (HTTPRedirectsClientHead302Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Head302"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.head302CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientHead302Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientHead302Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientHead302Response{}, err
	}
	return HTTPRedirectsClientHead302Response{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// head302CreateRequest creates the Head302 request.
func (client *HTTPRedirectsClient) head302CreateRequest(ctx context.Context, _ *HTTPRedirectsClientHead302Options) (*policy.Request, error) {
	urlPath := "/http/redirect/302"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Head307 - Redirect with 307, resulting in a 200 success
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientHead307Options contains the optional parameters for the HTTPRedirectsClient.Head307 method.
func (client *HTTPRedirectsClient) Head307(ctx context.Context, options *HTTPRedirectsClientHead307Options) (HTTPRedirectsClientHead307Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Head307"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.head307CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientHead307Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientHead307Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientHead307Response{}, err
	}
	return HTTPRedirectsClientHead307Response{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// head307CreateRequest creates the Head307 request.
func (client *HTTPRedirectsClient) head307CreateRequest(ctx context.Context, _ *HTTPRedirectsClientHead307Options) (*policy.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Options307 - options redirected with 307, resulting in a 200 after redirect
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientOptions307Options contains the optional parameters for the HTTPRedirectsClient.Options307
//     method.
func (client *HTTPRedirectsClient) Options307(ctx context.Context, options *HTTPRedirectsClientOptions307Options) (HTTPRedirectsClientOptions307Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Options307"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.options307CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientOptions307Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientOptions307Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientOptions307Response{}, err
	}
	return HTTPRedirectsClientOptions307Response{}, nil
}

// options307CreateRequest creates the Options307 request.
func (client *HTTPRedirectsClient) options307CreateRequest(ctx context.Context, _ *HTTPRedirectsClientOptions307Options) (*policy.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := runtime.NewRequest(ctx, http.MethodOptions, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Patch302 - Patch true Boolean value in request returns 302. This request should not be automatically redirected, but should
// return the received 302 to the caller for evaluation
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientPatch302Options contains the optional parameters for the HTTPRedirectsClient.Patch302 method.
func (client *HTTPRedirectsClient) Patch302(ctx context.Context, options *HTTPRedirectsClientPatch302Options) (HTTPRedirectsClientPatch302Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Patch302"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.patch302CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientPatch302Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientPatch302Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusFound) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientPatch302Response{}, err
	}
	resp, err := client.patch302HandleResponse(httpResp)
	return resp, err
}

// patch302CreateRequest creates the Patch302 request.
func (client *HTTPRedirectsClient) patch302CreateRequest(ctx context.Context, _ *HTTPRedirectsClientPatch302Options) (*policy.Request, error) {
	urlPath := "/http/redirect/302"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// patch302HandleResponse handles the Patch302 response.
func (client *HTTPRedirectsClient) patch302HandleResponse(resp *http.Response) (HTTPRedirectsClientPatch302Response, error) {
	result := HTTPRedirectsClientPatch302Response{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	return result, nil
}

// Patch307 - Patch redirected with 307, resulting in a 200 after redirect
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientPatch307Options contains the optional parameters for the HTTPRedirectsClient.Patch307 method.
func (client *HTTPRedirectsClient) Patch307(ctx context.Context, options *HTTPRedirectsClientPatch307Options) (HTTPRedirectsClientPatch307Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Patch307"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.patch307CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientPatch307Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientPatch307Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientPatch307Response{}, err
	}
	return HTTPRedirectsClientPatch307Response{}, nil
}

// patch307CreateRequest creates the Patch307 request.
func (client *HTTPRedirectsClient) patch307CreateRequest(ctx context.Context, _ *HTTPRedirectsClientPatch307Options) (*policy.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Post303 - Post true Boolean value in request returns 303. This request should be automatically redirected usign a get,
// ultimately returning a 200 status code
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientPost303Options contains the optional parameters for the HTTPRedirectsClient.Post303 method.
func (client *HTTPRedirectsClient) Post303(ctx context.Context, options *HTTPRedirectsClientPost303Options) (HTTPRedirectsClientPost303Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Post303"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.post303CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientPost303Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientPost303Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusSeeOther) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientPost303Response{}, err
	}
	resp, err := client.post303HandleResponse(httpResp)
	return resp, err
}

// post303CreateRequest creates the Post303 request.
func (client *HTTPRedirectsClient) post303CreateRequest(ctx context.Context, options *HTTPRedirectsClientPost303Options) (*policy.Request, error) {
	urlPath := "/http/redirect/303"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.BooleanValue != nil {
		if err := runtime.MarshalAsJSON(req, true); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// post303HandleResponse handles the Post303 response.
func (client *HTTPRedirectsClient) post303HandleResponse(resp *http.Response) (HTTPRedirectsClientPost303Response, error) {
	result := HTTPRedirectsClientPost303Response{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	return result, nil
}

// Post307 - Post redirected with 307, resulting in a 200 after redirect
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientPost307Options contains the optional parameters for the HTTPRedirectsClient.Post307 method.
func (client *HTTPRedirectsClient) Post307(ctx context.Context, options *HTTPRedirectsClientPost307Options) (HTTPRedirectsClientPost307Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Post307"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.post307CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientPost307Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientPost307Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientPost307Response{}, err
	}
	return HTTPRedirectsClientPost307Response{}, nil
}

// post307CreateRequest creates the Post307 request.
func (client *HTTPRedirectsClient) post307CreateRequest(ctx context.Context, options *HTTPRedirectsClientPost307Options) (*policy.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.BooleanValue != nil {
		if err := runtime.MarshalAsJSON(req, true); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// Put301 - Put true Boolean value in request returns 301. This request should not be automatically redirected, but should
// return the received 301 to the caller for evaluation
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientPut301Options contains the optional parameters for the HTTPRedirectsClient.Put301 method.
func (client *HTTPRedirectsClient) Put301(ctx context.Context, options *HTTPRedirectsClientPut301Options) (HTTPRedirectsClientPut301Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Put301"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.put301CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientPut301Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientPut301Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusMovedPermanently) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientPut301Response{}, err
	}
	resp, err := client.put301HandleResponse(httpResp)
	return resp, err
}

// put301CreateRequest creates the Put301 request.
func (client *HTTPRedirectsClient) put301CreateRequest(ctx context.Context, _ *HTTPRedirectsClientPut301Options) (*policy.Request, error) {
	urlPath := "/http/redirect/301"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// put301HandleResponse handles the Put301 response.
func (client *HTTPRedirectsClient) put301HandleResponse(resp *http.Response) (HTTPRedirectsClientPut301Response, error) {
	result := HTTPRedirectsClientPut301Response{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	return result, nil
}

// Put307 - Put redirected with 307, resulting in a 200 after redirect
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPRedirectsClientPut307Options contains the optional parameters for the HTTPRedirectsClient.Put307 method.
func (client *HTTPRedirectsClient) Put307(ctx context.Context, options *HTTPRedirectsClientPut307Options) (HTTPRedirectsClientPut307Response, error) {
	var err error
	const operationName = "HTTPRedirectsClient.Put307"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.put307CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsClientPut307Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRedirectsClientPut307Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRedirectsClientPut307Response{}, err
	}
	return HTTPRedirectsClientPut307Response{}, nil
}

// put307CreateRequest creates the Put307 request.
func (client *HTTPRedirectsClient) put307CreateRequest(ctx context.Context, _ *HTTPRedirectsClientPut307Options) (*policy.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}
