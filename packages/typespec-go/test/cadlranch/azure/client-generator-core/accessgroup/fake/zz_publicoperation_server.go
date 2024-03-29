// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"accessgroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
)

// PublicOperationServer is a fake server for instances of the accessgroup.PublicOperationClient type.
type PublicOperationServer struct {
	// NoDecoratorInPublic is the fake for method PublicOperationClient.NoDecoratorInPublic
	// HTTP status codes to indicate success: http.StatusOK
	NoDecoratorInPublic func(ctx context.Context, name string, options *accessgroup.PublicOperationClientNoDecoratorInPublicOptions) (resp azfake.Responder[accessgroup.PublicOperationClientNoDecoratorInPublicResponse], errResp azfake.ErrorResponder)

	// PublicDecoratorInPublic is the fake for method PublicOperationClient.PublicDecoratorInPublic
	// HTTP status codes to indicate success: http.StatusOK
	PublicDecoratorInPublic func(ctx context.Context, name string, options *accessgroup.PublicOperationClientPublicDecoratorInPublicOptions) (resp azfake.Responder[accessgroup.PublicOperationClientPublicDecoratorInPublicResponse], errResp azfake.ErrorResponder)
}

// NewPublicOperationServerTransport creates a new instance of PublicOperationServerTransport with the provided implementation.
// The returned PublicOperationServerTransport instance is connected to an instance of accessgroup.PublicOperationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPublicOperationServerTransport(srv *PublicOperationServer) *PublicOperationServerTransport {
	return &PublicOperationServerTransport{srv: srv}
}

// PublicOperationServerTransport connects instances of accessgroup.PublicOperationClient to instances of PublicOperationServer.
// Don't use this type directly, use NewPublicOperationServerTransport instead.
type PublicOperationServerTransport struct {
	srv *PublicOperationServer
}

// Do implements the policy.Transporter interface for PublicOperationServerTransport.
func (p *PublicOperationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PublicOperationServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "PublicOperationClient.NoDecoratorInPublic":
		resp, err = p.dispatchNoDecoratorInPublic(req)
	case "PublicOperationClient.PublicDecoratorInPublic":
		resp, err = p.dispatchPublicDecoratorInPublic(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (p *PublicOperationServerTransport) dispatchNoDecoratorInPublic(req *http.Request) (*http.Response, error) {
	if p.srv.NoDecoratorInPublic == nil {
		return nil, &nonRetriableError{errors.New("fake for method NoDecoratorInPublic not implemented")}
	}
	qp := req.URL.Query()
	nameParam, err := url.QueryUnescape(qp.Get("name"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.NoDecoratorInPublic(req.Context(), nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NoDecoratorModelInPublic, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PublicOperationServerTransport) dispatchPublicDecoratorInPublic(req *http.Request) (*http.Response, error) {
	if p.srv.PublicDecoratorInPublic == nil {
		return nil, &nonRetriableError{errors.New("fake for method PublicDecoratorInPublic not implemented")}
	}
	qp := req.URL.Query()
	nameParam, err := url.QueryUnescape(qp.Get("name"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.PublicDecoratorInPublic(req.Context(), nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PublicDecoratorModelInPublic, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
