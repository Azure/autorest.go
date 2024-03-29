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
	"nullablegroup"
)

// CollectionsModelServer is a fake server for instances of the nullablegroup.CollectionsModelClient type.
type CollectionsModelServer struct {
	// GetNonNull is the fake for method CollectionsModelClient.GetNonNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNonNull func(ctx context.Context, options *nullablegroup.CollectionsModelClientGetNonNullOptions) (resp azfake.Responder[nullablegroup.CollectionsModelClientGetNonNullResponse], errResp azfake.ErrorResponder)

	// GetNull is the fake for method CollectionsModelClient.GetNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNull func(ctx context.Context, options *nullablegroup.CollectionsModelClientGetNullOptions) (resp azfake.Responder[nullablegroup.CollectionsModelClientGetNullResponse], errResp azfake.ErrorResponder)

	// PatchNonNull is the fake for method CollectionsModelClient.PatchNonNull
	// HTTP status codes to indicate success: http.StatusNoContent
	PatchNonNull func(ctx context.Context, body nullablegroup.CollectionsModelProperty, options *nullablegroup.CollectionsModelClientPatchNonNullOptions) (resp azfake.Responder[nullablegroup.CollectionsModelClientPatchNonNullResponse], errResp azfake.ErrorResponder)

	// PatchNull is the fake for method CollectionsModelClient.PatchNull
	// HTTP status codes to indicate success: http.StatusNoContent
	PatchNull func(ctx context.Context, body nullablegroup.CollectionsModelProperty, options *nullablegroup.CollectionsModelClientPatchNullOptions) (resp azfake.Responder[nullablegroup.CollectionsModelClientPatchNullResponse], errResp azfake.ErrorResponder)
}

// NewCollectionsModelServerTransport creates a new instance of CollectionsModelServerTransport with the provided implementation.
// The returned CollectionsModelServerTransport instance is connected to an instance of nullablegroup.CollectionsModelClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCollectionsModelServerTransport(srv *CollectionsModelServer) *CollectionsModelServerTransport {
	return &CollectionsModelServerTransport{srv: srv}
}

// CollectionsModelServerTransport connects instances of nullablegroup.CollectionsModelClient to instances of CollectionsModelServer.
// Don't use this type directly, use NewCollectionsModelServerTransport instead.
type CollectionsModelServerTransport struct {
	srv *CollectionsModelServer
}

// Do implements the policy.Transporter interface for CollectionsModelServerTransport.
func (c *CollectionsModelServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *CollectionsModelServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "CollectionsModelClient.GetNonNull":
		resp, err = c.dispatchGetNonNull(req)
	case "CollectionsModelClient.GetNull":
		resp, err = c.dispatchGetNull(req)
	case "CollectionsModelClient.PatchNonNull":
		resp, err = c.dispatchPatchNonNull(req)
	case "CollectionsModelClient.PatchNull":
		resp, err = c.dispatchPatchNull(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (c *CollectionsModelServerTransport) dispatchGetNonNull(req *http.Request) (*http.Response, error) {
	if c.srv.GetNonNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNonNull not implemented")}
	}
	respr, errRespr := c.srv.GetNonNull(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CollectionsModelProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CollectionsModelServerTransport) dispatchGetNull(req *http.Request) (*http.Response, error) {
	if c.srv.GetNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNull not implemented")}
	}
	respr, errRespr := c.srv.GetNull(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CollectionsModelProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CollectionsModelServerTransport) dispatchPatchNonNull(req *http.Request) (*http.Response, error) {
	if c.srv.PatchNonNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method PatchNonNull not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[nullablegroup.CollectionsModelProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.PatchNonNull(req.Context(), body, nil)
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

func (c *CollectionsModelServerTransport) dispatchPatchNull(req *http.Request) (*http.Response, error) {
	if c.srv.PatchNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method PatchNull not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[nullablegroup.CollectionsModelProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.PatchNull(req.Context(), body, nil)
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
