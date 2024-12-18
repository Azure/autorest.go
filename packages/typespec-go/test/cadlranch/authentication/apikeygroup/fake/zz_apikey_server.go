// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"apikeygroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// APIKeyServer is a fake server for instances of the apikeygroup.APIKeyClient type.
type APIKeyServer struct {
	// Invalid is the fake for method APIKeyClient.Invalid
	// HTTP status codes to indicate success: http.StatusNoContent
	Invalid func(ctx context.Context, options *apikeygroup.APIKeyClientInvalidOptions) (resp azfake.Responder[apikeygroup.APIKeyClientInvalidResponse], errResp azfake.ErrorResponder)

	// Valid is the fake for method APIKeyClient.Valid
	// HTTP status codes to indicate success: http.StatusNoContent
	Valid func(ctx context.Context, options *apikeygroup.APIKeyClientValidOptions) (resp azfake.Responder[apikeygroup.APIKeyClientValidResponse], errResp azfake.ErrorResponder)
}

// NewAPIKeyServerTransport creates a new instance of APIKeyServerTransport with the provided implementation.
// The returned APIKeyServerTransport instance is connected to an instance of apikeygroup.APIKeyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAPIKeyServerTransport(srv *APIKeyServer) *APIKeyServerTransport {
	return &APIKeyServerTransport{srv: srv}
}

// APIKeyServerTransport connects instances of apikeygroup.APIKeyClient to instances of APIKeyServer.
// Don't use this type directly, use NewAPIKeyServerTransport instead.
type APIKeyServerTransport struct {
	srv *APIKeyServer
}

// Do implements the policy.Transporter interface for APIKeyServerTransport.
func (a *APIKeyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *APIKeyServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if apiKeyServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = apiKeyServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "APIKeyClient.Invalid":
				res.resp, res.err = a.dispatchInvalid(req)
			case "APIKeyClient.Valid":
				res.resp, res.err = a.dispatchValid(req)
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

func (a *APIKeyServerTransport) dispatchInvalid(req *http.Request) (*http.Response, error) {
	if a.srv.Invalid == nil {
		return nil, &nonRetriableError{errors.New("fake for method Invalid not implemented")}
	}
	respr, errRespr := a.srv.Invalid(req.Context(), nil)
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

func (a *APIKeyServerTransport) dispatchValid(req *http.Request) (*http.Response, error) {
	if a.srv.Valid == nil {
		return nil, &nonRetriableError{errors.New("fake for method Valid not implemented")}
	}
	respr, errRespr := a.srv.Valid(req.Context(), nil)
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

// set this to conditionally intercept incoming requests to APIKeyServerTransport
var apiKeyServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
