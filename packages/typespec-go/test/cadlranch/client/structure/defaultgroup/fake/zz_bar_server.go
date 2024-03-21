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

// BarServer is a fake server for instances of the defaultgroup.BarClient type.
type BarServer struct {
	// Five is the fake for method BarClient.Five
	// HTTP status codes to indicate success: http.StatusNoContent
	Five func(ctx context.Context, options *defaultgroup.BarClientFiveOptions) (resp azfake.Responder[defaultgroup.BarClientFiveResponse], errResp azfake.ErrorResponder)

	// Nine is the fake for method BarClient.Nine
	// HTTP status codes to indicate success: http.StatusNoContent
	Nine func(ctx context.Context, options *defaultgroup.BarClientNineOptions) (resp azfake.Responder[defaultgroup.BarClientNineResponse], errResp azfake.ErrorResponder)

	// Six is the fake for method BarClient.Six
	// HTTP status codes to indicate success: http.StatusNoContent
	Six func(ctx context.Context, options *defaultgroup.BarClientSixOptions) (resp azfake.Responder[defaultgroup.BarClientSixResponse], errResp azfake.ErrorResponder)
}

// NewBarServerTransport creates a new instance of BarServerTransport with the provided implementation.
// The returned BarServerTransport instance is connected to an instance of defaultgroup.BarClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBarServerTransport(srv *BarServer) *BarServerTransport {
	return &BarServerTransport{srv: srv}
}

// BarServerTransport connects instances of defaultgroup.BarClient to instances of BarServer.
// Don't use this type directly, use NewBarServerTransport instead.
type BarServerTransport struct {
	srv *BarServer
}

// Do implements the policy.Transporter interface for BarServerTransport.
func (b *BarServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BarClient.Five":
		resp, err = b.dispatchFive(req)
	case "BarClient.Nine":
		resp, err = b.dispatchNine(req)
	case "BarClient.Six":
		resp, err = b.dispatchSix(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BarServerTransport) dispatchFive(req *http.Request) (*http.Response, error) {
	if b.srv.Five == nil {
		return nil, &nonRetriableError{errors.New("fake for method Five not implemented")}
	}
	respr, errRespr := b.srv.Five(req.Context(), nil)
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

func (b *BarServerTransport) dispatchNine(req *http.Request) (*http.Response, error) {
	if b.srv.Nine == nil {
		return nil, &nonRetriableError{errors.New("fake for method Nine not implemented")}
	}
	respr, errRespr := b.srv.Nine(req.Context(), nil)
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

func (b *BarServerTransport) dispatchSix(req *http.Request) (*http.Response, error) {
	if b.srv.Six == nil {
		return nil, &nonRetriableError{errors.New("fake for method Six not implemented")}
	}
	respr, errRespr := b.srv.Six(req.Context(), nil)
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