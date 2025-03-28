// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"arraygroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ArrayFloat32ValueServer is a fake server for instances of the arraygroup.ArrayFloat32ValueClient type.
type ArrayFloat32ValueServer struct {
	// Get is the fake for method ArrayFloat32ValueClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *arraygroup.ArrayFloat32ValueClientGetOptions) (resp azfake.Responder[arraygroup.ArrayFloat32ValueClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ArrayFloat32ValueClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body []float32, options *arraygroup.ArrayFloat32ValueClientPutOptions) (resp azfake.Responder[arraygroup.ArrayFloat32ValueClientPutResponse], errResp azfake.ErrorResponder)
}

// NewArrayFloat32ValueServerTransport creates a new instance of ArrayFloat32ValueServerTransport with the provided implementation.
// The returned ArrayFloat32ValueServerTransport instance is connected to an instance of arraygroup.ArrayFloat32ValueClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewArrayFloat32ValueServerTransport(srv *ArrayFloat32ValueServer) *ArrayFloat32ValueServerTransport {
	return &ArrayFloat32ValueServerTransport{srv: srv}
}

// ArrayFloat32ValueServerTransport connects instances of arraygroup.ArrayFloat32ValueClient to instances of ArrayFloat32ValueServer.
// Don't use this type directly, use NewArrayFloat32ValueServerTransport instead.
type ArrayFloat32ValueServerTransport struct {
	srv *ArrayFloat32ValueServer
}

// Do implements the policy.Transporter interface for ArrayFloat32ValueServerTransport.
func (a *ArrayFloat32ValueServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *ArrayFloat32ValueServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if arrayFloat32ValueServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = arrayFloat32ValueServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ArrayFloat32ValueClient.Get":
				res.resp, res.err = a.dispatchGet(req)
			case "ArrayFloat32ValueClient.Put":
				res.resp, res.err = a.dispatchPut(req)
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

func (a *ArrayFloat32ValueServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := a.srv.Get(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Float32Array, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ArrayFloat32ValueServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if a.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[[]float32](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Put(req.Context(), body, nil)
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

// set this to conditionally intercept incoming requests to ArrayFloat32ValueServerTransport
var arrayFloat32ValueServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
