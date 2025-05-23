// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"armloadtestservice"
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

// LoadTestsServer is a fake server for instances of the armloadtestservice.LoadTestsClient type.
type LoadTestsServer struct {
	// BeginCreateOrUpdate is the fake for method LoadTestsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, loadTestName string, resource armloadtestservice.LoadTestResource, options *armloadtestservice.LoadTestsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armloadtestservice.LoadTestsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method LoadTestsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, loadTestName string, options *armloadtestservice.LoadTestsClientBeginDeleteOptions) (resp azfake.PollerResponder[armloadtestservice.LoadTestsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method LoadTestsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, loadTestName string, options *armloadtestservice.LoadTestsClientGetOptions) (resp azfake.Responder[armloadtestservice.LoadTestsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method LoadTestsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armloadtestservice.LoadTestsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armloadtestservice.LoadTestsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method LoadTestsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armloadtestservice.LoadTestsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armloadtestservice.LoadTestsClientListBySubscriptionResponse])

	// NewOutboundNetworkDependenciesEndpointsPager is the fake for method LoadTestsClient.NewOutboundNetworkDependenciesEndpointsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewOutboundNetworkDependenciesEndpointsPager func(resourceGroupName string, loadTestName string, options *armloadtestservice.LoadTestsClientOutboundNetworkDependenciesEndpointsOptions) (resp azfake.PagerResponder[armloadtestservice.LoadTestsClientOutboundNetworkDependenciesEndpointsResponse])

	// BeginUpdate is the fake for method LoadTestsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginUpdate func(ctx context.Context, resourceGroupName string, loadTestName string, properties armloadtestservice.LoadTestResourceUpdate, options *armloadtestservice.LoadTestsClientBeginUpdateOptions) (resp azfake.PollerResponder[armloadtestservice.LoadTestsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewLoadTestsServerTransport creates a new instance of LoadTestsServerTransport with the provided implementation.
// The returned LoadTestsServerTransport instance is connected to an instance of armloadtestservice.LoadTestsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLoadTestsServerTransport(srv *LoadTestsServer) *LoadTestsServerTransport {
	return &LoadTestsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armloadtestservice.LoadTestsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armloadtestservice.LoadTestsClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armloadtestservice.LoadTestsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armloadtestservice.LoadTestsClientListBySubscriptionResponse]](),
		newOutboundNetworkDependenciesEndpointsPager: newTracker[azfake.PagerResponder[armloadtestservice.LoadTestsClientOutboundNetworkDependenciesEndpointsResponse]](),
		beginUpdate: newTracker[azfake.PollerResponder[armloadtestservice.LoadTestsClientUpdateResponse]](),
	}
}

// LoadTestsServerTransport connects instances of armloadtestservice.LoadTestsClient to instances of LoadTestsServer.
// Don't use this type directly, use NewLoadTestsServerTransport instead.
type LoadTestsServerTransport struct {
	srv                                          *LoadTestsServer
	beginCreateOrUpdate                          *tracker[azfake.PollerResponder[armloadtestservice.LoadTestsClientCreateOrUpdateResponse]]
	beginDelete                                  *tracker[azfake.PollerResponder[armloadtestservice.LoadTestsClientDeleteResponse]]
	newListByResourceGroupPager                  *tracker[azfake.PagerResponder[armloadtestservice.LoadTestsClientListByResourceGroupResponse]]
	newListBySubscriptionPager                   *tracker[azfake.PagerResponder[armloadtestservice.LoadTestsClientListBySubscriptionResponse]]
	newOutboundNetworkDependenciesEndpointsPager *tracker[azfake.PagerResponder[armloadtestservice.LoadTestsClientOutboundNetworkDependenciesEndpointsResponse]]
	beginUpdate                                  *tracker[azfake.PollerResponder[armloadtestservice.LoadTestsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for LoadTestsServerTransport.
func (l *LoadTestsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return l.dispatchToMethodFake(req, method)
}

func (l *LoadTestsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if loadTestsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = loadTestsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "LoadTestsClient.BeginCreateOrUpdate":
				res.resp, res.err = l.dispatchBeginCreateOrUpdate(req)
			case "LoadTestsClient.BeginDelete":
				res.resp, res.err = l.dispatchBeginDelete(req)
			case "LoadTestsClient.Get":
				res.resp, res.err = l.dispatchGet(req)
			case "LoadTestsClient.NewListByResourceGroupPager":
				res.resp, res.err = l.dispatchNewListByResourceGroupPager(req)
			case "LoadTestsClient.NewListBySubscriptionPager":
				res.resp, res.err = l.dispatchNewListBySubscriptionPager(req)
			case "LoadTestsClient.NewOutboundNetworkDependenciesEndpointsPager":
				res.resp, res.err = l.dispatchNewOutboundNetworkDependenciesEndpointsPager(req)
			case "LoadTestsClient.BeginUpdate":
				res.resp, res.err = l.dispatchBeginUpdate(req)
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

func (l *LoadTestsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := l.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/loadTests/(?P<loadTestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armloadtestservice.LoadTestResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		loadTestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("loadTestName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, loadTestNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		l.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		l.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		l.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (l *LoadTestsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := l.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/loadTests/(?P<loadTestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		loadTestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("loadTestName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginDelete(req.Context(), resourceGroupNameParam, loadTestNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		l.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		l.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		l.beginDelete.remove(req)
	}

	return resp, nil
}

func (l *LoadTestsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/loadTests/(?P<loadTestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	loadTestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("loadTestName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), resourceGroupNameParam, loadTestNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LoadTestResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LoadTestsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := l.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/loadTests`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		l.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armloadtestservice.LoadTestsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		l.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (l *LoadTestsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := l.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/loadTests`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := l.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		l.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armloadtestservice.LoadTestsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		l.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (l *LoadTestsServerTransport) dispatchNewOutboundNetworkDependenciesEndpointsPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewOutboundNetworkDependenciesEndpointsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewOutboundNetworkDependenciesEndpointsPager not implemented")}
	}
	newOutboundNetworkDependenciesEndpointsPager := l.newOutboundNetworkDependenciesEndpointsPager.get(req)
	if newOutboundNetworkDependenciesEndpointsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/loadTests/(?P<loadTestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/outboundNetworkDependenciesEndpoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		loadTestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("loadTestName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewOutboundNetworkDependenciesEndpointsPager(resourceGroupNameParam, loadTestNameParam, nil)
		newOutboundNetworkDependenciesEndpointsPager = &resp
		l.newOutboundNetworkDependenciesEndpointsPager.add(req, newOutboundNetworkDependenciesEndpointsPager)
		server.PagerResponderInjectNextLinks(newOutboundNetworkDependenciesEndpointsPager, req, func(page *armloadtestservice.LoadTestsClientOutboundNetworkDependenciesEndpointsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newOutboundNetworkDependenciesEndpointsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newOutboundNetworkDependenciesEndpointsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newOutboundNetworkDependenciesEndpointsPager) {
		l.newOutboundNetworkDependenciesEndpointsPager.remove(req)
	}
	return resp, nil
}

func (l *LoadTestsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := l.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/loadTests/(?P<loadTestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armloadtestservice.LoadTestResourceUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		loadTestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("loadTestName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginUpdate(req.Context(), resourceGroupNameParam, loadTestNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		l.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		l.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		l.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to LoadTestsServerTransport
var loadTestsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
