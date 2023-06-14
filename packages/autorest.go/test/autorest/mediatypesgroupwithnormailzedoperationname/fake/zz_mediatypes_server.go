//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/mediatypesgroupwithnormailzedoperationname"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
	"reflect"
)

// MediaTypesServer is a fake server for instances of the mediatypesgroupwithnormailzedoperationname.MediaTypesClient type.
type MediaTypesServer struct {
	// AnalyzeBody is the fake for method MediaTypesClient.AnalyzeBody
	// HTTP status codes to indicate success: http.StatusOK
	AnalyzeBody func(ctx context.Context, options *mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyOptions) (resp azfake.Responder[mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyResponse], errResp azfake.ErrorResponder)

	// AnalyzeBodyNoAcceptHeader is the fake for method MediaTypesClient.AnalyzeBodyNoAcceptHeader
	// HTTP status codes to indicate success: http.StatusAccepted
	AnalyzeBodyNoAcceptHeader func(ctx context.Context, options *mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions) (resp azfake.Responder[mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse], errResp azfake.ErrorResponder)

	// AnalyzeBodyNoAcceptHeaderWithBinary is the fake for method MediaTypesClient.AnalyzeBodyNoAcceptHeaderWithBinary
	// HTTP status codes to indicate success: http.StatusAccepted
	AnalyzeBodyNoAcceptHeaderWithBinary func(ctx context.Context, contentType mediatypesgroupwithnormailzedoperationname.ContentType, options *mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyNoAcceptHeaderWithBinaryOptions) (resp azfake.Responder[mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyNoAcceptHeaderWithBinaryResponse], errResp azfake.ErrorResponder)

	// AnalyzeBodyWithBinary is the fake for method MediaTypesClient.AnalyzeBodyWithBinary
	// HTTP status codes to indicate success: http.StatusOK
	AnalyzeBodyWithBinary func(ctx context.Context, contentType mediatypesgroupwithnormailzedoperationname.ContentType, options *mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyWithBinaryOptions) (resp azfake.Responder[mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyWithBinaryResponse], errResp azfake.ErrorResponder)

	// BinaryBodyWithThreeContentTypesWithBinary is the fake for method MediaTypesClient.BinaryBodyWithThreeContentTypesWithBinary
	// HTTP status codes to indicate success: http.StatusOK
	BinaryBodyWithThreeContentTypesWithBinary func(ctx context.Context, contentType mediatypesgroupwithnormailzedoperationname.ContentType1AutoGenerated, message io.ReadSeekCloser, options *mediatypesgroupwithnormailzedoperationname.MediaTypesClientBinaryBodyWithThreeContentTypesWithBinaryOptions) (resp azfake.Responder[mediatypesgroupwithnormailzedoperationname.MediaTypesClientBinaryBodyWithThreeContentTypesWithBinaryResponse], errResp azfake.ErrorResponder)

	// BinaryBodyWithTwoContentTypesWithBinary is the fake for method MediaTypesClient.BinaryBodyWithTwoContentTypesWithBinary
	// HTTP status codes to indicate success: http.StatusOK
	BinaryBodyWithTwoContentTypesWithBinary func(ctx context.Context, contentType mediatypesgroupwithnormailzedoperationname.ContentType1, message io.ReadSeekCloser, options *mediatypesgroupwithnormailzedoperationname.MediaTypesClientBinaryBodyWithTwoContentTypesWithBinaryOptions) (resp azfake.Responder[mediatypesgroupwithnormailzedoperationname.MediaTypesClientBinaryBodyWithTwoContentTypesWithBinaryResponse], errResp azfake.ErrorResponder)

	// ContentTypeWithEncodingWithText is the fake for method MediaTypesClient.ContentTypeWithEncodingWithText
	// HTTP status codes to indicate success: http.StatusOK
	ContentTypeWithEncodingWithText func(ctx context.Context, options *mediatypesgroupwithnormailzedoperationname.MediaTypesClientContentTypeWithEncodingWithTextOptions) (resp azfake.Responder[mediatypesgroupwithnormailzedoperationname.MediaTypesClientContentTypeWithEncodingWithTextResponse], errResp azfake.ErrorResponder)

	// PutTextAndJSONBodyWithText is the fake for method MediaTypesClient.PutTextAndJSONBodyWithText
	// HTTP status codes to indicate success: http.StatusOK
	PutTextAndJSONBodyWithText func(ctx context.Context, contentType mediatypesgroupwithnormailzedoperationname.ContentType1AutoGenerated2, message string, options *mediatypesgroupwithnormailzedoperationname.MediaTypesClientPutTextAndJSONBodyWithTextOptions) (resp azfake.Responder[mediatypesgroupwithnormailzedoperationname.MediaTypesClientPutTextAndJSONBodyWithTextResponse], errResp azfake.ErrorResponder)
}

