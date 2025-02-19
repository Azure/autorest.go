// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"recursivegroup"
)

// RecursiveServer is a fake server for instances of the recursivegroup.RecursiveClient type.
type RecursiveServer struct {
	// Get is the fake for method RecursiveClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *recursivegroup.RecursiveClientGetOptions) (resp azfake.Responder[recursivegroup.RecursiveClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method RecursiveClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, input recursivegroup.Extension, options *recursivegroup.RecursiveClientPutOptions) (resp azfake.Responder[recursivegroup.RecursiveClientPutResponse], errResp azfake.ErrorResponder)
}

// NewRecursiveServerTransport creates a new instance of RecursiveServerTransport with the provided implementation.
// The returned RecursiveServerTransport instance is connected to an instance of recursivegroup.RecursiveClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRecursiveServerTransport(srv *RecursiveServer) *RecursiveServerTransport {
	return &RecursiveServerTransport{srv: srv}
}

// RecursiveServerTransport connects instances of recursivegroup.RecursiveClient to instances of RecursiveServer.
// Don't use this type directly, use NewRecursiveServerTransport instead.
type RecursiveServerTransport struct {
	srv *RecursiveServer
}

// Do implements the policy.Transporter interface for RecursiveServerTransport.
func (r *RecursiveServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *RecursiveServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if recursiveServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = recursiveServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "RecursiveClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "RecursiveClient.Put":
				res.resp, res.err = r.dispatchPut(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (r *RecursiveServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := r.srv.Get(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Extension, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RecursiveServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if r.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[recursivegroup.Extension](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Put(req.Context(), body, nil)
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

// set this to conditionally intercept incoming requests to RecursiveServerTransport
var recursiveServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
