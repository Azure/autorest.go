//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armcompute"
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

// CloudServiceRoleInstancesServer is a fake server for instances of the armcompute.CloudServiceRoleInstancesClient type.
type CloudServiceRoleInstancesServer struct {
	// BeginDelete is the fake for method CloudServiceRoleInstancesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServiceRoleInstancesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.CloudServiceRoleInstancesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CloudServiceRoleInstancesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServiceRoleInstancesClientGetOptions) (resp azfake.Responder[armcompute.CloudServiceRoleInstancesClientGetResponse], errResp azfake.ErrorResponder)

	// GetInstanceView is the fake for method CloudServiceRoleInstancesClient.GetInstanceView
	// HTTP status codes to indicate success: http.StatusOK
	GetInstanceView func(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServiceRoleInstancesClientGetInstanceViewOptions) (resp azfake.Responder[armcompute.CloudServiceRoleInstancesClientGetInstanceViewResponse], errResp azfake.ErrorResponder)

	// GetRemoteDesktopFile is the fake for method CloudServiceRoleInstancesClient.GetRemoteDesktopFile
	// HTTP status codes to indicate success: http.StatusOK
	GetRemoteDesktopFile func(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServiceRoleInstancesClientGetRemoteDesktopFileOptions) (resp azfake.Responder[armcompute.CloudServiceRoleInstancesClientGetRemoteDesktopFileResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method CloudServiceRoleInstancesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, cloudServiceName string, options *armcompute.CloudServiceRoleInstancesClientListOptions) (resp azfake.PagerResponder[armcompute.CloudServiceRoleInstancesClientListResponse])

	// BeginRebuild is the fake for method CloudServiceRoleInstancesClient.BeginRebuild
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRebuild func(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServiceRoleInstancesClientBeginRebuildOptions) (resp azfake.PollerResponder[armcompute.CloudServiceRoleInstancesClientRebuildResponse], errResp azfake.ErrorResponder)

	// BeginReimage is the fake for method CloudServiceRoleInstancesClient.BeginReimage
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginReimage func(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServiceRoleInstancesClientBeginReimageOptions) (resp azfake.PollerResponder[armcompute.CloudServiceRoleInstancesClientReimageResponse], errResp azfake.ErrorResponder)

	// BeginRestart is the fake for method CloudServiceRoleInstancesClient.BeginRestart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRestart func(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServiceRoleInstancesClientBeginRestartOptions) (resp azfake.PollerResponder[armcompute.CloudServiceRoleInstancesClientRestartResponse], errResp azfake.ErrorResponder)
}

// NewCloudServiceRoleInstancesServerTransport creates a new instance of CloudServiceRoleInstancesServerTransport with the provided implementation.
// The returned CloudServiceRoleInstancesServerTransport instance is connected to an instance of armcompute.CloudServiceRoleInstancesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCloudServiceRoleInstancesServerTransport(srv *CloudServiceRoleInstancesServer) *CloudServiceRoleInstancesServerTransport {
	return &CloudServiceRoleInstancesServerTransport{srv: srv}
}

// CloudServiceRoleInstancesServerTransport connects instances of armcompute.CloudServiceRoleInstancesClient to instances of CloudServiceRoleInstancesServer.
// Don't use this type directly, use NewCloudServiceRoleInstancesServerTransport instead.
type CloudServiceRoleInstancesServerTransport struct {
	srv          *CloudServiceRoleInstancesServer
	beginDelete  *azfake.PollerResponder[armcompute.CloudServiceRoleInstancesClientDeleteResponse]
	newListPager *azfake.PagerResponder[armcompute.CloudServiceRoleInstancesClientListResponse]
	beginRebuild *azfake.PollerResponder[armcompute.CloudServiceRoleInstancesClientRebuildResponse]
	beginReimage *azfake.PollerResponder[armcompute.CloudServiceRoleInstancesClientReimageResponse]
	beginRestart *azfake.PollerResponder[armcompute.CloudServiceRoleInstancesClientRestartResponse]
}

