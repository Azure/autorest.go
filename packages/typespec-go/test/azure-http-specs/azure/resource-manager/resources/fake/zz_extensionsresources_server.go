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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"net/url"
	"regexp"
	"resources"
)

// ExtensionsResourcesServer is a fake server for instances of the resources.ExtensionsResourcesClient type.
type ExtensionsResourcesServer struct {
	// BeginCreateOrUpdate is the fake for method ExtensionsResourcesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceURI string, extensionsResourceName string, resource resources.ExtensionsResource, options *resources.ExtensionsResourcesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[resources.ExtensionsResourcesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ExtensionsResourcesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceURI string, extensionsResourceName string, options *resources.ExtensionsResourcesClientDeleteOptions) (resp azfake.Responder[resources.ExtensionsResourcesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ExtensionsResourcesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceURI string, extensionsResourceName string, options *resources.ExtensionsResourcesClientGetOptions) (resp azfake.Responder[resources.ExtensionsResourcesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByScopePager is the fake for method ExtensionsResourcesClient.NewListByScopePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByScopePager func(resourceURI string, options *resources.ExtensionsResourcesClientListByScopeOptions) (resp azfake.PagerResponder[resources.ExtensionsResourcesClientListByScopeResponse])

	// Update is the fake for method ExtensionsResourcesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceURI string, extensionsResourceName string, properties resources.ExtensionsResource, options *resources.ExtensionsResourcesClientUpdateOptions) (resp azfake.Responder[resources.ExtensionsResourcesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewExtensionsResourcesServerTransport creates a new instance of ExtensionsResourcesServerTransport with the provided implementation.
// The returned ExtensionsResourcesServerTransport instance is connected to an instance of resources.ExtensionsResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExtensionsResourcesServerTransport(srv *ExtensionsResourcesServer) *ExtensionsResourcesServerTransport {
	return &ExtensionsResourcesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[resources.ExtensionsResourcesClientCreateOrUpdateResponse]](),
		newListByScopePager: newTracker[azfake.PagerResponder[resources.ExtensionsResourcesClientListByScopeResponse]](),
	}
}

// ExtensionsResourcesServerTransport connects instances of resources.ExtensionsResourcesClient to instances of ExtensionsResourcesServer.
// Don't use this type directly, use NewExtensionsResourcesServerTransport instead.
type ExtensionsResourcesServerTransport struct {
	srv                 *ExtensionsResourcesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[resources.ExtensionsResourcesClientCreateOrUpdateResponse]]
	newListByScopePager *tracker[azfake.PagerResponder[resources.ExtensionsResourcesClientListByScopeResponse]]
}

// Do implements the policy.Transporter interface for ExtensionsResourcesServerTransport.
func (e *ExtensionsResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *ExtensionsResourcesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if extensionsResourcesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = extensionsResourcesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ExtensionsResourcesClient.BeginCreateOrUpdate":
				res.resp, res.err = e.dispatchBeginCreateOrUpdate(req)
			case "ExtensionsResourcesClient.Delete":
				res.resp, res.err = e.dispatchDelete(req)
			case "ExtensionsResourcesClient.Get":
				res.resp, res.err = e.dispatchGet(req)
			case "ExtensionsResourcesClient.NewListByScopePager":
				res.resp, res.err = e.dispatchNewListByScopePager(req)
			case "ExtensionsResourcesClient.Update":
				res.resp, res.err = e.dispatchUpdate(req)
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

func (e *ExtensionsResourcesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := e.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Azure\.ResourceManager\.Resources/extensionsResources/(?P<extensionsResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[resources.ExtensionsResource](req)
		if err != nil {
			return nil, err
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		extensionsResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("extensionsResourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceURIParam, extensionsResourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		e.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		e.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		e.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (e *ExtensionsResourcesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if e.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Azure\.ResourceManager\.Resources/extensionsResources/(?P<extensionsResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	extensionsResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("extensionsResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Delete(req.Context(), resourceURIParam, extensionsResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExtensionsResourcesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Azure\.ResourceManager\.Resources/extensionsResources/(?P<extensionsResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	extensionsResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("extensionsResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceURIParam, extensionsResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExtensionsResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExtensionsResourcesServerTransport) dispatchNewListByScopePager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByScopePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByScopePager not implemented")}
	}
	newListByScopePager := e.newListByScopePager.get(req)
	if newListByScopePager == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Azure\.ResourceManager\.Resources/extensionsResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListByScopePager(resourceURIParam, nil)
		newListByScopePager = &resp
		e.newListByScopePager.add(req, newListByScopePager)
		server.PagerResponderInjectNextLinks(newListByScopePager, req, func(page *resources.ExtensionsResourcesClientListByScopeResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByScopePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByScopePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByScopePager) {
		e.newListByScopePager.remove(req)
	}
	return resp, nil
}

func (e *ExtensionsResourcesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Azure\.ResourceManager\.Resources/extensionsResources/(?P<extensionsResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[resources.ExtensionsResource](req)
	if err != nil {
		return nil, err
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	extensionsResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("extensionsResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Update(req.Context(), resourceURIParam, extensionsResourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExtensionsResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ExtensionsResourcesServerTransport
var extensionsResourcesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
