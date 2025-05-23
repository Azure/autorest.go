// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"armapicenter"
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
)

// DeletedServicesServer is a fake server for instances of the armapicenter.DeletedServicesClient type.
type DeletedServicesServer struct {
	// Delete is the fake for method DeletedServicesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, deletedServiceName string, options *armapicenter.DeletedServicesClientDeleteOptions) (resp azfake.Responder[armapicenter.DeletedServicesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DeletedServicesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, deletedServiceName string, options *armapicenter.DeletedServicesClientGetOptions) (resp azfake.Responder[armapicenter.DeletedServicesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DeletedServicesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armapicenter.DeletedServicesClientListOptions) (resp azfake.PagerResponder[armapicenter.DeletedServicesClientListResponse])

	// NewListBySubscriptionPager is the fake for method DeletedServicesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armapicenter.DeletedServicesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armapicenter.DeletedServicesClientListBySubscriptionResponse])
}

// NewDeletedServicesServerTransport creates a new instance of DeletedServicesServerTransport with the provided implementation.
// The returned DeletedServicesServerTransport instance is connected to an instance of armapicenter.DeletedServicesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDeletedServicesServerTransport(srv *DeletedServicesServer) *DeletedServicesServerTransport {
	return &DeletedServicesServerTransport{
		srv:                        srv,
		newListPager:               newTracker[azfake.PagerResponder[armapicenter.DeletedServicesClientListResponse]](),
		newListBySubscriptionPager: newTracker[azfake.PagerResponder[armapicenter.DeletedServicesClientListBySubscriptionResponse]](),
	}
}

// DeletedServicesServerTransport connects instances of armapicenter.DeletedServicesClient to instances of DeletedServicesServer.
// Don't use this type directly, use NewDeletedServicesServerTransport instead.
type DeletedServicesServerTransport struct {
	srv                        *DeletedServicesServer
	newListPager               *tracker[azfake.PagerResponder[armapicenter.DeletedServicesClientListResponse]]
	newListBySubscriptionPager *tracker[azfake.PagerResponder[armapicenter.DeletedServicesClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for DeletedServicesServerTransport.
func (d *DeletedServicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DeletedServicesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if deletedServicesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = deletedServicesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DeletedServicesClient.Delete":
				res.resp, res.err = d.dispatchDelete(req)
			case "DeletedServicesClient.Get":
				res.resp, res.err = d.dispatchGet(req)
			case "DeletedServicesClient.NewListPager":
				res.resp, res.err = d.dispatchNewListPager(req)
			case "DeletedServicesClient.NewListBySubscriptionPager":
				res.resp, res.err = d.dispatchNewListBySubscriptionPager(req)
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

func (d *DeletedServicesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if d.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiCenter/deletedServices/(?P<deletedServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	deletedServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deletedServiceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Delete(req.Context(), resourceGroupNameParam, deletedServiceNameParam, nil)
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

func (d *DeletedServicesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiCenter/deletedServices/(?P<deletedServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	deletedServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deletedServiceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, deletedServiceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeletedService, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (d *DeletedServicesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiCenter/deletedServices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armapicenter.DeletedServicesClientListOptions
		if filterParam != nil {
			options = &armapicenter.DeletedServicesClientListOptions{
				Filter: filterParam,
			}
		}
		resp := d.srv.NewListPager(resourceGroupNameParam, options)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armapicenter.DeletedServicesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		d.newListPager.remove(req)
	}
	return resp, nil
}

func (d *DeletedServicesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := d.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiCenter/deletedServices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		d.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armapicenter.DeletedServicesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		d.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to DeletedServicesServerTransport
var deletedServicesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
