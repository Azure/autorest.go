// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"bytesgroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// BytesResponseBodyServer is a fake server for instances of the bytesgroup.BytesResponseBodyClient type.
type BytesResponseBodyServer struct {
	// Base64 is the fake for method BytesResponseBodyClient.Base64
	// HTTP status codes to indicate success: http.StatusOK
	Base64 func(ctx context.Context, options *bytesgroup.BytesResponseBodyClientBase64Options) (resp azfake.Responder[bytesgroup.BytesResponseBodyClientBase64Response], errResp azfake.ErrorResponder)

	// Base64URL is the fake for method BytesResponseBodyClient.Base64URL
	// HTTP status codes to indicate success: http.StatusOK
	Base64URL func(ctx context.Context, options *bytesgroup.BytesResponseBodyClientBase64URLOptions) (resp azfake.Responder[bytesgroup.BytesResponseBodyClientBase64URLResponse], errResp azfake.ErrorResponder)

	// CustomContentType is the fake for method BytesResponseBodyClient.CustomContentType
	// HTTP status codes to indicate success: http.StatusOK
	CustomContentType func(ctx context.Context, options *bytesgroup.BytesResponseBodyClientCustomContentTypeOptions) (resp azfake.Responder[bytesgroup.BytesResponseBodyClientCustomContentTypeResponse], errResp azfake.ErrorResponder)

	// Default is the fake for method BytesResponseBodyClient.Default
	// HTTP status codes to indicate success: http.StatusOK
	Default func(ctx context.Context, options *bytesgroup.BytesResponseBodyClientDefaultOptions) (resp azfake.Responder[bytesgroup.BytesResponseBodyClientDefaultResponse], errResp azfake.ErrorResponder)

	// OctetStream is the fake for method BytesResponseBodyClient.OctetStream
	// HTTP status codes to indicate success: http.StatusOK
	OctetStream func(ctx context.Context, options *bytesgroup.BytesResponseBodyClientOctetStreamOptions) (resp azfake.Responder[bytesgroup.BytesResponseBodyClientOctetStreamResponse], errResp azfake.ErrorResponder)
}

// NewBytesResponseBodyServerTransport creates a new instance of BytesResponseBodyServerTransport with the provided implementation.
// The returned BytesResponseBodyServerTransport instance is connected to an instance of bytesgroup.BytesResponseBodyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBytesResponseBodyServerTransport(srv *BytesResponseBodyServer) *BytesResponseBodyServerTransport {
	return &BytesResponseBodyServerTransport{srv: srv}
}

// BytesResponseBodyServerTransport connects instances of bytesgroup.BytesResponseBodyClient to instances of BytesResponseBodyServer.
// Don't use this type directly, use NewBytesResponseBodyServerTransport instead.
type BytesResponseBodyServerTransport struct {
	srv *BytesResponseBodyServer
}

// Do implements the policy.Transporter interface for BytesResponseBodyServerTransport.
func (b *BytesResponseBodyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BytesResponseBodyServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "BytesResponseBodyClient.Base64":
		resp, err = b.dispatchBase64(req)
	case "BytesResponseBodyClient.Base64URL":
		resp, err = b.dispatchBase64URL(req)
	case "BytesResponseBodyClient.CustomContentType":
		resp, err = b.dispatchCustomContentType(req)
	case "BytesResponseBodyClient.Default":
		resp, err = b.dispatchDefault(req)
	case "BytesResponseBodyClient.OctetStream":
		resp, err = b.dispatchOctetStream(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (b *BytesResponseBodyServerTransport) dispatchBase64(req *http.Request) (*http.Response, error) {
	if b.srv.Base64 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Base64 not implemented")}
	}
	respr, errRespr := b.srv.Base64(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsByteArray(respContent, server.GetResponse(respr).Value, runtime.Base64StdFormat, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BytesResponseBodyServerTransport) dispatchBase64URL(req *http.Request) (*http.Response, error) {
	if b.srv.Base64URL == nil {
		return nil, &nonRetriableError{errors.New("fake for method Base64URL not implemented")}
	}
	respr, errRespr := b.srv.Base64URL(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsByteArray(respContent, server.GetResponse(respr).Value, runtime.Base64URLFormat, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BytesResponseBodyServerTransport) dispatchCustomContentType(req *http.Request) (*http.Response, error) {
	if b.srv.CustomContentType == nil {
		return nil, &nonRetriableError{errors.New("fake for method CustomContentType not implemented")}
	}
	respr, errRespr := b.srv.CustomContentType(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: req.Header.Get("Content-Type"),
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BytesResponseBodyServerTransport) dispatchDefault(req *http.Request) (*http.Response, error) {
	if b.srv.Default == nil {
		return nil, &nonRetriableError{errors.New("fake for method Default not implemented")}
	}
	respr, errRespr := b.srv.Default(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsByteArray(respContent, server.GetResponse(respr).Value, runtime.Base64StdFormat, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BytesResponseBodyServerTransport) dispatchOctetStream(req *http.Request) (*http.Response, error) {
	if b.srv.OctetStream == nil {
		return nil, &nonRetriableError{errors.New("fake for method OctetStream not implemented")}
	}
	respr, errRespr := b.srv.OctetStream(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: req.Header.Get("Content-Type"),
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}