// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"emptygroup"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// EmptyServer is a fake server for instances of the emptygroup.EmptyClient type.
type EmptyServer struct {
	// GetEmpty is the fake for method EmptyClient.GetEmpty
	// HTTP status codes to indicate success: http.StatusOK
	GetEmpty func(ctx context.Context, options *emptygroup.GetEmptyOptions) (resp azfake.Responder[emptygroup.GetEmptyResponse], errResp azfake.ErrorResponder)

	// PostRoundTripEmpty is the fake for method EmptyClient.PostRoundTripEmpty
	// HTTP status codes to indicate success: http.StatusOK
	PostRoundTripEmpty func(ctx context.Context, body emptygroup.EmptyInputOutput, options *emptygroup.PostRoundTripEmptyOptions) (resp azfake.Responder[emptygroup.PostRoundTripEmptyResponse], errResp azfake.ErrorResponder)

	// PutEmpty is the fake for method EmptyClient.PutEmpty
	// HTTP status codes to indicate success: http.StatusNoContent
	PutEmpty func(ctx context.Context, input emptygroup.EmptyInput, options *emptygroup.PutEmptyOptions) (resp azfake.Responder[emptygroup.PutEmptyResponse], errResp azfake.ErrorResponder)
}

// NewEmptyServerTransport creates a new instance of EmptyServerTransport with the provided implementation.
// The returned EmptyServerTransport instance is connected to an instance of emptygroup.EmptyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEmptyServerTransport(srv *EmptyServer) *EmptyServerTransport {
	return &EmptyServerTransport{srv: srv}
}

// EmptyServerTransport connects instances of emptygroup.EmptyClient to instances of EmptyServer.
// Don't use this type directly, use NewEmptyServerTransport instead.
type EmptyServerTransport struct {
	srv *EmptyServer
}

// Do implements the policy.Transporter interface for EmptyServerTransport.
func (e *EmptyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "EmptyClient.GetEmpty":
		resp, err = e.dispatchGetEmpty(req)
	case "EmptyClient.PostRoundTripEmpty":
		resp, err = e.dispatchPostRoundTripEmpty(req)
	case "EmptyClient.PutEmpty":
		resp, err = e.dispatchPutEmpty(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EmptyServerTransport) dispatchGetEmpty(req *http.Request) (*http.Response, error) {
	if e.srv.GetEmpty == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEmpty not implemented")}
	}
	respr, errRespr := e.srv.GetEmpty(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EmptyOutput, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EmptyServerTransport) dispatchPostRoundTripEmpty(req *http.Request) (*http.Response, error) {
	if e.srv.PostRoundTripEmpty == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostRoundTripEmpty not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[emptygroup.EmptyInputOutput](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PostRoundTripEmpty(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EmptyInputOutput, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EmptyServerTransport) dispatchPutEmpty(req *http.Request) (*http.Response, error) {
	if e.srv.PutEmpty == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutEmpty not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[emptygroup.EmptyInput](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PutEmpty(req.Context(), body, nil)
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