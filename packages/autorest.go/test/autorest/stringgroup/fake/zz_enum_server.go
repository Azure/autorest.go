// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/stringgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// EnumServer is a fake server for instances of the stringgroup.EnumClient type.
type EnumServer struct {
	// GetNotExpandable is the fake for method EnumClient.GetNotExpandable
	// HTTP status codes to indicate success: http.StatusOK
	GetNotExpandable func(ctx context.Context, options *stringgroup.EnumClientGetNotExpandableOptions) (resp azfake.Responder[stringgroup.EnumClientGetNotExpandableResponse], errResp azfake.ErrorResponder)

	// GetReferenced is the fake for method EnumClient.GetReferenced
	// HTTP status codes to indicate success: http.StatusOK
	GetReferenced func(ctx context.Context, options *stringgroup.EnumClientGetReferencedOptions) (resp azfake.Responder[stringgroup.EnumClientGetReferencedResponse], errResp azfake.ErrorResponder)

	// GetReferencedConstant is the fake for method EnumClient.GetReferencedConstant
	// HTTP status codes to indicate success: http.StatusOK
	GetReferencedConstant func(ctx context.Context, options *stringgroup.EnumClientGetReferencedConstantOptions) (resp azfake.Responder[stringgroup.EnumClientGetReferencedConstantResponse], errResp azfake.ErrorResponder)

	// PutNotExpandable is the fake for method EnumClient.PutNotExpandable
	// HTTP status codes to indicate success: http.StatusOK
	PutNotExpandable func(ctx context.Context, stringBody stringgroup.Colors, options *stringgroup.EnumClientPutNotExpandableOptions) (resp azfake.Responder[stringgroup.EnumClientPutNotExpandableResponse], errResp azfake.ErrorResponder)

	// PutReferenced is the fake for method EnumClient.PutReferenced
	// HTTP status codes to indicate success: http.StatusOK
	PutReferenced func(ctx context.Context, enumStringBody stringgroup.Colors, options *stringgroup.EnumClientPutReferencedOptions) (resp azfake.Responder[stringgroup.EnumClientPutReferencedResponse], errResp azfake.ErrorResponder)

	// PutReferencedConstant is the fake for method EnumClient.PutReferencedConstant
	// HTTP status codes to indicate success: http.StatusOK
	PutReferencedConstant func(ctx context.Context, enumStringBody stringgroup.RefColorConstant, options *stringgroup.EnumClientPutReferencedConstantOptions) (resp azfake.Responder[stringgroup.EnumClientPutReferencedConstantResponse], errResp azfake.ErrorResponder)
}

// NewEnumServerTransport creates a new instance of EnumServerTransport with the provided implementation.
// The returned EnumServerTransport instance is connected to an instance of stringgroup.EnumClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEnumServerTransport(srv *EnumServer) *EnumServerTransport {
	return &EnumServerTransport{srv: srv}
}

// EnumServerTransport connects instances of stringgroup.EnumClient to instances of EnumServer.
// Don't use this type directly, use NewEnumServerTransport instead.
type EnumServerTransport struct {
	srv *EnumServer
}

// Do implements the policy.Transporter interface for EnumServerTransport.
func (e *EnumServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *EnumServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if enumServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = enumServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "EnumClient.GetNotExpandable":
				res.resp, res.err = e.dispatchGetNotExpandable(req)
			case "EnumClient.GetReferenced":
				res.resp, res.err = e.dispatchGetReferenced(req)
			case "EnumClient.GetReferencedConstant":
				res.resp, res.err = e.dispatchGetReferencedConstant(req)
			case "EnumClient.PutNotExpandable":
				res.resp, res.err = e.dispatchPutNotExpandable(req)
			case "EnumClient.PutReferenced":
				res.resp, res.err = e.dispatchPutReferenced(req)
			case "EnumClient.PutReferencedConstant":
				res.resp, res.err = e.dispatchPutReferencedConstant(req)
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

func (e *EnumServerTransport) dispatchGetNotExpandable(req *http.Request) (*http.Response, error) {
	if e.srv.GetNotExpandable == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNotExpandable not implemented")}
	}
	respr, errRespr := e.srv.GetNotExpandable(req.Context(), nil)
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
	return resp, nil
}

func (e *EnumServerTransport) dispatchGetReferenced(req *http.Request) (*http.Response, error) {
	if e.srv.GetReferenced == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetReferenced not implemented")}
	}
	respr, errRespr := e.srv.GetReferenced(req.Context(), nil)
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
	return resp, nil
}

func (e *EnumServerTransport) dispatchGetReferencedConstant(req *http.Request) (*http.Response, error) {
	if e.srv.GetReferencedConstant == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetReferencedConstant not implemented")}
	}
	respr, errRespr := e.srv.GetReferencedConstant(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RefColorConstant, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnumServerTransport) dispatchPutNotExpandable(req *http.Request) (*http.Response, error) {
	if e.srv.PutNotExpandable == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutNotExpandable not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[stringgroup.Colors](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PutNotExpandable(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnumServerTransport) dispatchPutReferenced(req *http.Request) (*http.Response, error) {
	if e.srv.PutReferenced == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutReferenced not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[stringgroup.Colors](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PutReferenced(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnumServerTransport) dispatchPutReferencedConstant(req *http.Request) (*http.Response, error) {
	if e.srv.PutReferencedConstant == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutReferencedConstant not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[stringgroup.RefColorConstant](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.PutReferencedConstant(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to EnumServerTransport
var enumServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