// Do implements the policy.Transporter interface for CloudServiceRoleInstancesServerTransport.
func (c *CloudServiceRoleInstancesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CloudServiceRoleInstancesClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "CloudServiceRoleInstancesClient.Get":
		resp, err = c.dispatchGet(req)
	case "CloudServiceRoleInstancesClient.GetInstanceView":
		resp, err = c.dispatchGetInstanceView(req)
	case "CloudServiceRoleInstancesClient.GetRemoteDesktopFile":
		resp, err = c.dispatchGetRemoteDesktopFile(req)
	case "CloudServiceRoleInstancesClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	case "CloudServiceRoleInstancesClient.BeginRebuild":
		resp, err = c.dispatchBeginRebuild(req)
	case "CloudServiceRoleInstancesClient.BeginReimage":
		resp, err = c.dispatchBeginReimage(req)
	case "CloudServiceRoleInstancesClient.BeginRestart":
		resp, err = c.dispatchBeginRestart(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloudServiceRoleInstancesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if c.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roleInstances/(?P<roleInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		roleInstanceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("roleInstanceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), roleInstanceNameUnescaped, resourceGroupNameUnescaped, cloudServiceNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(c.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginDelete) {
		c.beginDelete = nil
	}

	return resp, nil
}

func (c *CloudServiceRoleInstancesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roleInstances/(?P<roleInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	roleInstanceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("roleInstanceName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(armcompute.InstanceViewTypes(expandUnescaped))
	var options *armcompute.CloudServiceRoleInstancesClientGetOptions
	if expandParam != nil {
		options = &armcompute.CloudServiceRoleInstancesClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := c.srv.Get(req.Context(), roleInstanceNameUnescaped, resourceGroupNameUnescaped, cloudServiceNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CloudServiceRoleInstancesServerTransport) dispatchGetInstanceView(req *http.Request) (*http.Response, error) {
	if c.srv.GetInstanceView == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetInstanceView not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roleInstances/(?P<roleInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instanceView`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	roleInstanceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("roleInstanceName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetInstanceView(req.Context(), roleInstanceNameUnescaped, resourceGroupNameUnescaped, cloudServiceNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleInstanceView, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CloudServiceRoleInstancesServerTransport) dispatchGetRemoteDesktopFile(req *http.Request) (*http.Response, error) {
	if c.srv.GetRemoteDesktopFile == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetRemoteDesktopFile not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roleInstances/(?P<roleInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/remoteDesktopFile`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	roleInstanceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("roleInstanceName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetRemoteDesktopFile(req.Context(), roleInstanceNameUnescaped, resourceGroupNameUnescaped, cloudServiceNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CloudServiceRoleInstancesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if c.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roleInstances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(armcompute.InstanceViewTypes(expandUnescaped))
		var options *armcompute.CloudServiceRoleInstancesClientListOptions
		if expandParam != nil {
			options = &armcompute.CloudServiceRoleInstancesClientListOptions{
				Expand: expandParam,
			}
		}
		resp := c.srv.NewListPager(resourceGroupNameUnescaped, cloudServiceNameUnescaped, options)
		c.newListPager = &resp
		server.PagerResponderInjectNextLinks(c.newListPager, req, func(page *armcompute.CloudServiceRoleInstancesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(c.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(c.newListPager) {
		c.newListPager = nil
	}
	return resp, nil
}

func (c *CloudServiceRoleInstancesServerTransport) dispatchBeginRebuild(req *http.Request) (*http.Response, error) {
	if c.srv.BeginRebuild == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRebuild not implemented")}
	}
	if c.beginRebuild == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roleInstances/(?P<roleInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rebuild`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		roleInstanceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("roleInstanceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginRebuild(req.Context(), roleInstanceNameUnescaped, resourceGroupNameUnescaped, cloudServiceNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginRebuild = &respr
	}

	resp, err := server.PollerResponderNext(c.beginRebuild, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginRebuild) {
		c.beginRebuild = nil
	}

	return resp, nil
}

func (c *CloudServiceRoleInstancesServerTransport) dispatchBeginReimage(req *http.Request) (*http.Response, error) {
	if c.srv.BeginReimage == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginReimage not implemented")}
	}
	if c.beginReimage == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roleInstances/(?P<roleInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reimage`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		roleInstanceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("roleInstanceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginReimage(req.Context(), roleInstanceNameUnescaped, resourceGroupNameUnescaped, cloudServiceNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginReimage = &respr
	}

	resp, err := server.PollerResponderNext(c.beginReimage, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginReimage) {
		c.beginReimage = nil
	}

	return resp, nil
}

func (c *CloudServiceRoleInstancesServerTransport) dispatchBeginRestart(req *http.Request) (*http.Response, error) {
	if c.srv.BeginRestart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRestart not implemented")}
	}
	if c.beginRestart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roleInstances/(?P<roleInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restart`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		roleInstanceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("roleInstanceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginRestart(req.Context(), roleInstanceNameUnescaped, resourceGroupNameUnescaped, cloudServiceNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginRestart = &respr
	}

	resp, err := server.PollerResponderNext(c.beginRestart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginRestart) {
		c.beginRestart = nil
	}

	return resp, nil
}