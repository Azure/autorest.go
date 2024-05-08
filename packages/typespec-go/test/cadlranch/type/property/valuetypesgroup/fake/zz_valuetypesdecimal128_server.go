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
	"valuetypesgroup"
)

// ValueTypesDecimal128Server is a fake server for instances of the valuetypesgroup.ValueTypesDecimal128Client type.
type ValueTypesDecimal128Server struct {
	// Get is the fake for method ValueTypesDecimal128Client.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *valuetypesgroup.ValueTypesDecimal128ClientGetOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesDecimal128ClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ValueTypesDecimal128Client.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body valuetypesgroup.Decimal128Property, options *valuetypesgroup.ValueTypesDecimal128ClientPutOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesDecimal128ClientPutResponse], errResp azfake.ErrorResponder)
}

// NewValueTypesDecimal128ServerTransport creates a new instance of ValueTypesDecimal128ServerTransport with the provided implementation.
// The returned ValueTypesDecimal128ServerTransport instance is connected to an instance of valuetypesgroup.ValueTypesDecimal128Client via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewValueTypesDecimal128ServerTransport(srv *ValueTypesDecimal128Server) *ValueTypesDecimal128ServerTransport {
	return &ValueTypesDecimal128ServerTransport{srv: srv}
}

// ValueTypesDecimal128ServerTransport connects instances of valuetypesgroup.ValueTypesDecimal128Client to instances of ValueTypesDecimal128Server.
// Don't use this type directly, use NewValueTypesDecimal128ServerTransport instead.
type ValueTypesDecimal128ServerTransport struct {
	srv *ValueTypesDecimal128Server
}

// Do implements the policy.Transporter interface for ValueTypesDecimal128ServerTransport.
func (v *ValueTypesDecimal128ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *ValueTypesDecimal128ServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ValueTypesDecimal128Client.Get":
		resp, err = v.dispatchGet(req)
	case "ValueTypesDecimal128Client.Put":
		resp, err = v.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (v *ValueTypesDecimal128ServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := v.srv.Get(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Decimal128Property, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *ValueTypesDecimal128ServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if v.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[valuetypesgroup.Decimal128Property](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Put(req.Context(), body, nil)
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