// NewMediaTypesServerTransport creates a new instance of MediaTypesServerTransport with the provided implementation.
// The returned MediaTypesServerTransport instance is connected to an instance of mediatypesgroupwithnormailzedoperationname.MediaTypesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMediaTypesServerTransport(srv *MediaTypesServer) *MediaTypesServerTransport {
	return &MediaTypesServerTransport{srv: srv}
}

// MediaTypesServerTransport connects instances of mediatypesgroupwithnormailzedoperationname.MediaTypesClient to instances of MediaTypesServer.
// Don't use this type directly, use NewMediaTypesServerTransport instead.
type MediaTypesServerTransport struct {
	srv *MediaTypesServer
}

// Do implements the policy.Transporter interface for MediaTypesServerTransport.
func (m *MediaTypesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "MediaTypesClient.AnalyzeBody":
		resp, err = m.dispatchAnalyzeBody(req)
	case "MediaTypesClient.AnalyzeBodyNoAcceptHeader":
		resp, err = m.dispatchAnalyzeBodyNoAcceptHeader(req)
	case "MediaTypesClient.AnalyzeBodyNoAcceptHeaderWithBinary":
		resp, err = m.dispatchAnalyzeBodyNoAcceptHeaderWithBinary(req)
	case "MediaTypesClient.AnalyzeBodyWithBinary":
		resp, err = m.dispatchAnalyzeBodyWithBinary(req)
	case "MediaTypesClient.BinaryBodyWithThreeContentTypesWithBinary":
		resp, err = m.dispatchBinaryBodyWithThreeContentTypesWithBinary(req)
	case "MediaTypesClient.BinaryBodyWithTwoContentTypesWithBinary":
		resp, err = m.dispatchBinaryBodyWithTwoContentTypesWithBinary(req)
	case "MediaTypesClient.ContentTypeWithEncodingWithText":
		resp, err = m.dispatchContentTypeWithEncodingWithText(req)
	case "MediaTypesClient.PutTextAndJSONBodyWithText":
		resp, err = m.dispatchPutTextAndJSONBodyWithText(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *MediaTypesServerTransport) dispatchAnalyzeBody(req *http.Request) (*http.Response, error) {
	if m.srv.AnalyzeBody == nil {
		return nil, &nonRetriableError{errors.New("fake for method AnalyzeBody not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[mediatypesgroupwithnormailzedoperationname.SourcePath](req)
	if err != nil {
		return nil, err
	}
	var options *mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyOptions{
			Input: &body,
		}
	}
	respr, errRespr := m.srv.AnalyzeBody(req.Context(), options)
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

func (m *MediaTypesServerTransport) dispatchAnalyzeBodyNoAcceptHeader(req *http.Request) (*http.Response, error) {
	if m.srv.AnalyzeBodyNoAcceptHeader == nil {
		return nil, &nonRetriableError{errors.New("fake for method AnalyzeBodyNoAcceptHeader not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[mediatypesgroupwithnormailzedoperationname.SourcePath](req)
	if err != nil {
		return nil, err
	}
	var options *mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions{
			Input: &body,
		}
	}
	respr, errRespr := m.srv.AnalyzeBodyNoAcceptHeader(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MediaTypesServerTransport) dispatchAnalyzeBodyNoAcceptHeaderWithBinary(req *http.Request) (*http.Response, error) {
	if m.srv.AnalyzeBodyNoAcceptHeaderWithBinary == nil {
		return nil, &nonRetriableError{errors.New("fake for method AnalyzeBodyNoAcceptHeaderWithBinary not implemented")}
	}
	var options *mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyNoAcceptHeaderWithBinaryOptions
	if req.Body != nil {
		options = &mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyNoAcceptHeaderWithBinaryOptions{
			Input: req.Body.(io.ReadSeekCloser),
		}
	}
	respr, errRespr := m.srv.AnalyzeBodyNoAcceptHeaderWithBinary(req.Context(), mediatypesgroupwithnormailzedoperationname.ContentType(getHeaderValue(req.Header, "Content-Type")), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MediaTypesServerTransport) dispatchAnalyzeBodyWithBinary(req *http.Request) (*http.Response, error) {
	if m.srv.AnalyzeBodyWithBinary == nil {
		return nil, &nonRetriableError{errors.New("fake for method AnalyzeBodyWithBinary not implemented")}
	}
	var options *mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyWithBinaryOptions
	if req.Body != nil {
		options = &mediatypesgroupwithnormailzedoperationname.MediaTypesClientAnalyzeBodyWithBinaryOptions{
			Input: req.Body.(io.ReadSeekCloser),
		}
	}
	respr, errRespr := m.srv.AnalyzeBodyWithBinary(req.Context(), mediatypesgroupwithnormailzedoperationname.ContentType(getHeaderValue(req.Header, "Content-Type")), options)
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

func (m *MediaTypesServerTransport) dispatchBinaryBodyWithThreeContentTypesWithBinary(req *http.Request) (*http.Response, error) {
	if m.srv.BinaryBodyWithThreeContentTypesWithBinary == nil {
		return nil, &nonRetriableError{errors.New("fake for method BinaryBodyWithThreeContentTypesWithBinary not implemented")}
	}
	respr, errRespr := m.srv.BinaryBodyWithThreeContentTypesWithBinary(req.Context(), mediatypesgroupwithnormailzedoperationname.ContentType1AutoGenerated(getHeaderValue(req.Header, "Content-Type")), req.Body.(io.ReadSeekCloser), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsText(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MediaTypesServerTransport) dispatchBinaryBodyWithTwoContentTypesWithBinary(req *http.Request) (*http.Response, error) {
	if m.srv.BinaryBodyWithTwoContentTypesWithBinary == nil {
		return nil, &nonRetriableError{errors.New("fake for method BinaryBodyWithTwoContentTypesWithBinary not implemented")}
	}
	respr, errRespr := m.srv.BinaryBodyWithTwoContentTypesWithBinary(req.Context(), mediatypesgroupwithnormailzedoperationname.ContentType1(getHeaderValue(req.Header, "Content-Type")), req.Body.(io.ReadSeekCloser), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsText(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MediaTypesServerTransport) dispatchContentTypeWithEncodingWithText(req *http.Request) (*http.Response, error) {
	if m.srv.ContentTypeWithEncodingWithText == nil {
		return nil, &nonRetriableError{errors.New("fake for method ContentTypeWithEncodingWithText not implemented")}
	}
	body, err := server.UnmarshalRequestAsText(req)
	if err != nil {
		return nil, err
	}
	var options *mediatypesgroupwithnormailzedoperationname.MediaTypesClientContentTypeWithEncodingWithTextOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &mediatypesgroupwithnormailzedoperationname.MediaTypesClientContentTypeWithEncodingWithTextOptions{
			Input: &body,
		}
	}
	respr, errRespr := m.srv.ContentTypeWithEncodingWithText(req.Context(), options)
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

func (m *MediaTypesServerTransport) dispatchPutTextAndJSONBodyWithText(req *http.Request) (*http.Response, error) {
	if m.srv.PutTextAndJSONBodyWithText == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutTextAndJSONBodyWithText not implemented")}
	}
	body, err := server.UnmarshalRequestAsText(req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.PutTextAndJSONBodyWithText(req.Context(), mediatypesgroupwithnormailzedoperationname.ContentType1AutoGenerated2(getHeaderValue(req.Header, "Content-Type")), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsText(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}