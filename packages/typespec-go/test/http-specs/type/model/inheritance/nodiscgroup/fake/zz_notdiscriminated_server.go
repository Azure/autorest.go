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
	"nodiscgroup"
)

// NotDiscriminatedServer is a fake server for instances of the nodiscgroup.NotDiscriminatedClient type.
type NotDiscriminatedServer struct {
	// GetValid is the fake for method NotDiscriminatedClient.GetValid
	// HTTP status codes to indicate success: http.StatusOK
	GetValid func(ctx context.Context, options *nodiscgroup.NotDiscriminatedClientGetValidOptions) (resp azfake.Responder[nodiscgroup.NotDiscriminatedClientGetValidResponse], errResp azfake.ErrorResponder)

	// PostValid is the fake for method NotDiscriminatedClient.PostValid
	// HTTP status codes to indicate success: http.StatusNoContent
	PostValid func(ctx context.Context, input nodiscgroup.Siamese, options *nodiscgroup.NotDiscriminatedClientPostValidOptions) (resp azfake.Responder[nodiscgroup.NotDiscriminatedClientPostValidResponse], errResp azfake.ErrorResponder)

	// PutValid is the fake for method NotDiscriminatedClient.PutValid
	// HTTP status codes to indicate success: http.StatusOK
	PutValid func(ctx context.Context, input nodiscgroup.Siamese, options *nodiscgroup.NotDiscriminatedClientPutValidOptions) (resp azfake.Responder[nodiscgroup.NotDiscriminatedClientPutValidResponse], errResp azfake.ErrorResponder)
}

// NewNotDiscriminatedServerTransport creates a new instance of NotDiscriminatedServerTransport with the provided implementation.
// The returned NotDiscriminatedServerTransport instance is connected to an instance of nodiscgroup.NotDiscriminatedClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNotDiscriminatedServerTransport(srv *NotDiscriminatedServer) *NotDiscriminatedServerTransport {
	return &NotDiscriminatedServerTransport{srv: srv}
}

// NotDiscriminatedServerTransport connects instances of nodiscgroup.NotDiscriminatedClient to instances of NotDiscriminatedServer.
// Don't use this type directly, use NewNotDiscriminatedServerTransport instead.
type NotDiscriminatedServerTransport struct {
	srv *NotDiscriminatedServer
}

// Do implements the policy.Transporter interface for NotDiscriminatedServerTransport.
func (n *NotDiscriminatedServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return n.dispatchToMethodFake(req, method)
}

func (n *NotDiscriminatedServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if notDiscriminatedServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = notDiscriminatedServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "NotDiscriminatedClient.GetValid":
				res.resp, res.err = n.dispatchGetValid(req)
			case "NotDiscriminatedClient.PostValid":
				res.resp, res.err = n.dispatchPostValid(req)
			case "NotDiscriminatedClient.PutValid":
				res.resp, res.err = n.dispatchPutValid(req)
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

func (n *NotDiscriminatedServerTransport) dispatchGetValid(req *http.Request) (*http.Response, error) {
	if n.srv.GetValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetValid not implemented")}
	}
	respr, errRespr := n.srv.GetValid(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Siamese, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NotDiscriminatedServerTransport) dispatchPostValid(req *http.Request) (*http.Response, error) {
	if n.srv.PostValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostValid not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[nodiscgroup.Siamese](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.PostValid(req.Context(), body, nil)
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

func (n *NotDiscriminatedServerTransport) dispatchPutValid(req *http.Request) (*http.Response, error) {
	if n.srv.PutValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutValid not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[nodiscgroup.Siamese](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.PutValid(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Siamese, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to NotDiscriminatedServerTransport
var notDiscriminatedServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}