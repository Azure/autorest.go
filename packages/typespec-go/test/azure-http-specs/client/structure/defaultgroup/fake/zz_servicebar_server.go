// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"defaultgroup"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ServiceBarServer is a fake server for instances of the defaultgroup.ServiceBarClient type.
type ServiceBarServer struct {
	// Five is the fake for method ServiceBarClient.Five
	// HTTP status codes to indicate success: http.StatusNoContent
	Five func(ctx context.Context, options *defaultgroup.ServiceBarClientFiveOptions) (resp azfake.Responder[defaultgroup.ServiceBarClientFiveResponse], errResp azfake.ErrorResponder)

	// Six is the fake for method ServiceBarClient.Six
	// HTTP status codes to indicate success: http.StatusNoContent
	Six func(ctx context.Context, options *defaultgroup.ServiceBarClientSixOptions) (resp azfake.Responder[defaultgroup.ServiceBarClientSixResponse], errResp azfake.ErrorResponder)
}

// NewServiceBarServerTransport creates a new instance of ServiceBarServerTransport with the provided implementation.
// The returned ServiceBarServerTransport instance is connected to an instance of defaultgroup.ServiceBarClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServiceBarServerTransport(srv *ServiceBarServer) *ServiceBarServerTransport {
	return &ServiceBarServerTransport{srv: srv}
}

// ServiceBarServerTransport connects instances of defaultgroup.ServiceBarClient to instances of ServiceBarServer.
// Don't use this type directly, use NewServiceBarServerTransport instead.
type ServiceBarServerTransport struct {
	srv *ServiceBarServer
}

// Do implements the policy.Transporter interface for ServiceBarServerTransport.
func (s *ServiceBarServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServiceBarServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if serviceBarServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = serviceBarServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ServiceBarClient.Five":
				res.resp, res.err = s.dispatchFive(req)
			case "ServiceBarClient.Six":
				res.resp, res.err = s.dispatchSix(req)
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

func (s *ServiceBarServerTransport) dispatchFive(req *http.Request) (*http.Response, error) {
	if s.srv.Five == nil {
		return nil, &nonRetriableError{errors.New("fake for method Five not implemented")}
	}
	respr, errRespr := s.srv.Five(req.Context(), nil)
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

func (s *ServiceBarServerTransport) dispatchSix(req *http.Request) (*http.Response, error) {
	if s.srv.Six == nil {
		return nil, &nonRetriableError{errors.New("fake for method Six not implemented")}
	}
	respr, errRespr := s.srv.Six(req.Context(), nil)
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

// set this to conditionally intercept incoming requests to ServiceBarServerTransport
var serviceBarServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}