// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"bytesgroup"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
)

// QueryServer is a fake server for instances of the bytesgroup.QueryClient type.
type QueryServer struct {
	// Base64 is the fake for method QueryClient.Base64
	// HTTP status codes to indicate success: http.StatusNoContent
	Base64 func(ctx context.Context, value []byte, options *bytesgroup.QueryClientBase64Options) (resp azfake.Responder[bytesgroup.QueryClientBase64Response], errResp azfake.ErrorResponder)

	// Base64URL is the fake for method QueryClient.Base64URL
	// HTTP status codes to indicate success: http.StatusNoContent
	Base64URL func(ctx context.Context, value []byte, options *bytesgroup.QueryClientBase64URLOptions) (resp azfake.Responder[bytesgroup.QueryClientBase64URLResponse], errResp azfake.ErrorResponder)

	// Base64URLArray is the fake for method QueryClient.Base64URLArray
	// HTTP status codes to indicate success: http.StatusNoContent
	Base64URLArray func(ctx context.Context, value [][]byte, options *bytesgroup.QueryClientBase64URLArrayOptions) (resp azfake.Responder[bytesgroup.QueryClientBase64URLArrayResponse], errResp azfake.ErrorResponder)

	// Default is the fake for method QueryClient.Default
	// HTTP status codes to indicate success: http.StatusNoContent
	Default func(ctx context.Context, value []byte, options *bytesgroup.QueryClientDefaultOptions) (resp azfake.Responder[bytesgroup.QueryClientDefaultResponse], errResp azfake.ErrorResponder)
}

// NewQueryServerTransport creates a new instance of QueryServerTransport with the provided implementation.
// The returned QueryServerTransport instance is connected to an instance of bytesgroup.QueryClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewQueryServerTransport(srv *QueryServer) *QueryServerTransport {
	return &QueryServerTransport{srv: srv}
}

// QueryServerTransport connects instances of bytesgroup.QueryClient to instances of QueryServer.
// Don't use this type directly, use NewQueryServerTransport instead.
type QueryServerTransport struct {
	srv *QueryServer
}

// Do implements the policy.Transporter interface for QueryServerTransport.
func (q *QueryServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return q.dispatchToMethodFake(req, method)
}

func (q *QueryServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "QueryClient.Base64":
		resp, err = q.dispatchBase64(req)
	case "QueryClient.Base64URL":
		resp, err = q.dispatchBase64URL(req)
	case "QueryClient.Base64URLArray":
		resp, err = q.dispatchBase64URLArray(req)
	case "QueryClient.Default":
		resp, err = q.dispatchDefault(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (q *QueryServerTransport) dispatchBase64(req *http.Request) (*http.Response, error) {
	if q.srv.Base64 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Base64 not implemented")}
	}
	qp := req.URL.Query()
	valueUnescaped, err := url.QueryUnescape(qp.Get("value"))
	if err != nil {
		return nil, err
	}
	valueParam, err := base64.StdEncoding.DecodeString(valueUnescaped)
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.Base64(req.Context(), valueParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QueryServerTransport) dispatchBase64URL(req *http.Request) (*http.Response, error) {
	if q.srv.Base64URL == nil {
		return nil, &nonRetriableError{errors.New("fake for method Base64URL not implemented")}
	}
	qp := req.URL.Query()
	valueUnescaped, err := url.QueryUnescape(qp.Get("value"))
	if err != nil {
		return nil, err
	}
	valueParam, err := base64.URLEncoding.DecodeString(valueUnescaped)
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.Base64URL(req.Context(), valueParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QueryServerTransport) dispatchBase64URLArray(req *http.Request) (*http.Response, error) {
	if q.srv.Base64URLArray == nil {
		return nil, &nonRetriableError{errors.New("fake for method Base64URLArray not implemented")}
	}
	qp := req.URL.Query()
	valueUnescaped, err := url.QueryUnescape(qp.Get("value"))
	if err != nil {
		return nil, err
	}
	valueElements := splitHelper(valueUnescaped, ",")
	valueParam := make([][]byte, len(valueElements))
	for i := 0; i < len(valueElements); i++ {
		parsedURL, parseErr := base64.URLEncoding.DecodeString(valueElements[i])
		if parseErr != nil {
			return nil, parseErr
		}
		valueParam[i] = []byte(parsedURL)
	}
	respr, errRespr := q.srv.Base64URLArray(req.Context(), valueParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QueryServerTransport) dispatchDefault(req *http.Request) (*http.Response, error) {
	if q.srv.Default == nil {
		return nil, &nonRetriableError{errors.New("fake for method Default not implemented")}
	}
	qp := req.URL.Query()
	valueUnescaped, err := url.QueryUnescape(qp.Get("value"))
	if err != nil {
		return nil, err
	}
	valueParam, err := base64.StdEncoding.DecodeString(valueUnescaped)
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.Default(req.Context(), valueParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
