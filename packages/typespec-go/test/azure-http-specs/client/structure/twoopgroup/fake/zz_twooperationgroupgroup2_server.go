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
	"twoopgroup"
)

// TwoOperationGroupGroup2Server is a fake server for instances of the twoopgroup.TwoOperationGroupGroup2Client type.
type TwoOperationGroupGroup2Server struct {
	// Five is the fake for method TwoOperationGroupGroup2Client.Five
	// HTTP status codes to indicate success: http.StatusNoContent
	Five func(ctx context.Context, options *twoopgroup.TwoOperationGroupGroup2ClientFiveOptions) (resp azfake.Responder[twoopgroup.TwoOperationGroupGroup2ClientFiveResponse], errResp azfake.ErrorResponder)

	// Six is the fake for method TwoOperationGroupGroup2Client.Six
	// HTTP status codes to indicate success: http.StatusNoContent
	Six func(ctx context.Context, options *twoopgroup.TwoOperationGroupGroup2ClientSixOptions) (resp azfake.Responder[twoopgroup.TwoOperationGroupGroup2ClientSixResponse], errResp azfake.ErrorResponder)

	// Two is the fake for method TwoOperationGroupGroup2Client.Two
	// HTTP status codes to indicate success: http.StatusNoContent
	Two func(ctx context.Context, options *twoopgroup.TwoOperationGroupGroup2ClientTwoOptions) (resp azfake.Responder[twoopgroup.TwoOperationGroupGroup2ClientTwoResponse], errResp azfake.ErrorResponder)
}

// NewTwoOperationGroupGroup2ServerTransport creates a new instance of TwoOperationGroupGroup2ServerTransport with the provided implementation.
// The returned TwoOperationGroupGroup2ServerTransport instance is connected to an instance of twoopgroup.TwoOperationGroupGroup2Client via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTwoOperationGroupGroup2ServerTransport(srv *TwoOperationGroupGroup2Server) *TwoOperationGroupGroup2ServerTransport {
	return &TwoOperationGroupGroup2ServerTransport{srv: srv}
}

// TwoOperationGroupGroup2ServerTransport connects instances of twoopgroup.TwoOperationGroupGroup2Client to instances of TwoOperationGroupGroup2Server.
// Don't use this type directly, use NewTwoOperationGroupGroup2ServerTransport instead.
type TwoOperationGroupGroup2ServerTransport struct {
	srv *TwoOperationGroupGroup2Server
}

// Do implements the policy.Transporter interface for TwoOperationGroupGroup2ServerTransport.
func (t *TwoOperationGroupGroup2ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return t.dispatchToMethodFake(req, method)
}

func (t *TwoOperationGroupGroup2ServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if twoOperationGroupGroup2ServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = twoOperationGroupGroup2ServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "TwoOperationGroupGroup2Client.Five":
				res.resp, res.err = t.dispatchFive(req)
			case "TwoOperationGroupGroup2Client.Six":
				res.resp, res.err = t.dispatchSix(req)
			case "TwoOperationGroupGroup2Client.Two":
				res.resp, res.err = t.dispatchTwo(req)
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

func (t *TwoOperationGroupGroup2ServerTransport) dispatchFive(req *http.Request) (*http.Response, error) {
	if t.srv.Five == nil {
		return nil, &nonRetriableError{errors.New("fake for method Five not implemented")}
	}
	respr, errRespr := t.srv.Five(req.Context(), nil)
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

func (t *TwoOperationGroupGroup2ServerTransport) dispatchSix(req *http.Request) (*http.Response, error) {
	if t.srv.Six == nil {
		return nil, &nonRetriableError{errors.New("fake for method Six not implemented")}
	}
	respr, errRespr := t.srv.Six(req.Context(), nil)
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

func (t *TwoOperationGroupGroup2ServerTransport) dispatchTwo(req *http.Request) (*http.Response, error) {
	if t.srv.Two == nil {
		return nil, &nonRetriableError{errors.New("fake for method Two not implemented")}
	}
	respr, errRespr := t.srv.Two(req.Context(), nil)
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

// set this to conditionally intercept incoming requests to TwoOperationGroupGroup2ServerTransport
var twoOperationGroupGroup2ServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
