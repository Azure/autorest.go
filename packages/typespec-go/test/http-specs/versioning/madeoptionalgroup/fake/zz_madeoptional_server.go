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
	"madeoptionalgroup"
	"net/http"
	"net/url"
)

// MadeOptionalServer is a fake server for instances of the madeoptionalgroup.MadeOptionalClient type.
type MadeOptionalServer struct {
	// Test is the fake for method MadeOptionalClient.Test
	// HTTP status codes to indicate success: http.StatusOK
	Test func(ctx context.Context, body madeoptionalgroup.TestModel, options *madeoptionalgroup.MadeOptionalClientTestOptions) (resp azfake.Responder[madeoptionalgroup.MadeOptionalClientTestResponse], errResp azfake.ErrorResponder)
}

// NewMadeOptionalServerTransport creates a new instance of MadeOptionalServerTransport with the provided implementation.
// The returned MadeOptionalServerTransport instance is connected to an instance of madeoptionalgroup.MadeOptionalClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMadeOptionalServerTransport(srv *MadeOptionalServer) *MadeOptionalServerTransport {
	return &MadeOptionalServerTransport{srv: srv}
}

// MadeOptionalServerTransport connects instances of madeoptionalgroup.MadeOptionalClient to instances of MadeOptionalServer.
// Don't use this type directly, use NewMadeOptionalServerTransport instead.
type MadeOptionalServerTransport struct {
	srv *MadeOptionalServer
}

// Do implements the policy.Transporter interface for MadeOptionalServerTransport.
func (m *MadeOptionalServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *MadeOptionalServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if madeOptionalServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = madeOptionalServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "MadeOptionalClient.Test":
				res.resp, res.err = m.dispatchTest(req)
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

func (m *MadeOptionalServerTransport) dispatchTest(req *http.Request) (*http.Response, error) {
	if m.srv.Test == nil {
		return nil, &nonRetriableError{errors.New("fake for method Test not implemented")}
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[madeoptionalgroup.TestModel](req)
	if err != nil {
		return nil, err
	}
	paramUnescaped, err := url.QueryUnescape(qp.Get("param"))
	if err != nil {
		return nil, err
	}
	paramParam := getOptional(paramUnescaped)
	var options *madeoptionalgroup.MadeOptionalClientTestOptions
	if paramParam != nil {
		options = &madeoptionalgroup.MadeOptionalClientTestOptions{
			Param: paramParam,
		}
	}
	respr, errRespr := m.srv.Test(req.Context(), body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TestModel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to MadeOptionalServerTransport
var madeOptionalServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
