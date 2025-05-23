// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"extensiblegroup"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ExtensibleStringServer is a fake server for instances of the extensiblegroup.ExtensibleStringClient type.
type ExtensibleStringServer struct {
	// GetKnownValue is the fake for method ExtensibleStringClient.GetKnownValue
	// HTTP status codes to indicate success: http.StatusOK
	GetKnownValue func(ctx context.Context, options *extensiblegroup.ExtensibleStringClientGetKnownValueOptions) (resp azfake.Responder[extensiblegroup.ExtensibleStringClientGetKnownValueResponse], errResp azfake.ErrorResponder)

	// GetUnknownValue is the fake for method ExtensibleStringClient.GetUnknownValue
	// HTTP status codes to indicate success: http.StatusOK
	GetUnknownValue func(ctx context.Context, options *extensiblegroup.ExtensibleStringClientGetUnknownValueOptions) (resp azfake.Responder[extensiblegroup.ExtensibleStringClientGetUnknownValueResponse], errResp azfake.ErrorResponder)

	// PutKnownValue is the fake for method ExtensibleStringClient.PutKnownValue
	// HTTP status codes to indicate success: http.StatusNoContent
	PutKnownValue func(ctx context.Context, body extensiblegroup.DaysOfWeekExtensibleEnum, options *extensiblegroup.ExtensibleStringClientPutKnownValueOptions) (resp azfake.Responder[extensiblegroup.ExtensibleStringClientPutKnownValueResponse], errResp azfake.ErrorResponder)

	// PutUnknownValue is the fake for method ExtensibleStringClient.PutUnknownValue
	// HTTP status codes to indicate success: http.StatusNoContent
	PutUnknownValue func(ctx context.Context, body extensiblegroup.DaysOfWeekExtensibleEnum, options *extensiblegroup.ExtensibleStringClientPutUnknownValueOptions) (resp azfake.Responder[extensiblegroup.ExtensibleStringClientPutUnknownValueResponse], errResp azfake.ErrorResponder)
}

// NewExtensibleStringServerTransport creates a new instance of ExtensibleStringServerTransport with the provided implementation.
// The returned ExtensibleStringServerTransport instance is connected to an instance of extensiblegroup.ExtensibleStringClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExtensibleStringServerTransport(srv *ExtensibleStringServer) *ExtensibleStringServerTransport {
	return &ExtensibleStringServerTransport{srv: srv}
}

// ExtensibleStringServerTransport connects instances of extensiblegroup.ExtensibleStringClient to instances of ExtensibleStringServer.
// Don't use this type directly, use NewExtensibleStringServerTransport instead.
type ExtensibleStringServerTransport struct {
	srv *ExtensibleStringServer
}

// Do implements the policy.Transporter interface for ExtensibleStringServerTransport.
func (e *ExtensibleStringServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *ExtensibleStringServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if extensibleStringServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = extensibleStringServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ExtensibleStringClient.GetKnownValue":
				res.resp, res.err = e.dispatchGetKnownValue(req)
			case "ExtensibleStringClient.GetUnknownValue":
				res.resp, res.err = e.dispatchGetUnknownValue(req)
			case "ExtensibleStringClient.PutKnownValue":
				res.resp, res.err = e.dispatchPutKnownValue(req)
			case "ExtensibleStringClient.PutUnknownValue":
				res.resp, res.err = e.dispatchPutUnknownValue(req)
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

func (e *ExtensibleStringServerTransport) dispatchGetKnownValue(req *http.Request) (*http.Response, error) {
	if e.srv.GetKnownValue == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetKnownValue not implemented")}
	}
	respr, errRespr := e.srv.GetKnownValue(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ContentType; val != nil {
		resp.Header.Set("content-type", "application/json")
	}
	return resp, nil
}

func (e *ExtensibleStringServerTransport) dispatchGetUnknownValue(req *http.Request) (*http.Response, error) {
	if e.srv.GetUnknownValue == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetUnknownValue not implemented")}
	}
	respr, errRespr := e.srv.GetUnknownValue(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ContentType; val != nil {
		resp.Header.Set("content-type", "application/json")
	}
	return resp, nil
}

func (e *ExtensibleStringServerTransport) dispatchPutKnownValue(req *http.Request) (*http.Response, error) {
	if e.srv.PutKnownValue == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutKnownValue not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[extensiblegroup.DaysOfWeekExtensibleEnum](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PutKnownValue(req.Context(), body, nil)
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

func (e *ExtensibleStringServerTransport) dispatchPutUnknownValue(req *http.Request) (*http.Response, error) {
	if e.srv.PutUnknownValue == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutUnknownValue not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[extensiblegroup.DaysOfWeekExtensibleEnum](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PutUnknownValue(req.Context(), body, nil)
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

// set this to conditionally intercept incoming requests to ExtensibleStringServerTransport
var extensibleStringServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
