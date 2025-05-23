// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"condreqgroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"time"
)

// ConditionalRequestServer is a fake server for instances of the condreqgroup.ConditionalRequestClient type.
type ConditionalRequestServer struct {
	// HeadIfModifiedSince is the fake for method ConditionalRequestClient.HeadIfModifiedSince
	// HTTP status codes to indicate success: http.StatusNoContent
	HeadIfModifiedSince func(ctx context.Context, options *condreqgroup.ConditionalRequestClientHeadIfModifiedSinceOptions) (resp azfake.Responder[condreqgroup.ConditionalRequestClientHeadIfModifiedSinceResponse], errResp azfake.ErrorResponder)

	// PostIfMatch is the fake for method ConditionalRequestClient.PostIfMatch
	// HTTP status codes to indicate success: http.StatusNoContent
	PostIfMatch func(ctx context.Context, options *condreqgroup.ConditionalRequestClientPostIfMatchOptions) (resp azfake.Responder[condreqgroup.ConditionalRequestClientPostIfMatchResponse], errResp azfake.ErrorResponder)

	// PostIfNoneMatch is the fake for method ConditionalRequestClient.PostIfNoneMatch
	// HTTP status codes to indicate success: http.StatusNoContent
	PostIfNoneMatch func(ctx context.Context, options *condreqgroup.ConditionalRequestClientPostIfNoneMatchOptions) (resp azfake.Responder[condreqgroup.ConditionalRequestClientPostIfNoneMatchResponse], errResp azfake.ErrorResponder)

	// PostIfUnmodifiedSince is the fake for method ConditionalRequestClient.PostIfUnmodifiedSince
	// HTTP status codes to indicate success: http.StatusNoContent
	PostIfUnmodifiedSince func(ctx context.Context, options *condreqgroup.ConditionalRequestClientPostIfUnmodifiedSinceOptions) (resp azfake.Responder[condreqgroup.ConditionalRequestClientPostIfUnmodifiedSinceResponse], errResp azfake.ErrorResponder)
}

// NewConditionalRequestServerTransport creates a new instance of ConditionalRequestServerTransport with the provided implementation.
// The returned ConditionalRequestServerTransport instance is connected to an instance of condreqgroup.ConditionalRequestClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConditionalRequestServerTransport(srv *ConditionalRequestServer) *ConditionalRequestServerTransport {
	return &ConditionalRequestServerTransport{srv: srv}
}

// ConditionalRequestServerTransport connects instances of condreqgroup.ConditionalRequestClient to instances of ConditionalRequestServer.
// Don't use this type directly, use NewConditionalRequestServerTransport instead.
type ConditionalRequestServerTransport struct {
	srv *ConditionalRequestServer
}

// Do implements the policy.Transporter interface for ConditionalRequestServerTransport.
func (c *ConditionalRequestServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ConditionalRequestServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if conditionalRequestServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = conditionalRequestServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ConditionalRequestClient.HeadIfModifiedSince":
				res.resp, res.err = c.dispatchHeadIfModifiedSince(req)
			case "ConditionalRequestClient.PostIfMatch":
				res.resp, res.err = c.dispatchPostIfMatch(req)
			case "ConditionalRequestClient.PostIfNoneMatch":
				res.resp, res.err = c.dispatchPostIfNoneMatch(req)
			case "ConditionalRequestClient.PostIfUnmodifiedSince":
				res.resp, res.err = c.dispatchPostIfUnmodifiedSince(req)
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

func (c *ConditionalRequestServerTransport) dispatchHeadIfModifiedSince(req *http.Request) (*http.Response, error) {
	if c.srv.HeadIfModifiedSince == nil {
		return nil, &nonRetriableError{errors.New("fake for method HeadIfModifiedSince not implemented")}
	}
	ifModifiedSinceParam, err := parseOptional(getHeaderValue(req.Header, "If-Modified-Since"), func(v string) (time.Time, error) { return time.Parse(time.RFC1123, v) })
	if err != nil {
		return nil, err
	}
	var options *condreqgroup.ConditionalRequestClientHeadIfModifiedSinceOptions
	if ifModifiedSinceParam != nil {
		options = &condreqgroup.ConditionalRequestClientHeadIfModifiedSinceOptions{
			IfModifiedSince: ifModifiedSinceParam,
		}
	}
	respr, errRespr := c.srv.HeadIfModifiedSince(req.Context(), options)
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

func (c *ConditionalRequestServerTransport) dispatchPostIfMatch(req *http.Request) (*http.Response, error) {
	if c.srv.PostIfMatch == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostIfMatch not implemented")}
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *condreqgroup.ConditionalRequestClientPostIfMatchOptions
	if ifMatchParam != nil {
		options = &condreqgroup.ConditionalRequestClientPostIfMatchOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := c.srv.PostIfMatch(req.Context(), options)
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

func (c *ConditionalRequestServerTransport) dispatchPostIfNoneMatch(req *http.Request) (*http.Response, error) {
	if c.srv.PostIfNoneMatch == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostIfNoneMatch not implemented")}
	}
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	var options *condreqgroup.ConditionalRequestClientPostIfNoneMatchOptions
	if ifNoneMatchParam != nil {
		options = &condreqgroup.ConditionalRequestClientPostIfNoneMatchOptions{
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := c.srv.PostIfNoneMatch(req.Context(), options)
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

func (c *ConditionalRequestServerTransport) dispatchPostIfUnmodifiedSince(req *http.Request) (*http.Response, error) {
	if c.srv.PostIfUnmodifiedSince == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostIfUnmodifiedSince not implemented")}
	}
	ifUnmodifiedSinceParam, err := parseOptional(getHeaderValue(req.Header, "If-Unmodified-Since"), func(v string) (time.Time, error) { return time.Parse(time.RFC1123, v) })
	if err != nil {
		return nil, err
	}
	var options *condreqgroup.ConditionalRequestClientPostIfUnmodifiedSinceOptions
	if ifUnmodifiedSinceParam != nil {
		options = &condreqgroup.ConditionalRequestClientPostIfUnmodifiedSinceOptions{
			IfUnmodifiedSince: ifUnmodifiedSinceParam,
		}
	}
	respr, errRespr := c.srv.PostIfUnmodifiedSince(req.Context(), options)
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

// set this to conditionally intercept incoming requests to ConditionalRequestServerTransport
var conditionalRequestServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
