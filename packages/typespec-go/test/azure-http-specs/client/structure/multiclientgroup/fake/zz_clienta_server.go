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
	"multiclientgroup"
	"net/http"
)

// ClientAServer is a fake server for instances of the multiclientgroup.ClientAClient type.
type ClientAServer struct {
	// RenamedFive is the fake for method ClientAClient.RenamedFive
	// HTTP status codes to indicate success: http.StatusNoContent
	RenamedFive func(ctx context.Context, options *multiclientgroup.ClientAClientRenamedFiveOptions) (resp azfake.Responder[multiclientgroup.ClientAClientRenamedFiveResponse], errResp azfake.ErrorResponder)

	// RenamedOne is the fake for method ClientAClient.RenamedOne
	// HTTP status codes to indicate success: http.StatusNoContent
	RenamedOne func(ctx context.Context, options *multiclientgroup.ClientAClientRenamedOneOptions) (resp azfake.Responder[multiclientgroup.ClientAClientRenamedOneResponse], errResp azfake.ErrorResponder)

	// RenamedThree is the fake for method ClientAClient.RenamedThree
	// HTTP status codes to indicate success: http.StatusNoContent
	RenamedThree func(ctx context.Context, options *multiclientgroup.ClientAClientRenamedThreeOptions) (resp azfake.Responder[multiclientgroup.ClientAClientRenamedThreeResponse], errResp azfake.ErrorResponder)
}

// NewClientAServerTransport creates a new instance of ClientAServerTransport with the provided implementation.
// The returned ClientAServerTransport instance is connected to an instance of multiclientgroup.ClientAClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewClientAServerTransport(srv *ClientAServer) *ClientAServerTransport {
	return &ClientAServerTransport{srv: srv}
}

// ClientAServerTransport connects instances of multiclientgroup.ClientAClient to instances of ClientAServer.
// Don't use this type directly, use NewClientAServerTransport instead.
type ClientAServerTransport struct {
	srv *ClientAServer
}

// Do implements the policy.Transporter interface for ClientAServerTransport.
func (c *ClientAServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ClientAServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if clientAServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = clientAServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ClientAClient.RenamedFive":
				res.resp, res.err = c.dispatchRenamedFive(req)
			case "ClientAClient.RenamedOne":
				res.resp, res.err = c.dispatchRenamedOne(req)
			case "ClientAClient.RenamedThree":
				res.resp, res.err = c.dispatchRenamedThree(req)
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

func (c *ClientAServerTransport) dispatchRenamedFive(req *http.Request) (*http.Response, error) {
	if c.srv.RenamedFive == nil {
		return nil, &nonRetriableError{errors.New("fake for method RenamedFive not implemented")}
	}
	respr, errRespr := c.srv.RenamedFive(req.Context(), nil)
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

func (c *ClientAServerTransport) dispatchRenamedOne(req *http.Request) (*http.Response, error) {
	if c.srv.RenamedOne == nil {
		return nil, &nonRetriableError{errors.New("fake for method RenamedOne not implemented")}
	}
	respr, errRespr := c.srv.RenamedOne(req.Context(), nil)
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

func (c *ClientAServerTransport) dispatchRenamedThree(req *http.Request) (*http.Response, error) {
	if c.srv.RenamedThree == nil {
		return nil, &nonRetriableError{errors.New("fake for method RenamedThree not implemented")}
	}
	respr, errRespr := c.srv.RenamedThree(req.Context(), nil)
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

// set this to conditionally intercept incoming requests to ClientAServerTransport
var clientAServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}